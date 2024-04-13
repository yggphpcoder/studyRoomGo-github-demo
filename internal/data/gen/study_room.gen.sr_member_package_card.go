package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrMemberPackageCardMgr struct {
	*_BaseMgr
}

// SrMemberPackageCardMgr open func
func SrMemberPackageCardMgr(db *gorm.DB) *_SrMemberPackageCardMgr {
	if db == nil {
		panic(fmt.Errorf("SrMemberPackageCardMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrMemberPackageCardMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_member_package_card"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrMemberPackageCardMgr) GetTableName() string {
	return "sr_member_package_card"
}

// Reset 重置gorm会话
func (obj *_SrMemberPackageCardMgr) Reset() *_SrMemberPackageCardMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrMemberPackageCardMgr) Get() (result SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrMemberPackageCardMgr) Gets() (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrMemberPackageCardMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 会员-套餐id
func (obj *_SrMemberPackageCardMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_SrMemberPackageCardMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrMemberPackageCardMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithMemberID member_id获取 会员id
func (obj *_SrMemberPackageCardMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithActualAmount actual_amount获取 实收价格
func (obj *_SrMemberPackageCardMgr) WithActualAmount(actualAmount float64) Option {
	return optionFunc(func(o *options) { o.query["actual_amount"] = actualAmount })
}

// WithRemainMinute remain_minute获取 剩余分钟
func (obj *_SrMemberPackageCardMgr) WithRemainMinute(remainMinute int) Option {
	return optionFunc(func(o *options) { o.query["remain_minute"] = remainMinute })
}

// WithRemainDay remain_day获取 剩余天数
func (obj *_SrMemberPackageCardMgr) WithRemainDay(remainDay int) Option {
	return optionFunc(func(o *options) { o.query["remain_day"] = remainDay })
}

// WithMinutePer minute_per获取 每天可用分钟
func (obj *_SrMemberPackageCardMgr) WithMinutePer(minutePer int) Option {
	return optionFunc(func(o *options) { o.query["minute_per"] = minutePer })
}

// WithMinuteMax minute_max获取 封顶可用分钟
func (obj *_SrMemberPackageCardMgr) WithMinuteMax(minuteMax int) Option {
	return optionFunc(func(o *options) { o.query["minute_max"] = minuteMax })
}

// WithMinuteMin minute_min获取 最低可用分钟
func (obj *_SrMemberPackageCardMgr) WithMinuteMin(minuteMin int) Option {
	return optionFunc(func(o *options) { o.query["minute_min"] = minuteMin })
}

// WithStatus status获取 状态
func (obj *_SrMemberPackageCardMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTypeChannel type_channel获取 渠道类型
func (obj *_SrMemberPackageCardMgr) WithTypeChannel(typeChannel int8) Option {
	return optionFunc(func(o *options) { o.query["type_channel"] = typeChannel })
}

// WithActiveAt active_at获取 激活时间
func (obj *_SrMemberPackageCardMgr) WithActiveAt(activeAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["active_at"] = activeAt })
}

// WithInvalidAt invalid_at获取 失效时间
func (obj *_SrMemberPackageCardMgr) WithInvalidAt(invalidAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["invalid_at"] = invalidAt })
}

// WithPackageID package_id获取 套餐id
func (obj *_SrMemberPackageCardMgr) WithPackageID(packageID *int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithPackageName package_name获取 套餐名
func (obj *_SrMemberPackageCardMgr) WithPackageName(packageName string) Option {
	return optionFunc(func(o *options) { o.query["package_name"] = packageName })
}

// WithTypePackage type_package获取 套餐类型
func (obj *_SrMemberPackageCardMgr) WithTypePackage(typePackage int8) Option {
	return optionFunc(func(o *options) { o.query["type_package"] = typePackage })
}

// WithTypeCard type_card获取 卡类型
func (obj *_SrMemberPackageCardMgr) WithTypeCard(typeCard int8) Option {
	return optionFunc(func(o *options) { o.query["type_card"] = typeCard })
}

// WithTypeSale type_sale获取 销售类型
func (obj *_SrMemberPackageCardMgr) WithTypeSale(typeSale int8) Option {
	return optionFunc(func(o *options) { o.query["type_sale"] = typeSale })
}

// WithPrice price获取 售价
func (obj *_SrMemberPackageCardMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithOriPrice ori_price获取 原价
func (obj *_SrMemberPackageCardMgr) WithOriPrice(oriPrice float64) Option {
	return optionFunc(func(o *options) { o.query["ori_price"] = oriPrice })
}

// WithUseShop use_shop获取 适用店铺
func (obj *_SrMemberPackageCardMgr) WithUseShop(useShop string) Option {
	return optionFunc(func(o *options) { o.query["use_shop"] = useShop })
}

// WithUseSeat use_seat获取 适用座位
func (obj *_SrMemberPackageCardMgr) WithUseSeat(useSeat string) Option {
	return optionFunc(func(o *options) { o.query["use_seat"] = useSeat })
}

// WithUseTime use_time获取 使用时段
func (obj *_SrMemberPackageCardMgr) WithUseTime(useTime string) Option {
	return optionFunc(func(o *options) { o.query["use_time"] = useTime })
}

// WithUseBalance use_balance获取 使用余额支付
func (obj *_SrMemberPackageCardMgr) WithUseBalance(useBalance *int8) Option {
	return optionFunc(func(o *options) { o.query["use_balance"] = useBalance })
}

// WithValidDay valid_day获取 有效天数
func (obj *_SrMemberPackageCardMgr) WithValidDay(validDay int) Option {
	return optionFunc(func(o *options) { o.query["valid_day"] = validDay })
}

// WithActiveLimit active_limit获取 激活期限天数
func (obj *_SrMemberPackageCardMgr) WithActiveLimit(activeLimit int) Option {
	return optionFunc(func(o *options) { o.query["active_limit"] = activeLimit })
}

// WithSort sort获取 排序
func (obj *_SrMemberPackageCardMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrMemberPackageCardMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrMemberPackageCardMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrMemberPackageCardMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrMemberPackageCardMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrMemberPackageCardMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrMemberPackageCardMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrMemberPackageCardMgr) GetByOption(opts ...Option) (result SrMemberPackageCard, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrMemberPackageCardMgr) GetByOptions(opts ...Option) (results []*SrMemberPackageCard, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 会员-套餐id
func (obj *_SrMemberPackageCardMgr) GetFromID(id int64) (result SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 会员-套餐id
func (obj *_SrMemberPackageCardMgr) GetBatchFromID(ids []int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_SrMemberPackageCardMgr) GetFromMerchantID(merchantID int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_SrMemberPackageCardMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrMemberPackageCardMgr) GetFromShopID(shopID int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrMemberPackageCardMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_SrMemberPackageCardMgr) GetFromMemberID(memberID int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_SrMemberPackageCardMgr) GetBatchFromMemberID(memberIDs []int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromActualAmount 通过actual_amount获取内容 实收价格
func (obj *_SrMemberPackageCardMgr) GetFromActualAmount(actualAmount float64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`actual_amount` = ?", actualAmount).Find(&results).Error

	return
}

// GetBatchFromActualAmount 批量查找 实收价格
func (obj *_SrMemberPackageCardMgr) GetBatchFromActualAmount(actualAmounts []float64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`actual_amount` IN (?)", actualAmounts).Find(&results).Error

	return
}

// GetFromRemainMinute 通过remain_minute获取内容 剩余分钟
func (obj *_SrMemberPackageCardMgr) GetFromRemainMinute(remainMinute int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`remain_minute` = ?", remainMinute).Find(&results).Error

	return
}

// GetBatchFromRemainMinute 批量查找 剩余分钟
func (obj *_SrMemberPackageCardMgr) GetBatchFromRemainMinute(remainMinutes []int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`remain_minute` IN (?)", remainMinutes).Find(&results).Error

	return
}

// GetFromRemainDay 通过remain_day获取内容 剩余天数
func (obj *_SrMemberPackageCardMgr) GetFromRemainDay(remainDay int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`remain_day` = ?", remainDay).Find(&results).Error

	return
}

// GetBatchFromRemainDay 批量查找 剩余天数
func (obj *_SrMemberPackageCardMgr) GetBatchFromRemainDay(remainDays []int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`remain_day` IN (?)", remainDays).Find(&results).Error

	return
}

// GetFromMinutePer 通过minute_per获取内容 每天可用分钟
func (obj *_SrMemberPackageCardMgr) GetFromMinutePer(minutePer int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`minute_per` = ?", minutePer).Find(&results).Error

	return
}

// GetBatchFromMinutePer 批量查找 每天可用分钟
func (obj *_SrMemberPackageCardMgr) GetBatchFromMinutePer(minutePers []int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`minute_per` IN (?)", minutePers).Find(&results).Error

	return
}

// GetFromMinuteMax 通过minute_max获取内容 封顶可用分钟
func (obj *_SrMemberPackageCardMgr) GetFromMinuteMax(minuteMax int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`minute_max` = ?", minuteMax).Find(&results).Error

	return
}

// GetBatchFromMinuteMax 批量查找 封顶可用分钟
func (obj *_SrMemberPackageCardMgr) GetBatchFromMinuteMax(minuteMaxs []int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`minute_max` IN (?)", minuteMaxs).Find(&results).Error

	return
}

// GetFromMinuteMin 通过minute_min获取内容 最低可用分钟
func (obj *_SrMemberPackageCardMgr) GetFromMinuteMin(minuteMin int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`minute_min` = ?", minuteMin).Find(&results).Error

	return
}

// GetBatchFromMinuteMin 批量查找 最低可用分钟
func (obj *_SrMemberPackageCardMgr) GetBatchFromMinuteMin(minuteMins []int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`minute_min` IN (?)", minuteMins).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SrMemberPackageCardMgr) GetFromStatus(status int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SrMemberPackageCardMgr) GetBatchFromStatus(statuss []int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTypeChannel 通过type_channel获取内容 渠道类型
func (obj *_SrMemberPackageCardMgr) GetFromTypeChannel(typeChannel int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`type_channel` = ?", typeChannel).Find(&results).Error

	return
}

// GetBatchFromTypeChannel 批量查找 渠道类型
func (obj *_SrMemberPackageCardMgr) GetBatchFromTypeChannel(typeChannels []int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`type_channel` IN (?)", typeChannels).Find(&results).Error

	return
}

// GetFromActiveAt 通过active_at获取内容 激活时间
func (obj *_SrMemberPackageCardMgr) GetFromActiveAt(activeAt *time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`active_at` = ?", activeAt).Find(&results).Error

	return
}

// GetBatchFromActiveAt 批量查找 激活时间
func (obj *_SrMemberPackageCardMgr) GetBatchFromActiveAt(activeAts []*time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`active_at` IN (?)", activeAts).Find(&results).Error

	return
}

// GetFromInvalidAt 通过invalid_at获取内容 失效时间
func (obj *_SrMemberPackageCardMgr) GetFromInvalidAt(invalidAt *time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`invalid_at` = ?", invalidAt).Find(&results).Error

	return
}

// GetBatchFromInvalidAt 批量查找 失效时间
func (obj *_SrMemberPackageCardMgr) GetBatchFromInvalidAt(invalidAts []*time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`invalid_at` IN (?)", invalidAts).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容 套餐id
func (obj *_SrMemberPackageCardMgr) GetFromPackageID(packageID *int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找 套餐id
func (obj *_SrMemberPackageCardMgr) GetBatchFromPackageID(packageIDs []*int64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromPackageName 通过package_name获取内容 套餐名
func (obj *_SrMemberPackageCardMgr) GetFromPackageName(packageName string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`package_name` = ?", packageName).Find(&results).Error

	return
}

// GetBatchFromPackageName 批量查找 套餐名
func (obj *_SrMemberPackageCardMgr) GetBatchFromPackageName(packageNames []string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`package_name` IN (?)", packageNames).Find(&results).Error

	return
}

// GetFromTypePackage 通过type_package获取内容 套餐类型
func (obj *_SrMemberPackageCardMgr) GetFromTypePackage(typePackage int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`type_package` = ?", typePackage).Find(&results).Error

	return
}

// GetBatchFromTypePackage 批量查找 套餐类型
func (obj *_SrMemberPackageCardMgr) GetBatchFromTypePackage(typePackages []int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`type_package` IN (?)", typePackages).Find(&results).Error

	return
}

// GetFromTypeCard 通过type_card获取内容 卡类型
func (obj *_SrMemberPackageCardMgr) GetFromTypeCard(typeCard int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`type_card` = ?", typeCard).Find(&results).Error

	return
}

// GetBatchFromTypeCard 批量查找 卡类型
func (obj *_SrMemberPackageCardMgr) GetBatchFromTypeCard(typeCards []int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`type_card` IN (?)", typeCards).Find(&results).Error

	return
}

// GetFromTypeSale 通过type_sale获取内容 销售类型
func (obj *_SrMemberPackageCardMgr) GetFromTypeSale(typeSale int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`type_sale` = ?", typeSale).Find(&results).Error

	return
}

// GetBatchFromTypeSale 批量查找 销售类型
func (obj *_SrMemberPackageCardMgr) GetBatchFromTypeSale(typeSales []int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`type_sale` IN (?)", typeSales).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 售价
func (obj *_SrMemberPackageCardMgr) GetFromPrice(price float64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 售价
func (obj *_SrMemberPackageCardMgr) GetBatchFromPrice(prices []float64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromOriPrice 通过ori_price获取内容 原价
func (obj *_SrMemberPackageCardMgr) GetFromOriPrice(oriPrice float64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`ori_price` = ?", oriPrice).Find(&results).Error

	return
}

// GetBatchFromOriPrice 批量查找 原价
func (obj *_SrMemberPackageCardMgr) GetBatchFromOriPrice(oriPrices []float64) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`ori_price` IN (?)", oriPrices).Find(&results).Error

	return
}

// GetFromUseShop 通过use_shop获取内容 适用店铺
func (obj *_SrMemberPackageCardMgr) GetFromUseShop(useShop string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`use_shop` = ?", useShop).Find(&results).Error

	return
}

// GetBatchFromUseShop 批量查找 适用店铺
func (obj *_SrMemberPackageCardMgr) GetBatchFromUseShop(useShops []string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`use_shop` IN (?)", useShops).Find(&results).Error

	return
}

// GetFromUseSeat 通过use_seat获取内容 适用座位
func (obj *_SrMemberPackageCardMgr) GetFromUseSeat(useSeat string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`use_seat` = ?", useSeat).Find(&results).Error

	return
}

// GetBatchFromUseSeat 批量查找 适用座位
func (obj *_SrMemberPackageCardMgr) GetBatchFromUseSeat(useSeats []string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`use_seat` IN (?)", useSeats).Find(&results).Error

	return
}

// GetFromUseTime 通过use_time获取内容 使用时段
func (obj *_SrMemberPackageCardMgr) GetFromUseTime(useTime string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`use_time` = ?", useTime).Find(&results).Error

	return
}

// GetBatchFromUseTime 批量查找 使用时段
func (obj *_SrMemberPackageCardMgr) GetBatchFromUseTime(useTimes []string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`use_time` IN (?)", useTimes).Find(&results).Error

	return
}

// GetFromUseBalance 通过use_balance获取内容 使用余额支付
func (obj *_SrMemberPackageCardMgr) GetFromUseBalance(useBalance *int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`use_balance` = ?", useBalance).Find(&results).Error

	return
}

// GetBatchFromUseBalance 批量查找 使用余额支付
func (obj *_SrMemberPackageCardMgr) GetBatchFromUseBalance(useBalances []*int8) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`use_balance` IN (?)", useBalances).Find(&results).Error

	return
}

// GetFromValidDay 通过valid_day获取内容 有效天数
func (obj *_SrMemberPackageCardMgr) GetFromValidDay(validDay int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`valid_day` = ?", validDay).Find(&results).Error

	return
}

// GetBatchFromValidDay 批量查找 有效天数
func (obj *_SrMemberPackageCardMgr) GetBatchFromValidDay(validDays []int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`valid_day` IN (?)", validDays).Find(&results).Error

	return
}

// GetFromActiveLimit 通过active_limit获取内容 激活期限天数
func (obj *_SrMemberPackageCardMgr) GetFromActiveLimit(activeLimit int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`active_limit` = ?", activeLimit).Find(&results).Error

	return
}

// GetBatchFromActiveLimit 批量查找 激活期限天数
func (obj *_SrMemberPackageCardMgr) GetBatchFromActiveLimit(activeLimits []int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`active_limit` IN (?)", activeLimits).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrMemberPackageCardMgr) GetFromSort(sort int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrMemberPackageCardMgr) GetBatchFromSort(sorts []int) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrMemberPackageCardMgr) GetFromRemark(remark *string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrMemberPackageCardMgr) GetBatchFromRemark(remarks []*string) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrMemberPackageCardMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrMemberPackageCardMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrMemberPackageCardMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrMemberPackageCardMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrMemberPackageCardMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrMemberPackageCardMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrMemberPackageCardMgr) GetFromCreateBy(createBy *uint) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrMemberPackageCardMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrMemberPackageCardMgr) GetFromUpdateBy(updateBy *uint) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrMemberPackageCardMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrMemberPackageCardMgr) FetchByPrimaryKey(id int64) (result SrMemberPackageCard, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrMemberPackageCard{}).Where("`id` = ?", id).First(&result).Error

	return
}
