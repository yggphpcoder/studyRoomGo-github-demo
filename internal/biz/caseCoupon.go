package biz

import (
	"context"
	"encoding/json"
	"fmt"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"strconv"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
	"time"
)

// ShopUsecase is a Shop usecase.
type CouponUsecase struct {
	repo        CouponRepo
	orderRepo   OrderRepo
	packageRepo PackageCardRepo
	shopRepo    ShopRepo
	log         *log.Helper
}

// NewCouponUsecase new a Coupon usecase.
func NewCouponUsecase(repo CouponRepo, orderRepo OrderRepo, packageRepo PackageCardRepo, shopRepo ShopRepo, logger log.Logger) *CouponUsecase {
	return &CouponUsecase{
		repo:        repo,
		orderRepo:   orderRepo,
		packageRepo: packageRepo,
		shopRepo:    shopRepo,
		log:         log.NewHelper(logger),
	}
}
func (uc *CouponUsecase) CouponCenter(ctx context.Context, req *CouponCenterRequest) (list []*model.Coupons, error error) {

	data, err := uc.repo.ListBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var pName = make(map[int64]string)
	var pTypeName = map[int64]string{
		1: "时卡",
		2: "次卡",
		3: "天卡",
	}
	var searchPackageCard []string
	for _, item := range data {
		if item.Applicable.ApplicablePackageCard != "" {
			var packageCardId []string
			_ = json.Unmarshal([]byte(item.Applicable.ApplicablePackageCard), &packageCardId)
			searchPackageCard = append(searchPackageCard, packageCardId...) //适用套餐label
		}
		if item.Gift.GiftPackageCard != "" {
			var packageCardId []string
			_ = json.Unmarshal([]byte(item.Gift.GiftPackageCard), &packageCardId)
			searchPackageCard = append(searchPackageCard, packageCardId...) //适用套餐label
		}
	}
	if len(data) > 0 {
		cardList, _ := uc.packageRepo.ListBy(ctx, &SearchPackageCardRequest{
			PId:        searchPackageCard,
			MerchantId: req.MerchantId,
			Status:     2,
		})
		for _, p := range cardList {
			pName[p.ID] = p.Name
		}
	}
	shopNameMap := make(map[int64]string)
	shopList, _ := uc.shopRepo.ListBy(ctx, &SearchShopRequest{MerchantId: req.MerchantId})
	for _, shop := range shopList {
		shopNameMap[shop.ID] = shop.Name
	}
	for _, item := range data {
		if item.Applicable.ApplicablePackageCard != "" {
			var packageCardId []string
			_ = json.Unmarshal([]byte(item.Applicable.ApplicablePackageCard), &packageCardId)
			for _, id := range packageCardId {
				id, _ := strconv.Atoi(id)
				item.PackageLabel = append(item.PackageLabel, pName[int64(id)])
			}
		}
		if item.Applicable.ApplicablePackageCardType != "" {
			var packageCardType []string
			_ = json.Unmarshal([]byte(item.Applicable.ApplicablePackageCardType), &packageCardType)
			for _, id := range packageCardType {
				id, _ := strconv.Atoi(id)
				item.PackageTypeLabel = append(item.PackageTypeLabel, pTypeName[int64(id)])
			}
		}

		if item.Applicable.ApplicableShop != "" {
			var shop []string
			_ = json.Unmarshal([]byte(item.Applicable.ApplicableShop), &shop)
			for _, id := range shop {
				id, _ := strconv.Atoi(id)
				item.ShopLabel = append(item.ShopLabel, shopNameMap[int64(id)])
			}
		}

		if item.Gift.GiftPackageCard != "" {
			var packageCardId []string
			_ = json.Unmarshal([]byte(item.Gift.GiftPackageCard), &packageCardId)
			for _, id := range packageCardId {
				id, _ := strconv.Atoi(id)
				item.GiftPackageCardLabel = append(item.GiftPackageCardLabel, pName[int64(id)])
			}
		}
		list = append(list, item)
	}

	return list, nil
}
func (uc *CouponUsecase) MyCouponList(ctx context.Context, req *SearchCouponReceiveRequest) (list []*model.CouponsReceive, error error) {
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}

	data, err := uc.repo.ListReceiveBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var pName = make(map[int64]string)
	var pTypeName = map[int64]string{
		1: "时卡",
		2: "次卡",
		3: "天卡",
	}
	var searchPackageCard []string
	for _, item := range data {
		if item.Applicable.ApplicablePackageCard != "" {
			var packageCardId []string
			_ = json.Unmarshal([]byte(item.Applicable.ApplicablePackageCard), &packageCardId)
			searchPackageCard = append(searchPackageCard, packageCardId...)
		}
		if item.Gift.GiftPackageCard != "" {
			var packageCardId []string
			_ = json.Unmarshal([]byte(item.Gift.GiftPackageCard), &packageCardId)
			searchPackageCard = append(searchPackageCard, packageCardId...) //适用套餐label
		}
	}
	if len(data) > 0 {
		cardList, _ := uc.packageRepo.ListBy(ctx, &SearchPackageCardRequest{
			PId:        searchPackageCard,
			MerchantId: int64(merchantId),
			Status:     2,
		})
		for _, p := range cardList {
			pName[p.ID] = p.Name
		}
	}
	shopNameMap := make(map[int64]string)
	shopList, _ := uc.shopRepo.ListBy(ctx, &SearchShopRequest{MerchantId: int64(merchantId)})
	for _, shop := range shopList {
		shopNameMap[shop.ID] = shop.Name
	}
	var ExpiredCoupon []int64
	now := time.Now()
	for _, item := range data {
		if item.Status == 1 && item.IsUsed == 0 && now.After(*item.EndDate) { //已领取，未使用，过期
			ExpiredCoupon = append(ExpiredCoupon, item.ID)
			item.Status = 2
		}
		if item.Applicable.ApplicablePackageCard != "" {
			var packageCardId []string
			_ = json.Unmarshal([]byte(item.Applicable.ApplicablePackageCard), &packageCardId)
			for _, id := range packageCardId {
				id, _ := strconv.Atoi(id)
				item.PackageLabel = append(item.PackageLabel, pName[int64(id)])
			}
		}
		if item.Applicable.ApplicablePackageCardType != "" {
			var packageCardType []string
			_ = json.Unmarshal([]byte(item.Applicable.ApplicablePackageCardType), &packageCardType)
			for _, id := range packageCardType {
				id, _ := strconv.Atoi(id)
				item.PackageTypeLabel = append(item.PackageTypeLabel, pTypeName[int64(id)])
			}
		}
		if item.Applicable.ApplicableShop != "" {
			var shop []string
			_ = json.Unmarshal([]byte(item.Applicable.ApplicableShop), &shop)
			for _, id := range shop {
				id, _ := strconv.Atoi(id)
				item.ShopLabel = append(item.ShopLabel, shopNameMap[int64(id)])
			}
		}

		if item.Gift.GiftPackageCard != "" {
			var packageCardId []string
			_ = json.Unmarshal([]byte(item.Gift.GiftPackageCard), &packageCardId)
			for _, id := range packageCardId {
				id, _ := strconv.Atoi(id)
				item.GiftPackageCardLabel = append(item.GiftPackageCardLabel, pName[int64(id)])
			}
		}
		list = append(list, item)
	}
	if len(ExpiredCoupon) > 0 {
		_, _ = uc.repo.UpdateCouponReceiveExpired(ctx, ExpiredCoupon) //更新已过期
	}

	return list, nil
}

func (uc *CouponUsecase) GetCouponInfoByOrderId(ctx context.Context, orderId int64) (*model.CouponsReceive, error) {
	data, err := uc.repo.FindReceiveByOrderId(ctx, orderId)
	return data, err
}

func (uc *CouponUsecase) ReceiveCoupon(ctx context.Context, req *CouponReceiveRequest) (c *model.Coupons, err error) {
	var couponInfo *model.Coupons
	if req.CouponId != 0 {
		couponInfo, err = uc.repo.FindByID(ctx, req.CouponId)
	} else {
		couponInfo, err = uc.repo.FindByCode(ctx, req.CouponCode)
	}
	if err != nil {
		log.Errorf("Service ReceiveCoupon error:%s \r\n", err)
		return nil, err
	}
	merchantId := couponInfo.MerchantID
	shopId := couponInfo.ShopID
	if req.MerchantId != 0 {
		merchantId = req.MerchantId
	}
	if req.ShopId != 0 {
		shopId = req.ShopId
		shop, _ := uc.shopRepo.FindByID(ctx, shopId)
		couponInfo.ShopLabel = []string{shop.Name}
	}
	coupon := &CouponReceive{}
	coupon.GenerateCouponReceive(&CouponReceiveRequest{
		CouponCode:    req.CouponCode,
		MemberId:      req.MemberId,
		MerchantId:    merchantId,
		ShopId:        shopId,
		Notice:        req.Notice,
		Coupon:        couponInfo,
		CreateByAdmin: req.CreateByAdmin,
	})

	if ok, err := uc.repo.SaveCouponReceive(ctx, coupon.CouponsReceive); !ok {
		log.Errorf("Log Insert error:%s \r\n", err)
		return nil, err
	}
	return couponInfo, nil
}

// ReceiveCouponValidate 领取coupon 校验
func (uc *CouponUsecase) ReceiveCouponValidate(ctx context.Context, req *CouponReceiveRequest) (ok bool, err *errors2.Error) {
	var couponInfo *model.Coupons
	var err2 error
	if req.CouponId != 0 {
		couponInfo, err2 = uc.repo.FindByID(ctx, req.CouponId)
	} else {
		couponInfo, err2 = uc.repo.FindByCode(ctx, req.CouponCode)
	}
	if err2 != nil {
		err = errors2.New(400, fmt.Sprintf("优惠券不存在"), "")
	}
	memberCoupons, _ := uc.repo.ListReceiveBy(ctx, &SearchCouponReceiveRequest{
		CouponId: couponInfo.ID,
		MemberId: req.MemberId,
	})
	validate := &CouponReceiveValidate{
		CouponInfo:    couponInfo,
		MemberId:      req.MemberId,
		MemberCoupons: memberCoupons,
	}
	if ok, err := validate.Validate(); !ok {
		return false, err
	}
	return true, nil
}

func (uc *CouponUsecase) UseCouponPreValidate(ctx context.Context, req *CouponUseRequest) (res *CouponPreValidate, err *errors2.Error) {

	res = &CouponPreValidate{IsValid: true}
	packageCardInfo, err2 := uc.packageRepo.FindByID(ctx, req.PackageId)
	if err2 != nil {
		err = errors2.New(400, fmt.Sprintf("套餐不存在"), "")
	}
	couponList, err2 := uc.MyCouponList(ctx, &SearchCouponReceiveRequest{
		MemberId: req.MemberId,
		NotUsed:  1,
	})
	for _, coupon := range couponList {
		var couponErr *errors2.Error
		var ok bool
		validate := &CouponUseValidate{
			MerchantId:  req.MerchantId,
			ShopId:      req.ShopId,
			CouponInfo:  coupon,
			PackageCard: packageCardInfo,
			MemberId:    req.MemberId,
			Qty:         req.Qty,
		}
		ok, couponErr = validate.Validate()
		couponValid := &CouponValidate{
			Code:        coupon.Code,
			IsValid:     true,
			InvalidCode: 0,
		}
		if !ok {
			couponValid.IsValid = false
			couponValid.Reason = couponErr.GetReason()
			couponValid.InvalidCode = couponErr.GetCode()
			goto res
		}
	res:
		res.Coupon = append(res.Coupon, couponValid)

	}
	return res, nil

}

// UseCouponValidate 使用coupon 校验
func (uc *CouponUsecase) UseCouponValidate(ctx context.Context, req *CouponUseRequest) (ok bool, err *errors2.Error) {
	couponInfo, err2 := uc.repo.FindReceiveByCode(ctx, req.CouponCode)
	if err2 != nil {
		err = errors2.New(400, fmt.Sprintf("优惠券不存在"), "")
	}
	packageCardInfo, err2 := uc.packageRepo.FindByID(ctx, req.PackageId)
	if err2 != nil {
		err = errors2.New(400, fmt.Sprintf("套餐不存在"), "")
	}
	validate := &CouponUseValidate{
		MerchantId:  req.MerchantId,
		ShopId:      req.ShopId,
		CouponInfo:  couponInfo,
		PackageCard: packageCardInfo,
		MemberId:    req.MemberId,
		Qty:         req.Qty,
	}
	if ok, err = validate.Validate(); !ok {
		return false, err
	}
	return true, nil
}

// UseCoupon 使用coupon
func (uc *CouponUsecase) UseCoupon(ctx context.Context, req *CouponUseRequest) (ok bool, err *errors2.Error) {
	remark := ""
	action := CouponOperateUsed
	var couponInfo *model.CouponsReceive
	var err2 error
	couponInfo, err2 = uc.repo.FindReceiveByCode(ctx, req.CouponCode)
	if err2 != nil {
		remark = err2.Error()
		goto remark
	}
	couponInfo.IsUsed = 1
	couponInfo.OrderID = req.OrderId
	if ok, err2 = uc.repo.UpdateCouponReceive(ctx, couponInfo); !ok {
		remark = err2.Error()
		goto remark
	}

remark:
	CreateBy := uint(req.MemberId)
	l := &gen.SrCouponsUseLog{
		MemberID:         req.MemberId,
		MerchantID:       req.MerchantId,
		ShopID:           req.ShopId,
		OrderID:          req.OrderId,
		CouponID:         0,
		CouponsReceiveID: 0,
		Action:           int8(action),
		Remark:           &remark,
		CreateBy:         &CreateBy,
	}
	if couponInfo != nil {
		l.CouponID = couponInfo.CouponID
		l.CouponsReceiveID = couponInfo.ID
	}
	_ = uc.repo.RecordUseLog(ctx, &CouponUseLog{
		SrCouponsUseLog: l,
	})
	return true, nil
}

// ReturnCoupon 退还coupon
func (uc *CouponUsecase) ReturnCoupon(ctx context.Context, req *CouponReturnRequest) (ok bool, err *errors2.Error) {
	remark := ""
	action := CouponOperateReturned
	var order *gen.SrOrder
	var couponInfo *model.CouponsReceive
	var err2 error
	order, err2 = uc.orderRepo.FindByOrderNo(ctx, req.OrderNo)
	if err2 != nil {
		if gen.IsNotFound(err2) {
			return true, nil
		}
		order.ID = 0
		remark = err2.Error()
		goto remark
	}
	couponInfo, err2 = uc.repo.FindReceiveByOrderId(ctx, order.ID)
	if err2 != nil {
		if gen.IsNotFound(err2) {
			return true, nil
		}
		remark = err2.Error()
		goto remark
	}
	couponInfo.IsUsed = 0
	couponInfo.OrderID = 0
	if ok, err2 = uc.repo.UpdateCouponReceive(ctx, couponInfo); !ok {
		remark = err2.Error()
		goto remark
	}

remark:
	CreateBy := uint(req.MemberId)
	l := &gen.SrCouponsUseLog{
		MemberID:         req.MemberId,
		MerchantID:       req.MerchantId,
		ShopID:           req.ShopId,
		OrderID:          order.ID,
		CouponID:         0,
		CouponsReceiveID: 0,
		Action:           int8(action),
		Remark:           &remark,
		CreateBy:         &CreateBy,
	}
	if couponInfo != nil {
		l.CouponID = couponInfo.CouponID
		l.CouponsReceiveID = couponInfo.ID
	}
	_ = uc.repo.RecordUseLog(ctx, &CouponUseLog{
		SrCouponsUseLog: l,
	})
	return true, nil
}
