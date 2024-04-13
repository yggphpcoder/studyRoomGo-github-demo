package biz

import "studyRoomGo/internal/data/gen"

const (
	PaymentStatusPending  = 1
	PaymentStatusPaying   = 2
	PaymentStatusPayed    = 3
	PaymentStatusChannel  = 4
	PaymentStatusClosed   = 5
	PaymentStatusRefund   = 6
	PaymentStatusPayError = 7
)

const (
	WxPaymentStatusSuccess    = "SUCCESS"    //支付成功
	WxPaymentStatusRefund     = "REFUND"     //转入退款
	WxPaymentStatusNotpay     = "NOTPAY"     //未支付
	WxPaymentStatusClosed     = "CLOSED"     //已关闭
	WxPaymentStatusRevoked    = "REVOKED"    //已撤销（付款码支付）
	WxPaymentStatusUserpaying = "USERPAYING" //用户支付中（付款码支付）
	WxPaymentStatusPayerror   = "PAYERROR"   //支付失败(其他原因，如银行返回失败)
)

var WxPaymentToCode = map[string]int{
	WxPaymentStatusSuccess:  PaymentStatusPayed,
	WxPaymentStatusNotpay:   PaymentStatusPending,
	WxPaymentStatusPayerror: PaymentStatusPayError,
	WxPaymentStatusClosed:   PaymentStatusClosed,
	WxPaymentStatusRefund:   PaymentStatusRefund,
}

type Payment struct {
	*gen.SrPayment
	WxPayment *gen.SrPaymentWechat
}

type PaymentLog struct {
	*gen.SrPaymentLog
}
