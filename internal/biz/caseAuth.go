package biz

import (
	"context"
	"errors"
	"fmt"
	_func "studyRoomGo/common/func"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/data/gen"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// AuthUsecase is an Auth usecase.
type AuthUsecase struct {
	repo MemberRepo
	log  *log.Helper
}

type WxUserInfo struct {
	*gen.WxUserinfo
}

// NewAuthUsecase new an Auth usecase.
func NewAuthUsecase(repo MemberRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger)}
}

// WxLogin 微信初始化登录
func (uc *AuthUsecase) WxLogin(ctx context.Context, merchantId int64, wx *wechat.OauthResp) error {
	tmp, err := uc.repo.FindByWXOpenId(ctx, merchantId, wx.OpenID)
	if err != nil {
		return err
	}

	if tmp.ID == 0 {
		_, err := uc.repo.CreateWXUserInfo(ctx, &WxUserInfo{
			WxUserinfo: &gen.WxUserinfo{
				Openid:     wx.OpenID,
				MerchantID: merchantId,
			},
		})
		tmp.Openid = wx.OpenID
		tmp.MerchantID = merchantId
		if err != nil {
			return err
		}
	}
	tmp.UnionID = wx.UnionID
	uc.repo.UpdateWXUserInfo(ctx, &WxUserInfo{
		WxUserinfo: tmp,
	})
	return nil
}

// LoginByWeChat  微信openid登录
func (uc *AuthUsecase) LoginByWeChat(ctx context.Context, merchantId int64, openId string, phone string) (*Member, error) {
	tmp, _ := uc.repo.FindByWXOpenId(ctx, merchantId, openId)
	if tmp.ID == 0 {
		return nil, errors.New("openid not exist")
	}
	memberPhone, _ := uc.repo.FindByPhone(ctx, merchantId, phone)
	if memberPhone.ID == 0 {
		return nil, nil //手机号没有注册，立即注册
	}
	if memberPhone.ID != 0 { //手机号已存在
		if tmp.MemberID != memberPhone.ID { //且当前微信号!=手机号的微信号
			tmp.Phone = memberPhone.Phone
			tmp.MemberID = memberPhone.ID
			_, _ = uc.repo.UpdateWXUserInfo(ctx, &WxUserInfo{tmp}) //更新当前微信号的账号
		}
	}
	//else { //手机号不存在
	//	if tmp.MemberID != 0 { //	且但当前微信号已经绑定过会员
	//		tmp.Phone = &phone
	//		_, _ = uc.repo.UpdateWXUserInfo(ctx, &WxUserInfo{tmp}) //更新当前微信号的绑定手机
	//		member, _ := uc.repo.FindByID(ctx, tmp.MemberID)
	//		member.Phone = &phone
	//		_, _ = uc.repo.Update(ctx, &Member{
	//			SrMember: member,
	//		}) //更新当前微信号的绑定手机
	//	}
	//}

	member, err := uc.repo.FindByID(ctx, tmp.MemberID)
	if err != nil {
		return nil, err
	}
	return &Member{SrMember: member}, nil

}
func (uc *AuthUsecase) LoginByWeChatUnionId(ctx context.Context, merchantId int64, unionId string) (*Member, error) {
	tmp, _ := uc.repo.FindByWXUnionId(ctx, merchantId, unionId)
	if tmp.ID == 0 {
		return nil, errors.New("unionId not exist")
	}
	member, err := uc.repo.FindByID(ctx, tmp.MemberID)
	if err != nil {
		return nil, err
	}
	return &Member{SrMember: member}, nil
}

// RegisterByWeChat  微信openid注册
func (uc *AuthUsecase) RegisterByWeChat(ctx context.Context, merchantId int64, openId string, phone string) (*Member, error) {
	tmp, _ := uc.repo.FindByWXOpenId(ctx, merchantId, openId)
	if tmp.ID == 0 {
		return nil, errors.New("openid not exist")
	}

	nickName := fmt.Sprintf("%v%v", "微信用户_", _func.RandAllString(6, _func.CHARS))

	now := time.Now()
	//第一次登录创建用户
	avatarPath := "avatar/default.png"
	member := &gen.SrMember{
		MerchantID: merchantId,
		Username:   openId,
		NickName:   &nickName,
		Phone:      &phone,
		AvatarPath: &avatarPath,
		RegisterAt: &now,
	}
	_, err := uc.repo.Save(ctx, &Member{
		SrMember: member,
	})
	if err != nil {
		return nil, err
	}
	//绑定用户id 与 微信openid
	tmp.MemberID = member.ID
	tmp.Phone = &phone
	_, err = uc.repo.UpdateWXUserInfo(ctx, &WxUserInfo{tmp})
	if err != nil {
		return nil, err
	}
	return &Member{SrMember: member}, nil
}
