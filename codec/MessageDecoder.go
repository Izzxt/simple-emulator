package codec

import (
	"bytes"
	"encoding/binary"
)

func Decode(b []byte) ([]byte, int32, int16) {
	buf := bytes.NewBuffer(b)

	if buf.Len() < 6 {
	}

	var length int32
	binary.Read(buf, binary.BigEndian, &length)
	length = length - 2

	if length+6 > int32(buf.Len()) {
	}

	buf = bytes.NewBuffer(b[4:])

	var messageId int16
	binary.Read(buf, binary.BigEndian, &messageId)

	return buf.Bytes(), length, messageId
}
