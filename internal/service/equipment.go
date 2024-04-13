package service

import (
	"context"
	"fmt"
	"strconv"
	pb "studyRoomGo/api/equipment/v1"
	"studyRoomGo/common/jwt"
	"studyRoomGo/common/ys7"
	"studyRoomGo/internal/biz"
	"time"
)

type EquipmentService struct {
	pb.UnimplementedEquipmentServer

	uc *biz.EquipmentUsecase
}

func NewEquipmentService(uc *biz.EquipmentUsecase) *EquipmentService {
	return &EquipmentService{uc: uc}
}

func (s *EquipmentService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	if ok, err := jwt.CheckIsAdmin(ctx); !ok {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	data, err := s.uc.ListBy(ctx, &biz.EquipmentRequest{
		ShoId: req.ShopId,
	})
	if err != nil {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	var list []*pb.Data
	for _, item := range data {
		list = append(list, &pb.Data{
			Id:            item.ID,
			Pid:           item.Pid,
			Name:          item.Name,
			TypeEquipment: int32(item.TypeEquipment),
			TypeRelate:    int32(item.TypeRelate),
			RelateID:      item.RelateID,
			Operation:     int32(item.Operation),
		})
	}
	return &pb.ListReply{
		Code: 200,
		Data: list,
	}, nil

}

func (s *EquipmentService) OpenDoor(ctx context.Context, req *pb.GetRequest) (*pb.DataReply, error) {
	if ok, err := jwt.CheckIsAdmin(ctx); !ok {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	ok, err := s.uc.OpenDoor(ctx, req.Id)
	if !ok {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.DataReply{
		Code: 200,
	}, nil
}

func (s *EquipmentService) Automate(ctx context.Context, req *pb.GetRequest) (*pb.DataReply, error) {
	if ok, err := jwt.CheckIsAdmin(ctx); !ok {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	ok, err := s.uc.Automate(ctx, req.Id, int8(req.Status))
	if !ok {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.DataReply{
		Code: 200,
	}, nil
}

func (s *EquipmentService) AutomateMain(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	if ok, err := jwt.CheckIsAdmin(ctx); !ok {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	ok, err := s.uc.AutomateMain(ctx, req.PId, int8(req.Status))
	if !ok {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.ListReply{
		Code: 200,
	}, nil
}

func (s *EquipmentService) AutomateAll(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	if ok, err := jwt.CheckIsAdmin(ctx); !ok {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	ok, err := s.uc.AutomateAll(ctx, req.ShopId, int8(req.Status))
	if !ok {
		return &pb.ListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	return &pb.ListReply{
		Code: 200,
	}, nil
}

func (s *EquipmentService) LiverCameraList(ctx context.Context, req *pb.ListRequest) (*pb.LiverCameraListReply, error) {
	if ok, err := jwt.CheckIsAdmin(ctx); !ok {
		return &pb.LiverCameraListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	equipmentList, err := s.uc.ListBy(ctx, &biz.EquipmentRequest{
		ShoId:         req.ShopId,
		TypeEquipment: 3,
	})
	if err != nil {
		return &pb.LiverCameraListReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	auth := ys7.NewAuth()
	token, _ := auth.GetAccessToken()
	ys7Device := ys7.Device{}
	liveAddress := ys7Device.LiveAddressList()
	var liveA = make(map[string]string)
	for _, datum := range liveAddress.Data {
		liveA[datum.DeviceSerial+strconv.Itoa(datum.ChannelNo)] = datum.LiveAddress
	}

	reply := pb.LiverCameraListReply{
		Code: 200,
		Msg:  "success",
	}
	for _, equipment := range equipmentList {
		if liveA[equipment.Mac+equipment.Code] != "" {
			data := &pb.LiverCameraData{
				Token:        token,
				DeviceSerial: equipment.Mac,
				ChannelNo:    equipment.Code,
				ChannelName:  equipment.Name,
				LiveAddress:  liveA[equipment.Mac+equipment.Code],
			}
			reply.Data = append(reply.Data, data)
		}
	}

	return &reply, nil
}

func (s *EquipmentService) SeatAutomate(ctx context.Context, req *pb.SeatAutomateRequest) (*pb.DataReply, error) {
	fmt.Printf("time %v,EquipmentService_SeatAutomate req %d status %v \n", time.Now().Unix(), req.SeatId, req.Status)

	data, err := s.uc.FindByRelateId(ctx, &biz.EquipmentRequest{
		TypeRelate: 1,
		RelateId:   req.SeatId,
	})

	if err != nil {
		fmt.Printf("EquipmentService_SeatAutomate error %d status %v", req.SeatId, req.Status)

		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}

	if req.Status == 1 {
		_, err = s.uc.Automate(ctx, data.ID, 1)
	} else {
		_, err = s.uc.Automate(ctx, data.ID, 0)

	}
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.DataReply{
		Code: 200,
	}, nil
}

func (s *EquipmentService) DoorAutomate(ctx context.Context, req *pb.DoorAutomateRequest) (*pb.DataReply, error) {
	_, err := s.uc.OpenDoor(ctx, req.DoorId)
	if err != nil {
		return &pb.DataReply{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &pb.DataReply{
		Code: 200,
	}, nil
}
