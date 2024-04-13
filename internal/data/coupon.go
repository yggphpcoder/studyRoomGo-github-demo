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

type CouponRepo struct {
	data *Data
	log  *log.Helper
}

// NewCouponRepo .
func NewCouponRepo(data *Data, logger log.Logger) biz.CouponRepo {
	return &CouponRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CouponRepo) FindByID(ctx context.Context, id int64) (data *model.Coupons, err error) {
	mgr := gen.SrCouponsMgr(r.data.db).Model(&data).Joins("Applicable").Joins("Gift")
	db := mgr.First(&data, id)
	if db.Error != nil {
		return nil, db.Error
	}
	return data, nil
}

func (r *CouponRepo) FindByCode(ctx context.Context, code string) (data *model.Coupons, err error) {
	mgr := gen.SrCouponsMgr(r.data.db).Model(&data).Joins("Applicable").Joins("Gift")
	db := mgr.Where("code = ?", code).Take(&data)
	if db.Error != nil {
		return nil, db.Error
	}
	return data, nil
}

func (r *CouponRepo) ListBy(ctx context.Context, req *biz.CouponCenterRequest) ([]*model.Coupons, error) {
	mgr := gen.SrCouponsMgr(r.data.db).Model(&model.Coupons{}).Joins("Applicable").Joins("Gift")
	if req.MerchantId != 0 {
		mgr.Where("merchant_id=?", req.MerchantId)
	}
	if req.ShopId != 0 {
		mgr.Where("shop_id=? or shop_id=0", req.ShopId)
	}
	if req.IsShow != nil {
		mgr.Where("is_show = ?", &req.IsShow)
	}
	//if req.Sort > 0 {
	//	mgr.Order(SortType(req.Sort, tableName))
	//} else {
	//	mgr.Order(SortType(sortStatusASort, tableName))
	//}
	if req.PageSize > 0 && req.Page > 0 {
		mgr.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	mgr.Where("status=1")
	mgr.Where("end_date > ?", time.Now())
	mgr.Where("deleted_at IS NULL")
	var modelList []*model.Coupons
	err := mgr.Scan(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}

func (r *CouponRepo) FindReceiveByCode(ctx context.Context, code string) (data *model.CouponsReceive, err error) {
	mgr := gen.SrCouponsReceiveMgr(r.data.db).Model(&data).Joins("Applicable").Joins("Gift")
	db := mgr.Where("code = ?", code).Take(&data)
	if db.Error != nil {
		return nil, db.Error
	}
	return data, nil
}

func (r *CouponRepo) FindReceiveByOrderId(ctx context.Context, id int64) (data *model.CouponsReceive, err error) {
	mgr := gen.SrCouponsReceiveMgr(r.data.db).Model(&data).Joins("Applicable").Joins("Gift")
	db := mgr.Where("order_id = ?", id).Take(&data)
	if db.Error != nil {
		return nil, db.Error
	}
	return data, nil
}
func (r *CouponRepo) FindReceiveById(ctx context.Context, id int64) (data *model.CouponsReceive, err error) {
	mgr := gen.SrCouponsReceiveMgr(r.data.db).Model(&data).Joins("Applicable").Joins("Gift")
	db := mgr.First(&data, id)
	if db.Error != nil {
		return nil, db.Error
	}
	return data, nil
}

func (r *CouponRepo) ListReceiveBy(ctx context.Context, req *biz.SearchCouponReceiveRequest) ([]*model.CouponsReceive, error) {
	mgr := gen.SrCouponsReceiveMgr(r.data.db).Model(&model.CouponsReceive{}).Joins("Applicable").Joins("Gift")
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	tableName := gen.SrCouponsReceiveMgr(r.data.db).GetTableName()
	if req.MemberId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrCouponsReceiveColumns.MemberID), req.MemberId)

	}
	if req.CouponId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrCouponsReceiveColumns.CouponID), req.CouponId)

	}
	if req.NotUsed != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrCouponsReceiveColumns.IsUsed), 0)
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrCouponsReceiveColumns.Status), 1) //未过期
	}
	//if req.Sort > 0 {
	//	mgr.Order(SortType(req.Sort, tableName))
	//} else {
	mgr.Order("is_used asc,id desc")
	//}
	if req.PageSize > 0 && req.Page > 0 {
		mgr.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}

	var modelList []*model.CouponsReceive
	err := mgr.Scan(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}

func (r *CouponRepo) SaveCouponReceive(ctx context.Context, couponReceive *model.CouponsReceive) (bool, error) {
	mgr := gen.SrCouponsReceiveMgr(r.data.db)
	mgr.Model(&model.CouponsReceive{}).Save(couponReceive)
	if mgr.Error != nil {
		return false, mgr.Error
	}

	return true, nil
}

func (r *CouponRepo) UpdateCouponReceive(ctx context.Context, couponReceive *model.CouponsReceive) (bool, error) {
	mgr := gen.SrCouponsReceiveMgr(r.data.db)
	err := mgr.Select("*").Omit(Omit...).Updates(&couponReceive.SrCouponsReceive).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *CouponRepo) UpdateCouponReceiveExpired(ctx context.Context, ids []int64) (bool, error) {
	mgr := gen.SrCouponsReceiveMgr(r.data.db)
	tableName := gen.SrCouponsReceiveMgr(r.data.db).GetTableName()
	mgr.Where(fmt.Sprintf("`%v`.%v in ?", tableName, gen.SrCouponsReceiveColumns.ID), ids) //过期
	err := mgr.Update(gen.SrCouponsReceiveColumns.Status, 2).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *CouponRepo) RecordUseLog(ctx context.Context, log *biz.CouponUseLog) error {
	mgr := gen.SrCouponsUseLogMgr(r.data.db)
	mgr.Save(log.SrCouponsUseLog)

	if mgr.Error != nil {
		return mgr.Error
	}
	return nil
}
