package gen

import "time"

// SrCoupons 优惠券总表
type SrCoupons struct {
	ID                int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'优惠券ID'"`      // 优惠券ID
	MerchantID        int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'创建商家id'"`                // 创建商家id
	ShopID            int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'创建店铺id'"`                    // 创建店铺id
	Code              string     `gorm:"unique;column:code;type:varchar(20);not null;default:'';comment:'优惠券代码'"`           // 优惠券代码
	Type              int8       `gorm:"column:type;type:tinyint;not null;default:0;comment:'1:即用型,2:领取型'"`                 // 1:即用型,2:领取型
	Name              string     `gorm:"column:name;type:varchar(100);not null;default:0;comment:'优惠券名称'"`                  // 优惠券名称
	Description       string     `gorm:"column:description;type:varchar(200);not null;default:'';comment:'优惠券描述'"`          // 优惠券描述
	DiscountType      int8       `gorm:"column:discount_type;type:tinyint;not null;default:0;comment:'优惠券类型：百分比折扣或固定金额折扣'"` // 优惠券类型：百分比折扣或固定金额折扣
	DiscountValue     float64    `gorm:"column:discount_value;type:decimal(10,2);not null;default:0.00;comment:'优惠券折扣值'"`   // 优惠券折扣值
	DiscountRule      string     `gorm:"column:discount_rule;type:varchar(1000);not null;default:'';comment:'优惠券折扣规则'"`     // 优惠券折扣规则
	MaxDiscount       float64    `gorm:"column:max_discount;type:decimal(10,2);not null;default:0.00;comment:'最大折扣金额'"`     // 最大折扣金额
	MinPurchase       float64    `gorm:"column:min_purchase;type:decimal(10,2);not null;default:0.00;comment:'最低购买金额限制'"`   // 最低购买金额限制
	StartDate         *time.Time `gorm:"column:start_date;type:datetime;comment:'优惠券有效开始日期'"`                               // 优惠券有效开始日期
	EndDate           *time.Time `gorm:"column:end_date;type:datetime;comment:'优惠券有效结束日期'"`                                 // 优惠券有效结束日期
	ValidDay          int        `gorm:"column:valid_day;type:int;not null;default:0;comment:'领取后有效天数'"`                    // 领取后有效天数
	Status            int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'优惠券状态'"`                     // 优惠券状态
	Quantity          int        `gorm:"column:quantity;type:int;not null;default:0;comment:'优惠券数量'"`                       // 优惠券数量
	UsageLimitPerUser int        `gorm:"column:usage_limit_per_user;type:int;not null;default:0;comment:'每个用户可使用的优惠券数量限制'"` // 每个用户可使用的优惠券数量限制
	UsageLimit        int        `gorm:"column:usage_limit;type:int;not null;default:0;comment:'优惠券总使用次数限制'"`               // 优惠券总使用次数限制
	IsReusable        int8       `gorm:"column:is_reusable;type:tinyint;not null;default:0;comment:'是否可重复使用'"`              // 是否可重复使用
	IsShow            int8       `gorm:"column:is_show;type:tinyint;not null;default:0;comment:'是否在前台显示'"`                  // 是否在前台显示
	Notes             string     `gorm:"column:notes;type:varchar(200);not null;default:'';comment:'备注'"`                   // 备注
	CreatedAt         *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt         *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy          *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy          *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrCoupons) TableName() string {
	return "sr_coupons"
}

// SrCouponsColumns get sql column name.获取数据库列名
var SrCouponsColumns = struct {
	ID                string
	MerchantID        string
	ShopID            string
	Code              string
	Type              string
	Name              string
	Description       string
	DiscountType      string
	DiscountValue     string
	DiscountRule      string
	MaxDiscount       string
	MinPurchase       string
	StartDate         string
	EndDate           string
	ValidDay          string
	Status            string
	Quantity          string
	UsageLimitPerUser string
	UsageLimit        string
	IsReusable        string
	IsShow            string
	Notes             string
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
	CreateBy          string
	UpdateBy          string
}{
	ID:                "id",
	MerchantID:        "merchant_id",
	ShopID:            "shop_id",
	Code:              "code",
	Type:              "type",
	Name:              "name",
	Description:       "description",
	DiscountType:      "discount_type",
	DiscountValue:     "discount_value",
	DiscountRule:      "discount_rule",
	MaxDiscount:       "max_discount",
	MinPurchase:       "min_purchase",
	StartDate:         "start_date",
	EndDate:           "end_date",
	ValidDay:          "valid_day",
	Status:            "status",
	Quantity:          "quantity",
	UsageLimitPerUser: "usage_limit_per_user",
	UsageLimit:        "usage_limit",
	IsReusable:        "is_reusable",
	IsShow:            "is_show",
	Notes:             "notes",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
	CreateBy:          "create_by",
	UpdateBy:          "update_by",
}
