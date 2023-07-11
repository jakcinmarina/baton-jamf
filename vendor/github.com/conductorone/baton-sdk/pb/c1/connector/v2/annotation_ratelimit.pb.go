// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: c1/connector/v2/annotation_ratelimit.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RateLimitDescription_Status int32

const (
	RateLimitDescription_STATUS_UNSPECIFIED RateLimitDescription_Status = 0
	RateLimitDescription_STATUS_OK          RateLimitDescription_Status = 1
	RateLimitDescription_STATUS_OVERLIMIT   RateLimitDescription_Status = 2
	RateLimitDescription_STATUS_ERROR       RateLimitDescription_Status = 3
)

// Enum value maps for RateLimitDescription_Status.
var (
	RateLimitDescription_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_OK",
		2: "STATUS_OVERLIMIT",
		3: "STATUS_ERROR",
	}
	RateLimitDescription_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_OK":          1,
		"STATUS_OVERLIMIT":   2,
		"STATUS_ERROR":       3,
	}
)

func (x RateLimitDescription_Status) Enum() *RateLimitDescription_Status {
	p := new(RateLimitDescription_Status)
	*p = x
	return p
}

func (x RateLimitDescription_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitDescription_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_c1_connector_v2_annotation_ratelimit_proto_enumTypes[0].Descriptor()
}

func (RateLimitDescription_Status) Type() protoreflect.EnumType {
	return &file_c1_connector_v2_annotation_ratelimit_proto_enumTypes[0]
}

func (x RateLimitDescription_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitDescription_Status.Descriptor instead.
func (RateLimitDescription_Status) EnumDescriptor() ([]byte, []int) {
	return file_c1_connector_v2_annotation_ratelimit_proto_rawDescGZIP(), []int{0, 0}
}

type RateLimitDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    RateLimitDescription_Status `protobuf:"varint,1,opt,name=status,proto3,enum=c1.connector.v2.RateLimitDescription_Status" json:"status,omitempty"`
	Limit     int64                       `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Remaining int64                       `protobuf:"varint,3,opt,name=remaining,proto3" json:"remaining,omitempty"`
	ResetAt   *timestamppb.Timestamp      `protobuf:"bytes,4,opt,name=reset_at,json=resetAt,proto3" json:"reset_at,omitempty"`
}

func (x *RateLimitDescription) Reset() {
	*x = RateLimitDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c1_connector_v2_annotation_ratelimit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitDescription) ProtoMessage() {}

func (x *RateLimitDescription) ProtoReflect() protoreflect.Message {
	mi := &file_c1_connector_v2_annotation_ratelimit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitDescription.ProtoReflect.Descriptor instead.
func (*RateLimitDescription) Descriptor() ([]byte, []int) {
	return file_c1_connector_v2_annotation_ratelimit_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimitDescription) GetStatus() RateLimitDescription_Status {
	if x != nil {
		return x.Status
	}
	return RateLimitDescription_STATUS_UNSPECIFIED
}

func (x *RateLimitDescription) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RateLimitDescription) GetRemaining() int64 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

func (x *RateLimitDescription) GetResetAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ResetAt
	}
	return nil
}

var File_c1_connector_v2_annotation_ratelimit_proto protoreflect.FileDescriptor

var file_c1_connector_v2_annotation_ratelimit_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x31,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0,
	0x02, 0x0a, 0x14, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x65, 0x74, 0x41, 0x74, 0x22, 0x57, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x6f, 0x6e, 0x65, 0x2f, 0x62, 0x61, 0x74,
	0x6f, 0x6e, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_c1_connector_v2_annotation_ratelimit_proto_rawDescOnce sync.Once
	file_c1_connector_v2_annotation_ratelimit_proto_rawDescData = file_c1_connector_v2_annotation_ratelimit_proto_rawDesc
)

func file_c1_connector_v2_annotation_ratelimit_proto_rawDescGZIP() []byte {
	file_c1_connector_v2_annotation_ratelimit_proto_rawDescOnce.Do(func() {
		file_c1_connector_v2_annotation_ratelimit_proto_rawDescData = protoimpl.X.CompressGZIP(file_c1_connector_v2_annotation_ratelimit_proto_rawDescData)
	})
	return file_c1_connector_v2_annotation_ratelimit_proto_rawDescData
}

var file_c1_connector_v2_annotation_ratelimit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_c1_connector_v2_annotation_ratelimit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_c1_connector_v2_annotation_ratelimit_proto_goTypes = []interface{}{
	(RateLimitDescription_Status)(0), // 0: c1.connector.v2.RateLimitDescription.Status
	(*RateLimitDescription)(nil),     // 1: c1.connector.v2.RateLimitDescription
	(*timestamppb.Timestamp)(nil),    // 2: google.protobuf.Timestamp
}
var file_c1_connector_v2_annotation_ratelimit_proto_depIdxs = []int32{
	0, // 0: c1.connector.v2.RateLimitDescription.status:type_name -> c1.connector.v2.RateLimitDescription.Status
	2, // 1: c1.connector.v2.RateLimitDescription.reset_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_c1_connector_v2_annotation_ratelimit_proto_init() }
func file_c1_connector_v2_annotation_ratelimit_proto_init() {
	if File_c1_connector_v2_annotation_ratelimit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_c1_connector_v2_annotation_ratelimit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitDescription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_c1_connector_v2_annotation_ratelimit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_c1_connector_v2_annotation_ratelimit_proto_goTypes,
		DependencyIndexes: file_c1_connector_v2_annotation_ratelimit_proto_depIdxs,
		EnumInfos:         file_c1_connector_v2_annotation_ratelimit_proto_enumTypes,
		MessageInfos:      file_c1_connector_v2_annotation_ratelimit_proto_msgTypes,
	}.Build()
	File_c1_connector_v2_annotation_ratelimit_proto = out.File
	file_c1_connector_v2_annotation_ratelimit_proto_rawDesc = nil
	file_c1_connector_v2_annotation_ratelimit_proto_goTypes = nil
	file_c1_connector_v2_annotation_ratelimit_proto_depIdxs = nil
}
