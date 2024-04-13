package biz

import (
	"context"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

// ConfigUsecase is a Config usecase.
type ConfigUsecase struct {
	repo ConfigRepo
	log  *log.Helper
}

type Config struct {
	Dict    *gen.SysDictData
	EdDict  *gen.EdDictData
	Setting *model.Setting
}

// NewConfigUsecase new a Config usecase.
func NewConfigUsecase(repo ConfigRepo, logger log.Logger) *ConfigUsecase {
	return &ConfigUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ConfigUsecase) FindDictById(ctx context.Context, id int64) (*Config, error) {
	data, err := uc.repo.FindDictByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Config{Dict: data}, nil
}

func (uc *ConfigUsecase) FindDictBy(ctx context.Context, key string, value string) (*Config, error) {
	data, err := uc.repo.FindDictBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	return &Config{Dict: data}, nil

}

func (uc *ConfigUsecase) ListDict(ctx context.Context) ([]*Config, error) {
	data, err := uc.repo.ListDictAll(ctx)
	if err != nil {
		return nil, err
	}
	var list []*Config
	for _, item := range data {
		list = append(list, &Config{Dict: item})
	}
	return list, nil
}

type SearchDictRequest struct {
	Type       string
	MerchantId int64
	ShopId     int64
}

func (uc *ConfigUsecase) SearchDict(ctx context.Context, req *SearchDictRequest) ([]*Config, error) {
	data, err := uc.repo.ListDictBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*Config
	for _, item := range data {
		list = append(list, &Config{Dict: item})
	}
	return list, nil
}

func (uc *ConfigUsecase) SearchEdDict(ctx context.Context, req *SearchDictRequest) ([]*Config, error) {
	data, err := uc.repo.ListEdDictBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*Config
	for _, item := range data {
		list = append(list, &Config{EdDict: item})
	}
	return list, nil
}

func (uc *ConfigUsecase) FindSettingBy(ctx context.Context, shopId int64, key string) (*Config, error) {
	data, err := uc.repo.FindSettingBy(ctx, shopId, key)
	if err != nil {
		return nil, err
	}
	return &Config{Setting: data}, nil
}

type SearchSettingRequest struct {
	MerchantId int64
	ShopId     int64
	Group      string
	Key        string
}

func (uc *ConfigUsecase) SearchSetting(ctx context.Context, req *SearchSettingRequest) ([]*Config, error) {
	data, err := uc.repo.ListSettingBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*Config
	for _, item := range data {
		list = append(list, &Config{Setting: item})
	}
	return list, nil
}
