package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdDictTypeMgr struct {
	*_BaseMgr
}

// EdDictTypeMgr open func
func EdDictTypeMgr(db *gorm.DB) *_EdDictTypeMgr {
	if db == nil {
		panic(fmt.Errorf("EdDictTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdDictTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_dict_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdDictTypeMgr) GetTableName() string {
	return "ed_dict_type"
}

// Reset 重置gorm会话
func (obj *_EdDictTypeMgr) Reset() *_EdDictTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdDictTypeMgr) Get() (result EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdDictTypeMgr) Gets() (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdDictTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDictID dict_id获取
func (obj *_EdDictTypeMgr) WithDictID(dictID int64) Option {
	return optionFunc(func(o *options) { o.query["dict_id"] = dictID })
}

// WithDictName dict_name获取
func (obj *_EdDictTypeMgr) WithDictName(dictName *string) Option {
	return optionFunc(func(o *options) { o.query["dict_name"] = dictName })
}

// WithDictType dict_type获取
func (obj *_EdDictTypeMgr) WithDictType(dictType *string) Option {
	return optionFunc(func(o *options) { o.query["dict_type"] = dictType })
}

// WithStatus status获取
func (obj *_EdDictTypeMgr) WithStatus(status *int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRemark remark获取
func (obj *_EdDictTypeMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateBy create_by获取 创建者
func (obj *_EdDictTypeMgr) WithCreateBy(createBy *int64) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取 更新者
func (obj *_EdDictTypeMgr) WithUpdateBy(updateBy *int64) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_EdDictTypeMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 最后更新时间
func (obj *_EdDictTypeMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdDictTypeMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithClass class获取 分类
func (obj *_EdDictTypeMgr) WithClass(class string) Option {
	return optionFunc(func(o *options) { o.query["class"] = class })
}

// GetByOption 功能选项模式获取
func (obj *_EdDictTypeMgr) GetByOption(opts ...Option) (result EdDictType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdDictTypeMgr) GetByOptions(opts ...Option) (results []*EdDictType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromDictID 通过dict_id获取内容
func (obj *_EdDictTypeMgr) GetFromDictID(dictID int64) (result EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`dict_id` = ?", dictID).First(&result).Error

	return
}

// GetBatchFromDictID 批量查找
func (obj *_EdDictTypeMgr) GetBatchFromDictID(dictIDs []int64) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`dict_id` IN (?)", dictIDs).Find(&results).Error

	return
}

// GetFromDictName 通过dict_name获取内容
func (obj *_EdDictTypeMgr) GetFromDictName(dictName *string) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`dict_name` = ?", dictName).Find(&results).Error

	return
}

// GetBatchFromDictName 批量查找
func (obj *_EdDictTypeMgr) GetBatchFromDictName(dictNames []*string) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`dict_name` IN (?)", dictNames).Find(&results).Error

	return
}

// GetFromDictType 通过dict_type获取内容
func (obj *_EdDictTypeMgr) GetFromDictType(dictType *string) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`dict_type` = ?", dictType).Find(&results).Error

	return
}

// GetBatchFromDictType 批量查找
func (obj *_EdDictTypeMgr) GetBatchFromDictType(dictTypes []*string) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`dict_type` IN (?)", dictTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_EdDictTypeMgr) GetFromStatus(status *int8) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_EdDictTypeMgr) GetBatchFromStatus(statuss []*int8) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容
func (obj *_EdDictTypeMgr) GetFromRemark(remark *string) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找
func (obj *_EdDictTypeMgr) GetBatchFromRemark(remarks []*string) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_EdDictTypeMgr) GetFromCreateBy(createBy *int64) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找 创建者
func (obj *_EdDictTypeMgr) GetBatchFromCreateBy(createBys []*int64) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_EdDictTypeMgr) GetFromUpdateBy(updateBy *int64) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找 更新者
func (obj *_EdDictTypeMgr) GetBatchFromUpdateBy(updateBys []*int64) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_EdDictTypeMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_EdDictTypeMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 最后更新时间
func (obj *_EdDictTypeMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 最后更新时间
func (obj *_EdDictTypeMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdDictTypeMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdDictTypeMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromClass 通过class获取内容 分类
func (obj *_EdDictTypeMgr) GetFromClass(class string) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`class` = ?", class).Find(&results).Error

	return
}

// GetBatchFromClass 批量查找 分类
func (obj *_EdDictTypeMgr) GetBatchFromClass(classs []string) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`class` IN (?)", classs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdDictTypeMgr) FetchByPrimaryKey(dictID int64) (result EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`dict_id` = ?", dictID).First(&result).Error

	return
}

// FetchIndexByIDxSysDictTypeCreateBy  获取多个内容
func (obj *_EdDictTypeMgr) FetchIndexByIDxSysDictTypeCreateBy(createBy *int64) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// FetchIndexByIDxSysDictTypeUpdateBy  获取多个内容
func (obj *_EdDictTypeMgr) FetchIndexByIDxSysDictTypeUpdateBy(updateBy *int64) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// FetchIndexByIDxSysDictTypeDeletedAt  获取多个内容
func (obj *_EdDictTypeMgr) FetchIndexByIDxSysDictTypeDeletedAt(deletedAt *time.Time) (results []*EdDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdDictType{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
