// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/certificate.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type Certificate struct {
	PemCert          string   `protobuf:"bytes,1,opt,name=pem_cert,json=pemCert,proto3" json:"pem_cert,omitempty"`
	SerialNumber     string   `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Issuer           string   `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	AuthorityKeyId   string   `protobuf:"bytes,4,opt,name=authority_key_id,json=authorityKeyId,proto3" json:"authority_key_id,omitempty"`
	RootSubject      string   `protobuf:"bytes,5,opt,name=root_subject,json=rootSubject,proto3" json:"root_subject,omitempty"`
	RootSubjectKeyId string   `protobuf:"bytes,6,opt,name=root_subject_key_id,json=rootSubjectKeyId,proto3" json:"root_subject_key_id,omitempty"`
	IsRoot           bool     `protobuf:"varint,7,opt,name=is_root,json=isRoot,proto3" json:"is_root,omitempty"`
	Owner            string   `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Subject          string   `protobuf:"bytes,9,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId     string   `protobuf:"bytes,10,opt,name=subject_key_id,json=subjectKeyId,proto3" json:"subject_key_id,omitempty"`
	Approvals        []*Grant `protobuf:"bytes,11,rep,name=approvals,proto3" json:"approvals,omitempty"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_2657e3d88fce7825, []int{0}
}
func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return m.Size()
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetPemCert() string {
	if m != nil {
		return m.PemCert
	}
	return ""
}

func (m *Certificate) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *Certificate) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Certificate) GetAuthorityKeyId() string {
	if m != nil {
		return m.AuthorityKeyId
	}
	return ""
}

func (m *Certificate) GetRootSubject() string {
	if m != nil {
		return m.RootSubject
	}
	return ""
}

func (m *Certificate) GetRootSubjectKeyId() string {
	if m != nil {
		return m.RootSubjectKeyId
	}
	return ""
}

func (m *Certificate) GetIsRoot() bool {
	if m != nil {
		return m.IsRoot
	}
	return false
}

func (m *Certificate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Certificate) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Certificate) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *Certificate) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func init() {
	proto.RegisterType((*Certificate)(nil), "zigbeealliance.distributedcomplianceledger.pki.Certificate")
}

func init() { proto.RegisterFile("pki/certificate.proto", fileDescriptor_2657e3d88fce7825) }

var fileDescriptor_2657e3d88fce7825 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbb, 0x8e, 0xd4, 0x30,
	0x18, 0x85, 0x27, 0x0c, 0x3b, 0x17, 0x67, 0x59, 0x56, 0xe6, 0xe6, 0xdd, 0x22, 0x1a, 0x2e, 0x45,
	0x9a, 0x24, 0x12, 0x88, 0x07, 0x60, 0x41, 0x42, 0x08, 0x89, 0x22, 0xd3, 0x51, 0x10, 0x39, 0xc9,
	0x4f, 0xd6, 0x4c, 0x12, 0x5b, 0xb6, 0x03, 0x84, 0xa7, 0xe0, 0x61, 0xe8, 0x69, 0x29, 0x57, 0x54,
	0x94, 0x68, 0xe6, 0x45, 0x90, 0xed, 0x0c, 0x13, 0xd1, 0x51, 0xfe, 0x9f, 0x8f, 0xcf, 0x39, 0xfa,
	0x6d, 0x74, 0x47, 0x6c, 0x58, 0x52, 0x80, 0xd4, 0xec, 0x3d, 0x2b, 0xa8, 0x86, 0x58, 0x48, 0xae,
	0x39, 0x8e, 0xbf, 0xb0, 0x2a, 0x07, 0xa0, 0x75, 0xcd, 0x68, 0x5b, 0x40, 0x5c, 0x32, 0xa5, 0x25,
	0xcb, 0x3b, 0x0d, 0x65, 0xc1, 0x1b, 0xe1, 0x68, 0x0d, 0x65, 0x05, 0x32, 0x16, 0x1b, 0x76, 0x7e,
	0x56, 0x70, 0xd5, 0x70, 0x95, 0xd9, 0xdb, 0x89, 0x1b, 0x9c, 0xd5, 0xf9, 0x4d, 0x93, 0x50, 0x49,
	0xda, 0x6a, 0x07, 0x1e, 0x7c, 0x9f, 0x22, 0xff, 0xf9, 0x21, 0x11, 0x9f, 0xa1, 0x85, 0x80, 0x26,
	0x33, 0x25, 0x88, 0xb7, 0xf2, 0xc2, 0x65, 0x3a, 0x17, 0xd0, 0x18, 0x05, 0x7e, 0x88, 0x6e, 0x28,
	0x90, 0x8c, 0xd6, 0x59, 0xdb, 0x35, 0x39, 0x48, 0x72, 0xcd, 0x9e, 0x1f, 0x3b, 0xf8, 0xc6, 0x32,
	0x7c, 0x17, 0xcd, 0x98, 0x52, 0x1d, 0x48, 0x32, 0xb5, 0xa7, 0xc3, 0x84, 0x43, 0x74, 0x4a, 0x3b,
	0x7d, 0xc9, 0x25, 0xd3, 0x7d, 0xb6, 0x81, 0x3e, 0x63, 0x25, 0xb9, 0x6e, 0x15, 0x27, 0x7f, 0xf9,
	0x6b, 0xe8, 0x5f, 0x95, 0xf8, 0x3e, 0x3a, 0x96, 0x9c, 0xeb, 0x4c, 0x75, 0xf9, 0x07, 0x28, 0x34,
	0x39, 0xb2, 0x2a, 0xdf, 0xb0, 0xb5, 0x43, 0x38, 0x42, 0xb7, 0xc6, 0x92, 0xbd, 0xdf, 0xcc, 0x2a,
	0x4f, 0x47, 0x4a, 0xe7, 0x78, 0x0f, 0xcd, 0x99, 0xca, 0x0c, 0x26, 0xf3, 0x95, 0x17, 0x2e, 0x4c,
	0xa9, 0x94, 0x73, 0x8d, 0x63, 0x74, 0xc4, 0x3f, 0xb5, 0x20, 0xc9, 0xc2, 0xdc, 0xbc, 0x20, 0x3f,
	0xbf, 0x45, 0xb7, 0x87, 0x75, 0x3d, 0x2b, 0x4b, 0x09, 0x4a, 0xad, 0xb5, 0x64, 0x6d, 0x95, 0x3a,
	0x19, 0x26, 0x68, 0xbe, 0x6f, 0xb5, 0x74, 0xbb, 0x19, 0x46, 0xfc, 0x08, 0x9d, 0xfc, 0x53, 0x06,
	0x0d, 0xcb, 0x19, 0x17, 0x59, 0xa3, 0x25, 0x15, 0x42, 0xf2, 0x8f, 0xb4, 0x56, 0xc4, 0x5f, 0x4d,
	0x43, 0xff, 0xf1, 0xd3, 0xff, 0x7c, 0xdc, 0xf8, 0xa5, 0x79, 0xbc, 0xf4, 0xe0, 0x73, 0xf1, 0xee,
	0xc7, 0x36, 0xf0, 0xae, 0xb6, 0x81, 0xf7, 0x7b, 0x1b, 0x78, 0x5f, 0x77, 0xc1, 0xe4, 0x6a, 0x17,
	0x4c, 0x7e, 0xed, 0x82, 0xc9, 0xdb, 0x17, 0x15, 0xd3, 0x97, 0x5d, 0x1e, 0x17, 0xbc, 0x49, 0x5c,
	0x4a, 0xb4, 0x8f, 0x49, 0x46, 0x31, 0xd1, 0x21, 0x27, 0x72, 0x41, 0xc9, 0xe7, 0xc4, 0xfc, 0x13,
	0xdd, 0x0b, 0x50, 0xf9, 0xcc, 0x7e, 0x94, 0x27, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xe2,
	0x23, 0x86, 0x9d, 0x02, 0x00, 0x00,
}

func (m *Certificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Certificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintCertificate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x42
	}
	if m.IsRoot {
		i--
		if m.IsRoot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.RootSubjectKeyId) > 0 {
		i -= len(m.RootSubjectKeyId)
		copy(dAtA[i:], m.RootSubjectKeyId)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.RootSubjectKeyId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RootSubject) > 0 {
		i -= len(m.RootSubject)
		copy(dAtA[i:], m.RootSubject)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.RootSubject)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AuthorityKeyId) > 0 {
		i -= len(m.AuthorityKeyId)
		copy(dAtA[i:], m.AuthorityKeyId)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.AuthorityKeyId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SerialNumber) > 0 {
		i -= len(m.SerialNumber)
		copy(dAtA[i:], m.SerialNumber)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.SerialNumber)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PemCert) > 0 {
		i -= len(m.PemCert)
		copy(dAtA[i:], m.PemCert)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.PemCert)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCertificate(dAtA []byte, offset int, v uint64) int {
	offset -= sovCertificate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Certificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PemCert)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.AuthorityKeyId)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.RootSubject)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.RootSubjectKeyId)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	if m.IsRoot {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovCertificate(uint64(l))
		}
	}
	return n
}

func sovCertificate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCertificate(x uint64) (n int) {
	return sovCertificate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Certificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificate
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
			return fmt.Errorf("proto: Certificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PemCert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PemCert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorityKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorityKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootSubject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootSubject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootSubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootSubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRoot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
			m.IsRoot = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
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
			skippy, err := skipCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificate
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
func skipCertificate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCertificate
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
					return 0, ErrIntOverflowCertificate
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
					return 0, ErrIntOverflowCertificate
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
				return 0, ErrInvalidLengthCertificate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCertificate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCertificate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCertificate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCertificate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCertificate = fmt.Errorf("proto: unexpected end of group")
)
