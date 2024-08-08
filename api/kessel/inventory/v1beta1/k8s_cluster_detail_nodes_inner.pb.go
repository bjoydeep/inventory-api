// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/k8s_cluster_detail_nodes_inner.proto

package v1beta1

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

type K8SClusterDetailNodesInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the node (this can contain private info)
	Name string `protobuf:"bytes,3373707,opt,name=name,proto3" json:"name,omitempty"`
	// CPU Capacity of the node defined in CPU units, e.g. \"0.5\"
	Cpu string `protobuf:"bytes,98728,opt,name=cpu,proto3" json:"cpu,omitempty"`
	// Memory Capacity of the node defined as MiB, e.g. \"50Mi\"
	Memory string `protobuf:"bytes,4014849,opt,name=memory,proto3" json:"memory,omitempty"`
	// Map of string keys and string values that can be used to organize and
	// categorize (scope and select) resources
	Labels []*ResourceLabel `protobuf:"bytes,36675587,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *K8SClusterDetailNodesInner) Reset() {
	*x = K8SClusterDetailNodesInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SClusterDetailNodesInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SClusterDetailNodesInner) ProtoMessage() {}

func (x *K8SClusterDetailNodesInner) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SClusterDetailNodesInner.ProtoReflect.Descriptor instead.
func (*K8SClusterDetailNodesInner) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDescGZIP(), []int{0}
}

func (x *K8SClusterDetailNodesInner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *K8SClusterDetailNodesInner) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *K8SClusterDetailNodesInner) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *K8SClusterDetailNodesInner) GetLabels() []*ResourceLabel {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x1a, 0x4b, 0x38, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x8b, 0xf5, 0xcd, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0xa8, 0x83, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x70, 0x75, 0x12, 0x19, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x81, 0x86, 0xf5,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x42, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x83, 0xc0, 0xbe, 0x11, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x42, 0x93, 0x01, 0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x1f,
	0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDescData = file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDesc
)

func file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDescData
}

var file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_goTypes = []any{
	(*K8SClusterDetailNodesInner)(nil), // 0: kessel.inventory.v1beta1.K8sClusterDetailNodesInner
	(*ResourceLabel)(nil),              // 1: kessel.inventory.v1beta1.ResourceLabel
}
var file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_depIdxs = []int32{
	1, // 0: kessel.inventory.v1beta1.K8sClusterDetailNodesInner.labels:type_name -> kessel.inventory.v1beta1.ResourceLabel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_init() }
func file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_init() {
	if File_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_resource_label_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*K8SClusterDetailNodesInner); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto = out.File
	file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_goTypes = nil
	file_kessel_inventory_v1beta1_k8s_cluster_detail_nodes_inner_proto_depIdxs = nil
}
