package qbranch

import (
	"bytes"
	"testing"
)

func TestCompactBinarySpec(t *testing.T) {
	buf := bytes.Buffer{}
	writer := NewCompactBinaryWriterV1(&buf)
	spec := writer.Spec()
	if spec.Endian != LITTLE_ENDIAN {
		t.Errorf("CompactBinaryWriterMustBeLittleEndian")
		return
	}
}
