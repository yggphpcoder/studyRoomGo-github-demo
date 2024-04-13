package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrSeatBookOperateLogMgr struct {
	*_BaseMgr
}

// SrSeatBookOperateLogMgr open func
func SrSeatBookOperateLogMgr(db *gorm.DB) *_SrSeatBookOperateLogMgr {
	if db == nil {
		panic(fmt.Errorf("SrSeatBookOperateLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrSeatBookOperateLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_seat_book_operate_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrSeatBookOperateLogMgr) GetTableName() string {
	return "sr_seat_book_operate_log"
}

// Reset 重置gorm会话
func (obj *_SrSeatBookOperateLogMgr) Reset() *_SrSeatBookOperateLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrSeatBookOperateLogMgr) Get() (result SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrSeatBookOperateLogMgr) Gets() (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrSeatBookOperateLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrSeatBookOperateLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSeatBookID seat_book_id获取 订座id
func (obj *_SrSeatBookOperateLogMgr) WithSeatBookID(seatBookID int64) Option {
	return optionFunc(func(o *options) { o.query["seat_book_id"] = seatBookID })
}

// WithMemberID member_id获取 会员id
func (obj *_SrSeatBookOperateLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithSeatID seat_id获取 座位id
func (obj *_SrSeatBookOperateLogMgr) WithSeatID(seatID int64) Option {
	return optionFunc(func(o *options) { o.query["seat_id"] = seatID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrSeatBookOperateLogMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithAction action获取 操作类型
func (obj *_SrSeatBookOperateLogMgr) WithAction(action int8) Option {
	return optionFunc(func(o *options) { o.query["action"] = action })
}

// WithExtra extra获取 额外信息
func (obj *_SrSeatBookOperateLogMgr) WithExtra(extra *string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithRemark remark获取 备注
func (obj *_SrSeatBookOperateLogMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrSeatBookOperateLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreateBy create_by获取
func (obj *_SrSeatBookOperateLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrSeatBookOperateLogMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrSeatBookOperateLogMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithUpdateBy update_by获取
func (obj *_SrSeatBookOperateLogMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrSeatBookOperateLogMgr) GetByOption(opts ...Option) (result SrSeatBookOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrSeatBookOperateLogMgr) GetByOptions(opts ...Option) (results []*SrSeatBookOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrSeatBookOperateLogMgr) GetFromID(id int64) (result SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromID(ids []int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSeatBookID 通过seat_book_id获取内容 订座id
func (obj *_SrSeatBookOperateLogMgr) GetFromSeatBookID(seatBookID int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`seat_book_id` = ?", seatBookID).Find(&results).Error

	return
}

// GetBatchFromSeatBookID 批量查找 订座id
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromSeatBookID(seatBookIDs []int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`seat_book_id` IN (?)", seatBookIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrSeatBookOperateLogMgr) GetFromMemberID(memberID int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromSeatID 通过seat_id获取内容 座位id
func (obj *_SrSeatBookOperateLogMgr) GetFromSeatID(seatID int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`seat_id` = ?", seatID).Find(&results).Error

	return
}

// GetBatchFromSeatID 批量查找 座位id
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromSeatID(seatIDs []int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`seat_id` IN (?)", seatIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrSeatBookOperateLogMgr) GetFromShopID(shopID int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromAction 通过action获取内容 操作类型
func (obj *_SrSeatBookOperateLogMgr) GetFromAction(action int8) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`action` = ?", action).Find(&results).Error

	return
}

// GetBatchFromAction 批量查找 操作类型
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromAction(actions []int8) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`action` IN (?)", actions).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 额外信息
func (obj *_SrSeatBookOperateLogMgr) GetFromExtra(extra *string) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 额外信息
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromExtra(extras []*string) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrSeatBookOperateLogMgr) GetFromRemark(remark *string) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromRemark(remarks []*string) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrSeatBookOperateLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrSeatBookOperateLogMgr) GetFromCreateBy(createBy *uint) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrSeatBookOperateLogMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrSeatBookOperateLogMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrSeatBookOperateLogMgr) GetFromUpdateBy(updateBy *uint) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrSeatBookOperateLogMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrSeatBookOperateLogMgr) FetchByPrimaryKey(id int64) (result SrSeatBookOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
