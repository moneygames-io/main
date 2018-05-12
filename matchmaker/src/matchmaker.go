package main

import(
	"fmt"
	"github.com/gorilla/websocket"
)

type Matchmaker struct {
	statusChannels []chan int
	gameserverChannel []chan string
	availableServers []string
	busyServers []string
	CurrentClients int
	TargetClients int
}

func NewMatchmaker(target int) *Matchmaker {
	availableServers := []string{"localhost:10000"}
	busyServers := []string{}
	return &Matchmaker{nil, nil, availableServers, busyServers, 0, target}
}

func (m *Matchmaker) PlayerJoined (conn *websocket.Conn) {
	m.CurrentClients++

	status := make(chan int)
	m.statusChannels = append(m.statusChannels, status)

	gameserver := make(chan string)
	m.gameserverChannel = append(m.gameserverChannel, gameserver)

	go m.syncMatchmaker(conn, status, gameserver)

	for _, statusChannel := range m.statusChannels {
		statusChannel <- 1
	}

	if m.CurrentClients == m.TargetClients {
		for _, gameChannel := range m.gameserverChannel {
			gameChannel <- "gameserverurl"
		}

		//TODO cleanup here
	}
}

func (m *Matchmaker) syncMatchmaker(conn *websocket.Conn, status chan int, gameserver chan string) {
	for {
		select {
			case <- status:
				if err := conn.WriteJSON(m); err != nil {
					fmt.Println(err)
				}
			case url := <-gameserver:
				if err := conn.WriteJSON(url); err != nil {
					fmt.Println(err)
				}
				conn.Close()
		}
	}
}
