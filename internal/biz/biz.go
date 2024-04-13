package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewMemberUsecase,
	NewAuthUsecase,
	NewPackageCardUsecase,
	NewShopUsecase,
	NewSeatUsecase,
	NewConfigUsecase,
	NewPaymentUsecase,
	NewOrderUsecase,
	NewEquipmentUsecase,
	NewCouponUsecase,
	NewCourseUsecase,
)
