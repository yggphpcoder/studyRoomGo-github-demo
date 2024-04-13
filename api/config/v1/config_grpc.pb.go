// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.3
// source: config/v1/config.proto

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
	Config_GetDict_FullMethodName       = "/api.config.v1.Config/GetDict"
	Config_FindDict_FullMethodName      = "/api.config.v1.Config/FindDict"
	Config_ListDict_FullMethodName      = "/api.config.v1.Config/ListDict"
	Config_SearchDict_FullMethodName    = "/api.config.v1.Config/SearchDict"
	Config_SearchEdDict_FullMethodName  = "/api.config.v1.Config/SearchEdDict"
	Config_SearchSetting_FullMethodName = "/api.config.v1.Config/SearchSetting"
)

// ConfigClient is the client API for Config service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigClient interface {
	// rpc Create (CreateRequest) returns (CreateReply);
	// rpc Update (UpdateRequest) returns (UpdateReply);
	// rpc Delete (DeleteRequest) returns (DeleteReply);
	GetDict(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataDictReply, error)
	FindDict(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*DataDictReply, error)
	ListDict(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListDictReply, error)
	SearchDict(ctx context.Context, in *SearchDictRequest, opts ...grpc.CallOption) (*ListDictReply, error)
	SearchEdDict(ctx context.Context, in *SearchDictRequest, opts ...grpc.CallOption) (*ListDictReply, error)
	SearchSetting(ctx context.Context, in *SearchSettingRequest, opts ...grpc.CallOption) (*ListSettingReply, error)
}

type configClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigClient(cc grpc.ClientConnInterface) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) GetDict(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DataDictReply, error) {
	out := new(DataDictReply)
	err := c.cc.Invoke(ctx, Config_GetDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) FindDict(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*DataDictReply, error) {
	out := new(DataDictReply)
	err := c.cc.Invoke(ctx, Config_FindDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) ListDict(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListDictReply, error) {
	out := new(ListDictReply)
	err := c.cc.Invoke(ctx, Config_ListDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) SearchDict(ctx context.Context, in *SearchDictRequest, opts ...grpc.CallOption) (*ListDictReply, error) {
	out := new(ListDictReply)
	err := c.cc.Invoke(ctx, Config_SearchDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) SearchEdDict(ctx context.Context, in *SearchDictRequest, opts ...grpc.CallOption) (*ListDictReply, error) {
	out := new(ListDictReply)
	err := c.cc.Invoke(ctx, Config_SearchEdDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) SearchSetting(ctx context.Context, in *SearchSettingRequest, opts ...grpc.CallOption) (*ListSettingReply, error) {
	out := new(ListSettingReply)
	err := c.cc.Invoke(ctx, Config_SearchSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServer is the server API for Config service.
// All implementations must embed UnimplementedConfigServer
// for forward compatibility
type ConfigServer interface {
	// rpc Create (CreateRequest) returns (CreateReply);
	// rpc Update (UpdateRequest) returns (UpdateReply);
	// rpc Delete (DeleteRequest) returns (DeleteReply);
	GetDict(context.Context, *GetRequest) (*DataDictReply, error)
	FindDict(context.Context, *FindRequest) (*DataDictReply, error)
	ListDict(context.Context, *ListRequest) (*ListDictReply, error)
	SearchDict(context.Context, *SearchDictRequest) (*ListDictReply, error)
	SearchEdDict(context.Context, *SearchDictRequest) (*ListDictReply, error)
	SearchSetting(context.Context, *SearchSettingRequest) (*ListSettingReply, error)
	mustEmbedUnimplementedConfigServer()
}

// UnimplementedConfigServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServer struct {
}

func (UnimplementedConfigServer) GetDict(context.Context, *GetRequest) (*DataDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDict not implemented")
}
func (UnimplementedConfigServer) FindDict(context.Context, *FindRequest) (*DataDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDict not implemented")
}
func (UnimplementedConfigServer) ListDict(context.Context, *ListRequest) (*ListDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDict not implemented")
}
func (UnimplementedConfigServer) SearchDict(context.Context, *SearchDictRequest) (*ListDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDict not implemented")
}
func (UnimplementedConfigServer) SearchEdDict(context.Context, *SearchDictRequest) (*ListDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchEdDict not implemented")
}
func (UnimplementedConfigServer) SearchSetting(context.Context, *SearchSettingRequest) (*ListSettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSetting not implemented")
}
func (UnimplementedConfigServer) mustEmbedUnimplementedConfigServer() {}

// UnsafeConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServer will
// result in compilation errors.
type UnsafeConfigServer interface {
	mustEmbedUnimplementedConfigServer()
}

func RegisterConfigServer(s grpc.ServiceRegistrar, srv ConfigServer) {
	s.RegisterService(&Config_ServiceDesc, srv)
}

func _Config_GetDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_GetDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetDict(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_FindDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).FindDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_FindDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).FindDict(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_ListDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).ListDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_ListDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).ListDict(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_SearchDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).SearchDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_SearchDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).SearchDict(ctx, req.(*SearchDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_SearchEdDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).SearchEdDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_SearchEdDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).SearchEdDict(ctx, req.(*SearchDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_SearchSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).SearchSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_SearchSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).SearchSetting(ctx, req.(*SearchSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Config_ServiceDesc is the grpc.ServiceDesc for Config service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Config_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.config.v1.Config",
	HandlerType: (*ConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDict",
			Handler:    _Config_GetDict_Handler,
		},
		{
			MethodName: "FindDict",
			Handler:    _Config_FindDict_Handler,
		},
		{
			MethodName: "ListDict",
			Handler:    _Config_ListDict_Handler,
		},
		{
			MethodName: "SearchDict",
			Handler:    _Config_SearchDict_Handler,
		},
		{
			MethodName: "SearchEdDict",
			Handler:    _Config_SearchEdDict_Handler,
		},
		{
			MethodName: "SearchSetting",
			Handler:    _Config_SearchSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/v1/config.proto",
}
