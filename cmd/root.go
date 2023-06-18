package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fontivate",
	Short: "Fontivate allows you to manage fonts completely from the command line.",
	Args:  cobra.NoArgs,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = "[0.1.0]"
}
