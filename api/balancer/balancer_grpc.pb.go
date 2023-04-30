// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/proto/balancer.proto

package balancer

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

// BalancerClient is the client API for Balancer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalancerClient interface {
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	Set(ctx context.Context, in *BalancerSetRequest, opts ...grpc.CallOption) (*BalancerSetResponse, error)
	SetToIndex(ctx context.Context, in *BalancerSetToIndexRequest, opts ...grpc.CallOption) (*BalancerSetResponse, error)
	Get(ctx context.Context, in *BalancerGetRequest, opts ...grpc.CallOption) (*BalancerGetResponse, error)
	GetFromIndex(ctx context.Context, in *BalancerGetFromIndexRequest, opts ...grpc.CallOption) (*BalancerGetResponse, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	Servers(ctx context.Context, in *ServersRequest, opts ...grpc.CallOption) (*ServersResponse, error)
}

type balancerClient struct {
	cc grpc.ClientConnInterface
}

func NewBalancerClient(cc grpc.ClientConnInterface) BalancerClient {
	return &balancerClient{cc}
}

func (c *balancerClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/api.Balancer/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerClient) Set(ctx context.Context, in *BalancerSetRequest, opts ...grpc.CallOption) (*BalancerSetResponse, error) {
	out := new(BalancerSetResponse)
	err := c.cc.Invoke(ctx, "/api.Balancer/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerClient) SetToIndex(ctx context.Context, in *BalancerSetToIndexRequest, opts ...grpc.CallOption) (*BalancerSetResponse, error) {
	out := new(BalancerSetResponse)
	err := c.cc.Invoke(ctx, "/api.Balancer/SetToIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerClient) Get(ctx context.Context, in *BalancerGetRequest, opts ...grpc.CallOption) (*BalancerGetResponse, error) {
	out := new(BalancerGetResponse)
	err := c.cc.Invoke(ctx, "/api.Balancer/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerClient) GetFromIndex(ctx context.Context, in *BalancerGetFromIndexRequest, opts ...grpc.CallOption) (*BalancerGetResponse, error) {
	out := new(BalancerGetResponse)
	err := c.cc.Invoke(ctx, "/api.Balancer/GetFromIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/api.Balancer/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/api.Balancer/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerClient) Servers(ctx context.Context, in *ServersRequest, opts ...grpc.CallOption) (*ServersResponse, error) {
	out := new(ServersResponse)
	err := c.cc.Invoke(ctx, "/api.Balancer/Servers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalancerServer is the server API for Balancer service.
// All implementations must embed UnimplementedBalancerServer
// for forward compatibility
type BalancerServer interface {
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
	Set(context.Context, *BalancerSetRequest) (*BalancerSetResponse, error)
	SetToIndex(context.Context, *BalancerSetToIndexRequest) (*BalancerSetResponse, error)
	Get(context.Context, *BalancerGetRequest) (*BalancerGetResponse, error)
	GetFromIndex(context.Context, *BalancerGetFromIndexRequest) (*BalancerGetResponse, error)
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	Servers(context.Context, *ServersRequest) (*ServersResponse, error)
	mustEmbedUnimplementedBalancerServer()
}

// UnimplementedBalancerServer must be embedded to have forward compatible implementations.
type UnimplementedBalancerServer struct {
}

func (UnimplementedBalancerServer) Index(context.Context, *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedBalancerServer) Set(context.Context, *BalancerSetRequest) (*BalancerSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedBalancerServer) SetToIndex(context.Context, *BalancerSetToIndexRequest) (*BalancerSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetToIndex not implemented")
}
func (UnimplementedBalancerServer) Get(context.Context, *BalancerGetRequest) (*BalancerGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBalancerServer) GetFromIndex(context.Context, *BalancerGetFromIndexRequest) (*BalancerGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFromIndex not implemented")
}
func (UnimplementedBalancerServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedBalancerServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedBalancerServer) Servers(context.Context, *ServersRequest) (*ServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Servers not implemented")
}
func (UnimplementedBalancerServer) mustEmbedUnimplementedBalancerServer() {}

// UnsafeBalancerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalancerServer will
// result in compilation errors.
type UnsafeBalancerServer interface {
	mustEmbedUnimplementedBalancerServer()
}

func RegisterBalancerServer(s grpc.ServiceRegistrar, srv BalancerServer) {
	s.RegisterService(&Balancer_ServiceDesc, srv)
}

func _Balancer_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balancer/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balancer_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalancerSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balancer/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).Set(ctx, req.(*BalancerSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balancer_SetToIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalancerSetToIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).SetToIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balancer/SetToIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).SetToIndex(ctx, req.(*BalancerSetToIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balancer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalancerGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balancer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).Get(ctx, req.(*BalancerGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balancer_GetFromIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalancerGetFromIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).GetFromIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balancer/GetFromIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).GetFromIndex(ctx, req.(*BalancerGetFromIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balancer_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balancer/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balancer_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balancer/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balancer_Servers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServer).Servers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balancer/Servers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServer).Servers(ctx, req.(*ServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Balancer_ServiceDesc is the grpc.ServiceDesc for Balancer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Balancer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Balancer",
	HandlerType: (*BalancerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _Balancer_Index_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Balancer_Set_Handler,
		},
		{
			MethodName: "SetToIndex",
			Handler:    _Balancer_SetToIndex_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Balancer_Get_Handler,
		},
		{
			MethodName: "GetFromIndex",
			Handler:    _Balancer_GetFromIndex_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Balancer_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _Balancer_Disconnect_Handler,
		},
		{
			MethodName: "Servers",
			Handler:    _Balancer_Servers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/balancer.proto",
}
