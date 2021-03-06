// Code generated by MockGen. DO NOT EDIT.
// Source: invitation.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	apperror "github.com/ispec-inc/sample/pkg/apperror"
	model "github.com/ispec-inc/sample/pkg/domain/model"
	reflect "reflect"
)

// MockInvitation is a mock of Invitation interface.
type MockInvitation struct {
	ctrl     *gomock.Controller
	recorder *MockInvitationMockRecorder
}

// MockInvitationMockRecorder is the mock recorder for MockInvitation.
type MockInvitationMockRecorder struct {
	mock *MockInvitation
}

// NewMockInvitation creates a new mock instance.
func NewMockInvitation(ctrl *gomock.Controller) *MockInvitation {
	mock := &MockInvitation{ctrl: ctrl}
	mock.recorder = &MockInvitationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvitation) EXPECT() *MockInvitationMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockInvitation) Find(id int64) (model.Invitation, apperror.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(model.Invitation)
	ret1, _ := ret[1].(apperror.Error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockInvitationMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockInvitation)(nil).Find), id)
}

// FindByUserID mocks base method.
func (m *MockInvitation) FindByUserID(uid int64) (model.Invitation, apperror.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", uid)
	ret0, _ := ret[0].(model.Invitation)
	ret1, _ := ret[1].(apperror.Error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID.
func (mr *MockInvitationMockRecorder) FindByUserID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockInvitation)(nil).FindByUserID), uid)
}

// Create mocks base method.
func (m *MockInvitation) Create(minv model.Invitation) apperror.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", minv)
	ret0, _ := ret[0].(apperror.Error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockInvitationMockRecorder) Create(minv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInvitation)(nil).Create), minv)
}
