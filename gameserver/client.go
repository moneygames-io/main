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

func (c *Client) collectInput(conn *websocket.Conn) {
	msg := &ClientUpdateMessage{}
	for {
		conn.ReadJSON(msg)
		c.Player.CurrentDirection = msg.CurrentDirection
		c.Player.CurrentSprint = msg.CurrentSprint

		c.CurrentZoomLevel = msg.CurrentZoomLevel
	}
}
