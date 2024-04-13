package gen

import "time"

// EdTeacherGrowup 教务-教师档案-成长记录
type EdTeacherGrowup struct {
	ID        int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	TeacherID int64      `gorm:"column:teacher_id;type:bigint;not null;default:0;comment:'教师id'"`             // 教师id
	Type      int8       `gorm:"column:type;type:tinyint;not null;comment:'记录分类'"`                            // 记录分类
	Title     string     `gorm:"column:title;type:varchar(255);not null;default:'';comment:'标题'"`             // 标题
	Content   *string    `gorm:"column:content;type:text;comment:'内容'"`                                       // 内容
	Files     *string    `gorm:"column:files;type:text;comment:'附件'"`                                         // 附件
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy  *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy  *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdTeacherGrowup) TableName() string {
	return "ed_teacher_growup"
}

// EdTeacherGrowupColumns get sql column name.获取数据库列名
var EdTeacherGrowupColumns = struct {
	ID        string
	TeacherID string
	Type      string
	Title     string
	Content   string
	Files     string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	CreateBy  string
	UpdateBy  string
}{
	ID:        "id",
	TeacherID: "teacher_id",
	Type:      "type",
	Title:     "title",
	Content:   "content",
	Files:     "files",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
}
