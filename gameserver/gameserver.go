package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type GameServer struct {
    Users map[*Client]*Player
	World *Map
}

var gameserver *GameServer

func main() {
    gameserver = &GameServer{make(map[*Client]*Player), NewMap(2)}

	http.HandleFunc("/ws", wsHandler)
	panic(http.ListenAndServe(":10000", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	gameserver.PlayerJoined(conn) // Should / can this be asyncronous
}

func (gs *GameServer) PlayerJoined(conn *websocket.Conn) {
	var message &RegisterMessage{}

	error := conn.readJSON(message)

	if error != nil || !validateToken(message.Token) {
		conn.Close()
	}

	// TODO token consumed

	c := NewClient(message, conn)
	c.Player = &Player{}
	gameserver.World.SpawnNewPlayer(c.Player)

	gameserver.Users[c] = c.Player
}

func validateToken(token string) bool {
	return true
}
