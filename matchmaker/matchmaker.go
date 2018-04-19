package main

import(
	"fmt"
	"github.com/gorilla/websocket"
)

type Matchmaker struct {
	channels []chan int
	CurrentClients int
	TargetClients int
}

func NewMatchmaker(target int) *Matchmaker {
	return &Matchmaker{nil, 0, target}
}

func (m *Matchmaker) PlayerJoined (conn *websocket.Conn) {
	m.CurrentClients++
	status := make(chan int)
	m.channels = append(m.channels, status)
	go m.syncMatchmaker(conn, status)

	for _, channel := range m.channels {
		channel <- 1
	}
}

func (m *Matchmaker) syncMatchmaker(conn *websocket.Conn, status chan int) {
	for {
		select {
			case <- status:
				if err := conn.WriteJSON(m); err != nil {
					fmt.Println(err)
				}
		}
	}
}
