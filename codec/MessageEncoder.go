package codec

import "encoding/binary"

func Encode(header uint16, bytes []byte) []byte {
	binary.BigEndian.PutUint32(bytes, uint32(len(bytes)-4))
	binary.BigEndian.PutUint16(bytes[4:], header)
	return bytes
}
