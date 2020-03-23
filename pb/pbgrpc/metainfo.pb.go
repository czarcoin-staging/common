// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metainfo.proto

package pbgrpc

import (
	context "context"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"

	. "storj.io/common/pb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetainfoClient is the client API for Metainfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetainfoClient interface {
	// Bucket
	CreateBucket(ctx context.Context, in *BucketCreateRequest, opts ...grpc.CallOption) (*BucketCreateResponse, error)
	GetBucket(ctx context.Context, in *BucketGetRequest, opts ...grpc.CallOption) (*BucketGetResponse, error)
	DeleteBucket(ctx context.Context, in *BucketDeleteRequest, opts ...grpc.CallOption) (*BucketDeleteResponse, error)
	ListBuckets(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error)
	SetBucketAttribution(ctx context.Context, in *BucketSetAttributionRequest, opts ...grpc.CallOption) (*BucketSetAttributionResponse, error)
	// Object
	BeginObject(ctx context.Context, in *ObjectBeginRequest, opts ...grpc.CallOption) (*ObjectBeginResponse, error)
	CommitObject(ctx context.Context, in *ObjectCommitRequest, opts ...grpc.CallOption) (*ObjectCommitResponse, error)
	GetObject(ctx context.Context, in *ObjectGetRequest, opts ...grpc.CallOption) (*ObjectGetResponse, error)
	ListObjects(ctx context.Context, in *ObjectListRequest, opts ...grpc.CallOption) (*ObjectListResponse, error)
	BeginDeleteObject(ctx context.Context, in *ObjectBeginDeleteRequest, opts ...grpc.CallOption) (*ObjectBeginDeleteResponse, error)
	FinishDeleteObject(ctx context.Context, in *ObjectFinishDeleteRequest, opts ...grpc.CallOption) (*ObjectFinishDeleteResponse, error)
	BeginSegment(ctx context.Context, in *SegmentBeginRequest, opts ...grpc.CallOption) (*SegmentBeginResponse, error)
	CommitSegment(ctx context.Context, in *SegmentCommitRequest, opts ...grpc.CallOption) (*SegmentCommitResponse, error)
	MakeInlineSegment(ctx context.Context, in *SegmentMakeInlineRequest, opts ...grpc.CallOption) (*SegmentMakeInlineResponse, error)
	BeginDeleteSegment(ctx context.Context, in *SegmentBeginDeleteRequest, opts ...grpc.CallOption) (*SegmentBeginDeleteResponse, error)
	FinishDeleteSegment(ctx context.Context, in *SegmentFinishDeleteRequest, opts ...grpc.CallOption) (*SegmentFinishDeleteResponse, error)
	ListSegments(ctx context.Context, in *SegmentListRequest, opts ...grpc.CallOption) (*SegmentListResponse, error)
	DownloadSegment(ctx context.Context, in *SegmentDownloadRequest, opts ...grpc.CallOption) (*SegmentDownloadResponse, error)
	Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error)
	CreateSegmentOld(ctx context.Context, in *SegmentWriteRequestOld, opts ...grpc.CallOption) (*SegmentWriteResponseOld, error)
	CommitSegmentOld(ctx context.Context, in *SegmentCommitRequestOld, opts ...grpc.CallOption) (*SegmentCommitResponseOld, error)
	SegmentInfoOld(ctx context.Context, in *SegmentInfoRequestOld, opts ...grpc.CallOption) (*SegmentInfoResponseOld, error)
	DownloadSegmentOld(ctx context.Context, in *SegmentDownloadRequestOld, opts ...grpc.CallOption) (*SegmentDownloadResponseOld, error)
	DeleteSegmentOld(ctx context.Context, in *SegmentDeleteRequestOld, opts ...grpc.CallOption) (*SegmentDeleteResponseOld, error)
	ListSegmentsOld(ctx context.Context, in *ListSegmentsRequestOld, opts ...grpc.CallOption) (*ListSegmentsResponseOld, error)
	SetAttributionOld(ctx context.Context, in *SetAttributionRequestOld, opts ...grpc.CallOption) (*SetAttributionResponseOld, error)
	ProjectInfo(ctx context.Context, in *ProjectInfoRequest, opts ...grpc.CallOption) (*ProjectInfoResponse, error)
}

type metainfoClient struct {
	cc *grpc.ClientConn
}

func NewMetainfoClient(cc *grpc.ClientConn) MetainfoClient {
	return &metainfoClient{cc}
}

func (c *metainfoClient) CreateBucket(ctx context.Context, in *BucketCreateRequest, opts ...grpc.CallOption) (*BucketCreateResponse, error) {
	out := new(BucketCreateResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) GetBucket(ctx context.Context, in *BucketGetRequest, opts ...grpc.CallOption) (*BucketGetResponse, error) {
	out := new(BucketGetResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/GetBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) DeleteBucket(ctx context.Context, in *BucketDeleteRequest, opts ...grpc.CallOption) (*BucketDeleteResponse, error) {
	out := new(BucketDeleteResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) ListBuckets(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error) {
	out := new(BucketListResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/ListBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) SetBucketAttribution(ctx context.Context, in *BucketSetAttributionRequest, opts ...grpc.CallOption) (*BucketSetAttributionResponse, error) {
	out := new(BucketSetAttributionResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/SetBucketAttribution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) BeginObject(ctx context.Context, in *ObjectBeginRequest, opts ...grpc.CallOption) (*ObjectBeginResponse, error) {
	out := new(ObjectBeginResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/BeginObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) CommitObject(ctx context.Context, in *ObjectCommitRequest, opts ...grpc.CallOption) (*ObjectCommitResponse, error) {
	out := new(ObjectCommitResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/CommitObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) GetObject(ctx context.Context, in *ObjectGetRequest, opts ...grpc.CallOption) (*ObjectGetResponse, error) {
	out := new(ObjectGetResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) ListObjects(ctx context.Context, in *ObjectListRequest, opts ...grpc.CallOption) (*ObjectListResponse, error) {
	out := new(ObjectListResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/ListObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) BeginDeleteObject(ctx context.Context, in *ObjectBeginDeleteRequest, opts ...grpc.CallOption) (*ObjectBeginDeleteResponse, error) {
	out := new(ObjectBeginDeleteResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/BeginDeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) FinishDeleteObject(ctx context.Context, in *ObjectFinishDeleteRequest, opts ...grpc.CallOption) (*ObjectFinishDeleteResponse, error) {
	out := new(ObjectFinishDeleteResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/FinishDeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) BeginSegment(ctx context.Context, in *SegmentBeginRequest, opts ...grpc.CallOption) (*SegmentBeginResponse, error) {
	out := new(SegmentBeginResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/BeginSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) CommitSegment(ctx context.Context, in *SegmentCommitRequest, opts ...grpc.CallOption) (*SegmentCommitResponse, error) {
	out := new(SegmentCommitResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/CommitSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) MakeInlineSegment(ctx context.Context, in *SegmentMakeInlineRequest, opts ...grpc.CallOption) (*SegmentMakeInlineResponse, error) {
	out := new(SegmentMakeInlineResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/MakeInlineSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) BeginDeleteSegment(ctx context.Context, in *SegmentBeginDeleteRequest, opts ...grpc.CallOption) (*SegmentBeginDeleteResponse, error) {
	out := new(SegmentBeginDeleteResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/BeginDeleteSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) FinishDeleteSegment(ctx context.Context, in *SegmentFinishDeleteRequest, opts ...grpc.CallOption) (*SegmentFinishDeleteResponse, error) {
	out := new(SegmentFinishDeleteResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/FinishDeleteSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) ListSegments(ctx context.Context, in *SegmentListRequest, opts ...grpc.CallOption) (*SegmentListResponse, error) {
	out := new(SegmentListResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/ListSegments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) DownloadSegment(ctx context.Context, in *SegmentDownloadRequest, opts ...grpc.CallOption) (*SegmentDownloadResponse, error) {
	out := new(SegmentDownloadResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/DownloadSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error) {
	out := new(BatchResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/Batch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) CreateSegmentOld(ctx context.Context, in *SegmentWriteRequestOld, opts ...grpc.CallOption) (*SegmentWriteResponseOld, error) {
	out := new(SegmentWriteResponseOld)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/CreateSegmentOld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) CommitSegmentOld(ctx context.Context, in *SegmentCommitRequestOld, opts ...grpc.CallOption) (*SegmentCommitResponseOld, error) {
	out := new(SegmentCommitResponseOld)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/CommitSegmentOld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) SegmentInfoOld(ctx context.Context, in *SegmentInfoRequestOld, opts ...grpc.CallOption) (*SegmentInfoResponseOld, error) {
	out := new(SegmentInfoResponseOld)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/SegmentInfoOld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) DownloadSegmentOld(ctx context.Context, in *SegmentDownloadRequestOld, opts ...grpc.CallOption) (*SegmentDownloadResponseOld, error) {
	out := new(SegmentDownloadResponseOld)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/DownloadSegmentOld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) DeleteSegmentOld(ctx context.Context, in *SegmentDeleteRequestOld, opts ...grpc.CallOption) (*SegmentDeleteResponseOld, error) {
	out := new(SegmentDeleteResponseOld)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/DeleteSegmentOld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) ListSegmentsOld(ctx context.Context, in *ListSegmentsRequestOld, opts ...grpc.CallOption) (*ListSegmentsResponseOld, error) {
	out := new(ListSegmentsResponseOld)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/ListSegmentsOld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) SetAttributionOld(ctx context.Context, in *SetAttributionRequestOld, opts ...grpc.CallOption) (*SetAttributionResponseOld, error) {
	out := new(SetAttributionResponseOld)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/SetAttributionOld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metainfoClient) ProjectInfo(ctx context.Context, in *ProjectInfoRequest, opts ...grpc.CallOption) (*ProjectInfoResponse, error) {
	out := new(ProjectInfoResponse)
	err := c.cc.Invoke(ctx, "/metainfo.Metainfo/ProjectInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetainfoServer is the server API for Metainfo service.
type MetainfoServer interface {
	// Bucket
	CreateBucket(context.Context, *BucketCreateRequest) (*BucketCreateResponse, error)
	GetBucket(context.Context, *BucketGetRequest) (*BucketGetResponse, error)
	DeleteBucket(context.Context, *BucketDeleteRequest) (*BucketDeleteResponse, error)
	ListBuckets(context.Context, *BucketListRequest) (*BucketListResponse, error)
	SetBucketAttribution(context.Context, *BucketSetAttributionRequest) (*BucketSetAttributionResponse, error)
	// Object
	BeginObject(context.Context, *ObjectBeginRequest) (*ObjectBeginResponse, error)
	CommitObject(context.Context, *ObjectCommitRequest) (*ObjectCommitResponse, error)
	GetObject(context.Context, *ObjectGetRequest) (*ObjectGetResponse, error)
	ListObjects(context.Context, *ObjectListRequest) (*ObjectListResponse, error)
	BeginDeleteObject(context.Context, *ObjectBeginDeleteRequest) (*ObjectBeginDeleteResponse, error)
	FinishDeleteObject(context.Context, *ObjectFinishDeleteRequest) (*ObjectFinishDeleteResponse, error)
	BeginSegment(context.Context, *SegmentBeginRequest) (*SegmentBeginResponse, error)
	CommitSegment(context.Context, *SegmentCommitRequest) (*SegmentCommitResponse, error)
	MakeInlineSegment(context.Context, *SegmentMakeInlineRequest) (*SegmentMakeInlineResponse, error)
	BeginDeleteSegment(context.Context, *SegmentBeginDeleteRequest) (*SegmentBeginDeleteResponse, error)
	FinishDeleteSegment(context.Context, *SegmentFinishDeleteRequest) (*SegmentFinishDeleteResponse, error)
	ListSegments(context.Context, *SegmentListRequest) (*SegmentListResponse, error)
	DownloadSegment(context.Context, *SegmentDownloadRequest) (*SegmentDownloadResponse, error)
	Batch(context.Context, *BatchRequest) (*BatchResponse, error)
	CreateSegmentOld(context.Context, *SegmentWriteRequestOld) (*SegmentWriteResponseOld, error)
	CommitSegmentOld(context.Context, *SegmentCommitRequestOld) (*SegmentCommitResponseOld, error)
	SegmentInfoOld(context.Context, *SegmentInfoRequestOld) (*SegmentInfoResponseOld, error)
	DownloadSegmentOld(context.Context, *SegmentDownloadRequestOld) (*SegmentDownloadResponseOld, error)
	DeleteSegmentOld(context.Context, *SegmentDeleteRequestOld) (*SegmentDeleteResponseOld, error)
	ListSegmentsOld(context.Context, *ListSegmentsRequestOld) (*ListSegmentsResponseOld, error)
	SetAttributionOld(context.Context, *SetAttributionRequestOld) (*SetAttributionResponseOld, error)
	ProjectInfo(context.Context, *ProjectInfoRequest) (*ProjectInfoResponse, error)
}

func RegisterMetainfoServer(s *grpc.Server, srv MetainfoServer) {
	s.RegisterService(&_Metainfo_serviceDesc, srv)
}

func _Metainfo_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/CreateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).CreateBucket(ctx, req.(*BucketCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_GetBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).GetBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/GetBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).GetBucket(ctx, req.(*BucketGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).DeleteBucket(ctx, req.(*BucketDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).ListBuckets(ctx, req.(*BucketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_SetBucketAttribution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketSetAttributionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).SetBucketAttribution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/SetBucketAttribution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).SetBucketAttribution(ctx, req.(*BucketSetAttributionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_BeginObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectBeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).BeginObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/BeginObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).BeginObject(ctx, req.(*ObjectBeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_CommitObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).CommitObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/CommitObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).CommitObject(ctx, req.(*ObjectCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).GetObject(ctx, req.(*ObjectGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/ListObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).ListObjects(ctx, req.(*ObjectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_BeginDeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectBeginDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).BeginDeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/BeginDeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).BeginDeleteObject(ctx, req.(*ObjectBeginDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_FinishDeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectFinishDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).FinishDeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/FinishDeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).FinishDeleteObject(ctx, req.(*ObjectFinishDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_BeginSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentBeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).BeginSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/BeginSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).BeginSegment(ctx, req.(*SegmentBeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_CommitSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).CommitSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/CommitSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).CommitSegment(ctx, req.(*SegmentCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_MakeInlineSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentMakeInlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).MakeInlineSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/MakeInlineSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).MakeInlineSegment(ctx, req.(*SegmentMakeInlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_BeginDeleteSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentBeginDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).BeginDeleteSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/BeginDeleteSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).BeginDeleteSegment(ctx, req.(*SegmentBeginDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_FinishDeleteSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentFinishDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).FinishDeleteSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/FinishDeleteSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).FinishDeleteSegment(ctx, req.(*SegmentFinishDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_ListSegments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).ListSegments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/ListSegments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).ListSegments(ctx, req.(*SegmentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_DownloadSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).DownloadSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/DownloadSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).DownloadSegment(ctx, req.(*SegmentDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_Batch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).Batch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/Batch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).Batch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_CreateSegmentOld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentWriteRequestOld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).CreateSegmentOld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/CreateSegmentOld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).CreateSegmentOld(ctx, req.(*SegmentWriteRequestOld))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_CommitSegmentOld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentCommitRequestOld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).CommitSegmentOld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/CommitSegmentOld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).CommitSegmentOld(ctx, req.(*SegmentCommitRequestOld))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_SegmentInfoOld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentInfoRequestOld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).SegmentInfoOld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/SegmentInfoOld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).SegmentInfoOld(ctx, req.(*SegmentInfoRequestOld))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_DownloadSegmentOld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentDownloadRequestOld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).DownloadSegmentOld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/DownloadSegmentOld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).DownloadSegmentOld(ctx, req.(*SegmentDownloadRequestOld))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_DeleteSegmentOld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SegmentDeleteRequestOld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).DeleteSegmentOld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/DeleteSegmentOld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).DeleteSegmentOld(ctx, req.(*SegmentDeleteRequestOld))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_ListSegmentsOld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSegmentsRequestOld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).ListSegmentsOld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/ListSegmentsOld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).ListSegmentsOld(ctx, req.(*ListSegmentsRequestOld))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_SetAttributionOld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAttributionRequestOld)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).SetAttributionOld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/SetAttributionOld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).SetAttributionOld(ctx, req.(*SetAttributionRequestOld))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metainfo_ProjectInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetainfoServer).ProjectInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metainfo.Metainfo/ProjectInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetainfoServer).ProjectInfo(ctx, req.(*ProjectInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metainfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metainfo.Metainfo",
	HandlerType: (*MetainfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBucket",
			Handler:    _Metainfo_CreateBucket_Handler,
		},
		{
			MethodName: "GetBucket",
			Handler:    _Metainfo_GetBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _Metainfo_DeleteBucket_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _Metainfo_ListBuckets_Handler,
		},
		{
			MethodName: "SetBucketAttribution",
			Handler:    _Metainfo_SetBucketAttribution_Handler,
		},
		{
			MethodName: "BeginObject",
			Handler:    _Metainfo_BeginObject_Handler,
		},
		{
			MethodName: "CommitObject",
			Handler:    _Metainfo_CommitObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _Metainfo_GetObject_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _Metainfo_ListObjects_Handler,
		},
		{
			MethodName: "BeginDeleteObject",
			Handler:    _Metainfo_BeginDeleteObject_Handler,
		},
		{
			MethodName: "FinishDeleteObject",
			Handler:    _Metainfo_FinishDeleteObject_Handler,
		},
		{
			MethodName: "BeginSegment",
			Handler:    _Metainfo_BeginSegment_Handler,
		},
		{
			MethodName: "CommitSegment",
			Handler:    _Metainfo_CommitSegment_Handler,
		},
		{
			MethodName: "MakeInlineSegment",
			Handler:    _Metainfo_MakeInlineSegment_Handler,
		},
		{
			MethodName: "BeginDeleteSegment",
			Handler:    _Metainfo_BeginDeleteSegment_Handler,
		},
		{
			MethodName: "FinishDeleteSegment",
			Handler:    _Metainfo_FinishDeleteSegment_Handler,
		},
		{
			MethodName: "ListSegments",
			Handler:    _Metainfo_ListSegments_Handler,
		},
		{
			MethodName: "DownloadSegment",
			Handler:    _Metainfo_DownloadSegment_Handler,
		},
		{
			MethodName: "Batch",
			Handler:    _Metainfo_Batch_Handler,
		},
		{
			MethodName: "CreateSegmentOld",
			Handler:    _Metainfo_CreateSegmentOld_Handler,
		},
		{
			MethodName: "CommitSegmentOld",
			Handler:    _Metainfo_CommitSegmentOld_Handler,
		},
		{
			MethodName: "SegmentInfoOld",
			Handler:    _Metainfo_SegmentInfoOld_Handler,
		},
		{
			MethodName: "DownloadSegmentOld",
			Handler:    _Metainfo_DownloadSegmentOld_Handler,
		},
		{
			MethodName: "DeleteSegmentOld",
			Handler:    _Metainfo_DeleteSegmentOld_Handler,
		},
		{
			MethodName: "ListSegmentsOld",
			Handler:    _Metainfo_ListSegmentsOld_Handler,
		},
		{
			MethodName: "SetAttributionOld",
			Handler:    _Metainfo_SetAttributionOld_Handler,
		},
		{
			MethodName: "ProjectInfo",
			Handler:    _Metainfo_ProjectInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metainfo.proto",
}
