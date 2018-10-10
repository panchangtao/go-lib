// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package pb

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

type TestProto struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64             `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	NextId               string            `protobuf:"bytes,4,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
	Filter               map[string]string `protobuf:"bytes,5,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TestProto) Reset()         { *m = TestProto{} }
func (m *TestProto) String() string { return proto.CompactTextString(m) }
func (*TestProto) ProtoMessage()    {}
func (*TestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestProto.Unmarshal(m, b)
}
func (m *TestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestProto.Marshal(b, m, deterministic)
}
func (m *TestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestProto.Merge(m, src)
}
func (m *TestProto) XXX_Size() int {
	return xxx_messageInfo_TestProto.Size(m)
}
func (m *TestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TestProto.DiscardUnknown(m)
}

var xxx_messageInfo_TestProto proto.InternalMessageInfo

func (m *TestProto) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestProto) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *TestProto) GetNextId() string {
	if m != nil {
		return m.NextId
	}
	return ""
}

func (m *TestProto) GetFilter() map[string]string {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*TestProto)(nil), "pb.test_proto")
	proto.RegisterMapType((map[string]string)(nil), "pb.test_proto.FilterEntry")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x3a, 0xc5, 0x08, 0x11, 0x8a,
	0x87, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65,
	0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x45, 0xc0, 0x6c, 0x21,
	0x01, 0x2e, 0xe6, 0xc4, 0xf4, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x10, 0x53, 0x48,
	0x9c, 0x8b, 0x3d, 0x2f, 0xb5, 0xa2, 0x24, 0x3e, 0x33, 0x45, 0x82, 0x05, 0xac, 0x90, 0x0d, 0xc4,
	0xf5, 0x4c, 0x11, 0x32, 0xe2, 0x62, 0x4b, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0x92, 0x60, 0x55, 0x60,
	0xd6, 0xe0, 0x36, 0x92, 0xd2, 0x2b, 0x48, 0xd2, 0x43, 0x58, 0xa7, 0xe7, 0x06, 0x96, 0x74, 0xcd,
	0x2b, 0x29, 0xaa, 0x0c, 0x82, 0xaa, 0x94, 0xb2, 0xe4, 0xe2, 0x46, 0x12, 0x06, 0xd9, 0x96, 0x9d,
	0x5a, 0x09, 0x75, 0x12, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x0a, 0x73, 0x14,
	0x84, 0x63, 0xc5, 0x64, 0xc1, 0x98, 0xc4, 0x06, 0x36, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x0c, 0xa3, 0xb6, 0x09, 0xe5, 0x00, 0x00, 0x00,
}
