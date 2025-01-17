// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-micro/registry/service/proto/registry.proto

package go_micro_registry

import (
	fmt "fmt"
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

// EventType defines the type of event
type EventType int32

const (
	EventType_Create EventType = 0
	EventType_Delete EventType = 1
	EventType_Update EventType = 2
)

var EventType_name = map[int32]string{
	0: "Create",
	1: "Delete",
	2: "Update",
}

var EventType_value = map[string]int32{
	"Create": 0,
	"Delete": 1,
	"Update": 2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{0}
}

// Service represents a go-micro service
type Service struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Endpoints            []*Endpoint       `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Nodes                []*Node           `protobuf:"bytes,5,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Options              *Options          `protobuf:"bytes,6,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Service) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

// Node represents the node the service is on
type Node struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string            `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port                 int64             `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Endpoint is a endpoint provided by a service
type Endpoint struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Request              *Value            `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Response             *Value            `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{2}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetRequest() *Value {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Endpoint) GetResponse() *Value {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Endpoint) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Value is an opaque value for a request or response
type Value struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Values               []*Value `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{3}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Value) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Value) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// Options are registry options
type Options struct {
	Ttl                  int64    `protobuf:"varint,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{4}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

// Result is returns by the watcher
type Result struct {
	Action               string   `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Service              *Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{5}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Result) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Result) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{6}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type GetRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{7}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type GetResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{8}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{9}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{10}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type WatchRequest struct {
	// service is optional
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{11}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

// Event is registry event
type Event struct {
	// Event Id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type of event
	Type EventType `protobuf:"varint,2,opt,name=type,proto3,enum=go.micro.registry.EventType" json:"type,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// service entry
	Service              *Service `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e39fd3e4b9b6e63, []int{12}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Create
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterEnum("go.micro.registry.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Service)(nil), "go.micro.registry.Service")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.registry.Service.MetadataEntry")
	proto.RegisterType((*Node)(nil), "go.micro.registry.Node")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.registry.Node.MetadataEntry")
	proto.RegisterType((*Endpoint)(nil), "go.micro.registry.Endpoint")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.registry.Endpoint.MetadataEntry")
	proto.RegisterType((*Value)(nil), "go.micro.registry.Value")
	proto.RegisterType((*Options)(nil), "go.micro.registry.Options")
	proto.RegisterType((*Result)(nil), "go.micro.registry.Result")
	proto.RegisterType((*EmptyResponse)(nil), "go.micro.registry.EmptyResponse")
	proto.RegisterType((*GetRequest)(nil), "go.micro.registry.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.registry.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.registry.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.registry.ListResponse")
	proto.RegisterType((*WatchRequest)(nil), "go.micro.registry.WatchRequest")
	proto.RegisterType((*Event)(nil), "go.micro.registry.Event")
}

func init() {
	proto.RegisterFile("go-micro/registry/service/proto/registry.proto", fileDescriptor_0e39fd3e4b9b6e63)
}

var fileDescriptor_0e39fd3e4b9b6e63 = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0xed, 0xfc, 0x4e, 0xda, 0x7e, 0xfd, 0x46, 0x08, 0x8c, 0x5b, 0x20, 0xb2, 0x04, 0x0a,
	0x48, 0x75, 0xaa, 0x50, 0x21, 0x7e, 0xae, 0x10, 0x0d, 0x95, 0x50, 0x0b, 0x62, 0xf9, 0xbb, 0x36,
	0xf1, 0xa8, 0x58, 0x24, 0xb6, 0xd9, 0xdd, 0x46, 0xca, 0x3b, 0x20, 0xf1, 0x04, 0xbc, 0x0d, 0x4f,
	0xc1, 0xd3, 0xa0, 0x5d, 0xaf, 0x93, 0x54, 0xb5, 0x03, 0x52, 0xe1, 0x6e, 0x66, 0xf7, 0x9c, 0xd9,
	0xd9, 0x33, 0x67, 0x6d, 0x08, 0x4e, 0xd3, 0xbd, 0x69, 0x3c, 0xe6, 0xe9, 0x80, 0xd3, 0x69, 0x2c,
	0x24, 0x9f, 0x0f, 0x04, 0xf1, 0x59, 0x3c, 0xa6, 0x41, 0xc6, 0x53, 0xb9, 0x5c, 0x0e, 0x74, 0x8a,
	0xff, 0x9f, 0xa6, 0x81, 0xc6, 0x07, 0xc5, 0x86, 0xff, 0xd3, 0x86, 0xd6, 0x9b, 0x9c, 0x83, 0x08,
	0xf5, 0x24, 0x9c, 0x92, 0x6b, 0xf5, 0xac, 0x7e, 0x87, 0xe9, 0x18, 0x5d, 0x68, 0xcd, 0x88, 0x8b,
	0x38, 0x4d, 0x5c, 0x5b, 0x2f, 0x17, 0x29, 0x1e, 0x42, 0x7b, 0x4a, 0x32, 0x8c, 0x42, 0x19, 0xba,
	0x4e, 0xcf, 0xe9, 0x77, 0x87, 0xfd, 0xe0, 0x42, 0xfd, 0xc0, 0xd4, 0x0e, 0x4e, 0x0c, 0x74, 0x94,
	0x48, 0x3e, 0x67, 0x0b, 0x26, 0x3e, 0x82, 0x0e, 0x25, 0x51, 0x96, 0xc6, 0x89, 0x14, 0x6e, 0x5d,
	0x97, 0xd9, 0x29, 0x29, 0x33, 0x32, 0x18, 0xb6, 0x44, 0xe3, 0x1e, 0x34, 0x92, 0x34, 0x22, 0xe1,
	0x36, 0x34, 0xed, 0x5a, 0x09, 0xed, 0x65, 0x1a, 0x11, 0xcb, 0x51, 0x78, 0x00, 0xad, 0x34, 0x93,
	0x71, 0x9a, 0x08, 0xb7, 0xd9, 0xb3, 0xfa, 0xdd, 0xa1, 0x57, 0x42, 0x78, 0x95, 0x23, 0x58, 0x01,
	0xf5, 0x9e, 0xc0, 0xe6, 0xb9, 0xd6, 0x71, 0x1b, 0x9c, 0xcf, 0x34, 0x37, 0x1a, 0xa9, 0x10, 0xaf,
	0x40, 0x63, 0x16, 0x4e, 0xce, 0xc8, 0x08, 0x94, 0x27, 0x8f, 0xed, 0x87, 0x96, 0xff, 0xc3, 0x82,
	0xba, 0x6a, 0x01, 0xb7, 0xc0, 0x8e, 0x23, 0xc3, 0xb1, 0xe3, 0x48, 0xa9, 0x1a, 0x46, 0x11, 0x27,
	0x21, 0x0a, 0x55, 0x4d, 0xaa, 0x66, 0x90, 0xa5, 0x5c, 0xba, 0x4e, 0xcf, 0xea, 0x3b, 0x4c, 0xc7,
	0xf8, 0x74, 0x45, 0xe9, 0x5c, 0xa2, 0xdb, 0x15, 0x77, 0xad, 0x92, 0xf9, 0x72, 0xd7, 0xf8, 0x6a,
	0x43, 0xbb, 0x18, 0x40, 0xa9, 0x49, 0x86, 0xd0, 0xe2, 0xf4, 0xe5, 0x8c, 0x84, 0xd4, 0xe4, 0xee,
	0xd0, 0x2d, 0xe9, 0xef, 0xbd, 0xaa, 0xc7, 0x0a, 0x20, 0x1e, 0x40, 0x9b, 0x93, 0xc8, 0xd2, 0x44,
	0x90, 0xbe, 0xec, 0x3a, 0xd2, 0x02, 0x89, 0xa3, 0x0b, 0x52, 0xdc, 0x5d, 0xe3, 0x96, 0x7f, 0x23,
	0x47, 0x08, 0x0d, 0xdd, 0x56, 0xa9, 0x14, 0x08, 0x75, 0x39, 0xcf, 0x0a, 0x96, 0x8e, 0x71, 0x1f,
	0x9a, 0x9a, 0x2d, 0xcc, 0x3b, 0xa9, 0xbe, 0xa8, 0xc1, 0xf9, 0x3b, 0xd0, 0x32, 0x4e, 0x54, 0x9d,
	0x49, 0x39, 0xd1, 0x67, 0x38, 0x4c, 0x85, 0xbe, 0x84, 0x26, 0x23, 0x71, 0x36, 0x91, 0x78, 0x15,
	0x9a, 0xe1, 0x58, 0xc1, 0x4c, 0x0b, 0x26, 0x53, 0x56, 0x37, 0xdf, 0x01, 0x33, 0x0f, 0xaf, 0xfa,
	0x65, 0xb2, 0x02, 0x8a, 0xbb, 0xd0, 0x91, 0xf1, 0x94, 0x84, 0x0c, 0xa7, 0x99, 0xf1, 0xdf, 0x72,
	0xc1, 0xff, 0x0f, 0x36, 0x47, 0xd3, 0x4c, 0xce, 0x99, 0x19, 0x85, 0x7f, 0x07, 0xe0, 0x88, 0x24,
	0x33, 0xe3, 0x74, 0x97, 0x47, 0xe6, 0xbd, 0x14, 0xa9, 0x3f, 0x82, 0xae, 0xc6, 0x99, 0x09, 0x3e,
	0x80, 0xb6, 0xd9, 0x11, 0xae, 0xa5, 0xe5, 0x58, 0xd7, 0xdc, 0x02, 0xeb, 0x6f, 0x42, 0xf7, 0x38,
	0x16, 0xc5, 0x79, 0xfe, 0x73, 0xd8, 0xc8, 0xd3, 0x4b, 0x96, 0xed, 0xc3, 0xc6, 0x87, 0x50, 0x8e,
	0x3f, 0xfd, 0xfe, 0x1e, 0xdf, 0x2d, 0x68, 0x8c, 0x66, 0x94, 0xc8, 0x0b, 0xaf, 0x79, 0x7f, 0x65,
	0xe6, 0x5b, 0xc3, 0xdd, 0x32, 0x43, 0x2a, 0xde, 0xdb, 0x79, 0x46, 0xc6, 0x11, 0x6b, 0xa5, 0x5e,
	0x1d, 0x5f, 0xfd, 0x8f, 0xc7, 0x77, 0x6f, 0x00, 0x9d, 0xc5, 0x31, 0x08, 0xd0, 0x7c, 0xc6, 0x29,
	0x94, 0xb4, 0x5d, 0x53, 0xf1, 0x21, 0x4d, 0x48, 0xd2, 0xb6, 0xa5, 0xe2, 0x77, 0x59, 0xa4, 0xd6,
	0xed, 0xe1, 0x37, 0x07, 0xda, 0xcc, 0x94, 0xc3, 0x13, 0x3d, 0xcd, 0xe2, 0x4f, 0x70, 0xa3, 0xe4,
	0xc0, 0xe5, 0xb0, 0xbd, 0x9b, 0x55, 0xdb, 0xc6, 0x1a, 0x35, 0x7c, 0x51, 0x94, 0x26, 0x8e, 0x6b,
	0xba, 0xf7, 0x7a, 0x65, 0x62, 0x9d, 0xb3, 0x59, 0x0d, 0x8f, 0x01, 0x0e, 0x89, 0xff, 0xad, 0x6a,
	0xaf, 0x73, 0xe3, 0x18, 0x8a, 0xc0, 0xb2, 0xbb, 0xac, 0x18, 0xcd, 0xbb, 0x55, 0xb9, 0xbf, 0x28,
	0x79, 0x04, 0x0d, 0xed, 0x21, 0x2c, 0xc3, 0xae, 0xba, 0xcb, 0xbb, 0x5e, 0x02, 0xc8, 0xdf, 0xb2,
	0x5f, 0xdb, 0xb7, 0x3e, 0x36, 0xf5, 0x6f, 0xfa, 0xfe, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c,
	0x3e, 0x92, 0x97, 0xd8, 0x07, 0x00, 0x00,
}
