// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wts.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Script struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interpretter         string   `protobuf:"bytes,2,opt,name=interpretter,proto3" json:"interpretter,omitempty"`
	Script               string   `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
	Checksum             string   `protobuf:"bytes,4,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Script) Reset()         { *m = Script{} }
func (m *Script) String() string { return proto.CompactTextString(m) }
func (*Script) ProtoMessage()    {}
func (*Script) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73afdd414cba0f9, []int{0}
}

func (m *Script) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Script.Unmarshal(m, b)
}
func (m *Script) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Script.Marshal(b, m, deterministic)
}
func (m *Script) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Script.Merge(m, src)
}
func (m *Script) XXX_Size() int {
	return xxx_messageInfo_Script.Size(m)
}
func (m *Script) XXX_DiscardUnknown() {
	xxx_messageInfo_Script.DiscardUnknown(m)
}

var xxx_messageInfo_Script proto.InternalMessageInfo

func (m *Script) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Script) GetInterpretter() string {
	if m != nil {
		return m.Interpretter
	}
	return ""
}

func (m *Script) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *Script) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

type ScriptRecord struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Public               bool     `protobuf:"varint,2,opt,name=public,proto3" json:"public,omitempty"`
	Script               *Script  `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScriptRecord) Reset()         { *m = ScriptRecord{} }
func (m *ScriptRecord) String() string { return proto.CompactTextString(m) }
func (*ScriptRecord) ProtoMessage()    {}
func (*ScriptRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73afdd414cba0f9, []int{1}
}

func (m *ScriptRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScriptRecord.Unmarshal(m, b)
}
func (m *ScriptRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScriptRecord.Marshal(b, m, deterministic)
}
func (m *ScriptRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScriptRecord.Merge(m, src)
}
func (m *ScriptRecord) XXX_Size() int {
	return xxx_messageInfo_ScriptRecord.Size(m)
}
func (m *ScriptRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ScriptRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ScriptRecord proto.InternalMessageInfo

func (m *ScriptRecord) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ScriptRecord) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func (m *ScriptRecord) GetScript() *Script {
	if m != nil {
		return m.Script
	}
	return nil
}

type PutResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73afdd414cba0f9, []int{2}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

func (m *PutResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73afdd414cba0f9, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type GetResponse struct {
	Record               *ScriptRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Err                  string        `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73afdd414cba0f9, []int{4}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetRecord() *ScriptRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *GetResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Script)(nil), "pb.Script")
	proto.RegisterType((*ScriptRecord)(nil), "pb.ScriptRecord")
	proto.RegisterType((*PutResponse)(nil), "pb.PutResponse")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
}

func init() { proto.RegisterFile("wts.proto", fileDescriptor_f73afdd414cba0f9) }

var fileDescriptor_f73afdd414cba0f9 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xdb, 0xa4, 0x84, 0x76, 0x52, 0xfe, 0xff, 0x32, 0x07, 0x09, 0xbd, 0x28, 0x7b, 0x0a,
	0x1e, 0x72, 0x88, 0x7e, 0x87, 0xe2, 0xad, 0x6c, 0x05, 0x0f, 0x82, 0x60, 0xe2, 0x80, 0x41, 0x9b,
	0xac, 0xbb, 0xb3, 0xf8, 0xf5, 0x25, 0x93, 0x4d, 0xda, 0xa2, 0xb7, 0x79, 0xfb, 0x32, 0xef, 0xfd,
	0x86, 0xc0, 0xea, 0x9b, 0x5d, 0x61, 0x6c, 0xc7, 0x1d, 0x46, 0xa6, 0x52, 0x0c, 0xc9, 0xa1, 0xb6,
	0x8d, 0x61, 0x44, 0x58, 0xb4, 0xaf, 0x47, 0xca, 0xe6, 0x37, 0xf3, 0x7c, 0xa5, 0x65, 0x46, 0x05,
	0xeb, 0xa6, 0x65, 0xb2, 0xc6, 0x12, 0x33, 0xd9, 0x2c, 0x12, 0xef, 0xe2, 0x0d, 0xaf, 0x20, 0x71,
	0x92, 0x90, 0xc5, 0xe2, 0x06, 0x85, 0x5b, 0x58, 0xd6, 0xef, 0x54, 0x7f, 0x38, 0x7f, 0xcc, 0x16,
	0xe2, 0x4c, 0x5a, 0xbd, 0xc0, 0x7a, 0x68, 0xd5, 0x54, 0x77, 0xf6, 0xad, 0xef, 0xf6, 0x8e, 0xec,
	0xd8, 0xdd, 0xcf, 0x7d, 0xae, 0xf1, 0xd5, 0x67, 0x53, 0x4b, 0xeb, 0x52, 0x07, 0x85, 0xea, 0xa2,
	0x2f, 0x2d, 0xa1, 0x30, 0x55, 0x11, 0xd2, 0x82, 0xa3, 0xae, 0x21, 0xdd, 0x7b, 0xd6, 0xe4, 0x4c,
	0xd7, 0x3a, 0xc2, 0x0d, 0xc4, 0x64, 0xc7, 0xf4, 0x7e, 0x54, 0xf7, 0x00, 0x3b, 0x62, 0x4d, 0x5f,
	0x9e, 0xdc, 0xdf, 0xa7, 0x8f, 0x48, 0xd1, 0x09, 0x49, 0x3d, 0x40, 0x2a, 0x5b, 0x21, 0x36, 0x87,
	0xc4, 0x0a, 0xbf, 0x2c, 0xa6, 0xe5, 0xe6, 0x8c, 0x44, 0xde, 0x75, 0xf0, 0x47, 0x80, 0x68, 0x02,
	0x28, 0x9f, 0x21, 0x7e, 0x7a, 0x3c, 0x60, 0x0e, 0xf1, 0x8e, 0x18, 0xff, 0xf5, 0x9b, 0x27, 0xa0,
	0xed, 0xff, 0x49, 0x0f, 0x55, 0x6a, 0x86, 0xb7, 0x10, 0xef, 0x3d, 0xe3, 0xaf, 0x8e, 0xe1, 0xdb,
	0xb3, 0x6b, 0xd5, 0xac, 0x4a, 0xe4, 0xff, 0xde, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xb0,
	0x4f, 0x41, 0xec, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WTSClient is the client API for WTS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WTSClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *ScriptRecord, opts ...grpc.CallOption) (*PutResponse, error)
}

type wTSClient struct {
	cc *grpc.ClientConn
}

func NewWTSClient(cc *grpc.ClientConn) WTSClient {
	return &wTSClient{cc}
}

func (c *wTSClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pb.WTS/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wTSClient) Put(ctx context.Context, in *ScriptRecord, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/pb.WTS/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WTSServer is the server API for WTS service.
type WTSServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *ScriptRecord) (*PutResponse, error)
}

// UnimplementedWTSServer can be embedded to have forward compatible implementations.
type UnimplementedWTSServer struct {
}

func (*UnimplementedWTSServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedWTSServer) Put(ctx context.Context, req *ScriptRecord) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}

func RegisterWTSServer(s *grpc.Server, srv WTSServer) {
	s.RegisterService(&_WTS_serviceDesc, srv)
}

func _WTS_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WTSServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WTS/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WTSServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WTS_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScriptRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WTSServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WTS/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WTSServer).Put(ctx, req.(*ScriptRecord))
	}
	return interceptor(ctx, in, info, handler)
}

var _WTS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WTS",
	HandlerType: (*WTSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _WTS_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _WTS_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wts.proto",
}
