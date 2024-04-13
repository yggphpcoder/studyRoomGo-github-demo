package gen

import "time"

// EdCoursesSchedulingAddress 教务-课程-排课-上课地址
type EdCoursesSchedulingAddress struct {
	ID          int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID  int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	ShopID      int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	AddressType int8       `gorm:"column:address_type;type:tinyint;not null;default:0;comment:'地址类型'"`          // 地址类型
	Name        string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'"`              // 名称
	Address     string     `gorm:"column:address;type:varchar(2000);not null;default:'';comment:'地址'"`          // 地址
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy    *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy    *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdCoursesSchedulingAddress) TableName() string {
	return "ed_courses_scheduling_address"
}

// EdCoursesSchedulingAddressColumns get sql column name.获取数据库列名
var EdCoursesSchedulingAddressColumns = struct {
	ID          string
	MerchantID  string
	ShopID      string
	AddressType string
	Name        string
	Address     string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	CreateBy    string
	UpdateBy    string
}{
	ID:          "id",
	MerchantID:  "merchant_id",
	ShopID:      "shop_id",
	AddressType: "address_type",
	Name:        "name",
	Address:     "address",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreateBy:    "create_by",
	UpdateBy:    "update_by",
}
