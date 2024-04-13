package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdCoursesSchedulingAddressMgr struct {
	*_BaseMgr
}

// EdCoursesSchedulingAddressMgr open func
func EdCoursesSchedulingAddressMgr(db *gorm.DB) *_EdCoursesSchedulingAddressMgr {
	if db == nil {
		panic(fmt.Errorf("EdCoursesSchedulingAddressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdCoursesSchedulingAddressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_courses_scheduling_address"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdCoursesSchedulingAddressMgr) GetTableName() string {
	return "ed_courses_scheduling_address"
}

// Reset 重置gorm会话
func (obj *_EdCoursesSchedulingAddressMgr) Reset() *_EdCoursesSchedulingAddressMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdCoursesSchedulingAddressMgr) Get() (result EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdCoursesSchedulingAddressMgr) Gets() (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdCoursesSchedulingAddressMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdCoursesSchedulingAddressMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdCoursesSchedulingAddressMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdCoursesSchedulingAddressMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithAddressType address_type获取 地址类型
func (obj *_EdCoursesSchedulingAddressMgr) WithAddressType(addressType int8) Option {
	return optionFunc(func(o *options) { o.query["address_type"] = addressType })
}

// WithName name获取 名称
func (obj *_EdCoursesSchedulingAddressMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAddress address获取 地址
func (obj *_EdCoursesSchedulingAddressMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithCreatedAt created_at获取
func (obj *_EdCoursesSchedulingAddressMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdCoursesSchedulingAddressMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdCoursesSchedulingAddressMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdCoursesSchedulingAddressMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdCoursesSchedulingAddressMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdCoursesSchedulingAddressMgr) GetByOption(opts ...Option) (result EdCoursesSchedulingAddress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdCoursesSchedulingAddressMgr) GetByOptions(opts ...Option) (results []*EdCoursesSchedulingAddress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdCoursesSchedulingAddressMgr) GetFromID(id int64) (result EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromID(ids []int64) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdCoursesSchedulingAddressMgr) GetFromMerchantID(merchantID int64) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdCoursesSchedulingAddressMgr) GetFromShopID(shopID int64) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromAddressType 通过address_type获取内容 地址类型
func (obj *_EdCoursesSchedulingAddressMgr) GetFromAddressType(addressType int8) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`address_type` = ?", addressType).Find(&results).Error

	return
}

// GetBatchFromAddressType 批量查找 地址类型
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromAddressType(addressTypes []int8) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`address_type` IN (?)", addressTypes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_EdCoursesSchedulingAddressMgr) GetFromName(name string) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromName(names []string) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 地址
func (obj *_EdCoursesSchedulingAddressMgr) GetFromAddress(address string) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 地址
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromAddress(addresss []string) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdCoursesSchedulingAddressMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdCoursesSchedulingAddressMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdCoursesSchedulingAddressMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdCoursesSchedulingAddressMgr) GetFromCreateBy(createBy *uint) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdCoursesSchedulingAddressMgr) GetFromUpdateBy(updateBy *uint) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdCoursesSchedulingAddressMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdCoursesSchedulingAddressMgr) FetchByPrimaryKey(id int64) (result EdCoursesSchedulingAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingAddress{}).Where("`id` = ?", id).First(&result).Error

	return
}
