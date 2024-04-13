package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"studyRoomGo/common/config"
	_func "studyRoomGo/common/func"
	"studyRoomGo/common/notice"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
	"time"

	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

const SharedKey = "gdfgouoihkjbxzcvmnoiy"

// SeatUsecase is a Seat usecase.
type SeatUsecase struct {
	repo               SeatRepo
	seatBookRepo       SeatBookRepo
	memberPackCardRepo MemberPackageCardRepo
	shopRepo           ShopRepo

	log *log.Helper
}

// NewSeatUsecase new a Seat usecase.
func NewSeatUsecase(repo SeatRepo, seatBookRepo SeatBookRepo, memberPackCardRepo MemberPackageCardRepo, shopRepo ShopRepo, logger log.Logger) *SeatUsecase {
	return &SeatUsecase{
		repo:               repo,
		seatBookRepo:       seatBookRepo,
		memberPackCardRepo: memberPackCardRepo,
		shopRepo:           shopRepo,
		log:                log.NewHelper(logger),
	}
}

type Seat struct {
	*model.Seat
	Booking []*SeatBook
	Locking []*SeatLock
}

func (uc *SeatUsecase) FindById(ctx context.Context, id int64) (*Seat, error) {
	data, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Seat{
		Seat: data,
	}, nil

}

func (uc *SeatUsecase) FindBy(ctx context.Context, key string, value string) (*Seat, error) {
	data, err := uc.repo.FindBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	return &Seat{
		Seat: data,
	}, nil
}

func (uc *SeatUsecase) List(ctx context.Context) ([]*Seat, error) {
	data, err := uc.repo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	var list []*Seat
	for _, item := range data {
		list = append(list, &Seat{Seat: item})
	}
	return list, nil
}

type SearchSeatRequest struct {
	Pid        int64
	Name       string
	TypeSeat   int32
	ShopId     int64
	ShopAreaId int64
	Status     int32
	Sort       int
}

func (uc *SeatUsecase) Search(ctx context.Context, req *SearchSeatRequest) ([]*Seat, error) {
	data, err := uc.repo.ListBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*Seat
	for _, item := range data {
		list = append(list, &Seat{Seat: item})
	}
	return list, nil
}

func (uc *SeatUsecase) ListBookingBy(ctx context.Context, req *SearchSeatBookingRequest) ([]*SeatBook, error) {
	data, err := uc.repo.ListBookingBy(ctx, req)
	if err != nil {
		log.Errorf("Service ListBookingBy error:%s \r\n", err)
		return nil, err
	}
	var list []*SeatBook
	for _, item := range data {
		list = append(list, &SeatBook{item})
	}
	return list, nil
}

func (uc *SeatUsecase) ListBookingDetailBy(ctx context.Context, req *SearchSeatBookingRequest) ([]*MySeatBook, error) {
	data, err := uc.seatBookRepo.ListBy(ctx, req)
	if err != nil {
		log.Errorf("Service ListBookingBy error:%s \r\n", err)
		return nil, err
	}
	var list []*MySeatBook
	for _, item := range data {
		list = append(list, &MySeatBook{SeatBook: item})
	}
	return list, nil
}
func (uc *SeatUsecase) ListLockingBy(ctx context.Context, req *SearchSeatLockingRequest) ([]*SeatLock, error) {
	data, err := uc.repo.ListLockingBy(ctx, req)
	if err != nil {
		log.Errorf("Service ListLockingBy error:%s \r\n", err)
		return nil, err
	}
	var list []*SeatLock
	for _, item := range data {
		list = append(list, &SeatLock{SrSeatLock: &item.SrSeatLock})
	}
	return list, nil
}

type TodayCanBookRange struct {
	SeatId        int64
	BookDate      string
	BookTimeStart string
	BookTimeEnd   string
	MaxBookMinute int64
}

type BookingRequest struct {
	MerchantId          int64
	SeatId              int64
	SeatLabel           string
	MemberId            int64
	MemberPackageCardId int64
	ShopId              int64
	ShopLabel           string
	ShopAreaId          int64
	BookingType         int32
	BookingDay          int32
	BookingMinute       int32
	BookingToCloseTime  bool
	BookingStartDay     time.Time
	BookingEndDay       time.Time
	BookingStart        time.Time
	BookingEnd          time.Time
	BookingWeek         []int32
	ValidDay            int32
	Status              int8
	NoEnoughUse         int32  //剩余时间不足最低消费
	UseRemain           int32  //使用剩余时间消费
	ReceiptCode         string //美团兑换
	//换座
	SeatBookId       int64   //原订座id
	ChangeSeatId     int64   //换座id
	ChangeShopAreaId int64   //换座后店铺区域
	ChildIds         []int64 //子座位

	//更新状态
	NewStatus int8

	//延长订座时间
	AddTimeToCloseTime bool
	AddTimeMin         int32

	CreateByAdmin bool
}

func (uc *SeatUsecase) Booking(ctx context.Context, req *BookingRequest) (ok bool, err error) {

	cardInfo, err := uc.memberPackCardRepo.FindByID(ctx, req.MemberPackageCardId)
	memberPackageCard := &MemberPackageCard{SrMemberPackageCard: cardInfo}
	useLog := &MemberPackageCardUseLog{
		MpID:     memberPackageCard.ID,
		MemberID: req.MemberId,
		Extra:    fmt.Sprintf("seatId: %v", req.SeatId),
	}
	oriRemainMinute := int32(memberPackageCard.RemainMinute)
	if memberPackageCard.RemainMinute != 0 {
		memberPackageCard.RemainMinute = memberPackageCard.RemainMinute - int(req.BookingMinute)
		useLog.Value = -int(req.BookingMinute)
		if req.NoEnoughUse == 1 { //最后一天用完所有时长
			memberPackageCard.RemainMinute = 0
			useLog.Value = -int(memberPackageCard.RemainMinute)
		}
	} else if memberPackageCard.RemainDay != 0 {
		memberPackageCard.RemainDay = memberPackageCard.RemainDay - int(req.BookingDay)
		useLog.Value = -int(req.BookingDay)
	} else {
		if memberPackageCard.TypePackage != 3 {
			err := errors.New("Remain  Error")
			log.Errorf("memberPackCardRepo Insert error:%s \r\n", err)
			return false, err
		}
	}

	_, err = uc.memberPackCardRepo.Update(ctx, memberPackageCard)
	if err != nil {
		log.Errorf("memberPackCardRepo Update error:%s \r\n", err)
		return false, err
	}
	uc.memberPackCardRepo.RecordUseLog(ctx, useLog)

	uniqueToken := GenerateUniqueToken(req.ShopId, req.MemberId)

	diffDay := int32(math.Ceil(req.BookingEndDay.Sub(req.BookingStartDay).Hours() / 24)) //相差多少天
	//按每天拆分订座单
	for i := 0; i <= int(diffDay); i++ {
		seatBook := &SeatBook{}
		seatBook.Generate(req)
		seatBook.UniqueToken = &uniqueToken
		startDay := seatBook.BookingStartDay.Add(time.Hour * time.Duration(24*i))
		if len(req.BookingWeek) > 0 && !_func.InInt32Splice(req.BookingWeek, int32(startDay.Weekday())) { //排除不选择星期
			continue
		}
		seatBook.BookingStartDay = startDay
		seatBook.BookingEndDay = startDay
		if req.BookingDay > 0 {
			seatBook.BookingDay = 1
		}
		if req.BookingMinute > 0 {
			seatBook.BookingMinute = int(req.BookingMinute / req.ValidDay) //时卡订座 拆分每天
			if req.NoEnoughUse == 1 && i == (int(req.ValidDay)-1) {        //最后一天用完所有时长
				lastRemainMin := int(oriRemainMinute - req.BookingMinute)
				seatBook.BookingMinute += lastRemainMin
				now := time.Now()
				newBookEnd, _ := time.ParseInLocation(time.DateTime, now.Format(time.DateOnly)+" "+seatBook.BookingEnd, time.Local)
				newBookEnd = newBookEnd.Add(time.Minute * time.Duration(lastRemainMin))
				seatBook.BookingEnd = newBookEnd.Format(time.TimeOnly)
			}
		}

		_, err = uc.seatBookRepo.Save(ctx, seatBook)
		if err != nil {
			log.Errorf("seatBookService Insert error:%s \r\n", err)
			return false, err
		}

		operateLog := &SeatBookOperateLog{
			SeatBookId:    seatBook.ID,
			ShopId:        req.ShopId,
			SeatId:        req.SeatId,
			MemberId:      req.MemberId,
			Remark:        fmt.Sprintf("订座初始状态,订座 %v 分钟", seatBook.BookingMinute),
			CreateByAdmin: req.CreateByAdmin,
		}
		operateLog.StatusToOperate(seatBook.Status)
		uc.seatBookRepo.RecordOperateLog(ctx, operateLog)

	}

	return true, nil
}

func (uc *SeatUsecase) FindBookingById(ctx context.Context, id int64) (book *SeatBook, err error) {

	data, err := uc.seatBookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &SeatBook{
		SrSeatBook: &data.SrSeatBook,
	}, nil
}

func (uc *SeatUsecase) FindLockingById(ctx context.Context, id int64) (look *SeatLock, err error) {
	data, err := uc.repo.FindLockingBy(ctx, &SearchSeatLockingRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &SeatLock{
		SrSeatLock: &data.SrSeatLock,
	}, nil
}

func (uc *SeatUsecase) FindShopAreaReception(ctx context.Context, shopId int64) (area *gen.SrShopArea, err error) {

	data, err := uc.shopRepo.FindShopAreaReception(ctx, shopId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *SeatUsecase) UpdateBooking(ctx context.Context, req *BookingRequest) (ok bool, err error) {
	if req.ChangeSeatId != 0 { //换座
		return uc.UpdateBookingChangeSeat(ctx, req)
	}
	if req.AddTimeMin != 0 || req.AddTimeToCloseTime { //延长订座时间
		return uc.UpdateBookingAddTime(ctx, req)
	}
	if req.NewStatus == SeatBook_StatusComplete { //结束订单
		return uc.UpdateBookingCompleted(ctx, req)
	}
	if req.NewStatus == SeatBook_StatusSeated { //入座
		ok, err = uc.UpdateBookingSeated(ctx, req)

		cardInfo, _ := uc.memberPackCardRepo.FindByID(ctx, req.MemberPackageCardId)
		lastCardInfo, _ := uc.memberPackCardRepo.ListBy(ctx, &SearchMemberPackageCardRequest{
			MemberId: int64(req.MemberId),
			Page:     1,
			PageSize: 1,
			Sort:     1,
		})
		//站内通知
		n := notice.Notice{
			Ctx:      ctx,
			MemberId: req.MemberId,
		}
		go n.SeatBookStatusSeatedNotice(notice.CardInfo{
			Name:          cardInfo.PackageName,
			TypeCard:      cardInfo.TypeCard,
			TypePackage:   cardInfo.TypePackage,
			RemainMin:     cardInfo.RemainMinute,
			RemainDay:     cardInfo.RemainDay,
			InvalidAt:     cardInfo.InvalidAt,
			LastCardId:    lastCardInfo[0].ID,
			CurrentCardId: int64(cardInfo.ID),
		})
		return ok, err
	}
	if req.NewStatus == SeatBook_StatusCustomerCancel { //用户取消订单
		return uc.UpdateBookingCustomerCancel(ctx, req)
	}
	if req.NewStatus == SeatBook_StatusSystemCancel { //系统取消订单
		return uc.UpdateBookingSystemCancel(ctx, req)
	}
	return false, errors.New("状态错误")
}

func (uc *SeatUsecase) UpdateBookingChangeSeat(ctx context.Context, req *BookingRequest) (ok bool, err error) {

	seat, _ := uc.seatBookRepo.FindByID(ctx, req.SeatBookId)

	seatBook := &SeatBook{
		SrSeatBook: &seat.SrSeatBook,
	}
	seatBook.SeatID = req.ChangeSeatId
	seatBook.ShopAreaID = req.ChangeShopAreaId

	_, err = uc.seatBookRepo.Update(ctx, seatBook)
	if err != nil {
		log.Errorf("seatBookService Update error:%s \r\n", err)
		return false, err
	}
	operateLog := &SeatBookOperateLog{
		SeatBookId:    seat.ID,
		ShopId:        req.ShopId,
		SeatId:        req.SeatId,
		MemberId:      req.MemberId,
		Operate:       SeatBook_OperateChangeSeat,
		Remark:        fmt.Sprintf("换座 %v -> %v", req.SeatId, req.ChangeSeatId),
		Extra:         nil,
		CreateByAdmin: req.CreateByAdmin,
	}

	uc.seatBookRepo.RecordOperateLog(ctx, operateLog)
	return true, nil
}

func (uc *SeatUsecase) UpdateBookingAddTime(ctx context.Context, req *BookingRequest) (ok bool, err error) {

	seat, _ := uc.seatBookRepo.FindByID(ctx, req.SeatBookId)

	seatBook := &SeatBook{
		SrSeatBook: &seat.SrSeatBook,
	}
	var remark string
	var bookingMinute int
	endTime, _ := time.ParseInLocation(time.DateTime, seatBook.BookingEndDay.Format(time.DateOnly)+" "+seatBook.BookingEnd, time.Local)

	if seatBook.Status == SeatBook_StatusSeated { //已入座
		endTime = *seatBook.TodayEndAt
	}

	if req.AddTimeMin > 0 {
		newEndTime := endTime.Add(time.Minute * time.Duration(req.AddTimeMin))
		seatBook.BookingEnd = newEndTime.Format(time.TimeOnly)
		seatBook.TodayEndAt = &newEndTime
		bookingMinute = int(req.AddTimeMin)
		seatBook.BookingMinute += bookingMinute

		remark = fmt.Sprintf("延长时间 %v 分钟", bookingMinute)

	}
	if req.AddTimeToCloseTime {
		shopSetting, _ := config.GetShopSetting(seatBook.ShopID)
		var closeTime string
		_ = json.Unmarshal([]byte(shopSetting["closeTime"]), &closeTime)

		newEndTime, _ := time.ParseInLocation(time.DateTime, seatBook.BookingEndDay.Format(time.DateOnly)+" "+closeTime, time.Local)
		seatBook.BookingEnd = newEndTime.Format(time.TimeOnly)
		seatBook.TodayEndAt = &newEndTime
		diff := newEndTime.Sub(endTime)
		bookingMinute = int(math.Ceil(diff.Minutes()))
		seatBook.BookingMinute += bookingMinute
		remark = fmt.Sprintf("延长时间 到营业结束 延长 %v 分钟", bookingMinute)

	}

	cardInfo, err := uc.memberPackCardRepo.FindByID(ctx, req.MemberPackageCardId)
	memberPackageCard := &MemberPackageCard{SrMemberPackageCard: cardInfo}
	useLog := &MemberPackageCardUseLog{
		MpID:     memberPackageCard.ID,
		MemberID: req.MemberId,
		Extra:    fmt.Sprintf("seatId: %v", req.SeatId),
	}

	if memberPackageCard.RemainMinute != 0 {
		memberPackageCard.RemainMinute = memberPackageCard.RemainMinute - bookingMinute
		useLog.Value = -int(req.BookingMinute)
	}

	_, err = uc.memberPackCardRepo.Update(ctx, memberPackageCard)
	if err != nil {
		log.Errorf("memberPackCardRepo Update error:%s \r\n", err)
		return false, err
	}
	uc.memberPackCardRepo.RecordUseLog(ctx, useLog)

	_, err = uc.seatBookRepo.Update(ctx, seatBook)
	if err != nil {
		log.Errorf("seatBookService Update error:%s \r\n", err)
		return false, err
	}
	operateLog := &SeatBookOperateLog{
		SeatBookId:    seatBook.ID,
		ShopId:        req.ShopId,
		SeatId:        req.SeatId,
		MemberId:      req.MemberId,
		Operate:       SeatBook_OperateAddTime,
		Remark:        remark,
		Extra:         nil,
		CreateByAdmin: req.CreateByAdmin,
	}
	uc.seatBookRepo.RecordOperateLog(ctx, operateLog)
	return true, nil
}

func (uc *SeatUsecase) UpdateBookingSeated(ctx context.Context, req *BookingRequest) (ok bool, err error) {

	seat, _ := uc.seatBookRepo.FindByID(ctx, req.SeatBookId)

	seatBook := &SeatBook{
		SrSeatBook: &seat.SrSeatBook,
	}
	seatBook.Status = SeatBook_StatusSeated

	now := time.Now()
	startAt, _ := time.ParseInLocation(time.DateTime, now.Format(time.DateOnly)+" "+seatBook.BookingStart, time.Local)
	endAt, _ := time.ParseInLocation(time.DateTime, now.Format(time.DateOnly)+" "+seatBook.BookingEnd, time.Local)
	remark := fmt.Sprintf("入座 %v -> %v", startAt, endAt)
	if seatBook.BookingMinute > 0 {
		if now.Before(startAt) { //提前入座
			endAt = now.Add(time.Minute * time.Duration(seatBook.BookingMinute)) //整段时间提前

			remark = fmt.Sprintf("提前入座 %v -> %v", startAt, endAt)
		}
	}
	nowS := now.Second()
	now = now.Add(time.Second * time.Duration(60-nowS)) //正秒开始
	secondE := endAt.Second()
	endAt = endAt.Add(time.Second * time.Duration(60-secondE)) //正秒结束
	seatBook.TodayEndAt = &endAt                               //计算离座时间
	seatBook.TodaySeatAt = &now                                //入座时间

	_, err = uc.seatBookRepo.Update(ctx, seatBook)
	if err != nil {
		log.Errorf("seatBookService Update error:%s \r\n", err)
		return false, err
	}

	operateLog := &SeatBookOperateLog{
		SeatBookId:    seat.ID,
		ShopId:        req.ShopId,
		SeatId:        req.SeatId,
		MemberId:      req.MemberId,
		Remark:        remark,
		Extra:         nil,
		CreateByAdmin: req.CreateByAdmin,
	}
	operateLog.StatusToOperate(seatBook.Status)

	uc.seatBookRepo.RecordOperateLog(ctx, operateLog)
	return true, nil
}

func (uc *SeatUsecase) UpdateBookingCompleted(ctx context.Context, req *BookingRequest) (ok bool, err error) {

	seat, _ := uc.seatBookRepo.FindByID(ctx, req.SeatBookId)

	seatBook := &SeatBook{
		SrSeatBook: &seat.SrSeatBook,
	}
	seatBook.Status = SeatBook_StatusComplete
	now := time.Now().Local()
	endAt := *seatBook.TodayEndAt
	seatBook.TodayEndAt = &now //用户点击完成，更新为实际结束时间
	_, err = uc.seatBookRepo.Update(ctx, seatBook)
	if err != nil {
		log.Errorf("seatBookService Update error:%s \r\n", err)
		return false, err
	}

	operateLog := &SeatBookOperateLog{
		SeatBookId:    seatBook.ID,
		ShopId:        req.ShopId,
		SeatId:        req.SeatId,
		MemberId:      req.MemberId,
		Remark:        fmt.Sprintf("结束订单"),
		Extra:         nil,
		CreateByAdmin: req.CreateByAdmin,
	}
	operateLog.StatusToOperate(seatBook.Status)

	uc.seatBookRepo.RecordOperateLog(ctx, operateLog)

	cardInfo, err := uc.memberPackCardRepo.FindByID(ctx, req.MemberPackageCardId)
	memberPackageCard := &MemberPackageCard{SrMemberPackageCard: cardInfo}
	useLog := &MemberPackageCardUseLog{
		MpID:     memberPackageCard.ID,
		MemberID: req.MemberId,
		Remark:   "用户结束订单",
		Extra:    fmt.Sprintf("seatId: %v", req.SeatId),
	}
	if memberPackageCard.TypePackage == 1 { //时卡
		minBookMin := GetMinBookMin(seatBook.ShopID)
		remainMin := int(math.Floor(endAt.Sub(now).Minutes())) //剩余分钟
		useMin := seatBook.BookingMinute - remainMin           //总分钟-剩余分钟=使用分钟
		refundMin := 0                                         //退还分钟
		if useMin <= minBookMin {                              //不足【minBookMin】分钟，按【minBookMin】分钟算
			refundMin = seatBook.BookingMinute - minBookMin
		} else {
			refundMin = remainMin //剩余分钟
		}
		if refundMin <= 0 { //防止修改数据误操作
			refundMin = 0
		}
		memberPackageCard.RemainMinute += refundMin
		useLog.Value = +refundMin
		useLog.Remark += fmt.Sprintf(" 退还: %v 分钟", refundMin)
		if refundMin > 0 && memberPackageCard.Status == MemberPackage_StatusClean {
			memberPackageCard.Status = MemberPackage_StatusNormal
		}

		_, err = uc.memberPackCardRepo.Update(ctx, memberPackageCard)
		if err != nil {
			log.Errorf("memberPackCardRepo Update error:%s \r\n", err)
			return false, err
		}
		uc.memberPackCardRepo.RecordUseLog(ctx, useLog)

	}

	return true, nil
}

func (uc *SeatUsecase) UpdateBookingCustomerCancel(ctx context.Context, req *BookingRequest) (ok bool, err error) {

	seat, _ := uc.seatBookRepo.FindByID(ctx, req.SeatBookId)

	seatBook := &SeatBook{
		SrSeatBook: &seat.SrSeatBook,
	}
	seatBook.Status = SeatBook_StatusCustomerCancel

	_, err = uc.seatBookRepo.Update(ctx, seatBook)
	if err != nil {
		log.Errorf("seatBookService Update error:%s \r\n", err)
		return false, err
	}

	operateLog := &SeatBookOperateLog{
		SeatBookId:    seat.ID,
		ShopId:        req.ShopId,
		SeatId:        req.SeatId,
		MemberId:      req.MemberId,
		Operate:       SeatBook_OperateCustomerCancel,
		Remark:        fmt.Sprintf("取消订单"),
		Extra:         nil,
		CreateByAdmin: req.CreateByAdmin,
	}
	operateLog.StatusToOperate(seatBook.Status)

	uc.seatBookRepo.RecordOperateLog(ctx, operateLog)

	cardInfo, err := uc.memberPackCardRepo.FindByID(ctx, req.MemberPackageCardId)
	memberPackageCard := &MemberPackageCard{SrMemberPackageCard: cardInfo}
	useLog := &MemberPackageCardUseLog{
		MpID:     memberPackageCard.ID,
		MemberID: req.MemberId,
		Remark:   "用户取消订单",
		Extra:    fmt.Sprintf("seatId: %v", req.SeatId),
	}

	if seatBook.BookingMinute != 0 {
		memberPackageCard.RemainMinute = memberPackageCard.RemainMinute + seatBook.BookingMinute
		useLog.Value = +seatBook.BookingMinute
	}
	if memberPackageCard.TypePackage == 2 && seatBook.BookingDay != 0 { //次卡
		memberPackageCard.RemainDay = memberPackageCard.RemainDay + seatBook.BookingDay
		useLog.Value = +seatBook.BookingDay
	}
	if memberPackageCard.Status == MemberPackage_StatusClean {
		memberPackageCard.Status = MemberPackage_StatusNormal
	}
	_, err = uc.memberPackCardRepo.Update(ctx, memberPackageCard)
	if err != nil {
		log.Errorf("memberPackCardRepo Update error:%s \r\n", err)
		return false, err
	}
	uc.memberPackCardRepo.RecordUseLog(ctx, useLog)

	return true, nil
}

func (uc *SeatUsecase) UpdateBookingSystemCancel(ctx context.Context, req *BookingRequest) (ok bool, err error) {

	seat, _ := uc.seatBookRepo.FindByID(ctx, req.SeatBookId)

	seatBook := &SeatBook{
		SrSeatBook: &seat.SrSeatBook,
	}
	seatBook.Status = SeatBook_StatusSystemCancel
	_, err = uc.seatBookRepo.Update(ctx, seatBook)
	if err != nil {
		log.Errorf("seatBookService Update error:%s \r\n", err)
		return false, err
	}

	operateLog := &SeatBookOperateLog{
		SeatBookId:    seatBook.ID,
		ShopId:        req.ShopId,
		SeatId:        req.SeatId,
		MemberId:      req.MemberId,
		Remark:        fmt.Sprintf("系统操作|超时未入座取消订单|"),
		Extra:         nil,
		CreateByAdmin: req.CreateByAdmin,
	}
	operateLog.StatusToOperate(seatBook.Status)

	uc.seatBookRepo.RecordOperateLog(ctx, operateLog)

	cardInfo, err := uc.memberPackCardRepo.FindByID(ctx, req.MemberPackageCardId)
	memberPackageCard := &MemberPackageCard{SrMemberPackageCard: cardInfo}
	useLog := &MemberPackageCardUseLog{
		MpID:     memberPackageCard.ID,
		MemberID: req.MemberId,
		Remark:   "系统操作｜超时未入座取消订单",
		Extra:    fmt.Sprintf("seatId: %v", req.SeatId),
	}
	if memberPackageCard.TypePackage == 1 { //时卡
		minBookMin := GetMinBookMin(seatBook.ShopID)

		refundMin := seatBook.BookingMinute //超时未入座 扣除【minBookMin】分钟
		refundMin -= minBookMin
		if refundMin <= 0 { //防止修改数据误操作
			refundMin = 0
		}
		memberPackageCard.RemainMinute += refundMin
		useLog.Value = +refundMin
		useLog.Remark += fmt.Sprintf(" 退还: %v 分钟", refundMin)
		if refundMin > 0 && memberPackageCard.Status == MemberPackage_StatusClean {
			memberPackageCard.Status = MemberPackage_StatusNormal
		}
		_, err = uc.memberPackCardRepo.Update(ctx, memberPackageCard)
		if err != nil {
			log.Errorf("memberPackCardRepo Update error:%s \r\n", err)
			return false, err
		}
		uc.memberPackCardRepo.RecordUseLog(ctx, useLog)
	}
	if memberPackageCard.TypePackage == 2 { //次卡
		if memberPackageCard.RemainDay == 0 && memberPackageCard.Status == MemberPackage_StatusClean { //剩余0天时,如果超时未入座,返还1天,并且有效期改为当天有效
			// 获取当前时间
			now := time.Now()
			// 设置时间为当天零点
			zeroTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

			memberPackageCard.RemainDay = 1
			memberPackageCard.Status = MemberPackage_StatusNormal
			memberPackageCard.InvalidAt = &zeroTime
			_, err = uc.memberPackCardRepo.Update(ctx, memberPackageCard)
			if err != nil {
				log.Errorf("memberPackCardRepo Update error:%s \r\n", err)
				return false, err
			}
			useLog.Value = memberPackageCard.RemainDay
			useLog.Remark += fmt.Sprintf("次卡最后一次超时未入座,返还%d天,失效时间设为%v,", memberPackageCard.RemainDay, zeroTime)
			uc.memberPackCardRepo.RecordUseLog(ctx, useLog)
		}
	}
	return true, nil
}

// BookingValidate 订座确定时校验
func (uc *SeatUsecase) BookingValidate(ctx context.Context, req *BookingRequest) (ok bool, err *errors2.Error) {

	validate := &BookingValidate{
		uc:  uc,
		req: req,
		ctx: ctx,
	}
	validate.card.SrMemberPackageCard, _ = uc.memberPackCardRepo.FindByID(ctx, req.MemberPackageCardId)
	validate.seat, _ = uc.repo.FindByID(ctx, req.SeatId)
	validate.req.ShopAreaId = validate.seat.ShopAreaID
	validate.req.SeatLabel = validate.seat.Name

	if ok, err := validate.BaseValidate(false); !ok {
		return false, err
	}
	oriSeatId := req.SeatId
	defer func() {
		validate.req.SeatId = oriSeatId //换座校验后，还原req.seatId
	}()
	if req.ChangeSeatId == 0 {
		if ok, err := validate.repeatMemberBookingValidate(); !ok {
			return false, err
		}
	} else {
		//换座id不为空，检查新座位
		changeSeat, _ := uc.repo.FindByID(ctx, req.ChangeSeatId)
		if ok, err := validate.changeSeatValidate(changeSeat); !ok {
			return false, err
		}
		validate.seat = changeSeat
		validate.req.SeatId = changeSeat.ID
		validate.req.ChangeShopAreaId = changeSeat.ShopAreaID
		validate.req.SeatLabel = validate.seat.Name

	}
	if ok, err := validate.seatValidate(); !ok {
		return false, err
	}
	if ok, err := validate.repeatSeatBookingValidate(); !ok {
		return false, err
	}

	validate.shop, _ = uc.shopRepo.FindByID(ctx, req.ShopId)
	validate.req.ShopLabel = validate.shop.Name
	if ok, err := validate.shopValidate(); !ok {
		return false, err
	}

	if req.ChangeSeatId == 0 {
		if ok, err := validate.PackageCardValidate(); !ok {
			return false, err
		}
		req.BookingMinute = validate.card.BookingMinute
		req.BookingDay = validate.card.BookingDay
	}

	return true, nil
}

const ErrorCodeNoActive = 410
const ErrorCodeRemainMinNoEnough = 411
const ErrorCodeUseRemainMinute = 412

type BookingPreValidate struct {
	IsValid       bool
	Reason        string
	BookingMinute int32
	BookingDay    int32
	PackageCard   []*PackageCardValidate
}

type PackageCardValidate struct {
	Id            int64
	IsValid       bool
	InvalidCode   int32
	Reason        string
	BookingMinute int32
	BookingDay    int32
}

// BookingPreValidate 订座前校验
func (uc *SeatUsecase) BookingPreValidate(ctx context.Context, req *BookingRequest) (res *BookingPreValidate, err error) {

	res = &BookingPreValidate{IsValid: true}

	validate := &BookingValidate{
		uc:  uc,
		req: req,
		ctx: ctx,
	}
	if ok, err := validate.BaseValidate(true); !ok {
		res.IsValid = false
		res.Reason = err.GetReason()
		return res, nil
	}
	if req.ChangeSeatId == 0 {
		if ok, err := validate.repeatMemberBookingValidate(); !ok {
			res.IsValid = false
			res.Reason = err.GetReason()
			return res, nil
		}
	}
	validate.seat, _ = uc.repo.FindByID(ctx, req.SeatId)
	if ok, err := validate.repeatSeatBookingValidate(); !ok {
		res.IsValid = false
		res.Reason = err.GetReason()
		return res, nil
	}

	if ok, err := validate.seatValidate(); !ok {
		res.IsValid = false
		res.Reason = err.GetReason()
		return res, nil
	}
	validate.shop, _ = uc.shopRepo.FindByID(ctx, req.ShopId)
	if ok, err := validate.shopValidate(); !ok {
		res.IsValid = false
		res.Reason = err.GetReason()
		return res, nil
	}
	if req.ChangeSeatId != 0 {
		//换座id不为空，检查新座位
		changeSeat, _ := uc.repo.FindByID(ctx, req.ChangeSeatId)
		validate.seat = changeSeat
	}
	cardList, _ := uc.memberPackCardRepo.ListBy(ctx, &SearchMemberPackageCardRequest{MemberId: req.MemberId, Status: []int32{0, 1}})
	for _, card := range cardList {
		var cardErr *errors2.Error
		var ok bool
		validate.card = struct {
			*gen.SrMemberPackageCard
			BookingMinute int32
			BookingDay    int32
		}{
			SrMemberPackageCard: card,
			BookingMinute:       0,
			BookingDay:          0,
		}
		cardValid := &PackageCardValidate{
			Id:          card.ID,
			IsValid:     true,
			InvalidCode: 0,
		}
		ok, cardErr = validate.PackageCardValidate()
		cardValid.BookingMinute = validate.card.BookingMinute
		cardValid.BookingDay = validate.card.BookingDay
		if !ok {
			cardValid.IsValid = false
			cardValid.Reason = cardErr.GetReason()
			cardValid.InvalidCode = cardErr.GetCode()
			goto res
		}

	res:
		res.PackageCard = append(res.PackageCard, cardValid)

	}
	return res, nil
}
func (uc *SeatUsecase) BookingChangeSeatValidate(ctx context.Context, req *BookingRequest) (ok bool, err error) {
	if req.ChangeSeatId == 0 {
		err = errors.New(fmt.Sprintf("换座id错误"))
		return false, err
	}
	list, _ := uc.seatBookRepo.ListOperateLog(ctx, req.SeatBookId, []int8{SeatBook_OperateChangeSeat})
	if len(list) >= 2 && !req.CreateByAdmin {
		err = errors.New("换座次数已超过2次,请联系店员")
		return false, err
	}
	return true, nil
}

func (uc *SeatUsecase) BookingAddTimeValidate(ctx context.Context, req *BookingRequest) (ok bool, err error) {
	if req.AddTimeMin == 0 && req.AddTimeToCloseTime != true {
		err = errors.New(fmt.Sprintf("数据错误"))
		return false, err
	}
	cardInfo, err := uc.memberPackCardRepo.FindByID(ctx, req.MemberPackageCardId)

	if cardInfo.TypePackage != 1 {
		err = errors.New(fmt.Sprintf("时卡才能延长时间 "))
		return false, err
	}

	if cardInfo.TypePackage == 1 { //时卡
		req.BookingEnd = req.BookingEnd.Add(time.Minute * time.Duration(req.AddTimeMin))
	}
	if req.AddTimeToCloseTime {
		shopSetting, _ := config.GetShopSetting(req.ShopId)
		var closeTime string
		_ = json.Unmarshal([]byte(shopSetting["closeTime"]), &closeTime)
		req.BookingEnd, _ = time.ParseInLocation(time.TimeOnly, closeTime, time.Local)
	}

	return true, nil
}

func (uc *SeatUsecase) BookingChangeStatusValidate(ctx context.Context, req *BookingRequest) (ok bool, err error) {
	if req.NewStatus == SeatBook_StatusSeated { //入座
		now := time.Now()
		nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		if nowDay.Before(req.BookingStartDay) || nowDay.After(req.BookingEndDay) {
			return false, errors.New("不在入座日期内")
		}
		startAt, _ := time.ParseInLocation(time.DateTime, req.BookingStartDay.Format(time.DateOnly)+" "+req.BookingStart.Format(time.TimeOnly), time.Local)

		beforeMin15 := startAt.Add(-time.Minute * 30)
		if now.Before(beforeMin15) {
			return false, errors.New("只能提前30分钟入座")
		}
		if now.Before(startAt) {
			if req.BookingMinute > 0 {
				endAt := now.Add(time.Minute * time.Duration(req.BookingMinute)) //整段时间提前
				req.BookingEnd = endAt
			}
			req.BookingStart = now
			validate := &BookingValidate{
				uc:  uc,
				req: req,
				ctx: ctx,
			}
			//提前入座校验是否被订座
			validate.seat, _ = uc.repo.FindByID(ctx, req.SeatId)
			if ok, _ := validate.repeatSeatBookingValidate(); !ok {
				return false, errors.New("该时间段还有人在使用，请稍后")
			}
		}
	}
	if req.NewStatus == SeatBook_StatusCustomerCancel { //取消订单
		if !req.CreateByAdmin && req.Status != SeatBook_StatusWait { //未开始才能取消
			err = errors.New(fmt.Sprintf("订单已开始，不能手动取消，请联系管理员"))
			return false, err
		}
		now := time.Now()
		startAt, _ := time.ParseInLocation(time.DateTime, req.BookingStartDay.Format(time.DateOnly)+" "+req.BookingStart.Format(time.TimeOnly), time.Local)
		afterMin1 := startAt.Add(-time.Minute * 1)
		if !req.CreateByAdmin && now.After(afterMin1) {
			return false, errors.New("开始前一分钟内，不能取消订单")
		}
	}
	return true, nil
}

func (uc *SeatUsecase) BookingOpenDoorValidate(ctx context.Context, book *SeatBook) (ok bool, err error) {
	if book.Status != SeatBook_StatusSeated {
		return false, errors.New("请先点击报到再开门")
	}
	//now := time.Now()
	//startAt := *book.TodaySeatAt
	//endAt := *book.TodayEndAt
	//if now.After(startAt) && now.Before(endAt) {
	return true, nil
	//}
	//return false, errors.New("不在入座时间内")
}

func (uc *SeatUsecase) BookingOpenBigDoorValidate(ctx context.Context, book *SeatBook) (ok bool, err error) {
	now := time.Now()
	endTime, _ := time.ParseInLocation(time.DateTime, book.BookingEndDay.Format(time.DateOnly)+" "+book.BookingEnd, time.Local)
	startAt, _ := time.ParseInLocation(time.DateTime, book.BookingStartDay.Format(time.DateOnly)+" "+book.BookingStart, time.Local)

	beforeMin15 := startAt.Add(-time.Minute * 30)
	if now.Before(beforeMin15) {
		return false, errors.New("只能提前30分钟入座")
	}
	if now.After(endTime) {
		return false, errors.New("不在入座时间内")
	}
	return true, nil

}

func (uc *SeatUsecase) BookingControlLightValidate(ctx context.Context, book *SeatBook) (ok bool, err error) {
	if book.Status != SeatBook_StatusSeated {
		return false, errors.New("请先点击报到再开灯")
	}
	//now := time.Now()
	//startAt := *book.TodaySeatAt
	//endAt := *book.TodayEndAt
	//if now.After(startAt) && now.Before(endAt) {
	return true, nil
	//}
	//return false, errors.New("不在入座时间内")
}

func (uc *SeatUsecase) FindSeatBookSharedLog(ctx context.Context, id int64, openid string) (shareLog *gen.SrSeatBookShareLog, err error) {
	data, _, err := uc.seatBookRepo.FindSharedByID(ctx, id, openid)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *SeatUsecase) LockingControlValidate(ctx context.Context, seatLock *SeatLock) (ok bool, err error) {
	now := time.Now()
	nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	nowTime, _ := time.ParseInLocation(time.TimeOnly, now.Format(time.TimeOnly), time.Local)
	if seatLock != nil && seatLock.CanControl == 1 {
		if !(nowDay.Before(*seatLock.LockStartDay) || nowDay.After(*seatLock.LockEndDay)) { //当前时间不在锁定日期范围之外
			sTime, _ := time.ParseInLocation(time.TimeOnly, seatLock.LockStart, time.Local)
			eTime, _ := time.ParseInLocation(time.TimeOnly, seatLock.LockEnd, time.Local)
			if !(nowTime.Before(sTime) || nowTime.After(eTime)) { //订座时间不在锁定时间范围之外
				type useTime struct {
					Week   []int8
					AllDay bool
					Hour   []string
				}
				var uTime []useTime
				if seatLock.UseTime != "" {
					var inWeek bool
					var inHour bool

					err2 := json.Unmarshal([]byte(seatLock.UseTime), &uTime)
					if err2 != nil {
						return false, errors2.New(400, err.Error(), "")
					}
					for _, u := range uTime {
						inWeek = false
						if _func.InInt8Splice(u.Week, int8(nowDay.Weekday())) {
							inWeek = true //只要有一个在，就要判断
						}
						if inWeek {
							inHour = false
							if u.AllDay == true {
								inHour = true
							}
						}
					}

					if inWeek && inHour {
						return true, nil
					}
				}
			}
		}
	}
	return false, errors.New("不在锁位时间内或无权控制")
}

type BookingValidate struct {
	uc  *SeatUsecase
	req *BookingRequest

	card struct {
		*gen.SrMemberPackageCard
		BookingMinute int32
		BookingDay    int32
	}
	seat *model.Seat
	shop *gen.SrShop
	ctx  context.Context
}

func (v *BookingValidate) BaseValidate(isPre bool) (ok bool, err *errors2.Error) {
	if v.req.SeatId == 0 || v.req.MemberId == 0 || v.req.ShopId == 0 || v.req.BookingType == 0 {
		err = errors2.New(400, fmt.Sprintf("数据错误"), "")
		return false, err
	}
	if !isPre {
		if v.req.MemberPackageCardId == 0 {
			err = errors2.New(400, fmt.Sprintf("请选择套餐"), "")
			return false, err
		}
	}
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	if date.After(v.req.BookingStartDay) || date.After(v.req.BookingEndDay) {
		err = errors2.New(400, fmt.Sprintf("预定日期小于当前日期"), "")
		return false, err
	}
	if v.req.BookingStartDay.After(v.req.BookingEndDay) {
		err = errors2.New(400, fmt.Sprintf("开始日期小于结束日期"), "")
		return false, err
	}

	if v.req.BookingStart.After(v.req.BookingEnd) || v.req.BookingStart.Equal(v.req.BookingEnd) {
		err = errors2.New(400, fmt.Sprintf("开始时间大于等于结束时间"), "")
		return false, err
	}
	if v.req.BookingStart.Day() != v.req.BookingEnd.Day() {
		err = errors2.New(400, fmt.Sprintf("结束时间超过当天时间"), "")
		return false, err
	}
	if len(v.req.BookingWeek) <= 0 || len(v.req.BookingWeek) > 7 {
		err = errors2.New(400, fmt.Sprintf("请选择订座星期"), "")
		return false, err
	}
	if len(v.req.BookingWeek) > 0 && len(v.req.BookingWeek) <= 7 {

		if v.req.ValidDay == 0 {
			err = errors2.New(400, fmt.Sprintf("请确定有效订座日期"), "")
			return false, err
		}

	}
	return true, nil
}
func (v *BookingValidate) PackageCardValidate() (ok bool, err *errors2.Error) {

	if strings.Contains(v.card.UseShop, "shopId") {
		i := fmt.Sprintf("|%d|", v.shop.ID)
		if !strings.Contains(v.card.UseShop, i) && !strings.Contains(v.card.UseShop, "|0|") {
			err = errors2.New(400, fmt.Sprintf("指定店铺不可用"), "")
			return false, err
		}
	}

	if strings.Contains(v.card.UseShop, "shopType") {
		i := fmt.Sprintf("|%d|", v.shop.ID)

		if !strings.Contains(v.card.UseShop, i) && !strings.Contains(v.card.UseShop, "|0|") {
			err = errors2.New(400, fmt.Sprintf("指定店铺类型不可用"), "")
			return false, err
		}
	}

	if !strings.Contains(v.card.UseSeat, fmt.Sprintf("\"%d\"", v.seat.TypeSeat)) && !strings.Contains(v.card.UseSeat, "\"0\"") {
		err = errors2.New(400, fmt.Sprintf("座位类型不可用"), "")
		return false, err
	}

	if ok, err := v.TimesPackageTypeValidate(); !ok {
		return false, err
	}

	if ok, err := v.HoursPackageTypeValidate(); !ok {
		return false, err
	}

	if ok, err := v.DayPackageTypeValidate(); !ok {
		return false, err
	}

	if v.card.SrMemberPackageCard == nil {
		err = errors2.New(400, fmt.Sprintf("套餐不存在"), "")
		return false, err
	}
	if v.card.Status == 0 {
		err = errors2.New(ErrorCodeNoActive, fmt.Sprintf("套餐未激活"), "")
		return false, err
	}
	if v.card.Status == 2 {
		err = errors2.New(400, fmt.Sprintf("套餐已过期"), "")
		return false, err
	}
	if v.card.Status == 3 {
		err = errors2.New(400, fmt.Sprintf("套餐已用完"), "")
		return false, err
	}
	if v.card.Status == 4 {
		err = errors2.New(400, fmt.Sprintf("套餐已冻结"), "")
		return false, err
	}
	if v.card.InvalidAt.Year() < 2000 {
		err = errors2.New(400, fmt.Sprintf("套餐未激活"), "")
		return false, err
	}

	now := time.Now()
	startAt, _ := time.ParseInLocation(time.DateTime, v.req.BookingStartDay.Format(time.DateOnly)+" "+v.req.BookingStart.Format(time.TimeOnly), time.Local)

	ActiveAtAdd10 := v.card.ActiveAt.Add(-10 * time.Second) //激活时间10秒前
	if startAt.Before(ActiveAtAdd10) {                      //开始时间在激活时间-10秒前，防止订座时激活失败
		err = errors2.New(400, fmt.Sprintf("订座时间小于激活时间"), "")
		return false, err
	}
	if now.After(*v.card.InvalidAt) {
		err = errors2.New(400, fmt.Sprintf("套餐已超过有效期"), "")
		return false, err
	}

	type useTime struct {
		Week   []int8
		AllDay bool
		Hour   []string
	}
	var uTime []useTime
	if v.card.UseTime != "" {
		var inWeek bool
		var inHour bool
		var canHour struct {
			s string
			e string
		}
		var canWeek []int8
		//sWeek := v.req.BookingStartDay.Weekday()
		//eWeek := v.req.BookingEndDay.Weekday()
		err2 := json.Unmarshal([]byte(v.card.UseTime), &uTime)
		if err2 != nil {
			return false, errors2.New(400, err.Error(), "")
		}
		for _, u := range uTime {
			canWeek = u.Week
			if !u.AllDay {
				canHour.s = u.Hour[0]
				canHour.e = u.Hour[1]
			}
			inWeek = true
			diffDay := int32(math.Ceil(v.req.BookingEndDay.Sub(v.req.BookingStartDay).Hours() / 24))

			if diffDay >= 0 {
				for i := int32(0); i <= diffDay; i++ {
					nextDay := v.req.BookingStartDay.Add(time.Hour * 24 * time.Duration(i))
					if _func.InInt32Splice(v.req.BookingWeek, int32(nextDay.Weekday())) {
						if !_func.InInt8Splice(u.Week, int8(nextDay.Weekday())) {
							inWeek = false
							break
						}
					}
				}
			}
			if inWeek {
				inHour = false
				if u.AllDay == true {
					inHour = true
				} else {

					sTime, _ := time.ParseInLocation(time.TimeOnly, u.Hour[0], time.Local)
					eTime, _ := time.ParseInLocation(time.TimeOnly, u.Hour[1], time.Local)
					if (v.req.BookingStart.After(sTime) || v.req.BookingStart.Equal(sTime)) && (v.req.BookingEnd.Before(eTime) || v.req.BookingEnd.Equal(eTime)) {
						inHour = true
						break
					}
				}
			}
		}

		if !inWeek || !inHour {
			if canHour.s == "" {
				err = errors2.New(400, fmt.Sprintf("适用时段:星期%v,营业时间", canWeek), "")
			} else {
				err = errors2.New(400, fmt.Sprintf("适用时段:星期%v,%s - %s", canWeek, canHour.s, canHour.e), "")
			}
			return false, err
		}
	}

	startTime := int(v.req.BookingStartDay.Unix()) + v.req.BookingStart.Hour()*3600 + v.req.BookingStart.Minute()*60
	endTime := int(v.req.BookingEndDay.Unix()) + v.req.BookingEnd.Hour()*3600 + v.req.BookingEnd.Minute()*60
	invalidAt := int(v.card.InvalidAt.Unix())
	if startTime > invalidAt {
		err = errors2.New(400, fmt.Sprintf("预定开始时间已超过套餐有效期"), "")
		return false, err
	}
	if endTime > invalidAt {
		err = errors2.New(400, fmt.Sprintf("预定结束时间已超过套餐有效期"), "")
		return false, err
	}

	return true, nil
}

// TimesPackageTypeValidate   次卡校验
func (v *BookingValidate) TimesPackageTypeValidate() (ok bool, err *errors2.Error) {
	if v.card.TypePackage != 2 {
		return true, nil
	}
	bookType := v.req.BookingType
	if bookType == 1 { //立马入座
		v.card.BookingDay = 1

	}

	if bookType == 2 { //当天预定
		v.card.BookingDay = 1
	}

	if bookType == 3 { //连续预定

		v.card.BookingDay = v.req.ValidDay
	}
	if v.card.RemainDay <= 0 || (v.card.RemainDay > 0 && v.card.RemainDay < int(v.card.BookingDay)) {
		err = errors2.New(400, fmt.Sprintf("超出剩余可订天数"), "")
		return false, err
	}
	if v.card.MinutePer > 0 {
		startTime := v.req.BookingStart.Hour()*3600 + v.req.BookingStart.Minute()*60
		endTime := v.req.BookingEnd.Hour()*3600 + v.req.BookingEnd.Minute()*60

		s := float64(endTime - startTime) //相差多少秒
		m := math.Floor(s / 60)           //分钟 //向下取整
		diffMinute := int(m)
		if diffMinute > v.card.MinutePer {
			err = errors2.New(400, fmt.Sprintf("超出每次可用分钟"), "")
			return false, err
		}
	}
	return true, nil
}

// HoursPackageTypeValidate  时卡校验
func (v *BookingValidate) HoursPackageTypeValidate() (ok bool, err *errors2.Error) {
	if v.card.TypePackage != 1 {
		return true, nil
	}
	diffDay := int32(math.Ceil(v.req.BookingEndDay.Sub(v.req.BookingStartDay).Hours() / 24)) //相差多少天

	startTime := v.req.BookingStart.Hour()*3600 + v.req.BookingStart.Minute()*60
	endTime := v.req.BookingEnd.Hour()*3600 + v.req.BookingEnd.Minute()*60

	s := float64(endTime - startTime) //相差多少秒
	m := math.Floor(s / 60)           //小时 //向下取整
	diffMinute := m
	bookType := v.req.BookingType
	if bookType == 1 { //立马入座
		v.card.BookingMinute = int32(diffMinute)
	}

	if bookType == 2 { //当天预定
		v.card.BookingMinute = int32(diffMinute) //当天为1次
	}

	if bookType == 3 { //连续预定
		v.card.BookingMinute = int32(diffMinute) * v.req.ValidDay
	}
	minBookMin := GetMinBookMin(v.req.ShopId)
	if v.card.BookingMinute < int32(minBookMin) && v.req.AddTimeMin == 0 {
		err = errors2.New(400, fmt.Sprintf("订座分钟少于最低订座%v分钟", minBookMin), "")
		return false, err
	}
	if diffDay == 0 && v.req.UseRemain != 1 && v.card.RemainMinute < minBookMin {
		v.card.BookingMinute = int32(v.card.RemainMinute)
		err = errors2.New(ErrorCodeUseRemainMinute, fmt.Sprintf("剩余时间不足%d分钟", minBookMin), "")
		return false, err
	}
	if v.req.UseRemain == 1 { //使用剩余时间订座
		v.card.BookingMinute = int32(v.card.RemainMinute)
		v.req.BookingEnd = v.req.BookingStart.Add(time.Minute * time.Duration(v.card.BookingMinute))
	}
	if v.card.RemainMinute <= 0 || (v.card.RemainMinute > 0 && v.card.RemainMinute < int(v.card.BookingMinute)) {
		err = errors2.New(400, fmt.Sprintf("超出剩余可订分钟数"), "")
		return false, err
	}

	//tt := v.card.RemainMinute - int(v.card.BookingMinute)
	//if v.req.NoEnoughUse == 0 && tt > 0 && tt < 60 {
	//	err = errors2.New(ErrorCodeRemainMinNoEnough, fmt.Sprintf("本次消费后剩余时间不足60分钟"), "")
	//	return false, err
	//}
	return true, nil
}

// DayPackageTypeValidate   天卡校验
func (v *BookingValidate) DayPackageTypeValidate() (ok bool, err *errors2.Error) {
	if v.card.TypePackage != 3 {
		return true, nil
	}

	bookType := v.req.BookingType

	if bookType == 1 { //立马入座
		v.card.BookingDay = 1
	}

	if bookType == 2 { //当天预定
		v.card.BookingDay = 1
	}

	if bookType == 3 { //连续预定
		v.card.BookingDay = v.req.ValidDay

	}
	return true, nil
}

// 座位校验
func (v *BookingValidate) seatValidate() (ok bool, err *errors2.Error) {
	if v.seat.Status == 1 {
		err = errors2.New(400, fmt.Sprintf("当前座位不可用"), "")
		return false, err
	}

	seatLock, _ := v.uc.repo.FindLockingBy(v.ctx, &SearchSeatLockingRequest{ //锁位检查
		SeatId:     v.req.SeatId,
		OrMemberId: v.req.MemberId,
		ShopId:     v.seat.ShopID,
		Status:     []int32{1},
	})

	if seatLock != nil {
		if !(v.req.BookingEndDay.Before(*seatLock.LockStartDay) || v.req.BookingStartDay.After(*seatLock.LockEndDay)) { //订座时间不在锁定日期范围之外
			sTime, _ := time.ParseInLocation(time.TimeOnly, seatLock.LockStart, time.Local)
			eTime, _ := time.ParseInLocation(time.TimeOnly, seatLock.LockEnd, time.Local)
			if !((v.req.BookingStart.Before(sTime) && (v.req.BookingEnd.Before(sTime) || v.req.BookingEnd.Equal(sTime))) ||
				(v.req.BookingStart.After(eTime) || v.req.BookingStart.Equal(eTime)) && v.req.BookingEnd.After(eTime)) { //订座时间不在锁定时间范围之外
				type useTime struct {
					Week   []int8
					AllDay bool
					Hour   []string
				}
				var uTime []useTime
				if seatLock.UseTime != "" {
					var inWeek bool
					var inHour bool

					err2 := json.Unmarshal([]byte(seatLock.UseTime), &uTime)
					if err2 != nil {
						return false, errors2.New(400, err.Error(), "")
					}
					for _, u := range uTime {
						inWeek = false
						diffDay := int32(math.Ceil(v.req.BookingEndDay.Sub(v.req.BookingStartDay).Hours() / 24))

						if diffDay >= 0 {
							for i := int32(0); i <= diffDay; i++ {
								nextDay := v.req.BookingStartDay.Add(time.Hour * 24 * time.Duration(i))
								if _func.InInt32Splice(v.req.BookingWeek, int32(nextDay.Weekday())) { //在订座的星期内
									if _func.InInt8Splice(u.Week, int8(nextDay.Weekday())) { //在锁位的星期内
										inWeek = true //只要有一个在，就要判断
										break
									}
								}
							}
						}

						if inWeek {
							inHour = false
							if u.AllDay == true {
								inHour = true
							}
						}

					}

					if inWeek && inHour {
						if seatLock.MemberID == v.req.MemberId && seatLock.SeatID != v.req.SeatId {
							err = errors2.New(400, fmt.Sprintf("该时间段您已锁定座位，不能选择其他座位"), "")
							return false, err
						}

						if seatLock.SeatID == v.req.SeatId && seatLock.MemberID != v.req.MemberId {
							err = errors2.New(400, fmt.Sprintf("该时间段座位已被锁定，请重新选择其他座位"), "")
							return false, err
						}
					}
				} else {
					if seatLock.MemberID == v.req.MemberId && seatLock.SeatID != v.req.SeatId {
						err = errors2.New(400, fmt.Sprintf("该时间段您已锁定座位，不能选择其他座位"), "")
						return false, err
					}

					if seatLock.SeatID == v.req.SeatId && seatLock.MemberID != v.req.MemberId {
						err = errors2.New(400, fmt.Sprintf("该时间段座位已被锁定，请重新选择其他座位"), "")
						return false, err
					}
				}
			}
		}
	}
	return true, nil
}

// 店铺校验
func (v *BookingValidate) shopValidate() (ok bool, err *errors2.Error) {
	shopSetting, _ := config.GetShopSetting(v.shop.ID)

	var closeTime string
	var openTime string
	bookDateS := v.req.BookingStartDay
	bookDateE := v.req.BookingEndDay
	bookTimeS := v.req.BookingStart
	bookTimeE := v.req.BookingEnd
	_ = json.Unmarshal([]byte(shopSetting["closeTime"]), &closeTime)
	_ = json.Unmarshal([]byte(shopSetting["openTime"]), &openTime)
	oTime, _ := time.ParseInLocation(time.TimeOnly, openTime, time.Local)
	cTime, _ := time.ParseInLocation(time.TimeOnly, closeTime, time.Local)
	if bookTimeS.Before(oTime) || bookTimeE.Before(oTime) || bookTimeE.After(cTime) {
		err = errors2.New(400, fmt.Sprintf("该时间段不在营业时间范围"), "")
		return false, err
	}
	isSetBreak := shopSetting["isSetBreak"]
	if isSetBreak == "1" {
		var breakDate []string
		var breakTime []string
		_ = json.Unmarshal([]byte(shopSetting["breakDate"]), &breakDate)
		_ = json.Unmarshal([]byte(shopSetting["breakTime"]), &breakTime)
		breakDateS, _ := time.ParseInLocation(time.DateOnly, breakDate[0], time.Local)
		breakDateE, _ := time.ParseInLocation(time.DateOnly, breakDate[1], time.Local)
		breakTimeS, _ := time.ParseInLocation(time.TimeOnly, breakTime[0], time.Local)
		breakTimeE, _ := time.ParseInLocation(time.TimeOnly, breakTime[1], time.Local)
		if (breakDateS.Before(bookDateE) || breakDateS.Equal(bookDateE) || breakDateS.Equal(bookDateS)) &&
			(breakDateE.After(bookDateE) || breakDateE.After(bookDateS) || breakDateE.Equal(bookDateE) || breakDateE.Equal(bookDateS)) {
			if breakTimeS.Before(bookTimeE) && (breakTimeE.After(bookTimeE) || breakTimeE.After(bookTimeS)) {
				err = errors2.New(400, fmt.Sprintf("该时间段休息,请重新选择订座"), "")
				return false, err
			}
		}
	}
	return true, nil

}

// repeatMemberBookingValidate 用户重复订座校验
func (v *BookingValidate) repeatMemberBookingValidate() (ok bool, err *errors2.Error) {
	booking, _ := v.uc.seatBookRepo.CheckMemberCanBooking(v.ctx, v.req)
	if booking != nil && booking.ID != 0 {
		//startTime := v.req.BookingStartDay.Unix() + int64(v.req.BookingStart.Hour()*3600+v.req.BookingStart.Minute()*60)
		//endTime := v.req.BookingEndDay.Unix() + int64(v.req.BookingEnd.Hour()*3600+v.req.BookingEnd.Minute()*60)
		//t1 := time.Unix(startTime, 0).Format(time.DateTime)
		//t2 := time.Unix(endTime, 0).Format(time.DateTime)
		seat, _ := v.uc.repo.FindByID(v.ctx, booking.SeatID)

		err = errors2.New(400, fmt.Sprintf("您在这时间段已经订座【%s】,请重新选择", seat.Name), "")
		return false, err
	}

	return true, nil
}

// repeatSeatBookingValidate 座位重复订座校验
func (v *BookingValidate) repeatSeatBookingValidate() (ok bool, err *errors2.Error) {
	startTime := v.req.BookingStartDay.Unix() + int64(v.req.BookingStart.Hour()*3600+v.req.BookingStart.Minute()*60)
	endTime := v.req.BookingEndDay.Unix() + int64(v.req.BookingEnd.Hour()*3600+v.req.BookingEnd.Minute()*60)
	t1 := time.Unix(startTime, 0).Format(time.TimeOnly)
	t2 := time.Unix(endTime, 0).Format(time.TimeOnly)

	booking, _ := v.uc.seatBookRepo.CheckSeatCanBooking(v.ctx, v.req)
	if booking != nil && booking.ID != 0 {
		err = errors2.New(400, fmt.Sprintf("该订座时间段 %s - %s 已经被订座,请重新选择", t1, t2), "")
		return false, err
	}

	//// 检查父座位
	//if v.seat.Pid != 0 { //有父座位,检查父座位
	//	newReq := v.req
	//	newReq.SeatId = v.seat.Pid
	//	booking, _ := v.uc.seatBookRepo.CheckSeatCanBooking(v.ctx, newReq)
	//	if booking != nil && booking.ID != 0 && (v.req.SeatBookId != 0 && booking.ID != v.req.SeatBookId) {
	//		err = errors.New(fmt.Sprintf("该区域在时间段 %s - %s 已经被订座,请重新选择", t1, t2))
	//		return false, err
	//	}
	//}
	//
	////检查子座位
	//var child struct {
	//	ids  []int64
	//	name map[int64]string
	//}
	//child.name = make(map[int64]string)
	//childSeat, _ := v.uc.repo.ListBy(v.ctx, &SearchSeatRequest{Pid: v.req.SeatId})
	//for _, i2 := range childSeat {
	//	child.ids = append(child.ids, i2.ID)
	//	child.name[i2.ID] = i2.Name
	//
	//}
	//if len(child.ids) > 0 {
	//	v.req.ChildIds = child.ids
	//	booking, _ := v.uc.seatBookRepo.CheckChildrenSeatCanBooking(v.ctx, v.req)
	//	if booking != nil && booking.ID != 0 && (v.req.SeatBookId != 0 && booking.ID != v.req.SeatBookId) {
	//		err = errors.New(fmt.Sprintf("该区域在时间段 %s - %s 已经有座位【%s】被订座,请重新选择", t1, t2, child.name[booking.SeatID]))
	//		return false, err
	//	}
	//}
	return true, nil
}

// ChangeSeatValidate 换座校验
func (v *BookingValidate) changeSeatValidate(newSeat *model.Seat) (ok bool, err *errors2.Error) {
	if !strings.Contains(v.card.UseSeat, fmt.Sprintf("\"%d\"", newSeat.TypeSeat)) && !strings.Contains(v.card.UseSeat, "\"0\"") {
		err = errors2.New(400, fmt.Sprintf("套餐不支持该座位类型"), "")
		return false, err
	}
	return true, nil
}

func (uc *SeatUsecase) RecordLockingControlLog(ctx context.Context, seatLock *SeatLock) {
	operateLog := &SeatBookOperateLog{
		SeatBookId: 0,
		ShopId:     seatLock.ShopID,
		SeatId:     seatLock.SeatID,
		MemberId:   seatLock.MemberID,
		Operate:    SeatBook_LockingControlDoor,
		Remark:     fmt.Sprintf("锁位,入座开门"),
		Extra:      nil,
	}

	uc.seatBookRepo.RecordOperateLog(ctx, operateLog)
}

func GenerateUniqueToken(shopId int64, customerId int64) string {

	customerStr := strconv.Itoa(int(customerId)) //后四位取模
	customerStr = fmt.Sprintf("%04s", customerStr)

	shopStr := strconv.Itoa(int(shopId)) //后四位取模
	shopStr = fmt.Sprintf("%02s", shopStr)

	timeUnix := strconv.Itoa(int(time.Now().Unix())) //单位秒
	cIndex := customerStr[len(customerStr)-4:]
	randStr := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)) //这里面前面的04v是和后面的1000相对应的
	var build strings.Builder
	build.Grow(20)
	build.WriteString(shopStr)
	build.WriteString(timeUnix)
	build.WriteString(cIndex)
	build.WriteString(randStr)
	tag := build.String()
	return tag
}

func GetMinBookMin(shopId int64) int {
	shopSetting, _ := config.GetShopSetting(shopId)
	str := ""
	_ = json.Unmarshal([]byte(shopSetting["minBookMin"]), &str)
	minBookMin, _ := strconv.Atoi(str)
	return minBookMin
}

func GetNotSeatMinCancel(shopId int64) int {
	shopSetting, _ := config.GetShopSetting(shopId)
	str := ""
	_ = json.Unmarshal([]byte(shopSetting["notSeatMinCancel"]), &str)
	notSeatMinCancel, _ := strconv.Atoi(str)
	return notSeatMinCancel
}
