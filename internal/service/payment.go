package service

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	pb "studyRoomGo/api/payment/v1"
	_func "studyRoomGo/common/func"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

type PaymentService struct {
	pb.UnimplementedPaymentServer
	uc            *biz.PaymentUsecase
	packageCardUc *biz.PackageCardUsecase
	orderUc       *biz.OrderUsecase
	couponUc      *biz.CouponUsecase
	courseUc      *biz.CourseUsecase
}

func NewPaymentService(uc *biz.PaymentUsecase, packageCardUc *biz.PackageCardUsecase, orderUc *biz.OrderUsecase, couponUc *biz.CouponUsecase, courseUc *biz.CourseUsecase) *PaymentService {
	return &PaymentService{uc: uc, packageCardUc: packageCardUc, orderUc: orderUc, couponUc: couponUc, courseUc: courseUc}
}

func (s *PaymentService) WxPayCheckout(ctx context.Context, req *pb.WxCheckoutRequest) (*pb.WxCheckoutReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	rep, err := s.uc.WxPaymentCheckout(ctx, &biz.WxCheckout{
		MemberId: memberId,
		OrderNo:  req.OrderNo,
	})
	if err != nil {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.WxCheckoutReply{
		Code: 200,
		Msg:  "",
		Data: &pb.WxCheckoutReply_Data{
			TimeStamp: rep.TimeStamp,
			NonceStr:  rep.NonceStr,
			Package:   rep.Package,
			SignType:  rep.SignType,
			PaySign:   rep.PaySign,
		},
	}, nil
}

func (s *PaymentService) WxPayCheckoutReturn(ctx context.Context, req *pb.WxCheckoutReturnRequest) (*pb.WxCheckoutReturnReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	// if config2.AppConf.Environment != "pro" { //测试用
	// 	order, _ := s.uc.OrderInfo(ctx, req.OrderNo)
	// 	s.PaymentCallback(ctx, &biz.PaymentCallback{
	// 		OrderNo:       req.OrderNo,
	// 		TransactionId: "",
	// 		PaymentStatus: biz.PaymentStatusPayed,
	// 		ActiveAmount:  order.Price,
	// 	})
	// 	if req.Status == 2 { //取消支付
	// 		_, _ = s.couponUc.ReturnCoupon(ctx, &biz.CouponReturnRequest{
	// 			OrderNo:    req.OrderNo,
	// 			MemberId:   memberId,
	// 			MerchantId: int64(merchantId),
	// 			ShopId:     int64(shopId),
	// 		})
	// 	}
	// 	return &pb.WxCheckoutReturnReply{
	// 		Code: 200,
	// 		Msg:  "",
	// 	}, nil
	// }
	if ok, err := s.uc.WxPaymentCheckoutReturn(ctx, &biz.WxCheckoutReturn{
		MemberId:   memberId,
		OutTradeNo: req.OrderNo,
		Status:     int(req.Status),
	}); !ok {
		return &pb.WxCheckoutReturnReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if req.Status == 2 { //取消支付
		_, _ = s.couponUc.ReturnCoupon(ctx, &biz.CouponReturnRequest{
			OrderNo:    req.OrderNo,
			MemberId:   memberId,
			MerchantId: int64(merchantId),
			ShopId:     int64(shopId),
		})
	}
	return &pb.WxCheckoutReturnReply{
		Code: 200,
		Msg:  "",
	}, nil
}
func (s *PaymentService) WxPayNotice(ctx context.Context, req *pb.WxNoticeRequest) (*pb.WxNoticeReply, error) {

	defer func() {
		transactionUnLock(req.TransactionId)
	}()
	if !transactionAddLock(req.TransactionId) {
		return nil, errors.BadRequest("FAIL", "用户操作频繁，请稍后")
	}
	if ok, err := s.uc.WxPaymentNotice(ctx, &biz.WxCheckoutNotice{
		SpMchid:        req.SubMchid,
		SpAppid:        req.SubAppid,
		SubMchid:       req.SubMchid,
		SubAppid:       req.SubAppid,
		TransactionId:  req.TransactionId,
		OutTradeNo:     req.OutTradeNo,
		TradeType:      req.TradeType,
		TradeState:     req.TradeState,
		TradeStateDesc: req.TradeStateDesc,
		BankType:       req.BankType,
		SuccessTime:    req.SuccessTime,
		Payer: struct {
			OpenId    string
			SpOpenid  string
			SubOpenid string
		}{
			OpenId:    req.Payer.Openid,
			SpOpenid:  req.Payer.SpOpenid,
			SubOpenid: req.Payer.SubOpenid,
		},
		Amount: struct {
			Total         int32
			PayerTotal    int32
			Currency      string
			PayerCurrency string
		}{
			Total:         req.Amount.Total,
			PayerTotal:    req.Amount.PayerTotal,
			Currency:      req.Amount.Currency,
			PayerCurrency: req.Amount.PayerCurrency,
		},
	}); !ok {
		return nil, errors.BadRequest("FAIL", err.Error())
	}

	if ok, err := s.PaymentCallback(ctx, &biz.PaymentCallback{
		OrderNo:       req.OutTradeNo,
		TransactionId: req.TransactionId,
		PaymentStatus: biz.WxPaymentToCode[req.TradeState],
		ActiveAmount:  float64(req.Amount.Total) / 100,
	}); !ok {
		return nil, errors.BadRequest("FAIL", err.Error())
	}
	return &pb.WxNoticeReply{
		Code:    "SUCCESS",
		Message: "",
	}, nil
}

func (s *PaymentService) PaymentCallback(ctx context.Context, callback *biz.PaymentCallback) (bool, error) {
	order, err := s.uc.OrderInfo(ctx, callback.OrderNo)
	if err != nil {
		return false, err
	}
	if callback.PaymentStatus == biz.PaymentStatusPayed {
		if order.OrderType == biz.OrderTypeBuyPackageCard { //购买套餐
			var err error
			coupon, err := s.couponUc.GetCouponInfoByOrderId(ctx, order.ID)

			var GiftCardPercent int16
			if coupon != nil && coupon.IsUsed != 0 {
				if coupon.Gift.GiftCardPercent != 0 { //赠送时长
					GiftCardPercent = coupon.Gift.GiftCardPercent

				}
			}
			relateId, _ := strconv.Atoi(order.OrderTypeRelateID)

			for i := 1; i <= order.Quantity; i++ { //数量购买
				_, err = s.packageCardUc.CreateMemberPackageCard(ctx, &biz.CreateMemberPackageCardRequest{
					Idx:             int32(i),
					MerchantId:      order.MerchantID,
					ShopId:          order.ShopID,
					MemberId:        order.MemberID,
					PackageId:       int64(relateId),
					OrderNo:         callback.OrderNo,
					TypeChannel:     2,
					ActiveAmount:    callback.ActiveAmount,
					Qty:             int32(order.Quantity),
					GiftCardPercent: GiftCardPercent,
					Coupon:          coupon,
				})
				if err != nil {
					return false, err
				}
			}
			if coupon != nil {
				if coupon.DiscountType == 4 { //买[M]件送[N]件
					var rule *biz.DiscountRule
					_ = json.Unmarshal([]byte(coupon.DiscountRule), &rule)
					ruleN := _func.Itoa(rule.N)
					for i := 1; i <= ruleN; i++ { //数量购买
						_, err = s.packageCardUc.CreateMemberPackageCard(ctx, &biz.CreateMemberPackageCardRequest{
							MerchantId:   order.MerchantID,
							ShopId:       order.ShopID,
							MemberId:     order.MemberID,
							PackageId:    int64(relateId),
							OrderNo:      callback.OrderNo,
							TypeChannel:  4,
							ActiveAmount: 0,
						})
					}
				}
				if coupon.Gift.GiftPackageCard != "" { //赠送卡
					var packageCardId []string
					_ = json.Unmarshal([]byte(coupon.Gift.GiftPackageCard), &packageCardId)
					for _, id := range packageCardId {
						_, err = s.packageCardUc.CreateMemberPackageCard(ctx, &biz.CreateMemberPackageCardRequest{
							MerchantId:   order.MerchantID,
							ShopId:       order.ShopID,
							MemberId:     order.MemberID,
							PackageId:    int64(_func.Itoa(id)),
							OrderNo:      callback.OrderNo,
							TypeChannel:  4,
							ActiveAmount: 0,
						})
					}
				}
			}
			//执行分账
			merchatPay := wechat.Merchant[order.MerchantID]
			if merchatPay.IsProfitSharing {
				_ = s.orderUc.CreateProfitSharingOrder(ctx, &biz.CreateProfitSharingOrderRequest{
					MerchantId:      order.MerchantID,
					ShopId:          order.ShopID,
					RelateOrderType: biz.OrderTypeBuyPackageCard,
					RelateOrderNo:   callback.OrderNo,
					TransactionId:   callback.TransactionId,
					TotalAmount:     callback.ActiveAmount,
				})
			}

		}
		if order.OrderType == biz.OrderTypeBuyCourse { //购买课程
			var err error
			relateId := strings.Split(order.OrderTypeRelateID, "|")
			saleId, _ := strconv.Atoi(relateId[0])
			studentId, _ := strconv.Atoi(relateId[1])
			for i := 1; i <= order.Quantity; i++ { //数量购买
				_, err = s.courseUc.CreateStudentCourse(ctx, &biz.CreateStudentCourseRequest{
					Idx:          int32(i),
					MerchantId:   order.MerchantID,
					ShopId:       order.ShopID,
					MemberId:     order.MemberID,
					SaleId:       int64(saleId),
					StudentId:    int64(studentId),
					Qty:          int32(order.Quantity),
					OrderNo:      callback.OrderNo,
					IsActive:     false,
					ActiveAmount: callback.ActiveAmount,
				})
				if err != nil {
					return false, err
				}
			}
		}
	}

	return true, nil
}

// WxPayCheckoutBuyCourse 微信支付购买课程，目前在用
func (s *PaymentService) WxPayCheckoutBuyCourse(ctx context.Context, req *pb.WxCheckoutBuyCourseRequest) (*pb.WxCheckoutReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	merchantId := 0
	// shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	// if md, ok := metadata.FromServerContext(ctx); ok {
	// 	shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	// }
	qty := int32(1)
	if req.Qty != 0 {
		qty = req.Qty
	}

	_, err := s.courseUc.BuyCourseValidate(ctx, &biz.CreateBuyCourseOrderRequest{
		MemberId:  memberId,
		CourseId:  req.CourseId,
		SaleId:    req.SaleId,
		StudentId: req.StudentId,
		Qty:       qty,
	})
	if err != nil {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	//创建订单
	createInfo, err := s.orderUc.CreateBuyCourseOrder(ctx, &biz.CreateBuyCourseOrderRequest{
		MerchantId: int64(merchantId),
		MemberId:   memberId,
		CourseId:   0,
		SaleId:     req.SaleId,
		StudentId:  req.StudentId,
		Qty:        qty,
	})
	if err != nil {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	orderNo := createInfo.OrderInfo.OrderNo
	// 发起 微信支付
	rep, err := s.uc.WxPaymentCheckout(ctx, &biz.WxCheckout{
		MemberId: memberId,
		OrderNo:  orderNo,
	})
	if err != nil {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	//修改状态为 支付中
	if ok, err := s.uc.WxPaymentCheckoutReturn(ctx, &biz.WxCheckoutReturn{
		MemberId:   memberId,
		OutTradeNo: orderNo,
		Status:     1, //支付中
	}); !ok {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.WxCheckoutReply{
		Code: 200,
		Msg:  "",
		Data: &pb.WxCheckoutReply_Data{
			OrderNo:   orderNo,
			TimeStamp: rep.TimeStamp,
			NonceStr:  rep.NonceStr,
			Package:   rep.Package,
			SignType:  rep.SignType,
			PaySign:   rep.PaySign,
		},
	}, nil
}

// WxPayCheckoutBuyCard 微信支付购卡，目前在用
func (s *PaymentService) WxPayCheckoutBuyCard(ctx context.Context, req *pb.WxCheckoutBuyCardRequest) (*pb.WxCheckoutReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	qty := int32(1)
	if req.Qty != 0 {
		qty = req.Qty
	}
	if req.CouponCode != "" {
		_, err := s.couponUc.UseCouponValidate(ctx, &biz.CouponUseRequest{
			CouponCode: req.CouponCode,
			MemberId:   memberId,
			PackageId:  req.CardId,
			MerchantId: int64(merchantId),
			ShopId:     int64(shopId),
			Qty:        qty,
		})
		if err != nil {
			return &pb.WxCheckoutReply{
				Code: err.GetCode(),
				Msg:  err.GetReason(),
			}, nil
		}
	}

	_, err := s.packageCardUc.BuyCardValidate(ctx, &biz.CreateBuyCardOrderRequest{
		MemberId: memberId,
		CardId:   req.CardId,
		Qty:      qty,
	})
	if err != nil {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	//创建订单
	createInfo, err := s.orderUc.CreateBuyCardOrder(ctx, &biz.CreateBuyCardOrderRequest{
		MerchantId: int64(merchantId),
		MemberId:   memberId,
		CardId:     req.CardId,
		CouponCode: req.CouponCode,
		Qty:        qty,
	})
	if err != nil {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if req.CouponCode != "" {
		_, _ = s.couponUc.UseCoupon(ctx, &biz.CouponUseRequest{
			OrderId:    createInfo.OrderInfo.ID,
			CouponCode: req.CouponCode,
			MemberId:   memberId,
			MerchantId: int64(merchantId),
			ShopId:     int64(shopId),
		})
	}
	orderNo := createInfo.OrderInfo.OrderNo
	// 发起 微信支付
	rep, err := s.uc.WxPaymentCheckout(ctx, &biz.WxCheckout{
		MemberId: memberId,
		OrderNo:  orderNo,
	})
	if err != nil {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	//修改状态为 支付中
	if ok, err := s.uc.WxPaymentCheckoutReturn(ctx, &biz.WxCheckoutReturn{
		MemberId:   memberId,
		OutTradeNo: orderNo,
		Status:     1, //支付中
	}); !ok {
		return &pb.WxCheckoutReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.WxCheckoutReply{
		Code: 200,
		Msg:  "",
		Data: &pb.WxCheckoutReply_Data{
			OrderNo:   orderNo,
			TimeStamp: rep.TimeStamp,
			NonceStr:  rep.NonceStr,
			Package:   rep.Package,
			SignType:  rep.SignType,
			PaySign:   rep.PaySign,
		},
	}, nil
}
