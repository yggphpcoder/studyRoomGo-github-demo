package biz

import (
	"encoding/json"
	"fmt"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
	"time"
)

const (
	CouponStatusReceived = 1 // 已领取

)
const (
	_                     = iota
	CouponOperateUsed     //已使用
	CouponOperateReturned //已使用
)

type Coupon struct {
	*model.Coupons
}

type CouponReceive struct {
	*model.CouponsReceive
}

type CouponUseLog struct {
	*gen.SrCouponsUseLog
}

type CouponCenterRequest struct {
	MerchantId int64
	ShopId     int64
	IsShow     *int8
	Sort       int
	Page       int
	PageSize   int
}
type SearchCouponReceiveRequest struct {
	MemberId int64
	CouponId int64
	NotUsed  int8
	Sort     int
	Page     int
	PageSize int
}

type CouponReceiveRequest struct {
	CouponCode    string
	CouponId      int64
	MemberId      int64
	MerchantId    int64
	ShopId        int64
	Notice        string
	CreateByAdmin bool
	Coupon        *model.Coupons
}

type CouponUseRequest struct {
	CouponCode string
	MemberId   int64
	PackageId  int64
	OrderId    int64
	MerchantId int64
	ShopId     int64
	Notice     string
	Qty        int32
	Coupon     *model.Coupons
}

type CouponReturnRequest struct {
	OrderNo    string
	CouponId   int64
	MemberId   int64
	MerchantId int64
	ShopId     int64
	Notice     string
	Coupon     *model.Coupons
}

func (r *CouponReceive) GenerateCouponReceive(req *CouponReceiveRequest) {
	createBy := uint(req.MemberId)
	if req.CreateByAdmin {
		createBy = uint(0)
	}
	startDate := time.Now()
	endDate := startDate.Add(time.Duration(req.Coupon.ValidDay) * time.Hour * 24)
	r.CouponsReceive = &model.CouponsReceive{
		SrCouponsReceive: gen.SrCouponsReceive{
			Code:          GenerateCouponReceiveCode(req.MerchantId, req.ShopId),
			CouponID:      req.Coupon.ID,
			MemberID:      req.MemberId,
			MerchantID:    req.MerchantId,
			ShopID:        req.ShopId,
			OrderID:       0,
			Name:          req.Coupon.Name,
			Description:   req.Coupon.Description,
			DiscountType:  req.Coupon.DiscountType,
			DiscountValue: req.Coupon.DiscountValue,
			DiscountRule:  req.Coupon.DiscountRule,
			MaxDiscount:   req.Coupon.MaxDiscount,
			MinPurchase:   req.Coupon.MinPurchase,
			StartDate:     &startDate,
			EndDate:       &endDate,
			UsageLimit:    req.Coupon.UsageLimit,
			Status:        CouponStatusReceived,
			IsUsed:        0,
			Notice:        req.Notice,
			Notes:         "",
			CreateBy:      &createBy,
			UpdateBy:      &createBy,
		},
		Applicable: gen.SrCouponsReceiveApplicable{
			ApplicableOrderType:       req.Coupon.Applicable.ApplicableOrderType,
			ApplicableShop:            req.Coupon.Applicable.ApplicableShop,
			ApplicableShopType:        req.Coupon.Applicable.ApplicableShopType,
			ApplicableSeat:            req.Coupon.Applicable.ApplicableSeat,
			ApplicableSeatType:        req.Coupon.Applicable.ApplicableSeatType,
			ApplicablePackageCard:     req.Coupon.Applicable.ApplicablePackageCard,
			ApplicablePackageCardType: req.Coupon.Applicable.ApplicablePackageCardType,
			ApplicableMemberType:      req.Coupon.Applicable.ApplicableMemberType,
		},
		Gift: gen.SrCouponsReceiveGift{
			GiftPackageCard: req.Coupon.Gift.GiftPackageCard,
			GiftCardPercent: req.Coupon.Gift.GiftCardPercent,
		},
	}
}

func GenerateCouponReceiveCode(merchantId int64, shopId int64) string {

	shopStr := strconv.Itoa(int(shopId)) //后四位取模
	shopStr = fmt.Sprintf("%04s", shopStr)

	merchantStr := strconv.Itoa(int(merchantId)) //后2位取模
	merchantStr = fmt.Sprintf("%02s", merchantStr)

	chars := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	randStr := _func.RandAllString(13, chars)
	var build strings.Builder
	build.Grow(20)
	build.WriteString("C")
	build.WriteString(merchantStr)
	build.WriteString(shopStr)
	build.WriteString(randStr)
	tag := build.String()
	return tag
}

type CouponReceiveValidate struct {
	CouponInfo    *model.Coupons
	MemberId      int64
	MemberCoupons []*model.CouponsReceive
}

func (v *CouponReceiveValidate) Validate() (ok bool, err *errors2.Error) {
	if ok, err := v.BaseValidate(); !ok {
		return false, err
	}
	if ok, err := v.MemberValidate(); !ok {
		return false, err
	}
	return true, nil
}
func (v *CouponReceiveValidate) BaseValidate() (ok bool, err *errors2.Error) {

	//优惠券状态校验
	if v.CouponInfo.Status != 1 {
		err = errors2.New(400, fmt.Sprintf("优惠券未开始"), "")
		return false, err
	}
	//活动开始时间校验
	now := time.Now()
	if now.Before(*v.CouponInfo.StartDate) || now.After(*v.CouponInfo.EndDate) {
		err = errors2.New(400, fmt.Sprintf("优惠券已过期"), "")
		return false, err
	}
	return true, nil
}
func (v *CouponReceiveValidate) MemberValidate() (ok bool, err *errors2.Error) {

	//优惠券状态校验
	if v.CouponInfo.UsageLimitPerUser != 0 {
		count := len(v.MemberCoupons)
		if count >= v.CouponInfo.UsageLimitPerUser {
			err = errors2.New(43001, fmt.Sprintf("优惠券已超过领取限制"), "")
			return false, err
		}
	}

	return true, nil
}

type CouponUseValidate struct {
	MerchantId  int64
	ShopId      int64
	CouponInfo  *model.CouponsReceive
	PackageCard *gen.SrPackageCard
	MemberId    int64
	Qty         int32
}

type CouponPreValidate struct {
	IsValid bool
	Reason  string
	Coupon  []*CouponValidate
}

type CouponValidate struct {
	Code        string
	IsValid     bool
	InvalidCode int32
	Reason      string
}

func (v *CouponUseValidate) Validate() (ok bool, err *errors2.Error) {
	if ok, err := v.BaseValidate(); !ok {
		return false, err
	}

	//是否适用订单类型校验
	if v.CouponInfo.Applicable.ApplicableOrderType == OrderTypeBuyPackageCard {
		if ok, err := v.ShopValidate(); !ok {
			return false, err
		}
		if ok, err := v.PackageCardValidate(); !ok {
			return false, err
		}
		if ok, err := v.PackageCardTypeValidate(); !ok {
			return false, err
		}
	}
	if ok, err := v.CouponTypeValidate(); !ok {
		return false, err
	}

	//是否适用座位类型校验
	//是否适用座位校验
	//是否适用店铺类型校验
	//是否适用用户类型校验
	return true, nil
}

func (v *CouponUseValidate) BaseValidate() (ok bool, err *errors2.Error) {
	//是否本人登录
	if v.CouponInfo.MemberID != v.MemberId {
		err = errors2.New(400, fmt.Sprintf("优惠券不是本人"), "")
		return false, err
	}
	//是否已使用校验
	if v.CouponInfo.IsUsed == 1 {
		err = errors2.New(400, fmt.Sprintf("优惠券已使用"), "")
		return false, err
	}
	//是否失效校验
	if v.CouponInfo.Status != 1 {
		err = errors2.New(400, fmt.Sprintf("优惠券已失效"), "")
		return false, err
	}
	//是否有效期内
	now := time.Now()
	if now.Before(*v.CouponInfo.StartDate) || now.After(*v.CouponInfo.EndDate) {
		err = errors2.New(400, fmt.Sprintf("优惠券已过期"), "")
		return false, err
	}
	return true, nil
}

func (v *CouponUseValidate) ShopValidate() (ok bool, err *errors2.Error) {
	//适用店铺校验
	if v.CouponInfo.Applicable.ApplicableShop != "" {
		var shop []string
		_ = json.Unmarshal([]byte(v.CouponInfo.Applicable.ApplicableShop), &shop)
		for _, id := range shop {
			if strings.Contains(v.PackageCard.UseShop, "shopId") {
				i := fmt.Sprintf("|%s|", id)
				if !strings.Contains(v.PackageCard.UseShop, i) {
					err = errors2.New(400, fmt.Sprintf("优惠券不适用该店铺"), "")
					return false, err
				}
			}
		}
	}
	return true, nil
}

func (v *CouponUseValidate) PackageCardTypeValidate() (ok bool, err *errors2.Error) {
	//是否适用套餐类型校验
	if v.CouponInfo.Applicable.ApplicablePackageCardType != "" {
		var packageCardType []string
		_ = json.Unmarshal([]byte(v.CouponInfo.Applicable.ApplicablePackageCardType), &packageCardType)
		typePackage := strconv.Itoa(int(v.PackageCard.TypePackage))
		if !_func.InStrSplice(packageCardType, typePackage) {
			err = errors2.New(400, fmt.Sprintf("优惠券不适用该套餐类型"), "")
			return false, err
		}
	}
	return true, nil
}

func (v *CouponUseValidate) PackageCardValidate() (ok bool, err *errors2.Error) {
	//是否适用套餐校验
	if v.CouponInfo.Applicable.ApplicablePackageCard != "" {
		var packageCardId []string
		_ = json.Unmarshal([]byte(v.CouponInfo.Applicable.ApplicablePackageCard), &packageCardId)
		pId := strconv.Itoa(int(v.PackageCard.ID))
		if !_func.InStrSplice(packageCardId, pId) {
			err = errors2.New(400, fmt.Sprintf("优惠券不适用该套餐"), "")
			return false, err
		}
	}

	return true, nil
}

type DiscountRule struct {
	M string
	N string
}

func (v *CouponUseValidate) CouponTypeValidate() (ok bool, err *errors2.Error) {
	if v.CouponInfo.DiscountType == 1 { //折扣型
		return true, nil
	}
	if v.CouponInfo.DiscountType == 2 { //立减型
		leastTotal := decimal.NewFromFloat(v.CouponInfo.DiscountValue)
		totalPrice := decimal.NewFromFloat(v.PackageCard.Price * float64(v.Qty))
		if totalPrice.LessThanOrEqual(leastTotal) {
			err = errors2.New(400, fmt.Sprintf("本单未超过¥%.2f不可用", v.CouponInfo.DiscountValue), "")
			return false, err
		}
		return true, nil

	}

	var rule *DiscountRule
	err2 := json.Unmarshal([]byte(v.CouponInfo.DiscountRule), &rule)
	if err2 != nil {
		err = errors2.New(400, fmt.Sprintf("内部错误,请联系管理员"), "")
		return false, err
	}
	ruleM := int32(_func.Itoa(rule.M))
	if err2 != nil {
		err = errors2.New(400, fmt.Sprintf("内部错误,请联系管理员"), "")
		return false, err
	}
	if _func.InInt8Splice([]int8{3, 4, 5, 7}, v.CouponInfo.DiscountType) { //3:第[M]件打折[N]%,4:买[M]件送[N]件,5:满[M]件立减¥[N],7:满[M]件打折[N]%
		if v.Qty < ruleM {
			err = errors2.New(400, fmt.Sprintf("本单未满%d件不可用", ruleM), "")
			return false, err
		}
		if v.CouponInfo.DiscountType == 5 { //	//满[M]件立减¥[N]
			ruleN, _ := strconv.ParseFloat(rule.N, 64)
			leastTotal := decimal.NewFromFloat(ruleN)

			totalPrice := decimal.NewFromFloat(v.PackageCard.Price * float64(v.Qty))
			if totalPrice.LessThanOrEqual(leastTotal) {
				err = errors2.New(400, fmt.Sprintf("本单未超过¥%.2f不可用", ruleN), "")
				return false, err
			}
		}
	}
	if v.CouponInfo.DiscountType == 6 { //	满¥[M]立减¥[N]
		ruleMM, _ := strconv.ParseFloat(rule.M, 64)
		totalPrice := decimal.NewFromFloat(v.PackageCard.Price * float64(v.Qty))
		leastTotal := decimal.NewFromFloat(ruleMM)
		if totalPrice.LessThan(leastTotal) {
			err = errors2.New(400, fmt.Sprintf("本单未满¥%.2f不可用", ruleMM), "")
			return false, err
		}
	}
	return true, nil
}
