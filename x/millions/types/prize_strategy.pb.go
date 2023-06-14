// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lumnetwork/millions/prize_strategy.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type PrizeStrategy struct {
	PrizeBatches []PrizeBatch `protobuf:"bytes,1,rep,name=prize_batches,json=prizeBatches,proto3" json:"prize_batches"`
}

func (m *PrizeStrategy) Reset()         { *m = PrizeStrategy{} }
func (m *PrizeStrategy) String() string { return proto.CompactTextString(m) }
func (*PrizeStrategy) ProtoMessage()    {}
func (*PrizeStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e63ed8a3d849f1, []int{0}
}
func (m *PrizeStrategy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrizeStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrizeStrategy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrizeStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrizeStrategy.Merge(m, src)
}
func (m *PrizeStrategy) XXX_Size() int {
	return m.Size()
}
func (m *PrizeStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_PrizeStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_PrizeStrategy proto.InternalMessageInfo

func (m *PrizeStrategy) GetPrizeBatches() []PrizeBatch {
	if m != nil {
		return m.PrizeBatches
	}
	return nil
}

func init() {
	proto.RegisterType((*PrizeStrategy)(nil), "lumnetwork.millions.PrizeStrategy")
}

func init() {
	proto.RegisterFile("lumnetwork/millions/prize_strategy.proto", fileDescriptor_49e63ed8a3d849f1)
}

var fileDescriptor_49e63ed8a3d849f1 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0x29, 0xcd, 0xcd,
	0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0xcf, 0xcd, 0xcc, 0xc9, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6,
	0x2f, 0x28, 0xca, 0xac, 0x4a, 0x8d, 0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x49, 0x4d, 0xaf, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x46, 0xa8, 0xd4, 0x83, 0xa9, 0x94, 0x12, 0x49, 0xcf, 0x4f,
	0xcf, 0x07, 0xcb, 0xeb, 0x83, 0x58, 0x10, 0xa5, 0x52, 0xaa, 0xb8, 0x0d, 0x4d, 0x4a, 0x2c, 0x49,
	0xce, 0x80, 0x28, 0x53, 0x8a, 0xe6, 0xe2, 0x0d, 0x00, 0x09, 0x06, 0x43, 0x2d, 0x12, 0xf2, 0xe2,
	0xe2, 0x45, 0x52, 0x95, 0x5a, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xaf, 0x87, 0xc5,
	0x6a, 0x3d, 0xb0, 0x56, 0x27, 0x90, 0x42, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0x78, 0x0a,
	0xe0, 0x22, 0xa9, 0xc5, 0x4e, 0xee, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10,
	0xa5, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x53, 0x9a, 0xab,
	0x0b, 0x73, 0x69, 0x72, 0x46, 0x62, 0x66, 0x9e, 0x7e, 0x05, 0xc2, 0xc5, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0x60, 0xc7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x70, 0x31, 0xa4,
	0x2a, 0x01, 0x00, 0x00,
}

func (m *PrizeStrategy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrizeStrategy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrizeStrategy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PrizeBatches) > 0 {
		for iNdEx := len(m.PrizeBatches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PrizeBatches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPrizeStrategy(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrizeStrategy(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrizeStrategy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PrizeStrategy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PrizeBatches) > 0 {
		for _, e := range m.PrizeBatches {
			l = e.Size()
			n += 1 + l + sovPrizeStrategy(uint64(l))
		}
	}
	return n
}

func sovPrizeStrategy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrizeStrategy(x uint64) (n int) {
	return sovPrizeStrategy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PrizeStrategy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrizeStrategy
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
			return fmt.Errorf("proto: PrizeStrategy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrizeStrategy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrizeBatches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrizeStrategy
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
				return ErrInvalidLengthPrizeStrategy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrizeStrategy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrizeBatches = append(m.PrizeBatches, PrizeBatch{})
			if err := m.PrizeBatches[len(m.PrizeBatches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrizeStrategy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrizeStrategy
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
func skipPrizeStrategy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrizeStrategy
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
					return 0, ErrIntOverflowPrizeStrategy
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
					return 0, ErrIntOverflowPrizeStrategy
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
				return 0, ErrInvalidLengthPrizeStrategy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrizeStrategy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrizeStrategy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrizeStrategy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrizeStrategy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrizeStrategy = fmt.Errorf("proto: unexpected end of group")
)
