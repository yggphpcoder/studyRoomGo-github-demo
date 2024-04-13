package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxOauthMgr struct {
	*_BaseMgr
}

// WxOauthMgr open func
func WxOauthMgr(db *gorm.DB) *_WxOauthMgr {
	if db == nil {
		panic(fmt.Errorf("WxOauthMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxOauthMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_oauth"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxOauthMgr) GetTableName() string {
	return "wx_oauth"
}

// Reset 重置gorm会话
func (obj *_WxOauthMgr) Reset() *_WxOauthMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxOauthMgr) Get() (result WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxOauthMgr) Gets() (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxOauthMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxOauthMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_WxOauthMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithOpenid openid获取 微信用户唯一标识符
//
// 微信用户唯一标识符
func (obj *_WxOauthMgr) WithOpenid(openid string) Option {
	return optionFunc(func(o *options) { o.query["openid"] = openid })
}

// WithAccessToken access_token获取
func (obj *_WxOauthMgr) WithAccessToken(accessToken string) Option {
	return optionFunc(func(o *options) { o.query["access_token"] = accessToken })
}

// WithRefreshToken refresh_token获取
func (obj *_WxOauthMgr) WithRefreshToken(refreshToken *string) Option {
	return optionFunc(func(o *options) { o.query["refresh_token"] = refreshToken })
}

// WithUnionID union_id获取
func (obj *_WxOauthMgr) WithUnionID(unionID *string) Option {
	return optionFunc(func(o *options) { o.query["union_id"] = unionID })
}

// WithExpires expires获取
func (obj *_WxOauthMgr) WithExpires(expires time.Time) Option {
	return optionFunc(func(o *options) { o.query["expires"] = expires })
}

// WithPhone phone获取 用户绑定的手机号
func (obj *_WxOauthMgr) WithPhone(phone *string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithMemberID member_id获取 会员id
func (obj *_WxOauthMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// GetByOption 功能选项模式获取
func (obj *_WxOauthMgr) GetByOption(opts ...Option) (result WxOauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxOauthMgr) GetByOptions(opts ...Option) (results []*WxOauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxOauthMgr) GetFromID(id int) (result WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxOauthMgr) GetBatchFromID(ids []int) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_WxOauthMgr) GetFromMerchantID(merchantID int64) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_WxOauthMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromOpenid 通过openid获取内容 微信用户唯一标识符
//
// 微信用户唯一标识符
func (obj *_WxOauthMgr) GetFromOpenid(openid string) (result WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`openid` = ?", openid).First(&result).Error

	return
}

// GetBatchFromOpenid 批量查找 微信用户唯一标识符
//
// 微信用户唯一标识符
func (obj *_WxOauthMgr) GetBatchFromOpenid(openids []string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`openid` IN (?)", openids).Find(&results).Error

	return
}

// GetFromAccessToken 通过access_token获取内容
func (obj *_WxOauthMgr) GetFromAccessToken(accessToken string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`access_token` = ?", accessToken).Find(&results).Error

	return
}

// GetBatchFromAccessToken 批量查找
func (obj *_WxOauthMgr) GetBatchFromAccessToken(accessTokens []string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`access_token` IN (?)", accessTokens).Find(&results).Error

	return
}

// GetFromRefreshToken 通过refresh_token获取内容
func (obj *_WxOauthMgr) GetFromRefreshToken(refreshToken *string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`refresh_token` = ?", refreshToken).Find(&results).Error

	return
}

// GetBatchFromRefreshToken 批量查找
func (obj *_WxOauthMgr) GetBatchFromRefreshToken(refreshTokens []*string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`refresh_token` IN (?)", refreshTokens).Find(&results).Error

	return
}

// GetFromUnionID 通过union_id获取内容
func (obj *_WxOauthMgr) GetFromUnionID(unionID *string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`union_id` = ?", unionID).Find(&results).Error

	return
}

// GetBatchFromUnionID 批量查找
func (obj *_WxOauthMgr) GetBatchFromUnionID(unionIDs []*string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`union_id` IN (?)", unionIDs).Find(&results).Error

	return
}

// GetFromExpires 通过expires获取内容
func (obj *_WxOauthMgr) GetFromExpires(expires time.Time) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`expires` = ?", expires).Find(&results).Error

	return
}

// GetBatchFromExpires 批量查找
func (obj *_WxOauthMgr) GetBatchFromExpires(expiress []time.Time) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`expires` IN (?)", expiress).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 用户绑定的手机号
func (obj *_WxOauthMgr) GetFromPhone(phone *string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 用户绑定的手机号
func (obj *_WxOauthMgr) GetBatchFromPhone(phones []*string) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_WxOauthMgr) GetFromMemberID(memberID int64) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_WxOauthMgr) GetBatchFromMemberID(memberIDs []int64) (results []*WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxOauthMgr) FetchByPrimaryKey(id int) (result WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByOpenid primary or index 获取唯一内容
func (obj *_WxOauthMgr) FetchUniqueByOpenid(openid string) (result WxOauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxOauth{}).Where("`openid` = ?", openid).First(&result).Error

	return
}
