package networking

import (
	"github.com/izzxt/simple-emulator/game"
	"github.com/izzxt/simple-emulator/packet"
)

type Packet interface {
	Execute(gameClient game.IGameClient, incomingPacket packet.IncomingPacket) Packet
}
