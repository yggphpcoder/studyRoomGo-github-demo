// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.21.3
// source: order/v1/order.proto

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

const OperationOrderCalOrderTotalPrice = "/api.order.v1.Order/CalOrderTotalPrice"

type OrderHTTPServer interface {
	CalOrderTotalPrice(context.Context, *BuyCardOrderRequest) (*BuyCardOrderReply, error)
}

func RegisterOrderHTTPServer(s *http.Server, srv OrderHTTPServer) {
	r := s.Route("/")
	r.POST("/order/v1/cal_order_total_price", _Order_CalOrderTotalPrice0_HTTP_Handler(srv))
}

func _Order_CalOrderTotalPrice0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BuyCardOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderCalOrderTotalPrice)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CalOrderTotalPrice(ctx, req.(*BuyCardOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BuyCardOrderReply)
		return ctx.Result(200, reply)
	}
}

type OrderHTTPClient interface {
	CalOrderTotalPrice(ctx context.Context, req *BuyCardOrderRequest, opts ...http.CallOption) (rsp *BuyCardOrderReply, err error)
}

type OrderHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderHTTPClient(client *http.Client) OrderHTTPClient {
	return &OrderHTTPClientImpl{client}
}

func (c *OrderHTTPClientImpl) CalOrderTotalPrice(ctx context.Context, in *BuyCardOrderRequest, opts ...http.CallOption) (*BuyCardOrderReply, error) {
	var out BuyCardOrderReply
	pattern := "/order/v1/cal_order_total_price"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderCalOrderTotalPrice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
