// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/cloud/datacatalog/v1/schema.proto

package datacatalogpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a schema, for example, a BigQuery, GoogleSQL, or Avro schema.
type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unified GoogleSQL-like schema of columns.
	//
	// The overall maximum number of columns and nested columns is 10,000.
	// The maximum nested depth is 15 levels.
	Columns []*ColumnSchema `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Schema) GetColumns() []*ColumnSchema {
	if x != nil {
		return x.Columns
	}
	return nil
}

// A column within a schema. Columns can be nested inside
// other columns.
type ColumnSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the column.
	//
	// Must be a UTF-8 string without dots (.).
	// The maximum size is 64 bytes.
	Column string `protobuf:"bytes,6,opt,name=column,proto3" json:"column,omitempty"`
	// Required. Type of the column.
	//
	// Must be a UTF-8 string with the maximum size of 128 bytes.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Optional. Description of the column. Default value is an empty string.
	//
	// The description must be a UTF-8 string with the maximum size of 2000
	// bytes.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. A column's mode indicates whether values in this column are required,
	// nullable, or repeated.
	//
	// Only `NULLABLE`, `REQUIRED`, and `REPEATED` values are supported.
	// Default mode is `NULLABLE`.
	Mode string `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	// Optional. Schema of sub-columns. A column can have zero or more sub-columns.
	Subcolumns []*ColumnSchema `protobuf:"bytes,7,rep,name=subcolumns,proto3" json:"subcolumns,omitempty"`
}

func (x *ColumnSchema) Reset() {
	*x = ColumnSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnSchema) ProtoMessage() {}

func (x *ColumnSchema) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnSchema.ProtoReflect.Descriptor instead.
func (*ColumnSchema) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *ColumnSchema) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *ColumnSchema) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ColumnSchema) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ColumnSchema) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *ColumnSchema) GetSubcolumns() []*ColumnSchema {
	if x != nil {
		return x.Subcolumns
	}
	return nil
}

var File_google_cloud_datacatalog_v1_schema_proto protoreflect.FileDescriptor

var file_google_cloud_datacatalog_v1_schema_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x43, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1b, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x4e,
	0x0a, 0x0a, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x42, 0xcb,
	0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31,
	0x3b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0xf8, 0x01, 0x01, 0xaa,
	0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datacatalog_v1_schema_proto_rawDescOnce sync.Once
	file_google_cloud_datacatalog_v1_schema_proto_rawDescData = file_google_cloud_datacatalog_v1_schema_proto_rawDesc
)

func file_google_cloud_datacatalog_v1_schema_proto_rawDescGZIP() []byte {
	file_google_cloud_datacatalog_v1_schema_proto_rawDescOnce.Do(func() {
		file_google_cloud_datacatalog_v1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datacatalog_v1_schema_proto_rawDescData)
	})
	return file_google_cloud_datacatalog_v1_schema_proto_rawDescData
}

var file_google_cloud_datacatalog_v1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_datacatalog_v1_schema_proto_goTypes = []interface{}{
	(*Schema)(nil),       // 0: google.cloud.datacatalog.v1.Schema
	(*ColumnSchema)(nil), // 1: google.cloud.datacatalog.v1.ColumnSchema
}
var file_google_cloud_datacatalog_v1_schema_proto_depIdxs = []int32{
	1, // 0: google.cloud.datacatalog.v1.Schema.columns:type_name -> google.cloud.datacatalog.v1.ColumnSchema
	1, // 1: google.cloud.datacatalog.v1.ColumnSchema.subcolumns:type_name -> google.cloud.datacatalog.v1.ColumnSchema
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_datacatalog_v1_schema_proto_init() }
func file_google_cloud_datacatalog_v1_schema_proto_init() {
	if File_google_cloud_datacatalog_v1_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datacatalog_v1_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema); i {
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
		file_google_cloud_datacatalog_v1_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnSchema); i {
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
			RawDescriptor: file_google_cloud_datacatalog_v1_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datacatalog_v1_schema_proto_goTypes,
		DependencyIndexes: file_google_cloud_datacatalog_v1_schema_proto_depIdxs,
		MessageInfos:      file_google_cloud_datacatalog_v1_schema_proto_msgTypes,
	}.Build()
	File_google_cloud_datacatalog_v1_schema_proto = out.File
	file_google_cloud_datacatalog_v1_schema_proto_rawDesc = nil
	file_google_cloud_datacatalog_v1_schema_proto_goTypes = nil
	file_google_cloud_datacatalog_v1_schema_proto_depIdxs = nil
}
