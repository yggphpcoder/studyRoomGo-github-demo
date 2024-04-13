package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WxNoticeSubscribeMgr struct {
	*_BaseMgr
}

// WxNoticeSubscribeMgr open func
func WxNoticeSubscribeMgr(db *gorm.DB) *_WxNoticeSubscribeMgr {
	if db == nil {
		panic(fmt.Errorf("WxNoticeSubscribeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxNoticeSubscribeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_notice_subscribe"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxNoticeSubscribeMgr) GetTableName() string {
	return "wx_notice_subscribe"
}

// Reset 重置gorm会话
func (obj *_WxNoticeSubscribeMgr) Reset() *_WxNoticeSubscribeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxNoticeSubscribeMgr) Get() (result WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxNoticeSubscribeMgr) Gets() (results []*WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxNoticeSubscribeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMemberID member_id获取 会员id
func (obj *_WxNoticeSubscribeMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithType type获取 订阅类型
func (obj *_WxNoticeSubscribeMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithIsSubscribe is_subscribe获取 是否订阅
func (obj *_WxNoticeSubscribeMgr) WithIsSubscribe(isSubscribe int) Option {
	return optionFunc(func(o *options) { o.query["is_subscribe"] = isSubscribe })
}

// GetByOption 功能选项模式获取
func (obj *_WxNoticeSubscribeMgr) GetByOption(opts ...Option) (result WxNoticeSubscribe, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxNoticeSubscribeMgr) GetByOptions(opts ...Option) (results []*WxNoticeSubscribe, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_WxNoticeSubscribeMgr) GetFromMemberID(memberID int64) (results []*WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_WxNoticeSubscribeMgr) GetBatchFromMemberID(memberIDs []int64) (results []*WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 订阅类型
func (obj *_WxNoticeSubscribeMgr) GetFromType(_type string) (results []*WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 订阅类型
func (obj *_WxNoticeSubscribeMgr) GetBatchFromType(_types []string) (results []*WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromIsSubscribe 通过is_subscribe获取内容 是否订阅
func (obj *_WxNoticeSubscribeMgr) GetFromIsSubscribe(isSubscribe int) (results []*WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where("`is_subscribe` = ?", isSubscribe).Find(&results).Error

	return
}

// GetBatchFromIsSubscribe 批量查找 是否订阅
func (obj *_WxNoticeSubscribeMgr) GetBatchFromIsSubscribe(isSubscribes []int) (results []*WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where("`is_subscribe` IN (?)", isSubscribes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueIndexByWxNoticeSubscribePk primary or index 获取唯一内容
func (obj *_WxNoticeSubscribeMgr) FetchUniqueIndexByWxNoticeSubscribePk(memberID int64, _type string) (result WxNoticeSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxNoticeSubscribe{}).Where("`member_id` = ? AND `type` = ?", memberID, _type).First(&result).Error

	return
}
