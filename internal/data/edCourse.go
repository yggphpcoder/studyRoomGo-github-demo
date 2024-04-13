package data

import (
	"context"
	"fmt"
	"strings"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type CourseRepo struct {
	data *Data
	log  *log.Helper
}

// NewCourseRepo .
func NewCourseRepo(data *Data, logger log.Logger) biz.CourseRepo {
	return &CourseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CourseRepo) FindByID(ctx context.Context, id int64) (*model.Course, error) {
	var data model.Course
	err := gen.EdCoursesInfoMgr(r.data.db).Model(model.Course{}).Preload("Sale").First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *CourseRepo) FindBy(ctx context.Context, key string, value string) (*model.Course, error) {
	var data *model.Course

	condition := gen.Condition{}
	condition.And(key, "=", value)

	where, out := condition.Get()

	mgr := gen.EdCoursesInfoMgr(r.data.db).Model(model.Course{}).Preload("Sale").Where(where, out...)
	mgr.Where("deleted_at IS NULL")

	err := mgr.First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *CourseRepo) ListBy(ctx context.Context, req *biz.SearchCourseRequest) ([]*model.Course, error) {
	mgr := gen.EdCoursesInfoMgr(r.data.db).Model(model.Course{}).Preload("Sale")

	condition := gen.Condition{}
	if len(req.CId) > 0 {
		condition.And(gen.EdCoursesInfoColumns.ID, "in", req.CId)
	}
	if req.MerchantId != 0 {
		condition.And(gen.EdCoursesInfoColumns.MerchantID, "=", req.MerchantId)
	}
	if req.Status != 0 {
		condition.And(gen.EdCoursesInfoColumns.Status, "=", req.Status)
	}
	if req.CourseType != 0 {
		condition.And(gen.EdCoursesInfoColumns.CourseType, "=", req.CourseType)
	}

	if req.TeachingType != 0 {
		condition.And(gen.EdCoursesInfoColumns.TeachingType, "=", req.TeachingType)
	}

	if req.UseShop != "" {
		condition1 := gen.Condition{}

		shop := strings.Split(req.UseShop, "|")
		str := fmt.Sprintf("%%|%v|%%", shop[1])
		condition1.And(gen.EdCoursesInfoColumns.UseShop, "like", str)
		str = fmt.Sprintf("%%|%v|%%", shop[2])
		condition1.Or(gen.EdCoursesInfoColumns.UseShop, "like", str)
		where1, out1 := condition1.Get()
		mgr.Where(where1, out1...)
	}
	var list []*model.Course

	where, out := condition.Get()
	mgr.Where(where, out...)
	mgr.Where("is_show = 1")
	mgr.Where("deleted_at IS NULL")
	mgr.Order(SortType(req.Sort))
	err := mgr.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *CourseRepo) ListAll(context.Context) ([]*model.Course, error) {
	var list []*model.Course

	mgr := gen.EdCoursesInfoMgr(r.data.db).Model(model.Course{}).Preload("Sale")
	mgr.Where("is_show = 1")
	mgr.Where("deleted_at IS NULL")
	mgr.Order(SortType(sortASort))
	err := mgr.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// FindSaleByID implements biz.CourseRepo.
func (r *CourseRepo) FindSaleByID(ctx context.Context, id int64) (*model.CourseSale, error) {
	var data model.CourseSale
	err := gen.EdCoursesSaleMgr(r.data.db).Model(model.CourseSale{}).Preload("Course").First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
