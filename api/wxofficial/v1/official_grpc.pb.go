// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.3
// source: wxofficial/v1/official.proto

package v1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// WxOfficialClient is the client API for WxOfficial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WxOfficialClient interface {
}

type wxOfficialClient struct {
	cc grpc.ClientConnInterface
}

func NewWxOfficialClient(cc grpc.ClientConnInterface) WxOfficialClient {
	return &wxOfficialClient{cc}
}

// WxOfficialServer is the server API for WxOfficial service.
// All implementations must embed UnimplementedWxOfficialServer
// for forward compatibility
type WxOfficialServer interface {
	mustEmbedUnimplementedWxOfficialServer()
}

// UnimplementedWxOfficialServer must be embedded to have forward compatible implementations.
type UnimplementedWxOfficialServer struct {
}

func (UnimplementedWxOfficialServer) mustEmbedUnimplementedWxOfficialServer() {}

// UnsafeWxOfficialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WxOfficialServer will
// result in compilation errors.
type UnsafeWxOfficialServer interface {
	mustEmbedUnimplementedWxOfficialServer()
}

func RegisterWxOfficialServer(s grpc.ServiceRegistrar, srv WxOfficialServer) {
	s.RegisterService(&WxOfficial_ServiceDesc, srv)
}

// WxOfficial_ServiceDesc is the grpc.ServiceDesc for WxOfficial service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WxOfficial_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.auth.v1.WxOfficial",
	HandlerType: (*WxOfficialServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "wxofficial/v1/official.proto",
}
