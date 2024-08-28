// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/cosmwasmpool/v1beta1/model/tx.proto

package model

import (
	context "context"
	fmt "fmt"
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

// ===================== MsgCreateCosmwasmPool
type MsgCreateCosmWasmPool struct {
	CodeId         uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id"`
	InstantiateMsg []byte `protobuf:"bytes,2,opt,name=instantiate_msg,json=instantiateMsg,proto3" json:"instantiate_msg,omitempty" yaml:"instantiate_msg"`
	Sender         string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
}

func (m *MsgCreateCosmWasmPool) Reset()         { *m = MsgCreateCosmWasmPool{} }
func (m *MsgCreateCosmWasmPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCosmWasmPool) ProtoMessage()    {}
func (*MsgCreateCosmWasmPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ff1ac8555d314d1, []int{0}
}
func (m *MsgCreateCosmWasmPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCosmWasmPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCosmWasmPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCosmWasmPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCosmWasmPool.Merge(m, src)
}
func (m *MsgCreateCosmWasmPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCosmWasmPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCosmWasmPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCosmWasmPool proto.InternalMessageInfo

func (m *MsgCreateCosmWasmPool) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *MsgCreateCosmWasmPool) GetInstantiateMsg() []byte {
	if m != nil {
		return m.InstantiateMsg
	}
	return nil
}

func (m *MsgCreateCosmWasmPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// Returns a unique poolID to identify the pool with.
type MsgCreateCosmWasmPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateCosmWasmPoolResponse) Reset()         { *m = MsgCreateCosmWasmPoolResponse{} }
func (m *MsgCreateCosmWasmPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCosmWasmPoolResponse) ProtoMessage()    {}
func (*MsgCreateCosmWasmPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ff1ac8555d314d1, []int{1}
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCosmWasmPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCosmWasmPoolResponse.Merge(m, src)
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCosmWasmPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCosmWasmPoolResponse proto.InternalMessageInfo

func (m *MsgCreateCosmWasmPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateCosmWasmPool)(nil), "osmosis.cosmwasmpool.v1beta1.MsgCreateCosmWasmPool")
	proto.RegisterType((*MsgCreateCosmWasmPoolResponse)(nil), "osmosis.cosmwasmpool.v1beta1.MsgCreateCosmWasmPoolResponse")
}

func init() {
	proto.RegisterFile("osmosis/cosmwasmpool/v1beta1/model/tx.proto", fileDescriptor_2ff1ac8555d314d1)
}

var fileDescriptor_2ff1ac8555d314d1 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x3f, 0x8b, 0xd4, 0x40,
	0x14, 0xdf, 0x51, 0xc9, 0xe1, 0xa0, 0x27, 0x37, 0xf8, 0x67, 0x09, 0x5e, 0xb2, 0xc4, 0x66, 0xbd,
	0xe3, 0x32, 0xec, 0x1d, 0x88, 0x9c, 0x5d, 0xf6, 0x9a, 0x2d, 0x0e, 0x24, 0x8d, 0x60, 0x73, 0x4c,
	0x92, 0x21, 0x06, 0x32, 0x79, 0x4b, 0xde, 0xb8, 0x9e, 0xad, 0x85, 0x85, 0x95, 0x8d, 0xdf, 0xe3,
	0x3e, 0x86, 0xe5, 0x96, 0x56, 0x41, 0x92, 0x62, 0xfb, 0xfd, 0x04, 0x92, 0x7f, 0xb8, 0x4a, 0xb0,
	0xb0, 0x49, 0xde, 0xfb, 0xfd, 0x99, 0x79, 0xef, 0xcd, 0xa3, 0xc7, 0x80, 0x0a, 0x30, 0x41, 0x1e,
	0x02, 0xaa, 0x0f, 0x02, 0xd5, 0x12, 0x20, 0xe5, 0xab, 0x59, 0x20, 0xb5, 0x98, 0x71, 0x05, 0x91,
	0x4c, 0xb9, 0xbe, 0x76, 0x97, 0x39, 0x68, 0x60, 0x4f, 0x3b, 0xb1, 0xbb, 0x2b, 0x76, 0x3b, 0xb1,
	0xf9, 0x24, 0x6c, 0x68, 0xae, 0x30, 0xe6, 0xab, 0x59, 0xfd, 0x6b, 0x6d, 0xe6, 0x81, 0x50, 0x49,
	0x06, 0xbc, 0xf9, 0x76, 0xd0, 0xc3, 0x18, 0x62, 0x68, 0x42, 0x5e, 0x47, 0x2d, 0xea, 0x54, 0x84,
	0x3e, 0xba, 0xc4, 0x78, 0x9e, 0x4b, 0xa1, 0xe5, 0x1c, 0x50, 0xbd, 0x11, 0xa8, 0x5e, 0x03, 0xa4,
	0xec, 0x98, 0xee, 0x85, 0x10, 0xc9, 0xab, 0x24, 0x1a, 0x93, 0x09, 0x99, 0xde, 0xf1, 0xd8, 0xb6,
	0xb0, 0xf7, 0x3f, 0x0a, 0x95, 0x9e, 0x3b, 0x1d, 0xe1, 0xf8, 0x46, 0x1d, 0x2d, 0x22, 0x36, 0xa7,
	0x0f, 0x92, 0x0c, 0xb5, 0xc8, 0x74, 0x22, 0xb4, 0xbc, 0x52, 0x18, 0x8f, 0x6f, 0x4d, 0xc8, 0xf4,
	0x9e, 0x67, 0x6e, 0x0b, 0xfb, 0x71, 0x6b, 0xfa, 0x4b, 0xe0, 0xf8, 0xfb, 0x3b, 0xc8, 0x25, 0xc6,
	0xec, 0x39, 0x35, 0x50, 0x66, 0x91, 0xcc, 0xc7, 0xb7, 0x27, 0x64, 0x7a, 0xd7, 0x3b, 0xd8, 0x16,
	0xf6, 0xfd, 0xd6, 0xdb, 0xe2, 0x8e, 0xdf, 0x09, 0xce, 0x4f, 0x3e, 0x6d, 0x6e, 0x8e, 0xba, 0xe4,
	0xcb, 0xe6, 0xe6, 0xe8, 0xb0, 0x9f, 0xe9, 0x60, 0x2f, 0xce, 0x05, 0x3d, 0x1c, 0x24, 0x7c, 0x89,
	0x4b, 0xc8, 0x50, 0xb2, 0x67, 0x74, 0xaf, 0x1e, 0xec, 0xef, 0x66, 0x69, 0x59, 0xd8, 0x46, 0x2d,
	0x59, 0x5c, 0xf8, 0x46, 0x4d, 0x2d, 0xa2, 0xd3, 0x6f, 0x84, 0xd2, 0xfe, 0x18, 0xc8, 0xd9, 0x67,
	0x42, 0xd9, 0xc0, 0xdc, 0xce, 0xdc, 0x7f, 0x3d, 0x99, 0x3b, 0x58, 0x87, 0xf9, 0xea, 0x3f, 0x4c,
	0x7d, 0xf1, 0x9e, 0xff, 0xbd, 0xb4, 0xc8, 0xba, 0xb4, 0xc8, 0xcf, 0xd2, 0x22, 0x5f, 0x2b, 0x6b,
	0xb4, 0xae, 0xac, 0xd1, 0x8f, 0xca, 0x1a, 0xbd, 0x7d, 0x19, 0x27, 0xfa, 0xdd, 0xfb, 0xc0, 0x0d,
	0x41, 0xf1, 0xee, 0x82, 0x93, 0x54, 0x04, 0xd8, 0x27, 0x7c, 0x75, 0xfa, 0x82, 0x5f, 0xff, 0xb9,
	0x88, 0xcd, 0x02, 0x06, 0x46, 0xb3, 0x1e, 0x67, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0x63,
	0xd9, 0x41, 0xad, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgCreatorClient is the client API for MsgCreator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgCreatorClient interface {
	CreateCosmWasmPool(ctx context.Context, in *MsgCreateCosmWasmPool, opts ...grpc.CallOption) (*MsgCreateCosmWasmPoolResponse, error)
}

type msgCreatorClient struct {
	cc grpc1.ClientConn
}

func NewMsgCreatorClient(cc grpc1.ClientConn) MsgCreatorClient {
	return &msgCreatorClient{cc}
}

func (c *msgCreatorClient) CreateCosmWasmPool(ctx context.Context, in *MsgCreateCosmWasmPool, opts ...grpc.CallOption) (*MsgCreateCosmWasmPoolResponse, error) {
	out := new(MsgCreateCosmWasmPoolResponse)
	err := c.cc.Invoke(ctx, "/osmosis.cosmwasmpool.v1beta1.MsgCreator/CreateCosmWasmPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgCreatorServer is the server API for MsgCreator service.
type MsgCreatorServer interface {
	CreateCosmWasmPool(context.Context, *MsgCreateCosmWasmPool) (*MsgCreateCosmWasmPoolResponse, error)
}

// UnimplementedMsgCreatorServer can be embedded to have forward compatible implementations.
type UnimplementedMsgCreatorServer struct {
}

func (*UnimplementedMsgCreatorServer) CreateCosmWasmPool(ctx context.Context, req *MsgCreateCosmWasmPool) (*MsgCreateCosmWasmPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCosmWasmPool not implemented")
}

func RegisterMsgCreatorServer(s grpc1.Server, srv MsgCreatorServer) {
	s.RegisterService(&_MsgCreator_serviceDesc, srv)
}

func _MsgCreator_CreateCosmWasmPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateCosmWasmPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgCreatorServer).CreateCosmWasmPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.cosmwasmpool.v1beta1.MsgCreator/CreateCosmWasmPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgCreatorServer).CreateCosmWasmPool(ctx, req.(*MsgCreateCosmWasmPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgCreator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.cosmwasmpool.v1beta1.MsgCreator",
	HandlerType: (*MsgCreatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCosmWasmPool",
			Handler:    _MsgCreator_CreateCosmWasmPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/cosmwasmpool/v1beta1/model/tx.proto",
}

func (m *MsgCreateCosmWasmPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCosmWasmPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCosmWasmPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InstantiateMsg) > 0 {
		i -= len(m.InstantiateMsg)
		copy(dAtA[i:], m.InstantiateMsg)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InstantiateMsg)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateCosmWasmPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCosmWasmPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCosmWasmPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
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
func (m *MsgCreateCosmWasmPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovTx(uint64(m.CodeId))
	}
	l = len(m.InstantiateMsg)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateCosmWasmPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateCosmWasmPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateCosmWasmPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCosmWasmPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstantiateMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InstantiateMsg = append(m.InstantiateMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.InstantiateMsg == nil {
				m.InstantiateMsg = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateCosmWasmPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateCosmWasmPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCosmWasmPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
