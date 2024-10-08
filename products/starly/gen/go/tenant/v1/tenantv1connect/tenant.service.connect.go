// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tenant/v1/tenant.service.proto

package tenantv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/techmely/models/tenant/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TenantServicePortName is the fully-qualified name of the TenantServicePort service.
	TenantServicePortName = "gen.go.tenant.v1.TenantServicePort"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TenantServicePortCreateProcedure is the fully-qualified name of the TenantServicePort's create
	// RPC.
	TenantServicePortCreateProcedure = "/gen.go.tenant.v1.TenantServicePort/create"
	// TenantServicePortGetProcedure is the fully-qualified name of the TenantServicePort's get RPC.
	TenantServicePortGetProcedure = "/gen.go.tenant.v1.TenantServicePort/get"
	// TenantServicePortGetAllProcedure is the fully-qualified name of the TenantServicePort's getAll
	// RPC.
	TenantServicePortGetAllProcedure = "/gen.go.tenant.v1.TenantServicePort/getAll"
	// TenantServicePortGetAvailableProcedure is the fully-qualified name of the TenantServicePort's
	// getAvailable RPC.
	TenantServicePortGetAvailableProcedure = "/gen.go.tenant.v1.TenantServicePort/getAvailable"
	// TenantServicePortUpdateProcedure is the fully-qualified name of the TenantServicePort's update
	// RPC.
	TenantServicePortUpdateProcedure = "/gen.go.tenant.v1.TenantServicePort/update"
	// TenantServicePortDeleteProcedure is the fully-qualified name of the TenantServicePort's delete
	// RPC.
	TenantServicePortDeleteProcedure = "/gen.go.tenant.v1.TenantServicePort/delete"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	tenantServicePortServiceDescriptor            = v1.File_tenant_v1_tenant_service_proto.Services().ByName("TenantServicePort")
	tenantServicePortCreateMethodDescriptor       = tenantServicePortServiceDescriptor.Methods().ByName("create")
	tenantServicePortGetMethodDescriptor          = tenantServicePortServiceDescriptor.Methods().ByName("get")
	tenantServicePortGetAllMethodDescriptor       = tenantServicePortServiceDescriptor.Methods().ByName("getAll")
	tenantServicePortGetAvailableMethodDescriptor = tenantServicePortServiceDescriptor.Methods().ByName("getAvailable")
	tenantServicePortUpdateMethodDescriptor       = tenantServicePortServiceDescriptor.Methods().ByName("update")
	tenantServicePortDeleteMethodDescriptor       = tenantServicePortServiceDescriptor.Methods().ByName("delete")
)

// TenantServicePortClient is a client for the gen.go.tenant.v1.TenantServicePort service.
type TenantServicePortClient interface {
	Create(context.Context, *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error)
	Get(context.Context, *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error)
	GetAll(context.Context, *connect.Request[v1.GetTenantsRequest]) (*connect.Response[v1.GetTenantsResponse], error)
	GetAvailable(context.Context, *connect.Request[v1.GetAvailableTenantsRequest]) (*connect.Response[v1.GetAvailableTenantsResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateTenantRequest]) (*connect.Response[v1.UpdateTenantResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error)
}

// NewTenantServicePortClient constructs a client for the gen.go.tenant.v1.TenantServicePort
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTenantServicePortClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TenantServicePortClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tenantServicePortClient{
		create: connect.NewClient[v1.CreateTenantRequest, v1.CreateTenantResponse](
			httpClient,
			baseURL+TenantServicePortCreateProcedure,
			connect.WithSchema(tenantServicePortCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[v1.GetTenantRequest, v1.GetTenantResponse](
			httpClient,
			baseURL+TenantServicePortGetProcedure,
			connect.WithSchema(tenantServicePortGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAll: connect.NewClient[v1.GetTenantsRequest, v1.GetTenantsResponse](
			httpClient,
			baseURL+TenantServicePortGetAllProcedure,
			connect.WithSchema(tenantServicePortGetAllMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAvailable: connect.NewClient[v1.GetAvailableTenantsRequest, v1.GetAvailableTenantsResponse](
			httpClient,
			baseURL+TenantServicePortGetAvailableProcedure,
			connect.WithSchema(tenantServicePortGetAvailableMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		update: connect.NewClient[v1.UpdateTenantRequest, v1.UpdateTenantResponse](
			httpClient,
			baseURL+TenantServicePortUpdateProcedure,
			connect.WithSchema(tenantServicePortUpdateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		delete: connect.NewClient[v1.DeleteTenantRequest, v1.DeleteTenantResponse](
			httpClient,
			baseURL+TenantServicePortDeleteProcedure,
			connect.WithSchema(tenantServicePortDeleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// tenantServicePortClient implements TenantServicePortClient.
type tenantServicePortClient struct {
	create       *connect.Client[v1.CreateTenantRequest, v1.CreateTenantResponse]
	get          *connect.Client[v1.GetTenantRequest, v1.GetTenantResponse]
	getAll       *connect.Client[v1.GetTenantsRequest, v1.GetTenantsResponse]
	getAvailable *connect.Client[v1.GetAvailableTenantsRequest, v1.GetAvailableTenantsResponse]
	update       *connect.Client[v1.UpdateTenantRequest, v1.UpdateTenantResponse]
	delete       *connect.Client[v1.DeleteTenantRequest, v1.DeleteTenantResponse]
}

// Create calls gen.go.tenant.v1.TenantServicePort.create.
func (c *tenantServicePortClient) Create(ctx context.Context, req *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Get calls gen.go.tenant.v1.TenantServicePort.get.
func (c *tenantServicePortClient) Get(ctx context.Context, req *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// GetAll calls gen.go.tenant.v1.TenantServicePort.getAll.
func (c *tenantServicePortClient) GetAll(ctx context.Context, req *connect.Request[v1.GetTenantsRequest]) (*connect.Response[v1.GetTenantsResponse], error) {
	return c.getAll.CallUnary(ctx, req)
}

// GetAvailable calls gen.go.tenant.v1.TenantServicePort.getAvailable.
func (c *tenantServicePortClient) GetAvailable(ctx context.Context, req *connect.Request[v1.GetAvailableTenantsRequest]) (*connect.Response[v1.GetAvailableTenantsResponse], error) {
	return c.getAvailable.CallUnary(ctx, req)
}

// Update calls gen.go.tenant.v1.TenantServicePort.update.
func (c *tenantServicePortClient) Update(ctx context.Context, req *connect.Request[v1.UpdateTenantRequest]) (*connect.Response[v1.UpdateTenantResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls gen.go.tenant.v1.TenantServicePort.delete.
func (c *tenantServicePortClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// TenantServicePortHandler is an implementation of the gen.go.tenant.v1.TenantServicePort service.
type TenantServicePortHandler interface {
	Create(context.Context, *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error)
	Get(context.Context, *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error)
	GetAll(context.Context, *connect.Request[v1.GetTenantsRequest]) (*connect.Response[v1.GetTenantsResponse], error)
	GetAvailable(context.Context, *connect.Request[v1.GetAvailableTenantsRequest]) (*connect.Response[v1.GetAvailableTenantsResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateTenantRequest]) (*connect.Response[v1.UpdateTenantResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error)
}

// NewTenantServicePortHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTenantServicePortHandler(svc TenantServicePortHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tenantServicePortCreateHandler := connect.NewUnaryHandler(
		TenantServicePortCreateProcedure,
		svc.Create,
		connect.WithSchema(tenantServicePortCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tenantServicePortGetHandler := connect.NewUnaryHandler(
		TenantServicePortGetProcedure,
		svc.Get,
		connect.WithSchema(tenantServicePortGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tenantServicePortGetAllHandler := connect.NewUnaryHandler(
		TenantServicePortGetAllProcedure,
		svc.GetAll,
		connect.WithSchema(tenantServicePortGetAllMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tenantServicePortGetAvailableHandler := connect.NewUnaryHandler(
		TenantServicePortGetAvailableProcedure,
		svc.GetAvailable,
		connect.WithSchema(tenantServicePortGetAvailableMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tenantServicePortUpdateHandler := connect.NewUnaryHandler(
		TenantServicePortUpdateProcedure,
		svc.Update,
		connect.WithSchema(tenantServicePortUpdateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tenantServicePortDeleteHandler := connect.NewUnaryHandler(
		TenantServicePortDeleteProcedure,
		svc.Delete,
		connect.WithSchema(tenantServicePortDeleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/gen.go.tenant.v1.TenantServicePort/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TenantServicePortCreateProcedure:
			tenantServicePortCreateHandler.ServeHTTP(w, r)
		case TenantServicePortGetProcedure:
			tenantServicePortGetHandler.ServeHTTP(w, r)
		case TenantServicePortGetAllProcedure:
			tenantServicePortGetAllHandler.ServeHTTP(w, r)
		case TenantServicePortGetAvailableProcedure:
			tenantServicePortGetAvailableHandler.ServeHTTP(w, r)
		case TenantServicePortUpdateProcedure:
			tenantServicePortUpdateHandler.ServeHTTP(w, r)
		case TenantServicePortDeleteProcedure:
			tenantServicePortDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTenantServicePortHandler returns CodeUnimplemented from all methods.
type UnimplementedTenantServicePortHandler struct{}

func (UnimplementedTenantServicePortHandler) Create(context.Context, *connect.Request[v1.CreateTenantRequest]) (*connect.Response[v1.CreateTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.tenant.v1.TenantServicePort.create is not implemented"))
}

func (UnimplementedTenantServicePortHandler) Get(context.Context, *connect.Request[v1.GetTenantRequest]) (*connect.Response[v1.GetTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.tenant.v1.TenantServicePort.get is not implemented"))
}

func (UnimplementedTenantServicePortHandler) GetAll(context.Context, *connect.Request[v1.GetTenantsRequest]) (*connect.Response[v1.GetTenantsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.tenant.v1.TenantServicePort.getAll is not implemented"))
}

func (UnimplementedTenantServicePortHandler) GetAvailable(context.Context, *connect.Request[v1.GetAvailableTenantsRequest]) (*connect.Response[v1.GetAvailableTenantsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.tenant.v1.TenantServicePort.getAvailable is not implemented"))
}

func (UnimplementedTenantServicePortHandler) Update(context.Context, *connect.Request[v1.UpdateTenantRequest]) (*connect.Response[v1.UpdateTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.tenant.v1.TenantServicePort.update is not implemented"))
}

func (UnimplementedTenantServicePortHandler) Delete(context.Context, *connect.Request[v1.DeleteTenantRequest]) (*connect.Response[v1.DeleteTenantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.tenant.v1.TenantServicePort.delete is not implemented"))
}
