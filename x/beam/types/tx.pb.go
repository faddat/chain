// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lumnetwork/chain/beam/tx.proto

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

type MsgOpenBeam struct {
	Id                  string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorAddress      string      `protobuf:"bytes,2,opt,name=creator_address,json=creatorAddress,proto3" json:"creator_address,omitempty"`
	Secret              string      `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	Amount              *types.Coin `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Schema              string      `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
	Data                *BeamData   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	ClaimAddress        string      `protobuf:"bytes,7,opt,name=claim_address,json=claimAddress,proto3" json:"claim_address,omitempty"`
	ClaimExpiresAtBlock int32       `protobuf:"varint,8,opt,name=claim_expires_at_block,json=claimExpiresAtBlock,proto3" json:"claim_expires_at_block,omitempty"`
	ClosesAtBlock       int32       `protobuf:"varint,9,opt,name=closes_at_block,json=closesAtBlock,proto3" json:"closes_at_block,omitempty"`
}

func (m *MsgOpenBeam) Reset()         { *m = MsgOpenBeam{} }
func (m *MsgOpenBeam) String() string { return proto.CompactTextString(m) }
func (*MsgOpenBeam) ProtoMessage()    {}
func (*MsgOpenBeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d811cb35a71f0c7, []int{0}
}
func (m *MsgOpenBeam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOpenBeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOpenBeam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOpenBeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOpenBeam.Merge(m, src)
}
func (m *MsgOpenBeam) XXX_Size() int {
	return m.Size()
}
func (m *MsgOpenBeam) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOpenBeam.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOpenBeam proto.InternalMessageInfo

func (m *MsgOpenBeam) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgOpenBeam) GetCreatorAddress() string {
	if m != nil {
		return m.CreatorAddress
	}
	return ""
}

func (m *MsgOpenBeam) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *MsgOpenBeam) GetAmount() *types.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgOpenBeam) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *MsgOpenBeam) GetData() *BeamData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MsgOpenBeam) GetClaimAddress() string {
	if m != nil {
		return m.ClaimAddress
	}
	return ""
}

func (m *MsgOpenBeam) GetClaimExpiresAtBlock() int32 {
	if m != nil {
		return m.ClaimExpiresAtBlock
	}
	return 0
}

func (m *MsgOpenBeam) GetClosesAtBlock() int32 {
	if m != nil {
		return m.ClosesAtBlock
	}
	return 0
}

type MsgUpdateBeam struct {
	Id                  string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdaterAddress      string      `protobuf:"bytes,2,opt,name=updater_address,json=updaterAddress,proto3" json:"updater_address,omitempty"`
	Amount              *types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Status              BeamState   `protobuf:"varint,4,opt,name=status,proto3,enum=lumnetwork.chain.beam.BeamState" json:"status,omitempty"`
	CancelReason        string      `protobuf:"bytes,5,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
	HideContent         bool        `protobuf:"varint,6,opt,name=hide_content,json=hideContent,proto3" json:"hide_content,omitempty"`
	Data                *BeamData   `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	ClaimAddress        string      `protobuf:"bytes,8,opt,name=claim_address,json=claimAddress,proto3" json:"claim_address,omitempty"`
	ClaimExpiresAtBlock int32       `protobuf:"varint,9,opt,name=claim_expires_at_block,json=claimExpiresAtBlock,proto3" json:"claim_expires_at_block,omitempty"`
	ClosesAtBlock       int32       `protobuf:"varint,10,opt,name=closes_at_block,json=closesAtBlock,proto3" json:"closes_at_block,omitempty"`
}

func (m *MsgUpdateBeam) Reset()         { *m = MsgUpdateBeam{} }
func (m *MsgUpdateBeam) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateBeam) ProtoMessage()    {}
func (*MsgUpdateBeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d811cb35a71f0c7, []int{1}
}
func (m *MsgUpdateBeam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateBeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateBeam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateBeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateBeam.Merge(m, src)
}
func (m *MsgUpdateBeam) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateBeam) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateBeam.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateBeam proto.InternalMessageInfo

func (m *MsgUpdateBeam) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgUpdateBeam) GetUpdaterAddress() string {
	if m != nil {
		return m.UpdaterAddress
	}
	return ""
}

func (m *MsgUpdateBeam) GetAmount() *types.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgUpdateBeam) GetStatus() BeamState {
	if m != nil {
		return m.Status
	}
	return BeamState_StateUnspecified
}

func (m *MsgUpdateBeam) GetCancelReason() string {
	if m != nil {
		return m.CancelReason
	}
	return ""
}

func (m *MsgUpdateBeam) GetHideContent() bool {
	if m != nil {
		return m.HideContent
	}
	return false
}

func (m *MsgUpdateBeam) GetData() *BeamData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MsgUpdateBeam) GetClaimAddress() string {
	if m != nil {
		return m.ClaimAddress
	}
	return ""
}

func (m *MsgUpdateBeam) GetClaimExpiresAtBlock() int32 {
	if m != nil {
		return m.ClaimExpiresAtBlock
	}
	return 0
}

func (m *MsgUpdateBeam) GetClosesAtBlock() int32 {
	if m != nil {
		return m.ClosesAtBlock
	}
	return 0
}

type MsgClaimBeam struct {
	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClaimerAddress string `protobuf:"bytes,2,opt,name=claimer_address,json=claimerAddress,proto3" json:"claimer_address,omitempty"`
	Secret         string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *MsgClaimBeam) Reset()         { *m = MsgClaimBeam{} }
func (m *MsgClaimBeam) String() string { return proto.CompactTextString(m) }
func (*MsgClaimBeam) ProtoMessage()    {}
func (*MsgClaimBeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d811cb35a71f0c7, []int{2}
}
func (m *MsgClaimBeam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimBeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimBeam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimBeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimBeam.Merge(m, src)
}
func (m *MsgClaimBeam) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimBeam) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimBeam.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimBeam proto.InternalMessageInfo

func (m *MsgClaimBeam) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgClaimBeam) GetClaimerAddress() string {
	if m != nil {
		return m.ClaimerAddress
	}
	return ""
}

func (m *MsgClaimBeam) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgOpenBeam)(nil), "lumnetwork.chain.beam.MsgOpenBeam")
	proto.RegisterType((*MsgUpdateBeam)(nil), "lumnetwork.chain.beam.MsgUpdateBeam")
	proto.RegisterType((*MsgClaimBeam)(nil), "lumnetwork.chain.beam.MsgClaimBeam")
}

func init() { proto.RegisterFile("lumnetwork/chain/beam/tx.proto", fileDescriptor_7d811cb35a71f0c7) }

var fileDescriptor_7d811cb35a71f0c7 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0x8c, 0x93, 0x34, 0x4d, 0x36, 0x1f, 0x95, 0x0c, 0x54, 0xa6, 0x07, 0x37, 0x14, 0x09, 0xc2,
	0x81, 0xb5, 0xda, 0x1e, 0x80, 0x63, 0x13, 0x38, 0x46, 0x48, 0x46, 0x5c, 0xb8, 0x58, 0xcf, 0xeb,
	0x27, 0xc7, 0xaa, 0xbd, 0x6b, 0x79, 0xd7, 0x10, 0xfe, 0x05, 0x3f, 0x89, 0x63, 0x8f, 0x3d, 0x72,
	0x42, 0x28, 0xb9, 0xf2, 0x23, 0x90, 0xd7, 0xdb, 0x16, 0xd4, 0x14, 0x14, 0x89, 0x4b, 0x94, 0x9d,
	0x19, 0x8f, 0xdf, 0xce, 0xe8, 0x99, 0xb8, 0x69, 0x99, 0x71, 0x54, 0x9f, 0x44, 0x71, 0xee, 0xb1,
	0x05, 0x24, 0xdc, 0x0b, 0x11, 0x32, 0x4f, 0x2d, 0x69, 0x5e, 0x08, 0x25, 0xec, 0x07, 0x37, 0x3c,
	0xd5, 0x3c, 0xad, 0xf8, 0x83, 0xfb, 0xb1, 0x88, 0x85, 0x56, 0x78, 0xd5, 0xbf, 0x5a, 0x7c, 0xe0,
	0x32, 0x21, 0x33, 0x21, 0xbd, 0x10, 0x24, 0x7a, 0x1f, 0x8f, 0x43, 0x54, 0x70, 0xec, 0x31, 0x91,
	0x70, 0xc3, 0x8f, 0x37, 0xbf, 0xac, 0xfa, 0xa9, 0x15, 0x47, 0x3f, 0x9b, 0xa4, 0x3f, 0x97, 0xf1,
	0xdb, 0x1c, 0xf9, 0x14, 0x21, 0xb3, 0x47, 0xa4, 0x99, 0x44, 0x8e, 0x35, 0xb6, 0x26, 0x3d, 0xbf,
	0x99, 0x44, 0xf6, 0x53, 0xb2, 0xc7, 0x0a, 0x04, 0x25, 0x8a, 0x00, 0xa2, 0xa8, 0x40, 0x29, 0x9d,
	0xa6, 0x26, 0x47, 0x06, 0x3e, 0xab, 0x51, 0x7b, 0x9f, 0x74, 0x24, 0xb2, 0x02, 0x95, 0xd3, 0xd2,
	0xbc, 0x39, 0xd9, 0x2f, 0x48, 0x07, 0x32, 0x51, 0x72, 0xe5, 0xb4, 0xc7, 0xd6, 0xa4, 0x7f, 0xf2,
	0x90, 0xd6, 0x33, 0xd3, 0x6a, 0x66, 0x6a, 0x66, 0xa6, 0x33, 0x91, 0xf0, 0x69, 0xfb, 0xe2, 0xfb,
	0xa1, 0xe5, 0x1b, 0xb9, 0x36, 0x64, 0x0b, 0xcc, 0xc0, 0xd9, 0x31, 0x86, 0xfa, 0x64, 0xbf, 0x22,
	0xed, 0x08, 0x14, 0x38, 0x1d, 0x6d, 0x77, 0x48, 0x37, 0xe6, 0x45, 0xab, 0xcb, 0xbc, 0x06, 0x05,
	0xc6, 0x54, 0x3f, 0x62, 0x3f, 0x26, 0x43, 0x96, 0x42, 0x92, 0x5d, 0x5f, 0x65, 0x57, 0x3b, 0x0f,
	0x34, 0x78, 0x75, 0x91, 0x53, 0xb2, 0x5f, 0x8b, 0x70, 0x99, 0x27, 0x05, 0xca, 0x00, 0x54, 0x10,
	0xa6, 0x82, 0x9d, 0x3b, 0xdd, 0xb1, 0x35, 0xd9, 0xf1, 0xef, 0x69, 0xf6, 0x4d, 0x4d, 0x9e, 0xa9,
	0x69, 0x45, 0xd9, 0x4f, 0xc8, 0x1e, 0x4b, 0x85, 0xfc, 0x5d, 0xdd, 0xd3, 0xea, 0x61, 0x0d, 0x1b,
	0xdd, 0xd1, 0xd7, 0x16, 0x19, 0xce, 0x65, 0xfc, 0x3e, 0x8f, 0x40, 0xe1, 0x5d, 0x81, 0x97, 0x9a,
	0xbd, 0x15, 0xb8, 0x81, 0xaf, 0xe6, 0xbc, 0x09, 0xb6, 0xb5, 0x5d, 0xb0, 0x2f, 0x49, 0x47, 0x2a,
	0x50, 0xa5, 0xd4, 0x8d, 0x8c, 0x4e, 0xc6, 0x7f, 0x89, 0xf0, 0x9d, 0x02, 0x85, 0xbe, 0xd1, 0xeb,
	0xfc, 0x80, 0x33, 0x4c, 0x83, 0x02, 0x41, 0x0a, 0x6e, 0x9a, 0x19, 0xd4, 0xa0, 0xaf, 0x31, 0xfb,
	0x11, 0x19, 0x2c, 0x92, 0x08, 0x03, 0x26, 0xb8, 0x42, 0xae, 0x74, 0x4f, 0x5d, 0xbf, 0x5f, 0x61,
	0xb3, 0x1a, 0xba, 0xae, 0x70, 0xf7, 0x3f, 0x54, 0xd8, 0xdd, 0xaa, 0xc2, 0xde, 0x56, 0x15, 0x92,
	0x4d, 0x15, 0x06, 0x64, 0x30, 0x97, 0xf1, 0xac, 0x72, 0xb8, 0x73, 0x63, 0x2a, 0xf2, 0x76, 0x81,
	0x06, 0xfe, 0xc7, 0xc6, 0x4c, 0x67, 0x17, 0x2b, 0xd7, 0xba, 0x5c, 0xb9, 0xd6, 0x8f, 0x95, 0x6b,
	0x7d, 0x59, 0xbb, 0x8d, 0xcb, 0xb5, 0xdb, 0xf8, 0xb6, 0x76, 0x1b, 0x1f, 0x9e, 0xc5, 0x89, 0x5a,
	0x94, 0x21, 0x65, 0x22, 0xf3, 0xd2, 0x32, 0x7b, 0xfe, 0xe7, 0x6a, 0x2f, 0xcd, 0x97, 0xe4, 0x73,
	0x8e, 0x32, 0xec, 0xe8, 0xf5, 0x3e, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xfb, 0x0a, 0xfc,
	0x6f, 0x04, 0x00, 0x00,
}

func (m *MsgOpenBeam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOpenBeam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOpenBeam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClosesAtBlock != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ClosesAtBlock))
		i--
		dAtA[i] = 0x48
	}
	if m.ClaimExpiresAtBlock != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ClaimExpiresAtBlock))
		i--
		dAtA[i] = 0x40
	}
	if len(m.ClaimAddress) > 0 {
		i -= len(m.ClaimAddress)
		copy(dAtA[i:], m.ClaimAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClaimAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Schema) > 0 {
		i -= len(m.Schema)
		copy(dAtA[i:], m.Schema)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Schema)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Amount != nil {
		{
			size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CreatorAddress) > 0 {
		i -= len(m.CreatorAddress)
		copy(dAtA[i:], m.CreatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CreatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateBeam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateBeam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateBeam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClosesAtBlock != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ClosesAtBlock))
		i--
		dAtA[i] = 0x50
	}
	if m.ClaimExpiresAtBlock != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ClaimExpiresAtBlock))
		i--
		dAtA[i] = 0x48
	}
	if len(m.ClaimAddress) > 0 {
		i -= len(m.ClaimAddress)
		copy(dAtA[i:], m.ClaimAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClaimAddress)))
		i--
		dAtA[i] = 0x42
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.HideContent {
		i--
		if m.HideContent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.CancelReason) > 0 {
		i -= len(m.CancelReason)
		copy(dAtA[i:], m.CancelReason)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CancelReason)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Status != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Amount != nil {
		{
			size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UpdaterAddress) > 0 {
		i -= len(m.UpdaterAddress)
		copy(dAtA[i:], m.UpdaterAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.UpdaterAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimBeam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimBeam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimBeam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClaimerAddress) > 0 {
		i -= len(m.ClaimerAddress)
		copy(dAtA[i:], m.ClaimerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClaimerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgOpenBeam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CreatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ClaimAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ClaimExpiresAtBlock != 0 {
		n += 1 + sovTx(uint64(m.ClaimExpiresAtBlock))
	}
	if m.ClosesAtBlock != 0 {
		n += 1 + sovTx(uint64(m.ClosesAtBlock))
	}
	return n
}

func (m *MsgUpdateBeam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.UpdaterAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTx(uint64(m.Status))
	}
	l = len(m.CancelReason)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.HideContent {
		n += 2
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ClaimAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ClaimExpiresAtBlock != 0 {
		n += 1 + sovTx(uint64(m.ClaimExpiresAtBlock))
	}
	if m.ClosesAtBlock != 0 {
		n += 1 + sovTx(uint64(m.ClosesAtBlock))
	}
	return n
}

func (m *MsgClaimBeam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ClaimerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgOpenBeam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgOpenBeam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOpenBeam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Amount == nil {
				m.Amount = &types.Coin{}
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &BeamData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimExpiresAtBlock", wireType)
			}
			m.ClaimExpiresAtBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimExpiresAtBlock |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosesAtBlock", wireType)
			}
			m.ClosesAtBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClosesAtBlock |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateBeam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateBeam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateBeam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdaterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdaterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Amount == nil {
				m.Amount = &types.Coin{}
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= BeamState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CancelReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CancelReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HideContent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.HideContent = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &BeamData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimExpiresAtBlock", wireType)
			}
			m.ClaimExpiresAtBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimExpiresAtBlock |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosesAtBlock", wireType)
			}
			m.ClosesAtBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClosesAtBlock |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimBeam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimBeam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimBeam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
