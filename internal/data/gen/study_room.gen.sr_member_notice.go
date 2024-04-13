package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrMemberNoticeMgr struct {
	*_BaseMgr
}

// SrMemberNoticeMgr open func
func SrMemberNoticeMgr(db *gorm.DB) *_SrMemberNoticeMgr {
	if db == nil {
		panic(fmt.Errorf("SrMemberNoticeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrMemberNoticeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_member_notice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrMemberNoticeMgr) GetTableName() string {
	return "sr_member_notice"
}

// Reset 重置gorm会话
func (obj *_SrMemberNoticeMgr) Reset() *_SrMemberNoticeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrMemberNoticeMgr) Get() (result SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrMemberNoticeMgr) Gets() (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrMemberNoticeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_SrMemberNoticeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *_SrMemberNoticeMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMerchantID merchant_id获取 创建商家id
func (obj *_SrMemberNoticeMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 创建店铺id
func (obj *_SrMemberNoticeMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithNoticeType notice_type获取 通知类型
func (obj *_SrMemberNoticeMgr) WithNoticeType(noticeType int8) Option {
	return optionFunc(func(o *options) { o.query["notice_type"] = noticeType })
}

// WithTitle title获取 标题
func (obj *_SrMemberNoticeMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_SrMemberNoticeMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithIsRead is_read获取 是否已读
func (obj *_SrMemberNoticeMgr) WithIsRead(isRead int8) Option {
	return optionFunc(func(o *options) { o.query["is_read"] = isRead })
}

// WithExtJSON ext_json获取 扩展内容
func (obj *_SrMemberNoticeMgr) WithExtJSON(extJSON string) Option {
	return optionFunc(func(o *options) { o.query["ext_json"] = extJSON })
}

// WithCreatedAt created_at获取
func (obj *_SrMemberNoticeMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrMemberNoticeMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrMemberNoticeMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrMemberNoticeMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrMemberNoticeMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrMemberNoticeMgr) GetByOption(opts ...Option) (result SrMemberNotice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrMemberNoticeMgr) GetByOptions(opts ...Option) (results []*SrMemberNotice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 id
func (obj *_SrMemberNoticeMgr) GetFromID(id int64) (result SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_SrMemberNoticeMgr) GetBatchFromID(ids []int64) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrMemberNoticeMgr) GetFromMemberID(memberID int64) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrMemberNoticeMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 创建商家id
func (obj *_SrMemberNoticeMgr) GetFromMerchantID(merchantID int64) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 创建商家id
func (obj *_SrMemberNoticeMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 创建店铺id
func (obj *_SrMemberNoticeMgr) GetFromShopID(shopID int64) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 创建店铺id
func (obj *_SrMemberNoticeMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromNoticeType 通过notice_type获取内容 通知类型
func (obj *_SrMemberNoticeMgr) GetFromNoticeType(noticeType int8) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`notice_type` = ?", noticeType).Find(&results).Error

	return
}

// GetBatchFromNoticeType 批量查找 通知类型
func (obj *_SrMemberNoticeMgr) GetBatchFromNoticeType(noticeTypes []int8) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`notice_type` IN (?)", noticeTypes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_SrMemberNoticeMgr) GetFromTitle(title string) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_SrMemberNoticeMgr) GetBatchFromTitle(titles []string) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_SrMemberNoticeMgr) GetFromContent(content string) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 内容
func (obj *_SrMemberNoticeMgr) GetBatchFromContent(contents []string) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromIsRead 通过is_read获取内容 是否已读
func (obj *_SrMemberNoticeMgr) GetFromIsRead(isRead int8) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`is_read` = ?", isRead).Find(&results).Error

	return
}

// GetBatchFromIsRead 批量查找 是否已读
func (obj *_SrMemberNoticeMgr) GetBatchFromIsRead(isReads []int8) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`is_read` IN (?)", isReads).Find(&results).Error

	return
}

// GetFromExtJSON 通过ext_json获取内容 扩展内容
func (obj *_SrMemberNoticeMgr) GetFromExtJSON(extJSON string) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`ext_json` = ?", extJSON).Find(&results).Error

	return
}

// GetBatchFromExtJSON 批量查找 扩展内容
func (obj *_SrMemberNoticeMgr) GetBatchFromExtJSON(extJSONs []string) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`ext_json` IN (?)", extJSONs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrMemberNoticeMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrMemberNoticeMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrMemberNoticeMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrMemberNoticeMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrMemberNoticeMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrMemberNoticeMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrMemberNoticeMgr) GetFromCreateBy(createBy *uint) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrMemberNoticeMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrMemberNoticeMgr) GetFromUpdateBy(updateBy *uint) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrMemberNoticeMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrMemberNoticeMgr) FetchByPrimaryKey(id int64) (result SrMemberNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberNotice{}).Where("`id` = ?", id).First(&result).Error

	return
}
