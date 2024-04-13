package gen

import (
	"time"
)

// EdCoursesScheduling 教务-课程-排课表
type EdCoursesScheduling struct {
	ID                int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID        int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	ShopID            int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	CourseID          int64      `gorm:"column:course_id;type:bigint;not null;default:0;comment:'课程id'"`              // 课程id
	StudentID         int64      `gorm:"column:student_id;type:bigint;not null;default:0;comment:'学生id'"`             // 学生id
	StudentCourseID   int64      `gorm:"column:student_course_id;type:bigint;not null;default:0;comment:'学生课程id'"`    // 学生课程id
	ClassID           int64      `gorm:"column:class_id;type:bigint;not null;default:0;comment:'班级id'"`               // 班级id
	ClassSchedulingID int64      `gorm:"column:class_scheduling_id;type:bigint;not null;default:0;comment:'班排课id'"`   // 班排课id
	TeacherID         int64      `gorm:"column:teacher_id;type:bigint;not null;default:0;comment:'教师id'"`             // 教师id
	StartDate         time.Time  `gorm:"column:start_date;type:date;not null;comment:'开始日期'"`                         // 开始日期
	EndDate           time.Time  `gorm:"column:end_date;type:date;not null;comment:'结束日期'"`                           // 结束日期
	StartTime         string     `gorm:"column:start_time;type:time;not null;comment:'开始时间'"`                         // 开始时间
	EndTime           string     `gorm:"column:end_time;type:time;not null;comment:'开始时间'"`                           // 开始时间
	ActualMin         int        `gorm:"column:actual_min;type:int;not null;default:0;comment:'实际上课分钟'"`              // 实际上课分钟
	CheckIn           int8       `gorm:"column:check_in;type:tinyint;not null;default:0;comment:'签到'"`                // 签到
	Status            int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`                  // 状态
	Remark            *string    `gorm:"column:remark;type:varchar(1000);comment:'备注'"`                               // 备注
	CreatedAt         *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt         *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy          *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy          *uint      `gorm:"column:update_by;type:int unsigned"`
	AddressID         int64      `gorm:"column:address_id;type:bigint;not null;default:0;comment:'上课地址'"` // 上课地址
}

// TableName get sql table name.获取数据库表名
func (m *EdCoursesScheduling) TableName() string {
	return "ed_courses_scheduling"
}

// EdCoursesSchedulingColumns get sql column name.获取数据库列名
var EdCoursesSchedulingColumns = struct {
	ID                string
	MerchantID        string
	ShopID            string
	CourseID          string
	StudentID         string
	StudentCourseID   string
	ClassID           string
	ClassSchedulingID string
	TeacherID         string
	StartDate         string
	EndDate           string
	StartTime         string
	EndTime           string
	ActualMin         string
	CheckIn           string
	Status            string
	Remark            string
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
	CreateBy          string
	UpdateBy          string
	AddressID         string
}{
	ID:                "id",
	MerchantID:        "merchant_id",
	ShopID:            "shop_id",
	CourseID:          "course_id",
	StudentID:         "student_id",
	StudentCourseID:   "student_course_id",
	ClassID:           "class_id",
	ClassSchedulingID: "class_scheduling_id",
	TeacherID:         "teacher_id",
	StartDate:         "start_date",
	EndDate:           "end_date",
	StartTime:         "start_time",
	EndTime:           "end_time",
	ActualMin:         "actual_min",
	CheckIn:           "check_in",
	Status:            "status",
	Remark:            "remark",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
	CreateBy:          "create_by",
	UpdateBy:          "update_by",
	AddressID:         "address_id",
}
