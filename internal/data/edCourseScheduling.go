package data

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
)

type CourseSchedulingRepo struct {
	data *Data
	log  *log.Helper
}

// NewCourseSchedulingRepo .
func NewCourseSchedulingRepo(data *Data, logger log.Logger) biz.CourseSchedulingRepo {
	return &CourseSchedulingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CourseSchedulingRepo) Update(ctx context.Context, data *gen.EdCoursesScheduling) (bool, error) {
	mgr := gen.EdCoursesSchedulingMgr(r.data.db)
	if err := mgr.Select("*").Omit(Omit...).Updates(&data).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (r *CourseSchedulingRepo) RecordLog(oldData *gen.EdCoursesScheduling, data *gen.EdCoursesScheduling, by int) (bool, error) {

	extra := make(map[string]string)

	if oldData.Status != data.Status {
		extra["状态"] = fmt.Sprintf("%v - %v", oldData.Status, data.Status)
	}

	if oldData.TeacherID != data.TeacherID {
		extra["老师"] = fmt.Sprintf("%v - %v", oldData.CheckIn, data.TeacherID)
	}
	if oldData.ClassID != data.ClassID {
		extra["课室"] = fmt.Sprintf("%v - %v", oldData.ClassID, data.ClassID)
	}
	if oldData.ActualMin != data.ActualMin {
		extra["实际分钟"] = fmt.Sprintf("%v - %v", oldData.ActualMin, data.ActualMin)
	}
	if oldData.StartDate != data.StartDate {
		extra["上课日期"] = fmt.Sprintf("%v - %v", oldData.StartDate, data.StartDate)
	}
	if oldData.StartTime != data.StartTime {
		extra["上课时间"] = fmt.Sprintf("%v - %v", oldData.StartTime, data.StartTime)
	}
	if oldData.EndTime != data.EndTime {
		extra["下课时间"] = fmt.Sprintf("%v - %v", oldData.EndTime, data.EndTime)
	}
	if len(extra) > 0 {
		uby := uint(by)
		jsonExtra, _ := json.Marshal(extra)
		r.data.db.Create(&gen.EdCoursesSchedulingLog{
			CoursesSchedulingID: int64(data.ID),
			Operate:             2,
			Extra:               string(jsonExtra),
			Remark:              "修改",
			CreateBy:            &uby,
			UpdateBy:            &uby,
		})

	}
	return true, nil
}
func (r *CourseSchedulingRepo) RecordCheckInLog(course biz.StudentCourseCheckIn, by int) (bool, error) {

	extra := make(map[string]string)

	extra["报到"] = fmt.Sprintf("%v - %v", course.OldCheckIn, course.NewCheckIn)

	if course.Err != nil {
		extra["发送服务号通知"] = fmt.Sprintf("%v", course.Err.Error())
	}
	if len(extra) > 0 {
		opType := int8(5)
		remark := "修改报到"
		if course.NewCheckIn == 2 {
			opType = 3
			remark = "修改报到:签到"
		}
		if course.NewCheckIn == 3 {
			opType = 4
			remark = "修改报到:签退"
		}
		uby := uint(by)
		jsonExtra, _ := json.Marshal(extra)
		operateLog := gen.EdCoursesSchedulingLog{
			CoursesSchedulingID: int64(course.Id),
			Operate:             opType,
			Extra:               string(jsonExtra),
			Remark:              remark,
			CreateBy:            &uby,
			UpdateBy:            &uby,
		}
		r.data.db.Create(&operateLog)
	}
	return true, nil
}
func (r *CourseSchedulingRepo) FindByID(ctx context.Context, id int64) (*gen.EdCoursesScheduling, error) {
	mgr := gen.EdCoursesSchedulingMgr(r.data.db)
	mgr.Where("id=?", id)
	mgr.WithDeletedAt(nil)
	var model *gen.EdCoursesScheduling

	err := mgr.First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (r *CourseSchedulingRepo) FindDetailByID(ctx context.Context, id int64) (*model.CourseScheduling, error) {
	mgr := gen.EdCoursesSchedulingMgr(r.data.db)
	mgr.Where("id=?", id)
	mgr.WithDeletedAt(nil)
	mgr.Preload("StudentCourses")
	mgr.Preload("Student")
	mgr.Preload("Teacher")
	mgr.Preload("Class")
	mgr.Preload("Address")
	var model *model.CourseScheduling

	err := mgr.First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (r *CourseSchedulingRepo) ListBy(ctx context.Context, req *biz.SearchSchedulingRequest) ([]*model.CourseScheduling, error) {
	mgr := gen.EdCoursesSchedulingMgr(r.data.db)
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	tableName := mgr.GetTableName()

	if req.IsStudent && req.MemberId != 0 {
		mgr.Select("ed_courses_scheduling.*,ed_student_relate_member.member_id as member_id")
		mgr.Joins(fmt.Sprintf("Right JOIN ed_student_relate_member ON ed_student_relate_member.student_id = `%v`.student_id AND ed_student_relate_member.member_id = %v", tableName, req.MemberId))
	}
	if req.IsTeacher && req.MemberId != 0 {
		mgr.Select("ed_courses_scheduling.*,ed_teacher_relate_member.member_id as member_id")
		mgr.Joins(fmt.Sprintf("Right JOIN ed_teacher_relate_member ON ed_teacher_relate_member.teacher_id = `%v`.teacher_id AND ed_teacher_relate_member.member_id = %v", tableName, req.MemberId))
	}
	if req.StudentId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdCoursesSchedulingColumns.StudentID), req.StudentId)

	}
	if req.TeacherId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdCoursesSchedulingColumns.TeacherID), req.TeacherId)

	}
	if req.ClassId != nil {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdCoursesSchedulingColumns.ClassID), req.ClassId)

	}
	if req.StudentCourseId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdCoursesSchedulingColumns.StudentCourseID), req.StudentCourseId)

	}
	if req.CourseId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdCoursesSchedulingColumns.CourseID), req.CourseId)

	}
	if len(req.Status) != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v in ?", tableName, gen.EdCoursesSchedulingColumns.Status), req.Status)
	}
	if req.Sort > 0 {
		mgr.Order(SortType(req.Sort, tableName))
	} else {
		mgr.Order(SortType(sortStatusASort, tableName))
	}
	if req.PageSize > 0 && req.Page > 0 {
		mgr.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	mgr.Preload("StudentCourses")
	mgr.Preload("Student")
	mgr.Preload("Teacher")
	mgr.Preload("Class")
	mgr.Preload("Address")
	var modelList []*model.CourseScheduling
	err := mgr.Find(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}
