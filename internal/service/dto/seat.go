package dto

import (
	"encoding/json"
	"math"
	pb "studyRoomGo/api/seat/v1"
	"studyRoomGo/common/config"
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/biz"
	"time"
)

type Seat struct {
	Seat         *biz.Seat
	TodayCanBook []*biz.TodayCanBookRange
}

func (r *Seat) Reply() *pb.Data {

	var bookList []*pb.Booking
	var lockList []*pb.Locking

	for _, item := range r.Seat.Booking {
		bookList = append(bookList, &pb.Booking{
			Id:              item.ID,
			MemberId:        item.MemberID,
			BookingType:     int32(item.BookingType),
			BookingDay:      int32(item.BookingDay),
			BookingMinute:   int32(item.BookingMinute),
			BookingStartDay: item.BookingStartDay.Format(time.DateOnly),
			BookingEndDay:   item.BookingEndDay.Format(time.DateOnly),
			BookingStart:    item.BookingStart,
			BookingEnd:      item.BookingEnd,
			Status:          int32(item.Status),
		})
	}
	for _, item := range r.Seat.Locking {
		lockList = append(lockList, &pb.Locking{
			//SeatId:     item.SeatID,
			MemberId: item.MemberID,
			//ShopId:     item.ShopID,
			//ShopAreaId: item.ShopAreaID,
			LockingStartDay: item.LockStartDay.Format(time.DateOnly),
			LockingEndDay:   item.LockEndDay.Format(time.DateOnly),
			LockingStart:    item.LockStart,
			LockingEnd:      item.LockEnd,
			UseTime:         item.UseTime,
			Status:          int32(item.Status),
		})
	}

	mapPos := ""
	if r.Seat.MapPosition != nil {
		mapPos = *r.Seat.MapPosition
	}
	return &pb.Data{
		Id:            r.Seat.ID,
		Pid:           r.Seat.Pid,
		Name:          r.Seat.Name,
		TypeSeat:      int32(r.Seat.TypeSeat),
		TypeSeatLabel: r.Seat.TypeSeatLabel,
		ShopId:        r.Seat.ShopID,
		ShopAreaId:    r.Seat.ShopAreaID,
		ShopAreaLabel: r.Seat.ShopAreaLabel,
		Status:        int32(r.Seat.Status),
		Sort:          int32(r.Seat.Sort),
		Remark:        "",
		MapPosition:   mapPos,
		Booking:       bookList,
		Locking:       lockList,
	}
}

func (r *Seat) OverviewReply() *pb.OverviewReply_Data {
	var bookList []*pb.Booking
	var lockList []*pb.Locking
	for _, item := range r.Seat.Booking {
		bookList = append(bookList, &pb.Booking{
			//SeatId:              item.SeatID,
			MemberId: item.MemberID,
			//MemberPackageCardId: item.MemberPackageCardID,
			//ShopId:              item.ShopID,
			//ShopAreaId:          item.ShopAreaID,
			BookingType:     int32(item.BookingType),
			BookingDay:      int32(item.BookingDay),
			BookingMinute:   int32(item.BookingMinute),
			BookingStartDay: item.BookingStartDay.Format(time.DateOnly),
			BookingEndDay:   item.BookingEndDay.Format(time.DateOnly),
			BookingStart:    item.BookingStart,
			BookingEnd:      item.BookingEnd,
			Status:          int32(item.Status),
		})
	}
	for _, item := range r.Seat.Locking {
		lockList = append(lockList, &pb.Locking{
			//SeatId:     item.SeatID,
			MemberId: item.MemberID,
			//ShopId:     item.ShopID,
			//ShopAreaId: item.ShopAreaID,
			LockingStartDay: item.LockStartDay.Format(time.DateOnly),
			LockingEndDay:   item.LockEndDay.Format(time.DateOnly),
			LockingStart:    item.LockStart,
			LockingEnd:      item.LockEnd,
			UseTime:         item.UseTime,
			Status:          int32(item.Status),
		})
	}
	mapPos := ""
	if r.Seat.MapPosition != nil {
		mapPos = *r.Seat.MapPosition
	}
	return &pb.OverviewReply_Data{
		Id:            r.Seat.ID,
		Pid:           r.Seat.Pid,
		Name:          r.Seat.Name,
		TypeSeat:      int32(r.Seat.TypeSeat),
		TypeSeatLabel: r.Seat.TypeSeatLabel,
		ShopId:        r.Seat.ShopID,
		ShopAreaId:    r.Seat.ShopAreaID,
		ShopAreaLabel: r.Seat.ShopAreaLabel,
		Status:        int32(r.Seat.Status),
		Sort:          int32(r.Seat.Sort),
		Remark:        "",
		MapPosition:   mapPos,
		Booking:       bookList,
		Locking:       lockList,
	}
}

type Booking struct {
}

func (r *Booking) BookingReq(memberId int64, s *pb.BookingRequest) *biz.BookingRequest {
	t1, _ := time.ParseInLocation(time.DateOnly, s.BookingStartDay, time.Local)
	t2, _ := time.ParseInLocation(time.DateOnly, s.BookingEndDay, time.Local)
	t3, _ := time.ParseInLocation(time.TimeOnly, s.BookingStart, time.Local)
	t4, _ := time.ParseInLocation(time.TimeOnly, s.BookingEnd, time.Local)
	bookMinute := s.BookingMinute
	if bookMinute == 0 { //兼容性
		bookMinute = s.BookingHours * 60
	}
	if s.BookingType == 1 { //	按时入座
		t := time.Now().Local().Format(time.TimeOnly)
		t3, _ = time.ParseInLocation(time.TimeOnly, t, time.Local)
		t4 = t3.Add(time.Minute * time.Duration(bookMinute))
	}
	if s.BookingToUseEndTime { //订座到适用时段结束时间
		t4, _ = time.ParseInLocation(time.TimeOnly, s.BookingEnd, time.Local)
	}
	if s.BookingToCloseTime { //订座到营业结束
		shopSetting, _ := config.GetShopSetting(s.ShopId)
		var closeTime string
		_ = json.Unmarshal([]byte(shopSetting["closeTime"]), &closeTime)

		t4, _ = time.ParseInLocation(time.TimeOnly, closeTime, time.Local)
	}
	if s.BookingWeek == nil {
		s.BookingWeek = []int32{0, 1, 2, 3, 4, 5, 6}
	}
	diffDay := int32(math.Ceil(t2.Sub(t1).Hours() / 24))
	validDay := 0 //有效日期数

	if diffDay >= 0 {
		for i := int32(0); i <= diffDay; i++ {
			nextDay := t1.Add(time.Hour * 24 * time.Duration(i))
			if _func.InInt32Splice(s.BookingWeek, int32(nextDay.Weekday())) {
				validDay++
			}
		}
	}
	return &biz.BookingRequest{
		MerchantId:          s.MerchantId,
		SeatId:              s.SeatId,
		MemberId:            memberId,
		MemberPackageCardId: s.MemberPackageCardId,
		ShopId:              s.ShopId,
		BookingType:         s.BookingType,
		BookingStartDay:     t1,
		BookingEndDay:       t2,
		BookingStart:        t3,
		BookingEnd:          t4,
		BookingToCloseTime:  s.BookingToCloseTime,
		BookingMinute:       bookMinute,
		BookingWeek:         s.BookingWeek,
		ValidDay:            int32(validDay),
		NoEnoughUse:         s.NoEnoughUse,
		ReceiptCode:         s.ReceiptCode,
		UseRemain:           s.UseRemain,
	}
}
func (r *Booking) BookingPreValidateReply(res *biz.BookingPreValidate) *pb.BookingValidateReply_Data {

	var cardlist []*pb.BookingValidateReply_PackageCard

	for _, card := range res.PackageCard {
		cardlist = append(cardlist, &pb.BookingValidateReply_PackageCard{
			Id:            card.Id,
			IsValid:       card.IsValid,
			InvalidCode:   card.InvalidCode,
			Reason:        card.Reason,
			BookingMinute: card.BookingMinute,
			BookingDay:    card.BookingDay,
		})
	}
	return &pb.BookingValidateReply_Data{
		IsValid:     res.IsValid,
		Reason:      res.Reason,
		PackageCard: cardlist,
	}
}
