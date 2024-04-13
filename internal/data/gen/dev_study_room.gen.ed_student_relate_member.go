package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _EdStudentRelateMemberMgr struct {
	*_BaseMgr
}

// EdStudentRelateMemberMgr open func
func EdStudentRelateMemberMgr(db *gorm.DB) *_EdStudentRelateMemberMgr {
	if db == nil {
		panic(fmt.Errorf("EdStudentRelateMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdStudentRelateMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_student_relate_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdStudentRelateMemberMgr) GetTableName() string {
	return "ed_student_relate_member"
}

// Reset 重置gorm会话
func (obj *_EdStudentRelateMemberMgr) Reset() *_EdStudentRelateMemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdStudentRelateMemberMgr) Get() (result EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdStudentRelateMemberMgr) Gets() (results []*EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdStudentRelateMemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdStudentRelateMemberMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *_EdStudentRelateMemberMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithStudentID student_id获取 学生id
func (obj *_EdStudentRelateMemberMgr) WithStudentID(studentID int64) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// GetByOption 功能选项模式获取
func (obj *_EdStudentRelateMemberMgr) GetByOption(opts ...Option) (result EdStudentRelateMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdStudentRelateMemberMgr) GetByOptions(opts ...Option) (results []*EdStudentRelateMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdStudentRelateMemberMgr) GetFromID(id int64) (result EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdStudentRelateMemberMgr) GetBatchFromID(ids []int64) (results []*EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EdStudentRelateMemberMgr) GetFromMemberID(memberID int64) (results []*EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EdStudentRelateMemberMgr) GetBatchFromMemberID(memberIDs []int64) (results []*EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_EdStudentRelateMemberMgr) GetFromStudentID(studentID int64) (results []*EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where("`student_id` = ?", studentID).Find(&results).Error

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_EdStudentRelateMemberMgr) GetBatchFromStudentID(studentIDs []int64) (results []*EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdStudentRelateMemberMgr) FetchByPrimaryKey(id int64) (result EdStudentRelateMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentRelateMember{}).Where("`id` = ?", id).First(&result).Error

	return
}
