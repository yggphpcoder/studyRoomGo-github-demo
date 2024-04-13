package gen

// SrCouponsReceiveApplicable 优惠券适用表
type SrCouponsReceiveApplicable struct {
	ID                        int64  `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'id'"`                     // id
	CouponsReceiveID          int64  `gorm:"column:coupons_receive_id;type:bigint;not null;default:0;comment:'优惠券领取表ID'"`                   // 优惠券领取表ID
	ApplicableOrderType       int8   `gorm:"column:applicable_order_type;type:tinyint;not null;default:0;comment:'适用订单类型'"`                 // 适用订单类型
	ApplicableShop            string `gorm:"column:applicable_shop;type:varchar(500);not null;default:'';comment:'适用店铺列表'"`                 // 适用店铺列表
	ApplicableShopType        string `gorm:"column:applicable_shop_type;type:varchar(500);not null;default:'';comment:'适用店铺类型列表'"`          // 适用店铺类型列表
	ApplicableSeat            string `gorm:"column:applicable_seat;type:varchar(500);not null;default:'';comment:'适用座位列表'"`                 // 适用座位列表
	ApplicableSeatType        string `gorm:"column:applicable_seat_type;type:varchar(500);not null;default:'';comment:'适用座位类型列表'"`          // 适用座位类型列表
	ApplicablePackageCard     string `gorm:"column:applicable_package_card;type:varchar(500);not null;default:'';comment:'适用套餐卡列表'"`        // 适用套餐卡列表
	ApplicablePackageCardType string `gorm:"column:applicable_package_card_type;type:varchar(500);not null;default:'';comment:'适用套餐卡类型列表'"` // 适用套餐卡类型列表
	ApplicableMemberType      string `gorm:"column:applicable_member_type;type:varchar(500);not null;default:'';comment:'适用用户类型列表'"`        // 适用用户类型列表
}

// TableName get sql table name.获取数据库表名
func (m *SrCouponsReceiveApplicable) TableName() string {
	return "sr_coupons_receive_applicable"
}

// SrCouponsReceiveApplicableColumns get sql column name.获取数据库列名
var SrCouponsReceiveApplicableColumns = struct {
	ID                        string
	CouponsReceiveID          string
	ApplicableOrderType       string
	ApplicableShop            string
	ApplicableShopType        string
	ApplicableSeat            string
	ApplicableSeatType        string
	ApplicablePackageCard     string
	ApplicablePackageCardType string
	ApplicableMemberType      string
}{
	ID:                        "id",
	CouponsReceiveID:          "coupons_receive_id",
	ApplicableOrderType:       "applicable_order_type",
	ApplicableShop:            "applicable_shop",
	ApplicableShopType:        "applicable_shop_type",
	ApplicableSeat:            "applicable_seat",
	ApplicableSeatType:        "applicable_seat_type",
	ApplicablePackageCard:     "applicable_package_card",
	ApplicablePackageCardType: "applicable_package_card_type",
	ApplicableMemberType:      "applicable_member_type",
}
