// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prometheus/alertmanager/api/v2/client/silence (interfaces: ClientService)
//
// Generated by this command:
//
//	mockgen -destination internal/mocks/alertmanager/silence/mock.go github.com/prometheus/alertmanager/api/v2/client/silence ClientService
//

// Package mock_silence is a generated GoMock package.
package mock_silence

import (
	reflect "reflect"

	runtime "github.com/go-openapi/runtime"
	silence "github.com/prometheus/alertmanager/api/v2/client/silence"
	gomock "go.uber.org/mock/gomock"
)

// MockClientService is a mock of ClientService interface.
type MockClientService struct {
	ctrl     *gomock.Controller
	recorder *MockClientServiceMockRecorder
}

// MockClientServiceMockRecorder is the mock recorder for MockClientService.
type MockClientServiceMockRecorder struct {
	mock *MockClientService
}

// NewMockClientService creates a new mock instance.
func NewMockClientService(ctrl *gomock.Controller) *MockClientService {
	mock := &MockClientService{ctrl: ctrl}
	mock.recorder = &MockClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientService) EXPECT() *MockClientServiceMockRecorder {
	return m.recorder
}

// DeleteSilence mocks base method.
func (m *MockClientService) DeleteSilence(arg0 *silence.DeleteSilenceParams, arg1 ...silence.ClientOption) (*silence.DeleteSilenceOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSilence", varargs...)
	ret0, _ := ret[0].(*silence.DeleteSilenceOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSilence indicates an expected call of DeleteSilence.
func (mr *MockClientServiceMockRecorder) DeleteSilence(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSilence", reflect.TypeOf((*MockClientService)(nil).DeleteSilence), varargs...)
}

// GetSilence mocks base method.
func (m *MockClientService) GetSilence(arg0 *silence.GetSilenceParams, arg1 ...silence.ClientOption) (*silence.GetSilenceOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSilence", varargs...)
	ret0, _ := ret[0].(*silence.GetSilenceOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSilence indicates an expected call of GetSilence.
func (mr *MockClientServiceMockRecorder) GetSilence(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSilence", reflect.TypeOf((*MockClientService)(nil).GetSilence), varargs...)
}

// GetSilences mocks base method.
func (m *MockClientService) GetSilences(arg0 *silence.GetSilencesParams, arg1 ...silence.ClientOption) (*silence.GetSilencesOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSilences", varargs...)
	ret0, _ := ret[0].(*silence.GetSilencesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSilences indicates an expected call of GetSilences.
func (mr *MockClientServiceMockRecorder) GetSilences(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSilences", reflect.TypeOf((*MockClientService)(nil).GetSilences), varargs...)
}

// PostSilences mocks base method.
func (m *MockClientService) PostSilences(arg0 *silence.PostSilencesParams, arg1 ...silence.ClientOption) (*silence.PostSilencesOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostSilences", varargs...)
	ret0, _ := ret[0].(*silence.PostSilencesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSilences indicates an expected call of PostSilences.
func (mr *MockClientServiceMockRecorder) PostSilences(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSilences", reflect.TypeOf((*MockClientService)(nil).PostSilences), varargs...)
}

// SetTransport mocks base method.
func (m *MockClientService) SetTransport(arg0 runtime.ClientTransport) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransport", arg0)
}

// SetTransport indicates an expected call of SetTransport.
func (mr *MockClientServiceMockRecorder) SetTransport(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransport", reflect.TypeOf((*MockClientService)(nil).SetTransport), arg0)
}