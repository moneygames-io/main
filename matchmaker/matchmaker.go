package main

import (
	"fmt"

	"github.com/gorilla/websocket"

	"github.com/go-redis/redis"
)

type Matchmaker struct {
	statusChannels    []chan int
	gameserverChannel []chan Server
	redisClient       *redis.Client
	CurrentClients    int
	TargetClients     int
}

// TODO https://stackoverflow.com/questions/7893776/the-most-accurate-way-to-check-js-objects-type
// This and type switches are the type of things you need to be dealing with, not these nested messages.
type MatchmakerMessage struct {
	Status Matchmaker
}

type Server struct {
	Url string
}

type ServerInfoMessage struct {
	Url Server
}

func NewMatchmaker(target int) *Matchmaker {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Matchmaker could not connect to redis")
		fmt.Println(err)
		return
	}
	return &Matchmaker{nil, nil, client, 0, target}
}

func (m *Matchmaker) popServer() string {
	c := m.redisClient

	keys, _ := c.Keys("*").Result()

	for _, key := range keys {
		status, _ := c.Get(key).Result()
		if status == "idle" {
			c.Set(key, "waiting for players", 0).Err() // TODO Handle error
			return key
		}
	}
}

func (m *Matchmaker) PlayerJoined(conn *websocket.Conn) {
	m.CurrentClients++

	status := make(chan int)
	m.statusChannels = append(m.statusChannels, status)

	gameserver := make(chan Server)
	m.gameserverChannel = append(m.gameserverChannel, gameserver)

	go m.syncMatchmaker(conn, status, gameserver)

	for _, statusChannel := range m.statusChannels {
		statusChannel <- 1
	}

	if m.CurrentClients == m.TargetClients {

		selectedServer := popServer()

		for _, gameChannel := range m.gameserverChannel {
			gameChannel <- selectedServer
		}

		m.CurrentClients = 0
		//TODO cleanup here
	}
}

func (m *Matchmaker) syncMatchmaker(conn *websocket.Conn, status chan int, gameserver chan Server) {
	for {
		select {
		case <-status:
			if err := conn.WriteJSON(MatchmakerMessage{*m}); err != nil {
				fmt.Println(err)
			}
		case gs := <-gameserver:
			if err := conn.WriteJSON(ServerInfoMessage{gs}); err != nil {
				fmt.Println(err)
			}
			conn.Close()
		}
	}
}
