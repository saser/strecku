// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/role.proto

package streckuv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// The `Role` resource represents a many-to-many mapping between `User` and
// `Store` resources.
type Role struct {
	// The `name` field contains the resource name of the role. It is of the form
	// `roles/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The `user_name` field contains the resource name of the `User` resource
	// this `Role` maps.
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// The `store_name` field contains the resource name of the `Store` resource
	// this `Role` maps.
	StoreName            string   `protobuf:"bytes,3,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_742aa1c3ad5f8a33, []int{0}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Role) GetStoreName() string {
	if m != nil {
		return m.StoreName
	}
	return ""
}

func init() {
	proto.RegisterType((*Role)(nil), "strecku.v1.Role")
}

func init() { proto.RegisterFile("v1/role.proto", fileDescriptor_742aa1c3ad5f8a33) }

var fileDescriptor_742aa1c3ad5f8a33 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x33, 0xd4, 0x2f,
	0xca, 0xcf, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x2e, 0x29, 0x4a, 0x4d,
	0xce, 0x2e, 0xd5, 0x2b, 0x33, 0x54, 0x0a, 0xe3, 0x62, 0x09, 0xca, 0xcf, 0x49, 0x15, 0x12, 0xe2,
	0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xa4,
	0xb9, 0x38, 0x4b, 0x8b, 0x53, 0x8b, 0xe2, 0xc1, 0x12, 0x4c, 0x60, 0x09, 0x0e, 0x90, 0x80, 0x1f,
	0x48, 0x52, 0x96, 0x8b, 0xab, 0xb8, 0x24, 0xbf, 0x28, 0x15, 0x22, 0xcb, 0x0c, 0x96, 0xe5, 0x04,
	0x8b, 0x80, 0xa4, 0x9d, 0x3c, 0xb8, 0xf8, 0x92, 0xf3, 0x73, 0xf5, 0x10, 0x36, 0x39, 0x71, 0x82,
	0xec, 0x09, 0x00, 0x39, 0x20, 0x80, 0x31, 0x8a, 0x13, 0x2a, 0x51, 0x66, 0xb8, 0x88, 0x89, 0x39,
	0x38, 0x22, 0x62, 0x15, 0x13, 0x57, 0x30, 0x54, 0x69, 0x98, 0xe1, 0x29, 0x38, 0x27, 0x26, 0xcc,
	0x30, 0x89, 0x0d, 0xec, 0x68, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x5d, 0x6a, 0x06,
	0xc5, 0x00, 0x00, 0x00,
}
