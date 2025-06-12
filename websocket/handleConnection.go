package websocket

import (
	"log"
	"net/http"

	"vtt_api/models"
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Obter nome de usuário e sala da URL
	query := r.URL.Query()
	username := query.Get("username")
	roomName := query.Get("room")

	if username == "" || roomName == "" {
		ws.Close()
		return
	}

	client := &models.Client{
		Conn:     ws,
		Username: username,
		Room:     roomName,
	}

	room := CreateRoomIfNotExists(roomName)
	room.register <- client

	defer func() {
		room.unregister <- client
	}()

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Erro na leitura: %v", err)
			break
		}

		msg.Username = client.Username
		msg.Room = client.Room
		room.broadcast <- msg
	}
}
