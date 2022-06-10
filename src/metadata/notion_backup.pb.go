// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: notion_backup.proto

package metadata

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

type NotionObjectType int32

const (
	NotionObjectType_UNKNOWN  NotionObjectType = 0
	NotionObjectType_ROOT     NotionObjectType = 1
	NotionObjectType_PAGE     NotionObjectType = 2
	NotionObjectType_DATABASE NotionObjectType = 3
	NotionObjectType_BLOCK    NotionObjectType = 4
)

// Enum value maps for NotionObjectType.
var (
	NotionObjectType_name = map[int32]string{
		0: "UNKNOWN",
		1: "ROOT",
		2: "PAGE",
		3: "DATABASE",
		4: "BLOCK",
	}
	NotionObjectType_value = map[string]int32{
		"UNKNOWN":  0,
		"ROOT":     1,
		"PAGE":     2,
		"DATABASE": 3,
		"BLOCK":    4,
	}
)

func (x NotionObjectType) Enum() *NotionObjectType {
	p := new(NotionObjectType)
	*p = x
	return p
}

func (x NotionObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotionObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_notion_backup_proto_enumTypes[0].Descriptor()
}

func (NotionObjectType) Type() protoreflect.EnumType {
	return &file_notion_backup_proto_enumTypes[0]
}

func (x NotionObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotionObjectType.Descriptor instead.
func (NotionObjectType) EnumDescriptor() ([]byte, []int) {
	return file_notion_backup_proto_rawDescGZIP(), []int{0}
}

type MetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map for storing NotionObject with uuid as a key and NotionObject as a value
	NotionObjectMap map[string]*NotionObject `protobuf:"bytes,1,rep,name=notion_object_map,json=notionObjectMap,proto3" json:"notion_object_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Map stores the mapping of each NotionObject UUID to different NotionObject
	// UUIDs
	ObjectMapping map[string]*ChildrenNotionObjectUuids `protobuf:"bytes,2,rep,name=object_mapping,json=objectMapping,proto3" json:"object_mapping,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetaData) Reset() {
	*x = MetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notion_backup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaData) ProtoMessage() {}

func (x *MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_notion_backup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaData.ProtoReflect.Descriptor instead.
func (*MetaData) Descriptor() ([]byte, []int) {
	return file_notion_backup_proto_rawDescGZIP(), []int{0}
}

func (x *MetaData) GetNotionObjectMap() map[string]*NotionObject {
	if x != nil {
		return x.NotionObjectMap
	}
	return nil
}

func (x *MetaData) GetObjectMapping() map[string]*ChildrenNotionObjectUuids {
	if x != nil {
		return x.ObjectMapping
	}
	return nil
}

// List of UUIDs of different NotionObject
type ChildrenNotionObjectUuids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChildrenUuidList []string `protobuf:"bytes,1,rep,name=children_uuid_list,json=childrenUuidList,proto3" json:"children_uuid_list,omitempty"`
}

func (x *ChildrenNotionObjectUuids) Reset() {
	*x = ChildrenNotionObjectUuids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notion_backup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildrenNotionObjectUuids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildrenNotionObjectUuids) ProtoMessage() {}

func (x *ChildrenNotionObjectUuids) ProtoReflect() protoreflect.Message {
	mi := &file_notion_backup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildrenNotionObjectUuids.ProtoReflect.Descriptor instead.
func (*ChildrenNotionObjectUuids) Descriptor() ([]byte, []int) {
	return file_notion_backup_proto_rawDescGZIP(), []int{1}
}

func (x *ChildrenNotionObjectUuids) GetChildrenUuidList() []string {
	if x != nil {
		return x.ChildrenUuidList
	}
	return nil
}

type NotionObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the object in notionbackup module
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Storage identifier of the object with which object can be retrived from
	// database or any storage system where notion objects are stored
	StorageIdentifier string `protobuf:"bytes,2,opt,name=storage_identifier,json=storageIdentifier,proto3" json:"storage_identifier,omitempty"`
	// Notion object type
	Type NotionObjectType `protobuf:"varint,3,opt,name=type,proto3,enum=NotionObjectType" json:"type,omitempty"`
	// ID of the notion object. This ID belongs to ID created by Notion App
	NotionObjectId string `protobuf:"bytes,4,opt,name=notion_object_id,json=notionObjectId,proto3" json:"notion_object_id,omitempty"`
}

func (x *NotionObject) Reset() {
	*x = NotionObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notion_backup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotionObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotionObject) ProtoMessage() {}

func (x *NotionObject) ProtoReflect() protoreflect.Message {
	mi := &file_notion_backup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotionObject.ProtoReflect.Descriptor instead.
func (*NotionObject) Descriptor() ([]byte, []int) {
	return file_notion_backup_proto_rawDescGZIP(), []int{2}
}

func (x *NotionObject) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *NotionObject) GetStorageIdentifier() string {
	if x != nil {
		return x.StorageIdentifier
	}
	return ""
}

func (x *NotionObject) GetType() NotionObjectType {
	if x != nil {
		return x.Type
	}
	return NotionObjectType_UNKNOWN
}

func (x *NotionObject) GetNotionObjectId() string {
	if x != nil {
		return x.NotionObjectId
	}
	return ""
}

var File_notion_backup_proto protoreflect.FileDescriptor

var file_notion_backup_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x4a, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x6e,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x43,
	0x0a, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x1a, 0x51, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x12, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x55, 0x75, 0x69, 0x64, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x49, 0x0a, 0x19, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x75, 0x69, 0x64,
	0x73, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0xa2, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x2a, 0x4c, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x4f, 0x54, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x50, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x41, 0x54,
	0x41, 0x42, 0x41, 0x53, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x4f, 0x43, 0x4b,
	0x10, 0x04, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notion_backup_proto_rawDescOnce sync.Once
	file_notion_backup_proto_rawDescData = file_notion_backup_proto_rawDesc
)

func file_notion_backup_proto_rawDescGZIP() []byte {
	file_notion_backup_proto_rawDescOnce.Do(func() {
		file_notion_backup_proto_rawDescData = protoimpl.X.CompressGZIP(file_notion_backup_proto_rawDescData)
	})
	return file_notion_backup_proto_rawDescData
}

var file_notion_backup_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notion_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notion_backup_proto_goTypes = []interface{}{
	(NotionObjectType)(0),             // 0: NotionObjectType
	(*MetaData)(nil),                  // 1: MetaData
	(*ChildrenNotionObjectUuids)(nil), // 2: ChildrenNotionObjectUuids
	(*NotionObject)(nil),              // 3: NotionObject
	nil,                               // 4: MetaData.NotionObjectMapEntry
	nil,                               // 5: MetaData.ObjectMappingEntry
}
var file_notion_backup_proto_depIdxs = []int32{
	4, // 0: MetaData.notion_object_map:type_name -> MetaData.NotionObjectMapEntry
	5, // 1: MetaData.object_mapping:type_name -> MetaData.ObjectMappingEntry
	0, // 2: NotionObject.type:type_name -> NotionObjectType
	3, // 3: MetaData.NotionObjectMapEntry.value:type_name -> NotionObject
	2, // 4: MetaData.ObjectMappingEntry.value:type_name -> ChildrenNotionObjectUuids
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_notion_backup_proto_init() }
func file_notion_backup_proto_init() {
	if File_notion_backup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notion_backup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaData); i {
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
		file_notion_backup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildrenNotionObjectUuids); i {
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
		file_notion_backup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotionObject); i {
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
			RawDescriptor: file_notion_backup_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notion_backup_proto_goTypes,
		DependencyIndexes: file_notion_backup_proto_depIdxs,
		EnumInfos:         file_notion_backup_proto_enumTypes,
		MessageInfos:      file_notion_backup_proto_msgTypes,
	}.Build()
	File_notion_backup_proto = out.File
	file_notion_backup_proto_rawDesc = nil
	file_notion_backup_proto_goTypes = nil
	file_notion_backup_proto_depIdxs = nil
}
