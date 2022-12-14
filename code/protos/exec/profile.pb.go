// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: profile.proto

package exec

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

type ProfileRequest_Type int32

const (
	ProfileRequest_nnMeter     ProfileRequest_Type = 0
	ProfileRequest_paddlelite  ProfileRequest_Type = 1
	ProfileRequest_tflite      ProfileRequest_Type = 2
	ProfileRequest_onnxruntime ProfileRequest_Type = 3
)

// Enum value maps for ProfileRequest_Type.
var (
	ProfileRequest_Type_name = map[int32]string{
		0: "nnMeter",
		1: "paddlelite",
		2: "tflite",
		3: "onnxruntime",
	}
	ProfileRequest_Type_value = map[string]int32{
		"nnMeter":     0,
		"paddlelite":  1,
		"tflite":      2,
		"onnxruntime": 3,
	}
)

func (x ProfileRequest_Type) Enum() *ProfileRequest_Type {
	p := new(ProfileRequest_Type)
	*p = x
	return p
}

func (x ProfileRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProfileRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_profile_proto_enumTypes[0].Descriptor()
}

func (ProfileRequest_Type) Type() protoreflect.EnumType {
	return &file_profile_proto_enumTypes[0]
}

func (x ProfileRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProfileRequest_Type.Descriptor instead.
func (ProfileRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0, 0}
}

type ProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelFile  *File  `protobuf:"bytes,1,opt,name=modelFile,proto3" json:"modelFile,omitempty"`
	DeviceName string `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	// Types that are assignable to Args:
	//
	//	*ProfileRequest_NnmeterArgs
	//	*ProfileRequest_PaddleLiteArgs
	Args isProfileRequest_Args `protobuf_oneof:"args"`
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileRequest) GetModelFile() *File {
	if x != nil {
		return x.ModelFile
	}
	return nil
}

func (x *ProfileRequest) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (m *ProfileRequest) GetArgs() isProfileRequest_Args {
	if m != nil {
		return m.Args
	}
	return nil
}

func (x *ProfileRequest) GetNnmeterArgs() *NNMeterArgs {
	if x, ok := x.GetArgs().(*ProfileRequest_NnmeterArgs); ok {
		return x.NnmeterArgs
	}
	return nil
}

func (x *ProfileRequest) GetPaddleLiteArgs() *PaddleLiteArgs {
	if x, ok := x.GetArgs().(*ProfileRequest_PaddleLiteArgs); ok {
		return x.PaddleLiteArgs
	}
	return nil
}

type isProfileRequest_Args interface {
	isProfileRequest_Args()
}

type ProfileRequest_NnmeterArgs struct {
	NnmeterArgs *NNMeterArgs `protobuf:"bytes,10,opt,name=nnmeterArgs,proto3,oneof"`
}

type ProfileRequest_PaddleLiteArgs struct {
	PaddleLiteArgs *PaddleLiteArgs `protobuf:"bytes,11,opt,name=paddleLiteArgs,proto3,oneof"`
}

func (*ProfileRequest_NnmeterArgs) isProfileRequest_Args() {}

func (*ProfileRequest_PaddleLiteArgs) isProfileRequest_Args() {}

type ProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProfileResponse) Reset() {
	*x = ProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResponse) ProtoMessage() {}

func (x *ProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResponse.ProtoReflect.Descriptor instead.
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Size     uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{2}
}

func (x *File) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *File) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type NNMeterArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Predictor string `protobuf:"bytes,1,opt,name=predictor,proto3" json:"predictor,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Framework string `protobuf:"bytes,3,opt,name=framework,proto3" json:"framework,omitempty"`
}

func (x *NNMeterArgs) Reset() {
	*x = NNMeterArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NNMeterArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NNMeterArgs) ProtoMessage() {}

func (x *NNMeterArgs) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NNMeterArgs.ProtoReflect.Descriptor instead.
func (*NNMeterArgs) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{3}
}

func (x *NNMeterArgs) GetPredictor() string {
	if x != nil {
		return x.Predictor
	}
	return ""
}

func (x *NNMeterArgs) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NNMeterArgs) GetFramework() string {
	if x != nil {
		return x.Framework
	}
	return ""
}

type PaddleLiteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *PaddleLiteArgs) Reset() {
	*x = PaddleLiteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaddleLiteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaddleLiteArgs) ProtoMessage() {}

func (x *PaddleLiteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaddleLiteArgs.ProtoReflect.Descriptor instead.
func (*PaddleLiteArgs) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{4}
}

func (x *PaddleLiteArgs) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_profile_proto protoreflect.FileDescriptor

var file_profile_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xa1, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x09, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x6e, 0x6e, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x41, 0x72, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x4e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73,
	0x48, 0x00, 0x52, 0x0b, 0x6e, 0x6e, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x12,
	0x40, 0x0a, 0x0e, 0x70, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x41, 0x72, 0x67,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x50, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x48,
	0x00, 0x52, 0x0e, 0x70, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x41, 0x72, 0x67,
	0x73, 0x22, 0x40, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x6e, 0x6e, 0x4d,
	0x65, 0x74, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x70, 0x61, 0x64, 0x64, 0x6c, 0x65,
	0x6c, 0x69, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65,
	0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x6f, 0x6e, 0x6e, 0x78, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x10, 0x03, 0x42, 0x06, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x23, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x4a, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x63, 0x0a, 0x0b,
	0x4e, 0x4e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x22, 0x2a, 0x0a, 0x0e, 0x50, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x97, 0x01,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x65, 0x78,
	0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_proto_rawDescOnce sync.Once
	file_profile_proto_rawDescData = file_profile_proto_rawDesc
)

func file_profile_proto_rawDescGZIP() []byte {
	file_profile_proto_rawDescOnce.Do(func() {
		file_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_proto_rawDescData)
	})
	return file_profile_proto_rawDescData
}

var file_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_profile_proto_goTypes = []interface{}{
	(ProfileRequest_Type)(0), // 0: protos.ProfileRequest.Type
	(*ProfileRequest)(nil),   // 1: protos.ProfileRequest
	(*ProfileResponse)(nil),  // 2: protos.ProfileResponse
	(*File)(nil),             // 3: protos.File
	(*NNMeterArgs)(nil),      // 4: protos.NNMeterArgs
	(*PaddleLiteArgs)(nil),   // 5: protos.PaddleLiteArgs
}
var file_profile_proto_depIdxs = []int32{
	3, // 0: protos.ProfileRequest.modelFile:type_name -> protos.File
	4, // 1: protos.ProfileRequest.nnmeterArgs:type_name -> protos.NNMeterArgs
	5, // 2: protos.ProfileRequest.paddleLiteArgs:type_name -> protos.PaddleLiteArgs
	1, // 3: protos.Profile.profileWithArgs:input_type -> protos.ProfileRequest
	1, // 4: protos.Profile.getProfileAbility:input_type -> protos.ProfileRequest
	2, // 5: protos.Profile.profileWithArgs:output_type -> protos.ProfileResponse
	2, // 6: protos.Profile.getProfileAbility:output_type -> protos.ProfileResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_profile_proto_init() }
func file_profile_proto_init() {
	if File_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileRequest); i {
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
		file_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileResponse); i {
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
		file_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NNMeterArgs); i {
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
		file_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaddleLiteArgs); i {
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
	file_profile_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProfileRequest_NnmeterArgs)(nil),
		(*ProfileRequest_PaddleLiteArgs)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_proto_goTypes,
		DependencyIndexes: file_profile_proto_depIdxs,
		EnumInfos:         file_profile_proto_enumTypes,
		MessageInfos:      file_profile_proto_msgTypes,
	}.Build()
	File_profile_proto = out.File
	file_profile_proto_rawDesc = nil
	file_profile_proto_goTypes = nil
	file_profile_proto_depIdxs = nil
}
