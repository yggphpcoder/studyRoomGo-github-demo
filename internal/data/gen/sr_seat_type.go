package gen

import "time"

// SrSeatType 座位类型信息
type SrSeatType struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	Name       string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'座位类型名'"`           // 座位类型名
	ShopID     string     `gorm:"column:shop_id;type:varchar(500);not null;default:'';comment:'适用店铺id列表'"`     // 适用店铺id列表
	ShopType   *string    `gorm:"column:shop_type;type:varchar(500);comment:'适用店铺类型'"`                         // 适用店铺类型
	Sort       int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark     *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrSeatType) TableName() string {
	return "sr_seat_type"
}

// SrSeatTypeColumns get sql column name.获取数据库列名
var SrSeatTypeColumns = struct {
	ID         string
	MerchantID string
	Name       string
	ShopID     string
	ShopType   string
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
	Name:       "name",
	ShopID:     "shop_id",
	ShopType:   "shop_type",
	Sort:       "sort",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
}
