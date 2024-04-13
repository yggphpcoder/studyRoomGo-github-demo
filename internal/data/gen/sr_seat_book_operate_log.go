package gen

import "time"

// SrSeatBookOperateLog 订座操作记录
type SrSeatBookOperateLog struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	SeatBookID int64      `gorm:"column:seat_book_id;type:bigint;not null;default:0;comment:'订座id'"`           // 订座id
	MemberID   int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`              // 会员id
	SeatID     int64      `gorm:"column:seat_id;type:bigint;not null;default:0;comment:'座位id'"`                // 座位id
	ShopID     int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	Action     int8       `gorm:"column:action;type:tinyint;not null;default:0;comment:'操作类型'"`                // 操作类型
	Extra      *string    `gorm:"column:extra;type:varchar(1000);comment:'额外信息'"`                              // 额外信息
	Remark     *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrSeatBookOperateLog) TableName() string {
	return "sr_seat_book_operate_log"
}

// SrSeatBookOperateLogColumns get sql column name.获取数据库列名
var SrSeatBookOperateLogColumns = struct {
	ID         string
	SeatBookID string
	MemberID   string
	SeatID     string
	ShopID     string
	Action     string
	Extra      string
	Remark     string
	CreatedAt  string
	CreateBy   string
	DeletedAt  string
	UpdatedAt  string
	UpdateBy   string
}{
	ID:         "id",
	SeatBookID: "seat_book_id",
	MemberID:   "member_id",
	SeatID:     "seat_id",
	ShopID:     "shop_id",
	Action:     "action",
	Extra:      "extra",
	Remark:     "remark",
	CreatedAt:  "created_at",
	CreateBy:   "create_by",
	DeletedAt:  "deleted_at",
	UpdatedAt:  "updated_at",
	UpdateBy:   "update_by",
}
