package ws

import "golang.org/x/net/websocket"

const (
	TYPE_MESSAGE = 0

	// TYPE_FETCH_REQUEST  = 1
	// TYPE_FETCH_RESPONSE = 2
	// TYPE_POLL_REQUEST   = 3
	// TYPE_POLL_RESPONSE  = 4
)

type Message struct {
	Type int    `json:"type"`
	Msg  string `json:"message"`
	from *websocket.Conn
}
