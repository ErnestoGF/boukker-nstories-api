// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: stories.proto

package storiespb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestChangeCover_Dst int32

const (
	RequestChangeCover_IGNORE  RequestChangeCover_Dst = 0
	RequestChangeCover_STORY   RequestChangeCover_Dst = 1
	RequestChangeCover_CHAPTER RequestChangeCover_Dst = 2
)

// Enum value maps for RequestChangeCover_Dst.
var (
	RequestChangeCover_Dst_name = map[int32]string{
		0: "IGNORE",
		1: "STORY",
		2: "CHAPTER",
	}
	RequestChangeCover_Dst_value = map[string]int32{
		"IGNORE":  0,
		"STORY":   1,
		"CHAPTER": 2,
	}
)

func (x RequestChangeCover_Dst) Enum() *RequestChangeCover_Dst {
	p := new(RequestChangeCover_Dst)
	*p = x
	return p
}

func (x RequestChangeCover_Dst) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestChangeCover_Dst) Descriptor() protoreflect.EnumDescriptor {
	return file_stories_proto_enumTypes[0].Descriptor()
}

func (RequestChangeCover_Dst) Type() protoreflect.EnumType {
	return &file_stories_proto_enumTypes[0]
}

func (x RequestChangeCover_Dst) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestChangeCover_Dst.Descriptor instead.
func (RequestChangeCover_Dst) EnumDescriptor() ([]byte, []int) {
	return file_stories_proto_rawDescGZIP(), []int{5, 0}
}

// RequestWriteStory ...
type RequestWriteStory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title              string                `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Author             string                `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
	IsMy               bool                  `protobuf:"varint,3,opt,name=IsMy,proto3" json:"IsMy,omitempty"`
	Tags               []string              `protobuf:"bytes,4,rep,name=Tags,proto3" json:"Tags,omitempty"`
	CategoryID         string                `protobuf:"bytes,5,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	LanguageID         string                `protobuf:"bytes,6,opt,name=LanguageID,proto3" json:"LanguageID,omitempty"`
	TokenCover         string                `protobuf:"bytes,7,opt,name=TokenCover,proto3" json:"TokenCover,omitempty"`
	Characters         []string              `protobuf:"bytes,8,rep,name=Characters,proto3" json:"Characters,omitempty"`
	Audience           AudienceType          `protobuf:"varint,9,opt,name=Audience,proto3,enum=stories.AudienceType" json:"Audience,omitempty"`
	Copyright          CopyrightType         `protobuf:"varint,10,opt,name=Copyright,proto3,enum=stories.CopyrightType" json:"Copyright,omitempty"`
	Clasification      ClasificationType     `protobuf:"varint,11,opt,name=Clasification,proto3,enum=stories.ClasificationType" json:"Clasification,omitempty"`
	CommentsModeration CommentModerationType `protobuf:"varint,12,opt,name=CommentsModeration,proto3,enum=stories.CommentModerationType" json:"CommentsModeration,omitempty"`
	Description        string                `protobuf:"bytes,13,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *RequestWriteStory) Reset() {
	*x = RequestWriteStory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestWriteStory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestWriteStory) ProtoMessage() {}

func (x *RequestWriteStory) ProtoReflect() protoreflect.Message {
	mi := &file_stories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestWriteStory.ProtoReflect.Descriptor instead.
func (*RequestWriteStory) Descriptor() ([]byte, []int) {
	return file_stories_proto_rawDescGZIP(), []int{0}
}

func (x *RequestWriteStory) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RequestWriteStory) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *RequestWriteStory) GetIsMy() bool {
	if x != nil {
		return x.IsMy
	}
	return false
}

func (x *RequestWriteStory) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RequestWriteStory) GetCategoryID() string {
	if x != nil {
		return x.CategoryID
	}
	return ""
}

func (x *RequestWriteStory) GetLanguageID() string {
	if x != nil {
		return x.LanguageID
	}
	return ""
}

func (x *RequestWriteStory) GetTokenCover() string {
	if x != nil {
		return x.TokenCover
	}
	return ""
}

func (x *RequestWriteStory) GetCharacters() []string {
	if x != nil {
		return x.Characters
	}
	return nil
}

func (x *RequestWriteStory) GetAudience() AudienceType {
	if x != nil {
		return x.Audience
	}
	return AudienceType_AT_IGNORE
}

func (x *RequestWriteStory) GetCopyright() CopyrightType {
	if x != nil {
		return x.Copyright
	}
	return CopyrightType_CRT_IGNORE
}

func (x *RequestWriteStory) GetClasification() ClasificationType {
	if x != nil {
		return x.Clasification
	}
	return ClasificationType_CT_IGNORE
}

func (x *RequestWriteStory) GetCommentsModeration() CommentModerationType {
	if x != nil {
		return x.CommentsModeration
	}
	return CommentModerationType_CMT_IGNORE
}

func (x *RequestWriteStory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// RequestEditStory ...
type RequestEditStory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string                `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title              string                `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Tags               []string              `protobuf:"bytes,3,rep,name=Tags,proto3" json:"Tags,omitempty"`
	CategoryID         string                `protobuf:"bytes,4,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	Characters         []string              `protobuf:"bytes,5,rep,name=Characters,proto3" json:"Characters,omitempty"`
	Audience           AudienceType          `protobuf:"varint,6,opt,name=Audience,proto3,enum=stories.AudienceType" json:"Audience,omitempty"`
	Copyright          CopyrightType         `protobuf:"varint,7,opt,name=Copyright,proto3,enum=stories.CopyrightType" json:"Copyright,omitempty"`
	Clasification      ClasificationType     `protobuf:"varint,8,opt,name=Clasification,proto3,enum=stories.ClasificationType" json:"Clasification,omitempty"`
	CommentsModeration CommentModerationType `protobuf:"varint,9,opt,name=CommentsModeration,proto3,enum=stories.CommentModerationType" json:"CommentsModeration,omitempty"`
	Description        string                `protobuf:"bytes,10,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *RequestEditStory) Reset() {
	*x = RequestEditStory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stories_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestEditStory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestEditStory) ProtoMessage() {}

func (x *RequestEditStory) ProtoReflect() protoreflect.Message {
	mi := &file_stories_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestEditStory.ProtoReflect.Descriptor instead.
func (*RequestEditStory) Descriptor() ([]byte, []int) {
	return file_stories_proto_rawDescGZIP(), []int{1}
}

func (x *RequestEditStory) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RequestEditStory) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RequestEditStory) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RequestEditStory) GetCategoryID() string {
	if x != nil {
		return x.CategoryID
	}
	return ""
}

func (x *RequestEditStory) GetCharacters() []string {
	if x != nil {
		return x.Characters
	}
	return nil
}

func (x *RequestEditStory) GetAudience() AudienceType {
	if x != nil {
		return x.Audience
	}
	return AudienceType_AT_IGNORE
}

func (x *RequestEditStory) GetCopyright() CopyrightType {
	if x != nil {
		return x.Copyright
	}
	return CopyrightType_CRT_IGNORE
}

func (x *RequestEditStory) GetClasification() ClasificationType {
	if x != nil {
		return x.Clasification
	}
	return ClasificationType_CT_IGNORE
}

func (x *RequestEditStory) GetCommentsModeration() CommentModerationType {
	if x != nil {
		return x.CommentsModeration
	}
	return CommentModerationType_CMT_IGNORE
}

func (x *RequestEditStory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// RequestChangeFinished
type RequestChangeFinished struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Finished bool `protobuf:"varint,1,opt,name=Finished,proto3" json:"Finished,omitempty"`
}

func (x *RequestChangeFinished) Reset() {
	*x = RequestChangeFinished{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stories_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChangeFinished) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChangeFinished) ProtoMessage() {}

func (x *RequestChangeFinished) ProtoReflect() protoreflect.Message {
	mi := &file_stories_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChangeFinished.ProtoReflect.Descriptor instead.
func (*RequestChangeFinished) Descriptor() ([]byte, []int) {
	return file_stories_proto_rawDescGZIP(), []int{2}
}

func (x *RequestChangeFinished) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

// RequestChangeStoryStatus ...
type RequestChangeStoryStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string     `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ChaptersID []string   `protobuf:"bytes,2,rep,name=ChaptersID,proto3" json:"ChaptersID,omitempty"`
	Status     StatusType `protobuf:"varint,3,opt,name=Status,proto3,enum=stories.StatusType" json:"Status,omitempty"`
}

func (x *RequestChangeStoryStatus) Reset() {
	*x = RequestChangeStoryStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stories_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChangeStoryStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChangeStoryStatus) ProtoMessage() {}

func (x *RequestChangeStoryStatus) ProtoReflect() protoreflect.Message {
	mi := &file_stories_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChangeStoryStatus.ProtoReflect.Descriptor instead.
func (*RequestChangeStoryStatus) Descriptor() ([]byte, []int) {
	return file_stories_proto_rawDescGZIP(), []int{3}
}

func (x *RequestChangeStoryStatus) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RequestChangeStoryStatus) GetChaptersID() []string {
	if x != nil {
		return x.ChaptersID
	}
	return nil
}

func (x *RequestChangeStoryStatus) GetStatus() StatusType {
	if x != nil {
		return x.Status
	}
	return StatusType_ST_IGNORE
}

// RequestRemoveStory ...
type RequestRemoveStory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RequestRemoveStory) Reset() {
	*x = RequestRemoveStory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stories_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRemoveStory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRemoveStory) ProtoMessage() {}

func (x *RequestRemoveStory) ProtoReflect() protoreflect.Message {
	mi := &file_stories_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRemoveStory.ProtoReflect.Descriptor instead.
func (*RequestRemoveStory) Descriptor() ([]byte, []int) {
	return file_stories_proto_rawDescGZIP(), []int{4}
}

func (x *RequestRemoveStory) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// RequestChangeCover cambiar el cover de la historia o el capitulo segun el destino del cover
type RequestChangeCover struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Token       string                 `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Destination RequestChangeCover_Dst `protobuf:"varint,3,opt,name=Destination,proto3,enum=stories.RequestChangeCover_Dst" json:"Destination,omitempty"`
}

func (x *RequestChangeCover) Reset() {
	*x = RequestChangeCover{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stories_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestChangeCover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestChangeCover) ProtoMessage() {}

func (x *RequestChangeCover) ProtoReflect() protoreflect.Message {
	mi := &file_stories_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestChangeCover.ProtoReflect.Descriptor instead.
func (*RequestChangeCover) Descriptor() ([]byte, []int) {
	return file_stories_proto_rawDescGZIP(), []int{5}
}

func (x *RequestChangeCover) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RequestChangeCover) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RequestChangeCover) GetDestination() RequestChangeCover_Dst {
	if x != nil {
		return x.Destination
	}
	return RequestChangeCover_IGNORE
}

var File_stories_proto protoreflect.FileDescriptor

var file_stories_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x73, 0x4d, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x49, 0x73, 0x4d, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x31, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x6c, 0x61,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x43, 0x6c,
	0x61, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x12, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa9, 0x03,
	0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x08,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x09, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x70,
	0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x43, 0x6f, 0x70, 0x79,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x43, 0x6c, 0x61, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x6f, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x15, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x77,
	0x0a, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x49, 0x44, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0xa8, 0x01,
	0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43,
	0x6f, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x73, 0x74,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x29, 0x0a,
	0x03, 0x44, 0x73, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x48, 0x41, 0x50, 0x54, 0x45, 0x52, 0x10, 0x02, 0x32, 0x92, 0x06, 0x0a, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x13,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x09, 0x45, 0x64, 0x69, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x49, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x41, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x12, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x1a, 0x13,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x0f, 0x52, 0x65, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x72, 0x6e, 0x65,
	0x73, 0x74, 0x6f, 0x47, 0x46, 0x2f, 0x62, 0x6f, 0x75, 0x6b, 0x6b, 0x65, 0x72, 0x2d, 0x6e, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stories_proto_rawDescOnce sync.Once
	file_stories_proto_rawDescData = file_stories_proto_rawDesc
)

func file_stories_proto_rawDescGZIP() []byte {
	file_stories_proto_rawDescOnce.Do(func() {
		file_stories_proto_rawDescData = protoimpl.X.CompressGZIP(file_stories_proto_rawDescData)
	})
	return file_stories_proto_rawDescData
}

var file_stories_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stories_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stories_proto_goTypes = []interface{}{
	(RequestChangeCover_Dst)(0),       // 0: stories.RequestChangeCover.Dst
	(*RequestWriteStory)(nil),         // 1: stories.RequestWriteStory
	(*RequestEditStory)(nil),          // 2: stories.RequestEditStory
	(*RequestChangeFinished)(nil),     // 3: stories.RequestChangeFinished
	(*RequestChangeStoryStatus)(nil),  // 4: stories.RequestChangeStoryStatus
	(*RequestRemoveStory)(nil),        // 5: stories.RequestRemoveStory
	(*RequestChangeCover)(nil),        // 6: stories.RequestChangeCover
	(AudienceType)(0),                 // 7: stories.AudienceType
	(CopyrightType)(0),                // 8: stories.CopyrightType
	(ClasificationType)(0),            // 9: stories.ClasificationType
	(CommentModerationType)(0),        // 10: stories.CommentModerationType
	(StatusType)(0),                   // 11: stories.StatusType
	(*RequestWriteChapter)(nil),       // 12: stories.RequestWriteChapter
	(*RequestEditChapter)(nil),        // 13: stories.RequestEditChapter
	(*RequestReorderChapters)(nil),    // 14: stories.RequestReorderChapters
	(*RequestRemoveChapter)(nil),      // 15: stories.RequestRemoveChapter
	(*RequestRemoveChapterCover)(nil), // 16: stories.RequestRemoveChapterCover
	(*ResponseID)(nil),                // 17: stories.ResponseID
	(*emptypb.Empty)(nil),             // 18: google.protobuf.Empty
}
var file_stories_proto_depIdxs = []int32{
	7,  // 0: stories.RequestWriteStory.Audience:type_name -> stories.AudienceType
	8,  // 1: stories.RequestWriteStory.Copyright:type_name -> stories.CopyrightType
	9,  // 2: stories.RequestWriteStory.Clasification:type_name -> stories.ClasificationType
	10, // 3: stories.RequestWriteStory.CommentsModeration:type_name -> stories.CommentModerationType
	7,  // 4: stories.RequestEditStory.Audience:type_name -> stories.AudienceType
	8,  // 5: stories.RequestEditStory.Copyright:type_name -> stories.CopyrightType
	9,  // 6: stories.RequestEditStory.Clasification:type_name -> stories.ClasificationType
	10, // 7: stories.RequestEditStory.CommentsModeration:type_name -> stories.CommentModerationType
	11, // 8: stories.RequestChangeStoryStatus.Status:type_name -> stories.StatusType
	0,  // 9: stories.RequestChangeCover.Destination:type_name -> stories.RequestChangeCover.Dst
	1,  // 10: stories.stories.WriteStory:input_type -> stories.RequestWriteStory
	5,  // 11: stories.stories.RemoveStory:input_type -> stories.RequestRemoveStory
	2,  // 12: stories.stories.EditStory:input_type -> stories.RequestEditStory
	3,  // 13: stories.stories.ChangeFinished:input_type -> stories.RequestChangeFinished
	4,  // 14: stories.stories.ChangeStatus:input_type -> stories.RequestChangeStoryStatus
	6,  // 15: stories.stories.ChangeCover:input_type -> stories.RequestChangeCover
	12, // 16: stories.stories.WriteChapter:input_type -> stories.RequestWriteChapter
	13, // 17: stories.stories.EditChapter:input_type -> stories.RequestEditChapter
	14, // 18: stories.stories.ReorderChapters:input_type -> stories.RequestReorderChapters
	15, // 19: stories.stories.RemoveChapter:input_type -> stories.RequestRemoveChapter
	16, // 20: stories.stories.RemoveChapterCover:input_type -> stories.RequestRemoveChapterCover
	17, // 21: stories.stories.WriteStory:output_type -> stories.ResponseID
	18, // 22: stories.stories.RemoveStory:output_type -> google.protobuf.Empty
	18, // 23: stories.stories.EditStory:output_type -> google.protobuf.Empty
	18, // 24: stories.stories.ChangeFinished:output_type -> google.protobuf.Empty
	18, // 25: stories.stories.ChangeStatus:output_type -> google.protobuf.Empty
	18, // 26: stories.stories.ChangeCover:output_type -> google.protobuf.Empty
	17, // 27: stories.stories.WriteChapter:output_type -> stories.ResponseID
	18, // 28: stories.stories.EditChapter:output_type -> google.protobuf.Empty
	18, // 29: stories.stories.ReorderChapters:output_type -> google.protobuf.Empty
	18, // 30: stories.stories.RemoveChapter:output_type -> google.protobuf.Empty
	18, // 31: stories.stories.RemoveChapterCover:output_type -> google.protobuf.Empty
	21, // [21:32] is the sub-list for method output_type
	10, // [10:21] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_stories_proto_init() }
func file_stories_proto_init() {
	if File_stories_proto != nil {
		return
	}
	file_commons_proto_init()
	file_chapters_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stories_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestWriteStory); i {
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
		file_stories_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestEditStory); i {
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
		file_stories_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChangeFinished); i {
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
		file_stories_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChangeStoryStatus); i {
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
		file_stories_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRemoveStory); i {
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
		file_stories_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestChangeCover); i {
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
			RawDescriptor: file_stories_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stories_proto_goTypes,
		DependencyIndexes: file_stories_proto_depIdxs,
		EnumInfos:         file_stories_proto_enumTypes,
		MessageInfos:      file_stories_proto_msgTypes,
	}.Build()
	File_stories_proto = out.File
	file_stories_proto_rawDesc = nil
	file_stories_proto_goTypes = nil
	file_stories_proto_depIdxs = nil
}
