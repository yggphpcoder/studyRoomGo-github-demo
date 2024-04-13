package service

import (
	"context"
	pb "studyRoomGo/api/dianping/v1"
	"studyRoomGo/common/dianping"
	"studyRoomGo/internal/biz"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

type DianpingService struct {
	pb.UnimplementedDianpingServer
	uc *biz.PackageCardUsecase
}

func NewDianpingService(uc *biz.PackageCardUsecase) *DianpingService {
	return &DianpingService{uc: uc}
}

func (s *DianpingService) Prepare(ctx context.Context, req *pb.PrepareRequest) (*pb.DataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	defer func() {
		memberUnLock(memberId)
	}()

	if !memberAddLock(memberId) {
		return &pb.DataReply{
			Code: 400,
			Msg:  "用户操作频繁，请稍后",
		}, nil
	}

	pId, count, err := dianping.Prepare(req.ShopId, memberId, req.ReceiptCode)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	_, err = s.uc.BuyCardValidate(ctx, &biz.CreateBuyCardOrderRequest{
		MemberId: memberId,
		CardId:   pId,
	})
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	_, err = dianping.Consume(req.ShopId, memberId, req.ReceiptCode, count)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	for i := 0; i < count; i++ {
		_, err = s.uc.CreateMemberPackageCard(ctx, &biz.CreateMemberPackageCardRequest{
			MemberId:     memberId,
			PackageId:    pId,
			TypeChannel:  3,
			ActiveAmount: 0,
			ReceiptCode:  req.ReceiptCode,
		})
		if err != nil {
			return &pb.DataReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
	}
	return &pb.DataReply{
		Code: 200,
		Msg:  "success",
	}, nil
}

func (s *DianpingService) ScanPrepare(ctx context.Context, req *pb.PrepareRequest) (*pb.DataReply, error) {
	req.ReceiptCode = req.QrCode
	return s.Prepare(ctx, req) //目前与输码验券校验接口
}

//func (s *DianpingService) ScanPrepare(ctx context.Context, req *pb.PrepareRequest) (*pb.DataReply, error) {
//	jwtCtx, _ := jwt.FromContext(ctx)
//	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
//
//	defer func() {
//		memberUnLock(memberId)
//	}()
//
//	if !memberAddLock(memberId) {
//		return &pb.DataReply{
//			Code: 400,
//			Msg:  "用户操作频繁，请稍后",
//		}, nil
//	}
//
//	pIds, count, err := dianping.ScanPrepare(req.ShopId, memberId, req.QrCode)
//	if err != nil {
//		return &pb.DataReply{
//			Code: 400,
//			Msg:  err.Error(),
//		}, nil
//
//	}
//
//	for _, pid := range pIds { //暂时只返回一组数据
//		_, err = s.uc.BuyCardValidate(ctx, &biz.CreateMemberPackageCardRequest{
//			MemberId:    memberId,
//			CardId:      pid,
//			TypeChannel: 3,
//			IsActive:    true,
//		})
//		if err != nil {
//			return &pb.DataReply{
//				Code: 400,
//				Msg:  err.Error(),
//			}, nil
//		}
//		for i := 0; i < count; i++ {
//			_, err := s.uc.CreateMemberPackageCard(ctx, &biz.CreateMemberPackageCardRequest{
//				MemberId:    memberId,
//				CardId:      pid,
//				TypeChannel: 3,
//			})
//			if err != nil {
//				return &pb.DataReply{
//					Code: 400,
//					Msg:  err.Error(),
//				}, nil
//			}
//		}
//	}
//
//	return &pb.DataReply{
//		Code: 200,
//		Msg:  "success",
//	}, nil
//}

func (s *DianpingService) AccessToken(ctx context.Context, req *pb.PrepareRequest) (*pb.DataReply, error) {
	err := dianping.CheckAuth("a87c1216347443cd")
	if err != nil {
		return nil, err
	}
	dianping.SaveAccessToken("a87c1216347443cd")
	if err != nil {
		return nil, err
	}
	return &pb.DataReply{
		Code: 200,
		Msg:  "success",
	}, nil
}
