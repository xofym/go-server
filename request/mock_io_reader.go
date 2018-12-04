// Code generated by mockery v1.0.0. DO NOT EDIT.

package request

import mock "github.com/stretchr/testify/mock"

// mockIoReader is an autogenerated mock type for the ioReader type
type mockIoReader struct {
	mock.Mock
}

// Read provides a mock function with given fields: p
func (_m *mockIoReader) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
