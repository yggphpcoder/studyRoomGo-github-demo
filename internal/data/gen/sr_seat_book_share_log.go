package gen

import "time"

// SrSeatBookShareLog 订座共享日志
type SrSeatBookShareLog struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	SeatBookID int64      `gorm:"column:seat_book_id;type:bigint;not null;default:0;comment:'订座id'"`           // 订座id
	Openid     string     `gorm:"column:openid;type:varchar(255);not null;comment:'微信用户唯一标识符'"`                // 微信用户唯一标识符
	MemberID   int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`              // 会员id
	Remark     *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrSeatBookShareLog) TableName() string {
	return "sr_seat_book_share_log"
}

// SrSeatBookShareLogColumns get sql column name.获取数据库列名
var SrSeatBookShareLogColumns = struct {
	ID         string
	SeatBookID string
	Openid     string
	MemberID   string
	Remark     string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	CreateBy   string
	UpdateBy   string
}{
	ID:         "id",
	SeatBookID: "seat_book_id",
	Openid:     "openid",
	MemberID:   "member_id",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
}
