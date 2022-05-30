// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api.proto

package api

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

// NewsfeedConfiguratorClient is the client API for NewsfeedConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsfeedConfiguratorClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUserFilter(ctx context.Context, in *UpdateUserFilterRequest, opts ...grpc.CallOption) (*UpdateUserFilterResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	AddSource(ctx context.Context, in *AddSourceRequest, opts ...grpc.CallOption) (*AddSourceResponse, error)
	GetSource(ctx context.Context, in *GetSourceRequest, opts ...grpc.CallOption) (*GetSourceResponse, error)
	GetUserSources(ctx context.Context, in *GetUserSourcesRequest, opts ...grpc.CallOption) (*GetUserSourcesResponse, error)
	UpdateSourceName(ctx context.Context, in *UpdateSourceNameRequest, opts ...grpc.CallOption) (*UpdateSourceNameResponse, error)
	UpdateSourceIsActive(ctx context.Context, in *UpdateSourceIsActiveRequest, opts ...grpc.CallOption) (*UpdateSourceIsActiveResponse, error)
	DeleteSource(ctx context.Context, in *DeleteSourceRequest, opts ...grpc.CallOption) (*DeleteSourceResponse, error)
}

type newsfeedConfiguratorClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsfeedConfiguratorClient(cc grpc.ClientConnInterface) NewsfeedConfiguratorClient {
	return &newsfeedConfiguratorClient{cc}
}

func (c *newsfeedConfiguratorClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/addUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) UpdateUserFilter(ctx context.Context, in *UpdateUserFilterRequest, opts ...grpc.CallOption) (*UpdateUserFilterResponse, error) {
	out := new(UpdateUserFilterResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/updateUserFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/deleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) AddSource(ctx context.Context, in *AddSourceRequest, opts ...grpc.CallOption) (*AddSourceResponse, error) {
	out := new(AddSourceResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/addSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) GetSource(ctx context.Context, in *GetSourceRequest, opts ...grpc.CallOption) (*GetSourceResponse, error) {
	out := new(GetSourceResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/getSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) GetUserSources(ctx context.Context, in *GetUserSourcesRequest, opts ...grpc.CallOption) (*GetUserSourcesResponse, error) {
	out := new(GetUserSourcesResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/getUserSources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) UpdateSourceName(ctx context.Context, in *UpdateSourceNameRequest, opts ...grpc.CallOption) (*UpdateSourceNameResponse, error) {
	out := new(UpdateSourceNameResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/updateSourceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) UpdateSourceIsActive(ctx context.Context, in *UpdateSourceIsActiveRequest, opts ...grpc.CallOption) (*UpdateSourceIsActiveResponse, error) {
	out := new(UpdateSourceIsActiveResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/updateSourceIsActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsfeedConfiguratorClient) DeleteSource(ctx context.Context, in *DeleteSourceRequest, opts ...grpc.CallOption) (*DeleteSourceResponse, error) {
	out := new(DeleteSourceResponse)
	err := c.cc.Invoke(ctx, "/api.NewsfeedConfigurator/deleteSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsfeedConfiguratorServer is the server API for NewsfeedConfigurator service.
// All implementations must embed UnimplementedNewsfeedConfiguratorServer
// for forward compatibility
type NewsfeedConfiguratorServer interface {
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUserFilter(context.Context, *UpdateUserFilterRequest) (*UpdateUserFilterResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	AddSource(context.Context, *AddSourceRequest) (*AddSourceResponse, error)
	GetSource(context.Context, *GetSourceRequest) (*GetSourceResponse, error)
	GetUserSources(context.Context, *GetUserSourcesRequest) (*GetUserSourcesResponse, error)
	UpdateSourceName(context.Context, *UpdateSourceNameRequest) (*UpdateSourceNameResponse, error)
	UpdateSourceIsActive(context.Context, *UpdateSourceIsActiveRequest) (*UpdateSourceIsActiveResponse, error)
	DeleteSource(context.Context, *DeleteSourceRequest) (*DeleteSourceResponse, error)
	mustEmbedUnimplementedNewsfeedConfiguratorServer()
}

// UnimplementedNewsfeedConfiguratorServer must be embedded to have forward compatible implementations.
type UnimplementedNewsfeedConfiguratorServer struct {
}

func (UnimplementedNewsfeedConfiguratorServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) UpdateUserFilter(context.Context, *UpdateUserFilterRequest) (*UpdateUserFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserFilter not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) AddSource(context.Context, *AddSourceRequest) (*AddSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSource not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) GetSource(context.Context, *GetSourceRequest) (*GetSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSource not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) GetUserSources(context.Context, *GetUserSourcesRequest) (*GetUserSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSources not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) UpdateSourceName(context.Context, *UpdateSourceNameRequest) (*UpdateSourceNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSourceName not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) UpdateSourceIsActive(context.Context, *UpdateSourceIsActiveRequest) (*UpdateSourceIsActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSourceIsActive not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) DeleteSource(context.Context, *DeleteSourceRequest) (*DeleteSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSource not implemented")
}
func (UnimplementedNewsfeedConfiguratorServer) mustEmbedUnimplementedNewsfeedConfiguratorServer() {}

// UnsafeNewsfeedConfiguratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsfeedConfiguratorServer will
// result in compilation errors.
type UnsafeNewsfeedConfiguratorServer interface {
	mustEmbedUnimplementedNewsfeedConfiguratorServer()
}

func RegisterNewsfeedConfiguratorServer(s grpc.ServiceRegistrar, srv NewsfeedConfiguratorServer) {
	s.RegisterService(&NewsfeedConfigurator_ServiceDesc, srv)
}

func _NewsfeedConfigurator_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/addUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/getUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_UpdateUserFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).UpdateUserFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/updateUserFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).UpdateUserFilter(ctx, req.(*UpdateUserFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/deleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_AddSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).AddSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/addSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).AddSource(ctx, req.(*AddSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_GetSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).GetSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/getSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).GetSource(ctx, req.(*GetSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_GetUserSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).GetUserSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/getUserSources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).GetUserSources(ctx, req.(*GetUserSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_UpdateSourceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSourceNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).UpdateSourceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/updateSourceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).UpdateSourceName(ctx, req.(*UpdateSourceNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_UpdateSourceIsActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSourceIsActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).UpdateSourceIsActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/updateSourceIsActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).UpdateSourceIsActive(ctx, req.(*UpdateSourceIsActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsfeedConfigurator_DeleteSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsfeedConfiguratorServer).DeleteSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NewsfeedConfigurator/deleteSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsfeedConfiguratorServer).DeleteSource(ctx, req.(*DeleteSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsfeedConfigurator_ServiceDesc is the grpc.ServiceDesc for NewsfeedConfigurator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsfeedConfigurator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.NewsfeedConfigurator",
	HandlerType: (*NewsfeedConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addUser",
			Handler:    _NewsfeedConfigurator_AddUser_Handler,
		},
		{
			MethodName: "getUser",
			Handler:    _NewsfeedConfigurator_GetUser_Handler,
		},
		{
			MethodName: "updateUserFilter",
			Handler:    _NewsfeedConfigurator_UpdateUserFilter_Handler,
		},
		{
			MethodName: "deleteUser",
			Handler:    _NewsfeedConfigurator_DeleteUser_Handler,
		},
		{
			MethodName: "addSource",
			Handler:    _NewsfeedConfigurator_AddSource_Handler,
		},
		{
			MethodName: "getSource",
			Handler:    _NewsfeedConfigurator_GetSource_Handler,
		},
		{
			MethodName: "getUserSources",
			Handler:    _NewsfeedConfigurator_GetUserSources_Handler,
		},
		{
			MethodName: "updateSourceName",
			Handler:    _NewsfeedConfigurator_UpdateSourceName_Handler,
		},
		{
			MethodName: "updateSourceIsActive",
			Handler:    _NewsfeedConfigurator_UpdateSourceIsActive_Handler,
		},
		{
			MethodName: "deleteSource",
			Handler:    _NewsfeedConfigurator_DeleteSource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
