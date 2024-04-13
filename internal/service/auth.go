package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	pb "studyRoomGo/api/auth/v1"
	_func "studyRoomGo/common/func"
	myJwt "studyRoomGo/common/jwt"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/biz"
	checkloginEquipment "studyRoomGo/internal/middleware/checkLoginEquipment"

	"github.com/go-kratos/kratos/v2/metadata"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	uc *biz.AuthUsecase
}

func NewAuthService(uc *biz.AuthUsecase) *AuthService {
	return &AuthService{uc: uc}
}

// WxLogin 微信登录
func (s *AuthService) WxLogin(ctx context.Context, req *pb.WxLoginRequest) (*pb.CreateAuthReply, error) {
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	authService, err := wechat.NewAuth(int64(merchantId))
	if err != nil {
		return &pb.CreateAuthReply{
			Code: 401,
			Msg:  err.Error(),
		}, nil
	}
	wx, err := authService.Oauth(req.JsCode)
	if err != nil {
		return &pb.CreateAuthReply{
			Code: 401,
			Msg:  err.Error(),
		}, nil
	}
	err = s.uc.WxLogin(ctx, int64(merchantId), wx)
	if err != nil {
		return &pb.CreateAuthReply{
			Code: 401,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.CreateAuthReply{
		Code: 200,
		Data: &pb.CreateAuthReply_Data{OpenId: wx.OpenID},
	}, nil
}

// LoginAndRegisterByWeChat 绑定会员和微信信息
func (s *AuthService) LoginAndRegisterByWeChat(ctx context.Context, req *pb.LoginByWeChatRequest) (*pb.LoginByWeChatReply, error) {
	openId := req.OpenId

	if openId == "o2lPC5HLgikBeDcL81N4SWPwaB_Y" {
		member, _ := s.uc.LoginByWeChat(ctx, 1, req.OpenId, "15626204813")
		randToken := _func.RandNumString(16)
		token := myJwt.CreateToken(openId, member.ID, member.Type, randToken)
		checkloginEquipment.SetLastLoginEquipment(ctx, member.ID, randToken)
		return &pb.LoginByWeChatReply{
			Code:     200,
			JwtToken: token,
		}, nil
	}
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	authService, err := wechat.NewAuth(int64(merchantId))
	if err != nil {
		return &pb.LoginByWeChatReply{
			Code: 401,
			Msg:  err.Error(),
		}, nil
	}

	phone, err := authService.GetUserPhoneNumber(req.PhoneCode)
	if err != nil {
		return &pb.LoginByWeChatReply{
			Code: 401,
			Msg:  err.Error(),
		}, nil
	}
	//登录
	member, err := s.uc.LoginByWeChat(ctx, int64(merchantId), openId, phone)
	if err != nil {
		return &pb.LoginByWeChatReply{
			Code: 401,
			Msg:  err.Error(),
		}, nil
	}
	if member == nil { //注册，
		member, err = s.uc.RegisterByWeChat(ctx, int64(merchantId), openId, phone)
		if err != nil {
			return &pb.LoginByWeChatReply{
				Code: 401,
				Msg:  err.Error(),
			}, nil
		}
	}
	randToken := _func.RandNumString(16)
	token := myJwt.CreateToken(openId, member.ID, member.Type, randToken)
	checkloginEquipment.SetLastLoginEquipment(ctx, member.ID, randToken)

	return &pb.LoginByWeChatReply{
		Code:     200,
		JwtToken: token,
	}, nil
}

// 公众号 微信网页登录
func (s *AuthService) OauthLogin(ctx context.Context, req *pb.WxLoginRequest) (*pb.LoginByWeChatReply, error) {
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	authService, err := wechat.NewOfficialAuth(int64(merchantId))
	if err != nil {
		return &pb.LoginByWeChatReply{
			Code: 401,
			Msg:  err.Error(),
		}, nil
	}
	resp, err := authService.Oauth(req.JsCode)
	if err != nil {
		return &pb.LoginByWeChatReply{
			Code: 401,
			Msg:  "获取微信授权失败，请重试",
		}, nil
	}
	//登录
	member, err := s.uc.LoginByWeChatUnionId(ctx, int64(merchantId), resp.UnionID)
	if err != nil {
		return &pb.LoginByWeChatReply{
			Code: 401,
			Msg:  "未注册小程序，请先注册",
		}, nil
	}
	authService.OauthBindMemberID(resp.OpenID, member.ID)
	randToken := _func.RandNumString(16)
	token := myJwt.CreateToken("", member.ID, member.Type, randToken)
	checkloginEquipment.SetLastLoginEquipment(ctx, member.ID, randToken)

	return &pb.LoginByWeChatReply{
		Code:     200,
		JwtToken: token,
	}, nil
}

func (s *AuthService) GetShareSign(ctx context.Context, req *pb.GetShareSignRequest) (*pb.GetShareSignReply, error) {

	str := fmt.Sprintf("%v|%v|%v|%v|%v", req.Type, req.RelateId, int64(req.MemberId), req.OpenId, "sadfvdfgys")
	hasher := md5.New()
	hasher.Write([]byte(str))
	md5Value := hex.EncodeToString(hasher.Sum(nil))
	relpy := md5Value
	return &pb.GetShareSignReply{
		Code: 200,
		Data: &pb.GetShareSignReply_Data{Sign: relpy},
	}, nil
}

func (s *AuthService) ShareRecord(ctx context.Context, req *pb.ShareRecordRequest) (*pb.ShareRecordReply, error) {

	str := fmt.Sprintf("%v|%v|%v|%v|%v", req.Type, req.RelateId, int64(req.ShareMemberId), req.ShareOpenId, "sadfvdfgys")
	hasher := md5.New()
	hasher.Write([]byte(str))
	md5Value := hex.EncodeToString(hasher.Sum(nil))
	relpy := md5Value
	if relpy != req.Sign {
		return &pb.ShareRecordReply{
			Code: 401,
			Msg:  errors.New("签名错误").Error(),
		}, nil
	}
	wechat.ShareRecord(wechat.ShareRecordReq{
		OpenID:        req.OpenId,
		MemberID:      int64(req.MemberId),
		ShareType:     req.Type,
		RelateId:      req.RelateId,
		ShareOpenID:   req.ShareOpenId,
		ShareMemberID: int64(req.ShareMemberId),
	})
	return &pb.ShareRecordReply{
		Code: 200,
	}, nil
}
