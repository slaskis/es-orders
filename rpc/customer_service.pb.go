// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/customer_service.proto

/*
	Package rpc is a generated protocol buffer package.

	It is generated from these files:
		rpc/customer_service.proto
		rpc/order_service.proto
		rpc/user_service.proto

	It has these top-level messages:
		Customer
		CustomerResponse
		GetCustomerRequest
		Order
		OrderItem
		Item
		NewItem
		OrderNewRequest
		OrderItemAddRequest
		OrderItemRemoveRequest
		OrderApproveRequest
		OrderRejectRequest
		GetOrderRequest
		OrderResponse
		User
		UserResponse
		GetUserRequest
*/
package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Customer struct {
	ID        string    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Version   int32     `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Name      string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt time.Time `protobuf:"bytes,4,opt,name=createdAt,stdtime" json:"createdAt"`
	UpdatedAt time.Time `protobuf:"bytes,5,opt,name=updatedAt,stdtime" json:"updatedAt"`
}

func (m *Customer) Reset()                    { *m = Customer{} }
func (m *Customer) String() string            { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()               {}
func (*Customer) Descriptor() ([]byte, []int) { return fileDescriptorCustomerService, []int{0} }

func (m *Customer) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Customer) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Customer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Customer) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

type CustomerResponse struct {
	Customer *Customer `protobuf:"bytes,1,opt,name=customer" json:"customer,omitempty"`
}

func (m *CustomerResponse) Reset()                    { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string            { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()               {}
func (*CustomerResponse) Descriptor() ([]byte, []int) { return fileDescriptorCustomerService, []int{1} }

func (m *CustomerResponse) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type GetCustomerRequest struct {
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *GetCustomerRequest) Reset()         { *m = GetCustomerRequest{} }
func (m *GetCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerRequest) ProtoMessage()    {}
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorCustomerService, []int{2}
}

func (m *GetCustomerRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*Customer)(nil), "acme.Customer")
	proto.RegisterType((*CustomerResponse)(nil), "acme.CustomerResponse")
	proto.RegisterType((*GetCustomerRequest)(nil), "acme.GetCustomerRequest")
}
func (m *Customer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Customer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomerService(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Version != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCustomerService(dAtA, i, uint64(m.Version))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCustomerService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintCustomerService(dAtA, i, uint64(types.SizeOfStdTime(m.CreatedAt)))
	n1, err := types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x2a
	i++
	i = encodeVarintCustomerService(dAtA, i, uint64(types.SizeOfStdTime(m.UpdatedAt)))
	n2, err := types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *CustomerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomerResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Customer != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomerService(dAtA, i, uint64(m.Customer.Size()))
		n3, err := m.Customer.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *GetCustomerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCustomerRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomerService(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	return i, nil
}

func encodeVarintCustomerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Customer) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovCustomerService(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovCustomerService(uint64(m.Version))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCustomerService(uint64(l))
	}
	l = types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovCustomerService(uint64(l))
	l = types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovCustomerService(uint64(l))
	return n
}

func (m *CustomerResponse) Size() (n int) {
	var l int
	_ = l
	if m.Customer != nil {
		l = m.Customer.Size()
		n += 1 + l + sovCustomerService(uint64(l))
	}
	return n
}

func (m *GetCustomerRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovCustomerService(uint64(l))
	}
	return n
}

func sovCustomerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCustomerService(x uint64) (n int) {
	return sovCustomerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Customer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Customer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Customer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCustomerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCustomerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCustomerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCustomerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomerService
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
func (m *CustomerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CustomerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCustomerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Customer == nil {
				m.Customer = &Customer{}
			}
			if err := m.Customer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomerService
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
func (m *GetCustomerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetCustomerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCustomerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCustomerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomerService
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
func skipCustomerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCustomerService
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
					return 0, ErrIntOverflowCustomerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCustomerService
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCustomerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCustomerService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCustomerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCustomerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCustomerService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpc/customer_service.proto", fileDescriptorCustomerService) }

var fileDescriptorCustomerService = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0xc6, 0x2d, 0x7f, 0x14, 0x4a, 0x82, 0xda, 0x83, 0x69, 0xf6, 0x00, 0x84, 0x78, 0x20, 0x26,
	0x76, 0x13, 0xbc, 0x6b, 0x44, 0x12, 0xc3, 0x75, 0xf5, 0xe4, 0xc5, 0x2c, 0x65, 0x5c, 0x37, 0xb1,
	0xb4, 0xb6, 0x5d, 0x9e, 0xc3, 0xc7, 0xe2, 0xc8, 0x13, 0xa8, 0xe1, 0x49, 0x0c, 0xdd, 0xed, 0xae,
	0xc1, 0x93, 0xb7, 0x99, 0xdd, 0xdf, 0x7c, 0xf3, 0xcd, 0x57, 0x1c, 0x68, 0xc5, 0x43, 0x9e, 0x19,
	0x2b, 0x05, 0xe8, 0x67, 0x03, 0x7a, 0x95, 0x72, 0x60, 0x4a, 0x4b, 0x2b, 0x49, 0x23, 0xe6, 0x02,
	0x82, 0x7e, 0x22, 0x65, 0xf2, 0x06, 0xa1, 0xfb, 0x36, 0xcf, 0x5e, 0x42, 0x9b, 0x0a, 0x30, 0x36,
	0x16, 0x2a, 0xc7, 0x82, 0xcb, 0x24, 0xb5, 0xaf, 0xd9, 0x9c, 0x71, 0x29, 0xc2, 0x44, 0x26, 0xb2,
	0x22, 0x77, 0x9d, 0x6b, 0x5c, 0x95, 0xe3, 0xc3, 0x0d, 0xc2, 0xad, 0xbb, 0x62, 0x21, 0xe9, 0xe2,
	0xda, 0x6c, 0x4a, 0xd1, 0x00, 0x8d, 0xda, 0x51, 0x6d, 0x36, 0x25, 0x14, 0x1f, 0xad, 0x40, 0x9b,
	0x54, 0x2e, 0x69, 0x6d, 0x80, 0x46, 0xcd, 0xc8, 0xb7, 0x84, 0xe0, 0xc6, 0x32, 0x16, 0x40, 0xeb,
	0x8e, 0x75, 0x35, 0x99, 0xe0, 0x36, 0xd7, 0x10, 0x5b, 0x58, 0xdc, 0x5a, 0xda, 0x18, 0xa0, 0x51,
	0x67, 0x1c, 0xb0, 0xdc, 0x2e, 0xf3, 0x26, 0xd8, 0xa3, 0xb7, 0x3b, 0x69, 0xad, 0x3f, 0xfb, 0x07,
	0x1f, 0x5f, 0x7d, 0x14, 0x55, 0x63, 0x3b, 0x8d, 0x4c, 0x2d, 0x0a, 0x8d, 0xe6, 0x7f, 0x34, 0xca,
	0xb1, 0xe1, 0x35, 0x3e, 0xf1, 0x17, 0x45, 0x60, 0x94, 0x5c, 0x1a, 0x20, 0x17, 0xb8, 0xe5, 0x63,
	0x75, 0xf7, 0x75, 0xc6, 0x5d, 0xb6, 0xcb, 0x93, 0x95, 0x64, 0xf9, 0x7f, 0x78, 0x8e, 0xc9, 0x3d,
	0xd8, 0x4a, 0xe2, 0x3d, 0x03, 0x63, 0xf7, 0xb3, 0x19, 0x47, 0xf8, 0xd8, 0x23, 0x0f, 0xf9, 0x3b,
	0x91, 0x1b, 0xdc, 0xf9, 0x35, 0x48, 0x68, 0xbe, 0xe1, 0xaf, 0x56, 0x70, 0xb6, 0xb7, 0xbb, 0x70,
	0x39, 0x39, 0x5d, 0x6f, 0x7b, 0x68, 0xb3, 0xed, 0xa1, 0xef, 0x6d, 0x0f, 0x3d, 0xd5, 0xb5, 0xe2,
	0xf3, 0x43, 0x77, 0xf5, 0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x6b, 0x3f, 0x11, 0x1a,
	0x02, 0x00, 0x00,
}
