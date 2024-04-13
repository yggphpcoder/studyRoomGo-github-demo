package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdSalemanSalaryCoursesMgr struct {
	*_BaseMgr
}

// EdSalemanSalaryCoursesMgr open func
func EdSalemanSalaryCoursesMgr(db *gorm.DB) *_EdSalemanSalaryCoursesMgr {
	if db == nil {
		panic(fmt.Errorf("EdSalemanSalaryCoursesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdSalemanSalaryCoursesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_saleman_salary_courses"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdSalemanSalaryCoursesMgr) GetTableName() string {
	return "ed_saleman_salary_courses"
}

// Reset 重置gorm会话
func (obj *_EdSalemanSalaryCoursesMgr) Reset() *_EdSalemanSalaryCoursesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdSalemanSalaryCoursesMgr) Get() (result EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdSalemanSalaryCoursesMgr) Gets() (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdSalemanSalaryCoursesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdSalemanSalaryCoursesMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSalemanSalaryID saleman_salary_id获取 销售员工资id
func (obj *_EdSalemanSalaryCoursesMgr) WithSalemanSalaryID(salemanSalaryID int64) Option {
	return optionFunc(func(o *options) { o.query["saleman_salary_id"] = salemanSalaryID })
}

// WithSaleType sale_type获取 销售类型
func (obj *_EdSalemanSalaryCoursesMgr) WithSaleType(saleType string) Option {
	return optionFunc(func(o *options) { o.query["sale_type"] = saleType })
}

// WithSaleChannel sale_channel获取 销售渠道
func (obj *_EdSalemanSalaryCoursesMgr) WithSaleChannel(saleChannel string) Option {
	return optionFunc(func(o *options) { o.query["sale_channel"] = saleChannel })
}

// WithCourseID course_id获取 课程id
func (obj *_EdSalemanSalaryCoursesMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithStudentID student_id获取 学生id
func (obj *_EdSalemanSalaryCoursesMgr) WithStudentID(studentID int64) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithStudentCourseID student_course_id获取 学生课程id
func (obj *_EdSalemanSalaryCoursesMgr) WithStudentCourseID(studentCourseID int64) Option {
	return optionFunc(func(o *options) { o.query["student_course_id"] = studentCourseID })
}

// WithSalemanID saleman_id获取 教师id
func (obj *_EdSalemanSalaryCoursesMgr) WithSalemanID(salemanID int64) Option {
	return optionFunc(func(o *options) { o.query["saleman_id"] = salemanID })
}

// WithSaleAt sale_at获取 销售日期
func (obj *_EdSalemanSalaryCoursesMgr) WithSaleAt(saleAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["sale_at"] = saleAt })
}

// WithCourseAmount course_amount获取 课程总额
func (obj *_EdSalemanSalaryCoursesMgr) WithCourseAmount(courseAmount float64) Option {
	return optionFunc(func(o *options) { o.query["course_amount"] = courseAmount })
}

// WithCoursePercent course_percent获取 课程提成比例
func (obj *_EdSalemanSalaryCoursesMgr) WithCoursePercent(coursePercent float64) Option {
	return optionFunc(func(o *options) { o.query["course_percent"] = coursePercent })
}

// WithCourseSalaryAmount course_salary_amount获取 课程工资总额
func (obj *_EdSalemanSalaryCoursesMgr) WithCourseSalaryAmount(courseSalaryAmount float64) Option {
	return optionFunc(func(o *options) { o.query["course_salary_amount"] = courseSalaryAmount })
}

// WithSubsidies subsidies获取 补贴
func (obj *_EdSalemanSalaryCoursesMgr) WithSubsidies(subsidies float64) Option {
	return optionFunc(func(o *options) { o.query["subsidies"] = subsidies })
}

// WithAmount amount获取 总额
func (obj *_EdSalemanSalaryCoursesMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithStatus status获取 状态
func (obj *_EdSalemanSalaryCoursesMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRemark remark获取 备注
func (obj *_EdSalemanSalaryCoursesMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdSalemanSalaryCoursesMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdSalemanSalaryCoursesMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdSalemanSalaryCoursesMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdSalemanSalaryCoursesMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdSalemanSalaryCoursesMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdSalemanSalaryCoursesMgr) GetByOption(opts ...Option) (result EdSalemanSalaryCourses, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdSalemanSalaryCoursesMgr) GetByOptions(opts ...Option) (results []*EdSalemanSalaryCourses, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdSalemanSalaryCoursesMgr) GetFromID(id int64) (result EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromID(ids []int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSalemanSalaryID 通过saleman_salary_id获取内容 销售员工资id
func (obj *_EdSalemanSalaryCoursesMgr) GetFromSalemanSalaryID(salemanSalaryID int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`saleman_salary_id` = ?", salemanSalaryID).Find(&results).Error

	return
}

// GetBatchFromSalemanSalaryID 批量查找 销售员工资id
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromSalemanSalaryID(salemanSalaryIDs []int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`saleman_salary_id` IN (?)", salemanSalaryIDs).Find(&results).Error

	return
}

// GetFromSaleType 通过sale_type获取内容 销售类型
func (obj *_EdSalemanSalaryCoursesMgr) GetFromSaleType(saleType string) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`sale_type` = ?", saleType).Find(&results).Error

	return
}

// GetBatchFromSaleType 批量查找 销售类型
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromSaleType(saleTypes []string) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`sale_type` IN (?)", saleTypes).Find(&results).Error

	return
}

// GetFromSaleChannel 通过sale_channel获取内容 销售渠道
func (obj *_EdSalemanSalaryCoursesMgr) GetFromSaleChannel(saleChannel string) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`sale_channel` = ?", saleChannel).Find(&results).Error

	return
}

// GetBatchFromSaleChannel 批量查找 销售渠道
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromSaleChannel(saleChannels []string) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`sale_channel` IN (?)", saleChannels).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 课程id
func (obj *_EdSalemanSalaryCoursesMgr) GetFromCourseID(courseID int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`course_id` = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量查找 课程id
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`course_id` IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_EdSalemanSalaryCoursesMgr) GetFromStudentID(studentID int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`student_id` = ?", studentID).Find(&results).Error

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromStudentID(studentIDs []int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

// GetFromStudentCourseID 通过student_course_id获取内容 学生课程id
func (obj *_EdSalemanSalaryCoursesMgr) GetFromStudentCourseID(studentCourseID int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`student_course_id` = ?", studentCourseID).Find(&results).Error

	return
}

// GetBatchFromStudentCourseID 批量查找 学生课程id
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromStudentCourseID(studentCourseIDs []int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`student_course_id` IN (?)", studentCourseIDs).Find(&results).Error

	return
}

// GetFromSalemanID 通过saleman_id获取内容 教师id
func (obj *_EdSalemanSalaryCoursesMgr) GetFromSalemanID(salemanID int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`saleman_id` = ?", salemanID).Find(&results).Error

	return
}

// GetBatchFromSalemanID 批量查找 教师id
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromSalemanID(salemanIDs []int64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`saleman_id` IN (?)", salemanIDs).Find(&results).Error

	return
}

// GetFromSaleAt 通过sale_at获取内容 销售日期
func (obj *_EdSalemanSalaryCoursesMgr) GetFromSaleAt(saleAt time.Time) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`sale_at` = ?", saleAt).Find(&results).Error

	return
}

// GetBatchFromSaleAt 批量查找 销售日期
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromSaleAt(saleAts []time.Time) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`sale_at` IN (?)", saleAts).Find(&results).Error

	return
}

// GetFromCourseAmount 通过course_amount获取内容 课程总额
func (obj *_EdSalemanSalaryCoursesMgr) GetFromCourseAmount(courseAmount float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`course_amount` = ?", courseAmount).Find(&results).Error

	return
}

// GetBatchFromCourseAmount 批量查找 课程总额
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromCourseAmount(courseAmounts []float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`course_amount` IN (?)", courseAmounts).Find(&results).Error

	return
}

// GetFromCoursePercent 通过course_percent获取内容 课程提成比例
func (obj *_EdSalemanSalaryCoursesMgr) GetFromCoursePercent(coursePercent float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`course_percent` = ?", coursePercent).Find(&results).Error

	return
}

// GetBatchFromCoursePercent 批量查找 课程提成比例
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromCoursePercent(coursePercents []float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`course_percent` IN (?)", coursePercents).Find(&results).Error

	return
}

// GetFromCourseSalaryAmount 通过course_salary_amount获取内容 课程工资总额
func (obj *_EdSalemanSalaryCoursesMgr) GetFromCourseSalaryAmount(courseSalaryAmount float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`course_salary_amount` = ?", courseSalaryAmount).Find(&results).Error

	return
}

// GetBatchFromCourseSalaryAmount 批量查找 课程工资总额
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromCourseSalaryAmount(courseSalaryAmounts []float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`course_salary_amount` IN (?)", courseSalaryAmounts).Find(&results).Error

	return
}

// GetFromSubsidies 通过subsidies获取内容 补贴
func (obj *_EdSalemanSalaryCoursesMgr) GetFromSubsidies(subsidies float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`subsidies` = ?", subsidies).Find(&results).Error

	return
}

// GetBatchFromSubsidies 批量查找 补贴
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromSubsidies(subsidiess []float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`subsidies` IN (?)", subsidiess).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 总额
func (obj *_EdSalemanSalaryCoursesMgr) GetFromAmount(amount float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 总额
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromAmount(amounts []float64) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EdSalemanSalaryCoursesMgr) GetFromStatus(status int8) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromStatus(statuss []int8) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdSalemanSalaryCoursesMgr) GetFromRemark(remark *string) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromRemark(remarks []*string) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdSalemanSalaryCoursesMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdSalemanSalaryCoursesMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdSalemanSalaryCoursesMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdSalemanSalaryCoursesMgr) GetFromCreateBy(createBy *uint) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdSalemanSalaryCoursesMgr) GetFromUpdateBy(updateBy *uint) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdSalemanSalaryCoursesMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdSalemanSalaryCoursesMgr) FetchByPrimaryKey(id int64) (result EdSalemanSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdSalemanSalaryCourses{}).Where("`id` = ?", id).First(&result).Error

	return
}
