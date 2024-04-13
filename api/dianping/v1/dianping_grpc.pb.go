// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.3
// source: dianping/v1/dianping.proto

package v1

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
	Dianping_Prepare_FullMethodName     = "/api.dianping.v1.Dianping/prepare"
	Dianping_ScanPrepare_FullMethodName = "/api.dianping.v1.Dianping/scanPrepare"
	Dianping_AccessToken_FullMethodName = "/api.dianping.v1.Dianping/accessToken"
)

// DianpingClient is the client API for Dianping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DianpingClient interface {
	Prepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*DataReply, error)
	ScanPrepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*DataReply, error)
	AccessToken(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*DataReply, error)
}

type dianpingClient struct {
	cc grpc.ClientConnInterface
}

func NewDianpingClient(cc grpc.ClientConnInterface) DianpingClient {
	return &dianpingClient{cc}
}

func (c *dianpingClient) Prepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*DataReply, error) {
	out := new(DataReply)
	err := c.cc.Invoke(ctx, Dianping_Prepare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dianpingClient) ScanPrepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*DataReply, error) {
	out := new(DataReply)
	err := c.cc.Invoke(ctx, Dianping_ScanPrepare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dianpingClient) AccessToken(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*DataReply, error) {
	out := new(DataReply)
	err := c.cc.Invoke(ctx, Dianping_AccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DianpingServer is the server API for Dianping service.
// All implementations must embed UnimplementedDianpingServer
// for forward compatibility
type DianpingServer interface {
	Prepare(context.Context, *PrepareRequest) (*DataReply, error)
	ScanPrepare(context.Context, *PrepareRequest) (*DataReply, error)
	AccessToken(context.Context, *PrepareRequest) (*DataReply, error)
	mustEmbedUnimplementedDianpingServer()
}

// UnimplementedDianpingServer must be embedded to have forward compatible implementations.
type UnimplementedDianpingServer struct {
}

func (UnimplementedDianpingServer) Prepare(context.Context, *PrepareRequest) (*DataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedDianpingServer) ScanPrepare(context.Context, *PrepareRequest) (*DataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanPrepare not implemented")
}
func (UnimplementedDianpingServer) AccessToken(context.Context, *PrepareRequest) (*DataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessToken not implemented")
}
func (UnimplementedDianpingServer) mustEmbedUnimplementedDianpingServer() {}

// UnsafeDianpingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DianpingServer will
// result in compilation errors.
type UnsafeDianpingServer interface {
	mustEmbedUnimplementedDianpingServer()
}

func RegisterDianpingServer(s grpc.ServiceRegistrar, srv DianpingServer) {
	s.RegisterService(&Dianping_ServiceDesc, srv)
}

func _Dianping_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DianpingServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dianping_Prepare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DianpingServer).Prepare(ctx, req.(*PrepareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dianping_ScanPrepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DianpingServer).ScanPrepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dianping_ScanPrepare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DianpingServer).ScanPrepare(ctx, req.(*PrepareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dianping_AccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DianpingServer).AccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dianping_AccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DianpingServer).AccessToken(ctx, req.(*PrepareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dianping_ServiceDesc is the grpc.ServiceDesc for Dianping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dianping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.dianping.v1.Dianping",
	HandlerType: (*DianpingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "prepare",
			Handler:    _Dianping_Prepare_Handler,
		},
		{
			MethodName: "scanPrepare",
			Handler:    _Dianping_ScanPrepare_Handler,
		},
		{
			MethodName: "accessToken",
			Handler:    _Dianping_AccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dianping/v1/dianping.proto",
}
