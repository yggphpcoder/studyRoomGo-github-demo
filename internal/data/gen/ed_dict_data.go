package gen

import "time"

// EdDictData 教务-字典数据表
type EdDictData struct {
	DictCode   int64      `gorm:"autoIncrement:true;primaryKey;column:dict_code;type:bigint;not null"`
	MerchantID int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"` // 商家id
	ShopID     int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`     // 店铺id
	DictSort   *int64     `gorm:"column:dict_sort;type:bigint"`
	DictLabel  *string    `gorm:"column:dict_label;type:varchar(128)"`
	DictValue  *string    `gorm:"column:dict_value;type:varchar(255)"`
	DictType   *string    `gorm:"column:dict_type;type:varchar(64)"`
	CSSClass   *string    `gorm:"column:css_class;type:varchar(128)"`
	ListClass  *string    `gorm:"column:list_class;type:varchar(128)"`
	IsDefault  *string    `gorm:"column:is_default;type:varchar(8)"`
	Status     *int8      `gorm:"column:status;type:tinyint"`
	Default    *string    `gorm:"column:default;type:varchar(8)"`
	Remark     *string    `gorm:"column:remark;type:varchar(255)"`
	CreateBy   *int64     `gorm:"index:idx_sys_dict_data_create_by;column:create_by;type:bigint;comment:'创建者'"`         // 创建者
	UpdateBy   *int64     `gorm:"index:idx_sys_dict_data_update_by;column:update_by;type:bigint;comment:'更新者'"`         // 更新者
	CreatedAt  *time.Time `gorm:"column:created_at;type:datetime(3);comment:'创建时间'"`                                    // 创建时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:datetime(3);comment:'最后更新时间'"`                                  // 最后更新时间
	DeletedAt  *time.Time `gorm:"index:idx_sys_dict_data_deleted_at;column:deleted_at;type:datetime(3);comment:'删除时间'"` // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *EdDictData) TableName() string {
	return "ed_dict_data"
}

// EdDictDataColumns get sql column name.获取数据库列名
var EdDictDataColumns = struct {
	DictCode   string
	MerchantID string
	ShopID     string
	DictSort   string
	DictLabel  string
	DictValue  string
	DictType   string
	CSSClass   string
	ListClass  string
	IsDefault  string
	Status     string
	Default    string
	Remark     string
	CreateBy   string
	UpdateBy   string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}{
	DictCode:   "dict_code",
	MerchantID: "merchant_id",
	ShopID:     "shop_id",
	DictSort:   "dict_sort",
	DictLabel:  "dict_label",
	DictValue:  "dict_value",
	DictType:   "dict_type",
	CSSClass:   "css_class",
	ListClass:  "list_class",
	IsDefault:  "is_default",
	Status:     "status",
	Default:    "default",
	Remark:     "remark",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}
