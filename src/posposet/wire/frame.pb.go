// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frame.proto

package wire

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventDescr struct {
	Creator              []byte   `protobuf:"bytes,1,opt,name=Creator,proto3" json:"Creator,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventDescr) Reset()         { *m = EventDescr{} }
func (m *EventDescr) String() string { return proto.CompactTextString(m) }
func (*EventDescr) ProtoMessage()    {}
func (*EventDescr) Descriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{0}
}

func (m *EventDescr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventDescr.Unmarshal(m, b)
}
func (m *EventDescr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventDescr.Marshal(b, m, deterministic)
}
func (m *EventDescr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDescr.Merge(m, src)
}
func (m *EventDescr) XXX_Size() int {
	return xxx_messageInfo_EventDescr.Size(m)
}
func (m *EventDescr) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDescr.DiscardUnknown(m)
}

var xxx_messageInfo_EventDescr proto.InternalMessageInfo

func (m *EventDescr) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *EventDescr) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Frame struct {
	Index                uint32        `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Events               []*EventDescr `protobuf:"bytes,2,rep,name=Events,proto3" json:"Events,omitempty"`
	Roots                []*EventDescr `protobuf:"bytes,3,rep,name=Roots,proto3" json:"Roots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{1}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Frame) GetEvents() []*EventDescr {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Frame) GetRoots() []*EventDescr {
	if m != nil {
		return m.Roots
	}
	return nil
}

func init() {
	proto.RegisterType((*EventDescr)(nil), "wire.EventDescr")
	proto.RegisterType((*Frame)(nil), "wire.Frame")
}

func init() { proto.RegisterFile("frame.proto", fileDescriptor_5379e2b825e15002) }

var fileDescriptor_5379e2b825e15002 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2b, 0x4a, 0xcc,
	0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xcf, 0x2c, 0x4a, 0x55, 0xb2, 0xe2,
	0xe2, 0x72, 0x2d, 0x4b, 0xcd, 0x2b, 0x71, 0x49, 0x2d, 0x4e, 0x2e, 0x12, 0x92, 0xe0, 0x62, 0x77,
	0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x71, 0x85,
	0x84, 0xb8, 0x58, 0x3c, 0x12, 0x8b, 0x33, 0x24, 0x98, 0xc0, 0xc2, 0x60, 0xb6, 0x52, 0x3e, 0x17,
	0xab, 0x1b, 0xc8, 0x40, 0x21, 0x11, 0x2e, 0x56, 0xcf, 0xbc, 0x94, 0xd4, 0x0a, 0xb0, 0x26, 0xde,
	0x20, 0x08, 0x47, 0x48, 0x83, 0x8b, 0x0d, 0x6c, 0x74, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7,
	0x91, 0x80, 0x1e, 0xc8, 0x46, 0x3d, 0x84, 0x75, 0x41, 0x50, 0x79, 0x21, 0x35, 0x2e, 0xd6, 0xa0,
	0xfc, 0xfc, 0x92, 0x62, 0x09, 0x66, 0x1c, 0x0a, 0x21, 0xd2, 0x49, 0x6c, 0x60, 0x97, 0x1b, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x0d, 0x69, 0x84, 0xc8, 0x00, 0x00, 0x00,
}
