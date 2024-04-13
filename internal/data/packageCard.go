package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"
	"strconv"
	"strings"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type PackageCardRepo struct {
	data *Data
	log  *log.Helper
}

// NewPackageCardRepo .
func NewPackageCardRepo(data *Data, logger log.Logger) biz.PackageCardRepo {
	return &PackageCardRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *PackageCardRepo) FindByID(ctx context.Context, id int64) (*gen.SrPackageCard, error) {
	mgr := gen.SrPackageCardMgr(r.data.db)
	data, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *PackageCardRepo) FindBy(ctx context.Context, key string, value string) (*gen.SrPackageCard, error) {
	condition := gen.Condition{}
	condition.And(key, "=", value)

	where, out := condition.Get()

	mgr := gen.SrPackageCardMgr(r.data.db.Where(where, out...))
	mgr.Where("deleted_at IS NULL")

	data, err := mgr.Get()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *PackageCardRepo) ListBy(ctx context.Context, req *biz.SearchPackageCardRequest) ([]*gen.SrPackageCard, error) {
	mgr := gen.SrPackageCardMgr(r.data.db.Order(SortType(req.Sort)))

	condition := gen.Condition{}
	if len(req.PId) > 0 {
		condition.And(gen.SrPackageCardColumns.ID, "in", req.PId)
	}
	if req.MerchantId != 0 {
		condition.And(gen.SrPackageCardColumns.MerchantID, "=", req.MerchantId)
	}
	if req.Status != 0 {
		condition.And(gen.SrPackageCardColumns.Status, "=", req.Status)
	}
	if req.TypePackage != 0 {
		condition.And(gen.SrPackageCardColumns.TypePackage, "=", req.TypePackage)
	}
	if req.TypeCard != 0 {
		condition.And(gen.SrPackageCardColumns.TypeCard, "=", req.TypeCard)
	}

	if req.UseShop != "" {
		condition1 := gen.Condition{}

		shop := strings.Split(req.UseShop, "|")
		str := fmt.Sprintf("%%|%v|%%", shop[1])
		condition1.And(gen.SrPackageCardColumns.UseShop, "like", str)
		str = fmt.Sprintf("%%|%v|%%", shop[2])
		condition1.Or(gen.SrPackageCardColumns.UseShop, "like", str)
		where1, out1 := condition1.Get()
		mgr.Where(where1, out1...)
	}
	if req.UseSeat != "" {
		str := fmt.Sprintf("%%\"%v\"%%", req.UseSeat)
		condition.And(gen.SrPackageCardColumns.UseSeat, "like", str)
	}
	if req.TypeSale != 0 {
		condition.And(gen.SrPackageCardColumns.TypeSale, "=", req.TypeSale)
	} else {
		condition.And(gen.SrPackageCardColumns.TypeSale, "=", 1)
	}
	where, out := condition.Get()
	mgr.Where(where, out...)
	mgr.Where("deleted_at IS NULL")
	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *PackageCardRepo) ListAll(context.Context) ([]*gen.SrPackageCard, error) {
	condition := gen.Condition{}
	condition.And(gen.SrPackageCardColumns.TypeSale, "=", 1) //线上销售
	where, out := condition.Get()
	mgr := gen.SrPackageCardMgr(r.data.db.Where(where, out...).Order(SortType(sortASort)))
	mgr.Where("deleted_at IS NULL")

	data, err := mgr.Gets()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *PackageCardRepo) RecordBuyLog(ctx context.Context, data *biz.PackageCardMemberBuyLog) (bool, error) {
	merchantId := 0
	shopId := 0
	if md, ok := metadata.FromServerContext(ctx); ok {
		merchantId, _ = strconv.Atoi(md.Get("x-md-global-merchantId"))
	}
	if md, ok := metadata.FromServerContext(ctx); ok {
		shopId, _ = strconv.Atoi(md.Get("x-md-global-shopId"))
	}
	if merchantId != 0 {
		data.MerchantID = int64(merchantId)
	}
	if shopId != 0 {
		data.ShopID = int64(shopId)
	}
	mgr := gen.SrPackageCardMemberBuyLogMgr(r.data.db.Order(SortType(sortCreatAtSort)))
	mgr.Save(data)

	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *PackageCardRepo) FindBuyLogBy(ctx context.Context, key string, value string) (*gen.SrPackageCardMemberBuyLog, error) {
	condition := gen.Condition{}
	condition.And(key, "=", value)

	where, out := condition.Get()

	mgr := gen.SrPackageCardMemberBuyLogMgr(r.data.db.Where(where, out...))
	mgr.Where("deleted_at IS NULL")

	data, err := mgr.Get()
	if err != nil {
		return nil, err
	}
	return &data, nil
}
