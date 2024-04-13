package biz

type SearchStudentCourseRequest struct {
	ShopId          int64
	MemberId        int64
	StudentCourseId int64
	StudentId       int64
	CourseId        int64
	Status          []int32

	Sort     int
	Page     int
	PageSize int
}

type SearchStudentRequest struct {
	ShopId   int64
	MemberId int64

	StudentCourseId int64

	Sort     int
	Page     int
	PageSize int
}

type SearchSchedulingRequest struct {
	ShopId   int64
	MemberId int64

	IsStudent bool
	IsTeacher bool

	StudentCourseId int64
	ClassId         *int64
	StudentId       int64
	TeacherId       int64
	CourseId        int64
	Status          []int32
	Sort            int
	Page            int
	PageSize        int
}

type StudentCourseCheckIn struct {
	Id              int
	StudentCourseId int
	StudentId       int
	OldCheckIn      int
	NewCheckIn      int
	Err             error
}
