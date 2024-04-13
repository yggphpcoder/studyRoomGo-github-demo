package gen

import "time"

// EdStudentInfoKey 教务-学生档案-属性
type EdStudentInfoKey struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	ShopID     int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	Key        string     `gorm:"column:key;type:varchar(64);not null;default:'';comment:'字段'"`                // 字段
	Value      string     `gorm:"column:value;type:varchar(64);not null;default:'';comment:'字段名'"`             // 字段名
	HTMLType   string     `gorm:"column:html_type;type:varchar(64);not null;default:'';comment:'字段类型'"`        // 字段类型
	DictType   string     `gorm:"column:dict_type;type:varchar(64);not null;default:'';comment:'字典'"`          // 字典
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *EdStudentInfoKey) TableName() string {
	return "ed_student_info_key"
}

// EdStudentInfoKeyColumns get sql column name.获取数据库列名
var EdStudentInfoKeyColumns = struct {
	ID         string
	MerchantID string
	ShopID     string
	Key        string
	Value      string
	HTMLType   string
	DictType   string
	CreateBy   string
	UpdateBy   string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}{
	ID:         "id",
	MerchantID: "merchant_id",
	ShopID:     "shop_id",
	Key:        "key",
	Value:      "value",
	HTMLType:   "html_type",
	DictType:   "dict_type",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}
