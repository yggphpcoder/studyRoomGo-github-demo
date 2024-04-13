package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SysDictDataMgr struct {
	*_BaseMgr
}

// SysDictDataMgr open func
func SysDictDataMgr(db *gorm.DB) *_SysDictDataMgr {
	if db == nil {
		panic(fmt.Errorf("SysDictDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysDictDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_dict_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysDictDataMgr) GetTableName() string {
	return "sys_dict_data"
}

// Reset 重置gorm会话
func (obj *_SysDictDataMgr) Reset() *_SysDictDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SysDictDataMgr) Get() (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysDictDataMgr) Gets() (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SysDictDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDictCode dict_code获取
func (obj *_SysDictDataMgr) WithDictCode(dictCode int64) Option {
	return optionFunc(func(o *options) { o.query["dict_code"] = dictCode })
}

// WithDictSort dict_sort获取
func (obj *_SysDictDataMgr) WithDictSort(dictSort *int64) Option {
	return optionFunc(func(o *options) { o.query["dict_sort"] = dictSort })
}

// WithDictLabel dict_label获取
func (obj *_SysDictDataMgr) WithDictLabel(dictLabel *string) Option {
	return optionFunc(func(o *options) { o.query["dict_label"] = dictLabel })
}

// WithDictValue dict_value获取
func (obj *_SysDictDataMgr) WithDictValue(dictValue *string) Option {
	return optionFunc(func(o *options) { o.query["dict_value"] = dictValue })
}

// WithDictType dict_type获取
func (obj *_SysDictDataMgr) WithDictType(dictType *string) Option {
	return optionFunc(func(o *options) { o.query["dict_type"] = dictType })
}

// WithCSSClass css_class获取
func (obj *_SysDictDataMgr) WithCSSClass(cssClass *string) Option {
	return optionFunc(func(o *options) { o.query["css_class"] = cssClass })
}

// WithListClass list_class获取
func (obj *_SysDictDataMgr) WithListClass(listClass *string) Option {
	return optionFunc(func(o *options) { o.query["list_class"] = listClass })
}

// WithIsDefault is_default获取
func (obj *_SysDictDataMgr) WithIsDefault(isDefault *string) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// WithStatus status获取
func (obj *_SysDictDataMgr) WithStatus(status *int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDefault default获取
func (obj *_SysDictDataMgr) WithDefault(_default *string) Option {
	return optionFunc(func(o *options) { o.query["default"] = _default })
}

// WithRemark remark获取
func (obj *_SysDictDataMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysDictDataMgr) WithCreateBy(createBy *int64) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysDictDataMgr) WithUpdateBy(updateBy *int64) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_SysDictDataMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 最后更新时间
func (obj *_SysDictDataMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SysDictDataMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_SysDictDataMgr) GetByOption(opts ...Option) (result SysDictData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysDictDataMgr) GetByOptions(opts ...Option) (results []*SysDictData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromDictCode 通过dict_code获取内容
func (obj *_SysDictDataMgr) GetFromDictCode(dictCode int64) (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_code` = ?", dictCode).First(&result).Error

	return
}

// GetBatchFromDictCode 批量查找
func (obj *_SysDictDataMgr) GetBatchFromDictCode(dictCodes []int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_code` IN (?)", dictCodes).Find(&results).Error

	return
}

// GetFromDictSort 通过dict_sort获取内容
func (obj *_SysDictDataMgr) GetFromDictSort(dictSort *int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_sort` = ?", dictSort).Find(&results).Error

	return
}

// GetBatchFromDictSort 批量查找
func (obj *_SysDictDataMgr) GetBatchFromDictSort(dictSorts []*int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_sort` IN (?)", dictSorts).Find(&results).Error

	return
}

// GetFromDictLabel 通过dict_label获取内容
func (obj *_SysDictDataMgr) GetFromDictLabel(dictLabel *string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_label` = ?", dictLabel).Find(&results).Error

	return
}

// GetBatchFromDictLabel 批量查找
func (obj *_SysDictDataMgr) GetBatchFromDictLabel(dictLabels []*string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_label` IN (?)", dictLabels).Find(&results).Error

	return
}

// GetFromDictValue 通过dict_value获取内容
func (obj *_SysDictDataMgr) GetFromDictValue(dictValue *string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_value` = ?", dictValue).Find(&results).Error

	return
}

// GetBatchFromDictValue 批量查找
func (obj *_SysDictDataMgr) GetBatchFromDictValue(dictValues []*string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_value` IN (?)", dictValues).Find(&results).Error

	return
}

// GetFromDictType 通过dict_type获取内容
func (obj *_SysDictDataMgr) GetFromDictType(dictType *string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_type` = ?", dictType).Find(&results).Error

	return
}

// GetBatchFromDictType 批量查找
func (obj *_SysDictDataMgr) GetBatchFromDictType(dictTypes []*string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_type` IN (?)", dictTypes).Find(&results).Error

	return
}

// GetFromCSSClass 通过css_class获取内容
func (obj *_SysDictDataMgr) GetFromCSSClass(cssClass *string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`css_class` = ?", cssClass).Find(&results).Error

	return
}

// GetBatchFromCSSClass 批量查找
func (obj *_SysDictDataMgr) GetBatchFromCSSClass(cssClasss []*string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`css_class` IN (?)", cssClasss).Find(&results).Error

	return
}

// GetFromListClass 通过list_class获取内容
func (obj *_SysDictDataMgr) GetFromListClass(listClass *string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`list_class` = ?", listClass).Find(&results).Error

	return
}

// GetBatchFromListClass 批量查找
func (obj *_SysDictDataMgr) GetBatchFromListClass(listClasss []*string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`list_class` IN (?)", listClasss).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容
func (obj *_SysDictDataMgr) GetFromIsDefault(isDefault *string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找
func (obj *_SysDictDataMgr) GetBatchFromIsDefault(isDefaults []*string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_SysDictDataMgr) GetFromStatus(status *int8) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_SysDictDataMgr) GetBatchFromStatus(statuss []*int8) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDefault 通过default获取内容
func (obj *_SysDictDataMgr) GetFromDefault(_default *string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`default` = ?", _default).Find(&results).Error

	return
}

// GetBatchFromDefault 批量查找
func (obj *_SysDictDataMgr) GetBatchFromDefault(_defaults []*string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`default` IN (?)", _defaults).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容
func (obj *_SysDictDataMgr) GetFromRemark(remark *string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找
func (obj *_SysDictDataMgr) GetBatchFromRemark(remarks []*string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_SysDictDataMgr) GetFromCreateBy(createBy *int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找 创建者
func (obj *_SysDictDataMgr) GetBatchFromCreateBy(createBys []*int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_SysDictDataMgr) GetFromUpdateBy(updateBy *int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找 更新者
func (obj *_SysDictDataMgr) GetBatchFromUpdateBy(updateBys []*int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_SysDictDataMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_SysDictDataMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 最后更新时间
func (obj *_SysDictDataMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 最后更新时间
func (obj *_SysDictDataMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SysDictDataMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SysDictDataMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SysDictDataMgr) FetchByPrimaryKey(dictCode int64) (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`dict_code` = ?", dictCode).First(&result).Error

	return
}

// FetchIndexByIDxSysDictDataCreateBy  获取多个内容
func (obj *_SysDictDataMgr) FetchIndexByIDxSysDictDataCreateBy(createBy *int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// FetchIndexByIDxSysDictDataUpdateBy  获取多个内容
func (obj *_SysDictDataMgr) FetchIndexByIDxSysDictDataUpdateBy(updateBy *int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// FetchIndexByIDxSysDictDataDeletedAt  获取多个内容
func (obj *_SysDictDataMgr) FetchIndexByIDxSysDictDataDeletedAt(deletedAt *time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysDictData{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
