package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdSalemanMgr struct {
	*_BaseMgr
}

// EdSalemanMgr open func
func EdSalemanMgr(db *gorm.DB) *_EdSalemanMgr {
	if db == nil {
		panic(fmt.Errorf("EdSalemanMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdSalemanMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_saleman"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdSalemanMgr) GetTableName() string {
	return "ed_saleman"
}

// Reset 重置gorm会话
func (obj *_EdSalemanMgr) Reset() *_EdSalemanMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdSalemanMgr) Get() (result EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdSalemanMgr) Gets() (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdSalemanMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdSalemanMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdSalemanMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdSalemanMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithSaleManType sale_man_type获取 销售人员类型
func (obj *_EdSalemanMgr) WithSaleManType(saleManType int8) Option {
	return optionFunc(func(o *options) { o.query["sale_man_type"] = saleManType })
}

// WithName name获取 名称
func (obj *_EdSalemanMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreatedAt created_at获取
func (obj *_EdSalemanMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdSalemanMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdSalemanMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdSalemanMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdSalemanMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdSalemanMgr) GetByOption(opts ...Option) (result EdSaleman, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdSalemanMgr) GetByOptions(opts ...Option) (results []*EdSaleman, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdSalemanMgr) GetFromID(id int64) (result EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdSalemanMgr) GetBatchFromID(ids []int64) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdSalemanMgr) GetFromMerchantID(merchantID int64) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdSalemanMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdSalemanMgr) GetFromShopID(shopID int64) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdSalemanMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromSaleManType 通过sale_man_type获取内容 销售人员类型
func (obj *_EdSalemanMgr) GetFromSaleManType(saleManType int8) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`sale_man_type` = ?", saleManType).Find(&results).Error

	return
}

// GetBatchFromSaleManType 批量查找 销售人员类型
func (obj *_EdSalemanMgr) GetBatchFromSaleManType(saleManTypes []int8) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`sale_man_type` IN (?)", saleManTypes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_EdSalemanMgr) GetFromName(name string) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_EdSalemanMgr) GetBatchFromName(names []string) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdSalemanMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdSalemanMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdSalemanMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdSalemanMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdSalemanMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdSalemanMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdSalemanMgr) GetFromCreateBy(createBy *uint) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdSalemanMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdSalemanMgr) GetFromUpdateBy(updateBy *uint) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdSalemanMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdSalemanMgr) FetchByPrimaryKey(id int64) (result EdSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSaleman{}).Where("`id` = ?", id).First(&result).Error

	return
}
