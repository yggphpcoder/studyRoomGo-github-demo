package biz

import (
	"encoding/json"
	"strconv"
	"studyRoomGo/internal/data/gen"
)

type PackageCardMemberBuyLog struct {
	*gen.SrPackageCardMemberBuyLog
}

func (r *PackageCardMemberBuyLog) GenerateByPackageCard(p *PackageCard, m *MemberPackageCard, req *CreateMemberPackageCardRequest) {
	createBy := uint(0)
	r.SrPackageCardMemberBuyLog = &gen.SrPackageCardMemberBuyLog{}
	r.MerchantID = m.MerchantID
	r.ShopID = m.ShopID
	r.OrderNo = req.OrderNo
	r.MpID = m.ID
	r.MemberID = m.MemberID
	r.PackageCardID = p.ID
	r.ActualAmount = m.ActualAmount
	r.Price = p.Price
	r.OriPrice = p.OriPrice
	r.Name = p.Name
	r.TypePackage = p.TypePackage
	r.TypeCard = p.TypeCard
	r.UseSeat = p.UseSeat
	r.UseShop = p.UseShop
	r.TypeSale = p.TypeSale
	r.ValidDay = p.ValidDay
	r.ActiveLimit = p.ActiveLimit
	r.BuyLimit = p.BuyLimit
	r.UseTime = p.UseTime
	r.BookingDay = p.BookingDay
	r.BookingMinute = p.BookingMinute
	r.BookingMinutePer = p.BookingMinutePer
	r.BookingMinuteMax = p.BookingMinuteMax
	r.BookingMinuteMin = p.BookingMinuteMin
	r.CanBalance = p.CanBalance
	r.Remark = p.Remark
	r.CreateBy = &createBy
	r.UpdateBy = &createBy
	str := make(map[string]string)
	if req.ReceiptCode != "" {
		str["ReceiptCode"] = req.ReceiptCode
	}
	if req.GiftCardPercent != 0 {
		str["GiftCardPercent"] = strconv.Itoa(int(req.GiftCardPercent))
		if m.Gift.Minute != 0 {
			str["GiftMinute"] = strconv.Itoa(m.Gift.Minute)
		}
		if m.Gift.Day != 0 {
			str["GiftDay"] = strconv.Itoa(m.Gift.Day)
		}
	}
	if req.TypeChannel == 3 {
		str["来源"] = "赠送卡;优惠券活动"
	}
	if len(str) > 0 {
		bytes, _ := json.Marshal(str)
		jsonStr := string(bytes)
		r.Remark = &jsonStr
	}

}
