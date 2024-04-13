package gen

import (
	"time"
)

// SrSeatBook 座位订座
type SrSeatBook struct {
	ID                  int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`   // 主键编码
	MerchantID          int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`              // 商家id
	UniqueToken         *string    `gorm:"column:unique_token;type:char(20);default:'';comment:'唯一标签，标注每次相同的订座操作'"`       // 唯一标签，标注每次相同的订座操作
	SeatID              int64      `gorm:"column:seat_id;type:bigint;not null;default:0;comment:'座位id'"`                  // 座位id
	MemberID            int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`                // 会员id
	MemberPackageCardID int64      `gorm:"column:member_package_card_id;type:bigint;not null;default:0;comment:'会员套餐id'"` // 会员套餐id
	ShopID              int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                  // 店铺id
	ShopAreaID          int64      `gorm:"column:shop_area_id;type:bigint;not null;default:0;comment:'店铺分区id'"`           // 店铺分区id
	BookingType         int8       `gorm:"column:booking_type;type:tinyint;not null;default:0;comment:'订座类型'"`            // 订座类型
	BookingDay          int        `gorm:"column:booking_day;type:int;not null;default:0;comment:'订座天数'"`                 // 订座天数
	BookingMinute       int        `gorm:"column:booking_minute;type:int;not null;default:0;comment:'订座分钟数'"`             // 订座分钟数
	BookingStartDay     time.Time  `gorm:"column:booking_start_day;type:date;not null;comment:'订座开始日期'"`                  // 订座开始日期
	BookingStart        string     `gorm:"column:booking_start;type:time;not null;comment:'订座开始时间'"`                      // 订座开始时间
	BookingEndDay       time.Time  `gorm:"column:booking_end_day;type:date;not null;comment:'预定结束日期'"`                    // 预定结束日期
	BookingEnd          string     `gorm:"column:booking_end;type:time;not null;comment:'订座结束时间'"`                        // 订座结束时间
	TodaySeatAt         *time.Time `gorm:"column:today_seat_at;type:timestamp;comment:'当天入座时间'"`                          // 当天入座时间
	TodayEndAt          *time.Time `gorm:"column:today_end_at;type:timestamp;comment:'当天结束时间'"`                           // 当天结束时间
	Status              int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'使用状态'"`                  // 使用状态
	Automation          string     `gorm:"column:automation;type:varchar(1000);not null;default:'';comment:'自动化操作'"`      // 自动化操作
	CreatedAt           *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt           *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy            *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy            *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrSeatBook) TableName() string {
	return "sr_seat_book"
}

// SrSeatBookColumns get sql column name.获取数据库列名
var SrSeatBookColumns = struct {
	ID                  string
	MerchantID          string
	UniqueToken         string
	SeatID              string
	MemberID            string
	MemberPackageCardID string
	ShopID              string
	ShopAreaID          string
	BookingType         string
	BookingDay          string
	BookingMinute       string
	BookingStartDay     string
	BookingStart        string
	BookingEndDay       string
	BookingEnd          string
	TodaySeatAt         string
	TodayEndAt          string
	Status              string
	Automation          string
	CreatedAt           string
	UpdatedAt           string
	DeletedAt           string
	CreateBy            string
	UpdateBy            string
}{
	ID:                  "id",
	MerchantID:          "merchant_id",
	UniqueToken:         "unique_token",
	SeatID:              "seat_id",
	MemberID:            "member_id",
	MemberPackageCardID: "member_package_card_id",
	ShopID:              "shop_id",
	ShopAreaID:          "shop_area_id",
	BookingType:         "booking_type",
	BookingDay:          "booking_day",
	BookingMinute:       "booking_minute",
	BookingStartDay:     "booking_start_day",
	BookingStart:        "booking_start",
	BookingEndDay:       "booking_end_day",
	BookingEnd:          "booking_end",
	TodaySeatAt:         "today_seat_at",
	TodayEndAt:          "today_end_at",
	Status:              "status",
	Automation:          "automation",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
	CreateBy:            "create_by",
	UpdateBy:            "update_by",
}
