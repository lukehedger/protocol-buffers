// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package tutorial

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

// Define enum types if a field should have one of a predefined list of values
type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Add a message for each data structure to serialise
type Person struct {
	// Specify a name and a type for each field in the message
	// The " = 1", " = 2" markers on each element identify the unique "tag"
	// that field uses in the binary encoding
	// If a field value isn't set, a default value is used: 0, "", false etc.
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	// Other message types can be used as field types
	// Person message contains PhoneNumber messages
	// If a field is `repeated`, the field may be repeated any number of times
	// `repeated` fields are similar to dynamically sized arrays
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

// It is possible to define message types nested inside other messages
type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these
type AddressBook struct {
	// AddressBook message contains Person messages
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x74, 0xd3, 0x74, 0x69, 0x5f, 0xa1, 0xc4, 0x87, 0x48, 0x28, 0x1e, 0x42, 0x4e, 0x01, 0x61,
	0x0f, 0x55, 0xf0, 0x6c, 0xa1, 0xa0, 0x68, 0x4d, 0x59, 0x14, 0xcf, 0x09, 0x79, 0x60, 0x68, 0x92,
	0xb7, 0x6c, 0xd2, 0x43, 0xff, 0xdd, 0x83, 0x64, 0x13, 0x45, 0xc4, 0xdb, 0xbc, 0x99, 0x61, 0x76,
	0x76, 0xe0, 0x3c, 0x2b, 0x0a, 0x4b, 0x6d, 0x9b, 0x33, 0x1f, 0x94, 0xb1, 0xdc, 0x31, 0xce, 0xba,
	0x63, 0xc7, 0xb6, 0xcc, 0xaa, 0xf8, 0x53, 0x80, 0xdc, 0x93, 0x6d, 0xb9, 0x41, 0x04, 0xbf, 0xc9,
	0x6a, 0x0a, 0x45, 0x24, 0x92, 0xb9, 0x76, 0x18, 0x97, 0xe0, 0x95, 0x45, 0xe8, 0x45, 0x22, 0x99,
	0x6a, 0xaf, 0x2c, 0xf0, 0x02, 0xa6, 0x54, 0x67, 0x65, 0x15, 0x4e, 0x9c, 0x69, 0x38, 0xf0, 0x16,
	0xa4, 0xf9, 0xe0, 0x86, 0xda, 0xd0, 0x8f, 0x26, 0xc9, 0x62, 0x7d, 0xa5, 0xbe, 0xf3, 0xd5, 0x90,
	0xad, 0xf6, 0xbd, 0xfc, 0x72, 0xac, 0x73, 0xb2, 0x7a, 0xf4, 0xae, 0xde, 0x60, 0xf1, 0x8b, 0xc6,
	0x4b, 0x90, 0x8d, 0x43, 0x63, 0x81, 0xf1, 0x42, 0x05, 0x7e, 0x77, 0x32, 0xe4, 0x4a, 0x2c, 0xd7,
	0xab, 0xff, 0xa3, 0x5f, 0x4f, 0x86, 0xb4, 0xf3, 0xc5, 0xd7, 0x30, 0xff, 0xa1, 0x10, 0x40, 0xee,
	0xd2, 0xcd, 0xe3, 0xf3, 0x36, 0x38, 0xc3, 0x19, 0xf8, 0x0f, 0xe9, 0x6e, 0x1b, 0x88, 0x1e, 0xbd,
	0xa7, 0xfa, 0x29, 0xf0, 0xe2, 0x3b, 0x58, 0xdc, 0x0f, 0xeb, 0x6c, 0x98, 0x0f, 0x98, 0x80, 0x34,
	0xc4, 0xa6, 0xea, 0x47, 0xe8, 0x3f, 0x12, 0xfc, 0x7d, 0x4d, 0x8f, 0x7a, 0x2e, 0xdd, 0x90, 0x37,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0xe9, 0x13, 0xff, 0x5d, 0x01, 0x00, 0x00,
}
