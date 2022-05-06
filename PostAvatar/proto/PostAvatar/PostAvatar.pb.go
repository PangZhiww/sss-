// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: proto/PostAvatar/PostAvatar.proto

package PostAvatar

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Say string `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_PostAvatar_PostAvatar_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSay() string {
	if x != nil {
		return x.Say
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar    []byte `protobuf:"bytes,1,opt,name=Avatar,proto3" json:"Avatar,omitempty"`      // 二进制图片流
	Filesize  int64  `protobuf:"varint,2,opt,name=Filesize,proto3" json:"Filesize,omitempty"` // 文件大小
	Fileext   string `protobuf:"bytes,3,opt,name=Fileext,proto3" json:"Fileext,omitempty"`    // 文件后缀
	SessionId string `protobuf:"bytes,4,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[1]
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
	return file_proto_PostAvatar_PostAvatar_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetAvatar() []byte {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *Request) GetFilesize() int64 {
	if x != nil {
		return x.Filesize
	}
	return 0
}

func (x *Request) GetFileext() string {
	if x != nil {
		return x.Fileext
	}
	return ""
}

func (x *Request) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno     string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg    string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	AvatarUrl string `protobuf:"bytes,3,opt,name=AvatarUrl,proto3" json:"AvatarUrl,omitempty"` // 不完整的头像地址
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[2]
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
	return file_proto_PostAvatar_PostAvatar_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *Response) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *Response) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type StreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StreamingRequest) Reset() {
	*x = StreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingRequest) ProtoMessage() {}

func (x *StreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingRequest.ProtoReflect.Descriptor instead.
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return file_proto_PostAvatar_PostAvatar_proto_rawDescGZIP(), []int{3}
}

func (x *StreamingRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StreamingResponse) Reset() {
	*x = StreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingResponse) ProtoMessage() {}

func (x *StreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingResponse.ProtoReflect.Descriptor instead.
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return file_proto_PostAvatar_PostAvatar_proto_rawDescGZIP(), []int{4}
}

func (x *StreamingResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_proto_PostAvatar_PostAvatar_proto_rawDescGZIP(), []int{5}
}

func (x *Ping) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_proto_PostAvatar_PostAvatar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_proto_PostAvatar_PostAvatar_proto_rawDescGZIP(), []int{6}
}

func (x *Pong) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

var File_proto_PostAvatar_PostAvatar_proto protoreflect.FileDescriptor

var file_proto_PostAvatar_PostAvatar_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x2f, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72,
	0x76, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x1b, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x61, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x61, 0x79, 0x22, 0x75, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x65,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x65,
	0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x56, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x45, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1e, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x22, 0x1e, 0x0a,
	0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x32, 0x61, 0x0a,
	0x0a, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x53, 0x0a, 0x0a, 0x50,
	0x6f, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x20, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_PostAvatar_PostAvatar_proto_rawDescOnce sync.Once
	file_proto_PostAvatar_PostAvatar_proto_rawDescData = file_proto_PostAvatar_PostAvatar_proto_rawDesc
)

func file_proto_PostAvatar_PostAvatar_proto_rawDescGZIP() []byte {
	file_proto_PostAvatar_PostAvatar_proto_rawDescOnce.Do(func() {
		file_proto_PostAvatar_PostAvatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_PostAvatar_PostAvatar_proto_rawDescData)
	})
	return file_proto_PostAvatar_PostAvatar_proto_rawDescData
}

var file_proto_PostAvatar_PostAvatar_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_PostAvatar_PostAvatar_proto_goTypes = []interface{}{
	(*Message)(nil),           // 0: go.micro.srv.PostAvatar.Message
	(*Request)(nil),           // 1: go.micro.srv.PostAvatar.Request
	(*Response)(nil),          // 2: go.micro.srv.PostAvatar.Response
	(*StreamingRequest)(nil),  // 3: go.micro.srv.PostAvatar.StreamingRequest
	(*StreamingResponse)(nil), // 4: go.micro.srv.PostAvatar.StreamingResponse
	(*Ping)(nil),              // 5: go.micro.srv.PostAvatar.Ping
	(*Pong)(nil),              // 6: go.micro.srv.PostAvatar.Pong
}
var file_proto_PostAvatar_PostAvatar_proto_depIdxs = []int32{
	1, // 0: go.micro.srv.PostAvatar.PostAvatar.PostAvatar:input_type -> go.micro.srv.PostAvatar.Request
	2, // 1: go.micro.srv.PostAvatar.PostAvatar.PostAvatar:output_type -> go.micro.srv.PostAvatar.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_PostAvatar_PostAvatar_proto_init() }
func file_proto_PostAvatar_PostAvatar_proto_init() {
	if File_proto_PostAvatar_PostAvatar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_PostAvatar_PostAvatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_PostAvatar_PostAvatar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_PostAvatar_PostAvatar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_PostAvatar_PostAvatar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingRequest); i {
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
		file_proto_PostAvatar_PostAvatar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingResponse); i {
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
		file_proto_PostAvatar_PostAvatar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_proto_PostAvatar_PostAvatar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_proto_PostAvatar_PostAvatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_PostAvatar_PostAvatar_proto_goTypes,
		DependencyIndexes: file_proto_PostAvatar_PostAvatar_proto_depIdxs,
		MessageInfos:      file_proto_PostAvatar_PostAvatar_proto_msgTypes,
	}.Build()
	File_proto_PostAvatar_PostAvatar_proto = out.File
	file_proto_PostAvatar_PostAvatar_proto_rawDesc = nil
	file_proto_PostAvatar_PostAvatar_proto_goTypes = nil
	file_proto_PostAvatar_PostAvatar_proto_depIdxs = nil
}
