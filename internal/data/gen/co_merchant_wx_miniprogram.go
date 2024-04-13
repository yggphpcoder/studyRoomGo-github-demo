package gen

import "time"

// CoMerchantWxMiniprogram 商家小程序信息
type CoMerchantWxMiniprogram struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`      // 主键编码
	MerchantID int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`                 // 商家id
	SignCode   string     `gorm:"column:sign_code;type:varchar(20);not null;default:'';comment:'商家标识码'"`            // 商家标识码
	AppID      string     `gorm:"column:app_id;type:varchar(20);not null;default:'';comment:'微信小程序app_id'"`         // 微信小程序app_id
	AppSecret  string     `gorm:"column:app_secret;type:varchar(50);not null;default:'';comment:'微信小程序app_secret'"` // 微信小程序app_secret
	Version    int8       `gorm:"column:version;type:tinyint;not null;comment:'合作版本'"`                              // 合作版本
	Status     int8       `gorm:"column:status;type:tinyint;not null;comment:'状态'"`                                 // 状态
	Sort       int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                             // 排序
	Remark     *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                     // 备注
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *CoMerchantWxMiniprogram) TableName() string {
	return "co_merchant_wx_miniprogram"
}

// CoMerchantWxMiniprogramColumns get sql column name.获取数据库列名
var CoMerchantWxMiniprogramColumns = struct {
	ID         string
	MerchantID string
	SignCode   string
	AppID      string
	AppSecret  string
	Version    string
	Status     string
	Sort       string
	Remark     string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	CreateBy   string
	UpdateBy   string
}{
	ID:         "id",
	MerchantID: "merchant_id",
	SignCode:   "sign_code",
	AppID:      "app_id",
	AppSecret:  "app_secret",
	Version:    "version",
	Status:     "status",
	Sort:       "sort",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
}
