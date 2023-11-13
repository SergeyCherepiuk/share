package ws

import (
	"encoding/json"

	"github.com/SergeyCherepiuk/share/client/pkg/diff"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
	"golang.org/x/net/websocket"
)

func Listen(conn *websocket.Conn) {
	var message string
	for {
		if err := websocket.Message.Receive(conn, &message); err != nil {
			continue
		}

		var deserialized struct {
			Type    int    `json:"type"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal([]byte(message), &deserialized); err != nil {
			continue
		}

		var operations []diff.Operation
		if err := json.Unmarshal([]byte(deserialized.Message), &operations); err != nil {
			continue
		}

		for _, operation := range operations {
			ot.Operations <- operation
		}
	}
}
