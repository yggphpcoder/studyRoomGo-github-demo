package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"strconv"
	pb "studyRoomGo/api/order/v1"
	"studyRoomGo/internal/biz"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	uc            *biz.OrderUsecase
	packageCardUc *biz.PackageCardUsecase
}

func NewOrderService(uc *biz.OrderUsecase, packageCardUc *biz.PackageCardUsecase) *OrderService {
	return &OrderService{uc: uc, packageCardUc: packageCardUc}
}

func (s *OrderService) CalOrderTotalPrice(ctx context.Context, req *pb.BuyCardOrderRequest) (*pb.BuyCardOrderReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	qty := int32(1)
	if req.Qty != 0 {
		qty = req.Qty
	}
	totalPrice := s.uc.CalOrderTotalPrice(ctx, &biz.CreateBuyCardOrderRequest{
		MerchantId: int64(merchantId),
		MemberId:   memberId,
		CardId:     req.CardId,
		CouponCode: req.CouponCode,
		Qty:        qty,
	})
	return &pb.BuyCardOrderReply{
		Code: 200,
		Msg:  "",
		Data: &pb.BuyCardOrderReply_Data{
			TotalPrice: fmt.Sprintf("%.2f", totalPrice),
		},
	}, nil
}
