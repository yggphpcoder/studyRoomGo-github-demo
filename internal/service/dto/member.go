package dto

import (
	pb "studyRoomGo/api/member/v1"
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"time"
)

type Member struct {
	Member          *biz.Member
	MyPackageCard   *biz.MyPackageCard
	MySeatBook      *biz.MySeatBook
	MySeatLock      *biz.MySeatLock
	HasNoReadNotice bool
}

func (r *Member) Reply() *pb.Data {
	avatar := ""
	nickName := ""
	phone := ""
	if r.Member.AvatarPath != nil {
		avatar = *r.Member.AvatarPath
	}
	if r.Member.NickName != nil {
		nickName = *r.Member.NickName
	}
	if r.Member.Phone != nil {
		phone = *r.Member.Phone
	}
	return &pb.Data{
		Id:              r.Member.ID,
		AvatarPath:      avatar,
		NickName:        nickName,
		Phone:           phone,
		Type:            int32(r.Member.Type),
		Status:          int32(r.Member.Status),
		NeedUpdate:      r.Member.NeedUpdate,
		HasNoReadNotice: r.HasNoReadNotice,
	}
}

func (r *Member) MyPackageCardReply() *pb.MyPackageCardData {
	return &pb.MyPackageCardData{
		Id:                    r.MyPackageCard.ID,
		MemberId:              r.MyPackageCard.MemberID,
		ActualAmount:          float32(r.MyPackageCard.ActualAmount),
		RemainMinute:          int32(r.MyPackageCard.RemainMinute),
		RemainDay:             int32(r.MyPackageCard.RemainDay),
		MinutePer:             int32(r.MyPackageCard.MinutePer),
		MinuteMax:             int32(r.MyPackageCard.MinuteMax),
		MinuteMin:             int32(r.MyPackageCard.MinuteMin),
		Status:                int32(r.MyPackageCard.Status),
		TypeChannel:           int32(r.MyPackageCard.TypeChannel),
		ActiveAt:              r.MyPackageCard.ActiveAt.Format(time.DateTime),
		InvalidAt:             r.MyPackageCard.InvalidAt.Format(time.DateTime),
		PackageId:             *r.MyPackageCard.PackageID,
		PackageName:           r.MyPackageCard.PackageName,
		TypePackage:           int32(r.MyPackageCard.TypePackage),
		TypeCard:              int32(r.MyPackageCard.TypeCard),
		TypeSale:              int32(r.MyPackageCard.TypeSale),
		Price:                 float32(r.MyPackageCard.Price),
		OriPrice:              float32(r.MyPackageCard.OriPrice),
		UseShop:               r.MyPackageCard.UseShopStr,
		UseSeat:               r.MyPackageCard.UseSeatStr,
		UseTime:               r.MyPackageCard.UseTime,
		UseBalance:            int32(*r.MyPackageCard.UseBalance),
		ValidDay:              int32(r.MyPackageCard.ValidDay),
		ActiveLimit:           int32(r.MyPackageCard.ActiveLimit),
		Sort:                  int32(r.MyPackageCard.Sort),
		Remark:                *r.MyPackageCard.Remark,
		CreatedAt:             r.MyPackageCard.CreatedAt.Format(time.DateTime),
		CanBookingToCloseTime: r.MyPackageCard.CanBookingToCloseTime,
		CanBookingToHour:      r.MyPackageCard.CanBookingToHour,
		UseShopId:             r.MyPackageCard.UseShopId,
	}
}

func (r *Member) MySeatBookReply() *pb.MySeatBookData {
	var todaySeatAt string
	var todayEndAt string
	if r.MySeatBook.TodaySeatAt != nil {
		todaySeatAt = r.MySeatBook.TodaySeatAt.Format(time.DateTime)
	}
	if r.MySeatBook.TodayEndAt != nil {
		todayEndAt = r.MySeatBook.TodayEndAt.Format(time.DateTime)
	}
	canCtrlLight := 1 //远程控制灯
	canShareLive := 0 //分享现场
	if _func.InIntSplice([]int{4, 5}, int(r.MySeatBook.ShopAreaID)) {
		canCtrlLight = 0
		canShareLive = 1
	}
	return &pb.MySeatBookData{
		Id:              r.MySeatBook.ID,
		ShopId:          r.MySeatBook.ShopID,
		ShopAreaId:      r.MySeatBook.ShopAreaID,
		SeatId:          r.MySeatBook.SeatID,
		MCardId:         r.MySeatBook.MemberPackageCardID,
		PackageType:     r.MySeatBook.PackageType,
		ShopLabel:       r.MySeatBook.ShopLabel,
		ShopAreaLabel:   r.MySeatBook.ShopAreaLabel,
		ShopAreaCover:   r.MySeatBook.ShopAreaCover,
		SeatLabel:       r.MySeatBook.SeatLabel,
		SeatTypeLabel:   r.MySeatBook.SeatTypeLabel,
		PackageLabel:    r.MySeatBook.PackageLabel,
		BookingType:     int32(r.MySeatBook.BookingType),
		BookingDay:      int32(r.MySeatBook.BookingDay),
		BookingMinute:   int32(r.MySeatBook.BookingMinute),
		BookingStartDay: r.MySeatBook.BookingStartDay.Format(time.DateOnly),
		BookingEndDay:   r.MySeatBook.BookingEndDay.Format(time.DateOnly),
		BookingStart:    r.MySeatBook.BookingStart,
		BookingEnd:      r.MySeatBook.BookingEnd,
		TodaySeatAt:     todaySeatAt,
		TodayEndAt:      todayEndAt,
		Status:          int32(r.MySeatBook.Status),
		CanShare:        int32(r.MySeatBook.Share),
		CanShareLive:    int32(canShareLive),
		SharedCount:     int32(r.MySeatBook.SharedCount),
		CanCtrlLight:    int32(canCtrlLight),
		LiveAddress:     r.MySeatBook.LiveAddress,
	}
}

func (r *Member) MySeatLockReply() *pb.MySeatLockData {
	return &pb.MySeatLockData{
		Id:                  r.MySeatLock.ID,
		MemberId:            r.MySeatLock.MemberID,
		ShopId:              r.MySeatLock.ShopID,
		ShopAreaId:          r.MySeatLock.ShopAreaID,
		SeatId:              r.MySeatLock.SeatID,
		LockingStartDay:     r.MySeatLock.LockStartDay.Format(time.DateOnly),
		LockingEndDay:       r.MySeatLock.LockEndDay.Format(time.DateOnly),
		LockingStart:        r.MySeatLock.LockStart,
		LockingEnd:          r.MySeatLock.LockEnd,
		UseTime:             r.MySeatLock.UseTime,
		Status:              int32(r.MySeatLock.Status),
		ShopLabel:           r.MySeatLock.ShopLabel,
		ShopAreaLabel:       r.MySeatLock.ShopAreaLabel,
		SeatLabel:           r.MySeatLock.SeatLabel,
		SeatTypeLabel:       r.MySeatLock.SeatTypeLabel,
		ShopAreaCover:       r.MySeatLock.ShopAreaCover,
		PackageLabel:        r.MySeatLock.PackageLabel,
		MemberPackageCardId: r.MySeatLock.MemberPackageCardID,
		CanControl:          int32(r.MySeatLock.CanControl),
	}
}

type Notice struct {
	Notice *gen.SrMemberNotice
}

func (r *Notice) Reply() *pb.MyNoticeData {
	return &pb.MyNoticeData{
		Id:         r.Notice.ID,
		Title:      r.Notice.Title,
		Content:    r.Notice.Content,
		NoticeType: int32(r.Notice.NoticeType),
		ExtJson:    r.Notice.ExtJSON,
		CreatedAt:  r.Notice.CreatedAt.Format(time.DateTime),
	}
}
