package ws

import (
	"encoding/json"

	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
	"golang.org/x/net/websocket"
)

func Send(conn *websocket.Conn, contents <-chan []byte) {
	prev := make([]byte, 0)
	for {
		curr := <-contents
		operations := ot.Diff(prev, curr)
		prev = curr

		if serialized, err := json.Marshal(operations); err == nil { // TODO: Handler potential error
			conn.Write(serialized) // TODO: Handler potential error
		}
	}
}
