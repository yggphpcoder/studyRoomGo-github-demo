package gen

import "time"

// EdSaleman 教务-销售人员
type EdSaleman struct {
	ID          int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID  int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	ShopID      int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	SaleManType int8       `gorm:"column:sale_man_type;type:tinyint;not null;default:0;comment:'销售人员类型'"`       // 销售人员类型
	Name        string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'"`              // 名称
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy    *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy    *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdSaleman) TableName() string {
	return "ed_saleman"
}

// EdSalemanColumns get sql column name.获取数据库列名
var EdSalemanColumns = struct {
	ID          string
	MerchantID  string
	ShopID      string
	SaleManType string
	Name        string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	CreateBy    string
	UpdateBy    string
}{
	ID:          "id",
	MerchantID:  "merchant_id",
	ShopID:      "shop_id",
	SaleManType: "sale_man_type",
	Name:        "name",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreateBy:    "create_by",
	UpdateBy:    "update_by",
}
