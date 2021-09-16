// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package storiespb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StoriesClient is the client API for Stories service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoriesClient interface {
	// WriteStory ...
	WriteStory(ctx context.Context, in *RequestWriteStory, opts ...grpc.CallOption) (*ResponseID, error)
	// EditStory actualiza los datos de la historia
	EditStory(ctx context.Context, in *RequestEditStory, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ChangeFinished pone o quita si una historia llego a su fin
	ChangeFinished(ctx context.Context, in *RequestChangeFinished, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// RemoveStory ...
	RemoveStory(ctx context.Context, in *RequestRemoveStory, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// MyStories lista las historias del usuario logueado
	MyStories(ctx context.Context, in *RequestMyStories, opts ...grpc.CallOption) (Stories_MyStoriesClient, error)
	// ListStories lista historias
	ListStories(ctx context.Context, in *RequestListStories, opts ...grpc.CallOption) (Stories_ListStoriesClient, error)
	// ChangeStatus cambia el estado de la historia y/o capitulos.
	//
	// Acciones:
	// * Cambia el estado de la historia y todos sus capitulos si ChaptersID esta
	// vacio.
	// * Cambia el estado de los capitulos q vienen en ChaptersID y si corresponde
	// tambien se cambiara el estado de la historia.
	ChangeStatus(ctx context.Context, in *RequestChangeStoryStatus, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ChangeCover cambia el cover a  la historia y/o capitulos segun los datos de
	// la peticion
	ChangeCover(ctx context.Context, in *RequestChangeCover, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// WriteChapter ...
	WriteChapter(ctx context.Context, in *RequestWriteChapter, opts ...grpc.CallOption) (*ResponseID, error)
	// EditChapter actualiza los datos de un capitulo
	EditChapter(ctx context.Context, in *RequestEditChapter, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ReorderChapters ...
	ReorderChapters(ctx context.Context, in *RequestReorderChapters, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// RemoveChapter ...
	RemoveChapter(ctx context.Context, in *RequestRemoveChapter, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// RemoveChapterCover ...
	RemoveChapterCover(ctx context.Context, in *RequestRemoveChapterCover, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ReadChapter leer un capitulo de una historia
	ReadChapter(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*ResponseReadChapter, error)
}

type storiesClient struct {
	cc grpc.ClientConnInterface
}

func NewStoriesClient(cc grpc.ClientConnInterface) StoriesClient {
	return &storiesClient{cc}
}

func (c *storiesClient) WriteStory(ctx context.Context, in *RequestWriteStory, opts ...grpc.CallOption) (*ResponseID, error) {
	out := new(ResponseID)
	err := c.cc.Invoke(ctx, "/stories.stories/WriteStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) EditStory(ctx context.Context, in *RequestEditStory, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/EditStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) ChangeFinished(ctx context.Context, in *RequestChangeFinished, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/ChangeFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) RemoveStory(ctx context.Context, in *RequestRemoveStory, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/RemoveStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) MyStories(ctx context.Context, in *RequestMyStories, opts ...grpc.CallOption) (Stories_MyStoriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stories_ServiceDesc.Streams[0], "/stories.stories/MyStories", opts...)
	if err != nil {
		return nil, err
	}
	x := &storiesMyStoriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stories_MyStoriesClient interface {
	Recv() (*ResponseMyStories, error)
	grpc.ClientStream
}

type storiesMyStoriesClient struct {
	grpc.ClientStream
}

func (x *storiesMyStoriesClient) Recv() (*ResponseMyStories, error) {
	m := new(ResponseMyStories)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storiesClient) ListStories(ctx context.Context, in *RequestListStories, opts ...grpc.CallOption) (Stories_ListStoriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stories_ServiceDesc.Streams[1], "/stories.stories/ListStories", opts...)
	if err != nil {
		return nil, err
	}
	x := &storiesListStoriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stories_ListStoriesClient interface {
	Recv() (*ResponseListStories, error)
	grpc.ClientStream
}

type storiesListStoriesClient struct {
	grpc.ClientStream
}

func (x *storiesListStoriesClient) Recv() (*ResponseListStories, error) {
	m := new(ResponseListStories)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storiesClient) ChangeStatus(ctx context.Context, in *RequestChangeStoryStatus, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) ChangeCover(ctx context.Context, in *RequestChangeCover, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/ChangeCover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) WriteChapter(ctx context.Context, in *RequestWriteChapter, opts ...grpc.CallOption) (*ResponseID, error) {
	out := new(ResponseID)
	err := c.cc.Invoke(ctx, "/stories.stories/WriteChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) EditChapter(ctx context.Context, in *RequestEditChapter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/EditChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) ReorderChapters(ctx context.Context, in *RequestReorderChapters, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/ReorderChapters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) RemoveChapter(ctx context.Context, in *RequestRemoveChapter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/RemoveChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) RemoveChapterCover(ctx context.Context, in *RequestRemoveChapterCover, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/stories.stories/RemoveChapterCover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storiesClient) ReadChapter(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*ResponseReadChapter, error) {
	out := new(ResponseReadChapter)
	err := c.cc.Invoke(ctx, "/stories.stories/ReadChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoriesServer is the server API for Stories service.
// All implementations must embed UnimplementedStoriesServer
// for forward compatibility
type StoriesServer interface {
	// WriteStory ...
	WriteStory(context.Context, *RequestWriteStory) (*ResponseID, error)
	// EditStory actualiza los datos de la historia
	EditStory(context.Context, *RequestEditStory) (*emptypb.Empty, error)
	// ChangeFinished pone o quita si una historia llego a su fin
	ChangeFinished(context.Context, *RequestChangeFinished) (*emptypb.Empty, error)
	// RemoveStory ...
	RemoveStory(context.Context, *RequestRemoveStory) (*emptypb.Empty, error)
	// MyStories lista las historias del usuario logueado
	MyStories(*RequestMyStories, Stories_MyStoriesServer) error
	// ListStories lista historias
	ListStories(*RequestListStories, Stories_ListStoriesServer) error
	// ChangeStatus cambia el estado de la historia y/o capitulos.
	//
	// Acciones:
	// * Cambia el estado de la historia y todos sus capitulos si ChaptersID esta
	// vacio.
	// * Cambia el estado de los capitulos q vienen en ChaptersID y si corresponde
	// tambien se cambiara el estado de la historia.
	ChangeStatus(context.Context, *RequestChangeStoryStatus) (*emptypb.Empty, error)
	// ChangeCover cambia el cover a  la historia y/o capitulos segun los datos de
	// la peticion
	ChangeCover(context.Context, *RequestChangeCover) (*emptypb.Empty, error)
	// WriteChapter ...
	WriteChapter(context.Context, *RequestWriteChapter) (*ResponseID, error)
	// EditChapter actualiza los datos de un capitulo
	EditChapter(context.Context, *RequestEditChapter) (*emptypb.Empty, error)
	// ReorderChapters ...
	ReorderChapters(context.Context, *RequestReorderChapters) (*emptypb.Empty, error)
	// RemoveChapter ...
	RemoveChapter(context.Context, *RequestRemoveChapter) (*emptypb.Empty, error)
	// RemoveChapterCover ...
	RemoveChapterCover(context.Context, *RequestRemoveChapterCover) (*emptypb.Empty, error)
	// ReadChapter leer un capitulo de una historia
	ReadChapter(context.Context, *RequestID) (*ResponseReadChapter, error)
	mustEmbedUnimplementedStoriesServer()
}

// UnimplementedStoriesServer must be embedded to have forward compatible implementations.
type UnimplementedStoriesServer struct {
}

func (UnimplementedStoriesServer) WriteStory(context.Context, *RequestWriteStory) (*ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteStory not implemented")
}
func (UnimplementedStoriesServer) EditStory(context.Context, *RequestEditStory) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditStory not implemented")
}
func (UnimplementedStoriesServer) ChangeFinished(context.Context, *RequestChangeFinished) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeFinished not implemented")
}
func (UnimplementedStoriesServer) RemoveStory(context.Context, *RequestRemoveStory) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStory not implemented")
}
func (UnimplementedStoriesServer) MyStories(*RequestMyStories, Stories_MyStoriesServer) error {
	return status.Errorf(codes.Unimplemented, "method MyStories not implemented")
}
func (UnimplementedStoriesServer) ListStories(*RequestListStories, Stories_ListStoriesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStories not implemented")
}
func (UnimplementedStoriesServer) ChangeStatus(context.Context, *RequestChangeStoryStatus) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedStoriesServer) ChangeCover(context.Context, *RequestChangeCover) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCover not implemented")
}
func (UnimplementedStoriesServer) WriteChapter(context.Context, *RequestWriteChapter) (*ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteChapter not implemented")
}
func (UnimplementedStoriesServer) EditChapter(context.Context, *RequestEditChapter) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditChapter not implemented")
}
func (UnimplementedStoriesServer) ReorderChapters(context.Context, *RequestReorderChapters) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReorderChapters not implemented")
}
func (UnimplementedStoriesServer) RemoveChapter(context.Context, *RequestRemoveChapter) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveChapter not implemented")
}
func (UnimplementedStoriesServer) RemoveChapterCover(context.Context, *RequestRemoveChapterCover) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveChapterCover not implemented")
}
func (UnimplementedStoriesServer) ReadChapter(context.Context, *RequestID) (*ResponseReadChapter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadChapter not implemented")
}
func (UnimplementedStoriesServer) mustEmbedUnimplementedStoriesServer() {}

// UnsafeStoriesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoriesServer will
// result in compilation errors.
type UnsafeStoriesServer interface {
	mustEmbedUnimplementedStoriesServer()
}

func RegisterStoriesServer(s grpc.ServiceRegistrar, srv StoriesServer) {
	s.RegisterService(&Stories_ServiceDesc, srv)
}

func _Stories_WriteStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWriteStory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).WriteStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/WriteStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).WriteStory(ctx, req.(*RequestWriteStory))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_EditStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditStory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).EditStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/EditStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).EditStory(ctx, req.(*RequestEditStory))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_ChangeFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChangeFinished)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).ChangeFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/ChangeFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).ChangeFinished(ctx, req.(*RequestChangeFinished))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_RemoveStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRemoveStory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).RemoveStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/RemoveStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).RemoveStory(ctx, req.(*RequestRemoveStory))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_MyStories_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestMyStories)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StoriesServer).MyStories(m, &storiesMyStoriesServer{stream})
}

type Stories_MyStoriesServer interface {
	Send(*ResponseMyStories) error
	grpc.ServerStream
}

type storiesMyStoriesServer struct {
	grpc.ServerStream
}

func (x *storiesMyStoriesServer) Send(m *ResponseMyStories) error {
	return x.ServerStream.SendMsg(m)
}

func _Stories_ListStories_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestListStories)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StoriesServer).ListStories(m, &storiesListStoriesServer{stream})
}

type Stories_ListStoriesServer interface {
	Send(*ResponseListStories) error
	grpc.ServerStream
}

type storiesListStoriesServer struct {
	grpc.ServerStream
}

func (x *storiesListStoriesServer) Send(m *ResponseListStories) error {
	return x.ServerStream.SendMsg(m)
}

func _Stories_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChangeStoryStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).ChangeStatus(ctx, req.(*RequestChangeStoryStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_ChangeCover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestChangeCover)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).ChangeCover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/ChangeCover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).ChangeCover(ctx, req.(*RequestChangeCover))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_WriteChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWriteChapter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).WriteChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/WriteChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).WriteChapter(ctx, req.(*RequestWriteChapter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_EditChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditChapter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).EditChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/EditChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).EditChapter(ctx, req.(*RequestEditChapter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_ReorderChapters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReorderChapters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).ReorderChapters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/ReorderChapters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).ReorderChapters(ctx, req.(*RequestReorderChapters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_RemoveChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRemoveChapter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).RemoveChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/RemoveChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).RemoveChapter(ctx, req.(*RequestRemoveChapter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_RemoveChapterCover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRemoveChapterCover)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).RemoveChapterCover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/RemoveChapterCover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).RemoveChapterCover(ctx, req.(*RequestRemoveChapterCover))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stories_ReadChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoriesServer).ReadChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stories.stories/ReadChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoriesServer).ReadChapter(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

// Stories_ServiceDesc is the grpc.ServiceDesc for Stories service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stories_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stories.stories",
	HandlerType: (*StoriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteStory",
			Handler:    _Stories_WriteStory_Handler,
		},
		{
			MethodName: "EditStory",
			Handler:    _Stories_EditStory_Handler,
		},
		{
			MethodName: "ChangeFinished",
			Handler:    _Stories_ChangeFinished_Handler,
		},
		{
			MethodName: "RemoveStory",
			Handler:    _Stories_RemoveStory_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _Stories_ChangeStatus_Handler,
		},
		{
			MethodName: "ChangeCover",
			Handler:    _Stories_ChangeCover_Handler,
		},
		{
			MethodName: "WriteChapter",
			Handler:    _Stories_WriteChapter_Handler,
		},
		{
			MethodName: "EditChapter",
			Handler:    _Stories_EditChapter_Handler,
		},
		{
			MethodName: "ReorderChapters",
			Handler:    _Stories_ReorderChapters_Handler,
		},
		{
			MethodName: "RemoveChapter",
			Handler:    _Stories_RemoveChapter_Handler,
		},
		{
			MethodName: "RemoveChapterCover",
			Handler:    _Stories_RemoveChapterCover_Handler,
		},
		{
			MethodName: "ReadChapter",
			Handler:    _Stories_ReadChapter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MyStories",
			Handler:       _Stories_MyStories_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListStories",
			Handler:       _Stories_ListStories_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stories.proto",
}
