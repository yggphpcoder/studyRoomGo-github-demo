package service

import (
	pb "studyRoomGo/api/wxofficial/v1"
)

type WxOfficialService struct {
	pb.UnimplementedWxOfficialServer
}

func NewWxOfficialService() *WxOfficialService {
	return &WxOfficialService{}
}
