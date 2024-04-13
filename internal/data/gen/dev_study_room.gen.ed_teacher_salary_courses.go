package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdTeacherSalaryCoursesMgr struct {
	*_BaseMgr
}

// EdTeacherSalaryCoursesMgr open func
func EdTeacherSalaryCoursesMgr(db *gorm.DB) *_EdTeacherSalaryCoursesMgr {
	if db == nil {
		panic(fmt.Errorf("EdTeacherSalaryCoursesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdTeacherSalaryCoursesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_teacher_salary_courses"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdTeacherSalaryCoursesMgr) GetTableName() string {
	return "ed_teacher_salary_courses"
}

// Reset 重置gorm会话
func (obj *_EdTeacherSalaryCoursesMgr) Reset() *_EdTeacherSalaryCoursesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdTeacherSalaryCoursesMgr) Get() (result EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdTeacherSalaryCoursesMgr) Gets() (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdTeacherSalaryCoursesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdTeacherSalaryCoursesMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTeacherSalaryID teacher_salary_id获取 教师工资id
func (obj *_EdTeacherSalaryCoursesMgr) WithTeacherSalaryID(teacherSalaryID int64) Option {
	return optionFunc(func(o *options) { o.query["teacher_salary_id"] = teacherSalaryID })
}

// WithTeachingType teaching_type获取 教学模式
func (obj *_EdTeacherSalaryCoursesMgr) WithTeachingType(teachingType int8) Option {
	return optionFunc(func(o *options) { o.query["teaching_type"] = teachingType })
}

// WithSchedulingID scheduling_id获取 课程排课id
func (obj *_EdTeacherSalaryCoursesMgr) WithSchedulingID(schedulingID int64) Option {
	return optionFunc(func(o *options) { o.query["scheduling_id"] = schedulingID })
}

// WithClassID class_id获取 班级id
func (obj *_EdTeacherSalaryCoursesMgr) WithClassID(classID int64) Option {
	return optionFunc(func(o *options) { o.query["class_id"] = classID })
}

// WithCourseID course_id获取 课程id
func (obj *_EdTeacherSalaryCoursesMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithStudentID student_id获取 学生id
func (obj *_EdTeacherSalaryCoursesMgr) WithStudentID(studentID int64) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithStudentCourseID student_course_id获取 学生课程id
func (obj *_EdTeacherSalaryCoursesMgr) WithStudentCourseID(studentCourseID int64) Option {
	return optionFunc(func(o *options) { o.query["student_course_id"] = studentCourseID })
}

// WithTeacherID teacher_id获取 教师id
func (obj *_EdTeacherSalaryCoursesMgr) WithTeacherID(teacherID int64) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithStartDate start_date获取 开始日期
func (obj *_EdTeacherSalaryCoursesMgr) WithStartDate(startDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 结束日期
func (obj *_EdTeacherSalaryCoursesMgr) WithEndDate(endDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithStartTime start_time获取 开始时间
func (obj *_EdTeacherSalaryCoursesMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 开始时间
func (obj *_EdTeacherSalaryCoursesMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithCourseSalaryPer course_salary_per获取 每课时工资(分钟)
func (obj *_EdTeacherSalaryCoursesMgr) WithCourseSalaryPer(courseSalaryPer float64) Option {
	return optionFunc(func(o *options) { o.query["course_salary_per"] = courseSalaryPer })
}

// WithActualMin actual_min获取 实际上课分钟
func (obj *_EdTeacherSalaryCoursesMgr) WithActualMin(actualMin int) Option {
	return optionFunc(func(o *options) { o.query["actual_min"] = actualMin })
}

// WithCourseSalaryAmount course_salary_amount获取 课时工资总额
func (obj *_EdTeacherSalaryCoursesMgr) WithCourseSalaryAmount(courseSalaryAmount float64) Option {
	return optionFunc(func(o *options) { o.query["course_salary_amount"] = courseSalaryAmount })
}

// WithSubsidies subsidies获取 补贴
func (obj *_EdTeacherSalaryCoursesMgr) WithSubsidies(subsidies float64) Option {
	return optionFunc(func(o *options) { o.query["subsidies"] = subsidies })
}

// WithAmount amount获取 总额
func (obj *_EdTeacherSalaryCoursesMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithStatus status获取 状态
func (obj *_EdTeacherSalaryCoursesMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRemark remark获取 备注
func (obj *_EdTeacherSalaryCoursesMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdTeacherSalaryCoursesMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdTeacherSalaryCoursesMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdTeacherSalaryCoursesMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdTeacherSalaryCoursesMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdTeacherSalaryCoursesMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdTeacherSalaryCoursesMgr) GetByOption(opts ...Option) (result EdTeacherSalaryCourses, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdTeacherSalaryCoursesMgr) GetByOptions(opts ...Option) (results []*EdTeacherSalaryCourses, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdTeacherSalaryCoursesMgr) GetFromID(id int64) (result EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromID(ids []int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTeacherSalaryID 通过teacher_salary_id获取内容 教师工资id
func (obj *_EdTeacherSalaryCoursesMgr) GetFromTeacherSalaryID(teacherSalaryID int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`teacher_salary_id` = ?", teacherSalaryID).Find(&results).Error

	return
}

// GetBatchFromTeacherSalaryID 批量查找 教师工资id
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromTeacherSalaryID(teacherSalaryIDs []int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`teacher_salary_id` IN (?)", teacherSalaryIDs).Find(&results).Error

	return
}

// GetFromTeachingType 通过teaching_type获取内容 教学模式
func (obj *_EdTeacherSalaryCoursesMgr) GetFromTeachingType(teachingType int8) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`teaching_type` = ?", teachingType).Find(&results).Error

	return
}

// GetBatchFromTeachingType 批量查找 教学模式
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromTeachingType(teachingTypes []int8) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`teaching_type` IN (?)", teachingTypes).Find(&results).Error

	return
}

// GetFromSchedulingID 通过scheduling_id获取内容 课程排课id
func (obj *_EdTeacherSalaryCoursesMgr) GetFromSchedulingID(schedulingID int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`scheduling_id` = ?", schedulingID).Find(&results).Error

	return
}

// GetBatchFromSchedulingID 批量查找 课程排课id
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromSchedulingID(schedulingIDs []int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`scheduling_id` IN (?)", schedulingIDs).Find(&results).Error

	return
}

// GetFromClassID 通过class_id获取内容 班级id
func (obj *_EdTeacherSalaryCoursesMgr) GetFromClassID(classID int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`class_id` = ?", classID).Find(&results).Error

	return
}

// GetBatchFromClassID 批量查找 班级id
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromClassID(classIDs []int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`class_id` IN (?)", classIDs).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 课程id
func (obj *_EdTeacherSalaryCoursesMgr) GetFromCourseID(courseID int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`course_id` = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量查找 课程id
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`course_id` IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_EdTeacherSalaryCoursesMgr) GetFromStudentID(studentID int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`student_id` = ?", studentID).Find(&results).Error

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromStudentID(studentIDs []int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

// GetFromStudentCourseID 通过student_course_id获取内容 学生课程id
func (obj *_EdTeacherSalaryCoursesMgr) GetFromStudentCourseID(studentCourseID int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`student_course_id` = ?", studentCourseID).Find(&results).Error

	return
}

// GetBatchFromStudentCourseID 批量查找 学生课程id
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromStudentCourseID(studentCourseIDs []int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`student_course_id` IN (?)", studentCourseIDs).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 教师id
func (obj *_EdTeacherSalaryCoursesMgr) GetFromTeacherID(teacherID int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 教师id
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromTeacherID(teacherIDs []int64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 开始日期
func (obj *_EdTeacherSalaryCoursesMgr) GetFromStartDate(startDate time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 开始日期
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromStartDate(startDates []time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 结束日期
func (obj *_EdTeacherSalaryCoursesMgr) GetFromEndDate(endDate time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 结束日期
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromEndDate(endDates []time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 开始时间
func (obj *_EdTeacherSalaryCoursesMgr) GetFromStartTime(startTime time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 开始时间
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 开始时间
func (obj *_EdTeacherSalaryCoursesMgr) GetFromEndTime(endTime time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 开始时间
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromCourseSalaryPer 通过course_salary_per获取内容 每课时工资(分钟)
func (obj *_EdTeacherSalaryCoursesMgr) GetFromCourseSalaryPer(courseSalaryPer float64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`course_salary_per` = ?", courseSalaryPer).Find(&results).Error

	return
}

// GetBatchFromCourseSalaryPer 批量查找 每课时工资(分钟)
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromCourseSalaryPer(courseSalaryPers []float64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`course_salary_per` IN (?)", courseSalaryPers).Find(&results).Error

	return
}

// GetFromActualMin 通过actual_min获取内容 实际上课分钟
func (obj *_EdTeacherSalaryCoursesMgr) GetFromActualMin(actualMin int) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`actual_min` = ?", actualMin).Find(&results).Error

	return
}

// GetBatchFromActualMin 批量查找 实际上课分钟
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromActualMin(actualMins []int) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`actual_min` IN (?)", actualMins).Find(&results).Error

	return
}

// GetFromCourseSalaryAmount 通过course_salary_amount获取内容 课时工资总额
func (obj *_EdTeacherSalaryCoursesMgr) GetFromCourseSalaryAmount(courseSalaryAmount float64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`course_salary_amount` = ?", courseSalaryAmount).Find(&results).Error

	return
}

// GetBatchFromCourseSalaryAmount 批量查找 课时工资总额
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromCourseSalaryAmount(courseSalaryAmounts []float64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`course_salary_amount` IN (?)", courseSalaryAmounts).Find(&results).Error

	return
}

// GetFromSubsidies 通过subsidies获取内容 补贴
func (obj *_EdTeacherSalaryCoursesMgr) GetFromSubsidies(subsidies float64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`subsidies` = ?", subsidies).Find(&results).Error

	return
}

// GetBatchFromSubsidies 批量查找 补贴
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromSubsidies(subsidiess []float64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`subsidies` IN (?)", subsidiess).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 总额
func (obj *_EdTeacherSalaryCoursesMgr) GetFromAmount(amount float64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 总额
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromAmount(amounts []float64) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EdTeacherSalaryCoursesMgr) GetFromStatus(status int8) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromStatus(statuss []int8) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdTeacherSalaryCoursesMgr) GetFromRemark(remark *string) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromRemark(remarks []*string) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdTeacherSalaryCoursesMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdTeacherSalaryCoursesMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdTeacherSalaryCoursesMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdTeacherSalaryCoursesMgr) GetFromCreateBy(createBy *uint) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdTeacherSalaryCoursesMgr) GetFromUpdateBy(updateBy *uint) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdTeacherSalaryCoursesMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdTeacherSalaryCoursesMgr) FetchByPrimaryKey(id int64) (result EdTeacherSalaryCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherSalaryCourses{}).Where("`id` = ?", id).First(&result).Error

	return
}
