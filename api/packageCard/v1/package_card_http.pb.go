// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.21.3
// source: packageCard/v1/package_card.proto

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

const OperationPackageCardBuyCard = "/api.packageCard.v1.PackageCard/BuyCard"
const OperationPackageCardFind = "/api.packageCard.v1.PackageCard/Find"
const OperationPackageCardFindBuyLog = "/api.packageCard.v1.PackageCard/FindBuyLog"
const OperationPackageCardGet = "/api.packageCard.v1.PackageCard/Get"
const OperationPackageCardList = "/api.packageCard.v1.PackageCard/List"
const OperationPackageCardMemberCardActive = "/api.packageCard.v1.PackageCard/MemberCardActive"
const OperationPackageCardSearch = "/api.packageCard.v1.PackageCard/Search"

type PackageCardHTTPServer interface {
	BuyCard(context.Context, *CardRequest) (*CardReply, error)
	Find(context.Context, *FindRequest) (*DataReply, error)
	FindBuyLog(context.Context, *FindRequest) (*BuyLogDataReply, error)
	Get(context.Context, *GetRequest) (*DataReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	MemberCardActive(context.Context, *CardRequest) (*CardReply, error)
	Search(context.Context, *SearchRequest) (*ListReply, error)
}

func RegisterPackageCardHTTPServer(s *http.Server, srv PackageCardHTTPServer) {
	r := s.Route("/")
	r.GET("/package_card/v1/package_card/{id}", _PackageCard_Get1_HTTP_Handler(srv))
	r.GET("/package_card/v1/find", _PackageCard_Find1_HTTP_Handler(srv))
	r.GET("/package_card/v1/list", _PackageCard_List1_HTTP_Handler(srv))
	r.GET("/package_card/v1/search", _PackageCard_Search1_HTTP_Handler(srv))
	r.POST("/package_card/v1/buyCard", _PackageCard_BuyCard0_HTTP_Handler(srv))
	r.POST("/package_card/v1/member_card_active", _PackageCard_MemberCardActive0_HTTP_Handler(srv))
	r.GET("/package_card/v1/findBuyLog", _PackageCard_FindBuyLog0_HTTP_Handler(srv))
}

func _PackageCard_Get1_HTTP_Handler(srv PackageCardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPackageCardGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*GetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DataReply)
		return ctx.Result(200, reply)
	}
}

func _PackageCard_Find1_HTTP_Handler(srv PackageCardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FindRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPackageCardFind)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Find(ctx, req.(*FindRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DataReply)
		return ctx.Result(200, reply)
	}
}

func _PackageCard_List1_HTTP_Handler(srv PackageCardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPackageCardList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListReply)
		return ctx.Result(200, reply)
	}
}

func _PackageCard_Search1_HTTP_Handler(srv PackageCardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPackageCardSearch)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*SearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListReply)
		return ctx.Result(200, reply)
	}
}

func _PackageCard_BuyCard0_HTTP_Handler(srv PackageCardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPackageCardBuyCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BuyCard(ctx, req.(*CardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CardReply)
		return ctx.Result(200, reply)
	}
}

func _PackageCard_MemberCardActive0_HTTP_Handler(srv PackageCardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPackageCardMemberCardActive)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberCardActive(ctx, req.(*CardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CardReply)
		return ctx.Result(200, reply)
	}
}

func _PackageCard_FindBuyLog0_HTTP_Handler(srv PackageCardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FindRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPackageCardFindBuyLog)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FindBuyLog(ctx, req.(*FindRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BuyLogDataReply)
		return ctx.Result(200, reply)
	}
}

type PackageCardHTTPClient interface {
	BuyCard(ctx context.Context, req *CardRequest, opts ...http.CallOption) (rsp *CardReply, err error)
	Find(ctx context.Context, req *FindRequest, opts ...http.CallOption) (rsp *DataReply, err error)
	FindBuyLog(ctx context.Context, req *FindRequest, opts ...http.CallOption) (rsp *BuyLogDataReply, err error)
	Get(ctx context.Context, req *GetRequest, opts ...http.CallOption) (rsp *DataReply, err error)
	List(ctx context.Context, req *ListRequest, opts ...http.CallOption) (rsp *ListReply, err error)
	MemberCardActive(ctx context.Context, req *CardRequest, opts ...http.CallOption) (rsp *CardReply, err error)
	Search(ctx context.Context, req *SearchRequest, opts ...http.CallOption) (rsp *ListReply, err error)
}

type PackageCardHTTPClientImpl struct {
	cc *http.Client
}

func NewPackageCardHTTPClient(client *http.Client) PackageCardHTTPClient {
	return &PackageCardHTTPClientImpl{client}
}

func (c *PackageCardHTTPClientImpl) BuyCard(ctx context.Context, in *CardRequest, opts ...http.CallOption) (*CardReply, error) {
	var out CardReply
	pattern := "/package_card/v1/buyCard"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPackageCardBuyCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PackageCardHTTPClientImpl) Find(ctx context.Context, in *FindRequest, opts ...http.CallOption) (*DataReply, error) {
	var out DataReply
	pattern := "/package_card/v1/find"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPackageCardFind))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PackageCardHTTPClientImpl) FindBuyLog(ctx context.Context, in *FindRequest, opts ...http.CallOption) (*BuyLogDataReply, error) {
	var out BuyLogDataReply
	pattern := "/package_card/v1/findBuyLog"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPackageCardFindBuyLog))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PackageCardHTTPClientImpl) Get(ctx context.Context, in *GetRequest, opts ...http.CallOption) (*DataReply, error) {
	var out DataReply
	pattern := "/package_card/v1/package_card/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPackageCardGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PackageCardHTTPClientImpl) List(ctx context.Context, in *ListRequest, opts ...http.CallOption) (*ListReply, error) {
	var out ListReply
	pattern := "/package_card/v1/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPackageCardList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PackageCardHTTPClientImpl) MemberCardActive(ctx context.Context, in *CardRequest, opts ...http.CallOption) (*CardReply, error) {
	var out CardReply
	pattern := "/package_card/v1/member_card_active"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPackageCardMemberCardActive))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PackageCardHTTPClientImpl) Search(ctx context.Context, in *SearchRequest, opts ...http.CallOption) (*ListReply, error) {
	var out ListReply
	pattern := "/package_card/v1/search"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPackageCardSearch))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
