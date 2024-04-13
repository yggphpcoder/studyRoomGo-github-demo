package data

import (
	"context"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"

	"studyRoomGo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ConfigRepo struct {
	data *Data
	log  *log.Helper
}

// NewConfigRepo .
func NewConfigRepo(data *Data, logger log.Logger) biz.ConfigRepo {
	return &ConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ConfigRepo) FindDictByID(ctx context.Context, id int64) (*gen.SysDictData, error) {
	mgr := gen.SysDictDataMgr(r.data.db)
	data, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ConfigRepo) FindDictBy(ctx context.Context, key string, value string) (*gen.SysDictData, error) {
	condition := gen.Condition{}
	condition.And(key, "=", value)
	where, out := condition.Get()

	mgr := gen.SysDictDataMgr(r.data.db.Where(where, out...))
	data, err := mgr.Get()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ConfigRepo) ListDictBy(ctx context.Context, req *biz.SearchDictRequest) ([]*gen.SysDictData, error) {
	condition := gen.Condition{}
	if req.Type != "" {
		condition.And(gen.SysDictDataColumns.DictType, "=", req.Type)
	}
	where, out := condition.Get()

	mgr := gen.SysDictDataMgr(r.data.db.Where(where, out...).Order(SortType(sortDictSort)))
	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *ConfigRepo) ListDictAll(context.Context) ([]*gen.SysDictData, error) {
	mgr := gen.SysDictDataMgr(r.data.db)
	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *ConfigRepo) UpdateByKey(ctx context.Context, shopId int64, key string, value string) (bool, error) {
	condition := gen.Condition{}
	condition.And(gen.SrSettingColumns.ShopID, "=", shopId)
	condition.And(gen.SrSettingColumns.Key, "=", key)

	where, out := condition.Get()

	mgr := gen.SrSettingMgr(r.data.db.Where(where, out...))
	if err := mgr.Select("value").Omit(Omit...).Update(gen.SrSettingColumns.Value, value).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *ConfigRepo) FindSettingBy(ctx context.Context, shopId int64, key string) (*model.Setting, error) {
	var setting *model.Setting
	condition := gen.Condition{}
	condition.And(gen.SrSettingColumns.ShopID, "=", shopId)
	condition.And(gen.SrSettingColumns.Key, "=", key)
	where, out := condition.Get()

	mgr := gen.SrSettingMgr(r.data.db)
	mgr.Joins("left join sr_setting_text as s on s.setting_id = sr_setting.id")
	mgr.Where(where, out...)
	err := mgr.Find(&setting).Error
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *ConfigRepo) ListSettingBy(ctx context.Context, req *biz.SearchSettingRequest) ([]*model.Setting, error) {
	var setting []*model.Setting
	condition := gen.Condition{}
	if req.MerchantId != 0 {
		condition.And(gen.SrSettingColumns.MerchantID, "=", req.MerchantId)
	}
	if req.ShopId != 0 {
		condition.And(gen.SrSettingColumns.ShopID, "=", req.ShopId)
	}
	if req.Group != "" {
		condition.And(gen.SrSettingColumns.Group, "=", req.Group)
	}
	if req.Key != "" {
		condition.And(gen.SrSettingColumns.Key, "=", req.Key)
	}
	where, out := condition.Get()

	mgr := gen.SrSettingMgr(r.data.db)
	mgr.Where(where, out...)
	mgr.Joins("left join sr_setting_text as s on s.setting_id = sr_setting.id")
	err := mgr.Scan(&setting).Error
	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (r *ConfigRepo) ListEdDictBy(ctx context.Context, req *biz.SearchDictRequest) ([]*gen.EdDictData, error) {
	var setting []*gen.EdDictData

	subQuery1 := gen.EdDictDataMgr(r.data.db).
		Select("dict_value", "dict_type", "MAX(`merchant_id`) as merchant_id", "MAX(`shop_id`) as shop_id").
		Group("dict_type,dict_value").
		Where("merchant_id IN (?) and shop_id in (?) ", []int64{req.MerchantId, 0}, []int64{req.ShopId, 0})
	mgr := gen.EdDictDataMgr(r.data.db).
		Where("status = 2").
		Order("dict_sort asc").
		Joins("inner join (?) as b on ed_dict_data.`dict_type` = b.`dict_type` and ed_dict_data.`dict_value` = b.`dict_value`", subQuery1).
		Where("ed_dict_data.merchant_id = b.merchant_id and ed_dict_data.shop_id = b.shop_id").
		Where("ed_dict_data.dict_type = ?", req.Type).Order(SortType(sortDictASort))
	err := mgr.Scan(&setting).Error
	if err != nil {
		return nil, err
	}

	return setting, nil
}
