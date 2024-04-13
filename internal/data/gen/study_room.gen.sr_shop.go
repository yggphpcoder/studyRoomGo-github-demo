package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrShopMgr struct {
	*_BaseMgr
}

// SrShopMgr open func
func SrShopMgr(db *gorm.DB) *_SrShopMgr {
	if db == nil {
		panic(fmt.Errorf("SrShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrShopMgr) GetTableName() string {
	return "sr_shop"
}

// Reset 重置gorm会话
func (obj *_SrShopMgr) Reset() *_SrShopMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrShopMgr) Get() (result SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrShopMgr) Gets() (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrShopMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrShop{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrShopMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_SrShopMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithName name获取 店铺名
func (obj *_SrShopMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTelephone telephone获取 联系电话
func (obj *_SrShopMgr) WithTelephone(telephone string) Option {
	return optionFunc(func(o *options) { o.query["telephone"] = telephone })
}

// WithWechat wechat获取 微信号
func (obj *_SrShopMgr) WithWechat(wechat string) Option {
	return optionFunc(func(o *options) { o.query["wechat"] = wechat })
}

// WithEmail email获取 邮箱
func (obj *_SrShopMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithProvince province获取 省份
func (obj *_SrShopMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 城市
func (obj *_SrShopMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithDistrict district获取 区
func (obj *_SrShopMgr) WithDistrict(district string) Option {
	return optionFunc(func(o *options) { o.query["district"] = district })
}

// WithAddress address获取 店铺地址
func (obj *_SrShopMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithLongitude longitude获取 经度
func (obj *_SrShopMgr) WithLongitude(longitude string) Option {
	return optionFunc(func(o *options) { o.query["longitude"] = longitude })
}

// WithLatitude latitude获取 纬度
func (obj *_SrShopMgr) WithLatitude(latitude string) Option {
	return optionFunc(func(o *options) { o.query["latitude"] = latitude })
}

// WithLeaderID leader_id获取 负责人id
func (obj *_SrShopMgr) WithLeaderID(leaderID int64) Option {
	return optionFunc(func(o *options) { o.query["leader_id"] = leaderID })
}

// WithTypeShop type_shop获取 店铺类型
func (obj *_SrShopMgr) WithTypeShop(typeShop int8) Option {
	return optionFunc(func(o *options) { o.query["type_shop"] = typeShop })
}

// WithBusinessHours business_hours获取 营业时间
func (obj *_SrShopMgr) WithBusinessHours(businessHours string) Option {
	return optionFunc(func(o *options) { o.query["business_hours"] = businessHours })
}

// WithSort sort获取 排序
func (obj *_SrShopMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrShopMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStatus status获取 状态
func (obj *_SrShopMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取
func (obj *_SrShopMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrShopMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrShopMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrShopMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrShopMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrShopMgr) GetByOption(opts ...Option) (result SrShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrShopMgr) GetByOptions(opts ...Option) (results []*SrShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrShopMgr) GetFromID(id int64) (result SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrShopMgr) GetBatchFromID(ids []int64) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_SrShopMgr) GetFromMerchantID(merchantID int64) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_SrShopMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 店铺名
func (obj *_SrShopMgr) GetFromName(name string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 店铺名
func (obj *_SrShopMgr) GetBatchFromName(names []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTelephone 通过telephone获取内容 联系电话
func (obj *_SrShopMgr) GetFromTelephone(telephone string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`telephone` = ?", telephone).Find(&results).Error

	return
}

// GetBatchFromTelephone 批量查找 联系电话
func (obj *_SrShopMgr) GetBatchFromTelephone(telephones []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`telephone` IN (?)", telephones).Find(&results).Error

	return
}

// GetFromWechat 通过wechat获取内容 微信号
func (obj *_SrShopMgr) GetFromWechat(wechat string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`wechat` = ?", wechat).Find(&results).Error

	return
}

// GetBatchFromWechat 批量查找 微信号
func (obj *_SrShopMgr) GetBatchFromWechat(wechats []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`wechat` IN (?)", wechats).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_SrShopMgr) GetFromEmail(email string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 邮箱
func (obj *_SrShopMgr) GetBatchFromEmail(emails []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 省份
func (obj *_SrShopMgr) GetFromProvince(province string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 省份
func (obj *_SrShopMgr) GetBatchFromProvince(provinces []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 城市
func (obj *_SrShopMgr) GetFromCity(city string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 城市
func (obj *_SrShopMgr) GetBatchFromCity(citys []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromDistrict 通过district获取内容 区
func (obj *_SrShopMgr) GetFromDistrict(district string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`district` = ?", district).Find(&results).Error

	return
}

// GetBatchFromDistrict 批量查找 区
func (obj *_SrShopMgr) GetBatchFromDistrict(districts []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`district` IN (?)", districts).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 店铺地址
func (obj *_SrShopMgr) GetFromAddress(address string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 店铺地址
func (obj *_SrShopMgr) GetBatchFromAddress(addresss []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromLongitude 通过longitude获取内容 经度
func (obj *_SrShopMgr) GetFromLongitude(longitude string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`longitude` = ?", longitude).Find(&results).Error

	return
}

// GetBatchFromLongitude 批量查找 经度
func (obj *_SrShopMgr) GetBatchFromLongitude(longitudes []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`longitude` IN (?)", longitudes).Find(&results).Error

	return
}

// GetFromLatitude 通过latitude获取内容 纬度
func (obj *_SrShopMgr) GetFromLatitude(latitude string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`latitude` = ?", latitude).Find(&results).Error

	return
}

// GetBatchFromLatitude 批量查找 纬度
func (obj *_SrShopMgr) GetBatchFromLatitude(latitudes []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`latitude` IN (?)", latitudes).Find(&results).Error

	return
}

// GetFromLeaderID 通过leader_id获取内容 负责人id
func (obj *_SrShopMgr) GetFromLeaderID(leaderID int64) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`leader_id` = ?", leaderID).Find(&results).Error

	return
}

// GetBatchFromLeaderID 批量查找 负责人id
func (obj *_SrShopMgr) GetBatchFromLeaderID(leaderIDs []int64) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`leader_id` IN (?)", leaderIDs).Find(&results).Error

	return
}

// GetFromTypeShop 通过type_shop获取内容 店铺类型
func (obj *_SrShopMgr) GetFromTypeShop(typeShop int8) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`type_shop` = ?", typeShop).Find(&results).Error

	return
}

// GetBatchFromTypeShop 批量查找 店铺类型
func (obj *_SrShopMgr) GetBatchFromTypeShop(typeShops []int8) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`type_shop` IN (?)", typeShops).Find(&results).Error

	return
}

// GetFromBusinessHours 通过business_hours获取内容 营业时间
func (obj *_SrShopMgr) GetFromBusinessHours(businessHours string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`business_hours` = ?", businessHours).Find(&results).Error

	return
}

// GetBatchFromBusinessHours 批量查找 营业时间
func (obj *_SrShopMgr) GetBatchFromBusinessHours(businessHourss []string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`business_hours` IN (?)", businessHourss).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrShopMgr) GetFromSort(sort int) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrShopMgr) GetBatchFromSort(sorts []int) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrShopMgr) GetFromRemark(remark *string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrShopMgr) GetBatchFromRemark(remarks []*string) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_SrShopMgr) GetFromStatus(status int8) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_SrShopMgr) GetBatchFromStatus(statuss []int8) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrShopMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrShopMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrShopMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrShopMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrShopMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrShopMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrShopMgr) GetFromCreateBy(createBy *uint) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrShopMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrShopMgr) GetFromUpdateBy(updateBy *uint) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrShopMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrShopMgr) FetchByPrimaryKey(id int64) (result SrShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrShop{}).Where("`id` = ?", id).First(&result).Error

	return
}
