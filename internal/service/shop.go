package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"strconv"
	"strings"
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/service/dto"

	pb "studyRoomGo/api/shop/v1"
)

type ShopService struct {
	pb.UnimplementedShopServer

	uc *biz.ShopUsecase
}

func NewShopService(uc *biz.ShopUsecase) *ShopService {
	return &ShopService{uc: uc}
}

func (s *ShopService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	return &pb.CreateReply{}, nil
}
func (s *ShopService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	return &pb.UpdateReply{}, nil
}
func (s *ShopService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return &pb.DeleteReply{}, nil
}
func (s *ShopService) Get(ctx context.Context, req *pb.GetRequest) (*pb.DataReply, error) {
	data, err := s.uc.FindById(ctx, req.Id)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	dto := dto.Shop{Shop: data}
	return &pb.DataReply{
		Code: 200,
		Data: dto.Reply(),
	}, nil
}

func (s *ShopService) Find(ctx context.Context, req *pb.FindRequest) (*pb.DataReply, error) {
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
	dto := dto.Shop{Shop: data}
	return &pb.DataReply{
		Code: 200,
		Data: dto.Reply(),
	}, nil
}
func (s *ShopService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.ListReply, error) {
	dto := dto.Shop{}
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	data, err := s.uc.Search(ctx, &biz.SearchShopRequest{
		MerchantId: int64(merchantId),
		Name:       req.Name,
		Province:   req.Province,
		City:       req.City,
		District:   req.District,
		Address:    req.Address,
		TypeShop:   req.TypeShop,
		Sort:       int(req.Sort),
		Status:     1,
	})
	if err != nil {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.Data
	for _, item := range data {
		dto.Shop = item
		list = append(list, dto.Reply())
	}
	return &pb.ListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *ShopService) UploadQrCode(ctx context.Context, req *pb.UploadQrCodeRequest) (*pb.QrCodeReply, error) {

	data, err := s.uc.UploadQrCode(ctx, req.ShopId, req.QrCodeBase64)
	if err != nil {
		return &pb.QrCodeReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.QrCodeReply{
		Code: 200,
		Msg:  "",
		Data: &pb.QrCodeReply_Data{QrCodePath: data},
	}, nil
}
