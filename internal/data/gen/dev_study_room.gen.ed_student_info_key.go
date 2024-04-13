package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdStudentInfoKeyMgr struct {
	*_BaseMgr
}

// EdStudentInfoKeyMgr open func
func EdStudentInfoKeyMgr(db *gorm.DB) *_EdStudentInfoKeyMgr {
	if db == nil {
		panic(fmt.Errorf("EdStudentInfoKeyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdStudentInfoKeyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_student_info_key"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdStudentInfoKeyMgr) GetTableName() string {
	return "ed_student_info_key"
}

// Reset 重置gorm会话
func (obj *_EdStudentInfoKeyMgr) Reset() *_EdStudentInfoKeyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdStudentInfoKeyMgr) Get() (result EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdStudentInfoKeyMgr) Gets() (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdStudentInfoKeyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdStudentInfoKeyMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdStudentInfoKeyMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdStudentInfoKeyMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithKey key获取 字段
func (obj *_EdStudentInfoKeyMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithValue value获取 字段名
func (obj *_EdStudentInfoKeyMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithHTMLType html_type获取 字段类型
func (obj *_EdStudentInfoKeyMgr) WithHTMLType(htmlType string) Option {
	return optionFunc(func(o *options) { o.query["html_type"] = htmlType })
}

// WithDictType dict_type获取 字典
func (obj *_EdStudentInfoKeyMgr) WithDictType(dictType string) Option {
	return optionFunc(func(o *options) { o.query["dict_type"] = dictType })
}

// WithCreateBy create_by获取
func (obj *_EdStudentInfoKeyMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdStudentInfoKeyMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithCreatedAt created_at获取
func (obj *_EdStudentInfoKeyMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdStudentInfoKeyMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdStudentInfoKeyMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_EdStudentInfoKeyMgr) GetByOption(opts ...Option) (result EdStudentInfoKey, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdStudentInfoKeyMgr) GetByOptions(opts ...Option) (results []*EdStudentInfoKey, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdStudentInfoKeyMgr) GetFromID(id int64) (result EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdStudentInfoKeyMgr) GetBatchFromID(ids []int64) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdStudentInfoKeyMgr) GetFromMerchantID(merchantID int64) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdStudentInfoKeyMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdStudentInfoKeyMgr) GetFromShopID(shopID int64) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdStudentInfoKeyMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容 字段
func (obj *_EdStudentInfoKeyMgr) GetFromKey(key string) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`key` = ?", key).Find(&results).Error

	return
}

// GetBatchFromKey 批量查找 字段
func (obj *_EdStudentInfoKeyMgr) GetBatchFromKey(keys []string) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`key` IN (?)", keys).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 字段名
func (obj *_EdStudentInfoKeyMgr) GetFromValue(value string) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找 字段名
func (obj *_EdStudentInfoKeyMgr) GetBatchFromValue(values []string) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromHTMLType 通过html_type获取内容 字段类型
func (obj *_EdStudentInfoKeyMgr) GetFromHTMLType(htmlType string) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`html_type` = ?", htmlType).Find(&results).Error

	return
}

// GetBatchFromHTMLType 批量查找 字段类型
func (obj *_EdStudentInfoKeyMgr) GetBatchFromHTMLType(htmlTypes []string) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`html_type` IN (?)", htmlTypes).Find(&results).Error

	return
}

// GetFromDictType 通过dict_type获取内容 字典
func (obj *_EdStudentInfoKeyMgr) GetFromDictType(dictType string) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`dict_type` = ?", dictType).Find(&results).Error

	return
}

// GetBatchFromDictType 批量查找 字典
func (obj *_EdStudentInfoKeyMgr) GetBatchFromDictType(dictTypes []string) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`dict_type` IN (?)", dictTypes).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdStudentInfoKeyMgr) GetFromCreateBy(createBy *uint) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdStudentInfoKeyMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdStudentInfoKeyMgr) GetFromUpdateBy(updateBy *uint) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdStudentInfoKeyMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdStudentInfoKeyMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdStudentInfoKeyMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdStudentInfoKeyMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdStudentInfoKeyMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdStudentInfoKeyMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdStudentInfoKeyMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdStudentInfoKeyMgr) FetchByPrimaryKey(id int64) (result EdStudentInfoKey, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoKey{}).Where("`id` = ?", id).First(&result).Error

	return
}
