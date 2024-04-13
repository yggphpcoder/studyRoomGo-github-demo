package biz

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/data/gen"
)

// PaymentUsecase is a Payment usecase.
type PaymentUsecase struct {
	repo      PaymentRepo
	orderRepo OrderRepo
	log       *log.Helper
}

// NewPaymentUsecase new a Payment usecase.
func NewPaymentUsecase(repo PaymentRepo, orderRepo OrderRepo, logger log.Logger) *PaymentUsecase {
	return &PaymentUsecase{
		repo:      repo,
		orderRepo: orderRepo,
		log:       log.NewHelper(logger),
	}
}

type WxCheckoutReply struct {
	// 预支付交易会话标识
	PrepayId string
	// 应用ID
	Appid string
	// 时间戳
	TimeStamp string
	// 随机字符串
	NonceStr string
	// 订单详情扩展字符串
	Package string
	// 签名方式
	SignType string
	// 签名
	PaySign string
}

type WxCheckout struct {
	MemberId int64
	OrderNo  string
}

func (uc *PaymentUsecase) WxPaymentCheckout(ctx context.Context, checkout *WxCheckout) (*WxCheckoutReply, error) {
	order, err := uc.orderRepo.FindByOrderNo(ctx, checkout.OrderNo)
	if err != nil {
		return nil, err
	}
	if order.MemberID != checkout.MemberId {
		return nil, errors.New("not your order")
	}
	if order.Status != OrderStatusPending {
		return nil, errors.New("order status not pending")
	}
	payment, err := uc.repo.FindByOrderNo(ctx, checkout.OrderNo)
	if err != nil {
		if !gen.IsNotFound(err) {
			return nil, err
		}
	}

	if payment == nil {
		createBy := uint(0)
		if ok, err := uc.repo.SavePayment(ctx, &Payment{
			SrPayment: &gen.SrPayment{
				OrderNo:     checkout.OrderNo,
				PaymentType: "wxpay",
				Currency:    "CNY",
				Amount:      order.Price,
				Status:      PaymentStatusPending,
				CreateBy:    &createBy,
				UpdateBy:    &createBy,
			},
		}); !ok {
			return nil, err
		}
	} else {
		updateBy := uint(0)
		payment.Status = PaymentStatusPending
		payment.UpdateBy = &updateBy
		if ok, err := uc.repo.UpdatePayment(ctx, &Payment{SrPayment: payment}); !ok {
			return nil, err
		}
	}
	order.Status = OrderStatusPaying

	if ok, err := uc.orderRepo.Update(ctx, &Order{SrOrder: order}); !ok {
		return nil, err
	}

	attachByte, err := json.Marshal(wechat.Attach{
		MerchantId: order.MerchantID,
		ShopId:     order.ShopID,
	})
	if err != nil {
		return nil, err
	}
	wxJSApi := wechat.JSApi{
		MerchantId:  order.MerchantID,
		ShopId:      order.ShopID,
		MemberId:    order.MemberID,
		Description: order.OrderDesc,
		OutTradeNo:  order.OrderNo,
		Attach:      string(attachByte),
		TotalPrice:  int64(order.Price * 100),
	}
	if order.MerchantID == 3 {
		rep, err := wxJSApi.PartnerRequest()
		if err != nil {
			return nil, err
		}
		return &WxCheckoutReply{
			PrepayId:  *rep.PrepayId,
			Appid:     *rep.Appid,
			TimeStamp: *rep.TimeStamp,
			NonceStr:  *rep.NonceStr,
			Package:   *rep.Package,
			SignType:  *rep.SignType,
			PaySign:   *rep.PaySign,
		}, err
	}
	rep, err := wxJSApi.Request()
	if err != nil {
		return nil, err
	}
	return &WxCheckoutReply{
		PrepayId:  *rep.PrepayId,
		Appid:     *rep.Appid,
		TimeStamp: *rep.TimeStamp,
		NonceStr:  *rep.NonceStr,
		Package:   *rep.Package,
		SignType:  *rep.SignType,
		PaySign:   *rep.PaySign,
	}, err
}

type WxCheckoutReturn struct {
	MemberId   int64
	OutTradeNo string
	Status     int
}

// WxPaymentCheckoutReturn  creates a Payment, and returns the new Payment.
func (uc *PaymentUsecase) WxPaymentCheckoutReturn(ctx context.Context, ret *WxCheckoutReturn) (bool, error) {
	order, err := uc.orderRepo.FindByOrderNo(ctx, ret.OutTradeNo)
	if err != nil {
		return false, err
	}
	if order.MemberID != ret.MemberId {
		return false, errors.New("not your order")
	}

	paymentInfo, err := uc.repo.FindByOrderNo(ctx, ret.OutTradeNo)
	if err != nil {
		return false, err
	}

	switch ret.Status {
	case 1:
		paymentInfo.Status = PaymentStatusPaying
	case 2:
		paymentInfo.Status = PaymentStatusChannel
		order.Status = OrderStatusCanncel
		if ok, err := uc.orderRepo.Update(ctx, &Order{SrOrder: order}); !ok {
			return false, err
		}
	default:
		return false, errors.New("status error")
	}

	payment := &Payment{
		SrPayment: paymentInfo,
	}
	if ok, err := uc.repo.UpdatePayment(ctx, payment); !ok {
		return false, err
	}
	return true, nil
}

type WxCheckoutNotice struct {
	SpMchid        string
	SpAppid        string
	SubMchid       string
	SubAppid       string
	TransactionId  string
	OutTradeNo     string
	TradeType      string
	TradeState     string
	TradeStateDesc string
	BankType       string
	SuccessTime    string
	Payer          struct {
		OpenId    string
		SpOpenid  string
		SubOpenid string
	}
	Amount struct {
		Total         int32
		PayerTotal    int32
		Currency      string
		PayerCurrency string
	}
}

func (uc *PaymentUsecase) WxPaymentNotice(ctx context.Context, notice *WxCheckoutNotice) (bool, error) {
	order, err := uc.orderRepo.FindByOrderNo(ctx, notice.OutTradeNo)
	if err != nil {
		return false, err
	}
	if order.Status != OrderStatusPaying {

		return false, errors.New("order status not paying")
	}

	paymentInfo, err := uc.repo.FindByOrderNo(ctx, notice.OutTradeNo)
	if err != nil {
		return false, err
	}
	if paymentInfo.Status != PaymentStatusPaying {
		return false, errors.New("payment status not paying")
	}
	order.PayProvider = "wxpay"
	order.Status = int8(WxPaymentToOrderCode[notice.TradeState])
	if ok, err := uc.orderRepo.Update(ctx, &Order{SrOrder: order}); !ok {
		return false, err
	}

	paymentInfo.TransactionID = &notice.TransactionId
	paymentInfo.Status = int8(WxPaymentToCode[notice.TradeState])
	payment := &Payment{
		SrPayment: paymentInfo,
		WxPayment: &gen.SrPaymentWechat{
			OutTradeNo:          notice.OutTradeNo,
			TransactionID:       notice.TransactionId,
			TradeType:           notice.TradeType,
			TradeState:          notice.TradeState,
			TradeStateDesc:      notice.TradeStateDesc,
			BankType:            notice.BankType,
			SuccessTime:         notice.SuccessTime,
			PayerOpenid:         notice.Payer.OpenId,
			AmountTotal:         int(notice.Amount.Total),
			AmountPayerTotal:    int(notice.Amount.PayerTotal),
			AmountCurrency:      notice.Amount.Currency,
			AmountPayerCurrency: notice.Amount.PayerCurrency,
			Attach:              "",
		},
	}
	if payment.WxPayment.PayerOpenid == "" {
		payment.WxPayment.PayerOpenid = notice.Payer.SubOpenid
	}
	if ok, err := uc.repo.UpdatePayment(ctx, payment); !ok {
		return false, err
	}
	if ok, err := uc.repo.SaveWechatPayment(ctx, payment); !ok {
		return false, err
	}
	return true, nil
}

type PaymentCallback struct {
	OrderNo       string
	TransactionId string
	PaymentStatus int
	ActiveAmount  float64
}

func (uc *PaymentUsecase) OrderInfo(ctx context.Context, orderNo string) (*Order, error) {
	data, err := uc.orderRepo.FindByOrderNo(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	return &Order{
		SrOrder: data,
	}, nil
}
