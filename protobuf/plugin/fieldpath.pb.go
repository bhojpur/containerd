// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/protobuf/plugin/fieldpath.proto

package plugin

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	math "math"
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

var E_FieldpathAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63300,
	Name:          "containerd.plugin.fieldpath_all",
	Tag:           "varint,63300,opt,name=fieldpath_all",
	Filename:      "github.com/containerd/containerd/protobuf/plugin/fieldpath.proto",
}

var E_Fieldpath = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64400,
	Name:          "containerd.plugin.fieldpath",
	Tag:           "varint,64400,opt,name=fieldpath",
	Filename:      "github.com/containerd/containerd/protobuf/plugin/fieldpath.proto",
}

func init() {
	proto.RegisterExtension(E_FieldpathAll)
	proto.RegisterExtension(E_Fieldpath)
}

func init() {
	proto.RegisterFile("github.com/containerd/containerd/protobuf/plugin/fieldpath.proto", fileDescriptor_604a244430167409)
}

var fileDescriptor_604a244430167409 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0x41, 0x66, 0x16, 0x14, 0xe5, 0x97, 0xe4, 0x27, 0x95, 0xa6, 0xe9, 0x17, 0xe4, 0x94, 0xa6,
	0x67, 0xe6, 0xe9, 0xa7, 0x65, 0xa6, 0xe6, 0xa4, 0x14, 0x24, 0x96, 0x64, 0xe8, 0x81, 0x65, 0x84,
	0x04, 0x11, 0x6a, 0xf5, 0x20, 0x4a, 0xa4, 0x14, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0x11, 0x5a,
	0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xf2, 0x8b, 0x20, 0x9a, 0xac, 0x9c, 0xb9, 0x78,
	0xe1, 0xe6, 0xc4, 0x27, 0xe6, 0xe4, 0x08, 0xc9, 0xe8, 0x41, 0xf4, 0xe8, 0xc1, 0xf4, 0xe8, 0xb9,
	0x65, 0xe6, 0xa4, 0xfa, 0x17, 0x94, 0x64, 0xe6, 0xe7, 0x15, 0x4b, 0x1c, 0x79, 0xc7, 0xac, 0xc0,
	0xa8, 0xc1, 0x11, 0xc4, 0x03, 0xd7, 0xe4, 0x98, 0x93, 0x63, 0x65, 0xcf, 0xc5, 0x09, 0xe7, 0x0b,
	0xc9, 0x63, 0x18, 0xe0, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x0e, 0x37, 0x63, 0xc2, 0x77, 0x88, 0x19,
	0x08, 0x3d, 0x4e, 0x5e, 0x27, 0x1e, 0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7, 0xd0, 0xf0, 0x48, 0x8e,
	0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x8c, 0x32, 0x20,
	0x35, 0x58, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x81, 0x40, 0x89, 0x49, 0x01, 0x00, 0x00,
}
