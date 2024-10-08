// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.180.0
//   protoc               unknown
// source: file/v1/file.model.proto

/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Value } from "../../google/protobuf/struct";

export const protobufPackage = "gen.go.file.v1";

export interface FileModel {
  id: string;
  title: string;
  description: string;
  contentType: string;
  bucket: string;
  fileName: string;
  filePath: string;
  fileUrl: string;
  metadata: any | undefined;
  userId?: string | undefined;
}

function createBaseFileModel(): FileModel {
  return {
    id: "",
    title: "",
    description: "",
    contentType: "",
    bucket: "",
    fileName: "",
    filePath: "",
    fileUrl: "",
    metadata: undefined,
    userId: undefined,
  };
}

export const FileModel = {
  encode(message: FileModel, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.contentType !== "") {
      writer.uint32(34).string(message.contentType);
    }
    if (message.bucket !== "") {
      writer.uint32(42).string(message.bucket);
    }
    if (message.fileName !== "") {
      writer.uint32(50).string(message.fileName);
    }
    if (message.filePath !== "") {
      writer.uint32(58).string(message.filePath);
    }
    if (message.fileUrl !== "") {
      writer.uint32(66).string(message.fileUrl);
    }
    if (message.metadata !== undefined) {
      Value.encode(Value.wrap(message.metadata), writer.uint32(74).fork()).ldelim();
    }
    if (message.userId !== undefined) {
      writer.uint32(82).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FileModel {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFileModel();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.title = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.description = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.contentType = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.bucket = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.fileName = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.filePath = reader.string();
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.fileUrl = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.metadata = Value.unwrap(Value.decode(reader, reader.uint32()));
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.userId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FileModel {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      title: isSet(object.title) ? globalThis.String(object.title) : "",
      description: isSet(object.description) ? globalThis.String(object.description) : "",
      contentType: isSet(object.contentType) ? globalThis.String(object.contentType) : "",
      bucket: isSet(object.bucket) ? globalThis.String(object.bucket) : "",
      fileName: isSet(object.fileName) ? globalThis.String(object.fileName) : "",
      filePath: isSet(object.filePath) ? globalThis.String(object.filePath) : "",
      fileUrl: isSet(object.fileUrl) ? globalThis.String(object.fileUrl) : "",
      metadata: isSet(object?.metadata) ? object.metadata : undefined,
      userId: isSet(object.userId) ? globalThis.String(object.userId) : undefined,
    };
  },

  toJSON(message: FileModel): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.contentType !== "") {
      obj.contentType = message.contentType;
    }
    if (message.bucket !== "") {
      obj.bucket = message.bucket;
    }
    if (message.fileName !== "") {
      obj.fileName = message.fileName;
    }
    if (message.filePath !== "") {
      obj.filePath = message.filePath;
    }
    if (message.fileUrl !== "") {
      obj.fileUrl = message.fileUrl;
    }
    if (message.metadata !== undefined) {
      obj.metadata = message.metadata;
    }
    if (message.userId !== undefined) {
      obj.userId = message.userId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<FileModel>, I>>(base?: I): FileModel {
    return FileModel.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<FileModel>, I>>(object: I): FileModel {
    const message = createBaseFileModel();
    message.id = object.id ?? "";
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.contentType = object.contentType ?? "";
    message.bucket = object.bucket ?? "";
    message.fileName = object.fileName ?? "";
    message.filePath = object.filePath ?? "";
    message.fileUrl = object.fileUrl ?? "";
    message.metadata = object.metadata ?? undefined;
    message.userId = object.userId ?? undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
