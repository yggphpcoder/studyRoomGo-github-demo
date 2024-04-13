package gen

import "time"

// SrMember 会员信息表
type SrMember struct {
	ID         int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	AvatarPath *string    `gorm:"column:avatar_path;type:varchar(255);comment:'头像路径'"`                         // 头像路径
	Username   string     `gorm:"column:username;type:varchar(64);not null;comment:'用户名'"`                     // 用户名
	Password   string     `gorm:"column:password;type:varchar(128);not null;comment:'密码'"`                     // 密码
	NickName   *string    `gorm:"column:nick_name;type:varchar(128);comment:'昵称'"`                             // 昵称
	Phone      *string    `gorm:"column:phone;type:varchar(11);comment:'手机号'"`                                 // 手机号
	Type       int8       `gorm:"column:type;type:tinyint;not null;comment:'会员类型'"`                            // 会员类型
	Status     int8       `gorm:"column:status;type:tinyint;not null;comment:'会员状态'"`                          // 会员状态
	RegisterAt *time.Time `gorm:"column:register_at;type:timestamp;comment:'创建时间'"`                            // 创建时间
	Remark     *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy   *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy   *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrMember) TableName() string {
	return "sr_member"
}

// SrMemberColumns get sql column name.获取数据库列名
var SrMemberColumns = struct {
	ID         string
	AvatarPath string
	Username   string
	Password   string
	NickName   string
	Phone      string
	Type       string
	Status     string
	RegisterAt string
	Remark     string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	CreateBy   string
	UpdateBy   string
}{
	ID:         "id",
	AvatarPath: "avatar_path",
	Username:   "username",
	Password:   "password",
	NickName:   "nick_name",
	Phone:      "phone",
	Type:       "type",
	Status:     "status",
	RegisterAt: "register_at",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreateBy:   "create_by",
	UpdateBy:   "update_by",
}
