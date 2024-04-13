package service

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"studyRoomGo/common/dianping"
	_func "studyRoomGo/common/func"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/service/dto"
	"time"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	pb "studyRoomGo/api/seat/v1"
	"studyRoomGo/common/equipment"
)

type SeatService struct {
	pb.UnimplementedSeatServer

	uc          *biz.SeatUsecase
	equipmentUC *biz.EquipmentUsecase
	packageUC   *biz.PackageCardUsecase
}

func NewSeatService(uc *biz.SeatUsecase, equipmentUC *biz.EquipmentUsecase, packageUC *biz.PackageCardUsecase) *SeatService {
	return &SeatService{uc: uc, equipmentUC: equipmentUC, packageUC: packageUC}
}

func (s *SeatService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	return &pb.CreateReply{}, nil
}
func (s *SeatService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	return &pb.UpdateReply{}, nil
}
func (s *SeatService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return &pb.DeleteReply{}, nil
}
func (s *SeatService) Get(ctx context.Context, req *pb.GetRequest) (*pb.DataReply, error) {
	data, err := s.uc.FindById(ctx, req.Id)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	dto := dto.Seat{Seat: data}
	return &pb.DataReply{
		Code: 200,
		Data: dto.Reply(),
	}, nil
}

func (s *SeatService) Find(ctx context.Context, req *pb.FindRequest) (*pb.DataReply, error) {
	var filterKey = []string{
		"id", "name",
	}
	if !_func.InStrSplice(filterKey, req.Key) {
		return &pb.DataReply{
			Code: 400,
			Msg:  fmt.Sprintf("key not in %v", strings.Join(filterKey, ",")),
		}, nil
	}
	data, err := s.uc.FindBy(ctx, req.Key, req.Value)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	t1, _ := time.ParseInLocation(time.DateTime, req.GteDateEnd, time.Local)
	t3, _ := time.ParseInLocation(time.DateTime, req.LteDateEnd, time.Local)

	if data.Seat != nil {

		bookReq := &biz.SearchSeatBookingRequest{
			SeatId: data.ID,
			Status: []int32{1, 2, 3},
		}
		lockReq := &biz.SearchSeatLockingRequest{
			SeatId: data.ID,
			Status: []int32{1},
		}

		if t1.Year() > 1 {
			bookReq.GteBookDateEnd = t1.Format(time.DateOnly)
			bookReq.GteBookTimeEnd = t1.Format(time.TimeOnly)

			//lockReq.GteLockDateEnd = t1.Format(time.DateOnly)
			lockReq.GteLockTimeEnd = t1.Format(time.TimeOnly)
		}

		if t3.Year() > 1 {
			bookReq.LteBookDateEnd = t3.Format(time.DateOnly)
			bookReq.LteBookTimeEnd = t3.Format(time.TimeOnly)

			lockReq.LteLockDateEnd = t3.Format(time.DateOnly)
			lockReq.LteLockTimeEnd = t3.Format(time.TimeOnly)
		}
		data.Booking, _ = s.uc.ListBookingBy(ctx, bookReq)
		data.Locking, _ = s.uc.ListLockingBy(ctx, lockReq)
	}
	dto := dto.Seat{Seat: data}
	return &pb.DataReply{
		Code: 200,
		Data: dto.Reply(),
	}, nil
}

func (s *SeatService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	dto := dto.Seat{}

	data, err := s.uc.List(ctx)
	if err != nil {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.Data
	for _, item := range data {
		dto.Seat = item
		list = append(list, dto.Reply())
	}
	return &pb.ListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *SeatService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.ListReply, error) {
	dto := dto.Seat{}

	data, err := s.uc.Search(ctx, &biz.SearchSeatRequest{
		Pid:        req.Pid,
		Name:       req.Name,
		TypeSeat:   req.TypeSeat,
		ShopId:     req.ShopId,
		ShopAreaId: req.ShopAreaId,
		Status:     req.Status,
		Sort:       int(req.Sort),
	})
	if err != nil {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.Data
	for _, item := range data {
		dto.Seat = item
		list = append(list, dto.Reply())
	}
	return &pb.ListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *SeatService) Overview(ctx context.Context, req *pb.OverviewRequest) (*pb.OverviewReply, error) {
	dto := dto.Seat{}

	data, err := s.uc.Search(ctx, &biz.SearchSeatRequest{
		ShopId: req.ShopId,
	})
	if err != nil {
		return &pb.OverviewReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	t1, _ := time.ParseInLocation(time.DateTime, req.GteDateEnd, time.Local)
	t2, _ := time.ParseInLocation(time.DateTime, req.GteDateStart, time.Local)
	t3, _ := time.ParseInLocation(time.DateTime, req.LteDateEnd, time.Local)
	t4, _ := time.ParseInLocation(time.DateTime, req.LteDateStart, time.Local)

	bookReq := &biz.SearchSeatBookingRequest{
		ShopId: req.ShopId,
		Status: []int32{1, 2, 3},
	}
	lockReq := &biz.SearchSeatLockingRequest{
		ShopId: req.ShopId,
		Status: []int32{1},
	}
	if t1.Year() > 1 {
		bookReq.GteBookDateEnd = t1.Format(time.DateOnly)
		bookReq.GteBookTimeEnd = t1.Format(time.TimeOnly)

		lockReq.GteLockDateEnd = t1.Format(time.DateOnly)
		lockReq.GteLockTimeEnd = t1.Format(time.TimeOnly)
	}
	if t2.Year() > 1 {
		bookReq.GteBookDateStart = t2.Format(time.DateOnly)
		bookReq.GteBookTimeStart = t2.Format(time.TimeOnly)

		lockReq.GteLockDateStart = t2.Format(time.DateOnly)
		lockReq.GteLockTimeStart = t2.Format(time.TimeOnly)
	}

	if t3.Year() > 1 {
		bookReq.LteBookDateEnd = t3.Format(time.DateOnly)
		bookReq.LteBookTimeEnd = t3.Format(time.TimeOnly)

		lockReq.LteLockDateEnd = t3.Format(time.DateOnly)
		lockReq.LteLockTimeEnd = t3.Format(time.TimeOnly)
	}
	if t4.Year() > 1 {
		bookReq.LteBookDateStart = t4.Format(time.DateOnly)
		bookReq.LteBookTimeStart = t4.Format(time.TimeOnly)

		lockReq.LteLockDateStart = t4.Format(time.DateOnly)
		lockReq.LteLockTimeStart = t4.Format(time.TimeOnly)
	}
	bookingList, _ := s.uc.ListBookingBy(ctx, bookReq)

	lockingList, _ := s.uc.ListLockingBy(ctx, lockReq)
	var list []*pb.OverviewReply_Data
	for _, item := range data {
		dto.Seat = item

		var seatBook []*biz.SeatBook
		var seatLock []*biz.SeatLock
		for _, book := range bookingList {
			if book.SeatID == item.ID {
				seatBook = append(seatBook, book)

			}
		}
		for _, lock := range lockingList {
			if lock.SeatID == item.ID {
				seatLock = append(seatLock, lock)

			}
		}

		dto.Seat.Booking = seatBook
		dto.Seat.Locking = seatLock
		list = append(list, dto.OverviewReply())

	}
	return &pb.OverviewReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *SeatService) BookingValidate(ctx context.Context, req *pb.BookingRequest) (*pb.BookingValidateReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	dto := dto.Booking{}
	bookReq := dto.BookingReq(memberId, req)

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

	res, err := s.uc.BookingPreValidate(ctx, bookReq)
	if err != nil {
		return &pb.BookingValidateReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.BookingValidateReply{
		Code: 200,
		Data: dto.BookingPreValidateReply(res),
	}, nil
}

func (s *SeatService) Booking(ctx context.Context, req *pb.BookingRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	defer func() {
		seatUnLock(req.SeatId)
		memberUnLock(memberId)
	}()

	if !seatAddLock(req.SeatId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  errors.New("座位操作频繁，请稍后").Error(),
		}, nil
	}

	if !memberAddLock(memberId) {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "用户操作频繁，请稍后",
		}, nil
	}
	if req.ReceiptCode != "" { //美团兑换，并立即使用
		pId, count, err := dianping.Prepare(req.ShopId, memberId, req.ReceiptCode)
		if err != nil {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		_, err = s.packageUC.BuyCardValidate(ctx, &biz.CreateBuyCardOrderRequest{
			MemberId: memberId,
			CardId:   pId,
		})
		if err != nil {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		_, err = dianping.Consume(req.ShopId, memberId, req.ReceiptCode, count)
		if err != nil {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		for i := 0; i < count; i++ {
			mId, err := s.packageUC.CreateMemberPackageCard(ctx, &biz.CreateMemberPackageCardRequest{
				MemberId:     memberId,
				PackageId:    pId,
				TypeChannel:  3,
				IsActive:     true,
				ActiveAmount: 0,
				ReceiptCode:  req.ReceiptCode,
			})
			if err != nil {
				return &pb.BookingReply{
					Code: 400,
					Msg:  err.Error(),
				}, nil
			}
			req.MemberPackageCardId = mId
		}
		card, _ := s.packageUC.FindMemberCardById(ctx, memberId, req.MemberPackageCardId)
		if card.TypePackage == 2 || card.TypePackage == 3 {
			req.BookingToCloseTime = true
		}
	}
	dto := dto.Booking{}
	bookReq := dto.BookingReq(memberId, req)
	bookReq.MerchantId = int64(merchantId)
	if ok, err := s.uc.BookingValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: err.GetCode(),
			Msg:  err.GetReason(),
		}, nil
	}

	if ok, err := s.uc.Booking(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	notSeatMinCancel := biz.GetNotSeatMinCancel(bookReq.ShopId)

	wx, _ := wechat.NewMessage(bookReq.MerchantId)
	_ = wx.SendSeatBookMessage(bookReq.MemberId, wechat.SeatBookMessage{
		ShopName:  bookReq.ShopLabel,
		SeatName:  bookReq.SeatLabel,
		StartDate: bookReq.BookingStartDay.Format(time.DateOnly),
		TimeRange: bookReq.BookingStart.Format("15:04") + "~" + bookReq.BookingEnd.Format("15:04"),
		Tip:       fmt.Sprintf("到期未入座,座位保留%d分钟", notSeatMinCancel),
	})
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *SeatService) BookingAddTime(ctx context.Context, req *pb.BookingAddTimeRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	book, err := s.uc.FindBookingById(ctx, req.BookId)
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
	if ok, err := s.uc.BookingAddTimeValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if ok, err := s.uc.BookingValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.GetReason(),
		}, nil
	}

	if ok, err := s.uc.UpdateBooking(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *SeatService) BookingChangeSeat(ctx context.Context, req *pb.BookingChangeSeatRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	book, err := s.uc.FindBookingById(ctx, req.BookId)
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
		ChangeSeatId:        req.ChangeSeatId,
		Status:              book.Status,
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
	if ok, err := s.uc.BookingChangeSeatValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if ok, err := s.uc.BookingValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.GetReason(),
		}, nil
	}

	if ok, err := s.uc.UpdateBooking(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	if bookReq.Status == biz.SeatBook_StatusSeated {
		//断电
		if ok, err := s.seatAutomate(ctx, bookReq.SeatId, false); !ok {
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
			_, _ = s.seatAutomate(ctx, ChangeSeatId, true)
		}(bookReq.ChangeSeatId)
	}

	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *SeatService) BookingChangeStatus(ctx context.Context, req *pb.BookingChangeStatusRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	book, err := s.uc.FindBookingById(ctx, req.BookId)
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
		NewStatus:           int8(req.ChangeStatus),
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

	if ok, err := s.uc.BookingChangeStatusValidate(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	if ok, err := s.uc.UpdateBooking(ctx, bookReq); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	if bookReq.NewStatus == biz.SeatBook_StatusSeated { //入座,通电
		if ok, _ := s.seatAutomate(ctx, bookReq.SeatId, true); !ok {
			//return &pb.BookingReply{
			//	Code: 400,
			//	Msg:  "开灯失败，请手动重试或联系管理员",
			//}, nil
		}
	}

	if bookReq.NewStatus == biz.SeatBook_StatusComplete { //结束,断电
		if ok, _ := s.seatAutomate(ctx, bookReq.SeatId, false); !ok {
			//return &pb.BookingReply{
			//	Code: 400,
			//	Msg:  "关灯失败，请手动重试或联系管理员",
			//}, nil
		}
	}
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *SeatService) BookingControlLight(ctx context.Context, req *pb.BookingControlSeatRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	book, err := s.uc.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if book.MemberID != memberId {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "非本人订座不能操作",
		}, nil
	}
	return s.controlLight(ctx, book)
}
func (s *SeatService) BookingControlLightShared(ctx context.Context, req *pb.BookingControlSeatRequest) (*pb.BookingReply, error) {
	str := fmt.Sprintf("%d|%s", req.BookId, biz.SharedKey)
	hash := hmac.New(md5.New, []byte(biz.SharedKey)) //创建对应的sha256哈希加密算法
	hash.Write([]byte(str))
	token := hex.EncodeToString(hash.Sum([]byte("")))
	if token != req.Token {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "认证失败",
		}, nil
	}
	_, err := s.uc.FindSeatBookSharedLog(ctx, req.BookId, req.OpenId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	book, err := s.uc.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return s.controlLight(ctx, book)
}

func (s *SeatService) BookingOpenDoor(ctx context.Context, req *pb.BookingOpenDoorRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	book, err := s.uc.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if book.MemberID != memberId {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "非本人订座不能操作",
		}, nil
	}
	return s.openDoor(ctx, book, req.Door)
}

func (s *SeatService) BookingOpenDoorShared(ctx context.Context, req *pb.BookingOpenDoorRequest) (*pb.BookingReply, error) {
	str := fmt.Sprintf("%d|%s", req.BookId, biz.SharedKey)
	hash := hmac.New(md5.New, []byte(biz.SharedKey)) //创建对应的sha256哈希加密算法
	hash.Write([]byte(str))
	token := hex.EncodeToString(hash.Sum([]byte("")))
	if token != req.Token {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "认证失败",
		}, nil
	}
	_, err := s.uc.FindSeatBookSharedLog(ctx, req.BookId, req.OpenId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	book, err := s.uc.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return s.openDoor(ctx, book, req.Door)
}
func (s *SeatService) BookingShareGenToken(ctx context.Context, req *pb.BookingControlSeatRequest) (*pb.BookingGenSharekDataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	book, err := s.uc.FindBookingById(ctx, req.BookId)
	if err != nil {
		return &pb.BookingGenSharekDataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if book.MemberID != memberId {
		return &pb.BookingGenSharekDataReply{
			Code: 400,
			Msg:  "非本人订座不能分享",
		}, nil
	}
	seat, err := s.uc.FindById(ctx, book.SeatID)
	if err != nil {
		return nil, err
	}
	if seat.Share == 0 {
		return &pb.BookingGenSharekDataReply{
			Code: 400,
			Msg:  "本座位不允许分享",
		}, nil
	}
	if !_func.InInt8Splice([]int8{1, 2, 3, 4}, book.Status) {
		return &pb.BookingGenSharekDataReply{
			Code: 400,
			Msg:  "订座已失效,不允许分享",
		}, nil
	}
	now := time.Now()
	endTime, _ := time.ParseInLocation(time.DateTime, book.BookingEndDay.Format(time.DateOnly)+" "+book.BookingEnd, time.Local)
	if now.After(endTime) {
		return &pb.BookingGenSharekDataReply{
			Code: 400,
			Msg:  "不在入座时间内",
		}, nil
	}
	str := fmt.Sprintf("%d|%s", book.ID, biz.SharedKey)
	hash := hmac.New(md5.New, []byte(biz.SharedKey)) //创建对应的sha256哈希加密算法
	hash.Write([]byte(str))
	token := hex.EncodeToString(hash.Sum([]byte("")))

	return &pb.BookingGenSharekDataReply{
		Code: 200,
		Msg:  "",
		Data: &pb.BookingGenSharekDataReply_Data{Token: token},
	}, nil
}

// LockingOpenDoor 锁位无需订座，开门
func (s *SeatService) LockingOpenDoor(ctx context.Context, req *pb.LockingOpenDoorRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	lock, err := s.uc.FindLockingById(ctx, req.LockId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if lock.MemberID != memberId {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "非本人订座不能操作",
		}, nil
	}
	return s.lockOpenDoor(ctx, lock, req.Door)
}

// LockingControlLight 锁位无需订座，开灯
func (s *SeatService) LockingControlLight(ctx context.Context, req *pb.LockingControlSeatRequest) (*pb.BookingReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	lock, err := s.uc.FindLockingById(ctx, req.LockId)
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if lock.MemberID != memberId {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "非本人订座不能操作",
		}, nil
	}
	return s.lockControlLight(ctx, lock)
}

func (s *SeatService) seatAutomate(ctx context.Context, seatId int64, status bool) (ok bool, err error) {
	data, err := s.equipmentUC.FindByRelateId(ctx, &biz.EquipmentRequest{
		TypeRelate: 1,
		RelateId:   seatId,
	})
	if err != nil {
		return false, err
	}

	fmt.Printf("SeatService_SeatAutomate %d status %v", seatId, status)
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

func (s *SeatService) openDoor(ctx context.Context, book *biz.SeatBook, door int64) (*pb.BookingReply, error) {
	var eq *gen.SrEquipment
	var err error
	if door == 1 {
		if ok, err := s.uc.BookingOpenBigDoorValidate(ctx, book); !ok {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		area, err := s.uc.FindShopAreaReception(ctx, book.ShopID)
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

		if ok, err := s.uc.BookingOpenDoorValidate(ctx, book); !ok {
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

func (s *SeatService) controlLight(ctx context.Context, book *biz.SeatBook) (*pb.BookingReply, error) {

	if ok, err := s.uc.BookingControlLightValidate(ctx, book); !ok {
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
	if ok, _ := s.seatAutomate(ctx, book.SeatID, operation); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "灯控失败，请手动重试或联系管理员",
		}, nil
	}

	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *SeatService) lockOpenDoor(ctx context.Context, lock *biz.SeatLock, door int64) (*pb.BookingReply, error) {
	var eq *gen.SrEquipment
	var err error
	if ok, err := s.uc.LockingControlValidate(ctx, lock); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	areaId := lock.ShopAreaID
	if door == 1 {
		area, err := s.uc.FindShopAreaReception(ctx, lock.ShopID)
		if err != nil {
			return &pb.BookingReply{
				Code: 400,
				Msg:  err.Error(),
			}, nil
		}
		areaId = area.ID
	}
	eq, err = s.equipmentUC.FindByRelateId(ctx, &biz.EquipmentRequest{
		TypeRelate: 2,
		RelateId:   areaId,
		ShoId:      lock.ShopID,
	})
	if err != nil {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if lock.ShopID == 2 {
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
	if door == 1 {
		//记录入座
		s.uc.RecordLockingControlLog(ctx, lock)
	}
	return &pb.BookingReply{
		Code: 200,
	}, nil
}

func (s *SeatService) lockControlLight(ctx context.Context, lock *biz.SeatLock) (*pb.BookingReply, error) {

	if ok, err := s.uc.LockingControlValidate(ctx, lock); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	eqData, err := s.equipmentUC.FindByRelateId(ctx, &biz.EquipmentRequest{
		TypeRelate: 1,
		RelateId:   lock.SeatID,
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
	if ok, _ := s.seatAutomate(ctx, lock.SeatID, operation); !ok {
		return &pb.BookingReply{
			Code: 400,
			Msg:  "灯控失败，请手动重试或联系管理员",
		}, nil
	}

	return &pb.BookingReply{
		Code: 200,
	}, nil
}
