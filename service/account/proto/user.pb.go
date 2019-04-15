// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package go_micro_service_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ReqSignup struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignup) Reset()         { *m = ReqSignup{} }
func (m *ReqSignup) String() string { return proto.CompactTextString(m) }
func (*ReqSignup) ProtoMessage()    {}
func (*ReqSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *ReqSignup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignup.Unmarshal(m, b)
}
func (m *ReqSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignup.Marshal(b, m, deterministic)
}
func (m *ReqSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignup.Merge(m, src)
}
func (m *ReqSignup) XXX_Size() int {
	return xxx_messageInfo_ReqSignup.Size(m)
}
func (m *ReqSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignup.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignup proto.InternalMessageInfo

func (m *ReqSignup) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignup) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignup struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignup) Reset()         { *m = RespSignup{} }
func (m *RespSignup) String() string { return proto.CompactTextString(m) }
func (*RespSignup) ProtoMessage()    {}
func (*RespSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RespSignup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignup.Unmarshal(m, b)
}
func (m *RespSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignup.Marshal(b, m, deterministic)
}
func (m *RespSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignup.Merge(m, src)
}
func (m *RespSignup) XXX_Size() int {
	return xxx_messageInfo_RespSignup.Size(m)
}
func (m *RespSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignup.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignup proto.InternalMessageInfo

func (m *RespSignup) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignup) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqSignin struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignin) Reset()         { *m = ReqSignin{} }
func (m *ReqSignin) String() string { return proto.CompactTextString(m) }
func (*ReqSignin) ProtoMessage()    {}
func (*ReqSignin) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *ReqSignin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignin.Unmarshal(m, b)
}
func (m *ReqSignin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignin.Marshal(b, m, deterministic)
}
func (m *ReqSignin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignin.Merge(m, src)
}
func (m *ReqSignin) XXX_Size() int {
	return xxx_messageInfo_ReqSignin.Size(m)
}
func (m *ReqSignin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignin proto.InternalMessageInfo

func (m *ReqSignin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignin struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignin) Reset()         { *m = RespSignin{} }
func (m *RespSignin) String() string { return proto.CompactTextString(m) }
func (*RespSignin) ProtoMessage()    {}
func (*RespSignin) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *RespSignin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignin.Unmarshal(m, b)
}
func (m *RespSignin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignin.Marshal(b, m, deterministic)
}
func (m *RespSignin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignin.Merge(m, src)
}
func (m *RespSignin) XXX_Size() int {
	return xxx_messageInfo_RespSignin.Size(m)
}
func (m *RespSignin) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignin.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignin proto.InternalMessageInfo

func (m *RespSignin) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RespSignin) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqUserInfo struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserInfo) Reset()         { *m = ReqUserInfo{} }
func (m *ReqUserInfo) String() string { return proto.CompactTextString(m) }
func (*ReqUserInfo) ProtoMessage()    {}
func (*ReqUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *ReqUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserInfo.Unmarshal(m, b)
}
func (m *ReqUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserInfo.Marshal(b, m, deterministic)
}
func (m *ReqUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserInfo.Merge(m, src)
}
func (m *ReqUserInfo) XXX_Size() int {
	return xxx_messageInfo_ReqUserInfo.Size(m)
}
func (m *ReqUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserInfo proto.InternalMessageInfo

func (m *ReqUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RespUserInfo struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	SignupAt             string   `protobuf:"bytes,6,opt,name=signupAt,proto3" json:"signupAt,omitempty"`
	LastActiveAt         string   `protobuf:"bytes,7,opt,name=lastActiveAt,proto3" json:"lastActiveAt,omitempty"`
	Status               int32    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUserInfo) Reset()         { *m = RespUserInfo{} }
func (m *RespUserInfo) String() string { return proto.CompactTextString(m) }
func (*RespUserInfo) ProtoMessage()    {}
func (*RespUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *RespUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUserInfo.Unmarshal(m, b)
}
func (m *RespUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUserInfo.Marshal(b, m, deterministic)
}
func (m *RespUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUserInfo.Merge(m, src)
}
func (m *RespUserInfo) XXX_Size() int {
	return xxx_messageInfo_RespUserInfo.Size(m)
}
func (m *RespUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RespUserInfo proto.InternalMessageInfo

func (m *RespUserInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RespUserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RespUserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RespUserInfo) GetSignupAt() string {
	if m != nil {
		return m.SignupAt
	}
	return ""
}

func (m *RespUserInfo) GetLastActiveAt() string {
	if m != nil {
		return m.LastActiveAt
	}
	return ""
}

func (m *RespUserInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReqUserFile struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserFile) Reset()         { *m = ReqUserFile{} }
func (m *ReqUserFile) String() string { return proto.CompactTextString(m) }
func (*ReqUserFile) ProtoMessage()    {}
func (*ReqUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *ReqUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserFile.Unmarshal(m, b)
}
func (m *ReqUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserFile.Marshal(b, m, deterministic)
}
func (m *ReqUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserFile.Merge(m, src)
}
func (m *ReqUserFile) XXX_Size() int {
	return xxx_messageInfo_ReqUserFile.Size(m)
}
func (m *ReqUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserFile proto.InternalMessageInfo

func (m *ReqUserFile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUserFile) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type RespUserFile struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FileData             []byte   `protobuf:"bytes,3,opt,name=fileData,proto3" json:"fileData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUserFile) Reset()         { *m = RespUserFile{} }
func (m *RespUserFile) String() string { return proto.CompactTextString(m) }
func (*RespUserFile) ProtoMessage()    {}
func (*RespUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *RespUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUserFile.Unmarshal(m, b)
}
func (m *RespUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUserFile.Marshal(b, m, deterministic)
}
func (m *RespUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUserFile.Merge(m, src)
}
func (m *RespUserFile) XXX_Size() int {
	return xxx_messageInfo_RespUserFile.Size(m)
}
func (m *RespUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_RespUserFile proto.InternalMessageInfo

func (m *RespUserFile) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserFile) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserFile) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqSignup)(nil), "go.micro.service.user.ReqSignup")
	proto.RegisterType((*RespSignup)(nil), "go.micro.service.user.RespSignup")
	proto.RegisterType((*ReqSignin)(nil), "go.micro.service.user.ReqSignin")
	proto.RegisterType((*RespSignin)(nil), "go.micro.service.user.RespSignin")
	proto.RegisterType((*ReqUserInfo)(nil), "go.micro.service.user.ReqUserInfo")
	proto.RegisterType((*RespUserInfo)(nil), "go.micro.service.user.RespUserInfo")
	proto.RegisterType((*ReqUserFile)(nil), "go.micro.service.user.ReqUserFile")
	proto.RegisterType((*RespUserFile)(nil), "go.micro.service.user.RespUserFile")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0xaf, 0x12, 0x31,
	0x10, 0xc6, 0x97, 0xf7, 0xd8, 0x05, 0x06, 0x4e, 0x0d, 0x9a, 0x86, 0x13, 0xd6, 0x8b, 0x5e, 0xf6,
	0xa0, 0x37, 0x2f, 0x86, 0x68, 0x4c, 0x3c, 0x69, 0x96, 0x60, 0xbc, 0xd6, 0x65, 0xc0, 0xc6, 0xdd,
	0x76, 0xd9, 0x16, 0xfc, 0x67, 0x8d, 0x7f, 0x8b, 0x69, 0xbb, 0x5d, 0xc0, 0x00, 0x1b, 0xde, 0x8d,
	0x6f, 0x66, 0xfa, 0xe3, 0xeb, 0xf4, 0xcb, 0x02, 0xec, 0x35, 0xd6, 0x69, 0x55, 0x2b, 0xa3, 0xc8,
	0xb3, 0xad, 0x4a, 0x4b, 0x91, 0xd7, 0x2a, 0xd5, 0x58, 0x1f, 0x44, 0x8e, 0xa9, 0x6d, 0xb2, 0x0f,
	0x30, 0xca, 0x70, 0xb7, 0x14, 0x5b, 0xb9, 0xaf, 0xc8, 0x0c, 0x86, 0xb6, 0x28, 0x79, 0x89, 0xb4,
	0x37, 0xef, 0xbd, 0x1a, 0x65, 0xad, 0xb6, 0xbd, 0x8a, 0x6b, 0xfd, 0x5b, 0xd5, 0x6b, 0xfa, 0xe0,
	0x7b, 0x41, 0xb3, 0x77, 0x00, 0x19, 0xea, 0xaa, 0xa1, 0x10, 0xe8, 0xe7, 0x6a, 0xed, 0x09, 0x71,
	0xe6, 0x7e, 0x13, 0x0a, 0x83, 0x12, 0xb5, 0xe6, 0x5b, 0x6c, 0x0e, 0x07, 0x79, 0x62, 0x40, 0xc8,
	0x27, 0x1b, 0xf8, 0x7a, 0x34, 0x20, 0xe4, 0x45, 0x03, 0x53, 0x88, 0x8d, 0xfa, 0x85, 0xb2, 0x39,
	0xea, 0xc5, 0xa9, 0xad, 0xc7, 0x73, 0x5b, 0xaf, 0x61, 0x9c, 0xe1, 0x6e, 0xa5, 0xb1, 0xfe, 0x2c,
	0x37, 0xea, 0x96, 0x31, 0xf6, 0xa7, 0x07, 0x13, 0xfb, 0xef, 0xed, 0xf0, 0x5d, 0x0b, 0x38, 0x43,
	0x3f, 0xfe, 0x77, 0xe7, 0x29, 0xc4, 0x58, 0x72, 0x51, 0xd0, 0xbe, 0x77, 0xed, 0x84, 0xad, 0x56,
	0x3f, 0x95, 0x44, 0x1a, 0xfb, 0xaa, 0x13, 0x96, 0xa3, 0xdd, 0x03, 0x2c, 0x0c, 0x4d, 0x3c, 0x27,
	0x68, 0xc2, 0x60, 0x52, 0x70, 0x6d, 0x16, 0xb9, 0x11, 0x07, 0x5c, 0x18, 0x3a, 0x70, 0xfd, 0xb3,
	0x1a, 0x79, 0x0e, 0x89, 0x36, 0xdc, 0xec, 0x35, 0x1d, 0x3a, 0xdf, 0x8d, 0x62, 0xef, 0xdb, 0x4d,
	0x7c, 0x12, 0x05, 0xde, 0x7c, 0xa2, 0x29, 0xc4, 0x85, 0x28, 0x85, 0x71, 0x57, 0x8c, 0x33, 0x2f,
	0xd8, 0xf7, 0xe3, 0x7a, 0x1c, 0xe1, 0xee, 0xf5, 0x6c, 0x44, 0x81, 0x1f, 0xb9, 0xe1, 0x6e, 0x3d,
	0x93, 0xac, 0xd5, 0x6f, 0xfe, 0x3e, 0xc0, 0xd8, 0x62, 0x97, 0x3e, 0xd1, 0xe4, 0x0b, 0x24, 0x4d,
	0x06, 0xe7, 0xe9, 0xc5, 0xb8, 0xa7, 0x6d, 0xd6, 0x67, 0x2f, 0xae, 0x4e, 0x84, 0x20, 0xb3, 0x28,
	0x00, 0x85, 0xec, 0x02, 0x0a, 0xd9, 0x09, 0x14, 0x92, 0x45, 0x64, 0x05, 0xc3, 0x36, 0x26, 0xec,
	0x3a, 0x32, 0xcc, 0xcc, 0x5e, 0xde, 0x80, 0x86, 0x21, 0x16, 0x91, 0x6f, 0x30, 0x0a, 0xeb, 0xd5,
	0x5d, 0x5c, 0x3b, 0xd4, 0xc9, 0xb5, 0x43, 0x2c, 0xfa, 0x91, 0xb8, 0x6f, 0xc7, 0xdb, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd0, 0x48, 0x87, 0x32, 0x49, 0x04, 0x00, 0x00,
}
