package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"github.com/SergeyCherepiuk/share/client/pkg/file"
	"github.com/spf13/cobra"
)

var (
	CreateCommand *cobra.Command
	path          string
)

func init() {
	CreateCommand = &cobra.Command{Use: "create", Run: create}
	CreateCommand.Flags().StringVarP(&path, "output", "o", "", "temporary file")
}

func create(cmd *cobra.Command, args []string) {
	// Create temporary file
	if _, err := os.Stat(path); err == nil {
		log.Fatal("file already exists")
	}

	if _, err := os.Create(path); err != nil {
		log.Fatal(err)
	}
	clean.Add(func() { os.Remove(path) })

	fileContents, err := file.Listen(path)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Printf("File's content has been updated: %s\n", <-fileContents)
	}

	// TODO: Open websocket connection with the server and create the room

	// TODO: Compute diff between current and previous version

	// TODO: Prepare OT (operational transformation) events

	// TODO: Send event over websocket connection
}
