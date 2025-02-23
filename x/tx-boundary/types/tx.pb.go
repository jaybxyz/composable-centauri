// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: centauri/txboundary/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateDelegateBoundary is the Msg/UpdateDelegateBoundary request type.
//
// Since: cosmos-sdk 0.47
type MsgUpdateDelegateBoundary struct {
	// authority is the address that controls the module (defaults to x/gov unless
	// overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// boundary defines the x/tx-boundary parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Boundary Boundary `protobuf:"bytes,2,opt,name=boundary,proto3" json:"boundary"`
}

func (m *MsgUpdateDelegateBoundary) Reset()         { *m = MsgUpdateDelegateBoundary{} }
func (m *MsgUpdateDelegateBoundary) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDelegateBoundary) ProtoMessage()    {}
func (*MsgUpdateDelegateBoundary) Descriptor() ([]byte, []int) {
	return fileDescriptor_148586c432693946, []int{0}
}
func (m *MsgUpdateDelegateBoundary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDelegateBoundary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDelegateBoundary.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDelegateBoundary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDelegateBoundary.Merge(m, src)
}
func (m *MsgUpdateDelegateBoundary) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDelegateBoundary) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDelegateBoundary.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDelegateBoundary proto.InternalMessageInfo

func (m *MsgUpdateDelegateBoundary) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateDelegateBoundary) GetBoundary() Boundary {
	if m != nil {
		return m.Boundary
	}
	return Boundary{}
}

// MsgUpdateDelegateBoundaryResponse defines the response structure for executing a
// MsgUpdateDelegateBoundary message.
//
// Since: cosmos-sdk 0.47
type MsgUpdateDelegateBoundaryResponse struct {
}

func (m *MsgUpdateDelegateBoundaryResponse) Reset()         { *m = MsgUpdateDelegateBoundaryResponse{} }
func (m *MsgUpdateDelegateBoundaryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDelegateBoundaryResponse) ProtoMessage()    {}
func (*MsgUpdateDelegateBoundaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_148586c432693946, []int{1}
}
func (m *MsgUpdateDelegateBoundaryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDelegateBoundaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDelegateBoundaryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDelegateBoundaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDelegateBoundaryResponse.Merge(m, src)
}
func (m *MsgUpdateDelegateBoundaryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDelegateBoundaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDelegateBoundaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDelegateBoundaryResponse proto.InternalMessageInfo

// MsgUpdateRedelegateBoundary is the Msg/MsgUpdateRedelegateBoundary request type.
//
// Since: cosmos-sdk 0.47
type MsgUpdateRedelegateBoundary struct {
	// authority is the address that controls the module (defaults to x/gov unless
	// overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// boundary defines the x/tx-boundary parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Boundary Boundary `protobuf:"bytes,2,opt,name=boundary,proto3" json:"boundary"`
}

func (m *MsgUpdateRedelegateBoundary) Reset()         { *m = MsgUpdateRedelegateBoundary{} }
func (m *MsgUpdateRedelegateBoundary) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateRedelegateBoundary) ProtoMessage()    {}
func (*MsgUpdateRedelegateBoundary) Descriptor() ([]byte, []int) {
	return fileDescriptor_148586c432693946, []int{2}
}
func (m *MsgUpdateRedelegateBoundary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateRedelegateBoundary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateRedelegateBoundary.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateRedelegateBoundary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateRedelegateBoundary.Merge(m, src)
}
func (m *MsgUpdateRedelegateBoundary) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateRedelegateBoundary) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateRedelegateBoundary.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateRedelegateBoundary proto.InternalMessageInfo

func (m *MsgUpdateRedelegateBoundary) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateRedelegateBoundary) GetBoundary() Boundary {
	if m != nil {
		return m.Boundary
	}
	return Boundary{}
}

// MsgUpdateRedelegateBoundaryResponse defines the response structure for executing a
// MsgUpdateRedelegateBoundary message.
//
// Since: cosmos-sdk 0.47
type MsgUpdateRedelegateBoundaryResponse struct {
}

func (m *MsgUpdateRedelegateBoundaryResponse) Reset()         { *m = MsgUpdateRedelegateBoundaryResponse{} }
func (m *MsgUpdateRedelegateBoundaryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateRedelegateBoundaryResponse) ProtoMessage()    {}
func (*MsgUpdateRedelegateBoundaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_148586c432693946, []int{3}
}
func (m *MsgUpdateRedelegateBoundaryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateRedelegateBoundaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateRedelegateBoundaryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateRedelegateBoundaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateRedelegateBoundaryResponse.Merge(m, src)
}
func (m *MsgUpdateRedelegateBoundaryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateRedelegateBoundaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateRedelegateBoundaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateRedelegateBoundaryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateDelegateBoundary)(nil), "centauri.txboundary.v1beta1.MsgUpdateDelegateBoundary")
	proto.RegisterType((*MsgUpdateDelegateBoundaryResponse)(nil), "centauri.txboundary.v1beta1.MsgUpdateDelegateBoundaryResponse")
	proto.RegisterType((*MsgUpdateRedelegateBoundary)(nil), "centauri.txboundary.v1beta1.MsgUpdateRedelegateBoundary")
	proto.RegisterType((*MsgUpdateRedelegateBoundaryResponse)(nil), "centauri.txboundary.v1beta1.MsgUpdateRedelegateBoundaryResponse")
}

func init() {
	proto.RegisterFile("centauri/txboundary/v1beta1/tx.proto", fileDescriptor_148586c432693946)
}

var fileDescriptor_148586c432693946 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x4e, 0xcd, 0x2b,
	0x49, 0x2c, 0x2d, 0xca, 0xd4, 0x2f, 0xa9, 0x48, 0xca, 0x2f, 0xcd, 0x4b, 0x49, 0x2c, 0xaa, 0xd4,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x86, 0xa9, 0xd2, 0x43, 0xa8, 0xd2, 0x83, 0xaa, 0x92, 0x12, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x33, 0x04, 0x51, 0x10, 0x5d, 0x52, 0x82, 0x89,
	0xb9, 0x99, 0x79, 0xf9, 0xfa, 0x60, 0x12, 0x2a, 0x24, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x66, 0xea,
	0x83, 0x58, 0x50, 0x51, 0x49, 0x88, 0x09, 0xf1, 0x10, 0x09, 0x08, 0x07, 0x2a, 0xa5, 0x85, 0xcf,
	0x7d, 0x70, 0xa7, 0x40, 0xd4, 0xca, 0x41, 0x1d, 0x92, 0x94, 0x58, 0x9c, 0x0a, 0x57, 0x93, 0x9c,
	0x9f, 0x99, 0x07, 0x91, 0x57, 0x7a, 0xca, 0xc8, 0x25, 0xe9, 0x5b, 0x9c, 0x1e, 0x5a, 0x90, 0x92,
	0x58, 0x92, 0xea, 0x92, 0x9a, 0x93, 0x9a, 0x9e, 0x58, 0x92, 0xea, 0x04, 0x35, 0x43, 0xc8, 0x8c,
	0x8b, 0x33, 0xb1, 0xb4, 0x24, 0x23, 0xbf, 0x28, 0xb3, 0xa4, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xd3, 0x49, 0xe2, 0xd2, 0x16, 0x5d, 0x11, 0xa8, 0x73, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b,
	0x83, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x83, 0x10, 0x4a, 0x85, 0x7c, 0xb8, 0x38, 0x60, 0xee, 0x90,
	0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd5, 0xc3, 0x13, 0x5c, 0x7a, 0x30, 0x0b, 0x9d, 0x38,
	0x4f, 0xdc, 0x93, 0x67, 0x58, 0xf1, 0x7c, 0x83, 0x16, 0x63, 0x10, 0xdc, 0x04, 0x2b, 0xa7, 0xa6,
	0xe7, 0x1b, 0xb4, 0x10, 0xa6, 0x77, 0x3d, 0xdf, 0xa0, 0xa5, 0x0f, 0x0f, 0x82, 0x0a, 0xe4, 0x40,
	0xc0, 0xe9, 0x13, 0x25, 0x65, 0x2e, 0x45, 0x9c, 0x92, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x4a, 0xcf, 0x19, 0xb9, 0xa4, 0xe1, 0xaa, 0x82, 0x52, 0x53, 0x86, 0x6f, 0x70, 0xa8, 0x72,
	0x29, 0xe3, 0xf1, 0x28, 0x2c, 0x40, 0x8c, 0x76, 0x30, 0x71, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x4d,
	0x60, 0xe4, 0x12, 0xc3, 0x95, 0x44, 0xf0, 0xfa, 0x04, 0xa7, 0x0b, 0xa4, 0xec, 0xc8, 0xd3, 0x07,
	0x73, 0x9a, 0xd0, 0x34, 0x46, 0x2e, 0x09, 0x9c, 0x11, 0x65, 0x41, 0x9c, 0xe1, 0x98, 0x3a, 0xa5,
	0x1c, 0xc8, 0xd5, 0x09, 0x73, 0x98, 0x14, 0x6b, 0x03, 0x28, 0xbe, 0x9c, 0x74, 0x4f, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x18, 0x14, 0x47, 0xba, 0xf0, 0x48, 0x2a, 0xa9,
	0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x67, 0x47, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e,
	0x6a, 0x87, 0xa9, 0x7c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	UpdateDelegateBoundary(ctx context.Context, in *MsgUpdateDelegateBoundary, opts ...grpc.CallOption) (*MsgUpdateDelegateBoundaryResponse, error)
	UpdateRedelegateBoundary(ctx context.Context, in *MsgUpdateRedelegateBoundary, opts ...grpc.CallOption) (*MsgUpdateRedelegateBoundaryResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateDelegateBoundary(ctx context.Context, in *MsgUpdateDelegateBoundary, opts ...grpc.CallOption) (*MsgUpdateDelegateBoundaryResponse, error) {
	out := new(MsgUpdateDelegateBoundaryResponse)
	err := c.cc.Invoke(ctx, "/centauri.txboundary.v1beta1.Msg/UpdateDelegateBoundary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRedelegateBoundary(ctx context.Context, in *MsgUpdateRedelegateBoundary, opts ...grpc.CallOption) (*MsgUpdateRedelegateBoundaryResponse, error) {
	out := new(MsgUpdateRedelegateBoundaryResponse)
	err := c.cc.Invoke(ctx, "/centauri.txboundary.v1beta1.Msg/UpdateRedelegateBoundary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	UpdateDelegateBoundary(context.Context, *MsgUpdateDelegateBoundary) (*MsgUpdateDelegateBoundaryResponse, error)
	UpdateRedelegateBoundary(context.Context, *MsgUpdateRedelegateBoundary) (*MsgUpdateRedelegateBoundaryResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateDelegateBoundary(ctx context.Context, req *MsgUpdateDelegateBoundary) (*MsgUpdateDelegateBoundaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDelegateBoundary not implemented")
}
func (*UnimplementedMsgServer) UpdateRedelegateBoundary(ctx context.Context, req *MsgUpdateRedelegateBoundary) (*MsgUpdateRedelegateBoundaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRedelegateBoundary not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateDelegateBoundary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDelegateBoundary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDelegateBoundary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/centauri.txboundary.v1beta1.Msg/UpdateDelegateBoundary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDelegateBoundary(ctx, req.(*MsgUpdateDelegateBoundary))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRedelegateBoundary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRedelegateBoundary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRedelegateBoundary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/centauri.txboundary.v1beta1.Msg/UpdateRedelegateBoundary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRedelegateBoundary(ctx, req.(*MsgUpdateRedelegateBoundary))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "centauri.txboundary.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDelegateBoundary",
			Handler:    _Msg_UpdateDelegateBoundary_Handler,
		},
		{
			MethodName: "UpdateRedelegateBoundary",
			Handler:    _Msg_UpdateRedelegateBoundary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "centauri/txboundary/v1beta1/tx.proto",
}

func (m *MsgUpdateDelegateBoundary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDelegateBoundary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDelegateBoundary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Boundary.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDelegateBoundaryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDelegateBoundaryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDelegateBoundaryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateRedelegateBoundary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateRedelegateBoundary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateRedelegateBoundary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Boundary.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateRedelegateBoundaryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateRedelegateBoundaryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateRedelegateBoundaryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateDelegateBoundary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Boundary.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateDelegateBoundaryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateRedelegateBoundary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Boundary.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateRedelegateBoundaryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateDelegateBoundary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateDelegateBoundary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDelegateBoundary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Boundary", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Boundary.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateDelegateBoundaryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateDelegateBoundaryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDelegateBoundaryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateRedelegateBoundary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateRedelegateBoundary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateRedelegateBoundary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Boundary", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Boundary.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateRedelegateBoundaryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateRedelegateBoundaryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateRedelegateBoundaryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
