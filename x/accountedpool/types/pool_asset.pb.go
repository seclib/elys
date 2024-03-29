// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/tvl/pool_asset.proto

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

type PoolAsset struct {
	Token  types.Coin                             `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	Weight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"weight"`
}

func (m *PoolAsset) Reset()         { *m = PoolAsset{} }
func (m *PoolAsset) String() string { return proto.CompactTextString(m) }
func (*PoolAsset) ProtoMessage()    {}
func (*PoolAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_6805af7a13dc8985, []int{0}
}
func (m *PoolAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolAsset.Merge(m, src)
}
func (m *PoolAsset) XXX_Size() int {
	return m.Size()
}
func (m *PoolAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolAsset.DiscardUnknown(m)
}

var xxx_messageInfo_PoolAsset proto.InternalMessageInfo

func (m *PoolAsset) GetToken() types.Coin {
	if m != nil {
		return m.Token
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*PoolAsset)(nil), "elys.tvl.PoolAsset")
}

func init() { proto.RegisterFile("elys/tvl/pool_asset.proto", fileDescriptor_6805af7a13dc8985) }

var fileDescriptor_6805af7a13dc8985 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0x04, 0x15, 0x0d, 0x5b, 0xc4, 0xd0, 0x76, 0x70, 0x2b, 0x06, 0x94, 0xa5, 0xbe,
	0x2a, 0x88, 0x07, 0x20, 0x48, 0x48, 0x6c, 0xa8, 0x23, 0x0b, 0x4a, 0x82, 0x95, 0x46, 0x71, 0x7d,
	0xa3, 0xfa, 0x92, 0xd2, 0x57, 0x60, 0xe2, 0xb1, 0x3a, 0x76, 0x44, 0x0c, 0x15, 0x4a, 0x5e, 0x04,
	0xd9, 0xce, 0xc0, 0xe4, 0x9f, 0xcf, 0xdf, 0x91, 0xcf, 0x0d, 0xc7, 0x52, 0xed, 0x0c, 0x50, 0xa3,
	0xa0, 0x46, 0x54, 0xaf, 0xa9, 0x31, 0x92, 0x44, 0xbd, 0x41, 0xc2, 0xe8, 0xdc, 0x22, 0x41, 0x8d,
	0x9a, 0x5c, 0x16, 0x58, 0xa0, 0xbb, 0x04, 0xbb, 0xf3, 0x7c, 0xc2, 0x73, 0x34, 0x6b, 0x34, 0x90,
	0xa5, 0x46, 0x42, 0xb3, 0xc8, 0x24, 0xa5, 0x0b, 0xc8, 0xb1, 0xd4, 0x9e, 0x5f, 0x7d, 0xb2, 0x70,
	0xf8, 0x8c, 0xa8, 0xee, 0x6d, 0x66, 0x74, 0x17, 0x9e, 0x11, 0x56, 0x52, 0x8f, 0xd8, 0x8c, 0xc5,
	0x17, 0x37, 0x63, 0xe1, 0x6d, 0x61, 0x6d, 0xd1, 0xdb, 0xe2, 0x01, 0x4b, 0x9d, 0x9c, 0xee, 0x8f,
	0xd3, 0x60, 0xe9, 0x5f, 0x47, 0x8f, 0xe1, 0x60, 0x2b, 0xcb, 0x62, 0x45, 0xa3, 0x93, 0x19, 0x8b,
	0x87, 0x89, 0xb0, 0xf0, 0xe7, 0x38, 0xbd, 0x2e, 0x4a, 0x5a, 0xbd, 0x67, 0x22, 0xc7, 0x35, 0xf4,
	0xff, 0xf0, 0xcb, 0xdc, 0xbc, 0x55, 0x40, 0xbb, 0x5a, 0x1a, 0xf1, 0xa4, 0x69, 0xd9, 0xdb, 0x49,
	0xb2, 0x6f, 0x39, 0x3b, 0xb4, 0x9c, 0xfd, 0xb6, 0x9c, 0x7d, 0x75, 0x3c, 0x38, 0x74, 0x3c, 0xf8,
	0xee, 0x78, 0xf0, 0x12, 0xff, 0x4b, 0xb2, 0x8d, 0xe7, 0x5a, 0xd2, 0x16, 0x37, 0x95, 0x3b, 0xc0,
	0x87, 0x9b, 0x8d, 0xcb, 0xcb, 0x06, 0xae, 0xd7, 0xed, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x59, 0xf8, 0x08, 0x34, 0x01, 0x00, 0x00,
}

func (m *PoolAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolAsset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPoolAsset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPoolAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Token.Size()
	n += 1 + l + sovPoolAsset(uint64(l))
	l = m.Weight.Size()
	n += 1 + l + sovPoolAsset(uint64(l))
	return n
}

func sovPoolAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolAsset(x uint64) (n int) {
	return sovPoolAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolAsset
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
			return fmt.Errorf("proto: PoolAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolAsset
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
				return ErrInvalidLengthPoolAsset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolAsset
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
				return ErrInvalidLengthPoolAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolAsset
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
func skipPoolAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolAsset
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
					return 0, ErrIntOverflowPoolAsset
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
					return 0, ErrIntOverflowPoolAsset
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
				return 0, ErrInvalidLengthPoolAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolAsset = fmt.Errorf("proto: unexpected end of group")
)
