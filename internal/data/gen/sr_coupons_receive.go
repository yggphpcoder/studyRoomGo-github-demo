package gen

import "time"

// SrCouponsReceive 优惠券领取表
type SrCouponsReceive struct {
	ID            int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'id'"`         // id
	Code          string     `gorm:"unique;column:code;type:varchar(20);not null;comment:'优惠券代码'"`                      // 优惠券代码
	CouponID      int64      `gorm:"column:coupon_id;type:bigint;not null;default:0;comment:'优惠券ID'"`                   // 优惠券ID
	MemberID      int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`                    // 会员id
	MerchantID    int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'创建商家id'"`                // 创建商家id
	ShopID        int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'创建店铺id'"`                    // 创建店铺id
	OrderID       int64      `gorm:"column:order_id;type:bigint;not null;default:0;comment:'订单id'"`                     // 订单id
	Name          string     `gorm:"column:name;type:varchar(100);not null;default:'';comment:'优惠券名称'"`                 // 优惠券名称
	Description   string     `gorm:"column:description;type:varchar(200);not null;default:'';comment:'优惠券描述'"`          // 优惠券描述
	DiscountType  int8       `gorm:"column:discount_type;type:tinyint;not null;default:0;comment:'优惠券类型：百分比折扣或固定金额折扣'"` // 优惠券类型：百分比折扣或固定金额折扣
	DiscountValue float64    `gorm:"column:discount_value;type:decimal(10,2);not null;default:0.00;comment:'优惠券折扣值'"`   // 优惠券折扣值
	DiscountRule  string     `gorm:"column:discount_rule;type:varchar(1000);not null;default:'';comment:'优惠券折扣规则'"`     // 优惠券折扣规则
	MaxDiscount   float64    `gorm:"column:max_discount;type:decimal(10,2);not null;default:0.00;comment:'最大折扣金额'"`     // 最大折扣金额
	MinPurchase   float64    `gorm:"column:min_purchase;type:decimal(10,2);not null;default:0.00;comment:'最低购买金额限制'"`   // 最低购买金额限制
	StartDate     *time.Time `gorm:"column:start_date;type:datetime;comment:'优惠券有效开始日期'"`                               // 优惠券有效开始日期
	EndDate       *time.Time `gorm:"column:end_date;type:datetime;comment:'优惠券有效结束日期'"`                                 // 优惠券有效结束日期
	UsageLimit    int        `gorm:"column:usage_limit;type:int;not null;default:0;comment:'优惠券使用次数限制'"`                // 优惠券使用次数限制
	Status        int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'优惠券状态'"`                     // 优惠券状态
	IsUsed        int8       `gorm:"column:is_used;type:tinyint;not null;default:0;comment:'是否已使用'"`                    // 是否已使用
	Notice        string     `gorm:"column:notice;type:varchar(500);not null;default:'';comment:'通知文案'"`                // 通知文案
	Notes         string     `gorm:"column:notes;type:varchar(200);not null;default:'';comment:'备注'"`                   // 备注
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy      *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy      *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrCouponsReceive) TableName() string {
	return "sr_coupons_receive"
}

// SrCouponsReceiveColumns get sql column name.获取数据库列名
var SrCouponsReceiveColumns = struct {
	ID            string
	Code          string
	CouponID      string
	MemberID      string
	MerchantID    string
	ShopID        string
	OrderID       string
	Name          string
	Description   string
	DiscountType  string
	DiscountValue string
	DiscountRule  string
	MaxDiscount   string
	MinPurchase   string
	StartDate     string
	EndDate       string
	UsageLimit    string
	Status        string
	IsUsed        string
	Notice        string
	Notes         string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	CreateBy      string
	UpdateBy      string
}{
	ID:            "id",
	Code:          "code",
	CouponID:      "coupon_id",
	MemberID:      "member_id",
	MerchantID:    "merchant_id",
	ShopID:        "shop_id",
	OrderID:       "order_id",
	Name:          "name",
	Description:   "description",
	DiscountType:  "discount_type",
	DiscountValue: "discount_value",
	DiscountRule:  "discount_rule",
	MaxDiscount:   "max_discount",
	MinPurchase:   "min_purchase",
	StartDate:     "start_date",
	EndDate:       "end_date",
	UsageLimit:    "usage_limit",
	Status:        "status",
	IsUsed:        "is_used",
	Notice:        "notice",
	Notes:         "notes",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	CreateBy:      "create_by",
	UpdateBy:      "update_by",
}
