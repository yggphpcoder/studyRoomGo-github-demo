package model

import (
	"studyRoomGo/internal/data/gen"
)

type Course struct {
	gen.EdCoursesInfo
	Sale []*gen.EdCoursesSale `json:"sale" gorm:"foreignKey:CourseID"`
}

type CourseSale struct {
	gen.EdCoursesSale
	Course *gen.EdCoursesInfo `json:"course" gorm:"foreignKey:CourseID"`
}

type CourseScheduling struct {
	MemberId int64 `json:"member_id" gorm:"column:member_id;"`
	gen.EdCoursesScheduling
	StudentCourses *gen.EdStudentCourses           `json:"studentCourses" gorm:"foreignKey:StudentCourseID"`
	Student        *gen.EdStudentInfo              `json:"student" gorm:"foreignKey:StudentID;references:ID;"`
	Teacher        *gen.EdTeacherInfo              `json:"teacher" gorm:"foreignKey:TeacherID;references:ID;"`
	Class          *gen.EdClassInfo                `json:"class" gorm:"foreignKey:ClassID;references:ID;"`
	Address        *gen.EdCoursesSchedulingAddress `json:"address" gorm:"foreignKey:AddressID;references:ID;"`
}

type MyStudent struct {
	gen.EdStudentRelateMember
	Student *gen.EdStudentInfo `json:"student" gorm:"foreignKey:StudentID"`
}

type StudentCourse struct {
	MemberId int64 `json:"member_id" gorm:"column:member_id;"`
	gen.EdStudentCourses
	Student gen.EdStudentInfo `json:"student" gorm:"foreignKey:StudentID;references:ID;"`
}

type ClassScheduling struct {
	MemberId int64 `json:"member_id" gorm:"column:member_id;"`
	gen.EdClassScheduling
	Course           *gen.EdCoursesInfo              `json:"course" gorm:"foreignKey:CourseID"`
	CourseScheduling []*ClassCourseScheduling        `json:"courseScheduling"  gorm:"references:ID;"`
	Teacher          *gen.EdTeacherInfo              `json:"teacher" gorm:"foreignKey:TeacherID;references:ID;"`
	Class            *gen.EdClassInfo                `json:"class" gorm:"foreignKey:ClassID;references:ID;"`
	Address          *gen.EdCoursesSchedulingAddress `json:"address" gorm:"foreignKey:AddressID;references:ID;"`
}

type ClassCourseScheduling struct {
	*gen.EdCoursesScheduling
	Student *gen.EdStudentInfo `json:"student" gorm:"foreignKey:StudentID;references:ID;"`
}
