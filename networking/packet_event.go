package networking

type PacketEvent struct {
	Packet Packet
}

func NewPacketEvent(packet Packet) *PacketEvent {
	return &PacketEvent{Packet: packet}
}
