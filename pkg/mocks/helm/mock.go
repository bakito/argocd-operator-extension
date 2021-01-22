// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/snorwin/argocd-operator-extension/pkg/helm (interfaces: Client)

// Package mock_helm is a generated GoMock package.
package mock_helm

import (
	gomock "github.com/golang/mock/gomock"
	chart "helm.sh/helm/v3/pkg/chart"
	chartutil "helm.sh/helm/v3/pkg/chartutil"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Install mocks base method
func (m *MockClient) Install(arg0 string, arg1 *chart.Chart, arg2 chartutil.Values) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockClientMockRecorder) Install(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockClient)(nil).Install), arg0, arg1, arg2)
}

// Uninstall mocks base method
func (m *MockClient) Uninstall(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall
func (mr *MockClientMockRecorder) Uninstall(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockClient)(nil).Uninstall), arg0)
}

// Upgrade mocks base method
func (m *MockClient) Upgrade(arg0 string, arg1 *chart.Chart, arg2 chartutil.Values, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade
func (mr *MockClientMockRecorder) Upgrade(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockClient)(nil).Upgrade), arg0, arg1, arg2, arg3)
}
