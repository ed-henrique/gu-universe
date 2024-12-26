// encoding should be: a1-b2-c16-d[len]
//
// Read as:
//
// - Section 'a' with 1 (one) byte for versioning
//
// - Section 'b' with 2 (two) bytes for operation (message type)
//
// - Section 'c' with 16 (sixteen) bytes for message length
//
// - Section 'd' with a variable number of bytes (equal to section 'c' value) for data
package encoding

type Encodeable interface {
	Encode() []byte
}

type Version uint8

const (
	V1 Version = iota
)

type MessageType uint8

const (
	MsgTypeUnidentified MessageType = iota
	MsgTypePlayer
)
