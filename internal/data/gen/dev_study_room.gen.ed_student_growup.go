package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdStudentGrowupMgr struct {
	*_BaseMgr
}

// EdStudentGrowupMgr open func
func EdStudentGrowupMgr(db *gorm.DB) *_EdStudentGrowupMgr {
	if db == nil {
		panic(fmt.Errorf("EdStudentGrowupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdStudentGrowupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_student_growup"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdStudentGrowupMgr) GetTableName() string {
	return "ed_student_growup"
}

// Reset 重置gorm会话
func (obj *_EdStudentGrowupMgr) Reset() *_EdStudentGrowupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdStudentGrowupMgr) Get() (result EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdStudentGrowupMgr) Gets() (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdStudentGrowupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdStudentGrowupMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStudentID student_id获取 学生id
func (obj *_EdStudentGrowupMgr) WithStudentID(studentID int64) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithType type获取 记录分类
func (obj *_EdStudentGrowupMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithTitle title获取 标题
func (obj *_EdStudentGrowupMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_EdStudentGrowupMgr) WithContent(content *string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithFiles files获取 附件
func (obj *_EdStudentGrowupMgr) WithFiles(files *string) Option {
	return optionFunc(func(o *options) { o.query["files"] = files })
}

// WithCreatedAt created_at获取
func (obj *_EdStudentGrowupMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdStudentGrowupMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdStudentGrowupMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdStudentGrowupMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdStudentGrowupMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdStudentGrowupMgr) GetByOption(opts ...Option) (result EdStudentGrowup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdStudentGrowupMgr) GetByOptions(opts ...Option) (results []*EdStudentGrowup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdStudentGrowupMgr) GetFromID(id int64) (result EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdStudentGrowupMgr) GetBatchFromID(ids []int64) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_EdStudentGrowupMgr) GetFromStudentID(studentID int64) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`student_id` = ?", studentID).Find(&results).Error

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_EdStudentGrowupMgr) GetBatchFromStudentID(studentIDs []int64) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 记录分类
func (obj *_EdStudentGrowupMgr) GetFromType(_type int8) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 记录分类
func (obj *_EdStudentGrowupMgr) GetBatchFromType(_types []int8) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_EdStudentGrowupMgr) GetFromTitle(title string) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_EdStudentGrowupMgr) GetBatchFromTitle(titles []string) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_EdStudentGrowupMgr) GetFromContent(content *string) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 内容
func (obj *_EdStudentGrowupMgr) GetBatchFromContent(contents []*string) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromFiles 通过files获取内容 附件
func (obj *_EdStudentGrowupMgr) GetFromFiles(files *string) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`files` = ?", files).Find(&results).Error

	return
}

// GetBatchFromFiles 批量查找 附件
func (obj *_EdStudentGrowupMgr) GetBatchFromFiles(filess []*string) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`files` IN (?)", filess).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdStudentGrowupMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdStudentGrowupMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdStudentGrowupMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdStudentGrowupMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdStudentGrowupMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdStudentGrowupMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdStudentGrowupMgr) GetFromCreateBy(createBy *uint) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdStudentGrowupMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdStudentGrowupMgr) GetFromUpdateBy(updateBy *uint) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdStudentGrowupMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdStudentGrowupMgr) FetchByPrimaryKey(id int64) (result EdStudentGrowup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentGrowup{}).Where("`id` = ?", id).First(&result).Error

	return
}
