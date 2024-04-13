package biz

import (
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/data/gen"
	"time"
)

const (
	_                                = iota
	MemberPackageCard_OperateActive  //激活
	MemberPackageCard_OperateFreeze  //冻结
	MemberPackageCard_OperateInvalid //失效
)
const (
	MemberPackage_StatusWaiteActive = iota
	MemberPackage_StatusNormal      //正常
	MemberPackage_StatusEnd         //过期
	MemberPackage_StatusClean       //用完
	MemberPackage_StatusFreeze      //冻结
	MemberPackage_StatusNoActive    //过期未激活

)

type MemberPackageCard struct {
	*gen.SrMemberPackageCard
	Gift struct {
		Minute int
		Day    int
	}
}

func (r *MemberPackageCard) GenerateByPackageCard(req *CreateMemberPackageCardRequest, p *PackageCard) {
	t1, _ := time.ParseInLocation(time.DateTime, "1991-07-31 00:00:00", time.Local)
	t2, _ := time.ParseInLocation(time.DateTime, "1991-07-31 00:00:00", time.Local)

	useBalance := int8(0)
	remark := string("")
	createBy := uint(0)
	status := MemberPackage_StatusWaiteActive
	if req.IsActive {
		now := time.Now()
		t1 = now
		t2 = now.Add(time.Hour * time.Duration(24*p.ValidDay))

		status = MemberPackage_StatusNormal
	}
	r.SrMemberPackageCard = &gen.SrMemberPackageCard{
		MerchantID:   req.MerchantId,
		ShopID:       req.ShopId,
		MemberID:     req.MemberId,
		ActualAmount: req.ActiveAmount,
		RemainMinute: p.BookingMinute,
		RemainDay:    p.BookingDay,
		MinutePer:    p.BookingMinutePer,
		MinuteMax:    p.BookingMinuteMax,
		MinuteMin:    p.BookingMinuteMin,
		Status:       int8(status),
		TypeChannel:  req.TypeChannel, //小程序购买
		ActiveAt:     &t1,
		InvalidAt:    &t2,
		PackageID:    &p.ID,
		PackageName:  p.Name,
		TypePackage:  p.TypePackage,
		TypeCard:     p.TypeCard,
		TypeSale:     p.TypeSale,
		Price:        p.Price,
		OriPrice:     p.OriPrice,
		UseShop:      p.UseShop,
		UseSeat:      p.UseSeat,
		UseTime:      p.UseTime,
		UseBalance:   &useBalance,
		ValidDay:     p.ValidDay,
		ActiveLimit:  p.ActiveLimit,
		Sort:         0,
		Remark:       &remark,
		CreateBy:     &createBy,
		UpdateBy:     &createBy,
	}
	r.GiftCardPercent(req.GiftCardPercent) //赠送时长
}
func (r *MemberPackageCard) GiftCardPercent(percent int16) {
	if percent == 0 {
		return
	}
	if r.TypePackage == 1 { //时卡
		oriMinute := r.RemainMinute
		giftMinute := int(_func.Round(float64(oriMinute*int(percent)/100), 0))
		r.Gift.Minute = giftMinute
		r.RemainMinute = oriMinute + giftMinute
	}
	if r.TypePackage == 2 { //次卡
		oriDay := r.RemainDay
		giftDay := int(_func.Round(float64(oriDay*int(percent)/100), 0))
		r.Gift.Day = giftDay
		r.RemainDay = oriDay + giftDay
	}
	if r.TypePackage == 3 { //天卡
		oriDay := r.ValidDay //有效日期
		giftDay := int(_func.Round(float64(oriDay*int(percent)/100), 0))
		r.Gift.Day = giftDay

		r.ValidDay = oriDay + giftDay
	}
}

type SearchMemberPackageCardRequest struct {
	ShopId    int64
	MemberId  int64
	PackageId int64
	Phone     string
	Status    []int32

	Sort     int
	Page     int
	PageSize int
}
