// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: parkplatz.proto

package parkplatz

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaId        string `protobuf:"bytes,1,opt,name=areaId,proto3" json:"areaId,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	StartDateTime int32  `protobuf:"varint,3,opt,name=startDateTime,proto3" json:"startDateTime,omitempty"`
	EndDateTime   int32  `protobuf:"varint,4,opt,name=endDateTime,proto3" json:"endDateTime,omitempty"`
}

func (x *ReservationRequest) Reset() {
	*x = ReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkplatz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationRequest) ProtoMessage() {}

func (x *ReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkplatz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationRequest.ProtoReflect.Descriptor instead.
func (*ReservationRequest) Descriptor() ([]byte, []int) {
	return file_parkplatz_proto_rawDescGZIP(), []int{0}
}

func (x *ReservationRequest) GetAreaId() string {
	if x != nil {
		return x.AreaId
	}
	return ""
}

func (x *ReservationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReservationRequest) GetStartDateTime() int32 {
	if x != nil {
		return x.StartDateTime
	}
	return 0
}

func (x *ReservationRequest) GetEndDateTime() int32 {
	if x != nil {
		return x.EndDateTime
	}
	return 0
}

type ReservationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationId string `protobuf:"bytes,1,opt,name=reservationId,proto3" json:"reservationId,omitempty"`
	Status        string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReservationDetails) Reset() {
	*x = ReservationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkplatz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationDetails) ProtoMessage() {}

func (x *ReservationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_parkplatz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationDetails.ProtoReflect.Descriptor instead.
func (*ReservationDetails) Descriptor() ([]byte, []int) {
	return file_parkplatz_proto_rawDescGZIP(), []int{1}
}

func (x *ReservationDetails) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

func (x *ReservationDetails) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type TerminationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationId string `protobuf:"bytes,1,opt,name=reservationId,proto3" json:"reservationId,omitempty"`
}

func (x *TerminationRequest) Reset() {
	*x = TerminationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkplatz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminationRequest) ProtoMessage() {}

func (x *TerminationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkplatz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminationRequest.ProtoReflect.Descriptor instead.
func (*TerminationRequest) Descriptor() ([]byte, []int) {
	return file_parkplatz_proto_rawDescGZIP(), []int{2}
}

func (x *TerminationRequest) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

type AreaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=displayName,proto3" json:"displayName,omitempty"`
}

func (x *AreaRequest) Reset() {
	*x = AreaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkplatz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaRequest) ProtoMessage() {}

func (x *AreaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkplatz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaRequest.ProtoReflect.Descriptor instead.
func (*AreaRequest) Descriptor() ([]byte, []int) {
	return file_parkplatz_proto_rawDescGZIP(), []int{3}
}

func (x *AreaRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type AreaDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaId     string `protobuf:"bytes,1,opt,name=areaId,proto3" json:"areaId,omitempty"`
	TotalSpots int32  `protobuf:"varint,2,opt,name=totalSpots,proto3" json:"totalSpots,omitempty"`
	FoodDrink  bool   `protobuf:"varint,3,opt,name=food_drink,json=foodDrink,proto3" json:"food_drink,omitempty"`
	Bar        bool   `protobuf:"varint,4,opt,name=bar,proto3" json:"bar,omitempty"`
	FastFood   bool   `protobuf:"varint,5,opt,name=fast_food,json=fastFood,proto3" json:"fast_food,omitempty"`
	Amenities  bool   `protobuf:"varint,6,opt,name=amenities,proto3" json:"amenities,omitempty"`
	Shopping   bool   `protobuf:"varint,7,opt,name=shopping,proto3" json:"shopping,omitempty"`
	Facilities bool   `protobuf:"varint,8,opt,name=facilities,proto3" json:"facilities,omitempty"`
	TwoWheeler bool   `protobuf:"varint,9,opt,name=two_wheeler,json=twoWheeler,proto3" json:"two_wheeler,omitempty"`
	Hotel      bool   `protobuf:"varint,10,opt,name=hotel,proto3" json:"hotel,omitempty"`
	Grill      bool   `protobuf:"varint,11,opt,name=grill,proto3" json:"grill,omitempty"`
	MedicalAid bool   `protobuf:"varint,12,opt,name=medical_aid,json=medicalAid,proto3" json:"medical_aid,omitempty"`
	Gas        bool   `protobuf:"varint,13,opt,name=gas,proto3" json:"gas,omitempty"`
	Charging   bool   `protobuf:"varint,14,opt,name=charging,proto3" json:"charging,omitempty"`
}

func (x *AreaDetails) Reset() {
	*x = AreaDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkplatz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaDetails) ProtoMessage() {}

func (x *AreaDetails) ProtoReflect() protoreflect.Message {
	mi := &file_parkplatz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaDetails.ProtoReflect.Descriptor instead.
func (*AreaDetails) Descriptor() ([]byte, []int) {
	return file_parkplatz_proto_rawDescGZIP(), []int{4}
}

func (x *AreaDetails) GetAreaId() string {
	if x != nil {
		return x.AreaId
	}
	return ""
}

func (x *AreaDetails) GetTotalSpots() int32 {
	if x != nil {
		return x.TotalSpots
	}
	return 0
}

func (x *AreaDetails) GetFoodDrink() bool {
	if x != nil {
		return x.FoodDrink
	}
	return false
}

func (x *AreaDetails) GetBar() bool {
	if x != nil {
		return x.Bar
	}
	return false
}

func (x *AreaDetails) GetFastFood() bool {
	if x != nil {
		return x.FastFood
	}
	return false
}

func (x *AreaDetails) GetAmenities() bool {
	if x != nil {
		return x.Amenities
	}
	return false
}

func (x *AreaDetails) GetShopping() bool {
	if x != nil {
		return x.Shopping
	}
	return false
}

func (x *AreaDetails) GetFacilities() bool {
	if x != nil {
		return x.Facilities
	}
	return false
}

func (x *AreaDetails) GetTwoWheeler() bool {
	if x != nil {
		return x.TwoWheeler
	}
	return false
}

func (x *AreaDetails) GetHotel() bool {
	if x != nil {
		return x.Hotel
	}
	return false
}

func (x *AreaDetails) GetGrill() bool {
	if x != nil {
		return x.Grill
	}
	return false
}

func (x *AreaDetails) GetMedicalAid() bool {
	if x != nil {
		return x.MedicalAid
	}
	return false
}

func (x *AreaDetails) GetGas() bool {
	if x != nil {
		return x.Gas
	}
	return false
}

func (x *AreaDetails) GetCharging() bool {
	if x != nil {
		return x.Charging
	}
	return false
}

type ExpulsionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaId string `protobuf:"bytes,1,opt,name=areaId,proto3" json:"areaId,omitempty"`
}

func (x *ExpulsionRequest) Reset() {
	*x = ExpulsionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkplatz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpulsionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpulsionRequest) ProtoMessage() {}

func (x *ExpulsionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkplatz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpulsionRequest.ProtoReflect.Descriptor instead.
func (*ExpulsionRequest) Descriptor() ([]byte, []int) {
	return file_parkplatz_proto_rawDescGZIP(), []int{5}
}

func (x *ExpulsionRequest) GetAreaId() string {
	if x != nil {
		return x.AreaId
	}
	return ""
}

type UtilizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
}

func (x *UtilizationRequest) Reset() {
	*x = UtilizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkplatz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtilizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtilizationRequest) ProtoMessage() {}

func (x *UtilizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkplatz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtilizationRequest.ProtoReflect.Descriptor instead.
func (*UtilizationRequest) Descriptor() ([]byte, []int) {
	return file_parkplatz_proto_rawDescGZIP(), []int{6}
}

func (x *UtilizationRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type UtilizationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName   string `protobuf:"bytes,1,opt,name=displayName,proto3" json:"displayName,omitempty"`
	TotalSpots    int32  `protobuf:"varint,2,opt,name=totalSpots,proto3" json:"totalSpots,omitempty"`
	UtilizedSpots int32  `protobuf:"varint,3,opt,name=utilizedSpots,proto3" json:"utilizedSpots,omitempty"`
}

func (x *UtilizationDetails) Reset() {
	*x = UtilizationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkplatz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtilizationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtilizationDetails) ProtoMessage() {}

func (x *UtilizationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_parkplatz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtilizationDetails.ProtoReflect.Descriptor instead.
func (*UtilizationDetails) Descriptor() ([]byte, []int) {
	return file_parkplatz_proto_rawDescGZIP(), []int{7}
}

func (x *UtilizationDetails) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *UtilizationDetails) GetTotalSpots() int32 {
	if x != nil {
		return x.TotalSpots
	}
	return 0
}

func (x *UtilizationDetails) GetUtilizedSpots() int32 {
	if x != nil {
		return x.UtilizedSpots
	}
	return 0
}

var File_parkplatz_proto protoreflect.FileDescriptor

var file_parkplatz_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74, 0x7a, 0x22, 0x8c, 0x01, 0x0a,
	0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x3a, 0x0a, 0x12, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0b, 0x41,
	0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x89, 0x03, 0x0a,
	0x0b, 0x41, 0x72, 0x65, 0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72,
	0x65, 0x61, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x6f,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x70, 0x6f, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x64, 0x72, 0x69,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x6f, 0x6f, 0x64, 0x44, 0x72,
	0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x62, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x66, 0x6f,
	0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x61, 0x73, 0x74, 0x46, 0x6f,
	0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6d, 0x65, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6d, 0x65, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a,
	0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x77, 0x6f, 0x5f, 0x77, 0x68, 0x65, 0x65, 0x6c, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x74, 0x77, 0x6f, 0x57, 0x68, 0x65, 0x65, 0x6c, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x69, 0x6c, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x67, 0x72, 0x69, 0x6c, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x2a, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x75,
	0x6c, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72,
	0x65, 0x61, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x12, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x12,
	0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x6f,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x70, 0x6f, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x53, 0x70, 0x6f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x75, 0x74, 0x69,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x70, 0x6f, 0x74, 0x73, 0x32, 0xfd, 0x02, 0x0a, 0x09, 0x50,
	0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74, 0x7a, 0x12, 0x4d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c,
	0x61, 0x74, 0x7a, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61,
	0x74, 0x7a, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61,
	0x74, 0x7a, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74,
	0x7a, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0b, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74,
	0x7a, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74, 0x7a,
	0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74, 0x7a,
	0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74, 0x7a, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x75, 0x6c, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74, 0x7a, 0x2e,
	0x45, 0x78, 0x70, 0x75, 0x6c, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x70, 0x6c, 0x61, 0x74, 0x7a, 0x2e, 0x41, 0x72, 0x65,
	0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_parkplatz_proto_rawDescOnce sync.Once
	file_parkplatz_proto_rawDescData = file_parkplatz_proto_rawDesc
)

func file_parkplatz_proto_rawDescGZIP() []byte {
	file_parkplatz_proto_rawDescOnce.Do(func() {
		file_parkplatz_proto_rawDescData = protoimpl.X.CompressGZIP(file_parkplatz_proto_rawDescData)
	})
	return file_parkplatz_proto_rawDescData
}

var file_parkplatz_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_parkplatz_proto_goTypes = []interface{}{
	(*ReservationRequest)(nil), // 0: parkplatz.ReservationRequest
	(*ReservationDetails)(nil), // 1: parkplatz.ReservationDetails
	(*TerminationRequest)(nil), // 2: parkplatz.TerminationRequest
	(*AreaRequest)(nil),        // 3: parkplatz.AreaRequest
	(*AreaDetails)(nil),        // 4: parkplatz.AreaDetails
	(*ExpulsionRequest)(nil),   // 5: parkplatz.ExpulsionRequest
	(*UtilizationRequest)(nil), // 6: parkplatz.UtilizationRequest
	(*UtilizationDetails)(nil), // 7: parkplatz.UtilizationDetails
}
var file_parkplatz_proto_depIdxs = []int32{
	0, // 0: parkplatz.Parkplatz.reservation:input_type -> parkplatz.ReservationRequest
	2, // 1: parkplatz.Parkplatz.termination:input_type -> parkplatz.TerminationRequest
	6, // 2: parkplatz.Parkplatz.utilization:input_type -> parkplatz.UtilizationRequest
	3, // 3: parkplatz.Parkplatz.provision:input_type -> parkplatz.AreaRequest
	5, // 4: parkplatz.Parkplatz.expulsion:input_type -> parkplatz.ExpulsionRequest
	1, // 5: parkplatz.Parkplatz.reservation:output_type -> parkplatz.ReservationDetails
	1, // 6: parkplatz.Parkplatz.termination:output_type -> parkplatz.ReservationDetails
	7, // 7: parkplatz.Parkplatz.utilization:output_type -> parkplatz.UtilizationDetails
	4, // 8: parkplatz.Parkplatz.provision:output_type -> parkplatz.AreaDetails
	4, // 9: parkplatz.Parkplatz.expulsion:output_type -> parkplatz.AreaDetails
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_parkplatz_proto_init() }
func file_parkplatz_proto_init() {
	if File_parkplatz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parkplatz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationRequest); i {
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
		file_parkplatz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationDetails); i {
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
		file_parkplatz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminationRequest); i {
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
		file_parkplatz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaRequest); i {
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
		file_parkplatz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaDetails); i {
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
		file_parkplatz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpulsionRequest); i {
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
		file_parkplatz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtilizationRequest); i {
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
		file_parkplatz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtilizationDetails); i {
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
			RawDescriptor: file_parkplatz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parkplatz_proto_goTypes,
		DependencyIndexes: file_parkplatz_proto_depIdxs,
		MessageInfos:      file_parkplatz_proto_msgTypes,
	}.Build()
	File_parkplatz_proto = out.File
	file_parkplatz_proto_rawDesc = nil
	file_parkplatz_proto_goTypes = nil
	file_parkplatz_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ParkplatzClient is the client API for Parkplatz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParkplatzClient interface {
	Reservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationDetails, error)
	Termination(ctx context.Context, in *TerminationRequest, opts ...grpc.CallOption) (*ReservationDetails, error)
	Utilization(ctx context.Context, in *UtilizationRequest, opts ...grpc.CallOption) (Parkplatz_UtilizationClient, error)
	Provision(ctx context.Context, in *AreaRequest, opts ...grpc.CallOption) (*AreaDetails, error)
	Expulsion(ctx context.Context, in *ExpulsionRequest, opts ...grpc.CallOption) (*AreaDetails, error)
}

type parkplatzClient struct {
	cc grpc.ClientConnInterface
}

func NewParkplatzClient(cc grpc.ClientConnInterface) ParkplatzClient {
	return &parkplatzClient{cc}
}

func (c *parkplatzClient) Reservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationDetails, error) {
	out := new(ReservationDetails)
	err := c.cc.Invoke(ctx, "/parkplatz.Parkplatz/reservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkplatzClient) Termination(ctx context.Context, in *TerminationRequest, opts ...grpc.CallOption) (*ReservationDetails, error) {
	out := new(ReservationDetails)
	err := c.cc.Invoke(ctx, "/parkplatz.Parkplatz/termination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkplatzClient) Utilization(ctx context.Context, in *UtilizationRequest, opts ...grpc.CallOption) (Parkplatz_UtilizationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Parkplatz_serviceDesc.Streams[0], "/parkplatz.Parkplatz/utilization", opts...)
	if err != nil {
		return nil, err
	}
	x := &parkplatzUtilizationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Parkplatz_UtilizationClient interface {
	Recv() (*UtilizationDetails, error)
	grpc.ClientStream
}

type parkplatzUtilizationClient struct {
	grpc.ClientStream
}

func (x *parkplatzUtilizationClient) Recv() (*UtilizationDetails, error) {
	m := new(UtilizationDetails)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *parkplatzClient) Provision(ctx context.Context, in *AreaRequest, opts ...grpc.CallOption) (*AreaDetails, error) {
	out := new(AreaDetails)
	err := c.cc.Invoke(ctx, "/parkplatz.Parkplatz/provision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkplatzClient) Expulsion(ctx context.Context, in *ExpulsionRequest, opts ...grpc.CallOption) (*AreaDetails, error) {
	out := new(AreaDetails)
	err := c.cc.Invoke(ctx, "/parkplatz.Parkplatz/expulsion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParkplatzServer is the server API for Parkplatz service.
type ParkplatzServer interface {
	Reservation(context.Context, *ReservationRequest) (*ReservationDetails, error)
	Termination(context.Context, *TerminationRequest) (*ReservationDetails, error)
	Utilization(*UtilizationRequest, Parkplatz_UtilizationServer) error
	Provision(context.Context, *AreaRequest) (*AreaDetails, error)
	Expulsion(context.Context, *ExpulsionRequest) (*AreaDetails, error)
}

// UnimplementedParkplatzServer can be embedded to have forward compatible implementations.
type UnimplementedParkplatzServer struct {
}

func (*UnimplementedParkplatzServer) Reservation(context.Context, *ReservationRequest) (*ReservationDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reservation not implemented")
}
func (*UnimplementedParkplatzServer) Termination(context.Context, *TerminationRequest) (*ReservationDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Termination not implemented")
}
func (*UnimplementedParkplatzServer) Utilization(*UtilizationRequest, Parkplatz_UtilizationServer) error {
	return status.Errorf(codes.Unimplemented, "method Utilization not implemented")
}
func (*UnimplementedParkplatzServer) Provision(context.Context, *AreaRequest) (*AreaDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provision not implemented")
}
func (*UnimplementedParkplatzServer) Expulsion(context.Context, *ExpulsionRequest) (*AreaDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expulsion not implemented")
}

func RegisterParkplatzServer(s *grpc.Server, srv ParkplatzServer) {
	s.RegisterService(&_Parkplatz_serviceDesc, srv)
}

func _Parkplatz_Reservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkplatzServer).Reservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkplatz.Parkplatz/Reservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkplatzServer).Reservation(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parkplatz_Termination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkplatzServer).Termination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkplatz.Parkplatz/Termination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkplatzServer).Termination(ctx, req.(*TerminationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parkplatz_Utilization_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UtilizationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ParkplatzServer).Utilization(m, &parkplatzUtilizationServer{stream})
}

type Parkplatz_UtilizationServer interface {
	Send(*UtilizationDetails) error
	grpc.ServerStream
}

type parkplatzUtilizationServer struct {
	grpc.ServerStream
}

func (x *parkplatzUtilizationServer) Send(m *UtilizationDetails) error {
	return x.ServerStream.SendMsg(m)
}

func _Parkplatz_Provision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkplatzServer).Provision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkplatz.Parkplatz/Provision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkplatzServer).Provision(ctx, req.(*AreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parkplatz_Expulsion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpulsionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkplatzServer).Expulsion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkplatz.Parkplatz/Expulsion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkplatzServer).Expulsion(ctx, req.(*ExpulsionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Parkplatz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parkplatz.Parkplatz",
	HandlerType: (*ParkplatzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "reservation",
			Handler:    _Parkplatz_Reservation_Handler,
		},
		{
			MethodName: "termination",
			Handler:    _Parkplatz_Termination_Handler,
		},
		{
			MethodName: "provision",
			Handler:    _Parkplatz_Provision_Handler,
		},
		{
			MethodName: "expulsion",
			Handler:    _Parkplatz_Expulsion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "utilization",
			Handler:       _Parkplatz_Utilization_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "parkplatz.proto",
}