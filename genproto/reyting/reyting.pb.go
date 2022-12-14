// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reyting/reyting.proto

package reyting

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Ranking struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Ranking              int64    `protobuf:"varint,2,opt,name=ranking,proto3" json:"ranking"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	PostId               int64    `protobuf:"varint,4,opt,name=post_id,json=postId,proto3" json:"post_id"`
	CustomerId           int64    `protobuf:"varint,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ranking) Reset()         { *m = Ranking{} }
func (m *Ranking) String() string { return proto.CompactTextString(m) }
func (*Ranking) ProtoMessage()    {}
func (*Ranking) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d96215f86e97cc8, []int{0}
}
func (m *Ranking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ranking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ranking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ranking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ranking.Merge(m, src)
}
func (m *Ranking) XXX_Size() int {
	return m.Size()
}
func (m *Ranking) XXX_DiscardUnknown() {
	xxx_messageInfo_Ranking.DiscardUnknown(m)
}

var xxx_messageInfo_Ranking proto.InternalMessageInfo

func (m *Ranking) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ranking) GetRanking() int64 {
	if m != nil {
		return m.Ranking
	}
	return 0
}

func (m *Ranking) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Ranking) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *Ranking) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

type Rankings struct {
	Rankings             []*Ranking `protobuf:"bytes,1,rep,name=rankings,proto3" json:"rankings"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Rankings) Reset()         { *m = Rankings{} }
func (m *Rankings) String() string { return proto.CompactTextString(m) }
func (*Rankings) ProtoMessage()    {}
func (*Rankings) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d96215f86e97cc8, []int{1}
}
func (m *Rankings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rankings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rankings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rankings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rankings.Merge(m, src)
}
func (m *Rankings) XXX_Size() int {
	return m.Size()
}
func (m *Rankings) XXX_DiscardUnknown() {
	xxx_messageInfo_Rankings.DiscardUnknown(m)
}

var xxx_messageInfo_Rankings proto.InternalMessageInfo

func (m *Rankings) GetRankings() []*Ranking {
	if m != nil {
		return m.Rankings
	}
	return nil
}

type Id struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d96215f86e97cc8, []int{2}
}
func (m *Id) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Id.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return m.Size()
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d96215f86e97cc8, []int{3}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Ranking)(nil), "reyting.Ranking")
	proto.RegisterType((*Rankings)(nil), "reyting.Rankings")
	proto.RegisterType((*Id)(nil), "reyting.id")
	proto.RegisterType((*Empty)(nil), "reyting.Empty")
}

func init() { proto.RegisterFile("reyting/reyting.proto", fileDescriptor_7d96215f86e97cc8) }

var fileDescriptor_7d96215f86e97cc8 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0x4d, 0xdb, 0x94, 0x17, 0x11, 0x95, 0x27, 0x50, 0xad, 0x0e, 0x21, 0xca, 0x94, 0x01,
	0x8a, 0x28, 0x0b, 0xac, 0xad, 0x10, 0xea, 0x1a, 0x0e, 0x80, 0x42, 0x6c, 0x55, 0x16, 0xca, 0x8f,
	0x6c, 0x83, 0x94, 0x73, 0xb0, 0x70, 0x09, 0xee, 0xc1, 0xc8, 0x11, 0x50, 0xb8, 0x08, 0x8a, 0x71,
	0xa2, 0x0a, 0x06, 0x16, 0xdb, 0xef, 0xfb, 0x79, 0x7e, 0x9f, 0x0d, 0xc7, 0x92, 0xd7, 0x5a, 0x14,
	0xdb, 0x73, 0xbb, 0x2f, 0x2a, 0x59, 0xea, 0x12, 0x5d, 0x5b, 0x46, 0x2f, 0x04, 0xdc, 0x24, 0x2d,
	0x1e, 0x45, 0xb1, 0x45, 0x84, 0x61, 0x91, 0xe6, 0x9c, 0x92, 0x90, 0xc4, 0xfb, 0x89, 0x39, 0x23,
	0x05, 0x57, 0xfe, 0xd0, 0x74, 0x10, 0x92, 0xd8, 0x49, 0xba, 0x12, 0x43, 0xf0, 0x18, 0x57, 0x99,
	0x14, 0x95, 0x16, 0x65, 0x41, 0x1d, 0x63, 0xda, 0x85, 0x70, 0x06, 0x6e, 0x55, 0x2a, 0x7d, 0x2f,
	0x18, 0x1d, 0x1a, 0xef, 0xb8, 0x2d, 0x37, 0x0c, 0x4f, 0xc0, 0xcb, 0x9e, 0x94, 0x2e, 0x73, 0x2e,
	0x5b, 0x72, 0x64, 0x48, 0xe8, 0xa0, 0x0d, 0x8b, 0xae, 0x60, 0x62, 0x87, 0x52, 0x78, 0x0a, 0x13,
	0x7b, 0xa5, 0xa2, 0x24, 0x74, 0x62, 0x6f, 0x39, 0x5d, 0x74, 0x61, 0xac, 0x28, 0xe9, 0x15, 0xd1,
	0x11, 0x0c, 0x04, 0x43, 0xbf, 0x5d, 0x4d, 0x0e, 0x27, 0x19, 0x08, 0x16, 0xb9, 0x30, 0xba, 0xc9,
	0x2b, 0x5d, 0x2f, 0xdf, 0x08, 0xf8, 0xd6, 0x74, 0xc7, 0xe5, 0xb3, 0xc8, 0x38, 0x5e, 0xc0, 0xc1,
	0x5a, 0xf2, 0x54, 0xf3, 0xee, 0x19, 0xfe, 0xb4, 0x9f, 0xfb, 0x3d, 0x62, 0xba, 0xe0, 0x19, 0x78,
	0xb7, 0x5c, 0xf7, 0x13, 0x7a, 0x3d, 0x2d, 0xd8, 0xfc, 0xf0, 0xb7, 0x5b, 0xe1, 0x35, 0xcc, 0x76,
	0xe4, 0xab, 0x7a, 0xdd, 0x07, 0xfd, 0xcf, 0xba, 0x9a, 0xbe, 0x37, 0x01, 0xf9, 0x68, 0x02, 0xf2,
	0xd9, 0x04, 0xe4, 0xf5, 0x2b, 0xd8, 0x7b, 0x18, 0x9b, 0x0f, 0xbc, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0x36, 0x4f, 0xf2, 0xd9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RankingServiceClient is the client API for RankingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RankingServiceClient interface {
	CreateRanking(ctx context.Context, in *Ranking, opts ...grpc.CallOption) (*Empty, error)
	GetRankings(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Rankings, error)
	GetRankingsByCustomerId(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Rankings, error)
}

type rankingServiceClient struct {
	cc *grpc.ClientConn
}

func NewRankingServiceClient(cc *grpc.ClientConn) RankingServiceClient {
	return &rankingServiceClient{cc}
}

func (c *rankingServiceClient) CreateRanking(ctx context.Context, in *Ranking, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/reyting.RankingService/CreateRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankingServiceClient) GetRankings(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Rankings, error) {
	out := new(Rankings)
	err := c.cc.Invoke(ctx, "/reyting.RankingService/GetRankings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankingServiceClient) GetRankingsByCustomerId(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Rankings, error) {
	out := new(Rankings)
	err := c.cc.Invoke(ctx, "/reyting.RankingService/GetRankingsByCustomerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankingServiceServer is the server API for RankingService service.
type RankingServiceServer interface {
	CreateRanking(context.Context, *Ranking) (*Empty, error)
	GetRankings(context.Context, *Id) (*Rankings, error)
	GetRankingsByCustomerId(context.Context, *Id) (*Rankings, error)
}

// UnimplementedRankingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRankingServiceServer struct {
}

func (*UnimplementedRankingServiceServer) CreateRanking(ctx context.Context, req *Ranking) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRanking not implemented")
}
func (*UnimplementedRankingServiceServer) GetRankings(ctx context.Context, req *Id) (*Rankings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRankings not implemented")
}
func (*UnimplementedRankingServiceServer) GetRankingsByCustomerId(ctx context.Context, req *Id) (*Rankings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRankingsByCustomerId not implemented")
}

func RegisterRankingServiceServer(s *grpc.Server, srv RankingServiceServer) {
	s.RegisterService(&_RankingService_serviceDesc, srv)
}

func _RankingService_CreateRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ranking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankingServiceServer).CreateRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reyting.RankingService/CreateRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankingServiceServer).CreateRanking(ctx, req.(*Ranking))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankingService_GetRankings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankingServiceServer).GetRankings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reyting.RankingService/GetRankings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankingServiceServer).GetRankings(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankingService_GetRankingsByCustomerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankingServiceServer).GetRankingsByCustomerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reyting.RankingService/GetRankingsByCustomerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankingServiceServer).GetRankingsByCustomerId(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _RankingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reyting.RankingService",
	HandlerType: (*RankingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRanking",
			Handler:    _RankingService_CreateRanking_Handler,
		},
		{
			MethodName: "GetRankings",
			Handler:    _RankingService_GetRankings_Handler,
		},
		{
			MethodName: "GetRankingsByCustomerId",
			Handler:    _RankingService_GetRankingsByCustomerId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reyting/reyting.proto",
}

func (m *Ranking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ranking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ranking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CustomerId != 0 {
		i = encodeVarintReyting(dAtA, i, uint64(m.CustomerId))
		i--
		dAtA[i] = 0x28
	}
	if m.PostId != 0 {
		i = encodeVarintReyting(dAtA, i, uint64(m.PostId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintReyting(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Ranking != 0 {
		i = encodeVarintReyting(dAtA, i, uint64(m.Ranking))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintReyting(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Rankings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rankings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rankings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Rankings) > 0 {
		for iNdEx := len(m.Rankings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rankings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReyting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Id) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Id) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Id) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintReyting(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Empty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintReyting(dAtA []byte, offset int, v uint64) int {
	offset -= sovReyting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ranking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovReyting(uint64(l))
	}
	if m.Ranking != 0 {
		n += 1 + sovReyting(uint64(m.Ranking))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovReyting(uint64(l))
	}
	if m.PostId != 0 {
		n += 1 + sovReyting(uint64(m.PostId))
	}
	if m.CustomerId != 0 {
		n += 1 + sovReyting(uint64(m.CustomerId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Rankings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rankings) > 0 {
		for _, e := range m.Rankings {
			l = e.Size()
			n += 1 + l + sovReyting(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Id) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovReyting(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovReyting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReyting(x uint64) (n int) {
	return sovReyting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ranking) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReyting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ranking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ranking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReyting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReyting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ranking", wireType)
			}
			m.Ranking = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ranking |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReyting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReyting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
			}
			m.PostId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomerId", wireType)
			}
			m.CustomerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CustomerId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReyting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReyting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Rankings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReyting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rankings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rankings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rankings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthReyting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReyting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rankings = append(m.Rankings, &Ranking{})
			if err := m.Rankings[len(m.Rankings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReyting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReyting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Id) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReyting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: id: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: id: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReyting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReyting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReyting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipReyting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReyting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipReyting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReyting
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReyting
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthReyting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReyting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReyting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReyting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReyting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReyting = fmt.Errorf("proto: unexpected end of group")
)
