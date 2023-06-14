// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lumnetwork/icacallbacks/callback_data.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

type CallbackData struct {
	CallbackKey  string `protobuf:"bytes,1,opt,name=callback_key,json=callbackKey,proto3" json:"callback_key,omitempty"`
	PortId       string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	ChannelId    string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Sequence     uint64 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	CallbackId   string `protobuf:"bytes,5,opt,name=callback_id,json=callbackId,proto3" json:"callback_id,omitempty"`
	CallbackArgs []byte `protobuf:"bytes,6,opt,name=callback_args,json=callbackArgs,proto3" json:"callback_args,omitempty"`
}

func (m *CallbackData) Reset()         { *m = CallbackData{} }
func (m *CallbackData) String() string { return proto.CompactTextString(m) }
func (*CallbackData) ProtoMessage()    {}
func (*CallbackData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0413fdcb772ec56c, []int{0}
}
func (m *CallbackData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CallbackData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CallbackData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CallbackData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackData.Merge(m, src)
}
func (m *CallbackData) XXX_Size() int {
	return m.Size()
}
func (m *CallbackData) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackData.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackData proto.InternalMessageInfo

func (m *CallbackData) GetCallbackKey() string {
	if m != nil {
		return m.CallbackKey
	}
	return ""
}

func (m *CallbackData) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *CallbackData) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *CallbackData) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *CallbackData) GetCallbackId() string {
	if m != nil {
		return m.CallbackId
	}
	return ""
}

func (m *CallbackData) GetCallbackArgs() []byte {
	if m != nil {
		return m.CallbackArgs
	}
	return nil
}

func init() {
	proto.RegisterType((*CallbackData)(nil), "lumnetwork.icacallbacks.CallbackData")
}

func init() {
	proto.RegisterFile("lumnetwork/icacallbacks/callback_data.proto", fileDescriptor_0413fdcb772ec56c)
}

var fileDescriptor_0413fdcb772ec56c = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x50, 0xbd, 0x4e, 0x84, 0x40,
	0x10, 0x66, 0xf5, 0x44, 0x6f, 0xc4, 0x66, 0x9b, 0x23, 0x26, 0xae, 0xa8, 0x0d, 0x89, 0x11, 0x62,
	0x7c, 0x02, 0x7f, 0x1a, 0x72, 0x1d, 0xa5, 0xcd, 0x65, 0x58, 0x36, 0x40, 0xf8, 0x15, 0x96, 0x28,
	0x6f, 0xe1, 0x63, 0x99, 0xd8, 0x5c, 0x69, 0x69, 0xe0, 0x45, 0x0c, 0xe4, 0x58, 0xb5, 0x9b, 0xef,
	0x6f, 0x66, 0xf2, 0xc1, 0x75, 0xd6, 0xe6, 0x85, 0x90, 0xaf, 0x65, 0x9d, 0xba, 0x09, 0x47, 0x8e,
	0x59, 0x16, 0x20, 0x4f, 0x1b, 0x77, 0x9e, 0x36, 0x21, 0x4a, 0x74, 0xaa, 0xba, 0x94, 0x25, 0x5d,
	0xfd, 0x9a, 0x9d, 0xbf, 0xe6, 0xcb, 0x4f, 0x02, 0xc6, 0xe3, 0x0e, 0x3d, 0xa1, 0x44, 0x7a, 0x01,
	0x86, 0x5a, 0x90, 0x8a, 0xce, 0x24, 0x16, 0xb1, 0x97, 0xfe, 0xf1, 0xcc, 0xad, 0x45, 0x47, 0x57,
	0x70, 0x58, 0x95, 0xb5, 0xdc, 0x24, 0xa1, 0xb9, 0x37, 0xa9, 0xfa, 0x08, 0xbd, 0x90, 0x9e, 0x01,
	0xf0, 0x18, 0x8b, 0x42, 0x64, 0xa3, 0xb6, 0x3f, 0x69, 0xcb, 0x1d, 0xe3, 0x85, 0xf4, 0x14, 0x8e,
	0x1a, 0xf1, 0xd2, 0x8a, 0x82, 0x0b, 0x73, 0x61, 0x11, 0x7b, 0xe1, 0x2b, 0x4c, 0xcf, 0x41, 0x9d,
	0x18, 0xb3, 0x07, 0x53, 0x16, 0x66, 0xca, 0x0b, 0xe9, 0x15, 0x9c, 0x28, 0x03, 0xd6, 0x51, 0x63,
	0xea, 0x16, 0xb1, 0x0d, 0x5f, 0x3d, 0x7b, 0x5f, 0x47, 0xcd, 0xc3, 0xfa, 0xa3, 0x67, 0x64, 0xdb,
	0x33, 0xf2, 0xdd, 0x33, 0xf2, 0x3e, 0x30, 0x6d, 0x3b, 0x30, 0xed, 0x6b, 0x60, 0xda, 0xf3, 0x6d,
	0x94, 0xc8, 0xb8, 0x0d, 0x1c, 0x5e, 0xe6, 0x6e, 0xd6, 0xe6, 0x37, 0x73, 0x73, 0x3c, 0xc6, 0xa4,
	0x70, 0xdf, 0xfe, 0x37, 0x28, 0xbb, 0x4a, 0x34, 0x81, 0x3e, 0x55, 0x77, 0xf7, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0xca, 0x8e, 0x4d, 0x69, 0x01, 0x00, 0x00,
}

func (m *CallbackData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallbackData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CallbackData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CallbackArgs) > 0 {
		i -= len(m.CallbackArgs)
		copy(dAtA[i:], m.CallbackArgs)
		i = encodeVarintCallbackData(dAtA, i, uint64(len(m.CallbackArgs)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CallbackId) > 0 {
		i -= len(m.CallbackId)
		copy(dAtA[i:], m.CallbackId)
		i = encodeVarintCallbackData(dAtA, i, uint64(len(m.CallbackId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Sequence != 0 {
		i = encodeVarintCallbackData(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintCallbackData(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintCallbackData(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CallbackKey) > 0 {
		i -= len(m.CallbackKey)
		copy(dAtA[i:], m.CallbackKey)
		i = encodeVarintCallbackData(dAtA, i, uint64(len(m.CallbackKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCallbackData(dAtA []byte, offset int, v uint64) int {
	offset -= sovCallbackData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CallbackData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CallbackKey)
	if l > 0 {
		n += 1 + l + sovCallbackData(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovCallbackData(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovCallbackData(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovCallbackData(uint64(m.Sequence))
	}
	l = len(m.CallbackId)
	if l > 0 {
		n += 1 + l + sovCallbackData(uint64(l))
	}
	l = len(m.CallbackArgs)
	if l > 0 {
		n += 1 + l + sovCallbackData(uint64(l))
	}
	return n
}

func sovCallbackData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCallbackData(x uint64) (n int) {
	return sovCallbackData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CallbackData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallbackData
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
			return fmt.Errorf("proto: CallbackData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallbackData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbackData
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
				return ErrInvalidLengthCallbackData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbackData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbackData
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
				return ErrInvalidLengthCallbackData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbackData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbackData
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
				return ErrInvalidLengthCallbackData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbackData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbackData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbackData
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
				return ErrInvalidLengthCallbackData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbackData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackArgs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbackData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCallbackData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCallbackData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackArgs = append(m.CallbackArgs[:0], dAtA[iNdEx:postIndex]...)
			if m.CallbackArgs == nil {
				m.CallbackArgs = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCallbackData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallbackData
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
func skipCallbackData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCallbackData
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
					return 0, ErrIntOverflowCallbackData
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
					return 0, ErrIntOverflowCallbackData
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
				return 0, ErrInvalidLengthCallbackData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCallbackData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCallbackData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCallbackData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCallbackData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCallbackData = fmt.Errorf("proto: unexpected end of group")
)
