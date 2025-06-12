package apihandlers

import (
	"encoding/json"
	"net/http"

	"vtt_api/websocket"
)

func ListActiveRooms(w http.ResponseWriter, r *http.Request) {
	var roomNames []string
	for name := range websocket.Rooms {
		roomNames = append(roomNames, name)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roomNames)
}
