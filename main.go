package main

import (
	"log"
	"net/http"

	apihandlers "vtt_api/apiHandlers"
	"vtt_api/websocket"
)

func main() {
	http.HandleFunc("/chat", websocket.HandleConnections)
	http.HandleFunc("/chat/rooms/active", apihandlers.ListActiveRooms)
	log.Println("Runnig at 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
