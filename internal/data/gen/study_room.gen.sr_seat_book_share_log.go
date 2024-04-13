package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrSeatBookShareLogMgr struct {
	*_BaseMgr
}

// SrSeatBookShareLogMgr open func
func SrSeatBookShareLogMgr(db *gorm.DB) *_SrSeatBookShareLogMgr {
	if db == nil {
		panic(fmt.Errorf("SrSeatBookShareLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrSeatBookShareLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_seat_book_share_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrSeatBookShareLogMgr) GetTableName() string {
	return "sr_seat_book_share_log"
}

// Reset 重置gorm会话
func (obj *_SrSeatBookShareLogMgr) Reset() *_SrSeatBookShareLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrSeatBookShareLogMgr) Get() (result SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrSeatBookShareLogMgr) Gets() (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrSeatBookShareLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrSeatBookShareLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSeatBookID seat_book_id获取 订座id
func (obj *_SrSeatBookShareLogMgr) WithSeatBookID(seatBookID int64) Option {
	return optionFunc(func(o *options) { o.query["seat_book_id"] = seatBookID })
}

// WithOpenid openid获取 微信用户唯一标识符
func (obj *_SrSeatBookShareLogMgr) WithOpenid(openid string) Option {
	return optionFunc(func(o *options) { o.query["openid"] = openid })
}

// WithMemberID member_id获取 会员id
func (obj *_SrSeatBookShareLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithRemark remark获取 备注
func (obj *_SrSeatBookShareLogMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrSeatBookShareLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrSeatBookShareLogMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrSeatBookShareLogMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrSeatBookShareLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrSeatBookShareLogMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrSeatBookShareLogMgr) GetByOption(opts ...Option) (result SrSeatBookShareLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrSeatBookShareLogMgr) GetByOptions(opts ...Option) (results []*SrSeatBookShareLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrSeatBookShareLogMgr) GetFromID(id int64) (result SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrSeatBookShareLogMgr) GetBatchFromID(ids []int64) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSeatBookID 通过seat_book_id获取内容 订座id
func (obj *_SrSeatBookShareLogMgr) GetFromSeatBookID(seatBookID int64) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`seat_book_id` = ?", seatBookID).Find(&results).Error

	return
}

// GetBatchFromSeatBookID 批量查找 订座id
func (obj *_SrSeatBookShareLogMgr) GetBatchFromSeatBookID(seatBookIDs []int64) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`seat_book_id` IN (?)", seatBookIDs).Find(&results).Error

	return
}

// GetFromOpenid 通过openid获取内容 微信用户唯一标识符
func (obj *_SrSeatBookShareLogMgr) GetFromOpenid(openid string) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`openid` = ?", openid).Find(&results).Error

	return
}

// GetBatchFromOpenid 批量查找 微信用户唯一标识符
func (obj *_SrSeatBookShareLogMgr) GetBatchFromOpenid(openids []string) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`openid` IN (?)", openids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrSeatBookShareLogMgr) GetFromMemberID(memberID int64) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrSeatBookShareLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrSeatBookShareLogMgr) GetFromRemark(remark *string) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrSeatBookShareLogMgr) GetBatchFromRemark(remarks []*string) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrSeatBookShareLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrSeatBookShareLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrSeatBookShareLogMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrSeatBookShareLogMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrSeatBookShareLogMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrSeatBookShareLogMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrSeatBookShareLogMgr) GetFromCreateBy(createBy *uint) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrSeatBookShareLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrSeatBookShareLogMgr) GetFromUpdateBy(updateBy *uint) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrSeatBookShareLogMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrSeatBookShareLogMgr) FetchByPrimaryKey(id int64) (result SrSeatBookShareLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBookShareLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
