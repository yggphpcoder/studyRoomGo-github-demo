package gen

// EdStudentInfoExtend 教务-学生档案-扩展信息
type EdStudentInfoExtend struct {
	StudentID     int64         `gorm:"primaryKey;column:student_id;type:bigint;not null;default:0;comment:'学生id'"` // 学生id
	EdStudentInfo EdStudentInfo `gorm:"joinForeignKey:student_id;foreignKey:id;references:StudentID"`               // 教务-学生档案
	Field0        *string       `gorm:"column:field_0;type:varchar(255);comment:'字段0'"`                             // 字段0
	Field1        *string       `gorm:"column:field_1;type:varchar(255);comment:'字段1'"`                             // 字段1
	Field2        *string       `gorm:"column:field_2;type:varchar(255);comment:'字段2'"`                             // 字段2
	Field3        *string       `gorm:"column:field_3;type:varchar(255);comment:'字段3'"`                             // 字段3
	Field4        *string       `gorm:"column:field_4;type:varchar(255);comment:'字段4'"`                             // 字段4
	Field5        *string       `gorm:"column:field_5;type:varchar(255);comment:'字段5'"`                             // 字段5
	Field6        *string       `gorm:"column:field_6;type:varchar(255);comment:'字段6'"`                             // 字段6
	Field7        *string       `gorm:"column:field_7;type:varchar(255);comment:'字段7'"`                             // 字段7
	Field8        *string       `gorm:"column:field_8;type:varchar(255);comment:'字段8'"`                             // 字段8
	Field9        *string       `gorm:"column:field_9;type:varchar(255);comment:'字段9'"`                             // 字段9
	Field10       *string       `gorm:"column:field_10;type:varchar(255);comment:'字段10'"`                           // 字段10
}

// TableName get sql table name.获取数据库表名
func (m *EdStudentInfoExtend) TableName() string {
	return "ed_student_info_extend"
}

// EdStudentInfoExtendColumns get sql column name.获取数据库列名
var EdStudentInfoExtendColumns = struct {
	StudentID string
	Field0    string
	Field1    string
	Field2    string
	Field3    string
	Field4    string
	Field5    string
	Field6    string
	Field7    string
	Field8    string
	Field9    string
	Field10   string
}{
	StudentID: "student_id",
	Field0:    "field_0",
	Field1:    "field_1",
	Field2:    "field_2",
	Field3:    "field_3",
	Field4:    "field_4",
	Field5:    "field_5",
	Field6:    "field_6",
	Field7:    "field_7",
	Field8:    "field_8",
	Field9:    "field_9",
	Field10:   "field_10",
}
