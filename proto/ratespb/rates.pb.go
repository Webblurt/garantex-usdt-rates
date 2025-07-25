// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.27.3
// source: proto/rates.proto

package ratespb

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

type GetRatesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRatesRequest) Reset() {
	*x = GetRatesRequest{}
	mi := &file_proto_rates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesRequest) ProtoMessage() {}

func (x *GetRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatesRequest.ProtoReflect.Descriptor instead.
func (*GetRatesRequest) Descriptor() ([]byte, []int) {
	return file_proto_rates_proto_rawDescGZIP(), []int{0}
}

type GetRatesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ask           float64                `protobuf:"fixed64,1,opt,name=ask,proto3" json:"ask,omitempty"`
	Bid           float64                `protobuf:"fixed64,2,opt,name=bid,proto3" json:"bid,omitempty"`
	Timestamp     string                 `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRatesResponse) Reset() {
	*x = GetRatesResponse{}
	mi := &file_proto_rates_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesResponse) ProtoMessage() {}

func (x *GetRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rates_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatesResponse.ProtoReflect.Descriptor instead.
func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return file_proto_rates_proto_rawDescGZIP(), []int{1}
}

func (x *GetRatesResponse) GetAsk() float64 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *GetRatesResponse) GetBid() float64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *GetRatesResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	mi := &file_proto_rates_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rates_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_rates_proto_rawDescGZIP(), []int{2}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp     string                 `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	mi := &file_proto_rates_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rates_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_rates_proto_rawDescGZIP(), []int{3}
}

func (x *HealthCheckResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HealthCheckResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_proto_rates_proto protoreflect.FileDescriptor

const file_proto_rates_proto_rawDesc = "" +
	"\n" +
	"\x11proto/rates.proto\x12\x05rates\"\x11\n" +
	"\x0fGetRatesRequest\"T\n" +
	"\x10GetRatesResponse\x12\x10\n" +
	"\x03ask\x18\x01 \x01(\x01R\x03ask\x12\x10\n" +
	"\x03bid\x18\x02 \x01(\x01R\x03bid\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\tR\ttimestamp\"\x14\n" +
	"\x12HealthCheckRequest\"K\n" +
	"\x13HealthCheckResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\tR\ttimestamp2\x90\x01\n" +
	"\vRateService\x12;\n" +
	"\bGetRates\x12\x16.rates.GetRatesRequest\x1a\x17.rates.GetRatesResponse\x12D\n" +
	"\vHealthCheck\x12\x19.rates.HealthCheckRequest\x1a\x1a.rates.HealthCheckResponseB7Z5github.com/Webblurt/garantex-usdt-rates/proto/ratespbb\x06proto3"

var (
	file_proto_rates_proto_rawDescOnce sync.Once
	file_proto_rates_proto_rawDescData []byte
)

func file_proto_rates_proto_rawDescGZIP() []byte {
	file_proto_rates_proto_rawDescOnce.Do(func() {
		file_proto_rates_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_rates_proto_rawDesc), len(file_proto_rates_proto_rawDesc)))
	})
	return file_proto_rates_proto_rawDescData
}

var file_proto_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_rates_proto_goTypes = []any{
	(*GetRatesRequest)(nil),     // 0: rates.GetRatesRequest
	(*GetRatesResponse)(nil),    // 1: rates.GetRatesResponse
	(*HealthCheckRequest)(nil),  // 2: rates.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 3: rates.HealthCheckResponse
}
var file_proto_rates_proto_depIdxs = []int32{
	0, // 0: rates.RateService.GetRates:input_type -> rates.GetRatesRequest
	2, // 1: rates.RateService.HealthCheck:input_type -> rates.HealthCheckRequest
	1, // 2: rates.RateService.GetRates:output_type -> rates.GetRatesResponse
	3, // 3: rates.RateService.HealthCheck:output_type -> rates.HealthCheckResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rates_proto_init() }
func file_proto_rates_proto_init() {
	if File_proto_rates_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_rates_proto_rawDesc), len(file_proto_rates_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rates_proto_goTypes,
		DependencyIndexes: file_proto_rates_proto_depIdxs,
		MessageInfos:      file_proto_rates_proto_msgTypes,
	}.Build()
	File_proto_rates_proto = out.File
	file_proto_rates_proto_goTypes = nil
	file_proto_rates_proto_depIdxs = nil
}
