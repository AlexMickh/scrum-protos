// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/projects/projects.proto

package projects

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Projects_CreateProject_FullMethodName  = "/projects.Projects/CreateProject"
	Projects_GetProject_FullMethodName     = "/projects.Projects/GetProject"
	Projects_AddParticipant_FullMethodName = "/projects.Projects/AddParticipant"
	Projects_DeleteProject_FullMethodName  = "/projects.Projects/DeleteProject"
)

// ProjectsClient is the client API for Projects service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectsClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error)
	AddParticipant(ctx context.Context, in *AddParticipantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type projectsClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectsClient(cc grpc.ClientConnInterface) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, Projects_CreateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProjectResponse)
	err := c.cc.Invoke(ctx, Projects_GetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) AddParticipant(ctx context.Context, in *AddParticipantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Projects_AddParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Projects_DeleteProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
// All implementations must embed UnimplementedProjectsServer
// for forward compatibility.
type ProjectsServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error)
	AddParticipant(context.Context, *AddParticipantRequest) (*emptypb.Empty, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedProjectsServer()
}

// UnimplementedProjectsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProjectsServer struct{}

func (UnimplementedProjectsServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectsServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectsServer) AddParticipant(context.Context, *AddParticipantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParticipant not implemented")
}
func (UnimplementedProjectsServer) DeleteProject(context.Context, *DeleteProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectsServer) mustEmbedUnimplementedProjectsServer() {}
func (UnimplementedProjectsServer) testEmbeddedByValue()                  {}

// UnsafeProjectsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectsServer will
// result in compilation errors.
type UnsafeProjectsServer interface {
	mustEmbedUnimplementedProjectsServer()
}

func RegisterProjectsServer(s grpc.ServiceRegistrar, srv ProjectsServer) {
	// If the following call pancis, it indicates UnimplementedProjectsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Projects_ServiceDesc, srv)
}

func _Projects_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_AddParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).AddParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_AddParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).AddParticipant(ctx, req.(*AddParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Projects_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Projects_ServiceDesc is the grpc.ServiceDesc for Projects service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Projects_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "projects.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Projects_CreateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Projects_GetProject_Handler,
		},
		{
			MethodName: "AddParticipant",
			Handler:    _Projects_AddParticipant_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _Projects_DeleteProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/projects/projects.proto",
}
