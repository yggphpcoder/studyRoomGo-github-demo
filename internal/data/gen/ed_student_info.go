package gen

import "time"

// EdStudentInfo 教务-学生档案
type EdStudentInfo struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	ShopID     int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	Photo      *string    `gorm:"column:photo;type:varchar(255);comment:'照片'"`                                 // 照片
	Name       string     `gorm:"column:name;type:varchar(64);not null;default:'';comment:'姓名'"`               // 姓名
	Gender     int8       `gorm:"column:gender;type:tinyint;not null;default:0;comment:'性别'"`                  // 性别
	Phone      string     `gorm:"column:phone;type:varchar(11);not null;default:'';comment:'手机号'"`             // 手机号
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdStudentInfo) TableName() string {
	return "ed_student_info"
}

// EdStudentInfoColumns get sql column name.获取数据库列名
var EdStudentInfoColumns = struct {
	ID         string
	MerchantID string
	ShopID     string
	Photo      string
	Name       string
	Gender     string
	Phone      string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	CreateBy   string
	UpdateBy   string
}{
	ID:         "id",
	MerchantID: "merchant_id",
	ShopID:     "shop_id",
	Photo:      "photo",
	Name:       "name",
	Gender:     "gender",
	Phone:      "phone",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
}
