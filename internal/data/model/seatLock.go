package model

import (
	"gorm.io/gorm"
	"studyRoomGo/internal/data/gen"
)

type SeatLock struct {
	gen.SrSeatLock
	ShopLabel     string `json:"shopLabel"`
	ShopAreaLabel string `json:"shopAreaLabel"`
	ShopAreaCover string `json:"shopAreaCover"`
	SeatLabel     string `json:"seatLabel"`
	SeatTypeLabel string `json:"seatTypeLabel"`
	PackageLabel  string `json:"packageLabel"`
}

func (m *SeatLock) ScopesJoin(db *gorm.DB) *gorm.DB {
	db.Select("sr_seat_lock.*,a.name AS shop_label,b.name AS shop_area_label,b.cover AS shop_area_cover,c.name AS seat_label,d.name AS seat_type_label,e.package_name AS package_label")
	db.Joins("left join sr_shop as a on a.id = sr_seat_lock.shop_id")
	db.Joins("left join sr_shop_area as b on b.id = sr_seat_lock.shop_area_id")
	db.Joins("left join sr_seat as c on c.id = sr_seat_lock.seat_id")
	db.Joins("left join sr_seat_type as d on d.id = c.type_seat")
	db.Joins("left join sr_member_package_card as e on e.id = sr_seat_lock.member_package_card_id")

	return db
}
