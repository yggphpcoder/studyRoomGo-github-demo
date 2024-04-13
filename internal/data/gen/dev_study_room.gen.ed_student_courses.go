package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdStudentCoursesMgr struct {
	*_BaseMgr
}

// EdStudentCoursesMgr open func
func EdStudentCoursesMgr(db *gorm.DB) *_EdStudentCoursesMgr {
	if db == nil {
		panic(fmt.Errorf("EdStudentCoursesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdStudentCoursesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_student_courses"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdStudentCoursesMgr) GetTableName() string {
	return "ed_student_courses"
}

// Reset 重置gorm会话
func (obj *_EdStudentCoursesMgr) Reset() *_EdStudentCoursesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdStudentCoursesMgr) Get() (result EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdStudentCoursesMgr) Gets() (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdStudentCoursesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdStudentCoursesMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdStudentCoursesMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdStudentCoursesMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithStudentID student_id获取 学生id
func (obj *_EdStudentCoursesMgr) WithStudentID(studentID int64) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithActualAmount actual_amount获取 实收价格
func (obj *_EdStudentCoursesMgr) WithActualAmount(actualAmount float64) Option {
	return optionFunc(func(o *options) { o.query["actual_amount"] = actualAmount })
}

// WithRemain remain获取 剩余课数
func (obj *_EdStudentCoursesMgr) WithRemain(remain int) Option {
	return optionFunc(func(o *options) { o.query["remain"] = remain })
}

// WithPerMin per_min获取 每堂分钟
func (obj *_EdStudentCoursesMgr) WithPerMin(perMin int) Option {
	return optionFunc(func(o *options) { o.query["per_min"] = perMin })
}

// WithStatus status获取 状态
func (obj *_EdStudentCoursesMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCourseID course_id获取 课程id
func (obj *_EdStudentCoursesMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithTitle title获取 标题
func (obj *_EdStudentCoursesMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithCourseType course_type获取 课程类型
func (obj *_EdStudentCoursesMgr) WithCourseType(courseType int8) Option {
	return optionFunc(func(o *options) { o.query["course_type"] = courseType })
}

// WithTeachingType teaching_type获取 教学模式
func (obj *_EdStudentCoursesMgr) WithTeachingType(teachingType int8) Option {
	return optionFunc(func(o *options) { o.query["teaching_type"] = teachingType })
}

// WithCourseSize course_size获取 课程人数
func (obj *_EdStudentCoursesMgr) WithCourseSize(courseSize int) Option {
	return optionFunc(func(o *options) { o.query["course_size"] = courseSize })
}

// WithStartDate start_date获取 开课日期
func (obj *_EdStudentCoursesMgr) WithStartDate(startDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 结课日期
func (obj *_EdStudentCoursesMgr) WithEndDate(endDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithCourseSaleID course_sale_id获取 课程营销id
func (obj *_EdStudentCoursesMgr) WithCourseSaleID(courseSaleID int64) Option {
	return optionFunc(func(o *options) { o.query["course_sale_id"] = courseSaleID })
}

// WithCourseSaleType course_sale_type获取 营销类型
func (obj *_EdStudentCoursesMgr) WithCourseSaleType(courseSaleType int8) Option {
	return optionFunc(func(o *options) { o.query["course_sale_type"] = courseSaleType })
}

// WithRule rule获取 营销规则
func (obj *_EdStudentCoursesMgr) WithRule(rule string) Option {
	return optionFunc(func(o *options) { o.query["rule"] = rule })
}

// WithPrice price获取 售价
func (obj *_EdStudentCoursesMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithOriPrice ori_price获取 原价
func (obj *_EdStudentCoursesMgr) WithOriPrice(oriPrice float64) Option {
	return optionFunc(func(o *options) { o.query["ori_price"] = oriPrice })
}

// WithActiveAt active_at获取 激活时间
func (obj *_EdStudentCoursesMgr) WithActiveAt(activeAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["active_at"] = activeAt })
}

// WithInvalidAt invalid_at获取 失效时间
func (obj *_EdStudentCoursesMgr) WithInvalidAt(invalidAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["invalid_at"] = invalidAt })
}

// WithValidDay valid_day获取 有效天数
func (obj *_EdStudentCoursesMgr) WithValidDay(validDay int) Option {
	return optionFunc(func(o *options) { o.query["valid_day"] = validDay })
}

// WithActiveLimit active_limit获取 激活期限天数
func (obj *_EdStudentCoursesMgr) WithActiveLimit(activeLimit int) Option {
	return optionFunc(func(o *options) { o.query["active_limit"] = activeLimit })
}

// WithSort sort获取 排序
func (obj *_EdStudentCoursesMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithCourseSalemanID course_saleman_id获取 课程销售员id
func (obj *_EdStudentCoursesMgr) WithCourseSalemanID(courseSalemanID int64) Option {
	return optionFunc(func(o *options) { o.query["course_saleman_id"] = courseSalemanID })
}

// WithRemark remark获取 备注
func (obj *_EdStudentCoursesMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdStudentCoursesMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdStudentCoursesMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdStudentCoursesMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdStudentCoursesMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdStudentCoursesMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdStudentCoursesMgr) GetByOption(opts ...Option) (result EdStudentCourses, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdStudentCoursesMgr) GetByOptions(opts ...Option) (results []*EdStudentCourses, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdStudentCoursesMgr) GetFromID(id int64) (result EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdStudentCoursesMgr) GetBatchFromID(ids []int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdStudentCoursesMgr) GetFromMerchantID(merchantID int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdStudentCoursesMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdStudentCoursesMgr) GetFromShopID(shopID int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdStudentCoursesMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_EdStudentCoursesMgr) GetFromStudentID(studentID int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`student_id` = ?", studentID).Find(&results).Error

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_EdStudentCoursesMgr) GetBatchFromStudentID(studentIDs []int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

// GetFromActualAmount 通过actual_amount获取内容 实收价格
func (obj *_EdStudentCoursesMgr) GetFromActualAmount(actualAmount float64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`actual_amount` = ?", actualAmount).Find(&results).Error

	return
}

// GetBatchFromActualAmount 批量查找 实收价格
func (obj *_EdStudentCoursesMgr) GetBatchFromActualAmount(actualAmounts []float64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`actual_amount` IN (?)", actualAmounts).Find(&results).Error

	return
}

// GetFromRemain 通过remain获取内容 剩余课数
func (obj *_EdStudentCoursesMgr) GetFromRemain(remain int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`remain` = ?", remain).Find(&results).Error

	return
}

// GetBatchFromRemain 批量查找 剩余课数
func (obj *_EdStudentCoursesMgr) GetBatchFromRemain(remains []int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`remain` IN (?)", remains).Find(&results).Error

	return
}

// GetFromPerMin 通过per_min获取内容 每堂分钟
func (obj *_EdStudentCoursesMgr) GetFromPerMin(perMin int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`per_min` = ?", perMin).Find(&results).Error

	return
}

// GetBatchFromPerMin 批量查找 每堂分钟
func (obj *_EdStudentCoursesMgr) GetBatchFromPerMin(perMins []int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`per_min` IN (?)", perMins).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EdStudentCoursesMgr) GetFromStatus(status int8) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EdStudentCoursesMgr) GetBatchFromStatus(statuss []int8) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 课程id
func (obj *_EdStudentCoursesMgr) GetFromCourseID(courseID int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_id` = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量查找 课程id
func (obj *_EdStudentCoursesMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_id` IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_EdStudentCoursesMgr) GetFromTitle(title string) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_EdStudentCoursesMgr) GetBatchFromTitle(titles []string) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromCourseType 通过course_type获取内容 课程类型
func (obj *_EdStudentCoursesMgr) GetFromCourseType(courseType int8) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_type` = ?", courseType).Find(&results).Error

	return
}

// GetBatchFromCourseType 批量查找 课程类型
func (obj *_EdStudentCoursesMgr) GetBatchFromCourseType(courseTypes []int8) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_type` IN (?)", courseTypes).Find(&results).Error

	return
}

// GetFromTeachingType 通过teaching_type获取内容 教学模式
func (obj *_EdStudentCoursesMgr) GetFromTeachingType(teachingType int8) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`teaching_type` = ?", teachingType).Find(&results).Error

	return
}

// GetBatchFromTeachingType 批量查找 教学模式
func (obj *_EdStudentCoursesMgr) GetBatchFromTeachingType(teachingTypes []int8) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`teaching_type` IN (?)", teachingTypes).Find(&results).Error

	return
}

// GetFromCourseSize 通过course_size获取内容 课程人数
func (obj *_EdStudentCoursesMgr) GetFromCourseSize(courseSize int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_size` = ?", courseSize).Find(&results).Error

	return
}

// GetBatchFromCourseSize 批量查找 课程人数
func (obj *_EdStudentCoursesMgr) GetBatchFromCourseSize(courseSizes []int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_size` IN (?)", courseSizes).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 开课日期
func (obj *_EdStudentCoursesMgr) GetFromStartDate(startDate *time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 开课日期
func (obj *_EdStudentCoursesMgr) GetBatchFromStartDate(startDates []*time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 结课日期
func (obj *_EdStudentCoursesMgr) GetFromEndDate(endDate *time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 结课日期
func (obj *_EdStudentCoursesMgr) GetBatchFromEndDate(endDates []*time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromCourseSaleID 通过course_sale_id获取内容 课程营销id
func (obj *_EdStudentCoursesMgr) GetFromCourseSaleID(courseSaleID int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_sale_id` = ?", courseSaleID).Find(&results).Error

	return
}

// GetBatchFromCourseSaleID 批量查找 课程营销id
func (obj *_EdStudentCoursesMgr) GetBatchFromCourseSaleID(courseSaleIDs []int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_sale_id` IN (?)", courseSaleIDs).Find(&results).Error

	return
}

// GetFromCourseSaleType 通过course_sale_type获取内容 营销类型
func (obj *_EdStudentCoursesMgr) GetFromCourseSaleType(courseSaleType int8) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_sale_type` = ?", courseSaleType).Find(&results).Error

	return
}

// GetBatchFromCourseSaleType 批量查找 营销类型
func (obj *_EdStudentCoursesMgr) GetBatchFromCourseSaleType(courseSaleTypes []int8) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_sale_type` IN (?)", courseSaleTypes).Find(&results).Error

	return
}

// GetFromRule 通过rule获取内容 营销规则
func (obj *_EdStudentCoursesMgr) GetFromRule(rule string) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`rule` = ?", rule).Find(&results).Error

	return
}

// GetBatchFromRule 批量查找 营销规则
func (obj *_EdStudentCoursesMgr) GetBatchFromRule(rules []string) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`rule` IN (?)", rules).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 售价
func (obj *_EdStudentCoursesMgr) GetFromPrice(price float64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 售价
func (obj *_EdStudentCoursesMgr) GetBatchFromPrice(prices []float64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromOriPrice 通过ori_price获取内容 原价
func (obj *_EdStudentCoursesMgr) GetFromOriPrice(oriPrice float64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`ori_price` = ?", oriPrice).Find(&results).Error

	return
}

// GetBatchFromOriPrice 批量查找 原价
func (obj *_EdStudentCoursesMgr) GetBatchFromOriPrice(oriPrices []float64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`ori_price` IN (?)", oriPrices).Find(&results).Error

	return
}

// GetFromActiveAt 通过active_at获取内容 激活时间
func (obj *_EdStudentCoursesMgr) GetFromActiveAt(activeAt *time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`active_at` = ?", activeAt).Find(&results).Error

	return
}

// GetBatchFromActiveAt 批量查找 激活时间
func (obj *_EdStudentCoursesMgr) GetBatchFromActiveAt(activeAts []*time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`active_at` IN (?)", activeAts).Find(&results).Error

	return
}

// GetFromInvalidAt 通过invalid_at获取内容 失效时间
func (obj *_EdStudentCoursesMgr) GetFromInvalidAt(invalidAt *time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`invalid_at` = ?", invalidAt).Find(&results).Error

	return
}

// GetBatchFromInvalidAt 批量查找 失效时间
func (obj *_EdStudentCoursesMgr) GetBatchFromInvalidAt(invalidAts []*time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`invalid_at` IN (?)", invalidAts).Find(&results).Error

	return
}

// GetFromValidDay 通过valid_day获取内容 有效天数
func (obj *_EdStudentCoursesMgr) GetFromValidDay(validDay int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`valid_day` = ?", validDay).Find(&results).Error

	return
}

// GetBatchFromValidDay 批量查找 有效天数
func (obj *_EdStudentCoursesMgr) GetBatchFromValidDay(validDays []int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`valid_day` IN (?)", validDays).Find(&results).Error

	return
}

// GetFromActiveLimit 通过active_limit获取内容 激活期限天数
func (obj *_EdStudentCoursesMgr) GetFromActiveLimit(activeLimit int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`active_limit` = ?", activeLimit).Find(&results).Error

	return
}

// GetBatchFromActiveLimit 批量查找 激活期限天数
func (obj *_EdStudentCoursesMgr) GetBatchFromActiveLimit(activeLimits []int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`active_limit` IN (?)", activeLimits).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_EdStudentCoursesMgr) GetFromSort(sort int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_EdStudentCoursesMgr) GetBatchFromSort(sorts []int) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromCourseSalemanID 通过course_saleman_id获取内容 课程销售员id
func (obj *_EdStudentCoursesMgr) GetFromCourseSalemanID(courseSalemanID int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_saleman_id` = ?", courseSalemanID).Find(&results).Error

	return
}

// GetBatchFromCourseSalemanID 批量查找 课程销售员id
func (obj *_EdStudentCoursesMgr) GetBatchFromCourseSalemanID(courseSalemanIDs []int64) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`course_saleman_id` IN (?)", courseSalemanIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdStudentCoursesMgr) GetFromRemark(remark *string) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdStudentCoursesMgr) GetBatchFromRemark(remarks []*string) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdStudentCoursesMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdStudentCoursesMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdStudentCoursesMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdStudentCoursesMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdStudentCoursesMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdStudentCoursesMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdStudentCoursesMgr) GetFromCreateBy(createBy *uint) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdStudentCoursesMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdStudentCoursesMgr) GetFromUpdateBy(updateBy *uint) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdStudentCoursesMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdStudentCoursesMgr) FetchByPrimaryKey(id int64) (result EdStudentCourses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCourses{}).Where("`id` = ?", id).First(&result).Error

	return
}
