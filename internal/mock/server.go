// Code generated by MockGen. DO NOT EDIT.
// Source: ./server.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/Donich1987/Image-loader/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// Mockcontroller is a mock of controller interface.
type Mockcontroller struct {
	ctrl     *gomock.Controller
	recorder *MockcontrollerMockRecorder
}

// MockcontrollerMockRecorder is the mock recorder for Mockcontroller.
type MockcontrollerMockRecorder struct {
	mock *Mockcontroller
}

// NewMockcontroller creates a new mock instance.
func NewMockcontroller(ctrl *gomock.Controller) *Mockcontroller {
	mock := &Mockcontroller{ctrl: ctrl}
	mock.recorder = &MockcontrollerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockcontroller) EXPECT() *MockcontrollerMockRecorder {
	return m.recorder
}

// AddFile mocks base method.
func (m *Mockcontroller) AddFile(ctx context.Context, image model.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFile", ctx, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFile indicates an expected call of AddFile.
func (mr *MockcontrollerMockRecorder) AddFile(ctx, image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFile", reflect.TypeOf((*Mockcontroller)(nil).AddFile), ctx, image)
}

// AddUser mocks base method.
func (m *Mockcontroller) AddUser(ctx context.Context, user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockcontrollerMockRecorder) AddUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*Mockcontroller)(nil).AddUser), ctx, user)
}

// Authorize mocks base method.
func (m *Mockcontroller) Authorize(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockcontrollerMockRecorder) Authorize(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*Mockcontroller)(nil).Authorize), ctx, login, password)
}

// DeleteUser mocks base method.
func (m *Mockcontroller) DeleteUser(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockcontrollerMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*Mockcontroller)(nil).DeleteUser), ctx, id)
}

// GetUser mocks base method.
func (m *Mockcontroller) GetUser(ctx context.Context, id int64) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockcontrollerMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*Mockcontroller)(nil).GetUser), ctx, id)
}

// UpdateUser mocks base method.
func (m *Mockcontroller) UpdateUser(ctx context.Context, modelUser model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, modelUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockcontrollerMockRecorder) UpdateUser(ctx, modelUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*Mockcontroller)(nil).UpdateUser), ctx, modelUser)
}
