// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lumnetwork/millions/draw.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type DrawState int32

const (
	DrawState_Unspecified        DrawState = 0
	DrawState_IcaWithdrawRewards DrawState = 1
	DrawState_IcqBalance         DrawState = 2
	DrawState_IbcTransfer        DrawState = 3
	DrawState_Drawing            DrawState = 4
	DrawState_Success            DrawState = 5
	DrawState_Failure            DrawState = 6
)

var DrawState_name = map[int32]string{
	0: "DRAW_STATE_UNSPECIFIED",
	1: "DRAW_STATE_ICA_WITHDRAWREWARDS",
	2: "DRAW_STATE_QUERY_BALANCE",
	3: "DRAW_STATE_IBC_TRANSFER",
	4: "DRAW_STATE_DRAWING",
	5: "DRAW_STATE_SUCCESS",
	6: "DRAW_STATE_FAILURE",
}

var DrawState_value = map[string]int32{
	"DRAW_STATE_UNSPECIFIED":         0,
	"DRAW_STATE_ICA_WITHDRAWREWARDS": 1,
	"DRAW_STATE_QUERY_BALANCE":       2,
	"DRAW_STATE_IBC_TRANSFER":        3,
	"DRAW_STATE_DRAWING":             4,
	"DRAW_STATE_SUCCESS":             5,
	"DRAW_STATE_FAILURE":             6,
}

func (x DrawState) String() string {
	return proto.EnumName(DrawState_name, int32(x))
}

func (DrawState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb40462a68b0e8bb, []int{0}
}

type Draw struct {
	// Draw IDs
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	DrawId uint64 `protobuf:"varint,2,opt,name=draw_id,json=drawId,proto3" json:"draw_id,omitempty"`
	// Draw states
	// error_state is only set in case of failure
	State      DrawState `protobuf:"varint,3,opt,name=state,proto3,enum=lumnetwork.millions.DrawState" json:"state,omitempty"`
	ErrorState DrawState `protobuf:"varint,4,opt,name=error_state,json=errorState,proto3,enum=lumnetwork.millions.DrawState" json:"error_state,omitempty"`
	// Draw state done data
	RandSeed               int64                                  `protobuf:"varint,5,opt,name=rand_seed,json=randSeed,proto3" json:"rand_seed,omitempty"`
	PrizePool              types.Coin                             `protobuf:"bytes,6,opt,name=prize_pool,json=prizePool,proto3" json:"prize_pool"`
	PrizePoolFreshAmount   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=prize_pool_fresh_amount,json=prizePoolFreshAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"prize_pool_fresh_amount"`
	PrizePoolRemainsAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=prize_pool_remains_amount,json=prizePoolRemainsAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"prize_pool_remains_amount"`
	PrizesRefs             []PrizeRef                             `protobuf:"bytes,11,rep,name=prizes_refs,json=prizesRefs,proto3" json:"prizes_refs"`
	TotalWinCount          uint64                                 `protobuf:"varint,12,opt,name=total_win_count,json=totalWinCount,proto3" json:"total_win_count,omitempty"`
	TotalWinAmount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,13,opt,name=total_win_amount,json=totalWinAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_win_amount"`
	// Draw creation and updates
	CreatedAtHeight int64     `protobuf:"varint,15,opt,name=created_at_height,json=createdAtHeight,proto3" json:"created_at_height,omitempty"`
	UpdatedAtHeight int64     `protobuf:"varint,16,opt,name=updated_at_height,json=updatedAtHeight,proto3" json:"updated_at_height,omitempty"`
	CreatedAt       time.Time `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	UpdatedAt       time.Time `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at"`
}

func (m *Draw) Reset()         { *m = Draw{} }
func (m *Draw) String() string { return proto.CompactTextString(m) }
func (*Draw) ProtoMessage()    {}
func (*Draw) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb40462a68b0e8bb, []int{0}
}
func (m *Draw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Draw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Draw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Draw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Draw.Merge(m, src)
}
func (m *Draw) XXX_Size() int {
	return m.Size()
}
func (m *Draw) XXX_DiscardUnknown() {
	xxx_messageInfo_Draw.DiscardUnknown(m)
}

var xxx_messageInfo_Draw proto.InternalMessageInfo

func (m *Draw) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *Draw) GetDrawId() uint64 {
	if m != nil {
		return m.DrawId
	}
	return 0
}

func (m *Draw) GetState() DrawState {
	if m != nil {
		return m.State
	}
	return DrawState_Unspecified
}

func (m *Draw) GetErrorState() DrawState {
	if m != nil {
		return m.ErrorState
	}
	return DrawState_Unspecified
}

func (m *Draw) GetRandSeed() int64 {
	if m != nil {
		return m.RandSeed
	}
	return 0
}

func (m *Draw) GetPrizePool() types.Coin {
	if m != nil {
		return m.PrizePool
	}
	return types.Coin{}
}

func (m *Draw) GetPrizesRefs() []PrizeRef {
	if m != nil {
		return m.PrizesRefs
	}
	return nil
}

func (m *Draw) GetTotalWinCount() uint64 {
	if m != nil {
		return m.TotalWinCount
	}
	return 0
}

func (m *Draw) GetCreatedAtHeight() int64 {
	if m != nil {
		return m.CreatedAtHeight
	}
	return 0
}

func (m *Draw) GetUpdatedAtHeight() int64 {
	if m != nil {
		return m.UpdatedAtHeight
	}
	return 0
}

func (m *Draw) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Draw) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("lumnetwork.millions.DrawState", DrawState_name, DrawState_value)
	proto.RegisterType((*Draw)(nil), "lumnetwork.millions.Draw")
}

func init() { proto.RegisterFile("lumnetwork/millions/draw.proto", fileDescriptor_bb40462a68b0e8bb) }

var fileDescriptor_bb40462a68b0e8bb = []byte{
	// 829 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x6d, 0xda, 0x96, 0x56, 0x8d, 0xc5, 0xb0, 0x81, 0xcd, 0xa8, 0x28, 0x4d, 0x34, 0x40,
	0x20, 0xa4, 0x31, 0x89, 0xb8, 0x3d, 0xf5, 0xd0, 0x82, 0xa2, 0xa8, 0x84, 0x46, 0x60, 0xb8, 0x4b,
	0x09, 0x6a, 0x7b, 0x21, 0x56, 0xe4, 0x4a, 0x5a, 0x84, 0xe4, 0xaa, 0xbb, 0xab, 0xaa, 0xed, 0x13,
	0x14, 0x3e, 0xe5, 0x5c, 0xc0, 0xa7, 0x3e, 0x4b, 0x81, 0x1c, 0x73, 0x2c, 0x7a, 0x48, 0x0b, 0xfb,
	0x45, 0x8a, 0x25, 0xa9, 0x1f, 0x18, 0x3e, 0x34, 0x39, 0x71, 0x77, 0xe6, 0x9b, 0xef, 0x9b, 0x99,
	0x9d, 0x21, 0x30, 0xd3, 0x79, 0x96, 0x63, 0xb1, 0xa0, 0xec, 0x95, 0x93, 0x91, 0x34, 0x25, 0x34,
	0xe7, 0x4e, 0xc2, 0xd0, 0xc2, 0x9e, 0x31, 0x2a, 0xa8, 0xfe, 0xf1, 0xda, 0x6f, 0x2f, 0xfd, 0x2d,
	0x33, 0xa6, 0x3c, 0xa3, 0xdc, 0x19, 0x21, 0x8e, 0x9d, 0x9f, 0x9e, 0x8d, 0xb0, 0x40, 0xcf, 0x9c,
	0x98, 0x92, 0xbc, 0x0c, 0x6a, 0x3d, 0x98, 0xd0, 0x09, 0x2d, 0x8e, 0x8e, 0x3c, 0x55, 0xd6, 0xe3,
	0x09, 0xa5, 0x93, 0x14, 0x3b, 0xc5, 0x6d, 0x34, 0x1f, 0x3b, 0x82, 0x64, 0x98, 0x0b, 0x94, 0xcd,
	0x2a, 0xc0, 0xa3, 0xbb, 0x72, 0x99, 0x31, 0xf2, 0x2b, 0x8e, 0x18, 0x1e, 0x97, 0xa0, 0xcf, 0x7e,
	0xdf, 0x07, 0x6a, 0x97, 0xa1, 0x85, 0x7e, 0x04, 0xf6, 0x67, 0x94, 0xa6, 0x11, 0x49, 0x0c, 0xc5,
	0x52, 0xda, 0x2a, 0xdc, 0x93, 0xd7, 0x20, 0x91, 0x0e, 0x59, 0x80, 0x74, 0x6c, 0x97, 0x0e, 0x79,
	0x0d, 0x12, 0xfd, 0x4b, 0xb0, 0xcb, 0x05, 0x12, 0xd8, 0xd8, 0xb1, 0x94, 0xf6, 0xc1, 0xa9, 0x69,
	0xdf, 0x51, 0x9b, 0x2d, 0xb9, 0x43, 0x89, 0x82, 0x25, 0x58, 0xff, 0x06, 0x34, 0x30, 0x63, 0x94,
	0x45, 0x65, 0xac, 0xfa, 0xbf, 0x62, 0x41, 0x11, 0x52, 0x9c, 0xf5, 0x4f, 0x40, 0x9d, 0xa1, 0x3c,
	0x89, 0x38, 0xc6, 0x89, 0xb1, 0x6b, 0x29, 0xed, 0x1d, 0x58, 0x93, 0x86, 0x10, 0xe3, 0x44, 0xff,
	0x1a, 0x80, 0xb2, 0x42, 0x99, 0xbc, 0xb1, 0x67, 0x29, 0xed, 0xc6, 0xe9, 0x43, 0xbb, 0xec, 0xaf,
	0x2d, 0xfb, 0x6b, 0x57, 0xfd, 0xb5, 0x3d, 0x4a, 0xf2, 0x8e, 0xfa, 0xe6, 0xdd, 0xf1, 0x16, 0xac,
	0x17, 0x21, 0x17, 0x94, 0xa6, 0x3a, 0x06, 0x47, 0xeb, 0xf8, 0x68, 0xcc, 0x30, 0x9f, 0x46, 0x28,
	0xa3, 0xf3, 0x5c, 0x18, 0xfb, 0x96, 0xd2, 0xae, 0x77, 0x6c, 0x19, 0xf1, 0xf7, 0xbb, 0xe3, 0xc7,
	0x13, 0x22, 0xa6, 0xf3, 0x91, 0x1d, 0xd3, 0xcc, 0xa9, 0x9e, 0xaf, 0xfc, 0x9c, 0xf0, 0xe4, 0x95,
	0x23, 0x7e, 0x99, 0x61, 0x6e, 0x07, 0xb9, 0x80, 0x0f, 0x56, 0xdc, 0x3d, 0x49, 0xe6, 0x16, 0x5c,
	0x3a, 0x01, 0x0f, 0x37, 0x64, 0x18, 0xce, 0x10, 0xc9, 0xf9, 0x52, 0xa8, 0xf6, 0x41, 0x42, 0x87,
	0x2b, 0x21, 0x58, 0xd2, 0x55, 0x52, 0x5d, 0xd0, 0x28, 0x3c, 0x5c, 0x3e, 0x3a, 0x37, 0x1a, 0xd6,
	0x4e, 0xbb, 0x71, 0xfa, 0xe9, 0x9d, 0xfd, 0xbe, 0x90, 0x38, 0x88, 0xc7, 0x55, 0x5b, 0xca, 0x4e,
	0x72, 0x88, 0xc7, 0x5c, 0x7f, 0x0c, 0x9a, 0x82, 0x0a, 0x94, 0x46, 0x0b, 0x92, 0x47, 0x71, 0x91,
	0xe6, 0x47, 0xc5, 0x30, 0xdc, 0x2b, 0xcc, 0x43, 0x92, 0x7b, 0x85, 0xda, 0x77, 0x40, 0x5b, 0xe3,
	0xaa, 0x7a, 0xee, 0x7d, 0x50, 0x3d, 0x07, 0x4b, 0xe2, 0xaa, 0x8e, 0x27, 0xe0, 0x7e, 0xcc, 0x30,
	0x12, 0x38, 0x89, 0x90, 0x88, 0xa6, 0x98, 0x4c, 0xa6, 0xc2, 0x68, 0x16, 0xcf, 0xdf, 0xac, 0x1c,
	0xae, 0x78, 0x51, 0x98, 0x25, 0x76, 0x3e, 0x4b, 0x6e, 0x61, 0xb5, 0x12, 0x5b, 0x39, 0x56, 0x58,
	0x0f, 0x80, 0x35, 0xaf, 0x71, 0xbf, 0x98, 0x98, 0x96, 0x5d, 0xee, 0x96, 0xbd, 0xdc, 0x2d, 0xbb,
	0xbf, 0xdc, 0xad, 0x4e, 0x4d, 0xd6, 0xf1, 0xfa, 0x9f, 0x63, 0x05, 0xd6, 0x57, 0xb2, 0x92, 0x64,
	0x2d, 0x68, 0xe8, 0xef, 0x43, 0xb2, 0xca, 0xe7, 0x4c, 0xad, 0xd5, 0x35, 0x70, 0xa6, 0xd6, 0x80,
	0xd6, 0x38, 0x53, 0x6b, 0x07, 0x5a, 0xf3, 0xc9, 0x9f, 0xdb, 0xa0, 0xbe, 0x5a, 0x02, 0xfd, 0x73,
	0x70, 0xd8, 0x85, 0xee, 0x30, 0x0a, 0xfb, 0x6e, 0xdf, 0x8f, 0x06, 0xe7, 0xe1, 0x85, 0xef, 0x05,
	0xbd, 0xc0, 0xef, 0x6a, 0x5b, 0xad, 0xe6, 0xe5, 0x95, 0xd5, 0x18, 0xe4, 0x7c, 0x86, 0x63, 0x32,
	0x26, 0x38, 0xd1, 0xbf, 0x02, 0xe6, 0x06, 0x38, 0xf0, 0xdc, 0x68, 0x18, 0xf4, 0x5f, 0x48, 0x13,
	0xf4, 0x87, 0x2e, 0xec, 0x86, 0x9a, 0xd2, 0x3a, 0xbc, 0xbc, 0xb2, 0xf4, 0x20, 0x46, 0x43, 0x22,
	0xa6, 0x72, 0xa7, 0x21, 0x5e, 0x20, 0x96, 0x70, 0xfd, 0x29, 0x30, 0x36, 0x62, 0xbf, 0x1d, 0xf8,
	0xf0, 0xfb, 0xa8, 0xe3, 0xbe, 0x74, 0xcf, 0x3d, 0x5f, 0xdb, 0x6e, 0x1d, 0x5c, 0x5e, 0x59, 0x20,
	0x88, 0x7f, 0xec, 0xa0, 0x14, 0xe5, 0x31, 0xd6, 0x9f, 0x82, 0xa3, 0x4d, 0xa5, 0x8e, 0x17, 0xf5,
	0xa1, 0x7b, 0x1e, 0xf6, 0x7c, 0xa8, 0xed, 0x94, 0x79, 0x05, 0xa3, 0xb8, 0xcf, 0x50, 0xce, 0xc7,
	0x98, 0xe9, 0x8f, 0x80, 0xbe, 0x81, 0x96, 0xc7, 0xe0, 0xfc, 0xb9, 0xa6, 0xb6, 0x1a, 0x97, 0x57,
	0xd6, 0xbe, 0xac, 0x95, 0xe4, 0x93, 0x5b, 0xa0, 0x70, 0xe0, 0x79, 0x7e, 0x18, 0x6a, 0xbb, 0x25,
	0x28, 0x9c, 0xc7, 0x31, 0xe6, 0xfc, 0x16, 0xa8, 0xe7, 0x06, 0x2f, 0x07, 0xd0, 0xd7, 0xf6, 0x4a,
	0x50, 0x0f, 0x91, 0x74, 0xce, 0x70, 0x4b, 0xfd, 0xed, 0x0f, 0x53, 0xe9, 0x3c, 0x7f, 0x73, 0x6d,
	0x2a, 0x6f, 0xaf, 0x4d, 0xe5, 0xdf, 0x6b, 0x53, 0x79, 0x7d, 0x63, 0x6e, 0xbd, 0xbd, 0x31, 0xb7,
	0xfe, 0xba, 0x31, 0xb7, 0x7e, 0x38, 0xd9, 0x98, 0xc6, 0x74, 0x9e, 0x9d, 0x2c, 0xff, 0x97, 0xf1,
	0x14, 0x91, 0xdc, 0xf9, 0x79, 0xfd, 0xdf, 0x2c, 0x06, 0x73, 0xb4, 0x57, 0xbc, 0xe5, 0x17, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xb9, 0xc5, 0x6a, 0xe7, 0x05, 0x00, 0x00,
}

func (m *Draw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Draw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Draw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.UpdatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintDraw(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x92
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintDraw(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x8a
	if m.UpdatedAtHeight != 0 {
		i = encodeVarintDraw(dAtA, i, uint64(m.UpdatedAtHeight))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.CreatedAtHeight != 0 {
		i = encodeVarintDraw(dAtA, i, uint64(m.CreatedAtHeight))
		i--
		dAtA[i] = 0x78
	}
	{
		size := m.TotalWinAmount.Size()
		i -= size
		if _, err := m.TotalWinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDraw(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if m.TotalWinCount != 0 {
		i = encodeVarintDraw(dAtA, i, uint64(m.TotalWinCount))
		i--
		dAtA[i] = 0x60
	}
	if len(m.PrizesRefs) > 0 {
		for iNdEx := len(m.PrizesRefs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PrizesRefs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDraw(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	{
		size := m.PrizePoolRemainsAmount.Size()
		i -= size
		if _, err := m.PrizePoolRemainsAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDraw(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.PrizePoolFreshAmount.Size()
		i -= size
		if _, err := m.PrizePoolFreshAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDraw(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.PrizePool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDraw(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.RandSeed != 0 {
		i = encodeVarintDraw(dAtA, i, uint64(m.RandSeed))
		i--
		dAtA[i] = 0x28
	}
	if m.ErrorState != 0 {
		i = encodeVarintDraw(dAtA, i, uint64(m.ErrorState))
		i--
		dAtA[i] = 0x20
	}
	if m.State != 0 {
		i = encodeVarintDraw(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if m.DrawId != 0 {
		i = encodeVarintDraw(dAtA, i, uint64(m.DrawId))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolId != 0 {
		i = encodeVarintDraw(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDraw(dAtA []byte, offset int, v uint64) int {
	offset -= sovDraw(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Draw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovDraw(uint64(m.PoolId))
	}
	if m.DrawId != 0 {
		n += 1 + sovDraw(uint64(m.DrawId))
	}
	if m.State != 0 {
		n += 1 + sovDraw(uint64(m.State))
	}
	if m.ErrorState != 0 {
		n += 1 + sovDraw(uint64(m.ErrorState))
	}
	if m.RandSeed != 0 {
		n += 1 + sovDraw(uint64(m.RandSeed))
	}
	l = m.PrizePool.Size()
	n += 1 + l + sovDraw(uint64(l))
	l = m.PrizePoolFreshAmount.Size()
	n += 1 + l + sovDraw(uint64(l))
	l = m.PrizePoolRemainsAmount.Size()
	n += 1 + l + sovDraw(uint64(l))
	if len(m.PrizesRefs) > 0 {
		for _, e := range m.PrizesRefs {
			l = e.Size()
			n += 1 + l + sovDraw(uint64(l))
		}
	}
	if m.TotalWinCount != 0 {
		n += 1 + sovDraw(uint64(m.TotalWinCount))
	}
	l = m.TotalWinAmount.Size()
	n += 1 + l + sovDraw(uint64(l))
	if m.CreatedAtHeight != 0 {
		n += 1 + sovDraw(uint64(m.CreatedAtHeight))
	}
	if m.UpdatedAtHeight != 0 {
		n += 2 + sovDraw(uint64(m.UpdatedAtHeight))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt)
	n += 2 + l + sovDraw(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.UpdatedAt)
	n += 2 + l + sovDraw(uint64(l))
	return n
}

func sovDraw(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDraw(x uint64) (n int) {
	return sovDraw(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Draw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDraw
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
			return fmt.Errorf("proto: Draw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Draw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrawId", wireType)
			}
			m.DrawId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DrawId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= DrawState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorState", wireType)
			}
			m.ErrorState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorState |= DrawState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandSeed", wireType)
			}
			m.RandSeed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandSeed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrizePool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
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
				return ErrInvalidLengthDraw
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrizePool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrizePoolFreshAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
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
				return ErrInvalidLengthDraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrizePoolFreshAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrizePoolRemainsAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
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
				return ErrInvalidLengthDraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrizePoolRemainsAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrizesRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
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
				return ErrInvalidLengthDraw
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrizesRefs = append(m.PrizesRefs, PrizeRef{})
			if err := m.PrizesRefs[len(m.PrizesRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWinCount", wireType)
			}
			m.TotalWinCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalWinCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWinAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
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
				return ErrInvalidLengthDraw
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAtHeight", wireType)
			}
			m.CreatedAtHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAtHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAtHeight", wireType)
			}
			m.UpdatedAtHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAtHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
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
				return ErrInvalidLengthDraw
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDraw
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
				return ErrInvalidLengthDraw
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDraw
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDraw(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDraw
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
func skipDraw(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDraw
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
					return 0, ErrIntOverflowDraw
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
					return 0, ErrIntOverflowDraw
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
				return 0, ErrInvalidLengthDraw
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDraw
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDraw
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDraw        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDraw          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDraw = fmt.Errorf("proto: unexpected end of group")
)
