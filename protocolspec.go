package qbranch

type SerializationEndian byte

const (
	_ SerializationEndian = iota
	LITTLE_ENDIAN
	BIG_ENDIAN
)

// ProtocolSpec is used to describe behavior of a given protocol like
// endian. This infomation is introduced to clarify behavior of
// protocols, which does not exist in official Bond CompactBinaryWriter.
// In Official Bond, it implicilty apply little endian. This can
// sometimes introduce confusions.
type ProtocolSpec struct {
	Endian SerializationEndian
}
