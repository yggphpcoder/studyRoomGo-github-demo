package biz

import (
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"math/rand"
	"strconv"
	"strings"
	"studyRoomGo/internal/data/gen"
	"time"
)

const (
	ProfitSharingOrderStatusPending    = 1  // 待处理
	ProfitSharingOrderStatusProcessing = 2  // 处理中
	ProfitSharingOrderStatusFinished   = 3  // 分账完成
	ProfitSharingOrderStatusFail       = 10 // 分账失败

)

var ProfitSharingToCode = map[profitsharing.OrderStatus]int8{
	profitsharing.ORDERSTATUS_PROCESSING: ProfitSharingOrderStatusProcessing,
	profitsharing.ORDERSTATUS_FINISHED:   ProfitSharingOrderStatusFinished,
}

type ProfitSharingWeChatOrder struct {
	*gen.MeProfitSharingWechatOrder
}

type ProfitSharingOrderRequest struct {
	ShopId          int64
	MerchantId      int64
	OrderDesc       string
	TransactionId   string
	RelateOrderType int8
	RelateOrderNo   string
	Status          int8
	UnfreezeUnsplit int8
	Amount          *ProfitSharingOrderAmount
	Receivers       *ProfitSharingOrderReceivers
}

type ProfitSharingOrderAmount struct {
	TotalAmount float64
	Ratio       int8
	Amount      float64
}

type ProfitSharingOrderReceivers struct {
	ReceiversType    string
	ReceiversAccount string
	ReceiversName    string
}

func (r *ProfitSharingWeChatOrder) GenerateProfitSharingOrder(req *ProfitSharingOrderRequest) {
	createBy := uint(0)
	r.MeProfitSharingWechatOrder = &gen.MeProfitSharingWechatOrder{
		OrderNo:          GenerateProfitSharingOrderNo(req.MerchantId, req.ShopId),
		TransactionID:    req.TransactionId,
		OrderDesc:        req.OrderDesc,
		ReceiversType:    req.Receivers.ReceiversType,
		ReceiversAccount: req.Receivers.ReceiversAccount,
		ReceiversName:    req.Receivers.ReceiversName,
		MerchantID:       req.MerchantId,
		ShopID:           req.ShopId,
		RelateOrderType:  req.RelateOrderType,
		RelateOrderNo:    req.RelateOrderNo,
		Status:           req.Status,
		UnfreezeUnsplit:  req.UnfreezeUnsplit,
		TotalAmount:      int(req.Amount.TotalAmount * 100),
		Ratio:            req.Amount.Ratio,
		Amount:           int(req.Amount.Amount * 100),
		CreateBy:         &createBy,
		UpdateBy:         &createBy,
	}
}

func GenerateProfitSharingOrderNo(merchantId int64, shopId int64) string {

	shopStr := strconv.Itoa(int(shopId)) //后四位取模
	shopStr = fmt.Sprintf("%04s", shopStr)

	merchantStr := strconv.Itoa(int(merchantId)) //后2位取模
	merchantStr = fmt.Sprintf("%02s", merchantStr)

	timeUnix := strconv.Itoa(int(time.Now().Unix()))                                             //单位秒
	randStr := fmt.Sprintf("%03v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000)) //这里面前面的03v是和后面的1000相对应的
	var build strings.Builder
	build.Grow(20)
	build.WriteString("P")
	build.WriteString(merchantStr)
	build.WriteString(shopStr)
	build.WriteString(timeUnix)
	build.WriteString(randStr)
	tag := build.String()
	return tag
}

type ProfitSharingWechatOrderLog struct {
	*gen.MeProfitSharingWechatOrderLog
}
