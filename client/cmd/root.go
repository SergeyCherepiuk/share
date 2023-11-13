package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
	"github.com/SergeyCherepiuk/share/client/pkg/file"
	"github.com/spf13/cobra"
	"golang.org/x/net/websocket"
)

var (
	RootCmd *cobra.Command
	path    string
	url     string
	origin  string
)

func init() {
	RootCmd = &cobra.Command{RunE: root}
	RootCmd.PersistentFlags().StringVarP(&path, "output", "o", "", "path to temporary file")
	RootCmd.AddCommand(createCmd, joinCmd)
}

func root(cmd *cobra.Command, args []string) error {
	// Create temporary file
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file already exists")
	}

	if _, err := os.Create(path); err != nil {
		return err
	}
	clean.Add(func() { os.Remove(path) })

	// Listen for file updates
	contents, err := file.Listen(path, 100*time.Millisecond)
	if err != nil {
		return err
	}

	// Open websocket connection with the server and create the room
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		return err
	}
	clean.Add(func() { ws.Close() })

	// TODO: Extract into separate method/package and optimize
	// Listening for changes received from server
	go func() {
		for {
			buf := make([]byte, 512)
			n, _ := ws.Read(buf)
			fmt.Println(string(buf[:n]))
		}
	}()

	// Compute difference between previous and current file states
	prev := []byte("")
	for {
		curr := <-contents
		operations := ot.Diff(prev, curr)
		prev = curr

		// Sending operations to the server
		if serialized, err := json.Marshal(operations); err == nil { // TODO: Handler potential error
			ws.Write(serialized) // TODO: Handler potential error
		}
	}
}
