package packet

import (
	"bytes"
	"encoding/binary"
)

type OutgoingPacket interface {
	Writebyte(value byte)
	WriteShort(value int16)
	WriteInt(value int32)
	WriteBool(value bool)
	WriteLong(value int64)
	WriteString(value string)
	WriteDouble(value []byte)
	GetHeader() uint16
	GetBytes() []byte
}

type outgoingPacket struct {
	header uint16
	bytes  bytes.Buffer
}

// GetHeader implements OutgoingPacket.
func (o *outgoingPacket) GetHeader() uint16 {
	return o.header
}

// GetBytes implements OutgoingPacket.
func (o *outgoingPacket) GetBytes() []byte {
	return o.bytes.Bytes()
}

// WriteLong implements OutgoingPacket.
func (o *outgoingPacket) WriteLong(value int64) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// WriteBool implements OutgoingPacket.
func (o *outgoingPacket) WriteBool(value bool) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// Writebyte implements OutgoingPacket.
func (o *outgoingPacket) Writebyte(value byte) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// WriteDouble implements OutgoingPacket.
func (*outgoingPacket) WriteDouble(value []byte) {
	panic("unimplemented")
}

// WriteInt implements OutgoingPacket.
func (o *outgoingPacket) WriteInt(value int32) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// WriteShort implements OutgoingPacket.
func (o *outgoingPacket) WriteShort(value int16) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// WriteString implements OutgoingPacket.
func (o *outgoingPacket) WriteString(value string) {
	if value == "" {
		o.WriteString("")
		return
	}
	o.WriteShort(int16(len(value)))
	binary.Write(&o.bytes, binary.BigEndian, []byte(value))
}

func NewOutgoingPacket(header uint16, b []byte) OutgoingPacket {
	return &outgoingPacket{
		header: header,
		bytes:  *bytes.NewBuffer(b),
	}
}
