// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: interface.proto

package gRPC

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{1}
}

func (x *PingReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request      *ItemList   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	NumDeliverer uint32      `protobuf:"varint,2,opt,name=num_deliverer,json=numDeliverer,proto3" json:"num_deliverer,omitempty"`
	Itemlists    []*ItemList `protobuf:"bytes,3,rep,name=itemlists,proto3" json:"itemlists,omitempty"`
	// id: 0 purchaser, 1 - n supplier, n+1 - m deliverer
	// (0, 1) (0, 2)...(0, n) (1, 2) (1, 3)...(1, m) (2, 3) (2, 4)...(2, m) ... (n, m)
	Distance []uint64 `protobuf:"varint,4,rep,packed,name=distance,proto3" json:"distance,omitempty"`
}

func (x *ScheduleRequest) Reset() {
	*x = ScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest) ProtoMessage() {}

func (x *ScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest.ProtoReflect.Descriptor instead.
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleRequest) GetRequest() *ItemList {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ScheduleRequest) GetNumDeliverer() uint32 {
	if x != nil {
		return x.NumDeliverer
	}
	return 0
}

func (x *ScheduleRequest) GetItemlists() []*ItemList {
	if x != nil {
		return x.Itemlists
	}
	return nil
}

func (x *ScheduleRequest) GetDistance() []uint64 {
	if x != nil {
		return x.Distance
	}
	return nil
}

type ScheduleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1 based index
	DelivererId uint32   `protobuf:"varint,1,opt,name=deliverer_id,json=delivererId,proto3" json:"deliverer_id,omitempty"`
	Route       []*Route `protobuf:"bytes,2,rep,name=route,proto3" json:"route,omitempty"`
}

func (x *ScheduleReply) Reset() {
	*x = ScheduleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleReply) ProtoMessage() {}

func (x *ScheduleReply) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleReply.ProtoReflect.Descriptor instead.
func (*ScheduleReply) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{3}
}

func (x *ScheduleReply) GetDelivererId() uint32 {
	if x != nil {
		return x.DelivererId
	}
	return 0
}

func (x *ScheduleReply) GetRoute() []*Route {
	if x != nil {
		return x.Route
	}
	return nil
}

type ItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[uint32]float64 `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *ItemList) Reset() {
	*x = ItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemList) ProtoMessage() {}

func (x *ItemList) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemList.ProtoReflect.Descriptor instead.
func (*ItemList) Descriptor() ([]byte, []int) {
	return file_interface_proto_rawDescGZIP(), []int{4}
}

func (x *ItemList) GetItems() map[uint32]float64 {
	if x != nil {
		return x.Items
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1 based index
	SupplierId uint32    `protobuf:"varint,1,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	Itemlist   *ItemList `protobuf:"bytes,2,opt,name=itemlist,proto3" json:"itemlist,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_interface_proto_msgTypes[5]
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
	return file_interface_proto_rawDescGZIP(), []int{5}
}

func (x *Route) GetSupplierId() uint32 {
	if x != nil {
		return x.SupplierId
	}
	return 0
}

func (x *Route) GetItemlist() *ItemList {
	if x != nil {
		return x.Itemlist
	}
	return nil
}

var File_interface_proto protoreflect.FileDescriptor

var file_interface_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75,
	0x6d, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x12,
	0x27, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x50, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x70, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x38,
	0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x5f, 0x0a, 0x09, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x22, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0c,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x08, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b,
	0x67, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interface_proto_rawDescOnce sync.Once
	file_interface_proto_rawDescData = file_interface_proto_rawDesc
)

func file_interface_proto_rawDescGZIP() []byte {
	file_interface_proto_rawDescOnce.Do(func() {
		file_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_interface_proto_rawDescData)
	})
	return file_interface_proto_rawDescData
}

var file_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_interface_proto_goTypes = []interface{}{
	(*PingRequest)(nil),     // 0: PingRequest
	(*PingReply)(nil),       // 1: PingReply
	(*ScheduleRequest)(nil), // 2: ScheduleRequest
	(*ScheduleReply)(nil),   // 3: ScheduleReply
	(*ItemList)(nil),        // 4: ItemList
	(*Route)(nil),           // 5: Route
	nil,                     // 6: ItemList.ItemsEntry
}
var file_interface_proto_depIdxs = []int32{
	4, // 0: ScheduleRequest.request:type_name -> ItemList
	4, // 1: ScheduleRequest.itemlists:type_name -> ItemList
	5, // 2: ScheduleReply.route:type_name -> Route
	6, // 3: ItemList.items:type_name -> ItemList.ItemsEntry
	4, // 4: Route.itemlist:type_name -> ItemList
	0, // 5: Algorithm.Ping:input_type -> PingRequest
	2, // 6: Algorithm.Schedule:input_type -> ScheduleRequest
	1, // 7: Algorithm.Ping:output_type -> PingReply
	3, // 8: Algorithm.Schedule:output_type -> ScheduleReply
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_interface_proto_init() }
func file_interface_proto_init() {
	if File_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReply); i {
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
		file_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleRequest); i {
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
		file_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleReply); i {
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
		file_interface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemList); i {
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
		file_interface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interface_proto_goTypes,
		DependencyIndexes: file_interface_proto_depIdxs,
		MessageInfos:      file_interface_proto_msgTypes,
	}.Build()
	File_interface_proto = out.File
	file_interface_proto_rawDesc = nil
	file_interface_proto_goTypes = nil
	file_interface_proto_depIdxs = nil
}
