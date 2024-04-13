package oplatform

import (
	"errors"
	"studyRoomGo/common/wechat/oplatform/token"

	"github.com/l306287405/wechat3rd"
)

const ( //微信服务商api配置
	componentAppId     = ""
	componentAppSecret = ""
	componentAESKey    = ""
	componentToken     = ""
)

// Auth 登录/用户信息
type Auth struct {
	Server *wechat3rd.Server
}

// NewAuth new auth
func NewAuth() *Auth {
	server, _ := wechat3rd.NewService(wechat3rd.Config{
		AppID:     componentAppId,
		AppSecret: componentAppSecret,
		AESKey:    componentAESKey,
		Token:     componentToken,
	}, nil, &token.AccessTokenServer{}, wechat3rd.DefaultErrorHandler)
	return &Auth{Server: server}
}

// Code2Session 登录凭证校验。
func (auth *Auth) Code2Session(AppId string, jsCode string) (result *wechat3rd.Jscode2sessionResp, err error) {
	resp := auth.Server.Jscode2session(AppId, jsCode)
	if resp.ErrCode != 0 {
		return nil, errors.New(resp.Error.ErrMsg)
	}
	return resp, nil
}
