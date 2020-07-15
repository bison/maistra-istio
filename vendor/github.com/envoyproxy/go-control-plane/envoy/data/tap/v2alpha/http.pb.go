// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/http.proto

package envoy_data_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
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

type HttpBufferedTrace struct {
	Request              *HttpBufferedTrace_Message `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *HttpBufferedTrace_Message `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpBufferedTrace) Reset()         { *m = HttpBufferedTrace{} }
func (m *HttpBufferedTrace) String() string { return proto.CompactTextString(m) }
func (*HttpBufferedTrace) ProtoMessage()    {}
func (*HttpBufferedTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d8a92b44eb7244, []int{0}
}

func (m *HttpBufferedTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpBufferedTrace.Unmarshal(m, b)
}
func (m *HttpBufferedTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpBufferedTrace.Marshal(b, m, deterministic)
}
func (m *HttpBufferedTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBufferedTrace.Merge(m, src)
}
func (m *HttpBufferedTrace) XXX_Size() int {
	return xxx_messageInfo_HttpBufferedTrace.Size(m)
}
func (m *HttpBufferedTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBufferedTrace.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBufferedTrace proto.InternalMessageInfo

func (m *HttpBufferedTrace) GetRequest() *HttpBufferedTrace_Message {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *HttpBufferedTrace) GetResponse() *HttpBufferedTrace_Message {
	if m != nil {
		return m.Response
	}
	return nil
}

type HttpBufferedTrace_Message struct {
	Headers              []*core.HeaderValue `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	Body                 *Body               `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Trailers             []*core.HeaderValue `protobuf:"bytes,3,rep,name=trailers,proto3" json:"trailers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HttpBufferedTrace_Message) Reset()         { *m = HttpBufferedTrace_Message{} }
func (m *HttpBufferedTrace_Message) String() string { return proto.CompactTextString(m) }
func (*HttpBufferedTrace_Message) ProtoMessage()    {}
func (*HttpBufferedTrace_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d8a92b44eb7244, []int{0, 0}
}

func (m *HttpBufferedTrace_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpBufferedTrace_Message.Unmarshal(m, b)
}
func (m *HttpBufferedTrace_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpBufferedTrace_Message.Marshal(b, m, deterministic)
}
func (m *HttpBufferedTrace_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBufferedTrace_Message.Merge(m, src)
}
func (m *HttpBufferedTrace_Message) XXX_Size() int {
	return xxx_messageInfo_HttpBufferedTrace_Message.Size(m)
}
func (m *HttpBufferedTrace_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBufferedTrace_Message.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBufferedTrace_Message proto.InternalMessageInfo

func (m *HttpBufferedTrace_Message) GetHeaders() []*core.HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpBufferedTrace_Message) GetBody() *Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HttpBufferedTrace_Message) GetTrailers() []*core.HeaderValue {
	if m != nil {
		return m.Trailers
	}
	return nil
}

type HttpStreamedTraceSegment struct {
	TraceId uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Types that are valid to be assigned to MessagePiece:
	//	*HttpStreamedTraceSegment_RequestHeaders
	//	*HttpStreamedTraceSegment_RequestBodyChunk
	//	*HttpStreamedTraceSegment_RequestTrailers
	//	*HttpStreamedTraceSegment_ResponseHeaders
	//	*HttpStreamedTraceSegment_ResponseBodyChunk
	//	*HttpStreamedTraceSegment_ResponseTrailers
	MessagePiece         isHttpStreamedTraceSegment_MessagePiece `protobuf_oneof:"message_piece"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *HttpStreamedTraceSegment) Reset()         { *m = HttpStreamedTraceSegment{} }
func (m *HttpStreamedTraceSegment) String() string { return proto.CompactTextString(m) }
func (*HttpStreamedTraceSegment) ProtoMessage()    {}
func (*HttpStreamedTraceSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d8a92b44eb7244, []int{1}
}

func (m *HttpStreamedTraceSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpStreamedTraceSegment.Unmarshal(m, b)
}
func (m *HttpStreamedTraceSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpStreamedTraceSegment.Marshal(b, m, deterministic)
}
func (m *HttpStreamedTraceSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStreamedTraceSegment.Merge(m, src)
}
func (m *HttpStreamedTraceSegment) XXX_Size() int {
	return xxx_messageInfo_HttpStreamedTraceSegment.Size(m)
}
func (m *HttpStreamedTraceSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStreamedTraceSegment.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStreamedTraceSegment proto.InternalMessageInfo

func (m *HttpStreamedTraceSegment) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

type isHttpStreamedTraceSegment_MessagePiece interface {
	isHttpStreamedTraceSegment_MessagePiece()
}

type HttpStreamedTraceSegment_RequestHeaders struct {
	RequestHeaders *core.HeaderMap `protobuf:"bytes,2,opt,name=request_headers,json=requestHeaders,proto3,oneof"`
}

type HttpStreamedTraceSegment_RequestBodyChunk struct {
	RequestBodyChunk *Body `protobuf:"bytes,3,opt,name=request_body_chunk,json=requestBodyChunk,proto3,oneof"`
}

type HttpStreamedTraceSegment_RequestTrailers struct {
	RequestTrailers *core.HeaderMap `protobuf:"bytes,4,opt,name=request_trailers,json=requestTrailers,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseHeaders struct {
	ResponseHeaders *core.HeaderMap `protobuf:"bytes,5,opt,name=response_headers,json=responseHeaders,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseBodyChunk struct {
	ResponseBodyChunk *Body `protobuf:"bytes,6,opt,name=response_body_chunk,json=responseBodyChunk,proto3,oneof"`
}

type HttpStreamedTraceSegment_ResponseTrailers struct {
	ResponseTrailers *core.HeaderMap `protobuf:"bytes,7,opt,name=response_trailers,json=responseTrailers,proto3,oneof"`
}

func (*HttpStreamedTraceSegment_RequestHeaders) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_RequestBodyChunk) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_RequestTrailers) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseHeaders) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseBodyChunk) isHttpStreamedTraceSegment_MessagePiece() {}

func (*HttpStreamedTraceSegment_ResponseTrailers) isHttpStreamedTraceSegment_MessagePiece() {}

func (m *HttpStreamedTraceSegment) GetMessagePiece() isHttpStreamedTraceSegment_MessagePiece {
	if m != nil {
		return m.MessagePiece
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetRequestHeaders() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestHeaders); ok {
		return x.RequestHeaders
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetRequestBodyChunk() *Body {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestBodyChunk); ok {
		return x.RequestBodyChunk
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetRequestTrailers() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestTrailers); ok {
		return x.RequestTrailers
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseHeaders() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseHeaders); ok {
		return x.ResponseHeaders
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseBodyChunk() *Body {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseBodyChunk); ok {
		return x.ResponseBodyChunk
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseTrailers() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseTrailers); ok {
		return x.ResponseTrailers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpStreamedTraceSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpStreamedTraceSegment_RequestHeaders)(nil),
		(*HttpStreamedTraceSegment_RequestBodyChunk)(nil),
		(*HttpStreamedTraceSegment_RequestTrailers)(nil),
		(*HttpStreamedTraceSegment_ResponseHeaders)(nil),
		(*HttpStreamedTraceSegment_ResponseBodyChunk)(nil),
		(*HttpStreamedTraceSegment_ResponseTrailers)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpBufferedTrace)(nil), "envoy.data.tap.v2alpha.HttpBufferedTrace")
	proto.RegisterType((*HttpBufferedTrace_Message)(nil), "envoy.data.tap.v2alpha.HttpBufferedTrace.Message")
	proto.RegisterType((*HttpStreamedTraceSegment)(nil), "envoy.data.tap.v2alpha.HttpStreamedTraceSegment")
}

func init() { proto.RegisterFile("envoy/data/tap/v2alpha/http.proto", fileDescriptor_90d8a92b44eb7244) }

var fileDescriptor_90d8a92b44eb7244 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x97, 0xb6, 0x34, 0xc3, 0x13, 0x6c, 0x33, 0x12, 0x0a, 0xd5, 0x40, 0x63, 0x70, 0xd8,
	0xc9, 0x81, 0x70, 0x41, 0x70, 0x0b, 0x07, 0x32, 0x8d, 0xa2, 0x29, 0x9b, 0xb8, 0x46, 0xff, 0x26,
	0xff, 0xad, 0x11, 0x4d, 0x6c, 0x6c, 0xa7, 0xa2, 0x37, 0x1e, 0x81, 0x57, 0xe1, 0xca, 0x13, 0x70,
	0xe5, 0x19, 0x78, 0x11, 0x94, 0xc4, 0xce, 0x40, 0xa5, 0x22, 0xda, 0xd1, 0xce, 0xf7, 0x7d, 0xfe,
	0x7e, 0xff, 0xd8, 0xe4, 0x31, 0x96, 0x4b, 0xbe, 0xf2, 0x33, 0xd0, 0xe0, 0x6b, 0x10, 0xfe, 0x32,
	0x80, 0x85, 0x98, 0x83, 0x3f, 0xd7, 0x5a, 0x30, 0x21, 0xb9, 0xe6, 0xf4, 0x7e, 0x23, 0x61, 0xb5,
	0x84, 0x69, 0x10, 0xcc, 0x48, 0x26, 0x07, 0xad, 0x15, 0x44, 0xee, 0x2f, 0x03, 0x3f, 0xe5, 0x12,
	0xfd, 0x19, 0x28, 0x6c, 0x5d, 0x93, 0x27, 0x1b, 0x82, 0x53, 0x5e, 0x14, 0xbc, 0x34, 0xa2, 0x87,
	0x55, 0x26, 0xc0, 0x87, 0xb2, 0xe4, 0x1a, 0x74, 0xce, 0x4b, 0xe5, 0x2b, 0x0d, 0xba, 0x52, 0xed,
	0xe7, 0xa3, 0x5f, 0x03, 0xb2, 0x1f, 0x69, 0x2d, 0xc2, 0xea, 0xf2, 0x12, 0x25, 0x66, 0x17, 0x12,
	0x52, 0xa4, 0xa7, 0xc4, 0x95, 0xf8, 0xa9, 0x42, 0xa5, 0x3d, 0xe7, 0xd0, 0x39, 0xde, 0x09, 0x9e,
	0xb3, 0x7f, 0x37, 0x64, 0x6b, 0x5e, 0x36, 0x45, 0xa5, 0xe0, 0x0a, 0x63, 0x9b, 0x40, 0xa7, 0x64,
	0x5b, 0xa2, 0x12, 0xbc, 0x54, 0xe8, 0x0d, 0x6e, 0x9a, 0xd6, 0x45, 0x4c, 0xbe, 0x39, 0xc4, 0x35,
	0xbb, 0xf4, 0x25, 0x71, 0xe7, 0x08, 0x19, 0x4a, 0xe5, 0x39, 0x87, 0xc3, 0xe3, 0x9d, 0xe0, 0x91,
	0x49, 0x06, 0x91, 0xb3, 0x65, 0xc0, 0xea, 0x89, 0xb1, 0xa8, 0x51, 0x7c, 0x80, 0x45, 0x85, 0xb1,
	0x95, 0xd3, 0x67, 0x64, 0x34, 0xe3, 0xd9, 0xca, 0x14, 0x3a, 0xd8, 0x54, 0x28, 0xe4, 0xd9, 0x2a,
	0x6e, 0x94, 0xf4, 0x15, 0xd9, 0xd6, 0x12, 0xf2, 0x45, 0x7d, 0xd8, 0xb0, 0xd7, 0x61, 0x9d, 0xfe,
	0xe8, 0xeb, 0x88, 0x78, 0x35, 0xdb, 0xb9, 0x96, 0x08, 0x85, 0x61, 0x3b, 0xc7, 0xab, 0x02, 0x4b,
	0x4d, 0x1f, 0x34, 0xc1, 0x29, 0x26, 0x79, 0xd6, 0x4c, 0x7b, 0x14, 0xbb, 0xcd, 0xfa, 0x24, 0xa3,
	0x6f, 0xc9, 0xae, 0x99, 0x62, 0x62, 0x39, 0xff, 0x2e, 0xbc, 0x7e, 0xf4, 0x14, 0x44, 0xb4, 0x15,
	0xdf, 0x35, 0xb6, 0xc8, 0xe0, 0xbe, 0x23, 0xd4, 0x06, 0xd5, 0x30, 0x49, 0x3a, 0xaf, 0xca, 0x8f,
	0xde, 0xf0, 0xff, 0xf0, 0xd1, 0x56, 0xbc, 0x67, 0x9c, 0xf5, 0xf2, 0x4d, 0xed, 0xa3, 0x27, 0xc4,
	0xee, 0x25, 0xdd, 0x48, 0x46, 0xbd, 0x7a, 0x59, 0x9c, 0x0b, 0x63, 0x6b, 0xa3, 0xda, 0x3f, 0xdb,
	0x21, 0xde, 0xea, 0x1b, 0xd5, 0xfa, 0x2c, 0xe3, 0x7b, 0x72, 0xaf, 0x8b, 0xfa, 0x03, 0x72, 0xdc,
	0x0b, 0x72, 0xdf, 0x5a, 0xaf, 0x29, 0x4f, 0x49, 0xb7, 0x79, 0x8d, 0xe9, 0xf6, 0xea, 0xd6, 0x31,
	0x59, 0xce, 0x70, 0x97, 0xdc, 0x29, 0xda, 0x4b, 0x9b, 0x88, 0x1c, 0x53, 0x0c, 0x5f, 0x7f, 0xff,
	0xf2, 0xe3, 0xe7, 0x78, 0xb0, 0xe7, 0x90, 0xa7, 0x39, 0x6f, 0xe3, 0x84, 0xe4, 0x9f, 0x57, 0x1b,
	0x7a, 0x86, 0xb7, 0xeb, 0xfb, 0x73, 0x56, 0xbf, 0xd9, 0x33, 0x67, 0x36, 0x6e, 0x1e, 0xef, 0x8b,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x0d, 0x2f, 0x04, 0x5b, 0x04, 0x00, 0x00,
}