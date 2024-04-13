package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"
	"studyRoomGo/internal/data/model"
)

type ShopRepo struct {
	data *Data
	log  *log.Helper
}

// NewShopRepo .
func NewShopRepo(data *Data, logger log.Logger) biz.ShopRepo {
	return &ShopRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ShopRepo) FindByID(ctx context.Context, id int64) (*gen.SrShop, error) {
	mgr := gen.SrShopMgr(r.data.db)
	data, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ShopRepo) FindBy(ctx context.Context, key string, value string) (*gen.SrShop, error) {
	condition := gen.Condition{}
	condition.And(key, "=", value)
	where, out := condition.Get()

	mgr := gen.SrShopMgr(r.data.db.Where(where, out...))
	data, err := mgr.Get()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ShopRepo) ListBy(ctx context.Context, req *biz.SearchShopRequest) ([]*gen.SrShop, error) {
	condition := gen.Condition{}
	if req.MerchantId != 0 {
		condition.And(gen.SrShopColumns.MerchantID, "=", req.MerchantId)
	}
	if req.Name != "" {
		str := fmt.Sprintf("%%%v%%", req.Name)
		condition.And(gen.SrShopColumns.Name, "like", str)
	}
	if req.Province != "" {
		str := fmt.Sprintf("%%%v%%", req.Province)
		condition.And(gen.SrShopColumns.Province, "like", str)
	}
	if req.City != "" {
		str := fmt.Sprintf("%%%v%%", req.City)
		condition.And(gen.SrShopColumns.City, "like", str)
	}
	if req.District != "" {
		str := fmt.Sprintf("%%%v%%", req.District)
		condition.And(gen.SrShopColumns.District, "like", str)
	}
	if req.Address != "" {
		str := fmt.Sprintf("%%%v%%", req.Address)
		condition.And(gen.SrShopColumns.Address, "like", str)
	}
	if req.TypeShop != 0 {
		condition.And(gen.SrShopColumns.TypeShop, "=", req.TypeShop)
	}
	if req.Status != 0 {
		condition.And(gen.SrShopColumns.Status, "=", req.Status)
	}
	where, out := condition.Get()

	mgr := gen.SrShopMgr(r.data.db.Where(where, out...).Order(SortType(req.Sort)))
	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ShopRepo) ListAll(context.Context) ([]*gen.SrShop, error) {
	mgr := gen.SrShopMgr(r.data.db)
	mgr.Order(SortType(sortASort))
	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *ShopRepo) CanUseSeat(ctx context.Context, shopId int64) (list []*biz.AreaSeatCount, total int32, using int32, err error) {
	mgr := gen.SrSeatMgr(r.data.db)

	mgr.Select("shop_id,shop_area_id, count(*) as total").Group("shop_id,shop_area_id")
	mgr.Where("shop_id = ?", shopId)
	mgr.Where("status in ?", []int{2, 3})
	mgr.Find(&list)

	for _, seatCount := range list {
		total += seatCount.Total
	}
	usingCount := int64(0)
	seatBookMgr := gen.SrSeatBookMgr(r.data.db)
	seatBookMgr.Where("shop_id = ?", shopId)
	seatBookMgr.Where("status in ?", []int32{biz.SeatBook_StatusNoSeat, biz.SeatBook_StatusSeated})
	seatBookMgr.Count(&usingCount)
	return list, total, int32(usingCount), nil
}

func (r *ShopRepo) ListSetting(ctx context.Context, shopId int64) (map[int64][]*model.Setting, error) {
	var setting []*model.Setting

	mgr := gen.SrSettingMgr(r.data.db)

	if shopId != 0 {
		condition := gen.Condition{}
		condition.And(gen.SrSettingColumns.ShopID, "=", shopId)
		where, out := condition.Get()
		mgr.Where(where, out...)
	}
	mgr.Joins("left join sr_setting_text as s on s.setting_id = sr_setting.id")
	err := mgr.Scan(&setting).Error
	if err != nil {
		return nil, err
	}
	var list = make(map[int64][]*model.Setting)

	for _, d := range setting {
		list[d.ShopID] = append(list[d.ShopID], d)
	}
	return list, nil
}

func (r *ShopRepo) FindShopAreaReception(ctx context.Context, shopId int64) (*gen.SrShopArea, error) {
	condition := gen.Condition{}
	condition.And(gen.SrShopAreaColumns.ShopID, "=", shopId)
	condition.And(gen.SrShopAreaColumns.TypeArea, "=", 2)
	where, out := condition.Get()

	mgr := gen.SrShopAreaMgr(r.data.db)
	mgr.Where(where, out...)
	data, err := mgr.Get()
	if err != nil {
		return nil, err
	}

	return &data, nil
}
