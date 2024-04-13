package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _EdStudentInfoExtendMgr struct {
	*_BaseMgr
}

// EdStudentInfoExtendMgr open func
func EdStudentInfoExtendMgr(db *gorm.DB) *_EdStudentInfoExtendMgr {
	if db == nil {
		panic(fmt.Errorf("EdStudentInfoExtendMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EdStudentInfoExtendMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ed_student_info_extend"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EdStudentInfoExtendMgr) GetTableName() string {
	return "ed_student_info_extend"
}

// Reset 重置gorm会话
func (obj *_EdStudentInfoExtendMgr) Reset() *_EdStudentInfoExtendMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EdStudentInfoExtendMgr) Get() (result EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("ed_student_info").Where("id = ?", result.StudentID).Find(&result.EdStudentInfo).Error; err != nil { // 教务-学生档案
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_EdStudentInfoExtendMgr) Gets() (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EdStudentInfoExtendMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithStudentID student_id获取 学生id
func (obj *_EdStudentInfoExtendMgr) WithStudentID(studentID int64) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithField0 field_0获取 字段0
func (obj *_EdStudentInfoExtendMgr) WithField0(field0 *string) Option {
	return optionFunc(func(o *options) { o.query["field_0"] = field0 })
}

// WithField1 field_1获取 字段1
func (obj *_EdStudentInfoExtendMgr) WithField1(field1 *string) Option {
	return optionFunc(func(o *options) { o.query["field_1"] = field1 })
}

// WithField2 field_2获取 字段2
func (obj *_EdStudentInfoExtendMgr) WithField2(field2 *string) Option {
	return optionFunc(func(o *options) { o.query["field_2"] = field2 })
}

// WithField3 field_3获取 字段3
func (obj *_EdStudentInfoExtendMgr) WithField3(field3 *string) Option {
	return optionFunc(func(o *options) { o.query["field_3"] = field3 })
}

// WithField4 field_4获取 字段4
func (obj *_EdStudentInfoExtendMgr) WithField4(field4 *string) Option {
	return optionFunc(func(o *options) { o.query["field_4"] = field4 })
}

// WithField5 field_5获取 字段5
func (obj *_EdStudentInfoExtendMgr) WithField5(field5 *string) Option {
	return optionFunc(func(o *options) { o.query["field_5"] = field5 })
}

// WithField6 field_6获取 字段6
func (obj *_EdStudentInfoExtendMgr) WithField6(field6 *string) Option {
	return optionFunc(func(o *options) { o.query["field_6"] = field6 })
}

// WithField7 field_7获取 字段7
func (obj *_EdStudentInfoExtendMgr) WithField7(field7 *string) Option {
	return optionFunc(func(o *options) { o.query["field_7"] = field7 })
}

// WithField8 field_8获取 字段8
func (obj *_EdStudentInfoExtendMgr) WithField8(field8 *string) Option {
	return optionFunc(func(o *options) { o.query["field_8"] = field8 })
}

// WithField9 field_9获取 字段9
func (obj *_EdStudentInfoExtendMgr) WithField9(field9 *string) Option {
	return optionFunc(func(o *options) { o.query["field_9"] = field9 })
}

// WithField10 field_10获取 字段10
func (obj *_EdStudentInfoExtendMgr) WithField10(field10 *string) Option {
	return optionFunc(func(o *options) { o.query["field_10"] = field10 })
}

// GetByOption 功能选项模式获取
func (obj *_EdStudentInfoExtendMgr) GetByOption(opts ...Option) (result EdStudentInfoExtend, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("ed_student_info").Where("id = ?", result.StudentID).Find(&result.EdStudentInfo).Error; err != nil { // 教务-学生档案
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EdStudentInfoExtendMgr) GetByOptions(opts ...Option) (results []*EdStudentInfoExtend, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromStudentID 通过student_id获取内容 学生id
func (obj *_EdStudentInfoExtendMgr) GetFromStudentID(studentID int64) (result EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`student_id` = ?", studentID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("ed_student_info").Where("id = ?", result.StudentID).Find(&result.EdStudentInfo).Error; err != nil { // 教务-学生档案
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromStudentID 批量查找 学生id
func (obj *_EdStudentInfoExtendMgr) GetBatchFromStudentID(studentIDs []int64) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField0 通过field_0获取内容 字段0
func (obj *_EdStudentInfoExtendMgr) GetFromField0(field0 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_0` = ?", field0).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField0 批量查找 字段0
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField0(field0s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_0` IN (?)", field0s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField1 通过field_1获取内容 字段1
func (obj *_EdStudentInfoExtendMgr) GetFromField1(field1 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_1` = ?", field1).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField1 批量查找 字段1
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField1(field1s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_1` IN (?)", field1s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField2 通过field_2获取内容 字段2
func (obj *_EdStudentInfoExtendMgr) GetFromField2(field2 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_2` = ?", field2).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField2 批量查找 字段2
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField2(field2s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_2` IN (?)", field2s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField3 通过field_3获取内容 字段3
func (obj *_EdStudentInfoExtendMgr) GetFromField3(field3 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_3` = ?", field3).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField3 批量查找 字段3
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField3(field3s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_3` IN (?)", field3s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField4 通过field_4获取内容 字段4
func (obj *_EdStudentInfoExtendMgr) GetFromField4(field4 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_4` = ?", field4).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField4 批量查找 字段4
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField4(field4s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_4` IN (?)", field4s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField5 通过field_5获取内容 字段5
func (obj *_EdStudentInfoExtendMgr) GetFromField5(field5 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_5` = ?", field5).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField5 批量查找 字段5
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField5(field5s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_5` IN (?)", field5s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField6 通过field_6获取内容 字段6
func (obj *_EdStudentInfoExtendMgr) GetFromField6(field6 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_6` = ?", field6).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField6 批量查找 字段6
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField6(field6s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_6` IN (?)", field6s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField7 通过field_7获取内容 字段7
func (obj *_EdStudentInfoExtendMgr) GetFromField7(field7 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_7` = ?", field7).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField7 批量查找 字段7
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField7(field7s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_7` IN (?)", field7s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField8 通过field_8获取内容 字段8
func (obj *_EdStudentInfoExtendMgr) GetFromField8(field8 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_8` = ?", field8).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField8 批量查找 字段8
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField8(field8s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_8` IN (?)", field8s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField9 通过field_9获取内容 字段9
func (obj *_EdStudentInfoExtendMgr) GetFromField9(field9 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_9` = ?", field9).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField9 批量查找 字段9
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField9(field9s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_9` IN (?)", field9s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromField10 通过field_10获取内容 字段10
func (obj *_EdStudentInfoExtendMgr) GetFromField10(field10 *string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_10` = ?", field10).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromField10 批量查找 字段10
func (obj *_EdStudentInfoExtendMgr) GetBatchFromField10(field10s []*string) (results []*EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`field_10` IN (?)", field10s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("ed_student_info").Where("id = ?", results[i].StudentID).Find(&results[i].EdStudentInfo).Error; err != nil { // 教务-学生档案
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EdStudentInfoExtendMgr) FetchByPrimaryKey(studentID int64) (result EdStudentInfoExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EdStudentInfoExtend{}).Where("`student_id` = ?", studentID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("ed_student_info").Where("id = ?", result.StudentID).Find(&result.EdStudentInfo).Error; err != nil { // 教务-学生档案
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
