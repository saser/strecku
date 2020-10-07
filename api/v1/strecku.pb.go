// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: saser/strecku/v1/strecku.proto

package pb

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

// User represents a user in the system.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the resource name of the user.
	// Format: users/{user}
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// email_address is the email address of the user.
	// The email address is unique among all users.
	// Required.
	EmailAddress string `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	// display_name is the name of the user as it would be displayed to a human.
	// Required.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// superuser is true if this user is a superuser, i.e., a user that is
	// authorized for any request.
	Superuser bool `protobuf:"varint,4,opt,name=superuser,proto3" json:"superuser,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetSuperuser() bool {
	if x != nil {
		return x.Superuser
	}
	return false
}

// Store represents a store in the system.
type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the resource name of the store.
	// Format: stores/{store}
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// display_name is the name of the store as it would be displayed to a human.
	// Required.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{1}
}

func (x *Store) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Store) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// Membership represents one instance of a many-to-many relation between users
// and stores, meaning that a user is a member of a store.
type Membership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the resource name of the membership.
	// Format: memberships/{membership}
	// Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// user is the resource name of the user.
	// Format: users/{user}
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// store is the resource name of the store.
	// Format: stores/{store}
	Store string `protobuf:"bytes,3,opt,name=store,proto3" json:"store,omitempty"`
	// administrator is true if the user is an administrator of the store.
	Administrator bool `protobuf:"varint,4,opt,name=administrator,proto3" json:"administrator,omitempty"`
	// discount is true if the user has a discount in the store.
	Discount bool `protobuf:"varint,5,opt,name=discount,proto3" json:"discount,omitempty"`
}

func (x *Membership) Reset() {
	*x = Membership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Membership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Membership) ProtoMessage() {}

func (x *Membership) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Membership.ProtoReflect.Descriptor instead.
func (*Membership) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{2}
}

func (x *Membership) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Membership) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Membership) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *Membership) GetAdministrator() bool {
	if x != nil {
		return x.Administrator
	}
	return false
}

func (x *Membership) GetDiscount() bool {
	if x != nil {
		return x.Discount
	}
	return false
}

// GetUserRequest is the request message for GetUser.
type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the resource name of the user to get.
	// Format: users/{user}
	// Required.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ListUsersRequest is the request message for ListUsers.
//
// (-- api-linter: core::0132::request-parent-required=disabled
//     aip.dev/not-precedent: Users are top-level resources. --)
type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page_size is the maximum number of users to return.
	// If unspecified, the server will choose a suitable number.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// page_token contains an opaque string used to get the next page of
	// results. It is usually provided by the previous call to ListUsers.
	// If unspecified, the first page will be returned.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{4}
}

func (x *ListUsersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListUsersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListUsersResponse is the response message for ListUsers.
type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// users contains the page of users.
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// next_page_token contains an opaque string used to get the next page of
	// results. Provide this in a subsequent call to ListUsers.
	// If this field is empty, there are no more pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{5}
}

func (x *ListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ListUsersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// CreateUserRequest is the request message for CreateUser.
//
// (-- api-linter: core::0133::request-parent-required=disabled
//     aip.dev/not-precedent: Users are top-level resources. --)
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user is the user to be created.
	// Required.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// password contains the password this user should use to authenticate.
	// Required.
	//
	// (-- api-linter: core::0133::request-unknown-fields=disabled
	//     aip.dev/not-precedent: A password is required for each user. --)
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// GetStoreRequest is the request message for GetStore.
type GetStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the resource name of the store to get.
	// Format: stores/{store}
	// Required.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetStoreRequest) Reset() {
	*x = GetStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStoreRequest) ProtoMessage() {}

func (x *GetStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStoreRequest.ProtoReflect.Descriptor instead.
func (*GetStoreRequest) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{7}
}

func (x *GetStoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ListStoresRequest is the request message for ListStores.
//
// (-- api-linter: core::0132::request-parent-required=disabled
//     aip.dev/not-precedent: Stores are top-level resources. --)
type ListStoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page_size is the maximum number of stores to return.
	// If unspecified, the server will choose a suitable number.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// page_token contains an opaque string used to get the next page of
	// results. It is usually provided by the previous call to ListStores.
	// If unspecified, the first page will be returned.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListStoresRequest) Reset() {
	*x = ListStoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStoresRequest) ProtoMessage() {}

func (x *ListStoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStoresRequest.ProtoReflect.Descriptor instead.
func (*ListStoresRequest) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{8}
}

func (x *ListStoresRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListStoresRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListStoresResponse is the response message for ListStores.
type ListStoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// stores contains the page of stores.
	Stores []*Store `protobuf:"bytes,1,rep,name=stores,proto3" json:"stores,omitempty"`
	// next_page_token contains an opaque string used to get the next page of
	// results. Provide this in a subsequent call to ListStores.
	// If this field is empty, there are no more pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListStoresResponse) Reset() {
	*x = ListStoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStoresResponse) ProtoMessage() {}

func (x *ListStoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStoresResponse.ProtoReflect.Descriptor instead.
func (*ListStoresResponse) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{9}
}

func (x *ListStoresResponse) GetStores() []*Store {
	if x != nil {
		return x.Stores
	}
	return nil
}

func (x *ListStoresResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// CreateStoreRequest is the request message for CreateStore.
//
// (-- api-linter: core::0133::request-parent-required=disabled
//     aip.dev/not-precedent: Stores are top-level resources. --)
type CreateStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store is the store to be created.
	// Required.
	Store *Store `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
}

func (x *CreateStoreRequest) Reset() {
	*x = CreateStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saser_strecku_v1_strecku_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStoreRequest) ProtoMessage() {}

func (x *CreateStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saser_strecku_v1_strecku_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStoreRequest.ProtoReflect.Descriptor instead.
func (*CreateStoreRequest) Descriptor() ([]byte, []int) {
	return file_saser_strecku_v1_strecku_proto_rawDescGZIP(), []int{10}
}

func (x *CreateStoreRequest) GetStore() *Store {
	if x != nil {
		return x.Store
	}
	return nil
}

var File_saser_strecku_v1_strecku_proto protoreflect.FileDescriptor

var file_saser_strecku_v1_strecku_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e,
	0x76, 0x31, 0x22, 0x80, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x69, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6d, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x32, 0xde,
	0x03, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x55, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x54, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x73,
	0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x73,
	0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x61,
	0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x24, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42,
	0x49, 0x0a, 0x13, 0x73, 0x65, 0x2e, 0x73, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x63, 0x6b, 0x75, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x55, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x63, 0x6b, 0x75,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_saser_strecku_v1_strecku_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_saser_strecku_v1_strecku_proto_goTypes = []interface{}{
	(*User)(nil),               // 0: saser.strecku.v1.User
	(*Store)(nil),              // 1: saser.strecku.v1.Store
	(*Membership)(nil),         // 2: saser.strecku.v1.Membership
	(*GetUserRequest)(nil),     // 3: saser.strecku.v1.GetUserRequest
	(*ListUsersRequest)(nil),   // 4: saser.strecku.v1.ListUsersRequest
	(*ListUsersResponse)(nil),  // 5: saser.strecku.v1.ListUsersResponse
	(*CreateUserRequest)(nil),  // 6: saser.strecku.v1.CreateUserRequest
	(*GetStoreRequest)(nil),    // 7: saser.strecku.v1.GetStoreRequest
	(*ListStoresRequest)(nil),  // 8: saser.strecku.v1.ListStoresRequest
	(*ListStoresResponse)(nil), // 9: saser.strecku.v1.ListStoresResponse
	(*CreateStoreRequest)(nil), // 10: saser.strecku.v1.CreateStoreRequest
}
var file_saser_strecku_v1_strecku_proto_depIdxs = []int32{
	0,  // 0: saser.strecku.v1.ListUsersResponse.users:type_name -> saser.strecku.v1.User
	0,  // 1: saser.strecku.v1.CreateUserRequest.user:type_name -> saser.strecku.v1.User
	1,  // 2: saser.strecku.v1.ListStoresResponse.stores:type_name -> saser.strecku.v1.Store
	1,  // 3: saser.strecku.v1.CreateStoreRequest.store:type_name -> saser.strecku.v1.Store
	3,  // 4: saser.strecku.v1.StreckU.GetUser:input_type -> saser.strecku.v1.GetUserRequest
	4,  // 5: saser.strecku.v1.StreckU.ListUsers:input_type -> saser.strecku.v1.ListUsersRequest
	6,  // 6: saser.strecku.v1.StreckU.CreateUser:input_type -> saser.strecku.v1.CreateUserRequest
	7,  // 7: saser.strecku.v1.StreckU.GetStore:input_type -> saser.strecku.v1.GetStoreRequest
	8,  // 8: saser.strecku.v1.StreckU.ListStores:input_type -> saser.strecku.v1.ListStoresRequest
	10, // 9: saser.strecku.v1.StreckU.CreateStore:input_type -> saser.strecku.v1.CreateStoreRequest
	0,  // 10: saser.strecku.v1.StreckU.GetUser:output_type -> saser.strecku.v1.User
	5,  // 11: saser.strecku.v1.StreckU.ListUsers:output_type -> saser.strecku.v1.ListUsersResponse
	0,  // 12: saser.strecku.v1.StreckU.CreateUser:output_type -> saser.strecku.v1.User
	1,  // 13: saser.strecku.v1.StreckU.GetStore:output_type -> saser.strecku.v1.Store
	9,  // 14: saser.strecku.v1.StreckU.ListStores:output_type -> saser.strecku.v1.ListStoresResponse
	1,  // 15: saser.strecku.v1.StreckU.CreateStore:output_type -> saser.strecku.v1.Store
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_saser_strecku_v1_strecku_proto_init() }
func file_saser_strecku_v1_strecku_proto_init() {
	if File_saser_strecku_v1_strecku_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_saser_strecku_v1_strecku_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			switch v := v.(*Store); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Membership); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResponse); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStoreRequest); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStoresRequest); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStoresResponse); i {
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
		file_saser_strecku_v1_strecku_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStoreRequest); i {
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
			NumMessages:   11,
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
