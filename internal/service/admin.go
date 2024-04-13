package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"studyRoomGo/common/dianping"
	"studyRoomGo/common/equipment"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/service/dto"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	pb "studyRoomGo/api/admin/v1"
	seatPb "studyRoomGo/api/seat/v1"
)

type AdminService struct {
	pb.UnimplementedAdminServer

	seatUC      *biz.SeatUsecase
	equipmentUC *biz.EquipmentUsecase
	packageUC   *biz.PackageCardUsecase
	orderUC     *biz.OrderUsecase
	couponUC    *biz.CouponUsecase
	courseUC    *biz.CourseUsecase
}

func NewAdminService(seatUC *biz.SeatUsecase, equipmentUC *biz.EquipmentUsecase, packageUC *biz.PackageCardUsecase, orderUC *biz.OrderUsecase, couponUC *biz.CouponUsecase, courseUC *biz.CourseUsecase) *AdminService {
	return &AdminService{seatUC: seatUC, equipmentUC: equipmentUC, packageUC: packageUC, orderUC: orderUC, couponUC: couponUC, courseUC: courseUC}
}

func (s *AdminService) BookingValidate(ctx context.Context, req *pb.BookingRequest) (*pb.BookingValidateReply, error) {
	dto := dto.Booking{}

	bookReq := dto.BookingReq(req.MemberId, &seatPb.BookingRequest{
		SeatId:              req.SeatId,
		MemberPackageCardId: req.MemberPackageCardId,
		ShopId:              req.ShopId,
		BookingType:         req.BookingType,
		BookingStartDay:     req.BookingStartDay,
		BookingEndDay:       req.BookingEndDay,
		BookingStart:        req.BookingStart,
		BookingEnd:          req.BookingEnd,
		BookingToCloseTime:  req.BookingToCloseTime,
		BookingHours:        req.BookingHours,
		BookingMinute:       req.BookingMinute,
		BookingWeek:         req.BookingWeek,
		NoEnoughUse:         req.NoEnoughUse,
		ReceiptCode:         req.ReceiptCode,
		UseRemain:           req.UseRemain,
	})

	defer func() {
		seatUnLock(bookReq.SeatId)
		memberUnLock(bookReq.MemberId)
	}()

	if !seatAddLock(bookReq.SeatId) {
		return &pb.BookingValidateReply{
			Code: 400,
			Msg:  errors.New("座位操作频繁，请稍后").Error(),
		}, nil
	}

	if !memberAddLock(bookReq.MemberId) {
		return &pb.BookingValidateReply{
			Code: 400,
			Msg:  "用户操作频繁，请稍后",
		}, nil
	}

	res, err := s.seatUC.BookingPreValidate(ctx, bookReq)
	if err != nil {
		return &pb.BookingValidateReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

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
	data := &pb.BookingValidateReply_Data{
		IsValid:     res.IsValid,
		Reason:      res.Reason,
		PackageCard: cardlist,
	}

	return &pb.BookingValidateReply{
		Code: 200,
		Data: data,
	}, nil
}

func (s *AdminService) Booking(ctx context.Context, req *pb.BookingRequest) (*pb.BookingReply, error) {

	dto := dto.Booking{}
	bookReq := dto.BookingReq(req.MemberId, &seatPb.BookingRequest{
		MerchantId:          req.MerchantId,
		SeatId:              req.SeatId,
		MemberPackageCardId: req.MemberPackageCardId,
		ShopId:              req.ShopId,
		BookingType:         req.BookingType,
		BookingStartDay:     req.BookingStartDay,
		BookingEndDay:       req.BookingEndDay,
		BookingStart:        req.BookingStart,
		BookingEnd:          req.BookingEnd,
		BookingToCloseTime:  req.BookingToCloseTime,
		BookingHours:        req.BookingHours,
		BookingMinute:       req.BookingMinute,
		BookingWeek:         req.BookingWeek,
		NoEnoughUse:         req.NoEnoughUse,
		ReceiptCode:         req.ReceiptCode,
		UseRemain:           req.UseRemain,
	})
	bookReq.CreateByAdmin = true //管理员订座
	defer func() {
		seatUnLock(bookReq.SeatId)
		memberUnLock(bookReq.MemberId)
	}()

	if !seatAddLock(bookReq.SeatId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  errors.New("座位操作频繁，请稍后").Error(),
		}, nil
	}

	if !memberAddLock(bookReq.MemberId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "用户操作频繁，请稍后",
		}, nil
	}
	if ok, err := s.seatUC.BookingValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: err.GetCode(),
			Msg:  err.GetReason(),
		}, nil
	}

	if ok, err := s.seatUC.Booking(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	notSeatMinCancel := biz.GetNotSeatMinCancel(bookReq.ShopId)

	wx, _ := wechat.NewMessage(req.MerchantId)
	_ = wx.SendSeatBookMessage(bookReq.MemberId, wechat.SeatBookMessage{
		Desc:      bookReq.ShopLabel + " - " + bookReq.SeatLabel,
		StartDate: bookReq.BookingStartDay.Format(time.DateOnly),
		TimeRange: bookReq.BookingStart.Format("15:04") + "~" + bookReq.BookingEnd.Format("15:04"),
		Tip:       fmt.Sprintf("到期未入座,座位保留%d分钟", notSeatMinCancel),
	})
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) BookingAddTime(ctx context.Context, req *pb.BookingAddTimeRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	book, err := s.seatUC.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	t3, _ := time.ParseInLocation(time.TimeOnly, book.BookingEnd, time.Local)
	t4, _ := time.ParseInLocation(time.TimeOnly, book.BookingEnd, time.Local)

	bookReq := &biz.BookingRequest{
		SeatId:              book.SeatID,
		MemberId:            memberId,
		MemberPackageCardId: book.MemberPackageCardID,
		ShopId:              book.ShopID,
		ShopAreaId:          book.ShopAreaID,
		BookingType:         int32(book.BookingType),
		BookingDay:          int32(book.BookingDay),
		BookingMinute:       int32(book.BookingMinute),
		BookingStartDay:     book.BookingStartDay,
		BookingEndDay:       book.BookingEndDay,
		BookingStart:        t3,
		BookingEnd:          t4,
		SeatBookId:          req.BookId,
		Status:              book.Status,
		AddTimeToCloseTime:  req.BookingToCloseTime,
		AddTimeMin:          req.BookingMin,
		CreateByAdmin:       true, //管理员操作
		BookingWeek:         []int32{0, 1, 2, 3, 4, 5, 6},
		ValidDay:            1,
	}

	defer func() {
		seatUnLock(bookReq.SeatId)
		memberUnLock(bookReq.MemberId)
	}()

	if !seatAddLock(bookReq.SeatId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  errors.New("座位操作频繁，请稍后").Error(),
		}, nil
	}

	if !memberAddLock(bookReq.MemberId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "用户操作频繁，请稍后",
		}, nil
	}
	if ok, err := s.seatUC.BookingAddTimeValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if ok, err := s.seatUC.BookingValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.GetReason(),
		}, nil
	}

	if ok, err := s.seatUC.UpdateBooking(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) BookingChangeSeat(ctx context.Context, req *pb.BookingChangeSeatRequest) (*pb.BookingReply, error) {
	book, err := s.seatUC.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	t3, _ := time.ParseInLocation(time.TimeOnly, book.BookingStart, time.Local)
	t4, _ := time.ParseInLocation(time.TimeOnly, book.BookingEnd, time.Local)

	bookReq := &biz.BookingRequest{
		SeatId:              book.SeatID,
		MemberId:            book.MemberID,
		MemberPackageCardId: book.MemberPackageCardID,
		ShopId:              book.ShopID,
		ShopAreaId:          book.ShopAreaID,
		BookingType:         int32(book.BookingType),
		BookingDay:          int32(book.BookingDay),
		BookingMinute:       int32(book.BookingMinute),
		BookingStartDay:     book.BookingStartDay,
		BookingEndDay:       book.BookingEndDay,
		BookingStart:        t3,
		BookingEnd:          t4,
		SeatBookId:          req.BookId,
		ChangeSeatId:        req.ChangeSeatId,
		Status:              book.Status,
		CreateByAdmin:       true, //管理员操作
		BookingWeek:         []int32{0, 1, 2, 3, 4, 5, 6},
		ValidDay:            1,
	}

	defer func() {
		seatUnLock(bookReq.SeatId)
		seatUnLock(bookReq.ChangeSeatId)
		memberUnLock(bookReq.MemberId)
	}()

	if !seatAddLock(bookReq.SeatId) || !seatAddLock(bookReq.ChangeSeatId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  errors.New("座位操作频繁，请稍后").Error(),
		}, nil
	}

	if !memberAddLock(bookReq.MemberId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "用户操作频繁，请稍后",
		}, nil
	}
	if ok, err := s.seatUC.BookingChangeSeatValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if ok, err := s.seatUC.BookingValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.GetReason(),
		}, nil
	}

	if ok, err := s.seatUC.UpdateBooking(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	//断电
	if ok, err := s.SeatAutomate(ctx, bookReq.SeatId, false); !ok {
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
	}
	//通电
	go func(ChangeSeatId int64) {
		<-time.After(time.Second * 2)
		_, _ = s.SeatAutomate(ctx, ChangeSeatId, true)
	}(bookReq.ChangeSeatId)

	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) BookingChangeStatus(ctx context.Context, req *pb.BookingChangeStatusRequest) (*pb.BookingReply, error) {
	book, err := s.seatUC.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	t3, _ := time.ParseInLocation(time.TimeOnly, book.BookingStart, time.Local)
	t4, _ := time.ParseInLocation(time.TimeOnly, book.BookingEnd, time.Local)

	bookReq := &biz.BookingRequest{
		SeatId:              book.SeatID,
		MemberId:            book.MemberID,
		MemberPackageCardId: book.MemberPackageCardID,
		ShopId:              book.ShopID,
		ShopAreaId:          book.ShopAreaID,
		BookingType:         int32(book.BookingType),
		BookingDay:          int32(book.BookingDay),
		BookingMinute:       int32(book.BookingMinute),
		BookingStartDay:     book.BookingStartDay,
		BookingEndDay:       book.BookingEndDay,
		BookingStart:        t3,
		BookingEnd:          t4,
		SeatBookId:          req.BookId,
		Status:              book.Status,
		NewStatus:           int8(req.ChangeStatus),
		CreateByAdmin:       true,
	}

	defer func() {
		seatUnLock(bookReq.SeatId)
		memberUnLock(bookReq.MemberId)
	}()

	if !seatAddLock(bookReq.SeatId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  errors.New("座位操作频繁，请稍后").Error(),
		}, nil
	}

	if !memberAddLock(bookReq.MemberId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "用户操作频繁，请稍后",
		}, nil
	}

	if ok, err := s.seatUC.BookingChangeStatusValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	if ok, err := s.seatUC.UpdateBooking(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	if bookReq.NewStatus == biz.SeatBook_StatusSeated { //入座,通电
		if ok, _ := s.SeatAutomate(ctx, bookReq.SeatId, true); !ok {
			//return &pb.BookingReply{
			//	Code: 400,
			//	Msg:  err.Error(),
			//}, nil
		}
	}

	if bookReq.NewStatus == biz.SeatBook_StatusComplete { //结束,断电
		if ok, _ := s.SeatAutomate(ctx, bookReq.SeatId, false); !ok {
			//return &pb.BookingReply{
			//	Code: 400,
			//	Msg:  err.Error(),
			//}, nil
		}
	}
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) BookingControlLight(ctx context.Context, req *pb.BookingControlSeatRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	_ = int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	book, err := s.seatUC.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if ok, err := s.seatUC.BookingControlLightValidate(ctx, book); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	eqData, err := s.equipmentUC.FindByRelateId(ctx, &biz.EquipmentRequest{
		TypeRelate: 1,
		RelateId:   book.SeatID,
	})
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "灯控失败，请手动重试或联系管理员",
		}, nil
	}
	operation := false
	if eqData.Operation == 0 {
		operation = true
	}
	if ok, _ := s.SeatAutomate(ctx, book.SeatID, operation); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "灯控失败，请手动重试或联系管理员",
		}, nil
	}

	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) BookingOpenDoor(ctx context.Context, req *pb.BookingOpenDoorRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	_ = int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	book, err := s.seatUC.FindBookingById(ctx, req.BookId)
	if err != nil {
		return nil, err
	}

	var eq *gen.SrEquipment
	if req.Door == 1 {
		if ok, err := s.seatUC.BookingOpenBigDoorValidate(ctx, book); !ok {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		area, err := s.seatUC.FindShopAreaReception(ctx, book.ShopID)
		if err != nil {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		eq, err = s.equipmentUC.FindByRelateId(ctx, &biz.EquipmentRequest{
			TypeRelate: 2,
			RelateId:   area.ID,
			ShoId:      book.ShopID,
		})
		if err != nil {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
	} else {

		if ok, err := s.seatUC.BookingOpenDoorValidate(ctx, book); !ok {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}

		eq, err = s.equipmentUC.FindByRelateId(ctx, &biz.EquipmentRequest{
			TypeRelate: 2,
			RelateId:   book.ShopAreaID,
		})
		if err != nil {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}

	}

	if book.ShopID == 2 {
		if ok, err := equipment.Servers["accesscontrol"].Automate(eq.Mac, eq.Code, 1); !ok {

			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
	} else {
		if ok, err := equipment.Servers["smyoo"].Automate(eq.Mac, eq.Code, 1); !ok {

			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
	}
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) SeatAutomate(ctx context.Context, seatId int64, status bool) (ok bool, err error) {
	data, err := s.equipmentUC.FindByRelateId(ctx, &biz.EquipmentRequest{
		TypeRelate: 1,
		RelateId:   seatId,
	})
	if err != nil {
		return false, err
	}

	fmt.Printf("AdminService_SeatAutomate %d status %v", seatId, status)
	if status {
		_, err = s.equipmentUC.Automate(ctx, data.ID, 1)
	} else {
		_, err = s.equipmentUC.Automate(ctx, data.ID, 0)

	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *AdminService) DianpingRefreshShop(ctx context.Context, req *pb.DianpingRequest) (reply *pb.DianpingReply, err error) {
	err = dianping.CreateShopList(req.Appkey)
	if err != nil {
		return &pb.DianpingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.DianpingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) DianpingRefreshTuanGou(ctx context.Context, req *pb.DianpingRequest) (reply *pb.DianpingReply, err error) {
	err = dianping.CreateTuangouList(req.Appkey)
	if err != nil {
		return &pb.DianpingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.DianpingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) SendSeatBookSubscribe(ctx context.Context, req *pb.BookingSubscribeRequest) (*pb.BookingReply, error) {

	list, err := s.seatUC.ListBookingDetailBy(ctx, &biz.SearchSeatBookingRequest{
		MerchantId: req.MerchantId,
		Status:     []int32{1, 2, 3, 4}, //待开始,未入座,已入座,待明天开始
	})
	if err != nil {
		log.Errorf("Service ListBookingBy error:%s \r\n", err)
		return nil, err
	}
	wx, _ := wechat.NewMessage(req.MerchantId)
	for _, book := range list {
		go func(b biz.MySeatBook, wx wechat.Message) {
			now := time.Now()
			nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
			if nowDay.Before(b.BookingStartDay) || nowDay.After(b.BookingEndDay) {
				return
			}

			startAt, _ := time.ParseInLocation(time.DateTime, b.BookingStartDay.Format(time.DateOnly)+" "+b.BookingStart, time.Local)

			beforeMin := startAt.Add(-time.Minute * 60)     //入座前60分钟通知
			beforeMinAdd2 := startAt.Add(-time.Minute * 59) //
			if b.Status == 1 && now.After(beforeMin) && now.Before(beforeMinAdd2) {
				_ = wx.SendSeatBookReadyMessage(b.MemberID, wechat.SendSeatBookReadyMessage{
					ShopName:      b.ShopLabel,
					SeatName:      b.SeatLabel,
					Countdown:     "60分钟",
					StartDateTime: startAt.Format("2006-01-02 15:04"),
					Tip:           "如无法到达,请提前取消,退还全部时长",
				})
			}
			if b.Status == 3 {
				beforeEndMin := b.TodayEndAt.Add(-time.Minute * 10)    //结束前前10分钟通知
				beforeEndMinAdd2 := b.TodayEndAt.Add(-time.Minute * 9) //
				if now.After(beforeEndMin) && now.Before(beforeEndMinAdd2) {
					_ = wx.SendSeatBookCompleteMessage(b.MemberID, wechat.SendSeatBookCompleteMessage{
						ShopName:      b.ShopLabel,
						SeatName:      b.SeatLabel,
						CloseDateTime: b.TodayEndAt.Format("2006-01-02 15:04"),
						Tip:           "剩余时间10分钟,请检查贵重物品",
					})
				}
			}
		}(*book, *wx)
	}
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) ExecProfitSharingOrder(ctx context.Context, req *pb.ProfitSharingRequest) (*pb.BookingReply, error) {
	err := s.orderUC.ExecProfitSharingOrder(ctx, req.ProfitSharingId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) UpdateProfitSharingOrderStatus(ctx context.Context, req *pb.ProfitSharingRequest) (*pb.BookingReply, error) {
	err := s.orderUC.UpdateProfitSharingOrderStatus(ctx, req.ProfitSharingId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *AdminService) SendCoupon(ctx context.Context, req *pb.SendCouponRequest) (*pb.SendCouponReply, error) {
	couponReceiveRequest := &biz.CouponReceiveRequest{
		CouponId:      req.CouponId,
		MemberId:      req.MemberId,
		Notice:        req.Notice,
		ShopId:        req.ShopId,
		CreateByAdmin: true,
	}
	if ok, err := s.couponUC.ReceiveCouponValidate(ctx, couponReceiveRequest); !ok {
		return &pb.SendCouponReply{
			Code: err.GetCode(),
			Msg:  err.GetReason(),
		}, nil
	}

	c, err := s.couponUC.ReceiveCoupon(ctx, couponReceiveRequest)
	if err != nil {
		return &pb.SendCouponReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	go func() {
		wx, _ := wechat.NewMessage(c.MerchantID)
		_ = wx.SendCouponReceiveMessage(req.MemberId, wechat.CouponReceiveMessage{
			ShopName:    strings.Join(c.ShopLabel, ","),
			CouponName:  c.Name,
			EndDateTime: c.EndDate.Format(time.DateTime),
			Tip:         "请在优惠券有效期内使用该券",
		})
	}()

	return &pb.SendCouponReply{
		Code: 200,
		Msg:  "",
	}, nil
}

func (s *AdminService) StudentChangeCheckIn(ctx context.Context, req *pb.StudentChangeCheckInRequest) (*pb.BookingReply, error) {
	scheduling, err := s.courseUC.FindCourseSchedulingDetailById(ctx, req.CoursesSchedulingId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	now := time.Now()

	// scheduling.CheckIn = int8(req.ChangeCheckIn)
	// if ok, err := s.courseUC.UpdateCourseScheduling(ctx, &scheduling.EdCoursesScheduling); !ok {
	// 	return &pb.BookingReply{
	// 		Code: 400,
	// 		Msg:  err.Error(),
	// 	}, nil
	// }
	wx, _ := wechat.NewMessage(wechat.MerchantToOfficalMerchantId(scheduling.MerchantID))
	if scheduling.CheckIn == 2 {
		err = wx.SendStudentCheckInSuccessMessage(scheduling.StudentID, wechat.SendStudentCheckInSuccess{
			CourseName:  scheduling.StudentCourses.Title,
			StudentName: scheduling.Student.Name,
			AddressName: scheduling.Address.Name,
			CourseTime:  scheduling.StartDate.Format(time.DateOnly) + " " + scheduling.StartTime,
			CheckIntime: now.Format(time.DateTime),
		})
	}
	if scheduling.CheckIn == 3 {
		err = wx.SendStudentCheckOutSuccessMessage(scheduling.StudentID, wechat.SendStudentCheckOutSuccess{
			StudentName:  scheduling.Student.Name,
			CheckOutDate: now.Format(time.DateOnly),
			CheckOutTime: now.Format(time.TimeOnly),
		})
	}
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.BookingReply{
		Code: 200,
	}, nil
}
