// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/jobs.proto

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

type Job struct {
	Type                 string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Parameters           []byte               `protobuf:"bytes,2,opt,name=parameters,proto3" json:"parameters,omitempty"`
	Id                   string               `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,15,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c98b8c82adc31b, []int{0}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Job) GetParameters() []byte {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type JobUpdate struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string               `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Result               string               `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	MetaData             string               `protobuf:"bytes,4,opt,name=metaData,proto3" json:"metaData,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,15,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *JobUpdate) Reset()         { *m = JobUpdate{} }
func (m *JobUpdate) String() string { return proto.CompactTextString(m) }
func (*JobUpdate) ProtoMessage()    {}
func (*JobUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c98b8c82adc31b, []int{1}
}

func (m *JobUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobUpdate.Unmarshal(m, b)
}
func (m *JobUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobUpdate.Marshal(b, m, deterministic)
}
func (m *JobUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobUpdate.Merge(m, src)
}
func (m *JobUpdate) XXX_Size() int {
	return xxx_messageInfo_JobUpdate.Size(m)
}
func (m *JobUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_JobUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_JobUpdate proto.InternalMessageInfo

func (m *JobUpdate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JobUpdate) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *JobUpdate) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *JobUpdate) GetMetaData() string {
	if m != nil {
		return m.MetaData
	}
	return ""
}

func (m *JobUpdate) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*JobUpdate)(nil), "JobUpdate")
}

func init() {
	proto.RegisterFile("protos/jobs.proto", fileDescriptor_54c98b8c82adc31b)
}

var fileDescriptor_54c98b8c82adc31b = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x3f, 0x4b, 0xc5, 0x30,
	0x14, 0xc5, 0x49, 0xde, 0xe3, 0xf1, 0x7a, 0x15, 0xc5, 0x0c, 0x12, 0x3a, 0x68, 0xe9, 0xd4, 0x29,
	0x05, 0x5d, 0x9c, 0xc5, 0xa9, 0x63, 0xd1, 0xc5, 0x2d, 0xa1, 0xd7, 0x52, 0x69, 0x48, 0x48, 0x6e,
	0x07, 0x67, 0x3f, 0x87, 0xdf, 0x55, 0x48, 0xff, 0xd8, 0xd9, 0x29, 0xe7, 0x24, 0x27, 0xe7, 0xfe,
	0xb8, 0x70, 0xe3, 0x83, 0x23, 0x17, 0xeb, 0x4f, 0x67, 0xa2, 0x4a, 0x3a, 0xbf, 0xef, 0x9d, 0xeb,
	0x47, 0xac, 0x93, 0x33, 0xd3, 0x47, 0x4d, 0x83, 0xc5, 0x48, 0xda, 0xfa, 0x39, 0x50, 0x7e, 0x33,
	0x38, 0x34, 0xce, 0x08, 0x01, 0x47, 0xfa, 0xf2, 0x28, 0x59, 0xc1, 0xaa, 0xac, 0x4d, 0x5a, 0xdc,
	0x01, 0x78, 0x1d, 0xb4, 0x45, 0xc2, 0x10, 0x25, 0x2f, 0x58, 0x75, 0xd9, 0xee, 0x6e, 0xc4, 0x15,
	0xf0, 0xa1, 0x93, 0x87, 0xf4, 0x83, 0x0f, 0x9d, 0x78, 0x82, 0x6c, 0xab, 0x97, 0xd7, 0x05, 0xab,
	0x2e, 0x1e, 0x72, 0x35, 0x03, 0xa8, 0x15, 0x40, 0xbd, 0xae, 0x89, 0xf6, 0x2f, 0x5c, 0xfe, 0x30,
	0xc8, 0x1a, 0x67, 0xde, 0x7c, 0xa7, 0x09, 0x97, 0x5e, 0xb6, 0xf5, 0xae, 0x6c, 0x7c, 0xc7, 0x76,
	0x0b, 0xa7, 0x80, 0x71, 0x1a, 0x69, 0x99, 0xbf, 0x38, 0x91, 0xc3, 0xd9, 0x22, 0xe9, 0x17, 0x4d,
	0x5a, 0x1e, 0xd3, 0xcb, 0xe6, 0xff, 0xcf, 0xf7, 0x0c, 0xef, 0x67, 0x35, 0xef, 0x30, 0x9a, 0x53,
	0x3a, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x21, 0x7c, 0x51, 0x6e, 0x01, 0x00, 0x00,
}
