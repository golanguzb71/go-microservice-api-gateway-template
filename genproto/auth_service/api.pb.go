// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: api.proto

package auth_service

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

type API struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Slug      string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug"`
	Method    string `protobuf:"bytes,3,opt,name=method,proto3" json:"method"`
	Path      string `protobuf:"bytes,4,opt,name=path,proto3" json:"path"`
	IsPublic  bool   `protobuf:"varint,5,opt,name=is_public,json=isPublic,proto3" json:"is_public"`
	HasAccess bool   `protobuf:"varint,6,opt,name=has_access,json=hasAccess,proto3" json:"has_access"`
}

func (x *API) Reset() {
	*x = API{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *API) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*API) ProtoMessage() {}

func (x *API) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use API.ProtoReflect.Descriptor instead.
func (*API) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *API) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *API) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *API) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *API) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *API) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *API) GetHasAccess() bool {
	if x != nil {
		return x.HasAccess
	}
	return false
}

type APIList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Items []*API `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *APIList) Reset() {
	*x = APIList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIList) ProtoMessage() {}

func (x *APIList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIList.ProtoReflect.Descriptor instead.
func (*APIList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *APIList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *APIList) GetItems() []*API {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpsertAPIPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	ApiId  string `protobuf:"bytes,2,opt,name=api_id,json=apiId,proto3" json:"api_id"`
}

func (x *UpsertAPIPermissionRequest) Reset() {
	*x = UpsertAPIPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAPIPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAPIPermissionRequest) ProtoMessage() {}

func (x *UpsertAPIPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAPIPermissionRequest.ProtoReflect.Descriptor instead.
func (*UpsertAPIPermissionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertAPIPermissionRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UpsertAPIPermissionRequest) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

type UpsertAPIPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOn bool `protobuf:"varint,1,opt,name=is_on,json=isOn,proto3" json:"is_on"`
}

func (x *UpsertAPIPermissionResponse) Reset() {
	*x = UpsertAPIPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAPIPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAPIPermissionResponse) ProtoMessage() {}

func (x *UpsertAPIPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAPIPermissionResponse.ProtoReflect.Descriptor instead.
func (*UpsertAPIPermissionResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpsertAPIPermissionResponse) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

type CheckAPIPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method"`
	Path   string `protobuf:"bytes,3,opt,name=path,proto3" json:"path"`
}

func (x *CheckAPIPermissionRequest) Reset() {
	*x = CheckAPIPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAPIPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAPIPermissionRequest) ProtoMessage() {}

func (x *CheckAPIPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAPIPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckAPIPermissionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *CheckAPIPermissionRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CheckAPIPermissionRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CheckAPIPermissionRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type CheckAPIPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAllowed bool `protobuf:"varint,1,opt,name=is_allowed,json=isAllowed,proto3" json:"is_allowed"`
}

func (x *CheckAPIPermissionResponse) Reset() {
	*x = CheckAPIPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAPIPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAPIPermissionResponse) ProtoMessage() {}

func (x *CheckAPIPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAPIPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckAPIPermissionResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *CheckAPIPermissionResponse) GetIsAllowed() bool {
	if x != nil {
		return x.IsAllowed
	}
	return false
}

type UpsertSpecificAPIPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId               string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	SpecificPermissionId string `protobuf:"bytes,2,opt,name=specific_permission_id,json=specificPermissionId,proto3" json:"specific_permission_id"`
}

func (x *UpsertSpecificAPIPermissionRequest) Reset() {
	*x = UpsertSpecificAPIPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertSpecificAPIPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertSpecificAPIPermissionRequest) ProtoMessage() {}

func (x *UpsertSpecificAPIPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertSpecificAPIPermissionRequest.ProtoReflect.Descriptor instead.
func (*UpsertSpecificAPIPermissionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *UpsertSpecificAPIPermissionRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UpsertSpecificAPIPermissionRequest) GetSpecificPermissionId() string {
	if x != nil {
		return x.SpecificPermissionId
	}
	return ""
}

type CheckSpecificAPIPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAllowed bool `protobuf:"varint,1,opt,name=is_allowed,json=isAllowed,proto3" json:"is_allowed"`
}

func (x *CheckSpecificAPIPermissionResponse) Reset() {
	*x = CheckSpecificAPIPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckSpecificAPIPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckSpecificAPIPermissionResponse) ProtoMessage() {}

func (x *CheckSpecificAPIPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckSpecificAPIPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckSpecificAPIPermissionResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *CheckSpecificAPIPermissionResponse) GetIsAllowed() bool {
	if x != nil {
		return x.IsAllowed
	}
	return false
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x03, 0x41, 0x50,
	0x49, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x1d,
	0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x48, 0x0a,
	0x07, 0x41, 0x50, 0x49, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x50, 0x49,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x1a, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x41, 0x50, 0x49, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x69, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x1b, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41,
	0x50, 0x49, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x22, 0x60, 0x0a, 0x19, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x50, 0x49, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x3b, 0x0a, 0x1a, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x50, 0x49, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x73, 0x0a, 0x22, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x41, 0x50, 0x49, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x43, 0x0a,
	0x22, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x41, 0x50,
	0x49, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []any{
	(*API)(nil),                                // 0: auth_service.API
	(*APIList)(nil),                            // 1: auth_service.APIList
	(*UpsertAPIPermissionRequest)(nil),         // 2: auth_service.UpsertAPIPermissionRequest
	(*UpsertAPIPermissionResponse)(nil),        // 3: auth_service.UpsertAPIPermissionResponse
	(*CheckAPIPermissionRequest)(nil),          // 4: auth_service.CheckAPIPermissionRequest
	(*CheckAPIPermissionResponse)(nil),         // 5: auth_service.CheckAPIPermissionResponse
	(*UpsertSpecificAPIPermissionRequest)(nil), // 6: auth_service.UpsertSpecificAPIPermissionRequest
	(*CheckSpecificAPIPermissionResponse)(nil), // 7: auth_service.CheckSpecificAPIPermissionResponse
}
var file_api_proto_depIdxs = []int32{
	0, // 0: auth_service.APIList.items:type_name -> auth_service.API
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*API); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*APIList); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertAPIPermissionRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertAPIPermissionResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CheckAPIPermissionRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CheckAPIPermissionResponse); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertSpecificAPIPermissionRequest); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CheckSpecificAPIPermissionResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
