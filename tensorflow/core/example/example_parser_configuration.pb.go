// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/example/example_parser_configuration.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow1 "github.com/desertbit/tfutils/tensorflow/core/framework"
import tensorflow4 "github.com/desertbit/tfutils/tensorflow/core/framework"
import tensorflow3 "github.com/desertbit/tfutils/tensorflow/core/framework"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VarLenFeatureProto struct {
	Dtype                   tensorflow3.DataType `protobuf:"varint,1,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
	ValuesOutputTensorName  string               `protobuf:"bytes,2,opt,name=values_output_tensor_name,json=valuesOutputTensorName" json:"values_output_tensor_name,omitempty"`
	IndicesOutputTensorName string               `protobuf:"bytes,3,opt,name=indices_output_tensor_name,json=indicesOutputTensorName" json:"indices_output_tensor_name,omitempty"`
	ShapesOutputTensorName  string               `protobuf:"bytes,4,opt,name=shapes_output_tensor_name,json=shapesOutputTensorName" json:"shapes_output_tensor_name,omitempty"`
}

func (m *VarLenFeatureProto) Reset()                    { *m = VarLenFeatureProto{} }
func (m *VarLenFeatureProto) String() string            { return proto.CompactTextString(m) }
func (*VarLenFeatureProto) ProtoMessage()               {}
func (*VarLenFeatureProto) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *VarLenFeatureProto) GetDtype() tensorflow3.DataType {
	if m != nil {
		return m.Dtype
	}
	return tensorflow3.DataType_DT_INVALID
}

func (m *VarLenFeatureProto) GetValuesOutputTensorName() string {
	if m != nil {
		return m.ValuesOutputTensorName
	}
	return ""
}

func (m *VarLenFeatureProto) GetIndicesOutputTensorName() string {
	if m != nil {
		return m.IndicesOutputTensorName
	}
	return ""
}

func (m *VarLenFeatureProto) GetShapesOutputTensorName() string {
	if m != nil {
		return m.ShapesOutputTensorName
	}
	return ""
}

type FixedLenFeatureProto struct {
	Dtype                  tensorflow3.DataType          `protobuf:"varint,1,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                  *tensorflow1.TensorShapeProto `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
	DefaultValue           *tensorflow4.TensorProto      `protobuf:"bytes,3,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	ValuesOutputTensorName string                        `protobuf:"bytes,4,opt,name=values_output_tensor_name,json=valuesOutputTensorName" json:"values_output_tensor_name,omitempty"`
}

func (m *FixedLenFeatureProto) Reset()                    { *m = FixedLenFeatureProto{} }
func (m *FixedLenFeatureProto) String() string            { return proto.CompactTextString(m) }
func (*FixedLenFeatureProto) ProtoMessage()               {}
func (*FixedLenFeatureProto) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *FixedLenFeatureProto) GetDtype() tensorflow3.DataType {
	if m != nil {
		return m.Dtype
	}
	return tensorflow3.DataType_DT_INVALID
}

func (m *FixedLenFeatureProto) GetShape() *tensorflow1.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *FixedLenFeatureProto) GetDefaultValue() *tensorflow4.TensorProto {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *FixedLenFeatureProto) GetValuesOutputTensorName() string {
	if m != nil {
		return m.ValuesOutputTensorName
	}
	return ""
}

type FeatureConfiguration struct {
	// Types that are valid to be assigned to Config:
	//	*FeatureConfiguration_FixedLenFeature
	//	*FeatureConfiguration_VarLenFeature
	Config isFeatureConfiguration_Config `protobuf_oneof:"config"`
}

func (m *FeatureConfiguration) Reset()                    { *m = FeatureConfiguration{} }
func (m *FeatureConfiguration) String() string            { return proto.CompactTextString(m) }
func (*FeatureConfiguration) ProtoMessage()               {}
func (*FeatureConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type isFeatureConfiguration_Config interface {
	isFeatureConfiguration_Config()
}

type FeatureConfiguration_FixedLenFeature struct {
	FixedLenFeature *FixedLenFeatureProto `protobuf:"bytes,1,opt,name=fixed_len_feature,json=fixedLenFeature,oneof"`
}
type FeatureConfiguration_VarLenFeature struct {
	VarLenFeature *VarLenFeatureProto `protobuf:"bytes,2,opt,name=var_len_feature,json=varLenFeature,oneof"`
}

func (*FeatureConfiguration_FixedLenFeature) isFeatureConfiguration_Config() {}
func (*FeatureConfiguration_VarLenFeature) isFeatureConfiguration_Config()   {}

func (m *FeatureConfiguration) GetConfig() isFeatureConfiguration_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *FeatureConfiguration) GetFixedLenFeature() *FixedLenFeatureProto {
	if x, ok := m.GetConfig().(*FeatureConfiguration_FixedLenFeature); ok {
		return x.FixedLenFeature
	}
	return nil
}

func (m *FeatureConfiguration) GetVarLenFeature() *VarLenFeatureProto {
	if x, ok := m.GetConfig().(*FeatureConfiguration_VarLenFeature); ok {
		return x.VarLenFeature
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FeatureConfiguration) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FeatureConfiguration_OneofMarshaler, _FeatureConfiguration_OneofUnmarshaler, _FeatureConfiguration_OneofSizer, []interface{}{
		(*FeatureConfiguration_FixedLenFeature)(nil),
		(*FeatureConfiguration_VarLenFeature)(nil),
	}
}

func _FeatureConfiguration_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FeatureConfiguration)
	// config
	switch x := m.Config.(type) {
	case *FeatureConfiguration_FixedLenFeature:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixedLenFeature); err != nil {
			return err
		}
	case *FeatureConfiguration_VarLenFeature:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VarLenFeature); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FeatureConfiguration.Config has unexpected type %T", x)
	}
	return nil
}

func _FeatureConfiguration_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FeatureConfiguration)
	switch tag {
	case 1: // config.fixed_len_feature
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FixedLenFeatureProto)
		err := b.DecodeMessage(msg)
		m.Config = &FeatureConfiguration_FixedLenFeature{msg}
		return true, err
	case 2: // config.var_len_feature
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VarLenFeatureProto)
		err := b.DecodeMessage(msg)
		m.Config = &FeatureConfiguration_VarLenFeature{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FeatureConfiguration_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FeatureConfiguration)
	// config
	switch x := m.Config.(type) {
	case *FeatureConfiguration_FixedLenFeature:
		s := proto.Size(x.FixedLenFeature)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FeatureConfiguration_VarLenFeature:
		s := proto.Size(x.VarLenFeature)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ExampleParserConfiguration struct {
	FeatureMap map[string]*FeatureConfiguration `protobuf:"bytes,1,rep,name=feature_map,json=featureMap" json:"feature_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ExampleParserConfiguration) Reset()                    { *m = ExampleParserConfiguration{} }
func (m *ExampleParserConfiguration) String() string            { return proto.CompactTextString(m) }
func (*ExampleParserConfiguration) ProtoMessage()               {}
func (*ExampleParserConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ExampleParserConfiguration) GetFeatureMap() map[string]*FeatureConfiguration {
	if m != nil {
		return m.FeatureMap
	}
	return nil
}

func init() {
	proto.RegisterType((*VarLenFeatureProto)(nil), "tensorflow.VarLenFeatureProto")
	proto.RegisterType((*FixedLenFeatureProto)(nil), "tensorflow.FixedLenFeatureProto")
	proto.RegisterType((*FeatureConfiguration)(nil), "tensorflow.FeatureConfiguration")
	proto.RegisterType((*ExampleParserConfiguration)(nil), "tensorflow.ExampleParserConfiguration")
}

func init() {
	proto.RegisterFile("tensorflow/core/example/example_parser_configuration.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0xe7, 0x76, 0x9d, 0xde, 0x9d, 0xbc, 0xa3, 0x60, 0x4d, 0x5b, 0x17, 0x21, 0x14, 0x55,
	0x02, 0x55, 0x08, 0xb5, 0x52, 0x90, 0x26, 0x36, 0x90, 0x90, 0x0a, 0x9b, 0x7a, 0x01, 0xa3, 0x0a,
	0xd3, 0xb8, 0xb4, 0x4c, 0x7b, 0x32, 0xa2, 0x25, 0xb1, 0xe5, 0x38, 0xdd, 0xfa, 0xd5, 0xf8, 0x20,
	0x7c, 0x0d, 0x2e, 0xb8, 0xe1, 0x12, 0xc5, 0x8e, 0xb6, 0xb4, 0xcd, 0xca, 0x05, 0x57, 0x89, 0x7c,
	0x7e, 0xcf, 0xf9, 0xf3, 0xc4, 0x27, 0x70, 0xac, 0x31, 0xcd, 0x84, 0x0a, 0x63, 0x71, 0x3d, 0x98,
	0x08, 0x85, 0x03, 0xbc, 0xe1, 0x89, 0x8c, 0x6f, 0x9f, 0x4c, 0x72, 0x95, 0xa1, 0x62, 0x13, 0x91,
	0x86, 0xd1, 0x65, 0xae, 0xb8, 0x8e, 0x44, 0xda, 0x97, 0x4a, 0x68, 0x41, 0xe1, 0x4e, 0xeb, 0xbe,
	0x58, 0xce, 0x13, 0x2a, 0x9e, 0xe0, 0xb5, 0x50, 0x57, 0x03, 0x1b, 0x61, 0xd9, 0x37, 0x2e, 0xd1,
	0x2a, 0xdd, 0x67, 0x7f, 0xa3, 0x4b, 0xee, 0xe9, 0x1a, 0x6e, 0x2e, 0x31, 0xb3, 0x58, 0xf7, 0x17,
	0x01, 0x7a, 0xc1, 0xd5, 0x07, 0x4c, 0x4f, 0x91, 0xeb, 0x5c, 0xe1, 0xd8, 0xf4, 0xf7, 0x1c, 0x5a,
	0xd3, 0x02, 0xeb, 0x10, 0x8f, 0xf4, 0x1e, 0xf8, 0xbb, 0xfd, 0xbb, 0x6c, 0xfd, 0xf7, 0x5c, 0xf3,
	0xf3, 0xb9, 0xc4, 0xc0, 0x22, 0xf4, 0x08, 0x0e, 0x66, 0x3c, 0xce, 0x31, 0x63, 0x22, 0xd7, 0x32,
	0xd7, 0xac, 0xec, 0x3a, 0xe5, 0x09, 0x76, 0x1a, 0x1e, 0xe9, 0x6d, 0x07, 0x7b, 0x16, 0xf8, 0x64,
	0xe2, 0xe7, 0x26, 0x7c, 0xc6, 0x13, 0xa4, 0xaf, 0xc1, 0x8d, 0xd2, 0x69, 0x34, 0xa9, 0xd7, 0x36,
	0x8d, 0x76, 0xbf, 0x24, 0x56, 0xc4, 0x47, 0x70, 0x60, 0x8c, 0xa9, 0xd5, 0x6e, 0xda, 0xba, 0x16,
	0x58, 0x96, 0x76, 0x7f, 0x12, 0xd8, 0x3d, 0x8d, 0x6e, 0x70, 0xfa, 0x2f, 0x73, 0xfb, 0xd0, 0x32,
	0xe9, 0xcd, 0x8c, 0x8e, 0xff, 0xb8, 0xca, 0xda, 0x5a, 0x9f, 0x8b, 0xb0, 0x49, 0x1c, 0x58, 0x94,
	0xbe, 0x81, 0x9d, 0x29, 0x86, 0x3c, 0x8f, 0x35, 0x33, 0x96, 0x98, 0x19, 0x1d, 0x7f, 0x7f, 0x55,
	0x6b, 0x65, 0xff, 0x97, 0xf4, 0x45, 0x01, 0xaf, 0x77, 0x7a, 0x73, 0x9d, 0xd3, 0xdd, 0xef, 0xc5,
	0xc4, 0x76, 0xd2, 0x77, 0xd5, 0xfb, 0x48, 0xcf, 0xe0, 0x51, 0x58, 0x38, 0xc1, 0x62, 0x4c, 0x59,
	0x68, 0x09, 0x33, 0xbd, 0xe3, 0x7b, 0xd5, 0xae, 0xea, 0xec, 0x1a, 0x6d, 0x04, 0xed, 0x70, 0xf1,
	0x9c, 0x8e, 0xa0, 0x3d, 0xe3, 0x6a, 0x21, 0x9b, 0xf5, 0xe7, 0x49, 0x35, 0xdb, 0xea, 0x95, 0x1b,
	0x6d, 0x04, 0x3b, 0xb3, 0xea, 0xe9, 0xf0, 0x3f, 0xd8, 0xb2, 0xab, 0xd3, 0xfd, 0x41, 0xc0, 0x3d,
	0xb1, 0x4b, 0x35, 0x36, 0x3b, 0xb5, 0x38, 0xc2, 0x17, 0x70, 0xca, 0x52, 0x2c, 0xe1, 0xb2, 0x43,
	0xbc, 0x66, 0xcf, 0xf1, 0x0f, 0xab, 0xe5, 0xee, 0x17, 0xf7, 0xcb, 0x6a, 0x1f, 0xb9, 0x3c, 0x49,
	0xb5, 0x9a, 0x07, 0x10, 0xde, 0x1e, 0xb8, 0x0c, 0xda, 0x4b, 0x61, 0xfa, 0x10, 0x9a, 0x57, 0x38,
	0x37, 0x06, 0x6d, 0x07, 0xc5, 0x2b, 0x3d, 0x84, 0x96, 0xfd, 0x94, 0x8d, 0x1a, 0xd3, 0x6a, 0x1c,
	0x0f, 0x2c, 0x7e, 0xdc, 0x78, 0x45, 0x86, 0x6f, 0x61, 0x4f, 0xa8, 0xcb, 0xaa, 0xa2, 0xfc, 0x77,
	0x0c, 0xbd, 0xfb, 0x5b, 0x36, 0x76, 0x65, 0x63, 0xf2, 0x9b, 0x90, 0xaf, 0x5b, 0x66, 0x8b, 0x5f,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x37, 0xba, 0xa0, 0xf8, 0x8c, 0x04, 0x00, 0x00,
}
