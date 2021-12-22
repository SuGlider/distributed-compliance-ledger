// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dclauth/account_stat.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type AccountStat struct {
	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (m *AccountStat) Reset()         { *m = AccountStat{} }
func (m *AccountStat) String() string { return proto.CompactTextString(m) }
func (*AccountStat) ProtoMessage()    {}
func (*AccountStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7f051f29b648b0, []int{0}
}
func (m *AccountStat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountStat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountStat.Merge(m, src)
}
func (m *AccountStat) XXX_Size() int {
	return m.Size()
}
func (m *AccountStat) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountStat.DiscardUnknown(m)
}

var xxx_messageInfo_AccountStat proto.InternalMessageInfo

func (m *AccountStat) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountStat)(nil), "zigbeealliance.distributedcomplianceledger.dclauth.AccountStat")
}

func init() { proto.RegisterFile("dclauth/account_stat.proto", fileDescriptor_2d7f051f29b648b0) }

var fileDescriptor_2d7f051f29b648b0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x49, 0xce, 0x49,
	0x2c, 0x2d, 0xc9, 0xd0, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x89, 0x2f, 0x2e, 0x49, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0xaa, 0xca, 0x4c, 0x4f, 0x4a, 0x4d, 0x4d, 0xcc,
	0xc9, 0xc9, 0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x4b, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a, 0x2d,
	0x49, 0x4d, 0x49, 0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6, 0xa4, 0xa7, 0x16, 0xe9, 0x41,
	0x8d, 0x51, 0x52, 0xe5, 0xe2, 0x76, 0x84, 0x98, 0x14, 0x5c, 0x92, 0x58, 0x22, 0x24, 0xc6, 0xc5,
	0x96, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0x04, 0xe5, 0x39,
	0x25, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x47, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc4, 0x7e, 0x5d, 0x98, 0x03, 0xf4, 0x91, 0x1c,
	0xa0, 0x8b, 0x70, 0x81, 0x2e, 0xc4, 0x09, 0xfa, 0x15, 0xfa, 0x30, 0xbf, 0x94, 0x54, 0x16, 0xa4,
	0x16, 0x27, 0xb1, 0x81, 0x7d, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x54, 0x33, 0x83, 0x86,
	0xe3, 0x00, 0x00, 0x00,
}

func (m *AccountStat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountStat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountStat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		i = encodeVarintAccountStat(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountStat(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountStat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountStat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovAccountStat(uint64(m.Number))
	}
	return n
}

func sovAccountStat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountStat(x uint64) (n int) {
	return sovAccountStat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountStat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountStat
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
			return fmt.Errorf("proto: AccountStat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountStat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountStat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccountStat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountStat
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
func skipAccountStat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountStat
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
					return 0, ErrIntOverflowAccountStat
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
					return 0, ErrIntOverflowAccountStat
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
				return 0, ErrInvalidLengthAccountStat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountStat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountStat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountStat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountStat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountStat = fmt.Errorf("proto: unexpected end of group")
)