package gen

// SrCouponsReceiveGift 优惠券赠送表
type SrCouponsReceiveGift struct {
	ID               int64  `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'id'"`     // id
	CouponsReceiveID int64  `gorm:"column:coupons_receive_id;type:bigint;not null;default:0;comment:'优惠券领取表ID'"`   // 优惠券领取表ID
	GiftPackageCard  string `gorm:"column:gift_package_card;type:varchar(500);not null;default:'';comment:'赠送套餐'"` // 赠送套餐
	GiftCardPercent  int16  `gorm:"column:gift_card_percent;type:smallint;not null;default:0;comment:'赠送时长%'"`     // 赠送时长%
}

// TableName get sql table name.获取数据库表名
func (m *SrCouponsReceiveGift) TableName() string {
	return "sr_coupons_receive_gift"
}

// SrCouponsReceiveGiftColumns get sql column name.获取数据库列名
var SrCouponsReceiveGiftColumns = struct {
	ID               string
	CouponsReceiveID string
	GiftPackageCard  string
	GiftCardPercent  string
}{
	ID:               "id",
	CouponsReceiveID: "coupons_receive_id",
	GiftPackageCard:  "gift_package_card",
	GiftCardPercent:  "gift_card_percent",
}
