package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdStudentCoursesSalemanMgr struct {
	*_BaseMgr
}

// EdStudentCoursesSalemanMgr open func
func EdStudentCoursesSalemanMgr(db *gorm.DB) *_EdStudentCoursesSalemanMgr {
	if db == nil {
		panic(fmt.Errorf("EdStudentCoursesSalemanMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdStudentCoursesSalemanMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_student_courses_saleman"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdStudentCoursesSalemanMgr) GetTableName() string {
	return "ed_student_courses_saleman"
}

// Reset 重置gorm会话
func (obj *_EdStudentCoursesSalemanMgr) Reset() *_EdStudentCoursesSalemanMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdStudentCoursesSalemanMgr) Get() (result EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdStudentCoursesSalemanMgr) Gets() (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdStudentCoursesSalemanMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithStudentCourseID student_course_id获取 学生-课程_id
func (obj *_EdStudentCoursesSalemanMgr) WithStudentCourseID(studentCourseID int64) Option {
	return optionFunc(func(o *options) { o.query["student_course_id"] = studentCourseID })
}

// WithSaleType sale_type获取 销售类型
func (obj *_EdStudentCoursesSalemanMgr) WithSaleType(saleType string) Option {
	return optionFunc(func(o *options) { o.query["sale_type"] = saleType })
}

// WithSaleChannel sale_channel获取 销售渠道
func (obj *_EdStudentCoursesSalemanMgr) WithSaleChannel(saleChannel string) Option {
	return optionFunc(func(o *options) { o.query["sale_channel"] = saleChannel })
}

// WithSalemanID saleman_id获取 销售人员id
func (obj *_EdStudentCoursesSalemanMgr) WithSalemanID(salemanID int64) Option {
	return optionFunc(func(o *options) { o.query["saleman_id"] = salemanID })
}

// WithRemark remark获取 备注
func (obj *_EdStudentCoursesSalemanMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdStudentCoursesSalemanMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdStudentCoursesSalemanMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdStudentCoursesSalemanMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdStudentCoursesSalemanMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdStudentCoursesSalemanMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdStudentCoursesSalemanMgr) GetByOption(opts ...Option) (result EdStudentCoursesSaleman, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdStudentCoursesSalemanMgr) GetByOptions(opts ...Option) (results []*EdStudentCoursesSaleman, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromStudentCourseID 通过student_course_id获取内容 学生-课程_id
func (obj *_EdStudentCoursesSalemanMgr) GetFromStudentCourseID(studentCourseID int64) (result EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`student_course_id` = ?", studentCourseID).First(&result).Error

	return
}

// GetBatchFromStudentCourseID 批量查找 学生-课程_id
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromStudentCourseID(studentCourseIDs []int64) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`student_course_id` IN (?)", studentCourseIDs).Find(&results).Error

	return
}

// GetFromSaleType 通过sale_type获取内容 销售类型
func (obj *_EdStudentCoursesSalemanMgr) GetFromSaleType(saleType string) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`sale_type` = ?", saleType).Find(&results).Error

	return
}

// GetBatchFromSaleType 批量查找 销售类型
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromSaleType(saleTypes []string) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`sale_type` IN (?)", saleTypes).Find(&results).Error

	return
}

// GetFromSaleChannel 通过sale_channel获取内容 销售渠道
func (obj *_EdStudentCoursesSalemanMgr) GetFromSaleChannel(saleChannel string) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`sale_channel` = ?", saleChannel).Find(&results).Error

	return
}

// GetBatchFromSaleChannel 批量查找 销售渠道
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromSaleChannel(saleChannels []string) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`sale_channel` IN (?)", saleChannels).Find(&results).Error

	return
}

// GetFromSalemanID 通过saleman_id获取内容 销售人员id
func (obj *_EdStudentCoursesSalemanMgr) GetFromSalemanID(salemanID int64) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`saleman_id` = ?", salemanID).Find(&results).Error

	return
}

// GetBatchFromSalemanID 批量查找 销售人员id
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromSalemanID(salemanIDs []int64) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`saleman_id` IN (?)", salemanIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdStudentCoursesSalemanMgr) GetFromRemark(remark *string) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromRemark(remarks []*string) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdStudentCoursesSalemanMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdStudentCoursesSalemanMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdStudentCoursesSalemanMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdStudentCoursesSalemanMgr) GetFromCreateBy(createBy *uint) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdStudentCoursesSalemanMgr) GetFromUpdateBy(updateBy *uint) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdStudentCoursesSalemanMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdStudentCoursesSalemanMgr) FetchByPrimaryKey(studentCourseID int64) (result EdStudentCoursesSaleman, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentCoursesSaleman{}).Where("`student_course_id` = ?", studentCourseID).First(&result).Error

	return
}
