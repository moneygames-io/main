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
	go gameserver.MapUpdater()

	http.HandleFunc("/ws", wsHandler)
	panic(http.ListenAndServe(":10000", nil))
}

func (gs *GameServer) MapUpdater() {
	// TODO Wait to start doing this channel? After the last connction is established?
	for {
		gs.World.Update()
		view := gs.World.render()

		for k := range gs.Users {
			k.Conn.writeJSON(k)
		}
	}
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
	gs.World.SpawnNewPlayer(c.Player)

	gs.Users[c] = c.Player
	go c.Player.collectInput(conn)
}

func validateToken(token string) bool {
	return true
}
