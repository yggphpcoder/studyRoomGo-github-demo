package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DianpingTuangouMgr struct {
	*_BaseMgr
}

// DianpingTuangouMgr open func
func DianpingTuangouMgr(db *gorm.DB) *_DianpingTuangouMgr {
	if db == nil {
		panic(fmt.Errorf("DianpingTuangouMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DianpingTuangouMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dianping_tuangou"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DianpingTuangouMgr) GetTableName() string {
	return "dianping_tuangou"
}

// Reset 重置gorm会话
func (obj *_DianpingTuangouMgr) Reset() *_DianpingTuangouMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DianpingTuangouMgr) Get() (result DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DianpingTuangouMgr) Gets() (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DianpingTuangouMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DianpingTuangouMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_DianpingTuangouMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithPackageCardID package_card_id获取 系统套餐id
func (obj *_DianpingTuangouMgr) WithPackageCardID(packageCardID int64) Option {
	return optionFunc(func(o *options) { o.query["package_card_id"] = packageCardID })
}

// WithOpenShopUUID open_shop_uuid获取 门店id的唯一标识
func (obj *_DianpingTuangouMgr) WithOpenShopUUID(openShopUUID string) Option {
	return optionFunc(func(o *options) { o.query["open_shop_uuid"] = openShopUUID })
}

// WithDealID deal_id获取 套餐id
func (obj *_DianpingTuangouMgr) WithDealID(dealID int64) Option {
	return optionFunc(func(o *options) { o.query["deal_id"] = dealID })
}

// WithDealgroupID dealgroup_id获取 团购id
func (obj *_DianpingTuangouMgr) WithDealgroupID(dealgroupID int64) Option {
	return optionFunc(func(o *options) { o.query["dealgroup_id"] = dealgroupID })
}

// WithTitle title获取 套餐名称
func (obj *_DianpingTuangouMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithPrice price获取 套餐价格
func (obj *_DianpingTuangouMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithMarketprice marketprice获取 套餐原价
func (obj *_DianpingTuangouMgr) WithMarketprice(marketprice float64) Option {
	return optionFunc(func(o *options) { o.query["marketprice"] = marketprice })
}

// GetByOption 功能选项模式获取
func (obj *_DianpingTuangouMgr) GetByOption(opts ...Option) (result DianpingTuangou, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DianpingTuangouMgr) GetByOptions(opts ...Option) (results []*DianpingTuangou, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DianpingTuangouMgr) GetFromID(id int) (result DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DianpingTuangouMgr) GetBatchFromID(ids []int) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_DianpingTuangouMgr) GetFromMerchantID(merchantID int64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_DianpingTuangouMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromPackageCardID 通过package_card_id获取内容 系统套餐id
func (obj *_DianpingTuangouMgr) GetFromPackageCardID(packageCardID int64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`package_card_id` = ?", packageCardID).Find(&results).Error

	return
}

// GetBatchFromPackageCardID 批量查找 系统套餐id
func (obj *_DianpingTuangouMgr) GetBatchFromPackageCardID(packageCardIDs []int64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`package_card_id` IN (?)", packageCardIDs).Find(&results).Error

	return
}

// GetFromOpenShopUUID 通过open_shop_uuid获取内容 门店id的唯一标识
func (obj *_DianpingTuangouMgr) GetFromOpenShopUUID(openShopUUID string) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`open_shop_uuid` = ?", openShopUUID).Find(&results).Error

	return
}

// GetBatchFromOpenShopUUID 批量查找 门店id的唯一标识
func (obj *_DianpingTuangouMgr) GetBatchFromOpenShopUUID(openShopUUIDs []string) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`open_shop_uuid` IN (?)", openShopUUIDs).Find(&results).Error

	return
}

// GetFromDealID 通过deal_id获取内容 套餐id
func (obj *_DianpingTuangouMgr) GetFromDealID(dealID int64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`deal_id` = ?", dealID).Find(&results).Error

	return
}

// GetBatchFromDealID 批量查找 套餐id
func (obj *_DianpingTuangouMgr) GetBatchFromDealID(dealIDs []int64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`deal_id` IN (?)", dealIDs).Find(&results).Error

	return
}

// GetFromDealgroupID 通过dealgroup_id获取内容 团购id
func (obj *_DianpingTuangouMgr) GetFromDealgroupID(dealgroupID int64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`dealgroup_id` = ?", dealgroupID).Find(&results).Error

	return
}

// GetBatchFromDealgroupID 批量查找 团购id
func (obj *_DianpingTuangouMgr) GetBatchFromDealgroupID(dealgroupIDs []int64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`dealgroup_id` IN (?)", dealgroupIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 套餐名称
func (obj *_DianpingTuangouMgr) GetFromTitle(title string) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 套餐名称
func (obj *_DianpingTuangouMgr) GetBatchFromTitle(titles []string) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 套餐价格
func (obj *_DianpingTuangouMgr) GetFromPrice(price float64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 套餐价格
func (obj *_DianpingTuangouMgr) GetBatchFromPrice(prices []float64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromMarketprice 通过marketprice获取内容 套餐原价
func (obj *_DianpingTuangouMgr) GetFromMarketprice(marketprice float64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`marketprice` = ?", marketprice).Find(&results).Error

	return
}

// GetBatchFromMarketprice 批量查找 套餐原价
func (obj *_DianpingTuangouMgr) GetBatchFromMarketprice(marketprices []float64) (results []*DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`marketprice` IN (?)", marketprices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DianpingTuangouMgr) FetchByPrimaryKey(id int) (result DianpingTuangou, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangou{}).Where("`id` = ?", id).First(&result).Error

	return
}
