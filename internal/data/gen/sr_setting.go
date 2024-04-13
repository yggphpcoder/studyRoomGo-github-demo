package gen

import "time"

// SrSetting 配置信息
type SrSetting struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	Name       string     `gorm:"column:name;type:varchar(50);not null;default:'';comment:'配置名'"`              // 配置名
	Key        string     `gorm:"column:key;type:varchar(50);not null;default:'';comment:'配置键'"`               // 配置键
	ValueType  int        `gorm:"column:value_type;type:int;not null;default:1;comment:'值类型'"`                 // 值类型
	Value      string     `gorm:"column:value;type:varchar(255);not null;default:'';comment:'配置值'"`            // 配置值
	Group      string     `gorm:"column:group;type:varchar(50);not null;default:'';comment:'配置组'"`             // 配置组
	MerchantID int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	ShopID     int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	Sort       int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark     *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrSetting) TableName() string {
	return "sr_setting"
}

// SrSettingColumns get sql column name.获取数据库列名
var SrSettingColumns = struct {
	ID         string
	Name       string
	Key        string
	ValueType  string
	Value      string
	Group      string
	MerchantID string
	ShopID     string
	Sort       string
	Remark     string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	CreateBy   string
	UpdateBy   string
}{
	ID:         "id",
	Name:       "name",
	Key:        "key",
	ValueType:  "value_type",
	Value:      "value",
	Group:      "group",
	MerchantID: "merchant_id",
	ShopID:     "shop_id",
	Sort:       "sort",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
}
