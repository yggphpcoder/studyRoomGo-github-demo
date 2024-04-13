package dto

import (
	pb "studyRoomGo/api/packageCard/v1"
	"studyRoomGo/internal/biz"
)

type PackageCard struct {
	PackageCard *biz.PackageCard
}

func (r *PackageCard) Reply() *pb.Data {
	return &pb.Data{
		Id:               r.PackageCard.ID,
		Name:             r.PackageCard.Name,
		TypePackage:      int32(r.PackageCard.TypePackage),
		TypeCard:         int32(r.PackageCard.TypeCard),
		TypeSale:         int32(r.PackageCard.TypeSale),
		Status:           int32(r.PackageCard.Status),
		Price:            float32(r.PackageCard.Price),
		OriPrice:         float32(r.PackageCard.OriPrice),
		ValidDay:         int64(r.PackageCard.ValidDay),
		ActiveLimit:      int64(r.PackageCard.ActiveLimit),
		BuyLimit:         int64(r.PackageCard.BuyLimit),
		UseShop:          r.PackageCard.UseShopStr,
		UseSeat:          r.PackageCard.UseSeatStr,
		UseTime:          r.PackageCard.UseTime,
		BookingDay:       int64(r.PackageCard.BookingDay),
		BookingMinute:    int64(r.PackageCard.BookingMinute),
		BookingMinutePer: int64(r.PackageCard.BookingMinutePer),
		BookingMinuteMax: int64(r.PackageCard.BookingMinuteMax),
		Remark:           *r.PackageCard.Remark,
	}
}
