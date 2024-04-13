package gen

import "time"

// EdStudentCoursesSaleman 教务-学生-课程-营销信息
type EdStudentCoursesSaleman struct {
	StudentCourseID int64      `gorm:"primaryKey;column:student_course_id;type:bigint;not null;default:0;comment:'学生-课程_id'"` // 学生-课程_id
	SaleType        string     `gorm:"column:sale_type;type:varchar(255);not null;default:'';comment:'销售类型'"`                 // 销售类型
	SaleChannel     string     `gorm:"column:sale_channel;type:varchar(255);not null;default:'';comment:'销售渠道'"`              // 销售渠道
	SalemanID       int64      `gorm:"column:saleman_id;type:bigint;not null;default:0;comment:'销售人员id'"`                     // 销售人员id
	Remark          *string    `gorm:"column:remark;type:varchar(1000);comment:'备注'"`                                         // 备注
	CreatedAt       *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt       *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy        *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy        *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdStudentCoursesSaleman) TableName() string {
	return "ed_student_courses_saleman"
}

// EdStudentCoursesSalemanColumns get sql column name.获取数据库列名
var EdStudentCoursesSalemanColumns = struct {
	StudentCourseID string
	SaleType        string
	SaleChannel     string
	SalemanID       string
	Remark          string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	CreateBy        string
	UpdateBy        string
}{
	StudentCourseID: "student_course_id",
	SaleType:        "sale_type",
	SaleChannel:     "sale_channel",
	SalemanID:       "saleman_id",
	Remark:          "remark",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	CreateBy:        "create_by",
	UpdateBy:        "update_by",
}
