// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type UserRequest struct {
	Auth                 *Authentication `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	User                 *User           `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_984ba6f534287b59, []int{0}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (dst *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(dst, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetAuth() *Authentication {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserIDRequest struct {
	Auth                 *Authentication `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	UserID               uint32          `protobuf:"varint,2,opt,name=userID" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserIDRequest) Reset()         { *m = UserIDRequest{} }
func (m *UserIDRequest) String() string { return proto.CompactTextString(m) }
func (*UserIDRequest) ProtoMessage()    {}
func (*UserIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_984ba6f534287b59, []int{1}
}
func (m *UserIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIDRequest.Unmarshal(m, b)
}
func (m *UserIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIDRequest.Marshal(b, m, deterministic)
}
func (dst *UserIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIDRequest.Merge(dst, src)
}
func (m *UserIDRequest) XXX_Size() int {
	return xxx_messageInfo_UserIDRequest.Size(m)
}
func (m *UserIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIDRequest proto.InternalMessageInfo

func (m *UserIDRequest) GetAuth() *Authentication {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UserIDRequest) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type UserEmailRequest struct {
	Auth                 *Authentication `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	UserEmail            string          `protobuf:"bytes,2,opt,name=userEmail" json:"userEmail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserEmailRequest) Reset()         { *m = UserEmailRequest{} }
func (m *UserEmailRequest) String() string { return proto.CompactTextString(m) }
func (*UserEmailRequest) ProtoMessage()    {}
func (*UserEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_984ba6f534287b59, []int{2}
}
func (m *UserEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEmailRequest.Unmarshal(m, b)
}
func (m *UserEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEmailRequest.Marshal(b, m, deterministic)
}
func (dst *UserEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEmailRequest.Merge(dst, src)
}
func (m *UserEmailRequest) XXX_Size() int {
	return xxx_messageInfo_UserEmailRequest.Size(m)
}
func (m *UserEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserEmailRequest proto.InternalMessageInfo

func (m *UserEmailRequest) GetAuth() *Authentication {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UserEmailRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

type UserUpdateRequest struct {
	Auth                 *Authentication `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	UserUpdate           *UserUpdate     `protobuf:"bytes,2,opt,name=userUpdate" json:"userUpdate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserUpdateRequest) Reset()         { *m = UserUpdateRequest{} }
func (m *UserUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UserUpdateRequest) ProtoMessage()    {}
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_984ba6f534287b59, []int{3}
}
func (m *UserUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateRequest.Unmarshal(m, b)
}
func (m *UserUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UserUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateRequest.Merge(dst, src)
}
func (m *UserUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UserUpdateRequest.Size(m)
}
func (m *UserUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateRequest proto.InternalMessageInfo

func (m *UserUpdateRequest) GetAuth() *Authentication {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UserUpdateRequest) GetUserUpdate() *UserUpdate {
	if m != nil {
		return m.UserUpdate
	}
	return nil
}

type User struct {
	// @inject_tag: storm:"id,increment"
	ID        uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty" storm:"id,increment"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName" json:"lastName,omitempty"`
	// @inject_tag: storm:"unique"
	Email                string               `protobuf:"bytes,4,opt,name=email" json:"email,omitempty" storm:"unique"`
	Password             string               `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	IsAdmin              bool                 `protobuf:"varint,6,opt,name=isAdmin" json:"isAdmin,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updatedAt" json:"updatedAt,omitempty"`
	LastSessionAt        *timestamp.Timestamp `protobuf:"bytes,9,opt,name=lastSessionAt" json:"lastSessionAt,omitempty"`
	HasAvatar            bool                 `protobuf:"varint,10,opt,name=hasAvatar" json:"hasAvatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_984ba6f534287b59, []int{4}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetLastSessionAt() *timestamp.Timestamp {
	if m != nil {
		return m.LastSessionAt
	}
	return nil
}

func (m *User) GetHasAvatar() bool {
	if m != nil {
		return m.HasAvatar
	}
	return false
}

type UserUpdate struct {
	// Types that are valid to be assigned to FirstNameOO:
	//	*UserUpdate_FirstName
	FirstNameOO isUserUpdate_FirstNameOO `protobuf_oneof:"firstNameOO"`
	// Types that are valid to be assigned to LastNameOO:
	//	*UserUpdate_LastName
	LastNameOO isUserUpdate_LastNameOO `protobuf_oneof:"lastNameOO"`
	// Types that are valid to be assigned to EmailOO:
	//	*UserUpdate_Email
	EmailOO isUserUpdate_EmailOO `protobuf_oneof:"emailOO"`
	// Types that are valid to be assigned to PasswordOO:
	//	*UserUpdate_Password
	PasswordOO isUserUpdate_PasswordOO `protobuf_oneof:"passwordOO"`
	// Types that are valid to be assigned to IsAdminOO:
	//	*UserUpdate_IsAdmin
	IsAdminOO isUserUpdate_IsAdminOO `protobuf_oneof:"isAdminOO"`
	// Types that are valid to be assigned to HasAvatarOO:
	//	*UserUpdate_HasAvatar
	HasAvatarOO          isUserUpdate_HasAvatarOO `protobuf_oneof:"hasAvatarOO"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UserUpdate) Reset()         { *m = UserUpdate{} }
func (m *UserUpdate) String() string { return proto.CompactTextString(m) }
func (*UserUpdate) ProtoMessage()    {}
func (*UserUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_984ba6f534287b59, []int{5}
}
func (m *UserUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdate.Unmarshal(m, b)
}
func (m *UserUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdate.Marshal(b, m, deterministic)
}
func (dst *UserUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdate.Merge(dst, src)
}
func (m *UserUpdate) XXX_Size() int {
	return xxx_messageInfo_UserUpdate.Size(m)
}
func (m *UserUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdate proto.InternalMessageInfo

type isUserUpdate_FirstNameOO interface {
	isUserUpdate_FirstNameOO()
}
type isUserUpdate_LastNameOO interface {
	isUserUpdate_LastNameOO()
}
type isUserUpdate_EmailOO interface {
	isUserUpdate_EmailOO()
}
type isUserUpdate_PasswordOO interface {
	isUserUpdate_PasswordOO()
}
type isUserUpdate_IsAdminOO interface {
	isUserUpdate_IsAdminOO()
}
type isUserUpdate_HasAvatarOO interface {
	isUserUpdate_HasAvatarOO()
}

type UserUpdate_FirstName struct {
	FirstName string `protobuf:"bytes,2,opt,name=firstName,oneof"`
}
type UserUpdate_LastName struct {
	LastName string `protobuf:"bytes,3,opt,name=lastName,oneof"`
}
type UserUpdate_Email struct {
	Email string `protobuf:"bytes,4,opt,name=email,oneof"`
}
type UserUpdate_Password struct {
	Password string `protobuf:"bytes,5,opt,name=password,oneof"`
}
type UserUpdate_IsAdmin struct {
	IsAdmin bool `protobuf:"varint,6,opt,name=isAdmin,oneof"`
}
type UserUpdate_HasAvatar struct {
	HasAvatar bool `protobuf:"varint,7,opt,name=hasAvatar,oneof"`
}

func (*UserUpdate_FirstName) isUserUpdate_FirstNameOO() {}
func (*UserUpdate_LastName) isUserUpdate_LastNameOO()   {}
func (*UserUpdate_Email) isUserUpdate_EmailOO()         {}
func (*UserUpdate_Password) isUserUpdate_PasswordOO()   {}
func (*UserUpdate_IsAdmin) isUserUpdate_IsAdminOO()     {}
func (*UserUpdate_HasAvatar) isUserUpdate_HasAvatarOO() {}

func (m *UserUpdate) GetFirstNameOO() isUserUpdate_FirstNameOO {
	if m != nil {
		return m.FirstNameOO
	}
	return nil
}
func (m *UserUpdate) GetLastNameOO() isUserUpdate_LastNameOO {
	if m != nil {
		return m.LastNameOO
	}
	return nil
}
func (m *UserUpdate) GetEmailOO() isUserUpdate_EmailOO {
	if m != nil {
		return m.EmailOO
	}
	return nil
}
func (m *UserUpdate) GetPasswordOO() isUserUpdate_PasswordOO {
	if m != nil {
		return m.PasswordOO
	}
	return nil
}
func (m *UserUpdate) GetIsAdminOO() isUserUpdate_IsAdminOO {
	if m != nil {
		return m.IsAdminOO
	}
	return nil
}
func (m *UserUpdate) GetHasAvatarOO() isUserUpdate_HasAvatarOO {
	if m != nil {
		return m.HasAvatarOO
	}
	return nil
}

func (m *UserUpdate) GetFirstName() string {
	if x, ok := m.GetFirstNameOO().(*UserUpdate_FirstName); ok {
		return x.FirstName
	}
	return ""
}

func (m *UserUpdate) GetLastName() string {
	if x, ok := m.GetLastNameOO().(*UserUpdate_LastName); ok {
		return x.LastName
	}
	return ""
}

func (m *UserUpdate) GetEmail() string {
	if x, ok := m.GetEmailOO().(*UserUpdate_Email); ok {
		return x.Email
	}
	return ""
}

func (m *UserUpdate) GetPassword() string {
	if x, ok := m.GetPasswordOO().(*UserUpdate_Password); ok {
		return x.Password
	}
	return ""
}

func (m *UserUpdate) GetIsAdmin() bool {
	if x, ok := m.GetIsAdminOO().(*UserUpdate_IsAdmin); ok {
		return x.IsAdmin
	}
	return false
}

func (m *UserUpdate) GetHasAvatar() bool {
	if x, ok := m.GetHasAvatarOO().(*UserUpdate_HasAvatar); ok {
		return x.HasAvatar
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UserUpdate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UserUpdate_OneofMarshaler, _UserUpdate_OneofUnmarshaler, _UserUpdate_OneofSizer, []interface{}{
		(*UserUpdate_FirstName)(nil),
		(*UserUpdate_LastName)(nil),
		(*UserUpdate_Email)(nil),
		(*UserUpdate_Password)(nil),
		(*UserUpdate_IsAdmin)(nil),
		(*UserUpdate_HasAvatar)(nil),
	}
}

func _UserUpdate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UserUpdate)
	// firstNameOO
	switch x := m.FirstNameOO.(type) {
	case *UserUpdate_FirstName:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FirstName)
	case nil:
	default:
		return fmt.Errorf("UserUpdate.FirstNameOO has unexpected type %T", x)
	}
	// lastNameOO
	switch x := m.LastNameOO.(type) {
	case *UserUpdate_LastName:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.LastName)
	case nil:
	default:
		return fmt.Errorf("UserUpdate.LastNameOO has unexpected type %T", x)
	}
	// emailOO
	switch x := m.EmailOO.(type) {
	case *UserUpdate_Email:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Email)
	case nil:
	default:
		return fmt.Errorf("UserUpdate.EmailOO has unexpected type %T", x)
	}
	// passwordOO
	switch x := m.PasswordOO.(type) {
	case *UserUpdate_Password:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Password)
	case nil:
	default:
		return fmt.Errorf("UserUpdate.PasswordOO has unexpected type %T", x)
	}
	// isAdminOO
	switch x := m.IsAdminOO.(type) {
	case *UserUpdate_IsAdmin:
		t := uint64(0)
		if x.IsAdmin {
			t = 1
		}
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("UserUpdate.IsAdminOO has unexpected type %T", x)
	}
	// hasAvatarOO
	switch x := m.HasAvatarOO.(type) {
	case *UserUpdate_HasAvatar:
		t := uint64(0)
		if x.HasAvatar {
			t = 1
		}
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("UserUpdate.HasAvatarOO has unexpected type %T", x)
	}
	return nil
}

func _UserUpdate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UserUpdate)
	switch tag {
	case 2: // firstNameOO.firstName
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.FirstNameOO = &UserUpdate_FirstName{x}
		return true, err
	case 3: // lastNameOO.lastName
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.LastNameOO = &UserUpdate_LastName{x}
		return true, err
	case 4: // emailOO.email
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.EmailOO = &UserUpdate_Email{x}
		return true, err
	case 5: // passwordOO.password
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PasswordOO = &UserUpdate_Password{x}
		return true, err
	case 6: // isAdminOO.isAdmin
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.IsAdminOO = &UserUpdate_IsAdmin{x != 0}
		return true, err
	case 7: // hasAvatarOO.hasAvatar
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.HasAvatarOO = &UserUpdate_HasAvatar{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _UserUpdate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UserUpdate)
	// firstNameOO
	switch x := m.FirstNameOO.(type) {
	case *UserUpdate_FirstName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.FirstName)))
		n += len(x.FirstName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// lastNameOO
	switch x := m.LastNameOO.(type) {
	case *UserUpdate_LastName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.LastName)))
		n += len(x.LastName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// emailOO
	switch x := m.EmailOO.(type) {
	case *UserUpdate_Email:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Email)))
		n += len(x.Email)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// passwordOO
	switch x := m.PasswordOO.(type) {
	case *UserUpdate_Password:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Password)))
		n += len(x.Password)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// isAdminOO
	switch x := m.IsAdminOO.(type) {
	case *UserUpdate_IsAdmin:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// hasAvatarOO
	switch x := m.HasAvatarOO.(type) {
	case *UserUpdate_HasAvatar:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "user.UserRequest")
	proto.RegisterType((*UserIDRequest)(nil), "user.UserIDRequest")
	proto.RegisterType((*UserEmailRequest)(nil), "user.UserEmailRequest")
	proto.RegisterType((*UserUpdateRequest)(nil), "user.UserUpdateRequest")
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserUpdate)(nil), "user.UserUpdate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetOwnUser(ctx context.Context, in *Authentication, opts ...grpc.CallOption) (*User, error)
	GetUserByID(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*User, error)
	GetUserByEmail(ctx context.Context, in *UserEmailRequest, opts ...grpc.CallOption) (*User, error)
	UpdateOwnUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*User, error)
	DeleteOwnUser(ctx context.Context, in *Authentication, opts ...grpc.CallOption) (*EmptyMessage, error)
	DeleteUserByID(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetOwnUser(ctx context.Context, in *Authentication, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetOwnUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByID(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByEmail(ctx context.Context, in *UserEmailRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateOwnUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateOwnUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteOwnUser(ctx context.Context, in *Authentication, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteOwnUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserByID(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	GetOwnUser(context.Context, *Authentication) (*User, error)
	GetUserByID(context.Context, *UserIDRequest) (*User, error)
	GetUserByEmail(context.Context, *UserEmailRequest) (*User, error)
	UpdateOwnUser(context.Context, *UserUpdateRequest) (*User, error)
	DeleteOwnUser(context.Context, *Authentication) (*EmptyMessage, error)
	DeleteUserByID(context.Context, *UserIDRequest) (*EmptyMessage, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetOwnUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authentication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOwnUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetOwnUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOwnUser(ctx, req.(*Authentication))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByID(ctx, req.(*UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByEmail(ctx, req.(*UserEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateOwnUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateOwnUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateOwnUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateOwnUser(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteOwnUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authentication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteOwnUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteOwnUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteOwnUser(ctx, req.(*Authentication))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserByID(ctx, req.(*UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOwnUser",
			Handler:    _UserService_GetOwnUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _UserService_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserService_GetUserByEmail_Handler,
		},
		{
			MethodName: "UpdateOwnUser",
			Handler:    _UserService_UpdateOwnUser_Handler,
		},
		{
			MethodName: "DeleteOwnUser",
			Handler:    _UserService_DeleteOwnUser_Handler,
		},
		{
			MethodName: "DeleteUserByID",
			Handler:    _UserService_DeleteUserByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_984ba6f534287b59) }

var fileDescriptor_user_984ba6f534287b59 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x15, 0x7f, 0x69, 0x5c, 0x99, 0x74, 0xdb, 0xba, 0x42, 0x18, 0x37, 0xe8, 0x14, 0x28,
	0x28, 0x21, 0x81, 0xd2, 0x53, 0xa9, 0x84, 0x43, 0xec, 0x43, 0x2b, 0x50, 0x92, 0x4b, 0xa0, 0x87,
	0x8d, 0x3d, 0xb1, 0x05, 0xfa, 0x70, 0xb5, 0xab, 0x84, 0xfc, 0x89, 0xfe, 0xcd, 0x5e, 0xfb, 0x13,
	0xca, 0xee, 0x5a, 0x1f, 0x71, 0xea, 0x06, 0x7c, 0xd3, 0xec, 0x9b, 0xf7, 0x66, 0xe6, 0xed, 0xac,
	0x00, 0x72, 0x86, 0x99, 0xb3, 0xca, 0x52, 0x9e, 0x92, 0xa6, 0xf8, 0xb6, 0x8c, 0x05, 0x26, 0x98,
	0xd1, 0x48, 0x1d, 0x5a, 0x1f, 0x16, 0x69, 0xba, 0x88, 0xf0, 0x58, 0x46, 0xb7, 0xf9, 0xdd, 0x31,
	0x0f, 0x63, 0x64, 0x9c, 0xc6, 0x2b, 0x95, 0x60, 0xdf, 0x40, 0xef, 0x9a, 0x61, 0x16, 0xe0, 0xcf,
	0x1c, 0x19, 0x27, 0x1f, 0xa1, 0x49, 0x73, 0xbe, 0x34, 0x1b, 0x87, 0x8d, 0xa3, 0xde, 0xe9, 0x7b,
	0xa7, 0x50, 0x73, 0x73, 0xbe, 0xc4, 0x84, 0x87, 0x33, 0xca, 0xc3, 0x34, 0x09, 0x64, 0x12, 0x19,
	0x81, 0xac, 0x69, 0x6a, 0x32, 0x19, 0x1c, 0xd9, 0x8c, 0x54, 0x93, 0xe7, 0xf6, 0x15, 0x18, 0x22,
	0x9a, 0x8e, 0x77, 0x52, 0x1f, 0x40, 0x3b, 0x97, 0x6c, 0xa9, 0x6f, 0x04, 0xeb, 0xc8, 0xfe, 0x01,
	0x07, 0x42, 0xf5, 0x3c, 0xa6, 0x61, 0xb4, 0x93, 0xf0, 0x10, 0xf4, 0xbc, 0x10, 0x90, 0xda, 0x7a,
	0x50, 0x1d, 0xd8, 0x19, 0xbc, 0x16, 0xf2, 0xd7, 0xab, 0x39, 0xe5, 0xb8, 0x93, 0xfe, 0x89, 0xba,
	0x16, 0xa5, 0xb0, 0x36, 0xe7, 0xa0, 0x32, 0x67, 0xad, 0x5c, 0xcb, 0xb1, 0xff, 0x68, 0xd0, 0x14,
	0x10, 0xe9, 0x83, 0x36, 0x1d, 0xcb, 0x2a, 0x46, 0xa0, 0x4d, 0xc7, 0xa2, 0xd5, 0xbb, 0x30, 0x63,
	0xfc, 0x3b, 0x8d, 0xb1, 0x68, 0xb5, 0x3c, 0x20, 0x16, 0x74, 0x23, 0xba, 0x06, 0xf7, 0x25, 0x58,
	0xc6, 0xe4, 0x2d, 0xb4, 0x50, 0x0e, 0xd8, 0x94, 0x80, 0x0a, 0x04, 0x63, 0x45, 0x19, 0x7b, 0x48,
	0xb3, 0xb9, 0xd9, 0x52, 0x8c, 0x22, 0x26, 0x26, 0x74, 0x42, 0xe6, 0xce, 0xe3, 0x30, 0x31, 0xdb,
	0x87, 0x8d, 0xa3, 0x6e, 0x50, 0x84, 0xe4, 0x33, 0xe8, 0xb3, 0x0c, 0x29, 0xc7, 0xb9, 0xcb, 0xcd,
	0x8e, 0x9c, 0xc7, 0x72, 0xd4, 0x62, 0x39, 0xc5, 0x62, 0x39, 0x57, 0xc5, 0x62, 0x05, 0x55, 0xb2,
	0x60, 0xe6, 0x72, 0x44, 0xc1, 0xec, 0xbe, 0xcc, 0x2c, 0x93, 0xc9, 0x57, 0x30, 0xc4, 0x2c, 0x97,
	0xc8, 0x58, 0x98, 0x26, 0x2e, 0x37, 0xf5, 0x17, 0xd9, 0x4f, 0x09, 0xc2, 0xbb, 0x25, 0x65, 0xee,
	0x3d, 0xe5, 0x34, 0x33, 0x41, 0x4e, 0x54, 0x1d, 0xd8, 0xbf, 0x34, 0x80, 0xea, 0x36, 0xc8, 0xe8,
	0x99, 0xd1, 0x93, 0xbd, 0xba, 0xd5, 0xc3, 0x4d, 0xab, 0x27, 0x8d, 0x9a, 0xd9, 0x83, 0x27, 0x66,
	0x4f, 0xb4, 0xc2, 0xee, 0xe1, 0xa6, 0xdd, 0x93, 0xfd, 0x9a, 0xe1, 0xd6, 0x86, 0xe1, 0x93, 0x66,
	0x65, 0xf9, 0xa8, 0xde, 0x7c, 0x47, 0xa2, 0xad, 0x5a, 0xfb, 0x9e, 0x01, 0xbd, 0xb2, 0x39, 0xdf,
	0xf7, 0x5e, 0x01, 0x14, 0xcd, 0xf8, 0xbe, 0xa7, 0x43, 0x47, 0xd6, 0x57, 0x40, 0x51, 0xcf, 0xf7,
	0xbd, 0x1e, 0xe8, 0xeb, 0x02, 0xbe, 0x2f, 0x24, 0x4a, 0x3d, 0xdf, 0x3f, 0xfd, 0xad, 0xa9, 0x3f,
	0xc1, 0x25, 0x66, 0xf7, 0xe1, 0x0c, 0xc9, 0x19, 0xc0, 0x05, 0x72, 0xff, 0x21, 0x91, 0x8b, 0xb9,
	0x6d, 0xe5, 0xad, 0xda, 0xab, 0xb7, 0xf7, 0xc8, 0x09, 0xf4, 0x2e, 0x90, 0x8b, 0xc0, 0x7b, 0x9c,
	0x8e, 0xc9, 0x9b, 0x0a, 0x2c, 0x7f, 0x02, 0x1b, 0x8c, 0x4f, 0xd0, 0x2f, 0x19, 0xf2, 0x01, 0x92,
	0x41, 0x85, 0xd7, 0xdf, 0xf8, 0x33, 0x9e, 0xa1, 0xae, 0xae, 0xea, 0x70, 0xf3, 0x85, 0xfd, 0x93,
	0xe7, 0x82, 0x31, 0xc6, 0x08, 0xeb, 0xbc, 0x2d, 0x93, 0xbd, 0x2b, 0x81, 0xf3, 0x78, 0xc5, 0x1f,
	0xbf, 0x21, 0x63, 0x74, 0x81, 0xf6, 0x1e, 0xf9, 0x02, 0x7d, 0x25, 0xf1, 0xff, 0x39, 0xb7, 0xf1,
	0xbd, 0xee, 0x4d, 0x3b, 0x4e, 0xe7, 0x18, 0xb1, 0xdb, 0xb6, 0xdc, 0xe2, 0xb3, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x30, 0xe3, 0x5a, 0x8f, 0xc7, 0x05, 0x00, 0x00,
}
