// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/store.proto

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

// A `Store` represents a store which can have products, purchases, etc.
type Store struct {
	// The `name` corresponds to the resource name of the store. It is of the form
	// `stores/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The `display_name` contains the name of the store as it should be displayed
	// to a human. For example: "Foobar Coffee".
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Store) Reset()         { *m = Store{} }
func (m *Store) String() string { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()    {}
func (*Store) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3638a3da90d9960, []int{0}
}

func (m *Store) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Store.Unmarshal(m, b)
}
func (m *Store) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Store.Marshal(b, m, deterministic)
}
func (m *Store) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Store.Merge(m, src)
}
func (m *Store) XXX_Size() int {
	return xxx_messageInfo_Store.Size(m)
}
func (m *Store) XXX_DiscardUnknown() {
	xxx_messageInfo_Store.DiscardUnknown(m)
}

var xxx_messageInfo_Store proto.InternalMessageInfo

func (m *Store) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Store) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func init() {
	proto.RegisterType((*Store)(nil), "strecku.v1.Store")
}

func init() { proto.RegisterFile("v1/store.proto", fileDescriptor_e3638a3da90d9960) }

var fileDescriptor_e3638a3da90d9960 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x33, 0xd4, 0x2f,
	0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x2e, 0x29, 0x4a,
	0x4d, 0xce, 0x2e, 0xd5, 0x2b, 0x33, 0x54, 0xb2, 0xe3, 0x62, 0x0d, 0x06, 0x49, 0x09, 0x09, 0x71,
	0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x8a,
	0x5c, 0x3c, 0x29, 0x99, 0xc5, 0x05, 0x39, 0x89, 0x95, 0xf1, 0x60, 0x39, 0x26, 0xb0, 0x1c, 0x37,
	0x54, 0xcc, 0x2f, 0x31, 0x37, 0xd5, 0xc9, 0x93, 0x8b, 0x2f, 0x39, 0x3f, 0x57, 0x0f, 0x61, 0xa2,
	0x13, 0x17, 0xd8, 0xbc, 0x00, 0x90, 0x4d, 0x01, 0x8c, 0x51, 0x9c, 0x50, 0x99, 0x32, 0xc3, 0x45,
	0x4c, 0xcc, 0xc1, 0x11, 0x11, 0xab, 0x98, 0xb8, 0x82, 0xa1, 0x6a, 0xc3, 0x0c, 0x4f, 0xc1, 0x39,
	0x31, 0x61, 0x86, 0x49, 0x6c, 0x60, 0xd7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x15,
	0xf0, 0x74, 0xaf, 0x00, 0x00, 0x00,
}