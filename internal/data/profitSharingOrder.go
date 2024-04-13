package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
)

type ProfitSharingOrderRepo struct {
	data *Data
	log  *log.Helper
}

// NewProfitSharingOrderRepo .
func NewProfitSharingOrderRepo(data *Data, logger log.Logger) biz.ProfitSharingOrderRepo {
	return &ProfitSharingOrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ProfitSharingOrderRepo) FindByID(ctx context.Context, id int64) (*gen.MeProfitSharingWechatOrder, error) {
	mgr := gen.MeProfitSharingWechatOrderMgr(r.data.db)
	data, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ProfitSharingOrderRepo) FindByOrderNo(ctx context.Context, orderNo string) (*gen.MeProfitSharingWechatOrder, error) {
	data := &gen.MeProfitSharingWechatOrder{}
	mgr := gen.MeProfitSharingWechatOrderMgr(r.data.db)
	mgr.Where("relate_order_no = ?", orderNo).First(&data)
	if mgr.Error != nil {
		return nil, mgr.Error
	}
	return data, nil
}

func (r *ProfitSharingOrderRepo) FindByTransactionId(ctx context.Context, transactionId string) (*gen.MeProfitSharingWechatOrder, error) {
	data := &gen.MeProfitSharingWechatOrder{}
	mgr := gen.MeProfitSharingWechatOrderMgr(r.data.db)
	mgr.Where("transaction_id = ?", transactionId).First(&data)
	if mgr.Error != nil {
		return nil, mgr.Error
	}
	return data, nil
}
func (r *ProfitSharingOrderRepo) FindReceiversBy(ctx context.Context, shopId int64, merchantId int64, relateOrderType int8) (*gen.MeProfitSharingWechatReceivers, error) {
	data := &gen.MeProfitSharingWechatReceivers{}
	mgr := gen.MeProfitSharingWechatReceiversMgr(r.data.db)
	mgr.Where("shop_id = ?", shopId)
	mgr.Where("merchant_id = ?", merchantId)
	mgr.Where("relate_order_type = ?", relateOrderType).First(&data)
	if mgr.Error != nil {
		return nil, mgr.Error
	}
	return data, nil
}

func (r *ProfitSharingOrderRepo) Save(ctx context.Context, order *gen.MeProfitSharingWechatOrder) (bool, error) {
	mgr := gen.MeProfitSharingWechatOrderMgr(r.data.db)
	mgr.Save(order)
	if mgr.Error != nil {
		return false, mgr.Error
	}

	return true, nil
}

func (r *ProfitSharingOrderRepo) Update(ctx context.Context, order *gen.MeProfitSharingWechatOrder) (bool, error) {
	mgr := gen.MeProfitSharingWechatOrderMgr(r.data.db)
	err := mgr.Select("*").Omit(Omit...).Updates(&order).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ProfitSharingOrderRepo) RecordLog(ctx context.Context, log *biz.ProfitSharingWechatOrderLog) error {
	mgr := gen.MeProfitSharingWechatOrderLogMgr(r.data.db)
	mgr.Save(log.MeProfitSharingWechatOrderLog)

	if mgr.Error != nil {
		return mgr.Error
	}
	return nil
}
