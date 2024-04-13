package gen

import (
	"time"
)

// EdSalemanSalaryCourses 教务-销售员工资-课程细表
type EdSalemanSalaryCourses struct {
	ID                 int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`           // 主键编码
	SalemanSalaryID    int64      `gorm:"column:saleman_salary_id;type:bigint;not null;default:0;comment:'销售员工资id'"`             // 销售员工资id
	SaleType           string     `gorm:"column:sale_type;type:varchar(255);not null;default:0;comment:'销售类型'"`                  // 销售类型
	SaleChannel        string     `gorm:"column:sale_channel;type:varchar(255);not null;default:0;comment:'销售渠道'"`               // 销售渠道
	CourseID           int64      `gorm:"column:course_id;type:bigint;not null;default:0;comment:'课程id'"`                        // 课程id
	StudentID          int64      `gorm:"column:student_id;type:bigint;not null;default:0;comment:'学生id'"`                       // 学生id
	StudentCourseID    int64      `gorm:"column:student_course_id;type:bigint;not null;default:0;comment:'学生课程id'"`              // 学生课程id
	SalemanID          int64      `gorm:"column:saleman_id;type:bigint;not null;default:0;comment:'教师id'"`                       // 教师id
	SaleAt             time.Time  `gorm:"column:sale_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'销售日期'"`       // 销售日期
	CourseAmount       float64    `gorm:"column:course_amount;type:decimal(10,2);not null;default:0.00;comment:'课程总额'"`          // 课程总额
	CoursePercent      float64    `gorm:"column:course_percent;type:decimal(10,2);not null;default:0.00;comment:'课程提成比例'"`       // 课程提成比例
	CourseSalaryAmount float64    `gorm:"column:course_salary_amount;type:decimal(10,2);not null;default:0.00;comment:'课程工资总额'"` // 课程工资总额
	Subsidies          float64    `gorm:"column:subsidies;type:decimal(10,2);not null;default:0.00;comment:'补贴'"`                // 补贴
	Amount             float64    `gorm:"column:amount;type:decimal(10,2);not null;default:0.00;comment:'总额'"`                   // 总额
	Status             int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`                            // 状态
	Remark             *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                          // 备注
	CreatedAt          *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt          *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt          *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy           *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy           *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdSalemanSalaryCourses) TableName() string {
	return "ed_saleman_salary_courses"
}

// EdSalemanSalaryCoursesColumns get sql column name.获取数据库列名
var EdSalemanSalaryCoursesColumns = struct {
	ID                 string
	SalemanSalaryID    string
	SaleType           string
	SaleChannel        string
	CourseID           string
	StudentID          string
	StudentCourseID    string
	SalemanID          string
	SaleAt             string
	CourseAmount       string
	CoursePercent      string
	CourseSalaryAmount string
	Subsidies          string
	Amount             string
	Status             string
	Remark             string
	CreatedAt          string
	UpdatedAt          string
	DeletedAt          string
	CreateBy           string
	UpdateBy           string
}{
	ID:                 "id",
	SalemanSalaryID:    "saleman_salary_id",
	SaleType:           "sale_type",
	SaleChannel:        "sale_channel",
	CourseID:           "course_id",
	StudentID:          "student_id",
	StudentCourseID:    "student_course_id",
	SalemanID:          "saleman_id",
	SaleAt:             "sale_at",
	CourseAmount:       "course_amount",
	CoursePercent:      "course_percent",
	CourseSalaryAmount: "course_salary_amount",
	Subsidies:          "subsidies",
	Amount:             "amount",
	Status:             "status",
	Remark:             "remark",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
	CreateBy:           "create_by",
	UpdateBy:           "update_by",
}
