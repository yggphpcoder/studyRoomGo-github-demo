package gen

import "time"

// SrShopAreaType 店铺分区类型信息
type SrShopAreaType struct {
	ID        int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	Name      string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'分区类型名'"`           // 分区类型名
	ShopID    string     `gorm:"column:shop_id;type:varchar(500);not null;default:'';comment:'适用店铺id列表'"`     // 适用店铺id列表
	Sort      int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark    *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy  *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy  *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrShopAreaType) TableName() string {
	return "sr_shop_area_type"
}

// SrShopAreaTypeColumns get sql column name.获取数据库列名
var SrShopAreaTypeColumns = struct {
	ID        string
	Name      string
	ShopID    string
	Sort      string
	Remark    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	CreateBy  string
	UpdateBy  string
}{
	ID:        "id",
	Name:      "name",
	ShopID:    "shop_id",
	Sort:      "sort",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
}
