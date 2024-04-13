package gen

import (
	"time"
)

// EdTeacherSalaryCourses 教务-教师档案-工资-课程明细表
type EdTeacherSalaryCourses struct {
	ID                 int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`           // 主键编码
	TeacherSalaryID    int64      `gorm:"column:teacher_salary_id;type:bigint;not null;default:0;comment:'教师工资id'"`              // 教师工资id
	TeachingType       int8       `gorm:"column:teaching_type;type:tinyint;not null;default:0;comment:'教学模式'"`                   // 教学模式
	SchedulingID       int64      `gorm:"column:scheduling_id;type:bigint;not null;default:0;comment:'课程排课id'"`                  // 课程排课id
	ClassID            int64      `gorm:"column:class_id;type:bigint;not null;default:0;comment:'班级id'"`                         // 班级id
	CourseID           int64      `gorm:"column:course_id;type:bigint;not null;default:0;comment:'课程id'"`                        // 课程id
	StudentID          int64      `gorm:"column:student_id;type:bigint;not null;default:0;comment:'学生id'"`                       // 学生id
	StudentCourseID    int64      `gorm:"column:student_course_id;type:bigint;not null;default:0;comment:'学生课程id'"`              // 学生课程id
	TeacherID          int64      `gorm:"column:teacher_id;type:bigint;not null;default:0;comment:'教师id'"`                       // 教师id
	StartDate          time.Time  `gorm:"column:start_date;type:date;not null;comment:'开始日期'"`                                   // 开始日期
	EndDate            time.Time  `gorm:"column:end_date;type:date;not null;comment:'结束日期'"`                                     // 结束日期
	StartTime          time.Time  `gorm:"column:start_time;type:time;not null;comment:'开始时间'"`                                   // 开始时间
	EndTime            time.Time  `gorm:"column:end_time;type:time;not null;comment:'开始时间'"`                                     // 开始时间
	CourseSalaryPer    float64    `gorm:"column:course_salary_per;type:decimal(10,2);not null;default:0.00;comment:'每课时工资(分钟)'"` // 每课时工资(分钟)
	ActualMin          int        `gorm:"column:actual_min;type:int;not null;default:0;comment:'实际上课分钟'"`                        // 实际上课分钟
	CourseSalaryAmount float64    `gorm:"column:course_salary_amount;type:decimal(10,2);not null;default:0.00;comment:'课时工资总额'"` // 课时工资总额
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
func (m *EdTeacherSalaryCourses) TableName() string {
	return "ed_teacher_salary_courses"
}

// EdTeacherSalaryCoursesColumns get sql column name.获取数据库列名
var EdTeacherSalaryCoursesColumns = struct {
	ID                 string
	TeacherSalaryID    string
	TeachingType       string
	SchedulingID       string
	ClassID            string
	CourseID           string
	StudentID          string
	StudentCourseID    string
	TeacherID          string
	StartDate          string
	EndDate            string
	StartTime          string
	EndTime            string
	CourseSalaryPer    string
	ActualMin          string
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
	TeacherSalaryID:    "teacher_salary_id",
	TeachingType:       "teaching_type",
	SchedulingID:       "scheduling_id",
	ClassID:            "class_id",
	CourseID:           "course_id",
	StudentID:          "student_id",
	StudentCourseID:    "student_course_id",
	TeacherID:          "teacher_id",
	StartDate:          "start_date",
	EndDate:            "end_date",
	StartTime:          "start_time",
	EndTime:            "end_time",
	CourseSalaryPer:    "course_salary_per",
	ActualMin:          "actual_min",
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
