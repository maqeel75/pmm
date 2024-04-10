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

package management

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	inventoryv1 "github.com/percona/pmm/api/inventory/v1"
	managementv1 "github.com/percona/pmm/api/management/v1"
	"github.com/percona/pmm/managed/models"
	"github.com/percona/pmm/managed/services"
)

// ManagementService allows to interact with services.
type ManagementService struct { //nolint:revive
	db            *reform.DB
	r             agentsRegistry
	state         agentsStateUpdater
	cc            connectionChecker
	sib           serviceInfoBroker
	vmdb          prometheusService
	vc            versionCache
	grafanaClient grafanaClient
	vmClient      victoriaMetricsClient

	managementv1.UnimplementedManagementServiceServer
}

// NewManagementService creates a ManagementService instance.
func NewManagementService(
	db *reform.DB,
	r agentsRegistry,
	state agentsStateUpdater,
	cc connectionChecker,
	sib serviceInfoBroker,
	vmdb prometheusService,
	vc versionCache,
	grafanaClient grafanaClient,
	vmClient victoriaMetricsClient,
) *ManagementService {
	return &ManagementService{
		db:            db,
		r:             r,
		state:         state,
		cc:            cc,
		sib:           sib,
		vmdb:          vmdb,
		vc:            vc,
		grafanaClient: grafanaClient,
		vmClient:      vmClient,
	}
}

// AddService add a Service and its Agents.
func (s *ManagementService) AddService(ctx context.Context, req *managementv1.AddServiceRequest) (*managementv1.AddServiceResponse, error) {
	switch req.Service.(type) {
	case *managementv1.AddServiceRequest_Mysql:
		return s.addMySQL(ctx, req.GetMysql())
	case *managementv1.AddServiceRequest_Mongodb:
		return s.addMongoDB(ctx, req.GetMongodb())
	case *managementv1.AddServiceRequest_Postgresql:
		return s.addPostgreSQL(ctx, req.GetPostgresql())
	case *managementv1.AddServiceRequest_Proxysql:
		return s.addProxySQL(ctx, req.GetProxysql())
	case *managementv1.AddServiceRequest_Haproxy:
		return s.addHAProxy(ctx, req.GetHaproxy())
	case *managementv1.AddServiceRequest_External:
		return s.addExternal(ctx, req.GetExternal())
	case *managementv1.AddServiceRequest_Rds:
		return s.addRDS(ctx, req.GetRds())
	default:
		return nil, errors.Errorf("invalid request %v", req.GetService())
	}
}

// RemoveService removes a Service along with its Agents.
func (s *ManagementService) RemoveService(ctx context.Context, req *managementv1.RemoveServiceRequest) (*managementv1.RemoveServiceResponse, error) {
	err := s.validateRequest(req)
	if err != nil {
		return nil, err
	}
	pmmAgentIDs := make(map[string]struct{})
	var reloadPrometheusConfig bool

	errTX := s.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		var service *models.Service
		var err error

		if LooksLikeID(req.ServiceId) {
			service, err = models.FindServiceByID(tx.Querier, models.NormalizeServiceID(req.ServiceId))
		} else {
			// if it's not a service ID, it is a service name then
			service, err = models.FindServiceByName(tx.Querier, req.ServiceId)
		}
		if err != nil {
			return err
		}

		if req.ServiceType != inventoryv1.ServiceType_SERVICE_TYPE_UNSPECIFIED {
			err := s.checkServiceType(service, req.ServiceType)
			if err != nil {
				return err
			}
		}

		agents, err := models.FindAgents(tx.Querier, models.AgentFilters{ServiceID: service.ServiceID})
		if err != nil {
			return err
		}
		for _, agent := range agents {
			_, err := models.RemoveAgent(tx.Querier, agent.AgentID, models.RemoveRestrict)
			if err != nil {
				return err
			}
			if agent.PMMAgentID != nil {
				pmmAgentIDs[pointer.GetString(agent.PMMAgentID)] = struct{}{}
			} else {
				reloadPrometheusConfig = true
			}
		}

		err = models.RemoveService(tx.Querier, service.ServiceID, models.RemoveCascade)
		if err != nil {
			return err
		}

		node, err := models.FindNodeByID(s.db.Querier, service.NodeID)
		if err != nil {
			return err
		}

		// For RDS and Azure we also want to remove the node.
		if node.NodeType == models.RemoteRDSNodeType || node.NodeType == models.RemoteAzureDatabaseNodeType {
			agents, err = models.FindAgents(tx.Querier, models.AgentFilters{NodeID: node.NodeID})
			if err != nil {
				return err
			}
			for _, a := range agents {
				_, err := models.RemoveAgent(s.db.Querier, a.AgentID, models.RemoveRestrict)
				if err != nil {
					return err
				}
				if a.PMMAgentID != nil {
					pmmAgentIDs[pointer.GetString(a.PMMAgentID)] = struct{}{}
				}
			}

			if len(pmmAgentIDs) <= 1 {
				if err = models.RemoveNode(tx.Querier, node.NodeID, models.RemoveCascade); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if errTX != nil {
		return nil, errTX
	}

	for agentID := range pmmAgentIDs {
		s.state.RequestStateUpdate(ctx, agentID)
	}
	if reloadPrometheusConfig {
		// It's required to regenerate victoriametrics config file for the agents which aren't run by pmm-agent.
		s.vmdb.RequestConfigurationUpdate()
	}
	return &managementv1.RemoveServiceResponse{}, nil
}

func (s *ManagementService) checkServiceType(service *models.Service, serviceType inventoryv1.ServiceType) error {
	if expected, ok := services.ServiceTypes[serviceType]; ok && expected == service.ServiceType {
		return nil
	}
	return status.Error(codes.InvalidArgument, "wrong service type")
}

func (s *ManagementService) validateRequest(request *managementv1.RemoveServiceRequest) error {
	if request.ServiceId == "" {
		return status.Error(codes.InvalidArgument, "service_id or service_name expected")
	}
	return nil
}
