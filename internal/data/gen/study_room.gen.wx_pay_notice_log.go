package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxPayNoticeLogMgr struct {
	*_BaseMgr
}

// WxPayNoticeLogMgr open func
func WxPayNoticeLogMgr(db *gorm.DB) *_WxPayNoticeLogMgr {
	if db == nil {
		panic(fmt.Errorf("WxPayNoticeLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxPayNoticeLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_pay_notice_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxPayNoticeLogMgr) GetTableName() string {
	return "wx_pay_notice_log"
}

// Reset 重置gorm会话
func (obj *_WxPayNoticeLogMgr) Reset() *_WxPayNoticeLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxPayNoticeLogMgr) Get() (result WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxPayNoticeLogMgr) Gets() (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxPayNoticeLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxPayNoticeLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 类型
func (obj *_WxPayNoticeLogMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithOrderNo order_no获取 订单编号
func (obj *_WxPayNoticeLogMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithTransactionID transaction_id获取 微信支付订单号
func (obj *_WxPayNoticeLogMgr) WithTransactionID(transactionID string) Option {
	return optionFunc(func(o *options) { o.query["transaction_id"] = transactionID })
}

// WithStatus status获取 状态
func (obj *_WxPayNoticeLogMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithHeader header获取 header数据
func (obj *_WxPayNoticeLogMgr) WithHeader(header string) Option {
	return optionFunc(func(o *options) { o.query["header"] = header })
}

// WithBody body获取 body数据
func (obj *_WxPayNoticeLogMgr) WithBody(body string) Option {
	return optionFunc(func(o *options) { o.query["body"] = body })
}

// WithCreatedAt created_at获取
func (obj *_WxPayNoticeLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_WxPayNoticeLogMgr) GetByOption(opts ...Option) (result WxPayNoticeLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxPayNoticeLogMgr) GetByOptions(opts ...Option) (results []*WxPayNoticeLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_WxPayNoticeLogMgr) GetFromID(id int) (result WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxPayNoticeLogMgr) GetBatchFromID(ids []int) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型
func (obj *_WxPayNoticeLogMgr) GetFromType(_type string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型
func (obj *_WxPayNoticeLogMgr) GetBatchFromType(_types []string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单编号
func (obj *_WxPayNoticeLogMgr) GetFromOrderNo(orderNo string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`order_no` = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量查找 订单编号
func (obj *_WxPayNoticeLogMgr) GetBatchFromOrderNo(orderNos []string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromTransactionID 通过transaction_id获取内容 微信支付订单号
func (obj *_WxPayNoticeLogMgr) GetFromTransactionID(transactionID string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`transaction_id` = ?", transactionID).Find(&results).Error

	return
}

// GetBatchFromTransactionID 批量查找 微信支付订单号
func (obj *_WxPayNoticeLogMgr) GetBatchFromTransactionID(transactionIDs []string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`transaction_id` IN (?)", transactionIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_WxPayNoticeLogMgr) GetFromStatus(status string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_WxPayNoticeLogMgr) GetBatchFromStatus(statuss []string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromHeader 通过header获取内容 header数据
func (obj *_WxPayNoticeLogMgr) GetFromHeader(header string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`header` = ?", header).Find(&results).Error

	return
}

// GetBatchFromHeader 批量查找 header数据
func (obj *_WxPayNoticeLogMgr) GetBatchFromHeader(headers []string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`header` IN (?)", headers).Find(&results).Error

	return
}

// GetFromBody 通过body获取内容 body数据
func (obj *_WxPayNoticeLogMgr) GetFromBody(body string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`body` = ?", body).Find(&results).Error

	return
}

// GetBatchFromBody 批量查找 body数据
func (obj *_WxPayNoticeLogMgr) GetBatchFromBody(bodys []string) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`body` IN (?)", bodys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_WxPayNoticeLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_WxPayNoticeLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxPayNoticeLogMgr) FetchByPrimaryKey(id int) (result WxPayNoticeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxPayNoticeLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
