package biz

import (
	"context"
	"errors"
	"studyRoomGo/common/equipment"
	"studyRoomGo/internal/data/gen"

	"github.com/go-kratos/kratos/v2/log"
)

// EquipmentUsecase is an Equipment usecase.
type EquipmentUsecase struct {
	repo EquipmentRepo
	log  *log.Helper
}

// NewEquipmentUsecase new an Equipment usecase.
func NewEquipmentUsecase(repo EquipmentRepo, logger log.Logger) *EquipmentUsecase {
	return &EquipmentUsecase{repo: repo, log: log.NewHelper(logger)}
}

type EquipmentRequest struct {
	ShoId         int64
	TypeRelate    int64
	TypeEquipment int64
	RelateId      int64
	Pid           *int64
}

func (uc *EquipmentUsecase) FindByRelateId(ctx context.Context, req *EquipmentRequest) (*gen.SrEquipment, error) {
	data, err := uc.repo.FindByRelateId(ctx, req)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *EquipmentUsecase) ListBy(ctx context.Context, req *EquipmentRequest) ([]*gen.SrEquipment, error) {
	data, err := uc.repo.ListBy(ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*gen.SrEquipment
	for _, item := range data {
		list = append(list, item)
	}
	return list, nil
}

func (uc *EquipmentUsecase) OpenDoor(ctx context.Context, id int64) (ok bool, err error) {

	data, _ := uc.repo.FindByID(ctx, id)
	if data.TypeEquipment != 2 {
		return false, errors.New("此设备不是门控，请找管理员")
	}
	if data.ShopID == 2 {
		if ok, _ := equipment.Servers["accesscontrol"].Automate(data.Mac, data.Code, 1); !ok {
			uc.log.Errorf("accesscontrol推送失败,%v_control,code:%s,status:%d", data.Mac, data.Code, 1)

			return false, errors.New("推送失败，请找管理员")
		}
	} else {
		if ok, _ := equipment.Servers["smyoo"].Automate(data.Mac, data.Code, 1); !ok {
			uc.log.Errorf("smyoo推送失败,%v_control,code:%s,status:%d", data.Mac, data.Code, 1)

			return false, errors.New("推送失败，请找管理员")
		}
	}

	return true, nil
}

// Automate 操作单个设备
func (uc *EquipmentUsecase) Automate(ctx context.Context, id int64, status int8) (ok bool, err error) {

	data, _ := uc.repo.FindByID(ctx, id)
	if data.ShopID == 2 {
		if ok, _ := equipment.Servers["mqtt"].Automate(data.Mac, data.Code, int(status)); !ok {
			return false, errors.New("推送失败，请找管理员")
		}
	} else {
		if ok, _ := equipment.Servers["smyoo"].Automate(data.Mac, data.Code, int(status)); !ok {
			return false, errors.New("推送失败，请找管理员")
		}
	}

	data.Operation = status
	ok, err = uc.repo.Save(ctx, data)
	if !ok {
		return false, err
	}
	return true, nil
}

// AutomateMain 操作主设备
func (uc *EquipmentUsecase) AutomateMain(ctx context.Context, pid int64, status int8) (ok bool, err error) {

	data, err := uc.repo.ListBy(ctx, &EquipmentRequest{Pid: &pid})
	if len(data) == 0 {
		return false, errors.New("操作失败，请找管理员")
	}
	ok, err = uc.repo.UpdatesOperation(ctx, &EquipmentRequest{Pid: &pid}, status)
	if !ok {
		return false, err
	}
	var code []string
	for _, item := range data {
		code = append(code, item.Code)
	}
	if data[0].ShopID == 2 {
		if ok, _ := equipment.Servers["mqtt"].AutomateMain(data[0].Mac, code, int(status)); !ok {
			return false, errors.New("推送失败，请找管理员")
		}
	} else {
		if ok, _ := equipment.Servers["smyoo"].AutomateMain(data[0].Mac, code, int(status)); !ok {
			return false, errors.New("推送失败，请找管理员")
		}
	}

	return true, nil
}

// AutomateAll 操作店铺所有设备
func (uc *EquipmentUsecase) AutomateAll(ctx context.Context, shopId int64, status int8) (ok bool, err error) {
	data, err := uc.repo.ListBy(ctx, &EquipmentRequest{ShoId: shopId})
	if len(data) == 0 {
		return false, err
	}

	ok, err = uc.repo.UpdatesOperation(ctx, &EquipmentRequest{ShoId: shopId}, status)
	if !ok {
		return false, err
	}

	macList := make(map[string][]string)
	for _, item := range data {
		macList[item.Mac] = append(macList[item.Mac], item.Code)
	}

	for mac, code := range macList {
		if data[0].ShopID == 2 {
			if ok, _ := equipment.Servers["mqtt"].AutomateMain(mac, code, int(status)); !ok {
				return false, errors.New("推送失败，请找管理员")
			}
		} else {
			if ok, _ := equipment.Servers["smyoo"].AutomateMain(mac, code, int(status)); !ok {
				return false, errors.New("推送失败，请找管理员")
			}
		}

	}

	return true, nil
}
