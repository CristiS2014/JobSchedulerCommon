// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/ingestion.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type SendJobRequest struct {
	Job                  *Job                 `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,15,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SendJobRequest) Reset()         { *m = SendJobRequest{} }
func (m *SendJobRequest) String() string { return proto.CompactTextString(m) }
func (*SendJobRequest) ProtoMessage()    {}
func (*SendJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd5582a5c5ad984e, []int{0}
}

func (m *SendJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendJobRequest.Unmarshal(m, b)
}
func (m *SendJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendJobRequest.Marshal(b, m, deterministic)
}
func (m *SendJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendJobRequest.Merge(m, src)
}
func (m *SendJobRequest) XXX_Size() int {
	return xxx_messageInfo_SendJobRequest.Size(m)
}
func (m *SendJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendJobRequest proto.InternalMessageInfo

func (m *SendJobRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *SendJobRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SendJobResponse struct {
	Id                   string               `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,15,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SendJobResponse) Reset()         { *m = SendJobResponse{} }
func (m *SendJobResponse) String() string { return proto.CompactTextString(m) }
func (*SendJobResponse) ProtoMessage()    {}
func (*SendJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd5582a5c5ad984e, []int{1}
}

func (m *SendJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendJobResponse.Unmarshal(m, b)
}
func (m *SendJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendJobResponse.Marshal(b, m, deterministic)
}
func (m *SendJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendJobResponse.Merge(m, src)
}
func (m *SendJobResponse) XXX_Size() int {
	return xxx_messageInfo_SendJobResponse.Size(m)
}
func (m *SendJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendJobResponse proto.InternalMessageInfo

func (m *SendJobResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SendJobResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*SendJobRequest)(nil), "SendJobRequest")
	proto.RegisterType((*SendJobResponse)(nil), "SendJobResponse")
}

func init() {
	proto.RegisterFile("protos/ingestion.proto", fileDescriptor_dd5582a5c5ad984e)
}

var fileDescriptor_dd5582a5c5ad984e = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0xcc, 0x4b, 0x4f, 0x2d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x0b, 0x48,
	0xc9, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x79, 0x49, 0xa5, 0x69, 0xfa, 0x25, 0x99,
	0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x50, 0x05, 0x82, 0x50, 0x8d, 0x59, 0xf9, 0x49, 0xc5,
	0x10, 0x21, 0xa5, 0x24, 0x2e, 0xbe, 0xe0, 0xd4, 0xbc, 0x14, 0xaf, 0xfc, 0xa4, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x31, 0x2e, 0xe6, 0xac, 0xfc, 0x24, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x16, 0x3d, 0x90, 0x0c, 0x48, 0x40, 0xc8, 0x82, 0x8b, 0x13, 0x6e, 0x9e, 0x04, 0x3f,
	0x58, 0x56, 0x4a, 0x0f, 0x62, 0xa3, 0x1e, 0xcc, 0x46, 0xbd, 0x10, 0x98, 0x8a, 0x20, 0x84, 0x62,
	0xa5, 0x68, 0x2e, 0x7e, 0xb8, 0x1d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c,
	0x9e, 0x29, 0x60, 0x3b, 0x38, 0x83, 0x98, 0x3c, 0x53, 0xc8, 0x37, 0xdc, 0xc8, 0x8e, 0x8b, 0x27,
	0x38, 0x39, 0x23, 0x35, 0xa5, 0x34, 0x27, 0xb5, 0xc8, 0x31, 0xc0, 0x53, 0x48, 0x8f, 0x8b, 0x1d,
	0x6a, 0x99, 0x10, 0xbf, 0x1e, 0xaa, 0xd7, 0xa4, 0x04, 0xf4, 0xd0, 0xdc, 0xa1, 0xc4, 0xe0, 0xc4,
	0x15, 0xc5, 0xa1, 0x07, 0x09, 0xb1, 0xe2, 0x24, 0x36, 0x30, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x11, 0x60, 0xb1, 0xc4, 0x61, 0x01, 0x00, 0x00,
}
