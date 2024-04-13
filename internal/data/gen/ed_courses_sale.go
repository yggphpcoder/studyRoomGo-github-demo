package gen

import "time"

// EdCoursesSale 教务-课程-销售
type EdCoursesSale struct {
	ID          int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	CourseID    int64      `gorm:"column:course_id;type:bigint;not null;default:0;comment:'课程id'"`              // 课程id
	UseShop     string     `gorm:"column:use_shop;type:varchar(500);not null;default:0;comment:'适用店铺'"`         // 适用店铺
	Status      int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'营销状态'"`                // 营销状态
	Type        int8       `gorm:"column:type;type:tinyint;not null;default:0;comment:'销售类型'"`                  // 销售类型
	Rule        string     `gorm:"column:rule;type:varchar(1000);not null;default:'';comment:'销售规则'"`           // 销售规则
	Price       float64    `gorm:"column:price;type:float(10,4);not null;default:0.0000;comment:'售价'"`          // 售价
	OriPrice    float64    `gorm:"column:ori_price;type:float(10,4);not null;default:0.0000;comment:'原价'"`      // 原价
	IsShow      int8       `gorm:"column:is_show;type:tinyint;not null;default:0;comment:'是否在前台显示'"`            // 是否在前台显示
	BuyLimit    int        `gorm:"column:buy_limit;type:int;not null;default:0;comment:'购买限制'"`                 // 购买限制
	StartDate   *time.Time `gorm:"column:start_date;type:datetime;comment:'售卖开始日期'"`                            // 售卖开始日期
	EndDate     *time.Time `gorm:"column:end_date;type:datetime;comment:'售卖结束日期'"`                              // 售卖结束日期
	ValidDay    int        `gorm:"column:valid_day;type:int;not null;default:0;comment:'默认有效天数'"`               // 默认有效天数
	ActiveLimit int        `gorm:"column:active_limit;type:int;not null;default:0;comment:'默认激活期限天数'"`          // 默认激活期限天数
	Sort        int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark      *string    `gorm:"column:remark;type:varchar(1000);comment:'备注'"`                               // 备注
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy    *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy    *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdCoursesSale) TableName() string {
	return "ed_courses_sale"
}

// EdCoursesSaleColumns get sql column name.获取数据库列名
var EdCoursesSaleColumns = struct {
	ID          string
	CourseID    string
	UseShop     string
	Status      string
	Type        string
	Rule        string
	Price       string
	OriPrice    string
	IsShow      string
	BuyLimit    string
	StartDate   string
	EndDate     string
	ValidDay    string
	ActiveLimit string
	Sort        string
	Remark      string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	CreateBy    string
	UpdateBy    string
}{
	ID:          "id",
	CourseID:    "course_id",
	UseShop:     "use_shop",
	Status:      "status",
	Type:        "type",
	Rule:        "rule",
	Price:       "price",
	OriPrice:    "ori_price",
	IsShow:      "is_show",
	BuyLimit:    "buy_limit",
	StartDate:   "start_date",
	EndDate:     "end_date",
	ValidDay:    "valid_day",
	ActiveLimit: "active_limit",
	Sort:        "sort",
	Remark:      "remark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreateBy:    "create_by",
	UpdateBy:    "update_by",
}
