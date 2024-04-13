package notice

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	config2 "studyRoomGo/common/config"
	"studyRoomGo/common/notice/gen"
	"studyRoomGo/common/wechat"
	"time"

	"github.com/go-kratos/kratos/v2/metadata"
)

const (
	_ = iota
	TypeSendCoupon
	TypeBuyExperienceCard //体验卡购买
	TypeLaoLaXin          //老拉新,老客购买卡时通知
	TypeXuFei             //续费提醒
)

var serviceMemberId = []int64{18, 21, 33}

type Info struct {
	MerchantId int64
	ShopId     int64
	NoticeType int8
	Title      string
	Content    string
	ExtJson    map[string]string
}

type Notice struct {
	Ctx      context.Context
	MemberId int64
}

type CardInfo struct {
	Name          string
	TypeCard      int8
	TypePackage   int8
	RemainMin     int
	RemainDay     int
	InvalidAt     *time.Time
	CurrentCardId int64
	LastCardId    int64
}

func (n *Notice) CreateMemberPackageCardNotice(merchantId int64, shopId int64, cardInfo CardInfo) {
	if cardInfo.TypeCard == 0 { //老客拉新
		n.SendNotice([]int64{n.MemberId}, &Info{
			MerchantId: merchantId,
			ShopId:     shopId,
			NoticeType: TypeLaoLaXin,
			Title:      "老会员拉新享优惠",
			Content:    "老会员推荐新会员成功购买时卡,赠送所购时长的15%;购买周卡送2日、月卡送5日、季卡送3周<br>详情咨询客服哦",
			ExtJson:    nil,
		})
	}
}

func (n *Notice) SeatBookStatusSeatedNotice(cardInfo CardInfo) {

	memberMgr := gen.SrMemberMgr(config2.AppConf.Data.DB)
	mInfo, _ := memberMgr.FetchByPrimaryKey(n.MemberId)
	MemberNameRune := []rune(*mInfo.NickName + "(" + *mInfo.Phone + ")")
	MemberName := string(MemberNameRune[:20])
	CardNameRune := []rune(cardInfo.Name)
	CardName := string(CardNameRune[:20])
	wx, _ := wechat.NewMessage(wechat.MerchantToOfficalMerchantId(mInfo.MerchantID))

	if cardInfo.TypeCard == 1 { //体验卡优惠
		n.SendNotice([]int64{n.MemberId}, &Info{
			NoticeType: TypeBuyExperienceCard,
			Title:      "体验享优惠",
			Content:    "美团好评再送4小时体验券;体验后即日购任何套餐享7.5折,次日享8.3折<br>详情咨询客服哦",
			ExtJson:    nil,
		})

		_ = wx.SendServiceNoticeMessage(serviceMemberId, wechat.ServiceNotice{
			ServiceName: "新客体验券使用提醒,请及时跟进",
			MemberName:  MemberName,
			ProductName: CardName,
			CreateAt:    time.Now().Format(time.DateTime),
		})
		return
	}
	if cardInfo.CurrentCardId != cardInfo.LastCardId {
		return //不是最新的卡，不用通知
	}
	if cardInfo.TypePackage == 1 && cardInfo.RemainMin <= 120 { //时卡
		n.SendNotice([]int64{n.MemberId}, &Info{
			NoticeType: TypeXuFei,
			Title:      "续费提醒",
			Content:    fmt.Sprintf("亲,您的套餐剩余%d分钟,今天续费可享7.8折哦!<br>详情咨询客服哦", cardInfo.RemainMin),
			ExtJson:    nil,
		})

		_ = wx.SendServiceNoticeMessage(serviceMemberId, wechat.ServiceNotice{
			ServiceName: "续费提醒：时卡不足2小时",
			MemberName:  MemberName,
			ProductName: CardName,
			CreateAt:    time.Now().Format(time.DateTime),
		})
	}
	if cardInfo.TypePackage == 2 && cardInfo.RemainDay == 1 { //次卡
		n.SendNotice([]int64{n.MemberId}, &Info{
			NoticeType: TypeXuFei,
			Title:      "续费提醒",
			Content:    fmt.Sprintf("亲,您的套餐剩余%d次,今天续费可享7.8折哦!<br>详情咨询客服哦", cardInfo.RemainDay),
			ExtJson:    nil,
		})
		_ = wx.SendServiceNoticeMessage(serviceMemberId, wechat.ServiceNotice{
			ServiceName: "续费提醒：次卡不足1次",
			MemberName:  MemberName,
			ProductName: CardName,
			CreateAt:    time.Now().Format(time.DateTime),
		})
	}
	if cardInfo.TypePackage == 3 { //天卡
		now := time.Now()
		sub := cardInfo.InvalidAt.Sub(now)
		if sub <= 86400*time.Second*7 { //小于7天
			n.SendNotice([]int64{n.MemberId}, &Info{
				NoticeType: TypeXuFei,
				Title:      "续费提醒",
				Content:    fmt.Sprintf("亲,您的套餐%s即将到期,续费可享7.8折哦!<br>详情咨询客服哦", cardInfo.InvalidAt.Format(time.DateTime)),
				ExtJson:    nil,
			})
			MemberName := (*mInfo.NickName + "(" + *mInfo.Phone + ")")
			_ = wx.SendServiceNoticeMessage(serviceMemberId, wechat.ServiceNotice{
				ServiceName: "续费提醒：天卡不足7天",
				MemberName:  MemberName,
				ProductName: CardName,
				CreateAt:    time.Now().Format(time.DateTime),
			})
		}

	}
}
func (n *Notice) SendNotice(memberId []int64, info *Info) {
	merchantId := int(info.MerchantId)
	shopId := int(info.ShopId)
	if merchantId == 0 {
		if md, ok := metadata.FromServerContext(n.Ctx); ok {
			merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
		}
	}
	if shopId == 0 {
		if md, ok := metadata.FromServerContext(n.Ctx); ok {
			shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
		}
	}

	if merchantId != 1 { //非鱼阅
		if info.NoticeType != TypeSendCoupon {
			return
		}
	}
	var Batch []*gen.SrMemberNotice
	for _, mId := range memberId {
		var data gen.SrMemberNotice
		data.IsRead = 0
		data.NoticeType = info.NoticeType
		data.MerchantID = int64(merchantId)
		data.ShopID = int64(shopId)
		data.MemberID = mId
		data.Title = info.Title
		data.Content = info.Content
		b, _ := json.Marshal(info.ExtJson)
		data.ExtJSON = string(b)
		Batch = append(Batch, &data)
	}
	memberMgr := gen.SrMemberNoticeMgr(config2.AppConf.Data.DB)

	err := memberMgr.CreateInBatches(Batch, 100).Error
	if err != nil {
		fmt.Printf("Send Notice Error %v", err)
	}
}
