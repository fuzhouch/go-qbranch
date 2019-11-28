package qbranch

import (
	"bytes"
	"fmt"
)

var (
	compactBinaryProtocolSpec ProtocolSpec = ProtocolSpec{
		Endian: LITTLE_ENDIAN,
	}
)

// CompactBinaryWriterV1 implements writer of v1 compact binary.
type CompactBinaryWriterV1 struct {
	buf  *bytes.Buffer
	spec ProtocolSpec
}

func NewCompactBinaryWriterV1(buf *bytes.Buffer) *CompactBinaryWriterV1 {
	return &CompactBinaryWriterV1{
		buf:  buf,
		spec: compactBinaryProtocolSpec,
	}
}

// Spec returns behavior descriptions of protocol
func (w *CompactBinaryWriterV1) Spec() ProtocolSpec {
	return w.spec
}

func (w *CompactBinaryWriterV1) WriteBool(index uint16, value bool) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteInt8(index uint16, value int8) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteUInt8(index uint16, value uint8) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteInt16(index uint16, value int16) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteUInt16(index uint16, value uint16) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteInt32(index uint16, value int32) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteUInt32(index uint16, value uint32) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteInt64(index uint16, value int64) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteUInt64(index uint16, value uint64) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteString(index uint16, value string) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteWString(index uint16, value string) (length uint) {
	length = 0
	return
}

func (w *CompactBinaryWriterV1) WriteNull(index uint16) (length uint) {
	length = 0
	return
}

func writeIndex(id uint, typeTag uint8, buf *bytes.Buffer) (length uint, err error) {
	if id <= 5 {
		buf.WriteByte(byte(id<<5) | typeTag)
		length = 1
		return
	} else if id <= 0xFF {
		buf.WriteByte(byte(id))
		buf.WriteByte(byte(0xC0) | typeTag)
		length = 2
		return
	} else if id <= 0xFFFF {
		buf.WriteByte(byte(id & 0xF))
		buf.WriteByte(byte(id & 0xF0 >> 8))
		buf.WriteByte(byte(id & 0xF00 >> 16))
		buf.WriteByte(byte(0xE0 | typeTag))
		length = 4
		return
	}
	length = 0
	err = fmt.Errorf("InvalidIndex:%v", id)
	return
}
