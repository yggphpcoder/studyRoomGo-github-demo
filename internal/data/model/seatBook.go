package model

import (
	"gorm.io/gorm"
	"math"
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/data/gen"
	"time"
)

type SeatBook struct {
	gen.SrSeatBook
	ShopLabel     string  `json:"shopLabel"`
	ShopAreaLabel string  `json:"shopAreaLabel"`
	ShopAreaCover string  `json:"shopAreaCover"`
	SeatLabel     string  `json:"seatLabel"`
	SeatTypeLabel string  `json:"seatTypeLabel"`
	PackageLabel  string  `json:"packageLabel"`
	PackageType   int32   `json:"packageType"`
	Share         int8    `json:"share"`
	ChildIds      []int64 `json:"childIds" gorm:"-"`
	BookingWeek   []int32 `json:"bookingWeek" gorm:"-"`
}

func (m *SeatBook) ScopesJoin(db *gorm.DB) *gorm.DB {
	db.Select("sr_seat_book.*,a.name AS shop_label,b.name AS shop_area_label,b.cover AS shop_area_cover,c.name AS seat_label,d.name AS seat_type_label,e.package_name AS package_label,e.type_package AS package_type,d.share as share")
	db.Joins("left join sr_shop as a on a.id = sr_seat_book.shop_id")
	db.Joins("left join sr_shop_area as b on b.id = sr_seat_book.shop_area_id")
	db.Joins("left join sr_seat as c on c.id = sr_seat_book.seat_id")
	db.Joins("left join sr_seat_type as d on d.id = c.type_seat")
	db.Joins("left join sr_member_package_card as e on e.id = sr_seat_book.member_package_card_id")
	return db
}

// ScopesSeatCanBooking 查询座位检查是否可用预定
func (m *SeatBook) ScopesSeatCanBooking() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		startDay := m.BookingStartDay
		endDay := m.BookingEndDay
		start := m.BookingStart
		end := m.BookingEnd
		db.Where("seat_id = ? ", m.SeatID)
		db.Where("shop_id = ? ", m.ShopID)
		db.Where("member_id <> ? ", m.MemberID)
		db.Where("status not in (5,6,7) ")
		db.Where("(booking_start <= ? and booking_end > ?) or (booking_start >= ? and booking_start <= ?)", start, start, start, end)
		if len(m.BookingWeek) == 7 {
			db.Where("booking_start_day >= ? and booking_start_day <= ?", startDay, endDay) //目前booking_start_day，booking_end_day都是同一天
		} else {
			diffDay := int32(math.Ceil(endDay.Sub(startDay).Hours() / 24))
			var validDay []time.Time //有效日期数

			if diffDay >= 0 {
				for i := int32(0); i <= diffDay; i++ {
					nextDay := startDay.Add(time.Hour * 24 * time.Duration(i))
					if _func.InInt32Splice(m.BookingWeek, int32(nextDay.Weekday())) {
						validDay = append(validDay, nextDay)
					}
				}
				db.Where("booking_start_day in ? ", validDay)

			}
		}
		return db
	}
}

// ScopesMemberCanBooking 查询用户检查是否可用预定
func (m *SeatBook) ScopesMemberCanBooking() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		startDay := m.BookingStartDay
		endDay := m.BookingEndDay
		start := m.BookingStart
		end := m.BookingEnd
		db.Where("member_id = ? ", m.MemberID)
		db.Where("shop_id = ? ", m.ShopID)
		db.Where("status not in (5,6,7) ")
		db.Where("(booking_start <= ? and booking_end > ?) or (booking_start >= ? and booking_start <= ?)", start, start, start, end)
		if len(m.BookingWeek) == 7 {
			db.Where("booking_start_day >= ? and booking_start_day <= ?", startDay, endDay) //目前booking_start_day，booking_end_day都是同一天
		} else {
			diffDay := int32(math.Ceil(endDay.Sub(startDay).Hours() / 24))
			var validDay []time.Time //有效日期数

			if diffDay >= 0 {
				for i := int32(0); i <= diffDay; i++ {
					nextDay := startDay.Add(time.Hour * 24 * time.Duration(i))
					if _func.InInt32Splice(m.BookingWeek, int32(nextDay.Weekday())) {
						validDay = append(validDay, nextDay)
					}
				}
				db.Where("booking_start_day in ? ", validDay)

			}
		}

		return db
	}
}

// ScopesChildrenSeat 查询所以子座位订座状态
func (m *SeatBook) ScopesChildrenSeat() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		startDay := m.BookingStartDay
		endDay := m.BookingEndDay
		start := m.BookingStart
		end := m.BookingEnd
		db.Where("seat_id in ? ", m.ChildIds)
		db.Where("shop_id = ? ", m.ShopID)
		db.Where("status not in (5,6,7) ")
		db.Where("(booking_start <= ? and booking_end > ?) or (booking_start >= ? and booking_start <= ?)", start, start, start, end)
		db.Where("(booking_start_day <= ? and booking_end_day >= ?) or (booking_start_day >= ? and booking_start_day <= ?)", startDay, startDay, startDay, endDay)

		return db
	}
}
