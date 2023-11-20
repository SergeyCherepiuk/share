package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
	"github.com/SergeyCherepiuk/share/client/pkg/file"
	"github.com/SergeyCherepiuk/share/client/pkg/ws"
	"github.com/spf13/cobra"
)

var (
	RootCmd  *cobra.Command
	path     string
	preserve bool
	url      string
	origin   string
)

func init() {
	RootCmd = &cobra.Command{RunE: root}
	RootCmd.PersistentFlags().StringVarP(&path, "output", "o", "", "path to temporary file")
	RootCmd.PersistentFlags().BoolVarP(&preserve, "preserve", "p", false, "preserves the file after quitting")
	RootCmd.AddCommand(createCmd, joinCmd)
}

func root(cmd *cobra.Command, args []string) error {
	// Create temporary file
	f, err := file.New(path, preserve)
	if err != nil {
		return err
	}

	// Open websocket connection with the server and create the room
	c, err := ws.New(url, origin)
	if err != nil {
		return err
	}

	tie(f, c) // Block forever
	return nil
}

func tie(f *file.File, c *ws.Connection) {
	for {
		select {
		case operations := <-f.Out:
			handleOperations(c.In, operations)
		case frame := <-c.Out:
			handleFrame(frame, f, c)
		}
	}
}

func handleOperations(ch chan<- ws.Frame, operations []ot.Operation) {
	if operationsSer, err := json.Marshal(operations); err == nil {
		ch <- ws.Frame{
			Opcode:  ws.OPCODE_OPERATIONS,
			Payload: operationsSer,
		}
	}
}

func handleFrame(frame ws.Frame, f *file.File, c *ws.Connection) {
	switch frame.Opcode {
	case ws.OPCODE_ROOM_INFO:
		fmt.Fprint(os.Stdout, string(frame.Payload))

	case ws.OPCODE_OPERATIONS:
		var operations []ot.Operation
		if err := json.Unmarshal(frame.Payload, &operations); err == nil {
			f.In <- operations
		}

	case ws.OPCODE_CONTENT_REQUEST:
		c.In <- ws.Frame{
			Opcode:  ws.OPCODE_CONTENT_RESPONSE,
			Payload: f.GetContent(),
		}

	case ws.OPCODE_CONTENT_RESPONSE:
		f.SetContent(frame.Payload)
	}
}
