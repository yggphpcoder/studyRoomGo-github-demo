package gen

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SrSeatMgr struct {
	*_BaseMgr
}

// SrSeatMgr open func
func SrSeatMgr(db *gorm.DB) *_SrSeatMgr {
	if db == nil {
		panic(fmt.Errorf("SrSeatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrSeatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_seat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrSeatMgr) GetTableName() string {
	return "sr_seat"
}

// Reset 重置gorm会话
func (obj *_SrSeatMgr) Reset() *_SrSeatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrSeatMgr) Get() (result SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrSeatMgr) Gets() (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrSeatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrSeatMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取 父id
func (obj *_SrSeatMgr) WithPid(pid int64) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithName name获取 座位名
func (obj *_SrSeatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTypeSeat type_seat获取 座位类型
func (obj *_SrSeatMgr) WithTypeSeat(typeSeat int8) Option {
	return optionFunc(func(o *options) { o.query["type_seat"] = typeSeat })
}

// WithShopID shop_id获取 店铺id
func (obj *_SrSeatMgr) WithShopID(shopID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopAreaID shop_area_id获取 店铺分区id
func (obj *_SrSeatMgr) WithShopAreaID(shopAreaID int64) Option {
	return optionFunc(func(o *options) { o.query["shop_area_id"] = shopAreaID })
}

// WithStatus status获取 可用状态
func (obj *_SrSeatMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSort sort获取 排序
func (obj *_SrSeatMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithRemark remark获取 备注
func (obj *_SrSeatMgr) WithRemark(remark *string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithMapPosition map_position获取 座位图坐标
func (obj *_SrSeatMgr) WithMapPosition(mapPosition *string) Option {
	return optionFunc(func(o *options) { o.query["map_position"] = mapPosition })
}

// WithCreatedAt created_at获取
func (obj *_SrSeatMgr) WithCreatedAt(createdAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_SrSeatMgr) WithUpdatedAt(updatedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_SrSeatMgr) WithDeletedAt(deletedAt *time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithCreateBy create_by获取
func (obj *_SrSeatMgr) WithCreateBy(createBy *uint) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithUpdateBy update_by获取
func (obj *_SrSeatMgr) WithUpdateBy(updateBy *uint) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_SrSeatMgr) GetByOption(opts ...Option) (result SrSeat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrSeatMgr) GetByOptions(opts ...Option) (results []*SrSeat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrSeatMgr) GetFromID(id int64) (result SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrSeatMgr) GetBatchFromID(ids []int64) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 父id
func (obj *_SrSeatMgr) GetFromPid(pid int64) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`pid` = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量查找 父id
func (obj *_SrSeatMgr) GetBatchFromPid(pids []int64) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`pid` IN (?)", pids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 座位名
func (obj *_SrSeatMgr) GetFromName(name string) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 座位名
func (obj *_SrSeatMgr) GetBatchFromName(names []string) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTypeSeat 通过type_seat获取内容 座位类型
func (obj *_SrSeatMgr) GetFromTypeSeat(typeSeat int8) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`type_seat` = ?", typeSeat).Find(&results).Error

	return
}

// GetBatchFromTypeSeat 批量查找 座位类型
func (obj *_SrSeatMgr) GetBatchFromTypeSeat(typeSeats []int8) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`type_seat` IN (?)", typeSeats).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_SrSeatMgr) GetFromShopID(shopID int64) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_SrSeatMgr) GetBatchFromShopID(shopIDs []int64) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopAreaID 通过shop_area_id获取内容 店铺分区id
func (obj *_SrSeatMgr) GetFromShopAreaID(shopAreaID int64) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`shop_area_id` = ?", shopAreaID).Find(&results).Error

	return
}

// GetBatchFromShopAreaID 批量查找 店铺分区id
func (obj *_SrSeatMgr) GetBatchFromShopAreaID(shopAreaIDs []int64) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`shop_area_id` IN (?)", shopAreaIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 可用状态
func (obj *_SrSeatMgr) GetFromStatus(status int8) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 可用状态
func (obj *_SrSeatMgr) GetBatchFromStatus(statuss []int8) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_SrSeatMgr) GetFromSort(sort int) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_SrSeatMgr) GetBatchFromSort(sorts []int) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SrSeatMgr) GetFromRemark(remark *string) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SrSeatMgr) GetBatchFromRemark(remarks []*string) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromMapPosition 通过map_position获取内容 座位图坐标
func (obj *_SrSeatMgr) GetFromMapPosition(mapPosition *string) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`map_position` = ?", mapPosition).Find(&results).Error

	return
}

// GetBatchFromMapPosition 批量查找 座位图坐标
func (obj *_SrSeatMgr) GetBatchFromMapPosition(mapPositions []*string) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`map_position` IN (?)", mapPositions).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_SrSeatMgr) GetFromCreatedAt(createdAt *time.Time) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_SrSeatMgr) GetBatchFromCreatedAt(createdAts []*time.Time) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_SrSeatMgr) GetFromUpdatedAt(updatedAt *time.Time) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_SrSeatMgr) GetBatchFromUpdatedAt(updatedAts []*time.Time) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_SrSeatMgr) GetFromDeletedAt(deletedAt *time.Time) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_SrSeatMgr) GetBatchFromDeletedAt(deletedAts []*time.Time) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容
func (obj *_SrSeatMgr) GetFromCreateBy(createBy *uint) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找
func (obj *_SrSeatMgr) GetBatchFromCreateBy(createBys []*uint) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_SrSeatMgr) GetFromUpdateBy(updateBy *uint) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_SrSeatMgr) GetBatchFromUpdateBy(updateBys []*uint) (results []*SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrSeatMgr) FetchByPrimaryKey(id int64) (result SrSeat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSeat{}).Where("`id` = ?", id).First(&result).Error

	return
}
