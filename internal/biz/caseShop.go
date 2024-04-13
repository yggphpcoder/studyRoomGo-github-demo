package biz

import (
	"bufio"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"studyRoomGo/common/config"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// ShopUsecase is a Shop usecase.
type ShopUsecase struct {
	repo       ShopRepo
	repoConfig ConfigRepo
	log        *log.Helper
}

type Shop struct {
	*gen.SrShop
	Seat    *ShopSeat
	Setting []*model.Setting
}

type AreaSeatCount struct {
	ShopId     int64
	ShopAreaId int64
	Total      int32
}
type ShopSeat struct {
	Total int32
	Using int32
	List  []*AreaSeatCount
}

// NewShopUsecase new a Shop usecase.
func NewShopUsecase(repo ShopRepo, repoConfig ConfigRepo, logger log.Logger) *ShopUsecase {
	return &ShopUsecase{repo: repo, repoConfig: repoConfig, log: log.NewHelper(logger)}
}

func (uc *ShopUsecase) FindById(ctx context.Context, id int64) (*Shop, error) {
	data, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	shopAreaCount, total, using, err := uc.repo.CanUseSeat(ctx, id)
	if err != nil {
		return nil, err
	}
	shop := &Shop{SrShop: data}
	shop.Seat = &ShopSeat{}
	shop.Seat.Total = total
	shop.Seat.Using = using
	shop.Seat.List = shopAreaCount
	return shop, nil
}

func (uc *ShopUsecase) FindBy(ctx context.Context, key string, value string) (*Shop, error) {
	data, err := uc.repo.FindBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	return &Shop{SrShop: data}, nil
}

func (uc *ShopUsecase) List(ctx context.Context) ([]*Shop, error) {
	data, err := uc.repo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	var list []*Shop
	shopSetting, err := uc.repo.ListSetting(ctx, 0)
	for _, item := range data {

		shopAreaCount, total, using, err := uc.repo.CanUseSeat(ctx, item.ID)
		if err != nil {
			continue
		}
		shop := &Shop{
			SrShop: item,
			Seat: &ShopSeat{
				Total: total,
				Using: using,
				List:  shopAreaCount,
			},
			Setting: shopSetting[item.ID],
		}
		list = append(list, shop)
	}

	return list, nil
}

type SearchShopRequest struct {
	MerchantId int64
	Name       string
	Province   string
	City       string
	District   string
	Address    string
	TypeShop   int32
	Sort       int
	Status     int8
}

func (uc *ShopUsecase) Search(ctx context.Context, req *SearchShopRequest) ([]*Shop, error) {
	data, err := uc.repo.ListBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*Shop
	shopSetting, err := uc.repo.ListSetting(ctx, 0)
	for _, item := range data {

		shopAreaCount, total, using, err := uc.repo.CanUseSeat(ctx, item.ID)
		if err != nil {
			continue
		}
		shop := &Shop{
			SrShop: item,
			Seat: &ShopSeat{
				Total: total,
				Using: using,
				List:  shopAreaCount,
			},
			Setting: shopSetting[item.ID],
		}
		list = append(list, shop)
	}

	return list, nil
}
func (uc *ShopUsecase) ListShopSetting(ctx context.Context, id int64) (map[string]string, error) {
	data, err := uc.repoConfig.ListSettingBy(ctx, &SearchSettingRequest{ShopId: id})
	if err != nil {
		return nil, err
	}
	setting := make(map[string]string)
	for _, item := range data {
		setting[item.Key] = item.Value
	}
	return setting, nil
}

func (uc *ShopUsecase) UploadQrCode(ctx context.Context, shopId int64, qrCodeBase64 string) (string, error) {
	//去除包装，获取到base64编码
	imgBase64 := qrCodeBase64
	//base64转码
	imgs, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", errors.New(fmt.Sprint("base64 decode error:", err))
	}

	timenow := time.Now().Unix()
	file, err := os.OpenFile(config.AppConf.Upload.GetUri()+"/"+config.AppConf.Upload.Qrcode.GetDir()+"/"+strconv.FormatInt(timenow, 10)+".jpg", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", errors.New(fmt.Sprint("create file error:", err))
	}
	w := bufio.NewWriter(file) //创建新的 Writer 对象
	_, err3 := w.WriteString(string(imgs))
	if err3 != nil {
		return "", errors.New(fmt.Sprint("write error:", err3))
	}
	w.Flush()
	defer file.Close()
	imgname := strconv.FormatInt(timenow, 10) + ".jpg"
	//t := struct {
	//	ImageName string `json:"imagename"`
	//}{imgname}
	//this.Data["json"] = t
	//this.ServeJSON()
	path := config.AppConf.Upload.Qrcode.GetDir() + "/" + imgname
	ok, err := uc.repoConfig.UpdateByKey(ctx, shopId, "build_qrcode", path)
	if !ok {
		return "", err
	}
	return path, nil
}
