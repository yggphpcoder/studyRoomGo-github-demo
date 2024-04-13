package dto

import (
	pb "studyRoomGo/api/course/v1"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/model"
	"time"
)

type Course struct {
	Course *biz.Course
}

func (r *Course) Reply() *pb.Data {
	var saleList []*pb.Sale

	for _, item := range r.Course.Sale {
		saleList = append(saleList, &pb.Sale{
			Id:          item.ID,
			UseShop:     item.UseShop,
			Status:      int32(item.Status),
			Type:        int32(item.Type),
			Rule:        item.Rule,
			Price:       float32(item.Price),
			OriPrice:    float32(item.OriPrice),
			BuyLimit:    int32(item.BuyLimit),
			StartDate:   item.StartDate.Format(time.DateOnly),
			EndDate:     item.EndDate.Format(time.DateOnly),
			ValidDay:    int32(item.ValidDay),
			ActiveLimit: int32(item.ActiveLimit),
			Sort:        int32(item.Sort),
			Remark:      *item.Remark,
		})
	}
	return &pb.Data{
		Id:           r.Course.ID,
		TeachingType: int32(r.Course.TeachingType),
		CourseType:   int32(r.Course.CourseType),
		BuyLimit:     int32(r.Course.BuyLimit),
		CourseSize:   int32(r.Course.CourseSize),
		Title:        r.Course.Title,
		Cover:        *r.Course.Cover,
		Carousel:     *r.Course.Carousel,
		Summary:      r.Course.Summary,
		Content:      r.Course.Content,
		StartDate:    r.Course.StartDate.Format(time.DateOnly),
		EndDate:      r.Course.EndDate.Format(time.DateOnly),
		Status:       int32(r.Course.Status),
		UseShop:      r.Course.UseShop,
		Sort:         int32(r.Course.Sort),
		Sale:         saleList,
	}
}

type Student struct {
	Student          *model.MyStudent
	StudentCourse    *model.StudentCourse
	CourseScheduling *model.CourseScheduling
	ClassScheduling  *model.ClassScheduling
}

func (r *Student) Reply() *pb.MyStudentData {
	return &pb.MyStudentData{
		Id:        r.Student.EdStudentRelateMember.ID,
		StudentId: r.Student.Student.ID,
		Name:      r.Student.Student.Name,
		Phone:     r.Student.Student.Phone,
		Photo:     *r.Student.Student.Photo,
		Gender:    int32(r.Student.Student.Gender),
	}
}

func (r *Student) StudentCourseReply() *pb.MyStudentCourseData {
	return &pb.MyStudentCourseData{
		Id:           r.StudentCourse.ID,
		Title:        r.StudentCourse.Title,
		StudentId:    r.StudentCourse.StudentID,
		StudentName:  r.StudentCourse.Student.Name,
		ActualAmount: float32(r.StudentCourse.ActualAmount),
		Remain:       int32(r.StudentCourse.Remain),
		PerMin:       int32(r.StudentCourse.PerMin),
		Status:       int32(r.StudentCourse.Status),
		CourseId:     r.StudentCourse.CourseID,
		CourseType:   int32(r.StudentCourse.CourseType),
		TeachingType: int32(r.StudentCourse.TeachingType),
		CourseSize:   int32(r.StudentCourse.CourseSize),
		StartDate:    r.StudentCourse.StartDate.Format(time.DateOnly),
		EndDate:      r.StudentCourse.EndDate.Format(time.DateOnly),
		Price:        float32(r.StudentCourse.Price),
		OriPrice:     float32(r.StudentCourse.OriPrice),
		Remark:       string(r.StudentCourse.Remark),
		ValidDay:     int32(r.StudentCourse.ValidDay),
		ActiveLimit:  int32(r.StudentCourse.ActiveLimit),
		ActiveAt:     r.StudentCourse.ActiveAt.Format(time.DateTime),
		InvalidAt:    r.StudentCourse.InvalidAt.Format(time.DateTime),
	}
}

func (r *Student) CourseSchedulingReply() *pb.MyCourseSchedulingData {
	p := &pb.MyCourseSchedulingData{
		Id:              r.CourseScheduling.ID,
		StudentId:       r.CourseScheduling.StudentCourses.StudentID,
		StudentName:     r.CourseScheduling.Student.Name,
		StudentCourseId: r.CourseScheduling.StudentCourseID,
		CourseName:      r.CourseScheduling.StudentCourses.Title,
		CourseType:      int32(r.CourseScheduling.StudentCourses.CourseType),
		TeachingType:    int32(r.CourseScheduling.StudentCourses.TeachingType),
		TeacherId:       r.CourseScheduling.TeacherID,
		ClassId:         r.CourseScheduling.ClassID,
		StartDate:       r.CourseScheduling.StartDate.Format(time.DateOnly),
		EndDate:         r.CourseScheduling.EndDate.Format(time.DateOnly),
		StartTime:       r.CourseScheduling.StartTime,
		EndTime:         r.CourseScheduling.EndTime,
		Status:          int32(r.CourseScheduling.Status),
		CheckIn:         int32(r.CourseScheduling.CheckIn),
		AddressId:       r.CourseScheduling.AddressID,
	}
	if r.CourseScheduling.Teacher != nil {
		p.TeacherName = r.CourseScheduling.Teacher.Name
	}
	if r.CourseScheduling.Class != nil {
		p.ClassName = r.CourseScheduling.Class.Name
	}
	if r.CourseScheduling.Address != nil {
		p.AddressName = r.CourseScheduling.Address.Name
	}
	return p
}

func (r *Student) ClassSchedulingReply() *pb.MyClassSchedulingData {
	p := &pb.MyClassSchedulingData{
		Id:           r.ClassScheduling.ID,
		CourseName:   r.ClassScheduling.Course.Title,
		CourseType:   int32(r.ClassScheduling.Course.CourseType),
		TeachingType: int32(r.ClassScheduling.Course.TeachingType),
		TeacherId:    r.ClassScheduling.TeacherID,
		ClassId:      r.ClassScheduling.ClassID,
		StartDate:    r.ClassScheduling.StartDate.Format(time.DateOnly),
		EndDate:      r.ClassScheduling.EndDate.Format(time.DateOnly),
		StartTime:    r.ClassScheduling.StartTime,
		EndTime:      r.ClassScheduling.EndTime,
		Status:       int32(r.ClassScheduling.Status),
		AddressId:    r.ClassScheduling.AddressID,
	}
	if r.ClassScheduling.Teacher != nil {
		p.TeacherName = r.ClassScheduling.Teacher.Name
	}
	if r.ClassScheduling.Class != nil {
		p.ClassName = r.ClassScheduling.Class.Name
	}
	if r.ClassScheduling.Address != nil {
		p.AddressName = r.ClassScheduling.Address.Name
	}
	var course []*pb.MyClassSchedulingData_CourseSchedulingCheckIn
	for _, item := range r.ClassScheduling.CourseScheduling {
		course = append(course, &pb.MyClassSchedulingData_CourseSchedulingCheckIn{
			Id:          item.ID,
			CheckIn:     int32(item.CheckIn),
			StudentName: item.Student.Name,
		})
	}
	p.Course = course
	return p
}
