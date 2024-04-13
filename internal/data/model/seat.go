package model

import (
	"fmt"
	"gorm.io/gorm"
	"studyRoomGo/internal/data/gen"
)

type Seat struct {
	gen.SrSeat
	TypeSeatLabel string `json:"typeSeatLabel"`
	ShopAreaLabel string `json:"shopAreaLabel"`
	Share         int8   `json:"share"`
}

func (m *Seat) ScopesJoin(db *gorm.DB) *gorm.DB {
	db.Select("sr_seat.*,b.name AS type_seat_label,c.name AS shop_area_label,b.share as share")
	db.Joins("left join sr_seat_type as b on b.id = sr_seat.type_seat")
	db.Joins("left join sr_shop_area as c on c.id = sr_seat.shop_area_id")
	return db
}

// ScopesDateTime 时间 查询
func (m *Seat) ScopesDateTime(op string, field string, time2 string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if time2 != "" {
			addTime := time2
			db.Where(fmt.Sprintf("`%s`.`%s` %s ?", db.Statement.Table, field, op), addTime)
		}
		return db
	}
}
