// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: user/v1/user.event.proto

package userv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string       `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Nickname        string       `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Status          UserStatus   `protobuf:"varint,4,opt,name=status,proto3,enum=gen.go.user.v1.UserStatus" json:"status,omitempty"`
	IsEmailVerified bool         `protobuf:"varint,5,opt,name=is_email_verified,json=isEmailVerified,proto3" json:"is_email_verified,omitempty"`
	Name            string       `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	AvatarUrl       string       `protobuf:"bytes,7,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	FirebaseUserId  string       `protobuf:"bytes,8,opt,name=firebase_user_id,json=firebaseUserId,proto3" json:"firebase_user_id,omitempty"`
	AuthStrategy    AuthStrategy `protobuf:"varint,9,opt,name=auth_strategy,json=authStrategy,proto3,enum=gen.go.user.v1.AuthStrategy" json:"auth_strategy,omitempty"`
	OpenPlatform    string       `protobuf:"bytes,10,opt,name=open_platform,json=openPlatform,proto3" json:"open_platform,omitempty"`
	UtmCampaign     string       `protobuf:"bytes,11,opt,name=utm_campaign,json=utmCampaign,proto3" json:"utm_campaign,omitempty"`
	UtmMedium       string       `protobuf:"bytes,12,opt,name=utm_medium,json=utmMedium,proto3" json:"utm_medium,omitempty"`
	UtmSource       string       `protobuf:"bytes,13,opt,name=utm_source,json=utmSource,proto3" json:"utm_source,omitempty"`
	Metadata        *anypb.Any   `protobuf:"bytes,14,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_event_proto_msgTypes[0]
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
	return file_user_v1_user_event_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateUserRequest) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_INACTIVE
}

func (x *CreateUserRequest) GetIsEmailVerified() bool {
	if x != nil {
		return x.IsEmailVerified
	}
	return false
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *CreateUserRequest) GetFirebaseUserId() string {
	if x != nil {
		return x.FirebaseUserId
	}
	return ""
}

func (x *CreateUserRequest) GetAuthStrategy() AuthStrategy {
	if x != nil {
		return x.AuthStrategy
	}
	return AuthStrategy_BASIC
}

func (x *CreateUserRequest) GetOpenPlatform() string {
	if x != nil {
		return x.OpenPlatform
	}
	return ""
}

func (x *CreateUserRequest) GetUtmCampaign() string {
	if x != nil {
		return x.UtmCampaign
	}
	return ""
}

func (x *CreateUserRequest) GetUtmMedium() string {
	if x != nil {
		return x.UtmMedium
	}
	return ""
}

func (x *CreateUserRequest) GetUtmSource() string {
	if x != nil {
		return x.UtmSource
	}
	return ""
}

func (x *CreateUserRequest) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_event_proto_msgTypes[1]
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
	return file_user_v1_user_event_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetUserRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetUserByAuthIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByAuthIdRequest) Reset() {
	*x = GetUserByAuthIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByAuthIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByAuthIdRequest) ProtoMessage() {}

func (x *GetUserByAuthIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByAuthIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByAuthIdRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_event_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByAuthIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUsersPaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	Limit    uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page     uint64 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetUsersPaginationRequest) Reset() {
	*x = GetUsersPaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersPaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersPaginationRequest) ProtoMessage() {}

func (x *GetUsersPaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersPaginationRequest.ProtoReflect.Descriptor instead.
func (*GetUsersPaginationRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_event_proto_rawDescGZIP(), []int{3}
}

func (x *GetUsersPaginationRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *GetUsersPaginationRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUsersPaginationRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetUsersPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUsersPaginationResponse) Reset() {
	*x = GetUsersPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersPaginationResponse) ProtoMessage() {}

func (x *GetUsersPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersPaginationResponse.ProtoReflect.Descriptor instead.
func (*GetUsersPaginationResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_event_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersPaginationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_event_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_event_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_user_v1_user_event_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_user_v1_user_event_proto protoreflect.FileDescriptor

var file_user_v1_user_event_proto_rawDesc = []byte{
	0x0a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x65, 0x6e, 0x2e,
	0x67, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfd, 0x03, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x69,
	0x73, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69,
	0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x70, 0x65, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x74, 0x6d, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x75, 0x74, 0x6d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x74, 0x6d, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x74, 0x6d, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x74, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x74, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0xaa, 0x01, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x65, 0x63, 0x68, 0x6d, 0x65, 0x6c, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x47, 0x47, 0x55, 0xaa, 0x02, 0x0e, 0x47, 0x65, 0x6e, 0x2e, 0x47, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x47, 0x65, 0x6e, 0x5c, 0x47, 0x6f, 0x5c, 0x55,
	0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x47, 0x65, 0x6e, 0x5c, 0x47, 0x6f, 0x5c,
	0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x47, 0x65, 0x6e, 0x3a, 0x3a, 0x47, 0x6f, 0x3a, 0x3a, 0x55,
	0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_event_proto_rawDescOnce sync.Once
	file_user_v1_user_event_proto_rawDescData = file_user_v1_user_event_proto_rawDesc
)

func file_user_v1_user_event_proto_rawDescGZIP() []byte {
	file_user_v1_user_event_proto_rawDescOnce.Do(func() {
		file_user_v1_user_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_event_proto_rawDescData)
	})
	return file_user_v1_user_event_proto_rawDescData
}

var file_user_v1_user_event_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_v1_user_event_proto_goTypes = []any{
	(*CreateUserRequest)(nil),          // 0: gen.go.user.v1.CreateUserRequest
	(*GetUserRequest)(nil),             // 1: gen.go.user.v1.GetUserRequest
	(*GetUserByAuthIdRequest)(nil),     // 2: gen.go.user.v1.GetUserByAuthIdRequest
	(*GetUsersPaginationRequest)(nil),  // 3: gen.go.user.v1.GetUsersPaginationRequest
	(*GetUsersPaginationResponse)(nil), // 4: gen.go.user.v1.GetUsersPaginationResponse
	(*UpdateUserRequest)(nil),          // 5: gen.go.user.v1.UpdateUserRequest
	(*DeleteUserRequest)(nil),          // 6: gen.go.user.v1.DeleteUserRequest
	(*DeleteUserResponse)(nil),         // 7: gen.go.user.v1.DeleteUserResponse
	(UserStatus)(0),                    // 8: gen.go.user.v1.UserStatus
	(AuthStrategy)(0),                  // 9: gen.go.user.v1.AuthStrategy
	(*anypb.Any)(nil),                  // 10: google.protobuf.Any
}
var file_user_v1_user_event_proto_depIdxs = []int32{
	8,  // 0: gen.go.user.v1.CreateUserRequest.status:type_name -> gen.go.user.v1.UserStatus
	9,  // 1: gen.go.user.v1.CreateUserRequest.auth_strategy:type_name -> gen.go.user.v1.AuthStrategy
	10, // 2: gen.go.user.v1.CreateUserRequest.metadata:type_name -> google.protobuf.Any
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_user_v1_user_event_proto_init() }
func file_user_v1_user_event_proto_init() {
	if File_user_v1_user_event_proto != nil {
		return
	}
	file_user_v1_user_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_v1_user_event_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_user_v1_user_event_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_user_v1_user_event_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserByAuthIdRequest); i {
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
		file_user_v1_user_event_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetUsersPaginationRequest); i {
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
		file_user_v1_user_event_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetUsersPaginationResponse); i {
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
		file_user_v1_user_event_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserRequest); i {
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
		file_user_v1_user_event_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserRequest); i {
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
		file_user_v1_user_event_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserResponse); i {
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
			RawDescriptor: file_user_v1_user_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_v1_user_event_proto_goTypes,
		DependencyIndexes: file_user_v1_user_event_proto_depIdxs,
		MessageInfos:      file_user_v1_user_event_proto_msgTypes,
	}.Build()
	File_user_v1_user_event_proto = out.File
	file_user_v1_user_event_proto_rawDesc = nil
	file_user_v1_user_event_proto_goTypes = nil
	file_user_v1_user_event_proto_depIdxs = nil
}
