package gen

import "time"

// SrShop 店铺信息
type SrShop struct {
	ID            int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID    int64      `gorm:"column:merchant_id;type:bigint;not null;comment:'商家id'"`                      // 商家id
	Name          string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'店铺名'"`             // 店铺名
	Telephone     string     `gorm:"column:telephone;type:varchar(20);not null;default:'';comment:'联系电话'"`        // 联系电话
	Wechat        string     `gorm:"column:wechat;type:varchar(50);not null;default:'';comment:'微信号'"`            // 微信号
	Email         string     `gorm:"column:email;type:varchar(200);not null;default:'';comment:'邮箱'"`             // 邮箱
	Province      string     `gorm:"column:province;type:varchar(10);not null;default:'';comment:'省份'"`           // 省份
	City          string     `gorm:"column:city;type:varchar(10);not null;default:'';comment:'城市'"`               // 城市
	District      string     `gorm:"column:district;type:varchar(10);not null;default:'';comment:'区'"`            // 区
	Address       string     `gorm:"column:address;type:varchar(255);not null;default:'';comment:'店铺地址'"`         // 店铺地址
	Longitude     string     `gorm:"column:longitude;type:varchar(30);not null;default:'';comment:'经度'"`          // 经度
	Latitude      string     `gorm:"column:latitude;type:varchar(30);not null;default:'';comment:'纬度'"`           // 纬度
	LeaderID      int64      `gorm:"column:leader_id;type:bigint;not null;default:0;comment:'负责人id'"`             // 负责人id
	TypeShop      int8       `gorm:"column:type_shop;type:tinyint;not null;default:0;comment:'店铺类型'"`             // 店铺类型
	BusinessHours string     `gorm:"column:business_hours;type:varchar(50);not null;comment:'营业时间'"`              // 营业时间
	Sort          int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark        *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	Status        int8       `gorm:"column:status;type:tinyint;not null;default:1;comment:'状态'"`                  // 状态
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy      *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy      *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrShop) TableName() string {
	return "sr_shop"
}

// SrShopColumns get sql column name.获取数据库列名
var SrShopColumns = struct {
	ID            string
	MerchantID    string
	Name          string
	Telephone     string
	Wechat        string
	Email         string
	Province      string
	City          string
	District      string
	Address       string
	Longitude     string
	Latitude      string
	LeaderID      string
	TypeShop      string
	BusinessHours string
	Sort          string
	Remark        string
	Status        string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	CreateBy      string
	UpdateBy      string
}{
	ID:            "id",
	MerchantID:    "merchant_id",
	Name:          "name",
	Telephone:     "telephone",
	Wechat:        "wechat",
	Email:         "email",
	Province:      "province",
	City:          "city",
	District:      "district",
	Address:       "address",
	Longitude:     "longitude",
	Latitude:      "latitude",
	LeaderID:      "leader_id",
	TypeShop:      "type_shop",
	BusinessHours: "business_hours",
	Sort:          "sort",
	Remark:        "remark",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	CreateBy:      "create_by",
	UpdateBy:      "update_by",
}
