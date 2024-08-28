// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentratedliquidity/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	// authorized_tick_spacing is an array of uint64s that represents the tick
	// spacing values concentrated-liquidity pools can be created with. For
	// example, an authorized_tick_spacing of [1, 10, 30] allows for pools
	// to be created with tick spacing of 1, 10, or 30.
	AuthorizedTickSpacing   []uint64                      `protobuf:"varint,1,rep,packed,name=authorized_tick_spacing,json=authorizedTickSpacing,proto3" json:"authorized_tick_spacing,omitempty" yaml:"authorized_tick_spacing"`
	AuthorizedSpreadFactors []cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,rep,name=authorized_spread_factors,json=authorizedSpreadFactors,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"authorized_spread_factors" yaml:"authorized_spread_factors"`
	// balancer_shares_reward_discount is the rate by which incentives flowing
	// from CL to Balancer pools will be discounted to encourage LPs to migrate.
	// e.g. a rate of 0.05 means Balancer LPs get 5% less incentives than full
	// range CL LPs.
	// This field can range from (0,1]. If set to 1, it indicates that all
	// incentives stay at cl pool.
	BalancerSharesRewardDiscount cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=balancer_shares_reward_discount,json=balancerSharesRewardDiscount,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"balancer_shares_reward_discount" yaml:"balancer_shares_reward_discount"`
	// DEPRECATED: authorized_quote_denoms is a list of quote denoms that can be
	// used as token1 when creating a pool. We limit the quote assets to a small
	// set for the purposes of having convenient price increments stemming from
	// tick to price conversion. These increments are in a human readable
	// magnitude only for token1 as a quote. For limit orders in the future, this
	// will be a desirable property in terms of UX as to allow users to set limit
	// orders at prices in terms of token1 (quote asset) that are easy to reason
	// about.
	AuthorizedQuoteDenoms []string        `protobuf:"bytes,4,rep,name=authorized_quote_denoms,json=authorizedQuoteDenoms,proto3" json:"authorized_quote_denoms,omitempty" yaml:"authorized_quote_denoms",deprecated:"true"` // Deprecated: Do not use.
	AuthorizedUptimes     []time.Duration `protobuf:"bytes,5,rep,name=authorized_uptimes,json=authorizedUptimes,proto3,stdduration" json:"duration,omitempty" yaml:"authorized_uptimes"`
	// is_permissionless_pool_creation_enabled is a boolean that determines if
	// concentrated liquidity pools can be created via message. At launch,
	// we consider allowing only governance to create pools, and then later
	// allowing permissionless pool creation by switching this flag to true
	// with a governance proposal.
	IsPermissionlessPoolCreationEnabled bool `protobuf:"varint,6,opt,name=is_permissionless_pool_creation_enabled,json=isPermissionlessPoolCreationEnabled,proto3" json:"is_permissionless_pool_creation_enabled,omitempty" yaml:"is_permissionless_pool_creation_enabled"`
	// unrestricted_pool_creator_whitelist is a list of addresses that are
	// allowed to bypass restrictions on permissionless supercharged pool
	// creation, like pool_creation_enabled, restricted quote assets, no
	// double creation of pools, etc.
	UnrestrictedPoolCreatorWhitelist []string `protobuf:"bytes,7,rep,name=unrestricted_pool_creator_whitelist,json=unrestrictedPoolCreatorWhitelist,proto3" json:"unrestricted_pool_creator_whitelist,omitempty" yaml:"unrestricted_pool_creator_whitelist"`
	HookGasLimit                     uint64   `protobuf:"varint,8,opt,name=hook_gas_limit,json=hookGasLimit,proto3" json:"hook_gas_limit,omitempty" yaml:"hook_gas_limit"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a3f6981164624c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAuthorizedTickSpacing() []uint64 {
	if m != nil {
		return m.AuthorizedTickSpacing
	}
	return nil
}

// Deprecated: Do not use.
func (m *Params) GetAuthorizedQuoteDenoms() []string {
	if m != nil {
		return m.AuthorizedQuoteDenoms
	}
	return nil
}

func (m *Params) GetAuthorizedUptimes() []time.Duration {
	if m != nil {
		return m.AuthorizedUptimes
	}
	return nil
}

func (m *Params) GetIsPermissionlessPoolCreationEnabled() bool {
	if m != nil {
		return m.IsPermissionlessPoolCreationEnabled
	}
	return false
}

func (m *Params) GetUnrestrictedPoolCreatorWhitelist() []string {
	if m != nil {
		return m.UnrestrictedPoolCreatorWhitelist
	}
	return nil
}

func (m *Params) GetHookGasLimit() uint64 {
	if m != nil {
		return m.HookGasLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.concentratedliquidity.Params")
}

func init() {
	proto.RegisterFile("osmosis/concentratedliquidity/params.proto", fileDescriptor_42a3f6981164624c)
}

var fileDescriptor_42a3f6981164624c = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4f, 0x6b, 0xd4, 0x4e,
	0x18, 0xc7, 0x37, 0xbf, 0xed, 0xaf, 0x7f, 0xa2, 0x08, 0x06, 0x8b, 0xd9, 0xaa, 0x49, 0x98, 0x82,
	0x2e, 0xa5, 0x4d, 0xa4, 0x42, 0x0f, 0xf5, 0x20, 0xc4, 0xd5, 0x5e, 0x2a, 0xd4, 0x54, 0x11, 0x8a,
	0x30, 0xcc, 0x4e, 0xa6, 0xd9, 0x61, 0x93, 0x3c, 0xe9, 0xcc, 0xc4, 0xba, 0x82, 0x27, 0x11, 0x3c,
	0x7a, 0xf0, 0xe0, 0x0b, 0xf2, 0xd0, 0x63, 0x8f, 0xe2, 0x21, 0x4a, 0x7b, 0xf3, 0xb8, 0xaf, 0x40,
	0x36, 0xd9, 0xb5, 0x59, 0xaa, 0xd8, 0x5b, 0x66, 0xbe, 0x9f, 0xe7, 0xcf, 0x3c, 0xf9, 0xf2, 0xe8,
	0x2b, 0x20, 0x13, 0x90, 0x5c, 0x7a, 0x14, 0x52, 0xca, 0x52, 0x25, 0x88, 0x62, 0x61, 0xcc, 0x0f,
	0x72, 0x1e, 0x72, 0x35, 0xf0, 0x32, 0x22, 0x48, 0x22, 0xdd, 0x4c, 0x80, 0x02, 0xe3, 0xd6, 0x98,
	0x75, 0xff, 0xc8, 0x2e, 0x5d, 0x8b, 0x20, 0x82, 0x92, 0xf4, 0x46, 0x5f, 0x55, 0xd0, 0x92, 0x15,
	0x01, 0x44, 0x31, 0xf3, 0xca, 0x53, 0x37, 0xdf, 0xf7, 0xc2, 0x5c, 0x10, 0xc5, 0x21, 0xad, 0x74,
	0xf4, 0x65, 0x4e, 0x9f, 0xdd, 0x29, 0xab, 0x18, 0x7b, 0xfa, 0x75, 0x92, 0xab, 0x1e, 0x08, 0xfe,
	0x86, 0x85, 0x58, 0x71, 0xda, 0xc7, 0x32, 0x23, 0x94, 0xa7, 0x91, 0xa9, 0x39, 0xcd, 0xf6, 0x8c,
	0x8f, 0x86, 0x85, 0x6d, 0x0d, 0x48, 0x12, 0x6f, 0xa2, 0xbf, 0x80, 0x28, 0x58, 0x3c, 0x53, 0x9e,
	0x71, 0xda, 0xdf, 0xad, 0xee, 0x8d, 0x77, 0x9a, 0xde, 0xaa, 0xc5, 0xc8, 0x4c, 0x30, 0x12, 0xe2,
	0x7d, 0x42, 0x15, 0x08, 0x69, 0xfe, 0xe7, 0x34, 0xdb, 0x0b, 0xfe, 0xd6, 0x51, 0x61, 0x37, 0xbe,
	0x15, 0xf6, 0x0d, 0x5a, 0x3e, 0x54, 0x86, 0x7d, 0x97, 0x83, 0x97, 0x10, 0xd5, 0x73, 0xb7, 0x59,
	0x44, 0xe8, 0xa0, 0xc3, 0xe8, 0xb0, 0xb0, 0x9d, 0x73, 0x1d, 0x4c, 0x67, 0x43, 0x41, 0xed, 0x19,
	0xbb, 0xa5, 0xf4, 0xb8, 0x52, 0x8c, 0x4f, 0x9a, 0x6e, 0x77, 0x49, 0x4c, 0x52, 0xca, 0x04, 0x96,
	0x3d, 0x22, 0x98, 0xc4, 0x82, 0x1d, 0x12, 0x11, 0xe2, 0x90, 0x4b, 0x0a, 0x79, 0xaa, 0xcc, 0xa6,
	0xa3, 0xb5, 0x17, 0xfc, 0x27, 0x17, 0xeb, 0xe5, 0x76, 0xd5, 0xcb, 0x3f, 0x72, 0xa2, 0xe0, 0xe6,
	0x84, 0xd8, 0x2d, 0x81, 0xa0, 0xd4, 0x3b, 0x63, 0xd9, 0x48, 0xa7, 0x06, 0x7f, 0x90, 0x83, 0x62,
	0x38, 0x64, 0x29, 0x24, 0xd2, 0x9c, 0x29, 0x27, 0xb3, 0x31, 0x2c, 0xec, 0xbb, 0xe7, 0x9e, 0x5d,
	0x07, 0xd1, 0x6a, 0xc8, 0x32, 0xc1, 0xe8, 0xc8, 0x12, 0x9b, 0x48, 0x89, 0x9c, 0x21, 0x53, 0xab,
	0xff, 0x8c, 0xa7, 0x23, 0xb8, 0x53, 0xb2, 0xc6, 0x7b, 0x4d, 0x37, 0x6a, 0x79, 0xf2, 0x4c, 0xf1,
	0x84, 0x49, 0xf3, 0x7f, 0xa7, 0xd9, 0xbe, 0xb4, 0xde, 0x72, 0x2b, 0xc7, 0xb8, 0x13, 0xc7, 0xb8,
	0x9d, 0xb1, 0x63, 0xfc, 0xfb, 0xa3, 0xa1, 0xfc, 0x2c, 0x6c, 0x63, 0xe2, 0xa1, 0x55, 0x48, 0xb8,
	0x62, 0x49, 0xa6, 0x06, 0xc3, 0xc2, 0x6e, 0x9d, 0x6b, 0x70, 0x9c, 0x18, 0x7d, 0xfe, 0x6e, 0x6b,
	0xc1, 0xd5, 0x33, 0xe1, 0x79, 0x75, 0x6f, 0x7c, 0xd0, 0xf4, 0x3b, 0x5c, 0xe2, 0x8c, 0x89, 0x84,
	0x4b, 0xc9, 0x21, 0x8d, 0x99, 0x94, 0x38, 0x03, 0x88, 0x31, 0x15, 0xac, 0xac, 0x80, 0x59, 0x4a,
	0xba, 0x31, 0x0b, 0xcd, 0x59, 0x47, 0x6b, 0xcf, 0xfb, 0xeb, 0xc3, 0xc2, 0x76, 0xab, 0x3a, 0x17,
	0x0c, 0x44, 0xc1, 0x32, 0x97, 0x3b, 0x53, 0xe0, 0x0e, 0x40, 0xfc, 0x70, 0x8c, 0x3d, 0xaa, 0x28,
	0xe3, 0xad, 0xbe, 0x9c, 0xa7, 0x82, 0x49, 0x25, 0x38, 0x55, 0x2c, 0xac, 0xe5, 0x02, 0x81, 0x0f,
	0x7b, 0x5c, 0xb1, 0x98, 0x4b, 0x65, 0xce, 0x95, 0xbf, 0xc3, 0x1d, 0x16, 0xf6, 0x4a, 0xd5, 0xc5,
	0x05, 0x82, 0x50, 0xe0, 0xd4, 0xa9, 0xdf, 0xd5, 0x41, 0xbc, 0x98, 0x20, 0xc6, 0x03, 0xfd, 0x4a,
	0x0f, 0xa0, 0x8f, 0x23, 0x22, 0x71, 0xcc, 0x13, 0xae, 0xcc, 0x79, 0x47, 0x6b, 0xcf, 0xf8, 0xad,
	0x61, 0x61, 0x2f, 0x56, 0x95, 0xa6, 0x75, 0x14, 0x5c, 0x1e, 0x5d, 0x6c, 0x11, 0xb9, 0x3d, 0x3a,
	0xfa, 0x2f, 0x8f, 0x4e, 0x2c, 0xed, 0xf8, 0xc4, 0xd2, 0x7e, 0x9c, 0x58, 0xda, 0xc7, 0x53, 0xab,
	0x71, 0x7c, 0x6a, 0x35, 0xbe, 0x9e, 0x5a, 0x8d, 0x3d, 0x3f, 0xe2, 0xaa, 0x97, 0x77, 0x5d, 0x0a,
	0x89, 0x37, 0x5e, 0x20, 0x6b, 0x31, 0xe9, 0xca, 0xc9, 0xc1, 0x7b, 0xb5, 0xbe, 0xe1, 0xbd, 0x9e,
	0xda, 0x3f, 0x6b, 0x67, 0x0b, 0x48, 0x0d, 0x32, 0x26, 0xbb, 0xb3, 0xa5, 0x17, 0xee, 0xfd, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x26, 0x79, 0x09, 0x52, 0xae, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HookGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HookGasLimit))
		i--
		dAtA[i] = 0x40
	}
	if len(m.UnrestrictedPoolCreatorWhitelist) > 0 {
		for iNdEx := len(m.UnrestrictedPoolCreatorWhitelist) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UnrestrictedPoolCreatorWhitelist[iNdEx])
			copy(dAtA[i:], m.UnrestrictedPoolCreatorWhitelist[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.UnrestrictedPoolCreatorWhitelist[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.IsPermissionlessPoolCreationEnabled {
		i--
		if m.IsPermissionlessPoolCreationEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.AuthorizedUptimes) > 0 {
		for iNdEx := len(m.AuthorizedUptimes) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.AuthorizedUptimes[iNdEx], dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.AuthorizedUptimes[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintParams(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.AuthorizedQuoteDenoms) > 0 {
		for iNdEx := len(m.AuthorizedQuoteDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthorizedQuoteDenoms[iNdEx])
			copy(dAtA[i:], m.AuthorizedQuoteDenoms[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AuthorizedQuoteDenoms[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.BalancerSharesRewardDiscount.Size()
		i -= size
		if _, err := m.BalancerSharesRewardDiscount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.AuthorizedSpreadFactors) > 0 {
		for iNdEx := len(m.AuthorizedSpreadFactors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.AuthorizedSpreadFactors[iNdEx].Size()
				i -= size
				if _, err := m.AuthorizedSpreadFactors[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.AuthorizedTickSpacing) > 0 {
		dAtA2 := make([]byte, len(m.AuthorizedTickSpacing)*10)
		var j1 int
		for _, num := range m.AuthorizedTickSpacing {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintParams(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuthorizedTickSpacing) > 0 {
		l = 0
		for _, e := range m.AuthorizedTickSpacing {
			l += sovParams(uint64(e))
		}
		n += 1 + sovParams(uint64(l)) + l
	}
	if len(m.AuthorizedSpreadFactors) > 0 {
		for _, e := range m.AuthorizedSpreadFactors {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.BalancerSharesRewardDiscount.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.AuthorizedQuoteDenoms) > 0 {
		for _, s := range m.AuthorizedQuoteDenoms {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.AuthorizedUptimes) > 0 {
		for _, e := range m.AuthorizedUptimes {
			l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(e)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.IsPermissionlessPoolCreationEnabled {
		n += 2
	}
	if len(m.UnrestrictedPoolCreatorWhitelist) > 0 {
		for _, s := range m.UnrestrictedPoolCreatorWhitelist {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.HookGasLimit != 0 {
		n += 1 + sovParams(uint64(m.HookGasLimit))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.AuthorizedTickSpacing = append(m.AuthorizedTickSpacing, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthParams
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthParams
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.AuthorizedTickSpacing) == 0 {
					m.AuthorizedTickSpacing = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.AuthorizedTickSpacing = append(m.AuthorizedTickSpacing, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedTickSpacing", wireType)
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedSpreadFactors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.LegacyDec
			m.AuthorizedSpreadFactors = append(m.AuthorizedSpreadFactors, v)
			if err := m.AuthorizedSpreadFactors[len(m.AuthorizedSpreadFactors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalancerSharesRewardDiscount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BalancerSharesRewardDiscount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedQuoteDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedQuoteDenoms = append(m.AuthorizedQuoteDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedUptimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedUptimes = append(m.AuthorizedUptimes, time.Duration(0))
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&(m.AuthorizedUptimes[len(m.AuthorizedUptimes)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPermissionlessPoolCreationEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPermissionlessPoolCreationEnabled = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnrestrictedPoolCreatorWhitelist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnrestrictedPoolCreatorWhitelist = append(m.UnrestrictedPoolCreatorWhitelist, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HookGasLimit", wireType)
			}
			m.HookGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HookGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
