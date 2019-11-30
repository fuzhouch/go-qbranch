package qbranch

import (
	"fmt"
	"io"
)

var (
	compactBinaryProtocolSpec ProtocolSpec = ProtocolSpec{
		Endian: LITTLE_ENDIAN,
	}
)

// CompactBinaryWriter implements writer of v1 compact binary.
type CompactBinaryWriter struct {
	version uint
	buf     io.Writer
	spec    ProtocolSpec
}

func NewCompactBinaryWriterV1(buf io.Writer) *CompactBinaryWriter {
	return &CompactBinaryWriter{
		version: 1, // Right now we support only v1
		buf:     buf,
		spec:    compactBinaryProtocolSpec,
	}
}

// Spec returns behavior descriptions of protocol
func (w *CompactBinaryWriter) Spec() ProtocolSpec {
	return w.spec
}

func (w *CompactBinaryWriter) WriteBool(i uint16, v bool) int {
	return 0
}

func (w *CompactBinaryWriter) WriteInt8(i uint16, v int8) int {
	return 0
}

func (w *CompactBinaryWriter) WriteUInt8(i uint16, v uint8) int {
	return 0
}

func (w *CompactBinaryWriter) WriteInt16(i uint16, v int16) int {
	return 0
}

func (w *CompactBinaryWriter) WriteUInt16(i uint16, v uint16) int {
	return 0
}

func (w *CompactBinaryWriter) WriteInt32(i uint16, v int32) int {
	return 0
}

func (w *CompactBinaryWriter) WriteUInt32(i uint16, v uint32) int {
	return 0
}

func (w *CompactBinaryWriter) WriteInt64(i uint16, v int64) int {
	return 0
}

func (w *CompactBinaryWriter) WriteUInt64(i uint16, v uint64) int {
	return 0
}

func (w *CompactBinaryWriter) WriteString(i uint16, v string) int {
	return 0
}

func (w *CompactBinaryWriter) WriteWString(i uint16, v string) int {
	return 0
}

func (w *CompactBinaryWriter) WriteNull(i uint16) int {
	return 0
}

func writeIndex(id uint, typeTag uint8, buf io.Writer) (int, error) {
	bytes := [4]byte{0, 0, 0, 0}
	if id <= 5 {
		bytes[0] = byte(id<<5) | typeTag
		return buf.Write(bytes[:1])
	} else if id <= 0xFF {
		bytes[0] = byte(id)
		bytes[1] = byte(0xC0) | typeTag
		return buf.Write(bytes[0:2])
	} else if id <= 0xFFFF {
		bytes[0] = byte(id & 0xF)
		bytes[1] = byte(id & 0xF0 >> 8)
		bytes[2] = byte(id & 0xF00 >> 16)
		bytes[3] = byte(0xE0 | typeTag)
		return buf.Write(bytes[0:])
		return 4, nil
	}
	return 0, fmt.Errorf("InvalidIndex:%v", id)
}
