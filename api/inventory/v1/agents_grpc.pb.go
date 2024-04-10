// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: inventory/v1/agents.proto

package inventoryv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AgentsService_ListAgents_FullMethodName   = "/inventory.v1.AgentsService/ListAgents"
	AgentsService_GetAgent_FullMethodName     = "/inventory.v1.AgentsService/GetAgent"
	AgentsService_GetAgentLogs_FullMethodName = "/inventory.v1.AgentsService/GetAgentLogs"
	AgentsService_AddAgent_FullMethodName     = "/inventory.v1.AgentsService/AddAgent"
	AgentsService_ChangeAgent_FullMethodName  = "/inventory.v1.AgentsService/ChangeAgent"
	AgentsService_RemoveAgent_FullMethodName  = "/inventory.v1.AgentsService/RemoveAgent"
)

// AgentsServiceClient is the client API for AgentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentsServiceClient interface {
	// ListAgents returns a list of all Agents.
	ListAgents(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error)
	// GetAgent returns a single Agent by ID.
	GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentResponse, error)
	// GetAgentLogs returns Agent logs by ID.
	GetAgentLogs(ctx context.Context, in *GetAgentLogsRequest, opts ...grpc.CallOption) (*GetAgentLogsResponse, error)
	// AddAgent adds an Agent to Inventory.
	AddAgent(ctx context.Context, in *AddAgentRequest, opts ...grpc.CallOption) (*AddAgentResponse, error)
	// ChangeAgent changes a subset of attributes of the Agent record in Inventory.
	ChangeAgent(ctx context.Context, in *ChangeAgentRequest, opts ...grpc.CallOption) (*ChangeAgentResponse, error)
	// RemoveAgent removes an Agent.
	RemoveAgent(ctx context.Context, in *RemoveAgentRequest, opts ...grpc.CallOption) (*RemoveAgentResponse, error)
}

type agentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentsServiceClient(cc grpc.ClientConnInterface) AgentsServiceClient {
	return &agentsServiceClient{cc}
}

func (c *agentsServiceClient) ListAgents(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error) {
	out := new(ListAgentsResponse)
	err := c.cc.Invoke(ctx, AgentsService_ListAgents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsServiceClient) GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentResponse, error) {
	out := new(GetAgentResponse)
	err := c.cc.Invoke(ctx, AgentsService_GetAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsServiceClient) GetAgentLogs(ctx context.Context, in *GetAgentLogsRequest, opts ...grpc.CallOption) (*GetAgentLogsResponse, error) {
	out := new(GetAgentLogsResponse)
	err := c.cc.Invoke(ctx, AgentsService_GetAgentLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsServiceClient) AddAgent(ctx context.Context, in *AddAgentRequest, opts ...grpc.CallOption) (*AddAgentResponse, error) {
	out := new(AddAgentResponse)
	err := c.cc.Invoke(ctx, AgentsService_AddAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsServiceClient) ChangeAgent(ctx context.Context, in *ChangeAgentRequest, opts ...grpc.CallOption) (*ChangeAgentResponse, error) {
	out := new(ChangeAgentResponse)
	err := c.cc.Invoke(ctx, AgentsService_ChangeAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsServiceClient) RemoveAgent(ctx context.Context, in *RemoveAgentRequest, opts ...grpc.CallOption) (*RemoveAgentResponse, error) {
	out := new(RemoveAgentResponse)
	err := c.cc.Invoke(ctx, AgentsService_RemoveAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentsServiceServer is the server API for AgentsService service.
// All implementations must embed UnimplementedAgentsServiceServer
// for forward compatibility
type AgentsServiceServer interface {
	// ListAgents returns a list of all Agents.
	ListAgents(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error)
	// GetAgent returns a single Agent by ID.
	GetAgent(context.Context, *GetAgentRequest) (*GetAgentResponse, error)
	// GetAgentLogs returns Agent logs by ID.
	GetAgentLogs(context.Context, *GetAgentLogsRequest) (*GetAgentLogsResponse, error)
	// AddAgent adds an Agent to Inventory.
	AddAgent(context.Context, *AddAgentRequest) (*AddAgentResponse, error)
	// ChangeAgent changes a subset of attributes of the Agent record in Inventory.
	ChangeAgent(context.Context, *ChangeAgentRequest) (*ChangeAgentResponse, error)
	// RemoveAgent removes an Agent.
	RemoveAgent(context.Context, *RemoveAgentRequest) (*RemoveAgentResponse, error)
	mustEmbedUnimplementedAgentsServiceServer()
}

// UnimplementedAgentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentsServiceServer struct{}

func (UnimplementedAgentsServiceServer) ListAgents(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgents not implemented")
}

func (UnimplementedAgentsServiceServer) GetAgent(context.Context, *GetAgentRequest) (*GetAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgent not implemented")
}

func (UnimplementedAgentsServiceServer) GetAgentLogs(context.Context, *GetAgentLogsRequest) (*GetAgentLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentLogs not implemented")
}

func (UnimplementedAgentsServiceServer) AddAgent(context.Context, *AddAgentRequest) (*AddAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAgent not implemented")
}

func (UnimplementedAgentsServiceServer) ChangeAgent(context.Context, *ChangeAgentRequest) (*ChangeAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAgent not implemented")
}

func (UnimplementedAgentsServiceServer) RemoveAgent(context.Context, *RemoveAgentRequest) (*RemoveAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAgent not implemented")
}
func (UnimplementedAgentsServiceServer) mustEmbedUnimplementedAgentsServiceServer() {}

// UnsafeAgentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentsServiceServer will
// result in compilation errors.
type UnsafeAgentsServiceServer interface {
	mustEmbedUnimplementedAgentsServiceServer()
}

func RegisterAgentsServiceServer(s grpc.ServiceRegistrar, srv AgentsServiceServer) {
	s.RegisterService(&AgentsService_ServiceDesc, srv)
}

func _AgentsService_ListAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServiceServer).ListAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentsService_ListAgents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServiceServer).ListAgents(ctx, req.(*ListAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentsService_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServiceServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentsService_GetAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServiceServer).GetAgent(ctx, req.(*GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentsService_GetAgentLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServiceServer).GetAgentLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentsService_GetAgentLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServiceServer).GetAgentLogs(ctx, req.(*GetAgentLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentsService_AddAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServiceServer).AddAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentsService_AddAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServiceServer).AddAgent(ctx, req.(*AddAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentsService_ChangeAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServiceServer).ChangeAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentsService_ChangeAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServiceServer).ChangeAgent(ctx, req.(*ChangeAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentsService_RemoveAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServiceServer).RemoveAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentsService_RemoveAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServiceServer).RemoveAgent(ctx, req.(*RemoveAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentsService_ServiceDesc is the grpc.ServiceDesc for AgentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.v1.AgentsService",
	HandlerType: (*AgentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAgents",
			Handler:    _AgentsService_ListAgents_Handler,
		},
		{
			MethodName: "GetAgent",
			Handler:    _AgentsService_GetAgent_Handler,
		},
		{
			MethodName: "GetAgentLogs",
			Handler:    _AgentsService_GetAgentLogs_Handler,
		},
		{
			MethodName: "AddAgent",
			Handler:    _AgentsService_AddAgent_Handler,
		},
		{
			MethodName: "ChangeAgent",
			Handler:    _AgentsService_ChangeAgent_Handler,
		},
		{
			MethodName: "RemoveAgent",
			Handler:    _AgentsService_RemoveAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/v1/agents.proto",
}