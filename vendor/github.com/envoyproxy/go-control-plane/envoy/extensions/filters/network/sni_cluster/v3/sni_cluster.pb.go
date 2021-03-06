// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/sni_cluster/v3/sni_cluster.proto

package envoy_extensions_filters_network_sni_cluster_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SniCluster struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SniCluster) Reset()         { *m = SniCluster{} }
func (m *SniCluster) String() string { return proto.CompactTextString(m) }
func (*SniCluster) ProtoMessage()    {}
func (*SniCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ea91e392e51dd8, []int{0}
}

func (m *SniCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SniCluster.Unmarshal(m, b)
}
func (m *SniCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SniCluster.Marshal(b, m, deterministic)
}
func (m *SniCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SniCluster.Merge(m, src)
}
func (m *SniCluster) XXX_Size() int {
	return xxx_messageInfo_SniCluster.Size(m)
}
func (m *SniCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_SniCluster.DiscardUnknown(m)
}

var xxx_messageInfo_SniCluster proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SniCluster)(nil), "envoy.extensions.filters.network.sni_cluster.v3.SniCluster")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/sni_cluster/v3/sni_cluster.proto", fileDescriptor_87ea91e392e51dd8)
}

var fileDescriptor_87ea91e392e51dd8 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0xce, 0xcb,
	0x8c, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0x46, 0xe6, 0xea, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0xe9, 0x83, 0x8d, 0xd0, 0x43, 0x18, 0xa1, 0x07, 0x35, 0x42, 0x0f, 0x6a,
	0x84, 0x1e, 0xb2, 0x9e, 0x32, 0x63, 0x29, 0xd9, 0xd2, 0x94, 0x82, 0x44, 0xfd, 0xc4, 0xbc, 0xbc,
	0xfc, 0x92, 0xc4, 0x12, 0xb0, 0x9d, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x10, 0xf3, 0xa4, 0x14,
	0x31, 0xa4, 0xcb, 0x52, 0x8b, 0x40, 0x06, 0x67, 0xe6, 0xa5, 0x43, 0x94, 0x28, 0x79, 0x71, 0x71,
	0x05, 0xe7, 0x65, 0x3a, 0x43, 0x8c, 0xb4, 0xb2, 0x99, 0x75, 0xb4, 0x43, 0xce, 0x9c, 0xcb, 0x14,
	0xe2, 0x8e, 0xe4, 0xfc, 0xbc, 0xb4, 0xcc, 0x74, 0xa8, 0x1b, 0xb0, 0x3b, 0xc1, 0x48, 0x0f, 0xa1,
	0xdb, 0x29, 0x6a, 0x57, 0xc3, 0x89, 0x8b, 0x6c, 0x4c, 0x02, 0x4c, 0x5c, 0xb6, 0x99, 0xf9, 0x7a,
	0x60, 0x33, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x48, 0xf4, 0x96, 0x13, 0x3f, 0xc2, 0xd0, 0x00,
	0x90, 0x2b, 0x03, 0x18, 0x93, 0xd8, 0xc0, 0xce, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x71,
	0x28, 0x00, 0xbf, 0x66, 0x01, 0x00, 0x00,
}
