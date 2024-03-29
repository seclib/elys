// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/leveragelp/types.proto

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

type Position struct {
	Address           string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Collateral        types.Coin                             `protobuf:"bytes,2,opt,name=collateral,proto3" json:"collateral"`
	Liabilities       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=liabilities,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liabilities"`
	InterestPaid      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=interest_paid,json=interestPaid,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"interest_paid"`
	Leverage          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=leverage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"leverage"`
	LeveragedLpAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=leveraged_lp_amount,json=leveragedLpAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"leveraged_lp_amount"`
	PositionHealth    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=position_health,json=positionHealth,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"position_health"`
	Id                uint64                                 `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	AmmPoolId         uint64                                 `protobuf:"varint,9,opt,name=amm_pool_id,json=ammPoolId,proto3" json:"amm_pool_id,omitempty"`
	StopLossPrice     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=stop_loss_price,json=stopLossPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stop_loss_price"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_992d513dd201f55b, []int{0}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Position.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return m.Size()
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Position) GetCollateral() types.Coin {
	if m != nil {
		return m.Collateral
	}
	return types.Coin{}
}

func (m *Position) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Position) GetAmmPoolId() uint64 {
	if m != nil {
		return m.AmmPoolId
	}
	return 0
}

type WhiteList struct {
	ValidatorList []string `protobuf:"bytes,1,rep,name=validator_list,json=validatorList,proto3" json:"validator_list,omitempty"`
}

func (m *WhiteList) Reset()         { *m = WhiteList{} }
func (m *WhiteList) String() string { return proto.CompactTextString(m) }
func (*WhiteList) ProtoMessage()    {}
func (*WhiteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_992d513dd201f55b, []int{1}
}
func (m *WhiteList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhiteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhiteList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhiteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhiteList.Merge(m, src)
}
func (m *WhiteList) XXX_Size() int {
	return m.Size()
}
func (m *WhiteList) XXX_DiscardUnknown() {
	xxx_messageInfo_WhiteList.DiscardUnknown(m)
}

var xxx_messageInfo_WhiteList proto.InternalMessageInfo

func (m *WhiteList) GetValidatorList() []string {
	if m != nil {
		return m.ValidatorList
	}
	return nil
}

func init() {
	proto.RegisterType((*Position)(nil), "elys.leveragelp.Position")
	proto.RegisterType((*WhiteList)(nil), "elys.leveragelp.WhiteList")
}

func init() { proto.RegisterFile("elys/leveragelp/types.proto", fileDescriptor_992d513dd201f55b) }

var fileDescriptor_992d513dd201f55b = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xae, 0xdb, 0x5a, 0x97, 0xb6, 0xc2, 0x70, 0x30, 0x43, 0xca, 0xaa, 0x49, 0xa0,
	0x5e, 0x96, 0x68, 0xe3, 0x03, 0x20, 0x0a, 0x07, 0x8a, 0x7a, 0x88, 0x82, 0xc4, 0x24, 0x0e, 0x44,
	0x4e, 0x6c, 0xb5, 0x4f, 0x73, 0xf2, 0xac, 0xd8, 0x2b, 0xec, 0x5b, 0x70, 0xe5, 0x1b, 0xed, 0xb8,
	0x23, 0xe2, 0x30, 0xa1, 0xf6, 0x8b, 0x20, 0xa7, 0x4d, 0xa9, 0xc4, 0x89, 0x9e, 0xf2, 0x5e, 0xfe,
	0xef, 0xfd, 0xf4, 0xb7, 0xad, 0x3f, 0x79, 0x2e, 0xd5, 0xad, 0x09, 0x95, 0x5c, 0xc8, 0x92, 0xcf,
	0xa4, 0xd2, 0xa1, 0xbd, 0xd5, 0xd2, 0x04, 0xba, 0x44, 0x8b, 0x74, 0xe0, 0xc4, 0xe0, 0xaf, 0x78,
	0xf2, 0x74, 0x86, 0x33, 0xac, 0xb4, 0xd0, 0x55, 0xeb, 0xb1, 0x13, 0x3f, 0x43, 0x93, 0xa3, 0x09,
	0x53, 0x6e, 0x64, 0xb8, 0xb8, 0x48, 0xa5, 0xe5, 0x17, 0x61, 0x86, 0x50, 0xac, 0xf5, 0xb3, 0x1f,
	0x87, 0xa4, 0x1d, 0xa1, 0x01, 0x0b, 0x58, 0x50, 0x46, 0x8e, 0xb9, 0x10, 0xa5, 0x34, 0x86, 0x79,
	0x43, 0x6f, 0xd4, 0x89, 0xeb, 0x96, 0xbe, 0x26, 0x24, 0x43, 0xa5, 0xb8, 0x95, 0x25, 0x57, 0xac,
	0x39, 0xf4, 0x46, 0xdd, 0xcb, 0x67, 0xc1, 0x9a, 0x1d, 0x38, 0x76, 0xb0, 0x61, 0x07, 0x6f, 0x11,
	0x8a, 0x71, 0xeb, 0xee, 0xe1, 0xb4, 0x11, 0xef, 0xac, 0xd0, 0x88, 0x74, 0x15, 0xf0, 0x14, 0x14,
	0x58, 0x90, 0x86, 0x1d, 0x38, 0xfc, 0x38, 0x70, 0x63, 0xbf, 0x1e, 0x4e, 0x5f, 0xce, 0xc0, 0xce,
	0x6f, 0xd2, 0x20, 0xc3, 0x3c, 0xdc, 0xf8, 0x5d, 0x7f, 0xce, 0x8d, 0xb8, 0xde, 0x9c, 0x7a, 0x52,
	0xd8, 0x78, 0x17, 0x41, 0x3f, 0x92, 0x1e, 0x14, 0x56, 0x96, 0xd2, 0xd8, 0x44, 0x73, 0x10, 0xac,
	0xb5, 0x17, 0xf3, 0x51, 0x0d, 0x89, 0x38, 0x08, 0xfa, 0x81, 0xb4, 0xeb, 0x2b, 0x65, 0x87, 0xff,
	0xcd, 0x7b, 0x27, 0xb3, 0x78, 0xbb, 0x4f, 0xbf, 0x90, 0x27, 0x75, 0x2d, 0x12, 0xa5, 0x13, 0x9e,
	0xe3, 0x4d, 0x61, 0xd9, 0xd1, 0x5e, 0x36, 0x1f, 0x6f, 0x51, 0x53, 0xfd, 0xa6, 0x02, 0xd1, 0x2b,
	0x32, 0xd0, 0x9b, 0x97, 0x4b, 0xe6, 0x92, 0x2b, 0x3b, 0x67, 0xc7, 0x7b, 0x59, 0xee, 0xd7, 0x98,
	0xf7, 0x15, 0x85, 0xf6, 0x49, 0x13, 0x04, 0x6b, 0x0f, 0xbd, 0x51, 0x2b, 0x6e, 0x82, 0xa0, 0x3e,
	0xe9, 0xf2, 0x3c, 0x4f, 0x34, 0xa2, 0x4a, 0x40, 0xb0, 0x4e, 0x25, 0x74, 0x78, 0x9e, 0x47, 0x88,
	0x6a, 0x22, 0xe8, 0x27, 0x32, 0x30, 0x16, 0x75, 0xa2, 0xd0, 0x98, 0x44, 0x97, 0x90, 0x49, 0x46,
	0xf6, 0x32, 0xd2, 0x73, 0x98, 0x29, 0x1a, 0x13, 0x39, 0xc8, 0xd9, 0x25, 0xe9, 0x5c, 0xcd, 0xc1,
	0xca, 0x29, 0x18, 0x4b, 0x5f, 0x90, 0xfe, 0x82, 0x2b, 0x10, 0xdc, 0x62, 0x99, 0x28, 0x30, 0x96,
	0x79, 0xc3, 0x83, 0x51, 0x27, 0xee, 0x6d, 0xff, 0xba, 0xb1, 0xf1, 0xe4, 0x6e, 0xe9, 0x7b, 0xf7,
	0x4b, 0xdf, 0xfb, 0xbd, 0xf4, 0xbd, 0xef, 0x2b, 0xbf, 0x71, 0xbf, 0xf2, 0x1b, 0x3f, 0x57, 0x7e,
	0xe3, 0x73, 0xb8, 0x63, 0xc2, 0x65, 0xe7, 0xbc, 0x90, 0xf6, 0x2b, 0x96, 0xd7, 0x55, 0x13, 0x7e,
	0xfb, 0x27, 0x67, 0xe9, 0x51, 0x95, 0x90, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xe5,
	0x47, 0x51, 0x87, 0x03, 0x00, 0x00,
}

func (m *Position) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Position) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Position) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.StopLossPrice.Size()
		i -= size
		if _, err := m.StopLossPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.AmmPoolId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AmmPoolId))
		i--
		dAtA[i] = 0x48
	}
	if m.Id != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.PositionHealth.Size()
		i -= size
		if _, err := m.PositionHealth.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.LeveragedLpAmount.Size()
		i -= size
		if _, err := m.LeveragedLpAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Leverage.Size()
		i -= size
		if _, err := m.Leverage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InterestPaid.Size()
		i -= size
		if _, err := m.InterestPaid.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Liabilities.Size()
		i -= size
		if _, err := m.Liabilities.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Collateral.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WhiteList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhiteList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhiteList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorList) > 0 {
		for iNdEx := len(m.ValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ValidatorList[iNdEx])
			copy(dAtA[i:], m.ValidatorList[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorList[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Position) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Collateral.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Liabilities.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.InterestPaid.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Leverage.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.LeveragedLpAmount.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.PositionHealth.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Id != 0 {
		n += 1 + sovTypes(uint64(m.Id))
	}
	if m.AmmPoolId != 0 {
		n += 1 + sovTypes(uint64(m.AmmPoolId))
	}
	l = m.StopLossPrice.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *WhiteList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorList) > 0 {
		for _, s := range m.ValidatorList {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Position) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Position: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Collateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liabilities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liabilities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestPaid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestPaid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leverage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Leverage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeveragedLpAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LeveragedLpAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionHealth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PositionHealth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmmPoolId", wireType)
			}
			m.AmmPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AmmPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StopLossPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StopLossPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *WhiteList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: WhiteList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhiteList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorList = append(m.ValidatorList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
