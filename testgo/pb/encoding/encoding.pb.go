// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encoding.proto

package encoding

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

type MsgInt32 struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgInt32) Reset()         { *m = MsgInt32{} }
func (m *MsgInt32) String() string { return proto.CompactTextString(m) }
func (*MsgInt32) ProtoMessage()    {}
func (*MsgInt32) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac330e3fa468db3c, []int{0}
}

func (m *MsgInt32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgInt32.Unmarshal(m, b)
}
func (m *MsgInt32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgInt32.Marshal(b, m, deterministic)
}
func (m *MsgInt32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInt32.Merge(m, src)
}
func (m *MsgInt32) XXX_Size() int {
	return xxx_messageInfo_MsgInt32.Size(m)
}
func (m *MsgInt32) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInt32.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInt32 proto.InternalMessageInfo

func (m *MsgInt32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MsgInt64 struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgInt64) Reset()         { *m = MsgInt64{} }
func (m *MsgInt64) String() string { return proto.CompactTextString(m) }
func (*MsgInt64) ProtoMessage()    {}
func (*MsgInt64) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac330e3fa468db3c, []int{1}
}

func (m *MsgInt64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgInt64.Unmarshal(m, b)
}
func (m *MsgInt64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgInt64.Marshal(b, m, deterministic)
}
func (m *MsgInt64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInt64.Merge(m, src)
}
func (m *MsgInt64) XXX_Size() int {
	return xxx_messageInfo_MsgInt64.Size(m)
}
func (m *MsgInt64) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInt64.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInt64 proto.InternalMessageInfo

func (m *MsgInt64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MsgSint32 struct {
	Value                int32    `protobuf:"zigzag32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSint32) Reset()         { *m = MsgSint32{} }
func (m *MsgSint32) String() string { return proto.CompactTextString(m) }
func (*MsgSint32) ProtoMessage()    {}
func (*MsgSint32) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac330e3fa468db3c, []int{2}
}

func (m *MsgSint32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSint32.Unmarshal(m, b)
}
func (m *MsgSint32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSint32.Marshal(b, m, deterministic)
}
func (m *MsgSint32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSint32.Merge(m, src)
}
func (m *MsgSint32) XXX_Size() int {
	return xxx_messageInfo_MsgSint32.Size(m)
}
func (m *MsgSint32) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSint32.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSint32 proto.InternalMessageInfo

func (m *MsgSint32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MsgSint64 struct {
	Value                int64    `protobuf:"zigzag64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSint64) Reset()         { *m = MsgSint64{} }
func (m *MsgSint64) String() string { return proto.CompactTextString(m) }
func (*MsgSint64) ProtoMessage()    {}
func (*MsgSint64) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac330e3fa468db3c, []int{3}
}

func (m *MsgSint64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSint64.Unmarshal(m, b)
}
func (m *MsgSint64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSint64.Marshal(b, m, deterministic)
}
func (m *MsgSint64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSint64.Merge(m, src)
}
func (m *MsgSint64) XXX_Size() int {
	return xxx_messageInfo_MsgSint64.Size(m)
}
func (m *MsgSint64) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSint64.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSint64 proto.InternalMessageInfo

func (m *MsgSint64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MsgString struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgString) Reset()         { *m = MsgString{} }
func (m *MsgString) String() string { return proto.CompactTextString(m) }
func (*MsgString) ProtoMessage()    {}
func (*MsgString) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac330e3fa468db3c, []int{4}
}

func (m *MsgString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgString.Unmarshal(m, b)
}
func (m *MsgString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgString.Marshal(b, m, deterministic)
}
func (m *MsgString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgString.Merge(m, src)
}
func (m *MsgString) XXX_Size() int {
	return xxx_messageInfo_MsgString.Size(m)
}
func (m *MsgString) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgString.DiscardUnknown(m)
}

var xxx_messageInfo_MsgString proto.InternalMessageInfo

func (m *MsgString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgInt32)(nil), "encoding.MsgInt32")
	proto.RegisterType((*MsgInt64)(nil), "encoding.MsgInt64")
	proto.RegisterType((*MsgSint32)(nil), "encoding.MsgSint32")
	proto.RegisterType((*MsgSint64)(nil), "encoding.MsgSint64")
	proto.RegisterType((*MsgString)(nil), "encoding.MsgString")
}

func init() { proto.RegisterFile("encoding.proto", fileDescriptor_ac330e3fa468db3c) }

var fileDescriptor_ac330e3fa468db3c = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x4b, 0xce,
	0x4f, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x14,
	0xb8, 0x38, 0x7c, 0x8b, 0xd3, 0x3d, 0xf3, 0x4a, 0x8c, 0x8d, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12,
	0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0x84, 0x0a, 0x33, 0x13,
	0x54, 0x15, 0xcc, 0x30, 0x15, 0x8a, 0x5c, 0x9c, 0xbe, 0xc5, 0xe9, 0xc1, 0x99, 0x98, 0x86, 0x08,
	0x62, 0x2a, 0x41, 0x37, 0x45, 0x08, 0x4d, 0x49, 0x49, 0x51, 0x66, 0x5e, 0x3a, 0xaa, 0x12, 0x4e,
	0xa8, 0x92, 0x24, 0x36, 0xb0, 0xeb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0x7e, 0x8b,
	0xab, 0xcf, 0x00, 0x00, 0x00,
}