// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inventory

import (
	"testing"

	"github.com/stretchr/testify/require"

	services "github.com/percona/pmm/api/inventory/v1/json/client/services_service"
)

func TestAddServiceExternal(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		res := &addServiceExternalResult{
			Service: &services.AddServiceOKBodyExternal{
				ServiceID:      "1",
				ServiceName:    "ClickHouse Service",
				NodeID:         "1",
				Environment:    "environment",
				Cluster:        "clickhouse-cluster",
				ReplicationSet: "clickhouse-replication-set",
				CustomLabels:   map[string]string{"key": "value", "foo": "bar"},
				Group:          "ClickHouse",
			},
		}
		expected := `External Service added.
Service ID     : 1
Service name   : ClickHouse Service
Node ID        : 1
Environment    : environment
Cluster name   : clickhouse-cluster
Replication set: clickhouse-replication-set
Custom labels  : map[foo:bar key:value]
Group          : ClickHouse
`
		require.Equal(t, expected, res.String())
	})
}
