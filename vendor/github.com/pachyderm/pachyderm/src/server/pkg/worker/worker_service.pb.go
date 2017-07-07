// Code generated by protoc-gen-gogo.
// source: server/pkg/worker/worker_service.proto
// DO NOT EDIT!

/*
Package worker is a generated protocol buffer package.

It is generated from these files:
	server/pkg/worker/worker_service.proto

It has these top-level messages:
	Input
	ProcessRequest
	ProcessResponse
	CancelRequest
	CancelResponse
*/
package worker

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import pps "github.com/pachyderm/pachyderm/src/client/pps"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf "github.com/gogo/protobuf/types"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Input struct {
	FileInfo     *pfs.FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo" json:"file_info,omitempty"`
	ParentCommit *pfs.Commit   `protobuf:"bytes,5,opt,name=parent_commit,json=parentCommit" json:"parent_commit,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lazy         bool          `protobuf:"varint,3,opt,name=lazy,proto3" json:"lazy,omitempty"`
	Branch       string        `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (m *Input) Reset()                    { *m = Input{} }
func (m *Input) String() string            { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()               {}
func (*Input) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{0} }

func (m *Input) GetFileInfo() *pfs.FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

func (m *Input) GetParentCommit() *pfs.Commit {
	if m != nil {
		return m.ParentCommit
	}
	return nil
}

func (m *Input) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Input) GetLazy() bool {
	if m != nil {
		return m.Lazy
	}
	return false
}

func (m *Input) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type ProcessRequest struct {
	// ID of the job for which we're processing 'data'. This is attached to logs
	// generated while processing 'data', so that they can be searched.
	JobID string `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// The datum to process
	Data []*Input `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	// The tag corresponding to the previous parent's run of this datum, used for
	// incremental jobs, may be nil.
	ParentOutput *pfs.Tag `protobuf:"bytes,3,opt,name=parent_output,json=parentOutput" json:"parent_output,omitempty"`
}

func (m *ProcessRequest) Reset()                    { *m = ProcessRequest{} }
func (m *ProcessRequest) String() string            { return proto.CompactTextString(m) }
func (*ProcessRequest) ProtoMessage()               {}
func (*ProcessRequest) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{1} }

func (m *ProcessRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *ProcessRequest) GetData() []*Input {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ProcessRequest) GetParentOutput() *pfs.Tag {
	if m != nil {
		return m.ParentOutput
	}
	return nil
}

// ProcessResponse contains a tag, only if the processing was successful.
type ProcessResponse struct {
	Tag *pfs.Tag `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	// If true, the user program has errored
	Failed bool `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
}

func (m *ProcessResponse) Reset()                    { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string            { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()               {}
func (*ProcessResponse) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{2} }

func (m *ProcessResponse) GetTag() *pfs.Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *ProcessResponse) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

type CancelRequest struct {
	JobID       string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	DataFilters []string `protobuf:"bytes,1,rep,name=data_filters,json=dataFilters" json:"data_filters,omitempty"`
}

func (m *CancelRequest) Reset()                    { *m = CancelRequest{} }
func (m *CancelRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()               {}
func (*CancelRequest) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{3} }

func (m *CancelRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *CancelRequest) GetDataFilters() []string {
	if m != nil {
		return m.DataFilters
	}
	return nil
}

type CancelResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *CancelResponse) Reset()                    { *m = CancelResponse{} }
func (m *CancelResponse) String() string            { return proto.CompactTextString(m) }
func (*CancelResponse) ProtoMessage()               {}
func (*CancelResponse) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{4} }

func (m *CancelResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Input)(nil), "worker.Input")
	proto.RegisterType((*ProcessRequest)(nil), "worker.ProcessRequest")
	proto.RegisterType((*ProcessResponse)(nil), "worker.ProcessResponse")
	proto.RegisterType((*CancelRequest)(nil), "worker.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "worker.CancelResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Worker service

type WorkerClient interface {
	Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	Status(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := grpc.Invoke(ctx, "/worker.Worker/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Status(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error) {
	out := new(pps.WorkerStatus)
	err := grpc.Invoke(ctx, "/worker.Worker/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := grpc.Invoke(ctx, "/worker.Worker/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Worker service

type WorkerServer interface {
	Process(context.Context, *ProcessRequest) (*ProcessResponse, error)
	Status(context.Context, *google_protobuf.Empty) (*pps.WorkerStatus, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Process(ctx, req.(*ProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Status(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "worker.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _Worker_Process_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Worker_Status_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Worker_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/pkg/worker/worker_service.proto",
}

func init() { proto.RegisterFile("server/pkg/worker/worker_service.proto", fileDescriptorWorkerService) }

var fileDescriptorWorkerService = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x49, 0xe2, 0x26, 0x93, 0xa6, 0x88, 0x15, 0x04, 0xcb, 0x1c, 0x30, 0x3e, 0xa0, 0xa8,
	0x12, 0x36, 0x0a, 0xe2, 0x80, 0xc4, 0x89, 0xd2, 0x4a, 0xe1, 0x02, 0x5a, 0x2a, 0x71, 0xb4, 0xd6,
	0xce, 0xd8, 0xb8, 0x75, 0xbc, 0xc6, 0xbb, 0x06, 0x95, 0x33, 0xbf, 0xc2, 0x3f, 0xf0, 0x35, 0x1c,
	0xf8, 0x12, 0xb4, 0xb3, 0x76, 0xab, 0xc2, 0xa9, 0x07, 0xcb, 0x33, 0x6f, 0x66, 0x67, 0xde, 0x7b,
	0x1a, 0x78, 0xaa, 0xb0, 0xfd, 0x8a, 0x6d, 0xdc, 0x5c, 0x14, 0xf1, 0x37, 0xd9, 0x5e, 0x60, 0xdb,
	0xff, 0x12, 0x53, 0x28, 0x33, 0x8c, 0x9a, 0x56, 0x6a, 0xc9, 0x5c, 0x8b, 0xfa, 0xf7, 0xb3, 0xaa,
	0xc4, 0x5a, 0xc7, 0x4d, 0xae, 0xcc, 0x67, 0xab, 0xd7, 0x68, 0xa3, 0xcc, 0x37, 0xa0, 0x85, 0x2c,
	0x24, 0x85, 0xb1, 0x89, 0x7a, 0xf4, 0x51, 0x21, 0x65, 0x51, 0x61, 0x4c, 0x59, 0xda, 0xe5, 0x31,
	0xee, 0x1a, 0x7d, 0x69, 0x8b, 0xe1, 0x4f, 0x07, 0x26, 0x9b, 0xba, 0xe9, 0x34, 0x3b, 0x82, 0x59,
	0x5e, 0x56, 0x98, 0x94, 0x75, 0x2e, 0x3d, 0x27, 0x70, 0x56, 0xf3, 0xf5, 0x22, 0x32, 0x1b, 0x4f,
	0xcb, 0x0a, 0x37, 0x75, 0x2e, 0xf9, 0x34, 0xef, 0x23, 0xf6, 0x1c, 0x16, 0x8d, 0x68, 0xb1, 0xd6,
	0x49, 0x26, 0x77, 0xbb, 0x52, 0x7b, 0x13, 0xea, 0x9f, 0x53, 0xff, 0x31, 0x41, 0xfc, 0xc0, 0x76,
	0xd8, 0x8c, 0x31, 0x18, 0xd7, 0x62, 0x87, 0xde, 0x9d, 0xc0, 0x59, 0xcd, 0x38, 0xc5, 0x06, 0xab,
	0xc4, 0xf7, 0x4b, 0x6f, 0x14, 0x38, 0xab, 0x29, 0xa7, 0x98, 0x2d, 0xc1, 0x4d, 0x5b, 0x51, 0x67,
	0x9f, 0xbd, 0x31, 0x75, 0xf6, 0x59, 0xf8, 0xc3, 0x81, 0xc3, 0x0f, 0xad, 0xcc, 0x50, 0x29, 0x8e,
	0x5f, 0x3a, 0x54, 0x9a, 0x05, 0xe0, 0x9e, 0xcb, 0x34, 0x29, 0xb7, 0x76, 0xe8, 0x9b, 0xd9, 0x9f,
	0xdf, 0x8f, 0x27, 0xef, 0x64, 0xba, 0x79, 0xcb, 0x27, 0xe7, 0x32, 0xdd, 0x6c, 0xd9, 0x13, 0x18,
	0x6f, 0x85, 0x16, 0x9e, 0x13, 0x8c, 0x48, 0x8d, 0xb5, 0x34, 0x22, 0xbd, 0x9c, 0x4a, 0xec, 0xd9,
	0x95, 0x12, 0xd9, 0xe9, 0xa6, 0xd3, 0x44, 0x66, 0xbe, 0x9e, 0x92, 0x92, 0x33, 0x51, 0x0c, 0x32,
	0xde, 0x53, 0x35, 0x3c, 0x81, 0xbb, 0x57, 0x2c, 0x54, 0x23, 0x6b, 0x85, 0xcc, 0x87, 0x91, 0x16,
	0x45, 0xef, 0xd8, 0xf5, 0x3b, 0x03, 0x1a, 0x35, 0xb9, 0x28, 0x2b, 0xb4, 0x14, 0xa7, 0xbc, 0xcf,
	0xc2, 0x33, 0x58, 0x1c, 0x8b, 0x3a, 0xc3, 0xea, 0x36, 0x5a, 0x0e, 0x0c, 0xe1, 0x24, 0x2f, 0x2b,
	0x8d, 0xad, 0x22, 0x4d, 0x33, 0x3e, 0x37, 0xd8, 0xa9, 0x85, 0xc2, 0x23, 0x38, 0x1c, 0xa6, 0xf6,
	0xdc, 0x3c, 0xd8, 0x57, 0x5d, 0x66, 0xe8, 0x12, 0xbf, 0x29, 0x1f, 0xd2, 0xf5, 0x2f, 0x07, 0xdc,
	0x4f, 0x64, 0x07, 0x7b, 0x0d, 0xfb, 0xbd, 0x26, 0xb6, 0x1c, 0x2c, 0xba, 0x69, 0xb5, 0xff, 0xf0,
	0x3f, 0xdc, 0x2e, 0x08, 0xf7, 0xd8, 0x4b, 0x70, 0x3f, 0x6a, 0xa1, 0x3b, 0xf3, 0xd8, 0x1e, 0x5a,
	0x34, 0x1c, 0x5a, 0x74, 0x62, 0x0e, 0xcd, 0xbf, 0x17, 0x99, 0x0b, 0xb5, 0xcb, 0x6c, 0x6b, 0xb8,
	0xc7, 0x5e, 0x81, 0x6b, 0xb9, 0xb2, 0x07, 0xc3, 0xec, 0x1b, 0x8e, 0xf8, 0xcb, 0x7f, 0xe1, 0x61,
	0x63, 0xea, 0xd2, 0xfc, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x85, 0x1c, 0x3e, 0x4a,
	0x03, 0x00, 0x00,
}