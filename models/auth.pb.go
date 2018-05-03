// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package models

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

type AuthResponse struct {
	Meta                 *DefaultResponse `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Token                string           `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_45e54ea74b933f01, []int{0}
}
func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (dst *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(dst, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetMeta() *DefaultResponse {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthResponse)(nil), "auth.AuthResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthService service

type AuthServiceClient interface {
	Signup(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Login(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Logout(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Signup(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/auth.AuthService/Signup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/auth.AuthService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := grpc.Invoke(ctx, "/auth.AuthService/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	Signup(context.Context, *UserRequest) (*AuthResponse, error)
	Login(context.Context, *UserRequest) (*AuthResponse, error)
	Logout(context.Context, *EmptyMessage) (*DefaultResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Signup(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _AuthService_Signup_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_45e54ea74b933f01) }

var fileDescriptor_auth_45e54ea74b933f01 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x2f, 0x72, 0x17, 0x74, 0x4e, 0x0b, 0x17, 0x85, 0x90, 0xea, 0x48, 0x75, 0x85, 0x04,
	0x3d, 0x4b, 0x2b, 0x45, 0xbb, 0xb3, 0xd9, 0xc3, 0xc6, 0x6e, 0xd5, 0xe7, 0x5e, 0x30, 0xd9, 0x8d,
	0x3b, 0xb3, 0x82, 0xbf, 0xca, 0xbf, 0x28, 0xb7, 0x21, 0x62, 0x23, 0xd8, 0xbd, 0x19, 0xde, 0x9b,
	0xf9, 0x66, 0x88, 0x4c, 0x94, 0x6d, 0xdd, 0x07, 0x2f, 0x5e, 0x4d, 0x77, 0xba, 0x3c, 0xb2, 0x70,
	0x08, 0xa6, 0x1d, 0x9a, 0x25, 0x45, 0x46, 0x18, 0x74, 0xa5, 0xe9, 0xf0, 0x3a, 0xca, 0x56, 0x83,
	0x7b, 0xef, 0x18, 0xea, 0x8c, 0xa6, 0x1d, 0xc4, 0x14, 0xd9, 0x22, 0x5b, 0xce, 0x57, 0x45, 0x3d,
	0x26, 0x6f, 0xf1, 0x6a, 0x62, 0x2b, 0xa3, 0x4f, 0x27, 0x97, 0x3a, 0xa1, 0x99, 0xf8, 0x37, 0xb8,
	0x62, 0x6f, 0x91, 0x2d, 0x0f, 0xf4, 0x50, 0xac, 0xbe, 0x32, 0x9a, 0xef, 0x86, 0x6e, 0x10, 0x3e,
	0x9a, 0x67, 0xa8, 0x0b, 0xca, 0x37, 0x8d, 0x75, 0xb1, 0x57, 0xc7, 0x75, 0x5a, 0xfd, 0xc0, 0x08,
	0x1a, 0xef, 0x11, 0x2c, 0xa5, 0xaa, 0x13, 0xee, 0x6f, 0x88, 0x6a, 0xa2, 0xce, 0x69, 0xb6, 0xf6,
	0xb6, 0x71, 0xff, 0x4f, 0x5c, 0x51, 0xbe, 0xf6, 0xd6, 0x47, 0x51, 0xa7, 0x3f, 0xd0, 0x77, 0x5d,
	0x2f, 0x9f, 0xf7, 0x60, 0x36, 0x16, 0xe5, 0x9f, 0xb7, 0x54, 0x93, 0x9b, 0xfd, 0xc7, 0xbc, 0xf3,
	0x2f, 0x68, 0xf9, 0x29, 0x4f, 0x6f, 0xb9, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x50, 0x4e,
	0xc8, 0x45, 0x01, 0x00, 0x00,
}
