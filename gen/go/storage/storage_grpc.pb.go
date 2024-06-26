// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: storage.proto

package storage

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

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	SaveFile(ctx context.Context, opts ...grpc.CallOption) (Storage_SaveFileClient, error)
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (Storage_GetFileClient, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) SaveFile(ctx context.Context, opts ...grpc.CallOption) (Storage_SaveFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[0], "/apollo.Storage/SaveFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageSaveFileClient{stream}
	return x, nil
}

type Storage_SaveFileClient interface {
	Send(*SaveFileRequest) error
	CloseAndRecv() (*SaveFileResponse, error)
	grpc.ClientStream
}

type storageSaveFileClient struct {
	grpc.ClientStream
}

func (x *storageSaveFileClient) Send(m *SaveFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageSaveFileClient) CloseAndRecv() (*SaveFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SaveFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (Storage_GetFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[1], "/apollo.Storage/GetFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_GetFileClient interface {
	Recv() (*GetFileResponse, error)
	grpc.ClientStream
}

type storageGetFileClient struct {
	grpc.ClientStream
}

func (x *storageGetFileClient) Recv() (*GetFileResponse, error) {
	m := new(GetFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	SaveFile(Storage_SaveFileServer) error
	GetFile(*GetFileRequest, Storage_GetFileServer) error
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) SaveFile(Storage_SaveFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveFile not implemented")
}
func (UnimplementedStorageServer) GetFile(*GetFileRequest, Storage_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_SaveFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).SaveFile(&storageSaveFileServer{stream})
}

type Storage_SaveFileServer interface {
	SendAndClose(*SaveFileResponse) error
	Recv() (*SaveFileRequest, error)
	grpc.ServerStream
}

type storageSaveFileServer struct {
	grpc.ServerStream
}

func (x *storageSaveFileServer) SendAndClose(m *SaveFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageSaveFileServer) Recv() (*SaveFileRequest, error) {
	m := new(SaveFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storage_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).GetFile(m, &storageGetFileServer{stream})
}

type Storage_GetFileServer interface {
	Send(*GetFileResponse) error
	grpc.ServerStream
}

type storageGetFileServer struct {
	grpc.ServerStream
}

func (x *storageGetFileServer) Send(m *GetFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apollo.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveFile",
			Handler:       _Storage_SaveFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetFile",
			Handler:       _Storage_GetFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "storage.proto",
}
