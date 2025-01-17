// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dclupgrade/proposed_upgrade.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ProposedUpgrade struct {
	Plan      types.Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan"`
	Creator   string     `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Approvals []*Grant   `protobuf:"bytes,3,rep,name=approvals,proto3" json:"approvals,omitempty"`
}

func (m *ProposedUpgrade) Reset()         { *m = ProposedUpgrade{} }
func (m *ProposedUpgrade) String() string { return proto.CompactTextString(m) }
func (*ProposedUpgrade) ProtoMessage()    {}
func (*ProposedUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b2f793e220daf3a, []int{0}
}
func (m *ProposedUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposedUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposedUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposedUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedUpgrade.Merge(m, src)
}
func (m *ProposedUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *ProposedUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedUpgrade proto.InternalMessageInfo

func (m *ProposedUpgrade) GetPlan() types.Plan {
	if m != nil {
		return m.Plan
	}
	return types.Plan{}
}

func (m *ProposedUpgrade) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ProposedUpgrade) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposedUpgrade)(nil), "zigbeealliance.distributedcomplianceledger.dclupgrade.ProposedUpgrade")
}

func init() { proto.RegisterFile("dclupgrade/proposed_upgrade.proto", fileDescriptor_0b2f793e220daf3a) }

var fileDescriptor_0b2f793e220daf3a = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x3f, 0x4f, 0x02, 0x31,
	0x18, 0xc6, 0xaf, 0x42, 0x34, 0x1c, 0x83, 0xc9, 0x85, 0x98, 0x93, 0x98, 0x13, 0x8d, 0x03, 0xcb,
	0xb5, 0x01, 0xa3, 0x93, 0x8b, 0x2c, 0x0e, 0x2e, 0x04, 0xe3, 0xc2, 0x42, 0x7a, 0xd7, 0x37, 0xf5,
	0x92, 0x72, 0x6d, 0xda, 0x42, 0xd4, 0x4f, 0xe1, 0x87, 0xf1, 0x43, 0x30, 0x12, 0x27, 0x5d, 0x8c,
	0x81, 0x2f, 0x62, 0xa0, 0x77, 0xe1, 0x5c, 0xdd, 0xfa, 0xfe, 0xe9, 0xf3, 0x7b, 0xde, 0xc7, 0x3f,
	0x63, 0xa9, 0x98, 0x29, 0xae, 0x29, 0x03, 0xa2, 0xb4, 0x54, 0xd2, 0x00, 0x9b, 0x14, 0x0d, 0xac,
	0xb4, 0xb4, 0x32, 0xb8, 0x7a, 0xcd, 0x78, 0x02, 0x40, 0x85, 0xc8, 0x68, 0x9e, 0x02, 0x66, 0x99,
	0xb1, 0x3a, 0x4b, 0x66, 0x16, 0x58, 0x2a, 0xa7, 0xca, 0x75, 0x05, 0x30, 0x0e, 0x1a, 0xef, 0xd4,
	0xda, 0x2d, 0x2e, 0xb9, 0xdc, 0x2a, 0x90, 0xcd, 0xcb, 0x89, 0xb5, 0x8f, 0x53, 0x69, 0xa6, 0xd2,
	0x4c, 0xdc, 0xc0, 0x15, 0xc5, 0xe8, 0xc2, 0x55, 0xa4, 0xb4, 0x33, 0xef, 0x25, 0x60, 0x69, 0x8f,
	0xfc, 0x71, 0xd3, 0x3e, 0xaa, 0x18, 0xe6, 0x9a, 0xe6, 0xd6, 0xf5, 0xcf, 0xbf, 0x90, 0x7f, 0x38,
	0x2c, 0x0e, 0x78, 0x74, 0xf3, 0xe0, 0xda, 0xaf, 0x2b, 0x41, 0xf3, 0x10, 0x75, 0x50, 0xb7, 0xd9,
	0x3f, 0xc1, 0x05, 0xae, 0x14, 0x2c, 0x00, 0x78, 0x28, 0x68, 0x3e, 0xa8, 0x2f, 0xbe, 0x4f, 0xbd,
	0xd1, 0x76, 0x3f, 0xe8, 0xfb, 0x07, 0xa9, 0x06, 0x6a, 0xa5, 0x0e, 0xf7, 0x3a, 0xa8, 0xdb, 0x18,
	0x84, 0x1f, 0xef, 0x71, 0xab, 0xf8, 0x7d, 0xcb, 0x98, 0x06, 0x63, 0x1e, 0xac, 0xce, 0x72, 0x3e,
	0x2a, 0x17, 0x83, 0xb1, 0xdf, 0xa0, 0x4a, 0x69, 0x39, 0xa7, 0xc2, 0x84, 0xb5, 0x4e, 0xad, 0xdb,
	0xec, 0xdf, 0xe0, 0x7f, 0x25, 0x87, 0xef, 0x36, 0x67, 0x8d, 0x76, 0x72, 0x03, 0x58, 0xac, 0x22,
	0xb4, 0x5c, 0x45, 0xe8, 0x67, 0x15, 0xa1, 0xb7, 0x75, 0xe4, 0x2d, 0xd7, 0x91, 0xf7, 0xb9, 0x8e,
	0xbc, 0xf1, 0x3d, 0xcf, 0xec, 0xd3, 0x2c, 0xc1, 0xa9, 0x9c, 0x12, 0x07, 0x8b, 0x4b, 0x1a, 0xa9,
	0xd0, 0xe2, 0x1d, 0x2e, 0x76, 0x3c, 0xf2, 0x4c, 0x2a, 0x41, 0xda, 0x17, 0x05, 0x26, 0xd9, 0xdf,
	0x26, 0x79, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xc2, 0x2a, 0x02, 0x14, 0x02, 0x00, 0x00,
}

func (m *ProposedUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposedUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposedUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Approvals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposedUpgrade(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProposedUpgrade(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposedUpgrade(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintProposedUpgrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposedUpgrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposedUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Plan.Size()
	n += 1 + l + sovProposedUpgrade(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProposedUpgrade(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovProposedUpgrade(uint64(l))
		}
	}
	return n
}

func sovProposedUpgrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposedUpgrade(x uint64) (n int) {
	return sovProposedUpgrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposedUpgrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposedUpgrade
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
			return fmt.Errorf("proto: ProposedUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposedUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedUpgrade
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
				return ErrInvalidLengthProposedUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedUpgrade
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
				return ErrInvalidLengthProposedUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedUpgrade
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
				return ErrInvalidLengthProposedUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, &Grant{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposedUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposedUpgrade
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
func skipProposedUpgrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposedUpgrade
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
					return 0, ErrIntOverflowProposedUpgrade
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
					return 0, ErrIntOverflowProposedUpgrade
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
				return 0, ErrInvalidLengthProposedUpgrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposedUpgrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposedUpgrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposedUpgrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposedUpgrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposedUpgrade = fmt.Errorf("proto: unexpected end of group")
)
