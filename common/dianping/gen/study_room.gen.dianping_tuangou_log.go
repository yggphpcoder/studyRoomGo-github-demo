package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _DianpingTuangouLogMgr struct {
	*_BaseMgr
}

// DianpingTuangouLogMgr open func
func DianpingTuangouLogMgr(db *gorm.DB) *_DianpingTuangouLogMgr {
	if db == nil {
		panic(fmt.Errorf("DianpingTuangouLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DianpingTuangouLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dianping_tuangou_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DianpingTuangouLogMgr) GetTableName() string {
	return "dianping_tuangou_log"
}

// Reset 重置gorm会话
func (obj *_DianpingTuangouLogMgr) Reset() *_DianpingTuangouLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DianpingTuangouLogMgr) Get() (result DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DianpingTuangouLogMgr) Gets() (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DianpingTuangouLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_DianpingTuangouLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTuangouID tuangou_id获取 团购表id
func (obj *_DianpingTuangouLogMgr) WithTuangouID(tuangouID int64) Option {
	return optionFunc(func(o *options) { o.query["tuangou_id"] = tuangouID })
}

// WithDealID deal_id获取 套餐id
func (obj *_DianpingTuangouLogMgr) WithDealID(dealID int64) Option {
	return optionFunc(func(o *options) { o.query["deal_id"] = dealID })
}

// WithMemberID member_id获取 会员id
func (obj *_DianpingTuangouLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithPaymentDetailID payment_detail_id获取 支付详情id
func (obj *_DianpingTuangouLogMgr) WithPaymentDetailID(paymentDetailID int64) Option {
	return optionFunc(func(o *options) { o.query["payment_detail_id"] = paymentDetailID })
}

// WithCode code获取 状态
func (obj *_DianpingTuangouLogMgr) WithCode(code int) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithReceiptCode receipt_code获取 输入券码
func (obj *_DianpingTuangouLogMgr) WithReceiptCode(receiptCode string) Option {
	return optionFunc(func(o *options) { o.query["receipt_code"] = receiptCode })
}

// WithPaymentAmount payment_amount获取 支付金额
func (obj *_DianpingTuangouLogMgr) WithPaymentAmount(paymentAmount int64) Option {
	return optionFunc(func(o *options) { o.query["payment_amount"] = paymentAmount })
}

// WithNote note获取 结果记录
func (obj *_DianpingTuangouLogMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

// WithCreatedAt created_at获取
func (obj *_DianpingTuangouLogMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreateBy create_by获取
func (obj *_DianpingTuangouLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithAction action获取 类型
func (obj *_DianpingTuangouLogMgr) WithAction(action string) Option {
	return optionFunc(func(o *options) { o.query["action"] = action })
}

// GetByOption 功能选项模式获取
func (obj *_DianpingTuangouLogMgr) GetByOption(opts ...Option) (result DianpingTuangouLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DianpingTuangouLogMgr) GetByOptions(opts ...Option) (results []*DianpingTuangouLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_DianpingTuangouLogMgr) GetFromID(id int64) (result DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_DianpingTuangouLogMgr) GetBatchFromID(ids []int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTuangouID 通过tuangou_id获取内容 团购表id
func (obj *_DianpingTuangouLogMgr) GetFromTuangouID(tuangouID int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`tuangou_id` = ?", tuangouID).Find(&results).Error

	return
}

// GetBatchFromTuangouID 批量查找 团购表id
func (obj *_DianpingTuangouLogMgr) GetBatchFromTuangouID(tuangouIDs []int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`tuangou_id` IN (?)", tuangouIDs).Find(&results).Error

	return
}

// GetFromDealID 通过deal_id获取内容 套餐id
func (obj *_DianpingTuangouLogMgr) GetFromDealID(dealID int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`deal_id` = ?", dealID).Find(&results).Error

	return
}

// GetBatchFromDealID 批量查找 套餐id
func (obj *_DianpingTuangouLogMgr) GetBatchFromDealID(dealIDs []int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`deal_id` IN (?)", dealIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_DianpingTuangouLogMgr) GetFromMemberID(memberID int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_DianpingTuangouLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromPaymentDetailID 通过payment_detail_id获取内容 支付详情id
func (obj *_DianpingTuangouLogMgr) GetFromPaymentDetailID(paymentDetailID int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`payment_detail_id` = ?", paymentDetailID).Find(&results).Error

	return
}

// GetBatchFromPaymentDetailID 批量查找 支付详情id
func (obj *_DianpingTuangouLogMgr) GetBatchFromPaymentDetailID(paymentDetailIDs []int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`payment_detail_id` IN (?)", paymentDetailIDs).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 状态
func (obj *_DianpingTuangouLogMgr) GetFromCode(code int) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 状态
func (obj *_DianpingTuangouLogMgr) GetBatchFromCode(codes []int) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromReceiptCode 通过receipt_code获取内容 输入券码
func (obj *_DianpingTuangouLogMgr) GetFromReceiptCode(receiptCode string) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`receipt_code` = ?", receiptCode).Find(&results).Error

	return
}

// GetBatchFromReceiptCode 批量查找 输入券码
func (obj *_DianpingTuangouLogMgr) GetBatchFromReceiptCode(receiptCodes []string) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`receipt_code` IN (?)", receiptCodes).Find(&results).Error

	return
}

// GetFromPaymentAmount 通过payment_amount获取内容 支付金额
func (obj *_DianpingTuangouLogMgr) GetFromPaymentAmount(paymentAmount int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`payment_amount` = ?", paymentAmount).Find(&results).Error

	return
}

// GetBatchFromPaymentAmount 批量查找 支付金额
func (obj *_DianpingTuangouLogMgr) GetBatchFromPaymentAmount(paymentAmounts []int64) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`payment_amount` IN (?)", paymentAmounts).Find(&results).Error

	return
}

// GetFromNote 通过note获取内容 结果记录
func (obj *_DianpingTuangouLogMgr) GetFromNote(note string) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`note` = ?", note).Find(&results).Error

	return
}

// GetBatchFromNote 批量查找 结果记录
func (obj *_DianpingTuangouLogMgr) GetBatchFromNote(notes []string) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`note` IN (?)", notes).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_DianpingTuangouLogMgr) GetFromCreatedAt(createdAt time.Time) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_DianpingTuangouLogMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_DianpingTuangouLogMgr) GetFromCreateBy(createBy *uint) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_DianpingTuangouLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromAction 通过action获取内容 类型
func (obj *_DianpingTuangouLogMgr) GetFromAction(action string) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`action` = ?", action).Find(&results).Error

	return
}

// GetBatchFromAction 批量查找 类型
func (obj *_DianpingTuangouLogMgr) GetBatchFromAction(actions []string) (results []*DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`action` IN (?)", actions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DianpingTuangouLogMgr) FetchByPrimaryKey(id int64) (result DianpingTuangouLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(DianpingTuangouLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
