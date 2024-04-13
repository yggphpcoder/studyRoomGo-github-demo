package gen

import "time"

// SrMemberNotice 用户消息通知
type SrMemberNotice struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'id'"` // id
	MemberID   int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`            // 会员id
	MerchantID int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'创建商家id'"`        // 创建商家id
	ShopID     int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'创建店铺id'"`            // 创建店铺id
	NoticeType int8       `gorm:"column:notice_type;type:tinyint;not null;default:0;comment:'通知类型'"`         // 通知类型
	Title      string     `gorm:"column:title;type:varchar(100);not null;default:'';comment:'标题'"`           // 标题
	Content    string     `gorm:"column:content;type:varchar(1000);not null;default:'';comment:'内容'"`        // 内容
	IsRead     int8       `gorm:"column:is_read;type:tinyint;not null;default:0;comment:'是否已读'"`             // 是否已读
	ExtJSON    string     `gorm:"column:ext_json;type:varchar(1000);not null;default:'';comment:'扩展内容'"`     // 扩展内容
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrMemberNotice) TableName() string {
	return "sr_member_notice"
}

// SrMemberNoticeColumns get sql column name.获取数据库列名
var SrMemberNoticeColumns = struct {
	ID         string
	MemberID   string
	MerchantID string
	ShopID     string
	NoticeType string
	Title      string
	Content    string
	IsRead     string
	ExtJSON    string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	CreateBy   string
	UpdateBy   string
}{
	ID:         "id",
	MemberID:   "member_id",
	MerchantID: "merchant_id",
	ShopID:     "shop_id",
	NoticeType: "notice_type",
	Title:      "title",
	Content:    "content",
	IsRead:     "is_read",
	ExtJSON:    "ext_json",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
}
