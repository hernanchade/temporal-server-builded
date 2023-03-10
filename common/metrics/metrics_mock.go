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
// Source: metrics.go

// Package metrics is a generated GoMock package.
package metrics

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	log "go.temporal.io/server/common/log"
)

// MockMetricsHandler is a mock of MetricsHandler interface.
type MockMetricsHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsHandlerMockRecorder
}

// MockMetricsHandlerMockRecorder is the mock recorder for MockMetricsHandler.
type MockMetricsHandlerMockRecorder struct {
	mock *MockMetricsHandler
}

// NewMockMetricsHandler creates a new mock instance.
func NewMockMetricsHandler(ctrl *gomock.Controller) *MockMetricsHandler {
	mock := &MockMetricsHandler{ctrl: ctrl}
	mock.recorder = &MockMetricsHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricsHandler) EXPECT() *MockMetricsHandlerMockRecorder {
	return m.recorder
}

// Counter mocks base method.
func (m *MockMetricsHandler) Counter(arg0 string) CounterMetric {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Counter", arg0)
	ret0, _ := ret[0].(CounterMetric)
	return ret0
}

// Counter indicates an expected call of Counter.
func (mr *MockMetricsHandlerMockRecorder) Counter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Counter", reflect.TypeOf((*MockMetricsHandler)(nil).Counter), arg0)
}

// Gauge mocks base method.
func (m *MockMetricsHandler) Gauge(arg0 string) GaugeMetric {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gauge", arg0)
	ret0, _ := ret[0].(GaugeMetric)
	return ret0
}

// Gauge indicates an expected call of Gauge.
func (mr *MockMetricsHandlerMockRecorder) Gauge(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockMetricsHandler)(nil).Gauge), arg0)
}

// Histogram mocks base method.
func (m *MockMetricsHandler) Histogram(arg0 string, arg1 MetricUnit) HistogramMetric {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Histogram", arg0, arg1)
	ret0, _ := ret[0].(HistogramMetric)
	return ret0
}

// Histogram indicates an expected call of Histogram.
func (mr *MockMetricsHandlerMockRecorder) Histogram(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Histogram", reflect.TypeOf((*MockMetricsHandler)(nil).Histogram), arg0, arg1)
}

// Stop mocks base method.
func (m *MockMetricsHandler) Stop(arg0 log.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", arg0)
}

// Stop indicates an expected call of Stop.
func (mr *MockMetricsHandlerMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockMetricsHandler)(nil).Stop), arg0)
}

// Timer mocks base method.
func (m *MockMetricsHandler) Timer(arg0 string) TimerMetric {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timer", arg0)
	ret0, _ := ret[0].(TimerMetric)
	return ret0
}

// Timer indicates an expected call of Timer.
func (mr *MockMetricsHandlerMockRecorder) Timer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timer", reflect.TypeOf((*MockMetricsHandler)(nil).Timer), arg0)
}

// WithTags mocks base method.
func (m *MockMetricsHandler) WithTags(arg0 ...Tag) MetricsHandler {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithTags", varargs...)
	ret0, _ := ret[0].(MetricsHandler)
	return ret0
}

// WithTags indicates an expected call of WithTags.
func (mr *MockMetricsHandlerMockRecorder) WithTags(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTags", reflect.TypeOf((*MockMetricsHandler)(nil).WithTags), arg0...)
}

// MockCounterMetric is a mock of CounterMetric interface.
type MockCounterMetric struct {
	ctrl     *gomock.Controller
	recorder *MockCounterMetricMockRecorder
}

// MockCounterMetricMockRecorder is the mock recorder for MockCounterMetric.
type MockCounterMetricMockRecorder struct {
	mock *MockCounterMetric
}

// NewMockCounterMetric creates a new mock instance.
func NewMockCounterMetric(ctrl *gomock.Controller) *MockCounterMetric {
	mock := &MockCounterMetric{ctrl: ctrl}
	mock.recorder = &MockCounterMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCounterMetric) EXPECT() *MockCounterMetricMockRecorder {
	return m.recorder
}

// Record mocks base method.
func (m *MockCounterMetric) Record(arg0 int64, arg1 ...Tag) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Record", varargs...)
}

// Record indicates an expected call of Record.
func (mr *MockCounterMetricMockRecorder) Record(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockCounterMetric)(nil).Record), varargs...)
}

// MockGaugeMetric is a mock of GaugeMetric interface.
type MockGaugeMetric struct {
	ctrl     *gomock.Controller
	recorder *MockGaugeMetricMockRecorder
}

// MockGaugeMetricMockRecorder is the mock recorder for MockGaugeMetric.
type MockGaugeMetricMockRecorder struct {
	mock *MockGaugeMetric
}

// NewMockGaugeMetric creates a new mock instance.
func NewMockGaugeMetric(ctrl *gomock.Controller) *MockGaugeMetric {
	mock := &MockGaugeMetric{ctrl: ctrl}
	mock.recorder = &MockGaugeMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGaugeMetric) EXPECT() *MockGaugeMetricMockRecorder {
	return m.recorder
}

// Record mocks base method.
func (m *MockGaugeMetric) Record(arg0 float64, arg1 ...Tag) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Record", varargs...)
}

// Record indicates an expected call of Record.
func (mr *MockGaugeMetricMockRecorder) Record(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockGaugeMetric)(nil).Record), varargs...)
}

// MockTimerMetric is a mock of TimerMetric interface.
type MockTimerMetric struct {
	ctrl     *gomock.Controller
	recorder *MockTimerMetricMockRecorder
}

// MockTimerMetricMockRecorder is the mock recorder for MockTimerMetric.
type MockTimerMetricMockRecorder struct {
	mock *MockTimerMetric
}

// NewMockTimerMetric creates a new mock instance.
func NewMockTimerMetric(ctrl *gomock.Controller) *MockTimerMetric {
	mock := &MockTimerMetric{ctrl: ctrl}
	mock.recorder = &MockTimerMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimerMetric) EXPECT() *MockTimerMetricMockRecorder {
	return m.recorder
}

// Record mocks base method.
func (m *MockTimerMetric) Record(arg0 time.Duration, arg1 ...Tag) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Record", varargs...)
}

// Record indicates an expected call of Record.
func (mr *MockTimerMetricMockRecorder) Record(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockTimerMetric)(nil).Record), varargs...)
}

// MockHistogramMetric is a mock of HistogramMetric interface.
type MockHistogramMetric struct {
	ctrl     *gomock.Controller
	recorder *MockHistogramMetricMockRecorder
}

// MockHistogramMetricMockRecorder is the mock recorder for MockHistogramMetric.
type MockHistogramMetricMockRecorder struct {
	mock *MockHistogramMetric
}

// NewMockHistogramMetric creates a new mock instance.
func NewMockHistogramMetric(ctrl *gomock.Controller) *MockHistogramMetric {
	mock := &MockHistogramMetric{ctrl: ctrl}
	mock.recorder = &MockHistogramMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHistogramMetric) EXPECT() *MockHistogramMetricMockRecorder {
	return m.recorder
}

// Record mocks base method.
func (m *MockHistogramMetric) Record(arg0 int64, arg1 ...Tag) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Record", varargs...)
}

// Record indicates an expected call of Record.
func (mr *MockHistogramMetricMockRecorder) Record(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockHistogramMetric)(nil).Record), varargs...)
}
