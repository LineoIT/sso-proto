// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var audit_pb = require('./audit_pb.js');
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

function serialize_proto_AuditLog(arg) {
  if (!(arg instanceof audit_pb.AuditLog)) {
    throw new Error('Expected argument of type proto.AuditLog');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_AuditLog(buffer_arg) {
  return audit_pb.AuditLog.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AuditLogFilter(arg) {
  if (!(arg instanceof audit_pb.AuditLogFilter)) {
    throw new Error('Expected argument of type proto.AuditLogFilter');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_AuditLogFilter(buffer_arg) {
  return audit_pb.AuditLogFilter.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AuditLogResponse(arg) {
  if (!(arg instanceof audit_pb.AuditLogResponse)) {
    throw new Error('Expected argument of type proto.AuditLogResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_AuditLogResponse(buffer_arg) {
  return audit_pb.AuditLogResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateAuditLogResponse(arg) {
  if (!(arg instanceof audit_pb.CreateAuditLogResponse)) {
    throw new Error('Expected argument of type proto.CreateAuditLogResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_CreateAuditLogResponse(buffer_arg) {
  return audit_pb.CreateAuditLogResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuditServiceService = exports.AuditServiceService = {
  createAuditLog: {
    path: '/proto.AuditService/CreateAuditLog',
    requestStream: false,
    responseStream: false,
    requestType: audit_pb.AuditLog,
    responseType: audit_pb.CreateAuditLogResponse,
    requestSerialize: serialize_proto_AuditLog,
    requestDeserialize: deserialize_proto_AuditLog,
    responseSerialize: serialize_proto_CreateAuditLogResponse,
    responseDeserialize: deserialize_proto_CreateAuditLogResponse,
  },
  updateAuditLog: {
    path: '/proto.AuditService/UpdateAuditLog',
    requestStream: false,
    responseStream: false,
    requestType: audit_pb.AuditLog,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_proto_AuditLog,
    requestDeserialize: deserialize_proto_AuditLog,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  listAuditLog: {
    path: '/proto.AuditService/ListAuditLog',
    requestStream: false,
    responseStream: false,
    requestType: audit_pb.AuditLogFilter,
    responseType: audit_pb.AuditLogResponse,
    requestSerialize: serialize_proto_AuditLogFilter,
    requestDeserialize: deserialize_proto_AuditLogFilter,
    responseSerialize: serialize_proto_AuditLogResponse,
    responseDeserialize: deserialize_proto_AuditLogResponse,
  },
  getAuditLog: {
    path: '/proto.AuditService/GetAuditLog',
    requestStream: false,
    responseStream: false,
    requestType: audit_pb.AuditLogFilter,
    responseType: audit_pb.AuditLog,
    requestSerialize: serialize_proto_AuditLogFilter,
    requestDeserialize: deserialize_proto_AuditLogFilter,
    responseSerialize: serialize_proto_AuditLog,
    responseDeserialize: deserialize_proto_AuditLog,
  },
};

exports.AuditServiceClient = grpc.makeGenericClientConstructor(AuditServiceService);
