package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
	"github.com/SergeyCherepiuk/share/client/pkg/file"
	"github.com/SergeyCherepiuk/share/client/pkg/ws"
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
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		return err
	}
	clean.Add(func() { conn.Close() })

	go ws.Listen(conn)         // Listening for changes from other user
	go ot.Apply(path)          // Applying those changes
	go ws.Send(conn, contents) // Sending client's changes to the server

	return <-make(chan error) // Block forever
}
