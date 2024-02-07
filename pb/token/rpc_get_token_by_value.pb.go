// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: token/rpc_get_token_by_value.proto

package token

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

type GetTokenByValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenByValueRequest) Reset() {
	*x = GetTokenByValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_rpc_get_token_by_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenByValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenByValueRequest) ProtoMessage() {}

func (x *GetTokenByValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_rpc_get_token_by_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenByValueRequest.ProtoReflect.Descriptor instead.
func (*GetTokenByValueRequest) Descriptor() ([]byte, []int) {
	return file_token_rpc_get_token_by_value_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenByValueRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetTokenByValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenByValueResponse) Reset() {
	*x = GetTokenByValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_rpc_get_token_by_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenByValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenByValueResponse) ProtoMessage() {}

func (x *GetTokenByValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_rpc_get_token_by_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenByValueResponse.ProtoReflect.Descriptor instead.
func (*GetTokenByValueResponse) Descriptor() ([]byte, []int) {
	return file_token_rpc_get_token_by_value_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenByValueResponse) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_token_rpc_get_token_by_value_proto protoreflect.FileDescriptor

var file_token_rpc_get_token_by_value_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3a, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_rpc_get_token_by_value_proto_rawDescOnce sync.Once
	file_token_rpc_get_token_by_value_proto_rawDescData = file_token_rpc_get_token_by_value_proto_rawDesc
)

func file_token_rpc_get_token_by_value_proto_rawDescGZIP() []byte {
	file_token_rpc_get_token_by_value_proto_rawDescOnce.Do(func() {
		file_token_rpc_get_token_by_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_rpc_get_token_by_value_proto_rawDescData)
	})
	return file_token_rpc_get_token_by_value_proto_rawDescData
}

var file_token_rpc_get_token_by_value_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_token_rpc_get_token_by_value_proto_goTypes = []interface{}{
	(*GetTokenByValueRequest)(nil),  // 0: pb.GetTokenByValueRequest
	(*GetTokenByValueResponse)(nil), // 1: pb.GetTokenByValueResponse
	(*Token)(nil),                   // 2: pb.Token
}
var file_token_rpc_get_token_by_value_proto_depIdxs = []int32{
	2, // 0: pb.GetTokenByValueResponse.token:type_name -> pb.Token
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_token_rpc_get_token_by_value_proto_init() }
func file_token_rpc_get_token_by_value_proto_init() {
	if File_token_rpc_get_token_by_value_proto != nil {
		return
	}
	file_token_token_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_rpc_get_token_by_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenByValueRequest); i {
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
		file_token_rpc_get_token_by_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenByValueResponse); i {
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
			RawDescriptor: file_token_rpc_get_token_by_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_rpc_get_token_by_value_proto_goTypes,
		DependencyIndexes: file_token_rpc_get_token_by_value_proto_depIdxs,
		MessageInfos:      file_token_rpc_get_token_by_value_proto_msgTypes,
	}.Build()
	File_token_rpc_get_token_by_value_proto = out.File
	file_token_rpc_get_token_by_value_proto_rawDesc = nil
	file_token_rpc_get_token_by_value_proto_goTypes = nil
	file_token_rpc_get_token_by_value_proto_depIdxs = nil
}
