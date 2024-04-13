package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdCoursesSaleMgr struct {
	*_BaseMgr
}

// EdCoursesSaleMgr open func
func EdCoursesSaleMgr(db *gorm.DB) *_EdCoursesSaleMgr {
	if db == nil {
		panic(fmt.Errorf("EdCoursesSaleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdCoursesSaleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_courses_sale"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdCoursesSaleMgr) GetTableName() string {
	return "ed_courses_sale"
}

// Reset 重置gorm会话
func (obj *_EdCoursesSaleMgr) Reset() *_EdCoursesSaleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdCoursesSaleMgr) Get() (result EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdCoursesSaleMgr) Gets() (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdCoursesSaleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdCoursesSaleMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCourseID course_id获取 课程id
func (obj *_EdCoursesSaleMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithUseShop use_shop获取 适用店铺
func (obj *_EdCoursesSaleMgr) WithUseShop(useShop string) Option {
	return optionFunc(func(o *options) { o.query["use_shop"] = useShop })
}

// WithStatus status获取 营销状态
func (obj *_EdCoursesSaleMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithType type获取 销售类型
func (obj *_EdCoursesSaleMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithRule rule获取 销售规则
func (obj *_EdCoursesSaleMgr) WithRule(rule string) Option {
	return optionFunc(func(o *options) { o.query["rule"] = rule })
}

// WithPrice price获取 售价
func (obj *_EdCoursesSaleMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithOriPrice ori_price获取 原价
func (obj *_EdCoursesSaleMgr) WithOriPrice(oriPrice float64) Option {
	return optionFunc(func(o *options) { o.query["ori_price"] = oriPrice })
}

// WithIsShow is_show获取 是否在前台显示
func (obj *_EdCoursesSaleMgr) WithIsShow(isShow int8) Option {
	return optionFunc(func(o *options) { o.query["is_show"] = isShow })
}

// WithBuyLimit buy_limit获取 购买限制
func (obj *_EdCoursesSaleMgr) WithBuyLimit(buyLimit int) Option {
	return optionFunc(func(o *options) { o.query["buy_limit"] = buyLimit })
}

// WithStartDate start_date获取 售卖开始日期
func (obj *_EdCoursesSaleMgr) WithStartDate(startDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_date"] = startDate })
}

// WithEndDate end_date获取 售卖结束日期
func (obj *_EdCoursesSaleMgr) WithEndDate(endDate *time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_date"] = endDate })
}

// WithValidDay valid_day获取 默认有效天数
func (obj *_EdCoursesSaleMgr) WithValidDay(validDay int) Option {
	return optionFunc(func(o *options) { o.query["valid_day"] = validDay })
}

// WithActiveLimit active_limit获取 默认激活期限天数
func (obj *_EdCoursesSaleMgr) WithActiveLimit(activeLimit int) Option {
	return optionFunc(func(o *options) { o.query["active_limit"] = activeLimit })
}

// WithSort sort获取 排序
func (obj *_EdCoursesSaleMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_EdCoursesSaleMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_EdCoursesSaleMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdCoursesSaleMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdCoursesSaleMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdCoursesSaleMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdCoursesSaleMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdCoursesSaleMgr) GetByOption(opts ...Option) (result EdCoursesSale, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdCoursesSaleMgr) GetByOptions(opts ...Option) (results []*EdCoursesSale, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdCoursesSaleMgr) GetFromID(id int64) (result EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdCoursesSaleMgr) GetBatchFromID(ids []int64) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 课程id
func (obj *_EdCoursesSaleMgr) GetFromCourseID(courseID int64) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`course_id` = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量查找 课程id
func (obj *_EdCoursesSaleMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`course_id` IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromUseShop 通过use_shop获取内容 适用店铺
func (obj *_EdCoursesSaleMgr) GetFromUseShop(useShop string) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`use_shop` = ?", useShop).Find(&results).Error

	return
}

// GetBatchFromUseShop 批量查找 适用店铺
func (obj *_EdCoursesSaleMgr) GetBatchFromUseShop(useShops []string) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`use_shop` IN (?)", useShops).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 营销状态
func (obj *_EdCoursesSaleMgr) GetFromStatus(status int8) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 营销状态
func (obj *_EdCoursesSaleMgr) GetBatchFromStatus(statuss []int8) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 销售类型
func (obj *_EdCoursesSaleMgr) GetFromType(_type int8) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 销售类型
func (obj *_EdCoursesSaleMgr) GetBatchFromType(_types []int8) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromRule 通过rule获取内容 销售规则
func (obj *_EdCoursesSaleMgr) GetFromRule(rule string) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`rule` = ?", rule).Find(&results).Error

	return
}

// GetBatchFromRule 批量查找 销售规则
func (obj *_EdCoursesSaleMgr) GetBatchFromRule(rules []string) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`rule` IN (?)", rules).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 售价
func (obj *_EdCoursesSaleMgr) GetFromPrice(price float64) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 售价
func (obj *_EdCoursesSaleMgr) GetBatchFromPrice(prices []float64) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromOriPrice 通过ori_price获取内容 原价
func (obj *_EdCoursesSaleMgr) GetFromOriPrice(oriPrice float64) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`ori_price` = ?", oriPrice).Find(&results).Error

	return
}

// GetBatchFromOriPrice 批量查找 原价
func (obj *_EdCoursesSaleMgr) GetBatchFromOriPrice(oriPrices []float64) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`ori_price` IN (?)", oriPrices).Find(&results).Error

	return
}

// GetFromIsShow 通过is_show获取内容 是否在前台显示
func (obj *_EdCoursesSaleMgr) GetFromIsShow(isShow int8) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`is_show` = ?", isShow).Find(&results).Error

	return
}

// GetBatchFromIsShow 批量查找 是否在前台显示
func (obj *_EdCoursesSaleMgr) GetBatchFromIsShow(isShows []int8) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`is_show` IN (?)", isShows).Find(&results).Error

	return
}

// GetFromBuyLimit 通过buy_limit获取内容 购买限制
func (obj *_EdCoursesSaleMgr) GetFromBuyLimit(buyLimit int) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`buy_limit` = ?", buyLimit).Find(&results).Error

	return
}

// GetBatchFromBuyLimit 批量查找 购买限制
func (obj *_EdCoursesSaleMgr) GetBatchFromBuyLimit(buyLimits []int) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`buy_limit` IN (?)", buyLimits).Find(&results).Error

	return
}

// GetFromStartDate 通过start_date获取内容 售卖开始日期
func (obj *_EdCoursesSaleMgr) GetFromStartDate(startDate *time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`start_date` = ?", startDate).Find(&results).Error

	return
}

// GetBatchFromStartDate 批量查找 售卖开始日期
func (obj *_EdCoursesSaleMgr) GetBatchFromStartDate(startDates []*time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`start_date` IN (?)", startDates).Find(&results).Error

	return
}

// GetFromEndDate 通过end_date获取内容 售卖结束日期
func (obj *_EdCoursesSaleMgr) GetFromEndDate(endDate *time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`end_date` = ?", endDate).Find(&results).Error

	return
}

// GetBatchFromEndDate 批量查找 售卖结束日期
func (obj *_EdCoursesSaleMgr) GetBatchFromEndDate(endDates []*time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`end_date` IN (?)", endDates).Find(&results).Error

	return
}

// GetFromValidDay 通过valid_day获取内容 默认有效天数
func (obj *_EdCoursesSaleMgr) GetFromValidDay(validDay int) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`valid_day` = ?", validDay).Find(&results).Error

	return
}

// GetBatchFromValidDay 批量查找 默认有效天数
func (obj *_EdCoursesSaleMgr) GetBatchFromValidDay(validDays []int) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`valid_day` IN (?)", validDays).Find(&results).Error

	return
}

// GetFromActiveLimit 通过active_limit获取内容 默认激活期限天数
func (obj *_EdCoursesSaleMgr) GetFromActiveLimit(activeLimit int) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`active_limit` = ?", activeLimit).Find(&results).Error

	return
}

// GetBatchFromActiveLimit 批量查找 默认激活期限天数
func (obj *_EdCoursesSaleMgr) GetBatchFromActiveLimit(activeLimits []int) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`active_limit` IN (?)", activeLimits).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_EdCoursesSaleMgr) GetFromSort(sort int) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_EdCoursesSaleMgr) GetBatchFromSort(sorts []int) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_EdCoursesSaleMgr) GetFromRemark(remark *string) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_EdCoursesSaleMgr) GetBatchFromRemark(remarks []*string) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdCoursesSaleMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdCoursesSaleMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdCoursesSaleMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdCoursesSaleMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdCoursesSaleMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdCoursesSaleMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdCoursesSaleMgr) GetFromCreateBy(createBy *uint) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdCoursesSaleMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdCoursesSaleMgr) GetFromUpdateBy(updateBy *uint) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdCoursesSaleMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdCoursesSaleMgr) FetchByPrimaryKey(id int64) (result EdCoursesSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdCoursesSale{}).Where("`id` = ?", id).First(&result).Error

	return
}
