// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: kessel/inventory/v1/health.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetLivezRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLivezRequest) Reset() {
	*x = GetLivezRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLivezRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLivezRequest) ProtoMessage() {}

func (x *GetLivezRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLivezRequest.ProtoReflect.Descriptor instead.
func (*GetLivezRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1_health_proto_rawDescGZIP(), []int{0}
}

type GetLivezResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code   uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetLivezResponse) Reset() {
	*x = GetLivezResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLivezResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLivezResponse) ProtoMessage() {}

func (x *GetLivezResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLivezResponse.ProtoReflect.Descriptor instead.
func (*GetLivezResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1_health_proto_rawDescGZIP(), []int{1}
}

func (x *GetLivezResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetLivezResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetReadyzRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetReadyzRequest) Reset() {
	*x = GetReadyzRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadyzRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadyzRequest) ProtoMessage() {}

func (x *GetReadyzRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadyzRequest.ProtoReflect.Descriptor instead.
func (*GetReadyzRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1_health_proto_rawDescGZIP(), []int{2}
}

type GetReadyzResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code   uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetReadyzResponse) Reset() {
	*x = GetReadyzResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadyzResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadyzResponse) ProtoMessage() {}

func (x *GetReadyzResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadyzResponse.ProtoReflect.Descriptor instead.
func (*GetReadyzResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1_health_proto_rawDescGZIP(), []int{3}
}

func (x *GetReadyzResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetReadyzResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_kessel_inventory_v1_health_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1_health_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x76, 0x65, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x7a, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0xa2, 0x02, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x7a, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x7a, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x7a, 0x12,
	0x84, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x7a, 0x12, 0x29, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x7a, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x7a, 0x42, 0x43, 0x0a, 0x23, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1_health_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1_health_proto_rawDescData = file_kessel_inventory_v1_health_proto_rawDesc
)

func file_kessel_inventory_v1_health_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1_health_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1_health_proto_rawDescData)
	})
	return file_kessel_inventory_v1_health_proto_rawDescData
}

var file_kessel_inventory_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kessel_inventory_v1_health_proto_goTypes = []any{
	(*GetLivezRequest)(nil),   // 0: api.kessel.inventory.v1.GetLivezRequest
	(*GetLivezResponse)(nil),  // 1: api.kessel.inventory.v1.GetLivezResponse
	(*GetReadyzRequest)(nil),  // 2: api.kessel.inventory.v1.GetReadyzRequest
	(*GetReadyzResponse)(nil), // 3: api.kessel.inventory.v1.GetReadyzResponse
}
var file_kessel_inventory_v1_health_proto_depIdxs = []int32{
	0, // 0: api.kessel.inventory.v1.InventoryHealthService.GetLivez:input_type -> api.kessel.inventory.v1.GetLivezRequest
	2, // 1: api.kessel.inventory.v1.InventoryHealthService.GetReadyz:input_type -> api.kessel.inventory.v1.GetReadyzRequest
	1, // 2: api.kessel.inventory.v1.InventoryHealthService.GetLivez:output_type -> api.kessel.inventory.v1.GetLivezResponse
	3, // 3: api.kessel.inventory.v1.InventoryHealthService.GetReadyz:output_type -> api.kessel.inventory.v1.GetReadyzResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1_health_proto_init() }
func file_kessel_inventory_v1_health_proto_init() {
	if File_kessel_inventory_v1_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1_health_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetLivezRequest); i {
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
		file_kessel_inventory_v1_health_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetLivezResponse); i {
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
		file_kessel_inventory_v1_health_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetReadyzRequest); i {
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
		file_kessel_inventory_v1_health_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetReadyzResponse); i {
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
			RawDescriptor: file_kessel_inventory_v1_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kessel_inventory_v1_health_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1_health_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1_health_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1_health_proto = out.File
	file_kessel_inventory_v1_health_proto_rawDesc = nil
	file_kessel_inventory_v1_health_proto_goTypes = nil
	file_kessel_inventory_v1_health_proto_depIdxs = nil
}