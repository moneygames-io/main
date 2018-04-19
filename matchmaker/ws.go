package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

type Msg struct {
	N int
	M int
}

var messages []chan int
var sockets int = 0
var poolsize int = 100

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
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

	message := make(chan int)
	messages = append(messages, message)

	go echo(conn, message)
	sockets++

	for _, m := range messages {
		m <- 1
	}
}

func echo(conn *websocket.Conn, message chan int) {
	for {
		select {
			case <- message:
				fmt.Println(Msg{sockets, poolsize})
				if err := conn.WriteJSON(Msg{sockets, poolsize}); err != nil {
					fmt.Println(err)
				}
		}
	}
}
