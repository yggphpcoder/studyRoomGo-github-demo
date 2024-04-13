package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrShopAreaSeatPriceMgr struct {
	*_BaseMgr
}

// SrShopAreaSeatPriceMgr open func
func SrShopAreaSeatPriceMgr(db *gorm.DB) *_SrShopAreaSeatPriceMgr {
	if db == nil {
		panic(fmt.Errorf("SrShopAreaSeatPriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrShopAreaSeatPriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_shop_area_seat_price"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrShopAreaSeatPriceMgr) GetTableName() string {
	return "sr_shop_area_seat_price"
}

// Reset 重置gorm会话
func (obj *_SrShopAreaSeatPriceMgr) Reset() *_SrShopAreaSeatPriceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrShopAreaSeatPriceMgr) Get() (result SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrShopAreaSeatPriceMgr) Gets() (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrShopAreaSeatPriceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrShopAreaSeatPriceMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAreaID area_id获取 区域id
func (obj *_SrShopAreaSeatPriceMgr) WithAreaID(areaID int64) Option {
	return optionFunc(func(o *options) { o.query["area_id"] = areaID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrShopAreaSeatPriceMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithPrice price获取 售价
func (obj *_SrShopAreaSeatPriceMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithUseTime use_time获取 使用时段
func (obj *_SrShopAreaSeatPriceMgr) WithUseTime(useTime string) Option {
	return optionFunc(func(o *options) { o.query["use_time"] = useTime })
}

// WithSort sort获取 排序
func (obj *_SrShopAreaSeatPriceMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrShopAreaSeatPriceMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrShopAreaSeatPriceMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrShopAreaSeatPriceMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrShopAreaSeatPriceMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrShopAreaSeatPriceMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrShopAreaSeatPriceMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrShopAreaSeatPriceMgr) GetByOption(opts ...Option) (result SrShopAreaSeatPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrShopAreaSeatPriceMgr) GetByOptions(opts ...Option) (results []*SrShopAreaSeatPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrShopAreaSeatPriceMgr) GetFromID(id int64) (result SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromID(ids []int64) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAreaID 通过area_id获取内容 区域id
func (obj *_SrShopAreaSeatPriceMgr) GetFromAreaID(areaID int64) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`area_id` = ?", areaID).Find(&results).Error

	return
}

// GetBatchFromAreaID 批量查找 区域id
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromAreaID(areaIDs []int64) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`area_id` IN (?)", areaIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrShopAreaSeatPriceMgr) GetFromShopID(shopID int64) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 售价
func (obj *_SrShopAreaSeatPriceMgr) GetFromPrice(price float64) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 售价
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromPrice(prices []float64) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromUseTime 通过use_time获取内容 使用时段
func (obj *_SrShopAreaSeatPriceMgr) GetFromUseTime(useTime string) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`use_time` = ?", useTime).Find(&results).Error

	return
}

// GetBatchFromUseTime 批量查找 使用时段
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromUseTime(useTimes []string) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`use_time` IN (?)", useTimes).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrShopAreaSeatPriceMgr) GetFromSort(sort int) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromSort(sorts []int) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrShopAreaSeatPriceMgr) GetFromRemark(remark *string) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromRemark(remarks []*string) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrShopAreaSeatPriceMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrShopAreaSeatPriceMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrShopAreaSeatPriceMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrShopAreaSeatPriceMgr) GetFromCreateBy(createBy *uint) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrShopAreaSeatPriceMgr) GetFromUpdateBy(updateBy *uint) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrShopAreaSeatPriceMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrShopAreaSeatPriceMgr) FetchByPrimaryKey(id int64) (result SrShopAreaSeatPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShopAreaSeatPrice{}).Where("`id` = ?", id).First(&result).Error

	return
}
