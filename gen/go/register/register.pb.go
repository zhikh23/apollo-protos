// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: register.proto

package register

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

type ResourceType int32

const (
	ResourceType_unknown ResourceType = 0
	ResourceType_file    ResourceType = 1
	ResourceType_url     ResourceType = 2
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "unknown",
		1: "file",
		2: "url",
	}
	ResourceType_value = map[string]int32{
		"unknown": 0,
		"file":    1,
		"url":     2,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_register_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_register_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{0}
}

type Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tags        []string     `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Type        ResourceType `protobuf:"varint,5,opt,name=type,proto3,enum=apollo.ResourceType" json:"type,omitempty"`
	FileId      *uint64      `protobuf:"varint,6,opt,name=file_id,json=fileId,proto3,oneof" json:"file_id,omitempty"`
	Url         *string      `protobuf:"bytes,7,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *Material) Reset() {
	*x = Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Material) ProtoMessage() {}

func (x *Material) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Material.ProtoReflect.Descriptor instead.
func (*Material) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{0}
}

func (x *Material) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Material) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Material) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Material) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Material) GetType() ResourceType {
	if x != nil {
		return x.Type
	}
	return ResourceType_unknown
}

func (x *Material) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *Material) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type RegisterResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tags        []string     `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Type        ResourceType `protobuf:"varint,4,opt,name=type,proto3,enum=apollo.ResourceType" json:"type,omitempty"`
	FileId      *uint64      `protobuf:"varint,5,opt,name=file_id,json=fileId,proto3,oneof" json:"file_id,omitempty"`
	Url         *string      `protobuf:"bytes,6,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *RegisterResourceRequest) Reset() {
	*x = RegisterResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResourceRequest) ProtoMessage() {}

func (x *RegisterResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResourceRequest.ProtoReflect.Descriptor instead.
func (*RegisterResourceRequest) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResourceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterResourceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RegisterResourceRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RegisterResourceRequest) GetType() ResourceType {
	if x != nil {
		return x.Type
	}
	return ResourceType_unknown
}

func (x *RegisterResourceRequest) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *RegisterResourceRequest) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type RegisterResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaterialId uint64 `protobuf:"varint,1,opt,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
}

func (x *RegisterResourceResponse) Reset() {
	*x = RegisterResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResourceResponse) ProtoMessage() {}

func (x *RegisterResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResourceResponse.ProtoReflect.Descriptor instead.
func (*RegisterResourceResponse) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResourceResponse) GetMaterialId() uint64 {
	if x != nil {
		return x.MaterialId
	}
	return 0
}

type GetResourceByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaterialId uint64 `protobuf:"varint,1,opt,name=material_id,json=materialId,proto3" json:"material_id,omitempty"`
}

func (x *GetResourceByIdRequest) Reset() {
	*x = GetResourceByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByIdRequest) ProtoMessage() {}

func (x *GetResourceByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByIdRequest.ProtoReflect.Descriptor instead.
func (*GetResourceByIdRequest) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{3}
}

func (x *GetResourceByIdRequest) GetMaterialId() uint64 {
	if x != nil {
		return x.MaterialId
	}
	return 0
}

type GetResourceByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Material *Material `protobuf:"bytes,1,opt,name=material,proto3" json:"material,omitempty"`
}

func (x *GetResourceByIdResponse) Reset() {
	*x = GetResourceByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceByIdResponse) ProtoMessage() {}

func (x *GetResourceByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceByIdResponse.ProtoReflect.Descriptor instead.
func (*GetResourceByIdResponse) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{4}
}

func (x *GetResourceByIdResponse) GetMaterial() *Material {
	if x != nil {
		return x.Material
	}
	return nil
}

type GetResourcesByTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GetResourcesByTagsRequest) Reset() {
	*x = GetResourcesByTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcesByTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesByTagsRequest) ProtoMessage() {}

func (x *GetResourcesByTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesByTagsRequest.ProtoReflect.Descriptor instead.
func (*GetResourcesByTagsRequest) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{5}
}

func (x *GetResourcesByTagsRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GetResourcesByTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Materials []*Material `protobuf:"bytes,1,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *GetResourcesByTagsResponse) Reset() {
	*x = GetResourcesByTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcesByTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesByTagsResponse) ProtoMessage() {}

func (x *GetResourcesByTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesByTagsResponse.ProtoReflect.Descriptor instead.
func (*GetResourcesByTagsResponse) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{6}
}

func (x *GetResourcesByTagsResponse) GetMaterials() []*Material {
	if x != nil {
		return x.Materials
	}
	return nil
}

var File_register_proto protoreflect.FileDescriptor

var file_register_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75,
	0x72, 0x6c, 0x22, 0xd6, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x3b, 0x0a, 0x18, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x2f, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x79, 0x54, 0x61,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x4c, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x79, 0x54,
	0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2a, 0x2e, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x10, 0x02, 0x32, 0x99, 0x02, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x55, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x6f, 0x6c,
	0x6c, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x6f, 0x6c,
	0x6c, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x21, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_register_proto_rawDescOnce sync.Once
	file_register_proto_rawDescData = file_register_proto_rawDesc
)

func file_register_proto_rawDescGZIP() []byte {
	file_register_proto_rawDescOnce.Do(func() {
		file_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_proto_rawDescData)
	})
	return file_register_proto_rawDescData
}

var file_register_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_register_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_register_proto_goTypes = []interface{}{
	(ResourceType)(0),                  // 0: apollo.ResourceType
	(*Material)(nil),                   // 1: apollo.Material
	(*RegisterResourceRequest)(nil),    // 2: apollo.RegisterResourceRequest
	(*RegisterResourceResponse)(nil),   // 3: apollo.RegisterResourceResponse
	(*GetResourceByIdRequest)(nil),     // 4: apollo.GetResourceByIdRequest
	(*GetResourceByIdResponse)(nil),    // 5: apollo.GetResourceByIdResponse
	(*GetResourcesByTagsRequest)(nil),  // 6: apollo.GetResourcesByTagsRequest
	(*GetResourcesByTagsResponse)(nil), // 7: apollo.GetResourcesByTagsResponse
}
var file_register_proto_depIdxs = []int32{
	0, // 0: apollo.Material.type:type_name -> apollo.ResourceType
	0, // 1: apollo.RegisterResourceRequest.type:type_name -> apollo.ResourceType
	1, // 2: apollo.GetResourceByIdResponse.material:type_name -> apollo.Material
	1, // 3: apollo.GetResourcesByTagsResponse.materials:type_name -> apollo.Material
	2, // 4: apollo.RegisterService.RegisterResource:input_type -> apollo.RegisterResourceRequest
	4, // 5: apollo.RegisterService.GetResourceById:input_type -> apollo.GetResourceByIdRequest
	6, // 6: apollo.RegisterService.GetResourcesByTags:input_type -> apollo.GetResourcesByTagsRequest
	3, // 7: apollo.RegisterService.RegisterResource:output_type -> apollo.RegisterResourceResponse
	5, // 8: apollo.RegisterService.GetResourceById:output_type -> apollo.GetResourceByIdResponse
	7, // 9: apollo.RegisterService.GetResourcesByTags:output_type -> apollo.GetResourcesByTagsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_register_proto_init() }
func file_register_proto_init() {
	if File_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Material); i {
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
		file_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResourceRequest); i {
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
		file_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResourceResponse); i {
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
		file_register_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByIdRequest); i {
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
		file_register_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceByIdResponse); i {
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
		file_register_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcesByTagsRequest); i {
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
		file_register_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcesByTagsResponse); i {
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
	file_register_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_register_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_register_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_register_proto_goTypes,
		DependencyIndexes: file_register_proto_depIdxs,
		EnumInfos:         file_register_proto_enumTypes,
		MessageInfos:      file_register_proto_msgTypes,
	}.Build()
	File_register_proto = out.File
	file_register_proto_rawDesc = nil
	file_register_proto_goTypes = nil
	file_register_proto_depIdxs = nil
}
