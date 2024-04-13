package gen

import "time"

// EdTeacherInfo 教务-教师档案
type EdTeacherInfo struct {
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
	Status     uint8      `gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:'状态'"` // 状态
}

// TableName get sql table name.获取数据库表名
func (m *EdTeacherInfo) TableName() string {
	return "ed_teacher_info"
}

// EdTeacherInfoColumns get sql column name.获取数据库列名
var EdTeacherInfoColumns = struct {
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
	Status     string
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
	Status:     "status",
}
