package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DianpingAuthMgr struct {
	*_BaseMgr
}

// DianpingAuthMgr open func
func DianpingAuthMgr(db *gorm.DB) *_DianpingAuthMgr {
	if db == nil {
		panic(fmt.Errorf("DianpingAuthMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DianpingAuthMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dianping_auth"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DianpingAuthMgr) GetTableName() string {
	return "dianping_auth"
}

// Reset 重置gorm会话
func (obj *_DianpingAuthMgr) Reset() *_DianpingAuthMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DianpingAuthMgr) Get() (result DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DianpingAuthMgr) Gets() (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DianpingAuthMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DianpingAuthMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAppKey app_key获取 北极星开放平台分配给应用的AppKey
func (obj *_DianpingAuthMgr) WithAppKey(appKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = appKey })
}

// WithAppSecret app_secret获取 密钥
func (obj *_DianpingAuthMgr) WithAppSecret(appSecret string) Option {
	return optionFunc(func(o *options) { o.query["app_secret"] = appSecret })
}

// WithAuthCode auth_code获取 授权码
func (obj *_DianpingAuthMgr) WithAuthCode(authCode string) Option {
	return optionFunc(func(o *options) { o.query["auth_code"] = authCode })
}

// WithAccessToken access_token获取 session
func (obj *_DianpingAuthMgr) WithAccessToken(accessToken string) Option {
	return optionFunc(func(o *options) { o.query["access_token"] = accessToken })
}

// WithExpiresIn expires_in获取 过期时间
func (obj *_DianpingAuthMgr) WithExpiresIn(expiresIn string) Option {
	return optionFunc(func(o *options) { o.query["expires_in"] = expiresIn })
}

// WithRefreshToken refresh_token获取 refresh_session
func (obj *_DianpingAuthMgr) WithRefreshToken(refreshToken string) Option {
	return optionFunc(func(o *options) { o.query["refresh_token"] = refreshToken })
}

// WithRemainRefreshCount remain_refresh_count获取 剩余刷新次数
func (obj *_DianpingAuthMgr) WithRemainRefreshCount(remainRefreshCount string) Option {
	return optionFunc(func(o *options) { o.query["remain_refresh_count"] = remainRefreshCount })
}

// WithBid bid获取 客户Id
func (obj *_DianpingAuthMgr) WithBid(bid string) Option {
	return optionFunc(func(o *options) { o.query["bid"] = bid })
}

// GetByOption 功能选项模式获取
func (obj *_DianpingAuthMgr) GetByOption(opts ...Option) (result DianpingAuth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DianpingAuthMgr) GetByOptions(opts ...Option) (results []*DianpingAuth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DianpingAuthMgr) GetFromID(id int) (result DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DianpingAuthMgr) GetBatchFromID(ids []int) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAppKey 通过app_key获取内容 北极星开放平台分配给应用的AppKey
func (obj *_DianpingAuthMgr) GetFromAppKey(appKey string) (result DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`app_key` = ?", appKey).First(&result).Error

	return
}

// GetBatchFromAppKey 批量查找 北极星开放平台分配给应用的AppKey
func (obj *_DianpingAuthMgr) GetBatchFromAppKey(appKeys []string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`app_key` IN (?)", appKeys).Find(&results).Error

	return
}

// GetFromAppSecret 通过app_secret获取内容 密钥
func (obj *_DianpingAuthMgr) GetFromAppSecret(appSecret string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`app_secret` = ?", appSecret).Find(&results).Error

	return
}

// GetBatchFromAppSecret 批量查找 密钥
func (obj *_DianpingAuthMgr) GetBatchFromAppSecret(appSecrets []string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`app_secret` IN (?)", appSecrets).Find(&results).Error

	return
}

// GetFromAuthCode 通过auth_code获取内容 授权码
func (obj *_DianpingAuthMgr) GetFromAuthCode(authCode string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`auth_code` = ?", authCode).Find(&results).Error

	return
}

// GetBatchFromAuthCode 批量查找 授权码
func (obj *_DianpingAuthMgr) GetBatchFromAuthCode(authCodes []string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`auth_code` IN (?)", authCodes).Find(&results).Error

	return
}

// GetFromAccessToken 通过access_token获取内容 session
func (obj *_DianpingAuthMgr) GetFromAccessToken(accessToken string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`access_token` = ?", accessToken).Find(&results).Error

	return
}

// GetBatchFromAccessToken 批量查找 session
func (obj *_DianpingAuthMgr) GetBatchFromAccessToken(accessTokens []string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`access_token` IN (?)", accessTokens).Find(&results).Error

	return
}

// GetFromExpiresIn 通过expires_in获取内容 过期时间
func (obj *_DianpingAuthMgr) GetFromExpiresIn(expiresIn string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`expires_in` = ?", expiresIn).Find(&results).Error

	return
}

// GetBatchFromExpiresIn 批量查找 过期时间
func (obj *_DianpingAuthMgr) GetBatchFromExpiresIn(expiresIns []string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`expires_in` IN (?)", expiresIns).Find(&results).Error

	return
}

// GetFromRefreshToken 通过refresh_token获取内容 refresh_session
func (obj *_DianpingAuthMgr) GetFromRefreshToken(refreshToken string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`refresh_token` = ?", refreshToken).Find(&results).Error

	return
}

// GetBatchFromRefreshToken 批量查找 refresh_session
func (obj *_DianpingAuthMgr) GetBatchFromRefreshToken(refreshTokens []string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`refresh_token` IN (?)", refreshTokens).Find(&results).Error

	return
}

// GetFromRemainRefreshCount 通过remain_refresh_count获取内容 剩余刷新次数
func (obj *_DianpingAuthMgr) GetFromRemainRefreshCount(remainRefreshCount string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`remain_refresh_count` = ?", remainRefreshCount).Find(&results).Error

	return
}

// GetBatchFromRemainRefreshCount 批量查找 剩余刷新次数
func (obj *_DianpingAuthMgr) GetBatchFromRemainRefreshCount(remainRefreshCounts []string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`remain_refresh_count` IN (?)", remainRefreshCounts).Find(&results).Error

	return
}

// GetFromBid 通过bid获取内容 客户Id
func (obj *_DianpingAuthMgr) GetFromBid(bid string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`bid` = ?", bid).Find(&results).Error

	return
}

// GetBatchFromBid 批量查找 客户Id
func (obj *_DianpingAuthMgr) GetBatchFromBid(bids []string) (results []*DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`bid` IN (?)", bids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DianpingAuthMgr) FetchByPrimaryKey(id int) (result DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByDianpingAuthAppKeyUIndex primary or index 获取唯一内容
func (obj *_DianpingAuthMgr) FetchUniqueByDianpingAuthAppKeyUIndex(appKey string) (result DianpingAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingAuth{}).Where("`app_key` = ?", appKey).First(&result).Error

	return
}
