package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"strconv"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
)

type OrderRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &OrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *OrderRepo) FindByID(ctx context.Context, id int64) (*gen.SrOrder, error) {
	mgr := gen.SrOrderMgr(r.data.db)
	data, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *OrderRepo) FindByOrderNo(ctx context.Context, orderNo string) (*gen.SrOrder, error) {
	mgr := gen.SrOrderMgr(r.data.db)
	data, err := mgr.FetchUniqueBySrOrderOrderNoUIndex(orderNo)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *OrderRepo) Save(ctx context.Context, order *biz.Order) (bool, error) {
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	order.SrOrder.MerchantID = int64(merchantId)
	order.SrOrder.ShopID = int64(shopId)
	modelOrder := &model.Order{
		SrOrder: order.SrOrder,
		Total:   order.Total,
	}
	mgr := gen.SrOrderMgr(r.data.db).Model(&modelOrder)
	mgr.Save(modelOrder)
	note := ""
	if mgr.Error != nil {
		note = mgr.Error.Error()
	}
	_, _ = r.RecordLog(ctx, &biz.OrderLog{
		SrOrderLog: &gen.SrOrderLog{
			OrderID: order.ID,
			OrderNo: order.OrderNo,
			Status:  order.Status,
			Note:    note,
		},
	})
	if mgr.Error != nil {
		return false, mgr.Error
	}

	return true, nil
}

func (r *OrderRepo) Update(ctx context.Context, order *biz.Order) (bool, error) {
	mgr := gen.SrOrderMgr(r.data.db)
	err := mgr.Select("*").Omit(Omit...).Updates(&order.SrOrder).Error
	note := ""
	if err != nil {
		note = err.Error()
	}

	_, _ = r.RecordLog(ctx, &biz.OrderLog{
		SrOrderLog: &gen.SrOrderLog{
			OrderID: order.ID,
			OrderNo: order.OrderNo,
			Status:  order.Status,
			Note:    note,
		},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *OrderRepo) RecordLog(ctx context.Context, log *biz.OrderLog) (bool, error) {
	mgr := gen.SrOrderLogMgr(r.data.db)
	mgr.Save(log.SrOrderLog)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}
