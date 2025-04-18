// Copyright (C) 2019-2021, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: admin/client.go

// Package admin is a generated GoMock package.
package admin

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// LoadVMs mocks base method.
func (m *MockClient) LoadVMs() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadVMs")
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadVMs indicates an expected call of LoadVMs.
func (mr *MockClientMockRecorder) LoadVMs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadVMs", reflect.TypeOf((*MockClient)(nil).LoadVMs))
}

// WhitelistSubnet mocks base method.
func (m *MockClient) WhitelistSubnet(subnetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhitelistSubnet", subnetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// WhitelistSubnet indicates an expected call of WhitelistSubnet.
func (mr *MockClientMockRecorder) WhitelistSubnet(subnetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhitelistSubnet", reflect.TypeOf((*MockClient)(nil).WhitelistSubnet), subnetID)
}
