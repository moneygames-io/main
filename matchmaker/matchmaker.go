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

// https://stackoverflow.com/questions/7893776/the-most-accurate-way-to-check-js-objects-type
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
	availableServers := []Server{Server{"ws://127.0.0.1:10000/ws"}}
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
