// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/tenant.proto

package talent

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum that represents how user data owned by the tenant is used.
type Tenant_DataUsageType int32

const (
	// Default value.
	Tenant_DATA_USAGE_TYPE_UNSPECIFIED Tenant_DataUsageType = 0
	// Data owned by this tenant is used to improve search/recommendation
	// quality across tenants.
	Tenant_AGGREGATED Tenant_DataUsageType = 1
	// Data owned by this tenant is used to improve search/recommendation
	// quality for this tenant only.
	Tenant_ISOLATED Tenant_DataUsageType = 2
)

var Tenant_DataUsageType_name = map[int32]string{
	0: "DATA_USAGE_TYPE_UNSPECIFIED",
	1: "AGGREGATED",
	2: "ISOLATED",
}

var Tenant_DataUsageType_value = map[string]int32{
	"DATA_USAGE_TYPE_UNSPECIFIED": 0,
	"AGGREGATED":                  1,
	"ISOLATED":                    2,
}

func (x Tenant_DataUsageType) String() string {
	return proto.EnumName(Tenant_DataUsageType_name, int32(x))
}

func (Tenant_DataUsageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b0da31c41aea5f3, []int{0, 0}
}

// A Tenant resource represents a tenant in the service. A tenant is a group or
// entity that shares common access with specific privileges for resources like
// profiles. Customer may create multiple tenants to provide data isolation for
// different groups.
type Tenant struct {
	// Required during tenant update.
	//
	// The resource name for a tenant. This is generated by the service when a
	// tenant is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/api-test-project/tenants/foo".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Client side tenant identifier, used to uniquely identify the
	// tenant.
	//
	// The maximum number of allowed characters is 255.
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// Optional. Indicates whether data owned by this tenant may be used to
	// provide product improvements across other tenants.
	//
	// Defaults behavior is
	// [DataUsageType.ISOLATED][google.cloud.talent.v4beta1.Tenant.DataUsageType.ISOLATED]
	// if it's unset.
	UsageType Tenant_DataUsageType `protobuf:"varint,3,opt,name=usage_type,json=usageType,proto3,enum=google.cloud.talent.v4beta1.Tenant_DataUsageType" json:"usage_type,omitempty"`
	// Optional. A list of keys of filterable
	// [Profile.custom_attributes][google.cloud.talent.v4beta1.Profile.custom_attributes],
	// whose corresponding `string_values` are used in keyword searches. Profiles
	// with `string_values` under these specified field keys are returned if any
	// of the values match the search keyword. Custom field values with
	// parenthesis, brackets and special symbols are not searchable as-is,
	// and must be surrounded by quotes.
	KeywordSearchableProfileCustomAttributes []string `protobuf:"bytes,4,rep,name=keyword_searchable_profile_custom_attributes,json=keywordSearchableProfileCustomAttributes,proto3" json:"keyword_searchable_profile_custom_attributes,omitempty"`
	XXX_NoUnkeyedLiteral                     struct{} `json:"-"`
	XXX_unrecognized                         []byte   `json:"-"`
	XXX_sizecache                            int32    `json:"-"`
}

func (m *Tenant) Reset()         { *m = Tenant{} }
func (m *Tenant) String() string { return proto.CompactTextString(m) }
func (*Tenant) ProtoMessage()    {}
func (*Tenant) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b0da31c41aea5f3, []int{0}
}

func (m *Tenant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tenant.Unmarshal(m, b)
}
func (m *Tenant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tenant.Marshal(b, m, deterministic)
}
func (m *Tenant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tenant.Merge(m, src)
}
func (m *Tenant) XXX_Size() int {
	return xxx_messageInfo_Tenant.Size(m)
}
func (m *Tenant) XXX_DiscardUnknown() {
	xxx_messageInfo_Tenant.DiscardUnknown(m)
}

var xxx_messageInfo_Tenant proto.InternalMessageInfo

func (m *Tenant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tenant) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *Tenant) GetUsageType() Tenant_DataUsageType {
	if m != nil {
		return m.UsageType
	}
	return Tenant_DATA_USAGE_TYPE_UNSPECIFIED
}

func (m *Tenant) GetKeywordSearchableProfileCustomAttributes() []string {
	if m != nil {
		return m.KeywordSearchableProfileCustomAttributes
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.talent.v4beta1.Tenant_DataUsageType", Tenant_DataUsageType_name, Tenant_DataUsageType_value)
	proto.RegisterType((*Tenant)(nil), "google.cloud.talent.v4beta1.Tenant")
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/tenant.proto", fileDescriptor_4b0da31c41aea5f3)
}

var fileDescriptor_4b0da31c41aea5f3 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x40, 0xc9, 0x6e, 0x55, 0xb1, 0x06, 0xaa, 0xca, 0x5c, 0x56, 0x2d, 0xd2, 0xae, 0x7a, 0xca,
	0x01, 0xc5, 0x2a, 0x70, 0xe3, 0x94, 0xee, 0x86, 0xd5, 0x4a, 0xa8, 0x44, 0x49, 0xf6, 0x00, 0x07,
	0xac, 0x49, 0x76, 0x1a, 0x22, 0x12, 0x3b, 0xb2, 0x27, 0xc0, 0x1e, 0xf9, 0x15, 0x7e, 0x84, 0x5f,
	0x43, 0xb5, 0xd3, 0x4a, 0x48, 0x55, 0x6f, 0x33, 0x9e, 0xf7, 0x3c, 0x33, 0xb2, 0x59, 0x58, 0x6b,
	0x5d, 0xb7, 0x28, 0xaa, 0x56, 0x0f, 0x7b, 0x41, 0xd0, 0xa2, 0x22, 0xf1, 0xe3, 0x5d, 0x89, 0x04,
	0x97, 0x82, 0x50, 0x81, 0xa2, 0xa8, 0x37, 0x9a, 0x34, 0x3f, 0xf7, 0x64, 0xe4, 0xc8, 0xc8, 0x93,
	0xd1, 0x48, 0x9e, 0xbd, 0x1a, 0xaf, 0x81, 0xbe, 0x11, 0xa0, 0x94, 0x26, 0xa0, 0x46, 0x2b, 0xeb,
	0xd5, 0xb3, 0xc5, 0x58, 0x75, 0x59, 0x39, 0xdc, 0x08, 0x6a, 0x3a, 0xb4, 0x04, 0x5d, 0xef, 0x81,
	0x8b, 0xbf, 0x13, 0x76, 0x5c, 0xb8, 0x66, 0x9c, 0xb3, 0x23, 0x05, 0x1d, 0xce, 0x83, 0x65, 0x10,
	0xce, 0x32, 0x17, 0xf3, 0x05, 0x7b, 0x86, 0xbf, 0x08, 0x8d, 0x82, 0x56, 0x36, 0xfb, 0xf9, 0xc4,
	0x95, 0xd8, 0xdd, 0xd1, 0x76, 0xcf, 0x53, 0xc6, 0x06, 0x0b, 0x35, 0x4a, 0x3a, 0xf4, 0x38, 0x9f,
	0x2e, 0x83, 0xf0, 0xe4, 0xcd, 0x65, 0xf4, 0xc8, 0xc0, 0x91, 0xef, 0x16, 0xad, 0x81, 0x60, 0x77,
	0x6b, 0x16, 0x87, 0x1e, 0xb3, 0xd9, 0x70, 0x17, 0xf2, 0xaf, 0xec, 0xf5, 0x77, 0x3c, 0xfc, 0xd4,
	0x66, 0x2f, 0x2d, 0x82, 0xa9, 0xbe, 0x41, 0xd9, 0xa2, 0xec, 0x8d, 0xbe, 0x69, 0x5a, 0x94, 0xd5,
	0x60, 0x49, 0x77, 0x12, 0x88, 0x4c, 0x53, 0x0e, 0x84, 0x76, 0x7e, 0xb4, 0x9c, 0x86, 0xb3, 0x2c,
	0x1c, 0x9d, 0xfc, 0x5e, 0x49, 0xbd, 0xb1, 0x72, 0x42, 0x7c, 0xcf, 0x5f, 0x5c, 0xb3, 0x17, 0xff,
	0xf5, 0xe6, 0x0b, 0x76, 0xbe, 0x8e, 0x8b, 0x58, 0xee, 0xf2, 0x78, 0x93, 0xc8, 0xe2, 0x73, 0x9a,
	0xc8, 0xdd, 0x75, 0x9e, 0x26, 0xab, 0xed, 0x87, 0x6d, 0xb2, 0x3e, 0x7d, 0xc2, 0x4f, 0x18, 0x8b,
	0x37, 0x9b, 0x2c, 0xd9, 0xc4, 0x45, 0xb2, 0x3e, 0x0d, 0xf8, 0x73, 0xf6, 0x74, 0x9b, 0x7f, 0xfa,
	0xe8, 0xb2, 0xc9, 0xd5, 0xef, 0x80, 0x2d, 0x2a, 0xdd, 0x3d, 0xb6, 0xf3, 0xd5, 0x4b, 0xbf, 0x74,
	0x86, 0x56, 0x0f, 0xa6, 0xba, 0x1d, 0x8d, 0x74, 0x1a, 0x7c, 0x89, 0x47, 0xa7, 0xd6, 0x2d, 0xa8,
	0x3a, 0xd2, 0xa6, 0x16, 0x35, 0x2a, 0xf7, 0x30, 0xc2, 0x97, 0xa0, 0x6f, 0xec, 0x83, 0x3f, 0xe4,
	0xbd, 0x4f, 0xff, 0x4c, 0xa6, 0xab, 0x22, 0x2f, 0x8f, 0x9d, 0xf3, 0xf6, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0xbf, 0xd9, 0x54, 0x54, 0x02, 0x00, 0x00,
}
