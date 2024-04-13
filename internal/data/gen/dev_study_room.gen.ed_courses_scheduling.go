package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdCoursesSchedulingMgr struct {
	*_BaseMgr
}

// EdCoursesSchedulingMgr open func
func EdCoursesSchedulingMgr(db *gorm.DB) *_EdCoursesSchedulingMgr {
	if db == nil {
		panic(fmt.Errorf("EdCoursesSchedulingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdCoursesSchedulingMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_courses_scheduling"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdCoursesSchedulingMgr) GetTableName() string {
	return "ed_courses_scheduling"
}

// Reset 重置gorm会话
func (obj *_EdCoursesSchedulingMgr) Reset() *_EdCoursesSchedulingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdCoursesSchedulingMgr) Get() (result EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdCoursesSchedulingMgr) Gets() (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdCoursesSchedulingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdCoursesSchedulingMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdCoursesSchedulingMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdCoursesSchedulingMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithCourseID course_id获取 课程id
func (obj *_EdCoursesSchedulingMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithStudentID student_id获取 学生id
func (obj *_EdCoursesSchedulingMgr) WithStudentID(studentID int64) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithStudentCourseID student_course_id获取 学生课程id
func (obj *_EdCoursesSchedulingMgr) WithStudentCourseID(studentCourseID int64) Option {
	return optionFunc(func(o *options) { o.query["student_course_id"] = studentCourseID })
}

// WithClassID class_id获取 班级id
func (obj *_EdCoursesSchedulingMgr) WithClassID(classID int64) Option {
	return optionFunc(func(o *options) { o.query["class_id"] = classID })
}

// WithClassSchedulingID class_scheduling_id获取 班排课id
func (obj *_EdCoursesSchedulingMgr) WithClassSchedulingID(classSchedulingID int64) Option {
	return optionFunc(func(o *options) { o.query["class_scheduling_id"] = classSchedulingID })
}

// WithTeacherID teacher_id获取 教师id
func (obj *_EdCoursesSchedulingMgr) WithTeacherID(teacherID int64) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithStartDate start_date获取 开始日期
func (obj *_EdCoursesSchedulingMgr) WithStartDate(startDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 结束日期
func (obj *_EdCoursesSchedulingMgr) WithEndDate(endDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithStartTime start_time获取 开始时间
func (obj *_EdCoursesSchedulingMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 开始时间
func (obj *_EdCoursesSchedulingMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithActualMin actual_min获取 实际上课分钟
func (obj *_EdCoursesSchedulingMgr) WithActualMin(actualMin int) Option {
	return optionFunc(func(o *options) { o.query["actual_min"] = actualMin })
}

// WithCheckIn check_in获取 签到
func (obj *_EdCoursesSchedulingMgr) WithCheckIn(checkIn int8) Option {
	return optionFunc(func(o *options) { o.query["check_in"] = checkIn })
}

// WithStatus status获取 状态
func (obj *_EdCoursesSchedulingMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRemark remark获取 备注
func (obj *_EdCoursesSchedulingMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdCoursesSchedulingMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdCoursesSchedulingMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdCoursesSchedulingMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdCoursesSchedulingMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdCoursesSchedulingMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithAddressID address_id获取 上课地址
func (obj *_EdCoursesSchedulingMgr) WithAddressID(addressID int64) Option {
	return optionFunc(func(o *options) { o.query["address_id"] = addressID })
}

// GetByOption 功能选项模式获取
func (obj *_EdCoursesSchedulingMgr) GetByOption(opts ...Option) (result EdCoursesScheduling, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdCoursesSchedulingMgr) GetByOptions(opts ...Option) (results []*EdCoursesScheduling, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdCoursesSchedulingMgr) GetFromID(id int64) (result EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdCoursesSchedulingMgr) GetBatchFromID(ids []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdCoursesSchedulingMgr) GetFromMerchantID(merchantID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdCoursesSchedulingMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdCoursesSchedulingMgr) GetFromShopID(shopID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdCoursesSchedulingMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 课程id
func (obj *_EdCoursesSchedulingMgr) GetFromCourseID(courseID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`course_id` = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量查找 课程id
func (obj *_EdCoursesSchedulingMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`course_id` IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_EdCoursesSchedulingMgr) GetFromStudentID(studentID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`student_id` = ?", studentID).Find(&results).Error

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_EdCoursesSchedulingMgr) GetBatchFromStudentID(studentIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

// GetFromStudentCourseID 通过student_course_id获取内容 学生课程id
func (obj *_EdCoursesSchedulingMgr) GetFromStudentCourseID(studentCourseID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`student_course_id` = ?", studentCourseID).Find(&results).Error

	return
}

// GetBatchFromStudentCourseID 批量查找 学生课程id
func (obj *_EdCoursesSchedulingMgr) GetBatchFromStudentCourseID(studentCourseIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`student_course_id` IN (?)", studentCourseIDs).Find(&results).Error

	return
}

// GetFromClassID 通过class_id获取内容 班级id
func (obj *_EdCoursesSchedulingMgr) GetFromClassID(classID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`class_id` = ?", classID).Find(&results).Error

	return
}

// GetBatchFromClassID 批量查找 班级id
func (obj *_EdCoursesSchedulingMgr) GetBatchFromClassID(classIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`class_id` IN (?)", classIDs).Find(&results).Error

	return
}

// GetFromClassSchedulingID 通过class_scheduling_id获取内容 班排课id
func (obj *_EdCoursesSchedulingMgr) GetFromClassSchedulingID(classSchedulingID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`class_scheduling_id` = ?", classSchedulingID).Find(&results).Error

	return
}

// GetBatchFromClassSchedulingID 批量查找 班排课id
func (obj *_EdCoursesSchedulingMgr) GetBatchFromClassSchedulingID(classSchedulingIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`class_scheduling_id` IN (?)", classSchedulingIDs).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 教师id
func (obj *_EdCoursesSchedulingMgr) GetFromTeacherID(teacherID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 教师id
func (obj *_EdCoursesSchedulingMgr) GetBatchFromTeacherID(teacherIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 开始日期
func (obj *_EdCoursesSchedulingMgr) GetFromStartDate(startDate time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 开始日期
func (obj *_EdCoursesSchedulingMgr) GetBatchFromStartDate(startDates []time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 结束日期
func (obj *_EdCoursesSchedulingMgr) GetFromEndDate(endDate time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 结束日期
func (obj *_EdCoursesSchedulingMgr) GetBatchFromEndDate(endDates []time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 开始时间
func (obj *_EdCoursesSchedulingMgr) GetFromStartTime(startTime time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 开始时间
func (obj *_EdCoursesSchedulingMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 开始时间
func (obj *_EdCoursesSchedulingMgr) GetFromEndTime(endTime time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 开始时间
func (obj *_EdCoursesSchedulingMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromActualMin 通过actual_min获取内容 实际上课分钟
func (obj *_EdCoursesSchedulingMgr) GetFromActualMin(actualMin int) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`actual_min` = ?", actualMin).Find(&results).Error

	return
}

// GetBatchFromActualMin 批量查找 实际上课分钟
func (obj *_EdCoursesSchedulingMgr) GetBatchFromActualMin(actualMins []int) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`actual_min` IN (?)", actualMins).Find(&results).Error

	return
}

// GetFromCheckIn 通过check_in获取内容 签到
func (obj *_EdCoursesSchedulingMgr) GetFromCheckIn(checkIn int8) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`check_in` = ?", checkIn).Find(&results).Error

	return
}

// GetBatchFromCheckIn 批量查找 签到
func (obj *_EdCoursesSchedulingMgr) GetBatchFromCheckIn(checkIns []int8) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`check_in` IN (?)", checkIns).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EdCoursesSchedulingMgr) GetFromStatus(status int8) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EdCoursesSchedulingMgr) GetBatchFromStatus(statuss []int8) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdCoursesSchedulingMgr) GetFromRemark(remark *string) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdCoursesSchedulingMgr) GetBatchFromRemark(remarks []*string) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdCoursesSchedulingMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdCoursesSchedulingMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdCoursesSchedulingMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdCoursesSchedulingMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdCoursesSchedulingMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdCoursesSchedulingMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdCoursesSchedulingMgr) GetFromCreateBy(createBy *uint) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdCoursesSchedulingMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdCoursesSchedulingMgr) GetFromUpdateBy(updateBy *uint) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdCoursesSchedulingMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromAddressID 通过address_id获取内容 上课地址
func (obj *_EdCoursesSchedulingMgr) GetFromAddressID(addressID int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`address_id` = ?", addressID).Find(&results).Error

	return
}

// GetBatchFromAddressID 批量查找 上课地址
func (obj *_EdCoursesSchedulingMgr) GetBatchFromAddressID(addressIDs []int64) (results []*EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`address_id` IN (?)", addressIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdCoursesSchedulingMgr) FetchByPrimaryKey(id int64) (result EdCoursesScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesScheduling{}).Where("`id` = ?", id).First(&result).Error

	return
}
