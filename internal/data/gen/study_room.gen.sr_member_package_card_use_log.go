package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrMemberPackageCardUseLogMgr struct {
	*_BaseMgr
}

// SrMemberPackageCardUseLogMgr open func
func SrMemberPackageCardUseLogMgr(db *gorm.DB) *_SrMemberPackageCardUseLogMgr {
	if db == nil {
		panic(fmt.Errorf("SrMemberPackageCardUseLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrMemberPackageCardUseLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_member_package_card_use_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrMemberPackageCardUseLogMgr) GetTableName() string {
	return "sr_member_package_card_use_log"
}

// Reset 重置gorm会话
func (obj *_SrMemberPackageCardUseLogMgr) Reset() *_SrMemberPackageCardUseLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrMemberPackageCardUseLogMgr) Get() (result SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrMemberPackageCardUseLogMgr) Gets() (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrMemberPackageCardUseLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrMemberPackageCardUseLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMpID mp_id获取 会员-套餐id
func (obj *_SrMemberPackageCardUseLogMgr) WithMpID(mpID int64) Option {
	return optionFunc(func(o *options) { o.query["mp_id"] = mpID })
}

// WithMemberID member_id获取 会员id
func (obj *_SrMemberPackageCardUseLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithValue value获取 值
func (obj *_SrMemberPackageCardUseLogMgr) WithValue(value int) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithExtra extra获取 额外信息
func (obj *_SrMemberPackageCardUseLogMgr) WithExtra(extra *string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithRemark remark获取 备注
func (obj *_SrMemberPackageCardUseLogMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrMemberPackageCardUseLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreateBy create_by获取
func (obj *_SrMemberPackageCardUseLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrMemberPackageCardUseLogMgr) GetByOption(opts ...Option) (result SrMemberPackageCardUseLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrMemberPackageCardUseLogMgr) GetByOptions(opts ...Option) (results []*SrMemberPackageCardUseLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrMemberPackageCardUseLogMgr) GetFromID(id int64) (result SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrMemberPackageCardUseLogMgr) GetBatchFromID(ids []int64) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMpID 通过mp_id获取内容 会员-套餐id
func (obj *_SrMemberPackageCardUseLogMgr) GetFromMpID(mpID int64) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`mp_id` = ?", mpID).Find(&results).Error

	return
}

// GetBatchFromMpID 批量查找 会员-套餐id
func (obj *_SrMemberPackageCardUseLogMgr) GetBatchFromMpID(mpIDs []int64) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`mp_id` IN (?)", mpIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrMemberPackageCardUseLogMgr) GetFromMemberID(memberID int64) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrMemberPackageCardUseLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 值
func (obj *_SrMemberPackageCardUseLogMgr) GetFromValue(value int) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找 值
func (obj *_SrMemberPackageCardUseLogMgr) GetBatchFromValue(values []int) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 额外信息
func (obj *_SrMemberPackageCardUseLogMgr) GetFromExtra(extra *string) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 额外信息
func (obj *_SrMemberPackageCardUseLogMgr) GetBatchFromExtra(extras []*string) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrMemberPackageCardUseLogMgr) GetFromRemark(remark *string) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrMemberPackageCardUseLogMgr) GetBatchFromRemark(remarks []*string) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrMemberPackageCardUseLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrMemberPackageCardUseLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrMemberPackageCardUseLogMgr) GetFromCreateBy(createBy *uint) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrMemberPackageCardUseLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrMemberPackageCardUseLogMgr) FetchByPrimaryKey(id int64) (result SrMemberPackageCardUseLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardUseLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
