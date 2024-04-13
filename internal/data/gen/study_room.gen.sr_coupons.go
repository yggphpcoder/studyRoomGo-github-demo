package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrCouponsMgr struct {
	*_BaseMgr
}

// SrCouponsMgr open func
func SrCouponsMgr(db *gorm.DB) *_SrCouponsMgr {
	if db == nil {
		panic(fmt.Errorf("SrCouponsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrCouponsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_coupons"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrCouponsMgr) GetTableName() string {
	return "sr_coupons"
}

// Reset 重置gorm会话
func (obj *_SrCouponsMgr) Reset() *_SrCouponsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrCouponsMgr) Get() (result SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrCouponsMgr) Gets() (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrCouponsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 优惠券ID
func (obj *_SrCouponsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 创建商家id
func (obj *_SrCouponsMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 创建店铺id
func (obj *_SrCouponsMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithCode code获取 优惠券代码
func (obj *_SrCouponsMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithType type获取 1:即用型,2:领取型
func (obj *_SrCouponsMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithName name获取 优惠券名称
func (obj *_SrCouponsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取 优惠券描述
func (obj *_SrCouponsMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithDiscountType discount_type获取 优惠券类型：百分比折扣或固定金额折扣
func (obj *_SrCouponsMgr) WithDiscountType(discountType int8) Option {
	return optionFunc(func(o *options) { o.query["discount_type"] = discountType })
}

// WithDiscountValue discount_value获取 优惠券折扣值
func (obj *_SrCouponsMgr) WithDiscountValue(discountValue float64) Option {
	return optionFunc(func(o *options) { o.query["discount_value"] = discountValue })
}

// WithDiscountRule discount_rule获取 优惠券折扣规则
func (obj *_SrCouponsMgr) WithDiscountRule(discountRule string) Option {
	return optionFunc(func(o *options) { o.query["discount_rule"] = discountRule })
}

// WithMaxDiscount max_discount获取 最大折扣金额
func (obj *_SrCouponsMgr) WithMaxDiscount(maxDiscount float64) Option {
	return optionFunc(func(o *options) { o.query["max_discount"] = maxDiscount })
}

// WithMinPurchase min_purchase获取 最低购买金额限制
func (obj *_SrCouponsMgr) WithMinPurchase(minPurchase float64) Option {
	return optionFunc(func(o *options) { o.query["min_purchase"] = minPurchase })
}

// WithStartDate start_date获取 优惠券有效开始日期
func (obj *_SrCouponsMgr) WithStartDate(startDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 优惠券有效结束日期
func (obj *_SrCouponsMgr) WithEndDate(endDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithValidDay valid_day获取 领取后有效天数
func (obj *_SrCouponsMgr) WithValidDay(validDay int) Option {
	return optionFunc(func(o *options) { o.query["valid_day"] = validDay })
}

// WithStatus status获取 优惠券状态
func (obj *_SrCouponsMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithQuantity quantity获取 优惠券数量
func (obj *_SrCouponsMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithUsageLimitPerUser usage_limit_per_user获取 每个用户可使用的优惠券数量限制
func (obj *_SrCouponsMgr) WithUsageLimitPerUser(usageLimitPerUser int) Option {
	return optionFunc(func(o *options) { o.query["usage_limit_per_user"] = usageLimitPerUser })
}

// WithUsageLimit usage_limit获取 优惠券总使用次数限制
func (obj *_SrCouponsMgr) WithUsageLimit(usageLimit int) Option {
	return optionFunc(func(o *options) { o.query["usage_limit"] = usageLimit })
}

// WithIsReusable is_reusable获取 是否可重复使用
func (obj *_SrCouponsMgr) WithIsReusable(isReusable int8) Option {
	return optionFunc(func(o *options) { o.query["is_reusable"] = isReusable })
}

// WithIsShow is_show获取 是否在前台显示
func (obj *_SrCouponsMgr) WithIsShow(isShow int8) Option {
	return optionFunc(func(o *options) { o.query["is_show"] = isShow })
}

// WithNotes notes获取 备注
func (obj *_SrCouponsMgr) WithNotes(notes string) Option {
	return optionFunc(func(o *options) { o.query["notes"] = notes })
}

// WithCreatedAt created_at获取
func (obj *_SrCouponsMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrCouponsMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrCouponsMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrCouponsMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrCouponsMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrCouponsMgr) GetByOption(opts ...Option) (result SrCoupons, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrCouponsMgr) GetByOptions(opts ...Option) (results []*SrCoupons, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 优惠券ID
func (obj *_SrCouponsMgr) GetFromID(id int64) (result SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 优惠券ID
func (obj *_SrCouponsMgr) GetBatchFromID(ids []int64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 创建商家id
func (obj *_SrCouponsMgr) GetFromMerchantID(merchantID int64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 创建商家id
func (obj *_SrCouponsMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 创建店铺id
func (obj *_SrCouponsMgr) GetFromShopID(shopID int64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 创建店铺id
func (obj *_SrCouponsMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 优惠券代码
func (obj *_SrCouponsMgr) GetFromCode(code string) (result SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 优惠券代码
func (obj *_SrCouponsMgr) GetBatchFromCode(codes []string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 1:即用型,2:领取型
func (obj *_SrCouponsMgr) GetFromType(_type int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 1:即用型,2:领取型
func (obj *_SrCouponsMgr) GetBatchFromType(_types []int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 优惠券名称
func (obj *_SrCouponsMgr) GetFromName(name string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 优惠券名称
func (obj *_SrCouponsMgr) GetBatchFromName(names []string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 优惠券描述
func (obj *_SrCouponsMgr) GetFromDescription(description string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 优惠券描述
func (obj *_SrCouponsMgr) GetBatchFromDescription(descriptions []string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromDiscountType 通过discount_type获取内容 优惠券类型：百分比折扣或固定金额折扣
func (obj *_SrCouponsMgr) GetFromDiscountType(discountType int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`discount_type` = ?", discountType).Find(&results).Error

	return
}

// GetBatchFromDiscountType 批量查找 优惠券类型：百分比折扣或固定金额折扣
func (obj *_SrCouponsMgr) GetBatchFromDiscountType(discountTypes []int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`discount_type` IN (?)", discountTypes).Find(&results).Error

	return
}

// GetFromDiscountValue 通过discount_value获取内容 优惠券折扣值
func (obj *_SrCouponsMgr) GetFromDiscountValue(discountValue float64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`discount_value` = ?", discountValue).Find(&results).Error

	return
}

// GetBatchFromDiscountValue 批量查找 优惠券折扣值
func (obj *_SrCouponsMgr) GetBatchFromDiscountValue(discountValues []float64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`discount_value` IN (?)", discountValues).Find(&results).Error

	return
}

// GetFromDiscountRule 通过discount_rule获取内容 优惠券折扣规则
func (obj *_SrCouponsMgr) GetFromDiscountRule(discountRule string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`discount_rule` = ?", discountRule).Find(&results).Error

	return
}

// GetBatchFromDiscountRule 批量查找 优惠券折扣规则
func (obj *_SrCouponsMgr) GetBatchFromDiscountRule(discountRules []string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`discount_rule` IN (?)", discountRules).Find(&results).Error

	return
}

// GetFromMaxDiscount 通过max_discount获取内容 最大折扣金额
func (obj *_SrCouponsMgr) GetFromMaxDiscount(maxDiscount float64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`max_discount` = ?", maxDiscount).Find(&results).Error

	return
}

// GetBatchFromMaxDiscount 批量查找 最大折扣金额
func (obj *_SrCouponsMgr) GetBatchFromMaxDiscount(maxDiscounts []float64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`max_discount` IN (?)", maxDiscounts).Find(&results).Error

	return
}

// GetFromMinPurchase 通过min_purchase获取内容 最低购买金额限制
func (obj *_SrCouponsMgr) GetFromMinPurchase(minPurchase float64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`min_purchase` = ?", minPurchase).Find(&results).Error

	return
}

// GetBatchFromMinPurchase 批量查找 最低购买金额限制
func (obj *_SrCouponsMgr) GetBatchFromMinPurchase(minPurchases []float64) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`min_purchase` IN (?)", minPurchases).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 优惠券有效开始日期
func (obj *_SrCouponsMgr) GetFromStartDate(startDate *time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 优惠券有效开始日期
func (obj *_SrCouponsMgr) GetBatchFromStartDate(startDates []*time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 优惠券有效结束日期
func (obj *_SrCouponsMgr) GetFromEndDate(endDate *time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 优惠券有效结束日期
func (obj *_SrCouponsMgr) GetBatchFromEndDate(endDates []*time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromValidDay 通过valid_day获取内容 领取后有效天数
func (obj *_SrCouponsMgr) GetFromValidDay(validDay int) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`valid_day` = ?", validDay).Find(&results).Error

	return
}

// GetBatchFromValidDay 批量查找 领取后有效天数
func (obj *_SrCouponsMgr) GetBatchFromValidDay(validDays []int) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`valid_day` IN (?)", validDays).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 优惠券状态
func (obj *_SrCouponsMgr) GetFromStatus(status int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 优惠券状态
func (obj *_SrCouponsMgr) GetBatchFromStatus(statuss []int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 优惠券数量
func (obj *_SrCouponsMgr) GetFromQuantity(quantity int) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 优惠券数量
func (obj *_SrCouponsMgr) GetBatchFromQuantity(quantitys []int) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromUsageLimitPerUser 通过usage_limit_per_user获取内容 每个用户可使用的优惠券数量限制
func (obj *_SrCouponsMgr) GetFromUsageLimitPerUser(usageLimitPerUser int) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`usage_limit_per_user` = ?", usageLimitPerUser).Find(&results).Error

	return
}

// GetBatchFromUsageLimitPerUser 批量查找 每个用户可使用的优惠券数量限制
func (obj *_SrCouponsMgr) GetBatchFromUsageLimitPerUser(usageLimitPerUsers []int) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`usage_limit_per_user` IN (?)", usageLimitPerUsers).Find(&results).Error

	return
}

// GetFromUsageLimit 通过usage_limit获取内容 优惠券总使用次数限制
func (obj *_SrCouponsMgr) GetFromUsageLimit(usageLimit int) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`usage_limit` = ?", usageLimit).Find(&results).Error

	return
}

// GetBatchFromUsageLimit 批量查找 优惠券总使用次数限制
func (obj *_SrCouponsMgr) GetBatchFromUsageLimit(usageLimits []int) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`usage_limit` IN (?)", usageLimits).Find(&results).Error

	return
}

// GetFromIsReusable 通过is_reusable获取内容 是否可重复使用
func (obj *_SrCouponsMgr) GetFromIsReusable(isReusable int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`is_reusable` = ?", isReusable).Find(&results).Error

	return
}

// GetBatchFromIsReusable 批量查找 是否可重复使用
func (obj *_SrCouponsMgr) GetBatchFromIsReusable(isReusables []int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`is_reusable` IN (?)", isReusables).Find(&results).Error

	return
}

// GetFromIsShow 通过is_show获取内容 是否在前台显示
func (obj *_SrCouponsMgr) GetFromIsShow(isShow int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`is_show` = ?", isShow).Find(&results).Error

	return
}

// GetBatchFromIsShow 批量查找 是否在前台显示
func (obj *_SrCouponsMgr) GetBatchFromIsShow(isShows []int8) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`is_show` IN (?)", isShows).Find(&results).Error

	return
}

// GetFromNotes 通过notes获取内容 备注
func (obj *_SrCouponsMgr) GetFromNotes(notes string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`notes` = ?", notes).Find(&results).Error

	return
}

// GetBatchFromNotes 批量查找 备注
func (obj *_SrCouponsMgr) GetBatchFromNotes(notess []string) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`notes` IN (?)", notess).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrCouponsMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrCouponsMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrCouponsMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrCouponsMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrCouponsMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrCouponsMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrCouponsMgr) GetFromCreateBy(createBy *uint) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrCouponsMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrCouponsMgr) GetFromUpdateBy(updateBy *uint) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrCouponsMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrCouponsMgr) FetchByPrimaryKey(id int64) (result SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_SrCouponsMgr) FetchUniqueByCode(code string) (result SrCoupons, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCoupons{}).Where("`code` = ?", code).First(&result).Error

	return
}
