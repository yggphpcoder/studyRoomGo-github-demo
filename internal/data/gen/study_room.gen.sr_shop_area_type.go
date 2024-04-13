package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrShopAreaTypeMgr struct {
	*_BaseMgr
}

// SrShopAreaTypeMgr open func
func SrShopAreaTypeMgr(db *gorm.DB) *_SrShopAreaTypeMgr {
	if db == nil {
		panic(fmt.Errorf("SrShopAreaTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrShopAreaTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_shop_area_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrShopAreaTypeMgr) GetTableName() string {
	return "sr_shop_area_type"
}

// Reset 重置gorm会话
func (obj *_SrShopAreaTypeMgr) Reset() *_SrShopAreaTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrShopAreaTypeMgr) Get() (result SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrShopAreaTypeMgr) Gets() (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrShopAreaTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrShopAreaTypeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 分区类型名
func (obj *_SrShopAreaTypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithShopID shop_id获取 适用店铺id列表
func (obj *_SrShopAreaTypeMgr) WithShopID(shopID string) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithSort sort获取 排序
func (obj *_SrShopAreaTypeMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrShopAreaTypeMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrShopAreaTypeMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrShopAreaTypeMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrShopAreaTypeMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrShopAreaTypeMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrShopAreaTypeMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrShopAreaTypeMgr) GetByOption(opts ...Option) (result SrShopAreaType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrShopAreaTypeMgr) GetByOptions(opts ...Option) (results []*SrShopAreaType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrShopAreaTypeMgr) GetFromID(id int64) (result SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrShopAreaTypeMgr) GetBatchFromID(ids []int64) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 分区类型名
func (obj *_SrShopAreaTypeMgr) GetFromName(name string) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 分区类型名
func (obj *_SrShopAreaTypeMgr) GetBatchFromName(names []string) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 适用店铺id列表
func (obj *_SrShopAreaTypeMgr) GetFromShopID(shopID string) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 适用店铺id列表
func (obj *_SrShopAreaTypeMgr) GetBatchFromShopID(shopIDs []string) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrShopAreaTypeMgr) GetFromSort(sort int) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrShopAreaTypeMgr) GetBatchFromSort(sorts []int) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrShopAreaTypeMgr) GetFromRemark(remark *string) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrShopAreaTypeMgr) GetBatchFromRemark(remarks []*string) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrShopAreaTypeMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrShopAreaTypeMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrShopAreaTypeMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrShopAreaTypeMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrShopAreaTypeMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrShopAreaTypeMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrShopAreaTypeMgr) GetFromCreateBy(createBy *uint) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrShopAreaTypeMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrShopAreaTypeMgr) GetFromUpdateBy(updateBy *uint) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrShopAreaTypeMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrShopAreaTypeMgr) FetchByPrimaryKey(id int64) (result SrShopAreaType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaType{}).Where("`id` = ?", id).First(&result).Error

	return
}
