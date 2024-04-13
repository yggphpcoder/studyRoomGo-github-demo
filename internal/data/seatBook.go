package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"strconv"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
	"time"
)

type SeatBookRepo struct {
	data *Data
	log  *log.Helper
}

// NewSeatRepo .
func NewSeatBookRepo(data *Data, logger log.Logger) biz.SeatBookRepo {
	return &SeatBookRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *SeatBookRepo) Save(ctx context.Context, book *biz.SeatBook) (bool, error) {
	mgr := gen.SrSeatBookMgr(r.data.db)
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		book.SrSeatBook.MerchantID = int64(merchantId)
	}
	mgr.Save(book.SrSeatBook)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *SeatBookRepo) Update(ctx context.Context, book *biz.SeatBook) (bool, error) {
	mgr := gen.SrSeatBookMgr(r.data.db)
	if err := mgr.Select("*").Omit(Omit...).Updates(&book.SrSeatBook).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *SeatBookRepo) FindByID(ctx context.Context, id int64) (*model.SeatBook, error) {
	mgr := gen.SrSeatBookMgr(r.data.db)
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	var modelSeatBook model.SeatBook

	err := mgr.First(&modelSeatBook, id).Error

	if err != nil {
		return nil, err
	}
	return &modelSeatBook, nil
}

func (r *SeatBookRepo) FindDetailByID(ctx context.Context, id int64) (*model.SeatBook, error) {
	mgr := gen.SrSeatBookMgr(r.data.db)
	tableName := mgr.GetTableName()

	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatBookColumns.MerchantID), merchantId)
	}
	var modelSeatBook *model.SeatBook

	mgr.Scopes(modelSeatBook.ScopesJoin)
	err := mgr.First(&modelSeatBook, id).Error

	if err != nil {
		return nil, err
	}
	return modelSeatBook, nil
}

func (r *SeatBookRepo) ListBy(ctx context.Context, req *biz.SearchSeatBookingRequest) ([]*model.SeatBook, error) {
	mgr := gen.SrSeatBookMgr(r.data.db)
	merchantId := 0
	if req.MerchantId != 0 {
		merchantId = int(req.MerchantId)
	} else {
		if md, ok := metadata.FromServerContext(ctx); ok {
			merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
		}

	}

	var modelSeatBook *model.SeatBook

	tableName := mgr.GetTableName()
	mgr.Scopes(modelSeatBook.ScopesJoin)
	if merchantId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatBookColumns.MerchantID), merchantId)
	}
	if req.MemberId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatBookColumns.MemberID), req.MemberId)

	}
	if req.ShopId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatBookColumns.ShopID), req.ShopId)

	}
	if len(req.Status) != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v in ?", tableName, gen.SrSeatColumns.Status), req.Status)
	}
	if req.Sort > 0 {
		mgr.Order(SortType(req.Sort, tableName))
	} else {
		mgr.Order(SortType(sortBookStartDayASort, tableName))
	}
	if req.PageSize > 0 && req.Page > 0 {
		mgr.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	var modelList []*model.SeatBook
	err := mgr.Scan(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}

func (r *SeatBookRepo) CheckMemberCanBooking(ctx context.Context, req *biz.BookingRequest) (*gen.SrSeatBook, error) {
	var seatBook *gen.SrSeatBook
	mgr := gen.SrSeatBookMgr(r.data.db)
	modelSeatBook := &model.SeatBook{
		SrSeatBook: gen.SrSeatBook{
			SeatID:          req.SeatId,
			ShopID:          req.ShopId,
			MemberID:        req.MemberId,
			BookingStartDay: req.BookingStartDay,
			BookingEndDay:   req.BookingEndDay,
			BookingStart:    req.BookingStart.Format(time.TimeOnly),
			BookingEnd:      req.BookingEnd.Format(time.TimeOnly),
		},
		BookingWeek: req.BookingWeek,
	}
	mgr.Scopes(modelSeatBook.ScopesMemberCanBooking())
	err := mgr.Last(&seatBook).Error

	if err != nil {
		return nil, err

	}
	return seatBook, nil
}

func (r *SeatBookRepo) CheckSeatCanBooking(ctx context.Context, req *biz.BookingRequest) (*gen.SrSeatBook, error) {
	var seatBook *gen.SrSeatBook
	mgr := gen.SrSeatBookMgr(r.data.db)
	modelSeatBook := &model.SeatBook{
		SrSeatBook: gen.SrSeatBook{
			SeatID:          req.SeatId,
			ShopID:          req.ShopId,
			MemberID:        req.MemberId,
			BookingStartDay: req.BookingStartDay,
			BookingEndDay:   req.BookingEndDay,
			BookingStart:    req.BookingStart.Format(time.TimeOnly),
			BookingEnd:      req.BookingEnd.Format(time.TimeOnly),
		},
		BookingWeek: req.BookingWeek,
	}
	mgr.Scopes(modelSeatBook.ScopesSeatCanBooking())
	err := mgr.Last(&seatBook).Error

	if err != nil {
		return nil, err

	}
	return seatBook, nil
}

func (r *SeatBookRepo) CheckChildrenSeatCanBooking(ctx context.Context, req *biz.BookingRequest) (*gen.SrSeatBook, error) {
	var seatBook *gen.SrSeatBook
	mgr := gen.SrSeatBookMgr(r.data.db)
	modelSeatBook := &model.SeatBook{
		SrSeatBook: gen.SrSeatBook{
			SeatID:          req.SeatId,
			ShopID:          req.ShopId,
			MemberID:        req.MemberId,
			BookingStartDay: req.BookingStartDay,
			BookingEndDay:   req.BookingEndDay,
			BookingStart:    req.BookingStart.Format(time.TimeOnly),
			BookingEnd:      req.BookingEnd.Format(time.TimeOnly),
		},
		ChildIds: req.ChildIds,
	}
	mgr.Scopes(modelSeatBook.ScopesChildrenSeat())
	err := mgr.Last(&seatBook).Error

	if err != nil {
		return nil, err

	}
	return seatBook, nil
}

func (r *SeatBookRepo) RecordOperateLog(ctx context.Context, log *biz.SeatBookOperateLog) {
	mgr := gen.SrSeatBookOperateLogMgr(r.data.db)
	data := &gen.SrSeatBookOperateLog{
		SeatBookID: log.SeatBookId,
		MemberID:   log.MemberId,
		SeatID:     log.SeatId,
		ShopID:     log.ShopId,
		Action:     log.Operate,
		Extra:      extra(log.Extra),
		Remark:     &log.Remark,
	}
	if !log.CreateByAdmin {
		creatBy := uint(log.MemberId)
		remark := fmt.Sprintf("用户操作｜ %s", *data.Remark)
		data.Remark = &remark
		data.CreateBy = &creatBy

	} else {
		creatBy := uint(0)
		remark := fmt.Sprintf("管理员操作｜ %s", *data.Remark)
		data.Remark = &remark
		data.CreateBy = &creatBy
	}
	mgr.Save(data)

}
func (r *SeatBookRepo) ListOperateLog(ctx context.Context, seatBookId int64, status []int8) ([]*gen.SrSeatBookOperateLog, error) {
	mgr := gen.SrSeatBookOperateLogMgr(r.data.db)

	if seatBookId != 0 {
		mgr.Where("seat_book_id = ?", seatBookId)

	}
	if len(status) != 0 {
		mgr.Where("action in ?", status)
	}

	var list []*gen.SrSeatBookOperateLog
	err := mgr.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *SeatBookRepo) FindSharedByID(ctx context.Context, seatBookId int64, openId string) (res *gen.SrSeatBookShareLog, total int64, err error) {
	mgr := gen.SrSeatBookShareLogMgr(r.data.db)
	mgr.Where("seat_book_id = ?", seatBookId)
	mgr.Count(&total)
	mgr2 := gen.SrSeatBookShareLogMgr(r.data.db)
	mgr2.Where("seat_book_id = ?", seatBookId)
	mgr2.Where("openid = ?", openId)
	data, err := mgr2.Get()
	if err != nil {
		return nil, total, err
	}
	return &data, total, nil
}

func (r *SeatBookRepo) SaveSharedLog(ctx context.Context, data *gen.SrSeatBookShareLog) (bool, error) {
	mgr := gen.SrSeatBookShareLogMgr(r.data.db)
	mgr.Save(&data)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}
