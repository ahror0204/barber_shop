// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: salon.proto

package protos

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Salon struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email"`
	Rating               int32    `protobuf:"varint,5,opt,name=rating,proto3" json:"rating"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address"`
	Latitude             string   `protobuf:"bytes,7,opt,name=latitude,proto3" json:"latitude"`
	Longitude            string   `protobuf:"bytes,8,opt,name=longitude,proto3" json:"longitude"`
	StartTime            string   `protobuf:"bytes,9,opt,name=start_time,json=startTime,proto3" json:"start_time"`
	EndTime              string   `protobuf:"bytes,10,opt,name=end_time,json=endTime,proto3" json:"end_time"`
	ImageUrl             string   `protobuf:"bytes,11,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	DeletedAt            string   `protobuf:"bytes,14,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Salon) Reset()         { *m = Salon{} }
func (m *Salon) String() string { return proto.CompactTextString(m) }
func (*Salon) ProtoMessage()    {}
func (*Salon) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddf36ac9db7a45b6, []int{0}
}
func (m *Salon) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Salon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Salon.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Salon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Salon.Merge(m, src)
}
func (m *Salon) XXX_Size() int {
	return m.Size()
}
func (m *Salon) XXX_DiscardUnknown() {
	xxx_messageInfo_Salon.DiscardUnknown(m)
}

var xxx_messageInfo_Salon proto.InternalMessageInfo

func (m *Salon) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Salon) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Salon) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Salon) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Salon) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Salon) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Salon) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *Salon) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *Salon) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Salon) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Salon) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *Salon) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Salon) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Salon) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type GetSalonsParams struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	Search               string   `protobuf:"bytes,3,opt,name=search,proto3" json:"search"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSalonsParams) Reset()         { *m = GetSalonsParams{} }
func (m *GetSalonsParams) String() string { return proto.CompactTextString(m) }
func (*GetSalonsParams) ProtoMessage()    {}
func (*GetSalonsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddf36ac9db7a45b6, []int{1}
}
func (m *GetSalonsParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSalonsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSalonsParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSalonsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSalonsParams.Merge(m, src)
}
func (m *GetSalonsParams) XXX_Size() int {
	return m.Size()
}
func (m *GetSalonsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSalonsParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetSalonsParams proto.InternalMessageInfo

func (m *GetSalonsParams) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetSalonsParams) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetSalonsParams) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type AllSalons struct {
	Salons               []*Salon `protobuf:"bytes,1,rep,name=salons,proto3" json:"salons"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllSalons) Reset()         { *m = AllSalons{} }
func (m *AllSalons) String() string { return proto.CompactTextString(m) }
func (*AllSalons) ProtoMessage()    {}
func (*AllSalons) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddf36ac9db7a45b6, []int{2}
}
func (m *AllSalons) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllSalons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllSalons.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllSalons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllSalons.Merge(m, src)
}
func (m *AllSalons) XXX_Size() int {
	return m.Size()
}
func (m *AllSalons) XXX_DiscardUnknown() {
	xxx_messageInfo_AllSalons.DiscardUnknown(m)
}

var xxx_messageInfo_AllSalons proto.InternalMessageInfo

func (m *AllSalons) GetSalons() []*Salon {
	if m != nil {
		return m.Salons
	}
	return nil
}

func (m *AllSalons) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Salon)(nil), "protos.Salon")
	proto.RegisterType((*GetSalonsParams)(nil), "protos.GetSalonsParams")
	proto.RegisterType((*AllSalons)(nil), "protos.AllSalons")
}

func init() { proto.RegisterFile("salon.proto", fileDescriptor_ddf36ac9db7a45b6) }

var fileDescriptor_ddf36ac9db7a45b6 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x49, 0xb3, 0xed, 0x36, 0x93, 0xee, 0x82, 0x2c, 0x84, 0x4c, 0x81, 0xaa, 0x54, 0x42,
	0xea, 0x69, 0x0f, 0xcb, 0x85, 0x6b, 0x01, 0x09, 0xb8, 0x20, 0x94, 0xc2, 0xb9, 0xf2, 0xd6, 0xa3,
	0xae, 0x25, 0xdb, 0x89, 0x6c, 0x87, 0x27, 0xe0, 0x21, 0x78, 0x24, 0x8e, 0x3c, 0x02, 0x2a, 0x2f,
	0x82, 0x32, 0xe3, 0x2e, 0xda, 0x53, 0xfc, 0x7f, 0x9f, 0x35, 0xa3, 0x78, 0x06, 0xea, 0xa8, 0x6c,
	0xeb, 0xaf, 0xba, 0xd0, 0xa6, 0x56, 0x4c, 0xe8, 0x13, 0xe7, 0xd0, 0x47, 0x0c, 0xcc, 0x56, 0x3f,
	0x4a, 0x18, 0x6f, 0x87, 0x3b, 0xe2, 0x12, 0x46, 0x46, 0xcb, 0x62, 0x59, 0xac, 0xab, 0x66, 0x64,
	0xb4, 0x10, 0x70, 0xe6, 0x95, 0x43, 0x39, 0x22, 0x42, 0x67, 0xf1, 0x12, 0x66, 0xdd, 0x6d, 0xeb,
	0x71, 0xe7, 0x7b, 0x77, 0x83, 0x41, 0x96, 0xe4, 0x6a, 0x62, 0x9f, 0x09, 0x89, 0xc7, 0x30, 0x46,
	0xa7, 0x8c, 0x95, 0x67, 0xe4, 0x38, 0x88, 0x27, 0x30, 0x09, 0x2a, 0x19, 0x7f, 0x90, 0xe3, 0x65,
	0xb1, 0x1e, 0x37, 0x39, 0x09, 0x09, 0xe7, 0x4a, 0xeb, 0x80, 0x31, 0xca, 0x09, 0xdd, 0x3f, 0x45,
	0x31, 0x87, 0xa9, 0x55, 0xc9, 0xa4, 0x5e, 0xa3, 0x3c, 0x27, 0x75, 0x97, 0xc5, 0x73, 0xa8, 0x6c,
	0xeb, 0x0f, 0x2c, 0xa7, 0x24, 0xff, 0x03, 0xf1, 0x02, 0x20, 0x26, 0x15, 0xd2, 0x2e, 0x19, 0x87,
	0xb2, 0x62, 0x4d, 0xe4, 0xab, 0x71, 0x28, 0x9e, 0xc2, 0x14, 0xbd, 0x66, 0x09, 0xdc, 0x13, 0xbd,
	0x26, 0xf5, 0x0c, 0x2a, 0xe3, 0xd4, 0x01, 0x77, 0x7d, 0xb0, 0xb2, 0xe6, 0xa6, 0x04, 0xbe, 0x05,
	0x3b, 0x94, 0xdd, 0x07, 0x54, 0x09, 0xf5, 0x4e, 0x25, 0x39, 0xe3, 0xb2, 0x99, 0x6c, 0xd2, 0xa0,
	0xfb, 0x4e, 0x9f, 0xf4, 0x05, 0xeb, 0x4c, 0x58, 0x6b, 0xb4, 0x98, 0xf5, 0x25, 0xeb, 0x4c, 0x36,
	0x69, 0xb5, 0x85, 0x87, 0x1f, 0x30, 0xd1, 0x20, 0xe2, 0x17, 0x15, 0x94, 0x8b, 0xc3, 0xfb, 0x77,
	0xea, 0x80, 0x34, 0x91, 0xb2, 0xa1, 0xf3, 0xf0, 0xb8, 0xd6, 0x38, 0x93, 0x68, 0x28, 0x65, 0xc3,
	0x61, 0x78, 0xdc, 0x88, 0x2a, 0xec, 0x6f, 0xf3, 0x3c, 0x72, 0x5a, 0x7d, 0x84, 0x6a, 0x63, 0x2d,
	0x17, 0x15, 0xaf, 0x60, 0x42, 0xbb, 0x10, 0x65, 0xb1, 0x2c, 0xd7, 0xf5, 0xf5, 0x05, 0x2f, 0x40,
	0xbc, 0x22, 0xdf, 0x64, 0x39, 0x74, 0xd8, 0xb7, 0xbd, 0xbf, 0xeb, 0x40, 0xe1, 0xfa, 0x0d, 0xcc,
	0xe8, 0xda, 0x16, 0xc3, 0x77, 0xb3, 0x47, 0xb1, 0x86, 0xfa, 0x1d, 0xfd, 0x39, 0xaf, 0xce, 0xfd,
	0x5a, 0x73, 0x38, 0xc5, 0x4f, 0xef, 0xdf, 0x3e, 0xfa, 0x75, 0x5c, 0x14, 0xbf, 0x8f, 0x8b, 0xe2,
	0xcf, 0x71, 0x51, 0xfc, 0xfc, 0xbb, 0x78, 0x70, 0xc3, 0x5b, 0xf8, 0xfa, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc8, 0x1d, 0x4a, 0x5a, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SalonServiceClient is the client API for SalonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SalonServiceClient interface {
	CreateSalon(ctx context.Context, in *Salon, opts ...grpc.CallOption) (*ID, error)
}

type salonServiceClient struct {
	cc *grpc.ClientConn
}

func NewSalonServiceClient(cc *grpc.ClientConn) SalonServiceClient {
	return &salonServiceClient{cc}
}

func (c *salonServiceClient) CreateSalon(ctx context.Context, in *Salon, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/protos.SalonService/CreateSalon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SalonServiceServer is the server API for SalonService service.
type SalonServiceServer interface {
	CreateSalon(context.Context, *Salon) (*ID, error)
}

// UnimplementedSalonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSalonServiceServer struct {
}

func (*UnimplementedSalonServiceServer) CreateSalon(ctx context.Context, req *Salon) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSalon not implemented")
}

func RegisterSalonServiceServer(s *grpc.Server, srv SalonServiceServer) {
	s.RegisterService(&_SalonService_serviceDesc, srv)
}

func _SalonService_CreateSalon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Salon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalonServiceServer).CreateSalon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.SalonService/CreateSalon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalonServiceServer).CreateSalon(ctx, req.(*Salon))
	}
	return interceptor(ctx, in, info, handler)
}

var _SalonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.SalonService",
	HandlerType: (*SalonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSalon",
			Handler:    _SalonService_CreateSalon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "salon.proto",
}

func (m *Salon) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Salon) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.PhoneNumber) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.PhoneNumber)))
		i += copy(dAtA[i:], m.PhoneNumber)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	if m.Rating != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSalon(dAtA, i, uint64(m.Rating))
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Latitude) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.Latitude)))
		i += copy(dAtA[i:], m.Latitude)
	}
	if len(m.Longitude) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.Longitude)))
		i += copy(dAtA[i:], m.Longitude)
	}
	if len(m.StartTime) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.StartTime)))
		i += copy(dAtA[i:], m.StartTime)
	}
	if len(m.EndTime) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.EndTime)))
		i += copy(dAtA[i:], m.EndTime)
	}
	if len(m.ImageUrl) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.ImageUrl)))
		i += copy(dAtA[i:], m.ImageUrl)
	}
	if len(m.CreatedAt) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.CreatedAt)))
		i += copy(dAtA[i:], m.CreatedAt)
	}
	if len(m.UpdatedAt) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.UpdatedAt)))
		i += copy(dAtA[i:], m.UpdatedAt)
	}
	if len(m.DeletedAt) > 0 {
		dAtA[i] = 0x72
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.DeletedAt)))
		i += copy(dAtA[i:], m.DeletedAt)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetSalonsParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSalonsParams) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Page != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSalon(dAtA, i, uint64(m.Page))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSalon(dAtA, i, uint64(m.Limit))
	}
	if len(m.Search) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSalon(dAtA, i, uint64(len(m.Search)))
		i += copy(dAtA[i:], m.Search)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AllSalons) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllSalons) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Salons) > 0 {
		for _, msg := range m.Salons {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSalon(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Count != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSalon(dAtA, i, uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSalon(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Salon) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.PhoneNumber)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	if m.Rating != 0 {
		n += 1 + sovSalon(uint64(m.Rating))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.Latitude)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.Longitude)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.StartTime)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.EndTime)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.ImageUrl)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	l = len(m.DeletedAt)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetSalonsParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Page != 0 {
		n += 1 + sovSalon(uint64(m.Page))
	}
	if m.Limit != 0 {
		n += 1 + sovSalon(uint64(m.Limit))
	}
	l = len(m.Search)
	if l > 0 {
		n += 1 + l + sovSalon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AllSalons) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Salons) > 0 {
		for _, e := range m.Salons {
			l = e.Size()
			n += 1 + l + sovSalon(uint64(l))
		}
	}
	if m.Count != 0 {
		n += 1 + sovSalon(uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSalon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSalon(x uint64) (n int) {
	return sovSalon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Salon) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSalon
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
			return fmt.Errorf("proto: Salon: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Salon: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhoneNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PhoneNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rating", wireType)
			}
			m.Rating = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rating |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Latitude = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Longitude = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSalon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSalon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSalon
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
func (m *GetSalonsParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSalon
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
			return fmt.Errorf("proto: GetSalonsParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSalonsParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Search", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Search = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSalon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSalon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSalon
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
func (m *AllSalons) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSalon
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
			return fmt.Errorf("proto: AllSalons: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllSalons: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salons", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
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
				return ErrInvalidLengthSalon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSalon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salons = append(m.Salons, &Salon{})
			if err := m.Salons[len(m.Salons)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSalon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSalon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSalon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSalon
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
func skipSalon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSalon
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
					return 0, ErrIntOverflowSalon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSalon
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
				return 0, ErrInvalidLengthSalon
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSalon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSalon
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSalon(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSalon
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSalon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSalon   = fmt.Errorf("proto: integer overflow")
)
