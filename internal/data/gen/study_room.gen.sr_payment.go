package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrPaymentMgr struct {
	*_BaseMgr
}

// SrPaymentMgr open func
func SrPaymentMgr(db *gorm.DB) *_SrPaymentMgr {
	if db == nil {
		panic(fmt.Errorf("SrPaymentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrPaymentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_payment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrPaymentMgr) GetTableName() string {
	return "sr_payment"
}

// Reset 重置gorm会话
func (obj *_SrPaymentMgr) Reset() *_SrPaymentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrPaymentMgr) Get() (result SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrPaymentMgr) Gets() (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrPaymentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrPaymentMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNo order_no获取 订单编号
func (obj *_SrPaymentMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithTransactionID transaction_id获取 交易号
func (obj *_SrPaymentMgr) WithTransactionID(transactionID *string) Option {
	return optionFunc(func(o *options) { o.query["transaction_id"] = transactionID })
}

// WithPaymentType payment_type获取 支付类型
func (obj *_SrPaymentMgr) WithPaymentType(paymentType string) Option {
	return optionFunc(func(o *options) { o.query["payment_type"] = paymentType })
}

// WithCurrency currency获取 币种
func (obj *_SrPaymentMgr) WithCurrency(currency string) Option {
	return optionFunc(func(o *options) { o.query["currency"] = currency })
}

// WithAmount amount获取 金额
func (obj *_SrPaymentMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithStatus status获取 状态
func (obj *_SrPaymentMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_SrPaymentMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrPaymentMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrPaymentMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrPaymentMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrPaymentMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrPaymentMgr) GetByOption(opts ...Option) (result SrPayment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrPaymentMgr) GetByOptions(opts ...Option) (results []*SrPayment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrPaymentMgr) GetFromID(id int64) (result SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrPaymentMgr) GetBatchFromID(ids []int64) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单编号
func (obj *_SrPaymentMgr) GetFromOrderNo(orderNo string) (result SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`order_no` = ?", orderNo).First(&result).Error

	return
}

// GetBatchFromOrderNo 批量查找 订单编号
func (obj *_SrPaymentMgr) GetBatchFromOrderNo(orderNos []string) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromTransactionID 通过transaction_id获取内容 交易号
func (obj *_SrPaymentMgr) GetFromTransactionID(transactionID *string) (result SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`transaction_id` = ?", transactionID).First(&result).Error

	return
}

// GetBatchFromTransactionID 批量查找 交易号
func (obj *_SrPaymentMgr) GetBatchFromTransactionID(transactionIDs []*string) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`transaction_id` IN (?)", transactionIDs).Find(&results).Error

	return
}

// GetFromPaymentType 通过payment_type获取内容 支付类型
func (obj *_SrPaymentMgr) GetFromPaymentType(paymentType string) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`payment_type` = ?", paymentType).Find(&results).Error

	return
}

// GetBatchFromPaymentType 批量查找 支付类型
func (obj *_SrPaymentMgr) GetBatchFromPaymentType(paymentTypes []string) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`payment_type` IN (?)", paymentTypes).Find(&results).Error

	return
}

// GetFromCurrency 通过currency获取内容 币种
func (obj *_SrPaymentMgr) GetFromCurrency(currency string) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`currency` = ?", currency).Find(&results).Error

	return
}

// GetBatchFromCurrency 批量查找 币种
func (obj *_SrPaymentMgr) GetBatchFromCurrency(currencys []string) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`currency` IN (?)", currencys).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 金额
func (obj *_SrPaymentMgr) GetFromAmount(amount float64) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 金额
func (obj *_SrPaymentMgr) GetBatchFromAmount(amounts []float64) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SrPaymentMgr) GetFromStatus(status int8) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SrPaymentMgr) GetBatchFromStatus(statuss []int8) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrPaymentMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrPaymentMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrPaymentMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrPaymentMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrPaymentMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrPaymentMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrPaymentMgr) GetFromCreateBy(createBy *uint) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrPaymentMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrPaymentMgr) GetFromUpdateBy(updateBy *uint) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrPaymentMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrPaymentMgr) FetchByPrimaryKey(id int64) (result SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueBySrPaymentOrderNoUIndex primary or index 获取唯一内容
func (obj *_SrPaymentMgr) FetchUniqueBySrPaymentOrderNoUIndex(orderNo string) (result SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`order_no` = ?", orderNo).First(&result).Error

	return
}

// FetchUniqueBySrPaymentTransactionIDUIndex primary or index 获取唯一内容
func (obj *_SrPaymentMgr) FetchUniqueBySrPaymentTransactionIDUIndex(transactionID *string) (result SrPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPayment{}).Where("`transaction_id` = ?", transactionID).First(&result).Error

	return
}
