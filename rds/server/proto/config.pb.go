// Configuration proto for ResourceDiscovery (rds) server.
// Example config:
//
// provider {
//   id: "gcp"
//   gcp_config {
//     project_id: 'test-project-id'
//     gce_instances {}
//     forwarding_rules {}
//   }
// }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/rds/server/proto/config.proto

package proto

import (
	proto "github.com/cloudprober/cloudprober/rds/file/proto"
	proto1 "github.com/cloudprober/cloudprober/rds/gcp/proto"
	proto2 "github.com/cloudprober/cloudprober/rds/kubernetes/proto"
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

	// List of providers that server supports.
	Provider []*Provider `protobuf:"bytes,1,rep,name=provider" json:"provider,omitempty"`
}

func (x *ServerConf) Reset() {
	*x = ServerConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConf) ProtoMessage() {}

func (x *ServerConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes[0]
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
	return file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ServerConf) GetProvider() []*Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provider identifier, e.g. "gcp". Server routes incoming requests to various
	// providers based on this id.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Types that are assignable to Config:
	//	*Provider_FileConfig
	//	*Provider_GcpConfig
	//	*Provider_KubernetesConfig
	Config isProvider_Config `protobuf_oneof:"config"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *Provider) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (m *Provider) GetConfig() isProvider_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *Provider) GetFileConfig() *proto.ProviderConfig {
	if x, ok := x.GetConfig().(*Provider_FileConfig); ok {
		return x.FileConfig
	}
	return nil
}

func (x *Provider) GetGcpConfig() *proto1.ProviderConfig {
	if x, ok := x.GetConfig().(*Provider_GcpConfig); ok {
		return x.GcpConfig
	}
	return nil
}

func (x *Provider) GetKubernetesConfig() *proto2.ProviderConfig {
	if x, ok := x.GetConfig().(*Provider_KubernetesConfig); ok {
		return x.KubernetesConfig
	}
	return nil
}

type isProvider_Config interface {
	isProvider_Config()
}

type Provider_FileConfig struct {
	FileConfig *proto.ProviderConfig `protobuf:"bytes,4,opt,name=file_config,json=fileConfig,oneof"`
}

type Provider_GcpConfig struct {
	GcpConfig *proto1.ProviderConfig `protobuf:"bytes,2,opt,name=gcp_config,json=gcpConfig,oneof"`
}

type Provider_KubernetesConfig struct {
	KubernetesConfig *proto2.ProviderConfig `protobuf:"bytes,3,opt,name=kubernetes_config,json=kubernetesConfig,oneof"`
}

func (*Provider_FileConfig) isProvider_Config() {}

func (*Provider_GcpConfig) isProvider_Config() {}

func (*Provider_KubernetesConfig) isProvider_Config() {}

var File_github_com_cloudprober_cloudprober_rds_server_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e,
	0x72, 0x64, 0x73, 0x1a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x67, 0x63, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x8e, 0x02,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x0b, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x0a, 0x67, 0x63, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09,
	0x67, 0x63, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x59, 0x0a, 0x11, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x48, 0x00, 0x52, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_goTypes = []interface{}{
	(*ServerConf)(nil),            // 0: cloudprober.rds.ServerConf
	(*Provider)(nil),              // 1: cloudprober.rds.Provider
	(*proto.ProviderConfig)(nil),  // 2: cloudprober.rds.file.ProviderConfig
	(*proto1.ProviderConfig)(nil), // 3: cloudprober.rds.gcp.ProviderConfig
	(*proto2.ProviderConfig)(nil), // 4: cloudprober.rds.kubernetes.ProviderConfig
}
var file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_depIdxs = []int32{
	1, // 0: cloudprober.rds.ServerConf.provider:type_name -> cloudprober.rds.Provider
	2, // 1: cloudprober.rds.Provider.file_config:type_name -> cloudprober.rds.file.ProviderConfig
	3, // 2: cloudprober.rds.Provider.gcp_config:type_name -> cloudprober.rds.gcp.ProviderConfig
	4, // 3: cloudprober.rds.Provider.kubernetes_config:type_name -> cloudprober.rds.kubernetes.ProviderConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_rds_server_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
	file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Provider_FileConfig)(nil),
		(*Provider_GcpConfig)(nil),
		(*Provider_KubernetesConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_rds_server_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_rds_server_proto_config_proto_depIdxs = nil
}
