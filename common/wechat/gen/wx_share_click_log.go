package gen

import "time"

// WxShareClickLog 分享点击日志
type WxShareClickLog struct {
	ID            int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`   // 主键编码
	ShareOpenID   string     `gorm:"column:share_open_id;type:varchar(255);not null;default:0;comment:'分享者openId'"` // 分享者openId
	ShareMemberID int64      `gorm:"column:share_member_id;type:bigint;not null;default:0;comment:'分享者会员id'"`       // 分享者会员id
	ShareType     string     `gorm:"column:share_type;type:varchar(20);not null;default:0;comment:'分享类型'"`          // 分享类型
	ShareRelateID int64      `gorm:"column:share_relate_id;type:bigint;not null;default:0;comment:'分享关联数据id'"`      // 分享关联数据id
	OpenID        string     `gorm:"column:open_id;type:varchar(255);not null;comment:'微信用户唯一标识符'"`                 // 微信用户唯一标识符
	MemberID      int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`                // 会员id
	Remark        *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                  // 备注
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy      *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy      *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *WxShareClickLog) TableName() string {
	return "wx_share_click_log"
}

// WxShareClickLogColumns get sql column name.获取数据库列名
var WxShareClickLogColumns = struct {
	ID            string
	ShareOpenID   string
	ShareMemberID string
	ShareType     string
	ShareRelateID string
	OpenID        string
	MemberID      string
	Remark        string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	CreateBy      string
	UpdateBy      string
}{
	ID:            "id",
	ShareOpenID:   "share_open_id",
	ShareMemberID: "share_member_id",
	ShareType:     "share_type",
	ShareRelateID: "share_relate_id",
	OpenID:        "open_id",
	MemberID:      "member_id",
	Remark:        "remark",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	CreateBy:      "create_by",
	UpdateBy:      "update_by",
}
