package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")

	http.Handle("/", http.FileServer(http.Dir("../../frontend/src")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server stoping...")
}
