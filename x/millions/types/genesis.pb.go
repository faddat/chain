// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lumnetwork/chain/millions/genesis.proto

package types

import (
	fmt "fmt"
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

type GenesisState struct {
	Params           Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	NextPoolId       uint64       `protobuf:"varint,2,opt,name=next_pool_id,json=nextPoolId,proto3" json:"next_pool_id,omitempty"`
	NextDepositId    uint64       `protobuf:"varint,3,opt,name=next_deposit_id,json=nextDepositId,proto3" json:"next_deposit_id,omitempty"`
	NextPrizeId      uint64       `protobuf:"varint,4,opt,name=next_prize_id,json=nextPrizeId,proto3" json:"next_prize_id,omitempty"`
	NextWithdrawalId uint64       `protobuf:"varint,5,opt,name=next_withdrawal_id,json=nextWithdrawalId,proto3" json:"next_withdrawal_id,omitempty"`
	Pools            []Pool       `protobuf:"bytes,6,rep,name=pools,proto3" json:"pools"`
	Deposits         []Deposit    `protobuf:"bytes,7,rep,name=deposits,proto3" json:"deposits"`
	Draws            []Draw       `protobuf:"bytes,8,rep,name=draws,proto3" json:"draws"`
	Prizes           []Prize      `protobuf:"bytes,9,rep,name=prizes,proto3" json:"prizes"`
	Withdrawals      []Withdrawal `protobuf:"bytes,10,rep,name=withdrawals,proto3" json:"withdrawals"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1408a4af22a0ae3, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetNextPoolId() uint64 {
	if m != nil {
		return m.NextPoolId
	}
	return 0
}

func (m *GenesisState) GetNextDepositId() uint64 {
	if m != nil {
		return m.NextDepositId
	}
	return 0
}

func (m *GenesisState) GetNextPrizeId() uint64 {
	if m != nil {
		return m.NextPrizeId
	}
	return 0
}

func (m *GenesisState) GetNextWithdrawalId() uint64 {
	if m != nil {
		return m.NextWithdrawalId
	}
	return 0
}

func (m *GenesisState) GetPools() []Pool {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *GenesisState) GetDeposits() []Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetDraws() []Draw {
	if m != nil {
		return m.Draws
	}
	return nil
}

func (m *GenesisState) GetPrizes() []Prize {
	if m != nil {
		return m.Prizes
	}
	return nil
}

func (m *GenesisState) GetWithdrawals() []Withdrawal {
	if m != nil {
		return m.Withdrawals
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lumnetwork.chain.millions.GenesisState")
}

func init() {
	proto.RegisterFile("lumnetwork/chain/millions/genesis.proto", fileDescriptor_e1408a4af22a0ae3)
}

var fileDescriptor_e1408a4af22a0ae3 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x1c, 0xc5, 0x13, 0x9b, 0xc6, 0x3a, 0xe9, 0xa2, 0x0c, 0x0a, 0xb1, 0x87, 0x6c, 0x0c, 0xb6, 0x2e,
	0x62, 0x13, 0xa8, 0x37, 0x05, 0x0f, 0xcb, 0x42, 0xc9, 0x41, 0x58, 0xd6, 0x83, 0xe0, 0xa5, 0x4c,
	0x9b, 0x21, 0x3b, 0x38, 0xc9, 0x84, 0xcc, 0x2c, 0x69, 0xfd, 0x14, 0x7e, 0xac, 0x82, 0x97, 0x1e,
	0x3d, 0x15, 0xd9, 0xfd, 0x06, 0x7e, 0x02, 0xf9, 0x4f, 0x66, 0x37, 0xf6, 0xb0, 0xd9, 0x5b, 0xf8,
	0xe7, 0xf7, 0x1e, 0xef, 0x3d, 0x06, 0xbd, 0xe1, 0x8b, 0xa2, 0xa4, 0xaa, 0x11, 0xf5, 0xf7, 0xe4,
	0x6a, 0x4e, 0x58, 0x99, 0x14, 0x8c, 0x73, 0x26, 0x4a, 0x99, 0xe4, 0xb4, 0xa4, 0x92, 0xc9, 0xb8,
	0xaa, 0x85, 0x12, 0xf8, 0x65, 0x07, 0xc6, 0x1a, 0x8c, 0xd7, 0xe0, 0xd1, 0xf3, 0x5c, 0xe4, 0x42,
	0x53, 0x09, 0x7c, 0xb5, 0x82, 0xa3, 0x1e, 0xe7, 0x8c, 0x56, 0x42, 0x32, 0x65, 0xc0, 0xd7, 0x3d,
	0x60, 0x4d, 0x1a, 0x43, 0x9d, 0x6c, 0xa7, 0x2a, 0x52, 0x93, 0x42, 0xee, 0x76, 0xab, 0x84, 0xe0,
	0x86, 0x3a, 0xee, 0xa1, 0x6a, 0xf6, 0x83, 0x1a, 0xec, 0xed, 0x76, 0xac, 0x61, 0x6a, 0x0e, 0xf1,
	0x88, 0xb1, 0x8c, 0x7e, 0x39, 0xe8, 0xf0, 0xbc, 0x9d, 0xec, 0x8b, 0x22, 0x8a, 0xe2, 0x29, 0x72,
	0xdb, 0x64, 0xbe, 0x1d, 0xda, 0x23, 0xef, 0xec, 0x55, 0xbc, 0x75, 0xc2, 0x78, 0xaa, 0xc1, 0xf1,
	0x8b, 0xdb, 0xfb, 0xa1, 0xf5, 0xf7, 0x7e, 0x38, 0xb8, 0x21, 0x05, 0xff, 0x10, 0xb5, 0xf2, 0x68,
	0x66, 0x7c, 0x70, 0x88, 0x0e, 0x4b, 0x7a, 0xad, 0x2e, 0xa0, 0xc8, 0x05, 0xcb, 0xfc, 0x47, 0xa1,
	0x3d, 0x72, 0x66, 0x08, 0x6e, 0x53, 0x21, 0x78, 0x9a, 0xe1, 0x13, 0xf4, 0x54, 0x13, 0x66, 0x61,
	0x80, 0xf6, 0x34, 0x34, 0x80, 0xf3, 0xa4, 0xbd, 0xa6, 0x19, 0x8e, 0xd0, 0xa0, 0x75, 0x82, 0xb2,
	0x40, 0x39, 0x9a, 0xf2, 0xb4, 0x15, 0xdc, 0xd2, 0x0c, 0xbf, 0x43, 0x58, 0x33, 0x5d, 0x53, 0x00,
	0xf7, 0x35, 0xf8, 0x0c, 0xfe, 0x7c, 0xdd, 0xfc, 0x48, 0x33, 0xfc, 0x11, 0xed, 0x43, 0x2c, 0xe9,
	0xbb, 0xe1, 0xde, 0xc8, 0x3b, 0x1b, 0xf6, 0x95, 0x15, 0x82, 0x8f, 0x1d, 0xa8, 0x3a, 0x6b, 0x35,
	0x78, 0x82, 0x0e, 0x4c, 0x62, 0xe9, 0x3f, 0xd6, 0xfa, 0xa8, 0x47, 0x6f, 0x6a, 0x18, 0x8b, 0x8d,
	0x12, 0x22, 0x40, 0x1c, 0xe9, 0x1f, 0xec, 0x8c, 0x30, 0xa9, 0x49, 0xb3, 0x8e, 0xa0, 0x35, 0xf8,
	0x13, 0x72, 0xf5, 0x18, 0xd2, 0x7f, 0xa2, 0xd5, 0x61, 0x5f, 0x01, 0x00, 0x8d, 0xdc, 0xa8, 0xf0,
	0x67, 0xe4, 0x75, 0x43, 0x49, 0x1f, 0x69, 0x93, 0xe3, 0x1e, 0x93, 0x6e, 0x3d, 0xe3, 0xf4, 0xbf,
	0x7e, 0x7c, 0x7e, 0xbb, 0x0c, 0xec, 0xbb, 0x65, 0x60, 0xff, 0x59, 0x06, 0xf6, 0xcf, 0x55, 0x60,
	0xdd, 0xad, 0x02, 0xeb, 0xf7, 0x2a, 0xb0, 0xbe, 0x9d, 0xe6, 0x4c, 0xcd, 0x17, 0x97, 0xf1, 0x95,
	0x28, 0x12, 0xbe, 0x28, 0x4e, 0x1f, 0xbe, 0xcf, 0xeb, 0xee, 0x85, 0xaa, 0x9b, 0x8a, 0xca, 0x4b,
	0x57, 0xbf, 0xce, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x18, 0x61, 0xe4, 0x38, 0xe9, 0x03,
	0x00, 0x00,
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
	if len(m.Withdrawals) > 0 {
		for iNdEx := len(m.Withdrawals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Withdrawals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Prizes) > 0 {
		for iNdEx := len(m.Prizes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prizes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Draws) > 0 {
		for iNdEx := len(m.Draws) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Draws[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.NextWithdrawalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextWithdrawalId))
		i--
		dAtA[i] = 0x28
	}
	if m.NextPrizeId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextPrizeId))
		i--
		dAtA[i] = 0x20
	}
	if m.NextDepositId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextDepositId))
		i--
		dAtA[i] = 0x18
	}
	if m.NextPoolId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextPoolId))
		i--
		dAtA[i] = 0x10
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.NextPoolId != 0 {
		n += 1 + sovGenesis(uint64(m.NextPoolId))
	}
	if m.NextDepositId != 0 {
		n += 1 + sovGenesis(uint64(m.NextDepositId))
	}
	if m.NextPrizeId != 0 {
		n += 1 + sovGenesis(uint64(m.NextPrizeId))
	}
	if m.NextWithdrawalId != 0 {
		n += 1 + sovGenesis(uint64(m.NextWithdrawalId))
	}
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Draws) > 0 {
		for _, e := range m.Draws {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Prizes) > 0 {
		for _, e := range m.Prizes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Withdrawals) > 0 {
		for _, e := range m.Withdrawals {
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPoolId", wireType)
			}
			m.NextPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextDepositId", wireType)
			}
			m.NextDepositId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextDepositId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPrizeId", wireType)
			}
			m.NextPrizeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextPrizeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextWithdrawalId", wireType)
			}
			m.NextWithdrawalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextWithdrawalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
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
			m.Pools = append(m.Pools, Pool{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
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
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Draws", wireType)
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
			m.Draws = append(m.Draws, Draw{})
			if err := m.Draws[len(m.Draws)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prizes", wireType)
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
			m.Prizes = append(m.Prizes, Prize{})
			if err := m.Prizes[len(m.Prizes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawals", wireType)
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
			m.Withdrawals = append(m.Withdrawals, Withdrawal{})
			if err := m.Withdrawals[len(m.Withdrawals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
