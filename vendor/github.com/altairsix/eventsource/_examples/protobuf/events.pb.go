// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	events.proto

It has these top-level messages:
	ItemAdded
	ItemRemoved
	ShoppingCart
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ItemAdded struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at" json:"at,omitempty"`
}

func (m *ItemAdded) Reset()                    { *m = ItemAdded{} }
func (m *ItemAdded) String() string            { return proto.CompactTextString(m) }
func (*ItemAdded) ProtoMessage()               {}
func (*ItemAdded) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ItemAdded) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ItemAdded) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ItemAdded) GetAt() int64 {
	if m != nil {
		return m.At
	}
	return 0
}

type ItemRemoved struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at" json:"at,omitempty"`
}

func (m *ItemRemoved) Reset()                    { *m = ItemRemoved{} }
func (m *ItemRemoved) String() string            { return proto.CompactTextString(m) }
func (*ItemRemoved) ProtoMessage()               {}
func (*ItemRemoved) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ItemRemoved) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ItemRemoved) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ItemRemoved) GetAt() int64 {
	if m != nil {
		return m.At
	}
	return 0
}

type ShoppingCart struct {
	Type int32        `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	A    *ItemAdded   `protobuf:"bytes,2,opt,name=a" json:"a,omitempty"`
	B    *ItemRemoved `protobuf:"bytes,3,opt,name=b" json:"b,omitempty"`
}

func (m *ShoppingCart) Reset()                    { *m = ShoppingCart{} }
func (m *ShoppingCart) String() string            { return proto.CompactTextString(m) }
func (*ShoppingCart) ProtoMessage()               {}
func (*ShoppingCart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ShoppingCart) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ShoppingCart) GetA() *ItemAdded {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *ShoppingCart) GetB() *ItemRemoved {
	if m != nil {
		return m.B
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemAdded)(nil), "main.item_added")
	proto.RegisterType((*ItemRemoved)(nil), "main.item_removed")
	proto.RegisterType((*ShoppingCart)(nil), "main.shopping_cart")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd0, 0xb1, 0x6a, 0xc4, 0x30,
	0x0c, 0x06, 0x60, 0xe4, 0x24, 0x2d, 0x55, 0xd2, 0x52, 0x34, 0x79, 0x2a, 0x26, 0x53, 0xa6, 0x0c,
	0xe9, 0x3b, 0x94, 0xce, 0x7e, 0x81, 0xe0, 0xd4, 0xa2, 0xe7, 0x21, 0xb6, 0x71, 0x4c, 0xe0, 0xde,
	0xfe, 0x88, 0xb9, 0xe3, 0x6e, 0xbf, 0x4d, 0x42, 0xe2, 0x43, 0xbf, 0xb0, 0xe3, 0x9d, 0x7d, 0xde,
	0xc6, 0x98, 0x42, 0x0e, 0x54, 0xaf, 0xc6, 0xf9, 0xfe, 0x07, 0xd1, 0x65, 0x5e, 0x67, 0x63, 0x2d,
	0x5b, 0xfa, 0x40, 0xe1, 0xac, 0x04, 0x05, 0xc3, 0x9b, 0x16, 0xce, 0x92, 0xc4, 0xd7, 0x9d, 0xd3,
	0xe6, 0x82, 0x97, 0x42, 0xc1, 0xd0, 0xe8, 0x5b, 0x7b, 0x6c, 0x9a, 0x2c, 0x2b, 0x05, 0x43, 0xa5,
	0x85, 0xc9, 0xfd, 0x2f, 0x76, 0xc5, 0x49, 0xbc, 0x86, 0xfd, 0x29, 0x89, 0xf1, 0x7d, 0x3b, 0x85,
	0x18, 0x9d, 0xff, 0x9f, 0xff, 0x4c, 0xca, 0x44, 0x58, 0xe7, 0x73, 0xe4, 0x82, 0x35, 0xba, 0xd4,
	0xf4, 0x85, 0x60, 0x0a, 0xd4, 0x4e, 0x9f, 0xe3, 0x11, 0x64, 0xbc, 0xa7, 0xd0, 0x60, 0x48, 0x21,
	0x2c, 0xc5, 0x6c, 0x27, 0x7a, 0x98, 0x5f, 0xaf, 0xd3, 0xb0, 0x2c, 0x2f, 0xe5, 0x0b, 0xdf, 0x97,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xed, 0xb8, 0x99, 0x15, 0x01, 0x00, 0x00,
}
