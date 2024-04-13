package data

import (
	"context"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/data/gen"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type MemberRepo struct {
	data *Data
	log  *log.Helper
}

// NewMemberRepo .
func NewMemberRepo(data *Data, logger log.Logger) biz.MemberRepo {
	return &MemberRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *MemberRepo) Save(ctx context.Context, info *biz.Member) (bool, error) {
	mgr := gen.SrMemberMgr(r.data.db)
	mgr.Save(info.SrMember)
	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *MemberRepo) Update(ctx context.Context, g *biz.Member) (bool, error) {
	mgr := gen.SrMemberMgr(r.data.db)
	if err := mgr.Select("*").Omit(Omit...).Updates(&g.SrMember).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *MemberRepo) FindByID(ctx context.Context, id int64) (*gen.SrMember, error) {
	mgr := gen.SrMemberMgr(r.data.db)
	member, err := mgr.FetchByPrimaryKey(id)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MemberRepo) FindByPhone(ctx context.Context, merchantId int64, phone string) (*gen.SrMember, error) {
	mgr := gen.SrMemberMgr(r.data.db)
	var tmp *gen.SrMember
	mgr.Where("merchant_id = ?", merchantId)
	mgr.Where("phone = ?", phone)
	mgr.Order("id desc")
	mgr.Find(&tmp)
	if tmp.ID == 0 {
		return tmp, mgr.Error
	}
	return tmp, nil
}
func (r *MemberRepo) ListBy(ctx context.Context, s string) ([]*gen.SrMember, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MemberRepo) ListAll(context.Context) ([]*gen.SrMember, error) {
	return nil, nil
}
func (r *MemberRepo) FindByWXOpenId(ctx context.Context, merchantId int64, openId string) (*gen.WxUserinfo, error) {
	mgr := gen.WxUserinfoMgr(r.data.db)
	var tmp *gen.WxUserinfo
	mgr.Where("merchant_id = ?", merchantId)
	mgr.Where("openid = ?", openId).Find(&tmp)
	if tmp.ID == 0 {
		return tmp, mgr.Error
	}
	return tmp, nil
}
func (r *MemberRepo) FindByWXUnionId(ctx context.Context, merchantId int64, unionId string) (*gen.WxUserinfo, error) {
	mgr := gen.WxUserinfoMgr(r.data.db)
	var tmp *gen.WxUserinfo
	mgr.Where("merchant_id = ?", merchantId)
	mgr.Where("union_id = ?", unionId).Find(&tmp)
	if tmp.ID == 0 {
		return tmp, mgr.Error
	}
	return tmp, nil
}
func (r *MemberRepo) FindByWXPhone(ctx context.Context, merchantId int64, phone string) (*gen.WxUserinfo, error) {
	mgr := gen.WxUserinfoMgr(r.data.db)
	var tmp *gen.WxUserinfo
	mgr.Where("merchant_id = ?", merchantId)
	mgr.Where("phone = ?", phone)
	mgr.Order(SortType(sortDId))
	mgr.Find(&tmp)
	if tmp.ID == 0 {
		return tmp, mgr.Error
	}
	return tmp, nil
}

func (r *MemberRepo) CreateWXUserInfo(ctx context.Context, info *biz.WxUserInfo) (bool, error) {
	mgr := gen.WxUserinfoMgr(r.data.db)
	mgr.Save(info.WxUserinfo)
	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func (r *MemberRepo) UpdateWXUserInfo(ctx context.Context, info *biz.WxUserInfo) (bool, error) {
	mgr := gen.WxUserinfoMgr(r.data.db)
	if err := mgr.Select("*").Omit(Omit...).Updates(&info.WxUserinfo).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *MemberRepo) HasNotice(ctx context.Context, memberId int64, isRead *int) (bool, error) {
	var tmp *gen.SrMemberNotice
	mgr := gen.SrMemberNoticeMgr(r.data.db)
	mgr.Where("member_id = ?", memberId)
	if isRead != nil {
		mgr.Where("is_read = ?", isRead)
	}
	mgr.Last(&tmp)

	if mgr.Error != nil {
		return false, mgr.Error
	}

	return true, nil
}

func (r *MemberRepo) LastNotice(ctx context.Context, memberId int64, isRead *int, limit int32) ([]*gen.SrMemberNotice, error) {
	var tmp []*gen.SrMemberNotice
	mgr := gen.SrMemberNoticeMgr(r.data.db)
	mgr.Where("member_id = ?", memberId)
	if isRead != nil {
		mgr.Where("is_read = ?", isRead)
	}
	mgr.Order("id desc")
	mgr.Limit(int(limit)).Find(&tmp)
	if mgr.Error != nil {
		return nil, mgr.Error
	}
	return tmp, nil
}

func (r *MemberRepo) ListNotice(ctx context.Context, memberId int64, Page int32, PageSize int32) ([]*gen.SrMemberNotice, error) {
	mgr := gen.SrMemberNoticeMgr(r.data.db)
	mgr.Where("member_id = ?", memberId)
	if PageSize > 0 && Page > 0 {
		mgr.Limit(int(PageSize)).Offset(int((Page - 1) * PageSize))
	}
	var modelList []*gen.SrMemberNotice
	err := mgr.Order("is_read asc,id desc").Scan(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil
}

func (r *MemberRepo) UpdateNotice(ctx context.Context, id []int64, isRead int) (bool, error) {
	mgr := gen.SrMemberNoticeMgr(r.data.db)
	err := mgr.Where("id in ?", id).Update("is_read", isRead).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *MemberRepo) SaveWXSubscribe(ctx context.Context, info *gen.WxNoticeSubscribe) (bool, error) {
	mgr := gen.WxNoticeSubscribeMgr(r.data.db)
	mgr.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "member_id"}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_subscribe"}),
	})
	mgr.Create(info)
	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}
