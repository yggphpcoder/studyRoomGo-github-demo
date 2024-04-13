// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.3
// source: payment/v1/payment.proto

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
	Payment_WxPayCheckout_FullMethodName          = "/api.payment.v1.Payment/WxPayCheckout"
	Payment_WxPayCheckoutReturn_FullMethodName    = "/api.payment.v1.Payment/WxPayCheckoutReturn"
	Payment_WxPayNotice_FullMethodName            = "/api.payment.v1.Payment/WxPayNotice"
	Payment_WxPayCheckoutBuyCard_FullMethodName   = "/api.payment.v1.Payment/WxPayCheckoutBuyCard"
	Payment_WxPayCheckoutBuyCourse_FullMethodName = "/api.payment.v1.Payment/WxPayCheckoutBuyCourse"
)

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	WxPayCheckout(ctx context.Context, in *WxCheckoutRequest, opts ...grpc.CallOption) (*WxCheckoutReply, error)
	WxPayCheckoutReturn(ctx context.Context, in *WxCheckoutReturnRequest, opts ...grpc.CallOption) (*WxCheckoutReturnReply, error)
	WxPayNotice(ctx context.Context, in *WxNoticeRequest, opts ...grpc.CallOption) (*WxNoticeReply, error)
	WxPayCheckoutBuyCard(ctx context.Context, in *WxCheckoutBuyCardRequest, opts ...grpc.CallOption) (*WxCheckoutReply, error)
	WxPayCheckoutBuyCourse(ctx context.Context, in *WxCheckoutBuyCourseRequest, opts ...grpc.CallOption) (*WxCheckoutReply, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) WxPayCheckout(ctx context.Context, in *WxCheckoutRequest, opts ...grpc.CallOption) (*WxCheckoutReply, error) {
	out := new(WxCheckoutReply)
	err := c.cc.Invoke(ctx, Payment_WxPayCheckout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) WxPayCheckoutReturn(ctx context.Context, in *WxCheckoutReturnRequest, opts ...grpc.CallOption) (*WxCheckoutReturnReply, error) {
	out := new(WxCheckoutReturnReply)
	err := c.cc.Invoke(ctx, Payment_WxPayCheckoutReturn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) WxPayNotice(ctx context.Context, in *WxNoticeRequest, opts ...grpc.CallOption) (*WxNoticeReply, error) {
	out := new(WxNoticeReply)
	err := c.cc.Invoke(ctx, Payment_WxPayNotice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) WxPayCheckoutBuyCard(ctx context.Context, in *WxCheckoutBuyCardRequest, opts ...grpc.CallOption) (*WxCheckoutReply, error) {
	out := new(WxCheckoutReply)
	err := c.cc.Invoke(ctx, Payment_WxPayCheckoutBuyCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) WxPayCheckoutBuyCourse(ctx context.Context, in *WxCheckoutBuyCourseRequest, opts ...grpc.CallOption) (*WxCheckoutReply, error) {
	out := new(WxCheckoutReply)
	err := c.cc.Invoke(ctx, Payment_WxPayCheckoutBuyCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility
type PaymentServer interface {
	WxPayCheckout(context.Context, *WxCheckoutRequest) (*WxCheckoutReply, error)
	WxPayCheckoutReturn(context.Context, *WxCheckoutReturnRequest) (*WxCheckoutReturnReply, error)
	WxPayNotice(context.Context, *WxNoticeRequest) (*WxNoticeReply, error)
	WxPayCheckoutBuyCard(context.Context, *WxCheckoutBuyCardRequest) (*WxCheckoutReply, error)
	WxPayCheckoutBuyCourse(context.Context, *WxCheckoutBuyCourseRequest) (*WxCheckoutReply, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServer struct {
}

func (UnimplementedPaymentServer) WxPayCheckout(context.Context, *WxCheckoutRequest) (*WxCheckoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxPayCheckout not implemented")
}
func (UnimplementedPaymentServer) WxPayCheckoutReturn(context.Context, *WxCheckoutReturnRequest) (*WxCheckoutReturnReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxPayCheckoutReturn not implemented")
}
func (UnimplementedPaymentServer) WxPayNotice(context.Context, *WxNoticeRequest) (*WxNoticeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxPayNotice not implemented")
}
func (UnimplementedPaymentServer) WxPayCheckoutBuyCard(context.Context, *WxCheckoutBuyCardRequest) (*WxCheckoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxPayCheckoutBuyCard not implemented")
}
func (UnimplementedPaymentServer) WxPayCheckoutBuyCourse(context.Context, *WxCheckoutBuyCourseRequest) (*WxCheckoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxPayCheckoutBuyCourse not implemented")
}
func (UnimplementedPaymentServer) mustEmbedUnimplementedPaymentServer() {}

// UnsafePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServer will
// result in compilation errors.
type UnsafePaymentServer interface {
	mustEmbedUnimplementedPaymentServer()
}

func RegisterPaymentServer(s grpc.ServiceRegistrar, srv PaymentServer) {
	s.RegisterService(&Payment_ServiceDesc, srv)
}

func _Payment_WxPayCheckout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxCheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).WxPayCheckout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_WxPayCheckout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).WxPayCheckout(ctx, req.(*WxCheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_WxPayCheckoutReturn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxCheckoutReturnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).WxPayCheckoutReturn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_WxPayCheckoutReturn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).WxPayCheckoutReturn(ctx, req.(*WxCheckoutReturnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_WxPayNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxNoticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).WxPayNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_WxPayNotice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).WxPayNotice(ctx, req.(*WxNoticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_WxPayCheckoutBuyCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxCheckoutBuyCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).WxPayCheckoutBuyCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_WxPayCheckoutBuyCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).WxPayCheckoutBuyCard(ctx, req.(*WxCheckoutBuyCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_WxPayCheckoutBuyCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxCheckoutBuyCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).WxPayCheckoutBuyCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_WxPayCheckoutBuyCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).WxPayCheckoutBuyCourse(ctx, req.(*WxCheckoutBuyCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_ServiceDesc is the grpc.ServiceDesc for Payment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.payment.v1.Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WxPayCheckout",
			Handler:    _Payment_WxPayCheckout_Handler,
		},
		{
			MethodName: "WxPayCheckoutReturn",
			Handler:    _Payment_WxPayCheckoutReturn_Handler,
		},
		{
			MethodName: "WxPayNotice",
			Handler:    _Payment_WxPayNotice_Handler,
		},
		{
			MethodName: "WxPayCheckoutBuyCard",
			Handler:    _Payment_WxPayCheckoutBuyCard_Handler,
		},
		{
			MethodName: "WxPayCheckoutBuyCourse",
			Handler:    _Payment_WxPayCheckoutBuyCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment/v1/payment.proto",
}
