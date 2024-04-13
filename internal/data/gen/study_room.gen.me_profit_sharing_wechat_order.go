package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MeProfitSharingWechatOrderMgr struct {
	*_BaseMgr
}

// MeProfitSharingWechatOrderMgr open func
func MeProfitSharingWechatOrderMgr(db *gorm.DB) *_MeProfitSharingWechatOrderMgr {
	if db == nil {
		panic(fmt.Errorf("MeProfitSharingWechatOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MeProfitSharingWechatOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("me_profit_sharing_wechat_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MeProfitSharingWechatOrderMgr) GetTableName() string {
	return "me_profit_sharing_wechat_order"
}

// Reset 重置gorm会话
func (obj *_MeProfitSharingWechatOrderMgr) Reset() *_MeProfitSharingWechatOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MeProfitSharingWechatOrderMgr) Get() (result MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MeProfitSharingWechatOrderMgr) Gets() (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MeProfitSharingWechatOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_MeProfitSharingWechatOrderMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNo order_no获取 分账单编号
func (obj *_MeProfitSharingWechatOrderMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithTransactionID transaction_id获取 交易号
func (obj *_MeProfitSharingWechatOrderMgr) WithTransactionID(transactionID string) Option {
	return optionFunc(func(o *options) { o.query["transaction_id"] = transactionID })
}

// WithWechatOrderID wechat_order_id获取 微信分账单号
func (obj *_MeProfitSharingWechatOrderMgr) WithWechatOrderID(wechatOrderID string) Option {
	return optionFunc(func(o *options) { o.query["wechat_order_id"] = wechatOrderID })
}

// WithOrderDesc order_desc获取 订单描述
func (obj *_MeProfitSharingWechatOrderMgr) WithOrderDesc(orderDesc string) Option {
	return optionFunc(func(o *options) { o.query["order_desc"] = orderDesc })
}

// WithReceiversType receivers_type获取 分账接收方类型:MERCHANT_ID,PERSONAL_OPENID
func (obj *_MeProfitSharingWechatOrderMgr) WithReceiversType(receiversType string) Option {
	return optionFunc(func(o *options) { o.query["receivers_type"] = receiversType })
}

// WithReceiversAccount receivers_account获取 分账接收方账号
func (obj *_MeProfitSharingWechatOrderMgr) WithReceiversAccount(receiversAccount string) Option {
	return optionFunc(func(o *options) { o.query["receivers_account"] = receiversAccount })
}

// WithReceiversName receivers_name获取 接受方名字
func (obj *_MeProfitSharingWechatOrderMgr) WithReceiversName(receiversName string) Option {
	return optionFunc(func(o *options) { o.query["receivers_name"] = receiversName })
}

// WithMerchantID merchant_id获取 商户id
func (obj *_MeProfitSharingWechatOrderMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_MeProfitSharingWechatOrderMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithRelateOrderType relate_order_type获取 关联订单类型 1:订座单
func (obj *_MeProfitSharingWechatOrderMgr) WithRelateOrderType(relateOrderType int8) Option {
	return optionFunc(func(o *options) { o.query["relate_order_type"] = relateOrderType })
}

// WithRelateOrderNo relate_order_no获取 分账关联单号
func (obj *_MeProfitSharingWechatOrderMgr) WithRelateOrderNo(relateOrderNo string) Option {
	return optionFunc(func(o *options) { o.query["relate_order_no"] = relateOrderNo })
}

// WithStatus status获取 分账订单状态
func (obj *_MeProfitSharingWechatOrderMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithUnfreezeUnsplit unfreeze_unsplit获取 是否解冻剩余未分资金
func (obj *_MeProfitSharingWechatOrderMgr) WithUnfreezeUnsplit(unfreezeUnsplit int8) Option {
	return optionFunc(func(o *options) { o.query["unfreeze_unsplit"] = unfreezeUnsplit })
}

// WithTotalAmount total_amount获取 订单总金额
func (obj *_MeProfitSharingWechatOrderMgr) WithTotalAmount(totalAmount int) Option {
	return optionFunc(func(o *options) { o.query["total_amount"] = totalAmount })
}

// WithRatio ratio获取 分账比例
func (obj *_MeProfitSharingWechatOrderMgr) WithRatio(ratio int8) Option {
	return optionFunc(func(o *options) { o.query["ratio"] = ratio })
}

// WithAmount amount获取 分账金额
func (obj *_MeProfitSharingWechatOrderMgr) WithAmount(amount int) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithCreatedAt created_at获取
func (obj *_MeProfitSharingWechatOrderMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_MeProfitSharingWechatOrderMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_MeProfitSharingWechatOrderMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_MeProfitSharingWechatOrderMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_MeProfitSharingWechatOrderMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_MeProfitSharingWechatOrderMgr) GetByOption(opts ...Option) (result MeProfitSharingWechatOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MeProfitSharingWechatOrderMgr) GetByOptions(opts ...Option) (results []*MeProfitSharingWechatOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_MeProfitSharingWechatOrderMgr) GetFromID(id int64) (result MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromID(ids []int64) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 分账单编号
func (obj *_MeProfitSharingWechatOrderMgr) GetFromOrderNo(orderNo string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`order_no` = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量查找 分账单编号
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromOrderNo(orderNos []string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromTransactionID 通过transaction_id获取内容 交易号
func (obj *_MeProfitSharingWechatOrderMgr) GetFromTransactionID(transactionID string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`transaction_id` = ?", transactionID).Find(&results).Error

	return
}

// GetBatchFromTransactionID 批量查找 交易号
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromTransactionID(transactionIDs []string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`transaction_id` IN (?)", transactionIDs).Find(&results).Error

	return
}

// GetFromWechatOrderID 通过wechat_order_id获取内容 微信分账单号
func (obj *_MeProfitSharingWechatOrderMgr) GetFromWechatOrderID(wechatOrderID string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`wechat_order_id` = ?", wechatOrderID).Find(&results).Error

	return
}

// GetBatchFromWechatOrderID 批量查找 微信分账单号
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromWechatOrderID(wechatOrderIDs []string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`wechat_order_id` IN (?)", wechatOrderIDs).Find(&results).Error

	return
}

// GetFromOrderDesc 通过order_desc获取内容 订单描述
func (obj *_MeProfitSharingWechatOrderMgr) GetFromOrderDesc(orderDesc string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`order_desc` = ?", orderDesc).Find(&results).Error

	return
}

// GetBatchFromOrderDesc 批量查找 订单描述
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromOrderDesc(orderDescs []string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`order_desc` IN (?)", orderDescs).Find(&results).Error

	return
}

// GetFromReceiversType 通过receivers_type获取内容 分账接收方类型:MERCHANT_ID,PERSONAL_OPENID
func (obj *_MeProfitSharingWechatOrderMgr) GetFromReceiversType(receiversType string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`receivers_type` = ?", receiversType).Find(&results).Error

	return
}

// GetBatchFromReceiversType 批量查找 分账接收方类型:MERCHANT_ID,PERSONAL_OPENID
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromReceiversType(receiversTypes []string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`receivers_type` IN (?)", receiversTypes).Find(&results).Error

	return
}

// GetFromReceiversAccount 通过receivers_account获取内容 分账接收方账号
func (obj *_MeProfitSharingWechatOrderMgr) GetFromReceiversAccount(receiversAccount string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`receivers_account` = ?", receiversAccount).Find(&results).Error

	return
}

// GetBatchFromReceiversAccount 批量查找 分账接收方账号
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromReceiversAccount(receiversAccounts []string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`receivers_account` IN (?)", receiversAccounts).Find(&results).Error

	return
}

// GetFromReceiversName 通过receivers_name获取内容 接受方名字
func (obj *_MeProfitSharingWechatOrderMgr) GetFromReceiversName(receiversName string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`receivers_name` = ?", receiversName).Find(&results).Error

	return
}

// GetBatchFromReceiversName 批量查找 接受方名字
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromReceiversName(receiversNames []string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`receivers_name` IN (?)", receiversNames).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商户id
func (obj *_MeProfitSharingWechatOrderMgr) GetFromMerchantID(merchantID int64) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商户id
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_MeProfitSharingWechatOrderMgr) GetFromShopID(shopID int64) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromShopID(shopIDs []int64) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromRelateOrderType 通过relate_order_type获取内容 关联订单类型 1:订座单
func (obj *_MeProfitSharingWechatOrderMgr) GetFromRelateOrderType(relateOrderType int8) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`relate_order_type` = ?", relateOrderType).Find(&results).Error

	return
}

// GetBatchFromRelateOrderType 批量查找 关联订单类型 1:订座单
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromRelateOrderType(relateOrderTypes []int8) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`relate_order_type` IN (?)", relateOrderTypes).Find(&results).Error

	return
}

// GetFromRelateOrderNo 通过relate_order_no获取内容 分账关联单号
func (obj *_MeProfitSharingWechatOrderMgr) GetFromRelateOrderNo(relateOrderNo string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`relate_order_no` = ?", relateOrderNo).Find(&results).Error

	return
}

// GetBatchFromRelateOrderNo 批量查找 分账关联单号
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromRelateOrderNo(relateOrderNos []string) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`relate_order_no` IN (?)", relateOrderNos).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 分账订单状态
func (obj *_MeProfitSharingWechatOrderMgr) GetFromStatus(status int8) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 分账订单状态
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromStatus(statuss []int8) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromUnfreezeUnsplit 通过unfreeze_unsplit获取内容 是否解冻剩余未分资金
func (obj *_MeProfitSharingWechatOrderMgr) GetFromUnfreezeUnsplit(unfreezeUnsplit int8) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`unfreeze_unsplit` = ?", unfreezeUnsplit).Find(&results).Error

	return
}

// GetBatchFromUnfreezeUnsplit 批量查找 是否解冻剩余未分资金
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromUnfreezeUnsplit(unfreezeUnsplits []int8) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`unfreeze_unsplit` IN (?)", unfreezeUnsplits).Find(&results).Error

	return
}

// GetFromTotalAmount 通过total_amount获取内容 订单总金额
func (obj *_MeProfitSharingWechatOrderMgr) GetFromTotalAmount(totalAmount int) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`total_amount` = ?", totalAmount).Find(&results).Error

	return
}

// GetBatchFromTotalAmount 批量查找 订单总金额
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromTotalAmount(totalAmounts []int) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`total_amount` IN (?)", totalAmounts).Find(&results).Error

	return
}

// GetFromRatio 通过ratio获取内容 分账比例
func (obj *_MeProfitSharingWechatOrderMgr) GetFromRatio(ratio int8) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`ratio` = ?", ratio).Find(&results).Error

	return
}

// GetBatchFromRatio 批量查找 分账比例
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromRatio(ratios []int8) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`ratio` IN (?)", ratios).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 分账金额
func (obj *_MeProfitSharingWechatOrderMgr) GetFromAmount(amount int) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 分账金额
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromAmount(amounts []int) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_MeProfitSharingWechatOrderMgr) GetFromCreatedAt(createdAt *time.Time) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_MeProfitSharingWechatOrderMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_MeProfitSharingWechatOrderMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_MeProfitSharingWechatOrderMgr) GetFromCreateBy(createBy *uint) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromCreateBy(createBys []*uint) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_MeProfitSharingWechatOrderMgr) GetFromUpdateBy(updateBy *uint) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_MeProfitSharingWechatOrderMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MeProfitSharingWechatOrderMgr) FetchByPrimaryKey(id int64) (result MeProfitSharingWechatOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}
