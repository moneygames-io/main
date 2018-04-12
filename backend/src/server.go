package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type msg struct {
	Msg string
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	go reply(conn)
}

func reply(conn *websocket.Conn) {
	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		r := msg{}
		switch m.Msg {
		case "hi":
			r.Msg = "hey"
		case "hello":
			r.Msg = "sup"
		default:
			r.Msg = "Don't know that one."
		}
		if err = conn.WriteJSON(r); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	fmt.Println("Server starting...")

	http.Handle("/", http.FileServer(http.Dir("../../frontend/src")))
	http.HandleFunc("/ws", wsHandler)

	panic(http.ListenAndServe(":8080", nil))

	fmt.Println("Server stoping...")
}
