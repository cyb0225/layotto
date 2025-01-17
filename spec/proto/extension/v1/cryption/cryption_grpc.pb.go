// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: cryption.proto

package cryption

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

// CryptionServiceClient is the client API for CryptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptionServiceClient interface {
	// Encrypt data
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
	// Decrypt data
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type cryptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptionServiceClient(cc grpc.ClientConnInterface) CryptionServiceClient {
	return &cryptionServiceClient{cc}
}

func (c *cryptionServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/spec.proto.extension.v1.cryption.CryptionService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptionServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/spec.proto.extension.v1.cryption.CryptionService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptionServiceServer is the server API for CryptionService service.
// All implementations should embed UnimplementedCryptionServiceServer
// for forward compatibility
type CryptionServiceServer interface {
	// Encrypt data
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	// Decrypt data
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
}

// UnimplementedCryptionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCryptionServiceServer struct {
}

func (UnimplementedCryptionServiceServer) Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedCryptionServiceServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}

// UnsafeCryptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptionServiceServer will
// result in compilation errors.
type UnsafeCryptionServiceServer interface {
	mustEmbedUnimplementedCryptionServiceServer()
}

func RegisterCryptionServiceServer(s grpc.ServiceRegistrar, srv CryptionServiceServer) {
	s.RegisterService(&CryptionService_ServiceDesc, srv)
}

func _CryptionService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptionServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.proto.extension.v1.cryption.CryptionService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptionServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptionService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptionServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.proto.extension.v1.cryption.CryptionService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptionServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptionService_ServiceDesc is the grpc.ServiceDesc for CryptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spec.proto.extension.v1.cryption.CryptionService",
	HandlerType: (*CryptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _CryptionService_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _CryptionService_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cryption.proto",
}
