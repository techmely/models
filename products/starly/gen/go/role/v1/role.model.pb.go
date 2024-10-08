// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: role/v1/role.model.proto

package rolev1

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

type UserRoles int32

const (
	UserRoles_SUPER_ADMIN UserRoles = 0
	UserRoles_MODERATOR   UserRoles = 1
	UserRoles_ADMIN       UserRoles = 2
	UserRoles_MEMBER      UserRoles = 3
	UserRoles_GUEST       UserRoles = 4
)

// Enum value maps for UserRoles.
var (
	UserRoles_name = map[int32]string{
		0: "SUPER_ADMIN",
		1: "MODERATOR",
		2: "ADMIN",
		3: "MEMBER",
		4: "GUEST",
	}
	UserRoles_value = map[string]int32{
		"SUPER_ADMIN": 0,
		"MODERATOR":   1,
		"ADMIN":       2,
		"MEMBER":      3,
		"GUEST":       4,
	}
)

func (x UserRoles) Enum() *UserRoles {
	p := new(UserRoles)
	*p = x
	return p
}

func (x UserRoles) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRoles) Descriptor() protoreflect.EnumDescriptor {
	return file_role_v1_role_model_proto_enumTypes[0].Descriptor()
}

func (UserRoles) Type() protoreflect.EnumType {
	return &file_role_v1_role_model_proto_enumTypes[0]
}

func (x UserRoles) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRoles.Descriptor instead.
func (UserRoles) EnumDescriptor() ([]byte, []int) {
	return file_role_v1_role_model_proto_rawDescGZIP(), []int{0}
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_v1_role_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_role_v1_role_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_role_v1_role_model_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_role_v1_role_model_proto protoreflect.FileDescriptor

var file_role_v1_role_model_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x65, 0x6e, 0x2e,
	0x67, 0x6f, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x16, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x2a, 0x4d, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x55, 0x50, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x4f, 0x44, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45,
	0x4d, 0x42, 0x45, 0x52, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x04, 0x42, 0xaa, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x6d, 0x65, 0x6c, 0x79, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x72,
	0x6f, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x47, 0x52, 0xaa, 0x02, 0x0e, 0x47, 0x65,
	0x6e, 0x2e, 0x47, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x47,
	0x65, 0x6e, 0x5c, 0x47, 0x6f, 0x5c, 0x52, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a,
	0x47, 0x65, 0x6e, 0x5c, 0x47, 0x6f, 0x5c, 0x52, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x47, 0x65, 0x6e,
	0x3a, 0x3a, 0x47, 0x6f, 0x3a, 0x3a, 0x52, 0x6f, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_v1_role_model_proto_rawDescOnce sync.Once
	file_role_v1_role_model_proto_rawDescData = file_role_v1_role_model_proto_rawDesc
)

func file_role_v1_role_model_proto_rawDescGZIP() []byte {
	file_role_v1_role_model_proto_rawDescOnce.Do(func() {
		file_role_v1_role_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_v1_role_model_proto_rawDescData)
	})
	return file_role_v1_role_model_proto_rawDescData
}

var file_role_v1_role_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_role_v1_role_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_role_v1_role_model_proto_goTypes = []any{
	(UserRoles)(0), // 0: gen.go.role.v1.UserRoles
	(*Role)(nil),   // 1: gen.go.role.v1.Role
}
var file_role_v1_role_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_role_v1_role_model_proto_init() }
func file_role_v1_role_model_proto_init() {
	if File_role_v1_role_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_v1_role_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Role); i {
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
			RawDescriptor: file_role_v1_role_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_v1_role_model_proto_goTypes,
		DependencyIndexes: file_role_v1_role_model_proto_depIdxs,
		EnumInfos:         file_role_v1_role_model_proto_enumTypes,
		MessageInfos:      file_role_v1_role_model_proto_msgTypes,
	}.Build()
	File_role_v1_role_model_proto = out.File
	file_role_v1_role_model_proto_rawDesc = nil
	file_role_v1_role_model_proto_goTypes = nil
	file_role_v1_role_model_proto_depIdxs = nil
}
