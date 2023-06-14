// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lumnetwork/chain/airdrop/claim.proto

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

type Action int32

const (
	ACTION_VOTE           Action = 0
	ACTION_DELEGATE_STAKE Action = 1
)

var Action_name = map[int32]string{
	0: "ACTION_VOTE",
	1: "ACTION_DELEGATE_STAKE",
}

var Action_value = map[string]int32{
	"ACTION_VOTE":           0,
	"ACTION_DELEGATE_STAKE": 1,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e293615af4b53f11, []int{0}
}

type ClaimRecord struct {
	Address                string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	InitialClaimableAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_claimable_amount,json=initialClaimableAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_amount" yaml:"initial_claimable_amount"`
	ActionCompleted        []bool                                   `protobuf:"varint,3,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_e293615af4b53f11, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetInitialClaimableAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialClaimableAmount
	}
	return nil
}

func (m *ClaimRecord) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("lumnetwork.chain.airdrop.Action", Action_name, Action_value)
	proto.RegisterType((*ClaimRecord)(nil), "lumnetwork.chain.airdrop.ClaimRecord")
}

func init() {
	proto.RegisterFile("lumnetwork/chain/airdrop/claim.proto", fileDescriptor_e293615af4b53f11)
}

var fileDescriptor_e293615af4b53f11 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6a, 0xd4, 0x40,
	0x00, 0xc6, 0x93, 0xae, 0x54, 0x9d, 0x05, 0xbb, 0x04, 0xff, 0xa4, 0x2b, 0x4c, 0x96, 0xe0, 0x61,
	0x91, 0x76, 0x86, 0xea, 0xad, 0xb7, 0x6c, 0x4c, 0x45, 0x14, 0x0b, 0x69, 0xf0, 0xe0, 0x25, 0x4c,
	0x26, 0xc3, 0x76, 0x68, 0x26, 0x13, 0x32, 0x13, 0xb5, 0x6f, 0xe0, 0xd1, 0x77, 0xf0, 0x22, 0x1e,
	0x7c, 0x8e, 0x1e, 0xf7, 0xe8, 0x29, 0xca, 0xee, 0x1b, 0xec, 0x13, 0xc8, 0x4e, 0xb2, 0xfe, 0x83,
	0x9e, 0x12, 0xbe, 0xf9, 0xf1, 0xe3, 0xfb, 0x98, 0x01, 0x8f, 0x8a, 0x46, 0x94, 0x4c, 0xbf, 0x97,
	0xf5, 0x05, 0xa6, 0xe7, 0x84, 0x97, 0x98, 0xf0, 0x3a, 0xaf, 0x65, 0x85, 0x69, 0x41, 0xb8, 0x40,
	0x55, 0x2d, 0xb5, 0x74, 0xdc, 0x3f, 0x14, 0x32, 0x14, 0xea, 0xa9, 0xf1, 0xdd, 0xb9, 0x9c, 0x4b,
	0x03, 0xe1, 0xcd, 0x5f, 0xc7, 0x8f, 0x21, 0x95, 0x4a, 0x48, 0x85, 0x33, 0xa2, 0x18, 0x7e, 0x77,
	0x94, 0x31, 0x4d, 0x8e, 0x30, 0x95, 0xbc, 0xec, 0xce, 0xfd, 0x6f, 0x3b, 0x60, 0x18, 0x6e, 0xfc,
	0x31, 0xa3, 0xb2, 0xce, 0x9d, 0x03, 0x70, 0x93, 0xe4, 0x79, 0xcd, 0x94, 0x72, 0xed, 0x89, 0x3d,
	0xbd, 0x3d, 0x73, 0xd6, 0xad, 0x77, 0xe7, 0x92, 0x88, 0xe2, 0xd8, 0xef, 0x0f, 0xfc, 0x78, 0x8b,
	0x38, 0x5f, 0x6c, 0xe0, 0xf2, 0x92, 0x6b, 0x4e, 0x8a, 0xd4, 0xb4, 0x24, 0x59, 0xc1, 0x52, 0x22,
	0x64, 0x53, 0x6a, 0x77, 0x67, 0x32, 0x98, 0x0e, 0x9f, 0xec, 0xa3, 0xae, 0x01, 0xda, 0x34, 0x40,
	0x7d, 0x03, 0x14, 0x4a, 0x5e, 0xce, 0xce, 0xae, 0x5a, 0xcf, 0x5a, 0xb7, 0x9e, 0xd7, 0xe9, 0xaf,
	0x13, 0xf9, 0x5f, 0x7f, 0x78, 0xd3, 0x39, 0xd7, 0xe7, 0x4d, 0x86, 0xa8, 0x14, 0xb8, 0x5f, 0xd4,
	0x7d, 0x0e, 0x55, 0x7e, 0x81, 0xf5, 0x65, 0xc5, 0x94, 0x71, 0xaa, 0xf8, 0x7e, 0xaf, 0x09, 0xb7,
	0x96, 0xc0, 0x48, 0x9c, 0x13, 0x30, 0x22, 0x54, 0x73, 0x59, 0xa6, 0x54, 0x8a, 0xaa, 0x60, 0x9a,
	0xe5, 0xee, 0x60, 0x32, 0x98, 0xde, 0x9a, 0x3d, 0x5c, 0xb7, 0xde, 0x83, 0x7e, 0xe1, 0x7f, 0x84,
	0x1f, 0xef, 0x75, 0x51, 0xb8, 0x4d, 0x1e, 0x1f, 0x83, 0xdd, 0xc0, 0x44, 0xce, 0x1e, 0x18, 0x06,
	0x61, 0xf2, 0xe2, 0xf4, 0x75, 0xfa, 0xe6, 0x34, 0x89, 0x46, 0x96, 0xb3, 0x0f, 0xee, 0xf5, 0xc1,
	0xb3, 0xe8, 0x55, 0xf4, 0x3c, 0x48, 0xa2, 0xf4, 0x2c, 0x09, 0x5e, 0x46, 0x23, 0x7b, 0x7c, 0xe3,
	0xe3, 0x67, 0x68, 0xcd, 0x4e, 0xae, 0x96, 0xd0, 0x5e, 0x2c, 0xa1, 0xfd, 0x73, 0x09, 0xed, 0x4f,
	0x2b, 0x68, 0x2d, 0x56, 0xd0, 0xfa, 0xbe, 0x82, 0xd6, 0xdb, 0x83, 0xbf, 0xf6, 0x15, 0x8d, 0x38,
	0xfc, 0xf7, 0x21, 0x7c, 0xf8, 0xfd, 0x14, 0xcc, 0xd2, 0x6c, 0xd7, 0xdc, 0xdd, 0xd3, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x19, 0x39, 0xb1, 0xf5, 0x33, 0x02, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InitialClaimableAmount) > 0 {
		for iNdEx := len(m.InitialClaimableAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.InitialClaimableAmount) > 0 {
		for _, e := range m.InitialClaimableAmount {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovClaim(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableAmount = append(m.InitialClaimableAmount, types.Coin{})
			if err := m.InitialClaimableAmount[len(m.InitialClaimableAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
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
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
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
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
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
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
