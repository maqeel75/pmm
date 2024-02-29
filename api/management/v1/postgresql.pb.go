// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: management/v1/postgresql.proto

package managementv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1 "github.com/percona/pmm/api/inventory/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddPostgreSQLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Create a new Node with those parameters.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	AddNode *AddNodeParams `protobuf:"bytes,3,opt,name=add_node,json=addNode,proto3" json:"add_node,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Node and Service access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Service Access port.
	// Port is required when the address present.
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	// Database name.
	Database string `protobuf:"bytes,7,opt,name=database,proto3" json:"database,omitempty"`
	// Service Access socket.
	// Address (and port) or socket is required.
	Socket string `protobuf:"bytes,8,opt,name=socket,proto3" json:"socket,omitempty"`
	// The "pmm-agent" identifier which should run agents. Required.
	PmmAgentId string `protobuf:"bytes,9,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,10,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,11,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,12,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// PostgreSQL username for scraping metrics.
	Username string `protobuf:"bytes,13,opt,name=username,proto3" json:"username,omitempty"`
	// PostgreSQL password for scraping metrics.
	Password string `protobuf:"bytes,14,opt,name=password,proto3" json:"password,omitempty"`
	// If true, adds qan-postgresql-pgstatements-agent for provided service.
	QanPostgresqlPgstatementsAgent bool `protobuf:"varint,15,opt,name=qan_postgresql_pgstatements_agent,json=qanPostgresqlPgstatementsAgent,proto3" json:"qan_postgresql_pgstatements_agent,omitempty"`
	// If true, adds qan-postgresql-pgstatmonitor-agent for provided service.
	QanPostgresqlPgstatmonitorAgent bool `protobuf:"varint,16,opt,name=qan_postgresql_pgstatmonitor_agent,json=qanPostgresqlPgstatmonitorAgent,proto3" json:"qan_postgresql_pgstatmonitor_agent,omitempty"`
	// Limit query length in QAN (default: server-defined; -1: no limit).
	MaxQueryLength int32 `protobuf:"varint,17,opt,name=max_query_length,json=maxQueryLength,proto3" json:"max_query_length,omitempty"`
	// Disable query examples.
	DisableQueryExamples bool `protobuf:"varint,18,opt,name=disable_query_examples,json=disableQueryExamples,proto3" json:"disable_query_examples,omitempty"`
	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `protobuf:"bytes,19,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Skip connection check.
	SkipConnectionCheck bool `protobuf:"varint,20,opt,name=skip_connection_check,json=skipConnectionCheck,proto3" json:"skip_connection_check,omitempty"`
	// Disable parsing comments from queries and showing them in QAN.
	DisableCommentsParsing bool `protobuf:"varint,21,opt,name=disable_comments_parsing,json=disableCommentsParsing,proto3" json:"disable_comments_parsing,omitempty"`
	// Use TLS for database connections.
	Tls bool `protobuf:"varint,22,opt,name=tls,proto3" json:"tls,omitempty"`
	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TlsSkipVerify bool `protobuf:"varint,23,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
	// Defines metrics flow model for this exporter.
	// Metrics could be pushed to the server with vmagent,
	// pulled by the server, or the server could choose behavior automatically.
	MetricsMode MetricsMode `protobuf:"varint,24,opt,name=metrics_mode,json=metricsMode,proto3,enum=management.v1.MetricsMode" json:"metrics_mode,omitempty"`
	// List of collector names to disable in this exporter.
	DisableCollectors []string `protobuf:"bytes,25,rep,name=disable_collectors,json=disableCollectors,proto3" json:"disable_collectors,omitempty"`
	// TLS CA certificate.
	TlsCa string `protobuf:"bytes,26,opt,name=tls_ca,json=tlsCa,proto3" json:"tls_ca,omitempty"`
	// TLS Certifcate.
	TlsCert string `protobuf:"bytes,27,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	// TLS Certificate Key.
	TlsKey string `protobuf:"bytes,28,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
	// Custom password for exporter endpoint /metrics.
	AgentPassword string `protobuf:"bytes,29,opt,name=agent_password,json=agentPassword,proto3" json:"agent_password,omitempty"`
	// Exporter log level
	LogLevel v1.LogLevel `protobuf:"varint,30,opt,name=log_level,json=logLevel,proto3,enum=inventory.v1.LogLevel" json:"log_level,omitempty"`
	// Limit for auto discovery.
	AutoDiscoveryLimit int32 `protobuf:"varint,31,opt,name=auto_discovery_limit,json=autoDiscoveryLimit,proto3" json:"auto_discovery_limit,omitempty"`
	// Optionally expose the exporter process on all public interfaces
	ExposeExporter bool `protobuf:"varint,32,opt,name=expose_exporter,json=exposeExporter,proto3" json:"expose_exporter,omitempty"`
}

func (x *AddPostgreSQLRequest) Reset() {
	*x = AddPostgreSQLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_v1_postgresql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostgreSQLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostgreSQLRequest) ProtoMessage() {}

func (x *AddPostgreSQLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_postgresql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostgreSQLRequest.ProtoReflect.Descriptor instead.
func (*AddPostgreSQLRequest) Descriptor() ([]byte, []int) {
	return file_management_v1_postgresql_proto_rawDescGZIP(), []int{0}
}

func (x *AddPostgreSQLRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetAddNode() *AddNodeParams {
	if x != nil {
		return x.AddNode
	}
	return nil
}

func (x *AddPostgreSQLRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AddPostgreSQLRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetSocket() string {
	if x != nil {
		return x.Socket
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetPmmAgentId() string {
	if x != nil {
		return x.PmmAgentId
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetReplicationSet() string {
	if x != nil {
		return x.ReplicationSet
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetQanPostgresqlPgstatementsAgent() bool {
	if x != nil {
		return x.QanPostgresqlPgstatementsAgent
	}
	return false
}

func (x *AddPostgreSQLRequest) GetQanPostgresqlPgstatmonitorAgent() bool {
	if x != nil {
		return x.QanPostgresqlPgstatmonitorAgent
	}
	return false
}

func (x *AddPostgreSQLRequest) GetMaxQueryLength() int32 {
	if x != nil {
		return x.MaxQueryLength
	}
	return 0
}

func (x *AddPostgreSQLRequest) GetDisableQueryExamples() bool {
	if x != nil {
		return x.DisableQueryExamples
	}
	return false
}

func (x *AddPostgreSQLRequest) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *AddPostgreSQLRequest) GetSkipConnectionCheck() bool {
	if x != nil {
		return x.SkipConnectionCheck
	}
	return false
}

func (x *AddPostgreSQLRequest) GetDisableCommentsParsing() bool {
	if x != nil {
		return x.DisableCommentsParsing
	}
	return false
}

func (x *AddPostgreSQLRequest) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *AddPostgreSQLRequest) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

func (x *AddPostgreSQLRequest) GetMetricsMode() MetricsMode {
	if x != nil {
		return x.MetricsMode
	}
	return MetricsMode_METRICS_MODE_UNSPECIFIED
}

func (x *AddPostgreSQLRequest) GetDisableCollectors() []string {
	if x != nil {
		return x.DisableCollectors
	}
	return nil
}

func (x *AddPostgreSQLRequest) GetTlsCa() string {
	if x != nil {
		return x.TlsCa
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetTlsCert() string {
	if x != nil {
		return x.TlsCert
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetTlsKey() string {
	if x != nil {
		return x.TlsKey
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetAgentPassword() string {
	if x != nil {
		return x.AgentPassword
	}
	return ""
}

func (x *AddPostgreSQLRequest) GetLogLevel() v1.LogLevel {
	if x != nil {
		return x.LogLevel
	}
	return v1.LogLevel(0)
}

func (x *AddPostgreSQLRequest) GetAutoDiscoveryLimit() int32 {
	if x != nil {
		return x.AutoDiscoveryLimit
	}
	return 0
}

func (x *AddPostgreSQLRequest) GetExposeExporter() bool {
	if x != nil {
		return x.ExposeExporter
	}
	return false
}

type AddPostgreSQLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service                         *v1.PostgreSQLService               `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	PostgresExporter                *v1.PostgresExporter                `protobuf:"bytes,2,opt,name=postgres_exporter,json=postgresExporter,proto3" json:"postgres_exporter,omitempty"`
	QanPostgresqlPgstatementsAgent  *v1.QANPostgreSQLPgStatementsAgent  `protobuf:"bytes,3,opt,name=qan_postgresql_pgstatements_agent,json=qanPostgresqlPgstatementsAgent,proto3" json:"qan_postgresql_pgstatements_agent,omitempty"`
	QanPostgresqlPgstatmonitorAgent *v1.QANPostgreSQLPgStatMonitorAgent `protobuf:"bytes,4,opt,name=qan_postgresql_pgstatmonitor_agent,json=qanPostgresqlPgstatmonitorAgent,proto3" json:"qan_postgresql_pgstatmonitor_agent,omitempty"`
	// Warning message.
	Warning string `protobuf:"bytes,5,opt,name=warning,proto3" json:"warning,omitempty"`
}

func (x *AddPostgreSQLResponse) Reset() {
	*x = AddPostgreSQLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_v1_postgresql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostgreSQLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostgreSQLResponse) ProtoMessage() {}

func (x *AddPostgreSQLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_postgresql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostgreSQLResponse.ProtoReflect.Descriptor instead.
func (*AddPostgreSQLResponse) Descriptor() ([]byte, []int) {
	return file_management_v1_postgresql_proto_rawDescGZIP(), []int{1}
}

func (x *AddPostgreSQLResponse) GetService() *v1.PostgreSQLService {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *AddPostgreSQLResponse) GetPostgresExporter() *v1.PostgresExporter {
	if x != nil {
		return x.PostgresExporter
	}
	return nil
}

func (x *AddPostgreSQLResponse) GetQanPostgresqlPgstatementsAgent() *v1.QANPostgreSQLPgStatementsAgent {
	if x != nil {
		return x.QanPostgresqlPgstatementsAgent
	}
	return nil
}

func (x *AddPostgreSQLResponse) GetQanPostgresqlPgstatmonitorAgent() *v1.QANPostgreSQLPgStatMonitorAgent {
	if x != nil {
		return x.QanPostgresqlPgstatmonitorAgent
	}
	return nil
}

func (x *AddPostgreSQLResponse) GetWarning() string {
	if x != nil {
		return x.Warning
	}
	return ""
}

var File_management_v1_postgresql_proto protoreflect.FileDescriptor

var file_management_v1_postgresql_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x0b, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x50, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2a,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x0c,
	0x70, 0x6d, 0x6d, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x6d, 0x6d,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x49, 0x0a,
	0x21, 0x71, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x5f,
	0x70, 0x67, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x71, 0x61, 0x6e, 0x50, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x50, 0x67, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x22, 0x71, 0x61, 0x6e, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x5f, 0x70, 0x67, 0x73, 0x74, 0x61,
	0x74, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x71, 0x61, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x71, 0x6c, 0x50, 0x67, 0x73, 0x74, 0x61, 0x74, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x6d, 0x61, 0x78, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x34, 0x0a, 0x16, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e,
	0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x6c,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x53,
	0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x3d, 0x0a, 0x0c, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x19,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6c, 0x73, 0x5f, 0x63,
	0x61, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6c, 0x73, 0x43, 0x61, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6c, 0x73,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6c, 0x73, 0x4b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x6c, 0x6f, 0x67,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x30,
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x61, 0x75,
	0x74, 0x6f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x6f, 0x73,
	0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xae, 0x03, 0x0a, 0x15, 0x41,
	0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4b, 0x0a, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x10, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x77, 0x0a, 0x21,
	0x71, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x5f, 0x70,
	0x67, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x41, 0x4e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x53, 0x51, 0x4c, 0x50, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x1e, 0x71, 0x61, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x71, 0x6c, 0x50, 0x67, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x7a, 0x0a, 0x22, 0x71, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x5f, 0x70, 0x67, 0x73, 0x74, 0x61, 0x74, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x41, 0x4e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x50, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x1f, 0x71, 0x61, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x50,
	0x67, 0x73, 0x74, 0x61, 0x74, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x32, 0x8e, 0x03, 0x0a, 0x11,
	0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xf8, 0x02, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x53, 0x51, 0x4c, 0x12, 0x23, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x53, 0x51,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9b,
	0x02, 0x92, 0x41, 0xef, 0x01, 0x12, 0x0e, 0x41, 0x64, 0x64, 0x20, 0x50, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x53, 0x51, 0x4c, 0x1a, 0xdc, 0x01, 0x41, 0x64, 0x64, 0x73, 0x20, 0x50, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x20, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x20, 0x49, 0x74,
	0x20, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x20, 0x61,
	0x64, 0x64, 0x73, 0x20, 0x61, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x74, 0x6f,
	0x20, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63,
	0x68, 0x20, 0x69, 0x73, 0x20, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x6e, 0x20,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x22, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x22, 0x2c, 0x20, 0x74, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x64, 0x64, 0x73, 0x20, 0x22, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x22, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20,
	0x22, 0x70, 0x6d, 0x6d, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x50, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x53, 0x51, 0x4c, 0x2f, 0x41, 0x64, 0x64, 0x42, 0xb0, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d,
	0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x19, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_v1_postgresql_proto_rawDescOnce sync.Once
	file_management_v1_postgresql_proto_rawDescData = file_management_v1_postgresql_proto_rawDesc
)

func file_management_v1_postgresql_proto_rawDescGZIP() []byte {
	file_management_v1_postgresql_proto_rawDescOnce.Do(func() {
		file_management_v1_postgresql_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_v1_postgresql_proto_rawDescData)
	})
	return file_management_v1_postgresql_proto_rawDescData
}

var (
	file_management_v1_postgresql_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_management_v1_postgresql_proto_goTypes  = []interface{}{
		(*AddPostgreSQLRequest)(nil),               // 0: management.v1.AddPostgreSQLRequest
		(*AddPostgreSQLResponse)(nil),              // 1: management.v1.AddPostgreSQLResponse
		nil,                                        // 2: management.v1.AddPostgreSQLRequest.CustomLabelsEntry
		(*AddNodeParams)(nil),                      // 3: management.v1.AddNodeParams
		(MetricsMode)(0),                           // 4: management.v1.MetricsMode
		(v1.LogLevel)(0),                           // 5: inventory.v1.LogLevel
		(*v1.PostgreSQLService)(nil),               // 6: inventory.v1.PostgreSQLService
		(*v1.PostgresExporter)(nil),                // 7: inventory.v1.PostgresExporter
		(*v1.QANPostgreSQLPgStatementsAgent)(nil),  // 8: inventory.v1.QANPostgreSQLPgStatementsAgent
		(*v1.QANPostgreSQLPgStatMonitorAgent)(nil), // 9: inventory.v1.QANPostgreSQLPgStatMonitorAgent
	}
)

var file_management_v1_postgresql_proto_depIdxs = []int32{
	3, // 0: management.v1.AddPostgreSQLRequest.add_node:type_name -> management.v1.AddNodeParams
	2, // 1: management.v1.AddPostgreSQLRequest.custom_labels:type_name -> management.v1.AddPostgreSQLRequest.CustomLabelsEntry
	4, // 2: management.v1.AddPostgreSQLRequest.metrics_mode:type_name -> management.v1.MetricsMode
	5, // 3: management.v1.AddPostgreSQLRequest.log_level:type_name -> inventory.v1.LogLevel
	6, // 4: management.v1.AddPostgreSQLResponse.service:type_name -> inventory.v1.PostgreSQLService
	7, // 5: management.v1.AddPostgreSQLResponse.postgres_exporter:type_name -> inventory.v1.PostgresExporter
	8, // 6: management.v1.AddPostgreSQLResponse.qan_postgresql_pgstatements_agent:type_name -> inventory.v1.QANPostgreSQLPgStatementsAgent
	9, // 7: management.v1.AddPostgreSQLResponse.qan_postgresql_pgstatmonitor_agent:type_name -> inventory.v1.QANPostgreSQLPgStatMonitorAgent
	0, // 8: management.v1.PostgreSQLService.AddPostgreSQL:input_type -> management.v1.AddPostgreSQLRequest
	1, // 9: management.v1.PostgreSQLService.AddPostgreSQL:output_type -> management.v1.AddPostgreSQLResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_management_v1_postgresql_proto_init() }
func file_management_v1_postgresql_proto_init() {
	if File_management_v1_postgresql_proto != nil {
		return
	}
	file_management_v1_metrics_proto_init()
	file_management_v1_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_management_v1_postgresql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostgreSQLRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_management_v1_postgresql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostgreSQLResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_management_v1_postgresql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_management_v1_postgresql_proto_goTypes,
		DependencyIndexes: file_management_v1_postgresql_proto_depIdxs,
		MessageInfos:      file_management_v1_postgresql_proto_msgTypes,
	}.Build()
	File_management_v1_postgresql_proto = out.File
	file_management_v1_postgresql_proto_rawDesc = nil
	file_management_v1_postgresql_proto_goTypes = nil
	file_management_v1_postgresql_proto_depIdxs = nil
}
