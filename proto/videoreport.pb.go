// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: videoreport.proto

package videoreport

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DailyReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime string `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"` // 按日统计，开始时间
	EndTime   string `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`       // 结束时间
}

func (x *DailyReportReq) Reset() {
	*x = DailyReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoreport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyReportReq) ProtoMessage() {}

func (x *DailyReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_videoreport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyReportReq.ProtoReflect.Descriptor instead.
func (*DailyReportReq) Descriptor() ([]byte, []int) {
	return file_videoreport_proto_rawDescGZIP(), []int{0}
}

func (x *DailyReportReq) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *DailyReportReq) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type ReportInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportData         string `protobuf:"bytes,1,opt,name=report_data,json=reportData,proto3" json:"report_data,omitempty"`                            //统计日期
	Dau                int32  `protobuf:"varint,2,opt,name=dau,proto3" json:"dau,omitempty"`                                                           // 日活
	TotalDuration      int64  `protobuf:"varint,3,opt,name=total_duration,json=totalDuration,proto3" json:"total_duration,omitempty"`                  // 总时长
	TotalValidDuration int64  `protobuf:"varint,4,opt,name=total_valid_duration,json=totalValidDuration,proto3" json:"total_valid_duration,omitempty"` // 总有效时长
	AvgDuration        string `protobuf:"bytes,5,opt,name=avg_duration,json=avgDuration,proto3" json:"avg_duration,omitempty"`                         // 人均时长
	AvgValidDuration   string `protobuf:"bytes,6,opt,name=avg_valid_duration,json=avgValidDuration,proto3" json:"avg_valid_duration,omitempty"`        // 人均有效时长 (超过8')
	AvgVv              string `protobuf:"bytes,7,opt,name=avg_vv,json=avgVv,proto3" json:"avg_vv,omitempty"`                                           // 人均vv
}

func (x *ReportInfo) Reset() {
	*x = ReportInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoreport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportInfo) ProtoMessage() {}

func (x *ReportInfo) ProtoReflect() protoreflect.Message {
	mi := &file_videoreport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportInfo.ProtoReflect.Descriptor instead.
func (*ReportInfo) Descriptor() ([]byte, []int) {
	return file_videoreport_proto_rawDescGZIP(), []int{1}
}

func (x *ReportInfo) GetReportData() string {
	if x != nil {
		return x.ReportData
	}
	return ""
}

func (x *ReportInfo) GetDau() int32 {
	if x != nil {
		return x.Dau
	}
	return 0
}

func (x *ReportInfo) GetTotalDuration() int64 {
	if x != nil {
		return x.TotalDuration
	}
	return 0
}

func (x *ReportInfo) GetTotalValidDuration() int64 {
	if x != nil {
		return x.TotalValidDuration
	}
	return 0
}

func (x *ReportInfo) GetAvgDuration() string {
	if x != nil {
		return x.AvgDuration
	}
	return ""
}

func (x *ReportInfo) GetAvgValidDuration() string {
	if x != nil {
		return x.AvgValidDuration
	}
	return ""
}

func (x *ReportInfo) GetAvgVv() string {
	if x != nil {
		return x.AvgVv
	}
	return ""
}

type DailyReportRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportInfo []*ReportInfo `protobuf:"bytes,1,rep,name=report_info,json=reportInfo,proto3" json:"report_info,omitempty"`
}

func (x *DailyReportRsp) Reset() {
	*x = DailyReportRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoreport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyReportRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyReportRsp) ProtoMessage() {}

func (x *DailyReportRsp) ProtoReflect() protoreflect.Message {
	mi := &file_videoreport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyReportRsp.ProtoReflect.Descriptor instead.
func (*DailyReportRsp) Descriptor() ([]byte, []int) {
	return file_videoreport_proto_rawDescGZIP(), []int{2}
}

func (x *DailyReportRsp) GetReportInfo() []*ReportInfo {
	if x != nil {
		return x.ReportInfo
	}
	return nil
}

var File_videoreport_proto protoreflect.FileDescriptor

var file_videoreport_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x4a, 0x0a, 0x0e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x80, 0x02, 0x0a,
	0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x61, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x75, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x67, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x76, 0x67, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x76,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x67, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x76, 0x67, 0x5f,
	0x76, 0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x76, 0x67, 0x56, 0x76, 0x22,
	0x4a, 0x0a, 0x0e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x38, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x5b, 0x0a, 0x0b, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4c, 0x0a, 0x10, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x73, 0x70, 0x42, 0x12, 0x5a, 0x0d, 0x2e, 0x3b, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x80, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_videoreport_proto_rawDescOnce sync.Once
	file_videoreport_proto_rawDescData = file_videoreport_proto_rawDesc
)

func file_videoreport_proto_rawDescGZIP() []byte {
	file_videoreport_proto_rawDescOnce.Do(func() {
		file_videoreport_proto_rawDescData = protoimpl.X.CompressGZIP(file_videoreport_proto_rawDescData)
	})
	return file_videoreport_proto_rawDescData
}

var file_videoreport_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_videoreport_proto_goTypes = []interface{}{
	(*DailyReportReq)(nil), // 0: videoreport.DailyReportReq
	(*ReportInfo)(nil),     // 1: videoreport.ReportInfo
	(*DailyReportRsp)(nil), // 2: videoreport.DailyReportRsp
}
var file_videoreport_proto_depIdxs = []int32{
	1, // 0: videoreport.DailyReportRsp.report_info:type_name -> videoreport.ReportInfo
	0, // 1: videoreport.VideoReport.queryDailyReport:input_type -> videoreport.DailyReportReq
	2, // 2: videoreport.VideoReport.queryDailyReport:output_type -> videoreport.DailyReportRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_videoreport_proto_init() }
func file_videoreport_proto_init() {
	if File_videoreport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_videoreport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyReportReq); i {
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
		file_videoreport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportInfo); i {
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
		file_videoreport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyReportRsp); i {
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
			RawDescriptor: file_videoreport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_videoreport_proto_goTypes,
		DependencyIndexes: file_videoreport_proto_depIdxs,
		MessageInfos:      file_videoreport_proto_msgTypes,
	}.Build()
	File_videoreport_proto = out.File
	file_videoreport_proto_rawDesc = nil
	file_videoreport_proto_goTypes = nil
	file_videoreport_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VideoReportClient is the client API for VideoReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoReportClient interface {
	// 查询每日报表统计数据
	QueryDailyReport(ctx context.Context, in *DailyReportReq, opts ...grpc.CallOption) (*DailyReportRsp, error)
}

type videoReportClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoReportClient(cc grpc.ClientConnInterface) VideoReportClient {
	return &videoReportClient{cc}
}

func (c *videoReportClient) QueryDailyReport(ctx context.Context, in *DailyReportReq, opts ...grpc.CallOption) (*DailyReportRsp, error) {
	out := new(DailyReportRsp)
	err := c.cc.Invoke(ctx, "/videoreport.VideoReport/queryDailyReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoReportServer is the server API for VideoReport service.
type VideoReportServer interface {
	// 查询每日报表统计数据
	QueryDailyReport(context.Context, *DailyReportReq) (*DailyReportRsp, error)
}

// UnimplementedVideoReportServer can be embedded to have forward compatible implementations.
type UnimplementedVideoReportServer struct {
}

func (*UnimplementedVideoReportServer) QueryDailyReport(context.Context, *DailyReportReq) (*DailyReportRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDailyReport not implemented")
}

func RegisterVideoReportServer(s *grpc.Server, srv VideoReportServer) {
	s.RegisterService(&_VideoReport_serviceDesc, srv)
}

func _VideoReport_QueryDailyReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DailyReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoReportServer).QueryDailyReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/videoreport.VideoReport/QueryDailyReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoReportServer).QueryDailyReport(ctx, req.(*DailyReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoReport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "videoreport.VideoReport",
	HandlerType: (*VideoReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "queryDailyReport",
			Handler:    _VideoReport_QueryDailyReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "videoreport.proto",
}
