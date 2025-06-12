package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: Criar validação
	},
}

var Rooms = make(map[string]*Room)
var RoomsMutex = &sync.Mutex{}
