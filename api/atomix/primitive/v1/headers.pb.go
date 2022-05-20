// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/primitive/v1/headers.proto

package v1

import (
	fmt "fmt"
	v1 "github.com/atomix/runtime/api/atomix/runtime/v1"
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

type RequestHeaders struct {
	PrimitiveID v1.PrimitiveId `protobuf:"bytes,1,opt,name=primitive_id,json=primitiveId,proto3" json:"primitive_id"`
}

func (m *RequestHeaders) Reset()         { *m = RequestHeaders{} }
func (m *RequestHeaders) String() string { return proto.CompactTextString(m) }
func (*RequestHeaders) ProtoMessage()    {}
func (*RequestHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_89cbf9c7799b548a, []int{0}
}
func (m *RequestHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeaders.Merge(m, src)
}
func (m *RequestHeaders) XXX_Size() int {
	return m.Size()
}
func (m *RequestHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeaders proto.InternalMessageInfo

func (m *RequestHeaders) GetPrimitiveID() v1.PrimitiveId {
	if m != nil {
		return m.PrimitiveID
	}
	return v1.PrimitiveId{}
}

type ResponseHeaders struct {
}

func (m *ResponseHeaders) Reset()         { *m = ResponseHeaders{} }
func (m *ResponseHeaders) String() string { return proto.CompactTextString(m) }
func (*ResponseHeaders) ProtoMessage()    {}
func (*ResponseHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_89cbf9c7799b548a, []int{1}
}
func (m *ResponseHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeaders.Merge(m, src)
}
func (m *ResponseHeaders) XXX_Size() int {
	return m.Size()
}
func (m *ResponseHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeaders proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RequestHeaders)(nil), "atomix.primitive.v1.RequestHeaders")
	proto.RegisterType((*ResponseHeaders)(nil), "atomix.primitive.v1.ResponseHeaders")
}

func init() { proto.RegisterFile("atomix/primitive/v1/headers.proto", fileDescriptor_89cbf9c7799b548a) }

var fileDescriptor_89cbf9c7799b548a = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0x2f, 0x28, 0xca, 0xcc, 0xcd, 0x2c, 0xc9, 0x2c, 0x4b, 0xd5, 0x2f, 0x33, 0xd4,
	0xcf, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86,
	0x28, 0xd1, 0x83, 0x2b, 0xd1, 0x2b, 0x33, 0x94, 0x82, 0xe9, 0x2b, 0x2a, 0xcd, 0x2b, 0xc9, 0xcc,
	0x05, 0xeb, 0x42, 0xc8, 0x83, 0xf5, 0x49, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20,
	0x16, 0x44, 0x54, 0x29, 0x83, 0x8b, 0x2f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0xc4, 0x03, 0x62,
	0x8b, 0x50, 0x18, 0x17, 0x0f, 0x5c, 0x6b, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7,
	0x91, 0x9c, 0x1e, 0xd4, 0x5a, 0xa8, 0x0d, 0x7a, 0x65, 0x86, 0x7a, 0x01, 0x30, 0x65, 0x9e, 0x29,
	0x4e, 0xc2, 0x27, 0xee, 0xc9, 0x33, 0x3c, 0xba, 0x27, 0xcf, 0x8d, 0x10, 0x74, 0x09, 0xe2, 0x2e,
	0x40, 0xa8, 0x50, 0x12, 0xe4, 0xe2, 0x0f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x85, 0x5a,
	0xe5, 0x24, 0x71, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0xd7,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x2c, 0xe0, 0x19, 0x10, 0x01, 0x00, 0x00,
}

func (m *RequestHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestHeaders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestHeaders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PrimitiveID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHeaders(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ResponseHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeaders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseHeaders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintHeaders(dAtA []byte, offset int, v uint64) int {
	offset -= sovHeaders(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PrimitiveID.Size()
	n += 1 + l + sovHeaders(uint64(l))
	return n
}

func (m *ResponseHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovHeaders(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeaders(x uint64) (n int) {
	return sovHeaders(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaders
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
			return fmt.Errorf("proto: RequestHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimitiveID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaders
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
				return ErrInvalidLengthHeaders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeaders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrimitiveID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeaders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHeaders
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
func (m *ResponseHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaders
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
			return fmt.Errorf("proto: ResponseHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHeaders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHeaders
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
func skipHeaders(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeaders
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
					return 0, ErrIntOverflowHeaders
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
					return 0, ErrIntOverflowHeaders
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
				return 0, ErrInvalidLengthHeaders
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHeaders
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHeaders
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHeaders        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeaders          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHeaders = fmt.Errorf("proto: unexpected end of group")
)
