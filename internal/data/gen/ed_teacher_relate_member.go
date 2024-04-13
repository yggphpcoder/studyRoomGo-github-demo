package gen

// EdTeacherRelateMember 教务-老师-关联会员
type EdTeacherRelateMember struct {
	ID        int64 `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MemberID  int64 `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`              // 会员id
	TeacherID int64 `gorm:"column:teacher_id;type:bigint;not null;default:0;comment:'老师id'"`             // 老师id
}

// TableName get sql table name.获取数据库表名
func (m *EdTeacherRelateMember) TableName() string {
	return "ed_teacher_relate_member"
}

// EdTeacherRelateMemberColumns get sql column name.获取数据库列名
var EdTeacherRelateMemberColumns = struct {
	ID        string
	MemberID  string
	TeacherID string
}{
	ID:        "id",
	MemberID:  "member_id",
	TeacherID: "teacher_id",
}
