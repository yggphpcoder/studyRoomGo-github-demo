package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdClassSchedulingMgr struct {
	*_BaseMgr
}

// EdClassSchedulingMgr open func
func EdClassSchedulingMgr(db *gorm.DB) *_EdClassSchedulingMgr {
	if db == nil {
		panic(fmt.Errorf("EdClassSchedulingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdClassSchedulingMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_class_scheduling"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdClassSchedulingMgr) GetTableName() string {
	return "ed_class_scheduling"
}

// Reset 重置gorm会话
func (obj *_EdClassSchedulingMgr) Reset() *_EdClassSchedulingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdClassSchedulingMgr) Get() (result EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdClassSchedulingMgr) Gets() (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdClassSchedulingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdClassSchedulingMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdClassSchedulingMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdClassSchedulingMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithClassID class_id获取 班级id
func (obj *_EdClassSchedulingMgr) WithClassID(classID int64) Option {
	return optionFunc(func(o *options) { o.query["class_id"] = classID })
}

// WithCourseID course_id获取 课程id
func (obj *_EdClassSchedulingMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithTeacherID teacher_id获取 教师id
func (obj *_EdClassSchedulingMgr) WithTeacherID(teacherID int64) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithStartDate start_date获取 开始日期
func (obj *_EdClassSchedulingMgr) WithStartDate(startDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 结束日期
func (obj *_EdClassSchedulingMgr) WithEndDate(endDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithStartTime start_time获取 开始时间
func (obj *_EdClassSchedulingMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 开始时间
func (obj *_EdClassSchedulingMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithActualMin actual_min获取 实际上课分钟
func (obj *_EdClassSchedulingMgr) WithActualMin(actualMin int) Option {
	return optionFunc(func(o *options) { o.query["actual_min"] = actualMin })
}

// WithStatus status获取 状态
func (obj *_EdClassSchedulingMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRemark remark获取 备注
func (obj *_EdClassSchedulingMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdClassSchedulingMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdClassSchedulingMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdClassSchedulingMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdClassSchedulingMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdClassSchedulingMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithAddressID address_id获取 上课地址
func (obj *_EdClassSchedulingMgr) WithAddressID(addressID int64) Option {
	return optionFunc(func(o *options) { o.query["address_id"] = addressID })
}

// GetByOption 功能选项模式获取
func (obj *_EdClassSchedulingMgr) GetByOption(opts ...Option) (result EdClassScheduling, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdClassSchedulingMgr) GetByOptions(opts ...Option) (results []*EdClassScheduling, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdClassSchedulingMgr) GetFromID(id int64) (result EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdClassSchedulingMgr) GetBatchFromID(ids []int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdClassSchedulingMgr) GetFromMerchantID(merchantID int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdClassSchedulingMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdClassSchedulingMgr) GetFromShopID(shopID int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdClassSchedulingMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromClassID 通过class_id获取内容 班级id
func (obj *_EdClassSchedulingMgr) GetFromClassID(classID int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`class_id` = ?", classID).Find(&results).Error

	return
}

// GetBatchFromClassID 批量查找 班级id
func (obj *_EdClassSchedulingMgr) GetBatchFromClassID(classIDs []int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`class_id` IN (?)", classIDs).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 课程id
func (obj *_EdClassSchedulingMgr) GetFromCourseID(courseID int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`course_id` = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量查找 课程id
func (obj *_EdClassSchedulingMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`course_id` IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 教师id
func (obj *_EdClassSchedulingMgr) GetFromTeacherID(teacherID int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 教师id
func (obj *_EdClassSchedulingMgr) GetBatchFromTeacherID(teacherIDs []int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 开始日期
func (obj *_EdClassSchedulingMgr) GetFromStartDate(startDate time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 开始日期
func (obj *_EdClassSchedulingMgr) GetBatchFromStartDate(startDates []time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 结束日期
func (obj *_EdClassSchedulingMgr) GetFromEndDate(endDate time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 结束日期
func (obj *_EdClassSchedulingMgr) GetBatchFromEndDate(endDates []time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 开始时间
func (obj *_EdClassSchedulingMgr) GetFromStartTime(startTime time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 开始时间
func (obj *_EdClassSchedulingMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 开始时间
func (obj *_EdClassSchedulingMgr) GetFromEndTime(endTime time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 开始时间
func (obj *_EdClassSchedulingMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromActualMin 通过actual_min获取内容 实际上课分钟
func (obj *_EdClassSchedulingMgr) GetFromActualMin(actualMin int) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`actual_min` = ?", actualMin).Find(&results).Error

	return
}

// GetBatchFromActualMin 批量查找 实际上课分钟
func (obj *_EdClassSchedulingMgr) GetBatchFromActualMin(actualMins []int) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`actual_min` IN (?)", actualMins).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EdClassSchedulingMgr) GetFromStatus(status int8) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EdClassSchedulingMgr) GetBatchFromStatus(statuss []int8) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdClassSchedulingMgr) GetFromRemark(remark *string) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdClassSchedulingMgr) GetBatchFromRemark(remarks []*string) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdClassSchedulingMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdClassSchedulingMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdClassSchedulingMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdClassSchedulingMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdClassSchedulingMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdClassSchedulingMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdClassSchedulingMgr) GetFromCreateBy(createBy *uint) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdClassSchedulingMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdClassSchedulingMgr) GetFromUpdateBy(updateBy *uint) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdClassSchedulingMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromAddressID 通过address_id获取内容 上课地址
func (obj *_EdClassSchedulingMgr) GetFromAddressID(addressID int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`address_id` = ?", addressID).Find(&results).Error

	return
}

// GetBatchFromAddressID 批量查找 上课地址
func (obj *_EdClassSchedulingMgr) GetBatchFromAddressID(addressIDs []int64) (results []*EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`address_id` IN (?)", addressIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdClassSchedulingMgr) FetchByPrimaryKey(id int64) (result EdClassScheduling, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassScheduling{}).Where("`id` = ?", id).First(&result).Error

	return
}
