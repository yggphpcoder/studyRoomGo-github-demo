package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrSeatLockMgr struct {
	*_BaseMgr
}

// SrSeatLockMgr open func
func SrSeatLockMgr(db *gorm.DB) *_SrSeatLockMgr {
	if db == nil {
		panic(fmt.Errorf("SrSeatLockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrSeatLockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_seat_lock"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrSeatLockMgr) GetTableName() string {
	return "sr_seat_lock"
}

// Reset 重置gorm会话
func (obj *_SrSeatLockMgr) Reset() *_SrSeatLockMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrSeatLockMgr) Get() (result SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrSeatLockMgr) Gets() (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrSeatLockMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrSeatLockMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSeatID seat_id获取 座位id
func (obj *_SrSeatLockMgr) WithSeatID(seatID int64) Option {
	return optionFunc(func(o *options) { o.query["seat_id"] = seatID })
}

// WithMemberID member_id获取 会员id
func (obj *_SrSeatLockMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrSeatLockMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopAreaID shop_area_id获取 店铺分区id
func (obj *_SrSeatLockMgr) WithShopAreaID(shopAreaID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_area_id"] = shopAreaID })
}

// WithLockStartDay lock_start_day获取 开始日期
func (obj *_SrSeatLockMgr) WithLockStartDay(lockStartDay *time.Time) Option {
	return optionFunc(func(o *options) { o.query["lock_start_day"] = lockStartDay })
}

// WithLockStart lock_start获取 开始时间
func (obj *_SrSeatLockMgr) WithLockStart(lockStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["lock_start"] = lockStart })
}

// WithLockEndDay lock_end_day获取 结束日期
func (obj *_SrSeatLockMgr) WithLockEndDay(lockEndDay *time.Time) Option {
	return optionFunc(func(o *options) { o.query["lock_end_day"] = lockEndDay })
}

// WithLockEnd lock_end获取 结束时间
func (obj *_SrSeatLockMgr) WithLockEnd(lockEnd time.Time) Option {
	return optionFunc(func(o *options) { o.query["lock_end"] = lockEnd })
}

// WithStatus status获取 锁定状态
func (obj *_SrSeatLockMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_SrSeatLockMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrSeatLockMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrSeatLockMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrSeatLockMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrSeatLockMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUseTime use_time获取 使用时段
func (obj *_SrSeatLockMgr) WithUseTime(useTime string) Option {
	return optionFunc(func(o *options) { o.query["use_time"] = useTime })
}

// WithMemberPackageCardID member_package_card_id获取 会员套餐id
func (obj *_SrSeatLockMgr) WithMemberPackageCardID(memberPackageCardID int64) Option {
	return optionFunc(func(o *options) { o.query["member_package_card_id"] = memberPackageCardID })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_SrSeatLockMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithCanControl can_control获取 是否可以控制门灯
func (obj *_SrSeatLockMgr) WithCanControl(canControl int8) Option {
	return optionFunc(func(o *options) { o.query["can_control"] = canControl })
}

// GetByOption 功能选项模式获取
func (obj *_SrSeatLockMgr) GetByOption(opts ...Option) (result SrSeatLock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrSeatLockMgr) GetByOptions(opts ...Option) (results []*SrSeatLock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrSeatLockMgr) GetFromID(id int64) (result SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrSeatLockMgr) GetBatchFromID(ids []int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSeatID 通过seat_id获取内容 座位id
func (obj *_SrSeatLockMgr) GetFromSeatID(seatID int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`seat_id` = ?", seatID).Find(&results).Error

	return
}

// GetBatchFromSeatID 批量查找 座位id
func (obj *_SrSeatLockMgr) GetBatchFromSeatID(seatIDs []int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`seat_id` IN (?)", seatIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrSeatLockMgr) GetFromMemberID(memberID int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrSeatLockMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrSeatLockMgr) GetFromShopID(shopID int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrSeatLockMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopAreaID 通过shop_area_id获取内容 店铺分区id
func (obj *_SrSeatLockMgr) GetFromShopAreaID(shopAreaID int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`shop_area_id` = ?", shopAreaID).Find(&results).Error

	return
}

// GetBatchFromShopAreaID 批量查找 店铺分区id
func (obj *_SrSeatLockMgr) GetBatchFromShopAreaID(shopAreaIDs []int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`shop_area_id` IN (?)", shopAreaIDs).Find(&results).Error

	return
}

// GetFromLockStartDay 通过lock_start_day获取内容 开始日期
func (obj *_SrSeatLockMgr) GetFromLockStartDay(lockStartDay *time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`lock_start_day` = ?", lockStartDay).Find(&results).Error

	return
}

// GetBatchFromLockStartDay 批量查找 开始日期
func (obj *_SrSeatLockMgr) GetBatchFromLockStartDay(lockStartDays []*time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`lock_start_day` IN (?)", lockStartDays).Find(&results).Error

	return
}

// GetFromLockStart 通过lock_start获取内容 开始时间
func (obj *_SrSeatLockMgr) GetFromLockStart(lockStart time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`lock_start` = ?", lockStart).Find(&results).Error

	return
}

// GetBatchFromLockStart 批量查找 开始时间
func (obj *_SrSeatLockMgr) GetBatchFromLockStart(lockStarts []time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`lock_start` IN (?)", lockStarts).Find(&results).Error

	return
}

// GetFromLockEndDay 通过lock_end_day获取内容 结束日期
func (obj *_SrSeatLockMgr) GetFromLockEndDay(lockEndDay *time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`lock_end_day` = ?", lockEndDay).Find(&results).Error

	return
}

// GetBatchFromLockEndDay 批量查找 结束日期
func (obj *_SrSeatLockMgr) GetBatchFromLockEndDay(lockEndDays []*time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`lock_end_day` IN (?)", lockEndDays).Find(&results).Error

	return
}

// GetFromLockEnd 通过lock_end获取内容 结束时间
func (obj *_SrSeatLockMgr) GetFromLockEnd(lockEnd time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`lock_end` = ?", lockEnd).Find(&results).Error

	return
}

// GetBatchFromLockEnd 批量查找 结束时间
func (obj *_SrSeatLockMgr) GetBatchFromLockEnd(lockEnds []time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`lock_end` IN (?)", lockEnds).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 锁定状态
func (obj *_SrSeatLockMgr) GetFromStatus(status int8) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 锁定状态
func (obj *_SrSeatLockMgr) GetBatchFromStatus(statuss []int8) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrSeatLockMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrSeatLockMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrSeatLockMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrSeatLockMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrSeatLockMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrSeatLockMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrSeatLockMgr) GetFromCreateBy(createBy *uint) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrSeatLockMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrSeatLockMgr) GetFromUpdateBy(updateBy *uint) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrSeatLockMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromUseTime 通过use_time获取内容 使用时段
func (obj *_SrSeatLockMgr) GetFromUseTime(useTime string) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`use_time` = ?", useTime).Find(&results).Error

	return
}

// GetBatchFromUseTime 批量查找 使用时段
func (obj *_SrSeatLockMgr) GetBatchFromUseTime(useTimes []string) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`use_time` IN (?)", useTimes).Find(&results).Error

	return
}

// GetFromMemberPackageCardID 通过member_package_card_id获取内容 会员套餐id
func (obj *_SrSeatLockMgr) GetFromMemberPackageCardID(memberPackageCardID int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`member_package_card_id` = ?", memberPackageCardID).Find(&results).Error

	return
}

// GetBatchFromMemberPackageCardID 批量查找 会员套餐id
func (obj *_SrSeatLockMgr) GetBatchFromMemberPackageCardID(memberPackageCardIDs []int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`member_package_card_id` IN (?)", memberPackageCardIDs).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_SrSeatLockMgr) GetFromMerchantID(merchantID int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_SrSeatLockMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromCanControl 通过can_control获取内容 是否可以控制门灯
func (obj *_SrSeatLockMgr) GetFromCanControl(canControl int8) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`can_control` = ?", canControl).Find(&results).Error

	return
}

// GetBatchFromCanControl 批量查找 是否可以控制门灯
func (obj *_SrSeatLockMgr) GetBatchFromCanControl(canControls []int8) (results []*SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`can_control` IN (?)", canControls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrSeatLockMgr) FetchByPrimaryKey(id int64) (result SrSeatLock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatLock{}).Where("`id` = ?", id).First(&result).Error

	return
}
