package gen

import "time"

// SrShopArea 店铺分区信息
type SrShopArea struct {
	ID        int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	Name      string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'区域名'"`             // 区域名
	Cover     string     `gorm:"column:cover;type:varchar(50);not null;default:'';comment:'封面图'"`             // 封面图
	ShopID    int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	TypeArea  int8       `gorm:"column:type_area;type:tinyint;not null;default:0;comment:'区域类型'"`             // 区域类型
	Sort      int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark    *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy  *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy  *uint      `gorm:"column:update_by;type:int unsigned"`
	IsCharge  int8       `gorm:"column:is_charge;type:tinyint;not null;default:0;comment:'是否可设置收费'"` // 是否可设置收费
}

// TableName get sql table name.获取数据库表名
func (m *SrShopArea) TableName() string {
	return "sr_shop_area"
}

// SrShopAreaColumns get sql column name.获取数据库列名
var SrShopAreaColumns = struct {
	ID        string
	Name      string
	Cover     string
	ShopID    string
	TypeArea  string
	Sort      string
	Remark    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	CreateBy  string
	UpdateBy  string
	IsCharge  string
}{
	ID:        "id",
	Name:      "name",
	Cover:     "cover",
	ShopID:    "shop_id",
	TypeArea:  "type_area",
	Sort:      "sort",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	IsCharge:  "is_charge",
}
