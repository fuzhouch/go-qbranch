package qbranch

type SerializationEndian byte

const (
	_ SerializationEndian = iota
	LITTLE_ENDIAN
	BIG_ENDIAN
)

type ProtocolSpec struct {
	Endian SerializationEndian
}
