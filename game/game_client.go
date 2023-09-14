package game

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: websocket.IsWebSocketUpgrade,
}

type GameClient struct {
	ctx     context.Context
	mutex   sync.RWMutex
	Habbos  map[*Habbo]bool
	Clients map[*websocket.Conn]bool
}

// AddClient implements IGameClient.
func (gc *GameClient) AddClient(habbo *Habbo) {
	gc.mutex.Lock()
	defer gc.mutex.Unlock()

	gc.Habbos[habbo] = true
}

// AddHabbo implements IGameClient.
func (*GameClient) AddHabbo(ssoTicket string) {
	panic("unimplemented")
}

// RemoveClient implements IGameClient.
func (gc *GameClient) RemoveClient(habbo *Habbo) {
	gc.mutex.Lock()
	defer gc.mutex.Unlock()

	if _, ok := gc.Habbos[habbo]; ok {
		println("habbo disconnected", habbo.connection.RemoteAddr().String())
		habbo.connection.Close()
		delete(gc.Clients, &habbo.connection)
		return
	}
}

// RemoveHabbo implements IGameClient.
func (*GameClient) RemoveHabbo() string {
	panic("unimplemented")
}

type IGameClient interface {
	AddHabbo(ssoTicket string)
	RemoveHabbo() string
	AddClient(habbo *Habbo)
	RemoveClient(habbo *Habbo)
	ServeWS(w http.ResponseWriter, r *http.Request)
}

func (gc *GameClient) ServeWS(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}

	habbo := NewHabbo(context.Background(), *c, gc)
	gc.AddClient(habbo)

	go habbo.ReadMessage()
	go habbo.WriteMessage()
}

func NewGameClient(ctx context.Context) IGameClient {
	return &GameClient{ctx: ctx, Habbos: make(map[*Habbo]bool), Clients: make(map[*websocket.Conn]bool)}
}
