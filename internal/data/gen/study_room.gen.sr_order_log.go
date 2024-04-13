package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrOrderLogMgr struct {
	*_BaseMgr
}

// SrOrderLogMgr open func
func SrOrderLogMgr(db *gorm.DB) *_SrOrderLogMgr {
	if db == nil {
		panic(fmt.Errorf("SrOrderLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrOrderLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_order_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrOrderLogMgr) GetTableName() string {
	return "sr_order_log"
}

// Reset 重置gorm会话
func (obj *_SrOrderLogMgr) Reset() *_SrOrderLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrOrderLogMgr) Get() (result SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrOrderLogMgr) Gets() (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrOrderLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SrOrderLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 订单id
func (obj *_SrOrderLogMgr) WithOrderID(orderID int64) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithOrderNo order_no获取 订单编号
func (obj *_SrOrderLogMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithStatus status获取 状态
func (obj *_SrOrderLogMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithNote note获取 其他记录
func (obj *_SrOrderLogMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

// WithCreatedAt created_at获取
func (obj *_SrOrderLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_SrOrderLogMgr) GetByOption(opts ...Option) (result SrOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrOrderLogMgr) GetByOptions(opts ...Option) (results []*SrOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_SrOrderLogMgr) GetFromID(id int) (result SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SrOrderLogMgr) GetBatchFromID(ids []int) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单id
func (obj *_SrOrderLogMgr) GetFromOrderID(orderID int64) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 订单id
func (obj *_SrOrderLogMgr) GetBatchFromOrderID(orderIDs []int64) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单编号
func (obj *_SrOrderLogMgr) GetFromOrderNo(orderNo string) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`order_no` = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量查找 订单编号
func (obj *_SrOrderLogMgr) GetBatchFromOrderNo(orderNos []string) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SrOrderLogMgr) GetFromStatus(status int8) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SrOrderLogMgr) GetBatchFromStatus(statuss []int8) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromNote 通过note获取内容 其他记录
func (obj *_SrOrderLogMgr) GetFromNote(note string) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`note` = ?", note).Find(&results).Error

	return
}

// GetBatchFromNote 批量查找 其他记录
func (obj *_SrOrderLogMgr) GetBatchFromNote(notes []string) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`note` IN (?)", notes).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrOrderLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrOrderLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrOrderLogMgr) FetchByPrimaryKey(id int) (result SrOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
