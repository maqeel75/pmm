// Code generated by mockery. DO NOT EDIT.

package management

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	reform "gopkg.in/reform.v1"

	models "github.com/percona/pmm/managed/models"
)

// mockConnectionChecker is an autogenerated mock type for the connectionChecker type
type mockConnectionChecker struct {
	mock.Mock
}

// CheckConnectionToService provides a mock function with given fields: ctx, q, service, agent
func (_m *mockConnectionChecker) CheckConnectionToService(ctx context.Context, q *reform.Querier, service *models.Service, agent *models.Agent) error {
	ret := _m.Called(ctx, q, service, agent)

	if len(ret) == 0 {
		panic("no return value specified for CheckConnectionToService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reform.Querier, *models.Service, *models.Agent) error); ok {
		r0 = rf(ctx, q, service, agent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockConnectionChecker creates a new instance of mockConnectionChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockConnectionChecker(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockConnectionChecker {
	mock := &mockConnectionChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
