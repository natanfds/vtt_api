package websocket

import (
	"log"

	"vtt_api/models"
)

type Room struct {
	clients    map[*models.Client]bool
	broadcast  chan models.Message
	register   chan *models.Client
	unregister chan *models.Client
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.register:
			r.clients[client] = true
			log.Printf("Client %s connected at %s", client.Username, client.Room)
		case client := <-r.unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				client.Conn.Close()
				log.Printf("Client %s leave %s", client.Username, client.Room)
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				err := client.Conn.WriteJSON(message)
				if err == nil {
					continue
				}
				log.Printf("Error to send message to %s, %v", client.Username, err)
				client.Conn.Close()
				delete(r.clients, client)
			}
		}
	}
}

func CreateRoomIfNotExists(roomName string) *Room {
	RoomsMutex.Lock()

	if room, exists := Rooms[roomName]; exists {
		return room
	}

	room := &Room{
		clients:    make(map[*models.Client]bool),
		broadcast:  make(chan models.Message),
		register:   make(chan *models.Client),
		unregister: make(chan *models.Client),
	}

	Rooms[roomName] = room
	go room.run()
	return room
}
