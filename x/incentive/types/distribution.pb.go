// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/incentive/distribution.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// FeePool is the global fee pool for distribution.
type FeePool struct {
	CommunityPool github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=community_pool,json=communityPool,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"community_pool"`
}

func (m *FeePool) Reset()         { *m = FeePool{} }
func (m *FeePool) String() string { return proto.CompactTextString(m) }
func (*FeePool) ProtoMessage()    {}
func (*FeePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd7627d4268669e, []int{0}
}
func (m *FeePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeePool.Merge(m, src)
}
func (m *FeePool) XXX_Size() int {
	return m.Size()
}
func (m *FeePool) XXX_DiscardUnknown() {
	xxx_messageInfo_FeePool.DiscardUnknown(m)
}

var xxx_messageInfo_FeePool proto.InternalMessageInfo

func (m *FeePool) GetCommunityPool() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.CommunityPool
	}
	return nil
}

func init() {
	proto.RegisterType((*FeePool)(nil), "elys.incentive.FeePool")
}

func init() { proto.RegisterFile("elys/incentive/distribution.proto", fileDescriptor_4fd7627d4268669e) }

var fileDescriptor_4fd7627d4268669e = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xcd, 0xa9, 0x2c,
	0xd6, 0xcf, 0xcc, 0x4b, 0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0xd5, 0x4f, 0xc9, 0x2c, 0x2e, 0x29,
	0xca, 0x4c, 0x2a, 0x2d, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03,
	0x29, 0xd1, 0x83, 0x2b, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xe9, 0x83, 0x58, 0x10,
	0x55, 0x52, 0x72, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xfa, 0x49, 0x89, 0xc5, 0xa9, 0xfa, 0x65,
	0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0xfa, 0xc9, 0xf9, 0x99, 0x50, 0x53, 0x94, 0x9a, 0x19, 0xb9,
	0xd8, 0xdd, 0x52, 0x53, 0x03, 0xf2, 0xf3, 0x73, 0x84, 0x2a, 0xb8, 0xf8, 0x92, 0xf3, 0x73, 0x73,
	0x4b, 0xf3, 0x32, 0x4b, 0x2a, 0xe3, 0x0b, 0xf2, 0xf3, 0x73, 0x24, 0x18, 0x15, 0x98, 0x35, 0xb8,
	0x8d, 0x64, 0xf4, 0x20, 0x86, 0xe8, 0x81, 0x0c, 0xd1, 0x83, 0x1a, 0xa2, 0xe7, 0x92, 0x9a, 0xec,
	0x9c, 0x9f, 0x99, 0xe7, 0x64, 0x7c, 0xe2, 0x9e, 0x3c, 0xc3, 0xaa, 0xfb, 0xf2, 0xda, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0x4b, 0x21, 0x94, 0x6e, 0x71, 0x4a,
	0xb6, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0x31, 0x4c, 0x4f, 0x71, 0x10, 0x2f, 0xdc, 0x22, 0x90, 0xcd,
	0x4e, 0x3e, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83,
	0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43,
	0x94, 0x1e, 0x92, 0xc9, 0x20, 0x4f, 0xeb, 0xe6, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0x83, 0x39,
	0xfa, 0x15, 0x48, 0xc1, 0x04, 0xb6, 0x25, 0x89, 0x0d, 0xec, 0x35, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x90, 0x97, 0x92, 0x2f, 0x45, 0x01, 0x00, 0x00,
}

func (this *FeePool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FeePool)
	if !ok {
		that2, ok := that.(FeePool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.CommunityPool) != len(that1.CommunityPool) {
		return false
	}
	for i := range this.CommunityPool {
		if !this.CommunityPool[i].Equal(&that1.CommunityPool[i]) {
			return false
		}
	}
	return true
}
func (m *FeePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CommunityPool) > 0 {
		for iNdEx := len(m.CommunityPool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CommunityPool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDistribution(dAtA []byte, offset int, v uint64) int {
	offset -= sovDistribution(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CommunityPool) > 0 {
		for _, e := range m.CommunityPool {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func sovDistribution(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDistribution(x uint64) (n int) {
	return sovDistribution(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: FeePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommunityPool = append(m.CommunityPool, types.DecCoin{})
			if err := m.CommunityPool[len(m.CommunityPool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func skipDistribution(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDistribution
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
					return 0, ErrIntOverflowDistribution
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
					return 0, ErrIntOverflowDistribution
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
				return 0, ErrInvalidLengthDistribution
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDistribution
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDistribution
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDistribution        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDistribution          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDistribution = fmt.Errorf("proto: unexpected end of group")
)
