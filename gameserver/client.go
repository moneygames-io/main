package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Name             string
	CurrentZoomLevel int
	Token            string
	Conn             *websocket.Conn
	Player           *Player
}

func NewClient(r *RegisterMessage, conn *websocket.Conn) *Client {
	c := &Client{}

	c.Name = r.Name
	c.Token = r.Token
	c.Conn = conn

	return c
}

func (c *Client) GetView(m *Map) [][]uint32 {
	head := c.Player.Snake.Head
	r0 := head.Row - c.CurrentZoomLevel
	c0 := head.Col - c.CurrentDirection

	colors := make([][]uint32, len(c.CurrentZoomLevel*2))

	for i := range m.Tiles {
		colors[i] = make([]uint32, len(c.CurrentZoomLevel*2))
	}

	for r := r0; r < c.CurrentZoomLevel*2; r++ {
		for c := c0; c < c.CurrentZoomLevel*2; c++ {
			colors[r][c] = m.GetColor(&m.Tiles[r][c])
		}
	}

	return colors
}

func (c *Client) CollectInput(conn *websocket.Conn) {
	msg := &ClientUpdateMessage{}
	for {
		conn.ReadJSON(msg)
		c.Player.CurrentDirection = msg.CurrentDirection
		c.Player.CurrentSprint = msg.CurrentSprint

		c.CurrentZoomLevel = msg.CurrentZoomLevel
	}
}
