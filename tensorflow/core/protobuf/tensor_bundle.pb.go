// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/tensor_bundle.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow2 "github.com/desertbit/tfutils/tensorflow/core/framework"
import tensorflow11 "github.com/desertbit/tfutils/tensorflow/core/framework"
import tensorflow3 "github.com/desertbit/tfutils/tensorflow/core/framework"
import tensorflow9 "github.com/desertbit/tfutils/tensorflow/core/framework"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An enum indicating the endianness of the platform that produced this
// bundle.  A bundle can only be read by a platform with matching endianness.
// Defaults to LITTLE, as most modern platforms are little-endian.
//
// Affects the binary tensor data bytes only, not the metadata in protobufs.
type BundleHeaderProto_Endianness int32

const (
	BundleHeaderProto_LITTLE BundleHeaderProto_Endianness = 0
	BundleHeaderProto_BIG    BundleHeaderProto_Endianness = 1
)

var BundleHeaderProto_Endianness_name = map[int32]string{
	0: "LITTLE",
	1: "BIG",
}
var BundleHeaderProto_Endianness_value = map[string]int32{
	"LITTLE": 0,
	"BIG":    1,
}

func (x BundleHeaderProto_Endianness) String() string {
	return proto.EnumName(BundleHeaderProto_Endianness_name, int32(x))
}
func (BundleHeaderProto_Endianness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0}
}

// Special header that is associated with a bundle.
//
// TODO(zongheng,zhifengc): maybe in the future, we can add information about
// which binary produced this checkpoint, timestamp, etc. Sometime, these can be
// valuable debugging information. And if needed, these can be used as defensive
// information ensuring reader (binary version) of the checkpoint and the writer
// (binary version) must match within certain range, etc.
type BundleHeaderProto struct {
	// Number of data files in the bundle.
	NumShards  int32                        `protobuf:"varint,1,opt,name=num_shards,json=numShards" json:"num_shards,omitempty"`
	Endianness BundleHeaderProto_Endianness `protobuf:"varint,2,opt,name=endianness,enum=tensorflow.BundleHeaderProto_Endianness" json:"endianness,omitempty"`
	// Versioning of the tensor bundle format.
	Version *tensorflow9.VersionDef `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *BundleHeaderProto) Reset()                    { *m = BundleHeaderProto{} }
func (m *BundleHeaderProto) String() string            { return proto.CompactTextString(m) }
func (*BundleHeaderProto) ProtoMessage()               {}
func (*BundleHeaderProto) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *BundleHeaderProto) GetNumShards() int32 {
	if m != nil {
		return m.NumShards
	}
	return 0
}

func (m *BundleHeaderProto) GetEndianness() BundleHeaderProto_Endianness {
	if m != nil {
		return m.Endianness
	}
	return BundleHeaderProto_LITTLE
}

func (m *BundleHeaderProto) GetVersion() *tensorflow9.VersionDef {
	if m != nil {
		return m.Version
	}
	return nil
}

// Describes the metadata related to a checkpointed tensor.
type BundleEntryProto struct {
	// The tensor dtype and shape.
	Dtype tensorflow3.DataType          `protobuf:"varint,1,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *tensorflow2.TensorShapeProto `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
	// The binary content of the tensor lies in:
	//   File "shard_id": bytes [offset, offset + size).
	ShardId int32 `protobuf:"varint,3,opt,name=shard_id,json=shardId" json:"shard_id,omitempty"`
	Offset  int64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	Size    int64 `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	// The CRC32C checksum of the tensor bytes.
	Crc32C uint32 `protobuf:"fixed32,6,opt,name=crc32c" json:"crc32c,omitempty"`
	// Iff present, this entry represents a partitioned tensor.  The previous
	// fields are interpreted as follows:
	//
	//   "dtype", "shape": describe the full tensor.
	//   "shard_id", "offset", "size", "crc32c": all IGNORED.
	//      These information for each slice can be looked up in their own
	//      BundleEntryProto, keyed by each "slice_name".
	Slices []*tensorflow11.TensorSliceProto `protobuf:"bytes,7,rep,name=slices" json:"slices,omitempty"`
}

func (m *BundleEntryProto) Reset()                    { *m = BundleEntryProto{} }
func (m *BundleEntryProto) String() string            { return proto.CompactTextString(m) }
func (*BundleEntryProto) ProtoMessage()               {}
func (*BundleEntryProto) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *BundleEntryProto) GetDtype() tensorflow3.DataType {
	if m != nil {
		return m.Dtype
	}
	return tensorflow3.DataType_DT_INVALID
}

func (m *BundleEntryProto) GetShape() *tensorflow2.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *BundleEntryProto) GetShardId() int32 {
	if m != nil {
		return m.ShardId
	}
	return 0
}

func (m *BundleEntryProto) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *BundleEntryProto) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BundleEntryProto) GetCrc32C() uint32 {
	if m != nil {
		return m.Crc32C
	}
	return 0
}

func (m *BundleEntryProto) GetSlices() []*tensorflow11.TensorSliceProto {
	if m != nil {
		return m.Slices
	}
	return nil
}

func init() {
	proto.RegisterType((*BundleHeaderProto)(nil), "tensorflow.BundleHeaderProto")
	proto.RegisterType((*BundleEntryProto)(nil), "tensorflow.BundleEntryProto")
	proto.RegisterEnum("tensorflow.BundleHeaderProto_Endianness", BundleHeaderProto_Endianness_name, BundleHeaderProto_Endianness_value)
}

func init() { proto.RegisterFile("tensorflow/core/protobuf/tensor_bundle.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0xaa, 0xd3, 0x40,
	0x10, 0x86, 0xdd, 0x93, 0x93, 0x44, 0xe7, 0xc0, 0xa1, 0xae, 0x72, 0x88, 0xa2, 0x10, 0x0b, 0x42,
	0x10, 0x49, 0x25, 0xf5, 0x09, 0x42, 0x8b, 0x2d, 0xf4, 0xa2, 0x6c, 0x83, 0xb7, 0x25, 0x4d, 0x36,
	0x1a, 0x4c, 0x77, 0xc3, 0x6e, 0x62, 0xa9, 0x2f, 0xe0, 0xf3, 0xf9, 0x36, 0x5e, 0x4a, 0x66, 0x53,
	0x1b, 0x28, 0x15, 0xef, 0x76, 0x66, 0xbe, 0x7f, 0x76, 0xfe, 0xd9, 0x85, 0xf7, 0x0d, 0x17, 0x5a,
	0xaa, 0xa2, 0x92, 0x87, 0x49, 0x26, 0x15, 0x9f, 0xd4, 0x4a, 0x36, 0x72, 0xd7, 0x16, 0x13, 0x53,
	0xd8, 0xee, 0x5a, 0x91, 0x57, 0x3c, 0xc4, 0x34, 0x85, 0x33, 0xfd, 0xf2, 0x42, 0x59, 0xa8, 0x74,
	0xcf, 0x0f, 0x52, 0x7d, 0x3b, 0x49, 0xf5, 0xd7, 0xb4, 0xee, 0x95, 0xff, 0x43, 0x57, 0x65, 0x76,
	0xa2, 0xdf, 0xfe, 0x83, 0x3e, 0xd6, 0x5c, 0xf7, 0x58, 0x70, 0x1d, 0xfb, 0xce, 0x95, 0x2e, 0xa5,
	0xe8, 0xc9, 0xf1, 0x2f, 0x02, 0x4f, 0x63, 0x74, 0xb2, 0xe0, 0x69, 0xce, 0xd5, 0x1a, 0xed, 0xbc,
	0x06, 0x10, 0xed, 0xbe, 0x9b, 0x53, 0xe5, 0xda, 0x23, 0x3e, 0x09, 0x6c, 0xf6, 0x44, 0xb4, 0xfb,
	0x0d, 0x26, 0xe8, 0x02, 0x80, 0x8b, 0xbc, 0x4c, 0x85, 0xe0, 0x5a, 0x7b, 0x37, 0x3e, 0x09, 0xee,
	0xa3, 0x20, 0x3c, 0xdf, 0x19, 0x5e, 0x74, 0x0c, 0xe7, 0x7f, 0x79, 0x36, 0xd0, 0xd2, 0x0f, 0xe0,
	0xf6, 0x03, 0x79, 0x96, 0x4f, 0x82, 0xbb, 0xe8, 0x61, 0xd8, 0xe6, 0xb3, 0x29, 0xcd, 0x78, 0xc1,
	0x4e, 0xd8, 0xf8, 0x0d, 0xc0, 0xb9, 0x17, 0x05, 0x70, 0x56, 0xcb, 0x24, 0x59, 0xcd, 0x47, 0x8f,
	0xa8, 0x0b, 0x56, 0xbc, 0xfc, 0x34, 0x22, 0xe3, 0x9f, 0x37, 0x30, 0x32, 0x13, 0xcc, 0x45, 0xa3,
	0x8e, 0xc6, 0xd2, 0x3b, 0xb0, 0xf3, 0x6e, 0x45, 0xe8, 0xe6, 0x3e, 0x7a, 0x3e, 0xbc, 0x67, 0x96,
	0x36, 0x69, 0x72, 0xac, 0x39, 0x33, 0x08, 0x8d, 0xc0, 0xc6, 0x27, 0x42, 0x6b, 0x77, 0xd1, 0xab,
	0x21, 0x9b, 0xe0, 0x71, 0xd3, 0x95, 0xb1, 0x31, 0x33, 0x28, 0x7d, 0x01, 0x8f, 0x71, 0x5d, 0xdb,
	0x32, 0x47, 0x2b, 0x36, 0x73, 0x31, 0x5e, 0xe6, 0xf4, 0x01, 0x1c, 0x59, 0x14, 0x9a, 0x37, 0xde,
	0xad, 0x4f, 0x02, 0x8b, 0xf5, 0x11, 0xa5, 0x70, 0xab, 0xcb, 0x1f, 0xdc, 0xb3, 0x31, 0x8b, 0xe7,
	0x8e, 0xcd, 0x54, 0x36, 0x8d, 0x32, 0xcf, 0xf1, 0x49, 0xe0, 0xb2, 0x3e, 0xa2, 0x1f, 0xc1, 0xc1,
	0x7f, 0xa0, 0x3d, 0xd7, 0xb7, 0xae, 0xcc, 0xd4, 0xd5, 0xcd, 0x4c, 0x3d, 0x1b, 0x87, 0xf0, 0x4c,
	0xaa, 0x2f, 0x43, 0xb4, 0x6d, 0xca, 0x2a, 0xa6, 0x46, 0x60, 0x76, 0x84, 0x0a, 0xbd, 0x26, 0xbf,
	0x09, 0xd9, 0x39, 0xf8, 0x29, 0xa6, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x75, 0xfa, 0x28,
	0xfd, 0x02, 0x00, 0x00,
}
