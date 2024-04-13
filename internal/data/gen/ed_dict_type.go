package gen

import "time"

// EdDictType 教务-字典类型表
type EdDictType struct {
	DictID    int64      `gorm:"autoIncrement:true;primaryKey;column:dict_id;type:bigint;not null"`
	DictName  *string    `gorm:"column:dict_name;type:varchar(128)"`
	DictType  *string    `gorm:"column:dict_type;type:varchar(128)"`
	Status    *int8      `gorm:"column:status;type:tinyint"`
	Remark    *string    `gorm:"column:remark;type:varchar(255)"`
	CreateBy  *int64     `gorm:"index:idx_sys_dict_type_create_by;column:create_by;type:bigint;comment:'创建者'"`         // 创建者
	UpdateBy  *int64     `gorm:"index:idx_sys_dict_type_update_by;column:update_by;type:bigint;comment:'更新者'"`         // 更新者
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime(3);comment:'创建时间'"`                                    // 创建时间
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime(3);comment:'最后更新时间'"`                                  // 最后更新时间
	DeletedAt *time.Time `gorm:"index:idx_sys_dict_type_deleted_at;column:deleted_at;type:datetime(3);comment:'删除时间'"` // 删除时间
	Class     string     `gorm:"column:class;type:varchar(255);not null;comment:'分类'"`                                 // 分类
}

// TableName get sql table name.获取数据库表名
func (m *EdDictType) TableName() string {
	return "ed_dict_type"
}

// EdDictTypeColumns get sql column name.获取数据库列名
var EdDictTypeColumns = struct {
	DictID    string
	DictName  string
	DictType  string
	Status    string
	Remark    string
	CreateBy  string
	UpdateBy  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Class     string
}{
	DictID:    "dict_id",
	DictName:  "dict_name",
	DictType:  "dict_type",
	Status:    "status",
	Remark:    "remark",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Class:     "class",
}
