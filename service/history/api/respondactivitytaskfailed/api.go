// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package respondactivitytaskfailed

import (
	"context"
	"time"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/workflowservice/v1"

	"go.temporal.io/server/api/historyservice/v1"
	"go.temporal.io/server/common"
	"go.temporal.io/server/common/definition"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/service/history/api"
	"go.temporal.io/server/service/history/consts"
	"go.temporal.io/server/service/history/shard"
)

func Invoke(
	ctx context.Context,
	req *historyservice.RespondActivityTaskFailedRequest,
	shard shard.Context,
	workflowConsistencyChecker api.WorkflowConsistencyChecker,
) (resp *historyservice.RespondActivityTaskFailedResponse, retError error) {
	namespaceEntry, err := api.GetActiveNamespace(shard, namespace.ID(req.GetNamespaceId()))
	if err != nil {
		return nil, err
	}
	namespace := namespaceEntry.Name()

	request := req.FailedRequest
	tokenSerializer := common.NewProtoTaskTokenSerializer()
	token, err0 := tokenSerializer.Deserialize(request.TaskToken)
	if err0 != nil {
		return nil, consts.ErrDeserializingToken
	}
	if err := api.SetActivityTaskRunID(ctx, token, workflowConsistencyChecker); err != nil {
		return nil, err
	}

	var activityStartedTime time.Time
	var taskQueue string
	var workflowTypeName string
	err = api.GetAndUpdateWorkflowWithNew(
		ctx,
		token.Clock,
		api.BypassMutableStateConsistencyPredicate,
		definition.NewWorkflowKey(
			token.NamespaceId,
			token.WorkflowId,
			token.RunId,
		),
		func(workflowContext api.WorkflowContext) (*api.UpdateWorkflowAction, error) {
			mutableState := workflowContext.GetMutableState()
			workflowTypeName = mutableState.GetWorkflowType().GetName()
			if !mutableState.IsWorkflowExecutionRunning() {
				return nil, consts.ErrWorkflowCompleted
			}

			scheduledEventID := token.GetScheduledEventId()
			if scheduledEventID == common.EmptyEventID { // client call CompleteActivityById, so get scheduledEventID by activityID
				scheduledEventID, err0 = api.GetActivityScheduledEventID(token.GetActivityId(), mutableState)
				if err0 != nil {
					return nil, err0
				}
			}
			ai, isRunning := mutableState.GetActivityInfo(scheduledEventID)

			// First check to see if cache needs to be refreshed as we could potentially have stale workflow execution in
			// some extreme cassandra failure cases.
			if !isRunning && scheduledEventID >= mutableState.GetNextEventID() {
				shard.GetMetricsHandler().Counter(metrics.StaleMutableStateCounter.GetMetricName()).Record(
					1,
					metrics.OperationTag(metrics.HistoryRespondActivityTaskFailedScope))
				return nil, consts.ErrStaleState
			}

			if !isRunning || ai.StartedEventId == common.EmptyEventID ||
				(token.GetScheduledEventId() != common.EmptyEventID && token.Attempt != ai.Attempt) {
				return nil, consts.ErrActivityTaskNotFound
			}

			if request.GetLastHeartbeatDetails() != nil {
				// Save heartbeat details as progress
				mutableState.UpdateActivityProgress(ai, &workflowservice.RecordActivityTaskHeartbeatRequest{
					TaskToken: request.GetTaskToken(),
					Details:   request.GetLastHeartbeatDetails(),
					Identity:  request.GetIdentity(),
					Namespace: request.GetNamespace(),
				})
			}

			postActions := &api.UpdateWorkflowAction{}
			failure := request.GetFailure()
			retryState, err := mutableState.RetryActivity(ai, failure)
			if err != nil {
				return nil, err
			}
			if retryState != enumspb.RETRY_STATE_IN_PROGRESS {
				// no more retry, and we want to record the failure event
				if _, err := mutableState.AddActivityTaskFailedEvent(scheduledEventID, ai.StartedEventId, failure, retryState, request.GetIdentity()); err != nil {
					// Unable to add ActivityTaskFailed event to history
					return nil, err
				}
				postActions.CreateWorkflowTask = true
			}

			activityStartedTime = *ai.StartedTime
			taskQueue = ai.TaskQueue
			return postActions, nil
		},
		nil,
		shard,
		workflowConsistencyChecker,
	)
	if err == nil && !activityStartedTime.IsZero() {
		shard.GetMetricsHandler().Timer(metrics.ActivityE2ELatency.GetMetricName()).Record(
			time.Since(activityStartedTime),
			metrics.OperationTag(metrics.HistoryRespondActivityTaskFailedScope),
			metrics.NamespaceTag(namespace.String()),
			metrics.WorkflowTypeTag(workflowTypeName),
			metrics.ActivityTypeTag(token.ActivityType),
			metrics.TaskQueueTag(taskQueue),
		)
	}
	return &historyservice.RespondActivityTaskFailedResponse{}, err
}
