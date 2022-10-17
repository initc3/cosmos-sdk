// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/base/node/v1beta1/query.proto

package node

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ConfigRequest defines the request structure for the Config gRPC query.
type ConfigRequest struct {
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8324226a07064341, []int{0}
}
func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

// ConfigResponse defines the response structure for the Config gRPC query.
type ConfigResponse struct {
	MinimumGasPrice string `protobuf:"bytes,1,opt,name=minimum_gas_price,json=minimumGasPrice,proto3" json:"minimum_gas_price,omitempty"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8324226a07064341, []int{1}
}
func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(m, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetMinimumGasPrice() string {
	if m != nil {
		return m.MinimumGasPrice
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "cosmos.base.node.v1beta1.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "cosmos.base.node.v1beta1.ConfigResponse")
}

func init() {
	proto.RegisterFile("cosmos/base/node/v1beta1/query.proto", fileDescriptor_8324226a07064341)
}

var fileDescriptor_8324226a07064341 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4b, 0xc3, 0x40,
	0x18, 0xc5, 0x73, 0x0e, 0x15, 0x0f, 0xb4, 0x98, 0xa9, 0x14, 0x39, 0x4a, 0x10, 0x0c, 0x42, 0xef,
	0xa8, 0xae, 0x4e, 0x3a, 0x74, 0x95, 0xba, 0xb9, 0x94, 0xcb, 0xf5, 0xf3, 0x3c, 0x6c, 0xee, 0x4b,
	0x73, 0x97, 0x82, 0xab, 0xe0, 0xae, 0xf8, 0x4f, 0x39, 0x16, 0x5c, 0x1c, 0x25, 0xf1, 0x0f, 0x91,
	0x24, 0xed, 0xe0, 0x50, 0x3a, 0x1d, 0xbc, 0xfb, 0xbd, 0xf7, 0x3d, 0x1e, 0x3d, 0x55, 0xe8, 0x52,
	0x74, 0x22, 0x91, 0x0e, 0x84, 0xc5, 0x19, 0x88, 0xe5, 0x28, 0x01, 0x2f, 0x47, 0x62, 0x51, 0x40,
	0xfe, 0xcc, 0xb3, 0x1c, 0x3d, 0x86, 0xbd, 0x96, 0xe2, 0x35, 0xc5, 0x6b, 0x8a, 0xaf, 0xa9, 0xfe,
	0x89, 0x46, 0xd4, 0x73, 0x10, 0x32, 0x33, 0x42, 0x5a, 0x8b, 0x5e, 0x7a, 0x83, 0xd6, 0xb5, 0xbe,
	0xa8, 0x4b, 0x0f, 0x6f, 0xd0, 0x3e, 0x18, 0x3d, 0x81, 0x45, 0x01, 0xce, 0x47, 0x57, 0xf4, 0x68,
	0x23, 0xb8, 0x0c, 0xad, 0x83, 0xf0, 0x9c, 0x1e, 0xa7, 0xc6, 0x9a, 0xb4, 0x48, 0xa7, 0x5a, 0xba,
	0x69, 0x96, 0x1b, 0x05, 0x3d, 0x32, 0x20, 0xf1, 0xc1, 0xa4, 0xbb, 0xfe, 0x18, 0x4b, 0x77, 0x5b,
	0xcb, 0x17, 0xef, 0x84, 0xee, 0xdf, 0x41, 0xbe, 0x34, 0x0a, 0xc2, 0x57, 0x42, 0x3b, 0x6d, 0x54,
	0x78, 0xc6, 0xb7, 0xd5, 0xe3, 0xff, 0xae, 0xf7, 0xe3, 0xdd, 0x60, 0xdb, 0x2a, 0x8a, 0x5f, 0xbe,
	0x7e, 0x3f, 0xf6, 0xa2, 0x70, 0x20, 0xb6, 0xee, 0xa3, 0x1a, 0xc7, 0xf5, 0xf8, 0xb3, 0x64, 0x64,
	0x55, 0x32, 0xf2, 0x53, 0x32, 0xf2, 0x56, 0xb1, 0x60, 0x55, 0xb1, 0xe0, 0xbb, 0x62, 0xc1, 0xfd,
	0x50, 0x1b, 0xff, 0x58, 0x24, 0x5c, 0x61, 0xba, 0x49, 0x69, 0x9f, 0xa1, 0x9b, 0x3d, 0x09, 0x35,
	0x37, 0x60, 0xbd, 0xd0, 0x79, 0xa6, 0x9a, 0xdc, 0xa4, 0xd3, 0x4c, 0x76, 0xf9, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x7d, 0x46, 0xb4, 0x93, 0x92, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// Config queries for the operator configuration.
	Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.node.v1beta1.Service/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// Config queries for the operator configuration.
	Config(context.Context, *ConfigRequest) (*ConfigResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Config(ctx context.Context, req *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.node.v1beta1.Service/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Config(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.node.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Config",
			Handler:    _Service_Config_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/node/v1beta1/query.proto",
}

func (m *ConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinimumGasPrice) > 0 {
		i -= len(m.MinimumGasPrice)
		copy(dAtA[i:], m.MinimumGasPrice)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MinimumGasPrice)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ConfigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MinimumGasPrice)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumGasPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinimumGasPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)