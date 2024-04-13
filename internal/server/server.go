package server

import (
	admin "studyRoomGo/api/admin/v1"
	auth "studyRoomGo/api/auth/v1"
	config "studyRoomGo/api/config/v1"
	coupon "studyRoomGo/api/coupon/v1"
	course "studyRoomGo/api/course/v1"
	dianping "studyRoomGo/api/dianping/v1"
	equipment "studyRoomGo/api/equipment/v1"
	member "studyRoomGo/api/member/v1"
	order "studyRoomGo/api/order/v1"
	packageCard "studyRoomGo/api/packageCard/v1"
	payment "studyRoomGo/api/payment/v1"
	seat "studyRoomGo/api/seat/v1"
	shop "studyRoomGo/api/shop/v1"
	"studyRoomGo/internal/service"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

// RegisterRPCServer 注册服务到server
func RegisterRPCServer(srv *grpc.Server, s *service.Service) {
	member.RegisterMemberServer(srv, s.Member)
	packageCard.RegisterPackageCardServer(srv, s.PackageCard)
	auth.RegisterAuthServer(srv, s.Auth)
	shop.RegisterShopServer(srv, s.Shop)
	seat.RegisterSeatServer(srv, s.Seat)
	config.RegisterConfigServer(srv, s.Config)
	dianping.RegisterDianpingServer(srv, s.Dianping)
	order.RegisterOrderServer(srv, s.Order)
	payment.RegisterPaymentServer(srv, s.Payment)
	equipment.RegisterEquipmentServer(srv, s.Equipment)
	admin.RegisterAdminServer(srv, s.Admin)
	coupon.RegisterCouponServer(srv, s.Coupon)
	// wxOfficial.RegisterWxOfficialServer(srv, s.WxOfficial)
	course.RegisterCourseServer(srv, s.Course)

}

// RegisterHTTPServer 注册服务到server
func RegisterHTTPServer(srv *http.Server, s *service.Service) {
	member.RegisterMemberHTTPServer(srv, s.Member)
	packageCard.RegisterPackageCardHTTPServer(srv, s.PackageCard)
	auth.RegisterAuthHTTPServer(srv, s.Auth)
	shop.RegisterShopHTTPServer(srv, s.Shop)
	seat.RegisterSeatHTTPServer(srv, s.Seat)
	config.RegisterConfigHTTPServer(srv, s.Config)
	dianping.RegisterDianpingHTTPServer(srv, s.Dianping)
	order.RegisterOrderHTTPServer(srv, s.Order)
	payment.RegisterPaymentHTTPServer(srv, s.Payment)
	equipment.RegisterEquipmentHTTPServer(srv, s.Equipment)
	admin.RegisterAdminHTTPServer(srv, s.Admin)
	coupon.RegisterCouponHTTPServer(srv, s.Coupon)
	// wxOfficial.RegisterWxOfficialHTTPServer(srv, s.WxOfficial)
	course.RegisterCourseHTTPServer(srv, s.Course)

}

// jwt url 校验名单
var jwtUrlCheckList = map[string]bool{
	"/api.packageCard.v1.PackageCard/BuyCard":          true,
	"/api.packageCard.v1.PackageCard/MemberCardActive": true,
	"/api.seat.v1.Seat/Booking":                        true,
	"/api.seat.v1.Seat/BookingValidate":                true,
	"/api.seat.v1.Seat/BookingChangeSeat":              true,
	"/api.seat.v1.Seat/BookingAddTime":                 true,
	"/api.seat.v1.Seat/BookingChangeStatus":            true,
	"/api.seat.v1.Seat/BookingOpenDoor":                true,
	"/api.seat.v1.Seat/BookingControlLight":            true,
	"/api.seat.v1.Seat/BookingShareGenToken":           true,
	"/api.order.v1.Order/CalOrderTotalPrice":           true,
	"/api.payment.v1.Payment/WxPayCheckout":            true,
	"/api.payment.v1.Payment/WxPayCheckoutBuyCard":     true,
	"/api.payment.v1.Payment/WxPayCheckoutBuyCourse":   true,
	"/api.payment.v1.Payment/WxPayCheckoutReturn":      true,
	"/api.seat.v1.Seat/LockingOpenDoor":                true,
	"/api.seat.v1.Seat/LockingControlLight":            true,
	"/api.coupon.v1.Coupon/MyCoupon":                   true,
	"/api.coupon.v1.Coupon/UseCouponValidate":          true,
	"/api.coupon.v1.Coupon/ReceiveCoupon":              true,

	"/api.dianping.v1.Dianping/prepare":     true,
	"/api.dianping.v1.Dianping/scanPrepare": true,
}

// jwt api前缀 校验名单
var jwtApiPrefixCheckList = []string{
	"/api.member.v1",
	"/api.equipment.v1",
	"/api.course.v1",
}

// jwt api前缀 校验  排除名单
var jwtUrlIgnoreList = map[string]bool{
	"/api.member.v1.Member/MySeatBookShared": true,
	"/api.course.v1.Course/Search":           true,
	"/api.course.v1.Course/Get":              true,
}
