package biz

import (
	"bufio"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	pb "studyRoomGo/api/member/v1"
	"studyRoomGo/common/config"
	_func "studyRoomGo/common/func"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
	"time"

	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// MemberUsecase is a Member usecase.
type MemberUsecase struct {
	repo               MemberRepo
	memberPackCardRepo MemberPackageCardRepo
	seatBookRepo       SeatBookRepo
	packageCardRepo    PackageCardRepo
	shopRepo           ShopRepo
	seatRepo           SeatRepo

	log *log.Helper
}

// Member is a Member model.
type Member struct {
	*gen.SrMember
	NeedUpdate bool
}

type MemberUpdateRequest struct {
	Id         int64
	AvatarPath string
	NickName   string
	Phone      string
}

// NewMemberUsecase new a Member usecase.
func NewMemberUsecase(repo MemberRepo, memberPackCardRepo MemberPackageCardRepo, seatBookRepo SeatBookRepo, packageCardRepo PackageCardRepo, shopRepo ShopRepo, seatRepo SeatRepo, logger log.Logger) *MemberUsecase {
	return &MemberUsecase{
		repo:               repo,
		memberPackCardRepo: memberPackCardRepo,
		seatBookRepo:       seatBookRepo,
		packageCardRepo:    packageCardRepo,
		shopRepo:           shopRepo,
		seatRepo:           seatRepo,
		log:                log.NewHelper(logger),
	}
}

// Create  creates a Member, and returns the new Member.
func (uc *MemberUsecase) Create(ctx context.Context, g *Member) (*Member, error) {
	uc.log.WithContext(ctx).Infof("CreateMember: %v", g.Username)
	_, err := uc.repo.Save(ctx, g)
	if err != nil {
		log.Errorf("seatBookService Insert error:%s \r\n", err)
		return g, err
	}
	return g, nil
}

func (uc *MemberUsecase) Update(ctx context.Context, g *MemberUpdateRequest) (*Member, error) {
	data, err := uc.repo.FindByID(ctx, g.Id)
	if err != nil {
		return nil, err
	}
	data.AvatarPath = &g.AvatarPath
	data.NickName = &g.NickName

	m := &Member{SrMember: data}
	ok, err := uc.repo.Update(ctx, m)
	if !ok {
		return nil, err
	}
	if data.NickName == nil || strings.Contains(*data.NickName, "微信用户") {
		m.NeedUpdate = true
	}
	return m, nil
}
func (uc *MemberUsecase) FindById(ctx context.Context, id int64) (*Member, error) {
	data, err := uc.repo.FindByID(ctx, id)
	m := &Member{SrMember: data}
	if data.NickName == nil || strings.Contains(*data.NickName, "微信用户") {
		m.NeedUpdate = true
	}
	return m, err
}
func (uc *MemberUsecase) HasNoReadNotice(ctx context.Context, memberId int64) (bool, error) {
	isRead := 0
	ok, err := uc.repo.HasNotice(ctx, memberId, &isRead)
	return ok, err
}

func (uc *MemberUsecase) LastNotice(ctx context.Context, memberId int64, limit int32) ([]*gen.SrMemberNotice, error) {
	isRead := 0
	data, err := uc.repo.LastNotice(ctx, memberId, &isRead, limit)
	return data, err
}

func (uc *MemberUsecase) ListNotice(ctx context.Context, memberId int64, Page int32, PageSize int32) ([]*gen.SrMemberNotice, error) {
	list, err := uc.repo.ListNotice(ctx, memberId, Page, PageSize)
	return list, err
}

func (uc *MemberUsecase) ReadNotice(ctx context.Context, id []int64) (bool, error) {
	ok, err := uc.repo.UpdateNotice(ctx, id, 1)
	return ok, err
}

type MyPackageCard struct {
	*gen.SrMemberPackageCard
	CanBookingToCloseTime bool
	CanBookingToHour      string
	UseShopStr            string
	UseSeatStr            string
	UseShopId             []int32
}

func (uc *MemberUsecase) MyPackageCard(ctx context.Context, memberId int64, id int64) (*MyPackageCard, error) {
	data, err := uc.memberPackCardRepo.FindByID(ctx, id)
	if data.MemberID != memberId {
		return nil, errors.New("认证失败")
	}

	type useTime struct {
		Week   []int8
		AllDay bool
		Hour   []string
	}
	var uTime []useTime
	var canBookingToHour string

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
	allDay := true
	err = json.Unmarshal([]byte(data.UseTime), &uTime)
	if err == nil {
		for _, u := range uTime {
			if !u.AllDay {
				allDay = false
				canBookingToHour = u.Hour[1]
			}
		}

	}
	var shopNameArr []string
	var seatNameArr []string
	var useShopId []int32
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
				useShopId = append(useShopId, int32(id))
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
	canBookingToCloseTime := false
	if allDay && data.MinutePer <= 0 {
		canBookingToCloseTime = true
	}
	return &MyPackageCard{
		SrMemberPackageCard:   data,
		CanBookingToCloseTime: canBookingToCloseTime,
		CanBookingToHour:      canBookingToHour,
		UseShopStr:            strings.Join(shopNameArr, ","),
		UseSeatStr:            strings.Join(seatNameArr, ","),
		UseShopId:             useShopId,
	}, err
}

func (uc *MemberUsecase) MyPackageCardList(ctx context.Context, req *SearchMemberPackageCardRequest) ([]*MyPackageCard, error) {
	data, err := uc.memberPackCardRepo.ListBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*MyPackageCard

	type useTime struct {
		Week   []int8
		AllDay bool
		Hour   []string
	}
	var uTime []useTime
	var canBookingToHour string

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
		allDay := true
		err := json.Unmarshal([]byte(item.UseTime), &uTime)
		if err == nil {
			for _, u := range uTime {
				if !u.AllDay {
					allDay = false
					canBookingToHour = u.Hour[1]
				}
			}

		}
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
		canBookingToCloseTime := false
		if item.MinutePer <= 0 {
			if allDay {
				canBookingToCloseTime = true //次卡订座到营业结束
			}
		} else {
			canBookingToHour = "" // 每次可订分钟不为空，则不能自动订座到可用时段结束
		}
		list = append(list, &MyPackageCard{
			SrMemberPackageCard:   item,
			CanBookingToCloseTime: canBookingToCloseTime,
			CanBookingToHour:      canBookingToHour,
			UseShopStr:            strings.Join(shopNameArr, ","),
			UseSeatStr:            strings.Join(seatNameArr, ","),
		})

	}

	return list, nil
}

type MySeatBook struct {
	*model.SeatBook
	SharedCount int64
	LiveAddress string
}

func (uc *MemberUsecase) MySeatBookList(ctx context.Context, req *SearchSeatBookingRequest) ([]*MySeatBook, error) {
	data, err := uc.seatBookRepo.ListBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*MySeatBook
	for _, item := range data {
		list = append(list, &MySeatBook{SeatBook: item})
	}
	return list, nil
}

func (uc *MemberUsecase) MySeatBook(ctx context.Context, id int64) (*MySeatBook, error) {
	data, err := uc.seatBookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &MySeatBook{SeatBook: data}, nil
}

type MySeatBookSharedValidateReq struct {
	Id       int64
	MemberId int64
	Token    string
	OpenId   string
}

func (uc *MemberUsecase) MySeatBookSharedValidate(ctx context.Context,
	req *MySeatBookSharedValidateReq) (ok bool, err *errors2.Error) {

	if req.OpenId == "" || req.Token == "" || req.Id == 0 {
		err = errors2.New(400, fmt.Sprintf("数据错误"), "")
		return false, err
	}
	str := fmt.Sprintf("%d|%s", req.Id, SharedKey)
	hash := hmac.New(md5.New, []byte(SharedKey)) //创建对应的sha256哈希加密算法
	hash.Write([]byte(str))
	token := hex.EncodeToString(hash.Sum([]byte("")))
	if token != req.Token {
		err = errors2.New(400, fmt.Sprintf("认证失败"), "")
		return false, err
	}
	book, _ := uc.MySeatBookShared(ctx, req.Id)
	if !_func.InInt8Splice([]int8{1, 2, 3, 4}, book.Status) {
		err = errors2.New(400, fmt.Sprintf("订座已失效,不允许分享"), "")
		return false, err
	}
	now := time.Now()
	endTime, _ := time.ParseInLocation(time.DateTime, book.BookingEndDay.Format(time.DateOnly)+" "+book.BookingEnd, time.Local)
	if now.After(endTime) {
		err = errors2.New(400, fmt.Sprintf("不在入座时间内"), "")
		return false, err
	}
	return ok, nil
}

func (uc *MemberUsecase) MySeatBookShared(ctx context.Context, id int64) (*MySeatBook, error) {
	data, err := uc.seatBookRepo.FindDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &MySeatBook{SeatBook: data}, nil
}

func (uc *MemberUsecase) MySeatBookSharedLogCount(ctx context.Context, req *pb.MySeatBookRequest, isMe bool) (count int64, err error) {
	data, count, err := uc.seatBookRepo.FindSharedByID(ctx, req.Id, req.OpenId)
	if err != nil && !gen.IsNotFound(err) {
		return count, err
	}
	if count < 10 || isMe {
		if data == nil || data.Openid != req.OpenId {
			if ok, _ := uc.SaveMySeatBookSharedLog(ctx, req); ok {
				count = count + 1
			}
		}
	} else {
		return count, errors.New("已超过最大分享次数（10次）")
	}

	return count, nil
}
func (uc *MemberUsecase) SaveMySeatBookSharedLog(ctx context.Context, req *pb.MySeatBookRequest) (bool bool, err error) {
	ok, err := uc.seatBookRepo.SaveSharedLog(ctx, &gen.SrSeatBookShareLog{
		SeatBookID: req.Id,
		Openid:     req.OpenId,
		MemberID:   req.MemberId,
	})
	if err != nil {
		return ok, err
	}
	return ok, err
}

func (uc *MemberUsecase) UploadImg(ctx context.Context, memberId int64, avatarBase64 string) (string, error) {
	//去除包装，获取到base64编码
	imgBase64 := avatarBase64
	//base64转码
	imgs, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", errors.New(fmt.Sprint("base64 decode error:", err))
	}

	timenow := time.Now().Unix()
	id := strconv.Itoa(int(memberId))
	file, err := os.OpenFile(config.AppConf.Upload.GetUri()+"/"+config.AppConf.Upload.Avatar.GetDir()+"/"+id+"_"+strconv.FormatInt(timenow, 10)+".jpg", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", errors.New(fmt.Sprint("create file error:", err))
	}
	w := bufio.NewWriter(file) //创建新的 Writer 对象
	_, err3 := w.WriteString(string(imgs))
	if err3 != nil {
		return "", errors.New(fmt.Sprint("write error:", err3))
	}
	w.Flush()
	defer file.Close()
	imgname := id + "_" + strconv.FormatInt(timenow, 10) + ".jpg"
	//t := struct {
	//	ImageName string `json:"imagename"`
	//}{imgname}
	//this.Data["json"] = t
	//this.ServeJSON()
	return config.AppConf.Upload.Avatar.GetDir() + "/" + imgname, nil
}

func (uc *MemberUsecase) MyCurrentUsePackageCard(ctx context.Context, memberId int64) (*MyPackageCard, error) {
	log, err := uc.memberPackCardRepo.LastUseLog(ctx, memberId)
	if err != nil {
		return nil, err
	}
	data, err := uc.memberPackCardRepo.FindByID(ctx, log.MpID)

	if err != nil {
		return nil, err
	}
	if data.MemberID != memberId {
		return nil, errors.New("认证失败")
	}

	return &MyPackageCard{SrMemberPackageCard: data}, err
}

type MySeatLock struct {
	*model.SeatLock
}

func (uc *MemberUsecase) MySeatLockList(ctx context.Context, req *SearchSeatLockingRequest) ([]*MySeatLock, error) {
	data, err := uc.seatRepo.ListLockingBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*MySeatLock
	for _, item := range data {
		list = append(list, &MySeatLock{item})
	}
	return list, nil
}

func (uc *MemberUsecase) SaveWXSubscribe(ctx context.Context, memberId int64, subType string, isSubscribe int) (bool, error) {
	ok, err := uc.repo.SaveWXSubscribe(ctx, &gen.WxNoticeSubscribe{MemberID: memberId, Type: subType, IsSubscribe: isSubscribe})
	return ok, err
}
