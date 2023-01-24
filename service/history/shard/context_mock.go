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

// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package shard is a generated GoMock package.
package shard

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v1 "go.temporal.io/api/common/v1"
	v10 "go.temporal.io/server/api/adminservice/v1"
	v11 "go.temporal.io/server/api/clock/v1"
	v12 "go.temporal.io/server/api/historyservice/v1"
	v13 "go.temporal.io/server/api/persistence/v1"
	archiver "go.temporal.io/server/common/archiver"
	clock "go.temporal.io/server/common/clock"
	cluster "go.temporal.io/server/common/cluster"
	definition "go.temporal.io/server/common/definition"
	log "go.temporal.io/server/common/log"
	metrics "go.temporal.io/server/common/metrics"
	namespace "go.temporal.io/server/common/namespace"
	persistence "go.temporal.io/server/common/persistence"
	serialization "go.temporal.io/server/common/persistence/serialization"
	searchattribute "go.temporal.io/server/common/searchattribute"
	configs "go.temporal.io/server/service/history/configs"
	events "go.temporal.io/server/service/history/events"
	tasks "go.temporal.io/server/service/history/tasks"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// AddTasks mocks base method.
func (m *MockContext) AddTasks(ctx context.Context, request *persistence.AddHistoryTasksRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTasks", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTasks indicates an expected call of AddTasks.
func (mr *MockContextMockRecorder) AddTasks(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTasks", reflect.TypeOf((*MockContext)(nil).AddTasks), ctx, request)
}

// AppendHistoryEvents mocks base method.
func (m *MockContext) AppendHistoryEvents(ctx context.Context, request *persistence.AppendHistoryNodesRequest, namespaceID namespace.ID, execution v1.WorkflowExecution) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendHistoryEvents", ctx, request, namespaceID, execution)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendHistoryEvents indicates an expected call of AppendHistoryEvents.
func (mr *MockContextMockRecorder) AppendHistoryEvents(ctx, request, namespaceID, execution interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendHistoryEvents", reflect.TypeOf((*MockContext)(nil).AppendHistoryEvents), ctx, request, namespaceID, execution)
}

// AssertOwnership mocks base method.
func (m *MockContext) AssertOwnership(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssertOwnership", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssertOwnership indicates an expected call of AssertOwnership.
func (mr *MockContextMockRecorder) AssertOwnership(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssertOwnership", reflect.TypeOf((*MockContext)(nil).AssertOwnership), ctx)
}

// ConflictResolveWorkflowExecution mocks base method.
func (m *MockContext) ConflictResolveWorkflowExecution(ctx context.Context, request *persistence.ConflictResolveWorkflowExecutionRequest) (*persistence.ConflictResolveWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConflictResolveWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*persistence.ConflictResolveWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConflictResolveWorkflowExecution indicates an expected call of ConflictResolveWorkflowExecution.
func (mr *MockContextMockRecorder) ConflictResolveWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConflictResolveWorkflowExecution", reflect.TypeOf((*MockContext)(nil).ConflictResolveWorkflowExecution), ctx, request)
}

// CreateWorkflowExecution mocks base method.
func (m *MockContext) CreateWorkflowExecution(ctx context.Context, request *persistence.CreateWorkflowExecutionRequest) (*persistence.CreateWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*persistence.CreateWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkflowExecution indicates an expected call of CreateWorkflowExecution.
func (mr *MockContextMockRecorder) CreateWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkflowExecution", reflect.TypeOf((*MockContext)(nil).CreateWorkflowExecution), ctx, request)
}

// CurrentVectorClock mocks base method.
func (m *MockContext) CurrentVectorClock() *v11.VectorClock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentVectorClock")
	ret0, _ := ret[0].(*v11.VectorClock)
	return ret0
}

// CurrentVectorClock indicates an expected call of CurrentVectorClock.
func (mr *MockContextMockRecorder) CurrentVectorClock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentVectorClock", reflect.TypeOf((*MockContext)(nil).CurrentVectorClock))
}

// DeleteFailoverLevel mocks base method.
func (m *MockContext) DeleteFailoverLevel(category tasks.Category, failoverID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFailoverLevel", category, failoverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFailoverLevel indicates an expected call of DeleteFailoverLevel.
func (mr *MockContextMockRecorder) DeleteFailoverLevel(category, failoverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFailoverLevel", reflect.TypeOf((*MockContext)(nil).DeleteFailoverLevel), category, failoverID)
}

// DeleteWorkflowExecution mocks base method.
func (m *MockContext) DeleteWorkflowExecution(ctx context.Context, workflowKey definition.WorkflowKey, branchToken []byte, startTime, closeTime *time.Time, closeExecutionVisibilityTaskID int64, stage *tasks.DeleteWorkflowExecutionStage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflowExecution", ctx, workflowKey, branchToken, startTime, closeTime, closeExecutionVisibilityTaskID, stage)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflowExecution indicates an expected call of DeleteWorkflowExecution.
func (mr *MockContextMockRecorder) DeleteWorkflowExecution(ctx, workflowKey, branchToken, startTime, closeTime, closeExecutionVisibilityTaskID, stage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflowExecution", reflect.TypeOf((*MockContext)(nil).DeleteWorkflowExecution), ctx, workflowKey, branchToken, startTime, closeTime, closeExecutionVisibilityTaskID, stage)
}

// GenerateTaskID mocks base method.
func (m *MockContext) GenerateTaskID() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTaskID")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTaskID indicates an expected call of GenerateTaskID.
func (mr *MockContextMockRecorder) GenerateTaskID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTaskID", reflect.TypeOf((*MockContext)(nil).GenerateTaskID))
}

// GenerateTaskIDs mocks base method.
func (m *MockContext) GenerateTaskIDs(number int) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTaskIDs", number)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTaskIDs indicates an expected call of GenerateTaskIDs.
func (mr *MockContextMockRecorder) GenerateTaskIDs(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTaskIDs", reflect.TypeOf((*MockContext)(nil).GenerateTaskIDs), number)
}

// GetAllFailoverLevels mocks base method.
func (m *MockContext) GetAllFailoverLevels(category tasks.Category) map[string]persistence.FailoverLevel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFailoverLevels", category)
	ret0, _ := ret[0].(map[string]persistence.FailoverLevel)
	return ret0
}

// GetAllFailoverLevels indicates an expected call of GetAllFailoverLevels.
func (mr *MockContextMockRecorder) GetAllFailoverLevels(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFailoverLevels", reflect.TypeOf((*MockContext)(nil).GetAllFailoverLevels), category)
}

// GetArchivalMetadata mocks base method.
func (m *MockContext) GetArchivalMetadata() archiver.ArchivalMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchivalMetadata")
	ret0, _ := ret[0].(archiver.ArchivalMetadata)
	return ret0
}

// GetArchivalMetadata indicates an expected call of GetArchivalMetadata.
func (mr *MockContextMockRecorder) GetArchivalMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchivalMetadata", reflect.TypeOf((*MockContext)(nil).GetArchivalMetadata))
}

// GetClusterMetadata mocks base method.
func (m *MockContext) GetClusterMetadata() cluster.Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterMetadata")
	ret0, _ := ret[0].(cluster.Metadata)
	return ret0
}

// GetClusterMetadata indicates an expected call of GetClusterMetadata.
func (mr *MockContextMockRecorder) GetClusterMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterMetadata", reflect.TypeOf((*MockContext)(nil).GetClusterMetadata))
}

// GetConfig mocks base method.
func (m *MockContext) GetConfig() *configs.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*configs.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockContextMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockContext)(nil).GetConfig))
}

// GetCurrentExecution mocks base method.
func (m *MockContext) GetCurrentExecution(ctx context.Context, request *persistence.GetCurrentExecutionRequest) (*persistence.GetCurrentExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentExecution", ctx, request)
	ret0, _ := ret[0].(*persistence.GetCurrentExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentExecution indicates an expected call of GetCurrentExecution.
func (mr *MockContextMockRecorder) GetCurrentExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentExecution", reflect.TypeOf((*MockContext)(nil).GetCurrentExecution), ctx, request)
}

// GetCurrentTime mocks base method.
func (m *MockContext) GetCurrentTime(cluster string) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentTime", cluster)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCurrentTime indicates an expected call of GetCurrentTime.
func (mr *MockContextMockRecorder) GetCurrentTime(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentTime", reflect.TypeOf((*MockContext)(nil).GetCurrentTime), cluster)
}

// GetEngine mocks base method.
func (m *MockContext) GetEngine(ctx context.Context) (Engine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEngine", ctx)
	ret0, _ := ret[0].(Engine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEngine indicates an expected call of GetEngine.
func (mr *MockContextMockRecorder) GetEngine(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEngine", reflect.TypeOf((*MockContext)(nil).GetEngine), ctx)
}

// GetEventsCache mocks base method.
func (m *MockContext) GetEventsCache() events.Cache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsCache")
	ret0, _ := ret[0].(events.Cache)
	return ret0
}

// GetEventsCache indicates an expected call of GetEventsCache.
func (mr *MockContextMockRecorder) GetEventsCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsCache", reflect.TypeOf((*MockContext)(nil).GetEventsCache))
}

// GetExecutionManager mocks base method.
func (m *MockContext) GetExecutionManager() persistence.ExecutionManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecutionManager")
	ret0, _ := ret[0].(persistence.ExecutionManager)
	return ret0
}

// GetExecutionManager indicates an expected call of GetExecutionManager.
func (mr *MockContextMockRecorder) GetExecutionManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecutionManager", reflect.TypeOf((*MockContext)(nil).GetExecutionManager))
}

// GetHistoryClient mocks base method.
func (m *MockContext) GetHistoryClient() v12.HistoryServiceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryClient")
	ret0, _ := ret[0].(v12.HistoryServiceClient)
	return ret0
}

// GetHistoryClient indicates an expected call of GetHistoryClient.
func (mr *MockContextMockRecorder) GetHistoryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryClient", reflect.TypeOf((*MockContext)(nil).GetHistoryClient))
}

// GetImmediateQueueExclusiveHighReadWatermark mocks base method.
func (m *MockContext) GetImmediateQueueExclusiveHighReadWatermark() tasks.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImmediateQueueExclusiveHighReadWatermark")
	ret0, _ := ret[0].(tasks.Key)
	return ret0
}

// GetImmediateQueueExclusiveHighReadWatermark indicates an expected call of GetImmediateQueueExclusiveHighReadWatermark.
func (mr *MockContextMockRecorder) GetImmediateQueueExclusiveHighReadWatermark() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImmediateQueueExclusiveHighReadWatermark", reflect.TypeOf((*MockContext)(nil).GetImmediateQueueExclusiveHighReadWatermark))
}

// GetLastUpdatedTime mocks base method.
func (m *MockContext) GetLastUpdatedTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastUpdatedTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLastUpdatedTime indicates an expected call of GetLastUpdatedTime.
func (mr *MockContextMockRecorder) GetLastUpdatedTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastUpdatedTime", reflect.TypeOf((*MockContext)(nil).GetLastUpdatedTime))
}

// GetLogger mocks base method.
func (m *MockContext) GetLogger() log.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(log.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockContextMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockContext)(nil).GetLogger))
}

// GetMaxTaskIDForCurrentRangeID mocks base method.
func (m *MockContext) GetMaxTaskIDForCurrentRangeID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxTaskIDForCurrentRangeID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetMaxTaskIDForCurrentRangeID indicates an expected call of GetMaxTaskIDForCurrentRangeID.
func (mr *MockContextMockRecorder) GetMaxTaskIDForCurrentRangeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxTaskIDForCurrentRangeID", reflect.TypeOf((*MockContext)(nil).GetMaxTaskIDForCurrentRangeID))
}

// GetMetricsHandler mocks base method.
func (m *MockContext) GetMetricsHandler() metrics.MetricsHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricsHandler")
	ret0, _ := ret[0].(metrics.MetricsHandler)
	return ret0
}

// GetMetricsHandler indicates an expected call of GetMetricsHandler.
func (mr *MockContextMockRecorder) GetMetricsHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsHandler", reflect.TypeOf((*MockContext)(nil).GetMetricsHandler))
}

// GetNamespaceNotificationVersion mocks base method.
func (m *MockContext) GetNamespaceNotificationVersion() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceNotificationVersion")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetNamespaceNotificationVersion indicates an expected call of GetNamespaceNotificationVersion.
func (mr *MockContextMockRecorder) GetNamespaceNotificationVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceNotificationVersion", reflect.TypeOf((*MockContext)(nil).GetNamespaceNotificationVersion))
}

// GetNamespaceRegistry mocks base method.
func (m *MockContext) GetNamespaceRegistry() namespace.Registry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceRegistry")
	ret0, _ := ret[0].(namespace.Registry)
	return ret0
}

// GetNamespaceRegistry indicates an expected call of GetNamespaceRegistry.
func (mr *MockContextMockRecorder) GetNamespaceRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceRegistry", reflect.TypeOf((*MockContext)(nil).GetNamespaceRegistry))
}

// GetPayloadSerializer mocks base method.
func (m *MockContext) GetPayloadSerializer() serialization.Serializer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayloadSerializer")
	ret0, _ := ret[0].(serialization.Serializer)
	return ret0
}

// GetPayloadSerializer indicates an expected call of GetPayloadSerializer.
func (mr *MockContextMockRecorder) GetPayloadSerializer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayloadSerializer", reflect.TypeOf((*MockContext)(nil).GetPayloadSerializer))
}

// GetQueueAckLevel mocks base method.
func (m *MockContext) GetQueueAckLevel(category tasks.Category) tasks.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueueAckLevel", category)
	ret0, _ := ret[0].(tasks.Key)
	return ret0
}

// GetQueueAckLevel indicates an expected call of GetQueueAckLevel.
func (mr *MockContextMockRecorder) GetQueueAckLevel(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueAckLevel", reflect.TypeOf((*MockContext)(nil).GetQueueAckLevel), category)
}

// GetQueueClusterAckLevel mocks base method.
func (m *MockContext) GetQueueClusterAckLevel(category tasks.Category, cluster string) tasks.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueueClusterAckLevel", category, cluster)
	ret0, _ := ret[0].(tasks.Key)
	return ret0
}

// GetQueueClusterAckLevel indicates an expected call of GetQueueClusterAckLevel.
func (mr *MockContextMockRecorder) GetQueueClusterAckLevel(category, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueClusterAckLevel", reflect.TypeOf((*MockContext)(nil).GetQueueClusterAckLevel), category, cluster)
}

// GetQueueState mocks base method.
func (m *MockContext) GetQueueState(category tasks.Category) (*v13.QueueState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueueState", category)
	ret0, _ := ret[0].(*v13.QueueState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetQueueState indicates an expected call of GetQueueState.
func (mr *MockContextMockRecorder) GetQueueState(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueState", reflect.TypeOf((*MockContext)(nil).GetQueueState), category)
}

// GetRemoteAdminClient mocks base method.
func (m *MockContext) GetRemoteAdminClient(cluster string) (v10.AdminServiceClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteAdminClient", cluster)
	ret0, _ := ret[0].(v10.AdminServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteAdminClient indicates an expected call of GetRemoteAdminClient.
func (mr *MockContextMockRecorder) GetRemoteAdminClient(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteAdminClient", reflect.TypeOf((*MockContext)(nil).GetRemoteAdminClient), cluster)
}

// GetReplicationStatus mocks base method.
func (m *MockContext) GetReplicationStatus(cluster []string) (map[string]*v12.ShardReplicationStatusPerCluster, map[string]*v12.HandoverNamespaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicationStatus", cluster)
	ret0, _ := ret[0].(map[string]*v12.ShardReplicationStatusPerCluster)
	ret1, _ := ret[1].(map[string]*v12.HandoverNamespaceInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetReplicationStatus indicates an expected call of GetReplicationStatus.
func (mr *MockContextMockRecorder) GetReplicationStatus(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicationStatus", reflect.TypeOf((*MockContext)(nil).GetReplicationStatus), cluster)
}

// GetReplicatorDLQAckLevel mocks base method.
func (m *MockContext) GetReplicatorDLQAckLevel(sourceCluster string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicatorDLQAckLevel", sourceCluster)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetReplicatorDLQAckLevel indicates an expected call of GetReplicatorDLQAckLevel.
func (mr *MockContextMockRecorder) GetReplicatorDLQAckLevel(sourceCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicatorDLQAckLevel", reflect.TypeOf((*MockContext)(nil).GetReplicatorDLQAckLevel), sourceCluster)
}

// GetSearchAttributesMapper mocks base method.
func (m *MockContext) GetSearchAttributesMapper() searchattribute.Mapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSearchAttributesMapper")
	ret0, _ := ret[0].(searchattribute.Mapper)
	return ret0
}

// GetSearchAttributesMapper indicates an expected call of GetSearchAttributesMapper.
func (mr *MockContextMockRecorder) GetSearchAttributesMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSearchAttributesMapper", reflect.TypeOf((*MockContext)(nil).GetSearchAttributesMapper))
}

// GetSearchAttributesProvider mocks base method.
func (m *MockContext) GetSearchAttributesProvider() searchattribute.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSearchAttributesProvider")
	ret0, _ := ret[0].(searchattribute.Provider)
	return ret0
}

// GetSearchAttributesProvider indicates an expected call of GetSearchAttributesProvider.
func (mr *MockContextMockRecorder) GetSearchAttributesProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSearchAttributesProvider", reflect.TypeOf((*MockContext)(nil).GetSearchAttributesProvider))
}

// GetShardID mocks base method.
func (m *MockContext) GetShardID() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardID")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetShardID indicates an expected call of GetShardID.
func (mr *MockContextMockRecorder) GetShardID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardID", reflect.TypeOf((*MockContext)(nil).GetShardID))
}

// GetThrottledLogger mocks base method.
func (m *MockContext) GetThrottledLogger() log.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThrottledLogger")
	ret0, _ := ret[0].(log.Logger)
	return ret0
}

// GetThrottledLogger indicates an expected call of GetThrottledLogger.
func (mr *MockContextMockRecorder) GetThrottledLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThrottledLogger", reflect.TypeOf((*MockContext)(nil).GetThrottledLogger))
}

// GetTimeSource mocks base method.
func (m *MockContext) GetTimeSource() clock.TimeSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeSource")
	ret0, _ := ret[0].(clock.TimeSource)
	return ret0
}

// GetTimeSource indicates an expected call of GetTimeSource.
func (mr *MockContextMockRecorder) GetTimeSource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeSource", reflect.TypeOf((*MockContext)(nil).GetTimeSource))
}

// GetWorkflowExecution mocks base method.
func (m *MockContext) GetWorkflowExecution(ctx context.Context, request *persistence.GetWorkflowExecutionRequest) (*persistence.GetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*persistence.GetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowExecution indicates an expected call of GetWorkflowExecution.
func (mr *MockContextMockRecorder) GetWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowExecution", reflect.TypeOf((*MockContext)(nil).GetWorkflowExecution), ctx, request)
}

// NewVectorClock mocks base method.
func (m *MockContext) NewVectorClock() (*v11.VectorClock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewVectorClock")
	ret0, _ := ret[0].(*v11.VectorClock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewVectorClock indicates an expected call of NewVectorClock.
func (mr *MockContextMockRecorder) NewVectorClock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewVectorClock", reflect.TypeOf((*MockContext)(nil).NewVectorClock))
}

// SetCurrentTime mocks base method.
func (m *MockContext) SetCurrentTime(cluster string, currentTime time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentTime", cluster, currentTime)
}

// SetCurrentTime indicates an expected call of SetCurrentTime.
func (mr *MockContextMockRecorder) SetCurrentTime(cluster, currentTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentTime", reflect.TypeOf((*MockContext)(nil).SetCurrentTime), cluster, currentTime)
}

// SetWorkflowExecution mocks base method.
func (m *MockContext) SetWorkflowExecution(ctx context.Context, request *persistence.SetWorkflowExecutionRequest) (*persistence.SetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*persistence.SetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetWorkflowExecution indicates an expected call of SetWorkflowExecution.
func (mr *MockContextMockRecorder) SetWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkflowExecution", reflect.TypeOf((*MockContext)(nil).SetWorkflowExecution), ctx, request)
}

// Unload mocks base method.
func (m *MockContext) Unload() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unload")
}

// Unload indicates an expected call of Unload.
func (mr *MockContextMockRecorder) Unload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unload", reflect.TypeOf((*MockContext)(nil).Unload))
}

// UpdateFailoverLevel mocks base method.
func (m *MockContext) UpdateFailoverLevel(category tasks.Category, failoverID string, level persistence.FailoverLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFailoverLevel", category, failoverID, level)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFailoverLevel indicates an expected call of UpdateFailoverLevel.
func (mr *MockContextMockRecorder) UpdateFailoverLevel(category, failoverID, level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFailoverLevel", reflect.TypeOf((*MockContext)(nil).UpdateFailoverLevel), category, failoverID, level)
}

// UpdateHandoverNamespaces mocks base method.
func (m *MockContext) UpdateHandoverNamespaces(newNamespaces []*namespace.Namespace, maxRepTaskID int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateHandoverNamespaces", newNamespaces, maxRepTaskID)
}

// UpdateHandoverNamespaces indicates an expected call of UpdateHandoverNamespaces.
func (mr *MockContextMockRecorder) UpdateHandoverNamespaces(newNamespaces, maxRepTaskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHandoverNamespaces", reflect.TypeOf((*MockContext)(nil).UpdateHandoverNamespaces), newNamespaces, maxRepTaskID)
}

// UpdateNamespaceNotificationVersion mocks base method.
func (m *MockContext) UpdateNamespaceNotificationVersion(namespaceNotificationVersion int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNamespaceNotificationVersion", namespaceNotificationVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNamespaceNotificationVersion indicates an expected call of UpdateNamespaceNotificationVersion.
func (mr *MockContextMockRecorder) UpdateNamespaceNotificationVersion(namespaceNotificationVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespaceNotificationVersion", reflect.TypeOf((*MockContext)(nil).UpdateNamespaceNotificationVersion), namespaceNotificationVersion)
}

// UpdateQueueAckLevel mocks base method.
func (m *MockContext) UpdateQueueAckLevel(category tasks.Category, ackLevel tasks.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQueueAckLevel", category, ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQueueAckLevel indicates an expected call of UpdateQueueAckLevel.
func (mr *MockContextMockRecorder) UpdateQueueAckLevel(category, ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQueueAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateQueueAckLevel), category, ackLevel)
}

// UpdateQueueClusterAckLevel mocks base method.
func (m *MockContext) UpdateQueueClusterAckLevel(category tasks.Category, cluster string, ackLevel tasks.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQueueClusterAckLevel", category, cluster, ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQueueClusterAckLevel indicates an expected call of UpdateQueueClusterAckLevel.
func (mr *MockContextMockRecorder) UpdateQueueClusterAckLevel(category, cluster, ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQueueClusterAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateQueueClusterAckLevel), category, cluster, ackLevel)
}

// UpdateQueueState mocks base method.
func (m *MockContext) UpdateQueueState(category tasks.Category, state *v13.QueueState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQueueState", category, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQueueState indicates an expected call of UpdateQueueState.
func (mr *MockContextMockRecorder) UpdateQueueState(category, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQueueState", reflect.TypeOf((*MockContext)(nil).UpdateQueueState), category, state)
}

// UpdateRemoteClusterInfo mocks base method.
func (m *MockContext) UpdateRemoteClusterInfo(cluster string, ackTaskID int64, ackTimestamp time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateRemoteClusterInfo", cluster, ackTaskID, ackTimestamp)
}

// UpdateRemoteClusterInfo indicates an expected call of UpdateRemoteClusterInfo.
func (mr *MockContextMockRecorder) UpdateRemoteClusterInfo(cluster, ackTaskID, ackTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemoteClusterInfo", reflect.TypeOf((*MockContext)(nil).UpdateRemoteClusterInfo), cluster, ackTaskID, ackTimestamp)
}

// UpdateReplicatorDLQAckLevel mocks base method.
func (m *MockContext) UpdateReplicatorDLQAckLevel(sourCluster string, ackLevel int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReplicatorDLQAckLevel", sourCluster, ackLevel)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReplicatorDLQAckLevel indicates an expected call of UpdateReplicatorDLQAckLevel.
func (mr *MockContextMockRecorder) UpdateReplicatorDLQAckLevel(sourCluster, ackLevel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReplicatorDLQAckLevel", reflect.TypeOf((*MockContext)(nil).UpdateReplicatorDLQAckLevel), sourCluster, ackLevel)
}

// UpdateScheduledQueueExclusiveHighReadWatermark mocks base method.
func (m *MockContext) UpdateScheduledQueueExclusiveHighReadWatermark(cluster string, singleProcessorMode bool) (tasks.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduledQueueExclusiveHighReadWatermark", cluster, singleProcessorMode)
	ret0, _ := ret[0].(tasks.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduledQueueExclusiveHighReadWatermark indicates an expected call of UpdateScheduledQueueExclusiveHighReadWatermark.
func (mr *MockContextMockRecorder) UpdateScheduledQueueExclusiveHighReadWatermark(cluster, singleProcessorMode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduledQueueExclusiveHighReadWatermark", reflect.TypeOf((*MockContext)(nil).UpdateScheduledQueueExclusiveHighReadWatermark), cluster, singleProcessorMode)
}

// UpdateWorkflowExecution mocks base method.
func (m *MockContext) UpdateWorkflowExecution(ctx context.Context, request *persistence.UpdateWorkflowExecutionRequest) (*persistence.UpdateWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowExecution", ctx, request)
	ret0, _ := ret[0].(*persistence.UpdateWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkflowExecution indicates an expected call of UpdateWorkflowExecution.
func (mr *MockContextMockRecorder) UpdateWorkflowExecution(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowExecution", reflect.TypeOf((*MockContext)(nil).UpdateWorkflowExecution), ctx, request)
}