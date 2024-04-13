package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrSettingMgr struct {
	*_BaseMgr
}

// SrSettingMgr open func
func SrSettingMgr(db *gorm.DB) *_SrSettingMgr {
	if db == nil {
		panic(fmt.Errorf("SrSettingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrSettingMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_setting"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrSettingMgr) GetTableName() string {
	return "sr_setting"
}

// Reset 重置gorm会话
func (obj *_SrSettingMgr) Reset() *_SrSettingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrSettingMgr) Get() (result SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrSettingMgr) Gets() (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrSettingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrSettingMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 配置名
func (obj *_SrSettingMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithKey key获取 配置键
func (obj *_SrSettingMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithValueType value_type获取 值类型
func (obj *_SrSettingMgr) WithValueType(valueType int) Option {
	return optionFunc(func(o *options) { o.query["value_type"] = valueType })
}

// WithValue value获取 配置值
func (obj *_SrSettingMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithGroup group获取 配置组
func (obj *_SrSettingMgr) WithGroup(group string) Option {
	return optionFunc(func(o *options) { o.query["group"] = group })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_SrSettingMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrSettingMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithSort sort获取 排序
func (obj *_SrSettingMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrSettingMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrSettingMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrSettingMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrSettingMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrSettingMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrSettingMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrSettingMgr) GetByOption(opts ...Option) (result SrSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrSettingMgr) GetByOptions(opts ...Option) (results []*SrSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrSettingMgr) GetFromID(id int64) (result SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrSettingMgr) GetBatchFromID(ids []int64) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 配置名
func (obj *_SrSettingMgr) GetFromName(name string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 配置名
func (obj *_SrSettingMgr) GetBatchFromName(names []string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容 配置键
func (obj *_SrSettingMgr) GetFromKey(key string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`key` = ?", key).Find(&results).Error

	return
}

// GetBatchFromKey 批量查找 配置键
func (obj *_SrSettingMgr) GetBatchFromKey(keys []string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`key` IN (?)", keys).Find(&results).Error

	return
}

// GetFromValueType 通过value_type获取内容 值类型
func (obj *_SrSettingMgr) GetFromValueType(valueType int) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`value_type` = ?", valueType).Find(&results).Error

	return
}

// GetBatchFromValueType 批量查找 值类型
func (obj *_SrSettingMgr) GetBatchFromValueType(valueTypes []int) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`value_type` IN (?)", valueTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 配置值
func (obj *_SrSettingMgr) GetFromValue(value string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找 配置值
func (obj *_SrSettingMgr) GetBatchFromValue(values []string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromGroup 通过group获取内容 配置组
func (obj *_SrSettingMgr) GetFromGroup(group string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`group` = ?", group).Find(&results).Error

	return
}

// GetBatchFromGroup 批量查找 配置组
func (obj *_SrSettingMgr) GetBatchFromGroup(groups []string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`group` IN (?)", groups).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_SrSettingMgr) GetFromMerchantID(merchantID int64) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_SrSettingMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrSettingMgr) GetFromShopID(shopID int64) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrSettingMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrSettingMgr) GetFromSort(sort int) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrSettingMgr) GetBatchFromSort(sorts []int) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrSettingMgr) GetFromRemark(remark *string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrSettingMgr) GetBatchFromRemark(remarks []*string) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrSettingMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrSettingMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrSettingMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrSettingMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrSettingMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrSettingMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrSettingMgr) GetFromCreateBy(createBy *uint) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrSettingMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrSettingMgr) GetFromUpdateBy(updateBy *uint) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrSettingMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrSettingMgr) FetchByPrimaryKey(id int64) (result SrSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}
