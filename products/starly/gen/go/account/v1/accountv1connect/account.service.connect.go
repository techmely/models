// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: account/v1/account.service.proto

package accountv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	account "github.com/techmely/models/account"
	v1 "github.com/techmely/models/account/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// AccountServicePortName is the fully-qualified name of the AccountServicePort service.
	AccountServicePortName = "gen.go.account.v1.AccountServicePort"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccountServicePortSignInProcedure is the fully-qualified name of the AccountServicePort's SignIn
	// RPC.
	AccountServicePortSignInProcedure = "/gen.go.account.v1.AccountServicePort/SignIn"
	// AccountServicePortSignInWithProviderProcedure is the fully-qualified name of the
	// AccountServicePort's SignInWithProvider RPC.
	AccountServicePortSignInWithProviderProcedure = "/gen.go.account.v1.AccountServicePort/SignInWithProvider"
	// AccountServicePortSignUpProcedure is the fully-qualified name of the AccountServicePort's SignUp
	// RPC.
	AccountServicePortSignUpProcedure = "/gen.go.account.v1.AccountServicePort/SignUp"
	// AccountServicePortSignOutProcedure is the fully-qualified name of the AccountServicePort's
	// SignOut RPC.
	AccountServicePortSignOutProcedure = "/gen.go.account.v1.AccountServicePort/SignOut"
	// AccountServicePortResendVerificationCodeProcedure is the fully-qualified name of the
	// AccountServicePort's ResendVerificationCode RPC.
	AccountServicePortResendVerificationCodeProcedure = "/gen.go.account.v1.AccountServicePort/ResendVerificationCode"
	// AccountServicePortUpdatePasswordProcedure is the fully-qualified name of the AccountServicePort's
	// UpdatePassword RPC.
	AccountServicePortUpdatePasswordProcedure = "/gen.go.account.v1.AccountServicePort/UpdatePassword"
	// AccountServicePortUpdateEmailProcedure is the fully-qualified name of the AccountServicePort's
	// UpdateEmail RPC.
	AccountServicePortUpdateEmailProcedure = "/gen.go.account.v1.AccountServicePort/UpdateEmail"
	// AccountServicePortVerifyAccountProcedure is the fully-qualified name of the AccountServicePort's
	// VerifyAccount RPC.
	AccountServicePortVerifyAccountProcedure = "/gen.go.account.v1.AccountServicePort/VerifyAccount"
	// AccountServicePortVerifyActivationLinkProcedure is the fully-qualified name of the
	// AccountServicePort's VerifyActivationLink RPC.
	AccountServicePortVerifyActivationLinkProcedure = "/gen.go.account.v1.AccountServicePort/VerifyActivationLink"
	// AccountServicePortForgotPasswordProcedure is the fully-qualified name of the AccountServicePort's
	// ForgotPassword RPC.
	AccountServicePortForgotPasswordProcedure = "/gen.go.account.v1.AccountServicePort/ForgotPassword"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	accountServicePortServiceDescriptor                      = v1.File_account_v1_account_service_proto.Services().ByName("AccountServicePort")
	accountServicePortSignInMethodDescriptor                 = accountServicePortServiceDescriptor.Methods().ByName("SignIn")
	accountServicePortSignInWithProviderMethodDescriptor     = accountServicePortServiceDescriptor.Methods().ByName("SignInWithProvider")
	accountServicePortSignUpMethodDescriptor                 = accountServicePortServiceDescriptor.Methods().ByName("SignUp")
	accountServicePortSignOutMethodDescriptor                = accountServicePortServiceDescriptor.Methods().ByName("SignOut")
	accountServicePortResendVerificationCodeMethodDescriptor = accountServicePortServiceDescriptor.Methods().ByName("ResendVerificationCode")
	accountServicePortUpdatePasswordMethodDescriptor         = accountServicePortServiceDescriptor.Methods().ByName("UpdatePassword")
	accountServicePortUpdateEmailMethodDescriptor            = accountServicePortServiceDescriptor.Methods().ByName("UpdateEmail")
	accountServicePortVerifyAccountMethodDescriptor          = accountServicePortServiceDescriptor.Methods().ByName("VerifyAccount")
	accountServicePortVerifyActivationLinkMethodDescriptor   = accountServicePortServiceDescriptor.Methods().ByName("VerifyActivationLink")
	accountServicePortForgotPasswordMethodDescriptor         = accountServicePortServiceDescriptor.Methods().ByName("ForgotPassword")
)

// AccountServicePortClient is a client for the gen.go.account.v1.AccountServicePort service.
type AccountServicePortClient interface {
	SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error)
	SignInWithProvider(context.Context, *connect.Request[v1.SignInWithProviderRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error)
	SignUp(context.Context, *connect.Request[v1.SignUpRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error)
	SignOut(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error)
	ResendVerificationCode(context.Context, *connect.Request[v1.ResendVerificationCodeRequest]) (*connect.Response[v1.ResendVerificationCodeResponse], error)
	UpdatePassword(context.Context, *connect.Request[v1.UpdatePasswordRequest]) (*connect.Response[v1.UpdatePasswordResponse], error)
	UpdateEmail(context.Context, *connect.Request[v1.UpdateEmailRequest]) (*connect.Response[v1.UpdateEmailResponse], error)
	VerifyAccount(context.Context, *connect.Request[v1.VerifyAccountRequest]) (*connect.Response[v1.VerifyAccountResponse], error)
	VerifyActivationLink(context.Context, *connect.Request[v1.VerifyActivationLinkRequest]) (*connect.Response[v1.VerifyActivationLinkResponse], error)
	ForgotPassword(context.Context, *connect.Request[v1.ForgotPasswordRequest]) (*connect.Response[v1.ForgotPasswordResponse], error)
}

// NewAccountServicePortClient constructs a client for the gen.go.account.v1.AccountServicePort
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccountServicePortClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AccountServicePortClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accountServicePortClient{
		signIn: connect.NewClient[v1.SignInRequest, account.AuthGoogleIdentityResponse](
			httpClient,
			baseURL+AccountServicePortSignInProcedure,
			connect.WithSchema(accountServicePortSignInMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		signInWithProvider: connect.NewClient[v1.SignInWithProviderRequest, account.AuthGoogleIdentityResponse](
			httpClient,
			baseURL+AccountServicePortSignInWithProviderProcedure,
			connect.WithSchema(accountServicePortSignInWithProviderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		signUp: connect.NewClient[v1.SignUpRequest, account.AuthGoogleIdentityResponse](
			httpClient,
			baseURL+AccountServicePortSignUpProcedure,
			connect.WithSchema(accountServicePortSignUpMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		signOut: connect.NewClient[emptypb.Empty, emptypb.Empty](
			httpClient,
			baseURL+AccountServicePortSignOutProcedure,
			connect.WithSchema(accountServicePortSignOutMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		resendVerificationCode: connect.NewClient[v1.ResendVerificationCodeRequest, v1.ResendVerificationCodeResponse](
			httpClient,
			baseURL+AccountServicePortResendVerificationCodeProcedure,
			connect.WithSchema(accountServicePortResendVerificationCodeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updatePassword: connect.NewClient[v1.UpdatePasswordRequest, v1.UpdatePasswordResponse](
			httpClient,
			baseURL+AccountServicePortUpdatePasswordProcedure,
			connect.WithSchema(accountServicePortUpdatePasswordMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateEmail: connect.NewClient[v1.UpdateEmailRequest, v1.UpdateEmailResponse](
			httpClient,
			baseURL+AccountServicePortUpdateEmailProcedure,
			connect.WithSchema(accountServicePortUpdateEmailMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		verifyAccount: connect.NewClient[v1.VerifyAccountRequest, v1.VerifyAccountResponse](
			httpClient,
			baseURL+AccountServicePortVerifyAccountProcedure,
			connect.WithSchema(accountServicePortVerifyAccountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		verifyActivationLink: connect.NewClient[v1.VerifyActivationLinkRequest, v1.VerifyActivationLinkResponse](
			httpClient,
			baseURL+AccountServicePortVerifyActivationLinkProcedure,
			connect.WithSchema(accountServicePortVerifyActivationLinkMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		forgotPassword: connect.NewClient[v1.ForgotPasswordRequest, v1.ForgotPasswordResponse](
			httpClient,
			baseURL+AccountServicePortForgotPasswordProcedure,
			connect.WithSchema(accountServicePortForgotPasswordMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// accountServicePortClient implements AccountServicePortClient.
type accountServicePortClient struct {
	signIn                 *connect.Client[v1.SignInRequest, account.AuthGoogleIdentityResponse]
	signInWithProvider     *connect.Client[v1.SignInWithProviderRequest, account.AuthGoogleIdentityResponse]
	signUp                 *connect.Client[v1.SignUpRequest, account.AuthGoogleIdentityResponse]
	signOut                *connect.Client[emptypb.Empty, emptypb.Empty]
	resendVerificationCode *connect.Client[v1.ResendVerificationCodeRequest, v1.ResendVerificationCodeResponse]
	updatePassword         *connect.Client[v1.UpdatePasswordRequest, v1.UpdatePasswordResponse]
	updateEmail            *connect.Client[v1.UpdateEmailRequest, v1.UpdateEmailResponse]
	verifyAccount          *connect.Client[v1.VerifyAccountRequest, v1.VerifyAccountResponse]
	verifyActivationLink   *connect.Client[v1.VerifyActivationLinkRequest, v1.VerifyActivationLinkResponse]
	forgotPassword         *connect.Client[v1.ForgotPasswordRequest, v1.ForgotPasswordResponse]
}

// SignIn calls gen.go.account.v1.AccountServicePort.SignIn.
func (c *accountServicePortClient) SignIn(ctx context.Context, req *connect.Request[v1.SignInRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error) {
	return c.signIn.CallUnary(ctx, req)
}

// SignInWithProvider calls gen.go.account.v1.AccountServicePort.SignInWithProvider.
func (c *accountServicePortClient) SignInWithProvider(ctx context.Context, req *connect.Request[v1.SignInWithProviderRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error) {
	return c.signInWithProvider.CallUnary(ctx, req)
}

// SignUp calls gen.go.account.v1.AccountServicePort.SignUp.
func (c *accountServicePortClient) SignUp(ctx context.Context, req *connect.Request[v1.SignUpRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error) {
	return c.signUp.CallUnary(ctx, req)
}

// SignOut calls gen.go.account.v1.AccountServicePort.SignOut.
func (c *accountServicePortClient) SignOut(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	return c.signOut.CallUnary(ctx, req)
}

// ResendVerificationCode calls gen.go.account.v1.AccountServicePort.ResendVerificationCode.
func (c *accountServicePortClient) ResendVerificationCode(ctx context.Context, req *connect.Request[v1.ResendVerificationCodeRequest]) (*connect.Response[v1.ResendVerificationCodeResponse], error) {
	return c.resendVerificationCode.CallUnary(ctx, req)
}

// UpdatePassword calls gen.go.account.v1.AccountServicePort.UpdatePassword.
func (c *accountServicePortClient) UpdatePassword(ctx context.Context, req *connect.Request[v1.UpdatePasswordRequest]) (*connect.Response[v1.UpdatePasswordResponse], error) {
	return c.updatePassword.CallUnary(ctx, req)
}

// UpdateEmail calls gen.go.account.v1.AccountServicePort.UpdateEmail.
func (c *accountServicePortClient) UpdateEmail(ctx context.Context, req *connect.Request[v1.UpdateEmailRequest]) (*connect.Response[v1.UpdateEmailResponse], error) {
	return c.updateEmail.CallUnary(ctx, req)
}

// VerifyAccount calls gen.go.account.v1.AccountServicePort.VerifyAccount.
func (c *accountServicePortClient) VerifyAccount(ctx context.Context, req *connect.Request[v1.VerifyAccountRequest]) (*connect.Response[v1.VerifyAccountResponse], error) {
	return c.verifyAccount.CallUnary(ctx, req)
}

// VerifyActivationLink calls gen.go.account.v1.AccountServicePort.VerifyActivationLink.
func (c *accountServicePortClient) VerifyActivationLink(ctx context.Context, req *connect.Request[v1.VerifyActivationLinkRequest]) (*connect.Response[v1.VerifyActivationLinkResponse], error) {
	return c.verifyActivationLink.CallUnary(ctx, req)
}

// ForgotPassword calls gen.go.account.v1.AccountServicePort.ForgotPassword.
func (c *accountServicePortClient) ForgotPassword(ctx context.Context, req *connect.Request[v1.ForgotPasswordRequest]) (*connect.Response[v1.ForgotPasswordResponse], error) {
	return c.forgotPassword.CallUnary(ctx, req)
}

// AccountServicePortHandler is an implementation of the gen.go.account.v1.AccountServicePort
// service.
type AccountServicePortHandler interface {
	SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error)
	SignInWithProvider(context.Context, *connect.Request[v1.SignInWithProviderRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error)
	SignUp(context.Context, *connect.Request[v1.SignUpRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error)
	SignOut(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error)
	ResendVerificationCode(context.Context, *connect.Request[v1.ResendVerificationCodeRequest]) (*connect.Response[v1.ResendVerificationCodeResponse], error)
	UpdatePassword(context.Context, *connect.Request[v1.UpdatePasswordRequest]) (*connect.Response[v1.UpdatePasswordResponse], error)
	UpdateEmail(context.Context, *connect.Request[v1.UpdateEmailRequest]) (*connect.Response[v1.UpdateEmailResponse], error)
	VerifyAccount(context.Context, *connect.Request[v1.VerifyAccountRequest]) (*connect.Response[v1.VerifyAccountResponse], error)
	VerifyActivationLink(context.Context, *connect.Request[v1.VerifyActivationLinkRequest]) (*connect.Response[v1.VerifyActivationLinkResponse], error)
	ForgotPassword(context.Context, *connect.Request[v1.ForgotPasswordRequest]) (*connect.Response[v1.ForgotPasswordResponse], error)
}

// NewAccountServicePortHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccountServicePortHandler(svc AccountServicePortHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	accountServicePortSignInHandler := connect.NewUnaryHandler(
		AccountServicePortSignInProcedure,
		svc.SignIn,
		connect.WithSchema(accountServicePortSignInMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortSignInWithProviderHandler := connect.NewUnaryHandler(
		AccountServicePortSignInWithProviderProcedure,
		svc.SignInWithProvider,
		connect.WithSchema(accountServicePortSignInWithProviderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortSignUpHandler := connect.NewUnaryHandler(
		AccountServicePortSignUpProcedure,
		svc.SignUp,
		connect.WithSchema(accountServicePortSignUpMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortSignOutHandler := connect.NewUnaryHandler(
		AccountServicePortSignOutProcedure,
		svc.SignOut,
		connect.WithSchema(accountServicePortSignOutMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortResendVerificationCodeHandler := connect.NewUnaryHandler(
		AccountServicePortResendVerificationCodeProcedure,
		svc.ResendVerificationCode,
		connect.WithSchema(accountServicePortResendVerificationCodeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortUpdatePasswordHandler := connect.NewUnaryHandler(
		AccountServicePortUpdatePasswordProcedure,
		svc.UpdatePassword,
		connect.WithSchema(accountServicePortUpdatePasswordMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortUpdateEmailHandler := connect.NewUnaryHandler(
		AccountServicePortUpdateEmailProcedure,
		svc.UpdateEmail,
		connect.WithSchema(accountServicePortUpdateEmailMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortVerifyAccountHandler := connect.NewUnaryHandler(
		AccountServicePortVerifyAccountProcedure,
		svc.VerifyAccount,
		connect.WithSchema(accountServicePortVerifyAccountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortVerifyActivationLinkHandler := connect.NewUnaryHandler(
		AccountServicePortVerifyActivationLinkProcedure,
		svc.VerifyActivationLink,
		connect.WithSchema(accountServicePortVerifyActivationLinkMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accountServicePortForgotPasswordHandler := connect.NewUnaryHandler(
		AccountServicePortForgotPasswordProcedure,
		svc.ForgotPassword,
		connect.WithSchema(accountServicePortForgotPasswordMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/gen.go.account.v1.AccountServicePort/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccountServicePortSignInProcedure:
			accountServicePortSignInHandler.ServeHTTP(w, r)
		case AccountServicePortSignInWithProviderProcedure:
			accountServicePortSignInWithProviderHandler.ServeHTTP(w, r)
		case AccountServicePortSignUpProcedure:
			accountServicePortSignUpHandler.ServeHTTP(w, r)
		case AccountServicePortSignOutProcedure:
			accountServicePortSignOutHandler.ServeHTTP(w, r)
		case AccountServicePortResendVerificationCodeProcedure:
			accountServicePortResendVerificationCodeHandler.ServeHTTP(w, r)
		case AccountServicePortUpdatePasswordProcedure:
			accountServicePortUpdatePasswordHandler.ServeHTTP(w, r)
		case AccountServicePortUpdateEmailProcedure:
			accountServicePortUpdateEmailHandler.ServeHTTP(w, r)
		case AccountServicePortVerifyAccountProcedure:
			accountServicePortVerifyAccountHandler.ServeHTTP(w, r)
		case AccountServicePortVerifyActivationLinkProcedure:
			accountServicePortVerifyActivationLinkHandler.ServeHTTP(w, r)
		case AccountServicePortForgotPasswordProcedure:
			accountServicePortForgotPasswordHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccountServicePortHandler returns CodeUnimplemented from all methods.
type UnimplementedAccountServicePortHandler struct{}

func (UnimplementedAccountServicePortHandler) SignIn(context.Context, *connect.Request[v1.SignInRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.SignIn is not implemented"))
}

func (UnimplementedAccountServicePortHandler) SignInWithProvider(context.Context, *connect.Request[v1.SignInWithProviderRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.SignInWithProvider is not implemented"))
}

func (UnimplementedAccountServicePortHandler) SignUp(context.Context, *connect.Request[v1.SignUpRequest]) (*connect.Response[account.AuthGoogleIdentityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.SignUp is not implemented"))
}

func (UnimplementedAccountServicePortHandler) SignOut(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.SignOut is not implemented"))
}

func (UnimplementedAccountServicePortHandler) ResendVerificationCode(context.Context, *connect.Request[v1.ResendVerificationCodeRequest]) (*connect.Response[v1.ResendVerificationCodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.ResendVerificationCode is not implemented"))
}

func (UnimplementedAccountServicePortHandler) UpdatePassword(context.Context, *connect.Request[v1.UpdatePasswordRequest]) (*connect.Response[v1.UpdatePasswordResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.UpdatePassword is not implemented"))
}

func (UnimplementedAccountServicePortHandler) UpdateEmail(context.Context, *connect.Request[v1.UpdateEmailRequest]) (*connect.Response[v1.UpdateEmailResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.UpdateEmail is not implemented"))
}

func (UnimplementedAccountServicePortHandler) VerifyAccount(context.Context, *connect.Request[v1.VerifyAccountRequest]) (*connect.Response[v1.VerifyAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.VerifyAccount is not implemented"))
}

func (UnimplementedAccountServicePortHandler) VerifyActivationLink(context.Context, *connect.Request[v1.VerifyActivationLinkRequest]) (*connect.Response[v1.VerifyActivationLinkResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.VerifyActivationLink is not implemented"))
}

func (UnimplementedAccountServicePortHandler) ForgotPassword(context.Context, *connect.Request[v1.ForgotPasswordRequest]) (*connect.Response[v1.ForgotPasswordResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gen.go.account.v1.AccountServicePort.ForgotPassword is not implemented"))
}
