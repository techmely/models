// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: account/v1/account.service.proto

package accountv1

import (
	account "github.com/techmely/models/account"
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

var File_account_v1_account_service_proto protoreflect.FileDescriptor

var file_account_v1_account_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x66,
	0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x8b, 0x08, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x58, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x12, 0x20, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x70, 0x0a, 0x12, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x20, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x07, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x16, 0x52, 0x65,
	0x73, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6e,
	0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x2e,
	0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x25, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x14, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x2e, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x67,
	0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc4,
	0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x63, 0x68,
	0x6d, 0x65, 0x6c, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x47, 0x47, 0x41, 0xaa, 0x02, 0x11, 0x47, 0x65, 0x6e, 0x2e, 0x47, 0x6f, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x47, 0x65, 0x6e,
	0x5c, 0x47, 0x6f, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1d, 0x47, 0x65, 0x6e, 0x5c, 0x47, 0x6f, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x14, 0x47, 0x65, 0x6e, 0x3a, 0x3a, 0x47, 0x6f, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_account_v1_account_service_proto_goTypes = []any{
	(*SignInRequest)(nil),                      // 0: gen.go.account.v1.SignInRequest
	(*SignInWithProviderRequest)(nil),          // 1: gen.go.account.v1.SignInWithProviderRequest
	(*SignUpRequest)(nil),                      // 2: gen.go.account.v1.SignUpRequest
	(*emptypb.Empty)(nil),                      // 3: google.protobuf.Empty
	(*ResendVerificationCodeRequest)(nil),      // 4: gen.go.account.v1.ResendVerificationCodeRequest
	(*UpdatePasswordRequest)(nil),              // 5: gen.go.account.v1.UpdatePasswordRequest
	(*UpdateEmailRequest)(nil),                 // 6: gen.go.account.v1.UpdateEmailRequest
	(*VerifyAccountRequest)(nil),               // 7: gen.go.account.v1.VerifyAccountRequest
	(*VerifyActivationLinkRequest)(nil),        // 8: gen.go.account.v1.VerifyActivationLinkRequest
	(*ForgotPasswordRequest)(nil),              // 9: gen.go.account.v1.ForgotPasswordRequest
	(*account.AuthGoogleIdentityResponse)(nil), // 10: gen.go.account.AuthGoogleIdentityResponse
	(*ResendVerificationCodeResponse)(nil),     // 11: gen.go.account.v1.ResendVerificationCodeResponse
	(*UpdatePasswordResponse)(nil),             // 12: gen.go.account.v1.UpdatePasswordResponse
	(*UpdateEmailResponse)(nil),                // 13: gen.go.account.v1.UpdateEmailResponse
	(*VerifyAccountResponse)(nil),              // 14: gen.go.account.v1.VerifyAccountResponse
	(*VerifyActivationLinkResponse)(nil),       // 15: gen.go.account.v1.VerifyActivationLinkResponse
	(*ForgotPasswordResponse)(nil),             // 16: gen.go.account.v1.ForgotPasswordResponse
}
var file_account_v1_account_service_proto_depIdxs = []int32{
	0,  // 0: gen.go.account.v1.AccountServicePort.SignIn:input_type -> gen.go.account.v1.SignInRequest
	1,  // 1: gen.go.account.v1.AccountServicePort.SignInWithProvider:input_type -> gen.go.account.v1.SignInWithProviderRequest
	2,  // 2: gen.go.account.v1.AccountServicePort.SignUp:input_type -> gen.go.account.v1.SignUpRequest
	3,  // 3: gen.go.account.v1.AccountServicePort.SignOut:input_type -> google.protobuf.Empty
	4,  // 4: gen.go.account.v1.AccountServicePort.ResendVerificationCode:input_type -> gen.go.account.v1.ResendVerificationCodeRequest
	5,  // 5: gen.go.account.v1.AccountServicePort.UpdatePassword:input_type -> gen.go.account.v1.UpdatePasswordRequest
	6,  // 6: gen.go.account.v1.AccountServicePort.UpdateEmail:input_type -> gen.go.account.v1.UpdateEmailRequest
	7,  // 7: gen.go.account.v1.AccountServicePort.VerifyAccount:input_type -> gen.go.account.v1.VerifyAccountRequest
	8,  // 8: gen.go.account.v1.AccountServicePort.VerifyActivationLink:input_type -> gen.go.account.v1.VerifyActivationLinkRequest
	9,  // 9: gen.go.account.v1.AccountServicePort.ForgotPassword:input_type -> gen.go.account.v1.ForgotPasswordRequest
	10, // 10: gen.go.account.v1.AccountServicePort.SignIn:output_type -> gen.go.account.AuthGoogleIdentityResponse
	10, // 11: gen.go.account.v1.AccountServicePort.SignInWithProvider:output_type -> gen.go.account.AuthGoogleIdentityResponse
	10, // 12: gen.go.account.v1.AccountServicePort.SignUp:output_type -> gen.go.account.AuthGoogleIdentityResponse
	3,  // 13: gen.go.account.v1.AccountServicePort.SignOut:output_type -> google.protobuf.Empty
	11, // 14: gen.go.account.v1.AccountServicePort.ResendVerificationCode:output_type -> gen.go.account.v1.ResendVerificationCodeResponse
	12, // 15: gen.go.account.v1.AccountServicePort.UpdatePassword:output_type -> gen.go.account.v1.UpdatePasswordResponse
	13, // 16: gen.go.account.v1.AccountServicePort.UpdateEmail:output_type -> gen.go.account.v1.UpdateEmailResponse
	14, // 17: gen.go.account.v1.AccountServicePort.VerifyAccount:output_type -> gen.go.account.v1.VerifyAccountResponse
	15, // 18: gen.go.account.v1.AccountServicePort.VerifyActivationLink:output_type -> gen.go.account.v1.VerifyActivationLinkResponse
	16, // 19: gen.go.account.v1.AccountServicePort.ForgotPassword:output_type -> gen.go.account.v1.ForgotPasswordResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_account_v1_account_service_proto_init() }
func file_account_v1_account_service_proto_init() {
	if File_account_v1_account_service_proto != nil {
		return
	}
	file_account_v1_account_event_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_v1_account_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_v1_account_service_proto_goTypes,
		DependencyIndexes: file_account_v1_account_service_proto_depIdxs,
	}.Build()
	File_account_v1_account_service_proto = out.File
	file_account_v1_account_service_proto_rawDesc = nil
	file_account_v1_account_service_proto_goTypes = nil
	file_account_v1_account_service_proto_depIdxs = nil
}
