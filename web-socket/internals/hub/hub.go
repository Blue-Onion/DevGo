package hub

import (
	"log"

	"github.com/Devgo/webSocket/internals/clients"
)

type Hub struct {
	clients    map[*clients.Client]bool
	broadcast  chan []byte
	register   chan *clients.Client
	unregister chan *clients.Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*clients.Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *clients.Client),
		unregister: make(chan *clients.Client),
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Println("User Register")

		case client := <-h.unregister:
			delete(h.clients, client)
			close(client.Send)
			log.Println("User Unregister")
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}
