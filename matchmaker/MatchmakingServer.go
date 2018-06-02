package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var matchmaker *Matchmaker

func main() {
	matchmaker = NewMatchmaker(2)

	http.HandleFunc("/ws", wsHandler)

	panic(http.ListenAndServe(":8001", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	matchmaker.PlayerJoined(conn)
}
