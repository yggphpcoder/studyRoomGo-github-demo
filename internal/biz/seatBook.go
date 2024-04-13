package biz

import (
	"studyRoomGo/internal/data/gen"
	"time"
)

const (
	_                              = iota
	SeatBook_OperateBooked         //已订座
	SeatBook_OperateWaiting        //等待入座
	SeatBook_OperateSeated         //入座
	SeatBook_OperateWaitNext       //连续订座，当天已结束
	SeatBook_OperateComplete       //取消
	SeatBook_OperateCustomerCancel //用户取消
	SeatBook_OperateSystemCancel   //结束
	SeatBook_OperateChangeSeat     //换座
	SeatBook_OperateAddTime        //换座
	SeatBook_LockingControlDoor    //换座

)

const (
	_                             = iota
	SeatBook_StatusWait           //待开始
	SeatBook_StatusNoSeat         //未入座
	SeatBook_StatusSeated         //入座
	SeatBook_StatusWaitNext       //连续订座，当天已结束
	SeatBook_StatusComplete       //已结束
	SeatBook_StatusCustomerCancel //用户已取消
	SeatBook_StatusSystemCancel   //系统已取消
)

type SeatBook struct {
	*gen.SrSeatBook
}

func (r *SeatBook) Generate(req *BookingRequest) {
	r.SrSeatBook = &gen.SrSeatBook{}
	r.SeatID = req.SeatId
	r.MemberID = req.MemberId
	r.MemberPackageCardID = req.MemberPackageCardId
	r.MerchantID = req.MerchantId
	r.ShopID = req.ShopId
	r.ShopAreaID = req.ShopAreaId
	r.BookingType = int8(req.BookingType)
	r.BookingDay = int(req.BookingDay)
	r.BookingMinute = int(req.BookingMinute)
	r.BookingStart = req.BookingStart.Format(time.TimeOnly)
	r.BookingEnd = req.BookingEnd.Format(time.TimeOnly)
	r.BookingStartDay = req.BookingStartDay
	r.BookingEndDay = req.BookingEndDay
	r.Status = SeatBook_StatusWait
	if r.BookingType == 1 { //立即入座
		r.Status = SeatBook_StatusNoSeat
	}
	if !req.CreateByAdmin {
		creatBy := uint(r.MemberID)
		r.CreateBy = &creatBy
	} else {
		creatBy := uint(0)
		r.CreateBy = &creatBy
	}

}

type SearchSeatBookingRequest struct {
	SeatId           int64
	ShopId           int64
	MerchantId       int64
	MemberId         int64
	Status           []int32
	GteBookDateEnd   string
	GteBookTimeEnd   string
	GteBookDateStart string
	GteBookTimeStart string
	LteBookDateEnd   string
	LteBookTimeEnd   string
	LteBookDateStart string
	LteBookTimeStart string

	Sort     int
	Page     int
	PageSize int
}
