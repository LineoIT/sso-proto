// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: account.service.proto

package sso_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_account_service_proto protoreflect.FileDescriptor

var file_account_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc2, 0x04, 0x0a, 0x0e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x44, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x06, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x72,
	0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x13, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c,
	0x69, 0x6e, 0x65, 0x6f, 0x49, 0x54, 0x2f, 0x73, 0x73, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_account_service_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),       // 0: proto.CreateAccountRequest
	(*UpdateFullnameRequest)(nil),      // 1: proto.UpdateFullnameRequest
	(*EnableAccountRequest)(nil),       // 2: proto.EnableAccountRequest
	(*AccountFilter)(nil),              // 3: proto.AccountFilter
	(*LoginRequest)(nil),               // 4: proto.LoginRequest
	(*ForgotPasswordRequest)(nil),      // 5: proto.ForgotPasswordRequest
	(*EnableResetPasswordRequest)(nil), // 6: proto.EnableResetPasswordRequest
	(*ResetPasswordRequest)(nil),       // 7: proto.ResetPasswordRequest
	(*ID)(nil),                         // 8: proto.ID
	(*Empty)(nil),                      // 9: proto.Empty
	(*AccountResponse)(nil),            // 10: proto.AccountResponse
	(*Account)(nil),                    // 11: proto.Account
	(*AccessToken)(nil),                // 12: proto.AccessToken
}
var file_account_service_proto_depIdxs = []int32{
	0,  // 0: proto.AccountService.CreateAccount:input_type -> proto.CreateAccountRequest
	1,  // 1: proto.AccountService.UpdateFullname:input_type -> proto.UpdateFullnameRequest
	2,  // 2: proto.AccountService.EnableAccount:input_type -> proto.EnableAccountRequest
	3,  // 3: proto.AccountService.ListAccounts:input_type -> proto.AccountFilter
	3,  // 4: proto.AccountService.GetAccount:input_type -> proto.AccountFilter
	4,  // 5: proto.AccountService.Signin:input_type -> proto.LoginRequest
	5,  // 6: proto.AccountService.ForgotPassword:input_type -> proto.ForgotPasswordRequest
	6,  // 7: proto.AccountService.EnableResetPassword:input_type -> proto.EnableResetPasswordRequest
	7,  // 8: proto.AccountService.ResetPassword:input_type -> proto.ResetPasswordRequest
	8,  // 9: proto.AccountService.CreateAccount:output_type -> proto.ID
	9,  // 10: proto.AccountService.UpdateFullname:output_type -> proto.Empty
	9,  // 11: proto.AccountService.EnableAccount:output_type -> proto.Empty
	10, // 12: proto.AccountService.ListAccounts:output_type -> proto.AccountResponse
	11, // 13: proto.AccountService.GetAccount:output_type -> proto.Account
	12, // 14: proto.AccountService.Signin:output_type -> proto.AccessToken
	9,  // 15: proto.AccountService.ForgotPassword:output_type -> proto.Empty
	12, // 16: proto.AccountService.EnableResetPassword:output_type -> proto.AccessToken
	9,  // 17: proto.AccountService.ResetPassword:output_type -> proto.Empty
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_account_service_proto_init() }
func file_account_service_proto_init() {
	if File_account_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_account_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_service_proto_goTypes,
		DependencyIndexes: file_account_service_proto_depIdxs,
	}.Build()
	File_account_service_proto = out.File
	file_account_service_proto_rawDesc = nil
	file_account_service_proto_goTypes = nil
	file_account_service_proto_depIdxs = nil
}
