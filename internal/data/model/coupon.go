package model

import (
	"studyRoomGo/internal/data/gen"
)

type CouponsReceive struct {
	gen.SrCouponsReceive
	Applicable           gen.SrCouponsReceiveApplicable `json:"applicable" gorm:"foreignKey:CouponsReceiveID"`
	Gift                 gen.SrCouponsReceiveGift       `json:"gift" gorm:"foreignKey:CouponsReceiveID"`
	ShopLabel            []string                       `json:"shopLabel" gorm:"-"`
	PackageLabel         []string                       `json:"packageLabel" gorm:"-"`
	PackageTypeLabel     []string                       `json:"packageTypeLabel" gorm:"-"`
	GiftPackageCardLabel []string                       `json:"giftPackageCardLabel" gorm:"-"`
}

type Coupons struct {
	gen.SrCoupons
	Applicable           gen.SrCouponsApplicable `json:"applicable" gorm:"foreignKey:CouponID"`
	Gift                 gen.SrCouponsGift       `json:"gift" gorm:"foreignKey:CouponID"`
	ShopLabel            []string                `json:"shopLabel" gorm:"-"`
	PackageLabel         []string                `json:"packageLabel" gorm:"-"`
	PackageTypeLabel     []string                `json:"packageTypeLabel" gorm:"-"`
	GiftPackageCardLabel []string                `json:"giftPackageCardLabel" gorm:"-"`
}
