// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	xcodebuild "github.com/bitrise-steplib/steps-xcode-test/xcodebuild"
	mock "github.com/stretchr/testify/mock"
)

// Xcodebuild is an autogenerated mock type for the Xcodebuild type
type Xcodebuild struct {
	mock.Mock
}

// RunTest provides a mock function with given fields: params
func (_m *Xcodebuild) RunTest(params xcodebuild.TestRunParams) (string, int, error) {
	ret := _m.Called(params)

	var r0 string
	if rf, ok := ret.Get(0).(func(xcodebuild.TestRunParams) string); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(xcodebuild.TestRunParams) int); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(xcodebuild.TestRunParams) error); ok {
		r2 = rf(params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Version provides a mock function with given fields:
func (_m *Xcodebuild) Version() (xcodebuild.Version, error) {
	ret := _m.Called()

	var r0 xcodebuild.Version
	if rf, ok := ret.Get(0).(func() xcodebuild.Version); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(xcodebuild.Version)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewXcodebuild interface {
	mock.TestingT
	Cleanup(func())
}

// NewXcodebuild creates a new instance of Xcodebuild. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewXcodebuild(t mockConstructorTestingTNewXcodebuild) *Xcodebuild {
	mock := &Xcodebuild{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
