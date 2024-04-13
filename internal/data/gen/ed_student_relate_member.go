package gen

// EdStudentRelateMember 教务-学生-关联会员
type EdStudentRelateMember struct {
	ID        int64 `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MemberID  int64 `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`              // 会员id
	StudentID int64 `gorm:"column:student_id;type:bigint;not null;default:0;comment:'学生id'"`             // 学生id
}

// TableName get sql table name.获取数据库表名
func (m *EdStudentRelateMember) TableName() string {
	return "ed_student_relate_member"
}

// EdStudentRelateMemberColumns get sql column name.获取数据库列名
var EdStudentRelateMemberColumns = struct {
	ID        string
	MemberID  string
	StudentID string
}{
	ID:        "id",
	MemberID:  "member_id",
	StudentID: "student_id",
}
