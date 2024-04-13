package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SrCouponsReceiveGiftMgr struct {
	*_BaseMgr
}

// SrCouponsReceiveGiftMgr open func
func SrCouponsReceiveGiftMgr(db *gorm.DB) *_SrCouponsReceiveGiftMgr {
	if db == nil {
		panic(fmt.Errorf("SrCouponsReceiveGiftMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrCouponsReceiveGiftMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_coupons_receive_gift"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrCouponsReceiveGiftMgr) GetTableName() string {
	return "sr_coupons_receive_gift"
}

// Reset 重置gorm会话
func (obj *_SrCouponsReceiveGiftMgr) Reset() *_SrCouponsReceiveGiftMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrCouponsReceiveGiftMgr) Get() (result SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrCouponsReceiveGiftMgr) Gets() (results []*SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrCouponsReceiveGiftMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_SrCouponsReceiveGiftMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCouponsReceiveID coupons_receive_id获取 优惠券领取表ID
func (obj *_SrCouponsReceiveGiftMgr) WithCouponsReceiveID(couponsReceiveID int64) Option {
	return optionFunc(func(o *options) { o.query["coupons_receive_id"] = couponsReceiveID })
}

// WithGiftPackageCard gift_package_card获取 赠送套餐
func (obj *_SrCouponsReceiveGiftMgr) WithGiftPackageCard(giftPackageCard string) Option {
	return optionFunc(func(o *options) { o.query["gift_package_card"] = giftPackageCard })
}

// WithGiftCardPercent gift_card_percent获取 赠送时长%
func (obj *_SrCouponsReceiveGiftMgr) WithGiftCardPercent(giftCardPercent int16) Option {
	return optionFunc(func(o *options) { o.query["gift_card_percent"] = giftCardPercent })
}

// GetByOption 功能选项模式获取
func (obj *_SrCouponsReceiveGiftMgr) GetByOption(opts ...Option) (result SrCouponsReceiveGift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrCouponsReceiveGiftMgr) GetByOptions(opts ...Option) (results []*SrCouponsReceiveGift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 id
func (obj *_SrCouponsReceiveGiftMgr) GetFromID(id int64) (result SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_SrCouponsReceiveGiftMgr) GetBatchFromID(ids []int64) (results []*SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCouponsReceiveID 通过coupons_receive_id获取内容 优惠券领取表ID
func (obj *_SrCouponsReceiveGiftMgr) GetFromCouponsReceiveID(couponsReceiveID int64) (results []*SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`coupons_receive_id` = ?", couponsReceiveID).Find(&results).Error

	return
}

// GetBatchFromCouponsReceiveID 批量查找 优惠券领取表ID
func (obj *_SrCouponsReceiveGiftMgr) GetBatchFromCouponsReceiveID(couponsReceiveIDs []int64) (results []*SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`coupons_receive_id` IN (?)", couponsReceiveIDs).Find(&results).Error

	return
}

// GetFromGiftPackageCard 通过gift_package_card获取内容 赠送套餐
func (obj *_SrCouponsReceiveGiftMgr) GetFromGiftPackageCard(giftPackageCard string) (results []*SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`gift_package_card` = ?", giftPackageCard).Find(&results).Error

	return
}

// GetBatchFromGiftPackageCard 批量查找 赠送套餐
func (obj *_SrCouponsReceiveGiftMgr) GetBatchFromGiftPackageCard(giftPackageCards []string) (results []*SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`gift_package_card` IN (?)", giftPackageCards).Find(&results).Error

	return
}

// GetFromGiftCardPercent 通过gift_card_percent获取内容 赠送时长%
func (obj *_SrCouponsReceiveGiftMgr) GetFromGiftCardPercent(giftCardPercent int16) (results []*SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`gift_card_percent` = ?", giftCardPercent).Find(&results).Error

	return
}

// GetBatchFromGiftCardPercent 批量查找 赠送时长%
func (obj *_SrCouponsReceiveGiftMgr) GetBatchFromGiftCardPercent(giftCardPercents []int16) (results []*SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`gift_card_percent` IN (?)", giftCardPercents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrCouponsReceiveGiftMgr) FetchByPrimaryKey(id int64) (result SrCouponsReceiveGift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceiveGift{}).Where("`id` = ?", id).First(&result).Error

	return
}
