// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.180.0
//   protoc               unknown
// source: tenant/v1/tenant.service.proto

/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import {
  CreateTenantRequest,
  CreateTenantResponse,
  DeleteTenantRequest,
  DeleteTenantResponse,
  GetAvailableTenantsRequest,
  GetAvailableTenantsResponse,
  GetTenantRequest,
  GetTenantResponse,
  GetTenantsRequest,
  GetTenantsResponse,
  UpdateTenantRequest,
  UpdateTenantResponse,
} from "./tenant.event";

export const protobufPackage = "gen.go.tenant.v1";

export interface TenantServicePort {
  create(request: CreateTenantRequest): Promise<CreateTenantResponse>;
  get(request: GetTenantRequest): Promise<GetTenantResponse>;
  getAll(request: GetTenantsRequest): Promise<GetTenantsResponse>;
  getAvailable(request: GetAvailableTenantsRequest): Promise<GetAvailableTenantsResponse>;
  update(request: UpdateTenantRequest): Promise<UpdateTenantResponse>;
  delete(request: DeleteTenantRequest): Promise<DeleteTenantResponse>;
}

export const TenantServicePortServiceName = "gen.go.tenant.v1.TenantServicePort";
export class TenantServicePortClientImpl implements TenantServicePort {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || TenantServicePortServiceName;
    this.rpc = rpc;
    this.create = this.create.bind(this);
    this.get = this.get.bind(this);
    this.getAll = this.getAll.bind(this);
    this.getAvailable = this.getAvailable.bind(this);
    this.update = this.update.bind(this);
    this.delete = this.delete.bind(this);
  }
  create(request: CreateTenantRequest): Promise<CreateTenantResponse> {
    const data = CreateTenantRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "create", data);
    return promise.then((data) => CreateTenantResponse.decode(_m0.Reader.create(data)));
  }

  get(request: GetTenantRequest): Promise<GetTenantResponse> {
    const data = GetTenantRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "get", data);
    return promise.then((data) => GetTenantResponse.decode(_m0.Reader.create(data)));
  }

  getAll(request: GetTenantsRequest): Promise<GetTenantsResponse> {
    const data = GetTenantsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "getAll", data);
    return promise.then((data) => GetTenantsResponse.decode(_m0.Reader.create(data)));
  }

  getAvailable(request: GetAvailableTenantsRequest): Promise<GetAvailableTenantsResponse> {
    const data = GetAvailableTenantsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "getAvailable", data);
    return promise.then((data) => GetAvailableTenantsResponse.decode(_m0.Reader.create(data)));
  }

  update(request: UpdateTenantRequest): Promise<UpdateTenantResponse> {
    const data = UpdateTenantRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "update", data);
    return promise.then((data) => UpdateTenantResponse.decode(_m0.Reader.create(data)));
  }

  delete(request: DeleteTenantRequest): Promise<DeleteTenantResponse> {
    const data = DeleteTenantRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "delete", data);
    return promise.then((data) => DeleteTenantResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
