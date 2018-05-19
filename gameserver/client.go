package main

type Client struct {
	Name             string
	CurrentZoomLevel int
	Token		     string
	Conn			 *websocket.Conn
	Player			 *Player
}

func NewClient(r *RegisterMessage, conn *websocket.Conn) {
	c := &Client{}

	c.Name = r.Name
	c.Token = r.Token
	c.Conn = conn
}
