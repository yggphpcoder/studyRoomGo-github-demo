package gen

import (
	"time"
)

// SrSeatLock 座位锁定
type SrSeatLock struct {
	ID                  int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	SeatID              int64      `gorm:"column:seat_id;type:bigint;not null;default:0;comment:'座位id'"`                // 座位id
	MemberID            int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`              // 会员id
	ShopID              int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	ShopAreaID          int64      `gorm:"column:shop_area_id;type:bigint;not null;default:0;comment:'店铺分区id'"`         // 店铺分区id
	LockStartDay        *time.Time `gorm:"column:lock_start_day;type:date;comment:'开始日期'"`                              // 开始日期
	LockStart           string     `gorm:"column:lock_start;type:time;not null;comment:'开始时间'"`                         // 开始时间
	LockEndDay          *time.Time `gorm:"column:lock_end_day;type:date;comment:'结束日期'"`                                // 结束日期
	LockEnd             string     `gorm:"column:lock_end;type:time;not null;comment:'结束时间'"`                           // 结束时间
	Status              int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'锁定状态'"`                // 锁定状态
	CreatedAt           *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt           *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy            *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy            *uint      `gorm:"column:update_by;type:int unsigned"`
	UseTime             string     `gorm:"column:use_time;type:varchar(1000);not null;default:'';comment:'使用时段'"`         // 使用时段
	MemberPackageCardID int64      `gorm:"column:member_package_card_id;type:bigint;not null;default:0;comment:'会员套餐id'"` // 会员套餐id
	MerchantID          int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`              // 商家id
	CanControl          int8       `gorm:"column:can_control;type:tinyint;not null;default:0;comment:'是否可以控制门灯'"`         // 是否可以控制门灯
}

// TableName get sql table name.获取数据库表名
func (m *SrSeatLock) TableName() string {
	return "sr_seat_lock"
}

// SrSeatLockColumns get sql column name.获取数据库列名
var SrSeatLockColumns = struct {
	ID                  string
	SeatID              string
	MemberID            string
	ShopID              string
	ShopAreaID          string
	LockStartDay        string
	LockStart           string
	LockEndDay          string
	LockEnd             string
	Status              string
	CreatedAt           string
	UpdatedAt           string
	DeletedAt           string
	CreateBy            string
	UpdateBy            string
	UseTime             string
	MemberPackageCardID string
	MerchantID          string
	CanControl          string
}{
	ID:                  "id",
	SeatID:              "seat_id",
	MemberID:            "member_id",
	ShopID:              "shop_id",
	ShopAreaID:          "shop_area_id",
	LockStartDay:        "lock_start_day",
	LockStart:           "lock_start",
	LockEndDay:          "lock_end_day",
	LockEnd:             "lock_end",
	Status:              "status",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
	CreateBy:            "create_by",
	UpdateBy:            "update_by",
	UseTime:             "use_time",
	MemberPackageCardID: "member_package_card_id",
	MerchantID:          "merchant_id",
	CanControl:          "can_control",
}
