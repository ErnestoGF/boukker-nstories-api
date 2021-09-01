// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: stories.proto

package storiespb

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

// RequestWriteStory ...
type RequestWriteStory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title              string                `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Author             string                `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	IsMy               bool                  `protobuf:"varint,4,opt,name=IsMy,proto3" json:"IsMy,omitempty"`
	Tags               []string              `protobuf:"bytes,5,rep,name=Tags,proto3" json:"Tags,omitempty"`
	CategoryID         string                `protobuf:"bytes,6,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	LanguageID         string                `protobuf:"bytes,7,opt,name=LanguageID,proto3" json:"LanguageID,omitempty"`
	TokenCover         string                `protobuf:"bytes,8,opt,name=TokenCover,proto3" json:"TokenCover,omitempty"` // StatusType Status = 9;
	Characters         []string              `protobuf:"bytes,12,rep,name=Characters,proto3" json:"Characters,omitempty"`
	Audience           AudienceType          `protobuf:"varint,13,opt,name=Audience,proto3,enum=stories.AudienceType" json:"Audience,omitempty"`
	Copyright          CopyrightType         `protobuf:"varint,14,opt,name=Copyright,proto3,enum=stories.CopyrightType" json:"Copyright,omitempty"`
	Clasification      ClasificationType     `protobuf:"varint,15,opt,name=Clasification,proto3,enum=stories.ClasificationType" json:"Clasification,omitempty"`
	CommentsModeration CommentModerationType `protobuf:"varint,16,opt,name=CommentsModeration,proto3,enum=stories.CommentModerationType" json:"CommentsModeration,omitempty"`
	Description        string                `protobuf:"bytes,17,opt,name=Description,proto3" json:"Description,omitempty"`
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

var File_stories_proto protoreflect.FileDescriptor

var file_stories_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x49,
	0x73, 0x4d, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x49, 0x73, 0x4d, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x40, 0x0a, 0x0d,
	0x43, 0x6c, 0x61, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6c,
	0x61, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0d, 0x43, 0x6c, 0x61, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e,
	0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x32, 0x8b, 0x01, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x41, 0x0a, 0x0c, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x42, 0x40,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x72, 0x6e,
	0x65, 0x73, 0x74, 0x6f, 0x47, 0x46, 0x2f, 0x62, 0x6f, 0x75, 0x6b, 0x6b, 0x65, 0x72, 0x2d, 0x6e,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_stories_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_stories_proto_goTypes = []interface{}{
	(*RequestWriteStory)(nil),   // 0: stories.RequestWriteStory
	(AudienceType)(0),           // 1: stories.AudienceType
	(CopyrightType)(0),          // 2: stories.CopyrightType
	(ClasificationType)(0),      // 3: stories.ClasificationType
	(CommentModerationType)(0),  // 4: stories.CommentModerationType
	(*RequestWriteChapter)(nil), // 5: stories.RequestWriteChapter
	(*ResponseID)(nil),          // 6: stories.ResponseID
}
var file_stories_proto_depIdxs = []int32{
	1, // 0: stories.RequestWriteStory.Audience:type_name -> stories.AudienceType
	2, // 1: stories.RequestWriteStory.Copyright:type_name -> stories.CopyrightType
	3, // 2: stories.RequestWriteStory.Clasification:type_name -> stories.ClasificationType
	4, // 3: stories.RequestWriteStory.CommentsModeration:type_name -> stories.CommentModerationType
	0, // 4: stories.stories.WriteStory:input_type -> stories.RequestWriteStory
	5, // 5: stories.stories.WriteChapter:input_type -> stories.RequestWriteChapter
	6, // 6: stories.stories.WriteStory:output_type -> stories.ResponseID
	6, // 7: stories.stories.WriteChapter:output_type -> stories.ResponseID
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stories_proto_goTypes,
		DependencyIndexes: file_stories_proto_depIdxs,
		MessageInfos:      file_stories_proto_msgTypes,
	}.Build()
	File_stories_proto = out.File
	file_stories_proto_rawDesc = nil
	file_stories_proto_goTypes = nil
	file_stories_proto_depIdxs = nil
}
