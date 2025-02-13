// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: Stream)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mockquic -destination quic/stream.go github.com/quic-go/quic-go Stream
//

// Package mockquic is a generated GoMock package.
package mockquic

import (
	context "context"
	reflect "reflect"
	time "time"

	quic "github.com/quic-go/quic-go"
	gomock "go.uber.org/mock/gomock"
)

// MockStream is a mock of Stream interface.
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
	isgomock struct{}
}

// MockStreamMockRecorder is the mock recorder for MockStream.
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance.
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// CancelRead mocks base method.
func (m *MockStream) CancelRead(arg0 quic.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelRead", arg0)
}

// CancelRead indicates an expected call of CancelRead.
func (mr *MockStreamMockRecorder) CancelRead(arg0 any) *MockStreamCancelReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRead", reflect.TypeOf((*MockStream)(nil).CancelRead), arg0)
	return &MockStreamCancelReadCall{Call: call}
}

// MockStreamCancelReadCall wrap *gomock.Call
type MockStreamCancelReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamCancelReadCall) Return() *MockStreamCancelReadCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamCancelReadCall) Do(f func(quic.StreamErrorCode)) *MockStreamCancelReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamCancelReadCall) DoAndReturn(f func(quic.StreamErrorCode)) *MockStreamCancelReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CancelWrite mocks base method.
func (m *MockStream) CancelWrite(arg0 quic.StreamErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelWrite", arg0)
}

// CancelWrite indicates an expected call of CancelWrite.
func (mr *MockStreamMockRecorder) CancelWrite(arg0 any) *MockStreamCancelWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWrite", reflect.TypeOf((*MockStream)(nil).CancelWrite), arg0)
	return &MockStreamCancelWriteCall{Call: call}
}

// MockStreamCancelWriteCall wrap *gomock.Call
type MockStreamCancelWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamCancelWriteCall) Return() *MockStreamCancelWriteCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamCancelWriteCall) Do(f func(quic.StreamErrorCode)) *MockStreamCancelWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamCancelWriteCall) DoAndReturn(f func(quic.StreamErrorCode)) *MockStreamCancelWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockStream) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStreamMockRecorder) Close() *MockStreamCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStream)(nil).Close))
	return &MockStreamCloseCall{Call: call}
}

// MockStreamCloseCall wrap *gomock.Call
type MockStreamCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamCloseCall) Return(arg0 error) *MockStreamCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamCloseCall) Do(f func() error) *MockStreamCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamCloseCall) DoAndReturn(f func() error) *MockStreamCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Context mocks base method.
func (m *MockStream) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockStreamMockRecorder) Context() *MockStreamContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockStream)(nil).Context))
	return &MockStreamContextCall{Call: call}
}

// MockStreamContextCall wrap *gomock.Call
type MockStreamContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamContextCall) Return(arg0 context.Context) *MockStreamContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamContextCall) Do(f func() context.Context) *MockStreamContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamContextCall) DoAndReturn(f func() context.Context) *MockStreamContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Read mocks base method.
func (m *MockStream) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStreamMockRecorder) Read(p any) *MockStreamReadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStream)(nil).Read), p)
	return &MockStreamReadCall{Call: call}
}

// MockStreamReadCall wrap *gomock.Call
type MockStreamReadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamReadCall) Return(n int, err error) *MockStreamReadCall {
	c.Call = c.Call.Return(n, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamReadCall) Do(f func([]byte) (int, error)) *MockStreamReadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamReadCall) DoAndReturn(f func([]byte) (int, error)) *MockStreamReadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDeadline mocks base method.
func (m *MockStream) SetDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockStreamMockRecorder) SetDeadline(t any) *MockStreamSetDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockStream)(nil).SetDeadline), t)
	return &MockStreamSetDeadlineCall{Call: call}
}

// MockStreamSetDeadlineCall wrap *gomock.Call
type MockStreamSetDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamSetDeadlineCall) Return(arg0 error) *MockStreamSetDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamSetDeadlineCall) Do(f func(time.Time) error) *MockStreamSetDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamSetDeadlineCall) DoAndReturn(f func(time.Time) error) *MockStreamSetDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetReadDeadline mocks base method.
func (m *MockStream) SetReadDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockStreamMockRecorder) SetReadDeadline(t any) *MockStreamSetReadDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockStream)(nil).SetReadDeadline), t)
	return &MockStreamSetReadDeadlineCall{Call: call}
}

// MockStreamSetReadDeadlineCall wrap *gomock.Call
type MockStreamSetReadDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamSetReadDeadlineCall) Return(arg0 error) *MockStreamSetReadDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamSetReadDeadlineCall) Do(f func(time.Time) error) *MockStreamSetReadDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamSetReadDeadlineCall) DoAndReturn(f func(time.Time) error) *MockStreamSetReadDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetWriteDeadline mocks base method.
func (m *MockStream) SetWriteDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline.
func (mr *MockStreamMockRecorder) SetWriteDeadline(t any) *MockStreamSetWriteDeadlineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockStream)(nil).SetWriteDeadline), t)
	return &MockStreamSetWriteDeadlineCall{Call: call}
}

// MockStreamSetWriteDeadlineCall wrap *gomock.Call
type MockStreamSetWriteDeadlineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamSetWriteDeadlineCall) Return(arg0 error) *MockStreamSetWriteDeadlineCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamSetWriteDeadlineCall) Do(f func(time.Time) error) *MockStreamSetWriteDeadlineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamSetWriteDeadlineCall) DoAndReturn(f func(time.Time) error) *MockStreamSetWriteDeadlineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StreamID mocks base method.
func (m *MockStream) StreamID() quic.StreamID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(quic.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID.
func (mr *MockStreamMockRecorder) StreamID() *MockStreamStreamIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockStream)(nil).StreamID))
	return &MockStreamStreamIDCall{Call: call}
}

// MockStreamStreamIDCall wrap *gomock.Call
type MockStreamStreamIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamStreamIDCall) Return(arg0 quic.StreamID) *MockStreamStreamIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamStreamIDCall) Do(f func() quic.StreamID) *MockStreamStreamIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamStreamIDCall) DoAndReturn(f func() quic.StreamID) *MockStreamStreamIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockStream) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockStreamMockRecorder) Write(p any) *MockStreamWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStream)(nil).Write), p)
	return &MockStreamWriteCall{Call: call}
}

// MockStreamWriteCall wrap *gomock.Call
type MockStreamWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStreamWriteCall) Return(n int, err error) *MockStreamWriteCall {
	c.Call = c.Call.Return(n, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStreamWriteCall) Do(f func([]byte) (int, error)) *MockStreamWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStreamWriteCall) DoAndReturn(f func([]byte) (int, error)) *MockStreamWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
