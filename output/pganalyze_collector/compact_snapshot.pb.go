// Code generated by protoc-gen-go.
// source: compact_snapshot.proto
// DO NOT EDIT!

package pganalyze_collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CompactSnapshot struct {
	// Basic information about this snapshot
	SnapshotVersionMajor int32                      `protobuf:"varint,1,opt,name=snapshot_version_major,json=snapshotVersionMajor" json:"snapshot_version_major,omitempty"`
	SnapshotVersionMinor int32                      `protobuf:"varint,2,opt,name=snapshot_version_minor,json=snapshotVersionMinor" json:"snapshot_version_minor,omitempty"`
	CollectorVersion     string                     `protobuf:"bytes,3,opt,name=collector_version,json=collectorVersion" json:"collector_version,omitempty"`
	SnapshotUuid         string                     `protobuf:"bytes,4,opt,name=snapshot_uuid,json=snapshotUuid" json:"snapshot_uuid,omitempty"`
	CollectedAt          *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=collected_at,json=collectedAt" json:"collected_at,omitempty"`
	BaseRefs             *CompactSnapshot_BaseRefs  `protobuf:"bytes,6,opt,name=base_refs,json=baseRefs" json:"base_refs,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*CompactSnapshot_LogSnapshot
	//	*CompactSnapshot_SystemSnapshot
	//	*CompactSnapshot_ActivitySnapshot
	Data isCompactSnapshot_Data `protobuf_oneof:"data"`
}

func (m *CompactSnapshot) Reset()                    { *m = CompactSnapshot{} }
func (m *CompactSnapshot) String() string            { return proto.CompactTextString(m) }
func (*CompactSnapshot) ProtoMessage()               {}
func (*CompactSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type isCompactSnapshot_Data interface {
	isCompactSnapshot_Data()
}

type CompactSnapshot_LogSnapshot struct {
	LogSnapshot *CompactLogSnapshot `protobuf:"bytes,10,opt,name=log_snapshot,json=logSnapshot,oneof"`
}
type CompactSnapshot_SystemSnapshot struct {
	SystemSnapshot *CompactSystemSnapshot `protobuf:"bytes,11,opt,name=system_snapshot,json=systemSnapshot,oneof"`
}
type CompactSnapshot_ActivitySnapshot struct {
	ActivitySnapshot *CompactActivitySnapshot `protobuf:"bytes,12,opt,name=activity_snapshot,json=activitySnapshot,oneof"`
}

func (*CompactSnapshot_LogSnapshot) isCompactSnapshot_Data()      {}
func (*CompactSnapshot_SystemSnapshot) isCompactSnapshot_Data()   {}
func (*CompactSnapshot_ActivitySnapshot) isCompactSnapshot_Data() {}

func (m *CompactSnapshot) GetData() isCompactSnapshot_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CompactSnapshot) GetSnapshotVersionMajor() int32 {
	if m != nil {
		return m.SnapshotVersionMajor
	}
	return 0
}

func (m *CompactSnapshot) GetSnapshotVersionMinor() int32 {
	if m != nil {
		return m.SnapshotVersionMinor
	}
	return 0
}

func (m *CompactSnapshot) GetCollectorVersion() string {
	if m != nil {
		return m.CollectorVersion
	}
	return ""
}

func (m *CompactSnapshot) GetSnapshotUuid() string {
	if m != nil {
		return m.SnapshotUuid
	}
	return ""
}

func (m *CompactSnapshot) GetCollectedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CollectedAt
	}
	return nil
}

func (m *CompactSnapshot) GetBaseRefs() *CompactSnapshot_BaseRefs {
	if m != nil {
		return m.BaseRefs
	}
	return nil
}

func (m *CompactSnapshot) GetLogSnapshot() *CompactLogSnapshot {
	if x, ok := m.GetData().(*CompactSnapshot_LogSnapshot); ok {
		return x.LogSnapshot
	}
	return nil
}

func (m *CompactSnapshot) GetSystemSnapshot() *CompactSystemSnapshot {
	if x, ok := m.GetData().(*CompactSnapshot_SystemSnapshot); ok {
		return x.SystemSnapshot
	}
	return nil
}

func (m *CompactSnapshot) GetActivitySnapshot() *CompactActivitySnapshot {
	if x, ok := m.GetData().(*CompactSnapshot_ActivitySnapshot); ok {
		return x.ActivitySnapshot
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CompactSnapshot) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CompactSnapshot_OneofMarshaler, _CompactSnapshot_OneofUnmarshaler, _CompactSnapshot_OneofSizer, []interface{}{
		(*CompactSnapshot_LogSnapshot)(nil),
		(*CompactSnapshot_SystemSnapshot)(nil),
		(*CompactSnapshot_ActivitySnapshot)(nil),
	}
}

func _CompactSnapshot_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CompactSnapshot)
	// data
	switch x := m.Data.(type) {
	case *CompactSnapshot_LogSnapshot:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogSnapshot); err != nil {
			return err
		}
	case *CompactSnapshot_SystemSnapshot:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SystemSnapshot); err != nil {
			return err
		}
	case *CompactSnapshot_ActivitySnapshot:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ActivitySnapshot); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CompactSnapshot.Data has unexpected type %T", x)
	}
	return nil
}

func _CompactSnapshot_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CompactSnapshot)
	switch tag {
	case 10: // data.log_snapshot
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CompactLogSnapshot)
		err := b.DecodeMessage(msg)
		m.Data = &CompactSnapshot_LogSnapshot{msg}
		return true, err
	case 11: // data.system_snapshot
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CompactSystemSnapshot)
		err := b.DecodeMessage(msg)
		m.Data = &CompactSnapshot_SystemSnapshot{msg}
		return true, err
	case 12: // data.activity_snapshot
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CompactActivitySnapshot)
		err := b.DecodeMessage(msg)
		m.Data = &CompactSnapshot_ActivitySnapshot{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CompactSnapshot_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CompactSnapshot)
	// data
	switch x := m.Data.(type) {
	case *CompactSnapshot_LogSnapshot:
		s := proto.Size(x.LogSnapshot)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CompactSnapshot_SystemSnapshot:
		s := proto.Size(x.SystemSnapshot)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CompactSnapshot_ActivitySnapshot:
		s := proto.Size(x.ActivitySnapshot)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CompactSnapshot_BaseRefs struct {
	RoleReferences     []*RoleReference     `protobuf:"bytes,1,rep,name=role_references,json=roleReferences" json:"role_references,omitempty"`
	DatabaseReferences []*DatabaseReference `protobuf:"bytes,2,rep,name=database_references,json=databaseReferences" json:"database_references,omitempty"`
	QueryReferences    []*QueryReference    `protobuf:"bytes,3,rep,name=query_references,json=queryReferences" json:"query_references,omitempty"`
	QueryInformations  []*QueryInformation  `protobuf:"bytes,4,rep,name=query_informations,json=queryInformations" json:"query_informations,omitempty"`
	RelationReferences []*RelationReference `protobuf:"bytes,5,rep,name=relation_references,json=relationReferences" json:"relation_references,omitempty"`
}

func (m *CompactSnapshot_BaseRefs) Reset()                    { *m = CompactSnapshot_BaseRefs{} }
func (m *CompactSnapshot_BaseRefs) String() string            { return proto.CompactTextString(m) }
func (*CompactSnapshot_BaseRefs) ProtoMessage()               {}
func (*CompactSnapshot_BaseRefs) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

func (m *CompactSnapshot_BaseRefs) GetRoleReferences() []*RoleReference {
	if m != nil {
		return m.RoleReferences
	}
	return nil
}

func (m *CompactSnapshot_BaseRefs) GetDatabaseReferences() []*DatabaseReference {
	if m != nil {
		return m.DatabaseReferences
	}
	return nil
}

func (m *CompactSnapshot_BaseRefs) GetQueryReferences() []*QueryReference {
	if m != nil {
		return m.QueryReferences
	}
	return nil
}

func (m *CompactSnapshot_BaseRefs) GetQueryInformations() []*QueryInformation {
	if m != nil {
		return m.QueryInformations
	}
	return nil
}

func (m *CompactSnapshot_BaseRefs) GetRelationReferences() []*RelationReference {
	if m != nil {
		return m.RelationReferences
	}
	return nil
}

func init() {
	proto.RegisterType((*CompactSnapshot)(nil), "pganalyze.collector.CompactSnapshot")
	proto.RegisterType((*CompactSnapshot_BaseRefs)(nil), "pganalyze.collector.CompactSnapshot.BaseRefs")
}

func init() { proto.RegisterFile("compact_snapshot.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xc9, 0xba, 0x56, 0x9d, 0x13, 0xd6, 0xd6, 0x43, 0x93, 0x15, 0x09, 0xad, 0xda, 0x04,
	0x54, 0xfc, 0xc9, 0xa4, 0xc1, 0x2d, 0x17, 0x1b, 0x5c, 0xf0, 0x67, 0x20, 0x61, 0x36, 0xb8, 0xe0,
	0x22, 0x72, 0x13, 0x37, 0x0b, 0x72, 0xe2, 0xd4, 0x76, 0x26, 0x95, 0x07, 0xe1, 0x35, 0x78, 0x45,
	0x54, 0x27, 0x4e, 0x33, 0x13, 0x7a, 0xd7, 0x73, 0xce, 0xf7, 0xfd, 0x6c, 0x7f, 0xa7, 0x01, 0x87,
	0x11, 0xcf, 0x0a, 0x12, 0xa9, 0x50, 0xe6, 0xa4, 0x90, 0x37, 0x5c, 0x05, 0x85, 0xe0, 0x8a, 0xc3,
	0x83, 0x22, 0x21, 0x39, 0x61, 0xab, 0x5f, 0x34, 0x88, 0x38, 0x63, 0x34, 0x52, 0x5c, 0xf8, 0x47,
	0x09, 0xe7, 0x09, 0xa3, 0xa7, 0x5a, 0x32, 0x2f, 0x17, 0xa7, 0x2a, 0xcd, 0xa8, 0x54, 0x24, 0x2b,
	0x2a, 0x97, 0x7f, 0x64, 0x68, 0x24, 0x52, 0xe9, 0x6d, 0xaa, 0x56, 0x16, 0xd6, 0xf7, 0x8d, 0x80,
	0xf1, 0xc4, 0x9e, 0x3d, 0x6c, 0xae, 0xb2, 0x92, 0x8a, 0x66, 0xf6, 0xd8, 0x93, 0x37, 0x44, 0xd0,
	0xb8, 0xaa, 0x8e, 0x7f, 0x0f, 0xc1, 0xe8, 0x4d, 0xa5, 0xff, 0x5a, 0xeb, 0xe0, 0x2b, 0x70, 0x68,
	0x3c, 0xe1, 0x2d, 0x15, 0x32, 0xe5, 0x79, 0x98, 0x91, 0x9f, 0x5c, 0x20, 0x67, 0xea, 0xcc, 0xfa,
	0xf8, 0x81, 0x99, 0x7e, 0xab, 0x86, 0x9f, 0xd6, 0xb3, 0x6e, 0x57, 0x9a, 0x73, 0x81, 0x76, 0xba,
	0x5d, 0xeb, 0x19, 0x7c, 0x06, 0x26, 0x4d, 0x2e, 0xc6, 0x86, 0x7a, 0x53, 0x67, 0xb6, 0x87, 0xc7,
	0xcd, 0xa0, 0x76, 0xc0, 0x13, 0x70, 0xbf, 0x39, 0xa2, 0x2c, 0xd3, 0x18, 0xed, 0x6a, 0xa1, 0x67,
	0x9a, 0xd7, 0x65, 0x1a, 0xc3, 0xd7, 0xc0, 0xab, 0x8d, 0x34, 0x0e, 0x89, 0x42, 0xfd, 0xa9, 0x33,
	0x73, 0xcf, 0xfc, 0xa0, 0xca, 0x3c, 0x30, 0x99, 0x07, 0x57, 0x26, 0x73, 0xec, 0x36, 0xfa, 0x73,
	0x05, 0x3f, 0x80, 0xbd, 0x39, 0x91, 0x34, 0x14, 0x74, 0x21, 0xd1, 0x40, 0x7b, 0x5f, 0x04, 0x1d,
	0x4b, 0x0c, 0xac, 0xd4, 0x82, 0x0b, 0x22, 0x29, 0xa6, 0x0b, 0x89, 0x87, 0xf3, 0xfa, 0x17, 0xbc,
	0x04, 0x5e, 0x7b, 0x3f, 0x08, 0x68, 0xdc, 0x93, 0x6d, 0xb8, 0x4b, 0x9e, 0x18, 0xe2, 0xbb, 0x7b,
	0xd8, 0x65, 0x9b, 0x12, 0x5e, 0x83, 0x91, 0xb5, 0x51, 0xe4, 0x6a, 0xe0, 0xd3, 0xad, 0xf7, 0xd3,
	0x96, 0x16, 0x73, 0x5f, 0xde, 0xe9, 0xc0, 0x1f, 0x60, 0xf2, 0xcf, 0xbf, 0x0c, 0x79, 0x1a, 0xfc,
	0x7c, 0x1b, 0xf8, 0xbc, 0x36, 0xb5, 0xd0, 0x63, 0x62, 0xf5, 0xfc, 0x3f, 0x3d, 0x30, 0x34, 0xc1,
	0xc0, 0x8f, 0x60, 0x24, 0x38, 0xd3, 0xd1, 0x52, 0x41, 0xf3, 0x88, 0x4a, 0xe4, 0x4c, 0x7b, 0x33,
	0xf7, 0xec, 0xb8, 0xf3, 0x1c, 0xcc, 0xd9, 0xda, 0x57, 0x49, 0xf1, 0xbe, 0x68, 0x97, 0x12, 0x7e,
	0x07, 0x07, 0x31, 0x51, 0xc4, 0xec, 0xca, 0x00, 0x77, 0x34, 0xf0, 0x71, 0x27, 0xf0, 0x6d, 0xad,
	0xdf, 0x40, 0x61, 0x6c, 0xb7, 0x24, 0xfc, 0x0c, 0xc6, 0xcb, 0x92, 0x8a, 0x55, 0x9b, 0xda, 0xd3,
	0xd4, 0x93, 0x4e, 0xea, 0x97, 0xb5, 0x78, 0x83, 0x1c, 0x2d, 0xef, 0xd4, 0x12, 0x5e, 0x01, 0x58,
	0xf1, 0xd2, 0x7c, 0xc1, 0x45, 0x46, 0x54, 0xca, 0x73, 0x89, 0x76, 0x35, 0xf1, 0xd1, 0xff, 0x89,
	0xef, 0x37, 0x6a, 0x3c, 0x59, 0x5a, 0x1d, 0xfd, 0x7c, 0x41, 0x99, 0x2e, 0xda, 0x17, 0xed, 0x6f,
	0x79, 0x3e, 0xae, 0xf5, 0xad, 0xe7, 0x0b, 0xbb, 0x25, 0x2f, 0x06, 0x60, 0x57, 0x87, 0x32, 0xd0,
	0x1f, 0xca, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0xfd, 0x62, 0xd1, 0xd9, 0x04, 0x00,
	0x00,
}