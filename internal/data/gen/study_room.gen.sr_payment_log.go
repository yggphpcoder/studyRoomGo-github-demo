package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrPaymentLogMgr struct {
	*_BaseMgr
}

// SrPaymentLogMgr open func
func SrPaymentLogMgr(db *gorm.DB) *_SrPaymentLogMgr {
	if db == nil {
		panic(fmt.Errorf("SrPaymentLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrPaymentLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_payment_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrPaymentLogMgr) GetTableName() string {
	return "sr_payment_log"
}

// Reset 重置gorm会话
func (obj *_SrPaymentLogMgr) Reset() *_SrPaymentLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrPaymentLogMgr) Get() (result SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrPaymentLogMgr) Gets() (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrPaymentLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SrPaymentLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPaymentID payment_id获取 支付表id
func (obj *_SrPaymentLogMgr) WithPaymentID(paymentID int64) Option {
	return optionFunc(func(o *options) { o.query["payment_id"] = paymentID })
}

// WithOrderNo order_no获取 订单编号
func (obj *_SrPaymentLogMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithTransactionID transaction_id获取 微信支付订单号
func (obj *_SrPaymentLogMgr) WithTransactionID(transactionID *string) Option {
	return optionFunc(func(o *options) { o.query["transaction_id"] = transactionID })
}

// WithStatus status获取 状态
func (obj *_SrPaymentLogMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithNote note获取 其他记录
func (obj *_SrPaymentLogMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

// WithCreatedAt created_at获取
func (obj *_SrPaymentLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_SrPaymentLogMgr) GetByOption(opts ...Option) (result SrPaymentLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrPaymentLogMgr) GetByOptions(opts ...Option) (results []*SrPaymentLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_SrPaymentLogMgr) GetFromID(id int) (result SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SrPaymentLogMgr) GetBatchFromID(ids []int) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPaymentID 通过payment_id获取内容 支付表id
func (obj *_SrPaymentLogMgr) GetFromPaymentID(paymentID int64) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`payment_id` = ?", paymentID).Find(&results).Error

	return
}

// GetBatchFromPaymentID 批量查找 支付表id
func (obj *_SrPaymentLogMgr) GetBatchFromPaymentID(paymentIDs []int64) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`payment_id` IN (?)", paymentIDs).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单编号
func (obj *_SrPaymentLogMgr) GetFromOrderNo(orderNo string) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`order_no` = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量查找 订单编号
func (obj *_SrPaymentLogMgr) GetBatchFromOrderNo(orderNos []string) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromTransactionID 通过transaction_id获取内容 微信支付订单号
func (obj *_SrPaymentLogMgr) GetFromTransactionID(transactionID *string) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`transaction_id` = ?", transactionID).Find(&results).Error

	return
}

// GetBatchFromTransactionID 批量查找 微信支付订单号
func (obj *_SrPaymentLogMgr) GetBatchFromTransactionID(transactionIDs []*string) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`transaction_id` IN (?)", transactionIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SrPaymentLogMgr) GetFromStatus(status int8) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SrPaymentLogMgr) GetBatchFromStatus(statuss []int8) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromNote 通过note获取内容 其他记录
func (obj *_SrPaymentLogMgr) GetFromNote(note string) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`note` = ?", note).Find(&results).Error

	return
}

// GetBatchFromNote 批量查找 其他记录
func (obj *_SrPaymentLogMgr) GetBatchFromNote(notes []string) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`note` IN (?)", notes).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrPaymentLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrPaymentLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrPaymentLogMgr) FetchByPrimaryKey(id int) (result SrPaymentLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
