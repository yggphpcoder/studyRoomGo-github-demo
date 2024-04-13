package gen

// SrCouponsGift 优惠券赠送表
type SrCouponsGift struct {
	CouponID        int64  `gorm:"primaryKey;column:coupon_id;type:bigint;not null;comment:'优惠券ID'"`              // 优惠券ID
	GiftPackageCard string `gorm:"column:gift_package_card;type:varchar(500);not null;default:'';comment:'赠送套餐'"` // 赠送套餐
	GiftCardPercent int16  `gorm:"column:gift_card_percent;type:smallint;not null;default:0;comment:'赠送时长%'"`     // 赠送时长%
}

// TableName get sql table name.获取数据库表名
func (m *SrCouponsGift) TableName() string {
	return "sr_coupons_gift"
}

// SrCouponsGiftColumns get sql column name.获取数据库列名
var SrCouponsGiftColumns = struct {
	CouponID        string
	GiftPackageCard string
	GiftCardPercent string
}{
	CouponID:        "coupon_id",
	GiftPackageCard: "gift_package_card",
	GiftCardPercent: "gift_card_percent",
}
