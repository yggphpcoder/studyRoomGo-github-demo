package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrCouponsReceiveMgr struct {
	*_BaseMgr
}

// SrCouponsReceiveMgr open func
func SrCouponsReceiveMgr(db *gorm.DB) *_SrCouponsReceiveMgr {
	if db == nil {
		panic(fmt.Errorf("SrCouponsReceiveMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrCouponsReceiveMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_coupons_receive"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrCouponsReceiveMgr) GetTableName() string {
	return "sr_coupons_receive"
}

// Reset 重置gorm会话
func (obj *_SrCouponsReceiveMgr) Reset() *_SrCouponsReceiveMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrCouponsReceiveMgr) Get() (result SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrCouponsReceiveMgr) Gets() (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrCouponsReceiveMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_SrCouponsReceiveMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 优惠券代码
func (obj *_SrCouponsReceiveMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithCouponID coupon_id获取 优惠券ID
func (obj *_SrCouponsReceiveMgr) WithCouponID(couponID int64) Option {
	return optionFunc(func(o *options) { o.query["coupon_id"] = couponID })
}

// WithMemberID member_id获取 会员id
func (obj *_SrCouponsReceiveMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMerchantID merchant_id获取 创建商家id
func (obj *_SrCouponsReceiveMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 创建店铺id
func (obj *_SrCouponsReceiveMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithOrderID order_id获取 订单id
func (obj *_SrCouponsReceiveMgr) WithOrderID(orderID int64) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithName name获取 优惠券名称
func (obj *_SrCouponsReceiveMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取 优惠券描述
func (obj *_SrCouponsReceiveMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithDiscountType discount_type获取 优惠券类型：百分比折扣或固定金额折扣
func (obj *_SrCouponsReceiveMgr) WithDiscountType(discountType int8) Option {
	return optionFunc(func(o *options) { o.query["discount_type"] = discountType })
}

// WithDiscountValue discount_value获取 优惠券折扣值
func (obj *_SrCouponsReceiveMgr) WithDiscountValue(discountValue float64) Option {
	return optionFunc(func(o *options) { o.query["discount_value"] = discountValue })
}

// WithDiscountRule discount_rule获取 优惠券折扣规则
func (obj *_SrCouponsReceiveMgr) WithDiscountRule(discountRule string) Option {
	return optionFunc(func(o *options) { o.query["discount_rule"] = discountRule })
}

// WithMaxDiscount max_discount获取 最大折扣金额
func (obj *_SrCouponsReceiveMgr) WithMaxDiscount(maxDiscount float64) Option {
	return optionFunc(func(o *options) { o.query["max_discount"] = maxDiscount })
}

// WithMinPurchase min_purchase获取 最低购买金额限制
func (obj *_SrCouponsReceiveMgr) WithMinPurchase(minPurchase float64) Option {
	return optionFunc(func(o *options) { o.query["min_purchase"] = minPurchase })
}

// WithStartDate start_date获取 优惠券有效开始日期
func (obj *_SrCouponsReceiveMgr) WithStartDate(startDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 优惠券有效结束日期
func (obj *_SrCouponsReceiveMgr) WithEndDate(endDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithUsageLimit usage_limit获取 优惠券使用次数限制
func (obj *_SrCouponsReceiveMgr) WithUsageLimit(usageLimit int) Option {
	return optionFunc(func(o *options) { o.query["usage_limit"] = usageLimit })
}

// WithStatus status获取 优惠券状态
func (obj *_SrCouponsReceiveMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsUsed is_used获取 是否已使用
func (obj *_SrCouponsReceiveMgr) WithIsUsed(isUsed int8) Option {
	return optionFunc(func(o *options) { o.query["is_used"] = isUsed })
}

// WithNotice notice获取 通知文案
func (obj *_SrCouponsReceiveMgr) WithNotice(notice string) Option {
	return optionFunc(func(o *options) { o.query["notice"] = notice })
}

// WithNotes notes获取 备注
func (obj *_SrCouponsReceiveMgr) WithNotes(notes string) Option {
	return optionFunc(func(o *options) { o.query["notes"] = notes })
}

// WithCreatedAt created_at获取
func (obj *_SrCouponsReceiveMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrCouponsReceiveMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrCouponsReceiveMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrCouponsReceiveMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrCouponsReceiveMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrCouponsReceiveMgr) GetByOption(opts ...Option) (result SrCouponsReceive, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrCouponsReceiveMgr) GetByOptions(opts ...Option) (results []*SrCouponsReceive, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 id
func (obj *_SrCouponsReceiveMgr) GetFromID(id int64) (result SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_SrCouponsReceiveMgr) GetBatchFromID(ids []int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 优惠券代码
func (obj *_SrCouponsReceiveMgr) GetFromCode(code string) (result SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 优惠券代码
func (obj *_SrCouponsReceiveMgr) GetBatchFromCode(codes []string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromCouponID 通过coupon_id获取内容 优惠券ID
func (obj *_SrCouponsReceiveMgr) GetFromCouponID(couponID int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`coupon_id` = ?", couponID).Find(&results).Error

	return
}

// GetBatchFromCouponID 批量查找 优惠券ID
func (obj *_SrCouponsReceiveMgr) GetBatchFromCouponID(couponIDs []int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`coupon_id` IN (?)", couponIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrCouponsReceiveMgr) GetFromMemberID(memberID int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrCouponsReceiveMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 创建商家id
func (obj *_SrCouponsReceiveMgr) GetFromMerchantID(merchantID int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 创建商家id
func (obj *_SrCouponsReceiveMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 创建店铺id
func (obj *_SrCouponsReceiveMgr) GetFromShopID(shopID int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 创建店铺id
func (obj *_SrCouponsReceiveMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单id
func (obj *_SrCouponsReceiveMgr) GetFromOrderID(orderID int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 订单id
func (obj *_SrCouponsReceiveMgr) GetBatchFromOrderID(orderIDs []int64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 优惠券名称
func (obj *_SrCouponsReceiveMgr) GetFromName(name string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 优惠券名称
func (obj *_SrCouponsReceiveMgr) GetBatchFromName(names []string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 优惠券描述
func (obj *_SrCouponsReceiveMgr) GetFromDescription(description string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 优惠券描述
func (obj *_SrCouponsReceiveMgr) GetBatchFromDescription(descriptions []string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromDiscountType 通过discount_type获取内容 优惠券类型：百分比折扣或固定金额折扣
func (obj *_SrCouponsReceiveMgr) GetFromDiscountType(discountType int8) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`discount_type` = ?", discountType).Find(&results).Error

	return
}

// GetBatchFromDiscountType 批量查找 优惠券类型：百分比折扣或固定金额折扣
func (obj *_SrCouponsReceiveMgr) GetBatchFromDiscountType(discountTypes []int8) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`discount_type` IN (?)", discountTypes).Find(&results).Error

	return
}

// GetFromDiscountValue 通过discount_value获取内容 优惠券折扣值
func (obj *_SrCouponsReceiveMgr) GetFromDiscountValue(discountValue float64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`discount_value` = ?", discountValue).Find(&results).Error

	return
}

// GetBatchFromDiscountValue 批量查找 优惠券折扣值
func (obj *_SrCouponsReceiveMgr) GetBatchFromDiscountValue(discountValues []float64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`discount_value` IN (?)", discountValues).Find(&results).Error

	return
}

// GetFromDiscountRule 通过discount_rule获取内容 优惠券折扣规则
func (obj *_SrCouponsReceiveMgr) GetFromDiscountRule(discountRule string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`discount_rule` = ?", discountRule).Find(&results).Error

	return
}

// GetBatchFromDiscountRule 批量查找 优惠券折扣规则
func (obj *_SrCouponsReceiveMgr) GetBatchFromDiscountRule(discountRules []string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`discount_rule` IN (?)", discountRules).Find(&results).Error

	return
}

// GetFromMaxDiscount 通过max_discount获取内容 最大折扣金额
func (obj *_SrCouponsReceiveMgr) GetFromMaxDiscount(maxDiscount float64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`max_discount` = ?", maxDiscount).Find(&results).Error

	return
}

// GetBatchFromMaxDiscount 批量查找 最大折扣金额
func (obj *_SrCouponsReceiveMgr) GetBatchFromMaxDiscount(maxDiscounts []float64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`max_discount` IN (?)", maxDiscounts).Find(&results).Error

	return
}

// GetFromMinPurchase 通过min_purchase获取内容 最低购买金额限制
func (obj *_SrCouponsReceiveMgr) GetFromMinPurchase(minPurchase float64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`min_purchase` = ?", minPurchase).Find(&results).Error

	return
}

// GetBatchFromMinPurchase 批量查找 最低购买金额限制
func (obj *_SrCouponsReceiveMgr) GetBatchFromMinPurchase(minPurchases []float64) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`min_purchase` IN (?)", minPurchases).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 优惠券有效开始日期
func (obj *_SrCouponsReceiveMgr) GetFromStartDate(startDate *time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 优惠券有效开始日期
func (obj *_SrCouponsReceiveMgr) GetBatchFromStartDate(startDates []*time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 优惠券有效结束日期
func (obj *_SrCouponsReceiveMgr) GetFromEndDate(endDate *time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 优惠券有效结束日期
func (obj *_SrCouponsReceiveMgr) GetBatchFromEndDate(endDates []*time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromUsageLimit 通过usage_limit获取内容 优惠券使用次数限制
func (obj *_SrCouponsReceiveMgr) GetFromUsageLimit(usageLimit int) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`usage_limit` = ?", usageLimit).Find(&results).Error

	return
}

// GetBatchFromUsageLimit 批量查找 优惠券使用次数限制
func (obj *_SrCouponsReceiveMgr) GetBatchFromUsageLimit(usageLimits []int) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`usage_limit` IN (?)", usageLimits).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 优惠券状态
func (obj *_SrCouponsReceiveMgr) GetFromStatus(status int8) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 优惠券状态
func (obj *_SrCouponsReceiveMgr) GetBatchFromStatus(statuss []int8) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsUsed 通过is_used获取内容 是否已使用
func (obj *_SrCouponsReceiveMgr) GetFromIsUsed(isUsed int8) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`is_used` = ?", isUsed).Find(&results).Error

	return
}

// GetBatchFromIsUsed 批量查找 是否已使用
func (obj *_SrCouponsReceiveMgr) GetBatchFromIsUsed(isUseds []int8) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`is_used` IN (?)", isUseds).Find(&results).Error

	return
}

// GetFromNotice 通过notice获取内容 通知文案
func (obj *_SrCouponsReceiveMgr) GetFromNotice(notice string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`notice` = ?", notice).Find(&results).Error

	return
}

// GetBatchFromNotice 批量查找 通知文案
func (obj *_SrCouponsReceiveMgr) GetBatchFromNotice(notices []string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`notice` IN (?)", notices).Find(&results).Error

	return
}

// GetFromNotes 通过notes获取内容 备注
func (obj *_SrCouponsReceiveMgr) GetFromNotes(notes string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`notes` = ?", notes).Find(&results).Error

	return
}

// GetBatchFromNotes 批量查找 备注
func (obj *_SrCouponsReceiveMgr) GetBatchFromNotes(notess []string) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`notes` IN (?)", notess).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrCouponsReceiveMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrCouponsReceiveMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrCouponsReceiveMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrCouponsReceiveMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrCouponsReceiveMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrCouponsReceiveMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrCouponsReceiveMgr) GetFromCreateBy(createBy *uint) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrCouponsReceiveMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrCouponsReceiveMgr) GetFromUpdateBy(updateBy *uint) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrCouponsReceiveMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrCouponsReceiveMgr) FetchByPrimaryKey(id int64) (result SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_SrCouponsReceiveMgr) FetchUniqueByCode(code string) (result SrCouponsReceive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsReceive{}).Where("`code` = ?", code).First(&result).Error

	return
}
