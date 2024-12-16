// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: decoder.proto

package decoder

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

type GeneralCosmosMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Signers []string `protobuf:"bytes,3,rep,name=signers,proto3" json:"signers,omitempty"`
}

func (x *GeneralCosmosMsg) Reset() {
	*x = GeneralCosmosMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_decoder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralCosmosMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralCosmosMsg) ProtoMessage() {}

func (x *GeneralCosmosMsg) ProtoReflect() protoreflect.Message {
	mi := &file_decoder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralCosmosMsg.ProtoReflect.Descriptor instead.
func (*GeneralCosmosMsg) Descriptor() ([]byte, []int) {
	return file_decoder_proto_rawDescGZIP(), []int{0}
}

func (x *GeneralCosmosMsg) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GeneralCosmosMsg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GeneralCosmosMsg) GetSigners() []string {
	if x != nil {
		return x.Signers
	}
	return nil
}

type DecodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxByte []byte `protobuf:"bytes,1,opt,name=txByte,proto3" json:"txByte,omitempty"`
}

func (x *DecodeRequest) Reset() {
	*x = DecodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_decoder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeRequest) ProtoMessage() {}

func (x *DecodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_decoder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeRequest.ProtoReflect.Descriptor instead.
func (*DecodeRequest) Descriptor() ([]byte, []int) {
	return file_decoder_proto_rawDescGZIP(), []int{1}
}

func (x *DecodeRequest) GetTxByte() []byte {
	if x != nil {
		return x.TxByte
	}
	return nil
}

type DecodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgs     []*GeneralCosmosMsg `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	TxResult string              `protobuf:"bytes,2,opt,name=txResult,proto3" json:"txResult,omitempty"`
}

func (x *DecodeResponse) Reset() {
	*x = DecodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_decoder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeResponse) ProtoMessage() {}

func (x *DecodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_decoder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeResponse.ProtoReflect.Descriptor instead.
func (*DecodeResponse) Descriptor() ([]byte, []int) {
	return file_decoder_proto_rawDescGZIP(), []int{2}
}

func (x *DecodeResponse) GetMsgs() []*GeneralCosmosMsg {
	if x != nil {
		return x.Msgs
	}
	return nil
}

func (x *DecodeResponse) GetTxResult() string {
	if x != nil {
		return x.TxResult
	}
	return ""
}

var File_decoder_proto protoreflect.FileDescriptor

var file_decoder_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x22, 0x5a, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x73, 0x22, 0x27, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x22, 0x5b, 0x0a,
	0x0e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x4d, 0x73, 0x67, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x4c, 0x0a, 0x0d, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x44,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2d, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x32, 0x38, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x64, 0x2f, 0x63, 0x6d, 0x64,
	0x2f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_decoder_proto_rawDescOnce sync.Once
	file_decoder_proto_rawDescData = file_decoder_proto_rawDesc
)

func file_decoder_proto_rawDescGZIP() []byte {
	file_decoder_proto_rawDescOnce.Do(func() {
		file_decoder_proto_rawDescData = protoimpl.X.CompressGZIP(file_decoder_proto_rawDescData)
	})
	return file_decoder_proto_rawDescData
}

var file_decoder_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_decoder_proto_goTypes = []interface{}{
	(*GeneralCosmosMsg)(nil), // 0: decoder.GeneralCosmosMsg
	(*DecodeRequest)(nil),    // 1: decoder.DecodeRequest
	(*DecodeResponse)(nil),   // 2: decoder.DecodeResponse
}
var file_decoder_proto_depIdxs = []int32{
	0, // 0: decoder.DecodeResponse.msgs:type_name -> decoder.GeneralCosmosMsg
	1, // 1: decoder.CosmosDecoder.Decode:input_type -> decoder.DecodeRequest
	2, // 2: decoder.CosmosDecoder.Decode:output_type -> decoder.DecodeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_decoder_proto_init() }
func file_decoder_proto_init() {
	if File_decoder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_decoder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralCosmosMsg); i {
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
		file_decoder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeRequest); i {
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
		file_decoder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeResponse); i {
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
			RawDescriptor: file_decoder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_decoder_proto_goTypes,
		DependencyIndexes: file_decoder_proto_depIdxs,
		MessageInfos:      file_decoder_proto_msgTypes,
	}.Build()
	File_decoder_proto = out.File
	file_decoder_proto_rawDesc = nil
	file_decoder_proto_goTypes = nil
	file_decoder_proto_depIdxs = nil
}
