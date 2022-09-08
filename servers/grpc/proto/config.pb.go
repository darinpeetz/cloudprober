// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/servers/grpc/proto/config.proto

package proto

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

type ServerConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *int32 `protobuf:"varint,1,opt,name=port,def=3142" json:"port,omitempty"`
	// Enables gRPC reflection for publicly visible services, allowing grpc_cli to
	// work. See https://grpc.io/grpc/core/md_doc_server_reflection_tutorial.html.
	EnableReflection *bool `protobuf:"varint,2,opt,name=enable_reflection,json=enableReflection,def=0" json:"enable_reflection,omitempty"`
	// If use_dedicated_server is set to true, then create a new gRPC server
	// to handle probes. Otherwise, attempt to reuse gRPC server from runconfig
	// if that was set.
	UseDedicatedServer *bool `protobuf:"varint,3,opt,name=use_dedicated_server,json=useDedicatedServer,def=1" json:"use_dedicated_server,omitempty"`
}

// Default values for ServerConf fields.
const (
	Default_ServerConf_Port               = int32(3142)
	Default_ServerConf_EnableReflection   = bool(false)
	Default_ServerConf_UseDedicatedServer = bool(true)
)

func (x *ServerConf) Reset() {
	*x = ServerConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConf) ProtoMessage() {}

func (x *ServerConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConf.ProtoReflect.Descriptor instead.
func (*ServerConf) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ServerConf) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return Default_ServerConf_Port
}

func (x *ServerConf) GetEnableReflection() bool {
	if x != nil && x.EnableReflection != nil {
		return *x.EnableReflection
	}
	return Default_ServerConf_EnableReflection
}

func (x *ServerConf) GetUseDedicatedServer() bool {
	if x != nil && x.UseDedicatedServer != nil {
		return *x.UseDedicatedServer
	}
	return Default_ServerConf_UseDedicatedServer
}

var File_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22, 0x92,
	0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x18, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x04, 0x33, 0x31, 0x34,
	0x32, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x14, 0x75,
	0x73, 0x65, 0x5f, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52,
	0x12, 0x75, 0x73, 0x65, 0x44, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_goTypes = []interface{}{
	(*ServerConf)(nil), // 0: cloudprober.servers.grpc.ServerConf
}
var file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConf); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_servers_grpc_proto_config_proto_depIdxs = nil
}
