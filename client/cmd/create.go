package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd *cobra.Command

func init() {
	createCmd = &cobra.Command{Use: "create", RunE: create}
}

func create(cmd *cobra.Command, args []string) error {
	url, _, origin = createCmdConnectionDetails()
	return RootCmd.RunE(cmd, args)
}

func createCmdConnectionDetails() (string, string, string) {
	url := fmt.Sprintf(
		"ws://%s:%s/%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
		os.Getenv("CREATE_ENDPOINT"),
	)
	origin := fmt.Sprintf(
		"http://%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	return url, "", origin
}
