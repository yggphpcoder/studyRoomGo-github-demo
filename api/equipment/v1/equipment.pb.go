// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.3
// source: equipment/v1/equipment.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DataReply) Reset() {
	*x = DataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataReply) ProtoMessage() {}

func (x *DataReply) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataReply.ProtoReflect.Descriptor instead.
func (*DataReply) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{1}
}

func (x *DataReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DataReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId int64 `protobuf:"varint,1,opt,name=shopId,proto3" json:"shopId,omitempty"`
	PId    int64 `protobuf:"varint,2,opt,name=pId,proto3" json:"pId,omitempty"`
	Status int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *ListRequest) GetPId() int64 {
	if x != nil {
		return x.PId
	}
	return 0
}

func (x *ListRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SeatAutomateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatId int64 `protobuf:"varint,1,opt,name=seatId,proto3" json:"seatId,omitempty"`
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SeatAutomateRequest) Reset() {
	*x = SeatAutomateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatAutomateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatAutomateRequest) ProtoMessage() {}

func (x *SeatAutomateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatAutomateRequest.ProtoReflect.Descriptor instead.
func (*SeatAutomateRequest) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{3}
}

func (x *SeatAutomateRequest) GetSeatId() int64 {
	if x != nil {
		return x.SeatId
	}
	return 0
}

func (x *SeatAutomateRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DoorAutomateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoorId int64 `protobuf:"varint,1,opt,name=doorId,proto3" json:"doorId,omitempty"`
}

func (x *DoorAutomateRequest) Reset() {
	*x = DoorAutomateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoorAutomateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoorAutomateRequest) ProtoMessage() {}

func (x *DoorAutomateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoorAutomateRequest.ProtoReflect.Descriptor instead.
func (*DoorAutomateRequest) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{4}
}

func (x *DoorAutomateRequest) GetDoorId() int64 {
	if x != nil {
		return x.DoorId
	}
	return 0
}

type ListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*Data `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListReply) Reset() {
	*x = ListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReply) ProtoMessage() {}

func (x *ListReply) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReply.ProtoReflect.Descriptor instead.
func (*ListReply) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{5}
}

func (x *ListReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListReply) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid           int64  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TypeEquipment int32  `protobuf:"varint,4,opt,name=typeEquipment,proto3" json:"typeEquipment,omitempty"`
	TypeRelate    int32  `protobuf:"varint,5,opt,name=typeRelate,proto3" json:"typeRelate,omitempty"`
	RelateID      int64  `protobuf:"varint,6,opt,name=relateID,proto3" json:"relateID,omitempty"`
	Operation     int32  `protobuf:"varint,7,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{6}
}

func (x *Data) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Data) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Data) GetTypeEquipment() int32 {
	if x != nil {
		return x.TypeEquipment
	}
	return 0
}

func (x *Data) GetTypeRelate() int32 {
	if x != nil {
		return x.TypeRelate
	}
	return 0
}

func (x *Data) GetRelateID() int64 {
	if x != nil {
		return x.RelateID
	}
	return 0
}

func (x *Data) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

type LiverCameraListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string             `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*LiverCameraData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *LiverCameraListReply) Reset() {
	*x = LiverCameraListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiverCameraListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiverCameraListReply) ProtoMessage() {}

func (x *LiverCameraListReply) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiverCameraListReply.ProtoReflect.Descriptor instead.
func (*LiverCameraListReply) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{7}
}

func (x *LiverCameraListReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LiverCameraListReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LiverCameraListReply) GetData() []*LiverCameraData {
	if x != nil {
		return x.Data
	}
	return nil
}

type LiverCameraData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	DeviceSerial string `protobuf:"bytes,2,opt,name=deviceSerial,proto3" json:"deviceSerial,omitempty"`
	ChannelNo    string `protobuf:"bytes,3,opt,name=channelNo,proto3" json:"channelNo,omitempty"`
	ChannelName  string `protobuf:"bytes,4,opt,name=channelName,proto3" json:"channelName,omitempty"`
	LiveAddress  string `protobuf:"bytes,5,opt,name=liveAddress,proto3" json:"liveAddress,omitempty"`
}

func (x *LiverCameraData) Reset() {
	*x = LiverCameraData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equipment_v1_equipment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiverCameraData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiverCameraData) ProtoMessage() {}

func (x *LiverCameraData) ProtoReflect() protoreflect.Message {
	mi := &file_equipment_v1_equipment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiverCameraData.ProtoReflect.Descriptor instead.
func (*LiverCameraData) Descriptor() ([]byte, []int) {
	return file_equipment_v1_equipment_proto_rawDescGZIP(), []int{8}
}

func (x *LiverCameraData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LiverCameraData) GetDeviceSerial() string {
	if x != nil {
		return x.DeviceSerial
	}
	return ""
}

func (x *LiverCameraData) GetChannelNo() string {
	if x != nil {
		return x.ChannelNo
	}
	return ""
}

func (x *LiverCameraData) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *LiverCameraData) GetLiveAddress() string {
	if x != nil {
		return x.LiveAddress
	}
	return ""
}

var File_equipment_v1_equipment_proto protoreflect.FileDescriptor

var file_equipment_v1_equipment_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x31, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x4f, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0x0a, 0x13, 0x73, 0x65, 0x61, 0x74,
	0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x2d, 0x0a, 0x13, 0x64, 0x6f, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x6f, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x5d,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbc, 0x01,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x74, 0x79, 0x70, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x14,
	0x4c, 0x69, 0x76, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65,
	0x72, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xad, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x76, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x32, 0xf9, 0x08, 0x0a, 0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x8c, 0x01, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x45, 0x92, 0x41, 0x1c, 0x1a, 0x1a, 0xe6, 0xa0,
	0xb9, 0xe6, 0x8d, 0xae, 0x69, 0x64, 0xe6, 0x8e, 0xa7, 0xe5, 0x88, 0xb6, 0xe9, 0x97, 0xa8, 0xe6,
	0x8e, 0xa7, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01,
	0x2a, 0x22, 0x1b, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x87,
	0x01, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x40, 0x92, 0x41, 0x16, 0x1a, 0x14, 0xe6, 0xa0, 0xb9,
	0xe6, 0x8d, 0xae, 0x69, 0x64, 0xe6, 0x8e, 0xa7, 0xe5, 0x88, 0xb6, 0xe8, 0xae, 0xbe, 0xe5, 0xa4,
	0x87, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x92, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x47, 0x92, 0x41, 0x14, 0x1a, 0x12, 0xe6, 0x8e, 0xa7, 0xe5, 0x88,
	0xb6, 0xe5, 0x85, 0xa8, 0xe9, 0x83, 0xa8, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x2f, 0x7b, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x7d, 0x12, 0x8d, 0x01,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x41, 0x92, 0x41, 0x11, 0x1a,
	0x0f, 0xe6, 0x8e, 0xa7, 0xe5, 0x88, 0xb6, 0xe4, 0xb8, 0xbb, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x7b, 0x70, 0x49, 0x64, 0x7d, 0x12, 0x83, 0x01,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x3f, 0x92, 0x41, 0x22, 0x12, 0x12, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe8,
	0xae, 0xbe, 0xe5, 0xa4, 0x87, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x1a, 0x0c, 0xe8, 0x8e, 0xb7,
	0xe5, 0x8f, 0x96, 0xe5, 0x85, 0xa8, 0xe9, 0x83, 0xa8, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12,
	0x12, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x61, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x37, 0x92, 0x41, 0x34, 0x12, 0x18, 0xe6,
	0xa0, 0xb9, 0xe6, 0x8d, 0xae, 0xe5, 0xba, 0xa7, 0xe4, 0xbd, 0x8d, 0xe6, 0x8e, 0xa7, 0xe5, 0x88,
	0xb6, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0x1a, 0x18, 0xe6, 0xa0, 0xb9, 0xe6, 0x8d, 0xae, 0xe5,
	0xba, 0xa7, 0xe4, 0xbd, 0x8d, 0xe6, 0x8e, 0xa7, 0xe5, 0x88, 0xb6, 0xe8, 0xae, 0xbe, 0xe5, 0xa4,
	0x87, 0x12, 0x79, 0x0a, 0x0c, 0x44, 0x6f, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6f, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x92, 0x41, 0x22, 0x12, 0x0f, 0xe6, 0x8e, 0xa7, 0xe5,
	0x88, 0xb6, 0xe9, 0x97, 0xa8, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0x1a, 0x0f, 0xe6, 0x8e, 0xa7,
	0xe5, 0x88, 0xb6, 0xe9, 0x97, 0xa8, 0xe8, 0xae, 0xbe, 0xe5, 0xa4, 0x87, 0x12, 0x9e, 0x01, 0x0a,
	0x0f, 0x4c, 0x69, 0x76, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x44, 0x92, 0x41, 0x1c, 0x12, 0x0c, 0xe7, 0x9b,
	0x91, 0xe6, 0x8e, 0xa7, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x1a, 0x0c, 0xe7, 0x9b, 0x91, 0xe6,
	0x8e, 0xa7, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d,
	0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x35, 0x0a,
	0x10, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x1f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x47, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_equipment_v1_equipment_proto_rawDescOnce sync.Once
	file_equipment_v1_equipment_proto_rawDescData = file_equipment_v1_equipment_proto_rawDesc
)

func file_equipment_v1_equipment_proto_rawDescGZIP() []byte {
	file_equipment_v1_equipment_proto_rawDescOnce.Do(func() {
		file_equipment_v1_equipment_proto_rawDescData = protoimpl.X.CompressGZIP(file_equipment_v1_equipment_proto_rawDescData)
	})
	return file_equipment_v1_equipment_proto_rawDescData
}

var file_equipment_v1_equipment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_equipment_v1_equipment_proto_goTypes = []interface{}{
	(*GetRequest)(nil),           // 0: api.equipment.v1.GetRequest
	(*DataReply)(nil),            // 1: api.equipment.v1.DataReply
	(*ListRequest)(nil),          // 2: api.equipment.v1.ListRequest
	(*SeatAutomateRequest)(nil),  // 3: api.equipment.v1.seatAutomateRequest
	(*DoorAutomateRequest)(nil),  // 4: api.equipment.v1.doorAutomateRequest
	(*ListReply)(nil),            // 5: api.equipment.v1.ListReply
	(*Data)(nil),                 // 6: api.equipment.v1.Data
	(*LiverCameraListReply)(nil), // 7: api.equipment.v1.LiverCameraListReply
	(*LiverCameraData)(nil),      // 8: api.equipment.v1.LiverCameraData
}
var file_equipment_v1_equipment_proto_depIdxs = []int32{
	6,  // 0: api.equipment.v1.ListReply.data:type_name -> api.equipment.v1.Data
	8,  // 1: api.equipment.v1.LiverCameraListReply.data:type_name -> api.equipment.v1.LiverCameraData
	0,  // 2: api.equipment.v1.Equipment.OpenDoor:input_type -> api.equipment.v1.GetRequest
	0,  // 3: api.equipment.v1.Equipment.Automate:input_type -> api.equipment.v1.GetRequest
	2,  // 4: api.equipment.v1.Equipment.AutomateAll:input_type -> api.equipment.v1.ListRequest
	2,  // 5: api.equipment.v1.Equipment.AutomateMain:input_type -> api.equipment.v1.ListRequest
	2,  // 6: api.equipment.v1.Equipment.List:input_type -> api.equipment.v1.ListRequest
	3,  // 7: api.equipment.v1.Equipment.SeatAutomate:input_type -> api.equipment.v1.seatAutomateRequest
	4,  // 8: api.equipment.v1.Equipment.DoorAutomate:input_type -> api.equipment.v1.doorAutomateRequest
	2,  // 9: api.equipment.v1.Equipment.LiverCameraList:input_type -> api.equipment.v1.ListRequest
	1,  // 10: api.equipment.v1.Equipment.OpenDoor:output_type -> api.equipment.v1.DataReply
	1,  // 11: api.equipment.v1.Equipment.Automate:output_type -> api.equipment.v1.DataReply
	5,  // 12: api.equipment.v1.Equipment.AutomateAll:output_type -> api.equipment.v1.ListReply
	5,  // 13: api.equipment.v1.Equipment.AutomateMain:output_type -> api.equipment.v1.ListReply
	5,  // 14: api.equipment.v1.Equipment.List:output_type -> api.equipment.v1.ListReply
	1,  // 15: api.equipment.v1.Equipment.SeatAutomate:output_type -> api.equipment.v1.DataReply
	1,  // 16: api.equipment.v1.Equipment.DoorAutomate:output_type -> api.equipment.v1.DataReply
	7,  // 17: api.equipment.v1.Equipment.LiverCameraList:output_type -> api.equipment.v1.LiverCameraListReply
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_equipment_v1_equipment_proto_init() }
func file_equipment_v1_equipment_proto_init() {
	if File_equipment_v1_equipment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_equipment_v1_equipment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_equipment_v1_equipment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_equipment_v1_equipment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_equipment_v1_equipment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeatAutomateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_equipment_v1_equipment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoorAutomateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_equipment_v1_equipment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_equipment_v1_equipment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_equipment_v1_equipment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiverCameraListReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_equipment_v1_equipment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiverCameraData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_equipment_v1_equipment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_equipment_v1_equipment_proto_goTypes,
		DependencyIndexes: file_equipment_v1_equipment_proto_depIdxs,
		MessageInfos:      file_equipment_v1_equipment_proto_msgTypes,
	}.Build()
	File_equipment_v1_equipment_proto = out.File
	file_equipment_v1_equipment_proto_rawDesc = nil
	file_equipment_v1_equipment_proto_goTypes = nil
	file_equipment_v1_equipment_proto_depIdxs = nil
}