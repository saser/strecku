// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/user_api.proto

package streckuv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for UserAPI.ListUsers.
type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6beeca62d8388557, []int{0}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

// Response message for UserAPI.ListUsers.
type ListUsersResponse struct {
	// The list of users.
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6beeca62d8388557, []int{1}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*ListUsersRequest)(nil), "strecku.v1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "strecku.v1.ListUsersResponse")
}

func init() { proto.RegisterFile("v1/user_api.proto", fileDescriptor_6beeca62d8388557) }

var fileDescriptor_6beeca62d8388557 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x33, 0xd4, 0x2f,
	0x2d, 0x4e, 0x2d, 0x8a, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a,
	0x2e, 0x29, 0x4a, 0x4d, 0xce, 0x2e, 0xd5, 0x2b, 0x33, 0x94, 0xe2, 0x85, 0x4a, 0x43, 0xa4, 0x94,
	0x84, 0xb8, 0x04, 0x7c, 0x32, 0x8b, 0x4b, 0x42, 0x8b, 0x53, 0x8b, 0x8a, 0x83, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x94, 0xac, 0xb9, 0x04, 0x91, 0xc4, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85,
	0xd4, 0xb8, 0x58, 0x41, 0xda, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x04, 0xf4, 0x10,
	0x66, 0xea, 0x81, 0x54, 0x06, 0x41, 0xa4, 0x8d, 0x82, 0xb9, 0xd8, 0x41, 0x5c, 0xc7, 0x00, 0x4f,
	0x21, 0x0f, 0x2e, 0x4e, 0xb8, 0x39, 0x42, 0x32, 0xc8, 0x1a, 0xd0, 0xad, 0x94, 0x92, 0xc5, 0x21,
	0x0b, 0xb1, 0xdc, 0xc9, 0x9b, 0x8b, 0x2f, 0x39, 0x3f, 0x17, 0x49, 0x8d, 0x13, 0x0f, 0xd8, 0x92,
	0x82, 0xcc, 0x00, 0x90, 0x2f, 0x02, 0x18, 0xa3, 0x38, 0xa1, 0x72, 0x65, 0x86, 0x8b, 0x98, 0x98,
	0x83, 0x23, 0x22, 0x56, 0x31, 0x71, 0x05, 0x43, 0x55, 0x87, 0x19, 0x9e, 0x82, 0x73, 0x62, 0xc2,
	0x0c, 0x93, 0xd8, 0xc0, 0x3e, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x23, 0x82, 0x2e, 0xda,
	0x29, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserAPIClient is the client API for UserAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserAPIClient interface {
	// ListUsers returns a list of users. The order is unspecified but deterministic.
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type userAPIClient struct {
	cc *grpc.ClientConn
}

func NewUserAPIClient(cc *grpc.ClientConn) UserAPIClient {
	return &userAPIClient{cc}
}

func (c *userAPIClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/strecku.v1.UserAPI/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAPIServer is the server API for UserAPI service.
type UserAPIServer interface {
	// ListUsers returns a list of users. The order is unspecified but deterministic.
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
}

// UnimplementedUserAPIServer can be embedded to have forward compatible implementations.
type UnimplementedUserAPIServer struct {
}

func (*UnimplementedUserAPIServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

func RegisterUserAPIServer(s *grpc.Server, srv UserAPIServer) {
	s.RegisterService(&_UserAPI_serviceDesc, srv)
}

func _UserAPI_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strecku.v1.UserAPI/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "strecku.v1.UserAPI",
	HandlerType: (*UserAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _UserAPI_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/user_api.proto",
}
