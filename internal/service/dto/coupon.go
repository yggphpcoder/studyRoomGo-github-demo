package dto

import (
	"strings"
	pb "studyRoomGo/api/coupon/v1"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/model"
	"time"
)

type CouponCenter struct {
	*model.Coupons
}

func (r *CouponCenter) CouponCenterReply() *pb.CouponData {
	return &pb.CouponData{
		Code:                           r.Code,
		Name:                           r.Name,
		Description:                    r.Description,
		DiscountType:                   int32(r.DiscountType),
		DiscountValue:                  int32(r.DiscountValue),
		DiscountRule:                   r.DiscountRule,
		MaxDiscount:                    int32(r.MaxDiscount),
		StartDate:                      r.StartDate.Format(time.DateTime),
		EndDate:                        r.EndDate.Format(time.DateTime),
		ValidDay:                       int32(r.ValidDay),
		UsageLimitPerUser:              int32(r.UsageLimitPerUser),
		Status:                         int32(r.Status),
		ApplicableShop:                 r.Applicable.ApplicableShop,
		ApplicableSeat:                 r.Applicable.ApplicableSeat,
		ApplicableSeatType:             r.Applicable.ApplicableSeatType,
		ApplicablePackageCard:          r.Applicable.ApplicablePackageCard,
		ApplicablePackageCardType:      r.Applicable.ApplicablePackageCardType,
		ApplicablePackageCardLabel:     strings.Join(r.PackageLabel, ","),
		ApplicablePackageCardTypeLabel: strings.Join(r.PackageTypeLabel, ","),
		ApplicableShopLabel:            strings.Join(r.ShopLabel, ","),
		GiftCardPercent:                int32(r.Gift.GiftCardPercent),
		GiftPackageCard:                r.Gift.GiftPackageCard,
		GiftPackageCardLabel:           strings.Join(r.GiftPackageCardLabel, ","),
	}
}

type MyCoupon struct {
	*model.CouponsReceive
}

func (r *MyCoupon) CouponReply() *pb.CouponData {
	return &pb.CouponData{
		Code:                           r.Code,
		MemberId:                       r.MemberID,
		Name:                           r.Name,
		Description:                    r.Description,
		DiscountType:                   int32(r.DiscountType),
		DiscountValue:                  int32(r.DiscountValue),
		DiscountRule:                   r.DiscountRule,
		MaxDiscount:                    int32(r.MaxDiscount),
		StartDate:                      r.StartDate.Format(time.DateTime),
		EndDate:                        r.EndDate.Format(time.DateTime),
		Status:                         int32(r.Status),
		IsUsed:                         int32(r.IsUsed),
		ApplicableShop:                 r.Applicable.ApplicableShop,
		ApplicableSeat:                 r.Applicable.ApplicableSeat,
		ApplicableSeatType:             r.Applicable.ApplicableSeatType,
		ApplicablePackageCard:          r.Applicable.ApplicablePackageCard,
		ApplicablePackageCardType:      r.Applicable.ApplicablePackageCardType,
		ApplicablePackageCardLabel:     strings.Join(r.PackageLabel, ","),
		ApplicablePackageCardTypeLabel: strings.Join(r.PackageTypeLabel, ","),
		ApplicableShopLabel:            strings.Join(r.ShopLabel, ","),
		GiftCardPercent:                int32(r.Gift.GiftCardPercent),
		GiftPackageCard:                r.Gift.GiftPackageCard,
		GiftPackageCardLabel:           strings.Join(r.GiftPackageCardLabel, ","),
	}
}

type Coupon struct {
	CouponReceive *biz.CouponReceive
}

func (r *Coupon) CouponPreValidateReply(res *biz.CouponPreValidate) *pb.CouponUseValidateReply_Data {

	var couponlist []*pb.CouponUseValidateReply_Coupon

	for _, coupon := range res.Coupon {
		couponlist = append(couponlist, &pb.CouponUseValidateReply_Coupon{
			Code:        coupon.Code,
			IsValid:     coupon.IsValid,
			InvalidCode: coupon.InvalidCode,
			Reason:      coupon.Reason,
		})
	}
	return &pb.CouponUseValidateReply_Data{
		IsValid: res.IsValid,
		Reason:  res.Reason,
		Coupon:  couponlist,
	}
}
