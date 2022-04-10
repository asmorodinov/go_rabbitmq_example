// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interface.proto

package protobuf

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

type Request struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Lang                 string   `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	Titles               bool     `protobuf:"varint,4,opt,name=titles,proto3" json:"titles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Request) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Request) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Request) GetTitles() bool {
	if m != nil {
		return m.Titles
	}
	return false
}

type Result struct {
	Length               int32    `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Path                 []string `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Result) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Result) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ef53c9e620778f1, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "helloworld.Request")
	proto.RegisterType((*Result)(nil), "helloworld.Result")
	proto.RegisterType((*Empty)(nil), "helloworld.Empty")
}

func init() { proto.RegisterFile("interface.proto", fileDescriptor_3ef53c9e620778f1) }

var fileDescriptor_3ef53c9e620778f1 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0xb4, 0x49, 0x75, 0x28, 0x4a, 0x47, 0x28, 0xa1, 0xa7, 0xb0, 0xa7, 0x9e, 0x22,
	0xe8, 0xc1, 0x9b, 0x07, 0xff, 0x7d, 0x80, 0xbd, 0x88, 0xde, 0x52, 0x33, 0x69, 0x16, 0x37, 0xbb,
	0x71, 0x33, 0x41, 0xfc, 0x18, 0x7e, 0x63, 0xd9, 0x4d, 0xc0, 0x0a, 0xde, 0xde, 0xbc, 0x79, 0xfb,
	0x5b, 0xde, 0xc0, 0xb9, 0x32, 0x4c, 0xae, 0x2e, 0xdf, 0xa8, 0xe8, 0x9c, 0x65, 0x8b, 0xd0, 0x90,
	0xd6, 0xf6, 0xd3, 0x3a, 0x5d, 0x89, 0x17, 0x58, 0x4a, 0xfa, 0x18, 0xa8, 0x67, 0x44, 0x58, 0xd4,
	0xce, 0xb6, 0x59, 0x94, 0x47, 0xbb, 0x53, 0x19, 0x34, 0x9e, 0x41, 0xcc, 0x36, 0x8b, 0x83, 0x13,
	0xb3, 0xf5, 0x19, 0x5d, 0x9a, 0x43, 0x36, 0x1f, 0x33, 0x5e, 0xe3, 0x06, 0x52, 0x56, 0xac, 0xa9,
	0xcf, 0x16, 0x79, 0xb4, 0x3b, 0x91, 0xd3, 0x24, 0x1e, 0x20, 0x95, 0xd4, 0x0f, 0x9a, 0x7d, 0x42,
	0x93, 0x39, 0x70, 0x13, 0xd8, 0x89, 0x9c, 0x26, 0x4f, 0xeb, 0x4a, 0x6e, 0xb2, 0x38, 0x9f, 0x7b,
	0x9a, 0xd7, 0xfe, 0x47, 0x55, 0x05, 0x7e, 0x22, 0x63, 0x55, 0x89, 0x25, 0x24, 0x8f, 0x6d, 0xc7,
	0x5f, 0x57, 0xdf, 0x11, 0x6c, 0x24, 0xf1, 0xe0, 0xcc, 0x48, 0x7d, 0x72, 0xb6, 0x7d, 0xb6, 0xee,
	0x9d, 0x1c, 0xde, 0xc0, 0xea, 0x78, 0x83, 0x58, 0xfc, 0x36, 0x2c, 0x46, 0x6f, 0xbb, 0x3e, 0xf6,
	0x02, 0x51, 0xcc, 0xf0, 0x16, 0xd6, 0x53, 0x7b, 0x4f, 0xbb, 0xd7, 0x8a, 0x0c, 0xe3, 0xc5, 0xdf,
	0xd7, 0x61, 0xbd, 0xfd, 0x07, 0x29, 0x66, 0x77, 0xab, 0x57, 0x28, 0x2e, 0xc3, 0x51, 0xf7, 0x43,
	0xbd, 0x4f, 0x83, 0xba, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x86, 0x57, 0xe1, 0x0f, 0x71, 0x01,
	0x00, 0x00,
}
