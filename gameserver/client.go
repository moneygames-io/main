package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Name             string
	CurrentZoomLevel int
	Token		     string
	Conn			 *websocket.Conn
	Player			 *Player
}

func NewClient(r *RegisterMessage, conn *websocket.Conn) *Client {
	c := &Client{}

	c.Name = r.Name
	c.Token = r.Token
	c.Conn = conn

	return c
}
