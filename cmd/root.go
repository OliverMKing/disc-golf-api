package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "dg-api",
	Short: "Runs and manages an API containing Disc Golf disc information",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
