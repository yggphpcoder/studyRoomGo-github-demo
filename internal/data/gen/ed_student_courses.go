package gen

import "time"

// EdStudentCourses 教务-学生-课程
type EdStudentCourses struct {
	ID              int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`  // 主键编码
	MerchantID      int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`             // 商家id
	ShopID          int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                 // 店铺id
	StudentID       int64      `gorm:"column:student_id;type:bigint;not null;default:0;comment:'学生id'"`              // 学生id
	ActualAmount    float64    `gorm:"column:actual_amount;type:float(10,4);not null;default:0.0000;comment:'实收价格'"` // 实收价格
	Remain          int        `gorm:"column:remain;type:int;not null;default:0;comment:'剩余课数'"`                     // 剩余课数
	PerMin          int        `gorm:"column:per_min;type:int;not null;default:0;comment:'每堂分钟'"`                    // 每堂分钟
	Status          int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`                   // 状态
	CourseID        int64      `gorm:"column:course_id;type:bigint;not null;default:0;comment:'课程id'"`               // 课程id
	Title           string     `gorm:"column:title;type:varchar(255);not null;default:'';comment:'标题'"`              // 标题
	CourseType      int8       `gorm:"column:course_type;type:tinyint;not null;default:0;comment:'课程类型'"`            // 课程类型
	TeachingType    int8       `gorm:"column:teaching_type;type:tinyint;not null;default:0;comment:'教学模式'"`          // 教学模式
	CourseSize      int        `gorm:"column:course_size;type:int;not null;default:0;comment:'课程人数'"`                // 课程人数
	StartDate       *time.Time `gorm:"column:start_date;type:datetime;comment:'开课日期'"`                               // 开课日期
	EndDate         *time.Time `gorm:"column:end_date;type:datetime;comment:'结课日期'"`                                 // 结课日期
	CourseSaleID    int64      `gorm:"column:course_sale_id;type:bigint;not null;default:0;comment:'课程营销id'"`        // 课程营销id
	CourseSaleType  int8       `gorm:"column:course_sale_type;type:tinyint;not null;default:0;comment:'营销类型'"`       // 营销类型
	Rule            string     `gorm:"column:rule;type:varchar(1000);not null;default:'';comment:'营销规则'"`            // 营销规则
	Price           float64    `gorm:"column:price;type:float(10,4);not null;default:0.0000;comment:'售价'"`           // 售价
	OriPrice        float64    `gorm:"column:ori_price;type:float(10,4);not null;default:0.0000;comment:'原价'"`       // 原价
	ActiveAt        *time.Time `gorm:"column:active_at;type:timestamp;comment:'激活时间'"`                               // 激活时间
	InvalidAt       *time.Time `gorm:"column:invalid_at;type:datetime;comment:'失效时间'"`                               // 失效时间
	ValidDay        int        `gorm:"column:valid_day;type:int;not null;default:0;comment:'有效天数'"`                  // 有效天数
	ActiveLimit     int        `gorm:"column:active_limit;type:int;not null;default:0;comment:'激活期限天数'"`             // 激活期限天数
	Sort            int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                         // 排序
	CourseSalemanID int64      `gorm:"column:course_saleman_id;type:bigint;not null;default:0;comment:'课程销售员id'"`    // 课程销售员id
	Remark          string     `gorm:"column:remark;type:varchar(1000);not null;default:'';comment:'备注'"`            // 备注
	CreatedAt       *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt       *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy        *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy        *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdStudentCourses) TableName() string {
	return "ed_student_courses"
}

// EdStudentCoursesColumns get sql column name.获取数据库列名
var EdStudentCoursesColumns = struct {
	ID              string
	MerchantID      string
	ShopID          string
	StudentID       string
	ActualAmount    string
	Remain          string
	PerMin          string
	Status          string
	CourseID        string
	Title           string
	CourseType      string
	TeachingType    string
	CourseSize      string
	StartDate       string
	EndDate         string
	CourseSaleID    string
	CourseSaleType  string
	Rule            string
	Price           string
	OriPrice        string
	ActiveAt        string
	InvalidAt       string
	ValidDay        string
	ActiveLimit     string
	Sort            string
	CourseSalemanID string
	Remark          string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	CreateBy        string
	UpdateBy        string
}{
	ID:              "id",
	MerchantID:      "merchant_id",
	ShopID:          "shop_id",
	StudentID:       "student_id",
	ActualAmount:    "actual_amount",
	Remain:          "remain",
	PerMin:          "per_min",
	Status:          "status",
	CourseID:        "course_id",
	Title:           "title",
	CourseType:      "course_type",
	TeachingType:    "teaching_type",
	CourseSize:      "course_size",
	StartDate:       "start_date",
	EndDate:         "end_date",
	CourseSaleID:    "course_sale_id",
	CourseSaleType:  "course_sale_type",
	Rule:            "rule",
	Price:           "price",
	OriPrice:        "ori_price",
	ActiveAt:        "active_at",
	InvalidAt:       "invalid_at",
	ValidDay:        "valid_day",
	ActiveLimit:     "active_limit",
	Sort:            "sort",
	CourseSalemanID: "course_saleman_id",
	Remark:          "remark",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	CreateBy:        "create_by",
	UpdateBy:        "update_by",
}
