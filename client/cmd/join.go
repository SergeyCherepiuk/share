package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	joinCmd *cobra.Command
	id      uuid.UUID
)

func init() {
	joinCmd = &cobra.Command{Use: "join", RunE: join}
}

func join(cmd *cobra.Command, args []string) error {
	var err error
	id, err = uuid.Parse(args[0])
	if err != nil {
		log.Fatal("invalid id")
	}

	url, _, origin = joinCmdConnectionDetails()
	return RootCmd.RunE(cmd, args)
}

func joinCmdConnectionDetails() (string, string, string) {
	url := fmt.Sprintf(
		"ws://%s:%s/%s/%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
		os.Getenv("JOIN_ENDPOINT"),
		id.String(),
	)
	origin := fmt.Sprintf(
		"http://%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	return url, "", origin
}
