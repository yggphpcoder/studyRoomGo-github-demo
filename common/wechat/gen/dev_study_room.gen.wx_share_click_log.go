package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxShareClickLogMgr struct {
	*_BaseMgr
}

// WxShareClickLogMgr open func
func WxShareClickLogMgr(db *gorm.DB) *_WxShareClickLogMgr {
	if db == nil {
		panic(fmt.Errorf("WxShareClickLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxShareClickLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_share_click_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxShareClickLogMgr) GetTableName() string {
	return "wx_share_click_log"
}

// Reset 重置gorm会话
func (obj *_WxShareClickLogMgr) Reset() *_WxShareClickLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxShareClickLogMgr) Get() (result WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxShareClickLogMgr) Gets() (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxShareClickLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_WxShareClickLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithShareOpenID share_open_id获取 分享者openId
func (obj *_WxShareClickLogMgr) WithShareOpenID(shareOpenID string) Option {
	return optionFunc(func(o *options) { o.query["share_open_id"] = shareOpenID })
}

// WithShareMemberID share_member_id获取 分享者会员id
func (obj *_WxShareClickLogMgr) WithShareMemberID(shareMemberID int64) Option {
	return optionFunc(func(o *options) { o.query["share_member_id"] = shareMemberID })
}

// WithShareType share_type获取 分享类型
func (obj *_WxShareClickLogMgr) WithShareType(shareType string) Option {
	return optionFunc(func(o *options) { o.query["share_type"] = shareType })
}

// WithShareRelateID share_relate_id获取 分享关联数据id
func (obj *_WxShareClickLogMgr) WithShareRelateID(shareRelateID int64) Option {
	return optionFunc(func(o *options) { o.query["share_relate_id"] = shareRelateID })
}

// WithOpenID open_id获取 微信用户唯一标识符
func (obj *_WxShareClickLogMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithMemberID member_id获取 会员id
func (obj *_WxShareClickLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithRemark remark获取 备注
func (obj *_WxShareClickLogMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_WxShareClickLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_WxShareClickLogMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_WxShareClickLogMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_WxShareClickLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_WxShareClickLogMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_WxShareClickLogMgr) GetByOption(opts ...Option) (result WxShareClickLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxShareClickLogMgr) GetByOptions(opts ...Option) (results []*WxShareClickLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_WxShareClickLogMgr) GetFromID(id int64) (result WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_WxShareClickLogMgr) GetBatchFromID(ids []int64) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromShareOpenID 通过share_open_id获取内容 分享者openId
func (obj *_WxShareClickLogMgr) GetFromShareOpenID(shareOpenID string) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`share_open_id` = ?", shareOpenID).Find(&results).Error

	return
}

// GetBatchFromShareOpenID 批量查找 分享者openId
func (obj *_WxShareClickLogMgr) GetBatchFromShareOpenID(shareOpenIDs []string) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`share_open_id` IN (?)", shareOpenIDs).Find(&results).Error

	return
}

// GetFromShareMemberID 通过share_member_id获取内容 分享者会员id
func (obj *_WxShareClickLogMgr) GetFromShareMemberID(shareMemberID int64) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`share_member_id` = ?", shareMemberID).Find(&results).Error

	return
}

// GetBatchFromShareMemberID 批量查找 分享者会员id
func (obj *_WxShareClickLogMgr) GetBatchFromShareMemberID(shareMemberIDs []int64) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`share_member_id` IN (?)", shareMemberIDs).Find(&results).Error

	return
}

// GetFromShareType 通过share_type获取内容 分享类型
func (obj *_WxShareClickLogMgr) GetFromShareType(shareType string) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`share_type` = ?", shareType).Find(&results).Error

	return
}

// GetBatchFromShareType 批量查找 分享类型
func (obj *_WxShareClickLogMgr) GetBatchFromShareType(shareTypes []string) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`share_type` IN (?)", shareTypes).Find(&results).Error

	return
}

// GetFromShareRelateID 通过share_relate_id获取内容 分享关联数据id
func (obj *_WxShareClickLogMgr) GetFromShareRelateID(shareRelateID int64) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`share_relate_id` = ?", shareRelateID).Find(&results).Error

	return
}

// GetBatchFromShareRelateID 批量查找 分享关联数据id
func (obj *_WxShareClickLogMgr) GetBatchFromShareRelateID(shareRelateIDs []int64) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`share_relate_id` IN (?)", shareRelateIDs).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 微信用户唯一标识符
func (obj *_WxShareClickLogMgr) GetFromOpenID(openID string) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`open_id` = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量查找 微信用户唯一标识符
func (obj *_WxShareClickLogMgr) GetBatchFromOpenID(openIDs []string) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`open_id` IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_WxShareClickLogMgr) GetFromMemberID(memberID int64) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_WxShareClickLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WxShareClickLogMgr) GetFromRemark(remark *string) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WxShareClickLogMgr) GetBatchFromRemark(remarks []*string) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_WxShareClickLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_WxShareClickLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_WxShareClickLogMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_WxShareClickLogMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_WxShareClickLogMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_WxShareClickLogMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_WxShareClickLogMgr) GetFromCreateBy(createBy *uint) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_WxShareClickLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_WxShareClickLogMgr) GetFromUpdateBy(updateBy *uint) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_WxShareClickLogMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxShareClickLogMgr) FetchByPrimaryKey(id int64) (result WxShareClickLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxShareClickLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
