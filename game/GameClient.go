package game

type GameClient struct {
	AuthTicket string
}

// GetAuthTicket implements IGameClient.
func (c *GameClient) GetAuthTicket() string {
	return c.AuthTicket
}

// SetAuthTicket implements IGameClient.
func (c *GameClient) SetAuthTicket(ssoTicket string) {
	c.AuthTicket = ssoTicket
}

type IGameClient interface {
	SetAuthTicket(ssoTicket string)
	GetAuthTicket() string
}

func NewGameClient() IGameClient {
	return &GameClient{AuthTicket: ""}
}
