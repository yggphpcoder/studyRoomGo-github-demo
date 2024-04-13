package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EdTeacherInfoMgr struct {
	*_BaseMgr
}

// EdTeacherInfoMgr open func
func EdTeacherInfoMgr(db *gorm.DB) *_EdTeacherInfoMgr {
	if db == nil {
		panic(fmt.Errorf("EdTeacherInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdTeacherInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_teacher_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdTeacherInfoMgr) GetTableName() string {
	return "ed_teacher_info"
}

// Reset 重置gorm会话
func (obj *_EdTeacherInfoMgr) Reset() *_EdTeacherInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdTeacherInfoMgr) Get() (result EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EdTeacherInfoMgr) Gets() (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdTeacherInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_EdTeacherInfoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMerchantID merchant_id获取 商家id
func (obj *_EdTeacherInfoMgr) WithMerchantID(merchantID int64) Option {
	return optionFunc(func(o *options) { o.query["merchant_id"] = merchantID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EdTeacherInfoMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithPhoto photo获取 照片
func (obj *_EdTeacherInfoMgr) WithPhoto(photo *string) Option {
	return optionFunc(func(o *options) { o.query["photo"] = photo })
}

// WithName name获取 姓名
func (obj *_EdTeacherInfoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithGender gender获取 性别
func (obj *_EdTeacherInfoMgr) WithGender(gender int8) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// WithPhone phone获取 手机号
func (obj *_EdTeacherInfoMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithCreatedAt created_at获取
func (obj *_EdTeacherInfoMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_EdTeacherInfoMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_EdTeacherInfoMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_EdTeacherInfoMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_EdTeacherInfoMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithStatus status获取 状态
func (obj *_EdTeacherInfoMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_EdTeacherInfoMgr) GetByOption(opts ...Option) (result EdTeacherInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdTeacherInfoMgr) GetByOptions(opts ...Option) (results []*EdTeacherInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_EdTeacherInfoMgr) GetFromID(id int64) (result EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_EdTeacherInfoMgr) GetBatchFromID(ids []int64) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMerchantID 通过merchant_id获取内容 商家id
func (obj *_EdTeacherInfoMgr) GetFromMerchantID(merchantID int64) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`merchant_id` = ?", merchantID).Find(&results).Error

	return
}

// GetBatchFromMerchantID 批量查找 商家id
func (obj *_EdTeacherInfoMgr) GetBatchFromMerchantID(merchantIDs []int64) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`merchant_id` IN (?)", merchantIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EdTeacherInfoMgr) GetFromShopID(shopID int64) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EdTeacherInfoMgr) GetBatchFromShopID(shopIDs []int64) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromPhoto 通过photo获取内容 照片
func (obj *_EdTeacherInfoMgr) GetFromPhoto(photo *string) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`photo` = ?", photo).Find(&results).Error

	return
}

// GetBatchFromPhoto 批量查找 照片
func (obj *_EdTeacherInfoMgr) GetBatchFromPhoto(photos []*string) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`photo` IN (?)", photos).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_EdTeacherInfoMgr) GetFromName(name string) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_EdTeacherInfoMgr) GetBatchFromName(names []string) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromGender 通过gender获取内容 性别
func (obj *_EdTeacherInfoMgr) GetFromGender(gender int8) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`gender` = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量查找 性别
func (obj *_EdTeacherInfoMgr) GetBatchFromGender(genders []int8) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`gender` IN (?)", genders).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机号
func (obj *_EdTeacherInfoMgr) GetFromPhone(phone string) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 手机号
func (obj *_EdTeacherInfoMgr) GetBatchFromPhone(phones []string) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_EdTeacherInfoMgr) GetFromCreatedAt(createdAt *time.Time) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_EdTeacherInfoMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_EdTeacherInfoMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_EdTeacherInfoMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_EdTeacherInfoMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_EdTeacherInfoMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_EdTeacherInfoMgr) GetFromCreateBy(createBy *uint) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_EdTeacherInfoMgr) GetBatchFromCreateBy(createBys []*uint) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_EdTeacherInfoMgr) GetFromUpdateBy(updateBy *uint) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_EdTeacherInfoMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EdTeacherInfoMgr) GetFromStatus(status uint8) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EdTeacherInfoMgr) GetBatchFromStatus(statuss []uint8) (results []*EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdTeacherInfoMgr) FetchByPrimaryKey(id int64) (result EdTeacherInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdTeacherInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}
