package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type GameServer struct {
    Users map[*Client]*Player
}

var gameserver *GameServer

func main() {
    gameserver = &GameServer{make(map[*Client]*Player)}

	http.HandleFunc("/ws", wsHandler)

	panic(http.ListenAndServe(":10000", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	gameserver.PlayerJoined(conn)
}

func (gs *GameServer) PlayerJoined(conn *websocket.Conn) {
	// https://blog.golang.org/json-and-go
	// Type switches exactly what you need
}
