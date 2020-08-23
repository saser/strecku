// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: saser/strecku/v1/strecku.proto

package streckuv1

import (
	proto "github.com/golang/protobuf/proto"
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

// AuthenticateUserRequest is the request message for AuthenticateUser.
type AuthenticateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// email_address contains the email address of the user that authenticates.
	// Required.
	EmailAddress string `protobuf:"bytes,1,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	// password contains the password of the user that authenticates.
	// Required.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthenticateUserRequest) Reset() {
	*x = AuthenticateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateUserRequest) ProtoMessage() {}

func (x *AuthenticateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateUserRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateUserRequest) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateUserRequest) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *AuthenticateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// AuthenticateUserResponse is the response message for AuthenticateUser.
type AuthenticateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user is the user that was authenticated.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// token is an opaque token that can be used to authenticate a request as this
	// user.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthenticateUserResponse) Reset() {
	*x = AuthenticateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateUserResponse) ProtoMessage() {}

func (x *AuthenticateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateUserResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateUserResponse) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *AuthenticateUserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_saser_strecku_v1_strecku_proto protoreflect.FileDescriptor

var file_saser_strecku_v1_strecku_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b,
	0x75, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5c, 0x0a, 0x18, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x74, 0x0a, 0x07, 0x53, 0x74, 0x72,
	0x65, 0x63, 0x6b, 0x55, 0x12, 0x69, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x41, 0x0a, 0x13, 0x73, 0x65, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x55, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x63, 0x6b, 0x75, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saser_strecku_v1_strecku_proto_rawDescOnce sync.Once
	file_saser_strecku_v1_strecku_proto_rawDescData = file_saser_strecku_v1_strecku_proto_rawDesc
)

func file_saser_strecku_v1_strecku_proto_rawDescGZIP() []byte {
	file_saser_strecku_v1_strecku_proto_rawDescOnce.Do(func() {
		file_saser_strecku_v1_strecku_proto_rawDescData = protoimpl.X.CompressGZIP(file_saser_strecku_v1_strecku_proto_rawDescData)
	})
	return file_saser_strecku_v1_strecku_proto_rawDescData
}

var file_saser_strecku_v1_strecku_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_saser_strecku_v1_strecku_proto_goTypes = []interface{}{
	(*AuthenticateUserRequest)(nil),  // 0: saser.strecku.v1.AuthenticateUserRequest
	(*AuthenticateUserResponse)(nil), // 1: saser.strecku.v1.AuthenticateUserResponse
	(*User)(nil),                     // 2: saser.strecku.v1.User
}
var file_saser_strecku_v1_strecku_proto_depIdxs = []int32{
	2, // 0: saser.strecku.v1.AuthenticateUserResponse.user:type_name -> saser.strecku.v1.User
	0, // 1: saser.strecku.v1.StreckU.AuthenticateUser:input_type -> saser.strecku.v1.AuthenticateUserRequest
	1, // 2: saser.strecku.v1.StreckU.AuthenticateUser:output_type -> saser.strecku.v1.AuthenticateUserResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_saser_strecku_v1_strecku_proto_init() }
func file_saser_strecku_v1_strecku_proto_init() {
	if File_saser_strecku_v1_strecku_proto != nil {
		return
	}
	file_saser_strecku_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_saser_strecku_v1_strecku_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateUserRequest); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateUserResponse); i {
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
			RawDescriptor: file_saser_strecku_v1_strecku_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_saser_strecku_v1_strecku_proto_goTypes,
		DependencyIndexes: file_saser_strecku_v1_strecku_proto_depIdxs,
		MessageInfos:      file_saser_strecku_v1_strecku_proto_msgTypes,
	}.Build()
	File_saser_strecku_v1_strecku_proto = out.File
	file_saser_strecku_v1_strecku_proto_rawDesc = nil
	file_saser_strecku_v1_strecku_proto_goTypes = nil
	file_saser_strecku_v1_strecku_proto_depIdxs = nil
}
