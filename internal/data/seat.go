package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
)

type SeatRepo struct {
	data *Data
	log  *log.Helper
}

// NewSeatRepo .
func NewSeatRepo(data *Data, logger log.Logger) biz.SeatRepo {
	return &SeatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *SeatRepo) FindByID(ctx context.Context, id int64) (*model.Seat, error) {
	mgr := gen.SrSeatMgr(r.data.db)
	var modelSeat *model.Seat

	mgr.Scopes(modelSeat.ScopesJoin)
	err := mgr.First(&modelSeat, id).Error

	if err != nil {
		return nil, err
	}
	return modelSeat, nil
}

func (r *SeatRepo) FindBy(ctx context.Context, key string, value string) (*model.Seat, error) {
	mgr := gen.SrSeatMgr(r.data.db)
	var modelSeat *model.Seat

	tableName := mgr.GetTableName()
	mgr.Scopes(modelSeat.ScopesJoin)

	mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, key), value)

	err := mgr.Scan(&modelSeat).Error
	if err != nil {
		return nil, err
	}
	return modelSeat, nil
}

func (r *SeatRepo) ListBy(ctx context.Context, req *biz.SearchSeatRequest) ([]*model.Seat, error) {
	mgr := gen.SrSeatMgr(r.data.db)
	var modelSeat *model.Seat

	tableName := mgr.GetTableName()
	mgr.Scopes(modelSeat.ScopesJoin)
	mgr.Order(SortType(req.Sort, tableName))
	if req.Name != "" {
		str := fmt.Sprintf("%%%v%%", req.Name)
		mgr.Where(fmt.Sprintf("`%v`.%v like ?", tableName, gen.SrSeatColumns.Name), str)
	}
	if req.Pid != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatColumns.Pid), req.Pid)

	}
	if req.ShopId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatColumns.ShopID), req.ShopId)

	}
	if req.ShopAreaId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatColumns.ShopAreaID), req.ShopAreaId)

	}
	if req.Status != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatColumns.Status), req.Status)

	}

	if req.TypeSeat != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatColumns.TypeSeat), req.TypeSeat)

	}
	var modelList []*model.Seat
	err := mgr.Scan(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}

func (r *SeatRepo) ListAll(context.Context) ([]*model.Seat, error) {
	var modelList []*model.Seat
	var modelSeat *model.Seat

	mgr := gen.SrSeatMgr(r.data.db)
	mgr.Scopes(modelSeat.ScopesJoin)
	err := mgr.Scan(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil
}

func (r *SeatRepo) ListBookingBy(ctx context.Context, req *biz.SearchSeatBookingRequest) ([]*gen.SrSeatBook, error) {
	condition := gen.Condition{}
	var modelSeat *model.Seat

	if req.SeatId != 0 {
		condition.And(gen.SrSeatBookColumns.SeatID, "=", req.SeatId)
	}
	if req.ShopId != 0 {
		condition.And(gen.SrSeatColumns.ShopID, "=", req.ShopId)
	}
	if req.Status != nil {
		condition.And(gen.SrSeatBookColumns.Status, "in", req.Status)
	}

	where, out := condition.Get()
	mgr := gen.SrSeatBookMgr(r.data.db)
	mgr.Scopes(
		modelSeat.ScopesDateTime("<=", "booking_end_day", req.GteBookDateEnd),
		modelSeat.ScopesDateTime("<=", "booking_end", req.GteBookTimeEnd),
		modelSeat.ScopesDateTime(">=", "booking_end_day", req.LteBookDateEnd),
		modelSeat.ScopesDateTime(">=", "booking_end", req.LteBookTimeEnd),
		modelSeat.ScopesDateTime("<=", "booking_start_day", req.GteBookDateStart),
		modelSeat.ScopesDateTime("<=", "booking_start", req.GteBookTimeStart),
		modelSeat.ScopesDateTime(">=", "booking_start_day", req.LteBookDateStart),
		modelSeat.ScopesDateTime(">=", "booking_start", req.LteBookTimeStart),
	).Where(where, out...).Order("booking_start_day asc,booking_start asc")
	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *SeatRepo) ListLockingBy(ctx context.Context, req *biz.SearchSeatLockingRequest) ([]*model.SeatLock, error) {
	var modelSeat *model.Seat
	modelSeatLock := &model.SeatLock{}
	tableName := modelSeatLock.SrSeatLock.TableName()
	mgr := gen.SrSeatLockMgr(r.data.db)

	mgr.Where(fmt.Sprintf("`%v`.%v IS NULL", tableName, gen.SrSeatLockColumns.DeletedAt))
	if req.Id != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatLockColumns.ID), req.Id)
	}
	if req.SeatId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ? or `%v`.%v = ?", tableName, gen.SrSeatLockColumns.SeatID, tableName, gen.SrSeatLockColumns.MemberID), req.SeatId, req.OrMemberId)
	}
	if req.MemberId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatLockColumns.MemberID), req.MemberId)
	}
	if req.ShopId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatLockColumns.ShopID), req.ShopId)
	}
	if req.Status != nil {

		mgr.Where(fmt.Sprintf("`%v`.%v in ?", tableName, gen.SrSeatLockColumns.Status), req.Status)
	}
	if req.Join {
		mgr.Scopes(modelSeatLock.ScopesJoin)
	}
	mgr.Scopes(
		modelSeat.ScopesDateTime("<=", "lock_end_day", req.GteLockDateEnd),
		modelSeat.ScopesDateTime("<=", "lock_end", req.GteLockTimeEnd),
		modelSeat.ScopesDateTime(">=", "lock_end_day", req.LteLockDateEnd),
		modelSeat.ScopesDateTime(">=", "lock_end", req.LteLockTimeEnd),
		modelSeat.ScopesDateTime("<=", "lock_start_day", req.GteLockDateStart),
		modelSeat.ScopesDateTime("<=", "lock_start", req.GteLockTimeStart),
		modelSeat.ScopesDateTime(">=", "lock_start_day", req.LteLockDateStart),
		modelSeat.ScopesDateTime(">=", "lock_start", req.LteLockTimeStart),
	).Order("lock_start_day asc,lock_start asc")
	var modelList []*model.SeatLock
	err := mgr.Scan(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}

func (r *SeatRepo) FindLockingBy(ctx context.Context, req *biz.SearchSeatLockingRequest) (*model.SeatLock, error) {
	var modelSeat *model.Seat
	modelSeatLock := &model.SeatLock{}
	tableName := modelSeatLock.SrSeatLock.TableName()
	mgr := gen.SrSeatLockMgr(r.data.db)

	mgr.Where(fmt.Sprintf("`%v`.%v IS NULL", tableName, gen.SrSeatLockColumns.DeletedAt))
	if req.Id != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatLockColumns.ID), req.Id)
	}
	if req.SeatId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ? or `%v`.%v = ?", tableName, gen.SrSeatLockColumns.SeatID, tableName, gen.SrSeatLockColumns.MemberID), req.SeatId, req.OrMemberId)
	}
	if req.MemberId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatLockColumns.MemberID), req.MemberId)
	}
	if req.ShopId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrSeatLockColumns.ShopID), req.ShopId)
	}
	if req.Status != nil {

		mgr.Where(fmt.Sprintf("`%v`.%v in ?", tableName, gen.SrSeatLockColumns.Status), req.Status)
	}
	if req.Join {
		mgr.Scopes(modelSeatLock.ScopesJoin)
	}
	mgr.Scopes(
		modelSeat.ScopesDateTime("<=", "lock_end_day", req.GteLockDateEnd),
		modelSeat.ScopesDateTime("<=", "lock_end", req.GteLockTimeEnd),
		modelSeat.ScopesDateTime(">=", "lock_end_day", req.LteLockDateEnd),
		modelSeat.ScopesDateTime(">=", "lock_end", req.LteLockTimeEnd),
		modelSeat.ScopesDateTime("<=", "lock_start_day", req.GteLockDateStart),
		modelSeat.ScopesDateTime("<=", "lock_start", req.GteLockTimeStart),
		modelSeat.ScopesDateTime(">=", "lock_start_day", req.LteLockDateStart),
		modelSeat.ScopesDateTime(">=", "lock_start", req.LteLockTimeStart),
	).Order("lock_start_day asc,lock_start asc")
	var data *model.SeatLock
	err := mgr.Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *SeatRepo) SeatTypeListAll(context.Context) ([]*gen.SrSeatType, error) {
	mgr := gen.SrSeatTypeMgr(r.data.db)
	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}
	return data, nil
}
