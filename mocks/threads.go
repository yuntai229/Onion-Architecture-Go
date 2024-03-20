// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/ports/threads.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	model "onion-architecrure-go/domain/model"
	dto "onion-architecrure-go/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockThreadRepo is a mock of ThreadRepo interface.
type MockThreadRepo struct {
	ctrl     *gomock.Controller
	recorder *MockThreadRepoMockRecorder
}

// MockThreadRepoMockRecorder is the mock recorder for MockThreadRepo.
type MockThreadRepoMockRecorder struct {
	mock *MockThreadRepo
}

// NewMockThreadRepo creates a new mock instance.
func NewMockThreadRepo(ctrl *gomock.Controller) *MockThreadRepo {
	mock := &MockThreadRepo{ctrl: ctrl}
	mock.recorder = &MockThreadRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThreadRepo) EXPECT() *MockThreadRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockThreadRepo) Create(requestId string, threadData model.Threads) *model.ErrorMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", requestId, threadData)
	ret0, _ := ret[0].(*model.ErrorMessage)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockThreadRepoMockRecorder) Create(requestId, threadData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockThreadRepo)(nil).Create), requestId, threadData)
}

// GetByUserId mocks base method.
func (m *MockThreadRepo) GetByUserId(requestId string, pagination *model.Pagination, userId uint) ([]model.Threads, *model.ErrorMessage) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", requestId, pagination, userId)
	ret0, _ := ret[0].([]model.Threads)
	ret1, _ := ret[1].(*model.ErrorMessage)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockThreadRepoMockRecorder) GetByUserId(requestId, pagination, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockThreadRepo)(nil).GetByUserId), requestId, pagination, userId)
}

// MockThreadApp is a mock of ThreadApp interface.
type MockThreadApp struct {
	ctrl     *gomock.Controller
	recorder *MockThreadAppMockRecorder
}

// MockThreadAppMockRecorder is the mock recorder for MockThreadApp.
type MockThreadAppMockRecorder struct {
	mock *MockThreadApp
}

// NewMockThreadApp creates a new mock instance.
func NewMockThreadApp(ctrl *gomock.Controller) *MockThreadApp {
	mock := &MockThreadApp{ctrl: ctrl}
	mock.recorder = &MockThreadAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThreadApp) EXPECT() *MockThreadAppMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockThreadApp) CreatePost(requestId string, requestBody dto.PostRequest, userId uint) *model.ErrorMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", requestId, requestBody, userId)
	ret0, _ := ret[0].(*model.ErrorMessage)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockThreadAppMockRecorder) CreatePost(requestId, requestBody, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockThreadApp)(nil).CreatePost), requestId, requestBody, userId)
}

// GetPost mocks base method.
func (m *MockThreadApp) GetPost(requestId string, pagination *model.Pagination, params dto.GetPostRequest) ([]model.Threads, *model.ErrorMessage) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", requestId, pagination, params)
	ret0, _ := ret[0].([]model.Threads)
	ret1, _ := ret[1].(*model.ErrorMessage)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockThreadAppMockRecorder) GetPost(requestId, pagination, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockThreadApp)(nil).GetPost), requestId, pagination, params)
}
