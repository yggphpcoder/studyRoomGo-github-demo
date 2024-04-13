package data

import (
	"context"
	"fmt"
	"strconv"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
)

type StudentCourseRepo struct {
	data *Data
	log  *log.Helper
}

// NewStudentCourseRepo .
func NewStudentCourseRepo(data *Data, logger log.Logger) biz.StudentCourseRepo {
	return &StudentCourseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *StudentCourseRepo) Save(ctx context.Context, model *model.StudentCourse) (bool, error) {
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	if merchantId != 0 {
		model.EdStudentCourses.MerchantID = int64(merchantId)
	}
	if shopId != 0 {
		model.EdStudentCourses.ShopID = int64(shopId)
	}
	mgr := gen.EdStudentCoursesMgr(r.data.db)
	mgr.Save(&model.EdStudentCourses)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *StudentCourseRepo) Update(ctx context.Context, model *model.StudentCourse) (bool, error) {
	mgr := gen.EdStudentCoursesMgr(r.data.db)
	if err := mgr.Select("*").Omit(Omit...).Updates(&model.EdStudentCourses).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *StudentCourseRepo) FindByID(ctx context.Context, id int64) (*model.StudentCourse, error) {
	mgr := gen.EdStudentCoursesMgr(r.data.db)
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	mgr.WithDeletedAt(nil)
	d, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &model.StudentCourse{EdStudentCourses: d}, nil
}

func (r *StudentCourseRepo) ListBy(ctx context.Context, req *biz.SearchStudentCourseRequest) ([]*model.StudentCourse, error) {
	mgr := gen.EdStudentCoursesMgr(r.data.db)
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	tableName := mgr.GetTableName()

	if req.MemberId != 0 {
		mgr.Select("ed_student_courses.*,ed_student_relate_member.member_id as member_id")
		mgr.Joins("Right JOIN ed_student_relate_member ON ed_student_relate_member.student_id = ed_student_courses.student_id AND ed_student_relate_member.member_id = ?", req.MemberId)
	}
	if req.StudentId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdStudentCoursesColumns.StudentID), req.StudentId)

	}

	if req.StudentCourseId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdStudentCoursesColumns.ID), req.StudentCourseId)

	}
	if req.CourseId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdStudentCoursesColumns.CourseID), req.CourseId)

	}
	if len(req.Status) != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v in ?", tableName, gen.EdStudentCoursesColumns.Status), req.Status)
	}
	if req.Sort > 0 {
		mgr.Order(SortType(req.Sort, tableName))
	} else {
		mgr.Order(SortType(sortStatusASort, tableName))
	}
	if req.PageSize > 0 && req.Page > 0 {
		mgr.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	mgr.Preload("Student")
	var modelList []*model.StudentCourse
	err := mgr.Find(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}
