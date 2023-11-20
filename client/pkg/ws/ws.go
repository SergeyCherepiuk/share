package ws

import (
	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"golang.org/x/net/websocket"
)

type Connection struct {
	In  chan Frame
	Out chan Frame

	conn *websocket.Conn
}

func New(url, origin string) (*Connection, error) {
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		return nil, err
	}
	clean.Add(func() { conn.Close() })

	connection := Connection{
		In:   make(chan Frame),
		Out:  make(chan Frame),
		conn: conn,
	}

	go connection.listen()
	go connection.send()

	return &connection, nil
}

func (c *Connection) listen() {
	var frame Frame
	for {
		if err := websocket.JSON.Receive(c.conn, &frame); err == nil {
			c.Out <- frame
		}
	}
}

func (c *Connection) send() {
	for {
		websocket.JSON.Send(c.conn, <-c.In)
	}
}
