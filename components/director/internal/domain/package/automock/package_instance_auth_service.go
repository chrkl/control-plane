// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// PackageInstanceAuthService is an autogenerated mock type for the PackageInstanceAuthService type
type PackageInstanceAuthService struct {
	mock.Mock
}

// GetForPackage provides a mock function with given fields: ctx, id, packageID
func (_m *PackageInstanceAuthService) GetForPackage(ctx context.Context, id string, packageID string) (*model.PackageInstanceAuth, error) {
	ret := _m.Called(ctx, id, packageID)

	var r0 *model.PackageInstanceAuth
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.PackageInstanceAuth); ok {
		r0 = rf(ctx, id, packageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PackageInstanceAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, packageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, id
func (_m *PackageInstanceAuthService) List(ctx context.Context, id string) ([]*model.PackageInstanceAuth, error) {
	ret := _m.Called(ctx, id)

	var r0 []*model.PackageInstanceAuth
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.PackageInstanceAuth); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PackageInstanceAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
