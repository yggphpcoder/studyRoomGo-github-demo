package service

import (
	"github.com/google/wire"
)

type Service struct {
	Member      *MemberService
	PackageCard *PackageCardService
	Auth        *AuthService
	Shop        *ShopService
	Seat        *SeatService
	Config      *ConfigService
	Dianping    *DianpingService
	Order       *OrderService
	Payment     *PaymentService
	Equipment   *EquipmentService
	Admin       *AdminService
	Coupon      *CouponService
	WxOfficial  *WxOfficialService
	Course      *CourseService
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(Service), "*"),
	NewMemberService,
	NewPackageCardService,
	NewAuthService,
	NewShopService,
	NewSeatService,
	NewConfigService,
	NewDianpingService,
	NewOrderService,
	NewPaymentService,
	NewEquipmentService,
	NewAdminService,
	NewCouponService,
	NewWxOfficialService,
	NewCourseService,
)

var seatLock = make(map[int64]bool)
var memberLock = make(map[int64]bool)
var transactionLock = make(map[string]bool)

func seatAddLock(seatId int64) bool {
	if seatId == 0 {
		return true
	}
	if seatLock[seatId] {
		return false
	}
	seatLock[seatId] = true
	//fmt.Println(seatLock)
	return true
}
func seatUnLock(seatId int64) {
	delete(seatLock, seatId)
	//fmt.Println(seatLock)

}

func memberAddLock(memberId int64) bool {
	if memberLock[memberId] {
		return false
	}
	memberLock[memberId] = true
	return true
}

func memberUnLock(memberId int64) {
	delete(memberLock, memberId)
}

func transactionAddLock(transactionId string) bool {
	if transactionLock[transactionId] {
		return false
	}
	transactionLock[transactionId] = true
	return true
}

func transactionUnLock(transactionId string) {
	delete(transactionLock, transactionId)
}
