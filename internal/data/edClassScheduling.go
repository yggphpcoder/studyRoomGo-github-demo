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

type ClassSchedulingRepo struct {
	data *Data
	log  *log.Helper
}

// NewClassSchedulingRepo .
func NewClassSchedulingRepo(data *Data, logger log.Logger) biz.ClassSchedulingRepo {
	return &ClassSchedulingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ClassSchedulingRepo) Update(ctx context.Context, data *gen.EdClassScheduling) (bool, error) {
	mgr := gen.EdClassSchedulingMgr(r.data.db)
	if err := mgr.Select("*").Omit(Omit...).Updates(&data).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *ClassSchedulingRepo) FindByID(ctx context.Context, id int64) (*gen.EdClassScheduling, error) {
	mgr := gen.EdClassSchedulingMgr(r.data.db)
	mgr.Where("id=?", id)
	mgr.WithDeletedAt(nil)
	var model *gen.EdClassScheduling

	err := mgr.First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *ClassSchedulingRepo) FindDetailByID(ctx context.Context, id int64) (*model.ClassScheduling, error) {
	mgr := gen.EdClassSchedulingMgr(r.data.db)
	mgr.Where("id=?", id)
	mgr.WithDeletedAt(nil)
	mgr.Preload("CourseScheduling")
	mgr.Preload("CourseScheduling.Student")
	mgr.Preload("Course")
	mgr.Preload("Teacher")
	mgr.Preload("Class")
	mgr.Preload("Address")
	var model *model.ClassScheduling

	err := mgr.First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (r *ClassSchedulingRepo) ListBy(ctx context.Context, req *biz.SearchSchedulingRequest) ([]*model.ClassScheduling, error) {
	mgr := gen.EdClassSchedulingMgr(r.data.db)
	merchantId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if merchantId != 0 {
		mgr.Where("merchant_id=?", merchantId)
	}
	tableName := mgr.GetTableName()

	if req.IsTeacher && req.MemberId != 0 {
		mgr.Select("ed_class_scheduling.*,ed_teacher_relate_member.member_id as member_id")
		mgr.Joins(fmt.Sprintf("Right JOIN ed_teacher_relate_member ON ed_teacher_relate_member.teacher_id = `%v`.teacher_id AND ed_teacher_relate_member.member_id = %v", tableName, req.MemberId))
	}

	if req.TeacherId != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v = ?", tableName, gen.EdClassSchedulingColumns.TeacherID), req.TeacherId)

	}

	if len(req.Status) != 0 {
		mgr.Where(fmt.Sprintf("`%v`.%v in ?", tableName, gen.EdClassSchedulingColumns.Status), req.Status)
	}
	if req.Sort > 0 {
		mgr.Order(SortType(req.Sort, tableName))
	} else {
		mgr.Order(SortType(sortStatusASort, tableName))
	}
	if req.PageSize > 0 && req.Page > 0 {
		mgr.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	mgr.Preload("CourseScheduling")
	mgr.Preload("CourseScheduling.Student")
	mgr.Preload("Course")
	mgr.Preload("Teacher")
	mgr.Preload("Class")
	mgr.Preload("Address")
	var modelList []*model.ClassScheduling
	err := mgr.Find(&modelList).Error
	if err != nil {
		return nil, err
	}

	return modelList, nil
}
