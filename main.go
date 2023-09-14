package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/izzxt/simple-emulator/game"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: websocket.IsWebSocketUpgrade,
}

func main() {
	gameClient := game.NewGameClient(context.Background())

	http.HandleFunc("/", gameClient.ServeWS)
	log.Fatal(http.ListenAndServe(":2096", nil))
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
