package main

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	CurrentDirection int
	CurrentSprint    bool
	Snake            *Snake
	Client           *Client
}

//TODO moved to Client YES
func (p *Player) collectInput(conn *websocket.Conn) {
	msg := &ClientUpdateMessage{}
	for {
		conn.ReadJSON(msg)
		p.CurrentDirection = msg.CurrentDirection
		p.CurrentSprint = msg.CurrentSprint
		p.Client.CurrentZoomLevel = msg.CurrentZoomLevel
	}
}
