package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SergeyCherepiuk/share/client/pkg/clean"
	"github.com/SergeyCherepiuk/share/client/pkg/diff/ot"
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

	// Listen for file updates
	contents, err := file.Listen(path, 100*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Open websocket connection with the server and create the room

	// Compute difference between previous and current file states
	prev := []byte("")
	for {
		curr := <-contents
		operations := ot.Diff(prev, curr) // ISSUE: Line swapping doesn't produce operations
		prev = curr

		for i, operation := range operations {
			fmt.Printf("%d: %+v (%T)\n", i, operation, operation)
		}

		// TODO: Send event over websocket connection
	}
}
