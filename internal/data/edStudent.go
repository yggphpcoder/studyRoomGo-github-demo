package data

import (
	"context"
	"strconv"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
)

type StudentRepo struct {
	data *Data
	log  *log.Helper
}

// NewStudentRepo .
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &StudentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *StudentRepo) FindByID(ctx context.Context, id int64) (*model.MyStudent, error) {
	mgr := gen.EdStudentRelateMemberMgr(r.data.db).Model(&model.MyStudent{}).Preload("Student")
	mgr.Where("id = ?", id)
	var model *model.MyStudent
	err := mgr.First(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *StudentRepo) ListBy(ctx context.Context, req *biz.SearchStudentRequest) ([]*model.MyStudent, error) {
	mgr := gen.EdStudentRelateMemberMgr(r.data.db).Model(&model.MyStudent{}).Preload("Student")
	mgr.Where("member_id = ?", req.MemberId)
	var modelList []*model.MyStudent
	err := mgr.Find(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil

}

// Save implements biz.StudentRepo.
func (r *StudentRepo) Save(ctx context.Context, data *model.MyStudent) (bool, error) {
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	if merchantId != 0 {
		data.Student.MerchantID = int64(merchantId)
	}
	if shopId != 0 {
		data.Student.ShopID = int64(shopId)
	}
	mgr := gen.EdStudentRelateMemberMgr(r.data.db).Model(&model.MyStudent{})
	mgr.Save(&data)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

// Update implements biz.StudentRepo.
func (r *StudentRepo) Update(ctx context.Context, data *model.MyStudent) (bool, error) {
	mgr := gen.EdStudentInfoMgr(r.data.db)
	mgr.Updates(&data.Student)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}
