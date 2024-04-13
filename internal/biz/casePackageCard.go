package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	_func "studyRoomGo/common/func"
	"studyRoomGo/common/notice"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
)

// PackageCardUsecase is a PackageCard usecase.
type PackageCardUsecase struct {
	repo               PackageCardRepo
	memberPackCardRepo MemberPackageCardRepo
	orderRepo          OrderRepo
	shopRepo           ShopRepo
	seatRepo           SeatRepo
	memberRepo         MemberRepo
	log                *log.Helper
}

// NewPackageCardUsecase new a PackageCard usecase.
func NewPackageCardUsecase(repo PackageCardRepo, memberPackCardRepo MemberPackageCardRepo, orderRepo OrderRepo, shopRepo ShopRepo, seatRepo SeatRepo, memberRepo MemberRepo, logger log.Logger) *PackageCardUsecase {
	return &PackageCardUsecase{
		repo:               repo,
		memberPackCardRepo: memberPackCardRepo,
		orderRepo:          orderRepo,
		shopRepo:           shopRepo,
		seatRepo:           seatRepo,
		memberRepo:         memberRepo,
		log:                log.NewHelper(logger),
	}
}

type PackageCard struct {
	*gen.SrPackageCard
	UseShopStr string
	UseSeatStr string
}

func (uc *PackageCardUsecase) FindById(ctx context.Context, id int64) (*PackageCard, error) {
	data, err := uc.repo.FindByID(ctx, id)

	shopNameMap := make(map[int64]string)
	seatTypeNameMap := make(map[int64]string)
	shopList, _ := uc.shopRepo.ListAll(ctx)
	for _, shop := range shopList {
		shopNameMap[shop.ID] = shop.Name
	}
	seatTypeList, _ := uc.seatRepo.SeatTypeListAll(ctx)
	for _, seat := range seatTypeList {
		seatTypeNameMap[seat.ID] = seat.Name
	}
	var shopNameArr []string
	var seatNameArr []string
	shopArr := strings.Split(data.UseShop, "|")
	if shopArr[0] == "shopId" {
		for i := 1; i < len(shopArr); i++ {
			if shopArr[i] != "" {
				id, _ := strconv.Atoi(shopArr[i])
				if id != 0 {
					shopNameArr = append(shopNameArr, shopNameMap[int64(id)])
				} else {
					shopNameArr = append(shopNameArr, "通用")
				}
			}

		}
	}
	var seat []string
	_ = json.Unmarshal([]byte(data.UseSeat), &seat)

	for _, id := range seat {
		id, _ := strconv.Atoi(id)
		if id != 0 {
			seatNameArr = append(seatNameArr, seatTypeNameMap[int64(id)])
		} else {
			seatNameArr = append(seatNameArr, "通用")
		}
	}

	if err != nil {
		return nil, err
	}
	return &PackageCard{
		SrPackageCard: data,
		UseShopStr:    strings.Join(shopNameArr, ","),
		UseSeatStr:    strings.Join(seatNameArr, ","),
	}, nil
}

func (uc *PackageCardUsecase) FindBy(ctx context.Context, key string, value string) (*PackageCard, error) {
	data, err := uc.repo.FindBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	return &PackageCard{SrPackageCard: data}, nil

}

func (uc *PackageCardUsecase) List(ctx context.Context) ([]*PackageCard, error) {
	data, err := uc.repo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	var list []*PackageCard
	for _, item := range data {
		list = append(list, &PackageCard{SrPackageCard: item})
	}
	return list, nil
}

type SearchPackageCardRequest struct {
	PId         []string
	MerchantId  int64
	TypePackage int32
	TypeCard    int32
	TypeSale    int32
	Status      int32
	UseShop     string
	UseSeat     string
	Sort        int
}

func (uc *PackageCardUsecase) Search(ctx context.Context, req *SearchPackageCardRequest) ([]*PackageCard, error) {
	data, err := uc.repo.ListBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*PackageCard

	shopNameMap := make(map[int64]string)
	seatTypeNameMap := make(map[int64]string)
	shopList, _ := uc.shopRepo.ListAll(ctx)
	for _, shop := range shopList {
		shopNameMap[shop.ID] = shop.Name
	}
	seatTypeList, _ := uc.seatRepo.SeatTypeListAll(ctx)
	for _, seat := range seatTypeList {
		seatTypeNameMap[seat.ID] = seat.Name
	}
	for _, item := range data {
		var shopNameArr []string
		var seatNameArr []string
		shopArr := strings.Split(item.UseShop, "|")
		if shopArr[0] == "shopId" {
			for i := 1; i < len(shopArr); i++ {
				if shopArr[i] != "" {
					id, _ := strconv.Atoi(shopArr[i])
					if id != 0 {
						shopNameArr = append(shopNameArr, shopNameMap[int64(id)])
					} else {
						shopNameArr = append(shopNameArr, "通用")
					}
				}
			}
		}
		var seat []string
		_ = json.Unmarshal([]byte(item.UseSeat), &seat)

		for _, id := range seat {
			id, _ := strconv.Atoi(id)
			if id != 0 {
				seatNameArr = append(seatNameArr, seatTypeNameMap[int64(id)])
			} else {
				seatNameArr = append(seatNameArr, "通用")
			}
		}

		list = append(list, &PackageCard{
			SrPackageCard: item,
			UseShopStr:    strings.Join(shopNameArr, ","),
			UseSeatStr:    strings.Join(seatNameArr, ","),
		})
	}
	return list, nil
}

type CreateMemberPackageCardRequest struct {
	Idx          int32 //创建时的索引
	MerchantId   int64
	ShopId       int64
	MemberId     int64
	PackageId    int64
	Qty          int32 //数量
	OrderNo      string
	TypeChannel  int8
	IsActive     bool
	ActiveAmount float64
	ReceiptCode  string //关联团购码

	GiftCardPercent int16 //赠送时长

	Coupon *model.CouponsReceive
}

func (uc *PackageCardUsecase) CreateMemberPackageCard(ctx context.Context, req *CreateMemberPackageCardRequest) (id int64, error error) {
	cardInfo, err := uc.repo.FindByID(ctx, req.PackageId)
	if err != nil {
		log.Errorf("Service GetSrPackageCard error:%s \r\n", err)
		return 0, err
	}
	mPackageCard := &MemberPackageCard{}

	perAmount := 0.00
	if req.ActiveAmount != 0 {
		if req.Coupon != nil && req.Coupon.DiscountType == 3 { //第[M]件打折[N]%
			var rule *DiscountRule
			_ = json.Unmarshal([]byte(req.Coupon.DiscountRule), &rule)
			ruleM := int32(_func.Itoa(rule.M))
			ruleN, _ := strconv.ParseFloat(rule.N, 64)
			perAmount = cardInfo.Price
			if req.Idx == ruleM {
				perAmount -= cardInfo.Price * ruleN / 100 //第[M]件打折[N]%
			}
		} else {
			decimalD1 := decimal.NewFromFloat(req.ActiveAmount)
			decimalD2 := decimal.NewFromFloat(float64(req.Qty))
			perAmount = _func.Round(decimalD1.Div(decimalD2).InexactFloat64(), 2) //保留2位小数
			if req.Idx == req.Qty {
				perAmount = req.ActiveAmount - (perAmount * float64(req.Qty-1))
			}
		}
	}
	req.ActiveAmount = perAmount //实收单价
	mPackageCard.GenerateByPackageCard(req, &PackageCard{SrPackageCard: cardInfo})
	_, err = uc.memberPackCardRepo.Save(ctx, mPackageCard)
	if err != nil {
		log.Errorf("SrMemberPackageCardService Insert error:%s \r\n", err)
		return 0, err
	}
	buyLog := &PackageCardMemberBuyLog{}
	buyLog.GenerateByPackageCard(&PackageCard{SrPackageCard: cardInfo}, mPackageCard, req)
	ok, err := uc.repo.RecordBuyLog(ctx, buyLog)
	if !ok {
		log.Errorf("Log Insert error:%s \r\n", err)
		return 0, err
	}
	if req.Idx == req.Qty {
		//站内通知
		n := notice.Notice{
			Ctx:      ctx,
			MemberId: req.MemberId,
		}
		go n.CreateMemberPackageCardNotice(req.MerchantId, req.ShopId, notice.CardInfo{
			Name:     cardInfo.Name,
			TypeCard: cardInfo.TypeCard,
		})
	}

	//更新用户类型
	member, _ := uc.memberRepo.FindByID(ctx, req.MemberId)
	if member.Type == 0 {
		member.Type = 2
		if ok, err := uc.memberRepo.Update(ctx, &Member{SrMember: member}); !ok {
			log.Errorf("GenerateByPackageCard  memberRepo Update error:%s \r\n", err)
		}
	}

	return mPackageCard.ID, nil
}

func (uc *PackageCardUsecase) BuyCardValidate(ctx context.Context, req *CreateBuyCardOrderRequest) (ok bool, err error) {
	cardInfo, err := uc.repo.FindByID(ctx, req.CardId)
	if err != nil {
		log.Errorf("Service GetSrPackageCard error:%s \r\n", err)
		return false, err
	}
	if cardInfo.Status != 2 {
		return false, errors.New("套餐已关闭，请联系管理员")
	}
	if cardInfo.BuyLimit > 0 {
		if req.Qty > int32(cardInfo.BuyLimit) {
			return false, errors.New(fmt.Sprintf("该套餐限购 %d 张，请联系管理员", cardInfo.BuyLimit))
		}
		list, err := uc.memberPackCardRepo.ListBy(ctx, &SearchMemberPackageCardRequest{MemberId: req.MemberId, PackageId: req.CardId})
		if err != nil {
			return false, err
		}
		if len(list) >= cardInfo.BuyLimit {
			return false, errors.New(fmt.Sprintf("该套餐限购 %d 张，请联系管理员", cardInfo.BuyLimit))
		}

	}
	return true, nil
}

type UpdateMemberPackageCardRequest struct {
	MpId     int64
	Status   int8
	ActiveAt *time.Time
}

func (uc *PackageCardUsecase) FindMemberCardById(ctx context.Context, memberId int64, id int64) (*MemberPackageCard, error) {
	data, err := uc.memberPackCardRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if data.MemberID != memberId {
		return nil, errors.New("认证失败")
	}
	return &MemberPackageCard{SrMemberPackageCard: data}, nil
}
func (uc *PackageCardUsecase) UpdateMemberPackageCardStatus(ctx context.Context, req *UpdateMemberPackageCardRequest) (ok bool, error error) {
	cardInfo, err := uc.memberPackCardRepo.FindByID(ctx, req.MpId)
	memberPackageCard := &MemberPackageCard{
		SrMemberPackageCard: cardInfo,
	}
	operateLog := &MemberPackageCardOperateLog{
		MpId:      cardInfo.ID,
		MemberId:  cardInfo.MemberID,
		PackageId: *cardInfo.PackageID,
	}

	if req.Status == 1 { //激活
		memberPackageCard.Status = req.Status
		if req.ActiveAt == nil {
			now := time.Now()
			req.ActiveAt = &now
		}
		if cardInfo.ValidDay == -1 {
			cardInfo.ValidDay = 36500 //100年
		}
		later := req.ActiveAt.Add(time.Hour * time.Duration(24*cardInfo.ValidDay))
		memberPackageCard.ActiveAt = req.ActiveAt
		memberPackageCard.InvalidAt = &later
		operateLog.StatusToOperate(req.Status)
		extra := make(map[string]interface{})
		extra["ActiveAt"] = memberPackageCard.ActiveAt.Format(time.DateTime)
		extra["InvalidAt"] = memberPackageCard.InvalidAt.Format(time.DateTime)
		extra["ValidDay"] = cardInfo.ValidDay
		operateLog.Extra = extra

	}
	_, err = uc.memberPackCardRepo.Update(ctx, memberPackageCard)
	if err != nil {
		log.Errorf("seatBookService Insert error:%s \r\n", err)
		return false, err
	}
	uc.memberPackCardRepo.RecordOperateLog(ctx, operateLog)
	return true, nil
}

type CardBuyLog struct {
	*gen.SrPackageCardMemberBuyLog
}

func (uc *PackageCardUsecase) FindBuyLogBy(ctx context.Context, key string, value string) (*CardBuyLog, error) {
	data, err := uc.repo.FindBuyLogBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	return &CardBuyLog{SrPackageCardMemberBuyLog: data}, nil

}
