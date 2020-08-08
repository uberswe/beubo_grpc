// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.1
// source: beubo.proto

// protoc -I grpc --go_out=plugins=grpc:grpc grpc/beubo.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request is a basic representation of a http request
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string    `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method  string    `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers []*Header `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beubo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_beubo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_beubo_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

// Request is a basic representation of a http request
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beubo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_beubo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_beubo_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// Header represents all the headers of a request
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beubo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_beubo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_beubo_proto_rawDescGZIP(), []int{2}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// The request message containing the user's name.
type PluginMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Identifier  string                    `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Host        string                    `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Distributed bool                      `protobuf:"varint,4,opt,name=distributed,proto3" json:"distributed,omitempty"` // Should Beubo act as a loadbalancer with multiple plugin nodes?
	Caching     bool                      `protobuf:"varint,5,opt,name=caching,proto3" json:"caching,omitempty"`         // Caching can be enabled so that Beubo does not repeat the same request multiple times but instead caches the response
	Endpoints   []*PluginMessage_Endpoint `protobuf:"bytes,6,rep,name=endpoints,proto3" json:"endpoints,omitempty"`      // These are endpoints that Beubo or other plugins can call
}

func (x *PluginMessage) Reset() {
	*x = PluginMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beubo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginMessage) ProtoMessage() {}

func (x *PluginMessage) ProtoReflect() protoreflect.Message {
	mi := &file_beubo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginMessage.ProtoReflect.Descriptor instead.
func (*PluginMessage) Descriptor() ([]byte, []int) {
	return file_beubo_proto_rawDescGZIP(), []int{3}
}

func (x *PluginMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginMessage) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *PluginMessage) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PluginMessage) GetDistributed() bool {
	if x != nil {
		return x.Distributed
	}
	return false
}

func (x *PluginMessage) GetCaching() bool {
	if x != nil {
		return x.Caching
	}
	return false
}

func (x *PluginMessage) GetEndpoints() []*PluginMessage_Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

// Event is the struct used for plugins to communicate with Beubo.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data   string     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Values []*any.Any `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beubo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_beubo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_beubo_proto_rawDescGZIP(), []int{4}
}

func (x *Event) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Event) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Event) GetValues() []*any.Any {
	if x != nil {
		return x.Values
	}
	return nil
}

type PluginMessage_Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Priority int32  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *PluginMessage_Endpoint) Reset() {
	*x = PluginMessage_Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beubo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginMessage_Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginMessage_Endpoint) ProtoMessage() {}

func (x *PluginMessage_Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_beubo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginMessage_Endpoint.ProtoReflect.Descriptor instead.
func (*PluginMessage_Endpoint) Descriptor() ([]byte, []int) {
	return file_beubo_proto_rawDescGZIP(), []int{3, 0}
}

func (x *PluginMessage_Endpoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginMessage_Endpoint) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

var File_beubo_proto protoreflect.FileDescriptor

var file_beubo_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x65, 0x75, 0x62, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x65, 0x75, 0x62, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x27, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x65, 0x75, 0x62, 0x6f, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0x24, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x8c, 0x02, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a,
	0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x62, 0x65, 0x75, 0x62, 0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x3a, 0x0a, 0x08, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x5b, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x32, 0x6e, 0x0a, 0x09, 0x42, 0x65, 0x75, 0x62, 0x6f, 0x47, 0x52, 0x50, 0x43,
	0x12, 0x2b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x62, 0x65,
	0x75, 0x62, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x65, 0x75, 0x62,
	0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x34, 0x0a,
	0x08, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x62, 0x65, 0x75, 0x62,
	0x6f, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x0e, 0x2e, 0x62, 0x65, 0x75, 0x62, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x45, 0x0a, 0x05, 0x62, 0x65, 0x75, 0x62, 0x6f, 0x42, 0x0a, 0x42, 0x65,
	0x75, 0x62, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x73, 0x74, 0x65, 0x6e,
	0x67, 0x68, 0x61, 0x6d, 0x6e, 0x2f, 0x62, 0x65, 0x75, 0x62, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_beubo_proto_rawDescOnce sync.Once
	file_beubo_proto_rawDescData = file_beubo_proto_rawDesc
)

func file_beubo_proto_rawDescGZIP() []byte {
	file_beubo_proto_rawDescOnce.Do(func() {
		file_beubo_proto_rawDescData = protoimpl.X.CompressGZIP(file_beubo_proto_rawDescData)
	})
	return file_beubo_proto_rawDescData
}

var file_beubo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_beubo_proto_goTypes = []interface{}{
	(*Request)(nil),                // 0: beubo.Request
	(*Response)(nil),               // 1: beubo.Response
	(*Header)(nil),                 // 2: beubo.Header
	(*PluginMessage)(nil),          // 3: beubo.PluginMessage
	(*Event)(nil),                  // 4: beubo.Event
	(*PluginMessage_Endpoint)(nil), // 5: beubo.PluginMessage.Endpoint
	(*any.Any)(nil),                // 6: google.protobuf.Any
}
var file_beubo_proto_depIdxs = []int32{
	2, // 0: beubo.Request.headers:type_name -> beubo.Header
	5, // 1: beubo.PluginMessage.endpoints:type_name -> beubo.PluginMessage.Endpoint
	6, // 2: beubo.Event.values:type_name -> google.protobuf.Any
	4, // 3: beubo.BeuboGRPC.Connect:input_type -> beubo.Event
	3, // 4: beubo.BeuboGRPC.Requests:input_type -> beubo.PluginMessage
	4, // 5: beubo.BeuboGRPC.Connect:output_type -> beubo.Event
	0, // 6: beubo.BeuboGRPC.Requests:output_type -> beubo.Request
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_beubo_proto_init() }
func file_beubo_proto_init() {
	if File_beubo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beubo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_beubo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_beubo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_beubo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginMessage); i {
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
		file_beubo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_beubo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginMessage_Endpoint); i {
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
			RawDescriptor: file_beubo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_beubo_proto_goTypes,
		DependencyIndexes: file_beubo_proto_depIdxs,
		MessageInfos:      file_beubo_proto_msgTypes,
	}.Build()
	File_beubo_proto = out.File
	file_beubo_proto_rawDesc = nil
	file_beubo_proto_goTypes = nil
	file_beubo_proto_depIdxs = nil
}
