package biz

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"studyRoomGo/internal/data/gen"
	"time"
)

const (
	OrderTypeSeatBook       = 1
	OrderTypeBuyPackageCard = 2
	OrderTypeBuyCourse      = 3
)

const (
	OrderStatusPending = 1 // 待支付
	OrderStatusPaying  = 2 // 支付中
	OrderStatusPayment = 3 // 已支付
	OrderStatusCanncel = 4 // 已取消
	OrderStatusClosed  = 5 // 已关闭
	OrderStatusRefund  = 6 // 已关闭
)

var WxPaymentToOrderCode = map[string]int{
	WxPaymentStatusSuccess:    OrderStatusPayment,
	WxPaymentStatusNotpay:     OrderStatusPending,
	WxPaymentStatusUserpaying: OrderStatusPaying,
	WxPaymentStatusPayerror:   OrderStatusPending,
	WxPaymentStatusClosed:     OrderStatusClosed,
	WxPaymentStatusRefund:     OrderStatusRefund,
}

type Order struct {
	*gen.SrOrder
	Total []*gen.SrOrderTotal
}

type OrderRequest struct {
	MemberId          int64
	MerchantId        int64
	OrderDesc         string
	OrderType         int8 // seatBook,buyPackageCard
	OrderTypeRelateID string
	Price             float64
	Qty               int32
	Status            int8
}

func (r *Order) GenerateOrder(req *OrderRequest) {
	createBy := uint(0)
	r.SrOrder = &gen.SrOrder{
		OrderNo:           GenerateOrderNo(req.MerchantId, req.MemberId),
		MemberID:          req.MemberId,
		MerchantID:        req.MerchantId,
		OrderDesc:         req.OrderDesc,
		OrderType:         req.OrderType,
		OrderTypeRelateID: req.OrderTypeRelateID,
		Status:            req.Status,
		Price:             req.Price,
		Quantity:          int(req.Qty),
		CreateBy:          &createBy,
		UpdateBy:          &createBy,
	}
}

func GenerateOrderNo(merchantId int64, customerId int64) string {

	customerStr := strconv.Itoa(int(customerId)) //后四位取模
	customerStr = fmt.Sprintf("%04s", customerStr)

	shopStr := strconv.Itoa(int(merchantId)) //后2位取模
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

type OrderLog struct {
	*gen.SrOrderLog
}

type CreateBuyCardOrderRequest struct {
	MerchantId int64
	MemberId   int64
	CardId     int64
	CouponCode string
	Qty        int32
}

type CreateOrderReply struct {
	CardInfo       *gen.SrPackageCard
	OrderInfo      *gen.SrOrder
	CourseSaleInfo *gen.EdCoursesSale
}

type CreateProfitSharingOrderRequest struct {
	MerchantId      int64
	ShopId          int64
	RelateOrderType int8
	RelateOrderNo   string
	TransactionId   string
	TotalAmount     float64
}

type CreateBuyCourseOrderRequest struct {
	MerchantId int64
	MemberId   int64
	CourseId   int64
	SaleId     int64
	StudentId  int64
	Qty        int32
}
