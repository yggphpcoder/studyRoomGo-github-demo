package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SrCouponsReceiveApplicableMgr struct {
	*_BaseMgr
}

// SrCouponsReceiveApplicableMgr open func
func SrCouponsReceiveApplicableMgr(db *gorm.DB) *_SrCouponsReceiveApplicableMgr {
	if db == nil {
		panic(fmt.Errorf("SrCouponsReceiveApplicableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrCouponsReceiveApplicableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_coupons_receive_applicable"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrCouponsReceiveApplicableMgr) GetTableName() string {
	return "sr_coupons_receive_applicable"
}

// Reset 重置gorm会话
func (obj *_SrCouponsReceiveApplicableMgr) Reset() *_SrCouponsReceiveApplicableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrCouponsReceiveApplicableMgr) Get() (result SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrCouponsReceiveApplicableMgr) Gets() (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrCouponsReceiveApplicableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_SrCouponsReceiveApplicableMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCouponsReceiveID coupons_receive_id获取 优惠券领取表ID
func (obj *_SrCouponsReceiveApplicableMgr) WithCouponsReceiveID(couponsReceiveID int64) Option {
	return optionFunc(func(o *options) { o.query["coupons_receive_id"] = couponsReceiveID })
}

// WithApplicableOrderType applicable_order_type获取 适用订单类型
func (obj *_SrCouponsReceiveApplicableMgr) WithApplicableOrderType(applicableOrderType int8) Option {
	return optionFunc(func(o *options) { o.query["applicable_order_type"] = applicableOrderType })
}

// WithApplicableShop applicable_shop获取 适用店铺列表
func (obj *_SrCouponsReceiveApplicableMgr) WithApplicableShop(applicableShop string) Option {
	return optionFunc(func(o *options) { o.query["applicable_shop"] = applicableShop })
}

// WithApplicableShopType applicable_shop_type获取 适用店铺类型列表
func (obj *_SrCouponsReceiveApplicableMgr) WithApplicableShopType(applicableShopType string) Option {
	return optionFunc(func(o *options) { o.query["applicable_shop_type"] = applicableShopType })
}

// WithApplicableSeat applicable_seat获取 适用座位列表
func (obj *_SrCouponsReceiveApplicableMgr) WithApplicableSeat(applicableSeat string) Option {
	return optionFunc(func(o *options) { o.query["applicable_seat"] = applicableSeat })
}

// WithApplicableSeatType applicable_seat_type获取 适用座位类型列表
func (obj *_SrCouponsReceiveApplicableMgr) WithApplicableSeatType(applicableSeatType string) Option {
	return optionFunc(func(o *options) { o.query["applicable_seat_type"] = applicableSeatType })
}

// WithApplicablePackageCard applicable_package_card获取 适用套餐卡列表
func (obj *_SrCouponsReceiveApplicableMgr) WithApplicablePackageCard(applicablePackageCard string) Option {
	return optionFunc(func(o *options) { o.query["applicable_package_card"] = applicablePackageCard })
}

// WithApplicablePackageCardType applicable_package_card_type获取 适用套餐卡类型列表
func (obj *_SrCouponsReceiveApplicableMgr) WithApplicablePackageCardType(applicablePackageCardType string) Option {
	return optionFunc(func(o *options) { o.query["applicable_package_card_type"] = applicablePackageCardType })
}

// WithApplicableMemberType applicable_member_type获取 适用用户类型列表
func (obj *_SrCouponsReceiveApplicableMgr) WithApplicableMemberType(applicableMemberType string) Option {
	return optionFunc(func(o *options) { o.query["applicable_member_type"] = applicableMemberType })
}

// GetByOption 功能选项模式获取
func (obj *_SrCouponsReceiveApplicableMgr) GetByOption(opts ...Option) (result SrCouponsReceiveApplicable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrCouponsReceiveApplicableMgr) GetByOptions(opts ...Option) (results []*SrCouponsReceiveApplicable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 id
func (obj *_SrCouponsReceiveApplicableMgr) GetFromID(id int64) (result SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromID(ids []int64) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCouponsReceiveID 通过coupons_receive_id获取内容 优惠券领取表ID
func (obj *_SrCouponsReceiveApplicableMgr) GetFromCouponsReceiveID(couponsReceiveID int64) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`coupons_receive_id` = ?", couponsReceiveID).Find(&results).Error

	return
}

// GetBatchFromCouponsReceiveID 批量查找 优惠券领取表ID
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromCouponsReceiveID(couponsReceiveIDs []int64) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`coupons_receive_id` IN (?)", couponsReceiveIDs).Find(&results).Error

	return
}

// GetFromApplicableOrderType 通过applicable_order_type获取内容 适用订单类型
func (obj *_SrCouponsReceiveApplicableMgr) GetFromApplicableOrderType(applicableOrderType int8) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_order_type` = ?", applicableOrderType).Find(&results).Error

	return
}

// GetBatchFromApplicableOrderType 批量查找 适用订单类型
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromApplicableOrderType(applicableOrderTypes []int8) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_order_type` IN (?)", applicableOrderTypes).Find(&results).Error

	return
}

// GetFromApplicableShop 通过applicable_shop获取内容 适用店铺列表
func (obj *_SrCouponsReceiveApplicableMgr) GetFromApplicableShop(applicableShop string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_shop` = ?", applicableShop).Find(&results).Error

	return
}

// GetBatchFromApplicableShop 批量查找 适用店铺列表
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromApplicableShop(applicableShops []string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_shop` IN (?)", applicableShops).Find(&results).Error

	return
}

// GetFromApplicableShopType 通过applicable_shop_type获取内容 适用店铺类型列表
func (obj *_SrCouponsReceiveApplicableMgr) GetFromApplicableShopType(applicableShopType string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_shop_type` = ?", applicableShopType).Find(&results).Error

	return
}

// GetBatchFromApplicableShopType 批量查找 适用店铺类型列表
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromApplicableShopType(applicableShopTypes []string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_shop_type` IN (?)", applicableShopTypes).Find(&results).Error

	return
}

// GetFromApplicableSeat 通过applicable_seat获取内容 适用座位列表
func (obj *_SrCouponsReceiveApplicableMgr) GetFromApplicableSeat(applicableSeat string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_seat` = ?", applicableSeat).Find(&results).Error

	return
}

// GetBatchFromApplicableSeat 批量查找 适用座位列表
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromApplicableSeat(applicableSeats []string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_seat` IN (?)", applicableSeats).Find(&results).Error

	return
}

// GetFromApplicableSeatType 通过applicable_seat_type获取内容 适用座位类型列表
func (obj *_SrCouponsReceiveApplicableMgr) GetFromApplicableSeatType(applicableSeatType string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_seat_type` = ?", applicableSeatType).Find(&results).Error

	return
}

// GetBatchFromApplicableSeatType 批量查找 适用座位类型列表
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromApplicableSeatType(applicableSeatTypes []string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_seat_type` IN (?)", applicableSeatTypes).Find(&results).Error

	return
}

// GetFromApplicablePackageCard 通过applicable_package_card获取内容 适用套餐卡列表
func (obj *_SrCouponsReceiveApplicableMgr) GetFromApplicablePackageCard(applicablePackageCard string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_package_card` = ?", applicablePackageCard).Find(&results).Error

	return
}

// GetBatchFromApplicablePackageCard 批量查找 适用套餐卡列表
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromApplicablePackageCard(applicablePackageCards []string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_package_card` IN (?)", applicablePackageCards).Find(&results).Error

	return
}

// GetFromApplicablePackageCardType 通过applicable_package_card_type获取内容 适用套餐卡类型列表
func (obj *_SrCouponsReceiveApplicableMgr) GetFromApplicablePackageCardType(applicablePackageCardType string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_package_card_type` = ?", applicablePackageCardType).Find(&results).Error

	return
}

// GetBatchFromApplicablePackageCardType 批量查找 适用套餐卡类型列表
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromApplicablePackageCardType(applicablePackageCardTypes []string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_package_card_type` IN (?)", applicablePackageCardTypes).Find(&results).Error

	return
}

// GetFromApplicableMemberType 通过applicable_member_type获取内容 适用用户类型列表
func (obj *_SrCouponsReceiveApplicableMgr) GetFromApplicableMemberType(applicableMemberType string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_member_type` = ?", applicableMemberType).Find(&results).Error

	return
}

// GetBatchFromApplicableMemberType 批量查找 适用用户类型列表
func (obj *_SrCouponsReceiveApplicableMgr) GetBatchFromApplicableMemberType(applicableMemberTypes []string) (results []*SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`applicable_member_type` IN (?)", applicableMemberTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrCouponsReceiveApplicableMgr) FetchByPrimaryKey(id int64) (result SrCouponsReceiveApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveApplicable{}).Where("`id` = ?", id).First(&result).Error

	return
}
