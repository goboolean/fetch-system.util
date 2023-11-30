// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/model/data.proto

package model

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

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price     float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Size      int64   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Timestamp int64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_model_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_api_model_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_api_model_data_proto_rawDescGZIP(), []int{0}
}

func (x *Trade) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Trade) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Trade) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Aggregate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open      float64 `protobuf:"fixed64,1,opt,name=open,proto3" json:"open,omitempty"`
	Closed    float64 `protobuf:"fixed64,2,opt,name=closed,proto3" json:"closed,omitempty"`
	Min       float64 `protobuf:"fixed64,3,opt,name=min,proto3" json:"min,omitempty"`
	Max       float64 `protobuf:"fixed64,4,opt,name=max,proto3" json:"max,omitempty"`
	Volume    int64   `protobuf:"varint,5,opt,name=volume,proto3" json:"volume,omitempty"`
	Timestamp int64   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Aggregate) Reset() {
	*x = Aggregate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_model_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aggregate) ProtoMessage() {}

func (x *Aggregate) ProtoReflect() protoreflect.Message {
	mi := &file_api_model_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aggregate.ProtoReflect.Descriptor instead.
func (*Aggregate) Descriptor() ([]byte, []int) {
	return file_api_model_data_proto_rawDescGZIP(), []int{1}
}

func (x *Aggregate) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Aggregate) GetClosed() float64 {
	if x != nil {
		return x.Closed
	}
	return 0
}

func (x *Aggregate) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Aggregate) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Aggregate) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Aggregate) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_api_model_data_proto protoreflect.FileDescriptor

var file_api_model_data_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x4f, 0x0a,
	0x05, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x91,
	0x01, 0x0a, 0x09, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x6f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x61, 0x43, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_model_data_proto_rawDescOnce sync.Once
	file_api_model_data_proto_rawDescData = file_api_model_data_proto_rawDesc
)

func file_api_model_data_proto_rawDescGZIP() []byte {
	file_api_model_data_proto_rawDescOnce.Do(func() {
		file_api_model_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_model_data_proto_rawDescData)
	})
	return file_api_model_data_proto_rawDescData
}

var file_api_model_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_model_data_proto_goTypes = []interface{}{
	(*Trade)(nil),     // 0: model.Trade
	(*Aggregate)(nil), // 1: model.Aggregate
}
var file_api_model_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_model_data_proto_init() }
func file_api_model_data_proto_init() {
	if File_api_model_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_model_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trade); i {
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
		file_api_model_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aggregate); i {
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
			RawDescriptor: file_api_model_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_model_data_proto_goTypes,
		DependencyIndexes: file_api_model_data_proto_depIdxs,
		MessageInfos:      file_api_model_data_proto_msgTypes,
	}.Build()
	File_api_model_data_proto = out.File
	file_api_model_data_proto_rawDesc = nil
	file_api_model_data_proto_goTypes = nil
	file_api_model_data_proto_depIdxs = nil
}
