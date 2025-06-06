// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: app/config/client.proto

package config

import (
	common "github.com/ydb-platform/fq-connector-go/api/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Connector client configuration
type TClientConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Connector GRPC API endpoint to connect
	ConnectorServerEndpoint *common.TGenericEndpoint `protobuf:"bytes,4,opt,name=connector_server_endpoint,json=connectorServerEndpoint,proto3" json:"connector_server_endpoint,omitempty"`
	// TLS credentials for Connector
	Tls *TClientTLSConfig `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty"`
	// Data source instance we read data from
	DataSourceInstance *common.TGenericDataSourceInstance `protobuf:"bytes,3,opt,name=data_source_instance,json=dataSourceInstance,proto3" json:"data_source_instance,omitempty"`
	// Solomon metrics endpoint to connect
	MetricsServerEndpoint *common.TGenericEndpoint `protobuf:"bytes,5,opt,name=metrics_server_endpoint,json=metricsServerEndpoint,proto3" json:"metrics_server_endpoint,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *TClientConfig) Reset() {
	*x = TClientConfig{}
	mi := &file_app_config_client_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TClientConfig) ProtoMessage() {}

func (x *TClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_client_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TClientConfig.ProtoReflect.Descriptor instead.
func (*TClientConfig) Descriptor() ([]byte, []int) {
	return file_app_config_client_proto_rawDescGZIP(), []int{0}
}

func (x *TClientConfig) GetConnectorServerEndpoint() *common.TGenericEndpoint {
	if x != nil {
		return x.ConnectorServerEndpoint
	}
	return nil
}

func (x *TClientConfig) GetTls() *TClientTLSConfig {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *TClientConfig) GetDataSourceInstance() *common.TGenericDataSourceInstance {
	if x != nil {
		return x.DataSourceInstance
	}
	return nil
}

func (x *TClientConfig) GetMetricsServerEndpoint() *common.TGenericEndpoint {
	if x != nil {
		return x.MetricsServerEndpoint
	}
	return nil
}

type TClientTLSConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// CA certificate path
	Ca string `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	// Disables certificate host name checking.
	// Should be used carefully only for debugging purposes.
	// See https://pkg.go.dev/crypto/tls#Config for more details.
	InsecureSkipVerify bool `protobuf:"varint,2,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecure_skip_verify,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *TClientTLSConfig) Reset() {
	*x = TClientTLSConfig{}
	mi := &file_app_config_client_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TClientTLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TClientTLSConfig) ProtoMessage() {}

func (x *TClientTLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_client_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TClientTLSConfig.ProtoReflect.Descriptor instead.
func (*TClientTLSConfig) Descriptor() ([]byte, []int) {
	return file_app_config_client_proto_rawDescGZIP(), []int{1}
}

func (x *TClientTLSConfig) GetCa() string {
	if x != nil {
		return x.Ca
	}
	return ""
}

func (x *TClientTLSConfig) GetInsecureSkipVerify() bool {
	if x != nil {
		return x.InsecureSkipVerify
	}
	return false
}

var File_app_config_client_proto protoreflect.FileDescriptor

var file_app_config_client_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x4e, 0x59, 0x71, 0x6c, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x3b, 0x79, 0x71, 0x6c, 0x2f, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x0d, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x52, 0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x54, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x17,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x52, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x54, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x12, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x17, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4e, 0x59,
	0x71, 0x6c, 0x2e, 0x54, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x15, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02,
	0x22, 0x54, 0x0a, 0x10, 0x54, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x63, 0x61, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x66, 0x71, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_app_config_client_proto_rawDescOnce sync.Once
	file_app_config_client_proto_rawDescData []byte
)

func file_app_config_client_proto_rawDescGZIP() []byte {
	file_app_config_client_proto_rawDescOnce.Do(func() {
		file_app_config_client_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_config_client_proto_rawDesc), len(file_app_config_client_proto_rawDesc)))
	})
	return file_app_config_client_proto_rawDescData
}

var file_app_config_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_config_client_proto_goTypes = []any{
	(*TClientConfig)(nil),                     // 0: NYql.Connector.App.Config.TClientConfig
	(*TClientTLSConfig)(nil),                  // 1: NYql.Connector.App.Config.TClientTLSConfig
	(*common.TGenericEndpoint)(nil),           // 2: NYql.TGenericEndpoint
	(*common.TGenericDataSourceInstance)(nil), // 3: NYql.TGenericDataSourceInstance
}
var file_app_config_client_proto_depIdxs = []int32{
	2, // 0: NYql.Connector.App.Config.TClientConfig.connector_server_endpoint:type_name -> NYql.TGenericEndpoint
	1, // 1: NYql.Connector.App.Config.TClientConfig.tls:type_name -> NYql.Connector.App.Config.TClientTLSConfig
	3, // 2: NYql.Connector.App.Config.TClientConfig.data_source_instance:type_name -> NYql.TGenericDataSourceInstance
	2, // 3: NYql.Connector.App.Config.TClientConfig.metrics_server_endpoint:type_name -> NYql.TGenericEndpoint
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_app_config_client_proto_init() }
func file_app_config_client_proto_init() {
	if File_app_config_client_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_config_client_proto_rawDesc), len(file_app_config_client_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_config_client_proto_goTypes,
		DependencyIndexes: file_app_config_client_proto_depIdxs,
		MessageInfos:      file_app_config_client_proto_msgTypes,
	}.Build()
	File_app_config_client_proto = out.File
	file_app_config_client_proto_goTypes = nil
	file_app_config_client_proto_depIdxs = nil
}
