// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/probes/ping/proto/config.proto

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

// Next tag: 1
type ProbeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Packets per probe
	PacketsPerProbe *int32 `protobuf:"varint,6,opt,name=packets_per_probe,json=packetsPerProbe,def=2" json:"packets_per_probe,omitempty"`
	// How long to wait between two packets to the same target
	PacketsIntervalMsec *int32 `protobuf:"varint,7,opt,name=packets_interval_msec,json=packetsIntervalMsec,def=25" json:"packets_interval_msec,omitempty"`
	// Resolve targets after these many probes
	ResolveTargetsInterval *int32 `protobuf:"varint,9,opt,name=resolve_targets_interval,json=resolveTargetsInterval,def=5" json:"resolve_targets_interval,omitempty"` // =10s
	// Ping payload size in bytes. It cannot be smaller than 8, number of bytes
	// required for the nanoseconds timestamp.
	PayloadSize *int32 `protobuf:"varint,10,opt,name=payload_size,json=payloadSize,def=56" json:"payload_size,omitempty"`
	// Use datagram socket for ICMP.
	// This option enables unprivileged pings (that is, you don't require root
	// privilege to send ICMP packets). Note that most of the Linux distributions
	// don't allow unprivileged pings by default. To enable unprivileged pings on
	// some Linux distributions, you may need to run the following command:
	//
	//     sudo sysctl -w net.ipv4.ping_group_range="0 <large valid group id>"
	//
	// net.ipv4.ping_group_range system setting takes two integers that specify
	// the group id range that is allowed to execute the unprivileged pings. Note
	// that the same setting (with ipv4 in the path) applies to IPv6 as well.
	//
	// Note: This option is not supported on Windows and is automatically
	// disabled there.
	UseDatagramSocket *bool `protobuf:"varint,12,opt,name=use_datagram_socket,json=useDatagramSocket,def=1" json:"use_datagram_socket,omitempty"`
	// Disable integrity checks. To detect data courruption in the network, we
	// craft the outgoing ICMP packet payload in a certain format and verify that
	// the reply payload matches the same format.
	DisableIntegrityCheck *bool `protobuf:"varint,13,opt,name=disable_integrity_check,json=disableIntegrityCheck,def=0" json:"disable_integrity_check,omitempty"`
	// Do not allow OS-level fragmentation, only works on unix systems.
	DisableFragmentation *bool `protobuf:"varint,14,opt,name=disable_fragmentation,json=disableFragmentation,def=0" json:"disable_fragmentation,omitempty"`
}

// Default values for ProbeConf fields.
const (
	Default_ProbeConf_PacketsPerProbe        = int32(2)
	Default_ProbeConf_PacketsIntervalMsec    = int32(25)
	Default_ProbeConf_ResolveTargetsInterval = int32(5)
	Default_ProbeConf_PayloadSize            = int32(56)
	Default_ProbeConf_UseDatagramSocket      = bool(true)
	Default_ProbeConf_DisableIntegrityCheck  = bool(false)
	Default_ProbeConf_DisableFragmentation   = bool(false)
)

func (x *ProbeConf) Reset() {
	*x = ProbeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf) ProtoMessage() {}

func (x *ProbeConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf.ProtoReflect.Descriptor instead.
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProbeConf) GetPacketsPerProbe() int32 {
	if x != nil && x.PacketsPerProbe != nil {
		return *x.PacketsPerProbe
	}
	return Default_ProbeConf_PacketsPerProbe
}

func (x *ProbeConf) GetPacketsIntervalMsec() int32 {
	if x != nil && x.PacketsIntervalMsec != nil {
		return *x.PacketsIntervalMsec
	}
	return Default_ProbeConf_PacketsIntervalMsec
}

func (x *ProbeConf) GetResolveTargetsInterval() int32 {
	if x != nil && x.ResolveTargetsInterval != nil {
		return *x.ResolveTargetsInterval
	}
	return Default_ProbeConf_ResolveTargetsInterval
}

func (x *ProbeConf) GetPayloadSize() int32 {
	if x != nil && x.PayloadSize != nil {
		return *x.PayloadSize
	}
	return Default_ProbeConf_PayloadSize
}

func (x *ProbeConf) GetUseDatagramSocket() bool {
	if x != nil && x.UseDatagramSocket != nil {
		return *x.UseDatagramSocket
	}
	return Default_ProbeConf_UseDatagramSocket
}

func (x *ProbeConf) GetDisableIntegrityCheck() bool {
	if x != nil && x.DisableIntegrityCheck != nil {
		return *x.DisableIntegrityCheck
	}
	return Default_ProbeConf_DisableIntegrityCheck
}

func (x *ProbeConf) GetDisableFragmentation() bool {
	if x != nil && x.DisableFragmentation != nil {
		return *x.DisableFragmentation
	}
	return Default_ProbeConf_DisableFragmentation
}

var File_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x87, 0x03, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x2d, 0x0a, 0x11, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x50, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x36, 0x0a, 0x15, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73,
	0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x32, 0x35, 0x52, 0x13, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x73, 0x65,
	0x63, 0x12, 0x3b, 0x0a, 0x18, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x3a, 0x01, 0x35, 0x52, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x25,
	0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x35, 0x36, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x11, 0x75, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x3d, 0x0a, 0x17, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x52, 0x15, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x3a, 0x0a, 0x15, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65,
	0x52, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_goTypes = []interface{}{
	(*ProbeConf)(nil), // 0: cloudprober.probes.ping.ProbeConf
}
var file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_probes_ping_proto_config_proto_depIdxs = nil
}
