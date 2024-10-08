// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.180.0
//   protoc               unknown
// source: tenant/v1/tenant.model.proto

/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Value } from "../../google/protobuf/struct";

export const protobufPackage = "gen.go.tenant.v1";

export interface TenantTable {
  id: string;
  name: string;
  slug: string;
  description: string;
  isVerified: boolean;
  ownerId: string;
  metadata: any | undefined;
  createdAt: string;
  updatedAt: string;
}

function createBaseTenantTable(): TenantTable {
  return {
    id: "",
    name: "",
    slug: "",
    description: "",
    isVerified: false,
    ownerId: "",
    metadata: undefined,
    createdAt: "",
    updatedAt: "",
  };
}

export const TenantTable = {
  encode(message: TenantTable, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.slug !== "") {
      writer.uint32(26).string(message.slug);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.isVerified !== false) {
      writer.uint32(40).bool(message.isVerified);
    }
    if (message.ownerId !== "") {
      writer.uint32(50).string(message.ownerId);
    }
    if (message.metadata !== undefined) {
      Value.encode(Value.wrap(message.metadata), writer.uint32(58).fork()).ldelim();
    }
    if (message.createdAt !== "") {
      writer.uint32(66).string(message.createdAt);
    }
    if (message.updatedAt !== "") {
      writer.uint32(74).string(message.updatedAt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TenantTable {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTenantTable();
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

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.slug = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.description = reader.string();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.isVerified = reader.bool();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.ownerId = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.metadata = Value.unwrap(Value.decode(reader, reader.uint32()));
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.createdAt = reader.string();
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.updatedAt = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TenantTable {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      slug: isSet(object.slug) ? globalThis.String(object.slug) : "",
      description: isSet(object.description) ? globalThis.String(object.description) : "",
      isVerified: isSet(object.isVerified) ? globalThis.Boolean(object.isVerified) : false,
      ownerId: isSet(object.ownerId) ? globalThis.String(object.ownerId) : "",
      metadata: isSet(object?.metadata) ? object.metadata : undefined,
      createdAt: isSet(object.createdAt) ? globalThis.String(object.createdAt) : "",
      updatedAt: isSet(object.updatedAt) ? globalThis.String(object.updatedAt) : "",
    };
  },

  toJSON(message: TenantTable): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.slug !== "") {
      obj.slug = message.slug;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.isVerified !== false) {
      obj.isVerified = message.isVerified;
    }
    if (message.ownerId !== "") {
      obj.ownerId = message.ownerId;
    }
    if (message.metadata !== undefined) {
      obj.metadata = message.metadata;
    }
    if (message.createdAt !== "") {
      obj.createdAt = message.createdAt;
    }
    if (message.updatedAt !== "") {
      obj.updatedAt = message.updatedAt;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TenantTable>, I>>(base?: I): TenantTable {
    return TenantTable.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TenantTable>, I>>(object: I): TenantTable {
    const message = createBaseTenantTable();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.slug = object.slug ?? "";
    message.description = object.description ?? "";
    message.isVerified = object.isVerified ?? false;
    message.ownerId = object.ownerId ?? "";
    message.metadata = object.metadata ?? undefined;
    message.createdAt = object.createdAt ?? "";
    message.updatedAt = object.updatedAt ?? "";
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
