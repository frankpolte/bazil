// Code generated by protoc-gen-gogo.
// source: bazil.org/bazil/cas/wire/manifest.proto
// DO NOT EDIT!

/*
	Package wire is a generated protocol buffer package.

	It is generated from these files:
		bazil.org/bazil/cas/wire/manifest.proto

	It has these top-level messages:
		Manifest
*/
package wire

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import bazil_org_bazil_cas "bazil.org/bazil/cas"

import io "io"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Manifest struct {
	Root             bazil_org_bazil_cas.Key `protobuf:"bytes,1,req,name=root,customtype=bazil.org/bazil/cas.Key" json:"root"`
	Size             uint64                  `protobuf:"varint,2,req,name=size" json:"size"`
	ChunkSize        uint32                  `protobuf:"varint,3,req,name=chunkSize" json:"chunkSize"`
	Fanout           uint32                  `protobuf:"varint,4,req,name=fanout" json:"fanout"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Manifest) Reset()         { *m = Manifest{} }
func (m *Manifest) String() string { return proto.CompactTextString(m) }
func (*Manifest) ProtoMessage()    {}

func (m *Manifest) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Manifest) GetChunkSize() uint32 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

func (m *Manifest) GetFanout() uint32 {
	if m != nil {
		return m.Fanout
	}
	return 0
}

func init() {
}
func (m *Manifest) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Root.Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 2:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Size |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.ChunkSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Fanout |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}