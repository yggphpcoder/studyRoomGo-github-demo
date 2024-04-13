package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrMemberPackageCardOperateLogMgr struct {
	*_BaseMgr
}

// SrMemberPackageCardOperateLogMgr open func
func SrMemberPackageCardOperateLogMgr(db *gorm.DB) *_SrMemberPackageCardOperateLogMgr {
	if db == nil {
		panic(fmt.Errorf("SrMemberPackageCardOperateLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrMemberPackageCardOperateLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_member_package_card_operate_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrMemberPackageCardOperateLogMgr) GetTableName() string {
	return "sr_member_package_card_operate_log"
}

// Reset 重置gorm会话
func (obj *_SrMemberPackageCardOperateLogMgr) Reset() *_SrMemberPackageCardOperateLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrMemberPackageCardOperateLogMgr) Get() (result SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrMemberPackageCardOperateLogMgr) Gets() (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrMemberPackageCardOperateLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrMemberPackageCardOperateLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMpID mp_id获取 会员-套餐id
func (obj *_SrMemberPackageCardOperateLogMgr) WithMpID(mpID int64) Option {
	return optionFunc(func(o *options) { o.query["mp_id"] = mpID })
}

// WithMemberID member_id获取 会员id
func (obj *_SrMemberPackageCardOperateLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithPackageCardID package_card_id获取 套餐id
func (obj *_SrMemberPackageCardOperateLogMgr) WithPackageCardID(packageCardID int64) Option {
	return optionFunc(func(o *options) { o.query["package_card_id"] = packageCardID })
}

// WithOperate operate获取 操作
func (obj *_SrMemberPackageCardOperateLogMgr) WithOperate(operate int8) Option {
	return optionFunc(func(o *options) { o.query["operate"] = operate })
}

// WithExtra extra获取 额外信息
func (obj *_SrMemberPackageCardOperateLogMgr) WithExtra(extra *string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithRemark remark获取 备注
func (obj *_SrMemberPackageCardOperateLogMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrMemberPackageCardOperateLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreateBy create_by获取
func (obj *_SrMemberPackageCardOperateLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrMemberPackageCardOperateLogMgr) GetByOption(opts ...Option) (result SrMemberPackageCardOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrMemberPackageCardOperateLogMgr) GetByOptions(opts ...Option) (results []*SrMemberPackageCardOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromID(id int64) (result SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromID(ids []int64) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMpID 通过mp_id获取内容 会员-套餐id
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromMpID(mpID int64) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`mp_id` = ?", mpID).Find(&results).Error

	return
}

// GetBatchFromMpID 批量查找 会员-套餐id
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromMpID(mpIDs []int64) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`mp_id` IN (?)", mpIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromMemberID(memberID int64) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromPackageCardID 通过package_card_id获取内容 套餐id
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromPackageCardID(packageCardID int64) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`package_card_id` = ?", packageCardID).Find(&results).Error

	return
}

// GetBatchFromPackageCardID 批量查找 套餐id
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromPackageCardID(packageCardIDs []int64) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`package_card_id` IN (?)", packageCardIDs).Find(&results).Error

	return
}

// GetFromOperate 通过operate获取内容 操作
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromOperate(operate int8) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`operate` = ?", operate).Find(&results).Error

	return
}

// GetBatchFromOperate 批量查找 操作
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromOperate(operates []int8) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`operate` IN (?)", operates).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 额外信息
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromExtra(extra *string) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 额外信息
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromExtra(extras []*string) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromRemark(remark *string) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromRemark(remarks []*string) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrMemberPackageCardOperateLogMgr) GetFromCreateBy(createBy *uint) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrMemberPackageCardOperateLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrMemberPackageCardOperateLogMgr) FetchByPrimaryKey(id int64) (result SrMemberPackageCardOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCardOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
