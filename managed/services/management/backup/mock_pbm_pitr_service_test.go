// Code generated by mockery. DO NOT EDIT.

package backup

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
	servicesbackup "github.com/percona/pmm/managed/services/backup"
)

// mockPbmPITRService is an autogenerated mock type for the pbmPITRService type
type mockPbmPITRService struct {
	mock.Mock
}

// ListPITRTimeranges provides a mock function with given fields: ctx, locationClient, location, artifact
func (_m *mockPbmPITRService) ListPITRTimeranges(ctx context.Context, locationClient servicesbackup.Storage, location *models.BackupLocation, artifact *models.Artifact) ([]servicesbackup.Timeline, error) {
	ret := _m.Called(ctx, locationClient, location, artifact)

	var r0 []servicesbackup.Timeline
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, servicesbackup.Storage, *models.BackupLocation, *models.Artifact) ([]servicesbackup.Timeline, error)); ok {
		return rf(ctx, locationClient, location, artifact)
	}
	if rf, ok := ret.Get(0).(func(context.Context, servicesbackup.Storage, *models.BackupLocation, *models.Artifact) []servicesbackup.Timeline); ok {
		r0 = rf(ctx, locationClient, location, artifact)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]servicesbackup.Timeline)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, servicesbackup.Storage, *models.BackupLocation, *models.Artifact) error); ok {
		r1 = rf(ctx, locationClient, location, artifact)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockPbmPITRService creates a new instance of mockPbmPITRService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockPbmPITRService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockPbmPITRService {
	mock := &mockPbmPITRService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
