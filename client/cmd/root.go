package cmd

import (
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
		case operation := <-f.Out:
			c.In <- operation
		case operation := <-c.Out:
			c.In <- operation
		}
	}
}
