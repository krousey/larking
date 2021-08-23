// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testpb

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingClient interface {
	// HTTP | gRPC
	// -----|-----
	// `GET /v1/messages/123456`  | `GetMessageOne(name: "messages/123456")`
	GetMessageOne(ctx context.Context, in *GetMessageRequestOne, opts ...grpc.CallOption) (*Message, error)
	// HTTP | gRPC
	// -----|-----
	// `GET /v1/messages/123456?revision=2&sub.subfield=foo` |
	// `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield:
	// "foo"))`
	// `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id:
	// "123456")`
	GetMessageTwo(ctx context.Context, in *GetMessageRequestTwo, opts ...grpc.CallOption) (*Message, error)
	// HTTP | gRPC
	// -----|-----
	// `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
	// "123456" message { text: "Hi!" })`
	UpdateMessage(ctx context.Context, in *UpdateMessageRequestOne, opts ...grpc.CallOption) (*Message, error)
	// HTTP | gRPC
	// -----|-----
	// `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
	// "123456" text: "Hi!")`
	UpdateMessageBody(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Action(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ActionSegment(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ActionSegments(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BatchGet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VariableOne(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VariableTwo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) GetMessageOne(ctx context.Context, in *GetMessageRequestOne, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/GetMessageOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) GetMessageTwo(ctx context.Context, in *GetMessageRequestTwo, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/GetMessageTwo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequestOne, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/UpdateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) UpdateMessageBody(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/UpdateMessageBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) Action(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) ActionSegment(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/ActionSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) ActionSegments(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/ActionSegments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) BatchGet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/BatchGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) VariableOne(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/VariableOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) VariableTwo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/VariableTwo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/GetShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/larking.testpb.Messaging/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServer is the server API for Messaging service.
// All implementations must embed UnimplementedMessagingServer
// for forward compatibility
type MessagingServer interface {
	// HTTP | gRPC
	// -----|-----
	// `GET /v1/messages/123456`  | `GetMessageOne(name: "messages/123456")`
	GetMessageOne(context.Context, *GetMessageRequestOne) (*Message, error)
	// HTTP | gRPC
	// -----|-----
	// `GET /v1/messages/123456?revision=2&sub.subfield=foo` |
	// `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield:
	// "foo"))`
	// `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id:
	// "123456")`
	GetMessageTwo(context.Context, *GetMessageRequestTwo) (*Message, error)
	// HTTP | gRPC
	// -----|-----
	// `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
	// "123456" message { text: "Hi!" })`
	UpdateMessage(context.Context, *UpdateMessageRequestOne) (*Message, error)
	// HTTP | gRPC
	// -----|-----
	// `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
	// "123456" text: "Hi!")`
	UpdateMessageBody(context.Context, *Message) (*Message, error)
	Action(context.Context, *Message) (*emptypb.Empty, error)
	ActionSegment(context.Context, *Message) (*emptypb.Empty, error)
	ActionSegments(context.Context, *Message) (*emptypb.Empty, error)
	BatchGet(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	VariableOne(context.Context, *Message) (*emptypb.Empty, error)
	VariableTwo(context.Context, *Message) (*emptypb.Empty, error)
	GetShelf(context.Context, *GetShelfRequest) (*Shelf, error)
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	mustEmbedUnimplementedMessagingServer()
}

// UnimplementedMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingServer struct {
}

func (UnimplementedMessagingServer) GetMessageOne(context.Context, *GetMessageRequestOne) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageOne not implemented")
}
func (UnimplementedMessagingServer) GetMessageTwo(context.Context, *GetMessageRequestTwo) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageTwo not implemented")
}
func (UnimplementedMessagingServer) UpdateMessage(context.Context, *UpdateMessageRequestOne) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessagingServer) UpdateMessageBody(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageBody not implemented")
}
func (UnimplementedMessagingServer) Action(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}
func (UnimplementedMessagingServer) ActionSegment(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionSegment not implemented")
}
func (UnimplementedMessagingServer) ActionSegments(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionSegments not implemented")
}
func (UnimplementedMessagingServer) BatchGet(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (UnimplementedMessagingServer) VariableOne(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VariableOne not implemented")
}
func (UnimplementedMessagingServer) VariableTwo(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VariableTwo not implemented")
}
func (UnimplementedMessagingServer) GetShelf(context.Context, *GetShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}
func (UnimplementedMessagingServer) GetBook(context.Context, *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedMessagingServer) CreateBook(context.Context, *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedMessagingServer) UpdateBook(context.Context, *UpdateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedMessagingServer) mustEmbedUnimplementedMessagingServer() {}

// UnsafeMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServer will
// result in compilation errors.
type UnsafeMessagingServer interface {
	mustEmbedUnimplementedMessagingServer()
}

func RegisterMessagingServer(s grpc.ServiceRegistrar, srv MessagingServer) {
	s.RegisterService(&Messaging_ServiceDesc, srv)
}

func _Messaging_GetMessageOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequestOne)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetMessageOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/GetMessageOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetMessageOne(ctx, req.(*GetMessageRequestOne))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_GetMessageTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequestTwo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetMessageTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/GetMessageTwo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetMessageTwo(ctx, req.(*GetMessageRequestTwo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequestOne)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/UpdateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).UpdateMessage(ctx, req.(*UpdateMessageRequestOne))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_UpdateMessageBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).UpdateMessageBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/UpdateMessageBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).UpdateMessageBody(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).Action(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_ActionSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).ActionSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/ActionSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).ActionSegment(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_ActionSegments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).ActionSegments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/ActionSegments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).ActionSegments(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/BatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).BatchGet(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_VariableOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).VariableOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/VariableOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).VariableOne(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_VariableTwo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).VariableTwo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/VariableTwo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).VariableTwo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/GetShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Messaging/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Messaging_ServiceDesc is the grpc.ServiceDesc for Messaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "larking.testpb.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessageOne",
			Handler:    _Messaging_GetMessageOne_Handler,
		},
		{
			MethodName: "GetMessageTwo",
			Handler:    _Messaging_GetMessageTwo_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _Messaging_UpdateMessage_Handler,
		},
		{
			MethodName: "UpdateMessageBody",
			Handler:    _Messaging_UpdateMessageBody_Handler,
		},
		{
			MethodName: "Action",
			Handler:    _Messaging_Action_Handler,
		},
		{
			MethodName: "ActionSegment",
			Handler:    _Messaging_ActionSegment_Handler,
		},
		{
			MethodName: "ActionSegments",
			Handler:    _Messaging_ActionSegments_Handler,
		},
		{
			MethodName: "BatchGet",
			Handler:    _Messaging_BatchGet_Handler,
		},
		{
			MethodName: "VariableOne",
			Handler:    _Messaging_VariableOne_Handler,
		},
		{
			MethodName: "VariableTwo",
			Handler:    _Messaging_VariableTwo_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _Messaging_GetShelf_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Messaging_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _Messaging_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _Messaging_UpdateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testpb/test.proto",
}

// FilesClient is the client API for Files service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesClient interface {
	// HTTP | gRPC
	// -----|-----
	// `POST /files/cat.jpg <body>` | `UploadDownload(filename: "cat.jpg", file:
	// { content_type: "image/jpeg", data: <body>})"`
	UploadDownload(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	LargeUploadDownload(ctx context.Context, opts ...grpc.CallOption) (Files_LargeUploadDownloadClient, error)
}

type filesClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesClient(cc grpc.ClientConnInterface) FilesClient {
	return &filesClient{cc}
}

func (c *filesClient) UploadDownload(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/larking.testpb.Files/UploadDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesClient) LargeUploadDownload(ctx context.Context, opts ...grpc.CallOption) (Files_LargeUploadDownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Files_ServiceDesc.Streams[0], "/larking.testpb.Files/LargeUploadDownload", opts...)
	if err != nil {
		return nil, err
	}
	x := &filesLargeUploadDownloadClient{stream}
	return x, nil
}

type Files_LargeUploadDownloadClient interface {
	Send(*UploadFileRequest) error
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type filesLargeUploadDownloadClient struct {
	grpc.ClientStream
}

func (x *filesLargeUploadDownloadClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *filesLargeUploadDownloadClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilesServer is the server API for Files service.
// All implementations must embed UnimplementedFilesServer
// for forward compatibility
type FilesServer interface {
	// HTTP | gRPC
	// -----|-----
	// `POST /files/cat.jpg <body>` | `UploadDownload(filename: "cat.jpg", file:
	// { content_type: "image/jpeg", data: <body>})"`
	UploadDownload(context.Context, *UploadFileRequest) (*httpbody.HttpBody, error)
	LargeUploadDownload(Files_LargeUploadDownloadServer) error
	mustEmbedUnimplementedFilesServer()
}

// UnimplementedFilesServer must be embedded to have forward compatible implementations.
type UnimplementedFilesServer struct {
}

func (UnimplementedFilesServer) UploadDownload(context.Context, *UploadFileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadDownload not implemented")
}
func (UnimplementedFilesServer) LargeUploadDownload(Files_LargeUploadDownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method LargeUploadDownload not implemented")
}
func (UnimplementedFilesServer) mustEmbedUnimplementedFilesServer() {}

// UnsafeFilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesServer will
// result in compilation errors.
type UnsafeFilesServer interface {
	mustEmbedUnimplementedFilesServer()
}

func RegisterFilesServer(s grpc.ServiceRegistrar, srv FilesServer) {
	s.RegisterService(&Files_ServiceDesc, srv)
}

func _Files_UploadDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).UploadDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.Files/UploadDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).UploadDownload(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Files_LargeUploadDownload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FilesServer).LargeUploadDownload(&filesLargeUploadDownloadServer{stream})
}

type Files_LargeUploadDownloadServer interface {
	Send(*httpbody.HttpBody) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type filesLargeUploadDownloadServer struct {
	grpc.ServerStream
}

func (x *filesLargeUploadDownloadServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

func (x *filesLargeUploadDownloadServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Files_ServiceDesc is the grpc.ServiceDesc for Files service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Files_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "larking.testpb.Files",
	HandlerType: (*FilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadDownload",
			Handler:    _Files_UploadDownload_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LargeUploadDownload",
			Handler:       _Files_LargeUploadDownload_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "testpb/test.proto",
}

// WellKnownClient is the client API for WellKnown service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WellKnownClient interface {
	// HTTP | gRPC
	// -----|-----
	// `GET /v1/wellknown/timestamp/2017-01-15T01:30:15.01Z` |
	// `Check(Timestamp{...})`
	Check(ctx context.Context, in *Scalars, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type wellKnownClient struct {
	cc grpc.ClientConnInterface
}

func NewWellKnownClient(cc grpc.ClientConnInterface) WellKnownClient {
	return &wellKnownClient{cc}
}

func (c *wellKnownClient) Check(ctx context.Context, in *Scalars, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/larking.testpb.WellKnown/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WellKnownServer is the server API for WellKnown service.
// All implementations must embed UnimplementedWellKnownServer
// for forward compatibility
type WellKnownServer interface {
	// HTTP | gRPC
	// -----|-----
	// `GET /v1/wellknown/timestamp/2017-01-15T01:30:15.01Z` |
	// `Check(Timestamp{...})`
	Check(context.Context, *Scalars) (*emptypb.Empty, error)
	mustEmbedUnimplementedWellKnownServer()
}

// UnimplementedWellKnownServer must be embedded to have forward compatible implementations.
type UnimplementedWellKnownServer struct {
}

func (UnimplementedWellKnownServer) Check(context.Context, *Scalars) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedWellKnownServer) mustEmbedUnimplementedWellKnownServer() {}

// UnsafeWellKnownServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WellKnownServer will
// result in compilation errors.
type UnsafeWellKnownServer interface {
	mustEmbedUnimplementedWellKnownServer()
}

func RegisterWellKnownServer(s grpc.ServiceRegistrar, srv WellKnownServer) {
	s.RegisterService(&WellKnown_ServiceDesc, srv)
}

func _WellKnown_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Scalars)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WellKnownServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/larking.testpb.WellKnown/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WellKnownServer).Check(ctx, req.(*Scalars))
	}
	return interceptor(ctx, in, info, handler)
}

// WellKnown_ServiceDesc is the grpc.ServiceDesc for WellKnown service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WellKnown_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "larking.testpb.WellKnown",
	HandlerType: (*WellKnownServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _WellKnown_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testpb/test.proto",
}
