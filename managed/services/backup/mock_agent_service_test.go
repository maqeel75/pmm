// Code generated by mockery. DO NOT EDIT.

package backup

import (
	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
)

// mockAgentService is an autogenerated mock type for the agentService type
type mockAgentService struct {
	mock.Mock
}

// PBMSwitchPITR provides a mock function with given fields: pmmAgentID, dsn, files, tdp, enabled
func (_m *mockAgentService) PBMSwitchPITR(pmmAgentID string, dsn string, files map[string]string, tdp *models.DelimiterPair, enabled bool) error {
	ret := _m.Called(pmmAgentID, dsn, files, tdp, enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]string, *models.DelimiterPair, bool) error); ok {
		r0 = rf(pmmAgentID, dsn, files, tdp, enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockAgentService creates a new instance of mockAgentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAgentService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAgentService {
	mock := &mockAgentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
