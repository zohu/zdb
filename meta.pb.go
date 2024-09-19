// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: meta.proto

package zdb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DatabaseType int32

const (
	DatabaseType_DATABASE_UNKNOWN    DatabaseType = 0
	DatabaseType_DATABASE_MYSQL      DatabaseType = 1
	DatabaseType_DATABASE_POSTGRESQL DatabaseType = 2
)

// Enum value maps for DatabaseType.
var (
	DatabaseType_name = map[int32]string{
		0: "DATABASE_UNKNOWN",
		1: "DATABASE_MYSQL",
		2: "DATABASE_POSTGRESQL",
	}
	DatabaseType_value = map[string]int32{
		"DATABASE_UNKNOWN":    0,
		"DATABASE_MYSQL":      1,
		"DATABASE_POSTGRESQL": 2,
	}
)

func (x DatabaseType) Enum() *DatabaseType {
	p := new(DatabaseType)
	*p = x
	return p
}

func (x DatabaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatabaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_meta_proto_enumTypes[0].Descriptor()
}

func (DatabaseType) Type() protoreflect.EnumType {
	return &file_meta_proto_enumTypes[0]
}

func (x DatabaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatabaseType.Descriptor instead.
func (DatabaseType) EnumDescriptor() ([]byte, []int) {
	return file_meta_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: yaml:"kind"
	Kind DatabaseType `protobuf:"varint,1,opt,name=kind,proto3,enum=zdb.DatabaseType" json:"kind,omitempty" yaml:"kind"`
	// @gotags: yaml:"host"
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty" yaml:"host"`
	// @gotags: yaml:"port"
	Port string `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty" yaml:"port"`
	// @gotags: yaml:"username"
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty" yaml:"username"`
	// @gotags: yaml:"password"
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty" yaml:"password"`
	// @gotags: yaml:"database"
	Database string `protobuf:"bytes,6,opt,name=database,proto3" json:"database,omitempty" yaml:"database"`
	// @gotags: yaml:"config"
	Config string `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty" yaml:"config"`
	// @gotags: yaml:"max_idle"
	MaxIdle int32 `protobuf:"varint,8,opt,name=max_idle,json=maxIdle,proto3" json:"max_idle,omitempty" yaml:"max_idle"`
	// @gotags: yaml:"max_open"
	MaxAlive int32 `protobuf:"varint,9,opt,name=max_alive,json=maxAlive,proto3" json:"max_alive,omitempty" yaml:"max_open"`
	// @gotags: yaml:"max_alive_minutes"
	MaxAliveMinutes int32 `protobuf:"varint,10,opt,name=max_alive_minutes,json=maxAliveMinutes,proto3" json:"max_alive_minutes,omitempty" yaml:"max_alive_minutes"`
	// @gotags: yaml:"slow_threshold_second"
	SlowThresholdSecond int32 `protobuf:"varint,11,opt,name=slow_threshold_second,json=slowThresholdSecond,proto3" json:"slow_threshold_second,omitempty" yaml:"slow_threshold_second"`
	// @gotags: yaml:"skip_caller_lookup"
	SkipCallerLookup bool `protobuf:"varint,12,opt,name=skip_caller_lookup,json=skipCallerLookup,proto3" json:"skip_caller_lookup,omitempty" yaml:"skip_caller_lookup"`
	// @gotags: yaml:"ignore_record_not_found_error"
	IgnoreRecordNotFoundError bool `protobuf:"varint,13,opt,name=ignore_record_not_found_error,json=ignoreRecordNotFoundError,proto3" json:"ignore_record_not_found_error,omitempty" yaml:"ignore_record_not_found_error"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_meta_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetKind() DatabaseType {
	if x != nil {
		return x.Kind
	}
	return DatabaseType_DATABASE_UNKNOWN
}

func (x *Config) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Config) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Config) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Config) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Config) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Config) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Config) GetMaxIdle() int32 {
	if x != nil {
		return x.MaxIdle
	}
	return 0
}

func (x *Config) GetMaxAlive() int32 {
	if x != nil {
		return x.MaxAlive
	}
	return 0
}

func (x *Config) GetMaxAliveMinutes() int32 {
	if x != nil {
		return x.MaxAliveMinutes
	}
	return 0
}

func (x *Config) GetSlowThresholdSecond() int32 {
	if x != nil {
		return x.SlowThresholdSecond
	}
	return 0
}

func (x *Config) GetSkipCallerLookup() bool {
	if x != nil {
		return x.SkipCallerLookup
	}
	return false
}

func (x *Config) GetIgnoreRecordNotFoundError() bool {
	if x != nil {
		return x.IgnoreRecordNotFoundError
	}
	return false
}

var File_meta_proto protoreflect.FileDescriptor

var file_meta_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x7a, 0x64,
	0x62, 0x22, 0xcb, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x7a, 0x64, 0x62,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x69,
	0x64, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x49, 0x64,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x41,
	0x6c, 0x69, 0x76, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x73,
	0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x73, 0x6c, 0x6f, 0x77,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12,
	0x2c, 0x0a, 0x12, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x6c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x6b, 0x69,
	0x70, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x40, 0x0a,
	0x1d, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2a,
	0x51, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53,
	0x45, 0x5f, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x41, 0x54,
	0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x51, 0x4c,
	0x10, 0x02, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x6f, 0x68, 0x75, 0x2f, 0x7a, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_meta_proto_rawDescOnce sync.Once
	file_meta_proto_rawDescData = file_meta_proto_rawDesc
)

func file_meta_proto_rawDescGZIP() []byte {
	file_meta_proto_rawDescOnce.Do(func() {
		file_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_proto_rawDescData)
	})
	return file_meta_proto_rawDescData
}

var file_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meta_proto_goTypes = []any{
	(DatabaseType)(0), // 0: zdb.DatabaseType
	(*Config)(nil),    // 1: zdb.Config
}
var file_meta_proto_depIdxs = []int32{
	0, // 0: zdb.Config.kind:type_name -> zdb.DatabaseType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_meta_proto_init() }
func file_meta_proto_init() {
	if File_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meta_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_meta_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_proto_goTypes,
		DependencyIndexes: file_meta_proto_depIdxs,
		EnumInfos:         file_meta_proto_enumTypes,
		MessageInfos:      file_meta_proto_msgTypes,
	}.Build()
	File_meta_proto = out.File
	file_meta_proto_rawDesc = nil
	file_meta_proto_goTypes = nil
	file_meta_proto_depIdxs = nil
}