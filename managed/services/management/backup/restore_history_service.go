// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package backup provides backup functionality.
package backup

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/reform.v1"

	backupv1 "github.com/percona/pmm/api/backup/v1"
	"github.com/percona/pmm/managed/models"
)

// RestoreService represents backup restore API.
type RestoreService struct {
	l  *logrus.Entry
	db *reform.DB

	backupv1.UnimplementedRestoreServiceServer
}

// NewRestoreService creates new restore API service.
func NewRestoreService(db *reform.DB) *RestoreService {
	return &RestoreService{
		l:  logrus.WithField("component", "management/backup/restore_history"),
		db: db,
	}
}

// Enabled returns if service is enabled and can be used.
func (s *RestoreService) Enabled() bool {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		s.l.WithError(err).Error("can't get settings")
		return false
	}
	return settings.IsBackupManagementEnabled()
}

// ListRestores returns a list of restores.
func (s *RestoreService) ListRestores(ctx context.Context, _ *backupv1.ListRestoresRequest) (*backupv1.ListRestoresResponse, error) {
	var items []*models.RestoreHistoryItem
	var services map[string]*models.Service
	var artifacts map[string]*models.Artifact
	var locationModels map[string]*models.BackupLocation

	err := s.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		q := tx.Querier

		var err error
		items, err = models.FindRestoreHistoryItems(q, models.RestoreHistoryItemFilters{})
		if err != nil {
			return err
		}

		artifactIDs := make([]string, 0, len(items))
		serviceIDs := make([]string, 0, len(items))
		for _, i := range items {
			artifactIDs = append(artifactIDs, i.ArtifactID)
			serviceIDs = append(serviceIDs, i.ServiceID)
		}
		artifacts, err = models.FindArtifactsByIDs(q, artifactIDs)
		if err != nil {
			return err
		}

		locationIDs := make([]string, 0, len(artifacts))
		for _, a := range artifacts {
			locationIDs = append(locationIDs, a.LocationID)
		}
		locationModels, err = models.FindBackupLocationsByIDs(q, locationIDs)
		if err != nil {
			return err
		}

		services, err = models.FindServicesByIDs(q, serviceIDs)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	artifactsResponse := make([]*backupv1.RestoreHistoryItem, 0, len(artifacts))
	for _, i := range items {
		convertedArtifact, err := convertRestoreHistoryItem(i, services, artifacts, locationModels)
		if err != nil {
			return nil, err
		}

		artifactsResponse = append(artifactsResponse, convertedArtifact)
	}
	return &backupv1.ListRestoresResponse{
		Items: artifactsResponse,
	}, nil
}

// GetLogs returns logs from the underlying tools for a backup/restore job.
func (s *RestoreService) GetLogs(_ context.Context, req *backupv1.RestoreGetLogsRequest) (*backupv1.RestoreGetLogsResponse, error) {
	jobsFilter := models.JobsFilter{
		Types: []models.JobType{
			models.MySQLBackupJob,
			models.MongoDBBackupJob,
			models.MongoDBRestoreBackupJob,
		},
	}

	jobsFilter.RestoreID = models.NormalizeRestoreID(req.RestoreId)

	jobs, err := models.FindJobs(s.db.Querier, jobsFilter)
	if err != nil {
		return nil, err
	}
	if len(jobs) == 0 {
		return nil, status.Error(codes.NotFound, "Job related to artifact was not found.")
	}
	if len(jobs) > 1 {
		s.l.Warn("provided ID appears in more than one job")
	}

	filter := models.JobLogsFilter{
		JobID:  jobs[0].ID,
		Offset: int(req.Offset),
	}
	if req.Limit > 0 {
		filter.Limit = pointer.ToInt(int(req.Limit))
	}

	jobLogs, err := models.FindJobLogs(s.db.Querier, filter)
	if err != nil {
		return nil, err
	}

	res := &backupv1.RestoreGetLogsResponse{
		Logs: make([]*backupv1.LogChunk, 0, len(jobLogs)),
	}
	for _, log := range jobLogs {
		if log.LastChunk {
			res.End = true
			break
		}
		res.Logs = append(res.Logs, &backupv1.LogChunk{
			ChunkId: uint32(log.ChunkID),
			Data:    log.Data,
		})
	}

	return res, nil
}

func convertRestoreStatus(status models.RestoreStatus) (*backupv1.RestoreStatus, error) {
	var s backupv1.RestoreStatus
	switch status {
	case models.InProgressRestoreStatus:
		s = backupv1.RestoreStatus_RESTORE_STATUS_IN_PROGRESS
	case models.SuccessRestoreStatus:
		s = backupv1.RestoreStatus_RESTORE_STATUS_SUCCESS
	case models.ErrorRestoreStatus:
		s = backupv1.RestoreStatus_RESTORE_STATUS_ERROR
	default:
		return nil, errors.Errorf("invalid status '%s'", status)
	}

	return &s, nil
}

//nolint:funlen
func convertRestoreHistoryItem(
	i *models.RestoreHistoryItem,
	services map[string]*models.Service,
	artifacts map[string]*models.Artifact,
	locations map[string]*models.BackupLocation,
) (*backupv1.RestoreHistoryItem, error) {
	startedAt := timestamppb.New(i.StartedAt)
	if err := startedAt.CheckValid(); err != nil {
		return nil, errors.Wrap(err, "failed to convert startedAt timestamp")
	}

	var finishedAt *timestamppb.Timestamp
	if i.FinishedAt != nil {
		finishedAt = timestamppb.New(*i.FinishedAt)
		if err := finishedAt.CheckValid(); err != nil {
			return nil, errors.Wrap(err, "failed to convert finishedAt timestamp")
		}
	}

	var pitrTimestamp *timestamppb.Timestamp
	if i.PITRTimestamp != nil {
		pitrTimestamp = timestamppb.New(*i.PITRTimestamp)
		if err := pitrTimestamp.CheckValid(); err != nil {
			return nil, errors.Wrap(err, "failed to convert PITR timestamp")
		}
	}

	artifact, ok := artifacts[i.ArtifactID]
	if !ok {
		return nil, errors.Errorf(
			"failed to convert restore history item with id '%s': no artifact id '%s' in the map", i.ID, i.ArtifactID)
	}

	l, ok := locations[artifact.LocationID]
	if !ok {
		return nil, errors.Errorf(
			"failed to convert restore history item with id '%s': no location id '%s' in the map",
			i.ID, artifact.LocationID)
	}

	s, ok := services[i.ServiceID]
	if !ok {
		return nil, errors.Errorf(
			"failed to convert restore history item with id '%s': no service id '%s' in the map", i.ID, i.ServiceID)
	}

	dm, err := convertDataModel(artifact.DataModel)
	if err != nil {
		return nil, errors.Wrapf(err, "restore history item id '%s'", i.ID)
	}

	status, err := convertRestoreStatus(i.Status)
	if err != nil {
		return nil, errors.Wrapf(err, "restore history item id '%s'", i.ID)
	}

	return &backupv1.RestoreHistoryItem{
		RestoreId:     i.ID,
		ArtifactId:    i.ArtifactID,
		Name:          artifact.Name,
		Vendor:        artifact.Vendor,
		LocationId:    artifact.LocationID,
		LocationName:  l.Name,
		ServiceId:     i.ServiceID,
		ServiceName:   s.ServiceName,
		DataModel:     dm,
		Status:        *status,
		StartedAt:     startedAt,
		FinishedAt:    finishedAt,
		PitrTimestamp: pitrTimestamp,
	}, nil
}

// Check interfaces.
var (
	_ backupv1.RestoreServiceServer = (*RestoreService)(nil)
)
