package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _EdTeacherRelateMemberMgr struct {
	*_BaseMgr
}

// EdTeacherRelateMemberMgr open func
func EdTeacherRelateMemberMgr(db *gorm.DB) *_EdTeacherRelateMemberMgr {
	if db == nil {
		panic(fmt.Errorf("EdTeacherRelateMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdTeacherRelateMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_teacher_relate_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdTeacherRelateMemberMgr) GetTableName() string {
	return "ed_teacher_relate_member"
}

// Reset 重置gorm会话
func (obj *_EdTeacherRelateMemberMgr) Reset() *_EdTeacherRelateMemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdTeacherRelateMemberMgr) Get() (result EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdTeacherRelateMemberMgr) Gets() (results []*EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdTeacherRelateMemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdTeacherRelateMemberMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *_EdTeacherRelateMemberMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithTeacherID teacher_id获取 老师id
func (obj *_EdTeacherRelateMemberMgr) WithTeacherID(teacherID int64) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// GetByOption 功能选项模式获取
func (obj *_EdTeacherRelateMemberMgr) GetByOption(opts ...Option) (result EdTeacherRelateMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdTeacherRelateMemberMgr) GetByOptions(opts ...Option) (results []*EdTeacherRelateMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdTeacherRelateMemberMgr) GetFromID(id int64) (result EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdTeacherRelateMemberMgr) GetBatchFromID(ids []int64) (results []*EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EdTeacherRelateMemberMgr) GetFromMemberID(memberID int64) (results []*EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EdTeacherRelateMemberMgr) GetBatchFromMemberID(memberIDs []int64) (results []*EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 老师id
func (obj *_EdTeacherRelateMemberMgr) GetFromTeacherID(teacherID int64) (results []*EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where("`teacher_id` = ?", teacherID).Find(&results).Error

	return
}

// GetBatchFromTeacherID 批量查找 老师id
func (obj *_EdTeacherRelateMemberMgr) GetBatchFromTeacherID(teacherIDs []int64) (results []*EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdTeacherRelateMemberMgr) FetchByPrimaryKey(id int64) (result EdTeacherRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherRelateMember{}).Where("`id` = ?", id).First(&result).Error

	return
}
