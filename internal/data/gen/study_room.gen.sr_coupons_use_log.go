package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrCouponsUseLogMgr struct {
	*_BaseMgr
}

// SrCouponsUseLogMgr open func
func SrCouponsUseLogMgr(db *gorm.DB) *_SrCouponsUseLogMgr {
	if db == nil {
		panic(fmt.Errorf("SrCouponsUseLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrCouponsUseLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_coupons_use_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrCouponsUseLogMgr) GetTableName() string {
	return "sr_coupons_use_log"
}

// Reset 重置gorm会话
func (obj *_SrCouponsUseLogMgr) Reset() *_SrCouponsUseLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrCouponsUseLogMgr) Get() (result SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrCouponsUseLogMgr) Gets() (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrCouponsUseLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_SrCouponsUseLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *_SrCouponsUseLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMerchantID merchant_id获取 创建商家id
func (obj *_SrCouponsUseLogMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 创建店铺id
func (obj *_SrCouponsUseLogMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithOrderID order_id获取 订单id
func (obj *_SrCouponsUseLogMgr) WithOrderID(orderID int64) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithCouponID coupon_id获取 优惠券ID
func (obj *_SrCouponsUseLogMgr) WithCouponID(couponID int64) Option {
	return optionFunc(func(o *options) { o.query["coupon_id"] = couponID })
}

// WithCouponsReceiveID coupons_receive_id获取 优惠券领取表ID
func (obj *_SrCouponsUseLogMgr) WithCouponsReceiveID(couponsReceiveID int64) Option {
	return optionFunc(func(o *options) { o.query["coupons_receive_id"] = couponsReceiveID })
}

// WithAction action获取 操作类型
func (obj *_SrCouponsUseLogMgr) WithAction(action int8) Option {
	return optionFunc(func(o *options) { o.query["action"] = action })
}

// WithExtra extra获取 额外信息
func (obj *_SrCouponsUseLogMgr) WithExtra(extra *string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithRemark remark获取 备注
func (obj *_SrCouponsUseLogMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrCouponsUseLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrCouponsUseLogMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrCouponsUseLogMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrCouponsUseLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrCouponsUseLogMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrCouponsUseLogMgr) GetByOption(opts ...Option) (result SrCouponsUseLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrCouponsUseLogMgr) GetByOptions(opts ...Option) (results []*SrCouponsUseLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 id
func (obj *_SrCouponsUseLogMgr) GetFromID(id int64) (result SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_SrCouponsUseLogMgr) GetBatchFromID(ids []int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrCouponsUseLogMgr) GetFromMemberID(memberID int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrCouponsUseLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 创建商家id
func (obj *_SrCouponsUseLogMgr) GetFromMerchantID(merchantID int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 创建商家id
func (obj *_SrCouponsUseLogMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 创建店铺id
func (obj *_SrCouponsUseLogMgr) GetFromShopID(shopID int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 创建店铺id
func (obj *_SrCouponsUseLogMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单id
func (obj *_SrCouponsUseLogMgr) GetFromOrderID(orderID int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 订单id
func (obj *_SrCouponsUseLogMgr) GetBatchFromOrderID(orderIDs []int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromCouponID 通过coupon_id获取内容 优惠券ID
func (obj *_SrCouponsUseLogMgr) GetFromCouponID(couponID int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`coupon_id` = ?", couponID).Find(&results).Error

	return
}

// GetBatchFromCouponID 批量查找 优惠券ID
func (obj *_SrCouponsUseLogMgr) GetBatchFromCouponID(couponIDs []int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`coupon_id` IN (?)", couponIDs).Find(&results).Error

	return
}

// GetFromCouponsReceiveID 通过coupons_receive_id获取内容 优惠券领取表ID
func (obj *_SrCouponsUseLogMgr) GetFromCouponsReceiveID(couponsReceiveID int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`coupons_receive_id` = ?", couponsReceiveID).Find(&results).Error

	return
}

// GetBatchFromCouponsReceiveID 批量查找 优惠券领取表ID
func (obj *_SrCouponsUseLogMgr) GetBatchFromCouponsReceiveID(couponsReceiveIDs []int64) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`coupons_receive_id` IN (?)", couponsReceiveIDs).Find(&results).Error

	return
}

// GetFromAction 通过action获取内容 操作类型
func (obj *_SrCouponsUseLogMgr) GetFromAction(action int8) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`action` = ?", action).Find(&results).Error

	return
}

// GetBatchFromAction 批量查找 操作类型
func (obj *_SrCouponsUseLogMgr) GetBatchFromAction(actions []int8) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`action` IN (?)", actions).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 额外信息
func (obj *_SrCouponsUseLogMgr) GetFromExtra(extra *string) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 额外信息
func (obj *_SrCouponsUseLogMgr) GetBatchFromExtra(extras []*string) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrCouponsUseLogMgr) GetFromRemark(remark *string) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrCouponsUseLogMgr) GetBatchFromRemark(remarks []*string) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrCouponsUseLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrCouponsUseLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrCouponsUseLogMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrCouponsUseLogMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrCouponsUseLogMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrCouponsUseLogMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrCouponsUseLogMgr) GetFromCreateBy(createBy *uint) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrCouponsUseLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrCouponsUseLogMgr) GetFromUpdateBy(updateBy *uint) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrCouponsUseLogMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrCouponsUseLogMgr) FetchByPrimaryKey(id int64) (result SrCouponsUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsUseLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
