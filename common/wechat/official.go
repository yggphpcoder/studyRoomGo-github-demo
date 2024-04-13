package wechat

import (
	"fmt"
	"strconv"
	config2 "studyRoomGo/common/config"
	merchant2 "studyRoomGo/common/wechat/merchant"
	"studyRoomGo/internal/data/gen"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
)

const (
	serviceAppId     = "" //微信服务号信息
	serviceAppSecret = ""
	serviceAESKey    = ""
	serviceToken     = ""
)

// 服务号
type OfficialAuth struct {
	MerchantId int64
	AppId      string
	AppSecret  string
	Version    int8
}

func NewOfficialAuth(merchantId int64) (*OfficialAuth, error) {
	merchant := merchant2.NewMerchant(merchantId)
	data, err := merchant.GetMerchantAppSecret()
	if err != nil {
		return nil, err
	}
	return &OfficialAuth{
		MerchantId: merchantId,
		AppId:      data.AppID,
		AppSecret:  data.AppSecret,
		Version:    data.Version,
	}, nil
}
func (a *OfficialAuth) Oauth(JsCode string) (*oauth.ResAccessToken, error) {

	wc := wechat.NewWechat()
	cfg := &offConfig.Config{
		AppID:          serviceAppId,
		AppSecret:      serviceAppSecret,
		Token:          serviceToken,
		EncodingAESKey: serviceAESKey,
		Cache:          cache.NewMemory(),
	}
	mini := wc.GetOfficialAccount(cfg)
	o := mini.GetOauth()
	resp, err := o.GetUserAccessToken(JsCode)
	fmt.Println(cfg, JsCode, resp)
	if err != nil {
		return nil, err
	}
	d := gen.WxOauthMgr(config2.AppConf.Data.DB)

	d.Save(&gen.WxOauth{
		MerchantID:   a.MerchantId,
		Openid:       resp.OpenID,
		UnionID:      resp.UnionID,
		AccessToken:  resp.AccessToken,
		Expires:      strconv.FormatInt(resp.ExpiresIn, 10),
		RefreshToken: resp.RefreshToken,
	})
	return &resp, nil

}
func (a *OfficialAuth) OauthBindMemberID(openid string, memberId int64) (bool, error) {
	d := gen.WxOauthMgr(config2.AppConf.Data.DB)

	err := d.Where("openid = ?", openid).Update("member_id", memberId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
