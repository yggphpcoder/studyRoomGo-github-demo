package biz

import (
	"context"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
)

// MemberRepo is a Greater repo.
type MemberRepo interface {
	Save(context.Context, *Member) (bool, error)
	Update(context.Context, *Member) (bool, error)
	FindByID(context.Context, int64) (*gen.SrMember, error)
	FindByPhone(ctx context.Context, merchantId int64, phone string) (*gen.SrMember, error)
	ListAll(context.Context) ([]*gen.SrMember, error)

	UpdateWXUserInfo(context.Context, *WxUserInfo) (bool, error)
	CreateWXUserInfo(context.Context, *WxUserInfo) (bool, error)
	FindByWXOpenId(ctx context.Context, merchantId int64, openId string) (*gen.WxUserinfo, error)
	FindByWXUnionId(ctx context.Context, merchantId int64, unionId string) (*gen.WxUserinfo, error)
	FindByWXPhone(ctx context.Context, merchantId int64, phone string) (*gen.WxUserinfo, error)

	HasNotice(ctx context.Context, memberId int64, isRead *int) (bool, error)
	LastNotice(ctx context.Context, memberId int64, isRead *int, limit int32) ([]*gen.SrMemberNotice, error)
	ListNotice(ctx context.Context, memberId int64, Page int32, PageSize int32) ([]*gen.SrMemberNotice, error)
	UpdateNotice(ctx context.Context, id []int64, isRead int) (bool, error)
	SaveWXSubscribe(ctx context.Context, info *gen.WxNoticeSubscribe) (bool, error)
}

// PackageCardRepo is a Greater repo.
type PackageCardRepo interface {
	FindByID(context.Context, int64) (*gen.SrPackageCard, error)
	FindBy(ctx context.Context, key string, value string) (*gen.SrPackageCard, error)
	ListBy(context.Context, *SearchPackageCardRequest) ([]*gen.SrPackageCard, error)
	ListAll(context.Context) ([]*gen.SrPackageCard, error)

	RecordBuyLog(context.Context, *PackageCardMemberBuyLog) (bool, error)
	FindBuyLogBy(ctx context.Context, key string, value string) (*gen.SrPackageCardMemberBuyLog, error)
}

type ShopRepo interface {
	FindByID(context.Context, int64) (*gen.SrShop, error)
	FindBy(ctx context.Context, key string, value string) (*gen.SrShop, error)
	ListBy(context.Context, *SearchShopRequest) ([]*gen.SrShop, error)
	ListAll(context.Context) ([]*gen.SrShop, error)
	CanUseSeat(ctx context.Context, shopId int64) (list []*AreaSeatCount, total int32, using int32, err error)
	ListSetting(ctx context.Context, shopId int64) (list map[int64][]*model.Setting, err error)
	FindShopAreaReception(ctx context.Context, shopId int64) (*gen.SrShopArea, error)
}

// MemberPackageCardRepo is a Greater repo.
type MemberPackageCardRepo interface {
	Save(context.Context, *MemberPackageCard) (bool, error)
	Update(context.Context, *MemberPackageCard) (bool, error)
	FindByID(ctx context.Context, id int64) (*gen.SrMemberPackageCard, error)
	ListBy(ctx context.Context, request *SearchMemberPackageCardRequest) ([]*gen.SrMemberPackageCard, error)

	RecordOperateLog(ctx context.Context, log *MemberPackageCardOperateLog)
	RecordUseLog(ctx context.Context, log *MemberPackageCardUseLog)
	LastUseLog(ctx context.Context, memberId int64) (*MemberPackageCardUseLog, error)
}

type SeatRepo interface {
	FindByID(context.Context, int64) (*model.Seat, error)
	FindBy(ctx context.Context, key string, value string) (*model.Seat, error)
	ListBy(context.Context, *SearchSeatRequest) ([]*model.Seat, error)
	ListAll(context.Context) ([]*model.Seat, error)
	SeatTypeListAll(context.Context) ([]*gen.SrSeatType, error)

	ListBookingBy(context.Context, *SearchSeatBookingRequest) ([]*gen.SrSeatBook, error)
	ListLockingBy(context.Context, *SearchSeatLockingRequest) ([]*model.SeatLock, error)
	FindLockingBy(context.Context, *SearchSeatLockingRequest) (*model.SeatLock, error)
}

type SeatBookRepo interface {
	FindByID(context.Context, int64) (*model.SeatBook, error)
	FindDetailByID(context.Context, int64) (*model.SeatBook, error)

	ListBy(context.Context, *SearchSeatBookingRequest) ([]*model.SeatBook, error)

	Save(context.Context, *SeatBook) (bool, error)
	Update(context.Context, *SeatBook) (bool, error)
	CheckMemberCanBooking(context.Context, *BookingRequest) (*gen.SrSeatBook, error)
	CheckSeatCanBooking(context.Context, *BookingRequest) (*gen.SrSeatBook, error)
	CheckChildrenSeatCanBooking(context.Context, *BookingRequest) (*gen.SrSeatBook, error)
	RecordOperateLog(context.Context, *SeatBookOperateLog)
	ListOperateLog(ctx context.Context, seatBookId int64, status []int8) ([]*gen.SrSeatBookOperateLog, error)
	FindSharedByID(ctx context.Context, seatBookId int64, openId string) (res *gen.SrSeatBookShareLog, total int64, err error)
	SaveSharedLog(ctx context.Context, data *gen.SrSeatBookShareLog) (bool, error)
}

type ConfigRepo interface {
	UpdateByKey(ctx context.Context, shopId int64, key string, value string) (bool, error)

	FindDictByID(context.Context, int64) (*gen.SysDictData, error)
	FindDictBy(ctx context.Context, key string, value string) (*gen.SysDictData, error)
	ListDictBy(context.Context, *SearchDictRequest) ([]*gen.SysDictData, error)
	ListDictAll(context.Context) ([]*gen.SysDictData, error)

	FindSettingBy(ctx context.Context, shopId int64, key string) (*model.Setting, error)
	ListSettingBy(context.Context, *SearchSettingRequest) ([]*model.Setting, error)

	ListEdDictBy(context.Context, *SearchDictRequest) ([]*gen.EdDictData, error)
}

type OrderRepo interface {
	FindByID(context.Context, int64) (*gen.SrOrder, error)
	FindByOrderNo(context.Context, string) (*gen.SrOrder, error)

	Save(context.Context, *Order) (bool, error)
	Update(context.Context, *Order) (bool, error)
}

type ProfitSharingOrderRepo interface {
	FindByID(context.Context, int64) (*gen.MeProfitSharingWechatOrder, error)
	FindByOrderNo(context.Context, string) (*gen.MeProfitSharingWechatOrder, error)
	FindByTransactionId(context.Context, string) (*gen.MeProfitSharingWechatOrder, error)

	FindReceiversBy(ctx context.Context, shopId int64, merchantId int64, relateOrderType int8) (*gen.MeProfitSharingWechatReceivers, error)

	Save(context.Context, *gen.MeProfitSharingWechatOrder) (bool, error)
	Update(context.Context, *gen.MeProfitSharingWechatOrder) (bool, error)

	RecordLog(ctx context.Context, log *ProfitSharingWechatOrderLog) error
}
type PaymentRepo interface {
	FindByOrderNo(ctx context.Context, orderNo string) (*gen.SrPayment, error)
	FindByTransactionId(ctx context.Context, TransactionId string) (*gen.SrPayment, error)

	SavePayment(context.Context, *Payment) (bool, error)
	UpdatePayment(context.Context, *Payment) (bool, error)

	SaveWechatPayment(context.Context, *Payment) (bool, error)
	UpdateWechatPayment(context.Context, *Payment) (bool, error)
}

type EquipmentRepo interface {
	Save(context.Context, *gen.SrEquipment) (bool, error)
	UpdatesOperation(ctx context.Context, req *EquipmentRequest, operation int8) (bool, error)

	FindByID(ctx context.Context, id int64) (*gen.SrEquipment, error)
	FindByRelateId(context.Context, *EquipmentRequest) (*gen.SrEquipment, error)

	ListBy(context.Context, *EquipmentRequest) ([]*gen.SrEquipment, error)
}

type CouponRepo interface {
	FindByID(ctx context.Context, id int64) (*model.Coupons, error)
	FindByCode(ctx context.Context, code string) (*model.Coupons, error)
	ListBy(ctx context.Context, request *CouponCenterRequest) ([]*model.Coupons, error)

	ListReceiveBy(ctx context.Context, request *SearchCouponReceiveRequest) ([]*model.CouponsReceive, error)
	FindReceiveById(ctx context.Context, id int64) (*model.CouponsReceive, error)
	FindReceiveByOrderId(ctx context.Context, id int64) (*model.CouponsReceive, error)
	FindReceiveByCode(ctx context.Context, code string) (*model.CouponsReceive, error)
	SaveCouponReceive(ctx context.Context, couponReceive *model.CouponsReceive) (bool, error)
	UpdateCouponReceive(ctx context.Context, couponReceive *model.CouponsReceive) (bool, error)
	UpdateCouponReceiveExpired(ctx context.Context, ids []int64) (bool, error)

	RecordUseLog(ctx context.Context, log *CouponUseLog) error
}

// CourseRepo is a Greater repo.
type CourseRepo interface {
	FindByID(context.Context, int64) (*model.Course, error)
	FindBy(ctx context.Context, key string, value string) (*model.Course, error)
	ListBy(context.Context, *SearchCourseRequest) ([]*model.Course, error)
	ListAll(context.Context) ([]*model.Course, error)

	FindSaleByID(context.Context, int64) (*model.CourseSale, error)
}

type StudentRepo interface {
	Save(context.Context, *model.MyStudent) (bool, error)
	Update(context.Context, *model.MyStudent) (bool, error)
	FindByID(ctx context.Context, id int64) (*model.MyStudent, error)
	ListBy(ctx context.Context, request *SearchStudentRequest) ([]*model.MyStudent, error)
}

type StudentCourseRepo interface {
	Save(context.Context, *model.StudentCourse) (bool, error)
	Update(context.Context, *model.StudentCourse) (bool, error)
	FindByID(ctx context.Context, id int64) (*model.StudentCourse, error)
	ListBy(ctx context.Context, request *SearchStudentCourseRequest) ([]*model.StudentCourse, error)
}

type CourseSchedulingRepo interface {
	FindByID(ctx context.Context, id int64) (*gen.EdCoursesScheduling, error)
	FindDetailByID(ctx context.Context, id int64) (*model.CourseScheduling, error)
	ListBy(ctx context.Context, request *SearchSchedulingRequest) ([]*model.CourseScheduling, error)
	Update(context.Context, *gen.EdCoursesScheduling) (bool, error)
	RecordLog(oldData *gen.EdCoursesScheduling, data *gen.EdCoursesScheduling, by int) (bool, error)
	RecordCheckInLog(course StudentCourseCheckIn, by int) (bool, error)
}
type ClassSchedulingRepo interface {
	FindByID(ctx context.Context, id int64) (*gen.EdClassScheduling, error)
	ListBy(ctx context.Context, request *SearchSchedulingRequest) ([]*model.ClassScheduling, error)
	Update(context.Context, *gen.EdClassScheduling) (bool, error)
}
