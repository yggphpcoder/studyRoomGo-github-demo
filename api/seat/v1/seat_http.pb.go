// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.21.3
// source: seat/v1/seat.proto

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

const OperationSeatBooking = "/api.seat.v1.Seat/Booking"
const OperationSeatBookingAddTime = "/api.seat.v1.Seat/BookingAddTime"
const OperationSeatBookingChangeSeat = "/api.seat.v1.Seat/BookingChangeSeat"
const OperationSeatBookingChangeStatus = "/api.seat.v1.Seat/BookingChangeStatus"
const OperationSeatBookingControlLight = "/api.seat.v1.Seat/BookingControlLight"
const OperationSeatBookingControlLightShared = "/api.seat.v1.Seat/BookingControlLightShared"
const OperationSeatBookingOpenDoor = "/api.seat.v1.Seat/BookingOpenDoor"
const OperationSeatBookingOpenDoorShared = "/api.seat.v1.Seat/BookingOpenDoorShared"
const OperationSeatBookingShareGenToken = "/api.seat.v1.Seat/BookingShareGenToken"
const OperationSeatBookingValidate = "/api.seat.v1.Seat/BookingValidate"
const OperationSeatFind = "/api.seat.v1.Seat/Find"
const OperationSeatGet = "/api.seat.v1.Seat/Get"
const OperationSeatList = "/api.seat.v1.Seat/List"
const OperationSeatLockingControlLight = "/api.seat.v1.Seat/LockingControlLight"
const OperationSeatLockingOpenDoor = "/api.seat.v1.Seat/LockingOpenDoor"
const OperationSeatOverview = "/api.seat.v1.Seat/Overview"
const OperationSeatSearch = "/api.seat.v1.Seat/Search"

type SeatHTTPServer interface {
	Booking(context.Context, *BookingRequest) (*BookingReply, error)
	BookingAddTime(context.Context, *BookingAddTimeRequest) (*BookingReply, error)
	BookingChangeSeat(context.Context, *BookingChangeSeatRequest) (*BookingReply, error)
	BookingChangeStatus(context.Context, *BookingChangeStatusRequest) (*BookingReply, error)
	BookingControlLight(context.Context, *BookingControlSeatRequest) (*BookingReply, error)
	BookingControlLightShared(context.Context, *BookingControlSeatRequest) (*BookingReply, error)
	BookingOpenDoor(context.Context, *BookingOpenDoorRequest) (*BookingReply, error)
	BookingOpenDoorShared(context.Context, *BookingOpenDoorRequest) (*BookingReply, error)
	BookingShareGenToken(context.Context, *BookingControlSeatRequest) (*BookingGenSharekDataReply, error)
	BookingValidate(context.Context, *BookingRequest) (*BookingValidateReply, error)
	Find(context.Context, *FindRequest) (*DataReply, error)
	Get(context.Context, *GetRequest) (*DataReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	LockingControlLight(context.Context, *LockingControlSeatRequest) (*BookingReply, error)
	LockingOpenDoor(context.Context, *LockingOpenDoorRequest) (*BookingReply, error)
	Overview(context.Context, *OverviewRequest) (*OverviewReply, error)
	Search(context.Context, *SearchRequest) (*ListReply, error)
}

func RegisterSeatHTTPServer(s *http.Server, srv SeatHTTPServer) {
	r := s.Route("/")
	r.GET("/seat/v1/seat/{id}", _Seat_Get2_HTTP_Handler(srv))
	r.GET("/seat/v1/find", _Seat_Find2_HTTP_Handler(srv))
	r.GET("/seat/v1/list", _Seat_List2_HTTP_Handler(srv))
	r.GET("/seat/v1/search", _Seat_Search2_HTTP_Handler(srv))
	r.GET("/seat/v1/overview", _Seat_Overview0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_validate", _Seat_BookingValidate0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking", _Seat_Booking0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_change_seat", _Seat_BookingChangeSeat0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_add_time", _Seat_BookingAddTime0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_change_status", _Seat_BookingChangeStatus0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_open_door", _Seat_BookingOpenDoor0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_control_light", _Seat_BookingControlLight0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_share_gen_token", _Seat_BookingShareGenToken0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_open_door_shared", _Seat_BookingOpenDoorShared0_HTTP_Handler(srv))
	r.POST("/seat/v1/booking_control_light_shared", _Seat_BookingControlLightShared0_HTTP_Handler(srv))
	r.POST("/seat/v1/locking_open_door", _Seat_LockingOpenDoor0_HTTP_Handler(srv))
	r.POST("/seat/v1/locking_control_light", _Seat_LockingControlLight0_HTTP_Handler(srv))
}

func _Seat_Get2_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatGet)
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

func _Seat_Find2_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FindRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatFind)
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

func _Seat_List2_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatList)
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

func _Seat_Search2_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatSearch)
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

func _Seat_Overview0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OverviewRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatOverview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Overview(ctx, req.(*OverviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OverviewReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingValidate0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingValidate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingValidate(ctx, req.(*BookingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingValidateReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_Booking0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBooking)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Booking(ctx, req.(*BookingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingChangeSeat0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingChangeSeatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingChangeSeat)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingChangeSeat(ctx, req.(*BookingChangeSeatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingAddTime0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingAddTimeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingAddTime)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingAddTime(ctx, req.(*BookingAddTimeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingChangeStatus0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingChangeStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingChangeStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingChangeStatus(ctx, req.(*BookingChangeStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingOpenDoor0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingOpenDoorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingOpenDoor)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingOpenDoor(ctx, req.(*BookingOpenDoorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingControlLight0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingControlSeatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingControlLight)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingControlLight(ctx, req.(*BookingControlSeatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingShareGenToken0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingControlSeatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingShareGenToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingShareGenToken(ctx, req.(*BookingControlSeatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingGenSharekDataReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingOpenDoorShared0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingOpenDoorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingOpenDoorShared)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingOpenDoorShared(ctx, req.(*BookingOpenDoorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_BookingControlLightShared0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BookingControlSeatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatBookingControlLightShared)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookingControlLightShared(ctx, req.(*BookingControlSeatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_LockingOpenDoor0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LockingOpenDoorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatLockingOpenDoor)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LockingOpenDoor(ctx, req.(*LockingOpenDoorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

func _Seat_LockingControlLight0_HTTP_Handler(srv SeatHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LockingControlSeatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSeatLockingControlLight)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LockingControlLight(ctx, req.(*LockingControlSeatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BookingReply)
		return ctx.Result(200, reply)
	}
}

type SeatHTTPClient interface {
	Booking(ctx context.Context, req *BookingRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	BookingAddTime(ctx context.Context, req *BookingAddTimeRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	BookingChangeSeat(ctx context.Context, req *BookingChangeSeatRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	BookingChangeStatus(ctx context.Context, req *BookingChangeStatusRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	BookingControlLight(ctx context.Context, req *BookingControlSeatRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	BookingControlLightShared(ctx context.Context, req *BookingControlSeatRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	BookingOpenDoor(ctx context.Context, req *BookingOpenDoorRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	BookingOpenDoorShared(ctx context.Context, req *BookingOpenDoorRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	BookingShareGenToken(ctx context.Context, req *BookingControlSeatRequest, opts ...http.CallOption) (rsp *BookingGenSharekDataReply, err error)
	BookingValidate(ctx context.Context, req *BookingRequest, opts ...http.CallOption) (rsp *BookingValidateReply, err error)
	Find(ctx context.Context, req *FindRequest, opts ...http.CallOption) (rsp *DataReply, err error)
	Get(ctx context.Context, req *GetRequest, opts ...http.CallOption) (rsp *DataReply, err error)
	List(ctx context.Context, req *ListRequest, opts ...http.CallOption) (rsp *ListReply, err error)
	LockingControlLight(ctx context.Context, req *LockingControlSeatRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	LockingOpenDoor(ctx context.Context, req *LockingOpenDoorRequest, opts ...http.CallOption) (rsp *BookingReply, err error)
	Overview(ctx context.Context, req *OverviewRequest, opts ...http.CallOption) (rsp *OverviewReply, err error)
	Search(ctx context.Context, req *SearchRequest, opts ...http.CallOption) (rsp *ListReply, err error)
}

type SeatHTTPClientImpl struct {
	cc *http.Client
}

func NewSeatHTTPClient(client *http.Client) SeatHTTPClient {
	return &SeatHTTPClientImpl{client}
}

func (c *SeatHTTPClientImpl) Booking(ctx context.Context, in *BookingRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/booking"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBooking))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingAddTime(ctx context.Context, in *BookingAddTimeRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/booking_add_time"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingAddTime))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingChangeSeat(ctx context.Context, in *BookingChangeSeatRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/booking_change_seat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingChangeSeat))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingChangeStatus(ctx context.Context, in *BookingChangeStatusRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/booking_change_status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingChangeStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingControlLight(ctx context.Context, in *BookingControlSeatRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/booking_control_light"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingControlLight))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingControlLightShared(ctx context.Context, in *BookingControlSeatRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/booking_control_light_shared"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingControlLightShared))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingOpenDoor(ctx context.Context, in *BookingOpenDoorRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/booking_open_door"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingOpenDoor))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingOpenDoorShared(ctx context.Context, in *BookingOpenDoorRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/booking_open_door_shared"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingOpenDoorShared))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingShareGenToken(ctx context.Context, in *BookingControlSeatRequest, opts ...http.CallOption) (*BookingGenSharekDataReply, error) {
	var out BookingGenSharekDataReply
	pattern := "/seat/v1/booking_share_gen_token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingShareGenToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) BookingValidate(ctx context.Context, in *BookingRequest, opts ...http.CallOption) (*BookingValidateReply, error) {
	var out BookingValidateReply
	pattern := "/seat/v1/booking_validate"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatBookingValidate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) Find(ctx context.Context, in *FindRequest, opts ...http.CallOption) (*DataReply, error) {
	var out DataReply
	pattern := "/seat/v1/find"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSeatFind))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) Get(ctx context.Context, in *GetRequest, opts ...http.CallOption) (*DataReply, error) {
	var out DataReply
	pattern := "/seat/v1/seat/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSeatGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) List(ctx context.Context, in *ListRequest, opts ...http.CallOption) (*ListReply, error) {
	var out ListReply
	pattern := "/seat/v1/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSeatList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) LockingControlLight(ctx context.Context, in *LockingControlSeatRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/locking_control_light"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatLockingControlLight))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) LockingOpenDoor(ctx context.Context, in *LockingOpenDoorRequest, opts ...http.CallOption) (*BookingReply, error) {
	var out BookingReply
	pattern := "/seat/v1/locking_open_door"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSeatLockingOpenDoor))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) Overview(ctx context.Context, in *OverviewRequest, opts ...http.CallOption) (*OverviewReply, error) {
	var out OverviewReply
	pattern := "/seat/v1/overview"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSeatOverview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SeatHTTPClientImpl) Search(ctx context.Context, in *SearchRequest, opts ...http.CallOption) (*ListReply, error) {
	var out ListReply
	pattern := "/seat/v1/search"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSeatSearch))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
