package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdCoursesInfoMgr struct {
	*_BaseMgr
}

// EdCoursesInfoMgr open func
func EdCoursesInfoMgr(db *gorm.DB) *_EdCoursesInfoMgr {
	if db == nil {
		panic(fmt.Errorf("EdCoursesInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdCoursesInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_courses_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdCoursesInfoMgr) GetTableName() string {
	return "ed_courses_info"
}

// Reset 重置gorm会话
func (obj *_EdCoursesInfoMgr) Reset() *_EdCoursesInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdCoursesInfoMgr) Get() (result EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdCoursesInfoMgr) Gets() (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdCoursesInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdCoursesInfoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdCoursesInfoMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithCourseType course_type获取 课程类型
func (obj *_EdCoursesInfoMgr) WithCourseType(courseType int8) Option {
	return optionFunc(func(o *options) { o.query["course_type"] = courseType })
}

// WithTeachingType teaching_type获取 教学模式
func (obj *_EdCoursesInfoMgr) WithTeachingType(teachingType int8) Option {
	return optionFunc(func(o *options) { o.query["teaching_type"] = teachingType })
}

// WithBuyLimit buy_limit获取 购买限制
func (obj *_EdCoursesInfoMgr) WithBuyLimit(buyLimit int) Option {
	return optionFunc(func(o *options) { o.query["buy_limit"] = buyLimit })
}

// WithCourseSize course_size获取 课程人数
func (obj *_EdCoursesInfoMgr) WithCourseSize(courseSize int) Option {
	return optionFunc(func(o *options) { o.query["course_size"] = courseSize })
}

// WithCover cover获取 封面
func (obj *_EdCoursesInfoMgr) WithCover(cover *string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithCarousel carousel获取 轮播图
func (obj *_EdCoursesInfoMgr) WithCarousel(carousel *string) Option {
	return optionFunc(func(o *options) { o.query["carousel"] = carousel })
}

// WithTitle title获取 标题
func (obj *_EdCoursesInfoMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithSummary summary获取 简介
func (obj *_EdCoursesInfoMgr) WithSummary(summary string) Option {
	return optionFunc(func(o *options) { o.query["summary"] = summary })
}

// WithContent content获取 详情
func (obj *_EdCoursesInfoMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithStartDate start_date获取 开课日期
func (obj *_EdCoursesInfoMgr) WithStartDate(startDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 课程结束日期
func (obj *_EdCoursesInfoMgr) WithEndDate(endDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithUseShop use_shop获取 适用店铺
func (obj *_EdCoursesInfoMgr) WithUseShop(useShop string) Option {
	return optionFunc(func(o *options) { o.query["use_shop"] = useShop })
}

// WithCreatedAt created_at获取
func (obj *_EdCoursesInfoMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdCoursesInfoMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdCoursesInfoMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdCoursesInfoMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdCoursesInfoMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithStatus status获取 课程状态
func (obj *_EdCoursesInfoMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsShow is_show获取 是否在前台显示
func (obj *_EdCoursesInfoMgr) WithIsShow(isShow int8) Option {
	return optionFunc(func(o *options) { o.query["is_show"] = isShow })
}

// WithSort sort获取 排序
func (obj *_EdCoursesInfoMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *_EdCoursesInfoMgr) GetByOption(opts ...Option) (result EdCoursesInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdCoursesInfoMgr) GetByOptions(opts ...Option) (results []*EdCoursesInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdCoursesInfoMgr) GetFromID(id int64) (result EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdCoursesInfoMgr) GetBatchFromID(ids []int64) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdCoursesInfoMgr) GetFromMerchantID(merchantID int64) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdCoursesInfoMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromCourseType 通过course_type获取内容 课程类型
func (obj *_EdCoursesInfoMgr) GetFromCourseType(courseType int8) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`course_type` = ?", courseType).Find(&results).Error

	return
}

// GetBatchFromCourseType 批量查找 课程类型
func (obj *_EdCoursesInfoMgr) GetBatchFromCourseType(courseTypes []int8) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`course_type` IN (?)", courseTypes).Find(&results).Error

	return
}

// GetFromTeachingType 通过teaching_type获取内容 教学模式
func (obj *_EdCoursesInfoMgr) GetFromTeachingType(teachingType int8) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`teaching_type` = ?", teachingType).Find(&results).Error

	return
}

// GetBatchFromTeachingType 批量查找 教学模式
func (obj *_EdCoursesInfoMgr) GetBatchFromTeachingType(teachingTypes []int8) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`teaching_type` IN (?)", teachingTypes).Find(&results).Error

	return
}

// GetFromBuyLimit 通过buy_limit获取内容 购买限制
func (obj *_EdCoursesInfoMgr) GetFromBuyLimit(buyLimit int) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`buy_limit` = ?", buyLimit).Find(&results).Error

	return
}

// GetBatchFromBuyLimit 批量查找 购买限制
func (obj *_EdCoursesInfoMgr) GetBatchFromBuyLimit(buyLimits []int) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`buy_limit` IN (?)", buyLimits).Find(&results).Error

	return
}

// GetFromCourseSize 通过course_size获取内容 课程人数
func (obj *_EdCoursesInfoMgr) GetFromCourseSize(courseSize int) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`course_size` = ?", courseSize).Find(&results).Error

	return
}

// GetBatchFromCourseSize 批量查找 课程人数
func (obj *_EdCoursesInfoMgr) GetBatchFromCourseSize(courseSizes []int) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`course_size` IN (?)", courseSizes).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容 封面
func (obj *_EdCoursesInfoMgr) GetFromCover(cover *string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`cover` = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量查找 封面
func (obj *_EdCoursesInfoMgr) GetBatchFromCover(covers []*string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`cover` IN (?)", covers).Find(&results).Error

	return
}

// GetFromCarousel 通过carousel获取内容 轮播图
func (obj *_EdCoursesInfoMgr) GetFromCarousel(carousel *string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`carousel` = ?", carousel).Find(&results).Error

	return
}

// GetBatchFromCarousel 批量查找 轮播图
func (obj *_EdCoursesInfoMgr) GetBatchFromCarousel(carousels []*string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`carousel` IN (?)", carousels).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_EdCoursesInfoMgr) GetFromTitle(title string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_EdCoursesInfoMgr) GetBatchFromTitle(titles []string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromSummary 通过summary获取内容 简介
func (obj *_EdCoursesInfoMgr) GetFromSummary(summary string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`summary` = ?", summary).Find(&results).Error

	return
}

// GetBatchFromSummary 批量查找 简介
func (obj *_EdCoursesInfoMgr) GetBatchFromSummary(summarys []string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`summary` IN (?)", summarys).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 详情
func (obj *_EdCoursesInfoMgr) GetFromContent(content string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 详情
func (obj *_EdCoursesInfoMgr) GetBatchFromContent(contents []string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 开课日期
func (obj *_EdCoursesInfoMgr) GetFromStartDate(startDate *time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 开课日期
func (obj *_EdCoursesInfoMgr) GetBatchFromStartDate(startDates []*time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 课程结束日期
func (obj *_EdCoursesInfoMgr) GetFromEndDate(endDate *time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 课程结束日期
func (obj *_EdCoursesInfoMgr) GetBatchFromEndDate(endDates []*time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromUseShop 通过use_shop获取内容 适用店铺
func (obj *_EdCoursesInfoMgr) GetFromUseShop(useShop string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`use_shop` = ?", useShop).Find(&results).Error

	return
}

// GetBatchFromUseShop 批量查找 适用店铺
func (obj *_EdCoursesInfoMgr) GetBatchFromUseShop(useShops []string) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`use_shop` IN (?)", useShops).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdCoursesInfoMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdCoursesInfoMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdCoursesInfoMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdCoursesInfoMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdCoursesInfoMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdCoursesInfoMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdCoursesInfoMgr) GetFromCreateBy(createBy *uint) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdCoursesInfoMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdCoursesInfoMgr) GetFromUpdateBy(updateBy *uint) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdCoursesInfoMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 课程状态
func (obj *_EdCoursesInfoMgr) GetFromStatus(status int8) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 课程状态
func (obj *_EdCoursesInfoMgr) GetBatchFromStatus(statuss []int8) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsShow 通过is_show获取内容 是否在前台显示
func (obj *_EdCoursesInfoMgr) GetFromIsShow(isShow int8) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`is_show` = ?", isShow).Find(&results).Error

	return
}

// GetBatchFromIsShow 批量查找 是否在前台显示
func (obj *_EdCoursesInfoMgr) GetBatchFromIsShow(isShows []int8) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`is_show` IN (?)", isShows).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_EdCoursesInfoMgr) GetFromSort(sort int) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_EdCoursesInfoMgr) GetBatchFromSort(sorts []int) (results []*EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdCoursesInfoMgr) FetchByPrimaryKey(id int64) (result EdCoursesInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}
