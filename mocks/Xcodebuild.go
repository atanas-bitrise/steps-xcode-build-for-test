// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	serialized "github.com/bitrise-io/go-xcode/xcodeproject/serialized"
	mock "github.com/stretchr/testify/mock"
)

// Xcodebuild is an autogenerated mock type for the Xcodebuild type
type Xcodebuild struct {
	mock.Mock
}

// ShowBuildSettings provides a mock function with given fields: projectPath, scheme, configuration, action, options
func (_m *Xcodebuild) ShowBuildSettings(projectPath string, scheme string, configuration string, action string, options []string) (serialized.Object, error) {
	ret := _m.Called(projectPath, scheme, configuration, action, options)

	var r0 serialized.Object
	if rf, ok := ret.Get(0).(func(string, string, string, string, []string) serialized.Object); ok {
		r0 = rf(projectPath, scheme, configuration, action, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(serialized.Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, []string) error); ok {
		r1 = rf(projectPath, scheme, configuration, action, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
