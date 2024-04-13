package gen

import "time"

// SrCouponsUseLog 优惠券使用表
type SrCouponsUseLog struct {
	ID               int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'id'"`   // id
	MemberID         int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`              // 会员id
	MerchantID       int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'创建商家id'"`          // 创建商家id
	ShopID           int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'创建店铺id'"`              // 创建店铺id
	OrderID          int64      `gorm:"column:order_id;type:bigint;not null;default:0;comment:'订单id'"`               // 订单id
	CouponID         int64      `gorm:"column:coupon_id;type:bigint;not null;default:0;comment:'优惠券ID'"`             // 优惠券ID
	CouponsReceiveID int64      `gorm:"column:coupons_receive_id;type:bigint;not null;default:0;comment:'优惠券领取表ID'"` // 优惠券领取表ID
	Action           int8       `gorm:"column:action;type:tinyint;not null;default:0;comment:'操作类型'"`                // 操作类型
	Extra            *string    `gorm:"column:extra;type:varchar(1000);comment:'额外信息'"`                              // 额外信息
	Remark           *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt        *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy         *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy         *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrCouponsUseLog) TableName() string {
	return "sr_coupons_use_log"
}

// SrCouponsUseLogColumns get sql column name.获取数据库列名
var SrCouponsUseLogColumns = struct {
	ID               string
	MemberID         string
	MerchantID       string
	ShopID           string
	OrderID          string
	CouponID         string
	CouponsReceiveID string
	Action           string
	Extra            string
	Remark           string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
	CreateBy         string
	UpdateBy         string
}{
	ID:               "id",
	MemberID:         "member_id",
	MerchantID:       "merchant_id",
	ShopID:           "shop_id",
	OrderID:          "order_id",
	CouponID:         "coupon_id",
	CouponsReceiveID: "coupons_receive_id",
	Action:           "action",
	Extra:            "extra",
	Remark:           "remark",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	CreateBy:         "create_by",
	UpdateBy:         "update_by",
}
