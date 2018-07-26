package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/parth/go-gameloop"
)

type GameServer struct {
	Users       map[*Client]*Player
	World       *Map
	RedisClient *redis.Client
	ID          string
	GL          *gameLoop.GameLoop
}

var gameserver *GameServer

func main() {
	gameserver = &GameServer{
		Users:       make(map[*Client]*Player),
		World:       NewMap(2),
		RedisClient: connectToRedis(),
		ID:          os.Getenv("GSPORT"),
	}
	gameserver.GL = gameLoop.New(2, gameserver.MapUpdater)

	http.HandleFunc("/ws", wsHandler)
	panic(http.ListenAndServe(":10000", nil))
}

func connectToRedis() *redis.Client {
	var client *redis.Client
	for {
		client = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		})
		_, err := client.Ping().Result()
		if err != nil {
			fmt.Println("gameserver could not connect to redis")
			fmt.Println(err)
		} else {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("gameserver connected to redis")

	return client
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	gameserver.PlayerJoined(conn)
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

	if len(gs.Users) > 1 && gs.GL.Running == false {
		gs.GL.Start()
		fmt.Println("started")
	}
}

func validateToken(token string) bool {
	// TODO validate token
	return true
}

func (gs *GameServer) PublishState(msg string) {
	gs.RedisClient.Set(gs.ID, msg, 0)
}

// TODO Need a better game start detection and game end detection.
func (gs *GameServer) MapUpdater(delta float64) {
	gs.PublishState("game started")
	gs.World.Update()
	view := gs.World.Render()

	for k := range gs.Users {
		//TODO reduce size of this
		go k.Conn.WriteJSON(&view)
	}
	//fmt.Println(time.Now())

	// TODO It's possible that the game ends before everyone joins
	if len(gs.World.Players) == 1 {
		// TODO Cleanup
		gs.PublishState("game finished")
		os.Exit(0)
	}
}
