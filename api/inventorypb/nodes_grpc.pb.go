// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: inventorypb/nodes.proto

package inventorypb

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
	Nodes_ListNodes_FullMethodName  = "/inventory.Nodes/ListNodes"
	Nodes_GetNode_FullMethodName    = "/inventory.Nodes/GetNode"
	Nodes_AddNode_FullMethodName    = "/inventory.Nodes/AddNode"
	Nodes_RemoveNode_FullMethodName = "/inventory.Nodes/RemoveNode"
)

// NodesClient is the client API for Nodes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodesClient interface {
	// ListNodes returns a list of all Nodes.
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	// GetNode returns a single Node by ID.
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	// AddNode adds any type of Node.
	AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*AddNodeResponse, error)
	// RemoveNode removes a Node.
	RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...grpc.CallOption) (*RemoveNodeResponse, error)
}

type nodesClient struct {
	cc grpc.ClientConnInterface
}

func NewNodesClient(cc grpc.ClientConnInterface) NodesClient {
	return &nodesClient{cc}
}

func (c *nodesClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := c.cc.Invoke(ctx, Nodes_ListNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, Nodes_GetNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesClient) AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*AddNodeResponse, error) {
	out := new(AddNodeResponse)
	err := c.cc.Invoke(ctx, Nodes_AddNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesClient) RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...grpc.CallOption) (*RemoveNodeResponse, error) {
	out := new(RemoveNodeResponse)
	err := c.cc.Invoke(ctx, Nodes_RemoveNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodesServer is the server API for Nodes service.
// All implementations must embed UnimplementedNodesServer
// for forward compatibility
type NodesServer interface {
	// ListNodes returns a list of all Nodes.
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	// GetNode returns a single Node by ID.
	GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	// AddNode adds any type of Node.
	AddNode(context.Context, *AddNodeRequest) (*AddNodeResponse, error)
	// RemoveNode removes a Node.
	RemoveNode(context.Context, *RemoveNodeRequest) (*RemoveNodeResponse, error)
	mustEmbedUnimplementedNodesServer()
}

// UnimplementedNodesServer must be embedded to have forward compatible implementations.
type UnimplementedNodesServer struct{}

func (UnimplementedNodesServer) ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}

func (UnimplementedNodesServer) GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}

func (UnimplementedNodesServer) AddNode(context.Context, *AddNodeRequest) (*AddNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}

func (UnimplementedNodesServer) RemoveNode(context.Context, *RemoveNodeRequest) (*RemoveNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNode not implemented")
}
func (UnimplementedNodesServer) mustEmbedUnimplementedNodesServer() {}

// UnsafeNodesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodesServer will
// result in compilation errors.
type UnsafeNodesServer interface {
	mustEmbedUnimplementedNodesServer()
}

func RegisterNodesServer(s grpc.ServiceRegistrar, srv NodesServer) {
	s.RegisterService(&Nodes_ServiceDesc, srv)
}

func _Nodes_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nodes_ListNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nodes_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nodes_GetNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nodes_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nodes_AddNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).AddNode(ctx, req.(*AddNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nodes_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).RemoveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nodes_RemoveNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).RemoveNode(ctx, req.(*RemoveNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nodes_ServiceDesc is the grpc.ServiceDesc for Nodes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nodes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Nodes",
	HandlerType: (*NodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNodes",
			Handler:    _Nodes_ListNodes_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _Nodes_GetNode_Handler,
		},
		{
			MethodName: "AddNode",
			Handler:    _Nodes_AddNode_Handler,
		},
		{
			MethodName: "RemoveNode",
			Handler:    _Nodes_RemoveNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventorypb/nodes.proto",
}
