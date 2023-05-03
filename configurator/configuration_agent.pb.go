// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: configurator/configuration_agent.proto

package configurator

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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Executed bool   `protobuf:"varint,1,opt,name=executed,proto3" json:"executed,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_configuration_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_configuration_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_configurator_configuration_agent_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetExecuted() bool {
	if x != nil {
		return x.Executed
	}
	return false
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IPAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip        string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (x *IPAssignment) Reset() {
	*x = IPAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_configuration_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPAssignment) ProtoMessage() {}

func (x *IPAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_configuration_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPAssignment.ProtoReflect.Descriptor instead.
func (*IPAssignment) Descriptor() ([]byte, []int) {
	return file_configurator_configuration_agent_proto_rawDescGZIP(), []int{1}
}

func (x *IPAssignment) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *IPAssignment) GetInterface() string {
	if x != nil {
		return x.Interface
	}
	return ""
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationNetwork string `protobuf:"bytes,1,opt,name=destinationNetwork,proto3" json:"destinationNetwork,omitempty"`
	NextHop            string `protobuf:"bytes,2,opt,name=nextHop,proto3" json:"nextHop,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_configuration_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_configuration_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_configurator_configuration_agent_proto_rawDescGZIP(), []int{2}
}

func (x *Route) GetDestinationNetwork() string {
	if x != nil {
		return x.DestinationNetwork
	}
	return ""
}

func (x *Route) GetNextHop() string {
	if x != nil {
		return x.NextHop
	}
	return ""
}

type NatConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalNetwork   string `protobuf:"bytes,1,opt,name=localNetwork,proto3" json:"localNetwork,omitempty"`
	RemoteNetwork  string `protobuf:"bytes,2,opt,name=remoteNetwork,proto3" json:"remoteNetwork,omitempty"`
	LocalInterface string `protobuf:"bytes,3,opt,name=localInterface,proto3" json:"localInterface,omitempty"`
	TableNumber    string `protobuf:"bytes,4,opt,name=tableNumber,proto3" json:"tableNumber,omitempty"`
	PacketMark     string `protobuf:"bytes,5,opt,name=packetMark,proto3" json:"packetMark,omitempty"`
	NextHop        string `protobuf:"bytes,6,opt,name=nextHop,proto3" json:"nextHop,omitempty"`
}

func (x *NatConfiguration) Reset() {
	*x = NatConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_configuration_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatConfiguration) ProtoMessage() {}

func (x *NatConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_configuration_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatConfiguration.ProtoReflect.Descriptor instead.
func (*NatConfiguration) Descriptor() ([]byte, []int) {
	return file_configurator_configuration_agent_proto_rawDescGZIP(), []int{3}
}

func (x *NatConfiguration) GetLocalNetwork() string {
	if x != nil {
		return x.LocalNetwork
	}
	return ""
}

func (x *NatConfiguration) GetRemoteNetwork() string {
	if x != nil {
		return x.RemoteNetwork
	}
	return ""
}

func (x *NatConfiguration) GetLocalInterface() string {
	if x != nil {
		return x.LocalInterface
	}
	return ""
}

func (x *NatConfiguration) GetTableNumber() string {
	if x != nil {
		return x.TableNumber
	}
	return ""
}

func (x *NatConfiguration) GetPacketMark() string {
	if x != nil {
		return x.PacketMark
	}
	return ""
}

func (x *NatConfiguration) GetNextHop() string {
	if x != nil {
		return x.NextHop
	}
	return ""
}

var File_configurator_configuration_agent_proto protoreflect.FileDescriptor

var file_configurator_configuration_agent_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x49, 0x50, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x22, 0x51, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2e, 0x0a,
	0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x22, 0xe0, 0x01, 0x0a, 0x10, 0x4e, 0x61, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x32, 0xe4, 0x01, 0x0a, 0x19, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x49,
	0x50, 0x54, 0x6f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x50, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configurator_configuration_agent_proto_rawDescOnce sync.Once
	file_configurator_configuration_agent_proto_rawDescData = file_configurator_configuration_agent_proto_rawDesc
)

func file_configurator_configuration_agent_proto_rawDescGZIP() []byte {
	file_configurator_configuration_agent_proto_rawDescOnce.Do(func() {
		file_configurator_configuration_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_configurator_configuration_agent_proto_rawDescData)
	})
	return file_configurator_configuration_agent_proto_rawDescData
}

var file_configurator_configuration_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_configurator_configuration_agent_proto_goTypes = []interface{}{
	(*Result)(nil),           // 0: configurator.Result
	(*IPAssignment)(nil),     // 1: configurator.IPAssignment
	(*Route)(nil),            // 2: configurator.Route
	(*NatConfiguration)(nil), // 3: configurator.NatConfiguration
}
var file_configurator_configuration_agent_proto_depIdxs = []int32{
	1, // 0: configurator.ConfigurationAgentService.AddIPToInterface:input_type -> configurator.IPAssignment
	2, // 1: configurator.ConfigurationAgentService.AddRoute:input_type -> configurator.Route
	3, // 2: configurator.ConfigurationAgentService.ConfigureNat:input_type -> configurator.NatConfiguration
	0, // 3: configurator.ConfigurationAgentService.AddIPToInterface:output_type -> configurator.Result
	0, // 4: configurator.ConfigurationAgentService.AddRoute:output_type -> configurator.Result
	0, // 5: configurator.ConfigurationAgentService.ConfigureNat:output_type -> configurator.Result
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_configurator_configuration_agent_proto_init() }
func file_configurator_configuration_agent_proto_init() {
	if File_configurator_configuration_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configurator_configuration_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_configurator_configuration_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPAssignment); i {
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
		file_configurator_configuration_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
		file_configurator_configuration_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NatConfiguration); i {
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
			RawDescriptor: file_configurator_configuration_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_configurator_configuration_agent_proto_goTypes,
		DependencyIndexes: file_configurator_configuration_agent_proto_depIdxs,
		MessageInfos:      file_configurator_configuration_agent_proto_msgTypes,
	}.Build()
	File_configurator_configuration_agent_proto = out.File
	file_configurator_configuration_agent_proto_rawDesc = nil
	file_configurator_configuration_agent_proto_goTypes = nil
	file_configurator_configuration_agent_proto_depIdxs = nil
}