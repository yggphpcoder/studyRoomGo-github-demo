package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/service/dto"

	pb "studyRoomGo/api/config/v1"

	"github.com/go-kratos/kratos/v2/metadata"
)

type ConfigService struct {
	pb.UnimplementedConfigServer

	uc *biz.ConfigUsecase
}

func NewConfigService(uc *biz.ConfigUsecase) *ConfigService {
	return &ConfigService{uc: uc}
}

func (s *ConfigService) GetDict(ctx context.Context, req *pb.GetRequest) (*pb.DataDictReply, error) {
	data, err := s.uc.FindDictById(ctx, req.Id)
	if err != nil {
		return &pb.DataDictReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	dictDto := dto.Dict{Config: data}
	return &pb.DataDictReply{
		Code: 200,
		Data: dictDto.DictReply(),
	}, nil
}

func (s *ConfigService) FindDict(ctx context.Context, req *pb.FindRequest) (*pb.DataDictReply, error) {
	var filterKey = []string{
		"id", "name",
	}
	if !_func.InStrSplice(filterKey, req.Key) {
		return &pb.DataDictReply{
			Code: 400,
			Msg:  fmt.Sprintf("key not in %v", strings.Join(filterKey, ",")),
		}, nil
	}
	data, err := s.uc.FindDictBy(ctx, req.Key, req.Value)
	if err != nil {
		return &pb.DataDictReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	dictDto := dto.Dict{Config: data}
	return &pb.DataDictReply{
		Code: 200,
		Data: dictDto.DictReply(),
	}, nil
}

func (s *ConfigService) ListDict(ctx context.Context, req *pb.ListRequest) (*pb.ListDictReply, error) {

	data, err := s.uc.ListDict(ctx)
	if err != nil {
		return &pb.ListDictReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.DataDict
	for _, item := range data {
		dictDto := dto.Dict{Config: item}
		list = append(list, dictDto.DictReply())
	}
	return &pb.ListDictReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *ConfigService) SearchDict(ctx context.Context, req *pb.SearchDictRequest) (*pb.ListDictReply, error) {
	data, err := s.uc.SearchDict(ctx, &biz.SearchDictRequest{
		Type: req.Type,
	})
	if err != nil {
		return &pb.ListDictReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.DataDict
	for _, item := range data {
		dictDto := dto.Dict{Config: item}
		list = append(list, dictDto.DictReply())
	}
	return &pb.ListDictReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *ConfigService) SearchEdDict(ctx context.Context, req *pb.SearchDictRequest) (*pb.ListDictReply, error) {
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	data, err := s.uc.SearchEdDict(ctx, &biz.SearchDictRequest{
		Type:       req.Type,
		MerchantId: int64(merchantId),
		ShopId:     int64(shopId),
	})
	if err != nil {
		return &pb.ListDictReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.DataDict
	for _, item := range data {
		dictDto := dto.Dict{Config: item}
		list = append(list, dictDto.EdDictReply())
	}
	return &pb.ListDictReply{
		Code: 200,
		Data: list,
	}, nil
}
func (s *ConfigService) SearchSetting(ctx context.Context, req *pb.SearchSettingRequest) (*pb.ListSettingReply, error) {
	data, err := s.uc.SearchSetting(ctx, &biz.SearchSettingRequest{
		MerchantId: req.MerchantId,
		ShopId:     req.ShopId,
		Group:      req.Group,
		Key:        req.Key,
	})
	if err != nil {
		return &pb.ListSettingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.DataSetting
	for _, item := range data {
		settingDto := dto.Setting{Config: item}
		list = append(list, settingDto.SettingReply())
	}
	return &pb.ListSettingReply{
		Code: 200,
		Data: list,
	}, nil
}
