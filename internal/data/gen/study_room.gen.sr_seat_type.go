package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrSeatTypeMgr struct {
	*_BaseMgr
}

// SrSeatTypeMgr open func
func SrSeatTypeMgr(db *gorm.DB) *_SrSeatTypeMgr {
	if db == nil {
		panic(fmt.Errorf("SrSeatTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrSeatTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_seat_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrSeatTypeMgr) GetTableName() string {
	return "sr_seat_type"
}

// Reset 重置gorm会话
func (obj *_SrSeatTypeMgr) Reset() *_SrSeatTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrSeatTypeMgr) Get() (result SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrSeatTypeMgr) Gets() (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrSeatTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrSeatTypeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_SrSeatTypeMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithName name获取 座位类型名
func (obj *_SrSeatTypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithShopID shop_id获取 适用店铺id列表
func (obj *_SrSeatTypeMgr) WithShopID(shopID string) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopType shop_type获取 适用店铺类型
func (obj *_SrSeatTypeMgr) WithShopType(shopType *string) Option {
	return optionFunc(func(o *options) { o.query["shop_type"] = shopType })
}

// WithSort sort获取 排序
func (obj *_SrSeatTypeMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrSeatTypeMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrSeatTypeMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrSeatTypeMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrSeatTypeMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrSeatTypeMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrSeatTypeMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrSeatTypeMgr) GetByOption(opts ...Option) (result SrSeatType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrSeatTypeMgr) GetByOptions(opts ...Option) (results []*SrSeatType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrSeatTypeMgr) GetFromID(id int64) (result SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrSeatTypeMgr) GetBatchFromID(ids []int64) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_SrSeatTypeMgr) GetFromMerchantID(merchantID int64) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_SrSeatTypeMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 座位类型名
func (obj *_SrSeatTypeMgr) GetFromName(name string) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 座位类型名
func (obj *_SrSeatTypeMgr) GetBatchFromName(names []string) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 适用店铺id列表
func (obj *_SrSeatTypeMgr) GetFromShopID(shopID string) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 适用店铺id列表
func (obj *_SrSeatTypeMgr) GetBatchFromShopID(shopIDs []string) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopType 通过shop_type获取内容 适用店铺类型
func (obj *_SrSeatTypeMgr) GetFromShopType(shopType *string) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`shop_type` = ?", shopType).Find(&results).Error

	return
}

// GetBatchFromShopType 批量查找 适用店铺类型
func (obj *_SrSeatTypeMgr) GetBatchFromShopType(shopTypes []*string) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`shop_type` IN (?)", shopTypes).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrSeatTypeMgr) GetFromSort(sort int) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrSeatTypeMgr) GetBatchFromSort(sorts []int) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrSeatTypeMgr) GetFromRemark(remark *string) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrSeatTypeMgr) GetBatchFromRemark(remarks []*string) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrSeatTypeMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrSeatTypeMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrSeatTypeMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrSeatTypeMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrSeatTypeMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrSeatTypeMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrSeatTypeMgr) GetFromCreateBy(createBy *uint) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrSeatTypeMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrSeatTypeMgr) GetFromUpdateBy(updateBy *uint) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrSeatTypeMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrSeatTypeMgr) FetchByPrimaryKey(id int64) (result SrSeatType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatType{}).Where("`id` = ?", id).First(&result).Error

	return
}
