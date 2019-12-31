// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

package streamergrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Control struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Control) Reset()         { *m = Control{} }
func (m *Control) String() string { return proto.CompactTextString(m) }
func (*Control) ProtoMessage()    {}
func (*Control) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{0}
}

func (m *Control) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Control.Unmarshal(m, b)
}
func (m *Control) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Control.Marshal(b, m, deterministic)
}
func (m *Control) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Control.Merge(m, src)
}
func (m *Control) XXX_Size() int {
	return xxx_messageInfo_Control.Size(m)
}
func (m *Control) XXX_DiscardUnknown() {
	xxx_messageInfo_Control.DiscardUnknown(m)
}

var xxx_messageInfo_Control proto.InternalMessageInfo

func (m *Control) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Control) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type NatsMeta struct {
	Sequence             uint64   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsMeta) Reset()         { *m = NatsMeta{} }
func (m *NatsMeta) String() string { return proto.CompactTextString(m) }
func (*NatsMeta) ProtoMessage()    {}
func (*NatsMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{1}
}

func (m *NatsMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsMeta.Unmarshal(m, b)
}
func (m *NatsMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsMeta.Marshal(b, m, deterministic)
}
func (m *NatsMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsMeta.Merge(m, src)
}
func (m *NatsMeta) XXX_Size() int {
	return xxx_messageInfo_NatsMeta.Size(m)
}
func (m *NatsMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsMeta.DiscardUnknown(m)
}

var xxx_messageInfo_NatsMeta proto.InternalMessageInfo

func (m *NatsMeta) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *NatsMeta) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type Message struct {
	NatsMeta             *NatsMeta `protobuf:"bytes,1,opt,name=natsMeta,proto3" json:"natsMeta,omitempty"`
	Data                 string    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetNatsMeta() *NatsMeta {
	if m != nil {
		return m.NatsMeta
	}
	return nil
}

func (m *Message) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Control)(nil), "streamergrpc.Control")
	proto.RegisterType((*NatsMeta)(nil), "streamergrpc.NatsMeta")
	proto.RegisterType((*Message)(nil), "streamergrpc.Message")
}

func init() { proto.RegisterFile("main.proto", fileDescriptor_7ed94b0a22d11796) }

var fileDescriptor_7ed94b0a22d11796 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x8d, 0xac, 0xdb, 0xec, 0xe8, 0x69, 0x50, 0x29, 0x8b, 0x07, 0xc9, 0x69, 0x4f, 0x45,
	0xea, 0x49, 0xf0, 0xa6, 0x07, 0x2f, 0x2b, 0x58, 0x3f, 0xc1, 0xd8, 0x1d, 0x4a, 0xa1, 0x49, 0x6a,
	0x32, 0xf5, 0xf3, 0x0b, 0x31, 0xad, 0xba, 0xb7, 0x79, 0x8f, 0x79, 0xbf, 0xf9, 0x03, 0x60, 0xa9,
	0x77, 0xd5, 0x18, 0xbc, 0x78, 0xbc, 0x88, 0x12, 0x98, 0x2c, 0x87, 0x2e, 0x8c, 0xad, 0x79, 0x80,
	0xe2, 0xc9, 0x3b, 0x09, 0x7e, 0xc0, 0x12, 0x8a, 0xd6, 0x5b, 0x4b, 0xee, 0x50, 0xaa, 0x5b, 0xb5,
	0xdb, 0x34, 0xb3, 0xc4, 0x4b, 0x38, 0xfb, 0xa2, 0x61, 0xe2, 0xf2, 0x34, 0xf9, 0x3f, 0xc2, 0x3c,
	0x83, 0x7e, 0x25, 0x89, 0x7b, 0x16, 0xc2, 0x2d, 0xe8, 0xc8, 0x9f, 0x13, 0xbb, 0x96, 0x53, 0x78,
	0xd5, 0x2c, 0x1a, 0x6f, 0x60, 0x23, 0xbd, 0xe5, 0x28, 0x64, 0xc7, 0x4c, 0xf8, 0x35, 0xcc, 0x1b,
	0x14, 0x7b, 0x8e, 0x91, 0x3a, 0xc6, 0x1a, 0xb4, 0xcb, 0xc0, 0x04, 0x39, 0xaf, 0xaf, 0xab, 0xbf,
	0xcb, 0x56, 0xf3, 0xb8, 0x66, 0xe9, 0x43, 0x84, 0xd5, 0x81, 0x84, 0x32, 0x37, 0xd5, 0xf5, 0x0b,
	0xe8, 0x8c, 0x8c, 0xf8, 0x08, 0xeb, 0xf7, 0x84, 0xc0, 0xab, 0xff, 0xac, 0x7c, 0xf5, 0xf6, 0xc8,
	0xce, 0x41, 0x73, 0xb2, 0x53, 0x77, 0xea, 0x63, 0x9d, 0x5e, 0x76, 0xff, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0xe2, 0x70, 0xcb, 0x89, 0x40, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessagesClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Messages_StreamClient, error)
}

type messagesClient struct {
	cc *grpc.ClientConn
}

func NewMessagesClient(cc *grpc.ClientConn) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Messages_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messages_serviceDesc.Streams[0], "/streamergrpc.Messages/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesStreamClient{stream}
	return x, nil
}

type Messages_StreamClient interface {
	Send(*Control) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type messagesStreamClient struct {
	grpc.ClientStream
}

func (x *messagesStreamClient) Send(m *Control) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messagesStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagesServer is the server API for Messages service.
type MessagesServer interface {
	Stream(Messages_StreamServer) error
}

func RegisterMessagesServer(s *grpc.Server, srv MessagesServer) {
	s.RegisterService(&_Messages_serviceDesc, srv)
}

func _Messages_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagesServer).Stream(&messagesStreamServer{stream})
}

type Messages_StreamServer interface {
	Send(*Message) error
	Recv() (*Control, error)
	grpc.ServerStream
}

type messagesStreamServer struct {
	grpc.ServerStream
}

func (x *messagesStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messagesStreamServer) Recv() (*Control, error) {
	m := new(Control)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Messages_serviceDesc = grpc.ServiceDesc{
	ServiceName: "streamergrpc.Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Messages_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "main.proto",
}