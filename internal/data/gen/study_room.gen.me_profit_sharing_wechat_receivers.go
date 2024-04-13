package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MeProfitSharingWechatReceiversMgr struct {
	*_BaseMgr
}

// MeProfitSharingWechatReceiversMgr open func
func MeProfitSharingWechatReceiversMgr(db *gorm.DB) *_MeProfitSharingWechatReceiversMgr {
	if db == nil {
		panic(fmt.Errorf("MeProfitSharingWechatReceiversMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MeProfitSharingWechatReceiversMgr{_BaseMgr: &_BaseMgr{DB: db.Table("me_profit_sharing_wechat_receivers"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MeProfitSharingWechatReceiversMgr) GetTableName() string {
	return "me_profit_sharing_wechat_receivers"
}

// Reset 重置gorm会话
func (obj *_MeProfitSharingWechatReceiversMgr) Reset() *_MeProfitSharingWechatReceiversMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MeProfitSharingWechatReceiversMgr) Get() (result MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MeProfitSharingWechatReceiversMgr) Gets() (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MeProfitSharingWechatReceiversMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_MeProfitSharingWechatReceiversMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商户id
func (obj *_MeProfitSharingWechatReceiversMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_MeProfitSharingWechatReceiversMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithReceiversType receivers_type获取 分账接收方类型:MERCHANT_ID,PERSONAL_OPENID
func (obj *_MeProfitSharingWechatReceiversMgr) WithReceiversType(receiversType string) Option {
	return optionFunc(func(o *options) { o.query["receivers_type"] = receiversType })
}

// WithReceiversAccount receivers_account获取 分账接收方账号
func (obj *_MeProfitSharingWechatReceiversMgr) WithReceiversAccount(receiversAccount string) Option {
	return optionFunc(func(o *options) { o.query["receivers_account"] = receiversAccount })
}

// WithReceiversName receivers_name获取 接受方名字
func (obj *_MeProfitSharingWechatReceiversMgr) WithReceiversName(receiversName string) Option {
	return optionFunc(func(o *options) { o.query["receivers_name"] = receiversName })
}

// WithRelateOrderType relate_order_type获取 关联订单类型 1:订座单
func (obj *_MeProfitSharingWechatReceiversMgr) WithRelateOrderType(relateOrderType int8) Option {
	return optionFunc(func(o *options) { o.query["relate_order_type"] = relateOrderType })
}

// WithRatio ratio获取 分账比例
func (obj *_MeProfitSharingWechatReceiversMgr) WithRatio(ratio int8) Option {
	return optionFunc(func(o *options) { o.query["ratio"] = ratio })
}

// WithStatus status获取 状态
func (obj *_MeProfitSharingWechatReceiversMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_MeProfitSharingWechatReceiversMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_MeProfitSharingWechatReceiversMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_MeProfitSharingWechatReceiversMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_MeProfitSharingWechatReceiversMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_MeProfitSharingWechatReceiversMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithRemark remark获取 备注
func (obj *_MeProfitSharingWechatReceiversMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// GetByOption 功能选项模式获取
func (obj *_MeProfitSharingWechatReceiversMgr) GetByOption(opts ...Option) (result MeProfitSharingWechatReceivers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MeProfitSharingWechatReceiversMgr) GetByOptions(opts ...Option) (results []*MeProfitSharingWechatReceivers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromID(id int64) (result MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromID(ids []int64) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商户id
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromMerchantID(merchantID int64) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商户id
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromShopID(shopID int64) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromShopID(shopIDs []int64) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromReceiversType 通过receivers_type获取内容 分账接收方类型:MERCHANT_ID,PERSONAL_OPENID
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromReceiversType(receiversType string) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`receivers_type` = ?", receiversType).Find(&results).Error

	return
}

// GetBatchFromReceiversType 批量查找 分账接收方类型:MERCHANT_ID,PERSONAL_OPENID
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromReceiversType(receiversTypes []string) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`receivers_type` IN (?)", receiversTypes).Find(&results).Error

	return
}

// GetFromReceiversAccount 通过receivers_account获取内容 分账接收方账号
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromReceiversAccount(receiversAccount string) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`receivers_account` = ?", receiversAccount).Find(&results).Error

	return
}

// GetBatchFromReceiversAccount 批量查找 分账接收方账号
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromReceiversAccount(receiversAccounts []string) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`receivers_account` IN (?)", receiversAccounts).Find(&results).Error

	return
}

// GetFromReceiversName 通过receivers_name获取内容 接受方名字
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromReceiversName(receiversName string) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`receivers_name` = ?", receiversName).Find(&results).Error

	return
}

// GetBatchFromReceiversName 批量查找 接受方名字
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromReceiversName(receiversNames []string) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`receivers_name` IN (?)", receiversNames).Find(&results).Error

	return
}

// GetFromRelateOrderType 通过relate_order_type获取内容 关联订单类型 1:订座单
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromRelateOrderType(relateOrderType int8) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`relate_order_type` = ?", relateOrderType).Find(&results).Error

	return
}

// GetBatchFromRelateOrderType 批量查找 关联订单类型 1:订座单
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromRelateOrderType(relateOrderTypes []int8) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`relate_order_type` IN (?)", relateOrderTypes).Find(&results).Error

	return
}

// GetFromRatio 通过ratio获取内容 分账比例
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromRatio(ratio int8) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`ratio` = ?", ratio).Find(&results).Error

	return
}

// GetBatchFromRatio 批量查找 分账比例
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromRatio(ratios []int8) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`ratio` IN (?)", ratios).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromStatus(status int8) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromStatus(statuss []int8) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromCreatedAt(createdAt *time.Time) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromCreateBy(createBy *uint) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromCreateBy(createBys []*uint) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromUpdateBy(updateBy *uint) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_MeProfitSharingWechatReceiversMgr) GetFromRemark(remark string) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_MeProfitSharingWechatReceiversMgr) GetBatchFromRemark(remarks []string) (results []*MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MeProfitSharingWechatReceiversMgr) FetchByPrimaryKey(id int64) (result MeProfitSharingWechatReceivers, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MeProfitSharingWechatReceivers{}).Where("`id` = ?", id).First(&result).Error

	return
}
