package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SrSettingTextMgr struct {
	*_BaseMgr
}

// SrSettingTextMgr open func
func SrSettingTextMgr(db *gorm.DB) *_SrSettingTextMgr {
	if db == nil {
		panic(fmt.Errorf("SrSettingTextMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrSettingTextMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_setting_text"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrSettingTextMgr) GetTableName() string {
	return "sr_setting_text"
}

// Reset 重置gorm会话
func (obj *_SrSettingTextMgr) Reset() *_SrSettingTextMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrSettingTextMgr) Get() (result SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SrSettingTextMgr) Gets() (results []*SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrSettingTextMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键编码
func (obj *_SrSettingTextMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSettingID setting_id获取 配置id
func (obj *_SrSettingTextMgr) WithSettingID(settingID int64) Option {
	return optionFunc(func(o *options) { o.query["setting_id"] = settingID })
}

// WithRichText rich_text获取 富文本
func (obj *_SrSettingTextMgr) WithRichText(richText string) Option {
	return optionFunc(func(o *options) { o.query["rich_text"] = richText })
}

// GetByOption 功能选项模式获取
func (obj *_SrSettingTextMgr) GetByOption(opts ...Option) (result SrSettingText, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrSettingTextMgr) GetByOptions(opts ...Option) (results []*SrSettingText, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键编码
func (obj *_SrSettingTextMgr) GetFromID(id int64) (result SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键编码
func (obj *_SrSettingTextMgr) GetBatchFromID(ids []int64) (results []*SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSettingID 通过setting_id获取内容 配置id
func (obj *_SrSettingTextMgr) GetFromSettingID(settingID int64) (results []*SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where("`setting_id` = ?", settingID).Find(&results).Error

	return
}

// GetBatchFromSettingID 批量查找 配置id
func (obj *_SrSettingTextMgr) GetBatchFromSettingID(settingIDs []int64) (results []*SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where("`setting_id` IN (?)", settingIDs).Find(&results).Error

	return
}

// GetFromRichText 通过rich_text获取内容 富文本
func (obj *_SrSettingTextMgr) GetFromRichText(richText string) (results []*SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where("`rich_text` = ?", richText).Find(&results).Error

	return
}

// GetBatchFromRichText 批量查找 富文本
func (obj *_SrSettingTextMgr) GetBatchFromRichText(richTexts []string) (results []*SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where("`rich_text` IN (?)", richTexts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrSettingTextMgr) FetchByPrimaryKey(id int64) (result SrSettingText, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrSettingText{}).Where("`id` = ?", id).First(&result).Error

	return
}
