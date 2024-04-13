// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.21.3
// source: payment/v1/payment.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPaymentWxPayCheckout = "/api.payment.v1.Payment/WxPayCheckout"
const OperationPaymentWxPayCheckoutBuyCard = "/api.payment.v1.Payment/WxPayCheckoutBuyCard"
const OperationPaymentWxPayCheckoutBuyCourse = "/api.payment.v1.Payment/WxPayCheckoutBuyCourse"
const OperationPaymentWxPayCheckoutReturn = "/api.payment.v1.Payment/WxPayCheckoutReturn"
const OperationPaymentWxPayNotice = "/api.payment.v1.Payment/WxPayNotice"

type PaymentHTTPServer interface {
	WxPayCheckout(context.Context, *WxCheckoutRequest) (*WxCheckoutReply, error)
	WxPayCheckoutBuyCard(context.Context, *WxCheckoutBuyCardRequest) (*WxCheckoutReply, error)
	WxPayCheckoutBuyCourse(context.Context, *WxCheckoutBuyCourseRequest) (*WxCheckoutReply, error)
	WxPayCheckoutReturn(context.Context, *WxCheckoutReturnRequest) (*WxCheckoutReturnReply, error)
	WxPayNotice(context.Context, *WxNoticeRequest) (*WxNoticeReply, error)
}

func RegisterPaymentHTTPServer(s *http.Server, srv PaymentHTTPServer) {
	r := s.Route("/")
	r.POST("/payment/v1/wxpay_checkout", _Payment_WxPayCheckout0_HTTP_Handler(srv))
	r.POST("/payment/v1/wxpay_checkout_return", _Payment_WxPayCheckoutReturn0_HTTP_Handler(srv))
	r.POST("/payment/v1/{merchant_sign}/wxpay_notice", _Payment_WxPayNotice0_HTTP_Handler(srv))
	r.POST("/payment/v1/wxpay_checkout_buy_card", _Payment_WxPayCheckoutBuyCard0_HTTP_Handler(srv))
	r.POST("/payment/v1/wxpay_checkout_buy_course", _Payment_WxPayCheckoutBuyCourse0_HTTP_Handler(srv))
}

func _Payment_WxPayCheckout0_HTTP_Handler(srv PaymentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxCheckoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPaymentWxPayCheckout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxPayCheckout(ctx, req.(*WxCheckoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxCheckoutReply)
		return ctx.Result(200, reply)
	}
}

func _Payment_WxPayCheckoutReturn0_HTTP_Handler(srv PaymentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxCheckoutReturnRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPaymentWxPayCheckoutReturn)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxPayCheckoutReturn(ctx, req.(*WxCheckoutReturnRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxCheckoutReturnReply)
		return ctx.Result(200, reply)
	}
}

func _Payment_WxPayNotice0_HTTP_Handler(srv PaymentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxNoticeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPaymentWxPayNotice)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxPayNotice(ctx, req.(*WxNoticeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxNoticeReply)
		return ctx.Result(200, reply)
	}
}

func _Payment_WxPayCheckoutBuyCard0_HTTP_Handler(srv PaymentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxCheckoutBuyCardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPaymentWxPayCheckoutBuyCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxPayCheckoutBuyCard(ctx, req.(*WxCheckoutBuyCardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxCheckoutReply)
		return ctx.Result(200, reply)
	}
}

func _Payment_WxPayCheckoutBuyCourse0_HTTP_Handler(srv PaymentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxCheckoutBuyCourseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPaymentWxPayCheckoutBuyCourse)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxPayCheckoutBuyCourse(ctx, req.(*WxCheckoutBuyCourseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxCheckoutReply)
		return ctx.Result(200, reply)
	}
}

type PaymentHTTPClient interface {
	WxPayCheckout(ctx context.Context, req *WxCheckoutRequest, opts ...http.CallOption) (rsp *WxCheckoutReply, err error)
	WxPayCheckoutBuyCard(ctx context.Context, req *WxCheckoutBuyCardRequest, opts ...http.CallOption) (rsp *WxCheckoutReply, err error)
	WxPayCheckoutBuyCourse(ctx context.Context, req *WxCheckoutBuyCourseRequest, opts ...http.CallOption) (rsp *WxCheckoutReply, err error)
	WxPayCheckoutReturn(ctx context.Context, req *WxCheckoutReturnRequest, opts ...http.CallOption) (rsp *WxCheckoutReturnReply, err error)
	WxPayNotice(ctx context.Context, req *WxNoticeRequest, opts ...http.CallOption) (rsp *WxNoticeReply, err error)
}

type PaymentHTTPClientImpl struct {
	cc *http.Client
}

func NewPaymentHTTPClient(client *http.Client) PaymentHTTPClient {
	return &PaymentHTTPClientImpl{client}
}

func (c *PaymentHTTPClientImpl) WxPayCheckout(ctx context.Context, in *WxCheckoutRequest, opts ...http.CallOption) (*WxCheckoutReply, error) {
	var out WxCheckoutReply
	pattern := "/payment/v1/wxpay_checkout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPaymentWxPayCheckout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PaymentHTTPClientImpl) WxPayCheckoutBuyCard(ctx context.Context, in *WxCheckoutBuyCardRequest, opts ...http.CallOption) (*WxCheckoutReply, error) {
	var out WxCheckoutReply
	pattern := "/payment/v1/wxpay_checkout_buy_card"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPaymentWxPayCheckoutBuyCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PaymentHTTPClientImpl) WxPayCheckoutBuyCourse(ctx context.Context, in *WxCheckoutBuyCourseRequest, opts ...http.CallOption) (*WxCheckoutReply, error) {
	var out WxCheckoutReply
	pattern := "/payment/v1/wxpay_checkout_buy_course"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPaymentWxPayCheckoutBuyCourse))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PaymentHTTPClientImpl) WxPayCheckoutReturn(ctx context.Context, in *WxCheckoutReturnRequest, opts ...http.CallOption) (*WxCheckoutReturnReply, error) {
	var out WxCheckoutReturnReply
	pattern := "/payment/v1/wxpay_checkout_return"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPaymentWxPayCheckoutReturn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PaymentHTTPClientImpl) WxPayNotice(ctx context.Context, in *WxNoticeRequest, opts ...http.CallOption) (*WxNoticeReply, error) {
	var out WxNoticeReply
	pattern := "/payment/v1/{merchant_sign}/wxpay_notice"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPaymentWxPayNotice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
