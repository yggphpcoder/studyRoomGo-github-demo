package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrSeatBookMgr struct {
	*_BaseMgr
}

// SrSeatBookMgr open func
func SrSeatBookMgr(db *gorm.DB) *_SrSeatBookMgr {
	if db == nil {
		panic(fmt.Errorf("SrSeatBookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrSeatBookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_seat_book"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrSeatBookMgr) GetTableName() string {
	return "sr_seat_book"
}

// Reset 重置gorm会话
func (obj *_SrSeatBookMgr) Reset() *_SrSeatBookMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrSeatBookMgr) Get() (result SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrSeatBookMgr) Gets() (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrSeatBookMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrSeatBookMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_SrSeatBookMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithUniqueToken unique_token获取 唯一标签，标注每次相同的订座操作
func (obj *_SrSeatBookMgr) WithUniqueToken(uniqueToken *string) Option {
	return optionFunc(func(o *options) { o.query["unique_token"] = uniqueToken })
}

// WithSeatID seat_id获取 座位id
func (obj *_SrSeatBookMgr) WithSeatID(seatID int64) Option {
	return optionFunc(func(o *options) { o.query["seat_id"] = seatID })
}

// WithMemberID member_id获取 会员id
func (obj *_SrSeatBookMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberPackageCardID member_package_card_id获取 会员套餐id
func (obj *_SrSeatBookMgr) WithMemberPackageCardID(memberPackageCardID int64) Option {
	return optionFunc(func(o *options) { o.query["member_package_card_id"] = memberPackageCardID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrSeatBookMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopAreaID shop_area_id获取 店铺分区id
func (obj *_SrSeatBookMgr) WithShopAreaID(shopAreaID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_area_id"] = shopAreaID })
}

// WithBookingType booking_type获取 订座类型
func (obj *_SrSeatBookMgr) WithBookingType(bookingType int8) Option {
	return optionFunc(func(o *options) { o.query["booking_type"] = bookingType })
}

// WithBookingDay booking_day获取 订座天数
func (obj *_SrSeatBookMgr) WithBookingDay(bookingDay int) Option {
	return optionFunc(func(o *options) { o.query["booking_day"] = bookingDay })
}

// WithBookingMinute booking_minute获取 订座分钟数
func (obj *_SrSeatBookMgr) WithBookingMinute(bookingMinute int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute"] = bookingMinute })
}

// WithBookingStartDay booking_start_day获取 订座开始日期
func (obj *_SrSeatBookMgr) WithBookingStartDay(bookingStartDay time.Time) Option {
	return optionFunc(func(o *options) { o.query["booking_start_day"] = bookingStartDay })
}

// WithBookingStart booking_start获取 订座开始时间
func (obj *_SrSeatBookMgr) WithBookingStart(bookingStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["booking_start"] = bookingStart })
}

// WithBookingEndDay booking_end_day获取 预定结束日期
func (obj *_SrSeatBookMgr) WithBookingEndDay(bookingEndDay time.Time) Option {
	return optionFunc(func(o *options) { o.query["booking_end_day"] = bookingEndDay })
}

// WithBookingEnd booking_end获取 订座结束时间
func (obj *_SrSeatBookMgr) WithBookingEnd(bookingEnd time.Time) Option {
	return optionFunc(func(o *options) { o.query["booking_end"] = bookingEnd })
}

// WithTodaySeatAt today_seat_at获取 当天入座时间
func (obj *_SrSeatBookMgr) WithTodaySeatAt(todaySeatAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["today_seat_at"] = todaySeatAt })
}

// WithTodayEndAt today_end_at获取 当天结束时间
func (obj *_SrSeatBookMgr) WithTodayEndAt(todayEndAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["today_end_at"] = todayEndAt })
}

// WithStatus status获取 使用状态
func (obj *_SrSeatBookMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithAutomation automation获取 自动化操作
func (obj *_SrSeatBookMgr) WithAutomation(automation string) Option {
	return optionFunc(func(o *options) { o.query["automation"] = automation })
}

// WithCreatedAt created_at获取
func (obj *_SrSeatBookMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrSeatBookMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrSeatBookMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrSeatBookMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrSeatBookMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrSeatBookMgr) GetByOption(opts ...Option) (result SrSeatBook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrSeatBookMgr) GetByOptions(opts ...Option) (results []*SrSeatBook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrSeatBookMgr) GetFromID(id int64) (result SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrSeatBookMgr) GetBatchFromID(ids []int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_SrSeatBookMgr) GetFromMerchantID(merchantID int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_SrSeatBookMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromUniqueToken 通过unique_token获取内容 唯一标签，标注每次相同的订座操作
func (obj *_SrSeatBookMgr) GetFromUniqueToken(uniqueToken *string) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`unique_token` = ?", uniqueToken).Find(&results).Error

	return
}

// GetBatchFromUniqueToken 批量查找 唯一标签，标注每次相同的订座操作
func (obj *_SrSeatBookMgr) GetBatchFromUniqueToken(uniqueTokens []*string) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`unique_token` IN (?)", uniqueTokens).Find(&results).Error

	return
}

// GetFromSeatID 通过seat_id获取内容 座位id
func (obj *_SrSeatBookMgr) GetFromSeatID(seatID int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`seat_id` = ?", seatID).Find(&results).Error

	return
}

// GetBatchFromSeatID 批量查找 座位id
func (obj *_SrSeatBookMgr) GetBatchFromSeatID(seatIDs []int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`seat_id` IN (?)", seatIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrSeatBookMgr) GetFromMemberID(memberID int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrSeatBookMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberPackageCardID 通过member_package_card_id获取内容 会员套餐id
func (obj *_SrSeatBookMgr) GetFromMemberPackageCardID(memberPackageCardID int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`member_package_card_id` = ?", memberPackageCardID).Find(&results).Error

	return
}

// GetBatchFromMemberPackageCardID 批量查找 会员套餐id
func (obj *_SrSeatBookMgr) GetBatchFromMemberPackageCardID(memberPackageCardIDs []int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`member_package_card_id` IN (?)", memberPackageCardIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrSeatBookMgr) GetFromShopID(shopID int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrSeatBookMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopAreaID 通过shop_area_id获取内容 店铺分区id
func (obj *_SrSeatBookMgr) GetFromShopAreaID(shopAreaID int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`shop_area_id` = ?", shopAreaID).Find(&results).Error

	return
}

// GetBatchFromShopAreaID 批量查找 店铺分区id
func (obj *_SrSeatBookMgr) GetBatchFromShopAreaID(shopAreaIDs []int64) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`shop_area_id` IN (?)", shopAreaIDs).Find(&results).Error

	return
}

// GetFromBookingType 通过booking_type获取内容 订座类型
func (obj *_SrSeatBookMgr) GetFromBookingType(bookingType int8) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_type` = ?", bookingType).Find(&results).Error

	return
}

// GetBatchFromBookingType 批量查找 订座类型
func (obj *_SrSeatBookMgr) GetBatchFromBookingType(bookingTypes []int8) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_type` IN (?)", bookingTypes).Find(&results).Error

	return
}

// GetFromBookingDay 通过booking_day获取内容 订座天数
func (obj *_SrSeatBookMgr) GetFromBookingDay(bookingDay int) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_day` = ?", bookingDay).Find(&results).Error

	return
}

// GetBatchFromBookingDay 批量查找 订座天数
func (obj *_SrSeatBookMgr) GetBatchFromBookingDay(bookingDays []int) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_day` IN (?)", bookingDays).Find(&results).Error

	return
}

// GetFromBookingMinute 通过booking_minute获取内容 订座分钟数
func (obj *_SrSeatBookMgr) GetFromBookingMinute(bookingMinute int) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_minute` = ?", bookingMinute).Find(&results).Error

	return
}

// GetBatchFromBookingMinute 批量查找 订座分钟数
func (obj *_SrSeatBookMgr) GetBatchFromBookingMinute(bookingMinutes []int) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_minute` IN (?)", bookingMinutes).Find(&results).Error

	return
}

// GetFromBookingStartDay 通过booking_start_day获取内容 订座开始日期
func (obj *_SrSeatBookMgr) GetFromBookingStartDay(bookingStartDay time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_start_day` = ?", bookingStartDay).Find(&results).Error

	return
}

// GetBatchFromBookingStartDay 批量查找 订座开始日期
func (obj *_SrSeatBookMgr) GetBatchFromBookingStartDay(bookingStartDays []time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_start_day` IN (?)", bookingStartDays).Find(&results).Error

	return
}

// GetFromBookingStart 通过booking_start获取内容 订座开始时间
func (obj *_SrSeatBookMgr) GetFromBookingStart(bookingStart time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_start` = ?", bookingStart).Find(&results).Error

	return
}

// GetBatchFromBookingStart 批量查找 订座开始时间
func (obj *_SrSeatBookMgr) GetBatchFromBookingStart(bookingStarts []time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_start` IN (?)", bookingStarts).Find(&results).Error

	return
}

// GetFromBookingEndDay 通过booking_end_day获取内容 预定结束日期
func (obj *_SrSeatBookMgr) GetFromBookingEndDay(bookingEndDay time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_end_day` = ?", bookingEndDay).Find(&results).Error

	return
}

// GetBatchFromBookingEndDay 批量查找 预定结束日期
func (obj *_SrSeatBookMgr) GetBatchFromBookingEndDay(bookingEndDays []time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_end_day` IN (?)", bookingEndDays).Find(&results).Error

	return
}

// GetFromBookingEnd 通过booking_end获取内容 订座结束时间
func (obj *_SrSeatBookMgr) GetFromBookingEnd(bookingEnd time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_end` = ?", bookingEnd).Find(&results).Error

	return
}

// GetBatchFromBookingEnd 批量查找 订座结束时间
func (obj *_SrSeatBookMgr) GetBatchFromBookingEnd(bookingEnds []time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`booking_end` IN (?)", bookingEnds).Find(&results).Error

	return
}

// GetFromTodaySeatAt 通过today_seat_at获取内容 当天入座时间
func (obj *_SrSeatBookMgr) GetFromTodaySeatAt(todaySeatAt *time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`today_seat_at` = ?", todaySeatAt).Find(&results).Error

	return
}

// GetBatchFromTodaySeatAt 批量查找 当天入座时间
func (obj *_SrSeatBookMgr) GetBatchFromTodaySeatAt(todaySeatAts []*time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`today_seat_at` IN (?)", todaySeatAts).Find(&results).Error

	return
}

// GetFromTodayEndAt 通过today_end_at获取内容 当天结束时间
func (obj *_SrSeatBookMgr) GetFromTodayEndAt(todayEndAt *time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`today_end_at` = ?", todayEndAt).Find(&results).Error

	return
}

// GetBatchFromTodayEndAt 批量查找 当天结束时间
func (obj *_SrSeatBookMgr) GetBatchFromTodayEndAt(todayEndAts []*time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`today_end_at` IN (?)", todayEndAts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 使用状态
func (obj *_SrSeatBookMgr) GetFromStatus(status int8) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 使用状态
func (obj *_SrSeatBookMgr) GetBatchFromStatus(statuss []int8) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromAutomation 通过automation获取内容 自动化操作
func (obj *_SrSeatBookMgr) GetFromAutomation(automation string) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`automation` = ?", automation).Find(&results).Error

	return
}

// GetBatchFromAutomation 批量查找 自动化操作
func (obj *_SrSeatBookMgr) GetBatchFromAutomation(automations []string) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`automation` IN (?)", automations).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrSeatBookMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrSeatBookMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrSeatBookMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrSeatBookMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrSeatBookMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrSeatBookMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrSeatBookMgr) GetFromCreateBy(createBy *uint) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrSeatBookMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrSeatBookMgr) GetFromUpdateBy(updateBy *uint) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrSeatBookMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrSeatBookMgr) FetchByPrimaryKey(id int64) (result SrSeatBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeatBook{}).Where("`id` = ?", id).First(&result).Error

	return
}
