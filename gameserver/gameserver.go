package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/kutase/go-gameloop"
)

type GameServer struct {
	Users       map[*Client]*Player
	World       *Map
	RedisClient *redis.Client
}

var gameserver *GameServer
var gameLoop *GameLoop

func getID() string {
	return os.Getenv("GSPORT")
}

func main() {
	var client *redis.Client
	for {
		client = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})
		_, err := client.Ping().Result()
		if err != nil {
			fmt.Println("Matchmaker could not connect to redis")
			fmt.Println(err)
		} else {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	gameserver = &GameServer{make(map[*Client]*Player), NewMap(2)}

	gameLoop = gameLoop.New(10, MapUpdater)
	go gameserver.MapUpdater()

	// TODO other cool routes, spectator mode, analytics mode
	http.HandleFunc("/ws", wsHandler)

	client.Set(getID(), "waiting for players")
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
func (gs *GameServer) MapUpdater(delta float64) {
	gameStarted := false
	// TODO Many ways to do this without polling
	if len(gs.Users) > 1 {
		gs.RedisClient.Set(getID(), "game started")
		gameStarted = true
		re
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
}
