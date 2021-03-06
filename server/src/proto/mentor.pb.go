// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mentor.proto

/*
Package mentor is a generated protocol buffer package.

It is generated from these files:
	mentor.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	Lesson
	ListLessonsResponse
	ListLessonsRequest
*/
package mentor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	IsSuccessful bool `protobuf:"varint,1,opt,name=isSuccessful" json:"isSuccessful,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetIsSuccessful() bool {
	if m != nil {
		return m.IsSuccessful
	}
	return false
}

type Lesson struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Score int32  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	Title string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	By    string `protobuf:"bytes,4,opt,name=by" json:"by,omitempty"`
	Time  int32  `protobuf:"varint,5,opt,name=time" json:"time,omitempty"`
	Url   string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
}

func (m *Lesson) Reset()                    { *m = Lesson{} }
func (m *Lesson) String() string            { return proto.CompactTextString(m) }
func (*Lesson) ProtoMessage()               {}
func (*Lesson) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Lesson) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Lesson) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Lesson) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Lesson) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

func (m *Lesson) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Lesson) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ListLessonsResponse struct {
	Lesson *Lesson `protobuf:"bytes,1,opt,name=lesson" json:"lesson,omitempty"`
}

func (m *ListLessonsResponse) Reset()                    { *m = ListLessonsResponse{} }
func (m *ListLessonsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLessonsResponse) ProtoMessage()               {}
func (*ListLessonsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListLessonsResponse) GetLesson() *Lesson {
	if m != nil {
		return m.Lesson
	}
	return nil
}

type ListLessonsRequest struct {
}

func (m *ListLessonsRequest) Reset()                    { *m = ListLessonsRequest{} }
func (m *ListLessonsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLessonsRequest) ProtoMessage()               {}
func (*ListLessonsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "mentor.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "mentor.LoginResponse")
	proto.RegisterType((*Lesson)(nil), "mentor.Lesson")
	proto.RegisterType((*ListLessonsResponse)(nil), "mentor.ListLessonsResponse")
	proto.RegisterType((*ListLessonsRequest)(nil), "mentor.ListLessonsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MentorService service

type MentorServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ListLessons(ctx context.Context, in *ListLessonsRequest, opts ...grpc.CallOption) (MentorService_ListLessonsClient, error)
}

type mentorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMentorServiceClient(cc *grpc.ClientConn) MentorServiceClient {
	return &mentorServiceClient{cc}
}

func (c *mentorServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/mentor.MentorService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mentorServiceClient) ListLessons(ctx context.Context, in *ListLessonsRequest, opts ...grpc.CallOption) (MentorService_ListLessonsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MentorService_serviceDesc.Streams[0], c.cc, "/mentor.MentorService/ListLessons", opts...)
	if err != nil {
		return nil, err
	}
	x := &mentorServiceListLessonsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MentorService_ListLessonsClient interface {
	Recv() (*ListLessonsResponse, error)
	grpc.ClientStream
}

type mentorServiceListLessonsClient struct {
	grpc.ClientStream
}

func (x *mentorServiceListLessonsClient) Recv() (*ListLessonsResponse, error) {
	m := new(ListLessonsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MentorService service

type MentorServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	ListLessons(*ListLessonsRequest, MentorService_ListLessonsServer) error
}

func RegisterMentorServiceServer(s *grpc.Server, srv MentorServiceServer) {
	s.RegisterService(&_MentorService_serviceDesc, srv)
}

func _MentorService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MentorServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mentor.MentorService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MentorServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MentorService_ListLessons_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListLessonsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MentorServiceServer).ListLessons(m, &mentorServiceListLessonsServer{stream})
}

type MentorService_ListLessonsServer interface {
	Send(*ListLessonsResponse) error
	grpc.ServerStream
}

type mentorServiceListLessonsServer struct {
	grpc.ServerStream
}

func (x *mentorServiceListLessonsServer) Send(m *ListLessonsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MentorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mentor.MentorService",
	HandlerType: (*MentorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _MentorService_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListLessons",
			Handler:       _MentorService_ListLessons_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mentor.proto",
}

func init() { proto.RegisterFile("mentor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x54, 0xd2, 0x26, 0xea, 0xb7, 0xfd, 0xd1, 0xa7, 0xa5, 0x48, 0x51, 0xb8, 0x20, 0x1f, 0x10,
	0xa7, 0x0a, 0xb5, 0x5c, 0xb9, 0x22, 0x0e, 0x85, 0x83, 0xfb, 0x04, 0x6d, 0xba, 0x20, 0x4b, 0x69,
	0x5c, 0xbc, 0x0e, 0xa8, 0xbc, 0x04, 0xaf, 0x8c, 0xb2, 0x0e, 0x51, 0x0b, 0xdc, 0x3c, 0xb3, 0xb3,
	0x33, 0xab, 0x31, 0x8c, 0x76, 0x54, 0x79, 0xeb, 0x66, 0x7b, 0x67, 0xbd, 0xc5, 0x34, 0x20, 0x75,
	0x0f, 0xa3, 0xa5, 0x7d, 0x31, 0x95, 0xa6, 0xd7, 0x9a, 0xd8, 0x63, 0x0e, 0x83, 0x9a, 0xc9, 0x3d,
	0xad, 0x77, 0x94, 0x45, 0x97, 0xd1, 0xf5, 0x3f, 0xdd, 0xe1, 0x66, 0xb6, 0x5f, 0x33, 0xbf, 0x5b,
	0xb7, 0xcd, 0xe2, 0x30, 0xfb, 0xc6, 0x6a, 0x01, 0xe3, 0xd6, 0x87, 0xf7, 0xb6, 0x62, 0x42, 0x05,
	0x23, 0xc3, 0xab, 0xba, 0x28, 0x88, 0xf9, 0xb9, 0x2e, 0xc5, 0x6c, 0xa0, 0x4f, 0x38, 0xf5, 0x01,
	0xe9, 0x92, 0x98, 0x6d, 0x85, 0x13, 0x88, 0xcd, 0x56, 0x34, 0x89, 0x8e, 0xcd, 0x16, 0xa7, 0x90,
	0x70, 0x61, 0x1d, 0x49, 0x4e, 0xa2, 0x03, 0x68, 0x58, 0x6f, 0x7c, 0x49, 0x59, 0x4f, 0xd2, 0x03,
	0x68, 0x76, 0x37, 0x87, 0xac, 0x2f, 0x54, 0xbc, 0x39, 0x20, 0x42, 0xdf, 0x9b, 0x1d, 0x65, 0x89,
	0xac, 0xca, 0x1b, 0xff, 0x43, 0xaf, 0x76, 0x65, 0x96, 0x8a, 0xa8, 0x79, 0xaa, 0x3b, 0x38, 0x5b,
	0x1a, 0xf6, 0x21, 0x9f, 0xbb, 0xb3, 0xaf, 0x20, 0x2d, 0x85, 0x92, 0x63, 0x86, 0xf3, 0xc9, 0xac,
	0xad, 0x2d, 0x08, 0x75, 0x3b, 0x55, 0x53, 0xc0, 0x93, 0x75, 0x69, 0x6f, 0xfe, 0x19, 0xc1, 0xf8,
	0x51, 0xf4, 0x2b, 0x72, 0x6f, 0xa6, 0x20, 0xbc, 0x85, 0x44, 0x7a, 0xc1, 0x69, 0x67, 0x74, 0x54,
	0x77, 0x7e, 0xfe, 0x83, 0x6d, 0xaf, 0x78, 0x80, 0xe1, 0x91, 0x3b, 0xe6, 0x9d, 0xea, 0x57, 0x64,
	0x7e, 0xf1, 0xe7, 0x2c, 0xf8, 0xdc, 0x44, 0x9b, 0x54, 0xbe, 0x7b, 0xf1, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x9d, 0x9e, 0xa3, 0x12, 0xfe, 0x01, 0x00, 0x00,
}
