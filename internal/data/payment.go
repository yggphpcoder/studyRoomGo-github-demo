package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
)

type PaymentRepo struct {
	data *Data
	log  *log.Helper
}

// NewPaymentRepo .
func NewPaymentRepo(data *Data, logger log.Logger) biz.PaymentRepo {
	return &PaymentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *PaymentRepo) FindByID(ctx context.Context, id int64) (*gen.SrPayment, error) {
	mgr := gen.SrPaymentMgr(r.data.db)
	data, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *PaymentRepo) FindByOrderNo(ctx context.Context, orderNo string) (*gen.SrPayment, error) {
	mgr := gen.SrPaymentMgr(r.data.db)
	data, err := mgr.FetchUniqueBySrPaymentOrderNoUIndex(orderNo)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *PaymentRepo) FindByTransactionId(ctx context.Context, TransactionId string) (*gen.SrPayment, error) {
	mgr := gen.SrPaymentMgr(r.data.db)
	data, err := mgr.FetchUniqueBySrPaymentTransactionIDUIndex(&TransactionId)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (r *PaymentRepo) SavePayment(ctx context.Context, payment *biz.Payment) (bool, error) {
	mgr := gen.SrPaymentMgr(r.data.db)
	mgr.Save(payment.SrPayment)
	note := ""
	if mgr.Error != nil {
		note = mgr.Error.Error()
	}
	_, _ = r.RecordLog(ctx, &biz.PaymentLog{
		SrPaymentLog: &gen.SrPaymentLog{
			PaymentID:     payment.ID,
			TransactionID: payment.TransactionID,
			OrderNo:       payment.OrderNo,
			Status:        payment.Status,
			Note:          note,
		},
	})
	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *PaymentRepo) UpdatePayment(ctx context.Context, payment *biz.Payment) (bool, error) {
	mgr := gen.SrPaymentMgr(r.data.db)
	err := mgr.Select("*").Omit(Omit...).Updates(&payment.SrPayment).Error
	note := ""
	if err != nil {
		note = err.Error()
	}
	_, _ = r.RecordLog(ctx, &biz.PaymentLog{
		SrPaymentLog: &gen.SrPaymentLog{
			PaymentID:     payment.ID,
			TransactionID: payment.TransactionID,
			OrderNo:       payment.OrderNo,
			Status:        payment.Status,
			Note:          note,
		},
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *PaymentRepo) SaveWechatPayment(ctx context.Context, payment *biz.Payment) (bool, error) {
	mgr := gen.SrPaymentWechatMgr(r.data.db)
	mgr.Save(payment.WxPayment)
	note := "wechat_payment|"
	if mgr.Error != nil {
		note = mgr.Error.Error()
	}
	_, _ = r.RecordLog(ctx, &biz.PaymentLog{
		SrPaymentLog: &gen.SrPaymentLog{
			PaymentID:     payment.ID,
			TransactionID: payment.TransactionID,
			OrderNo:       payment.OrderNo,
			Status:        payment.Status,
			Note:          note,
		},
	})
	if mgr.Error != nil {
		return false, mgr.Error
	}

	return true, nil
}

func (r *PaymentRepo) UpdateWechatPayment(ctx context.Context, payment *biz.Payment) (bool, error) {
	mgr := gen.SrPaymentWechatMgr(r.data.db)
	err := mgr.Select("*").Omit(Omit...).Updates(&payment.WxPayment).Error
	note := ""
	if err != nil {
		note = err.Error()
	}
	_, _ = r.RecordLog(ctx, &biz.PaymentLog{
		SrPaymentLog: &gen.SrPaymentLog{
			PaymentID:     payment.ID,
			TransactionID: payment.TransactionID,
			OrderNo:       payment.OrderNo,
			Status:        payment.Status,
			Note:          note,
		},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *PaymentRepo) RecordLog(ctx context.Context, log *biz.PaymentLog) (bool, error) {
	mgr := gen.SrPaymentLogMgr(r.data.db)
	mgr.Save(log.SrPaymentLog)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}
