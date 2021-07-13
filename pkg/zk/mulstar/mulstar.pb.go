// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/zk/mulstar/mulstar.proto

package zkmulstar

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_taurusgroup_cmp_ecdsa_pkg_math_curve "github.com/taurusgroup/cmp-ecdsa/pkg/math/curve"
	paillier "github.com/taurusgroup/cmp-ecdsa/pkg/paillier"
	github_com_taurusgroup_cmp_ecdsa_proto "github.com/taurusgroup/cmp-ecdsa/proto"
	io "io"
	math "math"
	math_big "math/big"
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

type Commitment struct {
	// A = (alpha ⊙ c ) ⊕ Enc(N0, beta, r)
	A *paillier.Ciphertext `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	// Bₓ = gᵃ
	Bx github_com_taurusgroup_cmp_ecdsa_pkg_math_curve.Point `protobuf:"bytes,2,opt,name=Bx,proto3,customtype=github.com/taurusgroup/cmp-ecdsa/pkg/math/curve.Point" json:"Bx"`
	// E = sᵃ tᵍ
	E *math_big.Int `protobuf:"bytes,4,opt,name=E,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"E,omitempty"`
	// S = sˣ tᵐ
	S *math_big.Int `protobuf:"bytes,5,opt,name=S,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"S,omitempty"`
}

func (m *Commitment) Reset()         { *m = Commitment{} }
func (m *Commitment) String() string { return proto.CompactTextString(m) }
func (*Commitment) ProtoMessage()    {}
func (*Commitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_01f4f37e5f048441, []int{0}
}
func (m *Commitment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Commitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Commitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commitment.Merge(m, src)
}
func (m *Commitment) XXX_Size() int {
	return m.Size()
}
func (m *Commitment) XXX_DiscardUnknown() {
	xxx_messageInfo_Commitment.DiscardUnknown(m)
}

var xxx_messageInfo_Commitment proto.InternalMessageInfo

type Proof struct {
	*Commitment `protobuf:"bytes,1,opt,name=C,proto3,embedded=C" json:"C,omitempty"`
	// Z = α + ex
	Z1 *math_big.Int `protobuf:"bytes,2,opt,name=Z1,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"Z1,omitempty"`
	// U = r⋅ρᵉ mod N
	Z2 *math_big.Int `protobuf:"bytes,3,opt,name=Z2,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"Z2,omitempty"`
	// V = s⋅ρₓᵉ
	W *math_big.Int `protobuf:"bytes,4,opt,name=W,proto3,casttypewith=math/big.Int;github.com/taurusgroup/cmp-ecdsa/proto.IntCaster" json:"W,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_01f4f37e5f048441, []int{1}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Commitment)(nil), "zkmulstar.Commitment")
	proto.RegisterType((*Proof)(nil), "zkmulstar.Proof")
}

func init() { proto.RegisterFile("pkg/zk/mulstar/mulstar.proto", fileDescriptor_01f4f37e5f048441) }

var fileDescriptor_01f4f37e5f048441 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xbf, 0x4e, 0xc2, 0x40,
	0x18, 0xc0, 0x7b, 0x15, 0x4c, 0x3c, 0x9d, 0x1a, 0x4d, 0x1a, 0xa2, 0x57, 0xc2, 0x84, 0x03, 0x77,
	0x01, 0xe3, 0x60, 0x08, 0x03, 0x6d, 0x18, 0x1c, 0x8c, 0x58, 0x06, 0x12, 0x9c, 0xda, 0x5a, 0x4a,
	0x03, 0xe5, 0x9a, 0xeb, 0xd5, 0x10, 0x9e, 0xc2, 0x57, 0x70, 0xd4, 0x27, 0x61, 0x64, 0x24, 0x0e,
	0xa8, 0xe5, 0x45, 0x4c, 0x8f, 0x7f, 0x6e, 0x3a, 0x74, 0x6a, 0xbf, 0xbb, 0xef, 0xfb, 0xdd, 0x77,
	0xbf, 0xbb, 0x83, 0xe7, 0xe1, 0xd0, 0x23, 0xd3, 0x21, 0x09, 0xe2, 0x51, 0xc4, 0x2d, 0xb6, 0xfd,
	0xe2, 0x90, 0x51, 0x4e, 0x95, 0xa3, 0xe9, 0x70, 0x33, 0x50, 0xa8, 0x78, 0x3e, 0x1f, 0xc4, 0x36,
	0x76, 0x68, 0x40, 0x3c, 0xea, 0x51, 0x22, 0x32, 0xec, 0xb8, 0x2f, 0x22, 0x11, 0x88, 0xbf, 0x75,
	0x65, 0xe1, 0x22, 0xe5, 0x86, 0x96, 0x3f, 0x1a, 0xf9, 0x2e, 0x23, 0x8e, 0x1f, 0x0e, 0x5c, 0xc6,
	0xdd, 0x09, 0x5f, 0x4f, 0x97, 0x5e, 0x65, 0x08, 0x0d, 0x1a, 0x04, 0x3e, 0x0f, 0xdc, 0x31, 0x57,
	0x4a, 0x10, 0x34, 0x55, 0x50, 0x04, 0xe5, 0xe3, 0xda, 0x29, 0xde, 0x56, 0x61, 0x63, 0x57, 0x65,
	0x82, 0xa6, 0x72, 0x07, 0x65, 0x7d, 0xa2, 0xca, 0x45, 0x50, 0x3e, 0xd1, 0x1b, 0xb3, 0xa5, 0x26,
	0x7d, 0x2c, 0xb5, 0xeb, 0x5f, 0x4d, 0x71, 0x2b, 0x66, 0x71, 0xe4, 0x31, 0x1a, 0x87, 0xc4, 0x09,
	0xc2, 0x8a, 0xeb, 0x3c, 0x45, 0x16, 0x49, 0xdb, 0x08, 0x2c, 0x3e, 0x20, 0x4e, 0xcc, 0x9e, 0x5d,
	0xdc, 0xa6, 0xfe, 0x98, 0x9b, 0xb2, 0x3e, 0x51, 0xee, 0x21, 0x68, 0xa9, 0x39, 0x41, 0x6b, 0xbe,
	0x7f, 0x6a, 0x0d, 0x91, 0x64, 0xfb, 0x1e, 0xbe, 0x1d, 0xf3, 0xfa, 0xdf, 0xd8, 0x74, 0x17, 0x69,
	0xaa, 0x61, 0x45, 0xdc, 0x65, 0x26, 0x68, 0xa5, 0xc0, 0x8e, 0x9a, 0xcf, 0x0c, 0xd8, 0x29, 0xbd,
	0xc9, 0x30, 0xdf, 0x66, 0x94, 0xf6, 0x95, 0x4b, 0x08, 0x8c, 0x8d, 0x9e, 0x33, 0xbc, 0x3b, 0x12,
	0xbc, 0x17, 0xa8, 0xe7, 0xe6, 0x4b, 0x0d, 0x98, 0xc0, 0x50, 0x1e, 0xa0, 0xdc, 0xab, 0x6e, 0x2c,
	0x65, 0xd0, 0x86, 0xdc, 0xab, 0x0a, 0x64, 0x4d, 0x3d, 0xc8, 0x0e, 0x59, 0x4b, 0x5d, 0x75, 0x33,
	0x94, 0xdf, 0xd5, 0x1f, 0x67, 0xdf, 0x48, 0x9a, 0x25, 0x08, 0xcc, 0x13, 0x04, 0x16, 0x09, 0x02,
	0x5f, 0x09, 0x02, 0x2f, 0x2b, 0x24, 0xcd, 0x57, 0x48, 0x5a, 0xac, 0x90, 0xd4, 0xbb, 0xf9, 0xd7,
	0x55, 0xd9, 0xbf, 0x84, 0xfa, 0xce, 0xb3, 0x7d, 0x28, 0x16, 0xbc, 0xfa, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0xd0, 0xd6, 0xb2, 0x2c, 0x03, 0x00, 0x00,
}

func (m *Commitment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Commitment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Commitment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.S)
		i -= size
		if _, err := __caster.MarshalTo(m.S, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMulstar(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.E)
		i -= size
		if _, err := __caster.MarshalTo(m.E, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMulstar(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Bx.Size()
		i -= size
		if _, err := m.Bx.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMulstar(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.A != nil {
		{
			size, err := m.A.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMulstar(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.W)
		i -= size
		if _, err := __caster.MarshalTo(m.W, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMulstar(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.Z2)
		i -= size
		if _, err := __caster.MarshalTo(m.Z2, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMulstar(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		size := __caster.Size(m.Z1)
		i -= size
		if _, err := __caster.MarshalTo(m.Z1, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMulstar(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Commitment != nil {
		{
			size, err := m.Commitment.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMulstar(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMulstar(dAtA []byte, offset int, v uint64) int {
	offset -= sovMulstar(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Commitment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovMulstar(uint64(l))
	}
	l = m.Bx.Size()
	n += 1 + l + sovMulstar(uint64(l))
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.E)
		n += 1 + l + sovMulstar(uint64(l))
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.S)
		n += 1 + l + sovMulstar(uint64(l))
	}
	return n
}

func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Commitment != nil {
		l = m.Commitment.Size()
		n += 1 + l + sovMulstar(uint64(l))
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.Z1)
		n += 1 + l + sovMulstar(uint64(l))
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.Z2)
		n += 1 + l + sovMulstar(uint64(l))
	}
	{
		__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
		l = __caster.Size(m.W)
		n += 1 + l + sovMulstar(uint64(l))
	}
	return n
}

func sovMulstar(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMulstar(x uint64) (n int) {
	return sovMulstar(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Commitment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMulstar
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
			return fmt.Errorf("proto: Commitment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Commitment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMulstar
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
				return ErrInvalidLengthMulstar
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMulstar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.A == nil {
				m.A = &paillier.Ciphertext{}
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMulstar
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
				return ErrInvalidLengthMulstar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMulstar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMulstar
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
				return ErrInvalidLengthMulstar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMulstar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.E = tmp
				}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMulstar
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
				return ErrInvalidLengthMulstar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMulstar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.S = tmp
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMulstar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMulstar
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
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMulstar
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMulstar
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
				return ErrInvalidLengthMulstar
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMulstar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Commitment == nil {
				m.Commitment = &Commitment{}
			}
			if err := m.Commitment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMulstar
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
				return ErrInvalidLengthMulstar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMulstar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Z1 = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z2", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMulstar
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
				return ErrInvalidLengthMulstar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMulstar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Z2 = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field W", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMulstar
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
				return ErrInvalidLengthMulstar
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMulstar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_taurusgroup_cmp_ecdsa_proto.IntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.W = tmp
				}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMulstar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMulstar
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
func skipMulstar(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMulstar
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
					return 0, ErrIntOverflowMulstar
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
					return 0, ErrIntOverflowMulstar
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
				return 0, ErrInvalidLengthMulstar
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMulstar
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMulstar
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMulstar        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMulstar          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMulstar = fmt.Errorf("proto: unexpected end of group")
)