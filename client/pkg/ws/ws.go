package ws

import (
	"encoding/json"

	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"github.com/SergeyCherepiuk/share/client/pkg/diff"
	"golang.org/x/net/websocket"
)

type Connection struct {
	In  chan diff.Operation
	Out chan diff.Operation

	conn *websocket.Conn
}

func New(url, origin string) (*Connection, error) {
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		return nil, err
	}
	clean.Add(func() { conn.Close() })

	connection := Connection{
		In:   make(chan diff.Operation),
		Out:  make(chan diff.Operation),
		conn: conn,
	}

	go connection.listen()
	go connection.send()

	return &connection, nil
}

func (c *Connection) listen() {
	var message string
	for {
		if err := websocket.Message.Receive(c.conn, &message); err != nil {
			continue
		}

		var deserialized struct {
			Type    int    `json:"type"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal([]byte(message), &deserialized); err != nil {
			continue
		}

		var operation diff.Operation
		if err := json.Unmarshal([]byte(deserialized.Message), &operation); err == nil {
			c.Out <- operation
		}
	}
}

func (c *Connection) send() {
	for {
		operation := <-c.In
		if serialized, err := json.Marshal(operation); err == nil { // TODO: Handle an error
			websocket.Message.Send(c.conn, serialized) // TODO: Handle an error
		}
	}
}
