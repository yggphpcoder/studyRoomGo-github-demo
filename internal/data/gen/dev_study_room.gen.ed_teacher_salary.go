package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdTeacherSalaryMgr struct {
	*_BaseMgr
}

// EdTeacherSalaryMgr open func
func EdTeacherSalaryMgr(db *gorm.DB) *_EdTeacherSalaryMgr {
	if db == nil {
		panic(fmt.Errorf("EdTeacherSalaryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdTeacherSalaryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_teacher_salary"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdTeacherSalaryMgr) GetTableName() string {
	return "ed_teacher_salary"
}

// Reset 重置gorm会话
func (obj *_EdTeacherSalaryMgr) Reset() *_EdTeacherSalaryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdTeacherSalaryMgr) Get() (result EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdTeacherSalaryMgr) Gets() (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdTeacherSalaryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdTeacherSalaryMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdTeacherSalaryMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdTeacherSalaryMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithTeacherID teacher_id获取 教师id
func (obj *_EdTeacherSalaryMgr) WithTeacherID(teacherID int64) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithSalaryAmount salary_amount获取 课时工资总额
func (obj *_EdTeacherSalaryMgr) WithSalaryAmount(salaryAmount float64) Option {
	return optionFunc(func(o *options) { o.query["salary_amount"] = salaryAmount })
}

// WithSalaryTime salary_time获取 发放工资时间
func (obj *_EdTeacherSalaryMgr) WithSalaryTime(salaryTime *time.Time) Option {
	return optionFunc(func(o *options) { o.query["salary_time"] = salaryTime })
}

// WithBonus bonus获取 奖金
func (obj *_EdTeacherSalaryMgr) WithBonus(bonus float64) Option {
	return optionFunc(func(o *options) { o.query["bonus"] = bonus })
}

// WithSubsidies subsidies获取 补贴
func (obj *_EdTeacherSalaryMgr) WithSubsidies(subsidies float64) Option {
	return optionFunc(func(o *options) { o.query["subsidies"] = subsidies })
}

// WithDeductions deductions获取 扣款
func (obj *_EdTeacherSalaryMgr) WithDeductions(deductions float64) Option {
	return optionFunc(func(o *options) { o.query["deductions"] = deductions })
}

// WithTotalSalary total_salary获取 总工资
func (obj *_EdTeacherSalaryMgr) WithTotalSalary(totalSalary float64) Option {
	return optionFunc(func(o *options) { o.query["total_salary"] = totalSalary })
}

// WithSalaryStatus salary_status获取 工资状态
func (obj *_EdTeacherSalaryMgr) WithSalaryStatus(salaryStatus int8) Option {
	return optionFunc(func(o *options) { o.query["salary_status"] = salaryStatus })
}

// WithSalaryRemark salary_remark获取 工资备注
func (obj *_EdTeacherSalaryMgr) WithSalaryRemark(salaryRemark *string) Option {
	return optionFunc(func(o *options) { o.query["salary_remark"] = salaryRemark })
}

// WithBonusRemark bonus_remark获取 奖金备注
func (obj *_EdTeacherSalaryMgr) WithBonusRemark(bonusRemark *string) Option {
	return optionFunc(func(o *options) { o.query["bonus_remark"] = bonusRemark })
}

// WithSubsidiesRemark subsidies_remark获取 补贴备注
func (obj *_EdTeacherSalaryMgr) WithSubsidiesRemark(subsidiesRemark *string) Option {
	return optionFunc(func(o *options) { o.query["subsidies_remark"] = subsidiesRemark })
}

// WithDeductionsRemark deductions_remark获取 扣款备注
func (obj *_EdTeacherSalaryMgr) WithDeductionsRemark(deductionsRemark *string) Option {
	return optionFunc(func(o *options) { o.query["deductions_remark"] = deductionsRemark })
}

// WithRemark remark获取 备注
func (obj *_EdTeacherSalaryMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdTeacherSalaryMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdTeacherSalaryMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdTeacherSalaryMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdTeacherSalaryMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdTeacherSalaryMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdTeacherSalaryMgr) GetByOption(opts ...Option) (result EdTeacherSalary, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdTeacherSalaryMgr) GetByOptions(opts ...Option) (results []*EdTeacherSalary, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdTeacherSalaryMgr) GetFromID(id int64) (result EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdTeacherSalaryMgr) GetBatchFromID(ids []int64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdTeacherSalaryMgr) GetFromMerchantID(merchantID int64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdTeacherSalaryMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdTeacherSalaryMgr) GetFromShopID(shopID int64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdTeacherSalaryMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 教师id
func (obj *_EdTeacherSalaryMgr) GetFromTeacherID(teacherID int64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 教师id
func (obj *_EdTeacherSalaryMgr) GetBatchFromTeacherID(teacherIDs []int64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromSalaryAmount 通过salary_amount获取内容 课时工资总额
func (obj *_EdTeacherSalaryMgr) GetFromSalaryAmount(salaryAmount float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`salary_amount` = ?", salaryAmount).Find(&results).Error

	return
}

// GetBatchFromSalaryAmount 批量查找 课时工资总额
func (obj *_EdTeacherSalaryMgr) GetBatchFromSalaryAmount(salaryAmounts []float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`salary_amount` IN (?)", salaryAmounts).Find(&results).Error

	return
}

// GetFromSalaryTime 通过salary_time获取内容 发放工资时间
func (obj *_EdTeacherSalaryMgr) GetFromSalaryTime(salaryTime *time.Time) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`salary_time` = ?", salaryTime).Find(&results).Error

	return
}

// GetBatchFromSalaryTime 批量查找 发放工资时间
func (obj *_EdTeacherSalaryMgr) GetBatchFromSalaryTime(salaryTimes []*time.Time) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`salary_time` IN (?)", salaryTimes).Find(&results).Error

	return
}

// GetFromBonus 通过bonus获取内容 奖金
func (obj *_EdTeacherSalaryMgr) GetFromBonus(bonus float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`bonus` = ?", bonus).Find(&results).Error

	return
}

// GetBatchFromBonus 批量查找 奖金
func (obj *_EdTeacherSalaryMgr) GetBatchFromBonus(bonuss []float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`bonus` IN (?)", bonuss).Find(&results).Error

	return
}

// GetFromSubsidies 通过subsidies获取内容 补贴
func (obj *_EdTeacherSalaryMgr) GetFromSubsidies(subsidies float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`subsidies` = ?", subsidies).Find(&results).Error

	return
}

// GetBatchFromSubsidies 批量查找 补贴
func (obj *_EdTeacherSalaryMgr) GetBatchFromSubsidies(subsidiess []float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`subsidies` IN (?)", subsidiess).Find(&results).Error

	return
}

// GetFromDeductions 通过deductions获取内容 扣款
func (obj *_EdTeacherSalaryMgr) GetFromDeductions(deductions float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`deductions` = ?", deductions).Find(&results).Error

	return
}

// GetBatchFromDeductions 批量查找 扣款
func (obj *_EdTeacherSalaryMgr) GetBatchFromDeductions(deductionss []float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`deductions` IN (?)", deductionss).Find(&results).Error

	return
}

// GetFromTotalSalary 通过total_salary获取内容 总工资
func (obj *_EdTeacherSalaryMgr) GetFromTotalSalary(totalSalary float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`total_salary` = ?", totalSalary).Find(&results).Error

	return
}

// GetBatchFromTotalSalary 批量查找 总工资
func (obj *_EdTeacherSalaryMgr) GetBatchFromTotalSalary(totalSalarys []float64) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`total_salary` IN (?)", totalSalarys).Find(&results).Error

	return
}

// GetFromSalaryStatus 通过salary_status获取内容 工资状态
func (obj *_EdTeacherSalaryMgr) GetFromSalaryStatus(salaryStatus int8) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`salary_status` = ?", salaryStatus).Find(&results).Error

	return
}

// GetBatchFromSalaryStatus 批量查找 工资状态
func (obj *_EdTeacherSalaryMgr) GetBatchFromSalaryStatus(salaryStatuss []int8) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`salary_status` IN (?)", salaryStatuss).Find(&results).Error

	return
}

// GetFromSalaryRemark 通过salary_remark获取内容 工资备注
func (obj *_EdTeacherSalaryMgr) GetFromSalaryRemark(salaryRemark *string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`salary_remark` = ?", salaryRemark).Find(&results).Error

	return
}

// GetBatchFromSalaryRemark 批量查找 工资备注
func (obj *_EdTeacherSalaryMgr) GetBatchFromSalaryRemark(salaryRemarks []*string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`salary_remark` IN (?)", salaryRemarks).Find(&results).Error

	return
}

// GetFromBonusRemark 通过bonus_remark获取内容 奖金备注
func (obj *_EdTeacherSalaryMgr) GetFromBonusRemark(bonusRemark *string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`bonus_remark` = ?", bonusRemark).Find(&results).Error

	return
}

// GetBatchFromBonusRemark 批量查找 奖金备注
func (obj *_EdTeacherSalaryMgr) GetBatchFromBonusRemark(bonusRemarks []*string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`bonus_remark` IN (?)", bonusRemarks).Find(&results).Error

	return
}

// GetFromSubsidiesRemark 通过subsidies_remark获取内容 补贴备注
func (obj *_EdTeacherSalaryMgr) GetFromSubsidiesRemark(subsidiesRemark *string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`subsidies_remark` = ?", subsidiesRemark).Find(&results).Error

	return
}

// GetBatchFromSubsidiesRemark 批量查找 补贴备注
func (obj *_EdTeacherSalaryMgr) GetBatchFromSubsidiesRemark(subsidiesRemarks []*string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`subsidies_remark` IN (?)", subsidiesRemarks).Find(&results).Error

	return
}

// GetFromDeductionsRemark 通过deductions_remark获取内容 扣款备注
func (obj *_EdTeacherSalaryMgr) GetFromDeductionsRemark(deductionsRemark *string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`deductions_remark` = ?", deductionsRemark).Find(&results).Error

	return
}

// GetBatchFromDeductionsRemark 批量查找 扣款备注
func (obj *_EdTeacherSalaryMgr) GetBatchFromDeductionsRemark(deductionsRemarks []*string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`deductions_remark` IN (?)", deductionsRemarks).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdTeacherSalaryMgr) GetFromRemark(remark *string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdTeacherSalaryMgr) GetBatchFromRemark(remarks []*string) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdTeacherSalaryMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdTeacherSalaryMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdTeacherSalaryMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdTeacherSalaryMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdTeacherSalaryMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdTeacherSalaryMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdTeacherSalaryMgr) GetFromCreateBy(createBy *uint) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdTeacherSalaryMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdTeacherSalaryMgr) GetFromUpdateBy(updateBy *uint) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdTeacherSalaryMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdTeacherSalaryMgr) FetchByPrimaryKey(id int64) (result EdTeacherSalary, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalary{}).Where("`id` = ?", id).First(&result).Error

	return
}
