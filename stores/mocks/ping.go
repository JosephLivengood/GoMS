// Code generated by MockGen. DO NOT EDIT.
// Source: ping.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "GoMS/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPingStore is a mock of PingStore interface.
type MockPingStore struct {
	ctrl     *gomock.Controller
	recorder *MockPingStoreMockRecorder
}

// MockPingStoreMockRecorder is the mock recorder for MockPingStore.
type MockPingStoreMockRecorder struct {
	mock *MockPingStore
}

// NewMockPingStore creates a new mock instance.
func NewMockPingStore(ctrl *gomock.Controller) *MockPingStore {
	mock := &MockPingStore{ctrl: ctrl}
	mock.recorder = &MockPingStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPingStore) EXPECT() *MockPingStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPingStore) Create(createdBy string) (*models.Ping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", createdBy)
	ret0, _ := ret[0].(*models.Ping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPingStoreMockRecorder) Create(createdBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPingStore)(nil).Create), createdBy)
}

// GetAll mocks base method.
func (m *MockPingStore) GetAll() ([]models.Ping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Ping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPingStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPingStore)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockPingStore) GetById(id int) (*models.Ping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.Ping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockPingStoreMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPingStore)(nil).GetById), id)
}
