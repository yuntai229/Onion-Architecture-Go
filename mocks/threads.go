// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/ports/threads.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	entity "onion-architecrure-go/domain/entity"
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
func (m *MockThreadRepo) Create(threadData entity.Threads) *entity.ErrorMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", threadData)
	ret0, _ := ret[0].(*entity.ErrorMessage)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockThreadRepoMockRecorder) Create(threadData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockThreadRepo)(nil).Create), threadData)
}

// GetByUserId mocks base method.
func (m *MockThreadRepo) GetByUserId(pagination *entity.Pagination, userId uint) ([]entity.Threads, *entity.ErrorMessage) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", pagination, userId)
	ret0, _ := ret[0].([]entity.Threads)
	ret1, _ := ret[1].(*entity.ErrorMessage)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockThreadRepoMockRecorder) GetByUserId(pagination, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockThreadRepo)(nil).GetByUserId), pagination, userId)
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
func (m *MockThreadApp) CreatePost(requestBody dto.PostRequest, userId uint) *entity.ErrorMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", requestBody, userId)
	ret0, _ := ret[0].(*entity.ErrorMessage)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockThreadAppMockRecorder) CreatePost(requestBody, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockThreadApp)(nil).CreatePost), requestBody, userId)
}

// GetPost mocks base method.
func (m *MockThreadApp) GetPost(pagination *entity.Pagination, params dto.GetPostRequest) ([]entity.Threads, *entity.ErrorMessage) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", pagination, params)
	ret0, _ := ret[0].([]entity.Threads)
	ret1, _ := ret[1].(*entity.ErrorMessage)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockThreadAppMockRecorder) GetPost(pagination, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockThreadApp)(nil).GetPost), pagination, params)
}
