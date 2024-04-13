package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdSalemanSalaryMgr struct {
	*_BaseMgr
}

// EdSalemanSalaryMgr open func
func EdSalemanSalaryMgr(db *gorm.DB) *_EdSalemanSalaryMgr {
	if db == nil {
		panic(fmt.Errorf("EdSalemanSalaryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdSalemanSalaryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_saleman_salary"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdSalemanSalaryMgr) GetTableName() string {
	return "ed_saleman_salary"
}

// Reset 重置gorm会话
func (obj *_EdSalemanSalaryMgr) Reset() *_EdSalemanSalaryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdSalemanSalaryMgr) Get() (result EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdSalemanSalaryMgr) Gets() (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdSalemanSalaryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdSalemanSalaryMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdSalemanSalaryMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdSalemanSalaryMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithSalemanID saleman_id获取 销售员id
func (obj *_EdSalemanSalaryMgr) WithSalemanID(salemanID int64) Option {
	return optionFunc(func(o *options) { o.query["saleman_id"] = salemanID })
}

// WithSalaryAmount salary_amount获取 课程工资总额
func (obj *_EdSalemanSalaryMgr) WithSalaryAmount(salaryAmount float64) Option {
	return optionFunc(func(o *options) { o.query["salary_amount"] = salaryAmount })
}

// WithSalaryTime salary_time获取 发放工资时间
func (obj *_EdSalemanSalaryMgr) WithSalaryTime(salaryTime *time.Time) Option {
	return optionFunc(func(o *options) { o.query["salary_time"] = salaryTime })
}

// WithBonus bonus获取 奖金
func (obj *_EdSalemanSalaryMgr) WithBonus(bonus float64) Option {
	return optionFunc(func(o *options) { o.query["bonus"] = bonus })
}

// WithSubsidies subsidies获取 补贴
func (obj *_EdSalemanSalaryMgr) WithSubsidies(subsidies float64) Option {
	return optionFunc(func(o *options) { o.query["subsidies"] = subsidies })
}

// WithDeductions deductions获取 扣款
func (obj *_EdSalemanSalaryMgr) WithDeductions(deductions float64) Option {
	return optionFunc(func(o *options) { o.query["deductions"] = deductions })
}

// WithTotalSalary total_salary获取 总工资
func (obj *_EdSalemanSalaryMgr) WithTotalSalary(totalSalary float64) Option {
	return optionFunc(func(o *options) { o.query["total_salary"] = totalSalary })
}

// WithSalaryStatus salary_status获取 工资状态
func (obj *_EdSalemanSalaryMgr) WithSalaryStatus(salaryStatus int8) Option {
	return optionFunc(func(o *options) { o.query["salary_status"] = salaryStatus })
}

// WithSalaryRemark salary_remark获取 工资备注
func (obj *_EdSalemanSalaryMgr) WithSalaryRemark(salaryRemark *string) Option {
	return optionFunc(func(o *options) { o.query["salary_remark"] = salaryRemark })
}

// WithBonusRemark bonus_remark获取 奖金备注
func (obj *_EdSalemanSalaryMgr) WithBonusRemark(bonusRemark *string) Option {
	return optionFunc(func(o *options) { o.query["bonus_remark"] = bonusRemark })
}

// WithSubsidiesRemark subsidies_remark获取 补贴备注
func (obj *_EdSalemanSalaryMgr) WithSubsidiesRemark(subsidiesRemark *string) Option {
	return optionFunc(func(o *options) { o.query["subsidies_remark"] = subsidiesRemark })
}

// WithDeductionsRemark deductions_remark获取 扣款备注
func (obj *_EdSalemanSalaryMgr) WithDeductionsRemark(deductionsRemark *string) Option {
	return optionFunc(func(o *options) { o.query["deductions_remark"] = deductionsRemark })
}

// WithRemark remark获取 备注
func (obj *_EdSalemanSalaryMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdSalemanSalaryMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdSalemanSalaryMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdSalemanSalaryMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdSalemanSalaryMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdSalemanSalaryMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdSalemanSalaryMgr) GetByOption(opts ...Option) (result EdSalemanSalary, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdSalemanSalaryMgr) GetByOptions(opts ...Option) (results []*EdSalemanSalary, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdSalemanSalaryMgr) GetFromID(id int64) (result EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdSalemanSalaryMgr) GetBatchFromID(ids []int64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdSalemanSalaryMgr) GetFromMerchantID(merchantID int64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdSalemanSalaryMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdSalemanSalaryMgr) GetFromShopID(shopID int64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdSalemanSalaryMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromSalemanID 通过saleman_id获取内容 销售员id
func (obj *_EdSalemanSalaryMgr) GetFromSalemanID(salemanID int64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`saleman_id` = ?", salemanID).Find(&results).Error

	return
}

// GetBatchFromSalemanID 批量查找 销售员id
func (obj *_EdSalemanSalaryMgr) GetBatchFromSalemanID(salemanIDs []int64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`saleman_id` IN (?)", salemanIDs).Find(&results).Error

	return
}

// GetFromSalaryAmount 通过salary_amount获取内容 课程工资总额
func (obj *_EdSalemanSalaryMgr) GetFromSalaryAmount(salaryAmount float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`salary_amount` = ?", salaryAmount).Find(&results).Error

	return
}

// GetBatchFromSalaryAmount 批量查找 课程工资总额
func (obj *_EdSalemanSalaryMgr) GetBatchFromSalaryAmount(salaryAmounts []float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`salary_amount` IN (?)", salaryAmounts).Find(&results).Error

	return
}

// GetFromSalaryTime 通过salary_time获取内容 发放工资时间
func (obj *_EdSalemanSalaryMgr) GetFromSalaryTime(salaryTime *time.Time) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`salary_time` = ?", salaryTime).Find(&results).Error

	return
}

// GetBatchFromSalaryTime 批量查找 发放工资时间
func (obj *_EdSalemanSalaryMgr) GetBatchFromSalaryTime(salaryTimes []*time.Time) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`salary_time` IN (?)", salaryTimes).Find(&results).Error

	return
}

// GetFromBonus 通过bonus获取内容 奖金
func (obj *_EdSalemanSalaryMgr) GetFromBonus(bonus float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`bonus` = ?", bonus).Find(&results).Error

	return
}

// GetBatchFromBonus 批量查找 奖金
func (obj *_EdSalemanSalaryMgr) GetBatchFromBonus(bonuss []float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`bonus` IN (?)", bonuss).Find(&results).Error

	return
}

// GetFromSubsidies 通过subsidies获取内容 补贴
func (obj *_EdSalemanSalaryMgr) GetFromSubsidies(subsidies float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`subsidies` = ?", subsidies).Find(&results).Error

	return
}

// GetBatchFromSubsidies 批量查找 补贴
func (obj *_EdSalemanSalaryMgr) GetBatchFromSubsidies(subsidiess []float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`subsidies` IN (?)", subsidiess).Find(&results).Error

	return
}

// GetFromDeductions 通过deductions获取内容 扣款
func (obj *_EdSalemanSalaryMgr) GetFromDeductions(deductions float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`deductions` = ?", deductions).Find(&results).Error

	return
}

// GetBatchFromDeductions 批量查找 扣款
func (obj *_EdSalemanSalaryMgr) GetBatchFromDeductions(deductionss []float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`deductions` IN (?)", deductionss).Find(&results).Error

	return
}

// GetFromTotalSalary 通过total_salary获取内容 总工资
func (obj *_EdSalemanSalaryMgr) GetFromTotalSalary(totalSalary float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`total_salary` = ?", totalSalary).Find(&results).Error

	return
}

// GetBatchFromTotalSalary 批量查找 总工资
func (obj *_EdSalemanSalaryMgr) GetBatchFromTotalSalary(totalSalarys []float64) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`total_salary` IN (?)", totalSalarys).Find(&results).Error

	return
}

// GetFromSalaryStatus 通过salary_status获取内容 工资状态
func (obj *_EdSalemanSalaryMgr) GetFromSalaryStatus(salaryStatus int8) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`salary_status` = ?", salaryStatus).Find(&results).Error

	return
}

// GetBatchFromSalaryStatus 批量查找 工资状态
func (obj *_EdSalemanSalaryMgr) GetBatchFromSalaryStatus(salaryStatuss []int8) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`salary_status` IN (?)", salaryStatuss).Find(&results).Error

	return
}

// GetFromSalaryRemark 通过salary_remark获取内容 工资备注
func (obj *_EdSalemanSalaryMgr) GetFromSalaryRemark(salaryRemark *string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`salary_remark` = ?", salaryRemark).Find(&results).Error

	return
}

// GetBatchFromSalaryRemark 批量查找 工资备注
func (obj *_EdSalemanSalaryMgr) GetBatchFromSalaryRemark(salaryRemarks []*string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`salary_remark` IN (?)", salaryRemarks).Find(&results).Error

	return
}

// GetFromBonusRemark 通过bonus_remark获取内容 奖金备注
func (obj *_EdSalemanSalaryMgr) GetFromBonusRemark(bonusRemark *string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`bonus_remark` = ?", bonusRemark).Find(&results).Error

	return
}

// GetBatchFromBonusRemark 批量查找 奖金备注
func (obj *_EdSalemanSalaryMgr) GetBatchFromBonusRemark(bonusRemarks []*string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`bonus_remark` IN (?)", bonusRemarks).Find(&results).Error

	return
}

// GetFromSubsidiesRemark 通过subsidies_remark获取内容 补贴备注
func (obj *_EdSalemanSalaryMgr) GetFromSubsidiesRemark(subsidiesRemark *string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`subsidies_remark` = ?", subsidiesRemark).Find(&results).Error

	return
}

// GetBatchFromSubsidiesRemark 批量查找 补贴备注
func (obj *_EdSalemanSalaryMgr) GetBatchFromSubsidiesRemark(subsidiesRemarks []*string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`subsidies_remark` IN (?)", subsidiesRemarks).Find(&results).Error

	return
}

// GetFromDeductionsRemark 通过deductions_remark获取内容 扣款备注
func (obj *_EdSalemanSalaryMgr) GetFromDeductionsRemark(deductionsRemark *string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`deductions_remark` = ?", deductionsRemark).Find(&results).Error

	return
}

// GetBatchFromDeductionsRemark 批量查找 扣款备注
func (obj *_EdSalemanSalaryMgr) GetBatchFromDeductionsRemark(deductionsRemarks []*string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`deductions_remark` IN (?)", deductionsRemarks).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdSalemanSalaryMgr) GetFromRemark(remark *string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdSalemanSalaryMgr) GetBatchFromRemark(remarks []*string) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdSalemanSalaryMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdSalemanSalaryMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdSalemanSalaryMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdSalemanSalaryMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdSalemanSalaryMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdSalemanSalaryMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdSalemanSalaryMgr) GetFromCreateBy(createBy *uint) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdSalemanSalaryMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdSalemanSalaryMgr) GetFromUpdateBy(updateBy *uint) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdSalemanSalaryMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdSalemanSalaryMgr) FetchByPrimaryKey(id int64) (result EdSalemanSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalary{}).Where("`id` = ?", id).First(&result).Error

	return
}
