package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdCoursesSchedulingLogMgr struct {
	*_BaseMgr
}

// EdCoursesSchedulingLogMgr open func
func EdCoursesSchedulingLogMgr(db *gorm.DB) *_EdCoursesSchedulingLogMgr {
	if db == nil {
		panic(fmt.Errorf("EdCoursesSchedulingLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdCoursesSchedulingLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_courses_scheduling_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdCoursesSchedulingLogMgr) GetTableName() string {
	return "ed_courses_scheduling_log"
}

// Reset 重置gorm会话
func (obj *_EdCoursesSchedulingLogMgr) Reset() *_EdCoursesSchedulingLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdCoursesSchedulingLogMgr) Get() (result EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdCoursesSchedulingLogMgr) Gets() (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdCoursesSchedulingLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdCoursesSchedulingLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCoursesSchedulingID courses_scheduling_id获取 订座id
func (obj *_EdCoursesSchedulingLogMgr) WithCoursesSchedulingID(coursesSchedulingID int64) Option {
	return optionFunc(func(o *options) { o.query["courses_scheduling_id"] = coursesSchedulingID })
}

// WithOperate operate获取 操作
func (obj *_EdCoursesSchedulingLogMgr) WithOperate(operate int8) Option {
	return optionFunc(func(o *options) { o.query["operate"] = operate })
}

// WithExtra extra获取 额外信息
func (obj *_EdCoursesSchedulingLogMgr) WithExtra(extra *string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithRemark remark获取 备注
func (obj *_EdCoursesSchedulingLogMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdCoursesSchedulingLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreateBy create_by获取
func (obj *_EdCoursesSchedulingLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdCoursesSchedulingLogMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdCoursesSchedulingLogMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdateBy update_by获取
func (obj *_EdCoursesSchedulingLogMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdCoursesSchedulingLogMgr) GetByOption(opts ...Option) (result EdCoursesSchedulingLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdCoursesSchedulingLogMgr) GetByOptions(opts ...Option) (results []*EdCoursesSchedulingLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdCoursesSchedulingLogMgr) GetFromID(id int64) (result EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromID(ids []int64) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCoursesSchedulingID 通过courses_scheduling_id获取内容 订座id
func (obj *_EdCoursesSchedulingLogMgr) GetFromCoursesSchedulingID(coursesSchedulingID int64) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`courses_scheduling_id` = ?", coursesSchedulingID).Find(&results).Error

	return
}

// GetBatchFromCoursesSchedulingID 批量查找 订座id
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromCoursesSchedulingID(coursesSchedulingIDs []int64) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`courses_scheduling_id` IN (?)", coursesSchedulingIDs).Find(&results).Error

	return
}

// GetFromOperate 通过operate获取内容 操作
func (obj *_EdCoursesSchedulingLogMgr) GetFromOperate(operate int8) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`operate` = ?", operate).Find(&results).Error

	return
}

// GetBatchFromOperate 批量查找 操作
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromOperate(operates []int8) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`operate` IN (?)", operates).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 额外信息
func (obj *_EdCoursesSchedulingLogMgr) GetFromExtra(extra *string) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 额外信息
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromExtra(extras []*string) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdCoursesSchedulingLogMgr) GetFromRemark(remark *string) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromRemark(remarks []*string) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdCoursesSchedulingLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdCoursesSchedulingLogMgr) GetFromCreateBy(createBy *uint) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdCoursesSchedulingLogMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdCoursesSchedulingLogMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdCoursesSchedulingLogMgr) GetFromUpdateBy(updateBy *uint) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdCoursesSchedulingLogMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdCoursesSchedulingLogMgr) FetchByPrimaryKey(id int64) (result EdCoursesSchedulingLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSchedulingLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
