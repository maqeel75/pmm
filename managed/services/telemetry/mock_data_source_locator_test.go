// Code generated by mockery. DO NOT EDIT.

package telemetry

import mock "github.com/stretchr/testify/mock"

// mockDataSourceLocator is an autogenerated mock type for the DataSourceLocator type
type mockDataSourceLocator struct {
	mock.Mock
}

// LocateTelemetryDataSource provides a mock function with given fields: name
func (_m *mockDataSourceLocator) LocateTelemetryDataSource(name string) (DataSource, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for LocateTelemetryDataSource")
	}

	var r0 DataSource
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (DataSource, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) DataSource); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(DataSource)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockDataSourceLocator creates a new instance of mockDataSourceLocator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockDataSourceLocator(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockDataSourceLocator {
	mock := &mockDataSourceLocator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
