package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SrOrderTotalMgr struct {
	*_BaseMgr
}

// SrOrderTotalMgr open func
func SrOrderTotalMgr(db *gorm.DB) *_SrOrderTotalMgr {
	if db == nil {
		panic(fmt.Errorf("SrOrderTotalMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrOrderTotalMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_order_total"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrOrderTotalMgr) GetTableName() string {
	return "sr_order_total"
}

// Reset 重置gorm会话
func (obj *_SrOrderTotalMgr) Reset() *_SrOrderTotalMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrOrderTotalMgr) Get() (result SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrOrderTotalMgr) Gets() (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrOrderTotalMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_SrOrderTotalMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 订单id
func (obj *_SrOrderTotalMgr) WithOrderID(orderID int64) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithType type获取 金额类型
func (obj *_SrOrderTotalMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithValue value获取 金额值
func (obj *_SrOrderTotalMgr) WithValue(value float64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithRelateType relate_type获取 关联类型
func (obj *_SrOrderTotalMgr) WithRelateType(relateType int8) Option {
	return optionFunc(func(o *options) { o.query["relate_type"] = relateType })
}

// WithRelateValue relate_value获取 关联值
func (obj *_SrOrderTotalMgr) WithRelateValue(relateValue string) Option {
	return optionFunc(func(o *options) { o.query["relate_value"] = relateValue })
}

// GetByOption 功能选项模式获取
func (obj *_SrOrderTotalMgr) GetByOption(opts ...Option) (result SrOrderTotal, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrOrderTotalMgr) GetByOptions(opts ...Option) (results []*SrOrderTotal, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 id
func (obj *_SrOrderTotalMgr) GetFromID(id int64) (result SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_SrOrderTotalMgr) GetBatchFromID(ids []int64) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单id
func (obj *_SrOrderTotalMgr) GetFromOrderID(orderID int64) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 订单id
func (obj *_SrOrderTotalMgr) GetBatchFromOrderID(orderIDs []int64) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 金额类型
func (obj *_SrOrderTotalMgr) GetFromType(_type string) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 金额类型
func (obj *_SrOrderTotalMgr) GetBatchFromType(_types []string) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 金额值
func (obj *_SrOrderTotalMgr) GetFromValue(value float64) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找 金额值
func (obj *_SrOrderTotalMgr) GetBatchFromValue(values []float64) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromRelateType 通过relate_type获取内容 关联类型
func (obj *_SrOrderTotalMgr) GetFromRelateType(relateType int8) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`relate_type` = ?", relateType).Find(&results).Error

	return
}

// GetBatchFromRelateType 批量查找 关联类型
func (obj *_SrOrderTotalMgr) GetBatchFromRelateType(relateTypes []int8) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`relate_type` IN (?)", relateTypes).Find(&results).Error

	return
}

// GetFromRelateValue 通过relate_value获取内容 关联值
func (obj *_SrOrderTotalMgr) GetFromRelateValue(relateValue string) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`relate_value` = ?", relateValue).Find(&results).Error

	return
}

// GetBatchFromRelateValue 批量查找 关联值
func (obj *_SrOrderTotalMgr) GetBatchFromRelateValue(relateValues []string) (results []*SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`relate_value` IN (?)", relateValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrOrderTotalMgr) FetchByPrimaryKey(id int64) (result SrOrderTotal, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrderTotal{}).Where("`id` = ?", id).First(&result).Error

	return
}
