// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: notification_models.proto

package notification_service

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_notification_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_notification_models_proto_rawDescGZIP(), []int{0}
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Timezone string `protobuf:"bytes,2,opt,name=timezone,proto3" json:"timezone"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_notification_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_notification_models_proto_rawDescGZIP(), []int{1}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Id) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

type GetListFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32             `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Limit    int32             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	Timezone string            `protobuf:"bytes,3,opt,name=timezone,proto3" json:"timezone"`
	Filters  []*Filters        `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters"`
	Sorts    []*SortBy         `protobuf:"bytes,5,rep,name=sorts,proto3" json:"sorts"`
	Extra    map[string]string `protobuf:"bytes,6,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetListFilter) Reset() {
	*x = GetListFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListFilter) ProtoMessage() {}

func (x *GetListFilter) ProtoReflect() protoreflect.Message {
	mi := &file_notification_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListFilter.ProtoReflect.Descriptor instead.
func (*GetListFilter) Descriptor() ([]byte, []int) {
	return file_notification_models_proto_rawDescGZIP(), []int{2}
}

func (x *GetListFilter) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListFilter) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *GetListFilter) GetFilters() []*Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *GetListFilter) GetSorts() []*SortBy {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *GetListFilter) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
}

func (x *Filters) Reset() {
	*x = Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_notification_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_notification_models_proto_rawDescGZIP(), []int{3}
}

func (x *Filters) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Filters) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Filters) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SortBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
}

func (x *SortBy) Reset() {
	*x = SortBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortBy) ProtoMessage() {}

func (x *SortBy) ProtoReflect() protoreflect.Message {
	mi := &file_notification_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortBy.ProtoReflect.Descriptor instead.
func (*SortBy) Descriptor() ([]byte, []int) {
	return file_notification_models_proto_rawDescGZIP(), []int{4}
}

func (x *SortBy) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SortBy) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_notification_models_proto protoreflect.FileDescriptor

var file_notification_models_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x30, 0x0a, 0x02, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0xc2, 0x02, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a,
	0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x44, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x49, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x32, 0x0a, 0x06,
	0x53, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_models_proto_rawDescOnce sync.Once
	file_notification_models_proto_rawDescData = file_notification_models_proto_rawDesc
)

func file_notification_models_proto_rawDescGZIP() []byte {
	file_notification_models_proto_rawDescOnce.Do(func() {
		file_notification_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_models_proto_rawDescData)
	})
	return file_notification_models_proto_rawDescData
}

var file_notification_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_notification_models_proto_goTypes = []any{
	(*Empty)(nil),         // 0: notification_service.Empty
	(*Id)(nil),            // 1: notification_service.Id
	(*GetListFilter)(nil), // 2: notification_service.GetListFilter
	(*Filters)(nil),       // 3: notification_service.Filters
	(*SortBy)(nil),        // 4: notification_service.SortBy
	nil,                   // 5: notification_service.GetListFilter.ExtraEntry
}
var file_notification_models_proto_depIdxs = []int32{
	3, // 0: notification_service.GetListFilter.filters:type_name -> notification_service.Filters
	4, // 1: notification_service.GetListFilter.sorts:type_name -> notification_service.SortBy
	5, // 2: notification_service.GetListFilter.extra:type_name -> notification_service.GetListFilter.ExtraEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notification_models_proto_init() }
func file_notification_models_proto_init() {
	if File_notification_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_models_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_notification_models_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Id); i {
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
		file_notification_models_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetListFilter); i {
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
		file_notification_models_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Filters); i {
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
		file_notification_models_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SortBy); i {
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
			RawDescriptor: file_notification_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notification_models_proto_goTypes,
		DependencyIndexes: file_notification_models_proto_depIdxs,
		MessageInfos:      file_notification_models_proto_msgTypes,
	}.Build()
	File_notification_models_proto = out.File
	file_notification_models_proto_rawDesc = nil
	file_notification_models_proto_goTypes = nil
	file_notification_models_proto_depIdxs = nil
}
