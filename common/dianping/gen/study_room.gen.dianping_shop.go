package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DianpingShopMgr struct {
	*_BaseMgr
}

// DianpingShopMgr open func
func DianpingShopMgr(db *gorm.DB) *_DianpingShopMgr {
	if db == nil {
		panic(fmt.Errorf("DianpingShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DianpingShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dianping_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DianpingShopMgr) GetTableName() string {
	return "dianping_shop"
}

// Reset 重置gorm会话
func (obj *_DianpingShopMgr) Reset() *_DianpingShopMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DianpingShopMgr) Get() (result DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DianpingShopMgr) Gets() (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DianpingShopMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DianpingShopMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 系统商家id
func (obj *_DianpingShopMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 系统shop_id
func (obj *_DianpingShopMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithOpenShopUUID open_shop_uuid获取 门店id的唯一标识
func (obj *_DianpingShopMgr) WithOpenShopUUID(openShopUUID string) Option {
	return optionFunc(func(o *options) { o.query["open_shop_uuid"] = openShopUUID })
}

// WithShopName shop_name获取 门店名称
func (obj *_DianpingShopMgr) WithShopName(shopName string) Option {
	return optionFunc(func(o *options) { o.query["shop_name"] = shopName })
}

// WithBranchName branch_name获取 分店名称
func (obj *_DianpingShopMgr) WithBranchName(branchName string) Option {
	return optionFunc(func(o *options) { o.query["branch_name"] = branchName })
}

// WithShopAddress shop_address获取 门店地址
func (obj *_DianpingShopMgr) WithShopAddress(shopAddress string) Option {
	return optionFunc(func(o *options) { o.query["shop_address"] = shopAddress })
}

// WithCityName city_name获取 所在城市
func (obj *_DianpingShopMgr) WithCityName(cityName string) Option {
	return optionFunc(func(o *options) { o.query["city_name"] = cityName })
}

// GetByOption 功能选项模式获取
func (obj *_DianpingShopMgr) GetByOption(opts ...Option) (result DianpingShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DianpingShopMgr) GetByOptions(opts ...Option) (results []*DianpingShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DianpingShopMgr) GetFromID(id int) (result DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DianpingShopMgr) GetBatchFromID(ids []int) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 系统商家id
func (obj *_DianpingShopMgr) GetFromMerchantID(merchantID int64) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 系统商家id
func (obj *_DianpingShopMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 系统shop_id
func (obj *_DianpingShopMgr) GetFromShopID(shopID int64) (result DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`shop_id` = ?", shopID).First(&result).Error

	return
}

// GetBatchFromShopID 批量查找 系统shop_id
func (obj *_DianpingShopMgr) GetBatchFromShopID(shopIDs []int64) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromOpenShopUUID 通过open_shop_uuid获取内容 门店id的唯一标识
func (obj *_DianpingShopMgr) GetFromOpenShopUUID(openShopUUID string) (result DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`open_shop_uuid` = ?", openShopUUID).First(&result).Error

	return
}

// GetBatchFromOpenShopUUID 批量查找 门店id的唯一标识
func (obj *_DianpingShopMgr) GetBatchFromOpenShopUUID(openShopUUIDs []string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`open_shop_uuid` IN (?)", openShopUUIDs).Find(&results).Error

	return
}

// GetFromShopName 通过shop_name获取内容 门店名称
func (obj *_DianpingShopMgr) GetFromShopName(shopName string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`shop_name` = ?", shopName).Find(&results).Error

	return
}

// GetBatchFromShopName 批量查找 门店名称
func (obj *_DianpingShopMgr) GetBatchFromShopName(shopNames []string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`shop_name` IN (?)", shopNames).Find(&results).Error

	return
}

// GetFromBranchName 通过branch_name获取内容 分店名称
func (obj *_DianpingShopMgr) GetFromBranchName(branchName string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`branch_name` = ?", branchName).Find(&results).Error

	return
}

// GetBatchFromBranchName 批量查找 分店名称
func (obj *_DianpingShopMgr) GetBatchFromBranchName(branchNames []string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`branch_name` IN (?)", branchNames).Find(&results).Error

	return
}

// GetFromShopAddress 通过shop_address获取内容 门店地址
func (obj *_DianpingShopMgr) GetFromShopAddress(shopAddress string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`shop_address` = ?", shopAddress).Find(&results).Error

	return
}

// GetBatchFromShopAddress 批量查找 门店地址
func (obj *_DianpingShopMgr) GetBatchFromShopAddress(shopAddresss []string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`shop_address` IN (?)", shopAddresss).Find(&results).Error

	return
}

// GetFromCityName 通过city_name获取内容 所在城市
func (obj *_DianpingShopMgr) GetFromCityName(cityName string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`city_name` = ?", cityName).Find(&results).Error

	return
}

// GetBatchFromCityName 批量查找 所在城市
func (obj *_DianpingShopMgr) GetBatchFromCityName(cityNames []string) (results []*DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`city_name` IN (?)", cityNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DianpingShopMgr) FetchByPrimaryKey(id int) (result DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByDianpingShopShopIDUIndex primary or index 获取唯一内容
func (obj *_DianpingShopMgr) FetchUniqueByDianpingShopShopIDUIndex(shopID int64) (result DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`shop_id` = ?", shopID).First(&result).Error

	return
}

// FetchUniqueByDianpingShopOpenShopUUIDUIndex primary or index 获取唯一内容
func (obj *_DianpingShopMgr) FetchUniqueByDianpingShopOpenShopUUIDUIndex(openShopUUID string) (result DianpingShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingShop{}).Where("`open_shop_uuid` = ?", openShopUUID).First(&result).Error

	return
}
