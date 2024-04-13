package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrPaymentWechatMgr struct {
	*_BaseMgr
}

// SrPaymentWechatMgr open func
func SrPaymentWechatMgr(db *gorm.DB) *_SrPaymentWechatMgr {
	if db == nil {
		panic(fmt.Errorf("SrPaymentWechatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrPaymentWechatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_payment_wechat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrPaymentWechatMgr) GetTableName() string {
	return "sr_payment_wechat"
}

// Reset 重置gorm会话
func (obj *_SrPaymentWechatMgr) Reset() *_SrPaymentWechatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrPaymentWechatMgr) Get() (result SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrPaymentWechatMgr) Gets() (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrPaymentWechatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrPaymentWechatMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOutTradeNo out_trade_no获取 商户订单号
func (obj *_SrPaymentWechatMgr) WithOutTradeNo(outTradeNo string) Option {
	return optionFunc(func(o *options) { o.query["out_trade_no"] = outTradeNo })
}

// WithTransactionID transaction_id获取 微信支付订单号
func (obj *_SrPaymentWechatMgr) WithTransactionID(transactionID string) Option {
	return optionFunc(func(o *options) { o.query["transaction_id"] = transactionID })
}

// WithTradeType trade_type获取 交易类型
func (obj *_SrPaymentWechatMgr) WithTradeType(tradeType string) Option {
	return optionFunc(func(o *options) { o.query["trade_type"] = tradeType })
}

// WithTradeState trade_state获取 交易状态
func (obj *_SrPaymentWechatMgr) WithTradeState(tradeState string) Option {
	return optionFunc(func(o *options) { o.query["trade_state"] = tradeState })
}

// WithTradeStateDesc trade_state_desc获取 交易状态描述
func (obj *_SrPaymentWechatMgr) WithTradeStateDesc(tradeStateDesc string) Option {
	return optionFunc(func(o *options) { o.query["trade_state_desc"] = tradeStateDesc })
}

// WithBankType bank_type获取 付款银行
func (obj *_SrPaymentWechatMgr) WithBankType(bankType string) Option {
	return optionFunc(func(o *options) { o.query["bank_type"] = bankType })
}

// WithSuccessTime success_time获取 支付完成时间
func (obj *_SrPaymentWechatMgr) WithSuccessTime(successTime string) Option {
	return optionFunc(func(o *options) { o.query["success_time"] = successTime })
}

// WithPayerOpenid payer_openid获取 用户标识
func (obj *_SrPaymentWechatMgr) WithPayerOpenid(payerOpenid string) Option {
	return optionFunc(func(o *options) { o.query["payer_openid"] = payerOpenid })
}

// WithAmountTotal amount_total获取 总金额,分
func (obj *_SrPaymentWechatMgr) WithAmountTotal(amountTotal int) Option {
	return optionFunc(func(o *options) { o.query["amount_total"] = amountTotal })
}

// WithAmountPayerTotal amount_payer_total获取 用户支付金额,分
func (obj *_SrPaymentWechatMgr) WithAmountPayerTotal(amountPayerTotal int) Option {
	return optionFunc(func(o *options) { o.query["amount_payer_total"] = amountPayerTotal })
}

// WithAmountCurrency amount_currency获取 货币类型
func (obj *_SrPaymentWechatMgr) WithAmountCurrency(amountCurrency string) Option {
	return optionFunc(func(o *options) { o.query["amount_currency"] = amountCurrency })
}

// WithAmountPayerCurrency amount_payer_currency获取 用户支付币种
func (obj *_SrPaymentWechatMgr) WithAmountPayerCurrency(amountPayerCurrency string) Option {
	return optionFunc(func(o *options) { o.query["amount_payer_currency"] = amountPayerCurrency })
}

// WithAttach attach获取 附加数据
func (obj *_SrPaymentWechatMgr) WithAttach(attach string) Option {
	return optionFunc(func(o *options) { o.query["attach"] = attach })
}

// WithCreatedAt created_at获取
func (obj *_SrPaymentWechatMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrPaymentWechatMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrPaymentWechatMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrPaymentWechatMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrPaymentWechatMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrPaymentWechatMgr) GetByOption(opts ...Option) (result SrPaymentWechat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrPaymentWechatMgr) GetByOptions(opts ...Option) (results []*SrPaymentWechat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrPaymentWechatMgr) GetFromID(id int64) (result SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrPaymentWechatMgr) GetBatchFromID(ids []int64) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOutTradeNo 通过out_trade_no获取内容 商户订单号
func (obj *_SrPaymentWechatMgr) GetFromOutTradeNo(outTradeNo string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`out_trade_no` = ?", outTradeNo).Find(&results).Error

	return
}

// GetBatchFromOutTradeNo 批量查找 商户订单号
func (obj *_SrPaymentWechatMgr) GetBatchFromOutTradeNo(outTradeNos []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`out_trade_no` IN (?)", outTradeNos).Find(&results).Error

	return
}

// GetFromTransactionID 通过transaction_id获取内容 微信支付订单号
func (obj *_SrPaymentWechatMgr) GetFromTransactionID(transactionID string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`transaction_id` = ?", transactionID).Find(&results).Error

	return
}

// GetBatchFromTransactionID 批量查找 微信支付订单号
func (obj *_SrPaymentWechatMgr) GetBatchFromTransactionID(transactionIDs []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`transaction_id` IN (?)", transactionIDs).Find(&results).Error

	return
}

// GetFromTradeType 通过trade_type获取内容 交易类型
func (obj *_SrPaymentWechatMgr) GetFromTradeType(tradeType string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`trade_type` = ?", tradeType).Find(&results).Error

	return
}

// GetBatchFromTradeType 批量查找 交易类型
func (obj *_SrPaymentWechatMgr) GetBatchFromTradeType(tradeTypes []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`trade_type` IN (?)", tradeTypes).Find(&results).Error

	return
}

// GetFromTradeState 通过trade_state获取内容 交易状态
func (obj *_SrPaymentWechatMgr) GetFromTradeState(tradeState string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`trade_state` = ?", tradeState).Find(&results).Error

	return
}

// GetBatchFromTradeState 批量查找 交易状态
func (obj *_SrPaymentWechatMgr) GetBatchFromTradeState(tradeStates []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`trade_state` IN (?)", tradeStates).Find(&results).Error

	return
}

// GetFromTradeStateDesc 通过trade_state_desc获取内容 交易状态描述
func (obj *_SrPaymentWechatMgr) GetFromTradeStateDesc(tradeStateDesc string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`trade_state_desc` = ?", tradeStateDesc).Find(&results).Error

	return
}

// GetBatchFromTradeStateDesc 批量查找 交易状态描述
func (obj *_SrPaymentWechatMgr) GetBatchFromTradeStateDesc(tradeStateDescs []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`trade_state_desc` IN (?)", tradeStateDescs).Find(&results).Error

	return
}

// GetFromBankType 通过bank_type获取内容 付款银行
func (obj *_SrPaymentWechatMgr) GetFromBankType(bankType string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`bank_type` = ?", bankType).Find(&results).Error

	return
}

// GetBatchFromBankType 批量查找 付款银行
func (obj *_SrPaymentWechatMgr) GetBatchFromBankType(bankTypes []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`bank_type` IN (?)", bankTypes).Find(&results).Error

	return
}

// GetFromSuccessTime 通过success_time获取内容 支付完成时间
func (obj *_SrPaymentWechatMgr) GetFromSuccessTime(successTime string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`success_time` = ?", successTime).Find(&results).Error

	return
}

// GetBatchFromSuccessTime 批量查找 支付完成时间
func (obj *_SrPaymentWechatMgr) GetBatchFromSuccessTime(successTimes []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`success_time` IN (?)", successTimes).Find(&results).Error

	return
}

// GetFromPayerOpenid 通过payer_openid获取内容 用户标识
func (obj *_SrPaymentWechatMgr) GetFromPayerOpenid(payerOpenid string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`payer_openid` = ?", payerOpenid).Find(&results).Error

	return
}

// GetBatchFromPayerOpenid 批量查找 用户标识
func (obj *_SrPaymentWechatMgr) GetBatchFromPayerOpenid(payerOpenids []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`payer_openid` IN (?)", payerOpenids).Find(&results).Error

	return
}

// GetFromAmountTotal 通过amount_total获取内容 总金额,分
func (obj *_SrPaymentWechatMgr) GetFromAmountTotal(amountTotal int) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`amount_total` = ?", amountTotal).Find(&results).Error

	return
}

// GetBatchFromAmountTotal 批量查找 总金额,分
func (obj *_SrPaymentWechatMgr) GetBatchFromAmountTotal(amountTotals []int) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`amount_total` IN (?)", amountTotals).Find(&results).Error

	return
}

// GetFromAmountPayerTotal 通过amount_payer_total获取内容 用户支付金额,分
func (obj *_SrPaymentWechatMgr) GetFromAmountPayerTotal(amountPayerTotal int) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`amount_payer_total` = ?", amountPayerTotal).Find(&results).Error

	return
}

// GetBatchFromAmountPayerTotal 批量查找 用户支付金额,分
func (obj *_SrPaymentWechatMgr) GetBatchFromAmountPayerTotal(amountPayerTotals []int) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`amount_payer_total` IN (?)", amountPayerTotals).Find(&results).Error

	return
}

// GetFromAmountCurrency 通过amount_currency获取内容 货币类型
func (obj *_SrPaymentWechatMgr) GetFromAmountCurrency(amountCurrency string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`amount_currency` = ?", amountCurrency).Find(&results).Error

	return
}

// GetBatchFromAmountCurrency 批量查找 货币类型
func (obj *_SrPaymentWechatMgr) GetBatchFromAmountCurrency(amountCurrencys []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`amount_currency` IN (?)", amountCurrencys).Find(&results).Error

	return
}

// GetFromAmountPayerCurrency 通过amount_payer_currency获取内容 用户支付币种
func (obj *_SrPaymentWechatMgr) GetFromAmountPayerCurrency(amountPayerCurrency string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`amount_payer_currency` = ?", amountPayerCurrency).Find(&results).Error

	return
}

// GetBatchFromAmountPayerCurrency 批量查找 用户支付币种
func (obj *_SrPaymentWechatMgr) GetBatchFromAmountPayerCurrency(amountPayerCurrencys []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`amount_payer_currency` IN (?)", amountPayerCurrencys).Find(&results).Error

	return
}

// GetFromAttach 通过attach获取内容 附加数据
func (obj *_SrPaymentWechatMgr) GetFromAttach(attach string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`attach` = ?", attach).Find(&results).Error

	return
}

// GetBatchFromAttach 批量查找 附加数据
func (obj *_SrPaymentWechatMgr) GetBatchFromAttach(attachs []string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`attach` IN (?)", attachs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrPaymentWechatMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrPaymentWechatMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrPaymentWechatMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrPaymentWechatMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrPaymentWechatMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrPaymentWechatMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrPaymentWechatMgr) GetFromCreateBy(createBy *uint) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrPaymentWechatMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrPaymentWechatMgr) GetFromUpdateBy(updateBy *uint) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrPaymentWechatMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrPaymentWechatMgr) FetchByPrimaryKey(id int64) (result SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexBySrPaymentWechatOutTradeNoIndex  获取多个内容
func (obj *_SrPaymentWechatMgr) FetchIndexBySrPaymentWechatOutTradeNoIndex(outTradeNo string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`out_trade_no` = ?", outTradeNo).Find(&results).Error

	return
}

// FetchIndexBySrPaymentWechatTransactionIDIndex  获取多个内容
func (obj *_SrPaymentWechatMgr) FetchIndexBySrPaymentWechatTransactionIDIndex(transactionID string) (results []*SrPaymentWechat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPaymentWechat{}).Where("`transaction_id` = ?", transactionID).Find(&results).Error

	return
}
