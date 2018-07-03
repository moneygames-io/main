package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
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

	// TODO other cool routes, spectator mode, analytics mode
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

// TODO Need a better game start detection and game end detection.
func (gs *GameServer) MapUpdater() {
	gameStarted := false
	// TODO Many ways to do this without polling
	for {
		if len(gs.Users) > 1 {
			gameStarted = true
			gs.World.Update()
			view := gs.World.Render()

			for k := range gs.Users {
				// TODO this is too large?
				// TODO Should this be async? Is it even blocking?
				k.Conn.WriteJSON(&view)
			}
		}

		if len(gs.World.Players) == 1 && gameStarted {
			// TODO Cleanup
			os.Exit(0)
		}

		// TODO Should be based on how much time has elapsed
		time.Sleep(500 * time.Millisecond)
	}
}
