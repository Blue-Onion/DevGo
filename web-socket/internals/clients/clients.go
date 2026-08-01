package clients

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	Send chan []byte
	conn *websocket.Conn
}

var upgarder = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
