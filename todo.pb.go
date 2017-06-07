// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

/*
Package requester is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	Task
	Config
	TaskList
	Response
	Options
*/
package requester

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Body  string `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Config struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type TaskList struct {
	List []*Task `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *TaskList) Reset()                    { *m = TaskList{} }
func (m *TaskList) String() string            { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()               {}
func (*TaskList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TaskList) GetList() []*Task {
	if m != nil {
		return m.List
	}
	return nil
}

type Response struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Options struct {
	Option string `protobuf:"bytes,1,opt,name=option" json:"option,omitempty"`
}

func (m *Options) Reset()                    { *m = Options{} }
func (m *Options) String() string            { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()               {}
func (*Options) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Options) GetOption() string {
	if m != nil {
		return m.Option
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "requester.Task")
	proto.RegisterType((*Config)(nil), "requester.Config")
	proto.RegisterType((*TaskList)(nil), "requester.TaskList")
	proto.RegisterType((*Response)(nil), "requester.Response")
	proto.RegisterType((*Options)(nil), "requester.Options")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x9b, 0x34, 0x4d, 0x93, 0x11, 0x14, 0x46, 0x91, 0x20, 0x1e, 0xe2, 0xea, 0x21, 0xa7,
	0x08, 0x55, 0xef, 0x42, 0x05, 0x2f, 0x82, 0xb0, 0xf8, 0x07, 0x52, 0x76, 0x2c, 0x8b, 0x31, 0x1b,
	0x33, 0xe3, 0xc1, 0x3f, 0xe5, 0x6f, 0x94, 0x5d, 0x13, 0xa9, 0xde, 0xbc, 0xbd, 0x37, 0xf3, 0xf2,
	0xbe, 0x0c, 0x0b, 0x20, 0xce, 0xb8, 0xba, 0x1f, 0x9c, 0x38, 0xcc, 0x07, 0x7a, 0x7b, 0x27, 0x16,
	0x1a, 0xd4, 0x2d, 0x24, 0x4f, 0x0d, 0xbf, 0xe0, 0x3e, 0xc4, 0xd6, 0x14, 0x51, 0x19, 0x55, 0x0b,
	0x1d, 0x5b, 0x83, 0x47, 0xb0, 0x10, 0x2b, 0x2d, 0x15, 0x71, 0x19, 0x55, 0xb9, 0xfe, 0x36, 0x88,
	0x90, 0x6c, 0x9c, 0xf9, 0x28, 0xe6, 0x61, 0x18, 0xb4, 0x3a, 0x85, 0x74, 0xed, 0xba, 0x67, 0xbb,
	0xf5, 0x5b, 0xd3, 0x48, 0x13, 0x5a, 0x72, 0x1d, 0xb4, 0xba, 0x84, 0xcc, 0xf7, 0x3f, 0x58, 0x16,
	0x3c, 0x87, 0xa4, 0xb5, 0x2c, 0x45, 0x54, 0xce, 0xab, 0xbd, 0xd5, 0x41, 0xfd, 0xf3, 0x17, 0xb5,
	0x8f, 0xe8, 0xb0, 0x54, 0x17, 0x90, 0x69, 0xe2, 0xde, 0x75, 0x4c, 0x58, 0xc0, 0xf2, 0x95, 0x98,
	0x9b, 0x2d, 0x8d, 0x9d, 0x93, 0x55, 0x67, 0xb0, 0x7c, 0xec, 0xc5, 0xba, 0x8e, 0xf1, 0x18, 0x52,
	0x17, 0xe4, 0x98, 0x19, 0xdd, 0xea, 0x33, 0x82, 0x5c, 0x4f, 0x04, 0xbc, 0x06, 0x58, 0x0f, 0xd4,
	0x08, 0x85, 0x6b, 0xff, 0xb2, 0x4f, 0x0e, 0x77, 0x06, 0x13, 0x5e, 0xcd, 0xf0, 0x06, 0xb2, 0x7b,
	0x12, 0x9f, 0x60, 0xc4, 0x9d, 0xc8, 0xc8, 0xfe, 0xf5, 0xd9, 0x74, 0xa6, 0x9a, 0x79, 0xd8, 0x1d,
	0xb5, 0xf4, 0x3f, 0xd8, 0x26, 0x0d, 0x8f, 0x73, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0x04, 0x03,
	0x4a, 0x43, 0xaa, 0x01, 0x00, 0x00,
}
