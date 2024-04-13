package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrMemberMgr struct {
	*_BaseMgr
}

// SrMemberMgr open func
func SrMemberMgr(db *gorm.DB) *_SrMemberMgr {
	if db == nil {
		panic(fmt.Errorf("SrMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrMemberMgr) GetTableName() string {
	return "sr_member"
}

// Reset 重置gorm会话
func (obj *_SrMemberMgr) Reset() *_SrMemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrMemberMgr) Get() (result SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrMemberMgr) Gets() (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrMemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrMemberMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAvatarPath avatar_path获取 头像路径
func (obj *_SrMemberMgr) WithAvatarPath(avatarPath *string) Option {
	return optionFunc(func(o *options) { o.query["avatar_path"] = avatarPath })
}

// WithUsername username获取 用户名
func (obj *_SrMemberMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码
func (obj *_SrMemberMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithNickName nick_name获取 昵称
func (obj *_SrMemberMgr) WithNickName(nickName *string) Option {
	return optionFunc(func(o *options) { o.query["nick_name"] = nickName })
}

// WithPhone phone获取 手机号
func (obj *_SrMemberMgr) WithPhone(phone *string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithType type获取 会员类型
func (obj *_SrMemberMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithStatus status获取 会员状态
func (obj *_SrMemberMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRegisterAt register_at获取 创建时间
func (obj *_SrMemberMgr) WithRegisterAt(registerAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["register_at"] = registerAt })
}

// WithRemark remark获取 备注
func (obj *_SrMemberMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrMemberMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrMemberMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrMemberMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrMemberMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrMemberMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrMemberMgr) GetByOption(opts ...Option) (result SrMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrMemberMgr) GetByOptions(opts ...Option) (results []*SrMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrMemberMgr) GetFromID(id int64) (result SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrMemberMgr) GetBatchFromID(ids []int64) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAvatarPath 通过avatar_path获取内容 头像路径
func (obj *_SrMemberMgr) GetFromAvatarPath(avatarPath *string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`avatar_path` = ?", avatarPath).Find(&results).Error

	return
}

// GetBatchFromAvatarPath 批量查找 头像路径
func (obj *_SrMemberMgr) GetBatchFromAvatarPath(avatarPaths []*string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`avatar_path` IN (?)", avatarPaths).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_SrMemberMgr) GetFromUsername(username string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 用户名
func (obj *_SrMemberMgr) GetBatchFromUsername(usernames []string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_SrMemberMgr) GetFromPassword(password string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_SrMemberMgr) GetBatchFromPassword(passwords []string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromNickName 通过nick_name获取内容 昵称
func (obj *_SrMemberMgr) GetFromNickName(nickName *string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`nick_name` = ?", nickName).Find(&results).Error

	return
}

// GetBatchFromNickName 批量查找 昵称
func (obj *_SrMemberMgr) GetBatchFromNickName(nickNames []*string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`nick_name` IN (?)", nickNames).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机号
func (obj *_SrMemberMgr) GetFromPhone(phone *string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 手机号
func (obj *_SrMemberMgr) GetBatchFromPhone(phones []*string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 会员类型
func (obj *_SrMemberMgr) GetFromType(_type int8) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 会员类型
func (obj *_SrMemberMgr) GetBatchFromType(_types []int8) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 会员状态
func (obj *_SrMemberMgr) GetFromStatus(status int8) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 会员状态
func (obj *_SrMemberMgr) GetBatchFromStatus(statuss []int8) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRegisterAt 通过register_at获取内容 创建时间
func (obj *_SrMemberMgr) GetFromRegisterAt(registerAt *time.Time) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`register_at` = ?", registerAt).Find(&results).Error

	return
}

// GetBatchFromRegisterAt 批量查找 创建时间
func (obj *_SrMemberMgr) GetBatchFromRegisterAt(registerAts []*time.Time) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`register_at` IN (?)", registerAts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrMemberMgr) GetFromRemark(remark *string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrMemberMgr) GetBatchFromRemark(remarks []*string) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrMemberMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrMemberMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrMemberMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrMemberMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrMemberMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrMemberMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrMemberMgr) GetFromCreateBy(createBy *uint) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrMemberMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrMemberMgr) GetFromUpdateBy(updateBy *uint) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrMemberMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrMemberMgr) FetchByPrimaryKey(id int64) (result SrMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMember{}).Where("`id` = ?", id).First(&result).Error

	return
}
