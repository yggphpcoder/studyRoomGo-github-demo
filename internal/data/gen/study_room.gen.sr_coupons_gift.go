package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SrCouponsGiftMgr struct {
	*_BaseMgr
}

// SrCouponsGiftMgr open func
func SrCouponsGiftMgr(db *gorm.DB) *_SrCouponsGiftMgr {
	if db == nil {
		panic(fmt.Errorf("SrCouponsGiftMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrCouponsGiftMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_coupons_gift"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrCouponsGiftMgr) GetTableName() string {
	return "sr_coupons_gift"
}

// Reset 重置gorm会话
func (obj *_SrCouponsGiftMgr) Reset() *_SrCouponsGiftMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrCouponsGiftMgr) Get() (result SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrCouponsGiftMgr) Gets() (results []*SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrCouponsGiftMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCouponID coupon_id获取 优惠券ID
func (obj *_SrCouponsGiftMgr) WithCouponID(couponID int64) Option {
	return optionFunc(func(o *options) { o.query["coupon_id"] = couponID })
}

// WithGiftPackageCard gift_package_card获取 赠送套餐
func (obj *_SrCouponsGiftMgr) WithGiftPackageCard(giftPackageCard string) Option {
	return optionFunc(func(o *options) { o.query["gift_package_card"] = giftPackageCard })
}

// WithGiftCardPercent gift_card_percent获取 赠送时长%
func (obj *_SrCouponsGiftMgr) WithGiftCardPercent(giftCardPercent int16) Option {
	return optionFunc(func(o *options) { o.query["gift_card_percent"] = giftCardPercent })
}

// GetByOption 功能选项模式获取
func (obj *_SrCouponsGiftMgr) GetByOption(opts ...Option) (result SrCouponsGift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrCouponsGiftMgr) GetByOptions(opts ...Option) (results []*SrCouponsGift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromCouponID 通过coupon_id获取内容 优惠券ID
func (obj *_SrCouponsGiftMgr) GetFromCouponID(couponID int64) (result SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where("`coupon_id` = ?", couponID).First(&result).Error

	return
}

// GetBatchFromCouponID 批量查找 优惠券ID
func (obj *_SrCouponsGiftMgr) GetBatchFromCouponID(couponIDs []int64) (results []*SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where("`coupon_id` IN (?)", couponIDs).Find(&results).Error

	return
}

// GetFromGiftPackageCard 通过gift_package_card获取内容 赠送套餐
func (obj *_SrCouponsGiftMgr) GetFromGiftPackageCard(giftPackageCard string) (results []*SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where("`gift_package_card` = ?", giftPackageCard).Find(&results).Error

	return
}

// GetBatchFromGiftPackageCard 批量查找 赠送套餐
func (obj *_SrCouponsGiftMgr) GetBatchFromGiftPackageCard(giftPackageCards []string) (results []*SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where("`gift_package_card` IN (?)", giftPackageCards).Find(&results).Error

	return
}

// GetFromGiftCardPercent 通过gift_card_percent获取内容 赠送时长%
func (obj *_SrCouponsGiftMgr) GetFromGiftCardPercent(giftCardPercent int16) (results []*SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where("`gift_card_percent` = ?", giftCardPercent).Find(&results).Error

	return
}

// GetBatchFromGiftCardPercent 批量查找 赠送时长%
func (obj *_SrCouponsGiftMgr) GetBatchFromGiftCardPercent(giftCardPercents []int16) (results []*SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where("`gift_card_percent` IN (?)", giftCardPercents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrCouponsGiftMgr) FetchByPrimaryKey(couponID int64) (result SrCouponsGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsGift{}).Where("`coupon_id` = ?", couponID).First(&result).Error

	return
}
