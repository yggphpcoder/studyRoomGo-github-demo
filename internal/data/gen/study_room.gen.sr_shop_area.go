package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrShopAreaMgr struct {
	*_BaseMgr
}

// SrShopAreaMgr open func
func SrShopAreaMgr(db *gorm.DB) *_SrShopAreaMgr {
	if db == nil {
		panic(fmt.Errorf("SrShopAreaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrShopAreaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_shop_area"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrShopAreaMgr) GetTableName() string {
	return "sr_shop_area"
}

// Reset 重置gorm会话
func (obj *_SrShopAreaMgr) Reset() *_SrShopAreaMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrShopAreaMgr) Get() (result SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrShopAreaMgr) Gets() (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrShopAreaMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrShopAreaMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 区域名
func (obj *_SrShopAreaMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCover cover获取 封面图
func (obj *_SrShopAreaMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrShopAreaMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithTypeArea type_area获取 区域类型
func (obj *_SrShopAreaMgr) WithTypeArea(typeArea int8) Option {
	return optionFunc(func(o *options) { o.query["type_area"] = typeArea })
}

// WithSort sort获取 排序
func (obj *_SrShopAreaMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrShopAreaMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrShopAreaMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrShopAreaMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrShopAreaMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrShopAreaMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrShopAreaMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithIsCharge is_charge获取 是否可设置收费
func (obj *_SrShopAreaMgr) WithIsCharge(isCharge int8) Option {
	return optionFunc(func(o *options) { o.query["is_charge"] = isCharge })
}

// GetByOption 功能选项模式获取
func (obj *_SrShopAreaMgr) GetByOption(opts ...Option) (result SrShopArea, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrShopAreaMgr) GetByOptions(opts ...Option) (results []*SrShopArea, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrShopAreaMgr) GetFromID(id int64) (result SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrShopAreaMgr) GetBatchFromID(ids []int64) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 区域名
func (obj *_SrShopAreaMgr) GetFromName(name string) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 区域名
func (obj *_SrShopAreaMgr) GetBatchFromName(names []string) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容 封面图
func (obj *_SrShopAreaMgr) GetFromCover(cover string) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`cover` = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量查找 封面图
func (obj *_SrShopAreaMgr) GetBatchFromCover(covers []string) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`cover` IN (?)", covers).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrShopAreaMgr) GetFromShopID(shopID int64) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrShopAreaMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromTypeArea 通过type_area获取内容 区域类型
func (obj *_SrShopAreaMgr) GetFromTypeArea(typeArea int8) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`type_area` = ?", typeArea).Find(&results).Error

	return
}

// GetBatchFromTypeArea 批量查找 区域类型
func (obj *_SrShopAreaMgr) GetBatchFromTypeArea(typeAreas []int8) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`type_area` IN (?)", typeAreas).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrShopAreaMgr) GetFromSort(sort int) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrShopAreaMgr) GetBatchFromSort(sorts []int) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrShopAreaMgr) GetFromRemark(remark *string) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrShopAreaMgr) GetBatchFromRemark(remarks []*string) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrShopAreaMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrShopAreaMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrShopAreaMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrShopAreaMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrShopAreaMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrShopAreaMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrShopAreaMgr) GetFromCreateBy(createBy *uint) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrShopAreaMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrShopAreaMgr) GetFromUpdateBy(updateBy *uint) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrShopAreaMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromIsCharge 通过is_charge获取内容 是否可设置收费
func (obj *_SrShopAreaMgr) GetFromIsCharge(isCharge int8) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`is_charge` = ?", isCharge).Find(&results).Error

	return
}

// GetBatchFromIsCharge 批量查找 是否可设置收费
func (obj *_SrShopAreaMgr) GetBatchFromIsCharge(isCharges []int8) (results []*SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`is_charge` IN (?)", isCharges).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrShopAreaMgr) FetchByPrimaryKey(id int64) (result SrShopArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopArea{}).Where("`id` = ?", id).First(&result).Error

	return
}
