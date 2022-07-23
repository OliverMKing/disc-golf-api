package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var (
	serverCmd = &cobra.Command{
		Use:   "server [flags]",
		Short: "Starts the Disc Golf API server",
		RunE:  RunServer,
	}
)

func RunServer(_ *cobra.Command, _ []string) error {
	fmt.Println("Hello world!")

	return nil
}
