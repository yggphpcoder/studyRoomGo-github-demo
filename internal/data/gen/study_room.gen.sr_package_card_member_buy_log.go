package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrPackageCardMemberBuyLogMgr struct {
	*_BaseMgr
}

// SrPackageCardMemberBuyLogMgr open func
func SrPackageCardMemberBuyLogMgr(db *gorm.DB) *_SrPackageCardMemberBuyLogMgr {
	if db == nil {
		panic(fmt.Errorf("SrPackageCardMemberBuyLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrPackageCardMemberBuyLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_package_card_member_buy_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrPackageCardMemberBuyLogMgr) GetTableName() string {
	return "sr_package_card_member_buy_log"
}

// Reset 重置gorm会话
func (obj *_SrPackageCardMemberBuyLogMgr) Reset() *_SrPackageCardMemberBuyLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrPackageCardMemberBuyLogMgr) Get() (result SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrPackageCardMemberBuyLogMgr) Gets() (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrPackageCardMemberBuyLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 套餐-会员购买记录id
func (obj *_SrPackageCardMemberBuyLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_SrPackageCardMemberBuyLogMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrPackageCardMemberBuyLogMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithOrderNo order_no获取 订单编号
func (obj *_SrPackageCardMemberBuyLogMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithMemberID member_id获取 会员id
func (obj *_SrPackageCardMemberBuyLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithPackageCardID package_card_id获取 套餐id
func (obj *_SrPackageCardMemberBuyLogMgr) WithPackageCardID(packageCardID int64) Option {
	return optionFunc(func(o *options) { o.query["package_card_id"] = packageCardID })
}

// WithMpID mp_id获取 会员-套餐id
func (obj *_SrPackageCardMemberBuyLogMgr) WithMpID(mpID int64) Option {
	return optionFunc(func(o *options) { o.query["mp_id"] = mpID })
}

// WithActualAmount actual_amount获取 实收金额
func (obj *_SrPackageCardMemberBuyLogMgr) WithActualAmount(actualAmount float64) Option {
	return optionFunc(func(o *options) { o.query["actual_amount"] = actualAmount })
}

// WithPrice price获取 售价
func (obj *_SrPackageCardMemberBuyLogMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithOriPrice ori_price获取 原价
func (obj *_SrPackageCardMemberBuyLogMgr) WithOriPrice(oriPrice float64) Option {
	return optionFunc(func(o *options) { o.query["ori_price"] = oriPrice })
}

// WithName name获取 套餐名
func (obj *_SrPackageCardMemberBuyLogMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTypePackage type_package获取 套餐类型
func (obj *_SrPackageCardMemberBuyLogMgr) WithTypePackage(typePackage int8) Option {
	return optionFunc(func(o *options) { o.query["type_package"] = typePackage })
}

// WithTypeCard type_card获取 卡类型
func (obj *_SrPackageCardMemberBuyLogMgr) WithTypeCard(typeCard int8) Option {
	return optionFunc(func(o *options) { o.query["type_card"] = typeCard })
}

// WithUseSeat use_seat获取 适用座位类型
func (obj *_SrPackageCardMemberBuyLogMgr) WithUseSeat(useSeat string) Option {
	return optionFunc(func(o *options) { o.query["use_seat"] = useSeat })
}

// WithUseShop use_shop获取 适用店铺
func (obj *_SrPackageCardMemberBuyLogMgr) WithUseShop(useShop string) Option {
	return optionFunc(func(o *options) { o.query["use_shop"] = useShop })
}

// WithTypeSale type_sale获取 销售类型
func (obj *_SrPackageCardMemberBuyLogMgr) WithTypeSale(typeSale int8) Option {
	return optionFunc(func(o *options) { o.query["type_sale"] = typeSale })
}

// WithValidDay valid_day获取 有效天数
func (obj *_SrPackageCardMemberBuyLogMgr) WithValidDay(validDay int) Option {
	return optionFunc(func(o *options) { o.query["valid_day"] = validDay })
}

// WithActiveLimit active_limit获取 激活期限天数
func (obj *_SrPackageCardMemberBuyLogMgr) WithActiveLimit(activeLimit int) Option {
	return optionFunc(func(o *options) { o.query["active_limit"] = activeLimit })
}

// WithBuyLimit buy_limit获取 每人限购次数
func (obj *_SrPackageCardMemberBuyLogMgr) WithBuyLimit(buyLimit int) Option {
	return optionFunc(func(o *options) { o.query["buy_limit"] = buyLimit })
}

// WithUseTime use_time获取 使用时段
func (obj *_SrPackageCardMemberBuyLogMgr) WithUseTime(useTime string) Option {
	return optionFunc(func(o *options) { o.query["use_time"] = useTime })
}

// WithBookingDay booking_day获取 总可用天数
func (obj *_SrPackageCardMemberBuyLogMgr) WithBookingDay(bookingDay int) Option {
	return optionFunc(func(o *options) { o.query["booking_day"] = bookingDay })
}

// WithBookingMinute booking_minute获取 总可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) WithBookingMinute(bookingMinute int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute"] = bookingMinute })
}

// WithBookingMinutePer booking_minute_per获取 每天可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) WithBookingMinutePer(bookingMinutePer int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute_per"] = bookingMinutePer })
}

// WithBookingMinuteMax booking_minute_max获取 封顶可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) WithBookingMinuteMax(bookingMinuteMax int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute_max"] = bookingMinuteMax })
}

// WithBookingMinuteMin booking_minute_min获取 最低可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) WithBookingMinuteMin(bookingMinuteMin int) Option {
	return optionFunc(func(o *options) { o.query["booking_minute_min"] = bookingMinuteMin })
}

// WithCanBalance can_balance获取 可否用余额购买
func (obj *_SrPackageCardMemberBuyLogMgr) WithCanBalance(canBalance *int8) Option {
	return optionFunc(func(o *options) { o.query["can_balance"] = canBalance })
}

// WithRemark remark获取 备注
func (obj *_SrPackageCardMemberBuyLogMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrPackageCardMemberBuyLogMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrPackageCardMemberBuyLogMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrPackageCardMemberBuyLogMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrPackageCardMemberBuyLogMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrPackageCardMemberBuyLogMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrPackageCardMemberBuyLogMgr) GetByOption(opts ...Option) (result SrPackageCardMemberBuyLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrPackageCardMemberBuyLogMgr) GetByOptions(opts ...Option) (results []*SrPackageCardMemberBuyLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 套餐-会员购买记录id
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromID(id int64) (result SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 套餐-会员购买记录id
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromID(ids []int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromMerchantID(merchantID int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromShopID(shopID int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单编号
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromOrderNo(orderNo string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`order_no` = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量查找 订单编号
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromOrderNo(orderNos []string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromMemberID(memberID int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromPackageCardID 通过package_card_id获取内容 套餐id
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromPackageCardID(packageCardID int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`package_card_id` = ?", packageCardID).Find(&results).Error

	return
}

// GetBatchFromPackageCardID 批量查找 套餐id
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromPackageCardID(packageCardIDs []int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`package_card_id` IN (?)", packageCardIDs).Find(&results).Error

	return
}

// GetFromMpID 通过mp_id获取内容 会员-套餐id
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromMpID(mpID int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`mp_id` = ?", mpID).Find(&results).Error

	return
}

// GetBatchFromMpID 批量查找 会员-套餐id
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromMpID(mpIDs []int64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`mp_id` IN (?)", mpIDs).Find(&results).Error

	return
}

// GetFromActualAmount 通过actual_amount获取内容 实收金额
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromActualAmount(actualAmount float64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`actual_amount` = ?", actualAmount).Find(&results).Error

	return
}

// GetBatchFromActualAmount 批量查找 实收金额
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromActualAmount(actualAmounts []float64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`actual_amount` IN (?)", actualAmounts).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 售价
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromPrice(price float64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 售价
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromPrice(prices []float64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromOriPrice 通过ori_price获取内容 原价
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromOriPrice(oriPrice float64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`ori_price` = ?", oriPrice).Find(&results).Error

	return
}

// GetBatchFromOriPrice 批量查找 原价
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromOriPrice(oriPrices []float64) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`ori_price` IN (?)", oriPrices).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 套餐名
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromName(name string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 套餐名
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromName(names []string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTypePackage 通过type_package获取内容 套餐类型
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromTypePackage(typePackage int8) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`type_package` = ?", typePackage).Find(&results).Error

	return
}

// GetBatchFromTypePackage 批量查找 套餐类型
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromTypePackage(typePackages []int8) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`type_package` IN (?)", typePackages).Find(&results).Error

	return
}

// GetFromTypeCard 通过type_card获取内容 卡类型
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromTypeCard(typeCard int8) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`type_card` = ?", typeCard).Find(&results).Error

	return
}

// GetBatchFromTypeCard 批量查找 卡类型
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromTypeCard(typeCards []int8) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`type_card` IN (?)", typeCards).Find(&results).Error

	return
}

// GetFromUseSeat 通过use_seat获取内容 适用座位类型
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromUseSeat(useSeat string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`use_seat` = ?", useSeat).Find(&results).Error

	return
}

// GetBatchFromUseSeat 批量查找 适用座位类型
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromUseSeat(useSeats []string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`use_seat` IN (?)", useSeats).Find(&results).Error

	return
}

// GetFromUseShop 通过use_shop获取内容 适用店铺
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromUseShop(useShop string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`use_shop` = ?", useShop).Find(&results).Error

	return
}

// GetBatchFromUseShop 批量查找 适用店铺
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromUseShop(useShops []string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`use_shop` IN (?)", useShops).Find(&results).Error

	return
}

// GetFromTypeSale 通过type_sale获取内容 销售类型
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromTypeSale(typeSale int8) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`type_sale` = ?", typeSale).Find(&results).Error

	return
}

// GetBatchFromTypeSale 批量查找 销售类型
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromTypeSale(typeSales []int8) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`type_sale` IN (?)", typeSales).Find(&results).Error

	return
}

// GetFromValidDay 通过valid_day获取内容 有效天数
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromValidDay(validDay int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`valid_day` = ?", validDay).Find(&results).Error

	return
}

// GetBatchFromValidDay 批量查找 有效天数
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromValidDay(validDays []int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`valid_day` IN (?)", validDays).Find(&results).Error

	return
}

// GetFromActiveLimit 通过active_limit获取内容 激活期限天数
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromActiveLimit(activeLimit int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`active_limit` = ?", activeLimit).Find(&results).Error

	return
}

// GetBatchFromActiveLimit 批量查找 激活期限天数
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromActiveLimit(activeLimits []int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`active_limit` IN (?)", activeLimits).Find(&results).Error

	return
}

// GetFromBuyLimit 通过buy_limit获取内容 每人限购次数
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromBuyLimit(buyLimit int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`buy_limit` = ?", buyLimit).Find(&results).Error

	return
}

// GetBatchFromBuyLimit 批量查找 每人限购次数
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromBuyLimit(buyLimits []int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`buy_limit` IN (?)", buyLimits).Find(&results).Error

	return
}

// GetFromUseTime 通过use_time获取内容 使用时段
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromUseTime(useTime string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`use_time` = ?", useTime).Find(&results).Error

	return
}

// GetBatchFromUseTime 批量查找 使用时段
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromUseTime(useTimes []string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`use_time` IN (?)", useTimes).Find(&results).Error

	return
}

// GetFromBookingDay 通过booking_day获取内容 总可用天数
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromBookingDay(bookingDay int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_day` = ?", bookingDay).Find(&results).Error

	return
}

// GetBatchFromBookingDay 批量查找 总可用天数
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromBookingDay(bookingDays []int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_day` IN (?)", bookingDays).Find(&results).Error

	return
}

// GetFromBookingMinute 通过booking_minute获取内容 总可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromBookingMinute(bookingMinute int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_minute` = ?", bookingMinute).Find(&results).Error

	return
}

// GetBatchFromBookingMinute 批量查找 总可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromBookingMinute(bookingMinutes []int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_minute` IN (?)", bookingMinutes).Find(&results).Error

	return
}

// GetFromBookingMinutePer 通过booking_minute_per获取内容 每天可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromBookingMinutePer(bookingMinutePer int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_minute_per` = ?", bookingMinutePer).Find(&results).Error

	return
}

// GetBatchFromBookingMinutePer 批量查找 每天可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromBookingMinutePer(bookingMinutePers []int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_minute_per` IN (?)", bookingMinutePers).Find(&results).Error

	return
}

// GetFromBookingMinuteMax 通过booking_minute_max获取内容 封顶可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromBookingMinuteMax(bookingMinuteMax int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_minute_max` = ?", bookingMinuteMax).Find(&results).Error

	return
}

// GetBatchFromBookingMinuteMax 批量查找 封顶可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromBookingMinuteMax(bookingMinuteMaxs []int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_minute_max` IN (?)", bookingMinuteMaxs).Find(&results).Error

	return
}

// GetFromBookingMinuteMin 通过booking_minute_min获取内容 最低可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromBookingMinuteMin(bookingMinuteMin int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_minute_min` = ?", bookingMinuteMin).Find(&results).Error

	return
}

// GetBatchFromBookingMinuteMin 批量查找 最低可用分钟
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromBookingMinuteMin(bookingMinuteMins []int) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`booking_minute_min` IN (?)", bookingMinuteMins).Find(&results).Error

	return
}

// GetFromCanBalance 通过can_balance获取内容 可否用余额购买
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromCanBalance(canBalance *int8) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`can_balance` = ?", canBalance).Find(&results).Error

	return
}

// GetBatchFromCanBalance 批量查找 可否用余额购买
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromCanBalance(canBalances []*int8) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`can_balance` IN (?)", canBalances).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromRemark(remark *string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromRemark(remarks []*string) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromCreateBy(createBy *uint) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrPackageCardMemberBuyLogMgr) GetFromUpdateBy(updateBy *uint) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrPackageCardMemberBuyLogMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrPackageCardMemberBuyLogMgr) FetchByPrimaryKey(id int64) (result SrPackageCardMemberBuyLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrPackageCardMemberBuyLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
