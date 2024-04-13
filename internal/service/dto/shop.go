package dto

import (
	pb "studyRoomGo/api/shop/v1"
	"studyRoomGo/internal/biz"
)

type Shop struct {
	*biz.Shop
}

func (r *Shop) Reply() *pb.Data {
	var setting []*pb.Data_Setting

	for _, s := range r.Setting {
		pbData := &pb.Data_Setting{
			Id:        s.SrSetting.ID,
			Key:       s.SrSetting.Key,
			Name:      s.SrSetting.Name,
			Value:     s.SrSetting.Value,
			ValueType: int32(s.SrSetting.ValueType),
		}
		if s.ValueType == 2 || s.ValueType == 3 {
			pbData.Value = s.SrSettingText.RichText
		}
		setting = append(setting, pbData)
	}
	var shopAreaCount []*pb.Data_ShopAreaCount

	for _, area := range r.Seat.List {
		pbShopAreaCount := &pb.Data_ShopAreaCount{
			ShopId:     area.ShopId,
			ShopAreaId: area.ShopAreaId,
			Total:      area.Total,
		}
		shopAreaCount = append(shopAreaCount, pbShopAreaCount)

	}

	return &pb.Data{
		Id:            r.ID,
		Name:          r.Name,
		Telephone:     r.Telephone,
		Wechat:        r.Wechat,
		Email:         r.Email,
		Province:      r.Province,
		City:          r.City,
		District:      r.District,
		Address:       r.Address,
		Longitude:     r.Longitude,
		Latitude:      r.Latitude,
		LeaderId:      r.LeaderID,
		TypeShop:      int32(r.TypeShop),
		BusinessHours: r.BusinessHours,
		Sort:          int64(r.Sort),
		Remark:        *r.Remark,
		Seat: &pb.Data_Seat{
			Total:         r.Seat.Total,
			Using:         r.Seat.Using,
			ShopAreaCount: shopAreaCount,
		},
		Setting: setting,
	}

}
