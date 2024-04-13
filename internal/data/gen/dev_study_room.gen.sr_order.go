package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrOrderMgr struct {
	*_BaseMgr
}

// SrOrderMgr open func
func SrOrderMgr(db *gorm.DB) *_SrOrderMgr {
	if db == nil {
		panic(fmt.Errorf("SrOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrOrderMgr) GetTableName() string {
	return "sr_order"
}

// Reset 重置gorm会话
func (obj *_SrOrderMgr) Reset() *_SrOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrOrderMgr) Get() (result SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrOrderMgr) Gets() (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrOrderMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNo order_no获取 订单编号
func (obj *_SrOrderMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithOrderDesc order_desc获取 订单描述
func (obj *_SrOrderMgr) WithOrderDesc(orderDesc string) Option {
	return optionFunc(func(o *options) { o.query["order_desc"] = orderDesc })
}

// WithMemberID member_id获取 会员id
func (obj *_SrOrderMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMerchantID merchant_id获取 商户id
func (obj *_SrOrderMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrOrderMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithOrderType order_type获取 订单类型 1:购买套餐，2:直接订座
func (obj *_SrOrderMgr) WithOrderType(orderType int8) Option {
	return optionFunc(func(o *options) { o.query["order_type"] = orderType })
}

// WithOrderTypeRelateID order_type_relate_id获取 订单类型关联信息
func (obj *_SrOrderMgr) WithOrderTypeRelateID(orderTypeRelateID string) Option {
	return optionFunc(func(o *options) { o.query["order_type_relate_id"] = orderTypeRelateID })
}

// WithPayProvider pay_provider获取 支付方式：wxpay
func (obj *_SrOrderMgr) WithPayProvider(payProvider string) Option {
	return optionFunc(func(o *options) { o.query["pay_provider"] = payProvider })
}

// WithStatus status获取 订单状态
func (obj *_SrOrderMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPrice price获取 订单金额
func (obj *_SrOrderMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithQuantity quantity获取 数量
func (obj *_SrOrderMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithCreatedAt created_at获取
func (obj *_SrOrderMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrOrderMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrOrderMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrOrderMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrOrderMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrOrderMgr) GetByOption(opts ...Option) (result SrOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrOrderMgr) GetByOptions(opts ...Option) (results []*SrOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrOrderMgr) GetFromID(id int64) (result SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrOrderMgr) GetBatchFromID(ids []int64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单编号
func (obj *_SrOrderMgr) GetFromOrderNo(orderNo string) (result SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_no` = ?", orderNo).First(&result).Error

	return
}

// GetBatchFromOrderNo 批量查找 订单编号
func (obj *_SrOrderMgr) GetBatchFromOrderNo(orderNos []string) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromOrderDesc 通过order_desc获取内容 订单描述
func (obj *_SrOrderMgr) GetFromOrderDesc(orderDesc string) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_desc` = ?", orderDesc).Find(&results).Error

	return
}

// GetBatchFromOrderDesc 批量查找 订单描述
func (obj *_SrOrderMgr) GetBatchFromOrderDesc(orderDescs []string) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_desc` IN (?)", orderDescs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrOrderMgr) GetFromMemberID(memberID int64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrOrderMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商户id
func (obj *_SrOrderMgr) GetFromMerchantID(merchantID int64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商户id
func (obj *_SrOrderMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrOrderMgr) GetFromShopID(shopID int64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrOrderMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromOrderType 通过order_type获取内容 订单类型 1:购买套餐，2:直接订座
func (obj *_SrOrderMgr) GetFromOrderType(orderType int8) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_type` = ?", orderType).Find(&results).Error

	return
}

// GetBatchFromOrderType 批量查找 订单类型 1:购买套餐，2:直接订座
func (obj *_SrOrderMgr) GetBatchFromOrderType(orderTypes []int8) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_type` IN (?)", orderTypes).Find(&results).Error

	return
}

// GetFromOrderTypeRelateID 通过order_type_relate_id获取内容 订单类型关联信息
func (obj *_SrOrderMgr) GetFromOrderTypeRelateID(orderTypeRelateID string) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_type_relate_id` = ?", orderTypeRelateID).Find(&results).Error

	return
}

// GetBatchFromOrderTypeRelateID 批量查找 订单类型关联信息
func (obj *_SrOrderMgr) GetBatchFromOrderTypeRelateID(orderTypeRelateIDs []string) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_type_relate_id` IN (?)", orderTypeRelateIDs).Find(&results).Error

	return
}

// GetFromPayProvider 通过pay_provider获取内容 支付方式：wxpay
func (obj *_SrOrderMgr) GetFromPayProvider(payProvider string) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`pay_provider` = ?", payProvider).Find(&results).Error

	return
}

// GetBatchFromPayProvider 批量查找 支付方式：wxpay
func (obj *_SrOrderMgr) GetBatchFromPayProvider(payProviders []string) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`pay_provider` IN (?)", payProviders).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 订单状态
func (obj *_SrOrderMgr) GetFromStatus(status int8) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 订单状态
func (obj *_SrOrderMgr) GetBatchFromStatus(statuss []int8) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 订单金额
func (obj *_SrOrderMgr) GetFromPrice(price float64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 订单金额
func (obj *_SrOrderMgr) GetBatchFromPrice(prices []float64) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 数量
func (obj *_SrOrderMgr) GetFromQuantity(quantity int) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 数量
func (obj *_SrOrderMgr) GetBatchFromQuantity(quantitys []int) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrOrderMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrOrderMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrOrderMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrOrderMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrOrderMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrOrderMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrOrderMgr) GetFromCreateBy(createBy *uint) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrOrderMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrOrderMgr) GetFromUpdateBy(updateBy *uint) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrOrderMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrOrderMgr) FetchByPrimaryKey(id int64) (result SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueBySrOrderOrderNoUIndex primary or index 获取唯一内容
func (obj *_SrOrderMgr) FetchUniqueBySrOrderOrderNoUIndex(orderNo string) (result SrOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrOrder{}).Where("`order_no` = ?", orderNo).First(&result).Error

	return
}
