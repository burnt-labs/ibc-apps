// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cosmos/ibc-go/v6/modules/core/05-port/types (interfaces: IBCModule)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/capability/types"
	types1 "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	exported "github.com/cosmos/ibc-go/v6/modules/core/exported"
	gomock "go.uber.org/mock/gomock"
)

// MockIBCModule is a mock of IBCModule interface.
type MockIBCModule struct {
	ctrl     *gomock.Controller
	recorder *MockIBCModuleMockRecorder
}

// MockIBCModuleMockRecorder is the mock recorder for MockIBCModule.
type MockIBCModuleMockRecorder struct {
	mock *MockIBCModule
}

// NewMockIBCModule creates a new mock instance.
func NewMockIBCModule(ctrl *gomock.Controller) *MockIBCModule {
	mock := &MockIBCModule{ctrl: ctrl}
	mock.recorder = &MockIBCModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBCModule) EXPECT() *MockIBCModuleMockRecorder {
	return m.recorder
}

// OnAcknowledgementPacket mocks base method.
func (m *MockIBCModule) OnAcknowledgementPacket(arg0 types.Context, arg1 types1.Packet, arg2 []byte, arg3 types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnAcknowledgementPacket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnAcknowledgementPacket indicates an expected call of OnAcknowledgementPacket.
func (mr *MockIBCModuleMockRecorder) OnAcknowledgementPacket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAcknowledgementPacket", reflect.TypeOf((*MockIBCModule)(nil).OnAcknowledgementPacket), arg0, arg1, arg2, arg3)
}

// OnChanCloseConfirm mocks base method.
func (m *MockIBCModule) OnChanCloseConfirm(arg0 types.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChanCloseConfirm", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnChanCloseConfirm indicates an expected call of OnChanCloseConfirm.
func (mr *MockIBCModuleMockRecorder) OnChanCloseConfirm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChanCloseConfirm", reflect.TypeOf((*MockIBCModule)(nil).OnChanCloseConfirm), arg0, arg1, arg2)
}

// OnChanCloseInit mocks base method.
func (m *MockIBCModule) OnChanCloseInit(arg0 types.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChanCloseInit", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnChanCloseInit indicates an expected call of OnChanCloseInit.
func (mr *MockIBCModuleMockRecorder) OnChanCloseInit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChanCloseInit", reflect.TypeOf((*MockIBCModule)(nil).OnChanCloseInit), arg0, arg1, arg2)
}

// OnChanOpenAck mocks base method.
func (m *MockIBCModule) OnChanOpenAck(arg0 types.Context, arg1, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChanOpenAck", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnChanOpenAck indicates an expected call of OnChanOpenAck.
func (mr *MockIBCModuleMockRecorder) OnChanOpenAck(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChanOpenAck", reflect.TypeOf((*MockIBCModule)(nil).OnChanOpenAck), arg0, arg1, arg2, arg3, arg4)
}

// OnChanOpenConfirm mocks base method.
func (m *MockIBCModule) OnChanOpenConfirm(arg0 types.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChanOpenConfirm", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnChanOpenConfirm indicates an expected call of OnChanOpenConfirm.
func (mr *MockIBCModuleMockRecorder) OnChanOpenConfirm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChanOpenConfirm", reflect.TypeOf((*MockIBCModule)(nil).OnChanOpenConfirm), arg0, arg1, arg2)
}

// OnChanOpenInit mocks base method.
func (m *MockIBCModule) OnChanOpenInit(arg0 types.Context, arg1 types1.Order, arg2 []string, arg3, arg4 string, arg5 *types0.Capability, arg6 types1.Counterparty, arg7 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChanOpenInit", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnChanOpenInit indicates an expected call of OnChanOpenInit.
func (mr *MockIBCModuleMockRecorder) OnChanOpenInit(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChanOpenInit", reflect.TypeOf((*MockIBCModule)(nil).OnChanOpenInit), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// OnChanOpenTry mocks base method.
func (m *MockIBCModule) OnChanOpenTry(arg0 types.Context, arg1 types1.Order, arg2 []string, arg3, arg4 string, arg5 *types0.Capability, arg6 types1.Counterparty, arg7 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChanOpenTry", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnChanOpenTry indicates an expected call of OnChanOpenTry.
func (mr *MockIBCModuleMockRecorder) OnChanOpenTry(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChanOpenTry", reflect.TypeOf((*MockIBCModule)(nil).OnChanOpenTry), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// OnRecvPacket mocks base method.
func (m *MockIBCModule) OnRecvPacket(arg0 types.Context, arg1 types1.Packet, arg2 types.AccAddress) exported.Acknowledgement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnRecvPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(exported.Acknowledgement)
	return ret0
}

// OnRecvPacket indicates an expected call of OnRecvPacket.
func (mr *MockIBCModuleMockRecorder) OnRecvPacket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRecvPacket", reflect.TypeOf((*MockIBCModule)(nil).OnRecvPacket), arg0, arg1, arg2)
}

// OnTimeoutPacket mocks base method.
func (m *MockIBCModule) OnTimeoutPacket(arg0 types.Context, arg1 types1.Packet, arg2 types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnTimeoutPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnTimeoutPacket indicates an expected call of OnTimeoutPacket.
func (mr *MockIBCModuleMockRecorder) OnTimeoutPacket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnTimeoutPacket", reflect.TypeOf((*MockIBCModule)(nil).OnTimeoutPacket), arg0, arg1, arg2)
}
