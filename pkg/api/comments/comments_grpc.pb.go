// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/comments/comments.proto

package comments

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Comments_CreateComment_FullMethodName = "/comments.Comments/CreateComment"
	Comments_GetComments_FullMethodName   = "/comments.Comments/GetComments"
)

// CommentsClient is the client API for Comments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentsClient interface {
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetCommentsResponse], error)
}

type commentsClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentsClient(cc grpc.ClientConnInterface) CommentsClient {
	return &commentsClient{cc}
}

func (c *commentsClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, Comments_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetCommentsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Comments_ServiceDesc.Streams[0], Comments_GetComments_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetCommentsRequest, GetCommentsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Comments_GetCommentsClient = grpc.ServerStreamingClient[GetCommentsResponse]

// CommentsServer is the server API for Comments service.
// All implementations must embed UnimplementedCommentsServer
// for forward compatibility.
type CommentsServer interface {
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	GetComments(*GetCommentsRequest, grpc.ServerStreamingServer[GetCommentsResponse]) error
	mustEmbedUnimplementedCommentsServer()
}

// UnimplementedCommentsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentsServer struct{}

func (UnimplementedCommentsServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentsServer) GetComments(*GetCommentsRequest, grpc.ServerStreamingServer[GetCommentsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedCommentsServer) mustEmbedUnimplementedCommentsServer() {}
func (UnimplementedCommentsServer) testEmbeddedByValue()                  {}

// UnsafeCommentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentsServer will
// result in compilation errors.
type UnsafeCommentsServer interface {
	mustEmbedUnimplementedCommentsServer()
}

func RegisterCommentsServer(s grpc.ServiceRegistrar, srv CommentsServer) {
	// If the following call pancis, it indicates UnimplementedCommentsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Comments_ServiceDesc, srv)
}

func _Comments_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comments_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_GetComments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCommentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommentsServer).GetComments(m, &grpc.GenericServerStream[GetCommentsRequest, GetCommentsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Comments_GetCommentsServer = grpc.ServerStreamingServer[GetCommentsResponse]

// Comments_ServiceDesc is the grpc.ServiceDesc for Comments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comments.Comments",
	HandlerType: (*CommentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _Comments_CreateComment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetComments",
			Handler:       _Comments_GetComments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/comments/comments.proto",
}
