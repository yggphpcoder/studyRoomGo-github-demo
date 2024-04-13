package dto

import (
	pb "studyRoomGo/api/config/v1"
	"studyRoomGo/internal/biz"
)

type Dict struct {
	Config *biz.Config
}

func (r *Dict) DictReply() *pb.DataDict {
	return &pb.DataDict{
		Id:        r.Config.Dict.DictCode,
		DictSort:  *r.Config.Dict.DictSort,
		DictLabel: *r.Config.Dict.DictLabel,
		DictValue: *r.Config.Dict.DictValue,
		DictType:  *r.Config.Dict.DictType,
	}
}

func (r *Dict) EdDictReply() *pb.DataDict {
	return &pb.DataDict{
		Id:        r.Config.EdDict.DictCode,
		DictSort:  *r.Config.EdDict.DictSort,
		DictLabel: *r.Config.EdDict.DictLabel,
		DictValue: *r.Config.EdDict.DictValue,
		DictType:  *r.Config.EdDict.DictType,
	}
}

type Setting struct {
	Config *biz.Config
}

func (r *Setting) SettingReply() *pb.DataSetting {
	re := &pb.DataSetting{
		Id:        r.Config.Setting.SrSetting.ID,
		Name:      r.Config.Setting.Name,
		Key:       r.Config.Setting.Key,
		Value:     r.Config.Setting.Value,
		ValueType: int32(r.Config.Setting.ValueType),
	}
	if r.Config.Setting.ValueType == 2 || r.Config.Setting.ValueType == 3 {
		re.Value = r.Config.Setting.RichText
	}
	return re
}
