// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: category_message.proto

package product_service

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

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Img         string `protobuf:"bytes,4,opt,name=img,proto3" json:"img,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_category_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_category_message_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Category) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *Category) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Category) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Category) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type CategoryCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryName string `protobuf:"bytes,1,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Img          string `protobuf:"bytes,3,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *CategoryCreateReq) Reset() {
	*x = CategoryCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryCreateReq) ProtoMessage() {}

func (x *CategoryCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_category_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryCreateReq.ProtoReflect.Descriptor instead.
func (*CategoryCreateReq) Descriptor() ([]byte, []int) {
	return file_category_message_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryCreateReq) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *CategoryCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CategoryCreateReq) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

type CategoryGetListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int32       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Category []*Category `protobuf:"bytes,2,rep,name=category,proto3" json:"category,omitempty"`
}

func (x *CategoryGetListResp) Reset() {
	*x = CategoryGetListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryGetListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryGetListResp) ProtoMessage() {}

func (x *CategoryGetListResp) ProtoReflect() protoreflect.Message {
	mi := &file_category_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryGetListResp.ProtoReflect.Descriptor instead.
func (*CategoryGetListResp) Descriptor() ([]byte, []int) {
	return file_category_message_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryGetListResp) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CategoryGetListResp) GetCategory() []*Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type CategoryUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryName string `protobuf:"bytes,1,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Id           string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CategoryUpdateReq) Reset() {
	*x = CategoryUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpdateReq) ProtoMessage() {}

func (x *CategoryUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_category_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpdateReq.ProtoReflect.Descriptor instead.
func (*CategoryUpdateReq) Descriptor() ([]byte, []int) {
	return file_category_message_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryUpdateReq) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *CategoryUpdateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CategoryUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SubCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId  string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Img         string `protobuf:"bytes,5,opt,name=img,proto3" json:"img,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *SubCategory) Reset() {
	*x = SubCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCategory) ProtoMessage() {}

func (x *SubCategory) ProtoReflect() protoreflect.Message {
	mi := &file_category_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCategory.ProtoReflect.Descriptor instead.
func (*SubCategory) Descriptor() ([]byte, []int) {
	return file_category_message_proto_rawDescGZIP(), []int{4}
}

func (x *SubCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubCategory) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *SubCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubCategory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SubCategory) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *SubCategory) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SubCategory) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *SubCategory) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type SubCategoryCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryName string `protobuf:"bytes,1,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Img          string `protobuf:"bytes,3,opt,name=img,proto3" json:"img,omitempty"`
	CategoryId   string `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *SubCategoryCreateReq) Reset() {
	*x = SubCategoryCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCategoryCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCategoryCreateReq) ProtoMessage() {}

func (x *SubCategoryCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_category_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCategoryCreateReq.ProtoReflect.Descriptor instead.
func (*SubCategoryCreateReq) Descriptor() ([]byte, []int) {
	return file_category_message_proto_rawDescGZIP(), []int{5}
}

func (x *SubCategoryCreateReq) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *SubCategoryCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SubCategoryCreateReq) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *SubCategoryCreateReq) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type SubCategoryGetListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int32          `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Subcategory []*SubCategory `protobuf:"bytes,2,rep,name=subcategory,proto3" json:"subcategory,omitempty"`
}

func (x *SubCategoryGetListResp) Reset() {
	*x = SubCategoryGetListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCategoryGetListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCategoryGetListResp) ProtoMessage() {}

func (x *SubCategoryGetListResp) ProtoReflect() protoreflect.Message {
	mi := &file_category_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCategoryGetListResp.ProtoReflect.Descriptor instead.
func (*SubCategoryGetListResp) Descriptor() ([]byte, []int) {
	return file_category_message_proto_rawDescGZIP(), []int{6}
}

func (x *SubCategoryGetListResp) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SubCategoryGetListResp) GetSubcategory() []*SubCategory {
	if x != nil {
		return x.Subcategory
	}
	return nil
}

type SubCategoryUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryName string `protobuf:"bytes,2,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CategoryId   string `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *SubCategoryUpdateReq) Reset() {
	*x = SubCategoryUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCategoryUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCategoryUpdateReq) ProtoMessage() {}

func (x *SubCategoryUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_category_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCategoryUpdateReq.ProtoReflect.Descriptor instead.
func (*SubCategoryUpdateReq) Descriptor() ([]byte, []int) {
	return file_category_message_proto_rawDescGZIP(), []int{7}
}

func (x *SubCategoryUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubCategoryUpdateReq) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *SubCategoryUpdateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SubCategoryUpdateReq) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

var File_category_message_proto protoreflect.FileDescriptor

var file_category_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x08, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x6d, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6c, 0x0a, 0x11, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x22, 0x62, 0x0a, 0x13, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x6a, 0x0a,
	0x11, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xe3, 0x01, 0x0a, 0x0b, 0x53, 0x75,
	0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x6d, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x90, 0x01, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d,
	0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x22, 0x6e, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_category_message_proto_rawDescOnce sync.Once
	file_category_message_proto_rawDescData = file_category_message_proto_rawDesc
)

func file_category_message_proto_rawDescGZIP() []byte {
	file_category_message_proto_rawDescOnce.Do(func() {
		file_category_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_category_message_proto_rawDescData)
	})
	return file_category_message_proto_rawDescData
}

var file_category_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_category_message_proto_goTypes = []any{
	(*Category)(nil),               // 0: product_service.Category
	(*CategoryCreateReq)(nil),      // 1: product_service.CategoryCreateReq
	(*CategoryGetListResp)(nil),    // 2: product_service.CategoryGetListResp
	(*CategoryUpdateReq)(nil),      // 3: product_service.CategoryUpdateReq
	(*SubCategory)(nil),            // 4: product_service.SubCategory
	(*SubCategoryCreateReq)(nil),   // 5: product_service.SubCategoryCreateReq
	(*SubCategoryGetListResp)(nil), // 6: product_service.SubCategoryGetListResp
	(*SubCategoryUpdateReq)(nil),   // 7: product_service.SubCategoryUpdateReq
}
var file_category_message_proto_depIdxs = []int32{
	0, // 0: product_service.CategoryGetListResp.category:type_name -> product_service.Category
	4, // 1: product_service.SubCategoryGetListResp.subcategory:type_name -> product_service.SubCategory
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_category_message_proto_init() }
func file_category_message_proto_init() {
	if File_category_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_category_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Category); i {
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
		file_category_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryCreateReq); i {
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
		file_category_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryGetListResp); i {
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
		file_category_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryUpdateReq); i {
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
		file_category_message_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SubCategory); i {
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
		file_category_message_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SubCategoryCreateReq); i {
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
		file_category_message_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SubCategoryGetListResp); i {
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
		file_category_message_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SubCategoryUpdateReq); i {
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
			RawDescriptor: file_category_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_category_message_proto_goTypes,
		DependencyIndexes: file_category_message_proto_depIdxs,
		MessageInfos:      file_category_message_proto_msgTypes,
	}.Build()
	File_category_message_proto = out.File
	file_category_message_proto_rawDesc = nil
	file_category_message_proto_goTypes = nil
	file_category_message_proto_depIdxs = nil
}
