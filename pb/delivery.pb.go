// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/delivery.proto

package pb

import (
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

type AssignOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int64                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Latitude      float64                `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float64                `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssignOrderRequest) Reset() {
	*x = AssignOrderRequest{}
	mi := &file_proto_delivery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignOrderRequest) ProtoMessage() {}

func (x *AssignOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_delivery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignOrderRequest.ProtoReflect.Descriptor instead.
func (*AssignOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *AssignOrderRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *AssignOrderRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *AssignOrderRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type AssignOrderResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Message          string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	DeliveryPersonId int64                  `protobuf:"varint,2,opt,name=delivery_person_id,json=deliveryPersonId,proto3" json:"delivery_person_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *AssignOrderResponse) Reset() {
	*x = AssignOrderResponse{}
	mi := &file_proto_delivery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignOrderResponse) ProtoMessage() {}

func (x *AssignOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_delivery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignOrderResponse.ProtoReflect.Descriptor instead.
func (*AssignOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *AssignOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AssignOrderResponse) GetDeliveryPersonId() int64 {
	if x != nil {
		return x.DeliveryPersonId
	}
	return 0
}

type UpdateOrderStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int64                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusRequest) Reset() {
	*x = UpdateOrderStatusRequest{}
	mi := &file_proto_delivery_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusRequest) ProtoMessage() {}

func (x *UpdateOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_delivery_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_delivery_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOrderStatusRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *UpdateOrderStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateOrderStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusResponse) Reset() {
	*x = UpdateOrderStatusResponse{}
	mi := &file_proto_delivery_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusResponse) ProtoMessage() {}

func (x *UpdateOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_delivery_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_delivery_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOrderStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_delivery_proto protoreflect.FileDescriptor

var file_proto_delivery_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x69, 0x0a, 0x12, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa3, 0x01, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x0b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_delivery_proto_rawDescOnce sync.Once
	file_proto_delivery_proto_rawDescData []byte
)

func file_proto_delivery_proto_rawDescGZIP() []byte {
	file_proto_delivery_proto_rawDescOnce.Do(func() {
		file_proto_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_delivery_proto_rawDesc), len(file_proto_delivery_proto_rawDesc)))
	})
	return file_proto_delivery_proto_rawDescData
}

var file_proto_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_delivery_proto_goTypes = []any{
	(*AssignOrderRequest)(nil),        // 0: pb.AssignOrderRequest
	(*AssignOrderResponse)(nil),       // 1: pb.AssignOrderResponse
	(*UpdateOrderStatusRequest)(nil),  // 2: pb.UpdateOrderStatusRequest
	(*UpdateOrderStatusResponse)(nil), // 3: pb.UpdateOrderStatusResponse
}
var file_proto_delivery_proto_depIdxs = []int32{
	0, // 0: pb.DeliveryService.AssignOrder:input_type -> pb.AssignOrderRequest
	2, // 1: pb.DeliveryService.UpdateOrderStatus:input_type -> pb.UpdateOrderStatusRequest
	1, // 2: pb.DeliveryService.AssignOrder:output_type -> pb.AssignOrderResponse
	3, // 3: pb.DeliveryService.UpdateOrderStatus:output_type -> pb.UpdateOrderStatusResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_delivery_proto_init() }
func file_proto_delivery_proto_init() {
	if File_proto_delivery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_delivery_proto_rawDesc), len(file_proto_delivery_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_delivery_proto_goTypes,
		DependencyIndexes: file_proto_delivery_proto_depIdxs,
		MessageInfos:      file_proto_delivery_proto_msgTypes,
	}.Build()
	File_proto_delivery_proto = out.File
	file_proto_delivery_proto_goTypes = nil
	file_proto_delivery_proto_depIdxs = nil
}
