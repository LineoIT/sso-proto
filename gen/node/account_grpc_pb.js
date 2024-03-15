// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var account_pb = require('./account_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AccessToken(arg) {
  if (!(arg instanceof account_pb.AccessToken)) {
    throw new Error('Expected argument of type proto.AccessToken');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_AccessToken(buffer_arg) {
  return account_pb.AccessToken.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_Account(arg) {
  if (!(arg instanceof account_pb.Account)) {
    throw new Error('Expected argument of type proto.Account');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_Account(buffer_arg) {
  return account_pb.Account.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AccountFilter(arg) {
  if (!(arg instanceof account_pb.AccountFilter)) {
    throw new Error('Expected argument of type proto.AccountFilter');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_AccountFilter(buffer_arg) {
  return account_pb.AccountFilter.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AccountResponse(arg) {
  if (!(arg instanceof account_pb.AccountResponse)) {
    throw new Error('Expected argument of type proto.AccountResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_AccountResponse(buffer_arg) {
  return account_pb.AccountResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateAccountRequest(arg) {
  if (!(arg instanceof account_pb.CreateAccountRequest)) {
    throw new Error('Expected argument of type proto.CreateAccountRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_CreateAccountRequest(buffer_arg) {
  return account_pb.CreateAccountRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateAccountResponse(arg) {
  if (!(arg instanceof account_pb.CreateAccountResponse)) {
    throw new Error('Expected argument of type proto.CreateAccountResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_CreateAccountResponse(buffer_arg) {
  return account_pb.CreateAccountResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_EnableAccountRequest(arg) {
  if (!(arg instanceof account_pb.EnableAccountRequest)) {
    throw new Error('Expected argument of type proto.EnableAccountRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_EnableAccountRequest(buffer_arg) {
  return account_pb.EnableAccountRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_EnableResetPasswordRequest(arg) {
  if (!(arg instanceof account_pb.EnableResetPasswordRequest)) {
    throw new Error('Expected argument of type proto.EnableResetPasswordRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_EnableResetPasswordRequest(buffer_arg) {
  return account_pb.EnableResetPasswordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_ForgotPasswordRequest(arg) {
  if (!(arg instanceof account_pb.ForgotPasswordRequest)) {
    throw new Error('Expected argument of type proto.ForgotPasswordRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_ForgotPasswordRequest(buffer_arg) {
  return account_pb.ForgotPasswordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_LoginRequest(arg) {
  if (!(arg instanceof account_pb.LoginRequest)) {
    throw new Error('Expected argument of type proto.LoginRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_LoginRequest(buffer_arg) {
  return account_pb.LoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_ResetPasswordRequest(arg) {
  if (!(arg instanceof account_pb.ResetPasswordRequest)) {
    throw new Error('Expected argument of type proto.ResetPasswordRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_ResetPasswordRequest(buffer_arg) {
  return account_pb.ResetPasswordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateFullnameRequest(arg) {
  if (!(arg instanceof account_pb.UpdateFullnameRequest)) {
    throw new Error('Expected argument of type proto.UpdateFullnameRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_UpdateFullnameRequest(buffer_arg) {
  return account_pb.UpdateFullnameRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var AccountServiceService = exports.AccountServiceService = {
  createAccount: {
    path: '/proto.AccountService/CreateAccount',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.CreateAccountRequest,
    responseType: account_pb.CreateAccountResponse,
    requestSerialize: serialize_proto_CreateAccountRequest,
    requestDeserialize: deserialize_proto_CreateAccountRequest,
    responseSerialize: serialize_proto_CreateAccountResponse,
    responseDeserialize: deserialize_proto_CreateAccountResponse,
  },
  updateFullname: {
    path: '/proto.AccountService/UpdateFullname',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.UpdateFullnameRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_proto_UpdateFullnameRequest,
    requestDeserialize: deserialize_proto_UpdateFullnameRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  enableAccount: {
    path: '/proto.AccountService/EnableAccount',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.EnableAccountRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_proto_EnableAccountRequest,
    requestDeserialize: deserialize_proto_EnableAccountRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  listAccounts: {
    path: '/proto.AccountService/ListAccounts',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.AccountFilter,
    responseType: account_pb.AccountResponse,
    requestSerialize: serialize_proto_AccountFilter,
    requestDeserialize: deserialize_proto_AccountFilter,
    responseSerialize: serialize_proto_AccountResponse,
    responseDeserialize: deserialize_proto_AccountResponse,
  },
  getAccount: {
    path: '/proto.AccountService/GetAccount',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.AccountFilter,
    responseType: account_pb.Account,
    requestSerialize: serialize_proto_AccountFilter,
    requestDeserialize: deserialize_proto_AccountFilter,
    responseSerialize: serialize_proto_Account,
    responseDeserialize: deserialize_proto_Account,
  },
  signin: {
    path: '/proto.AccountService/Signin',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.LoginRequest,
    responseType: account_pb.AccessToken,
    requestSerialize: serialize_proto_LoginRequest,
    requestDeserialize: deserialize_proto_LoginRequest,
    responseSerialize: serialize_proto_AccessToken,
    responseDeserialize: deserialize_proto_AccessToken,
  },
  forgotPassword: {
    path: '/proto.AccountService/ForgotPassword',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.ForgotPasswordRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_proto_ForgotPasswordRequest,
    requestDeserialize: deserialize_proto_ForgotPasswordRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  enableResetPassword: {
    path: '/proto.AccountService/EnableResetPassword',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.EnableResetPasswordRequest,
    responseType: account_pb.AccessToken,
    requestSerialize: serialize_proto_EnableResetPasswordRequest,
    requestDeserialize: deserialize_proto_EnableResetPasswordRequest,
    responseSerialize: serialize_proto_AccessToken,
    responseDeserialize: deserialize_proto_AccessToken,
  },
  resetPassword: {
    path: '/proto.AccountService/ResetPassword',
    requestStream: false,
    responseStream: false,
    requestType: account_pb.ResetPasswordRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_proto_ResetPasswordRequest,
    requestDeserialize: deserialize_proto_ResetPasswordRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.AccountServiceClient = grpc.makeGenericClientConstructor(AccountServiceService);
