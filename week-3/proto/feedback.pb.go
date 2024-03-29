// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feedback.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PassengerFeedback struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerFeedback) Reset()         { *m = PassengerFeedback{} }
func (m *PassengerFeedback) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedback) ProtoMessage()    {}
func (*PassengerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{0}
}

func (m *PassengerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedback.Unmarshal(m, b)
}
func (m *PassengerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedback.Marshal(b, m, deterministic)
}
func (m *PassengerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedback.Merge(m, src)
}
func (m *PassengerFeedback) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedback.Size(m)
}
func (m *PassengerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedback proto.InternalMessageInfo

func (m *PassengerFeedback) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *PassengerFeedback) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *PassengerFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type AddFeedbackRequest struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFeedbackRequest) Reset()         { *m = AddFeedbackRequest{} }
func (m *AddFeedbackRequest) String() string { return proto.CompactTextString(m) }
func (*AddFeedbackRequest) ProtoMessage()    {}
func (*AddFeedbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{1}
}

func (m *AddFeedbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeedbackRequest.Unmarshal(m, b)
}
func (m *AddFeedbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeedbackRequest.Marshal(b, m, deterministic)
}
func (m *AddFeedbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeedbackRequest.Merge(m, src)
}
func (m *AddFeedbackRequest) XXX_Size() int {
	return xxx_messageInfo_AddFeedbackRequest.Size(m)
}
func (m *AddFeedbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeedbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeedbackRequest proto.InternalMessageInfo

func (m *AddFeedbackRequest) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *AddFeedbackRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *AddFeedbackRequest) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type AddFeedbackResponse struct {
	ReturnCode           int32    `protobuf:"varint,1,opt,name=returnCode,proto3" json:"returnCode,omitempty"`
	ReturnMessage        string   `protobuf:"bytes,2,opt,name=returnMessage,proto3" json:"returnMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFeedbackResponse) Reset()         { *m = AddFeedbackResponse{} }
func (m *AddFeedbackResponse) String() string { return proto.CompactTextString(m) }
func (*AddFeedbackResponse) ProtoMessage()    {}
func (*AddFeedbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{2}
}

func (m *AddFeedbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeedbackResponse.Unmarshal(m, b)
}
func (m *AddFeedbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeedbackResponse.Marshal(b, m, deterministic)
}
func (m *AddFeedbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeedbackResponse.Merge(m, src)
}
func (m *AddFeedbackResponse) XXX_Size() int {
	return xxx_messageInfo_AddFeedbackResponse.Size(m)
}
func (m *AddFeedbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeedbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeedbackResponse proto.InternalMessageInfo

func (m *AddFeedbackResponse) GetReturnCode() int32 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

func (m *AddFeedbackResponse) GetReturnMessage() string {
	if m != nil {
		return m.ReturnMessage
	}
	return ""
}

type GetFeedbackByPassengerIDRequest struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedbackByPassengerIDRequest) Reset()         { *m = GetFeedbackByPassengerIDRequest{} }
func (m *GetFeedbackByPassengerIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackByPassengerIDRequest) ProtoMessage()    {}
func (*GetFeedbackByPassengerIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{3}
}

func (m *GetFeedbackByPassengerIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackByPassengerIDRequest.Unmarshal(m, b)
}
func (m *GetFeedbackByPassengerIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackByPassengerIDRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedbackByPassengerIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackByPassengerIDRequest.Merge(m, src)
}
func (m *GetFeedbackByPassengerIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackByPassengerIDRequest.Size(m)
}
func (m *GetFeedbackByPassengerIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackByPassengerIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackByPassengerIDRequest proto.InternalMessageInfo

func (m *GetFeedbackByPassengerIDRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

type GetFeedbackByPassengerIDResponse struct {
	ReturnCode           int32                `protobuf:"varint,1,opt,name=returnCode,proto3" json:"returnCode,omitempty"`
	ReturnMessage        string               `protobuf:"bytes,2,opt,name=returnMessage,proto3" json:"returnMessage,omitempty"`
	Feedbacks            []*PassengerFeedback `protobuf:"bytes,3,rep,name=feedbacks,proto3" json:"feedbacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetFeedbackByPassengerIDResponse) Reset()         { *m = GetFeedbackByPassengerIDResponse{} }
func (m *GetFeedbackByPassengerIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackByPassengerIDResponse) ProtoMessage()    {}
func (*GetFeedbackByPassengerIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{4}
}

func (m *GetFeedbackByPassengerIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackByPassengerIDResponse.Unmarshal(m, b)
}
func (m *GetFeedbackByPassengerIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackByPassengerIDResponse.Marshal(b, m, deterministic)
}
func (m *GetFeedbackByPassengerIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackByPassengerIDResponse.Merge(m, src)
}
func (m *GetFeedbackByPassengerIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackByPassengerIDResponse.Size(m)
}
func (m *GetFeedbackByPassengerIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackByPassengerIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackByPassengerIDResponse proto.InternalMessageInfo

func (m *GetFeedbackByPassengerIDResponse) GetReturnCode() int32 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

func (m *GetFeedbackByPassengerIDResponse) GetReturnMessage() string {
	if m != nil {
		return m.ReturnMessage
	}
	return ""
}

func (m *GetFeedbackByPassengerIDResponse) GetFeedbacks() []*PassengerFeedback {
	if m != nil {
		return m.Feedbacks
	}
	return nil
}

type GetFeedbackByBookingCodeRequest struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedbackByBookingCodeRequest) Reset()         { *m = GetFeedbackByBookingCodeRequest{} }
func (m *GetFeedbackByBookingCodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackByBookingCodeRequest) ProtoMessage()    {}
func (*GetFeedbackByBookingCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{5}
}

func (m *GetFeedbackByBookingCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackByBookingCodeRequest.Unmarshal(m, b)
}
func (m *GetFeedbackByBookingCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackByBookingCodeRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedbackByBookingCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackByBookingCodeRequest.Merge(m, src)
}
func (m *GetFeedbackByBookingCodeRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackByBookingCodeRequest.Size(m)
}
func (m *GetFeedbackByBookingCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackByBookingCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackByBookingCodeRequest proto.InternalMessageInfo

func (m *GetFeedbackByBookingCodeRequest) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

type GetFeedbackByBookingCodeResponse struct {
	ReturnCode           int32              `protobuf:"varint,1,opt,name=returnCode,proto3" json:"returnCode,omitempty"`
	ReturnMessage        string             `protobuf:"bytes,2,opt,name=returnMessage,proto3" json:"returnMessage,omitempty"`
	Feedback             *PassengerFeedback `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetFeedbackByBookingCodeResponse) Reset()         { *m = GetFeedbackByBookingCodeResponse{} }
func (m *GetFeedbackByBookingCodeResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackByBookingCodeResponse) ProtoMessage()    {}
func (*GetFeedbackByBookingCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{6}
}

func (m *GetFeedbackByBookingCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackByBookingCodeResponse.Unmarshal(m, b)
}
func (m *GetFeedbackByBookingCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackByBookingCodeResponse.Marshal(b, m, deterministic)
}
func (m *GetFeedbackByBookingCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackByBookingCodeResponse.Merge(m, src)
}
func (m *GetFeedbackByBookingCodeResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackByBookingCodeResponse.Size(m)
}
func (m *GetFeedbackByBookingCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackByBookingCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackByBookingCodeResponse proto.InternalMessageInfo

func (m *GetFeedbackByBookingCodeResponse) GetReturnCode() int32 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

func (m *GetFeedbackByBookingCodeResponse) GetReturnMessage() string {
	if m != nil {
		return m.ReturnMessage
	}
	return ""
}

func (m *GetFeedbackByBookingCodeResponse) GetFeedback() *PassengerFeedback {
	if m != nil {
		return m.Feedback
	}
	return nil
}

type DeleteFeedbackByPassengerIDRequest struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFeedbackByPassengerIDRequest) Reset()         { *m = DeleteFeedbackByPassengerIDRequest{} }
func (m *DeleteFeedbackByPassengerIDRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFeedbackByPassengerIDRequest) ProtoMessage()    {}
func (*DeleteFeedbackByPassengerIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{7}
}

func (m *DeleteFeedbackByPassengerIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDRequest.Unmarshal(m, b)
}
func (m *DeleteFeedbackByPassengerIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFeedbackByPassengerIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFeedbackByPassengerIDRequest.Merge(m, src)
}
func (m *DeleteFeedbackByPassengerIDRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDRequest.Size(m)
}
func (m *DeleteFeedbackByPassengerIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFeedbackByPassengerIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFeedbackByPassengerIDRequest proto.InternalMessageInfo

func (m *DeleteFeedbackByPassengerIDRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

type DeleteFeedbackByPassengerIDResponse struct {
	ReturnCode           int32    `protobuf:"varint,1,opt,name=returnCode,proto3" json:"returnCode,omitempty"`
	ReturnMessage        string   `protobuf:"bytes,2,opt,name=returnMessage,proto3" json:"returnMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFeedbackByPassengerIDResponse) Reset()         { *m = DeleteFeedbackByPassengerIDResponse{} }
func (m *DeleteFeedbackByPassengerIDResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteFeedbackByPassengerIDResponse) ProtoMessage()    {}
func (*DeleteFeedbackByPassengerIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{8}
}

func (m *DeleteFeedbackByPassengerIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDResponse.Unmarshal(m, b)
}
func (m *DeleteFeedbackByPassengerIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDResponse.Marshal(b, m, deterministic)
}
func (m *DeleteFeedbackByPassengerIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFeedbackByPassengerIDResponse.Merge(m, src)
}
func (m *DeleteFeedbackByPassengerIDResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDResponse.Size(m)
}
func (m *DeleteFeedbackByPassengerIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFeedbackByPassengerIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFeedbackByPassengerIDResponse proto.InternalMessageInfo

func (m *DeleteFeedbackByPassengerIDResponse) GetReturnCode() int32 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

func (m *DeleteFeedbackByPassengerIDResponse) GetReturnMessage() string {
	if m != nil {
		return m.ReturnMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*PassengerFeedback)(nil), "proto.PassengerFeedback")
	proto.RegisterType((*AddFeedbackRequest)(nil), "proto.AddFeedbackRequest")
	proto.RegisterType((*AddFeedbackResponse)(nil), "proto.AddFeedbackResponse")
	proto.RegisterType((*GetFeedbackByPassengerIDRequest)(nil), "proto.GetFeedbackByPassengerIDRequest")
	proto.RegisterType((*GetFeedbackByPassengerIDResponse)(nil), "proto.GetFeedbackByPassengerIDResponse")
	proto.RegisterType((*GetFeedbackByBookingCodeRequest)(nil), "proto.GetFeedbackByBookingCodeRequest")
	proto.RegisterType((*GetFeedbackByBookingCodeResponse)(nil), "proto.GetFeedbackByBookingCodeResponse")
	proto.RegisterType((*DeleteFeedbackByPassengerIDRequest)(nil), "proto.DeleteFeedbackByPassengerIDRequest")
	proto.RegisterType((*DeleteFeedbackByPassengerIDResponse)(nil), "proto.DeleteFeedbackByPassengerIDResponse")
}

func init() { proto.RegisterFile("feedback.proto", fileDescriptor_7b189e8c8330c03e) }

var fileDescriptor_7b189e8c8330c03e = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0xab, 0xd3, 0x40,
	0x10, 0xc0, 0xc9, 0xab, 0x15, 0x3b, 0x55, 0xc1, 0xf1, 0x12, 0xa3, 0xf8, 0xf2, 0xf6, 0x89, 0xaf,
	0xef, 0x1d, 0x1a, 0xa8, 0xe2, 0xdd, 0xb6, 0x54, 0x3c, 0x08, 0x25, 0x5e, 0x04, 0x4f, 0x69, 0x33,
	0x86, 0xd0, 0x92, 0x8d, 0xd9, 0xad, 0x20, 0xa5, 0x17, 0xbf, 0x82, 0xe2, 0x51, 0xfd, 0x34, 0x7e,
	0x01, 0xbf, 0x82, 0x1f, 0x44, 0xcc, 0xdf, 0x4d, 0x93, 0x36, 0x05, 0xeb, 0xa9, 0x64, 0x3a, 0x3b,
	0xf3, 0x9b, 0xd9, 0xdf, 0xc2, 0xed, 0x77, 0x44, 0xee, 0xcc, 0x99, 0x2f, 0xfa, 0x61, 0xc4, 0x25,
	0xc7, 0x76, 0xfc, 0x63, 0x3c, 0xf0, 0x38, 0xf7, 0x96, 0x64, 0x39, 0xa1, 0x6f, 0x39, 0x41, 0xc0,
	0xa5, 0x23, 0x7d, 0x1e, 0x88, 0x24, 0x89, 0x09, 0xb8, 0x33, 0x75, 0x84, 0xa0, 0xc0, 0xa3, 0x68,
	0x92, 0x9e, 0x47, 0x13, 0xba, 0x33, 0xce, 0x17, 0x7e, 0xe0, 0x8d, 0xb8, 0x4b, 0xba, 0x66, 0x6a,
	0xbd, 0x8e, 0xad, 0x86, 0xfe, 0x66, 0x84, 0xd9, 0xb1, 0x97, 0x63, 0xfd, 0xc4, 0xd4, 0x7a, 0x6d,
	0x5b, 0x0d, 0xa1, 0x01, 0x37, 0x32, 0x1e, 0xbd, 0x15, 0x17, 0xc8, 0xbf, 0x99, 0x04, 0x7c, 0xee,
	0xba, 0x59, 0x3b, 0x9b, 0xde, 0xaf, 0x48, 0xc8, 0xff, 0xde, 0xf5, 0x2d, 0xdc, 0x2d, 0x75, 0x15,
	0x21, 0x0f, 0x04, 0xe1, 0x43, 0x80, 0x88, 0xe4, 0x2a, 0x0a, 0xf2, 0xae, 0x6d, 0x5b, 0x89, 0xe0,
	0x23, 0xb8, 0x95, 0x7c, 0xbd, 0x22, 0x21, 0x1c, 0x8f, 0xe2, 0xb6, 0x1d, 0xbb, 0x1c, 0x64, 0x23,
	0x38, 0x7d, 0x41, 0x32, 0x2b, 0x3e, 0xfc, 0x38, 0x2d, 0xa0, 0x94, 0xf9, 0x54, 0x7a, 0xad, 0x42,
	0xcf, 0x7e, 0x68, 0x60, 0xee, 0xae, 0x72, 0x4c, 0x5e, 0x7c, 0x06, 0x9d, 0x6c, 0x31, 0x42, 0x6f,
	0x99, 0xad, 0x5e, 0x77, 0xa0, 0x27, 0x4a, 0xf4, 0x2b, 0x3e, 0xd8, 0x45, 0x6a, 0x65, 0xce, 0x61,
	0x71, 0x3d, 0x07, 0xdf, 0x23, 0xfb, 0xb6, 0x3d, 0x67, 0xa9, 0xca, 0x51, 0xe7, 0x7c, 0xba, 0x25,
	0xc4, 0xbe, 0x31, 0x0b, 0x55, 0x26, 0xc0, 0xc6, 0xb4, 0x24, 0x49, 0xff, 0x78, 0xa1, 0x0b, 0x38,
	0xdf, 0x5b, 0xe7, 0x98, 0xa3, 0x0e, 0x7e, 0x5e, 0x03, 0xbd, 0x32, 0xd4, 0x6b, 0x8a, 0x3e, 0xf8,
	0x73, 0xc2, 0x37, 0xd0, 0x55, 0xe4, 0xc7, 0x7b, 0xe9, 0x12, 0xaa, 0xcf, 0xd0, 0x30, 0xea, 0xfe,
	0x4a, 0x40, 0x19, 0x7e, 0xfa, 0xf5, 0xfb, 0xf3, 0xc9, 0x4d, 0x04, 0x2b, 0x37, 0x02, 0xbf, 0x6a,
	0xa0, 0xef, 0x92, 0x16, 0x1f, 0xa7, 0xc5, 0x1a, 0xde, 0x86, 0x71, 0xd1, 0x98, 0x97, 0x12, 0x5c,
	0xc6, 0x04, 0xe7, 0x78, 0x96, 0x13, 0x58, 0xf9, 0xc2, 0xad, 0xb5, 0xb2, 0xfb, 0x0d, 0x7e, 0xd9,
	0x06, 0x53, 0x2c, 0xab, 0x07, 0xab, 0xca, 0x5c, 0x0f, 0x56, 0xa3, 0x2b, 0xbb, 0x88, 0xc1, 0xce,
	0xf0, 0xb4, 0x00, 0x4b, 0x95, 0xb7, 0xd6, 0x8a, 0xfb, 0x1b, 0xfc, 0xae, 0xc1, 0xfd, 0x3d, 0x52,
	0xe0, 0x65, 0xda, 0xb1, 0x59, 0x40, 0xe3, 0xea, 0x90, 0xd4, 0xf2, 0xe2, 0xae, 0x9a, 0x17, 0x37,
	0xbb, 0x1e, 0x57, 0x7d, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x37, 0x0d, 0x7c, 0x51, 0x06,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PassengerFeedbackServiceClient is the client API for PassengerFeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PassengerFeedbackServiceClient interface {
	AddFeedback(ctx context.Context, in *AddFeedbackRequest, opts ...grpc.CallOption) (*AddFeedbackResponse, error)
	GetFeedbackByPassengerID(ctx context.Context, in *GetFeedbackByPassengerIDRequest, opts ...grpc.CallOption) (*GetFeedbackByPassengerIDResponse, error)
	GetFeedbackByBookingCode(ctx context.Context, in *GetFeedbackByBookingCodeRequest, opts ...grpc.CallOption) (*GetFeedbackByBookingCodeResponse, error)
	DeleteFeedbackByPassengerID(ctx context.Context, in *DeleteFeedbackByPassengerIDRequest, opts ...grpc.CallOption) (*DeleteFeedbackByPassengerIDResponse, error)
}

type passengerFeedbackServiceClient struct {
	cc *grpc.ClientConn
}

func NewPassengerFeedbackServiceClient(cc *grpc.ClientConn) PassengerFeedbackServiceClient {
	return &passengerFeedbackServiceClient{cc}
}

func (c *passengerFeedbackServiceClient) AddFeedback(ctx context.Context, in *AddFeedbackRequest, opts ...grpc.CallOption) (*AddFeedbackResponse, error) {
	out := new(AddFeedbackResponse)
	err := c.cc.Invoke(ctx, "/proto.PassengerFeedbackService/AddFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) GetFeedbackByPassengerID(ctx context.Context, in *GetFeedbackByPassengerIDRequest, opts ...grpc.CallOption) (*GetFeedbackByPassengerIDResponse, error) {
	out := new(GetFeedbackByPassengerIDResponse)
	err := c.cc.Invoke(ctx, "/proto.PassengerFeedbackService/GetFeedbackByPassengerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) GetFeedbackByBookingCode(ctx context.Context, in *GetFeedbackByBookingCodeRequest, opts ...grpc.CallOption) (*GetFeedbackByBookingCodeResponse, error) {
	out := new(GetFeedbackByBookingCodeResponse)
	err := c.cc.Invoke(ctx, "/proto.PassengerFeedbackService/GetFeedbackByBookingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) DeleteFeedbackByPassengerID(ctx context.Context, in *DeleteFeedbackByPassengerIDRequest, opts ...grpc.CallOption) (*DeleteFeedbackByPassengerIDResponse, error) {
	out := new(DeleteFeedbackByPassengerIDResponse)
	err := c.cc.Invoke(ctx, "/proto.PassengerFeedbackService/DeleteFeedbackByPassengerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassengerFeedbackServiceServer is the server API for PassengerFeedbackService service.
type PassengerFeedbackServiceServer interface {
	AddFeedback(context.Context, *AddFeedbackRequest) (*AddFeedbackResponse, error)
	GetFeedbackByPassengerID(context.Context, *GetFeedbackByPassengerIDRequest) (*GetFeedbackByPassengerIDResponse, error)
	GetFeedbackByBookingCode(context.Context, *GetFeedbackByBookingCodeRequest) (*GetFeedbackByBookingCodeResponse, error)
	DeleteFeedbackByPassengerID(context.Context, *DeleteFeedbackByPassengerIDRequest) (*DeleteFeedbackByPassengerIDResponse, error)
}

// UnimplementedPassengerFeedbackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPassengerFeedbackServiceServer struct {
}

func (*UnimplementedPassengerFeedbackServiceServer) AddFeedback(ctx context.Context, req *AddFeedbackRequest) (*AddFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeedback not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) GetFeedbackByPassengerID(ctx context.Context, req *GetFeedbackByPassengerIDRequest) (*GetFeedbackByPassengerIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedbackByPassengerID not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) GetFeedbackByBookingCode(ctx context.Context, req *GetFeedbackByBookingCodeRequest) (*GetFeedbackByBookingCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedbackByBookingCode not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) DeleteFeedbackByPassengerID(ctx context.Context, req *DeleteFeedbackByPassengerIDRequest) (*DeleteFeedbackByPassengerIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeedbackByPassengerID not implemented")
}

func RegisterPassengerFeedbackServiceServer(s *grpc.Server, srv PassengerFeedbackServiceServer) {
	s.RegisterService(&_PassengerFeedbackService_serviceDesc, srv)
}

func _PassengerFeedbackService_AddFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).AddFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PassengerFeedbackService/AddFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).AddFeedback(ctx, req.(*AddFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_GetFeedbackByPassengerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedbackByPassengerIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).GetFeedbackByPassengerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PassengerFeedbackService/GetFeedbackByPassengerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).GetFeedbackByPassengerID(ctx, req.(*GetFeedbackByPassengerIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_GetFeedbackByBookingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedbackByBookingCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).GetFeedbackByBookingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PassengerFeedbackService/GetFeedbackByBookingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).GetFeedbackByBookingCode(ctx, req.(*GetFeedbackByBookingCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_DeleteFeedbackByPassengerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeedbackByPassengerIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).DeleteFeedbackByPassengerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PassengerFeedbackService/DeleteFeedbackByPassengerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).DeleteFeedbackByPassengerID(ctx, req.(*DeleteFeedbackByPassengerIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PassengerFeedbackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PassengerFeedbackService",
	HandlerType: (*PassengerFeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeedback",
			Handler:    _PassengerFeedbackService_AddFeedback_Handler,
		},
		{
			MethodName: "GetFeedbackByPassengerID",
			Handler:    _PassengerFeedbackService_GetFeedbackByPassengerID_Handler,
		},
		{
			MethodName: "GetFeedbackByBookingCode",
			Handler:    _PassengerFeedbackService_GetFeedbackByBookingCode_Handler,
		},
		{
			MethodName: "DeleteFeedbackByPassengerID",
			Handler:    _PassengerFeedbackService_DeleteFeedbackByPassengerID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedback.proto",
}
