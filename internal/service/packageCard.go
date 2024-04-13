package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/service/dto"
	"time"

	pb "studyRoomGo/api/packageCard/v1"
)

type PackageCardService struct {
	pb.UnimplementedPackageCardServer

	uc *biz.PackageCardUsecase
}

func NewPackageCardService(uc *biz.PackageCardUsecase) *PackageCardService {
	return &PackageCardService{uc: uc}
}

func (s *PackageCardService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	return &pb.CreateReply{}, nil
}
func (s *PackageCardService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	return &pb.UpdateReply{}, nil
}
func (s *PackageCardService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return &pb.DeleteReply{}, nil
}
func (s *PackageCardService) Get(ctx context.Context, req *pb.GetRequest) (*pb.DataReply, error) {
	data, err := s.uc.FindById(ctx, req.Id)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	dto := dto.PackageCard{PackageCard: data}
	return &pb.DataReply{
		Code: 200,
		Data: dto.Reply(),
	}, nil
}

func (s *PackageCardService) Find(ctx context.Context, req *pb.FindRequest) (*pb.DataReply, error) {
	var filterKey = []string{
		"id", "name",
	}
	if !_func.InStrSplice(filterKey, req.Key) {
		return &pb.DataReply{
			Code: 400,
			Msg:  fmt.Sprintf("key not in %v", strings.Join(filterKey, ",")),
		}, nil
	}
	data, err := s.uc.FindBy(ctx, req.Key, req.Value)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	dto := dto.PackageCard{PackageCard: data}
	return &pb.DataReply{
		Code: 200,
		Data: dto.Reply(),
	}, nil
}

func (s *PackageCardService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {

	data, err := s.uc.List(ctx)
	if err != nil {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.Data
	for _, item := range data {
		dto := dto.PackageCard{PackageCard: item}
		list = append(list, dto.Reply())
	}
	return &pb.ListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *PackageCardService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.ListReply, error) {
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	data, err := s.uc.Search(ctx, &biz.SearchPackageCardRequest{
		MerchantId:  int64(merchantId),
		TypePackage: req.TypePackage,
		TypeCard:    req.TypeCard,
		TypeSale:    req.TypeSale,
		Status:      2,
		UseShop:     req.UseShop,
		UseSeat:     req.UseSeat,
		Sort:        int(req.Sort),
	})
	if err != nil {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.Data
	for _, item := range data {
		dto := dto.PackageCard{PackageCard: item}
		list = append(list, dto.Reply())
	}
	return &pb.ListReply{
		Code: 200,
		Data: list,
	}, nil
}

//func (s *PackageCardService) BuyCard(ctx context.Context, req *pb.CardRequest) (*pb.CardReply, error) {
//	jwtCtx, _ := jwt.FromContext(ctx)
//	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
//	_, err := s.uc.BuyCardValidate(ctx, &biz.CreateMemberPackageCardRequest{
//		MemberId:    memberId,
//		CardId:      req.CardId,
//		TypeChannel: 2,
//	})
//	if err != nil {
//		return &pb.CardReply{
//			Code: 400,
//			Msg:  err.Error(),
//		}, nil
//	}
//	_, err = s.uc.CreateMemberPackageCard(ctx, &biz.CreateMemberPackageCardRequest{
//		MemberId:    memberId,
//		CardId:      req.CardId,
//		TypeChannel: 2,
//	})
//	if err != nil {
//		return &pb.CardReply{
//			Code: 400,
//			Msg:  err.Error(),
//		}, nil
//	}
//
//	return &pb.CardReply{
//		Code: 200,
//	}, nil
//}

func (s *PackageCardService) MemberCardActive(ctx context.Context, req *pb.CardRequest) (*pb.CardReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	card, err := s.uc.FindMemberCardById(ctx, memberId, req.CardId)
	if err != nil {
		return &pb.CardReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	now := time.Now()
	activeAt := &now
	if req.ActiveDate != "" && req.ActiveTime != "" {
		t1, _ := time.ParseInLocation(time.DateTime, req.ActiveDate+" "+req.ActiveTime, time.Local)
		fmt.Println(t1.Year())
		if t1.Year() > 1 {
			activeAt = &t1
		}

	}
	if ok, err := s.uc.UpdateMemberPackageCardStatus(ctx, &biz.UpdateMemberPackageCardRequest{
		MpId:     card.ID,
		Status:   1,
		ActiveAt: activeAt,
	}); !ok {
		return &pb.CardReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.CardReply{
		Code: 200,
	}, nil
}

func (s *PackageCardService) FindBuyLog(ctx context.Context, req *pb.FindRequest) (*pb.BuyLogDataReply, error) {
	var filterKey = []string{
		"id", "order_no",
	}
	if !_func.InStrSplice(filterKey, req.Key) {
		return &pb.BuyLogDataReply{
			Code: 400,
			Msg:  fmt.Sprintf("key not in %v", strings.Join(filterKey, ",")),
		}, nil
	}
	data, err := s.uc.FindBuyLogBy(ctx, req.Key, req.Value)
	if err != nil {
		return &pb.BuyLogDataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.BuyLogDataReply{
		Code: 200,
		Data: &pb.BuyLogData{
			Id:            data.ID,
			OrderNo:       data.OrderNo,
			MemberId:      data.MemberID,
			PackageCardId: data.PackageCardID,
			MpId:          data.MpID,
		},
	}, nil
}
