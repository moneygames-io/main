package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type GameServer struct {
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	panic(http.ListenAndServe(":10000", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	matchmaker.PlayerJoined(conn)
}
