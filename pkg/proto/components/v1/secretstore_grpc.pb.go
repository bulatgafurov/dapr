//
//Copyright 2023 The Dapr Authors
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: dapr/proto/components/v1/secretstore.proto

package components

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SecretStore_Init_FullMethodName     = "/dapr.proto.components.v1.SecretStore/Init"
	SecretStore_Features_FullMethodName = "/dapr.proto.components.v1.SecretStore/Features"
	SecretStore_Get_FullMethodName      = "/dapr.proto.components.v1.SecretStore/Get"
	SecretStore_BulkGet_FullMethodName  = "/dapr.proto.components.v1.SecretStore/BulkGet"
	SecretStore_Ping_FullMethodName     = "/dapr.proto.components.v1.SecretStore/Ping"
)

// SecretStoreClient is the client API for SecretStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretStoreClient interface {
	// Initializes the secret store with the given metadata.
	Init(ctx context.Context, in *SecretStoreInitRequest, opts ...grpc.CallOption) (*SecretStoreInitResponse, error)
	// Returns a list of implemented secret store features.
	Features(ctx context.Context, in *FeaturesRequest, opts ...grpc.CallOption) (*FeaturesResponse, error)
	// Get an individual secret from the store.
	Get(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
	// Get all secrets from the store.
	BulkGet(ctx context.Context, in *BulkGetSecretRequest, opts ...grpc.CallOption) (*BulkGetSecretResponse, error)
	// Ping the pubsub. Used for liveness porpuses.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type secretStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretStoreClient(cc grpc.ClientConnInterface) SecretStoreClient {
	return &secretStoreClient{cc}
}

func (c *secretStoreClient) Init(ctx context.Context, in *SecretStoreInitRequest, opts ...grpc.CallOption) (*SecretStoreInitResponse, error) {
	out := new(SecretStoreInitResponse)
	err := c.cc.Invoke(ctx, SecretStore_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretStoreClient) Features(ctx context.Context, in *FeaturesRequest, opts ...grpc.CallOption) (*FeaturesResponse, error) {
	out := new(FeaturesResponse)
	err := c.cc.Invoke(ctx, SecretStore_Features_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretStoreClient) Get(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, SecretStore_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretStoreClient) BulkGet(ctx context.Context, in *BulkGetSecretRequest, opts ...grpc.CallOption) (*BulkGetSecretResponse, error) {
	out := new(BulkGetSecretResponse)
	err := c.cc.Invoke(ctx, SecretStore_BulkGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretStoreClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, SecretStore_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretStoreServer is the server API for SecretStore service.
// All implementations should embed UnimplementedSecretStoreServer
// for forward compatibility
type SecretStoreServer interface {
	// Initializes the secret store with the given metadata.
	Init(context.Context, *SecretStoreInitRequest) (*SecretStoreInitResponse, error)
	// Returns a list of implemented secret store features.
	Features(context.Context, *FeaturesRequest) (*FeaturesResponse, error)
	// Get an individual secret from the store.
	Get(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	// Get all secrets from the store.
	BulkGet(context.Context, *BulkGetSecretRequest) (*BulkGetSecretResponse, error)
	// Ping the pubsub. Used for liveness porpuses.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedSecretStoreServer should be embedded to have forward compatible implementations.
type UnimplementedSecretStoreServer struct {
}

func (UnimplementedSecretStoreServer) Init(context.Context, *SecretStoreInitRequest) (*SecretStoreInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedSecretStoreServer) Features(context.Context, *FeaturesRequest) (*FeaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Features not implemented")
}
func (UnimplementedSecretStoreServer) Get(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSecretStoreServer) BulkGet(context.Context, *BulkGetSecretRequest) (*BulkGetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkGet not implemented")
}
func (UnimplementedSecretStoreServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

// UnsafeSecretStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretStoreServer will
// result in compilation errors.
type UnsafeSecretStoreServer interface {
	mustEmbedUnimplementedSecretStoreServer()
}

func RegisterSecretStoreServer(s grpc.ServiceRegistrar, srv SecretStoreServer) {
	s.RegisterService(&SecretStore_ServiceDesc, srv)
}

func _SecretStore_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretStoreInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretStoreServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretStore_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretStoreServer).Init(ctx, req.(*SecretStoreInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretStore_Features_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretStoreServer).Features(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretStore_Features_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretStoreServer).Features(ctx, req.(*FeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretStore_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretStoreServer).Get(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretStore_BulkGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkGetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretStoreServer).BulkGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretStore_BulkGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretStoreServer).BulkGet(ctx, req.(*BulkGetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretStore_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretStoreServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretStore_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretStoreServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretStore_ServiceDesc is the grpc.ServiceDesc for SecretStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.components.v1.SecretStore",
	HandlerType: (*SecretStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _SecretStore_Init_Handler,
		},
		{
			MethodName: "Features",
			Handler:    _SecretStore_Features_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SecretStore_Get_Handler,
		},
		{
			MethodName: "BulkGet",
			Handler:    _SecretStore_BulkGet_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _SecretStore_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapr/proto/components/v1/secretstore.proto",
}