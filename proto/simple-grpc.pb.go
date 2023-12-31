//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/simple-grpc.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.1
// source: proto/simple-grpc.proto

package proto

import (
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

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Campo1  float32 `protobuf:"fixed32,1,opt,name=Campo1,proto3" json:"Campo1,omitempty"`
	Campo2  float32 `protobuf:"fixed32,2,opt,name=Campo2,proto3" json:"Campo2,omitempty"`
	Campo3  float32 `protobuf:"fixed32,3,opt,name=Campo3,proto3" json:"Campo3,omitempty"`
	Campo4  float32 `protobuf:"fixed32,4,opt,name=Campo4,proto3" json:"Campo4,omitempty"`
	Campo5  float32 `protobuf:"fixed32,5,opt,name=Campo5,proto3" json:"Campo5,omitempty"`
	Campo6  float32 `protobuf:"fixed32,6,opt,name=Campo6,proto3" json:"Campo6,omitempty"`
	Campo7  float32 `protobuf:"fixed32,7,opt,name=Campo7,proto3" json:"Campo7,omitempty"`
	Campo8  float32 `protobuf:"fixed32,8,opt,name=Campo8,proto3" json:"Campo8,omitempty"`
	Campo9  float32 `protobuf:"fixed32,9,opt,name=Campo9,proto3" json:"Campo9,omitempty"`
	Campo10 float32 `protobuf:"fixed32,10,opt,name=Campo10,proto3" json:"Campo10,omitempty"`
	Campo11 float32 `protobuf:"fixed32,11,opt,name=Campo11,proto3" json:"Campo11,omitempty"`
	Campo12 float32 `protobuf:"fixed32,12,opt,name=Campo12,proto3" json:"Campo12,omitempty"`
	Campo13 float32 `protobuf:"fixed32,13,opt,name=Campo13,proto3" json:"Campo13,omitempty"`
	Campo14 float32 `protobuf:"fixed32,14,opt,name=Campo14,proto3" json:"Campo14,omitempty"`
	Campo15 float32 `protobuf:"fixed32,15,opt,name=Campo15,proto3" json:"Campo15,omitempty"`
	Campo16 float32 `protobuf:"fixed32,16,opt,name=Campo16,proto3" json:"Campo16,omitempty"`
	Campo17 float32 `protobuf:"fixed32,17,opt,name=Campo17,proto3" json:"Campo17,omitempty"`
	Campo18 float32 `protobuf:"fixed32,18,opt,name=Campo18,proto3" json:"Campo18,omitempty"`
	Campo19 float32 `protobuf:"fixed32,19,opt,name=Campo19,proto3" json:"Campo19,omitempty"`
	Campo20 float32 `protobuf:"fixed32,20,opt,name=Campo20,proto3" json:"Campo20,omitempty"`
	Campo21 float32 `protobuf:"fixed32,21,opt,name=Campo21,proto3" json:"Campo21,omitempty"`
	Campo22 float32 `protobuf:"fixed32,22,opt,name=Campo22,proto3" json:"Campo22,omitempty"`
	Campo23 float32 `protobuf:"fixed32,23,opt,name=Campo23,proto3" json:"Campo23,omitempty"`
	Campo24 float32 `protobuf:"fixed32,24,opt,name=Campo24,proto3" json:"Campo24,omitempty"`
	Campo25 float32 `protobuf:"fixed32,25,opt,name=Campo25,proto3" json:"Campo25,omitempty"`
	Campo26 float32 `protobuf:"fixed32,26,opt,name=Campo26,proto3" json:"Campo26,omitempty"`
	Campo27 float32 `protobuf:"fixed32,27,opt,name=Campo27,proto3" json:"Campo27,omitempty"`
	Campo28 float32 `protobuf:"fixed32,28,opt,name=Campo28,proto3" json:"Campo28,omitempty"`
	Campo29 float32 `protobuf:"fixed32,29,opt,name=Campo29,proto3" json:"Campo29,omitempty"`
	Campo30 float32 `protobuf:"fixed32,30,opt,name=Campo30,proto3" json:"Campo30,omitempty"`
	Campo31 float32 `protobuf:"fixed32,31,opt,name=Campo31,proto3" json:"Campo31,omitempty"`
	Campo32 float32 `protobuf:"fixed32,32,opt,name=Campo32,proto3" json:"Campo32,omitempty"`
	Campo33 float32 `protobuf:"fixed32,33,opt,name=Campo33,proto3" json:"Campo33,omitempty"`
	Campo34 float32 `protobuf:"fixed32,34,opt,name=Campo34,proto3" json:"Campo34,omitempty"`
	Campo35 float32 `protobuf:"fixed32,35,opt,name=Campo35,proto3" json:"Campo35,omitempty"`
	Campo36 float32 `protobuf:"fixed32,36,opt,name=Campo36,proto3" json:"Campo36,omitempty"`
	Campo37 float32 `protobuf:"fixed32,37,opt,name=Campo37,proto3" json:"Campo37,omitempty"`
	Campo38 float32 `protobuf:"fixed32,38,opt,name=Campo38,proto3" json:"Campo38,omitempty"`
	Campo39 float32 `protobuf:"fixed32,39,opt,name=Campo39,proto3" json:"Campo39,omitempty"`
	Campo40 float32 `protobuf:"fixed32,40,opt,name=Campo40,proto3" json:"Campo40,omitempty"`
	Campo41 float32 `protobuf:"fixed32,41,opt,name=Campo41,proto3" json:"Campo41,omitempty"`
	Campo42 float32 `protobuf:"fixed32,42,opt,name=Campo42,proto3" json:"Campo42,omitempty"`
	Campo43 float32 `protobuf:"fixed32,43,opt,name=Campo43,proto3" json:"Campo43,omitempty"`
	Campo44 float32 `protobuf:"fixed32,44,opt,name=Campo44,proto3" json:"Campo44,omitempty"`
	Campo45 float32 `protobuf:"fixed32,45,opt,name=Campo45,proto3" json:"Campo45,omitempty"`
	Campo46 float32 `protobuf:"fixed32,46,opt,name=Campo46,proto3" json:"Campo46,omitempty"`
	Campo47 float32 `protobuf:"fixed32,47,opt,name=Campo47,proto3" json:"Campo47,omitempty"`
	Campo48 float32 `protobuf:"fixed32,48,opt,name=Campo48,proto3" json:"Campo48,omitempty"`
	Campo49 float32 `protobuf:"fixed32,49,opt,name=Campo49,proto3" json:"Campo49,omitempty"`
	Campo50 float32 `protobuf:"fixed32,50,opt,name=Campo50,proto3" json:"Campo50,omitempty"`
	Campo51 float32 `protobuf:"fixed32,51,opt,name=Campo51,proto3" json:"Campo51,omitempty"`
	Campo52 float32 `protobuf:"fixed32,52,opt,name=Campo52,proto3" json:"Campo52,omitempty"`
	Campo53 float32 `protobuf:"fixed32,53,opt,name=Campo53,proto3" json:"Campo53,omitempty"`
	Campo54 float32 `protobuf:"fixed32,54,opt,name=Campo54,proto3" json:"Campo54,omitempty"`
	Campo55 float32 `protobuf:"fixed32,55,opt,name=Campo55,proto3" json:"Campo55,omitempty"`
	Campo56 float32 `protobuf:"fixed32,56,opt,name=Campo56,proto3" json:"Campo56,omitempty"`
	Campo57 float32 `protobuf:"fixed32,57,opt,name=Campo57,proto3" json:"Campo57,omitempty"`
	Campo58 float32 `protobuf:"fixed32,58,opt,name=Campo58,proto3" json:"Campo58,omitempty"`
	Campo59 float32 `protobuf:"fixed32,59,opt,name=Campo59,proto3" json:"Campo59,omitempty"`
	Campo60 float32 `protobuf:"fixed32,60,opt,name=Campo60,proto3" json:"Campo60,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_proto_simple_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetCampo1() float32 {
	if x != nil {
		return x.Campo1
	}
	return 0
}

func (x *Model) GetCampo2() float32 {
	if x != nil {
		return x.Campo2
	}
	return 0
}

func (x *Model) GetCampo3() float32 {
	if x != nil {
		return x.Campo3
	}
	return 0
}

func (x *Model) GetCampo4() float32 {
	if x != nil {
		return x.Campo4
	}
	return 0
}

func (x *Model) GetCampo5() float32 {
	if x != nil {
		return x.Campo5
	}
	return 0
}

func (x *Model) GetCampo6() float32 {
	if x != nil {
		return x.Campo6
	}
	return 0
}

func (x *Model) GetCampo7() float32 {
	if x != nil {
		return x.Campo7
	}
	return 0
}

func (x *Model) GetCampo8() float32 {
	if x != nil {
		return x.Campo8
	}
	return 0
}

func (x *Model) GetCampo9() float32 {
	if x != nil {
		return x.Campo9
	}
	return 0
}

func (x *Model) GetCampo10() float32 {
	if x != nil {
		return x.Campo10
	}
	return 0
}

func (x *Model) GetCampo11() float32 {
	if x != nil {
		return x.Campo11
	}
	return 0
}

func (x *Model) GetCampo12() float32 {
	if x != nil {
		return x.Campo12
	}
	return 0
}

func (x *Model) GetCampo13() float32 {
	if x != nil {
		return x.Campo13
	}
	return 0
}

func (x *Model) GetCampo14() float32 {
	if x != nil {
		return x.Campo14
	}
	return 0
}

func (x *Model) GetCampo15() float32 {
	if x != nil {
		return x.Campo15
	}
	return 0
}

func (x *Model) GetCampo16() float32 {
	if x != nil {
		return x.Campo16
	}
	return 0
}

func (x *Model) GetCampo17() float32 {
	if x != nil {
		return x.Campo17
	}
	return 0
}

func (x *Model) GetCampo18() float32 {
	if x != nil {
		return x.Campo18
	}
	return 0
}

func (x *Model) GetCampo19() float32 {
	if x != nil {
		return x.Campo19
	}
	return 0
}

func (x *Model) GetCampo20() float32 {
	if x != nil {
		return x.Campo20
	}
	return 0
}

func (x *Model) GetCampo21() float32 {
	if x != nil {
		return x.Campo21
	}
	return 0
}

func (x *Model) GetCampo22() float32 {
	if x != nil {
		return x.Campo22
	}
	return 0
}

func (x *Model) GetCampo23() float32 {
	if x != nil {
		return x.Campo23
	}
	return 0
}

func (x *Model) GetCampo24() float32 {
	if x != nil {
		return x.Campo24
	}
	return 0
}

func (x *Model) GetCampo25() float32 {
	if x != nil {
		return x.Campo25
	}
	return 0
}

func (x *Model) GetCampo26() float32 {
	if x != nil {
		return x.Campo26
	}
	return 0
}

func (x *Model) GetCampo27() float32 {
	if x != nil {
		return x.Campo27
	}
	return 0
}

func (x *Model) GetCampo28() float32 {
	if x != nil {
		return x.Campo28
	}
	return 0
}

func (x *Model) GetCampo29() float32 {
	if x != nil {
		return x.Campo29
	}
	return 0
}

func (x *Model) GetCampo30() float32 {
	if x != nil {
		return x.Campo30
	}
	return 0
}

func (x *Model) GetCampo31() float32 {
	if x != nil {
		return x.Campo31
	}
	return 0
}

func (x *Model) GetCampo32() float32 {
	if x != nil {
		return x.Campo32
	}
	return 0
}

func (x *Model) GetCampo33() float32 {
	if x != nil {
		return x.Campo33
	}
	return 0
}

func (x *Model) GetCampo34() float32 {
	if x != nil {
		return x.Campo34
	}
	return 0
}

func (x *Model) GetCampo35() float32 {
	if x != nil {
		return x.Campo35
	}
	return 0
}

func (x *Model) GetCampo36() float32 {
	if x != nil {
		return x.Campo36
	}
	return 0
}

func (x *Model) GetCampo37() float32 {
	if x != nil {
		return x.Campo37
	}
	return 0
}

func (x *Model) GetCampo38() float32 {
	if x != nil {
		return x.Campo38
	}
	return 0
}

func (x *Model) GetCampo39() float32 {
	if x != nil {
		return x.Campo39
	}
	return 0
}

func (x *Model) GetCampo40() float32 {
	if x != nil {
		return x.Campo40
	}
	return 0
}

func (x *Model) GetCampo41() float32 {
	if x != nil {
		return x.Campo41
	}
	return 0
}

func (x *Model) GetCampo42() float32 {
	if x != nil {
		return x.Campo42
	}
	return 0
}

func (x *Model) GetCampo43() float32 {
	if x != nil {
		return x.Campo43
	}
	return 0
}

func (x *Model) GetCampo44() float32 {
	if x != nil {
		return x.Campo44
	}
	return 0
}

func (x *Model) GetCampo45() float32 {
	if x != nil {
		return x.Campo45
	}
	return 0
}

func (x *Model) GetCampo46() float32 {
	if x != nil {
		return x.Campo46
	}
	return 0
}

func (x *Model) GetCampo47() float32 {
	if x != nil {
		return x.Campo47
	}
	return 0
}

func (x *Model) GetCampo48() float32 {
	if x != nil {
		return x.Campo48
	}
	return 0
}

func (x *Model) GetCampo49() float32 {
	if x != nil {
		return x.Campo49
	}
	return 0
}

func (x *Model) GetCampo50() float32 {
	if x != nil {
		return x.Campo50
	}
	return 0
}

func (x *Model) GetCampo51() float32 {
	if x != nil {
		return x.Campo51
	}
	return 0
}

func (x *Model) GetCampo52() float32 {
	if x != nil {
		return x.Campo52
	}
	return 0
}

func (x *Model) GetCampo53() float32 {
	if x != nil {
		return x.Campo53
	}
	return 0
}

func (x *Model) GetCampo54() float32 {
	if x != nil {
		return x.Campo54
	}
	return 0
}

func (x *Model) GetCampo55() float32 {
	if x != nil {
		return x.Campo55
	}
	return 0
}

func (x *Model) GetCampo56() float32 {
	if x != nil {
		return x.Campo56
	}
	return 0
}

func (x *Model) GetCampo57() float32 {
	if x != nil {
		return x.Campo57
	}
	return 0
}

func (x *Model) GetCampo58() float32 {
	if x != nil {
		return x.Campo58
	}
	return 0
}

func (x *Model) GetCampo59() float32 {
	if x != nil {
		return x.Campo59
	}
	return 0
}

func (x *Model) GetCampo60() float32 {
	if x != nil {
		return x.Campo60
	}
	return 0
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_proto_simple_grpc_proto_rawDescGZIP(), []int{1}
}

type GetAllReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model []*Model `protobuf:"bytes,1,rep,name=model,proto3" json:"model,omitempty"`
}

func (x *GetAllReply) Reset() {
	*x = GetAllReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReply) ProtoMessage() {}

func (x *GetAllReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReply.ProtoReflect.Descriptor instead.
func (*GetAllReply) Descriptor() ([]byte, []int) {
	return file_proto_simple_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllReply) GetModel() []*Model {
	if x != nil {
		return x.Model
	}
	return nil
}

type GetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOneRequest) Reset() {
	*x = GetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneRequest) ProtoMessage() {}

func (x *GetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneRequest.ProtoReflect.Descriptor instead.
func (*GetOneRequest) Descriptor() ([]byte, []int) {
	return file_proto_simple_grpc_proto_rawDescGZIP(), []int{3}
}

type GetOneReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model *Model `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *GetOneReply) Reset() {
	*x = GetOneReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneReply) ProtoMessage() {}

func (x *GetOneReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneReply.ProtoReflect.Descriptor instead.
func (*GetOneReply) Descriptor() ([]byte, []int) {
	return file_proto_simple_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *GetOneReply) GetModel() *Model {
	if x != nil {
		return x.Model
	}
	return nil
}

var File_proto_simple_grpc_proto protoreflect.FileDescriptor

var file_proto_simple_grpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x67, 0x72, 0x70, 0x63, 0x22, 0x8d, 0x0c, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x36, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x37, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x38, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x38, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x39, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x31, 0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31,
	0x30, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x61, 0x6d, 0x70, 0x6f, 0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61,
	0x6d, 0x70, 0x6f, 0x31, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x33,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x33, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d,
	0x70, 0x6f, 0x31, 0x35, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70,
	0x6f, 0x31, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x36, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x36, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x37, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x37, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x31, 0x38, 0x18, 0x12, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31,
	0x38, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x39, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x31, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x61, 0x6d, 0x70, 0x6f, 0x32, 0x30, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61,
	0x6d, 0x70, 0x6f, 0x32, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x31,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x31, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x32, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d,
	0x70, 0x6f, 0x32, 0x33, 0x18, 0x17, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70,
	0x6f, 0x32, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x34, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x34, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x35, 0x18, 0x19, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x32, 0x36, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32,
	0x36, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x37, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x37, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x61, 0x6d, 0x70, 0x6f, 0x32, 0x38, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61,
	0x6d, 0x70, 0x6f, 0x32, 0x38, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x39,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x32, 0x39, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x30, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d,
	0x70, 0x6f, 0x33, 0x31, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70,
	0x6f, 0x33, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x32, 0x18, 0x20,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x32, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x33, 0x18, 0x21, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x33, 0x34, 0x18, 0x22, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33,
	0x34, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x35, 0x18, 0x23, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x61, 0x6d, 0x70, 0x6f, 0x33, 0x36, 0x18, 0x24, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61,
	0x6d, 0x70, 0x6f, 0x33, 0x36, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x37,
	0x18, 0x25, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x37, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x38, 0x18, 0x26, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x33, 0x38, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d,
	0x70, 0x6f, 0x33, 0x39, 0x18, 0x27, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70,
	0x6f, 0x33, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x30, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x30, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x31, 0x18, 0x29, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x34, 0x32, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34,
	0x32, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x33, 0x18, 0x2b, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x61, 0x6d, 0x70, 0x6f, 0x34, 0x34, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61,
	0x6d, 0x70, 0x6f, 0x34, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x35,
	0x18, 0x2d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x35, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x36, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x36, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d,
	0x70, 0x6f, 0x34, 0x37, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70,
	0x6f, 0x34, 0x37, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x38, 0x18, 0x30,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x38, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x39, 0x18, 0x31, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x43, 0x61, 0x6d, 0x70, 0x6f, 0x34, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x35, 0x30, 0x18, 0x32, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35,
	0x30, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x31, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x61, 0x6d, 0x70, 0x6f, 0x35, 0x32, 0x18, 0x34, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61,
	0x6d, 0x70, 0x6f, 0x35, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x33,
	0x18, 0x35, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x33, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x34, 0x18, 0x36, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d,
	0x70, 0x6f, 0x35, 0x35, 0x18, 0x37, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70,
	0x6f, 0x35, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x36, 0x18, 0x38,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x36, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x37, 0x18, 0x39, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x37, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f,
	0x35, 0x38, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35,
	0x38, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x39, 0x18, 0x3b, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x43, 0x61, 0x6d, 0x70, 0x6f, 0x35, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x61, 0x6d, 0x70, 0x6f, 0x36, 0x30, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x43, 0x61,
	0x6d, 0x70, 0x6f, 0x36, 0x30, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x0f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x32, 0x8a, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x19,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x67, 0x6e, 0x69, 0x74, 0x65, 0x39, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_simple_grpc_proto_rawDescOnce sync.Once
	file_proto_simple_grpc_proto_rawDescData = file_proto_simple_grpc_proto_rawDesc
)

func file_proto_simple_grpc_proto_rawDescGZIP() []byte {
	file_proto_simple_grpc_proto_rawDescOnce.Do(func() {
		file_proto_simple_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_simple_grpc_proto_rawDescData)
	})
	return file_proto_simple_grpc_proto_rawDescData
}

var file_proto_simple_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_simple_grpc_proto_goTypes = []interface{}{
	(*Model)(nil),         // 0: simplegrpc.Model
	(*GetAllRequest)(nil), // 1: simplegrpc.GetAllRequest
	(*GetAllReply)(nil),   // 2: simplegrpc.GetAllReply
	(*GetOneRequest)(nil), // 3: simplegrpc.GetOneRequest
	(*GetOneReply)(nil),   // 4: simplegrpc.GetOneReply
}
var file_proto_simple_grpc_proto_depIdxs = []int32{
	0, // 0: simplegrpc.GetAllReply.model:type_name -> simplegrpc.Model
	0, // 1: simplegrpc.GetOneReply.model:type_name -> simplegrpc.Model
	1, // 2: simplegrpc.Services.GetAll:input_type -> simplegrpc.GetAllRequest
	3, // 3: simplegrpc.Services.GetOne:input_type -> simplegrpc.GetOneRequest
	2, // 4: simplegrpc.Services.GetAll:output_type -> simplegrpc.GetAllReply
	4, // 5: simplegrpc.Services.GetOne:output_type -> simplegrpc.GetOneReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_simple_grpc_proto_init() }
func file_proto_simple_grpc_proto_init() {
	if File_proto_simple_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_simple_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_proto_simple_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_proto_simple_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReply); i {
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
		file_proto_simple_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneRequest); i {
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
		file_proto_simple_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneReply); i {
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
			RawDescriptor: file_proto_simple_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_simple_grpc_proto_goTypes,
		DependencyIndexes: file_proto_simple_grpc_proto_depIdxs,
		MessageInfos:      file_proto_simple_grpc_proto_msgTypes,
	}.Build()
	File_proto_simple_grpc_proto = out.File
	file_proto_simple_grpc_proto_rawDesc = nil
	file_proto_simple_grpc_proto_goTypes = nil
	file_proto_simple_grpc_proto_depIdxs = nil
}
