package gen

import "time"

// SrShopAreaSeatPrice 区域座位价格信息
type SrShopAreaSeatPrice struct {
	ID        int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	AreaID    int64      `gorm:"column:area_id;type:bigint;not null;default:0;comment:'区域id'"`                // 区域id
	ShopID    int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	Price     float64    `gorm:"column:price;type:float(10,4);not null;default:0.0000;comment:'售价'"`          // 售价
	UseTime   string     `gorm:"column:use_time;type:varchar(1000);not null;default:'';comment:'使用时段'"`       // 使用时段
	Sort      int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark    *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy  *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy  *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrShopAreaSeatPrice) TableName() string {
	return "sr_shop_area_seat_price"
}

// SrShopAreaSeatPriceColumns get sql column name.获取数据库列名
var SrShopAreaSeatPriceColumns = struct {
	ID        string
	AreaID    string
	ShopID    string
	Price     string
	UseTime   string
	Sort      string
	Remark    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	CreateBy  string
	UpdateBy  string
}{
	ID:        "id",
	AreaID:    "area_id",
	ShopID:    "shop_id",
	Price:     "price",
	UseTime:   "use_time",
	Sort:      "sort",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
}
