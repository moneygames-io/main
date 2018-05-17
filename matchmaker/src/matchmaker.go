package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Matchmaker struct {
	statusChannels    []chan int
	gameserverChannel []chan Server
	availableServers  []Server
	busyServers       []Server
	CurrentClients    int
	TargetClients     int
}

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
	availableServers := []Server{Server{"localhost:10000"}}
	busyServers := []Server{}
	return &Matchmaker{nil, nil, availableServers, busyServers, 0, target}
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

		selectedServer := m.availableServers[0]

		m.availableServers = m.availableServers[1:]
		m.busyServers = append(m.busyServers, selectedServer)

		// TODO no available game servers. Spin up some more game servers. Find the new ones, dispatch a game

		for _, gameChannel := range m.gameserverChannel {
			gameChannel <- selectedServer
		}

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