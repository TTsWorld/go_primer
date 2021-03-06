// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: usericon.proto

package grpc

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

type UserIconRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UId []int64 `protobuf:"varint,1,rep,packed,name=uId,proto3" json:"uId,omitempty"`
}

func (x *UserIconRequest) Reset() {
	*x = UserIconRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usericon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIconRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIconRequest) ProtoMessage() {}

func (x *UserIconRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usericon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIconRequest.ProtoReflect.Descriptor instead.
func (*UserIconRequest) Descriptor() ([]byte, []int) {
	return file_usericon_proto_rawDescGZIP(), []int{0}
}

func (x *UserIconRequest) GetUId() []int64 {
	if x != nil {
		return x.UId
	}
	return nil
}

type UserIconResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RspHead *RspHead `protobuf:"bytes,1,opt,name=rspHead,proto3" json:"rspHead,omitempty"`
	Icon    []string `protobuf:"bytes,2,rep,name=icon,proto3" json:"icon,omitempty"` //????????????
}

func (x *UserIconResponse) Reset() {
	*x = UserIconResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usericon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIconResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIconResponse) ProtoMessage() {}

func (x *UserIconResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usericon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIconResponse.ProtoReflect.Descriptor instead.
func (*UserIconResponse) Descriptor() ([]byte, []int) {
	return file_usericon_proto_rawDescGZIP(), []int{1}
}

func (x *UserIconResponse) GetRspHead() *RspHead {
	if x != nil {
		return x.RspHead
	}
	return nil
}

func (x *UserIconResponse) GetIcon() []string {
	if x != nil {
		return x.Icon
	}
	return nil
}

var File_usericon_proto protoreflect.FileDescriptor

var file_usericon_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x0e, 0x72, 0x73, 0x70, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x75, 0x49, 0x64, 0x22,
	0x53, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x52, 0x07, 0x72, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x32, 0x50, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e,
	0x12, 0x44, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x12,
	0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x63, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x69,
	0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usericon_proto_rawDescOnce sync.Once
	file_usericon_proto_rawDescData = file_usericon_proto_rawDesc
)

func file_usericon_proto_rawDescGZIP() []byte {
	file_usericon_proto_rawDescOnce.Do(func() {
		file_usericon_proto_rawDescData = protoimpl.X.CompressGZIP(file_usericon_proto_rawDescData)
	})
	return file_usericon_proto_rawDescData
}

var file_usericon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_usericon_proto_goTypes = []interface{}{
	(*UserIconRequest)(nil),  // 0: grpc.api.UserIconRequest
	(*UserIconResponse)(nil), // 1: grpc.api.UserIconResponse
	(*RspHead)(nil),          // 2: grpc.api.RspHead
}
var file_usericon_proto_depIdxs = []int32{
	2, // 0: grpc.api.UserIconResponse.rspHead:type_name -> grpc.api.RspHead
	0, // 1: grpc.api.UserIcon.getUserIcon:input_type -> grpc.api.UserIconRequest
	1, // 2: grpc.api.UserIcon.getUserIcon:output_type -> grpc.api.UserIconResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_usericon_proto_init() }
func file_usericon_proto_init() {
	if File_usericon_proto != nil {
		return
	}
	file_rsp_head_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_usericon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIconRequest); i {
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
		file_usericon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIconResponse); i {
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
			RawDescriptor: file_usericon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usericon_proto_goTypes,
		DependencyIndexes: file_usericon_proto_depIdxs,
		MessageInfos:      file_usericon_proto_msgTypes,
	}.Build()
	File_usericon_proto = out.File
	file_usericon_proto_rawDesc = nil
	file_usericon_proto_goTypes = nil
	file_usericon_proto_depIdxs = nil
}
