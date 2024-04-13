package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdDictDataMgr struct {
	*_BaseMgr
}

// EdDictDataMgr open func
func EdDictDataMgr(db *gorm.DB) *_EdDictDataMgr {
	if db == nil {
		panic(fmt.Errorf("EdDictDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdDictDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_dict_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdDictDataMgr) GetTableName() string {
	return "ed_dict_data"
}

// Reset 重置gorm会话
func (obj *_EdDictDataMgr) Reset() *_EdDictDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdDictDataMgr) Get() (result EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdDictDataMgr) Gets() (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdDictDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDictCode dict_code获取
func (obj *_EdDictDataMgr) WithDictCode(dictCode int64) Option {
	return optionFunc(func(o *options) { o.query["dict_code"] = dictCode })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdDictDataMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdDictDataMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithDictSort dict_sort获取
func (obj *_EdDictDataMgr) WithDictSort(dictSort *int64) Option {
	return optionFunc(func(o *options) { o.query["dict_sort"] = dictSort })
}

// WithDictLabel dict_label获取
func (obj *_EdDictDataMgr) WithDictLabel(dictLabel *string) Option {
	return optionFunc(func(o *options) { o.query["dict_label"] = dictLabel })
}

// WithDictValue dict_value获取
func (obj *_EdDictDataMgr) WithDictValue(dictValue *string) Option {
	return optionFunc(func(o *options) { o.query["dict_value"] = dictValue })
}

// WithDictType dict_type获取
func (obj *_EdDictDataMgr) WithDictType(dictType *string) Option {
	return optionFunc(func(o *options) { o.query["dict_type"] = dictType })
}

// WithCSSClass css_class获取
func (obj *_EdDictDataMgr) WithCSSClass(cssClass *string) Option {
	return optionFunc(func(o *options) { o.query["css_class"] = cssClass })
}

// WithListClass list_class获取
func (obj *_EdDictDataMgr) WithListClass(listClass *string) Option {
	return optionFunc(func(o *options) { o.query["list_class"] = listClass })
}

// WithIsDefault is_default获取
func (obj *_EdDictDataMgr) WithIsDefault(isDefault *string) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// WithStatus status获取
func (obj *_EdDictDataMgr) WithStatus(status *int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDefault default获取
func (obj *_EdDictDataMgr) WithDefault(_default *string) Option {
	return optionFunc(func(o *options) { o.query["default"] = _default })
}

// WithRemark remark获取
func (obj *_EdDictDataMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateBy create_by获取 创建者
func (obj *_EdDictDataMgr) WithCreateBy(createBy *int64) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取 更新者
func (obj *_EdDictDataMgr) WithUpdateBy(updateBy *int64) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_EdDictDataMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 最后更新时间
func (obj *_EdDictDataMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdDictDataMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_EdDictDataMgr) GetByOption(opts ...Option) (result EdDictData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdDictDataMgr) GetByOptions(opts ...Option) (results []*EdDictData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromDictCode 通过dict_code获取内容
func (obj *_EdDictDataMgr) GetFromDictCode(dictCode int64) (result EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_code` = ?", dictCode).First(&result).Error

	return
}

// GetBatchFromDictCode 批量查找
func (obj *_EdDictDataMgr) GetBatchFromDictCode(dictCodes []int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_code` IN (?)", dictCodes).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdDictDataMgr) GetFromMerchantID(merchantID int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdDictDataMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdDictDataMgr) GetFromShopID(shopID int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdDictDataMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromDictSort 通过dict_sort获取内容
func (obj *_EdDictDataMgr) GetFromDictSort(dictSort *int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_sort` = ?", dictSort).Find(&results).Error

	return
}

// GetBatchFromDictSort 批量查找
func (obj *_EdDictDataMgr) GetBatchFromDictSort(dictSorts []*int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_sort` IN (?)", dictSorts).Find(&results).Error

	return
}

// GetFromDictLabel 通过dict_label获取内容
func (obj *_EdDictDataMgr) GetFromDictLabel(dictLabel *string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_label` = ?", dictLabel).Find(&results).Error

	return
}

// GetBatchFromDictLabel 批量查找
func (obj *_EdDictDataMgr) GetBatchFromDictLabel(dictLabels []*string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_label` IN (?)", dictLabels).Find(&results).Error

	return
}

// GetFromDictValue 通过dict_value获取内容
func (obj *_EdDictDataMgr) GetFromDictValue(dictValue *string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_value` = ?", dictValue).Find(&results).Error

	return
}

// GetBatchFromDictValue 批量查找
func (obj *_EdDictDataMgr) GetBatchFromDictValue(dictValues []*string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_value` IN (?)", dictValues).Find(&results).Error

	return
}

// GetFromDictType 通过dict_type获取内容
func (obj *_EdDictDataMgr) GetFromDictType(dictType *string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_type` = ?", dictType).Find(&results).Error

	return
}

// GetBatchFromDictType 批量查找
func (obj *_EdDictDataMgr) GetBatchFromDictType(dictTypes []*string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_type` IN (?)", dictTypes).Find(&results).Error

	return
}

// GetFromCSSClass 通过css_class获取内容
func (obj *_EdDictDataMgr) GetFromCSSClass(cssClass *string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`css_class` = ?", cssClass).Find(&results).Error

	return
}

// GetBatchFromCSSClass 批量查找
func (obj *_EdDictDataMgr) GetBatchFromCSSClass(cssClasss []*string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`css_class` IN (?)", cssClasss).Find(&results).Error

	return
}

// GetFromListClass 通过list_class获取内容
func (obj *_EdDictDataMgr) GetFromListClass(listClass *string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`list_class` = ?", listClass).Find(&results).Error

	return
}

// GetBatchFromListClass 批量查找
func (obj *_EdDictDataMgr) GetBatchFromListClass(listClasss []*string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`list_class` IN (?)", listClasss).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容
func (obj *_EdDictDataMgr) GetFromIsDefault(isDefault *string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找
func (obj *_EdDictDataMgr) GetBatchFromIsDefault(isDefaults []*string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_EdDictDataMgr) GetFromStatus(status *int8) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_EdDictDataMgr) GetBatchFromStatus(statuss []*int8) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDefault 通过default获取内容
func (obj *_EdDictDataMgr) GetFromDefault(_default *string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`default` = ?", _default).Find(&results).Error

	return
}

// GetBatchFromDefault 批量查找
func (obj *_EdDictDataMgr) GetBatchFromDefault(_defaults []*string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`default` IN (?)", _defaults).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容
func (obj *_EdDictDataMgr) GetFromRemark(remark *string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找
func (obj *_EdDictDataMgr) GetBatchFromRemark(remarks []*string) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_EdDictDataMgr) GetFromCreateBy(createBy *int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找 创建者
func (obj *_EdDictDataMgr) GetBatchFromCreateBy(createBys []*int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_EdDictDataMgr) GetFromUpdateBy(updateBy *int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找 更新者
func (obj *_EdDictDataMgr) GetBatchFromUpdateBy(updateBys []*int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_EdDictDataMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_EdDictDataMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 最后更新时间
func (obj *_EdDictDataMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 最后更新时间
func (obj *_EdDictDataMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdDictDataMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdDictDataMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdDictDataMgr) FetchByPrimaryKey(dictCode int64) (result EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`dict_code` = ?", dictCode).First(&result).Error

	return
}

// FetchIndexByIDxSysDictDataCreateBy  获取多个内容
func (obj *_EdDictDataMgr) FetchIndexByIDxSysDictDataCreateBy(createBy *int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// FetchIndexByIDxSysDictDataUpdateBy  获取多个内容
func (obj *_EdDictDataMgr) FetchIndexByIDxSysDictDataUpdateBy(updateBy *int64) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// FetchIndexByIDxSysDictDataDeletedAt  获取多个内容
func (obj *_EdDictDataMgr) FetchIndexByIDxSysDictDataDeletedAt(deletedAt *time.Time) (results []*EdDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictData{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
