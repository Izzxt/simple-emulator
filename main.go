package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/izzxt/simple-emulator/codec"
	"github.com/izzxt/simple-emulator/game"
	"github.com/izzxt/simple-emulator/networking/incoming"
	"github.com/izzxt/simple-emulator/networking/outgoing"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: websocket.IsWebSocketUpgrade,
}

var (
	clients = make(map[*websocket.Conn]game.IGameClient)
)

func main() {
	http.HandleFunc("/", ws)
	log.Fatal(http.ListenAndServe(":2096", nil))
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}
	defer c.Close()

	receivedFile := writeIntoHexFile("received.hex")
	defer receivedFile.Close()
	sendFile := writeIntoHexFile("send.hex")
	defer sendFile.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		data, _, header := codec.Decode(message)

		incomingPacket := incoming.NewIncomingPacket(header, data)
		// fmt.Println("data :", incomingPacket.ReadString())

		client := game.NewGameClient()
		if header == 2419 {
			_, ok := clients[c]
			if !ok {
				clients[c] = client
				client.SetAuthTicket(incomingPacket.ReadString())
			}
		}
		fmt.Println("AuthTicket :", client.GetAuthTicket())

		if header == 357 {
			break
		}

		_, err = receivedFile.Write(message) // write to file received.hex
		if err != nil {
			log.Fatal(err)
		}

		bytes := make([]byte, 6)
		outgoingPacket := outgoing.NewOutgoingPacket(2491, bytes)
		outgoingPacket.WriteString("Hello World")
		b := codec.Encode(outgoingPacket.GetHeader(), outgoingPacket.GetBytes())

		err = c.WriteMessage(websocket.BinaryMessage, outgoingPacket.GetBytes())
		if err != nil {
			log.Println("write:", err)
			break
		}

		_, err = sendFile.Write(b) // write to file send.hex
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeIntoHexFile(filename string) *os.File {
	file, err := os.OpenFile(
		filename,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
