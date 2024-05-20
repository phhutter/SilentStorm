// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prometheus/alertmanager/api/v2/client/alert (interfaces: ClientService)
//
// Generated by this command:
//
//	mockgen -destination internal/mocks/alertmanager/alert/mock.go github.com/prometheus/alertmanager/api/v2/client/alert ClientService
//

// Package mock_alert is a generated GoMock package.
package mock_alert

import (
	reflect "reflect"

	runtime "github.com/go-openapi/runtime"
	alert "github.com/prometheus/alertmanager/api/v2/client/alert"
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

// GetAlerts mocks base method.
func (m *MockClientService) GetAlerts(arg0 *alert.GetAlertsParams, arg1 ...alert.ClientOption) (*alert.GetAlertsOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAlerts", varargs...)
	ret0, _ := ret[0].(*alert.GetAlertsOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlerts indicates an expected call of GetAlerts.
func (mr *MockClientServiceMockRecorder) GetAlerts(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlerts", reflect.TypeOf((*MockClientService)(nil).GetAlerts), varargs...)
}

// PostAlerts mocks base method.
func (m *MockClientService) PostAlerts(arg0 *alert.PostAlertsParams, arg1 ...alert.ClientOption) (*alert.PostAlertsOK, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostAlerts", varargs...)
	ret0, _ := ret[0].(*alert.PostAlertsOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostAlerts indicates an expected call of PostAlerts.
func (mr *MockClientServiceMockRecorder) PostAlerts(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostAlerts", reflect.TypeOf((*MockClientService)(nil).PostAlerts), varargs...)
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
