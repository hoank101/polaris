// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polaris/evm/v1alpha1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// WrappedPayloadEnvelope encapsulates an Ethereum transaction as an SDK message.
type WrappedPayloadEnvelope struct {
	// data is inner transaction data of the Ethereum transaction.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WrappedPayloadEnvelope) Reset()         { *m = WrappedPayloadEnvelope{} }
func (m *WrappedPayloadEnvelope) String() string { return proto.CompactTextString(m) }
func (*WrappedPayloadEnvelope) ProtoMessage()    {}
func (*WrappedPayloadEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{0}
}
func (m *WrappedPayloadEnvelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedPayloadEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedPayloadEnvelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedPayloadEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedPayloadEnvelope.Merge(m, src)
}
func (m *WrappedPayloadEnvelope) XXX_Size() int {
	return m.Size()
}
func (m *WrappedPayloadEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedPayloadEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedPayloadEnvelope proto.InternalMessageInfo

func (m *WrappedPayloadEnvelope) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// WrappedPayloadEnvelopeResponse defines the block response.
type WrappedPayloadEnvelopeResponse struct {
}

func (m *WrappedPayloadEnvelopeResponse) Reset()         { *m = WrappedPayloadEnvelopeResponse{} }
func (m *WrappedPayloadEnvelopeResponse) String() string { return proto.CompactTextString(m) }
func (*WrappedPayloadEnvelopeResponse) ProtoMessage()    {}
func (*WrappedPayloadEnvelopeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{1}
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedPayloadEnvelopeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedPayloadEnvelopeResponse.Merge(m, src)
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Size() int {
	return m.Size()
}
func (m *WrappedPayloadEnvelopeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedPayloadEnvelopeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedPayloadEnvelopeResponse proto.InternalMessageInfo

// WrappedEthereumTransaction encapsulates an Ethereum transaction as an SDK message.
type WrappedEthereumTransaction struct {
	// data is inner transaction data of the Ethereum transaction.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WrappedEthereumTransaction) Reset()         { *m = WrappedEthereumTransaction{} }
func (m *WrappedEthereumTransaction) String() string { return proto.CompactTextString(m) }
func (*WrappedEthereumTransaction) ProtoMessage()    {}
func (*WrappedEthereumTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b33d2a2c64400f, []int{2}
}
func (m *WrappedEthereumTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedEthereumTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedEthereumTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedEthereumTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedEthereumTransaction.Merge(m, src)
}
func (m *WrappedEthereumTransaction) XXX_Size() int {
	return m.Size()
}
func (m *WrappedEthereumTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedEthereumTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedEthereumTransaction proto.InternalMessageInfo

func (m *WrappedEthereumTransaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*WrappedPayloadEnvelope)(nil), "polaris.evm.v1alpha1.WrappedPayloadEnvelope")
	proto.RegisterType((*WrappedPayloadEnvelopeResponse)(nil), "polaris.evm.v1alpha1.WrappedPayloadEnvelopeResponse")
	proto.RegisterType((*WrappedEthereumTransaction)(nil), "polaris.evm.v1alpha1.WrappedEthereumTransaction")
}

func init() { proto.RegisterFile("polaris/evm/v1alpha1/tx.proto", fileDescriptor_d8b33d2a2c64400f) }

var fileDescriptor_d8b33d2a2c64400f = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0xf4, 0x40,
	0x14, 0x85, 0x77, 0xe0, 0xff, 0x2d, 0x06, 0xab, 0x20, 0xab, 0x04, 0x1c, 0x96, 0xad, 0x44, 0x96,
	0x19, 0x57, 0x7d, 0x02, 0x61, 0x0b, 0x0b, 0x61, 0x59, 0x05, 0xc1, 0xee, 0x66, 0x72, 0x49, 0x02,
	0x99, 0xcc, 0x30, 0x33, 0x3b, 0xec, 0x82, 0x85, 0x8f, 0xa0, 0x6f, 0xe2, 0x63, 0x58, 0x6e, 0x69,
	0x29, 0x49, 0xe1, 0x6b, 0x88, 0x31, 0x69, 0x24, 0x8d, 0xd5, 0xbd, 0x70, 0x3e, 0x2e, 0xe7, 0x9e,
	0x43, 0x8f, 0x8d, 0x2e, 0xc1, 0x16, 0x4e, 0x60, 0x50, 0x22, 0xcc, 0xa1, 0x34, 0x39, 0xcc, 0x85,
	0xdf, 0x70, 0x63, 0xb5, 0xd7, 0xd1, 0x41, 0x27, 0x73, 0x0c, 0x8a, 0xf7, 0x72, 0x7c, 0x28, 0xb5,
	0x53, 0xda, 0x09, 0xe5, 0x32, 0x11, 0xe6, 0xdf, 0xe3, 0x07, 0x9f, 0xce, 0xe8, 0xf8, 0xde, 0x82,
	0x31, 0x98, 0x2e, 0x61, 0x5b, 0x6a, 0x48, 0x17, 0x55, 0xc0, 0x52, 0x1b, 0x8c, 0x22, 0xfa, 0x2f,
	0x05, 0x0f, 0x47, 0x64, 0x42, 0x4e, 0xf6, 0x57, 0xed, 0x3e, 0x9d, 0x50, 0x36, 0x4c, 0xaf, 0xd0,
	0x19, 0x5d, 0x39, 0x9c, 0x9e, 0xd1, 0xb8, 0x23, 0x16, 0x3e, 0x47, 0x8b, 0x6b, 0x75, 0x67, 0xa1,
	0x72, 0x20, 0x7d, 0xa1, 0xab, 0xa1, 0x9b, 0xe7, 0x2f, 0x84, 0xd2, 0x1b, 0x97, 0xdd, 0xa2, 0x0d,
	0x85, 0xc4, 0xe8, 0x91, 0x8e, 0x97, 0x56, 0x4b, 0x74, 0xee, 0xb7, 0xa1, 0x19, 0x1f, 0x7a, 0x8d,
	0x0f, 0x1b, 0x8a, 0x2f, 0xff, 0x42, 0xf7, 0xf6, 0xe3, 0xff, 0x4f, 0x9f, 0xaf, 0xa7, 0xe4, 0xea,
	0xfa, 0xad, 0x66, 0x64, 0x57, 0x33, 0xf2, 0x51, 0x33, 0xf2, 0xdc, 0xb0, 0xd1, 0xae, 0x61, 0xa3,
	0xf7, 0x86, 0x8d, 0x1e, 0x44, 0x56, 0xf8, 0x7c, 0x9d, 0x70, 0xa9, 0x95, 0x48, 0xd0, 0x82, 0xcc,
	0xa1, 0xa8, 0x44, 0x5f, 0x49, 0x97, 0xf2, 0xa6, 0xed, 0xc6, 0x6f, 0x0d, 0xba, 0x64, 0xaf, 0xcd,
	0xf9, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x58, 0x64, 0x6d, 0xb7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	// ProcessPayloadEnvelope defines a method to process CL paylods.
	ProcessPayloadEnvelope(ctx context.Context, in *WrappedPayloadEnvelope, opts ...grpc.CallOption) (*WrappedPayloadEnvelopeResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) ProcessPayloadEnvelope(ctx context.Context, in *WrappedPayloadEnvelope, opts ...grpc.CallOption) (*WrappedPayloadEnvelopeResponse, error) {
	out := new(WrappedPayloadEnvelopeResponse)
	err := c.cc.Invoke(ctx, "/polaris.evm.v1alpha1.MsgService/ProcessPayloadEnvelope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// ProcessPayloadEnvelope defines a method to process CL paylods.
	ProcessPayloadEnvelope(context.Context, *WrappedPayloadEnvelope) (*WrappedPayloadEnvelopeResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) ProcessPayloadEnvelope(ctx context.Context, req *WrappedPayloadEnvelope) (*WrappedPayloadEnvelopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayloadEnvelope not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_ProcessPayloadEnvelope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrappedPayloadEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ProcessPayloadEnvelope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polaris.evm.v1alpha1.MsgService/ProcessPayloadEnvelope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ProcessPayloadEnvelope(ctx, req.(*WrappedPayloadEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "polaris.evm.v1alpha1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPayloadEnvelope",
			Handler:    _MsgService_ProcessPayloadEnvelope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "polaris/evm/v1alpha1/tx.proto",
}

func (m *WrappedPayloadEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedPayloadEnvelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedPayloadEnvelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WrappedPayloadEnvelopeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedPayloadEnvelopeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedPayloadEnvelopeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *WrappedEthereumTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedEthereumTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedEthereumTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WrappedPayloadEnvelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *WrappedPayloadEnvelopeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *WrappedEthereumTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WrappedPayloadEnvelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: WrappedPayloadEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedPayloadEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *WrappedPayloadEnvelopeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: WrappedPayloadEnvelopeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedPayloadEnvelopeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *WrappedEthereumTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: WrappedEthereumTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedEthereumTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
