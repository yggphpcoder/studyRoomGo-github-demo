package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdClassInfoMgr struct {
	*_BaseMgr
}

// EdClassInfoMgr open func
func EdClassInfoMgr(db *gorm.DB) *_EdClassInfoMgr {
	if db == nil {
		panic(fmt.Errorf("EdClassInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdClassInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_class_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdClassInfoMgr) GetTableName() string {
	return "ed_class_info"
}

// Reset 重置gorm会话
func (obj *_EdClassInfoMgr) Reset() *_EdClassInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdClassInfoMgr) Get() (result EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdClassInfoMgr) Gets() (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdClassInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdClassInfoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdClassInfoMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdClassInfoMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithCourseID course_id获取 课程id
func (obj *_EdClassInfoMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithTeacherID teacher_id获取 教师id
func (obj *_EdClassInfoMgr) WithTeacherID(teacherID int64) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithStartDate start_date获取 开始日期
func (obj *_EdClassInfoMgr) WithStartDate(startDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 结束日期
func (obj *_EdClassInfoMgr) WithEndDate(endDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithUseTime use_time获取 上课时段
func (obj *_EdClassInfoMgr) WithUseTime(useTime string) Option {
	return optionFunc(func(o *options) { o.query["use_time"] = useTime })
}

// WithStatus status获取 状态
func (obj *_EdClassInfoMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithName name获取 名称
func (obj *_EdClassInfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSummary summary获取 简介
func (obj *_EdClassInfoMgr) WithSummary(summary string) Option {
	return optionFunc(func(o *options) { o.query["summary"] = summary })
}

// WithRemark remark获取 备注
func (obj *_EdClassInfoMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdClassInfoMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdClassInfoMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdClassInfoMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdClassInfoMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdClassInfoMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdClassInfoMgr) GetByOption(opts ...Option) (result EdClassInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdClassInfoMgr) GetByOptions(opts ...Option) (results []*EdClassInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdClassInfoMgr) GetFromID(id int64) (result EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdClassInfoMgr) GetBatchFromID(ids []int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdClassInfoMgr) GetFromMerchantID(merchantID int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdClassInfoMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdClassInfoMgr) GetFromShopID(shopID int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdClassInfoMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 课程id
func (obj *_EdClassInfoMgr) GetFromCourseID(courseID int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`course_id` = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量查找 课程id
func (obj *_EdClassInfoMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`course_id` IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 教师id
func (obj *_EdClassInfoMgr) GetFromTeacherID(teacherID int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 教师id
func (obj *_EdClassInfoMgr) GetBatchFromTeacherID(teacherIDs []int64) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 开始日期
func (obj *_EdClassInfoMgr) GetFromStartDate(startDate time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 开始日期
func (obj *_EdClassInfoMgr) GetBatchFromStartDate(startDates []time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 结束日期
func (obj *_EdClassInfoMgr) GetFromEndDate(endDate time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 结束日期
func (obj *_EdClassInfoMgr) GetBatchFromEndDate(endDates []time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromUseTime 通过use_time获取内容 上课时段
func (obj *_EdClassInfoMgr) GetFromUseTime(useTime string) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`use_time` = ?", useTime).Find(&results).Error

	return
}

// GetBatchFromUseTime 批量查找 上课时段
func (obj *_EdClassInfoMgr) GetBatchFromUseTime(useTimes []string) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`use_time` IN (?)", useTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EdClassInfoMgr) GetFromStatus(status int8) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EdClassInfoMgr) GetBatchFromStatus(statuss []int8) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_EdClassInfoMgr) GetFromName(name string) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_EdClassInfoMgr) GetBatchFromName(names []string) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSummary 通过summary获取内容 简介
func (obj *_EdClassInfoMgr) GetFromSummary(summary string) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`summary` = ?", summary).Find(&results).Error

	return
}

// GetBatchFromSummary 批量查找 简介
func (obj *_EdClassInfoMgr) GetBatchFromSummary(summarys []string) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`summary` IN (?)", summarys).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdClassInfoMgr) GetFromRemark(remark *string) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdClassInfoMgr) GetBatchFromRemark(remarks []*string) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdClassInfoMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdClassInfoMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdClassInfoMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdClassInfoMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdClassInfoMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdClassInfoMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdClassInfoMgr) GetFromCreateBy(createBy *uint) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdClassInfoMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdClassInfoMgr) GetFromUpdateBy(updateBy *uint) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdClassInfoMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdClassInfoMgr) FetchByPrimaryKey(id int64) (result EdClassInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdClassInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}
