// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Todo struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// @inject_tag: sql:",notnull,default:false"
	Completed bool `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	// @inject_tag: sql:"type:timestamptz,default:now()"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: sql:"type:timestamptz"
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Todo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Todo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateTodoRequest struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(m, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type CreateTodoResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(m, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

func (m *CreateTodoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoRequest) Reset()         { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()    {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *GetTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoRequest.Unmarshal(m, b)
}
func (m *GetTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoRequest.Marshal(b, m, deterministic)
}
func (m *GetTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoRequest.Merge(m, src)
}
func (m *GetTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoRequest.Size(m)
}
func (m *GetTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoRequest proto.InternalMessageInfo

func (m *GetTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodoResponse struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoResponse) Reset()         { *m = GetTodoResponse{} }
func (m *GetTodoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoResponse) ProtoMessage()    {}
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *GetTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResponse.Unmarshal(m, b)
}
func (m *GetTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResponse.Marshal(b, m, deterministic)
}
func (m *GetTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResponse.Merge(m, src)
}
func (m *GetTodoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoResponse.Size(m)
}
func (m *GetTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResponse proto.InternalMessageInfo

func (m *GetTodoResponse) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type ListTodoRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	NotCompleted         bool     `protobuf:"varint,2,opt,name=not_completed,json=notCompleted,proto3" json:"not_completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoRequest) Reset()         { *m = ListTodoRequest{} }
func (m *ListTodoRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodoRequest) ProtoMessage()    {}
func (*ListTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{5}
}

func (m *ListTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoRequest.Unmarshal(m, b)
}
func (m *ListTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoRequest.Marshal(b, m, deterministic)
}
func (m *ListTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoRequest.Merge(m, src)
}
func (m *ListTodoRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodoRequest.Size(m)
}
func (m *ListTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoRequest proto.InternalMessageInfo

func (m *ListTodoRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListTodoRequest) GetNotCompleted() bool {
	if m != nil {
		return m.NotCompleted
	}
	return false
}

type ListTodoResponse struct {
	Items                []*Todo  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoResponse) Reset()         { *m = ListTodoResponse{} }
func (m *ListTodoResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodoResponse) ProtoMessage()    {}
func (*ListTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{6}
}

func (m *ListTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoResponse.Unmarshal(m, b)
}
func (m *ListTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoResponse.Marshal(b, m, deterministic)
}
func (m *ListTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoResponse.Merge(m, src)
}
func (m *ListTodoResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodoResponse.Size(m)
}
func (m *ListTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoResponse proto.InternalMessageInfo

func (m *ListTodoResponse) GetItems() []*Todo {
	if m != nil {
		return m.Items
	}
	return nil
}

type DeleteTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRequest) Reset()         { *m = DeleteTodoRequest{} }
func (m *DeleteTodoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRequest) ProtoMessage()    {}
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{7}
}

func (m *DeleteTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRequest.Unmarshal(m, b)
}
func (m *DeleteTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRequest.Merge(m, src)
}
func (m *DeleteTodoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRequest.Size(m)
}
func (m *DeleteTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRequest proto.InternalMessageInfo

func (m *DeleteTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoResponse) Reset()         { *m = DeleteTodoResponse{} }
func (m *DeleteTodoResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoResponse) ProtoMessage()    {}
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{8}
}

func (m *DeleteTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoResponse.Unmarshal(m, b)
}
func (m *DeleteTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoResponse.Merge(m, src)
}
func (m *DeleteTodoResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoResponse.Size(m)
}
func (m *DeleteTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Todo)(nil), "todo.v1.Todo")
	proto.RegisterType((*CreateTodoRequest)(nil), "todo.v1.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "todo.v1.CreateTodoResponse")
	proto.RegisterType((*GetTodoRequest)(nil), "todo.v1.GetTodoRequest")
	proto.RegisterType((*GetTodoResponse)(nil), "todo.v1.GetTodoResponse")
	proto.RegisterType((*ListTodoRequest)(nil), "todo.v1.ListTodoRequest")
	proto.RegisterType((*ListTodoResponse)(nil), "todo.v1.ListTodoResponse")
	proto.RegisterType((*DeleteTodoRequest)(nil), "todo.v1.DeleteTodoRequest")
	proto.RegisterType((*DeleteTodoResponse)(nil), "todo.v1.DeleteTodoResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0x96, 0xdd, 0xb8, 0x4d, 0x5e, 0x48, 0x93, 0x3e, 0x05, 0x30, 0xa6, 0x12, 0xc6, 0x61, 0x11,
	0xb1, 0xb0, 0xd5, 0x80, 0x40, 0xb0, 0x2b, 0x45, 0x62, 0xd3, 0x0d, 0x6e, 0x57, 0x08, 0x54, 0x39,
	0x99, 0xa1, 0x1a, 0x29, 0xf6, 0x98, 0xf8, 0x25, 0x1b, 0xc4, 0x86, 0x2b, 0x70, 0x34, 0x6e, 0x80,
	0x38, 0x01, 0x27, 0x40, 0x1e, 0x8f, 0x7f, 0x12, 0x57, 0xd0, 0x9d, 0xe7, 0x7d, 0x3f, 0xef, 0x9b,
	0x6f, 0x64, 0x00, 0x92, 0x4c, 0xfa, 0xe9, 0x4a, 0x92, 0xc4, 0x03, 0xf5, 0xbd, 0x39, 0x71, 0x8e,
	0xaf, 0xa5, 0xbc, 0x5e, 0xf2, 0x20, 0x4a, 0x45, 0x10, 0x25, 0x89, 0xa4, 0x88, 0x84, 0x4c, 0xb2,
	0x82, 0xe6, 0x3c, 0xd2, 0xa8, 0x3a, 0xcd, 0xd7, 0x9f, 0x03, 0x12, 0x31, 0xcf, 0x28, 0x8a, 0xd3,
	0x82, 0xe0, 0xfd, 0x32, 0xa0, 0x73, 0x29, 0x99, 0xc4, 0x43, 0x30, 0x05, 0xb3, 0x0d, 0xd7, 0x98,
	0xf6, 0x42, 0x53, 0x30, 0x1c, 0x83, 0x45, 0x82, 0x96, 0xdc, 0x36, 0xd5, 0xa8, 0x38, 0xa0, 0x0b,
	0x7d, 0xc6, 0xb3, 0xc5, 0x4a, 0xa4, 0xf9, 0x16, 0x7b, 0x4f, 0x61, 0xcd, 0x11, 0x1e, 0x43, 0x6f,
	0x21, 0xe3, 0x74, 0xc9, 0x89, 0x33, 0xbb, 0xe3, 0x1a, 0xd3, 0x6e, 0x58, 0x0f, 0xf0, 0x15, 0xc0,
	0x62, 0xc5, 0x23, 0xe2, 0xec, 0x2a, 0x22, 0xdb, 0x72, 0x8d, 0x69, 0x7f, 0xe6, 0xf8, 0x45, 0x48,
	0xbf, 0x0c, 0xe9, 0x5f, 0x96, 0x21, 0xc3, 0x9e, 0x66, 0x9f, 0x52, 0x2e, 0x5d, 0xa7, 0xac, 0x94,
	0xee, 0xff, 0x5f, 0xaa, 0xd9, 0xa7, 0xe4, 0xbd, 0x80, 0xa3, 0x33, 0xe5, 0x93, 0xdf, 0x34, 0xe4,
	0x5f, 0xd6, 0x3c, 0x23, 0x7c, 0x0c, 0x1d, 0x41, 0x3c, 0x56, 0x57, 0xee, 0xcf, 0x06, 0xbe, 0x2e,
	0xd4, 0x57, 0x1c, 0x05, 0x79, 0x4f, 0x00, 0x9b, 0xba, 0x2c, 0x95, 0x49, 0xc6, 0x77, 0x9b, 0xf2,
	0x5c, 0x38, 0x7c, 0xc7, 0xa9, 0x69, 0xbd, 0xcb, 0x78, 0x0e, 0xc3, 0x8a, 0xa1, 0x4d, 0x6e, 0xb1,
	0xfd, 0x1c, 0x86, 0xe7, 0x22, 0xdb, 0x32, 0x1e, 0x83, 0xb5, 0x14, 0xb1, 0x20, 0x25, 0xb3, 0xc2,
	0xe2, 0x80, 0x13, 0x18, 0x24, 0x92, 0xae, 0xea, 0xda, 0x4d, 0x55, 0xfb, 0x9d, 0x44, 0xd2, 0x59,
	0x39, 0xf3, 0x5e, 0xc2, 0xa8, 0x76, 0xd3, 0x21, 0x26, 0x60, 0xe5, 0x9b, 0x32, 0xdb, 0x70, 0xf7,
	0xda, 0x29, 0x0a, 0xcc, 0x9b, 0xc0, 0xd1, 0x5b, 0x9e, 0x7b, 0xfc, 0xeb, 0x86, 0x63, 0xc0, 0x26,
	0xa9, 0xf0, 0x9f, 0xfd, 0x31, 0xa1, 0x9f, 0x0f, 0x2e, 0xf8, 0x6a, 0x23, 0x16, 0x1c, 0x3f, 0x01,
	0xd4, 0x7d, 0xa2, 0x53, 0xad, 0x6b, 0x3d, 0x8e, 0xf3, 0xf0, 0x46, 0xac, 0xb0, 0xf5, 0xee, 0x7d,
	0xff, 0xf9, 0xfb, 0x87, 0x39, 0xf2, 0xba, 0xc1, 0xe6, 0x24, 0xc8, 0x79, 0xaf, 0x55, 0x61, 0x78,
	0x01, 0x07, 0xba, 0x66, 0xbc, 0x5f, 0xe9, 0xb7, 0x9f, 0xc6, 0xb1, 0xdb, 0x80, 0x76, 0xbd, 0xab,
	0x5c, 0x87, 0x38, 0x28, 0x5d, 0x83, 0xaf, 0x82, 0x7d, 0xc3, 0xf7, 0xd0, 0x2d, 0x7b, 0xc3, 0x5a,
	0xbc, 0xf3, 0x30, 0xce, 0x83, 0x1b, 0x10, 0xed, 0x3b, 0x52, 0xbe, 0x80, 0x55, 0x5a, 0xfc, 0x08,
	0x50, 0x97, 0xd5, 0xa8, 0xa1, 0x55, 0x73, 0xa3, 0x86, 0x76, 0xbb, 0x65, 0xe0, 0xa7, 0xdb, 0x81,
	0xdf, 0x74, 0x3e, 0x98, 0xe9, 0x7c, 0xbe, 0xaf, 0xfe, 0x88, 0x67, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x18, 0x0f, 0x74, 0xd3, 0x34, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error)
	ListTodo(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (*ListTodoResponse, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error) {
	out := new(GetTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoService/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodo(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (*ListTodoResponse, error) {
	out := new(ListTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoService/ListTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	GetTodo(context.Context, *GetTodoRequest) (*GetTodoResponse, error)
	ListTodo(context.Context, *ListTodoRequest) (*ListTodoResponse, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) GetTodo(ctx context.Context, req *GetTodoRequest) (*GetTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (*UnimplementedTodoServiceServer) ListTodo(ctx context.Context, req *ListTodoRequest) (*ListTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodo not implemented")
}
func (*UnimplementedTodoServiceServer) DeleteTodo(ctx context.Context, req *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoService/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoService/ListTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodo(ctx, req.(*ListTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.v1.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoService_GetTodo_Handler,
		},
		{
			MethodName: "ListTodo",
			Handler:    _TodoService_ListTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
