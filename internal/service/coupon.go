package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"strconv"
	pb "studyRoomGo/api/coupon/v1"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/service/dto"
)

type CouponService struct {
	pb.UnimplementedCouponServer
	uc       *biz.CouponUsecase
	memberUc *biz.MemberUsecase
}

func NewCouponService(uc *biz.CouponUsecase, memberdUc *biz.MemberUsecase) *CouponService {
	return &CouponService{uc: uc, memberUc: memberdUc}
}

func (s *CouponService) CouponCenter(ctx context.Context, req *pb.CouponListRequest) (*pb.CouponListReply, error) {

	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	couponDTO := dto.CouponCenter{}
	isShow := int8(1)
	data, err := s.uc.CouponCenter(ctx, &biz.CouponCenterRequest{
		MerchantId: int64(merchantId),
		ShopId:     int64(shopId),
		Page:       int(req.Page),
		PageSize:   int(req.PageSize),
		IsShow:     &isShow,
	})
	if err != nil {
		return &pb.CouponListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	var list []*pb.CouponData
	for _, item := range data {
		couponDTO.Coupons = item
		list = append(list, couponDTO.CouponCenterReply())
	}
	return &pb.CouponListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *CouponService) ReceiveCoupon(ctx context.Context, req *pb.ReceiveCouponRequest) (*pb.ReceiveCouponReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	couponReceiveRequest := &biz.CouponReceiveRequest{
		CouponCode: req.CouponCode,
		MemberId:   memberId,
		MerchantId: int64(merchantId),
		ShopId:     int64(shopId),
		Notice:     req.Notice,
	}
	if ok, err := s.uc.ReceiveCouponValidate(ctx, couponReceiveRequest); !ok {
		return &pb.ReceiveCouponReply{
			Code: err.GetCode(),
			Msg:  err.GetReason(),
		}, nil
	}

	_, err := s.uc.ReceiveCoupon(ctx, couponReceiveRequest)
	if err != nil {
		return &pb.ReceiveCouponReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.ReceiveCouponReply{
		Code: 200,
		Msg:  "",
	}, nil
}

func (s *CouponService) MyCoupon(ctx context.Context, req *pb.CouponListRequest) (*pb.CouponListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	couponDTO := dto.MyCoupon{}

	data, err := s.uc.MyCouponList(ctx, &biz.SearchCouponReceiveRequest{
		MemberId: memberId,
		NotUsed:  int8(req.NotUsed),
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	})
	if err != nil {
		return &pb.CouponListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	var list []*pb.CouponData
	for _, item := range data {
		couponDTO.CouponsReceive = item
		list = append(list, couponDTO.CouponReply())
	}
	return &pb.CouponListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *CouponService) UseCouponValidate(ctx context.Context, req *pb.CouponUseRequest) (*pb.CouponUseValidateReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	dto := dto.Coupon{}
	qty := int32(1)
	if req.Qty != 0 {
		qty = req.Qty
	}
	couponUseRequest := &biz.CouponUseRequest{
		CouponCode: req.CouponCode,
		MemberId:   memberId,
		PackageId:  req.CardId,
		MerchantId: int64(merchantId),
		ShopId:     int64(shopId),
		Qty:        qty,
	}
	res, err := s.uc.UseCouponPreValidate(ctx, couponUseRequest)
	if err != nil {
		return &pb.CouponUseValidateReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.CouponUseValidateReply{
		Code: 200,
		Data: dto.CouponPreValidateReply(res),
	}, nil
}
