package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdTeacherInfoKeyMgr struct {
	*_BaseMgr
}

// EdTeacherInfoKeyMgr open func
func EdTeacherInfoKeyMgr(db *gorm.DB) *_EdTeacherInfoKeyMgr {
	if db == nil {
		panic(fmt.Errorf("EdTeacherInfoKeyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdTeacherInfoKeyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_teacher_info_key"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdTeacherInfoKeyMgr) GetTableName() string {
	return "ed_teacher_info_key"
}

// Reset 重置gorm会话
func (obj *_EdTeacherInfoKeyMgr) Reset() *_EdTeacherInfoKeyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdTeacherInfoKeyMgr) Get() (result EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdTeacherInfoKeyMgr) Gets() (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdTeacherInfoKeyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdTeacherInfoKeyMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdTeacherInfoKeyMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdTeacherInfoKeyMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithKey key获取 字段
func (obj *_EdTeacherInfoKeyMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithValue value获取 字段名
func (obj *_EdTeacherInfoKeyMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithHTMLType html_type获取 字段类型
func (obj *_EdTeacherInfoKeyMgr) WithHTMLType(htmlType string) Option {
	return optionFunc(func(o *options) { o.query["html_type"] = htmlType })
}

// WithDictType dict_type获取 字典
func (obj *_EdTeacherInfoKeyMgr) WithDictType(dictType string) Option {
	return optionFunc(func(o *options) { o.query["dict_type"] = dictType })
}

// WithCreateBy create_by获取
func (obj *_EdTeacherInfoKeyMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdTeacherInfoKeyMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithCreatedAt created_at获取
func (obj *_EdTeacherInfoKeyMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdTeacherInfoKeyMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdTeacherInfoKeyMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_EdTeacherInfoKeyMgr) GetByOption(opts ...Option) (result EdTeacherInfoKey, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdTeacherInfoKeyMgr) GetByOptions(opts ...Option) (results []*EdTeacherInfoKey, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdTeacherInfoKeyMgr) GetFromID(id int64) (result EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromID(ids []int64) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdTeacherInfoKeyMgr) GetFromMerchantID(merchantID int64) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdTeacherInfoKeyMgr) GetFromShopID(shopID int64) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容 字段
func (obj *_EdTeacherInfoKeyMgr) GetFromKey(key string) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`key` = ?", key).Find(&results).Error

	return
}

// GetBatchFromKey 批量查找 字段
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromKey(keys []string) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`key` IN (?)", keys).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 字段名
func (obj *_EdTeacherInfoKeyMgr) GetFromValue(value string) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找 字段名
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromValue(values []string) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromHTMLType 通过html_type获取内容 字段类型
func (obj *_EdTeacherInfoKeyMgr) GetFromHTMLType(htmlType string) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`html_type` = ?", htmlType).Find(&results).Error

	return
}

// GetBatchFromHTMLType 批量查找 字段类型
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromHTMLType(htmlTypes []string) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`html_type` IN (?)", htmlTypes).Find(&results).Error

	return
}

// GetFromDictType 通过dict_type获取内容 字典
func (obj *_EdTeacherInfoKeyMgr) GetFromDictType(dictType string) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`dict_type` = ?", dictType).Find(&results).Error

	return
}

// GetBatchFromDictType 批量查找 字典
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromDictType(dictTypes []string) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`dict_type` IN (?)", dictTypes).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdTeacherInfoKeyMgr) GetFromCreateBy(createBy *uint) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdTeacherInfoKeyMgr) GetFromUpdateBy(updateBy *uint) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdTeacherInfoKeyMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdTeacherInfoKeyMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdTeacherInfoKeyMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdTeacherInfoKeyMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdTeacherInfoKeyMgr) FetchByPrimaryKey(id int64) (result EdTeacherInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfoKey{}).Where("`id` = ?", id).First(&result).Error

	return
}
