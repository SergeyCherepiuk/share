package ws

import "golang.org/x/net/websocket"

const (
	OPCODE_ROOM_INFO        = 0
	OPCODE_OPERATIONS       = 1
	OPCODE_CONTENT_REQUEST  = 2
	OPCODE_CONTENT_RESPONSE = 3
)

type Frame struct {
	Opcode  int    `json:"opcode"`
	Payload []byte `json:"payload"`
	from    *websocket.Conn
}
