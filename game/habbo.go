package game

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/izzxt/simple-emulator/codec"
	incoming "github.com/izzxt/simple-emulator/messages/incoming"
	outgoing "github.com/izzxt/simple-emulator/messages/outgoing"
)

type Habbo struct {
	ctx        context.Context
	connection websocket.Conn
	gameClient IGameClient
	incoming   chan incoming.IncomingPacket
}

func (h *Habbo) ReadMessage() {
	for {
		_, message, err := h.connection.ReadMessage()
		if err != nil {
			h.connection.Close()
			h.gameClient.RemoveClient(h)
			break
		}

		data, _, header := codec.Decode(message)
		incomingPacket := incoming.NewIncomingPacket(header, data)
		h.incoming <- incomingPacket
	}
}

// WriteMessage implements IGameClient.
func (h *Habbo) WriteMessage() {
	for {
		select {
		case incoming, ok := <-h.incoming:
			if !ok {
				panic("not ok")
			}

			if incoming.GetHeader() == 4000 {
				println("ReleaseVersionEvent :", incoming.ReadString())
			}

			if incoming.GetHeader() == 2419 {
				println("SecureLoginEvent :", incoming.ReadString())

				bytes := make([]byte, 6)
				outgoingPacket := outgoing.NewOutgoingPacket(2491, bytes)
				_ = codec.Encode(outgoingPacket.GetHeader(), outgoingPacket.GetBytes())
				err := h.connection.WriteMessage(websocket.BinaryMessage, outgoingPacket.GetBytes())
				if err != nil {
					h.gameClient.RemoveClient(h)
					break
				}
			}
		}
	}
}

func NewHabbo(ctx context.Context, conn websocket.Conn, gameClient IGameClient) *Habbo {
	return &Habbo{ctx: ctx, connection: conn, gameClient: gameClient, incoming: make(chan incoming.IncomingPacket)}
}
