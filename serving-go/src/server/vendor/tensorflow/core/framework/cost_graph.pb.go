// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/cost_graph.proto

package tensorflow

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

type CostGraphDef struct {
	Node                 []*CostGraphDef_Node `protobuf:"bytes,1,rep,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CostGraphDef) Reset()         { *m = CostGraphDef{} }
func (m *CostGraphDef) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef) ProtoMessage()    {}
func (*CostGraphDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0}
}

func (m *CostGraphDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef.Unmarshal(m, b)
}
func (m *CostGraphDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef.Marshal(b, m, deterministic)
}
func (m *CostGraphDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef.Merge(m, src)
}
func (m *CostGraphDef) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef.Size(m)
}
func (m *CostGraphDef) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef proto.InternalMessageInfo

func (m *CostGraphDef) GetNode() []*CostGraphDef_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type CostGraphDef_Node struct {
	// The name of the node. Names are globally unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The device of the node. Can be empty if the node is mapped to the
	// default partition or partitioning hasn't been run yet.
	Device string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	// The id of the node. Node ids are only unique inside a partition.
	Id         int32                           `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	InputInfo  []*CostGraphDef_Node_InputInfo  `protobuf:"bytes,4,rep,name=input_info,json=inputInfo,proto3" json:"input_info,omitempty"`
	OutputInfo []*CostGraphDef_Node_OutputInfo `protobuf:"bytes,5,rep,name=output_info,json=outputInfo,proto3" json:"output_info,omitempty"`
	// Temporary memory used by this node.
	TemporaryMemorySize int64 `protobuf:"varint,6,opt,name=temporary_memory_size,json=temporaryMemorySize,proto3" json:"temporary_memory_size,omitempty"`
	// Estimate of the computational cost of this node.
	ComputeCost int64 `protobuf:"varint,9,opt,name=compute_cost,json=computeCost,proto3" json:"compute_cost,omitempty"`
	// If true, the output is permanent: it can't be discarded, because this
	// node is part of the "final output". Nodes may depend on final nodes.
	IsFinal bool `protobuf:"varint,7,opt,name=is_final,json=isFinal,proto3" json:"is_final,omitempty"`
	// Ids of the control inputs for this node.
	ControlInput         []int32  `protobuf:"varint,8,rep,packed,name=control_input,json=controlInput,proto3" json:"control_input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostGraphDef_Node) Reset()         { *m = CostGraphDef_Node{} }
func (m *CostGraphDef_Node) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_Node) ProtoMessage()    {}
func (*CostGraphDef_Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0, 0}
}

func (m *CostGraphDef_Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef_Node.Unmarshal(m, b)
}
func (m *CostGraphDef_Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef_Node.Marshal(b, m, deterministic)
}
func (m *CostGraphDef_Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef_Node.Merge(m, src)
}
func (m *CostGraphDef_Node) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef_Node.Size(m)
}
func (m *CostGraphDef_Node) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef_Node.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef_Node proto.InternalMessageInfo

func (m *CostGraphDef_Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CostGraphDef_Node) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *CostGraphDef_Node) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CostGraphDef_Node) GetInputInfo() []*CostGraphDef_Node_InputInfo {
	if m != nil {
		return m.InputInfo
	}
	return nil
}

func (m *CostGraphDef_Node) GetOutputInfo() []*CostGraphDef_Node_OutputInfo {
	if m != nil {
		return m.OutputInfo
	}
	return nil
}

func (m *CostGraphDef_Node) GetTemporaryMemorySize() int64 {
	if m != nil {
		return m.TemporaryMemorySize
	}
	return 0
}

func (m *CostGraphDef_Node) GetComputeCost() int64 {
	if m != nil {
		return m.ComputeCost
	}
	return 0
}

func (m *CostGraphDef_Node) GetIsFinal() bool {
	if m != nil {
		return m.IsFinal
	}
	return false
}

func (m *CostGraphDef_Node) GetControlInput() []int32 {
	if m != nil {
		return m.ControlInput
	}
	return nil
}

// Inputs of this node. They must be executed before this node can be
// executed. An input is a particular output of another node, specified
// by the node id and the output index.
type CostGraphDef_Node_InputInfo struct {
	PrecedingNode        int32    `protobuf:"varint,1,opt,name=preceding_node,json=precedingNode,proto3" json:"preceding_node,omitempty"`
	PrecedingPort        int32    `protobuf:"varint,2,opt,name=preceding_port,json=precedingPort,proto3" json:"preceding_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostGraphDef_Node_InputInfo) Reset()         { *m = CostGraphDef_Node_InputInfo{} }
func (m *CostGraphDef_Node_InputInfo) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_Node_InputInfo) ProtoMessage()    {}
func (*CostGraphDef_Node_InputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0, 0, 0}
}

func (m *CostGraphDef_Node_InputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef_Node_InputInfo.Unmarshal(m, b)
}
func (m *CostGraphDef_Node_InputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef_Node_InputInfo.Marshal(b, m, deterministic)
}
func (m *CostGraphDef_Node_InputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef_Node_InputInfo.Merge(m, src)
}
func (m *CostGraphDef_Node_InputInfo) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef_Node_InputInfo.Size(m)
}
func (m *CostGraphDef_Node_InputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef_Node_InputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef_Node_InputInfo proto.InternalMessageInfo

func (m *CostGraphDef_Node_InputInfo) GetPrecedingNode() int32 {
	if m != nil {
		return m.PrecedingNode
	}
	return 0
}

func (m *CostGraphDef_Node_InputInfo) GetPrecedingPort() int32 {
	if m != nil {
		return m.PrecedingPort
	}
	return 0
}

// Outputs of this node.
type CostGraphDef_Node_OutputInfo struct {
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// If >= 0, the output is an alias of an input. Note that an alias input
	// may itself be an alias. The algorithm will therefore need to follow
	// those pointers.
	AliasInputPort       int64             `protobuf:"varint,2,opt,name=alias_input_port,json=aliasInputPort,proto3" json:"alias_input_port,omitempty"`
	Shape                *TensorShapeProto `protobuf:"bytes,3,opt,name=shape,proto3" json:"shape,omitempty"`
	Dtype                DataType          `protobuf:"varint,4,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CostGraphDef_Node_OutputInfo) Reset()         { *m = CostGraphDef_Node_OutputInfo{} }
func (m *CostGraphDef_Node_OutputInfo) String() string { return proto.CompactTextString(m) }
func (*CostGraphDef_Node_OutputInfo) ProtoMessage()    {}
func (*CostGraphDef_Node_OutputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f8948141565ace8, []int{0, 0, 1}
}

func (m *CostGraphDef_Node_OutputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostGraphDef_Node_OutputInfo.Unmarshal(m, b)
}
func (m *CostGraphDef_Node_OutputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostGraphDef_Node_OutputInfo.Marshal(b, m, deterministic)
}
func (m *CostGraphDef_Node_OutputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostGraphDef_Node_OutputInfo.Merge(m, src)
}
func (m *CostGraphDef_Node_OutputInfo) XXX_Size() int {
	return xxx_messageInfo_CostGraphDef_Node_OutputInfo.Size(m)
}
func (m *CostGraphDef_Node_OutputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CostGraphDef_Node_OutputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CostGraphDef_Node_OutputInfo proto.InternalMessageInfo

func (m *CostGraphDef_Node_OutputInfo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CostGraphDef_Node_OutputInfo) GetAliasInputPort() int64 {
	if m != nil {
		return m.AliasInputPort
	}
	return 0
}

func (m *CostGraphDef_Node_OutputInfo) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *CostGraphDef_Node_OutputInfo) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func init() {
	proto.RegisterType((*CostGraphDef)(nil), "tensorflow.CostGraphDef")
	proto.RegisterType((*CostGraphDef_Node)(nil), "tensorflow.CostGraphDef.Node")
	proto.RegisterType((*CostGraphDef_Node_InputInfo)(nil), "tensorflow.CostGraphDef.Node.InputInfo")
	proto.RegisterType((*CostGraphDef_Node_OutputInfo)(nil), "tensorflow.CostGraphDef.Node.OutputInfo")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/cost_graph.proto", fileDescriptor_5f8948141565ace8)
}

var fileDescriptor_5f8948141565ace8 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb5, 0x8d, 0x9d, 0x26, 0x93, 0x34, 0xa0, 0xe5, 0x8f, 0x4c, 0x04, 0x92, 0x01, 0x55,
	0x58, 0x15, 0x4a, 0x20, 0xbc, 0x41, 0xa9, 0x8a, 0x72, 0x00, 0x22, 0xb7, 0x17, 0x4e, 0xd6, 0x62,
	0xaf, 0xd3, 0x15, 0xb1, 0x67, 0xb5, 0xbb, 0xa1, 0x4a, 0x1f, 0x89, 0x27, 0xe0, 0xc1, 0x38, 0x70,
	0x44, 0x3b, 0x0e, 0x4e, 0x38, 0x90, 0xdb, 0xec, 0xcc, 0x6f, 0xbe, 0xf1, 0xf8, 0x1b, 0x38, 0x73,
	0xb2, 0xb6, 0x68, 0xca, 0x15, 0xde, 0x4e, 0x73, 0x34, 0x72, 0x5a, 0x1a, 0x51, 0xc9, 0x5b, 0x34,
	0xdf, 0xa6, 0x39, 0x5a, 0x97, 0x2d, 0x8d, 0xd0, 0x37, 0x13, 0x6d, 0xd0, 0x21, 0x87, 0x1d, 0x3b,
	0x7e, 0xfd, 0xff, 0xbe, 0xa6, 0x92, 0xd9, 0x1b, 0xa1, 0x65, 0xd3, 0x39, 0x3e, 0x3d, 0x40, 0x6f,
	0xb4, 0xb4, 0x0d, 0xf6, 0xe2, 0x67, 0x08, 0xc3, 0xf7, 0x68, 0xdd, 0x07, 0x3f, 0xf4, 0x42, 0x96,
	0xfc, 0x2d, 0x04, 0x35, 0x16, 0x32, 0x62, 0x71, 0x27, 0x19, 0xcc, 0x9e, 0x4d, 0x76, 0x32, 0x93,
	0x7d, 0x6e, 0xf2, 0x09, 0x0b, 0x99, 0x12, 0x3a, 0xfe, 0x15, 0x40, 0xe0, 0x9f, 0x9c, 0x43, 0x50,
	0x8b, 0xca, 0xf7, 0xb2, 0xa4, 0x9f, 0x52, 0xcc, 0x1f, 0x43, 0xb7, 0x90, 0xdf, 0x55, 0x2e, 0xa3,
	0x23, 0xca, 0x6e, 0x5f, 0x7c, 0x04, 0x47, 0xaa, 0x88, 0x3a, 0x31, 0x4b, 0xc2, 0xf4, 0x48, 0x15,
	0xfc, 0x12, 0x40, 0xd5, 0x7a, 0xed, 0x32, 0x55, 0x97, 0x18, 0x05, 0x34, 0xfd, 0xd5, 0xc1, 0xe9,
	0x93, 0xb9, 0xe7, 0xe7, 0x75, 0x89, 0x69, 0x5f, 0xfd, 0x0d, 0xf9, 0x1c, 0x06, 0xb8, 0x76, 0xad,
	0x50, 0x48, 0x42, 0xc9, 0x61, 0xa1, 0xcf, 0xd4, 0x40, 0x4a, 0x80, 0x6d, 0xcc, 0x67, 0xf0, 0xc8,
	0xc9, 0x4a, 0xa3, 0x11, 0x66, 0x93, 0x55, 0xb2, 0x42, 0xb3, 0xc9, 0xac, 0xba, 0x93, 0x51, 0x37,
	0x66, 0x49, 0x27, 0x7d, 0xd0, 0x16, 0x3f, 0x52, 0xed, 0x4a, 0xdd, 0x49, 0xfe, 0x1c, 0x86, 0x39,
	0x56, 0x7a, 0xed, 0x64, 0xe6, 0xcd, 0x8c, 0xfa, 0x84, 0x0e, 0xb6, 0x39, 0x3f, 0x9a, 0x3f, 0x81,
	0x9e, 0xb2, 0x59, 0xa9, 0x6a, 0xb1, 0x8a, 0x8e, 0x63, 0x96, 0xf4, 0xd2, 0x63, 0x65, 0x2f, 0xfd,
	0x93, 0xbf, 0x84, 0x93, 0x1c, 0x6b, 0x67, 0x70, 0x95, 0xd1, 0x46, 0x51, 0x2f, 0xee, 0x24, 0x61,
	0x3a, 0xdc, 0x26, 0x69, 0xe1, 0xf1, 0x17, 0xe8, 0xb7, 0x9b, 0xf3, 0x53, 0x18, 0x69, 0x23, 0x73,
	0x59, 0xa8, 0x7a, 0x99, 0x6d, 0x8d, 0xf3, 0xbf, 0xf4, 0xa4, 0xcd, 0x92, 0x33, 0xff, 0x60, 0x1a,
	0x8d, 0x23, 0x37, 0xf6, 0xb1, 0x05, 0x1a, 0x37, 0xfe, 0xc1, 0x00, 0x76, 0x3f, 0xc3, 0xfb, 0x49,
	0xfb, 0x32, 0x5a, 0x82, 0x62, 0x9e, 0xc0, 0x7d, 0xb1, 0x52, 0xc2, 0x36, 0x1f, 0xb8, 0xd3, 0xea,
	0xa4, 0x23, 0xca, 0xd3, 0xa7, 0x79, 0x31, 0x3e, 0x83, 0x90, 0x0e, 0x92, 0x4c, 0x1e, 0xcc, 0x9e,
	0xee, 0x7b, 0x70, 0x4d, 0xe1, 0x95, 0x2f, 0x2f, 0xfc, 0x1d, 0xa6, 0x0d, 0xca, 0xcf, 0x20, 0x2c,
	0xfc, 0x79, 0x46, 0x41, 0xcc, 0x92, 0xd1, 0xec, 0xe1, 0x7e, 0xcf, 0x85, 0x70, 0xe2, 0x7a, 0xa3,
	0x65, 0xda, 0x20, 0xe7, 0x6f, 0x20, 0x42, 0xb3, 0xdc, 0x27, 0xda, 0x13, 0x3f, 0xbf, 0xd7, 0x9a,
	0x4c, 0xf2, 0x76, 0xc1, 0x7e, 0x33, 0xf6, 0xb5, 0x4b, 0x37, 0xff, 0xee, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe1, 0x94, 0x65, 0xed, 0x82, 0x03, 0x00, 0x00,
}
