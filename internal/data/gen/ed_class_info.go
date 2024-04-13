package gen

import (
	"time"
)

// EdClassInfo 教务-班级-排课表
type EdClassInfo struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	ShopID     int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	CourseID   int64      `gorm:"column:course_id;type:bigint;not null;default:0;comment:'课程id'"`              // 课程id
	TeacherID  int64      `gorm:"column:teacher_id;type:bigint;not null;default:0;comment:'教师id'"`             // 教师id
	StartDate  time.Time  `gorm:"column:start_date;type:date;not null;comment:'开始日期'"`                         // 开始日期
	EndDate    time.Time  `gorm:"column:end_date;type:date;not null;comment:'结束日期'"`                           // 结束日期
	UseTime    string     `gorm:"column:use_time;type:varchar(1000);not null;default:'';comment:'上课时段'"`       // 上课时段
	Status     int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`                  // 状态
	Name       string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'名称'"`              // 名称
	Summary    string     `gorm:"column:summary;type:varchar(1000);not null;default:'';comment:'简介'"`          // 简介
	Remark     *string    `gorm:"column:remark;type:varchar(1000);comment:'备注'"`                               // 备注
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdClassInfo) TableName() string {
	return "ed_class_info"
}

// EdClassInfoColumns get sql column name.获取数据库列名
var EdClassInfoColumns = struct {
	ID         string
	MerchantID string
	ShopID     string
	CourseID   string
	TeacherID  string
	StartDate  string
	EndDate    string
	UseTime    string
	Status     string
	Name       string
	Summary    string
	Remark     string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	CreateBy   string
	UpdateBy   string
}{
	ID:         "id",
	MerchantID: "merchant_id",
	ShopID:     "shop_id",
	CourseID:   "course_id",
	TeacherID:  "teacher_id",
	StartDate:  "start_date",
	EndDate:    "end_date",
	UseTime:    "use_time",
	Status:     "status",
	Name:       "name",
	Summary:    "summary",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
}
