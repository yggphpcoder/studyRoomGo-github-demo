package data

import (
	"context"
	"studyRoomGo/internal/data/gen"

	"studyRoomGo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type EquipmentRepo struct {
	data *Data
	log  *log.Helper
}

// NewEquipmentRepo .
func NewEquipmentRepo(data *Data, logger log.Logger) biz.EquipmentRepo {
	return &EquipmentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *EquipmentRepo) Save(ctx context.Context, info *gen.SrEquipment) (bool, error) {
	mgr := gen.SrEquipmentMgr(r.data.db)
	mgr.Save(info)
	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *EquipmentRepo) UpdatesOperation(ctx context.Context, req *biz.EquipmentRequest, operation int8) (bool, error) {
	condition := gen.Condition{}

	if req.ShoId != 0 {
		condition.And(gen.SrEquipmentColumns.ShopID, "=", req.ShoId)
	}
	if req.Pid != nil {
		condition.And(gen.SrEquipmentColumns.Pid, "=", &req.Pid)
	}
	if req.TypeRelate != 0 {
		condition.And(gen.SrEquipmentColumns.TypeRelate, "=", req.TypeRelate)
	}
	if req.TypeEquipment != 0 {
		condition.And(gen.SrEquipmentColumns.TypeEquipment, "=", req.TypeEquipment)
	}
	where, out := condition.Get()

	mgr := gen.SrEquipmentMgr(r.data.db.Where(where, out...))
	mgr.Select("operation").Updates(&gen.SrEquipment{Operation: operation})
	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *EquipmentRepo) FindByID(ctx context.Context, id int64) (*gen.SrEquipment, error) {

	condition := gen.Condition{}
	condition.And(gen.SrEquipmentColumns.Status, "=", 2)
	condition.And(gen.SrEquipmentColumns.ID, "=", id)
	where, out := condition.Get()
	mgr := gen.SrEquipmentMgr(r.data.db.Where(where, out...))
	data, err := mgr.Get()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *EquipmentRepo) FindByRelateId(ctx context.Context, req *biz.EquipmentRequest) (*gen.SrEquipment, error) {
	condition := gen.Condition{}

	if req.TypeRelate != 0 {
		condition.And(gen.SrEquipmentColumns.TypeRelate, "=", req.TypeRelate)
	}
	if req.TypeEquipment != 0 {
		condition.And(gen.SrEquipmentColumns.TypeEquipment, "=", req.TypeEquipment)
	}
	condition.And(gen.SrEquipmentColumns.RelateID, "=", req.RelateId)
	condition.And(gen.SrEquipmentColumns.Status, "=", 2)

	where, out := condition.Get()

	mgr := gen.SrEquipmentMgr(r.data.db.Where(where, out...))
	data, err := mgr.Get()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *EquipmentRepo) ListBy(ctx context.Context, req *biz.EquipmentRequest) ([]*gen.SrEquipment, error) {
	condition := gen.Condition{}
	if req.ShoId != 0 {
		condition.And(gen.SrEquipmentColumns.ShopID, "=", req.ShoId)
	}
	if req.Pid != nil {
		condition.And(gen.SrEquipmentColumns.Pid, "=", &req.Pid)
	}
	if req.TypeRelate != 0 {
		condition.And(gen.SrEquipmentColumns.TypeRelate, "=", req.TypeRelate)
	}
	if req.TypeEquipment != 0 {
		condition.And(gen.SrEquipmentColumns.TypeEquipment, "=", req.TypeEquipment)
	}
	condition.And(gen.SrEquipmentColumns.Status, "=", 2)
	where, out := condition.Get()

	mgr := gen.SrEquipmentMgr(r.data.db.Where(where, out...).Order(SortType(sortASort)))
	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}
	return data, nil
}
