// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.180.0
//   protoc               unknown
// source: account/v1/account.service.proto

/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../google/protobuf/empty";
import { AuthGoogleIdentityResponse } from "../firebase.model";
import {
  ForgotPasswordRequest,
  ForgotPasswordResponse,
  ResendVerificationCodeRequest,
  ResendVerificationCodeResponse,
  SignInRequest,
  SignInWithProviderRequest,
  SignUpRequest,
  UpdateEmailRequest,
  UpdateEmailResponse,
  UpdatePasswordRequest,
  UpdatePasswordResponse,
  VerifyAccountRequest,
  VerifyAccountResponse,
  VerifyActivationLinkRequest,
  VerifyActivationLinkResponse,
} from "./account.event";

export const protobufPackage = "gen.go.account.v1";

export interface AccountServicePort {
  SignIn(request: SignInRequest): Promise<AuthGoogleIdentityResponse>;
  SignInWithProvider(request: SignInWithProviderRequest): Promise<AuthGoogleIdentityResponse>;
  SignUp(request: SignUpRequest): Promise<AuthGoogleIdentityResponse>;
  SignOut(request: Empty): Promise<Empty>;
  ResendVerificationCode(request: ResendVerificationCodeRequest): Promise<ResendVerificationCodeResponse>;
  UpdatePassword(request: UpdatePasswordRequest): Promise<UpdatePasswordResponse>;
  UpdateEmail(request: UpdateEmailRequest): Promise<UpdateEmailResponse>;
  VerifyAccount(request: VerifyAccountRequest): Promise<VerifyAccountResponse>;
  VerifyActivationLink(request: VerifyActivationLinkRequest): Promise<VerifyActivationLinkResponse>;
  ForgotPassword(request: ForgotPasswordRequest): Promise<ForgotPasswordResponse>;
}

export const AccountServicePortServiceName = "gen.go.account.v1.AccountServicePort";
export class AccountServicePortClientImpl implements AccountServicePort {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || AccountServicePortServiceName;
    this.rpc = rpc;
    this.SignIn = this.SignIn.bind(this);
    this.SignInWithProvider = this.SignInWithProvider.bind(this);
    this.SignUp = this.SignUp.bind(this);
    this.SignOut = this.SignOut.bind(this);
    this.ResendVerificationCode = this.ResendVerificationCode.bind(this);
    this.UpdatePassword = this.UpdatePassword.bind(this);
    this.UpdateEmail = this.UpdateEmail.bind(this);
    this.VerifyAccount = this.VerifyAccount.bind(this);
    this.VerifyActivationLink = this.VerifyActivationLink.bind(this);
    this.ForgotPassword = this.ForgotPassword.bind(this);
  }
  SignIn(request: SignInRequest): Promise<AuthGoogleIdentityResponse> {
    const data = SignInRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SignIn", data);
    return promise.then((data) => AuthGoogleIdentityResponse.decode(_m0.Reader.create(data)));
  }

  SignInWithProvider(request: SignInWithProviderRequest): Promise<AuthGoogleIdentityResponse> {
    const data = SignInWithProviderRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SignInWithProvider", data);
    return promise.then((data) => AuthGoogleIdentityResponse.decode(_m0.Reader.create(data)));
  }

  SignUp(request: SignUpRequest): Promise<AuthGoogleIdentityResponse> {
    const data = SignUpRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "SignUp", data);
    return promise.then((data) => AuthGoogleIdentityResponse.decode(_m0.Reader.create(data)));
  }

  SignOut(request: Empty): Promise<Empty> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request(this.service, "SignOut", data);
    return promise.then((data) => Empty.decode(_m0.Reader.create(data)));
  }

  ResendVerificationCode(request: ResendVerificationCodeRequest): Promise<ResendVerificationCodeResponse> {
    const data = ResendVerificationCodeRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ResendVerificationCode", data);
    return promise.then((data) => ResendVerificationCodeResponse.decode(_m0.Reader.create(data)));
  }

  UpdatePassword(request: UpdatePasswordRequest): Promise<UpdatePasswordResponse> {
    const data = UpdatePasswordRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdatePassword", data);
    return promise.then((data) => UpdatePasswordResponse.decode(_m0.Reader.create(data)));
  }

  UpdateEmail(request: UpdateEmailRequest): Promise<UpdateEmailResponse> {
    const data = UpdateEmailRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateEmail", data);
    return promise.then((data) => UpdateEmailResponse.decode(_m0.Reader.create(data)));
  }

  VerifyAccount(request: VerifyAccountRequest): Promise<VerifyAccountResponse> {
    const data = VerifyAccountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "VerifyAccount", data);
    return promise.then((data) => VerifyAccountResponse.decode(_m0.Reader.create(data)));
  }

  VerifyActivationLink(request: VerifyActivationLinkRequest): Promise<VerifyActivationLinkResponse> {
    const data = VerifyActivationLinkRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "VerifyActivationLink", data);
    return promise.then((data) => VerifyActivationLinkResponse.decode(_m0.Reader.create(data)));
  }

  ForgotPassword(request: ForgotPasswordRequest): Promise<ForgotPasswordResponse> {
    const data = ForgotPasswordRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ForgotPassword", data);
    return promise.then((data) => ForgotPasswordResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}