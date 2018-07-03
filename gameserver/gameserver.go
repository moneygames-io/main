package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type GameServer struct {
	Users map[*Client]*Player
	World *Map
}

var gameserver *GameServer

func getID() string {
	return ":" + os.Getenv("GSPORT")
}

func main() {
	gameserver = &GameServer{make(map[*Client]*Player), NewMap(2)}
	go gameserver.MapUpdater()

	http.HandleFunc("/ws", wsHandler)
	fmt.Println("ID", getID())
	panic(http.ListenAndServe(":10000", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	gameserver.PlayerJoined(conn) // TODO Should / can this be asyncronous
}

func (gs *GameServer) PlayerJoined(conn *websocket.Conn) {
	fmt.Println("Player Joined")
	message := &RegisterMessage{}

	error := conn.ReadJSON(message)

	if error != nil || !validateToken(message.Token) {
		conn.Close()
	}

	// TODO token consumed

	c := NewClient(message, conn)
	c.Player = &Player{}
	gs.World.SpawnNewPlayer(c.Player)

	gs.Users[c] = c.Player
	go c.collectInput(conn)
}

func validateToken(token string) bool {
	return true
}

func (gs *GameServer) MapUpdater() {
	// TODO Wait to start doing this channel? After the last connction is established?
	for {
		if len(gs.Users) > 1 {
			// TODO somewhere we are going to a nil snake location and calling move on that
			gs.World.Update()
			view := gs.World.Render()

			for k := range gs.Users {
				// TODO this is too large?
				k.Conn.WriteJSON(&view)
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}
