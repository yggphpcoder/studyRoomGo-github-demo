package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"strconv"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
)

type MemberPackageCardRepo struct {
	data *Data
	log  *log.Helper
}

// NewMemberPackageCardRepo .
func NewMemberPackageCardRepo(data *Data, logger log.Logger) biz.MemberPackageCardRepo {
	return &MemberPackageCardRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *MemberPackageCardRepo) Save(ctx context.Context, card *biz.MemberPackageCard) (bool, error) {
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	if merchantId != 0 {
		card.SrMemberPackageCard.MerchantID = int64(merchantId)
	}
	if shopId != 0 {
		card.SrMemberPackageCard.ShopID = int64(shopId)
	}
	mgr := gen.SrMemberPackageCardMgr(r.data.db)
	mgr.Save(card.SrMemberPackageCard)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *MemberPackageCardRepo) Update(ctx context.Context, card *biz.MemberPackageCard) (bool, error) {
	mgr := gen.SrMemberPackageCardMgr(r.data.db)
	if err := mgr.Select("*").Omit(Omit...).Updates(&card.SrMemberPackageCard).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *MemberPackageCardRepo) FindByID(ctx context.Context, id int64) (*gen.SrMemberPackageCard, error) {
	mgr := gen.SrMemberPackageCardMgr(r.data.db)
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	mgr.WithDeletedAt(nil)
	d, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *MemberPackageCardRepo) ListBy(ctx context.Context, req *biz.SearchMemberPackageCardRequest) ([]*gen.SrMemberPackageCard, error) {
	mgr := gen.SrMemberPackageCardMgr(r.data.db)
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	tableName := mgr.GetTableName()
	if req.MemberId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrMemberPackageCardColumns.MemberID), req.MemberId)

	}

	if req.PackageId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.SrMemberPackageCardColumns.PackageID), req.PackageId)

	}
	if len(req.Status) != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v in ?", tableName, gen.SrMemberPackageCardColumns.Status), req.Status)
	}
	if req.Sort > 0 {
		mgr.Order(SortType(req.Sort, tableName))
	} else {
		mgr.Order(SortType(sortStatusASort, tableName))
	}
	if req.PageSize > 0 && req.Page > 0 {
		mgr.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}

	var modelList []*gen.SrMemberPackageCard
	err := mgr.Scan(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}

func (r *MemberPackageCardRepo) RecordUseLog(ctx context.Context, log *biz.MemberPackageCardUseLog) {
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	mgr := gen.SrMemberPackageCardUseLogMgr(r.data.db)
	createBy := uint(0)
	mgr.Save(&gen.SrMemberPackageCardUseLog{
		MerchantID: int64(merchantId),
		MpID:       log.MpID,
		MemberID:   log.MemberID,
		Value:      log.Value,
		Extra:      extra(log.Extra),
		Remark:     &log.Remark,
		CreateBy:   &createBy,
	})
}

func (r *MemberPackageCardRepo) LastUseLog(ctx context.Context, memberId int64) (*biz.MemberPackageCardUseLog, error) {
	mgr := gen.SrMemberPackageCardUseLogMgr(r.data.db)

	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	mgr.Order(SortType(sortCreatAtDSort))
	d, err := mgr.GetByOption(mgr.WithMemberID(memberId))
	if err != nil {
		return nil, err
	}
	return &biz.MemberPackageCardUseLog{
		MpID:     d.MpID,
		MemberID: d.MemberID,
		Value:    d.Value,
		Remark:   *d.Remark,
		Extra:    nil,
		CreatAt:  d.CreatedAt,
	}, nil
}

func (r *MemberPackageCardRepo) RecordOperateLog(ctx context.Context, log *biz.MemberPackageCardOperateLog) {
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	mgr := gen.SrMemberPackageCardOperateLogMgr(r.data.db)
	createBy := uint(0)
	mgr.Save(&gen.SrMemberPackageCardOperateLog{
		MerchantID:    int64(merchantId),
		MpID:          log.MpId,
		MemberID:      log.MemberId,
		PackageCardID: log.PackageId,
		Operate:       log.Operate,
		Extra:         extra(log.Extra),
		CreateBy:      &createBy,
	})
}
