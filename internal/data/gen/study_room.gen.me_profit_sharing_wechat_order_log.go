package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MeProfitSharingWechatOrderLogMgr struct {
	*_BaseMgr
}

// MeProfitSharingWechatOrderLogMgr open func
func MeProfitSharingWechatOrderLogMgr(db *gorm.DB) *_MeProfitSharingWechatOrderLogMgr {
	if db == nil {
		panic(fmt.Errorf("MeProfitSharingWechatOrderLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MeProfitSharingWechatOrderLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("me_profit_sharing_wechat_order_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MeProfitSharingWechatOrderLogMgr) GetTableName() string {
	return "me_profit_sharing_wechat_order_log"
}

// Reset 重置gorm会话
func (obj *_MeProfitSharingWechatOrderLogMgr) Reset() *_MeProfitSharingWechatOrderLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MeProfitSharingWechatOrderLogMgr) Get() (result MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MeProfitSharingWechatOrderLogMgr) Gets() (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MeProfitSharingWechatOrderLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MeProfitSharingWechatOrderLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithProfitSharingID profit_sharing_id获取 分账单id
func (obj *_MeProfitSharingWechatOrderLogMgr) WithProfitSharingID(profitSharingID int64) Option {
	return optionFunc(func(o *options) { o.query["profit_sharing_id"] = profitSharingID })
}

// WithOrderNo order_no获取 分账单订单编号
func (obj *_MeProfitSharingWechatOrderLogMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithStatus status获取 状态
func (obj *_MeProfitSharingWechatOrderLogMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBody body获取 响应内容
func (obj *_MeProfitSharingWechatOrderLogMgr) WithBody(body *string) Option {
	return optionFunc(func(o *options) { o.query["body"] = body })
}

// WithHeader header获取 响应
func (obj *_MeProfitSharingWechatOrderLogMgr) WithHeader(header *string) Option {
	return optionFunc(func(o *options) { o.query["header"] = header })
}

// WithCreatedAt created_at获取
func (obj *_MeProfitSharingWechatOrderLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_MeProfitSharingWechatOrderLogMgr) GetByOption(opts ...Option) (result MeProfitSharingWechatOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MeProfitSharingWechatOrderLogMgr) GetByOptions(opts ...Option) (results []*MeProfitSharingWechatOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_MeProfitSharingWechatOrderLogMgr) GetFromID(id int) (result MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MeProfitSharingWechatOrderLogMgr) GetBatchFromID(ids []int) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromProfitSharingID 通过profit_sharing_id获取内容 分账单id
func (obj *_MeProfitSharingWechatOrderLogMgr) GetFromProfitSharingID(profitSharingID int64) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`profit_sharing_id` = ?", profitSharingID).Find(&results).Error

	return
}

// GetBatchFromProfitSharingID 批量查找 分账单id
func (obj *_MeProfitSharingWechatOrderLogMgr) GetBatchFromProfitSharingID(profitSharingIDs []int64) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`profit_sharing_id` IN (?)", profitSharingIDs).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 分账单订单编号
func (obj *_MeProfitSharingWechatOrderLogMgr) GetFromOrderNo(orderNo string) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`order_no` = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量查找 分账单订单编号
func (obj *_MeProfitSharingWechatOrderLogMgr) GetBatchFromOrderNo(orderNos []string) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_MeProfitSharingWechatOrderLogMgr) GetFromStatus(status int8) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_MeProfitSharingWechatOrderLogMgr) GetBatchFromStatus(statuss []int8) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBody 通过body获取内容 响应内容
func (obj *_MeProfitSharingWechatOrderLogMgr) GetFromBody(body *string) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`body` = ?", body).Find(&results).Error

	return
}

// GetBatchFromBody 批量查找 响应内容
func (obj *_MeProfitSharingWechatOrderLogMgr) GetBatchFromBody(bodys []*string) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`body` IN (?)", bodys).Find(&results).Error

	return
}

// GetFromHeader 通过header获取内容 响应
func (obj *_MeProfitSharingWechatOrderLogMgr) GetFromHeader(header *string) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`header` = ?", header).Find(&results).Error

	return
}

// GetBatchFromHeader 批量查找 响应
func (obj *_MeProfitSharingWechatOrderLogMgr) GetBatchFromHeader(headers []*string) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`header` IN (?)", headers).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_MeProfitSharingWechatOrderLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_MeProfitSharingWechatOrderLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MeProfitSharingWechatOrderLogMgr) FetchByPrimaryKey(id int) (result MeProfitSharingWechatOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
