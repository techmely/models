// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: tenant/v1/tenant.event.proto

package tenantv1

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

type GetTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{0}
}

func (x *GetTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTenantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTenantsRequest) Reset() {
	*x = GetTenantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantsRequest) ProtoMessage() {}

func (x *GetTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantsRequest.ProtoReflect.Descriptor instead.
func (*GetTenantsRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{1}
}

func (x *GetTenantsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAvailableTenantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAvailableTenantsRequest) Reset() {
	*x = GetAvailableTenantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableTenantsRequest) ProtoMessage() {}

func (x *GetAvailableTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableTenantsRequest.ProtoReflect.Descriptor instead.
func (*GetAvailableTenantsRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{2}
}

func (x *GetAvailableTenantsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTenantRequest) Reset() {
	*x = CreateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantRequest) ProtoMessage() {}

func (x *CreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantRequest.ProtoReflect.Descriptor instead.
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateTenantRequest) Reset() {
	*x = UpdateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantRequest) ProtoMessage() {}

func (x *UpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTenantRequest) Reset() {
	*x = DeleteTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantRequest) ProtoMessage() {}

func (x *DeleteTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantRequest.ProtoReflect.Descriptor instead.
func (*DeleteTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTenantResponse) Reset() {
	*x = GetTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantResponse) ProtoMessage() {}

func (x *GetTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantResponse.ProtoReflect.Descriptor instead.
func (*GetTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{6}
}

type GetTenantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTenantsResponse) Reset() {
	*x = GetTenantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantsResponse) ProtoMessage() {}

func (x *GetTenantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantsResponse.ProtoReflect.Descriptor instead.
func (*GetTenantsResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{7}
}

type CreateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTenantResponse) Reset() {
	*x = CreateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantResponse) ProtoMessage() {}

func (x *CreateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantResponse.ProtoReflect.Descriptor instead.
func (*CreateTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{8}
}

type UpdateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTenantResponse) Reset() {
	*x = UpdateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantResponse) ProtoMessage() {}

func (x *UpdateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantResponse.ProtoReflect.Descriptor instead.
func (*UpdateTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{9}
}

type GetAvailableTenantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAvailableTenantsResponse) Reset() {
	*x = GetAvailableTenantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableTenantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableTenantsResponse) ProtoMessage() {}

func (x *GetAvailableTenantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableTenantsResponse.ProtoReflect.Descriptor instead.
func (*GetAvailableTenantsResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{10}
}

type DeleteTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTenantResponse) Reset() {
	*x = DeleteTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v1_tenant_event_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantResponse) ProtoMessage() {}

func (x *DeleteTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v1_tenant_event_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantResponse.ProtoReflect.Descriptor instead.
func (*DeleteTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v1_tenant_event_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteTenantResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_tenant_v1_tenant_event_proto protoreflect.FileDescriptor

var file_tenant_v1_tenant_event_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0xba,
	0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x6d, 0x65, 0x6c, 0x79,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x47, 0x54,
	0xaa, 0x02, 0x10, 0x47, 0x65, 0x6e, 0x2e, 0x47, 0x6f, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x47, 0x65, 0x6e, 0x5c, 0x47, 0x6f, 0x5c, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x47, 0x65, 0x6e, 0x5c, 0x47, 0x6f, 0x5c,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x47, 0x65, 0x6e, 0x3a, 0x3a, 0x47, 0x6f, 0x3a,
	0x3a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tenant_v1_tenant_event_proto_rawDescOnce sync.Once
	file_tenant_v1_tenant_event_proto_rawDescData = file_tenant_v1_tenant_event_proto_rawDesc
)

func file_tenant_v1_tenant_event_proto_rawDescGZIP() []byte {
	file_tenant_v1_tenant_event_proto_rawDescOnce.Do(func() {
		file_tenant_v1_tenant_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_tenant_v1_tenant_event_proto_rawDescData)
	})
	return file_tenant_v1_tenant_event_proto_rawDescData
}

var file_tenant_v1_tenant_event_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_tenant_v1_tenant_event_proto_goTypes = []any{
	(*GetTenantRequest)(nil),            // 0: gen.go.tenant.v1.GetTenantRequest
	(*GetTenantsRequest)(nil),           // 1: gen.go.tenant.v1.GetTenantsRequest
	(*GetAvailableTenantsRequest)(nil),  // 2: gen.go.tenant.v1.GetAvailableTenantsRequest
	(*CreateTenantRequest)(nil),         // 3: gen.go.tenant.v1.CreateTenantRequest
	(*UpdateTenantRequest)(nil),         // 4: gen.go.tenant.v1.UpdateTenantRequest
	(*DeleteTenantRequest)(nil),         // 5: gen.go.tenant.v1.DeleteTenantRequest
	(*GetTenantResponse)(nil),           // 6: gen.go.tenant.v1.GetTenantResponse
	(*GetTenantsResponse)(nil),          // 7: gen.go.tenant.v1.GetTenantsResponse
	(*CreateTenantResponse)(nil),        // 8: gen.go.tenant.v1.CreateTenantResponse
	(*UpdateTenantResponse)(nil),        // 9: gen.go.tenant.v1.UpdateTenantResponse
	(*GetAvailableTenantsResponse)(nil), // 10: gen.go.tenant.v1.GetAvailableTenantsResponse
	(*DeleteTenantResponse)(nil),        // 11: gen.go.tenant.v1.DeleteTenantResponse
}
var file_tenant_v1_tenant_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tenant_v1_tenant_event_proto_init() }
func file_tenant_v1_tenant_event_proto_init() {
	if File_tenant_v1_tenant_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tenant_v1_tenant_event_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetTenantRequest); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetTenantsRequest); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAvailableTenantsRequest); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTenantRequest); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTenantRequest); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTenantRequest); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetTenantResponse); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetTenantsResponse); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTenantResponse); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTenantResponse); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetAvailableTenantsResponse); i {
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
		file_tenant_v1_tenant_event_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTenantResponse); i {
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
			RawDescriptor: file_tenant_v1_tenant_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tenant_v1_tenant_event_proto_goTypes,
		DependencyIndexes: file_tenant_v1_tenant_event_proto_depIdxs,
		MessageInfos:      file_tenant_v1_tenant_event_proto_msgTypes,
	}.Build()
	File_tenant_v1_tenant_event_proto = out.File
	file_tenant_v1_tenant_event_proto_rawDesc = nil
	file_tenant_v1_tenant_event_proto_goTypes = nil
	file_tenant_v1_tenant_event_proto_depIdxs = nil
}