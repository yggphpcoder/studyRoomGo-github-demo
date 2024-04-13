package service

import (
	"context"
	"strconv"
	"strings"
	pb "studyRoomGo/api/member/v1"
	"studyRoomGo/common/ys7"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/service/dto"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
)

type MemberService struct {
	pb.UnimplementedMemberServer

	uc   *biz.MemberUsecase
	eqUc *biz.EquipmentUsecase
}

func NewMemberService(uc *biz.MemberUsecase, eqUc *biz.EquipmentUsecase) *MemberService {
	return &MemberService{uc: uc, eqUc: eqUc}
}

func (s *MemberService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	return &pb.CreateReply{}, nil
}
func (s *MemberService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.DataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	memberDTO := dto.Member{}
	if strings.Contains(req.NickName, "微信用户") {
		return &pb.DataReply{
			Code: 400,
			Msg:  "请不要用【微信用户】作为昵称",
		}, nil
	}
	data, err := s.uc.Update(ctx, &biz.MemberUpdateRequest{
		Id:         memberId,
		AvatarPath: req.AvatarPath,
		NickName:   req.NickName,
		//Phone:      req.Phone,
	})
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	memberDTO.Member = data
	return &pb.DataReply{
		Code: 200,
		Data: memberDTO.Reply(),
	}, nil
}
func (s *MemberService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return &pb.DeleteReply{}, nil
}

func (s *MemberService) Get(ctx context.Context, req *pb.GetRequest) (*pb.DataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	data, err := s.uc.FindById(ctx, int64(memberId))
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	hadNotRead, _ := s.uc.HasNoReadNotice(ctx, int64(memberId))
	d := dto.Member{Member: data, HasNoReadNotice: hadNotRead}
	return &pb.DataReply{
		Code: 200,
		Data: d.Reply(),
	}, nil
}

func (s *MemberService) MyPackageCard(ctx context.Context, req *pb.GetRequest) (*pb.MyPackageCardDataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	memberDTO := dto.Member{}

	data, err := s.uc.MyPackageCard(ctx, memberId, req.Id)
	if err != nil {
		return &pb.MyPackageCardDataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	memberDTO.MyPackageCard = data
	return &pb.MyPackageCardDataReply{
		Code: 200,
		Data: memberDTO.MyPackageCardReply(),
	}, nil
}

func (s *MemberService) MyPackageCardList(ctx context.Context, req *pb.ListRequest) (*pb.MyPackageCardDataListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)

	memberDTO := dto.Member{}

	var status []int32
	if req.Status == 1 {
		status = []int32{0, 1}
	}

	data, err := s.uc.MyPackageCardList(ctx, &biz.SearchMemberPackageCardRequest{
		MemberId: int64(memberId),
		Status:   status,
	})
	if err != nil {
		return &pb.MyPackageCardDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	var list []*pb.MyPackageCardData
	for _, item := range data {
		memberDTO.MyPackageCard = item
		list = append(list, memberDTO.MyPackageCardReply())
	}
	return &pb.MyPackageCardDataListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *MemberService) MySeatBookList(ctx context.Context, req *pb.MySeatBookListRequest) (*pb.MySeatBookDataListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	memberDTO := dto.Member{}

	var status []int32
	if req.Status == 1 {
		status = []int32{1, 2, 3, 4} //待开始,未入座,已入座,待明天开始
	}
	data, err := s.uc.MySeatBookList(ctx, &biz.SearchSeatBookingRequest{
		MemberId: memberId,
		Status:   status,
		Sort:     int(req.Sort),
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	})
	if err != nil {
		return &pb.MySeatBookDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	var list []*pb.MySeatBookData
	for _, item := range data {
		memberDTO.MySeatBook = item
		list = append(list, memberDTO.MySeatBookReply())
	}
	return &pb.MySeatBookDataListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *MemberService) MySeatLockList(ctx context.Context, req *pb.MySeatLockRequest) (*pb.MySeatLockDataListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	memberDTO := dto.Member{}

	data, err := s.uc.MySeatLockList(ctx, &biz.SearchSeatLockingRequest{
		MemberId: memberId,
		Status:   []int32{req.Status},
		Join:     true,
	})
	if err != nil {
		return &pb.MySeatLockDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	var list []*pb.MySeatLockData
	for _, item := range data {
		memberDTO.MySeatLock = item
		list = append(list, memberDTO.MySeatLockReply())
	}
	return &pb.MySeatLockDataListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *MemberService) MySeatBook(ctx context.Context, req *pb.MySeatBookRequest) (*pb.MySeatBookDataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	data, err := s.uc.MySeatBook(ctx, req.Id)
	if err != nil {
		return &pb.MySeatBookDataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	if data.MemberID != memberId {
		return &pb.MySeatBookDataReply{
			Code: 400,
			Msg:  "认证失败",
		}, nil
	}
	memberDTO := dto.Member{MySeatBook: data}

	return &pb.MySeatBookDataReply{
		Code: 200,
		Data: memberDTO.MySeatBookReply(),
	}, nil
}
func (s *MemberService) MySeatBookShared(ctx context.Context, req *pb.MySeatBookRequest) (*pb.MySeatBookDataReply, error) {
	_, err := s.uc.MySeatBookSharedValidate(ctx, &biz.MySeatBookSharedValidateReq{
		Id:       req.Id,
		MemberId: req.MemberId,
		Token:    req.Token,
		OpenId:   req.OpenId,
	})
	if err != nil {
		return &pb.MySeatBookDataReply{
			Code: err.GetCode(),
			Msg:  err.GetReason(),
		}, nil
	}
	book, _ := s.uc.MySeatBookShared(ctx, req.Id)

	isMe := false
	if req.MemberId == book.MemberID {
		isMe = true //是否本人访问
	}
	count, err2 := s.uc.MySeatBookSharedLogCount(ctx, req, isMe)
	if err2 != nil {
		return &pb.MySeatBookDataReply{
			Code: 400,
			Msg:  err2.Error(),
		}, nil
	}

	book.SharedCount = count
	memberDTO := dto.Member{MySeatBook: book}
	return &pb.MySeatBookDataReply{
		Code: 200,
		Data: memberDTO.MySeatBookReply(),
	}, nil
}

func (s *MemberService) MySeatBookLiveShared(ctx context.Context, req *pb.MySeatBookRequest) (*pb.MySeatBookDataReply, error) {
	_, err := s.uc.MySeatBookSharedValidate(ctx, &biz.MySeatBookSharedValidateReq{
		Id:       req.Id,
		MemberId: req.MemberId,
		Token:    req.Token,
		OpenId:   req.OpenId,
	})
	if err != nil {
		return &pb.MySeatBookDataReply{
			Code: err.GetCode(),
			Msg:  err.GetReason(),
		}, nil
	}
	book, _ := s.uc.MySeatBookShared(ctx, req.Id)

	isMe := false
	if req.MemberId == book.MemberID {
		isMe = true //是否本人访问
	}
	count, err2 := s.uc.MySeatBookSharedLogCount(ctx, req, isMe)
	if err2 != nil {
		return &pb.MySeatBookDataReply{
			Code: 400,
			Msg:  err2.Error(),
		}, nil
	}
	eq, _ := s.eqUc.FindByRelateId(ctx, &biz.EquipmentRequest{
		TypeRelate:    2,
		TypeEquipment: 3, //监控
		RelateId:      book.ShopAreaID,
	})
	if eq.ID != 0 {
		ys7Device := ys7.Device{}
		endTime, _ := time.ParseInLocation(time.DateTime, book.BookingEndDay.Format(time.DateOnly)+" "+book.BookingEnd, time.Local)
		expireTime := endTime.Unix() - time.Now().Unix()
		live := ys7Device.LiveAddress(eq.Mac, eq.Code, strconv.FormatInt(expireTime, 10))
		book.SharedCount = count
		book.LiveAddress = live.Data.HdAddress

	}
	memberDTO := dto.Member{MySeatBook: book}
	return &pb.MySeatBookDataReply{
		Code: 200,
		Data: memberDTO.MySeatBookReply(),
	}, nil
}

func (s *MemberService) MyCurrentUsePackageCard(ctx context.Context, req *pb.MyCurrentPackageCardRequest) (*pb.MyPackageCardDataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))
	memberDTO := dto.Member{}

	data, err := s.uc.MyCurrentUsePackageCard(ctx, memberId)
	if err != nil {
		return &pb.MyPackageCardDataReply{
			Code: 404,
			Msg:  err.Error(),
		}, nil
	}

	memberDTO.MyPackageCard = data
	return &pb.MyPackageCardDataReply{
		Code: 200,
		Data: memberDTO.MyPackageCardReply(),
	}, nil
}
func (s *MemberService) UploadAvatar(ctx context.Context, req *pb.UploadAvatarRequest) (*pb.AvatarReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := int64(jwtCtx.(jwtV4.MapClaims)["memberId"].(float64))

	data, err := s.uc.UploadImg(ctx, memberId, req.AvatarBase64)
	if err != nil {
		return &pb.AvatarReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.AvatarReply{
		Code: 200,
		Msg:  "",
		Data: &pb.AvatarReply_Data{AvatarPath: data},
	}, nil
}
func (s *MemberService) MyNotice(ctx context.Context, req *pb.MyNoticeRequest) (*pb.MyNoticeDataListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	data, err := s.uc.ListNotice(ctx, int64(memberId), req.Page, req.PageSize)
	if err != nil {
		return &pb.MyNoticeDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d := dto.Notice{}
	var list []*pb.MyNoticeData
	for _, item := range data {
		d.Notice = item
		list = append(list, d.Reply())
	}
	return &pb.MyNoticeDataListReply{
		Code: 200,
		Data: list,
	}, nil
}
func (s *MemberService) LastNotice(ctx context.Context, req *pb.MyNoticeRequest) (*pb.MyNoticeDataListReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	data, err := s.uc.LastNotice(ctx, int64(memberId), 3)
	if err != nil {
		return &pb.MyNoticeDataListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	d := dto.Notice{}
	var list []*pb.MyNoticeData
	for _, item := range data {
		d.Notice = item
		list = append(list, d.Reply())
	}
	return &pb.MyNoticeDataListReply{
		Code: 200,
		Data: list,
	}, nil
}

func (s *MemberService) ReadNotice(ctx context.Context, req *pb.MyNoticeRequest) (*pb.MyNoticeDataReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	_ = jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	_, err := s.uc.ReadNotice(ctx, req.Id)
	if err != nil {
		return &pb.MyNoticeDataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.MyNoticeDataReply{
		Code: 200,
	}, nil
}

func (s *MemberService) WxSubscribe(ctx context.Context, req *pb.WxSubscribeRequest) (*pb.WxSubscribeReply, error) {
	jwtCtx, _ := jwt.FromContext(ctx)
	memberId := jwtCtx.(jwtV4.MapClaims)["memberId"].(float64)
	for i, t := range req.Type {
		_, _ = s.uc.SaveWXSubscribe(ctx, int64(memberId), t, int(req.IsSubscribe[i]))
	}
	return &pb.WxSubscribeReply{
		Code: 200,
	}, nil
}
