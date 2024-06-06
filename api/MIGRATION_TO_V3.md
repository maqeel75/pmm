## Migrations of API endpoints to make them more RESTful

| Current (v2)                                    | Migrate to (v3)                                | Comments                        |
| ----------------------------------------------- | ---------------------------------------------- | ------------------------------- |

**ServerService**                                   **ServerService**
GET /logz.zip                                       GET /v1/server/logs.zip                          ✅ nginx redirects /logs.zip to /v1/server/logs.zip
GET /v1/version                                     GET /v1/server/version                           ✅ nginx redirects /v1/version to /v1/server/version
GET /v1/readyz                                      GET /v1/server/readyz                            ✅ nginx redirects /v1/readyz to /v1/server/readyz
POST /v1/AWSInstanceCheck                           GET /v1/server/AWSInstance                       ✅
POST /v1/leaderHealthCheck                          GET /v1/server/leaderHealthCheck                 ✅
POST /v1/settings/Change                            PUT /v1/server/settings                          ✅
POST /v1/settings/Get                               GET /v1/server/settings                          ✅
POST /v1/updates/Check                              GET /v1/server/updates                           ✅
POST /v1/updates/Start                              POST /v1/server/updates:start                    ✅
POST /v1/updates/Status                             POST /v1/server/updates:getStatus                ✅ auth_token is passed in the body

**UserService**                                     **UserService**
GET /v1/user                                        GET /v1/users/me                                 ✅
PUT /v1/user                                        PUT /v1/users/me                                 ✅
POST /v1/user/list                                  GET /v1/users                                    ✅ 

**AgentsService**                                   **AgentsService**
POST /v1/inventory/Agents/Add                       POST /v1/inventory/agents                        ✅
POST /v1/inventory/Agents/Change                    PUT /v1/inventory/agents/{agent_id}              ✅
POST /v1/inventory/Agents/Get                       GET /v1/inventory/agents/{agent_id}              ✅
POST /v1/inventory/Agents/List                      GET /v1/inventory/agents                         ✅ Query param filters: service_id, node_id, 
                                                                                                        pmm_agent_id and agent_type
POST /v1/inventory/Agents/Remove                    DELETE /v1/inventory/agents/{agent_id}           ✅
POST /v1/inventory/Agents/GetLogs                   GET /v1/inventory/agents/{agent_id}/logs         ✅

**NodesService**                                   **NodesService**
POST /v1/inventory/Nodes/Add                        POST /v1/inventory/nodes                         ✅
POST /v1/inventory/Nodes/Get                        GET /v1/inventory/nodes/{node_id}                ✅
POST /v1/inventory/Nodes/Delete                     DELETE /v1/inventory/nodes/{node_id}             ✅
POST /v1/inventory/Nodes/List                       GET /v1/inventory/nodes                          ✅

**ServicesService**                                 **ServicesService**
POST /v1/inventory/Services/Add                     POST /v1/inventory/services                      ✅
POST /v1/inventory/Services/Change                  PUT /v1/inventory/services/{service_id}          ✅
POST /v1/inventory/Servicse/Get                     GET /v1/inventory/services/{service_id}          ✅
POST /v1/inventory/Services/List                    GET /v1/inventory/services                       ✅
POST /v1/inventory/Services/Remove                  DELETE /v1/inventory/services/{service_id}       ✅ pass ?force=true to remove a service with agents
POST /v1/inventory/Services/ListTypes               POST /v1/inventory/services:getTypes             ✅
POST /v1/inventory/Services/CustomLabels/Add        PUT /v1/inventory/services/{service_id}          ✅ NOTE: merged into PUT /v1/inventory/services/{id}
POST /v1/inventory/Services/CustomLabels/Remove     PUT /v1/inventory/services/{service_id}          ✅ NOTE: merged into PUT /v1/inventory/services/{id}

**ManagementService**                               **ManagementService**
POST /v1/management/Annotations/Add                 POST /v1/management/annotations                  ✅
POST /v1/management/Agent/List                      GET /v1/management/agents                        ✅ Moved from MgmtService
POST /v1/management/Node/Register                   POST /v1/management/nodes                        ✅
POST /v1/management/Node/Unregister                 DELETE /v1/management/nodes/{node_id}            ✅ ?force=true
POST /v1/management/Node/Get                        GET /v1/management/nodes/{node_id}               ✅ Moved from MgmtService
POST /v1/management/Node/List                       GET /v1/management/nodes                         ✅ Moved from MgmtService
POST /v1/management/Service/List                    GET /v1/management/services                      ✅ Moved from MgmtService
POST /v1/management/External/Add                    POST /v1/management/services                     ✅ NOTE: several endpoints merged into one
POST /v1/management/HAProxy/Add                     POST /v1/management/services                     ✅
POST /v1/management/MongoDB/Add                     POST /v1/management/services                     ✅
POST /v1/management/MySQL/Add                       POST /v1/management/services                     ✅
POST /v1/management/PostgreSQL/Add                  POST /v1/management/services                     ✅
POST /v1/management/ProxySQL/Add                    POST /v1/management/services                     ✅
POST /v1/management/RDS/Add                         POST /v1/management/services                     ✅
POST /v1/management/AzureDatabase/Add               POST /v1/management/services/azure               ✅ Moved from MgmtService
POST /v1/management/AzureDatabase/Discover          POST /v1/management/services:discoverAzure       ✅ Moved from MgmtService
POST /v1/management/RDS/Discover                    POST /v1/management/services:discoverRDS         ✅
POST /v1/management/Service/Remove                  DELETE /v1/management/services/{service_id}      ✅ NOTE: in addition, it accepts ?service_type=

**MgmtService**                                     **MgmtService**                                  NOTE: promoted to v1 from v1beta1
POST /v1/management/Agent/List                      GET /v1/management/agents                        ✅ Moved to ManagementService
POST /v1/management/Node/Get                        GET /v1/management/nodes/{node_id}               ✅ Moved to ManagementService
POST /v1/management/Node/List                       GET /v1/management/nodes                         ✅ Moved to ManagementService
POST /v1/management/AzureDatabase/Add               POST /v1/management/services                     ✅ Moved to ManagementService
POST /v1/management/AzureDatabase/Discover          POST /v1/management/services:discoverAzure       ✅ Moved to ManagementService
POST /v1/management/Service/List                    GET /v1/management/services                      ✅ Moved to ManagementService

**ActionsService**                                  **ActionService**
POST /v1/actions/Cancel                             POST /v1/actions:cancelAction                    ✅
POST /v1/actions/Get                                GET /v1/actions/{action_id}                      ✅
POST /v1/actions/StartMySQLExplain                  POST /v1/actions:startServiceAction              ✅ NOTE: several endpoints merged into one
POST /v1/actions/StartMySQLExplainJSON              POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartMySQLExplainTraditionalJSON   POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartMySQLShowIndex                POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartMySQLShowCreateTable          POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartMySQLShowTableStatus          POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartPostgreSQLShowCreateTable     POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartPostgreSQLShowIndex           POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartMongoDBExplain                POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartPTMongoDBSummary              POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartPTMySQLSummary                POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartPTPgSummary                   POST /v1/actions:startServiceAction              ✅
POST /v1/actions/StartPTSummary                     POST /v1/actions:startNodeAction                 ✅

**AlertingService**                                 **AlertingService**
POST /v1/alerting/Rules/Create                      POST /v1/alerting/rules                          ✅
POST /v1/alerting/Templates/Create                  POST /v1/alerting/templates                      ✅
POST /v1/alerting/Templates/Update                  PUT /v1/alerting/templates/{name}                ✅
POST /v1/alerting/Templates/List                    GET /v1/alerting/templates                       ✅
POST /v1/alerting/Templates/Delete                  DELETE /v1/alerting/templates/{name}             ✅

**AdvisorService**                                 **AdvisorService**
POST /v1/advisors/Change                            POST /v1/advisors/checks:batchChange             ✅
POST /v1/advisors/FailedChecks                      GET /v1/advisors/checks/failed                   ✅ ?service_id=1234&page_size=100&page_index=1
POST /v1/advisors/List                              GET /v1/advisors                                 ✅
POST /v1/advisors/ListChecks                        GET /v1/advisors/checks                          ✅
POST /v1/advisors/StartChecks                       POST /v1/advisors/checks:start                   ✅
POST /v1/advisors/ListFailedServices                GET /v1/advisors/failedServices                  ✅

**ArtifactsService**                                **ArtifactsService**                             TODO: merge to BackupService
POST /v1/backup/Artifacts/List                      GET /v1/backups/artifacts                        ✅
POST /v1/backup/Artifacts/Delete                    DELETE /v1/backups/artifacts/{artifact_id}       ✅ ?remove_files=true
POST /v1/backup/Artifacts/PITRTimeranges            GET /v1/backups/artifacts/{artifact_id}/pitr-timeranges ✅

**BackupsService**                                  **BackupService**                                NOTE: BackupsService renamed to BackupService
POST /v1/backup/Backups/ChangeScheduled             PUT /v1/backups:changeScheduled                  ✅
POST /v1/backup/Backups/GetLogs                     GET /v1/backups/{artifact_id}/logs               ✅
POST /v1/backup/Backups/ListArtifactCompatibleServices GET /v1/backups/{artifact_id}/compatible-services
POST /v1/backup/Backups/ListScheduled               GET /v1/backups/scheduled                        ✅
POST /v1/backup/Backups/RemoveScheduled             DELETE /v1/backups/scheduled/{scheduled_backup_id} ✅
POST /v1/backup/Backups/Restore                                                                      NOTE: Moved to RestoreService
POST /v1/backup/Backups/Schedule                    POST /v1/backups:schedule                        ✅
POST /v1/backup/Backups/Start                       POST /v1/backups:start                           ✅

**LocationsService**                                **LocationsService**
POST /v1/backup/Locations/Add                       POST /v1/backups/locations                       ✅
POST /v1/backup/Locations/Change                    PUT /v1/backups/locations/{location_id}          ✅
POST /v1/backup/Locations/List                      GET /v1/backups/locations                        ✅
POST /v1/backup/Locations/Remove                    DELETE /v1/backups/locations/{location_id}       ✅ ?force=true
POST /v1/backup/Locations/TestConfig                POST /v1/backups/locations:testConfig            ✅

**RestoreHistoryService**                           **RestoreService**
POST /v1/backup/RestoreHistory/List                 GET /v1/backups/restores                         ✅
POST /v1/backup/Backups/Restore                     POST /v1/backups/restores:start                  ✅
                                                    GET /v1/backups/restores/{restore_id}/logs       ✅ new, copied from /v1/backups/{artifact_id}/logs

**DumpsService**                                    **DumpService**                                  NOTE: renamed to DumpService
POST /v1/dump/List                                  GET /v1/dumps                                    ✅
POST /v1/dump/Delete                                POST /v1/dumps:batchDelete                       ✅ accepts an array in body
POST /v1/dump/GetLogs                               GET /v1/dumps/{dump_id}/logs                     ✅ ?offset=0&limit=100
POST /v1/dump/Start                                 POST /v1/dumps:start                             ✅              
POST /v1/dump/Upload                                POST /v1/dumps:upload                            ✅

**RoleService**                                     **AccessControlService**                         NOTE: renamed to AccessControlService
POST /v1/role/Assign                                POST /v1/accesscontrol/roles:assign              ✅
POST /v1/role/Create                                POST /v1/accesscontrol/roles                     ✅
POST /v1/role/Delete                                DELETE /v1/accesscontrol/roles/{role_id}         ✅ ?replacement_role_id=abcdedf-123456
POST /v1/role/Get                                   GET /v1/accesscontrol/roles/{role_id}            ✅
POST /v1/role/List                                  GET /v1/accesscontrol/roles                      ✅
POST /v1/role/SetDefault                            POST /v1/accesscontrol/roles:setDefault          ✅
POST /v1/role/Update                                PUT /v1/accesscontrol/roles/{role_id}            ✅

**QANService**                                      **QANService**
POST /v0/qan/Filters/Get                            POST /v1/qan/metrics:getFilters                  ✅
POST /v0/qan/GetMetricsNames                        POST /v1/qan/metrics:getNames                    ✅
POST /v0/qan/GetReport                              POST /v1/qan/metrics:getReport                   ✅
POST /v0/qan/ObjectDetails/ExplainFingerprintByQueryId POST /v1/qan:explainFingerprint               ✅
POST /v0/qan/ObjectDetails/GetHistogram             POST /v1/qan:getHistogram                        ✅
POST /v0/qan/ObjectDetails/GetLables                POST /v1/qan:getLabels                           ✅
POST /v0/qan/ObjectDetails/GetMetrics               POST /v1/qan:getMetrics                          ✅
POST /v0/qan/ObjectDetails/GetQueryPlan             GET /v1/qan/query/{queryid}/plan                 ✅
POST /v0/qan/ObjectDetails/QueryExists              POST /v1/qan/query:exists                        ✅ 
POST /v0/qan/ObjectDetails/GetQueryExample          POST /v1/qan/query:getExample                    ✅
POST /v0/qan/ObjectDetails/SchemaByQueryId          POST /v1/qan/query:getSchema                     ✅

**PlatformService**                                 **PlatformService**
POST /v1/platform/Connect                           POST /v1/platform:connect                        ✅
POST /v1/platform/Disconnect                        POST /v1/platform:disconnect                     ✅
POST /v1/platform/GetContactInformation             GET /v1/platform/contact                         ✅
POST /v1/platform/SearchOganizationEntitlemenets    GET /v1/platform/organization/entitlements       ✅
POST /v1/platform/SearchOganizationTickets          GET /v1/platform/organization/tickets            ✅
POST /v1/platform/ServerInfo                        GET /v1/platform/server                          ✅
POST /v1/platform/UserInfo                          GET /v1/platform/user                            ✅

// TODO: rename `period_start_from` to `start_from` and `period_start_to` to `start_to`