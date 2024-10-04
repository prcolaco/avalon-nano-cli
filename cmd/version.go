package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "development"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show program version",
	Long:  `Show program version`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
