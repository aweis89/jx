// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	GetLogsRequest
	GetLogsResponse
	AppArchive
	Version
	UpOptions
	UpRequest
	UpSummary
	UpMessage
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import hapi_chart "k8s.io/helm/pkg/proto/hapi/chart"
import hapi_chart3 "k8s.io/helm/pkg/proto/hapi/chart"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StatusCode is the enumeration of the possible status codes
// returned for a draft up. TODO: Flush this out.
type UpSummary_StatusCode int32

const (
	UpSummary_UNKNOWN UpSummary_StatusCode = 0
	UpSummary_LOGGING UpSummary_StatusCode = 1
	UpSummary_STARTED UpSummary_StatusCode = 2
	UpSummary_ONGOING UpSummary_StatusCode = 3
	UpSummary_SUCCESS UpSummary_StatusCode = 4
	UpSummary_FAILURE UpSummary_StatusCode = 5
)

var UpSummary_StatusCode_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOGGING",
	2: "STARTED",
	3: "ONGOING",
	4: "SUCCESS",
	5: "FAILURE",
}
var UpSummary_StatusCode_value = map[string]int32{
	"UNKNOWN": 0,
	"LOGGING": 1,
	"STARTED": 2,
	"ONGOING": 3,
	"SUCCESS": 4,
	"FAILURE": 5,
}

func (x UpSummary_StatusCode) String() string {
	return proto.EnumName(UpSummary_StatusCode_name, int32(x))
}
func (UpSummary_StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

// GetLogsRequest is the message sent to draftd to retrieve logs for a draft build.
type GetLogsRequest struct {
	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName" json:"app_name,omitempty"`
	BuildID string `protobuf:"bytes,2,opt,name=buildID" json:"buildID,omitempty"`
	Limit   int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *GetLogsRequest) Reset()                    { *m = GetLogsRequest{} }
func (m *GetLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLogsRequest) ProtoMessage()               {}
func (*GetLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetLogsRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *GetLogsRequest) GetBuildID() string {
	if m != nil {
		return m.BuildID
	}
	return ""
}

func (m *GetLogsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// GetLogsResponse is the message returned by draftd containing the logs (up to limit)
// for a draft build.
type GetLogsResponse struct {
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *GetLogsResponse) Reset()                    { *m = GetLogsResponse{} }
func (m *GetLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLogsResponse) ProtoMessage()               {}
func (*GetLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetLogsResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// AppArchive represents the archived application included in a `draft up`.
type AppArchive struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *AppArchive) Reset()                    { *m = AppArchive{} }
func (m *AppArchive) String() string            { return proto.CompactTextString(m) }
func (*AppArchive) ProtoMessage()               {}
func (*AppArchive) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppArchive) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppArchive) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// Version represents the wire description of a server's version.
type Version struct {
	SemVer       string `protobuf:"bytes,1,opt,name=sem_ver,json=semVer" json:"sem_ver,omitempty"`
	GitCommit    string `protobuf:"bytes,2,opt,name=git_commit,json=gitCommit" json:"git_commit,omitempty"`
	GitTreeState string `protobuf:"bytes,3,opt,name=git_tree_state,json=gitTreeState" json:"git_tree_state,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Version) GetSemVer() string {
	if m != nil {
		return m.SemVer
	}
	return ""
}

func (m *Version) GetGitCommit() string {
	if m != nil {
		return m.GitCommit
	}
	return ""
}

func (m *Version) GetGitTreeState() string {
	if m != nil {
		return m.GitTreeState
	}
	return ""
}

// UpOptions are configurable settings to use when issuing a `draft up`.
type UpOptions struct {
	ReleaseWait bool `protobuf:"varint,1,opt,name=release_wait,json=releaseWait" json:"release_wait,omitempty"`
}

func (m *UpOptions) Reset()                    { *m = UpOptions{} }
func (m *UpOptions) String() string            { return proto.CompactTextString(m) }
func (*UpOptions) ProtoMessage()               {}
func (*UpOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpOptions) GetReleaseWait() bool {
	if m != nil {
		return m.ReleaseWait
	}
	return false
}

// UpRequest indicates the message sent by a draft client and received
// by a draft server to carry out a draft up command.
type UpRequest struct {
	AppName    string             `protobuf:"bytes,1,opt,name=app_name,json=appName" json:"app_name,omitempty"`
	Namespace  string             `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Options    *UpOptions         `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	Chart      *hapi_chart3.Chart `protobuf:"bytes,4,opt,name=chart" json:"chart,omitempty"`
	Values     *hapi_chart.Config `protobuf:"bytes,5,opt,name=values" json:"values,omitempty"`
	AppArchive *AppArchive        `protobuf:"bytes,6,opt,name=app_archive,json=appArchive" json:"app_archive,omitempty"`
}

func (m *UpRequest) Reset()                    { *m = UpRequest{} }
func (m *UpRequest) String() string            { return proto.CompactTextString(m) }
func (*UpRequest) ProtoMessage()               {}
func (*UpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *UpRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpRequest) GetOptions() *UpOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *UpRequest) GetChart() *hapi_chart3.Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

func (m *UpRequest) GetValues() *hapi_chart.Config {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *UpRequest) GetAppArchive() *AppArchive {
	if m != nil {
		return m.AppArchive
	}
	return nil
}

// UpSummary is the message returned when executing a draft up.
type UpSummary struct {
	// stage_desc describes the particular stage this summary
	// represents, e.g. "Build Docker Image." This is meant
	// to be a canonical summary of the stage's intent.
	StageDesc string `protobuf:"bytes,1,opt,name=stage_desc,json=stageDesc" json:"stage_desc,omitempty"`
	// status_text indicates a string description of the progress
	// or completion of draft up.
	StatusText string `protobuf:"bytes,2,opt,name=status_text,json=statusText" json:"status_text,omitempty"`
	// status_code indicates the status of the progress or
	// completion of a draft up.
	StatusCode UpSummary_StatusCode `protobuf:"varint,3,opt,name=status_code,json=statusCode,enum=rpc.UpSummary_StatusCode" json:"status_code,omitempty"`
	// build_id is the build identifier associated with this draft up build.
	BuildId string `protobuf:"bytes,4,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
}

func (m *UpSummary) Reset()                    { *m = UpSummary{} }
func (m *UpSummary) String() string            { return proto.CompactTextString(m) }
func (*UpSummary) ProtoMessage()               {}
func (*UpSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpSummary) GetStageDesc() string {
	if m != nil {
		return m.StageDesc
	}
	return ""
}

func (m *UpSummary) GetStatusText() string {
	if m != nil {
		return m.StatusText
	}
	return ""
}

func (m *UpSummary) GetStatusCode() UpSummary_StatusCode {
	if m != nil {
		return m.StatusCode
	}
	return UpSummary_UNKNOWN
}

func (m *UpSummary) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

// UpMessage is the discriminated union of draft up messages
// exchanged between the client and server.
type UpMessage struct {
	// Types that are valid to be assigned to Message:
	//	*UpMessage_UpRequest
	//	*UpMessage_UpSummary
	Message isUpMessage_Message `protobuf_oneof:"Message"`
}

func (m *UpMessage) Reset()                    { *m = UpMessage{} }
func (m *UpMessage) String() string            { return proto.CompactTextString(m) }
func (*UpMessage) ProtoMessage()               {}
func (*UpMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isUpMessage_Message interface {
	isUpMessage_Message()
}

type UpMessage_UpRequest struct {
	UpRequest *UpRequest `protobuf:"bytes,1,opt,name=up_request,json=upRequest,oneof"`
}
type UpMessage_UpSummary struct {
	UpSummary *UpSummary `protobuf:"bytes,2,opt,name=up_summary,json=upSummary,oneof"`
}

func (*UpMessage_UpRequest) isUpMessage_Message() {}
func (*UpMessage_UpSummary) isUpMessage_Message() {}

func (m *UpMessage) GetMessage() isUpMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *UpMessage) GetUpRequest() *UpRequest {
	if x, ok := m.GetMessage().(*UpMessage_UpRequest); ok {
		return x.UpRequest
	}
	return nil
}

func (m *UpMessage) GetUpSummary() *UpSummary {
	if x, ok := m.GetMessage().(*UpMessage_UpSummary); ok {
		return x.UpSummary
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UpMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UpMessage_OneofMarshaler, _UpMessage_OneofUnmarshaler, _UpMessage_OneofSizer, []interface{}{
		(*UpMessage_UpRequest)(nil),
		(*UpMessage_UpSummary)(nil),
	}
}

func _UpMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UpMessage)
	// Message
	switch x := m.Message.(type) {
	case *UpMessage_UpRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpRequest); err != nil {
			return err
		}
	case *UpMessage_UpSummary:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpSummary); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UpMessage.Message has unexpected type %T", x)
	}
	return nil
}

func _UpMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UpMessage)
	switch tag {
	case 1: // Message.up_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpRequest)
		err := b.DecodeMessage(msg)
		m.Message = &UpMessage_UpRequest{msg}
		return true, err
	case 2: // Message.up_summary
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpSummary)
		err := b.DecodeMessage(msg)
		m.Message = &UpMessage_UpSummary{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UpMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UpMessage)
	// Message
	switch x := m.Message.(type) {
	case *UpMessage_UpRequest:
		s := proto.Size(x.UpRequest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpMessage_UpSummary:
		s := proto.Size(x.UpSummary)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*GetLogsRequest)(nil), "rpc.GetLogsRequest")
	proto.RegisterType((*GetLogsResponse)(nil), "rpc.GetLogsResponse")
	proto.RegisterType((*AppArchive)(nil), "rpc.AppArchive")
	proto.RegisterType((*Version)(nil), "rpc.Version")
	proto.RegisterType((*UpOptions)(nil), "rpc.UpOptions")
	proto.RegisterType((*UpRequest)(nil), "rpc.UpRequest")
	proto.RegisterType((*UpSummary)(nil), "rpc.UpSummary")
	proto.RegisterType((*UpMessage)(nil), "rpc.UpMessage")
	proto.RegisterEnum("rpc.UpSummary_StatusCode", UpSummary_StatusCode_name, UpSummary_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Draft service

type DraftClient interface {
	// GetVersion returns the current version of the server.
	GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Version, error)
	// GetLogs returns the logs for a particular draft build.
	GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error)
	// UpBuild returns a stream of the UpSummary within
	// for a given draft upload.
	//
	// Results are streamed rather than returned at once so
	// each stage of a draft up may be distinguished by the
	// client as a distinct message.
	UpBuild(ctx context.Context, in *UpMessage, opts ...grpc.CallOption) (Draft_UpBuildClient, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of UpMessages each representing a separate
	// draft up. This is the api invoked by the draft client when
	// doing a draft up with watch enabled.
	UpStream(ctx context.Context, opts ...grpc.CallOption) (Draft_UpStreamClient, error)
}

type draftClient struct {
	cc *grpc.ClientConn
}

func NewDraftClient(cc *grpc.ClientConn) DraftClient {
	return &draftClient{cc}
}

func (c *draftClient) GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := grpc.Invoke(ctx, "/rpc.Draft/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *draftClient) GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error) {
	out := new(GetLogsResponse)
	err := grpc.Invoke(ctx, "/rpc.Draft/GetLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *draftClient) UpBuild(ctx context.Context, in *UpMessage, opts ...grpc.CallOption) (Draft_UpBuildClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Draft_serviceDesc.Streams[0], c.cc, "/rpc.Draft/UpBuild", opts...)
	if err != nil {
		return nil, err
	}
	x := &draftUpBuildClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Draft_UpBuildClient interface {
	Recv() (*UpMessage, error)
	grpc.ClientStream
}

type draftUpBuildClient struct {
	grpc.ClientStream
}

func (x *draftUpBuildClient) Recv() (*UpMessage, error) {
	m := new(UpMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *draftClient) UpStream(ctx context.Context, opts ...grpc.CallOption) (Draft_UpStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Draft_serviceDesc.Streams[1], c.cc, "/rpc.Draft/UpStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &draftUpStreamClient{stream}
	return x, nil
}

type Draft_UpStreamClient interface {
	Send(*UpMessage) error
	Recv() (*UpMessage, error)
	grpc.ClientStream
}

type draftUpStreamClient struct {
	grpc.ClientStream
}

func (x *draftUpStreamClient) Send(m *UpMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *draftUpStreamClient) Recv() (*UpMessage, error) {
	m := new(UpMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Draft service

type DraftServer interface {
	// GetVersion returns the current version of the server.
	GetVersion(context.Context, *google_protobuf.Empty) (*Version, error)
	// GetLogs returns the logs for a particular draft build.
	GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error)
	// UpBuild returns a stream of the UpSummary within
	// for a given draft upload.
	//
	// Results are streamed rather than returned at once so
	// each stage of a draft up may be distinguished by the
	// client as a distinct message.
	UpBuild(*UpMessage, Draft_UpBuildServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of UpMessages each representing a separate
	// draft up. This is the api invoked by the draft client when
	// doing a draft up with watch enabled.
	UpStream(Draft_UpStreamServer) error
}

func RegisterDraftServer(s *grpc.Server, srv DraftServer) {
	s.RegisterService(&_Draft_serviceDesc, srv)
}

func _Draft_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DraftServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Draft/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DraftServer).GetVersion(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Draft_GetLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DraftServer).GetLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Draft/GetLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DraftServer).GetLogs(ctx, req.(*GetLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Draft_UpBuild_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DraftServer).UpBuild(m, &draftUpBuildServer{stream})
}

type Draft_UpBuildServer interface {
	Send(*UpMessage) error
	grpc.ServerStream
}

type draftUpBuildServer struct {
	grpc.ServerStream
}

func (x *draftUpBuildServer) Send(m *UpMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Draft_UpStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DraftServer).UpStream(&draftUpStreamServer{stream})
}

type Draft_UpStreamServer interface {
	Send(*UpMessage) error
	Recv() (*UpMessage, error)
	grpc.ServerStream
}

type draftUpStreamServer struct {
	grpc.ServerStream
}

func (x *draftUpStreamServer) Send(m *UpMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *draftUpStreamServer) Recv() (*UpMessage, error) {
	m := new(UpMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Draft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Draft",
	HandlerType: (*DraftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Draft_GetVersion_Handler,
		},
		{
			MethodName: "GetLogs",
			Handler:    _Draft_GetLogs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpBuild",
			Handler:       _Draft_UpBuild_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpStream",
			Handler:       _Draft_UpStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x6f, 0xda, 0x4e,
	0x10, 0xc5, 0x21, 0xe0, 0x78, 0x40, 0x84, 0xdf, 0xfe, 0xa2, 0x84, 0xd0, 0x56, 0x4d, 0xad, 0x4a,
	0x45, 0xad, 0x6a, 0x50, 0x5a, 0xf5, 0x90, 0x1b, 0x01, 0x4a, 0xa3, 0xa6, 0x20, 0x19, 0x48, 0xa4,
	0xf4, 0x60, 0x6d, 0xcc, 0xc4, 0xb1, 0x84, 0xf1, 0xd6, 0xbb, 0x4e, 0x93, 0x2f, 0xda, 0x6f, 0xd2,
	0x6b, 0x55, 0xed, 0xae, 0x0d, 0xa1, 0x97, 0xf6, 0x62, 0x79, 0xde, 0xcc, 0xdb, 0x7d, 0xf3, 0x67,
	0x07, 0xac, 0x84, 0xf9, 0x0e, 0x4b, 0x62, 0x11, 0x93, 0x62, 0xc2, 0xfc, 0xe6, 0x93, 0x20, 0x8e,
	0x83, 0x05, 0xb6, 0x15, 0x74, 0x9d, 0xde, 0xb4, 0x31, 0x62, 0xe2, 0x41, 0x47, 0x34, 0x0f, 0x6e,
	0x29, 0x0b, 0xdb, 0xfe, 0x2d, 0x4d, 0x44, 0xdb, 0x8f, 0x97, 0x37, 0x61, 0x90, 0x39, 0xf6, 0x1f,
	0x3b, 0xe4, 0x57, 0xe3, 0xf6, 0x57, 0xa8, 0x0d, 0x51, 0x9c, 0xc7, 0x01, 0x77, 0xf1, 0x5b, 0x8a,
	0x5c, 0x90, 0x43, 0xd8, 0xa1, 0x8c, 0x79, 0x4b, 0x1a, 0x61, 0xc3, 0x38, 0x32, 0x5a, 0x96, 0x6b,
	0x52, 0xc6, 0x46, 0x34, 0x42, 0xd2, 0x00, 0xf3, 0x3a, 0x0d, 0x17, 0xf3, 0xb3, 0x7e, 0x63, 0x4b,
	0x7b, 0x32, 0x93, 0xec, 0x41, 0x69, 0x11, 0x46, 0xa1, 0x68, 0x14, 0x8f, 0x8c, 0x56, 0xd1, 0xd5,
	0x86, 0xfd, 0x06, 0x76, 0x57, 0x87, 0x73, 0x16, 0x2f, 0xb9, 0x3a, 0xc2, 0x8f, 0x97, 0x02, 0x97,
	0x42, 0x1d, 0x5e, 0x75, 0x73, 0xd3, 0x3e, 0x01, 0xe8, 0x32, 0xd6, 0x4d, 0xfc, 0xdb, 0xf0, 0x0e,
	0x09, 0x81, 0xed, 0x47, 0x0a, 0xd4, 0xff, 0x63, 0xee, 0xd6, 0x26, 0x37, 0x00, 0xf3, 0x02, 0x13,
	0x1e, 0xc6, 0x4b, 0x72, 0x00, 0x26, 0xc7, 0xc8, 0xbb, 0xc3, 0x24, 0xe3, 0x96, 0x39, 0x46, 0x17,
	0x98, 0x90, 0x67, 0x00, 0x41, 0x28, 0x3c, 0x3f, 0x8e, 0xa4, 0x4e, 0xad, 0xdf, 0x0a, 0x42, 0xd1,
	0x53, 0x00, 0x79, 0x09, 0x35, 0xe9, 0x16, 0x09, 0xa2, 0xc7, 0x05, 0x15, 0xa8, 0x52, 0xb1, 0xdc,
	0x6a, 0x10, 0x8a, 0x69, 0x82, 0x38, 0x91, 0x98, 0xed, 0x80, 0x35, 0x63, 0x63, 0x26, 0xc2, 0x78,
	0xc9, 0xc9, 0x0b, 0xa8, 0x26, 0xb8, 0x40, 0xca, 0xd1, 0xfb, 0x4e, 0x43, 0x9d, 0xd0, 0x8e, 0x5b,
	0xc9, 0xb0, 0x4b, 0x1a, 0x0a, 0xfb, 0xa7, 0x21, 0x09, 0xff, 0x50, 0xda, 0xa7, 0x60, 0x49, 0x98,
	0x33, 0xea, 0x63, 0x2e, 0x6e, 0x05, 0x90, 0x16, 0x98, 0xb1, 0xbe, 0x54, 0xa9, 0xaa, 0x1c, 0xd7,
	0x1c, 0x39, 0x15, 0x2b, 0x29, 0x6e, 0xee, 0x26, 0xaf, 0xa0, 0xa4, 0xda, 0xdb, 0xd8, 0x56, 0x71,
	0xff, 0x39, 0xb2, 0xef, 0x8e, 0xee, 0x78, 0x4f, 0x7e, 0x5d, 0xed, 0x27, 0xaf, 0xa1, 0x7c, 0x47,
	0x17, 0x29, 0xf2, 0x46, 0x49, 0x45, 0x92, 0x8d, 0x48, 0x35, 0x3a, 0x6e, 0x16, 0x41, 0x3a, 0x50,
	0x91, 0xba, 0xa9, 0xee, 0x4d, 0xa3, 0xac, 0x08, 0xbb, 0x4a, 0xc2, 0xba, 0x65, 0x2e, 0xd0, 0xd5,
	0xbf, 0xfd, 0x4b, 0xe5, 0x3d, 0x49, 0xa3, 0x88, 0x26, 0x0f, 0xb2, 0xf4, 0x5c, 0xd0, 0x00, 0xbd,
	0x39, 0x72, 0x3f, 0xcb, 0xdc, 0x52, 0x48, 0x1f, 0xb9, 0x4f, 0x9e, 0x43, 0x45, 0x56, 0x3c, 0xe5,
	0x9e, 0xc0, 0xfb, 0xbc, 0x35, 0xa0, 0xa1, 0x29, 0xde, 0x0b, 0x72, 0xb2, 0x0a, 0xf0, 0xe3, 0xb9,
	0x6e, 0x4c, 0xed, 0xf8, 0x30, 0x2b, 0x41, 0x76, 0x89, 0x33, 0x51, 0x11, 0xbd, 0x78, 0x8e, 0x39,
	0x57, 0xfe, 0xcb, 0x9a, 0xab, 0x21, 0xf5, 0xc2, 0xb9, 0xaa, 0xc9, 0x6a, 0x68, 0xe7, 0xf6, 0x15,
	0xc0, 0x9a, 0x44, 0x2a, 0x60, 0xce, 0x46, 0x9f, 0x47, 0xe3, 0xcb, 0x51, 0xbd, 0x20, 0x8d, 0xf3,
	0xf1, 0x70, 0x78, 0x36, 0x1a, 0xd6, 0x0d, 0x69, 0x4c, 0xa6, 0x5d, 0x77, 0x3a, 0xe8, 0xd7, 0xb7,
	0xa4, 0x31, 0x1e, 0x0d, 0xc7, 0xd2, 0x53, 0x54, 0x9e, 0x59, 0xaf, 0x37, 0x98, 0x4c, 0xea, 0xdb,
	0xd2, 0xf8, 0xd8, 0x3d, 0x3b, 0x9f, 0xb9, 0x83, 0x7a, 0xc9, 0xbe, 0x97, 0xf9, 0x7f, 0x41, 0xce,
	0x69, 0x80, 0xa4, 0x0d, 0x90, 0x32, 0x2f, 0xd1, 0x53, 0xa0, 0xf2, 0x5f, 0x77, 0x30, 0x9b, 0x8d,
	0x4f, 0x05, 0xd7, 0x4a, 0x57, 0x83, 0xa2, 0x09, 0x5c, 0x67, 0xa6, 0x0a, 0xb2, 0x26, 0x64, 0xf9,
	0x6a, 0x42, 0x66, 0x9c, 0x5a, 0x60, 0x66, 0x97, 0x1d, 0xff, 0x30, 0xa0, 0xd4, 0x4f, 0xe8, 0x8d,
	0x20, 0xef, 0x01, 0x86, 0x28, 0xf2, 0x87, 0xb1, 0xef, 0xe8, 0xc5, 0xe1, 0xe4, 0x8b, 0xc3, 0x19,
	0xc8, 0xc5, 0xd1, 0xac, 0xaa, 0x73, 0xb3, 0x28, 0xbb, 0x40, 0x3e, 0x80, 0x99, 0x3d, 0x5a, 0xf2,
	0xbf, 0x72, 0x6d, 0xee, 0x87, 0xe6, 0xde, 0x26, 0xa8, 0xdf, 0xb5, 0x5d, 0x20, 0x6f, 0xc1, 0x9c,
	0xb1, 0x53, 0x59, 0x5a, 0x92, 0x4b, 0xcd, 0x24, 0x35, 0xff, 0xb0, 0xed, 0x42, 0xc7, 0x20, 0x1d,
	0xd8, 0x99, 0xb1, 0x89, 0x48, 0x90, 0x46, 0x7f, 0x8f, 0x6f, 0x19, 0x1d, 0xe3, 0xb4, 0x74, 0x25,
	0xf7, 0xdf, 0x75, 0x59, 0xe9, 0x7f, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x94, 0x44, 0xd6, 0xe6,
	0x18, 0x05, 0x00, 0x00,
}