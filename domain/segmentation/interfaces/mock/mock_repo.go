// Code generated by MockGen. DO NOT EDIT.
// Source: segmentation/interfaces/repository.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/tokopedia/go-test/domain/segmentation/model"
)

// MockSegmentationDBRepository is a mock of SegmentationDBRepository interface.
type MockSegmentationDBRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSegmentationDBRepositoryMockRecorder
}

// MockSegmentationDBRepositoryMockRecorder is the mock recorder for MockSegmentationDBRepository.
type MockSegmentationDBRepositoryMockRecorder struct {
	mock *MockSegmentationDBRepository
}

// NewMockSegmentationDBRepository creates a new mock instance.
func NewMockSegmentationDBRepository(ctrl *gomock.Controller) *MockSegmentationDBRepository {
	mock := &MockSegmentationDBRepository{ctrl: ctrl}
	mock.recorder = &MockSegmentationDBRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSegmentationDBRepository) EXPECT() *MockSegmentationDBRepositoryMockRecorder {
	return m.recorder
}

// GetSegmentData mocks base method.
func (m *MockSegmentationDBRepository) GetSegmentData(param model.SegmentDBRepoGetSegmentDataParam) (model.SegmentDBRepoGetSegmentDataResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegmentData", param)
	ret0, _ := ret[0].(model.SegmentDBRepoGetSegmentDataResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegmentData indicates an expected call of GetSegmentData.
func (mr *MockSegmentationDBRepositoryMockRecorder) GetSegmentData(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegmentData", reflect.TypeOf((*MockSegmentationDBRepository)(nil).GetSegmentData), param)
}