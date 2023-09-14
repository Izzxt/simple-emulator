package networking

import (
	"github.com/izzxt/simple-emulator/game"
	incoming "github.com/izzxt/simple-emulator/messages/incoming"
)

type Packet interface {
	Execute(gameClient game.IGameClient, incomingPacket incoming.IncomingPacket) Packet
}
