// Copyright (C) 2019-2021, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: checksum/checksum.go

// Package checksum is a generated GoMock package.
package checksum

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChecksummer is a mock of Checksummer interface.
type MockChecksummer struct {
	ctrl     *gomock.Controller
	recorder *MockChecksummerMockRecorder
}

// MockChecksummerMockRecorder is the mock recorder for MockChecksummer.
type MockChecksummerMockRecorder struct {
	mock *MockChecksummer
}

// NewMockChecksummer creates a new mock instance.
func NewMockChecksummer(ctrl *gomock.Controller) *MockChecksummer {
	mock := &MockChecksummer{ctrl: ctrl}
	mock.recorder = &MockChecksummerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChecksummer) EXPECT() *MockChecksummerMockRecorder {
	return m.recorder
}

// Checksum mocks base method.
func (m *MockChecksummer) Checksum(path string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checksum", path)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Checksum indicates an expected call of Checksum.
func (mr *MockChecksummerMockRecorder) Checksum(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checksum", reflect.TypeOf((*MockChecksummer)(nil).Checksum), path)
}
