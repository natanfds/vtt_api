package main

import (
	"log"
	"net/http"

	"vtt_api/websocket"
)

func main() {
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)
	http.HandleFunc("/chat", websocket.HandleConnections)

	log.Println("Runnig at 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
