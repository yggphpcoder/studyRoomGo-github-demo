package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CoMerchantWxMiniprogramMgr struct {
	*_BaseMgr
}

// CoMerchantWxMiniprogramMgr open func
func CoMerchantWxMiniprogramMgr(db *gorm.DB) *_CoMerchantWxMiniprogramMgr {
	if db == nil {
		panic(fmt.Errorf("CoMerchantWxMiniprogramMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CoMerchantWxMiniprogramMgr{_BaseMgr: &_BaseMgr{DB: db.Table("co_merchant_wx_miniprogram"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CoMerchantWxMiniprogramMgr) GetTableName() string {
	return "co_merchant_wx_miniprogram"
}

// Reset 重置gorm会话
func (obj *_CoMerchantWxMiniprogramMgr) Reset() *_CoMerchantWxMiniprogramMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CoMerchantWxMiniprogramMgr) Get() (result CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CoMerchantWxMiniprogramMgr) Gets() (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CoMerchantWxMiniprogramMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_CoMerchantWxMiniprogramMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_CoMerchantWxMiniprogramMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithSignCode sign_code获取 商家标识码
func (obj *_CoMerchantWxMiniprogramMgr) WithSignCode(signCode string) Option {
	return optionFunc(func(o *options) { o.query["sign_code"] = signCode })
}

// WithAppID app_id获取 微信小程序app_id
func (obj *_CoMerchantWxMiniprogramMgr) WithAppID(appID string) Option {
	return optionFunc(func(o *options) { o.query["app_id"] = appID })
}

// WithAppSecret app_secret获取 微信小程序app_secret
func (obj *_CoMerchantWxMiniprogramMgr) WithAppSecret(appSecret string) Option {
	return optionFunc(func(o *options) { o.query["app_secret"] = appSecret })
}

// WithVersion version获取 合作版本
func (obj *_CoMerchantWxMiniprogramMgr) WithVersion(version int8) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithStatus status获取 状态
func (obj *_CoMerchantWxMiniprogramMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSort sort获取 排序
func (obj *_CoMerchantWxMiniprogramMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_CoMerchantWxMiniprogramMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_CoMerchantWxMiniprogramMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_CoMerchantWxMiniprogramMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_CoMerchantWxMiniprogramMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_CoMerchantWxMiniprogramMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_CoMerchantWxMiniprogramMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_CoMerchantWxMiniprogramMgr) GetByOption(opts ...Option) (result CoMerchantWxMiniprogram, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CoMerchantWxMiniprogramMgr) GetByOptions(opts ...Option) (results []*CoMerchantWxMiniprogram, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_CoMerchantWxMiniprogramMgr) GetFromID(id int64) (result CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromID(ids []int64) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_CoMerchantWxMiniprogramMgr) GetFromMerchantID(merchantID int64) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromSignCode 通过sign_code获取内容 商家标识码
func (obj *_CoMerchantWxMiniprogramMgr) GetFromSignCode(signCode string) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`sign_code` = ?", signCode).Find(&results).Error

	return
}

// GetBatchFromSignCode 批量查找 商家标识码
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromSignCode(signCodes []string) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`sign_code` IN (?)", signCodes).Find(&results).Error

	return
}

// GetFromAppID 通过app_id获取内容 微信小程序app_id
func (obj *_CoMerchantWxMiniprogramMgr) GetFromAppID(appID string) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`app_id` = ?", appID).Find(&results).Error

	return
}

// GetBatchFromAppID 批量查找 微信小程序app_id
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromAppID(appIDs []string) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`app_id` IN (?)", appIDs).Find(&results).Error

	return
}

// GetFromAppSecret 通过app_secret获取内容 微信小程序app_secret
func (obj *_CoMerchantWxMiniprogramMgr) GetFromAppSecret(appSecret string) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`app_secret` = ?", appSecret).Find(&results).Error

	return
}

// GetBatchFromAppSecret 批量查找 微信小程序app_secret
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromAppSecret(appSecrets []string) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`app_secret` IN (?)", appSecrets).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 合作版本
func (obj *_CoMerchantWxMiniprogramMgr) GetFromVersion(version int8) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 合作版本
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromVersion(versions []int8) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_CoMerchantWxMiniprogramMgr) GetFromStatus(status int8) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromStatus(statuss []int8) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_CoMerchantWxMiniprogramMgr) GetFromSort(sort int) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromSort(sorts []int) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_CoMerchantWxMiniprogramMgr) GetFromRemark(remark *string) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromRemark(remarks []*string) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_CoMerchantWxMiniprogramMgr) GetFromCreatedAt(createdAt *time.Time) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_CoMerchantWxMiniprogramMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_CoMerchantWxMiniprogramMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_CoMerchantWxMiniprogramMgr) GetFromCreateBy(createBy *uint) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromCreateBy(createBys []*uint) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_CoMerchantWxMiniprogramMgr) GetFromUpdateBy(updateBy *uint) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_CoMerchantWxMiniprogramMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CoMerchantWxMiniprogramMgr) FetchByPrimaryKey(id int64) (result CoMerchantWxMiniprogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CoMerchantWxMiniprogram{}).Where("`id` = ?", id).First(&result).Error

	return
}
