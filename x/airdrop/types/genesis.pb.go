// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lumnetwork/chain/airdrop/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the airdrop module's genesis state.
type GenesisState struct {
	ModuleAccountBalance types.Coin    `protobuf:"bytes,1,opt,name=module_account_balance,json=moduleAccountBalance,proto3" json:"module_account_balance" yaml:"module_account_balance"`
	Params               Params        `protobuf:"bytes,2,opt,name=params,proto3" json:"params" yaml:"params"`
	ClaimRecords         []ClaimRecord `protobuf:"bytes,3,rep,name=claim_records,json=claimRecords,proto3" json:"claim_records" yaml:"claim_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_935e0de71e22ed43, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetModuleAccountBalance() types.Coin {
	if m != nil {
		return m.ModuleAccountBalance
	}
	return types.Coin{}
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClaimRecords() []ClaimRecord {
	if m != nil {
		return m.ClaimRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lumnetwork.chain.airdrop.GenesisState")
}

func init() {
	proto.RegisterFile("lumnetwork/chain/airdrop/genesis.proto", fileDescriptor_935e0de71e22ed43)
}

var fileDescriptor_935e0de71e22ed43 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0x5b, 0x48, 0x58, 0x14, 0xd8, 0x34, 0xfc, 0x7f, 0x2a, 0xd1, 0x42, 0x1a, 0x31, 0x2c,
	0x74, 0x26, 0xe0, 0xce, 0x9d, 0x25, 0xd1, 0xa5, 0xa6, 0xee, 0xdc, 0x90, 0xe9, 0x30, 0x29, 0x8d,
	0x9d, 0xde, 0xa6, 0x33, 0x45, 0x79, 0x0b, 0x5f, 0xc9, 0x1d, 0x4b, 0x96, 0xae, 0x88, 0x81, 0x37,
	0xf0, 0x09, 0x0c, 0x9d, 0x89, 0x84, 0xc4, 0xee, 0x26, 0xf7, 0x7e, 0x73, 0xce, 0xc9, 0x3d, 0xd6,
	0x45, 0x52, 0xf0, 0x94, 0xc9, 0x57, 0xc8, 0x5f, 0x30, 0x9d, 0x93, 0x38, 0xc5, 0x24, 0xce, 0x67,
	0x39, 0x64, 0x38, 0x62, 0x29, 0x13, 0xb1, 0x40, 0x59, 0x0e, 0x12, 0x6c, 0xe7, 0xc0, 0xa1, 0x92,
	0x43, 0x9a, 0xeb, 0x76, 0x22, 0x88, 0xa0, 0x84, 0xf0, 0xfe, 0xa5, 0xf8, 0xae, 0x4b, 0x41, 0x70,
	0x10, 0x38, 0x24, 0x82, 0xe1, 0xc5, 0x28, 0x64, 0x92, 0x8c, 0x30, 0x85, 0x38, 0xd5, 0xfb, 0xf3,
	0x4a, 0x5f, 0x9a, 0x90, 0x98, 0x6b, 0x6a, 0x50, 0x49, 0x65, 0x24, 0x27, 0x5c, 0x87, 0xf3, 0x3e,
	0x6a, 0x56, 0xeb, 0x5e, 0xc5, 0x7d, 0x92, 0x44, 0x32, 0x7b, 0x61, 0xfd, 0xe7, 0x30, 0x2b, 0x12,
	0x36, 0x25, 0x94, 0x42, 0x91, 0xca, 0x69, 0x48, 0x12, 0x92, 0x52, 0xe6, 0x98, 0x7d, 0x73, 0xd8,
	0x1c, 0x9f, 0x20, 0x15, 0x0f, 0xed, 0xe3, 0x21, 0x1d, 0x0f, 0x4d, 0x20, 0x4e, 0xfd, 0xc1, 0x6a,
	0xd3, 0x33, 0xbe, 0x37, 0xbd, 0xb3, 0x25, 0xe1, 0xc9, 0x8d, 0xf7, 0xb7, 0x8c, 0x17, 0x74, 0xd4,
	0xe2, 0x56, 0xcd, 0x7d, 0x35, 0xb6, 0x1f, 0xac, 0x86, 0x0a, 0xe6, 0xd4, 0x4a, 0x9f, 0x3e, 0xaa,
	0x3a, 0x1b, 0x7a, 0x2c, 0x39, 0xff, 0x9f, 0xb6, 0x6b, 0x2b, 0x3b, 0xf5, 0xdb, 0x0b, 0xb4, 0x8c,
	0x3d, 0xb7, 0xda, 0xe5, 0x3d, 0xa6, 0x39, 0xa3, 0x90, 0xcf, 0x84, 0x53, 0xef, 0xd7, 0x87, 0xcd,
	0xf1, 0xa0, 0x5a, 0x77, 0xb2, 0xc7, 0x83, 0x92, 0xf6, 0x4f, 0xb5, 0x78, 0x47, 0x89, 0x1f, 0x29,
	0x79, 0x41, 0x8b, 0x1e, 0x50, 0xe1, 0xdf, 0xad, 0xb6, 0xae, 0xb9, 0xde, 0xba, 0xe6, 0xd7, 0xd6,
	0x35, 0xdf, 0x77, 0xae, 0xb1, 0xde, 0xb9, 0xc6, 0xe7, 0xce, 0x35, 0x9e, 0x2f, 0xa3, 0x58, 0xce,
	0x8b, 0x10, 0x51, 0xe0, 0x38, 0x29, 0xf8, 0xd5, 0x71, 0x21, 0x6f, 0xbf, 0x95, 0xc8, 0x65, 0xc6,
	0x44, 0xd8, 0x28, 0x2b, 0xb9, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xde, 0xc6, 0x32, 0x26, 0x59,
	0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimRecords) > 0 {
		for iNdEx := len(m.ClaimRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ModuleAccountBalance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ModuleAccountBalance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClaimRecords) > 0 {
		for _, e := range m.ClaimRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccountBalance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ModuleAccountBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimRecords = append(m.ClaimRecords, ClaimRecord{})
			if err := m.ClaimRecords[len(m.ClaimRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
