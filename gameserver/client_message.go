package main

type ClientMessage struct {
	Name             string
	Angle			 int
	CurrentZoomLevel int
	CurrentSprint    bool
	Token		     string
	Conn			 *websocket.Conn
}
