package incoming

import (
	"github.com/izzxt/simple-emulator/game"
	"github.com/izzxt/simple-emulator/networking"
	"github.com/izzxt/simple-emulator/packet"
)

type IncomingMessage struct {
	gameClient      game.GameClient
	incomingMessage packet.IncomingPacket
}

// func NewIncomingMessage(gc game.GameClient, in packet.IncomingPacket) *IncomingMessage {
// 	return &IncomingMessage{gameClient: gc, incomingMessage: in}
// }

func (m IncomingMessage) RegisterIncomingMessage(packet networking.Packet) {
	packet.Execute(&m.gameClient, m.incomingMessage)
}
