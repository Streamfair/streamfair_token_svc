// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: token_svc.proto

package pb

import (
	refresh_token "github.com/Streamfair/streamfair_token_svc/pb/refresh_token"
	token "github.com/Streamfair/streamfair_token_svc/pb/token"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_token_svc_proto protoreflect.FileDescriptor

var file_token_svc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79,
	0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2f, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62,
	0x79, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2f, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x93, 0x1e, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x92, 0x41, 0x23, 0x12, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x9f, 0x01, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5c, 0x92, 0x41, 0x29, 0x12, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20,
	0x62, 0x79, 0x20, 0x49, 0x44, 0x1a, 0x16, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12, 0xb1, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x92, 0x41, 0x2f, 0x12,
	0x12, 0x47, 0x65, 0x74, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x19, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x20, 0x62, 0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x67,
	0x65, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0xbb, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x6e, 0x92, 0x41, 0x35, 0x12, 0x15, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x20, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1c, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x20, 0x61, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x69,
	0x74, 0x73, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a,
	0x01, 0x2a, 0x22, 0x2b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66,
	0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0xac, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x65, 0x92, 0x41, 0x2f, 0x12, 0x12, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44,
	0x1a, 0x19, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x20, 0x61, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x20, 0x62, 0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12, 0xac,
	0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x65, 0x92, 0x41, 0x2f, 0x12, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x1a,
	0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20,
	0x62, 0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d,
	0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12, 0xbb, 0x01,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6e, 0x92, 0x41, 0x35,
	0x12, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62,
	0x79, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x01, 0x2a, 0x22, 0x2b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0xb5, 0x01, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63,
	0x92, 0x41, 0x2f, 0x12, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x64, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61,
	0x6c, 0x6c, 0x20, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x22, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0x88, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4b, 0x92, 0x41, 0x1f, 0x12, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x1a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x21, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x8f,
	0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x4f, 0x92, 0x41, 0x1f, 0x12, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x1a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x32, 0x22, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x8f, 0x01, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4f, 0x92, 0x41, 0x1f, 0x12, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x61, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0xc0, 0x01, 0x0a, 0x15, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6d, 0x92, 0x41, 0x35, 0x12, 0x17, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1a, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x61, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x62, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0xc0, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x92, 0x41, 0x33,
	0x12, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22, 0x2a, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0xc6, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x79, 0x92, 0x41, 0x3f, 0x12, 0x1a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x20, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x1a, 0x21, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20,
	0x62, 0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31,
	0x22, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69,
	0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0xce, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x76, 0x92, 0x41, 0x39, 0x12,
	0x17, 0x47, 0x65, 0x74, 0x20, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x1a, 0x1e, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79,
	0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x22, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0xec, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8a, 0x01, 0x92, 0x41, 0x3f, 0x12, 0x1a, 0x47, 0x65, 0x74, 0x20,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79,
	0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x21, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x62, 0x79, 0x20, 0x69,
	0x74, 0x73, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x42, 0x22,
	0x40, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2f, 0x7b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x7d, 0x12, 0xdc, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x22, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x64, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x78, 0x92, 0x41, 0x3d, 0x12, 0x1a, 0x47, 0x65, 0x74,
	0x20, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x20, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x1f, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c,
	0x20, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x22, 0x30,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0xb5, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x63, 0x92, 0x41, 0x2f, 0x12, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x18, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x22, 0x29, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0xb1, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x64, 0x92, 0x41, 0x2f, 0x12, 0x14, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x20, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x17, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x20, 0x61, 0x20, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x22, 0x2a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69,
	0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0xbc, 0x01, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x67, 0x92, 0x41, 0x2f, 0x12, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x32,
	0x2a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0xbc, 0x01, 0x0a, 0x12,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x67, 0x92, 0x41, 0x2f, 0x12, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x17, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x20, 0x61, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22, 0x2a,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x91, 0x01, 0x92, 0x41, 0x5f,
	0x12, 0x5d, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x20, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3c, 0x0a, 0x0a,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x12, 0x1d, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x1a, 0x0f, 0x6e, 0x65, 0x6c, 0x69, 0x78,
	0x40, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x6f, 0x2e, 0x64, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x31, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x66, 0x61, 0x69,
	0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_token_svc_proto_goTypes = []interface{}{
	(*token.CreateTokenRequest)(nil),                      // 0: pb.CreateTokenRequest
	(*token.GetTokenByIdRequest)(nil),                     // 1: pb.GetTokenByIdRequest
	(*token.GetTokenByValueRequest)(nil),                  // 2: pb.GetTokenByValueRequest
	(*token.RevokeTokenByValueRequest)(nil),               // 3: pb.RevokeTokenByValueRequest
	(*token.RevokeTokenByIdRequest)(nil),                  // 4: pb.RevokeTokenByIdRequest
	(*token.DeleteTokenByIdRequest)(nil),                  // 5: pb.DeleteTokenByIdRequest
	(*token.DeleteTokenByValueRequest)(nil),               // 6: pb.DeleteTokenByValueRequest
	(*token.ListRevokedTokensRequest)(nil),                // 7: pb.ListRevokedTokensRequest
	(*token.ListTokensRequest)(nil),                       // 8: pb.ListTokensRequest
	(*token.UpdateTokenRequest)(nil),                      // 9: pb.UpdateTokenRequest
	(*token.VerifyTokenRequest)(nil),                      // 10: pb.VerifyTokenRequest
	(*refresh_token.BlacklistRefreshTokenRequest)(nil),    // 11: pb.BlacklistRefreshTokenRequest
	(*refresh_token.CreateRefreshTokenRequest)(nil),       // 12: pb.CreateRefreshTokenRequest
	(*refresh_token.DeleteRefreshTokenRequest)(nil),       // 13: pb.DeleteRefreshTokenRequest
	(*refresh_token.GetRefreshTokenByIDRequest)(nil),      // 14: pb.GetRefreshTokenByIDRequest
	(*refresh_token.GetRefreshTokenByValueRequest)(nil),   // 15: pb.GetRefreshTokenByValueRequest
	(*refresh_token.GetRevokedRefreshTokensRequest)(nil),  // 16: pb.GetRevokedRefreshTokensRequest
	(*refresh_token.ListRefreshTokensRequest)(nil),        // 17: pb.ListRefreshTokensRequest
	(*refresh_token.RevokeRefreshTokenRequest)(nil),       // 18: pb.RevokeRefreshTokenRequest
	(*refresh_token.UpdateRefreshTokenRequest)(nil),       // 19: pb.UpdateRefreshTokenRequest
	(*refresh_token.VerifyRefreshTokenRequest)(nil),       // 20: pb.VerifyRefreshTokenRequest
	(*token.CreateTokenResponse)(nil),                     // 21: pb.CreateTokenResponse
	(*token.GetTokenByIdResponse)(nil),                    // 22: pb.GetTokenByIdResponse
	(*token.GetTokenByValueResponse)(nil),                 // 23: pb.GetTokenByValueResponse
	(*emptypb.Empty)(nil),                                 // 24: google.protobuf.Empty
	(*token.ListRevokedTokensResponse)(nil),               // 25: pb.ListRevokedTokensResponse
	(*token.ListTokensResponse)(nil),                      // 26: pb.ListTokensResponse
	(*token.UpdateTokenResponse)(nil),                     // 27: pb.UpdateTokenResponse
	(*token.VerifyTokenResponse)(nil),                     // 28: pb.VerifyTokenResponse
	(*refresh_token.CreateRefreshTokenResponse)(nil),      // 29: pb.CreateRefreshTokenResponse
	(*refresh_token.GetRefreshTokenByIDResponse)(nil),     // 30: pb.GetRefreshTokenByIDResponse
	(*refresh_token.GetRefreshTokenByValueResponse)(nil),  // 31: pb.GetRefreshTokenByValueResponse
	(*refresh_token.GetRevokedRefreshTokensResponse)(nil), // 32: pb.GetRevokedRefreshTokensResponse
	(*refresh_token.ListRefreshTokensResponse)(nil),       // 33: pb.ListRefreshTokensResponse
	(*refresh_token.UpdateRefreshTokenResponse)(nil),      // 34: pb.UpdateRefreshTokenResponse
	(*refresh_token.VerifyRefreshTokenResponse)(nil),      // 35: pb.VerifyRefreshTokenResponse
}
var file_token_svc_proto_depIdxs = []int32{
	0,  // 0: pb.TokenService.CreateToken:input_type -> pb.CreateTokenRequest
	1,  // 1: pb.TokenService.GetTokenById:input_type -> pb.GetTokenByIdRequest
	2,  // 2: pb.TokenService.GetTokenByValue:input_type -> pb.GetTokenByValueRequest
	3,  // 3: pb.TokenService.RevokeTokenByValue:input_type -> pb.RevokeTokenByValueRequest
	4,  // 4: pb.TokenService.RevokeTokenById:input_type -> pb.RevokeTokenByIdRequest
	5,  // 5: pb.TokenService.DeleteTokenById:input_type -> pb.DeleteTokenByIdRequest
	6,  // 6: pb.TokenService.DeleteTokenByValue:input_type -> pb.DeleteTokenByValueRequest
	7,  // 7: pb.TokenService.ListRevokedTokens:input_type -> pb.ListRevokedTokensRequest
	8,  // 8: pb.TokenService.ListTokens:input_type -> pb.ListTokensRequest
	9,  // 9: pb.TokenService.UpdateToken:input_type -> pb.UpdateTokenRequest
	10, // 10: pb.TokenService.VerifyToken:input_type -> pb.VerifyTokenRequest
	11, // 11: pb.TokenService.BlacklistRefreshToken:input_type -> pb.BlacklistRefreshTokenRequest
	12, // 12: pb.TokenService.CreateRefreshToken:input_type -> pb.CreateRefreshTokenRequest
	13, // 13: pb.TokenService.DeleteRefreshToken:input_type -> pb.DeleteRefreshTokenRequest
	14, // 14: pb.TokenService.GetRefreshTokenByID:input_type -> pb.GetRefreshTokenByIDRequest
	15, // 15: pb.TokenService.GetRefreshTokenByValue:input_type -> pb.GetRefreshTokenByValueRequest
	16, // 16: pb.TokenService.GetRevokedRefreshTokens:input_type -> pb.GetRevokedRefreshTokensRequest
	17, // 17: pb.TokenService.ListRefreshTokens:input_type -> pb.ListRefreshTokensRequest
	18, // 18: pb.TokenService.RevokeRefreshToken:input_type -> pb.RevokeRefreshTokenRequest
	19, // 19: pb.TokenService.UpdateRefreshToken:input_type -> pb.UpdateRefreshTokenRequest
	20, // 20: pb.TokenService.VerifyRefreshToken:input_type -> pb.VerifyRefreshTokenRequest
	21, // 21: pb.TokenService.CreateToken:output_type -> pb.CreateTokenResponse
	22, // 22: pb.TokenService.GetTokenById:output_type -> pb.GetTokenByIdResponse
	23, // 23: pb.TokenService.GetTokenByValue:output_type -> pb.GetTokenByValueResponse
	24, // 24: pb.TokenService.RevokeTokenByValue:output_type -> google.protobuf.Empty
	24, // 25: pb.TokenService.RevokeTokenById:output_type -> google.protobuf.Empty
	24, // 26: pb.TokenService.DeleteTokenById:output_type -> google.protobuf.Empty
	24, // 27: pb.TokenService.DeleteTokenByValue:output_type -> google.protobuf.Empty
	25, // 28: pb.TokenService.ListRevokedTokens:output_type -> pb.ListRevokedTokensResponse
	26, // 29: pb.TokenService.ListTokens:output_type -> pb.ListTokensResponse
	27, // 30: pb.TokenService.UpdateToken:output_type -> pb.UpdateTokenResponse
	28, // 31: pb.TokenService.VerifyToken:output_type -> pb.VerifyTokenResponse
	24, // 32: pb.TokenService.BlacklistRefreshToken:output_type -> google.protobuf.Empty
	29, // 33: pb.TokenService.CreateRefreshToken:output_type -> pb.CreateRefreshTokenResponse
	24, // 34: pb.TokenService.DeleteRefreshToken:output_type -> google.protobuf.Empty
	30, // 35: pb.TokenService.GetRefreshTokenByID:output_type -> pb.GetRefreshTokenByIDResponse
	31, // 36: pb.TokenService.GetRefreshTokenByValue:output_type -> pb.GetRefreshTokenByValueResponse
	32, // 37: pb.TokenService.GetRevokedRefreshTokens:output_type -> pb.GetRevokedRefreshTokensResponse
	33, // 38: pb.TokenService.ListRefreshTokens:output_type -> pb.ListRefreshTokensResponse
	24, // 39: pb.TokenService.RevokeRefreshToken:output_type -> google.protobuf.Empty
	34, // 40: pb.TokenService.UpdateRefreshToken:output_type -> pb.UpdateRefreshTokenResponse
	35, // 41: pb.TokenService.VerifyRefreshToken:output_type -> pb.VerifyRefreshTokenResponse
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_token_svc_proto_init() }
func file_token_svc_proto_init() {
	if File_token_svc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_token_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_svc_proto_goTypes,
		DependencyIndexes: file_token_svc_proto_depIdxs,
	}.Build()
	File_token_svc_proto = out.File
	file_token_svc_proto_rawDesc = nil
	file_token_svc_proto_goTypes = nil
	file_token_svc_proto_depIdxs = nil
}
