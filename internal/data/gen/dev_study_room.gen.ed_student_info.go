package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdStudentInfoMgr struct {
	*_BaseMgr
}

// EdStudentInfoMgr open func
func EdStudentInfoMgr(db *gorm.DB) *_EdStudentInfoMgr {
	if db == nil {
		panic(fmt.Errorf("EdStudentInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdStudentInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_student_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdStudentInfoMgr) GetTableName() string {
	return "ed_student_info"
}

// Reset 重置gorm会话
func (obj *_EdStudentInfoMgr) Reset() *_EdStudentInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdStudentInfoMgr) Get() (result EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdStudentInfoMgr) Gets() (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdStudentInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdStudentInfoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdStudentInfoMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdStudentInfoMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithPhoto photo获取 照片
func (obj *_EdStudentInfoMgr) WithPhoto(photo *string) Option {
	return optionFunc(func(o *options) { o.query["photo"] = photo })
}

// WithName name获取 姓名
func (obj *_EdStudentInfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithGender gender获取 性别
func (obj *_EdStudentInfoMgr) WithGender(gender int8) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// WithPhone phone获取 手机号
func (obj *_EdStudentInfoMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithCreatedAt created_at获取
func (obj *_EdStudentInfoMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdStudentInfoMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdStudentInfoMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdStudentInfoMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdStudentInfoMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_EdStudentInfoMgr) GetByOption(opts ...Option) (result EdStudentInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdStudentInfoMgr) GetByOptions(opts ...Option) (results []*EdStudentInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdStudentInfoMgr) GetFromID(id int64) (result EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdStudentInfoMgr) GetBatchFromID(ids []int64) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdStudentInfoMgr) GetFromMerchantID(merchantID int64) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdStudentInfoMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdStudentInfoMgr) GetFromShopID(shopID int64) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdStudentInfoMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromPhoto 通过photo获取内容 照片
func (obj *_EdStudentInfoMgr) GetFromPhoto(photo *string) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`photo` = ?", photo).Find(&results).Error

	return
}

// GetBatchFromPhoto 批量查找 照片
func (obj *_EdStudentInfoMgr) GetBatchFromPhoto(photos []*string) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`photo` IN (?)", photos).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_EdStudentInfoMgr) GetFromName(name string) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_EdStudentInfoMgr) GetBatchFromName(names []string) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromGender 通过gender获取内容 性别
func (obj *_EdStudentInfoMgr) GetFromGender(gender int8) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`gender` = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量查找 性别
func (obj *_EdStudentInfoMgr) GetBatchFromGender(genders []int8) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`gender` IN (?)", genders).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机号
func (obj *_EdStudentInfoMgr) GetFromPhone(phone string) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 手机号
func (obj *_EdStudentInfoMgr) GetBatchFromPhone(phones []string) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdStudentInfoMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdStudentInfoMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdStudentInfoMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdStudentInfoMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdStudentInfoMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdStudentInfoMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdStudentInfoMgr) GetFromCreateBy(createBy *uint) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdStudentInfoMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdStudentInfoMgr) GetFromUpdateBy(updateBy *uint) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdStudentInfoMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdStudentInfoMgr) FetchByPrimaryKey(id int64) (result EdStudentInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}
