// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PathChecker is an autogenerated mock type for the PathChecker type
type PathChecker struct {
	mock.Mock
}

// IsPathExists provides a mock function with given fields: pth
func (_m *PathChecker) IsPathExists(pth string) (bool, error) {
	ret := _m.Called(pth)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(pth)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
