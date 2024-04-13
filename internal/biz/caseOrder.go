package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	_func "studyRoomGo/common/func"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"gorm.io/gorm"
)

// OrderUsecase is an Order usecase.
type OrderUsecase struct {
	repo                   OrderRepo
	packageCardRepo        PackageCardRepo
	profitSharingOrderRepo ProfitSharingOrderRepo
	couponRepo             CouponRepo
	courseRepo             CourseRepo

	log *log.Helper
}

// NewOrderUsecase new an Order usecase.
func NewOrderUsecase(repo OrderRepo, packageCardRepo PackageCardRepo, couponRepo CouponRepo, profitSharingOrderRepo ProfitSharingOrderRepo, courseRepo CourseRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{
		repo:                   repo,
		packageCardRepo:        packageCardRepo,
		couponRepo:             couponRepo,
		profitSharingOrderRepo: profitSharingOrderRepo,
		courseRepo:             courseRepo,
		log:                    log.NewHelper(logger),
	}
}
func (uc *OrderUsecase) CalOrderTotalPrice(ctx context.Context, req *CreateBuyCardOrderRequest) (totalPrice float64) {
	cardInfo, err := uc.packageCardRepo.FindByID(ctx, req.CardId)
	if err != nil {
		log.Errorf("Service FindByID error:%s \r\n", err)
		return 0.00
	}
	totalPrice = cardInfo.Price * float64(req.Qty) //单价*数量
	if req.CouponCode != "" {
		couponInfo, err := uc.couponRepo.FindReceiveByCode(ctx, req.CouponCode)
		if err != nil {
			log.Errorf("Service FindByCode error:%s \r\n", err)
			return 0.00
		}
		discount := CalCouponDiscount(couponInfo, CalCouponDiscountProduct{
			price:      cardInfo.Price,
			qty:        req.Qty,
			totalPrice: totalPrice,
		})
		totalPrice -= discount

		if totalPrice < 0 {
			totalPrice = 0.00
		}
	}
	return totalPrice
}
func (uc *OrderUsecase) CreateBuyCardOrder(ctx context.Context, req *CreateBuyCardOrderRequest) (rep *CreateOrderReply, error error) {
	cardInfo, err := uc.packageCardRepo.FindByID(ctx, req.CardId)
	if err != nil {
		log.Errorf("Service FindByID error:%s \r\n", err)
		return nil, err
	}
	var orderTotal []*gen.SrOrderTotal
	if req.CouponCode == "" && _func.InIntSplice([]int{18, 19, 21, 22, 24, 26, 33, 298}, int(req.MemberId)) {
		cardInfo.Price = 0.1
	}

	totalPrice := cardInfo.Price * float64(req.Qty) //单价*数量
	orderTotal = append(orderTotal, &gen.SrOrderTotal{
		Type:        "packageCard",
		Value:       totalPrice,
		RelateType:  1,
		RelateValue: strconv.Itoa(int(cardInfo.ID)),
	})

	if req.CouponCode != "" {
		couponInfo, err := uc.couponRepo.FindReceiveByCode(ctx, req.CouponCode)
		if err != nil {
			log.Errorf("Service FindByCode error:%s \r\n", err)
			return nil, err
		}
		discount := CalCouponDiscount(couponInfo, CalCouponDiscountProduct{
			price:      cardInfo.Price,
			qty:        req.Qty,
			totalPrice: totalPrice,
		})
		totalPrice -= discount
		if totalPrice < 0 {
			totalPrice = 0.00
		}
		orderTotal = append(orderTotal, &gen.SrOrderTotal{
			Type:        "couponDiscount",
			Value:       -discount,
			RelateType:  2,
			RelateValue: strconv.Itoa(int(couponInfo.ID)),
		})
	}
	order := &Order{}
	order.GenerateOrder(&OrderRequest{
		MerchantId:        req.MerchantId,
		MemberId:          req.MemberId,
		OrderDesc:         fmt.Sprintf("%s * %d", cardInfo.Name, req.Qty),
		OrderType:         OrderTypeBuyPackageCard,
		OrderTypeRelateID: strconv.Itoa(int(req.CardId)),
		Price:             totalPrice,
		Qty:               req.Qty,
		Status:            OrderStatusPending,
	})
	orderTotal = append(orderTotal, &gen.SrOrderTotal{
		Type:       "total",
		Value:      totalPrice,
		RelateType: 99,
	})
	order.Total = orderTotal
	if ok, err := uc.repo.Save(ctx, order); !ok {
		log.Errorf("Log Insert error:%s \r\n", err)
		return nil, err
	}

	return &CreateOrderReply{
		CardInfo:  cardInfo,
		OrderInfo: order.SrOrder,
	}, nil
}

func (uc *OrderUsecase) CreateProfitSharingOrder(ctx context.Context, req *CreateProfitSharingOrderRequest) (error error) {
	receiver, err := uc.profitSharingOrderRepo.FindReceiversBy(ctx, req.ShopId, req.MerchantId, req.RelateOrderType) //查找分账接受方信息
	if err != nil {
		log.Errorf("ExecProfitSharingOrder FindReceiversBy not exist shopId:%v,merchantId:%v,RelateOrderType:%v \r\n", req.ShopId, req.MerchantId, req.RelateOrderType)
		return err
	}
	if receiver.Status != 1 { //没有开启,跳过
		return nil
	}
	info, err := uc.profitSharingOrderRepo.FindByTransactionId(ctx, req.TransactionId) //查找是否已经分账
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("ExecProfitSharingOrder FindByTransactionId  err %d \r\n", err)
		return err
	}
	if info != nil && info.ID != 0 { //如果存在,则不能创建
		log.Errorf("ExecProfitSharingOrder FindByTransactionId  exist ID: %d \r\n", info.ID)
		return err
	}
	num := decimal.NewFromFloat(req.TotalAmount * (float64(receiver.Ratio) / 100))
	amount, _ := num.Round(2).Float64()

	order := &ProfitSharingWeChatOrder{}
	order.GenerateProfitSharingOrder(&ProfitSharingOrderRequest{
		ShopId:          req.ShopId,
		MerchantId:      req.MerchantId,
		OrderDesc:       fmt.Sprintf("分账给%s,总额:%.2f,分账率:%d%%,分账金额:%.2f", receiver.ReceiversName, req.TotalAmount, receiver.Ratio, amount),
		TransactionId:   req.TransactionId,
		RelateOrderType: req.RelateOrderType,
		RelateOrderNo:   req.RelateOrderNo,
		Status:          ProfitSharingOrderStatusPending,
		UnfreezeUnsplit: 1,
		Amount: &ProfitSharingOrderAmount{
			TotalAmount: req.TotalAmount,
			Ratio:       receiver.Ratio,
			Amount:      amount,
		},
		Receivers: &ProfitSharingOrderReceivers{
			ReceiversType:    receiver.ReceiversType,
			ReceiversAccount: receiver.ReceiversAccount,
			ReceiversName:    receiver.ReceiversName,
		},
	})

	if ok, err := uc.profitSharingOrderRepo.Save(ctx, order.MeProfitSharingWechatOrder); !ok {
		log.Errorf("GenerateProfitSharingOrder save error:%s ,RelateOrderNo: %s \r\n ", err, order.RelateOrderNo)
		return err
	}
	_ = uc.profitSharingOrderRepo.RecordLog(ctx, &ProfitSharingWechatOrderLog{
		MeProfitSharingWechatOrderLog: &gen.MeProfitSharingWechatOrderLog{
			ProfitSharingID: order.ID,
			OrderNo:         order.OrderNo,
			Status:          order.Status,
		},
	})
	return
}

func (uc *OrderUsecase) CreateBuyCourseOrder(ctx context.Context, req *CreateBuyCourseOrderRequest) (rep *CreateOrderReply, error error) {
	saleInfo, err := uc.courseRepo.FindSaleByID(ctx, req.SaleId)
	if err != nil {
		log.Errorf("Service FindByID error:%s \r\n", err)
		return nil, err
	}
	var orderTotal []*gen.SrOrderTotal

	totalPrice := saleInfo.Price * float64(req.Qty) //单价*数量
	orderTotal = append(orderTotal, &gen.SrOrderTotal{
		Type:        "courseSale",
		Value:       totalPrice,
		RelateType:  1,
		RelateValue: strconv.Itoa(int(saleInfo.ID)),
	})

	order := &Order{}
	order.GenerateOrder(&OrderRequest{
		MerchantId:        req.MerchantId,
		MemberId:          req.MemberId,
		OrderDesc:         fmt.Sprintf("%s * %d", saleInfo.Course.Title, req.Qty),
		OrderType:         OrderTypeBuyCourse,
		OrderTypeRelateID: fmt.Sprintf("%d|%d", req.SaleId, req.StudentId),
		Price:             totalPrice,
		Qty:               req.Qty,
		Status:            OrderStatusPending,
	})
	orderTotal = append(orderTotal, &gen.SrOrderTotal{
		Type:       "total",
		Value:      totalPrice,
		RelateType: 99,
	})
	order.Total = orderTotal
	if ok, err := uc.repo.Save(ctx, order); !ok {
		log.Errorf("Log Insert error:%s \r\n", err)
		return nil, err
	}

	return &CreateOrderReply{
		CourseSaleInfo: &saleInfo.EdCoursesSale,
		OrderInfo:      order.SrOrder,
	}, nil
}

// ExecProfitSharingOrder 执行分账
func (uc *OrderUsecase) ExecProfitSharingOrder(ctx context.Context, profitSharingId int64) (error error) {
	order, err := uc.profitSharingOrderRepo.FindByID(ctx, profitSharingId) //查找分账单
	if err != nil {
		log.Errorf("ExecProfitSharingOrder FindByID  err %d \r\n", err)
		return err
	}

	if order.Status != ProfitSharingOrderStatusPending {
		return errors.New("分账单状态不是待处理")
	}
	profitSharing := wechat.ProfitSharing{
		ShopId:          order.ShopID,
		MerchantId:      order.MerchantID,
		OrderNo:         order.OrderNo,
		OrderDesc:       order.OrderDesc,
		TransactionId:   order.TransactionID,
		UnfreezeUnsplit: order.UnfreezeUnsplit,
		Amount: &wechat.ProfitSharingOrderAmount{
			TotalAmount: order.TotalAmount,
			Ratio:       order.Ratio,
			Amount:      order.Amount,
		},
		Receivers: &wechat.ProfitSharingOrderReceivers{
			ReceiversType:    order.ReceiversType,
			ReceiversAccount: order.ReceiversAccount,
			ReceiversName:    order.ReceiversName,
		},
	}

	response, result, err := profitSharing.ProfitSharingCreateOrder()

	bodyStr := ""
	headerStr := ""
	header, _ := json.Marshal(result.Response.Header)
	headerStr = string(header)
	if err != nil {
		switch err.(type) {
		case *core.APIError:
			value := err.(*core.APIError)
			body, _ := json.Marshal(struct {
				StatusCode int
				Code       string
				Message    string
			}{
				StatusCode: value.StatusCode,
				Code:       value.Code,
				Message:    value.Message,
			})
			bodyStr = string(body)
		default:
			bodyStr = err.Error()
		}
		_ = uc.profitSharingOrderRepo.RecordLog(ctx, &ProfitSharingWechatOrderLog{
			MeProfitSharingWechatOrderLog: &gen.MeProfitSharingWechatOrderLog{
				ProfitSharingID: order.ID,
				OrderNo:         order.OrderNo,
				Status:          ProfitSharingOrderStatusFail,
				Body:            &bodyStr,
				Header:          &headerStr,
			},
		})
		order.Status = ProfitSharingOrderStatusFail
		if ok, err := uc.profitSharingOrderRepo.Update(ctx, order); !ok {
			log.Errorf("GenerateProfitSharingOrder Update error:%s \r\n", err)
			return err
		}
		return err
	}

	body, _ := json.Marshal(response)
	bodyStr = string(body)
	status := ProfitSharingToCode[*response.State]

	order.WechatOrderID = *response.OrderId
	order.Status = status
	if ok, err := uc.profitSharingOrderRepo.Update(ctx, order); !ok {
		log.Errorf("GenerateProfitSharingOrder Update error:%s \r\n", err)
		return err
	}
	_ = uc.profitSharingOrderRepo.RecordLog(ctx, &ProfitSharingWechatOrderLog{
		MeProfitSharingWechatOrderLog: &gen.MeProfitSharingWechatOrderLog{
			ProfitSharingID: order.ID,
			OrderNo:         order.OrderNo,
			Status:          status,
			Body:            &bodyStr,
			Header:          &headerStr,
		},
	})
	return nil
}

func (uc *OrderUsecase) UpdateProfitSharingOrderStatus(ctx context.Context, profitSharingId int64) (error error) {
	order, err := uc.profitSharingOrderRepo.FindByID(ctx, profitSharingId) //查找分账单
	if err != nil {
		log.Errorf("ExecProfitSharingOrder FindByID  err %d \r\n", err)
		return err
	}

	if order.Status != ProfitSharingOrderStatusProcessing {
		return errors.New("分账单状态不是分账中")
	}
	profitSharing := wechat.ProfitSharing{MerchantId: order.MerchantID}
	response, result, err := profitSharing.ProfitSharingGetOrder(order.OrderNo, order.TransactionID)

	bodyStr := ""
	headerStr := ""
	if result != nil {
		header, _ := json.Marshal(result.Response.Header)
		headerStr = string(header)
	}

	if err != nil {
		switch err.(type) {
		case *core.APIError:
			value := err.(*core.APIError)
			body, _ := json.Marshal(struct {
				StatusCode int
				Code       string
				Message    string
			}{
				StatusCode: value.StatusCode,
				Code:       value.Code,
				Message:    value.Message,
			})
			bodyStr = string(body)
		default:
			bodyStr = err.Error()
		}
		_ = uc.profitSharingOrderRepo.RecordLog(ctx, &ProfitSharingWechatOrderLog{
			MeProfitSharingWechatOrderLog: &gen.MeProfitSharingWechatOrderLog{
				ProfitSharingID: order.ID,
				OrderNo:         order.OrderNo,
				Status:          ProfitSharingOrderStatusProcessing,
				Body:            &bodyStr,
				Header:          &headerStr,
			},
		})
		order.Status = ProfitSharingOrderStatusProcessing
		if ok, err := uc.profitSharingOrderRepo.Update(ctx, order); !ok {
			log.Errorf("GenerateProfitSharingOrder Update error:%s \r\n", err)
			return err
		}
		return err
	}

	body, _ := json.Marshal(response)
	bodyStr = string(body)
	status := ProfitSharingToCode[*response.State]

	order.WechatOrderID = *response.OrderId
	order.Status = status
	if ok, err := uc.profitSharingOrderRepo.Update(ctx, order); !ok {
		log.Errorf("GenerateProfitSharingOrder Update error:%s \r\n", err)
		return err
	}
	_ = uc.profitSharingOrderRepo.RecordLog(ctx, &ProfitSharingWechatOrderLog{
		MeProfitSharingWechatOrderLog: &gen.MeProfitSharingWechatOrderLog{
			ProfitSharingID: order.ID,
			OrderNo:         order.OrderNo,
			Status:          status,
			Body:            &bodyStr,
			Header:          &headerStr,
		},
	})
	return nil
}

type CalCouponDiscountProduct struct {
	price      float64
	qty        int32
	totalPrice float64
}

func CalCouponDiscount(c *model.CouponsReceive, p CalCouponDiscountProduct) (discountPrice float64) {
	discountPrice = 0.00
	switch c.DiscountType {
	case 1: //打折
		discountPrice = p.totalPrice * c.DiscountValue / 100
	case 2: // 立减
		discountPrice = c.DiscountValue
	default:
		var rule *DiscountRule
		err2 := json.Unmarshal([]byte(c.DiscountRule), &rule)
		if err2 != nil {
			return
		}
		ruleM := int32(_func.Itoa(rule.M))
		ruleN, _ := strconv.ParseFloat(rule.N, 64)
		switch c.DiscountType {
		case 3: //第[M]件打折[N]%
			if p.qty >= ruleM {
				discountPrice = p.price * ruleN / 100
			}
		case 4: //买[M]件送[N]件
		case 5: //满[M]件立减¥[N]
			if p.qty >= ruleM {
				discountPrice = ruleN
			}
		case 6: //满¥[M]立减¥[N]
			ruleMM, _ := strconv.ParseFloat(rule.M, 64)
			totalPrice := decimal.NewFromFloat(p.totalPrice)
			leastTotal := decimal.NewFromFloat(ruleMM)
			if totalPrice.GreaterThanOrEqual(leastTotal) {
				discountPrice = ruleN
			}
		case 7: //满[M]件打折[N]%
			if p.qty >= ruleM {
				discountPrice = p.totalPrice * ruleN / 100
			}
		}
	}
	return discountPrice
}
