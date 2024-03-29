// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/incentive/pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Pool Info
type PoolInfo struct {
	// reward amount
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// reward wallet address
	RewardWallet string `protobuf:"bytes,2,opt,name=reward_wallet,json=rewardWallet,proto3" json:"reward_wallet,omitempty"`
	// multiplier for lp rewards
	Multiplier github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=multiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"multiplier"`
	// Block number since the creation of PoolInfo
	NumBlocks github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=num_blocks,json=numBlocks,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"num_blocks"`
	// Total dex rewards given
	DexRewardAmountGiven github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=dex_reward_amount_given,json=dexRewardAmountGiven,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"dex_reward_amount_given"`
	// Total eden rewards given
	EdenRewardAmountGiven github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=eden_reward_amount_given,json=edenRewardAmountGiven,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"eden_reward_amount_given"`
	// Eden APR, updated at every distribution
	EdenApr github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=eden_apr,json=edenApr,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"eden_apr"`
	// Dex APR, updated at every distribution
	DexApr github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=dex_apr,json=dexApr,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"dex_apr"`
}

func (m *PoolInfo) Reset()         { *m = PoolInfo{} }
func (m *PoolInfo) String() string { return proto.CompactTextString(m) }
func (*PoolInfo) ProtoMessage()    {}
func (*PoolInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dce8e3541398851, []int{0}
}
func (m *PoolInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolInfo.Merge(m, src)
}
func (m *PoolInfo) XXX_Size() int {
	return m.Size()
}
func (m *PoolInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PoolInfo proto.InternalMessageInfo

func (m *PoolInfo) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *PoolInfo) GetRewardWallet() string {
	if m != nil {
		return m.RewardWallet
	}
	return ""
}

func init() {
	proto.RegisterType((*PoolInfo)(nil), "elys.incentive.PoolInfo")
}

func init() { proto.RegisterFile("elys/incentive/pool.proto", fileDescriptor_1dce8e3541398851) }

var fileDescriptor_1dce8e3541398851 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x93, 0xd6, 0x26, 0x3a, 0xb4, 0x3d, 0x0c, 0x16, 0xa3, 0x87, 0x28, 0x2d, 0x14, 0x2f,
	0x26, 0x87, 0x7e, 0x02, 0xa5, 0x60, 0x73, 0x68, 0x29, 0xb9, 0x14, 0x7a, 0x09, 0x31, 0xf3, 0x9a,
	0x06, 0x27, 0x33, 0x61, 0x32, 0xd1, 0xf8, 0x2d, 0xf6, 0x63, 0x09, 0x7b, 0xf1, 0xb8, 0xec, 0x41,
	0x16, 0xfd, 0x22, 0xcb, 0x4c, 0x64, 0x11, 0x76, 0x2f, 0x9b, 0x53, 0xf2, 0xde, 0xff, 0xcd, 0xef,
	0x3f, 0x6f, 0xf8, 0xa3, 0x21, 0xd0, 0x5d, 0xe9, 0x67, 0x2c, 0x01, 0x26, 0xb3, 0x0d, 0xf8, 0x05,
	0xe7, 0xd4, 0x2b, 0x04, 0x97, 0x1c, 0x7f, 0x54, 0x92, 0xf7, 0x24, 0x8d, 0xfa, 0x29, 0x4f, 0xb9,
	0x96, 0x7c, 0xf5, 0xd7, 0x4c, 0x8d, 0x86, 0x09, 0x2f, 0x73, 0x5e, 0x46, 0x8d, 0xd0, 0x14, 0x8d,
	0xf4, 0xf9, 0xb6, 0x83, 0xba, 0xbf, 0x39, 0xa7, 0x01, 0xfb, 0xc7, 0xf1, 0x00, 0xd9, 0x8a, 0x1d,
	0x65, 0xc4, 0x31, 0x27, 0xe6, 0xb4, 0x13, 0x5a, 0xaa, 0x0c, 0x08, 0xfe, 0x82, 0x3e, 0x08, 0xd8,
	0xc6, 0x82, 0x44, 0xdb, 0x98, 0x52, 0x90, 0xce, 0x9b, 0x89, 0x39, 0xed, 0x85, 0xef, 0x9b, 0xe6,
	0x1f, 0xdd, 0xc3, 0xbf, 0x10, 0xca, 0x2b, 0x2a, 0xb3, 0x82, 0x66, 0x20, 0x9c, 0xb7, 0x6a, 0x62,
	0xe1, 0xed, 0x8f, 0x63, 0xe3, 0xfe, 0x38, 0xfe, 0x9a, 0x66, 0xf2, 0x7f, 0xb5, 0xf2, 0x12, 0x9e,
	0x5f, 0xfc, 0x2f, 0x9f, 0x59, 0x49, 0xd6, 0xbe, 0xdc, 0x15, 0x50, 0x7a, 0xdf, 0x21, 0x09, 0xaf,
	0x08, 0xf8, 0x27, 0x42, 0xac, 0xca, 0xa3, 0x15, 0xe5, 0xc9, 0xba, 0x74, 0x3a, 0xaf, 0xe6, 0x05,
	0x4c, 0x86, 0x3d, 0x56, 0xe5, 0x0b, 0x0d, 0xc0, 0x80, 0x06, 0x04, 0xea, 0xe8, 0xb2, 0x47, 0x9c,
	0xf3, 0x8a, 0xc9, 0x28, 0xcd, 0x36, 0xc0, 0x9c, 0x77, 0xad, 0xee, 0xda, 0x27, 0x50, 0x87, 0x9a,
	0x36, 0xd7, 0xb0, 0xa5, 0x62, 0xe1, 0x14, 0x39, 0x40, 0x80, 0xbd, 0xe8, 0x63, 0xb5, 0xda, 0xe1,
	0x93, 0xe2, 0x3d, 0x37, 0x0a, 0x50, 0x57, 0x1b, 0xc5, 0x85, 0x70, 0xec, 0x56, 0x0b, 0xd8, 0xea,
	0xfc, 0xbc, 0x10, 0x78, 0x89, 0x6c, 0xf5, 0x34, 0x8a, 0xd4, 0x6d, 0x45, 0xb2, 0x08, 0xd4, 0xf3,
	0x42, 0x2c, 0x7e, 0xec, 0x4f, 0xae, 0x79, 0x38, 0xb9, 0xe6, 0xc3, 0xc9, 0x35, 0x6f, 0xce, 0xae,
	0x71, 0x38, 0xbb, 0xc6, 0xdd, 0xd9, 0x35, 0xfe, 0x7a, 0x57, 0x24, 0x95, 0xd9, 0x19, 0x03, 0xb9,
	0xe5, 0x62, 0xad, 0x0b, 0xbf, 0xbe, 0x4a, 0xb7, 0xa6, 0xae, 0x2c, 0x1d, 0xcf, 0x6f, 0x8f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x78, 0x1a, 0x96, 0xfc, 0x02, 0x00, 0x00,
}

func (m *PoolInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.DexApr.Size()
		i -= size
		if _, err := m.DexApr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.EdenApr.Size()
		i -= size
		if _, err := m.EdenApr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.EdenRewardAmountGiven.Size()
		i -= size
		if _, err := m.EdenRewardAmountGiven.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.DexRewardAmountGiven.Size()
		i -= size
		if _, err := m.DexRewardAmountGiven.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.NumBlocks.Size()
		i -= size
		if _, err := m.NumBlocks.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Multiplier.Size()
		i -= size
		if _, err := m.Multiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.RewardWallet) > 0 {
		i -= len(m.RewardWallet)
		copy(dAtA[i:], m.RewardWallet)
		i = encodeVarintPool(dAtA, i, uint64(len(m.RewardWallet)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovPool(uint64(m.PoolId))
	}
	l = len(m.RewardWallet)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.Multiplier.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.NumBlocks.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.DexRewardAmountGiven.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.EdenRewardAmountGiven.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.EdenApr.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.DexApr.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: PoolInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardWallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardWallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Multiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumBlocks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DexRewardAmountGiven", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DexRewardAmountGiven.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenRewardAmountGiven", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenRewardAmountGiven.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenApr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenApr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DexApr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DexApr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
