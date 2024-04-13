package wechat

import (
	config2 "studyRoomGo/common/config"
	"studyRoomGo/common/wechat/gen"
	merchant2 "studyRoomGo/common/wechat/merchant"
	"studyRoomGo/common/wechat/oplatform"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

//var memory *cache.Memory
//
//func init() {
//	memory = cache.NewMemory()
//}

// OauthResp 微信授权返回信息
type OauthResp struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

type Auth struct {
	AppId     string
	AppSecret string
	Version   int8
}

func NewAuth(merchantId int64) (*Auth, error) {
	merchant := merchant2.NewMerchant(merchantId)
	data, err := merchant.GetMerchantAppSecret()
	if err != nil {
		return nil, err
	}
	return &Auth{
		AppId:     data.AppID,
		AppSecret: data.AppSecret,
		Version:   data.Version,
	}, nil
}
func (a *Auth) Oauth(JsCode string) (*OauthResp, error) {
	if a.Version == 1 {
		return Oauth(a.AppId, a.AppSecret, JsCode) //定制版，密钥授权
	}
	return OplatformOauth(a.AppId, JsCode)
}

func (a *Auth) GetUserPhoneNumber(code string) (phoneNumber string, err error) {
	if a.Version == 1 {
		return GetUserPhoneNumber(a.AppId, a.AppSecret, code) //定制版，密钥授权
	}
	//return OplatformOauth(a.AppId, JsCode)
	return "", err
}

// Oauth 微信小程序登录
func Oauth(AppId string, AppSecret string, JsCode string) (*OauthResp, error) {

	wc := wechat.NewWechat()
	cfg := &miniConfig.Config{
		AppID:     AppId,
		AppSecret: AppSecret,
		Cache:     cache.NewMemory(),
	}
	mini := wc.GetMiniProgram(cfg)
	a := mini.GetAuth()
	session, err := a.Code2Session(JsCode)

	if err != nil {
		return nil, err
	}
	resp := &OauthResp{
		OpenID:     session.OpenID,
		SessionKey: session.SessionKey,
		UnionID:    session.UnionID,
	}
	return resp, nil
}

func GetUserPhoneNumber(AppId string, AppSecret string, code string) (phoneNumber string, err error) {
	wc := wechat.NewWechat()
	cfg := &miniConfig.Config{
		AppID:     AppId,
		AppSecret: AppSecret,
		Cache:     cache.NewMemory(),
	}
	mini := wc.GetMiniProgram(cfg)
	a := mini.GetAuth()
	phone, err := a.GetPhoneNumber(code)
	if err != nil {
		return "", err
	}
	return phone.PhoneInfo.PurePhoneNumber, nil
}

func OplatformOauth(AppId string, JsCode string) (*OauthResp, error) {
	auth := oplatform.NewAuth()
	resp, err := auth.Code2Session(AppId, JsCode)
	if err != nil {
		return nil, err
	}
	return &OauthResp{
		OpenID:     resp.Openid,
		SessionKey: resp.SessionKey,
		UnionID:    resp.Unionid,
	}, nil

}

type ShareRecordReq struct {
	OpenID        string `json:"openid"`    // 用户唯一标识
	MemberID      int64  `json:"member_id"` // 用户id
	ShareType     string `json:"share_type"`
	RelateId      int64  `json:"relate_id"`       // 关联id
	ShareOpenID   string `json:"share_openid"`    // 分享人openid
	ShareMemberID int64  `json:"share_member_id"` // 分享人id

}

func ShareRecord(req ShareRecordReq) error {
	d := gen.WxShareClickLogMgr(config2.AppConf.Data.DB)

	err := d.Save(&gen.WxShareClickLog{
		ShareOpenID:   req.ShareOpenID,
		ShareMemberID: req.ShareMemberID,
		ShareType:     req.ShareType,
		ShareRelateID: req.RelateId,
		OpenID:        req.OpenID,
		MemberID:      req.MemberID,
	}).Error
	return err
}
