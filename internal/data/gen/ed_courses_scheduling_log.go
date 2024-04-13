package gen

import "time"

// EdCoursesSchedulingLog 排课操作
type EdCoursesSchedulingLog struct {
	ID                  int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	CoursesSchedulingID int64      `gorm:"column:courses_scheduling_id;type:bigint;not null;default:0;comment:'订座id'"`  // 订座id
	Operate             int8       `gorm:"column:operate;type:tinyint;not null;default:0;comment:'操作'"`                 // 操作
	Extra               string     `gorm:"column:extra;type:varchar(1000);comment:'额外信息'"`                              // 额外信息
	Remark              string     `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt           *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	CreateBy            *uint      `gorm:"column:create_by;type:int unsigned"`
	DeletedAt           *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	UpdatedAt           *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateBy            *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdCoursesSchedulingLog) TableName() string {
	return "ed_courses_scheduling_log"
}

// EdCoursesSchedulingLogColumns get sql column name.获取数据库列名
var EdCoursesSchedulingLogColumns = struct {
	ID                  string
	CoursesSchedulingID string
	Operate             string
	Extra               string
	Remark              string
	CreatedAt           string
	CreateBy            string
	DeletedAt           string
	UpdatedAt           string
	UpdateBy            string
}{
	ID:                  "id",
	CoursesSchedulingID: "courses_scheduling_id",
	Operate:             "operate",
	Extra:               "extra",
	Remark:              "remark",
	CreatedAt:           "created_at",
	CreateBy:            "create_by",
	DeletedAt:           "deleted_at",
	UpdatedAt:           "updated_at",
	UpdateBy:            "update_by",
}
