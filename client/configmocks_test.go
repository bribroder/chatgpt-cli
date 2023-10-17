// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bribroder/chatgpt-cli/config (interfaces: ConfigStore)

// Package client_test is a generated GoMock package.
package client_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/bribroder/chatgpt-cli/types"
)

// MockConfigStore is a mock of ConfigStore interface.
type MockConfigStore struct {
	ctrl     *gomock.Controller
	recorder *MockConfigStoreMockRecorder
}

// MockConfigStoreMockRecorder is the mock recorder for MockConfigStore.
type MockConfigStoreMockRecorder struct {
	mock *MockConfigStore
}

// NewMockConfigStore creates a new mock instance.
func NewMockConfigStore(ctrl *gomock.Controller) *MockConfigStore {
	mock := &MockConfigStore{ctrl: ctrl}
	mock.recorder = &MockConfigStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigStore) EXPECT() *MockConfigStoreMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockConfigStore) Read() (types.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].(types.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockConfigStoreMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockConfigStore)(nil).Read))
}

// ReadDefaults mocks base method.
func (m *MockConfigStore) ReadDefaults() types.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDefaults")
	ret0, _ := ret[0].(types.Config)
	return ret0
}

// ReadDefaults indicates an expected call of ReadDefaults.
func (mr *MockConfigStoreMockRecorder) ReadDefaults() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDefaults", reflect.TypeOf((*MockConfigStore)(nil).ReadDefaults))
}

// Write mocks base method.
func (m *MockConfigStore) Write(arg0 types.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockConfigStoreMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockConfigStore)(nil).Write), arg0)
}
