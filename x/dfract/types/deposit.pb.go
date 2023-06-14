// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lumnetwork/dfract/deposit.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

type Deposit struct {
	DepositorAddress string     `protobuf:"bytes,1,opt,name=depositor_address,json=depositorAddress,proto3" json:"depositor_address,omitempty"`
	Amount           types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
	CreatedAt        time.Time  `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae190bb6b609260f, []int{0}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func (m *Deposit) GetDepositorAddress() string {
	if m != nil {
		return m.DepositorAddress
	}
	return ""
}

func (m *Deposit) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *Deposit) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Deposit)(nil), "lumnetwork.dfract.Deposit")
}

func init() { proto.RegisterFile("lumnetwork/dfract/deposit.proto", fileDescriptor_ae190bb6b609260f) }

var fileDescriptor_ae190bb6b609260f = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xe3, 0xef, 0x43, 0x85, 0x86, 0x85, 0x46, 0x0c, 0xa5, 0x83, 0x53, 0x31, 0x55, 0xaa,
	0xb0, 0x55, 0x18, 0x98, 0xdb, 0xc2, 0x0d, 0x54, 0x4c, 0x2c, 0x95, 0xe3, 0xb8, 0xa9, 0x45, 0x9d,
	0x13, 0xd9, 0x27, 0xfc, 0xdc, 0x45, 0x2f, 0x86, 0x8b, 0xe8, 0xd8, 0x91, 0x09, 0x50, 0x7b, 0x23,
	0xa8, 0x71, 0x02, 0x9b, 0xed, 0xf3, 0xbc, 0x7a, 0xfc, 0x9e, 0x30, 0x5e, 0x95, 0x26, 0x57, 0xf8,
	0x02, 0xf6, 0x89, 0xa7, 0x0b, 0x2b, 0x24, 0xf2, 0x54, 0x15, 0xe0, 0x34, 0xb2, 0xc2, 0x02, 0x42,
	0xd4, 0xf9, 0x03, 0x98, 0x07, 0x7a, 0xe7, 0x19, 0x64, 0x50, 0x4d, 0xf9, 0xe1, 0xe4, 0xc1, 0x5e,
	0x9c, 0x01, 0x64, 0x2b, 0xc5, 0xab, 0x5b, 0x52, 0x2e, 0x38, 0x6a, 0xa3, 0x1c, 0x0a, 0x53, 0xd4,
	0x00, 0x95, 0xe0, 0x0c, 0x38, 0x9e, 0x08, 0xa7, 0xf8, 0xf3, 0x28, 0x51, 0x28, 0x46, 0x5c, 0x82,
	0xce, 0xfd, 0xfc, 0xf2, 0x9d, 0x84, 0xc7, 0x77, 0xde, 0x1d, 0x0d, 0xc3, 0x4e, 0xfd, 0x0d, 0xb0,
	0x73, 0x91, 0xa6, 0x56, 0x39, 0xd7, 0x25, 0x7d, 0x32, 0x68, 0xcf, 0xce, 0x7e, 0x07, 0x63, 0xff,
	0x1e, 0xdd, 0x86, 0x2d, 0x61, 0xa0, 0xcc, 0xb1, 0xfb, 0xaf, 0x4f, 0x06, 0xa7, 0xd7, 0x17, 0xcc,
	0x9b, 0xd8, 0xc1, 0xc4, 0x6a, 0x13, 0x9b, 0x82, 0xce, 0x27, 0x47, 0x9b, 0xcf, 0x38, 0x98, 0xd5,
	0x78, 0x34, 0x0d, 0x43, 0x69, 0x95, 0x40, 0x95, 0xce, 0x05, 0x76, 0xff, 0x57, 0xe1, 0x1e, 0xf3,
	0x3d, 0x58, 0xd3, 0x83, 0x3d, 0x34, 0x3d, 0x26, 0x27, 0x87, 0xf4, 0xfa, 0x2b, 0x26, 0xb3, 0x76,
	0x9d, 0x1b, 0xe3, 0xe4, 0x7e, 0xb3, 0xa3, 0x64, 0xbb, 0xa3, 0xe4, 0x7b, 0x47, 0xc9, 0x7a, 0x4f,
	0x83, 0xed, 0x9e, 0x06, 0x1f, 0x7b, 0x1a, 0x3c, 0x0e, 0x33, 0x8d, 0xcb, 0x32, 0x61, 0x12, 0x0c,
	0x5f, 0x95, 0xe6, 0xaa, 0xd9, 0xb3, 0x5c, 0x0a, 0x9d, 0xf3, 0xd7, 0x66, 0xdf, 0xf8, 0x56, 0x28,
	0x97, 0xb4, 0x2a, 0xdf, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x51, 0x6b, 0x3e, 0x91,
	0x01, 0x00, 0x00,
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintDeposit(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DepositorAddress) > 0 {
		i -= len(m.DepositorAddress)
		copy(dAtA[i:], m.DepositorAddress)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.DepositorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeposit(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeposit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepositorAddress)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovDeposit(uint64(l))
	return n
}

func sovDeposit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeposit(x uint64) (n int) {
	return sovDeposit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeposit
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
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeposit
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
func skipDeposit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
				return 0, ErrInvalidLengthDeposit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeposit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeposit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeposit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeposit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeposit = fmt.Errorf("proto: unexpected end of group")
)
