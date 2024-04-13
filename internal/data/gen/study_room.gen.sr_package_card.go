package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrPackageCardMgr struct {
	*_BaseMgr
}

// SrPackageCardMgr open func
func SrPackageCardMgr(db *gorm.DB) *_SrPackageCardMgr {
	if db == nil {
		panic(fmt.Errorf("SrPackageCardMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrPackageCardMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_package_card"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrPackageCardMgr) GetTableName() string {
	return "sr_package_card"
}

// Reset 重置gorm会话
func (obj *_SrPackageCardMgr) Reset() *_SrPackageCardMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrPackageCardMgr) Get() (result SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrPackageCardMgr) Gets() (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrPackageCardMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrPackageCardMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_SrPackageCardMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithName name获取 套餐名
func (obj *_SrPackageCardMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTypePackage type_package获取 套餐类型
func (obj *_SrPackageCardMgr) WithTypePackage(typePackage int8) Option {
	return optionFunc(func(o *options) { o.query["type_package"] = typePackage })
}

// WithTypeCard type_card获取 卡类型
func (obj *_SrPackageCardMgr) WithTypeCard(typeCard int8) Option {
	return optionFunc(func(o *options) { o.query["type_card"] = typeCard })
}

// WithTypeSale type_sale获取 销售类型
func (obj *_SrPackageCardMgr) WithTypeSale(typeSale int8) Option {
	return optionFunc(func(o *options) { o.query["type_sale"] = typeSale })
}

// WithStatus status获取 状态
func (obj *_SrPackageCardMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPrice price获取 售价
func (obj *_SrPackageCardMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithOriPrice ori_price获取 原价
func (obj *_SrPackageCardMgr) WithOriPrice(oriPrice float64) Option {
	return optionFunc(func(o *options) { o.query["ori_price"] = oriPrice })
}

// WithValidDay valid_day获取 有效天数
func (obj *_SrPackageCardMgr) WithValidDay(validDay int) Option {
	return optionFunc(func(o *options) { o.query["valid_day"] = validDay })
}

// WithActiveLimit active_limit获取 激活期限天数
func (obj *_SrPackageCardMgr) WithActiveLimit(activeLimit int) Option {
	return optionFunc(func(o *options) { o.query["active_limit"] = activeLimit })
}

// WithBuyLimit buy_limit获取 每人限购次数
func (obj *_SrPackageCardMgr) WithBuyLimit(buyLimit int) Option {
	return optionFunc(func(o *options) { o.query["buy_limit"] = buyLimit })
}

// WithUseShop use_shop获取 适用店铺
func (obj *_SrPackageCardMgr) WithUseShop(useShop string) Option {
	return optionFunc(func(o *options) { o.query["use_shop"] = useShop })
}

// WithUseSeat use_seat获取 适用座位类型
func (obj *_SrPackageCardMgr) WithUseSeat(useSeat string) Option {
	return optionFunc(func(o *options) { o.query["use_seat"] = useSeat })
}

// WithUseTime use_time获取 使用时段
func (obj *_SrPackageCardMgr) WithUseTime(useTime string) Option {
	return optionFunc(func(o *options) { o.query["use_time"] = useTime })
}

// WithBookingDay booking_day获取 总可用天数
func (obj *_SrPackageCardMgr) WithBookingDay(bookingDay int) Option {
	return optionFunc(func(o *options) { o.query["booking_day"] = bookingDay })
}

// WithBookingMinute booking_minute获取 总可用分钟
func (obj *_SrPackageCardMgr) WithBookingMinute(bookingMinute int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute"] = bookingMinute })
}

// WithBookingMinutePer booking_minute_per获取 每天可用分钟
func (obj *_SrPackageCardMgr) WithBookingMinutePer(bookingMinutePer int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute_per"] = bookingMinutePer })
}

// WithBookingMinuteMax booking_minute_max获取 封顶可用分钟
func (obj *_SrPackageCardMgr) WithBookingMinuteMax(bookingMinuteMax int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute_max"] = bookingMinuteMax })
}

// WithBookingMinuteMin booking_minute_min获取 最低可用分钟
func (obj *_SrPackageCardMgr) WithBookingMinuteMin(bookingMinuteMin int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute_min"] = bookingMinuteMin })
}

// WithCanBalance can_balance获取 可否用余额购买
func (obj *_SrPackageCardMgr) WithCanBalance(canBalance *int8) Option {
	return optionFunc(func(o *options) { o.query["can_balance"] = canBalance })
}

// WithSort sort获取 排序
func (obj *_SrPackageCardMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrPackageCardMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrPackageCardMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrPackageCardMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrPackageCardMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrPackageCardMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrPackageCardMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrPackageCardMgr) GetByOption(opts ...Option) (result SrPackageCard, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrPackageCardMgr) GetByOptions(opts ...Option) (results []*SrPackageCard, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrPackageCardMgr) GetFromID(id int64) (result SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrPackageCardMgr) GetBatchFromID(ids []int64) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_SrPackageCardMgr) GetFromMerchantID(merchantID int64) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_SrPackageCardMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 套餐名
func (obj *_SrPackageCardMgr) GetFromName(name string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 套餐名
func (obj *_SrPackageCardMgr) GetBatchFromName(names []string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTypePackage 通过type_package获取内容 套餐类型
func (obj *_SrPackageCardMgr) GetFromTypePackage(typePackage int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`type_package` = ?", typePackage).Find(&results).Error

	return
}

// GetBatchFromTypePackage 批量查找 套餐类型
func (obj *_SrPackageCardMgr) GetBatchFromTypePackage(typePackages []int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`type_package` IN (?)", typePackages).Find(&results).Error

	return
}

// GetFromTypeCard 通过type_card获取内容 卡类型
func (obj *_SrPackageCardMgr) GetFromTypeCard(typeCard int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`type_card` = ?", typeCard).Find(&results).Error

	return
}

// GetBatchFromTypeCard 批量查找 卡类型
func (obj *_SrPackageCardMgr) GetBatchFromTypeCard(typeCards []int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`type_card` IN (?)", typeCards).Find(&results).Error

	return
}

// GetFromTypeSale 通过type_sale获取内容 销售类型
func (obj *_SrPackageCardMgr) GetFromTypeSale(typeSale int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`type_sale` = ?", typeSale).Find(&results).Error

	return
}

// GetBatchFromTypeSale 批量查找 销售类型
func (obj *_SrPackageCardMgr) GetBatchFromTypeSale(typeSales []int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`type_sale` IN (?)", typeSales).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SrPackageCardMgr) GetFromStatus(status int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SrPackageCardMgr) GetBatchFromStatus(statuss []int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 售价
func (obj *_SrPackageCardMgr) GetFromPrice(price float64) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 售价
func (obj *_SrPackageCardMgr) GetBatchFromPrice(prices []float64) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromOriPrice 通过ori_price获取内容 原价
func (obj *_SrPackageCardMgr) GetFromOriPrice(oriPrice float64) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`ori_price` = ?", oriPrice).Find(&results).Error

	return
}

// GetBatchFromOriPrice 批量查找 原价
func (obj *_SrPackageCardMgr) GetBatchFromOriPrice(oriPrices []float64) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`ori_price` IN (?)", oriPrices).Find(&results).Error

	return
}

// GetFromValidDay 通过valid_day获取内容 有效天数
func (obj *_SrPackageCardMgr) GetFromValidDay(validDay int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`valid_day` = ?", validDay).Find(&results).Error

	return
}

// GetBatchFromValidDay 批量查找 有效天数
func (obj *_SrPackageCardMgr) GetBatchFromValidDay(validDays []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`valid_day` IN (?)", validDays).Find(&results).Error

	return
}

// GetFromActiveLimit 通过active_limit获取内容 激活期限天数
func (obj *_SrPackageCardMgr) GetFromActiveLimit(activeLimit int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`active_limit` = ?", activeLimit).Find(&results).Error

	return
}

// GetBatchFromActiveLimit 批量查找 激活期限天数
func (obj *_SrPackageCardMgr) GetBatchFromActiveLimit(activeLimits []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`active_limit` IN (?)", activeLimits).Find(&results).Error

	return
}

// GetFromBuyLimit 通过buy_limit获取内容 每人限购次数
func (obj *_SrPackageCardMgr) GetFromBuyLimit(buyLimit int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`buy_limit` = ?", buyLimit).Find(&results).Error

	return
}

// GetBatchFromBuyLimit 批量查找 每人限购次数
func (obj *_SrPackageCardMgr) GetBatchFromBuyLimit(buyLimits []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`buy_limit` IN (?)", buyLimits).Find(&results).Error

	return
}

// GetFromUseShop 通过use_shop获取内容 适用店铺
func (obj *_SrPackageCardMgr) GetFromUseShop(useShop string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`use_shop` = ?", useShop).Find(&results).Error

	return
}

// GetBatchFromUseShop 批量查找 适用店铺
func (obj *_SrPackageCardMgr) GetBatchFromUseShop(useShops []string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`use_shop` IN (?)", useShops).Find(&results).Error

	return
}

// GetFromUseSeat 通过use_seat获取内容 适用座位类型
func (obj *_SrPackageCardMgr) GetFromUseSeat(useSeat string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`use_seat` = ?", useSeat).Find(&results).Error

	return
}

// GetBatchFromUseSeat 批量查找 适用座位类型
func (obj *_SrPackageCardMgr) GetBatchFromUseSeat(useSeats []string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`use_seat` IN (?)", useSeats).Find(&results).Error

	return
}

// GetFromUseTime 通过use_time获取内容 使用时段
func (obj *_SrPackageCardMgr) GetFromUseTime(useTime string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`use_time` = ?", useTime).Find(&results).Error

	return
}

// GetBatchFromUseTime 批量查找 使用时段
func (obj *_SrPackageCardMgr) GetBatchFromUseTime(useTimes []string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`use_time` IN (?)", useTimes).Find(&results).Error

	return
}

// GetFromBookingDay 通过booking_day获取内容 总可用天数
func (obj *_SrPackageCardMgr) GetFromBookingDay(bookingDay int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_day` = ?", bookingDay).Find(&results).Error

	return
}

// GetBatchFromBookingDay 批量查找 总可用天数
func (obj *_SrPackageCardMgr) GetBatchFromBookingDay(bookingDays []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_day` IN (?)", bookingDays).Find(&results).Error

	return
}

// GetFromBookingMinute 通过booking_minute获取内容 总可用分钟
func (obj *_SrPackageCardMgr) GetFromBookingMinute(bookingMinute int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_minute` = ?", bookingMinute).Find(&results).Error

	return
}

// GetBatchFromBookingMinute 批量查找 总可用分钟
func (obj *_SrPackageCardMgr) GetBatchFromBookingMinute(bookingMinutes []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_minute` IN (?)", bookingMinutes).Find(&results).Error

	return
}

// GetFromBookingMinutePer 通过booking_minute_per获取内容 每天可用分钟
func (obj *_SrPackageCardMgr) GetFromBookingMinutePer(bookingMinutePer int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_minute_per` = ?", bookingMinutePer).Find(&results).Error

	return
}

// GetBatchFromBookingMinutePer 批量查找 每天可用分钟
func (obj *_SrPackageCardMgr) GetBatchFromBookingMinutePer(bookingMinutePers []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_minute_per` IN (?)", bookingMinutePers).Find(&results).Error

	return
}

// GetFromBookingMinuteMax 通过booking_minute_max获取内容 封顶可用分钟
func (obj *_SrPackageCardMgr) GetFromBookingMinuteMax(bookingMinuteMax int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_minute_max` = ?", bookingMinuteMax).Find(&results).Error

	return
}

// GetBatchFromBookingMinuteMax 批量查找 封顶可用分钟
func (obj *_SrPackageCardMgr) GetBatchFromBookingMinuteMax(bookingMinuteMaxs []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_minute_max` IN (?)", bookingMinuteMaxs).Find(&results).Error

	return
}

// GetFromBookingMinuteMin 通过booking_minute_min获取内容 最低可用分钟
func (obj *_SrPackageCardMgr) GetFromBookingMinuteMin(bookingMinuteMin int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_minute_min` = ?", bookingMinuteMin).Find(&results).Error

	return
}

// GetBatchFromBookingMinuteMin 批量查找 最低可用分钟
func (obj *_SrPackageCardMgr) GetBatchFromBookingMinuteMin(bookingMinuteMins []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`booking_minute_min` IN (?)", bookingMinuteMins).Find(&results).Error

	return
}

// GetFromCanBalance 通过can_balance获取内容 可否用余额购买
func (obj *_SrPackageCardMgr) GetFromCanBalance(canBalance *int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`can_balance` = ?", canBalance).Find(&results).Error

	return
}

// GetBatchFromCanBalance 批量查找 可否用余额购买
func (obj *_SrPackageCardMgr) GetBatchFromCanBalance(canBalances []*int8) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`can_balance` IN (?)", canBalances).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrPackageCardMgr) GetFromSort(sort int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrPackageCardMgr) GetBatchFromSort(sorts []int) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrPackageCardMgr) GetFromRemark(remark *string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrPackageCardMgr) GetBatchFromRemark(remarks []*string) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrPackageCardMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrPackageCardMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrPackageCardMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrPackageCardMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrPackageCardMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrPackageCardMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrPackageCardMgr) GetFromCreateBy(createBy *uint) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrPackageCardMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrPackageCardMgr) GetFromUpdateBy(updateBy *uint) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrPackageCardMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrPackageCardMgr) FetchByPrimaryKey(id int64) (result SrPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCard{}).Where("`id` = ?", id).First(&result).Error

	return
}
