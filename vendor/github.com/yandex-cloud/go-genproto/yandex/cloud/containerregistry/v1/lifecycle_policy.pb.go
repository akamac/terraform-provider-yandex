// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/containerregistry/v1/lifecycle_policy.proto

package containerregistry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type LifecyclePolicy_Status int32

const (
	LifecyclePolicy_STATUS_UNSPECIFIED LifecyclePolicy_Status = 0
	LifecyclePolicy_ACTIVE             LifecyclePolicy_Status = 1
	LifecyclePolicy_DISABLED           LifecyclePolicy_Status = 2
)

var LifecyclePolicy_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "ACTIVE",
	2: "DISABLED",
}

var LifecyclePolicy_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"ACTIVE":             1,
	"DISABLED":           2,
}

func (x LifecyclePolicy_Status) String() string {
	return proto.EnumName(LifecyclePolicy_Status_name, int32(x))
}

func (LifecyclePolicy_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_400d7b415ccde714, []int{0, 0}
}

type LifecyclePolicy struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RepositoryId         string                 `protobuf:"bytes,3,opt,name=repository_id,json=repositoryId,proto3" json:"repository_id,omitempty"`
	Description          string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status               LifecyclePolicy_Status `protobuf:"varint,5,opt,name=status,proto3,enum=yandex.cloud.containerregistry.v1.LifecyclePolicy_Status" json:"status,omitempty"`
	CreatedAt            *timestamp.Timestamp   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Rules                []*LifecycleRule       `protobuf:"bytes,7,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *LifecyclePolicy) Reset()         { *m = LifecyclePolicy{} }
func (m *LifecyclePolicy) String() string { return proto.CompactTextString(m) }
func (*LifecyclePolicy) ProtoMessage()    {}
func (*LifecyclePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_400d7b415ccde714, []int{0}
}

func (m *LifecyclePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecyclePolicy.Unmarshal(m, b)
}
func (m *LifecyclePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecyclePolicy.Marshal(b, m, deterministic)
}
func (m *LifecyclePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecyclePolicy.Merge(m, src)
}
func (m *LifecyclePolicy) XXX_Size() int {
	return xxx_messageInfo_LifecyclePolicy.Size(m)
}
func (m *LifecyclePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecyclePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_LifecyclePolicy proto.InternalMessageInfo

func (m *LifecyclePolicy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LifecyclePolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LifecyclePolicy) GetRepositoryId() string {
	if m != nil {
		return m.RepositoryId
	}
	return ""
}

func (m *LifecyclePolicy) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LifecyclePolicy) GetStatus() LifecyclePolicy_Status {
	if m != nil {
		return m.Status
	}
	return LifecyclePolicy_STATUS_UNSPECIFIED
}

func (m *LifecyclePolicy) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *LifecyclePolicy) GetRules() []*LifecycleRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type LifecycleRule struct {
	Description          string             `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	ExpirePeriod         *duration.Duration `protobuf:"bytes,2,opt,name=expire_period,json=expirePeriod,proto3" json:"expire_period,omitempty"`
	TagRegexp            string             `protobuf:"bytes,3,opt,name=tag_regexp,json=tagRegexp,proto3" json:"tag_regexp,omitempty"`
	Untagged             bool               `protobuf:"varint,4,opt,name=untagged,proto3" json:"untagged,omitempty"`
	RetainedTop          int64              `protobuf:"varint,5,opt,name=retained_top,json=retainedTop,proto3" json:"retained_top,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LifecycleRule) Reset()         { *m = LifecycleRule{} }
func (m *LifecycleRule) String() string { return proto.CompactTextString(m) }
func (*LifecycleRule) ProtoMessage()    {}
func (*LifecycleRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_400d7b415ccde714, []int{1}
}

func (m *LifecycleRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleRule.Unmarshal(m, b)
}
func (m *LifecycleRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleRule.Marshal(b, m, deterministic)
}
func (m *LifecycleRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleRule.Merge(m, src)
}
func (m *LifecycleRule) XXX_Size() int {
	return xxx_messageInfo_LifecycleRule.Size(m)
}
func (m *LifecycleRule) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleRule.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleRule proto.InternalMessageInfo

func (m *LifecycleRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LifecycleRule) GetExpirePeriod() *duration.Duration {
	if m != nil {
		return m.ExpirePeriod
	}
	return nil
}

func (m *LifecycleRule) GetTagRegexp() string {
	if m != nil {
		return m.TagRegexp
	}
	return ""
}

func (m *LifecycleRule) GetUntagged() bool {
	if m != nil {
		return m.Untagged
	}
	return false
}

func (m *LifecycleRule) GetRetainedTop() int64 {
	if m != nil {
		return m.RetainedTop
	}
	return 0
}

func init() {
	proto.RegisterEnum("yandex.cloud.containerregistry.v1.LifecyclePolicy_Status", LifecyclePolicy_Status_name, LifecyclePolicy_Status_value)
	proto.RegisterType((*LifecyclePolicy)(nil), "yandex.cloud.containerregistry.v1.LifecyclePolicy")
	proto.RegisterType((*LifecycleRule)(nil), "yandex.cloud.containerregistry.v1.LifecycleRule")
}

func init() {
	proto.RegisterFile("yandex/cloud/containerregistry/v1/lifecycle_policy.proto", fileDescriptor_400d7b415ccde714)
}

var fileDescriptor_400d7b415ccde714 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0x6f, 0x12, 0x41,
	0x14, 0xc5, 0x5d, 0xfe, 0x15, 0x06, 0xa8, 0x64, 0x1e, 0xcc, 0x4a, 0xa2, 0x22, 0xc6, 0x84, 0x68,
	0xba, 0x5b, 0xd6, 0x3f, 0xb1, 0x6a, 0x9b, 0x40, 0xa1, 0x91, 0xa4, 0x31, 0xb8, 0x50, 0x4d, 0x7c,
	0xd9, 0x0c, 0x3b, 0xb7, 0xdb, 0x49, 0x96, 0x9d, 0xc9, 0xec, 0x2c, 0x81, 0x37, 0x9f, 0x7d, 0xf0,
	0x6b, 0xd5, 0x8f, 0x64, 0xfa, 0x64, 0x9c, 0x05, 0x85, 0x62, 0xd2, 0xf4, 0x6d, 0xe7, 0xde, 0x7b,
	0xee, 0x9e, 0xfd, 0x9d, 0x1d, 0xf4, 0x66, 0x41, 0x22, 0x0a, 0x73, 0xdb, 0x0f, 0x79, 0x42, 0x6d,
	0x9f, 0x47, 0x8a, 0xb0, 0x08, 0xa4, 0x84, 0x80, 0xc5, 0x4a, 0x2e, 0xec, 0x59, 0xdb, 0x0e, 0xd9,
	0x39, 0xf8, 0x0b, 0x3f, 0x04, 0x4f, 0xf0, 0x90, 0xf9, 0x0b, 0x4b, 0x48, 0xae, 0x38, 0x7e, 0x9c,
	0x2a, 0x2d, 0xad, 0xb4, 0xb6, 0x94, 0xd6, 0xac, 0x5d, 0x7f, 0xb0, 0xb1, 0x7c, 0x46, 0x42, 0x46,
	0x89, 0x62, 0x3c, 0x4a, 0x37, 0xd4, 0x1f, 0x06, 0x9c, 0x07, 0x21, 0xd8, 0xfa, 0x34, 0x49, 0xce,
	0x6d, 0x9a, 0xc8, 0xf5, 0xfe, 0xa3, 0xeb, 0x7d, 0xc5, 0xa6, 0x10, 0x2b, 0x32, 0x15, 0xe9, 0x40,
	0xf3, 0x47, 0x16, 0xdd, 0x3d, 0x5d, 0xb9, 0x1b, 0x6a, 0x73, 0x78, 0x17, 0x65, 0x18, 0x35, 0x8d,
	0x86, 0xd1, 0x2a, 0xb9, 0x19, 0x46, 0x31, 0x46, 0xb9, 0x88, 0x4c, 0xc1, 0xcc, 0xe8, 0x8a, 0x7e,
	0xc6, 0x4f, 0x50, 0x55, 0x82, 0xe0, 0x31, 0x53, 0x5c, 0x2e, 0x3c, 0x46, 0xcd, 0xac, 0x6e, 0x56,
	0xfe, 0x15, 0x07, 0x14, 0x37, 0x50, 0x99, 0x42, 0xec, 0x4b, 0x26, 0xfe, 0x58, 0x32, 0x73, 0x7a,
	0x64, 0xbd, 0x84, 0x3f, 0xa1, 0x42, 0xac, 0x88, 0x4a, 0x62, 0x33, 0xdf, 0x30, 0x5a, 0xbb, 0xce,
	0x81, 0x75, 0x23, 0x12, 0xeb, 0x9a, 0x5d, 0x6b, 0xa4, 0x17, 0xb8, 0xcb, 0x45, 0xf8, 0x00, 0x21,
	0x5f, 0x02, 0x51, 0x40, 0x3d, 0xa2, 0xcc, 0x42, 0xc3, 0x68, 0x95, 0x9d, 0xba, 0x95, 0x72, 0xb0,
	0x56, 0x1c, 0xac, 0xf1, 0x8a, 0x83, 0x5b, 0x5a, 0x4e, 0x77, 0x14, 0x3e, 0x41, 0x79, 0x99, 0x84,
	0x10, 0x9b, 0x3b, 0x8d, 0x6c, 0xab, 0xec, 0xec, 0xdf, 0xc6, 0x8c, 0x9b, 0x84, 0xe0, 0xa6, 0xf2,
	0xe6, 0x5b, 0x54, 0x48, 0x4d, 0xe1, 0x7b, 0x08, 0x8f, 0xc6, 0x9d, 0xf1, 0xd9, 0xc8, 0x3b, 0xfb,
	0x38, 0x1a, 0xf6, 0x8f, 0x07, 0x27, 0x83, 0x7e, 0xaf, 0x76, 0x07, 0x23, 0x54, 0xe8, 0x1c, 0x8f,
	0x07, 0x9f, 0xfb, 0x35, 0x03, 0x57, 0x50, 0xb1, 0x37, 0x18, 0x75, 0xba, 0xa7, 0xfd, 0x5e, 0x2d,
	0xd3, 0xfc, 0x65, 0xa0, 0xea, 0xc6, 0x52, 0xfc, 0x7c, 0x93, 0xa2, 0xce, 0xa5, 0x5b, 0xfa, 0xfe,
	0xb3, 0x9d, 0x7f, 0x7f, 0xe8, 0xbc, 0x7a, 0xbd, 0x09, 0xf4, 0x03, 0xaa, 0xc2, 0x5c, 0x30, 0x09,
	0x9e, 0x00, 0xc9, 0x38, 0xd5, 0xa1, 0x95, 0x9d, 0xfb, 0x5b, 0x00, 0x7a, 0xcb, 0x1f, 0xa5, 0x5b,
	0xbc, 0xba, 0x6c, 0xe7, 0x8e, 0x9c, 0x97, 0x17, 0x6e, 0x25, 0x55, 0x0e, 0xb5, 0x10, 0xb7, 0x10,
	0x52, 0x24, 0xf0, 0x24, 0x04, 0x30, 0x17, 0x69, 0xbc, 0xeb, 0x6f, 0x2d, 0x29, 0x12, 0xb8, 0xba,
	0x87, 0xeb, 0xa8, 0x98, 0x44, 0x8a, 0x04, 0x01, 0x50, 0x9d, 0x71, 0xd1, 0xfd, 0x7b, 0xc6, 0xcf,
	0x50, 0x45, 0x82, 0xc6, 0x46, 0x3d, 0xc5, 0x85, 0x8e, 0x39, 0xdb, 0xdd, 0xb9, 0xba, 0x6c, 0x67,
	0x8f, 0x0e, 0xf7, 0xdd, 0xf2, 0xaa, 0x39, 0xe6, 0xa2, 0xfb, 0xcd, 0x40, 0x4f, 0x37, 0x88, 0x13,
	0xc1, 0xfe, 0x4b, 0xfd, 0xeb, 0x97, 0x80, 0xa9, 0x8b, 0x64, 0x62, 0xf9, 0x7c, 0x6a, 0xa7, 0x8a,
	0xbd, 0xf4, 0x82, 0x04, 0x7c, 0x2f, 0x80, 0x48, 0x7f, 0xa4, 0x7d, 0xe3, 0xb5, 0x7c, 0xb7, 0x55,
	0x9c, 0x14, 0xb4, 0xf4, 0xc5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x39, 0x0f, 0xa6, 0xd4,
	0x03, 0x00, 0x00,
}