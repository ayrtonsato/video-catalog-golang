// Code generated by MockGen. DO NOT EDIT.
// Source: category_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	models "github.com/ayrtonsato/video-catalog-golang/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCategory is a mock of Category interface
type MockCategory struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryMockRecorder
}

// MockCategoryMockRecorder is the mock recorder for MockCategory
type MockCategoryMockRecorder struct {
	mock *MockCategory
}

// NewMockCategory creates a new mock instance
func NewMockCategory(ctrl *gomock.Controller) *MockCategory {
	mock := &MockCategory{ctrl: ctrl}
	mock.recorder = &MockCategoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCategory) EXPECT() *MockCategoryMockRecorder {
	return m.recorder
}

// GetCategories mocks base method
func (m *MockCategory) GetCategories() ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories")
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories
func (mr *MockCategoryMockRecorder) GetCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategory)(nil).GetCategories))
}

// Save mocks base method
func (m *MockCategory) Save(name, description string) (models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", name, description)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockCategoryMockRecorder) Save(name, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCategory)(nil).Save), name, description)
}
