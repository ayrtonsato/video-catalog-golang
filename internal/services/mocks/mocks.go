// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/category_service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	models "github.com/ayrtonsato/video-catalog-golang/internal/models"
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReaderCategory is a mock of ReaderCategory interface
type MockReaderCategory struct {
	ctrl     *gomock.Controller
	recorder *MockReaderCategoryMockRecorder
}

// MockReaderCategoryMockRecorder is the mock recorder for MockReaderCategory
type MockReaderCategoryMockRecorder struct {
	mock *MockReaderCategory
}

// NewMockReaderCategory creates a new mock instance
func NewMockReaderCategory(ctrl *gomock.Controller) *MockReaderCategory {
	mock := &MockReaderCategory{ctrl: ctrl}
	mock.recorder = &MockReaderCategoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReaderCategory) EXPECT() *MockReaderCategoryMockRecorder {
	return m.recorder
}

// GetCategories mocks base method
func (m *MockReaderCategory) GetCategories() ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories")
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories
func (mr *MockReaderCategoryMockRecorder) GetCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockReaderCategory)(nil).GetCategories))
}

// MockWriterCategory is a mock of WriterCategory interface
type MockWriterCategory struct {
	ctrl     *gomock.Controller
	recorder *MockWriterCategoryMockRecorder
}

// MockWriterCategoryMockRecorder is the mock recorder for MockWriterCategory
type MockWriterCategoryMockRecorder struct {
	mock *MockWriterCategory
}

// NewMockWriterCategory creates a new mock instance
func NewMockWriterCategory(ctrl *gomock.Controller) *MockWriterCategory {
	mock := &MockWriterCategory{ctrl: ctrl}
	mock.recorder = &MockWriterCategoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriterCategory) EXPECT() *MockWriterCategoryMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockWriterCategory) Save(name, description string) (models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", name, description)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockWriterCategoryMockRecorder) Save(name, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockWriterCategory)(nil).Save), name, description)
}

// Update mocks base method
func (m *MockWriterCategory) Update(id uuid.UUID, fields ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockWriterCategoryMockRecorder) Update(id interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWriterCategory)(nil).Update), varargs...)
}
