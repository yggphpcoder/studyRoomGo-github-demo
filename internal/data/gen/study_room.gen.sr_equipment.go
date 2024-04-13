package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrEquipmentMgr struct {
	*_BaseMgr
}

// SrEquipmentMgr open func
func SrEquipmentMgr(db *gorm.DB) *_SrEquipmentMgr {
	if db == nil {
		panic(fmt.Errorf("SrEquipmentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrEquipmentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_equipment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrEquipmentMgr) GetTableName() string {
	return "sr_equipment"
}

// Reset 重置gorm会话
func (obj *_SrEquipmentMgr) Reset() *_SrEquipmentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrEquipmentMgr) Get() (result SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrEquipmentMgr) Gets() (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrEquipmentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrEquipmentMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取 父键
func (obj *_SrEquipmentMgr) WithPid(pid int64) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithName name获取 设备名
func (obj *_SrEquipmentMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTypeEquipment type_equipment获取 设备类型
func (obj *_SrEquipmentMgr) WithTypeEquipment(typeEquipment int8) Option {
	return optionFunc(func(o *options) { o.query["type_equipment"] = typeEquipment })
}

// WithTypeRelate type_relate获取 关联类型
func (obj *_SrEquipmentMgr) WithTypeRelate(typeRelate int8) Option {
	return optionFunc(func(o *options) { o.query["type_relate"] = typeRelate })
}

// WithRelateID relate_id获取 关联id
func (obj *_SrEquipmentMgr) WithRelateID(relateID int64) Option {
	return optionFunc(func(o *options) { o.query["relate_id"] = relateID })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrEquipmentMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithCode code获取 设备码
func (obj *_SrEquipmentMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithMac mac获取 mac地址
func (obj *_SrEquipmentMgr) WithMac(mac string) Option {
	return optionFunc(func(o *options) { o.query["mac"] = mac })
}

// WithPort port获取 端口数
func (obj *_SrEquipmentMgr) WithPort(port string) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithStatus status获取 可用状态
func (obj *_SrEquipmentMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithOperation operation获取 操作状态
func (obj *_SrEquipmentMgr) WithOperation(operation int8) Option {
	return optionFunc(func(o *options) { o.query["operation"] = operation })
}

// WithSort sort获取 排序
func (obj *_SrEquipmentMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrEquipmentMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreatedAt created_at获取
func (obj *_SrEquipmentMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrEquipmentMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrEquipmentMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrEquipmentMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrEquipmentMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrEquipmentMgr) GetByOption(opts ...Option) (result SrEquipment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrEquipmentMgr) GetByOptions(opts ...Option) (results []*SrEquipment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrEquipmentMgr) GetFromID(id int64) (result SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrEquipmentMgr) GetBatchFromID(ids []int64) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 父键
func (obj *_SrEquipmentMgr) GetFromPid(pid int64) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`pid` = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量查找 父键
func (obj *_SrEquipmentMgr) GetBatchFromPid(pids []int64) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`pid` IN (?)", pids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 设备名
func (obj *_SrEquipmentMgr) GetFromName(name string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 设备名
func (obj *_SrEquipmentMgr) GetBatchFromName(names []string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTypeEquipment 通过type_equipment获取内容 设备类型
func (obj *_SrEquipmentMgr) GetFromTypeEquipment(typeEquipment int8) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`type_equipment` = ?", typeEquipment).Find(&results).Error

	return
}

// GetBatchFromTypeEquipment 批量查找 设备类型
func (obj *_SrEquipmentMgr) GetBatchFromTypeEquipment(typeEquipments []int8) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`type_equipment` IN (?)", typeEquipments).Find(&results).Error

	return
}

// GetFromTypeRelate 通过type_relate获取内容 关联类型
func (obj *_SrEquipmentMgr) GetFromTypeRelate(typeRelate int8) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`type_relate` = ?", typeRelate).Find(&results).Error

	return
}

// GetBatchFromTypeRelate 批量查找 关联类型
func (obj *_SrEquipmentMgr) GetBatchFromTypeRelate(typeRelates []int8) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`type_relate` IN (?)", typeRelates).Find(&results).Error

	return
}

// GetFromRelateID 通过relate_id获取内容 关联id
func (obj *_SrEquipmentMgr) GetFromRelateID(relateID int64) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`relate_id` = ?", relateID).Find(&results).Error

	return
}

// GetBatchFromRelateID 批量查找 关联id
func (obj *_SrEquipmentMgr) GetBatchFromRelateID(relateIDs []int64) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`relate_id` IN (?)", relateIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrEquipmentMgr) GetFromShopID(shopID int64) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrEquipmentMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 设备码
func (obj *_SrEquipmentMgr) GetFromCode(code string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 设备码
func (obj *_SrEquipmentMgr) GetBatchFromCode(codes []string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromMac 通过mac获取内容 mac地址
func (obj *_SrEquipmentMgr) GetFromMac(mac string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`mac` = ?", mac).Find(&results).Error

	return
}

// GetBatchFromMac 批量查找 mac地址
func (obj *_SrEquipmentMgr) GetBatchFromMac(macs []string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`mac` IN (?)", macs).Find(&results).Error

	return
}

// GetFromPort 通过port获取内容 端口数
func (obj *_SrEquipmentMgr) GetFromPort(port string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`port` = ?", port).Find(&results).Error

	return
}

// GetBatchFromPort 批量查找 端口数
func (obj *_SrEquipmentMgr) GetBatchFromPort(ports []string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`port` IN (?)", ports).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 可用状态
func (obj *_SrEquipmentMgr) GetFromStatus(status int8) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 可用状态
func (obj *_SrEquipmentMgr) GetBatchFromStatus(statuss []int8) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromOperation 通过operation获取内容 操作状态
func (obj *_SrEquipmentMgr) GetFromOperation(operation int8) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`operation` = ?", operation).Find(&results).Error

	return
}

// GetBatchFromOperation 批量查找 操作状态
func (obj *_SrEquipmentMgr) GetBatchFromOperation(operations []int8) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`operation` IN (?)", operations).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrEquipmentMgr) GetFromSort(sort int) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrEquipmentMgr) GetBatchFromSort(sorts []int) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrEquipmentMgr) GetFromRemark(remark *string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrEquipmentMgr) GetBatchFromRemark(remarks []*string) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrEquipmentMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrEquipmentMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrEquipmentMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrEquipmentMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrEquipmentMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrEquipmentMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrEquipmentMgr) GetFromCreateBy(createBy *uint) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrEquipmentMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrEquipmentMgr) GetFromUpdateBy(updateBy *uint) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrEquipmentMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrEquipmentMgr) FetchByPrimaryKey(id int64) (result SrEquipment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrEquipment{}).Where("`id` = ?", id).First(&result).Error

	return
}
