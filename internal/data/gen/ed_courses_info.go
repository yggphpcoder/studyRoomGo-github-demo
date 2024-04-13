package gen

import "time"

// EdCoursesInfo 教务-课程档案
type EdCoursesInfo struct {
	ID           int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID   int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	CourseType   int8       `gorm:"column:course_type;type:tinyint;not null;default:0;comment:'课程类型'"`           // 课程类型
	TeachingType int8       `gorm:"column:teaching_type;type:tinyint;not null;default:0;comment:'教学模式'"`         // 教学模式
	BuyLimit     int        `gorm:"column:buy_limit;type:int;not null;default:0;comment:'购买限制'"`                 // 购买限制
	CourseSize   int        `gorm:"column:course_size;type:int;not null;default:0;comment:'课程人数'"`               // 课程人数
	Cover        *string    `gorm:"column:cover;type:varchar(255);comment:'封面'"`                                 // 封面
	Carousel     *string    `gorm:"column:carousel;type:varchar(1000);comment:'轮播图'"`                            // 轮播图
	Title        string     `gorm:"column:title;type:varchar(255);not null;default:'';comment:'标题'"`             // 标题
	Summary      string     `gorm:"column:summary;type:varchar(255);not null;default:'';comment:'简介'"`           // 简介
	Content      string     `gorm:"column:content;type:varchar(1000);not null;default:'';comment:'详情'"`          // 详情
	StartDate    *time.Time `gorm:"column:start_date;type:datetime;comment:'开课日期'"`                              // 开课日期
	EndDate      *time.Time `gorm:"column:end_date;type:datetime;comment:'课程结束日期'"`                              // 课程结束日期
	UseShop      string     `gorm:"column:use_shop;type:varchar(500);not null;default:0;comment:'适用店铺'"`         // 适用店铺
	CreatedAt    *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt    *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy     *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy     *uint      `gorm:"column:update_by;type:int unsigned"`
	Status       int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'课程状态'"`     // 课程状态
	IsShow       int8       `gorm:"column:is_show;type:tinyint;not null;default:0;comment:'是否在前台显示'"` // 是否在前台显示
	Sort         int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`             // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EdCoursesInfo) TableName() string {
	return "ed_courses_info"
}

// EdCoursesInfoColumns get sql column name.获取数据库列名
var EdCoursesInfoColumns = struct {
	ID           string
	MerchantID   string
	CourseType   string
	TeachingType string
	BuyLimit     string
	CourseSize   string
	Cover        string
	Carousel     string
	Title        string
	Summary      string
	Content      string
	StartDate    string
	EndDate      string
	UseShop      string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	CreateBy     string
	UpdateBy     string
	Status       string
	IsShow       string
	Sort         string
}{
	ID:           "id",
	MerchantID:   "merchant_id",
	CourseType:   "course_type",
	TeachingType: "teaching_type",
	BuyLimit:     "buy_limit",
	CourseSize:   "course_size",
	Cover:        "cover",
	Carousel:     "carousel",
	Title:        "title",
	Summary:      "summary",
	Content:      "content",
	StartDate:    "start_date",
	EndDate:      "end_date",
	UseShop:      "use_shop",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	CreateBy:     "create_by",
	UpdateBy:     "update_by",
	Status:       "status",
	IsShow:       "is_show",
	Sort:         "sort",
}
