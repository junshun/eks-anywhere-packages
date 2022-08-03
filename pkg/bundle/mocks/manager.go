// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/bundle/manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// ProcessBundle mocks base method.
func (m *MockManager) ProcessBundle(ctx context.Context, newBundle *v1alpha1.PackageBundle) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBundle", ctx, newBundle)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessBundle indicates an expected call of ProcessBundle.
func (mr *MockManagerMockRecorder) ProcessBundle(ctx, newBundle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBundle", reflect.TypeOf((*MockManager)(nil).ProcessBundle), ctx, newBundle)
}

// ProcessBundleController mocks base method.
func (m *MockManager) ProcessBundleController(ctx context.Context, pbc *v1alpha1.PackageBundleController) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBundleController", ctx, pbc)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessBundleController indicates an expected call of ProcessBundleController.
func (mr *MockManagerMockRecorder) ProcessBundleController(ctx, pbc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBundleController", reflect.TypeOf((*MockManager)(nil).ProcessBundleController), ctx, pbc)
}

// SortBundlesDescending mocks base method.
func (m *MockManager) SortBundlesDescending(bundles []v1alpha1.PackageBundle) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SortBundlesDescending", bundles)
}

// SortBundlesDescending indicates an expected call of SortBundlesDescending.
func (mr *MockManagerMockRecorder) SortBundlesDescending(bundles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortBundlesDescending", reflect.TypeOf((*MockManager)(nil).SortBundlesDescending), bundles)
}
