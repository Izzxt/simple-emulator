package networking

import (
	"github.com/izzxt/simple-emulator/game"
	"github.com/izzxt/simple-emulator/packet"
)

type PacketEvent struct {
	gc game.GameClient
	in packet.IncomingPacket
}

func NewPacketEvent(gc game.GameClient, in packet.IncomingPacket) *PacketEvent {
	return &PacketEvent{gc, in}
}
