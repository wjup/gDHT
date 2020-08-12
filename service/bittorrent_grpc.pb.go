// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BitTorrentClient is the client API for BitTorrent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BitTorrentClient interface {
	Send(ctx context.Context, in *BitTorrent, opts ...grpc.CallOption) (*Status, error)
}

type bitTorrentClient struct {
	cc grpc.ClientConnInterface
}

func NewBitTorrentClient(cc grpc.ClientConnInterface) BitTorrentClient {
	return &bitTorrentClient{cc}
}

func (c *bitTorrentClient) Send(ctx context.Context, in *BitTorrent, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service.bitTorrent/send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BitTorrentServer is the server API for BitTorrent service.
// All implementations must embed UnimplementedBitTorrentServer
// for forward compatibility
type BitTorrentServer interface {
	Send(context.Context, *BitTorrent) (*Status, error)
	mustEmbedUnimplementedBitTorrentServer()
}

// UnimplementedBitTorrentServer must be embedded to have forward compatible implementations.
type UnimplementedBitTorrentServer struct {
}

func (*UnimplementedBitTorrentServer) Send(context.Context, *BitTorrent) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedBitTorrentServer) mustEmbedUnimplementedBitTorrentServer() {}

func RegisterBitTorrentServer(s *grpc.Server, srv BitTorrentServer) {
	s.RegisterService(&_BitTorrent_serviceDesc, srv)
}

func _BitTorrent_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BitTorrent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitTorrentServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.bitTorrent/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitTorrentServer).Send(ctx, req.(*BitTorrent))
	}
	return interceptor(ctx, in, info, handler)
}

var _BitTorrent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.bitTorrent",
	HandlerType: (*BitTorrentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send",
			Handler:    _BitTorrent_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/bittorrent.proto",
}
