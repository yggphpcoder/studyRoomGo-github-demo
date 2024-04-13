package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdTeacherGrowupMgr struct {
	*_BaseMgr
}

// EdTeacherGrowupMgr open func
func EdTeacherGrowupMgr(db *gorm.DB) *_EdTeacherGrowupMgr {
	if db == nil {
		panic(fmt.Errorf("EdTeacherGrowupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdTeacherGrowupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_teacher_growup"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdTeacherGrowupMgr) GetTableName() string {
	return "ed_teacher_growup"
}

// Reset 重置gorm会话
func (obj *_EdTeacherGrowupMgr) Reset() *_EdTeacherGrowupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdTeacherGrowupMgr) Get() (result EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdTeacherGrowupMgr) Gets() (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdTeacherGrowupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdTeacherGrowupMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTeacherID teacher_id获取 教师id
func (obj *_EdTeacherGrowupMgr) WithTeacherID(teacherID int64) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithType type获取 记录分类
func (obj *_EdTeacherGrowupMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithTitle title获取 标题
func (obj *_EdTeacherGrowupMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_EdTeacherGrowupMgr) WithContent(content *string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithFiles files获取 附件
func (obj *_EdTeacherGrowupMgr) WithFiles(files *string) Option {
	return optionFunc(func(o *options) { o.query["files"] = files })
}

// WithCreatedAt created_at获取
func (obj *_EdTeacherGrowupMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdTeacherGrowupMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdTeacherGrowupMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdTeacherGrowupMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdTeacherGrowupMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdTeacherGrowupMgr) GetByOption(opts ...Option) (result EdTeacherGrowup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdTeacherGrowupMgr) GetByOptions(opts ...Option) (results []*EdTeacherGrowup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdTeacherGrowupMgr) GetFromID(id int64) (result EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdTeacherGrowupMgr) GetBatchFromID(ids []int64) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 教师id
func (obj *_EdTeacherGrowupMgr) GetFromTeacherID(teacherID int64) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 教师id
func (obj *_EdTeacherGrowupMgr) GetBatchFromTeacherID(teacherIDs []int64) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 记录分类
func (obj *_EdTeacherGrowupMgr) GetFromType(_type int8) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 记录分类
func (obj *_EdTeacherGrowupMgr) GetBatchFromType(_types []int8) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_EdTeacherGrowupMgr) GetFromTitle(title string) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_EdTeacherGrowupMgr) GetBatchFromTitle(titles []string) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_EdTeacherGrowupMgr) GetFromContent(content *string) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 内容
func (obj *_EdTeacherGrowupMgr) GetBatchFromContent(contents []*string) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromFiles 通过files获取内容 附件
func (obj *_EdTeacherGrowupMgr) GetFromFiles(files *string) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`files` = ?", files).Find(&results).Error

	return
}

// GetBatchFromFiles 批量查找 附件
func (obj *_EdTeacherGrowupMgr) GetBatchFromFiles(filess []*string) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`files` IN (?)", filess).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdTeacherGrowupMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdTeacherGrowupMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdTeacherGrowupMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdTeacherGrowupMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdTeacherGrowupMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdTeacherGrowupMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdTeacherGrowupMgr) GetFromCreateBy(createBy *uint) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdTeacherGrowupMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdTeacherGrowupMgr) GetFromUpdateBy(updateBy *uint) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdTeacherGrowupMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdTeacherGrowupMgr) FetchByPrimaryKey(id int64) (result EdTeacherGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherGrowup{}).Where("`id` = ?", id).First(&result).Error

	return
}
