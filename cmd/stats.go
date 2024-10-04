package cmd

import (
	"fmt"
	"strings"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats [flags] miner_host",
	Short: "Gets statistics form the miner",
	Long:  `Gets statistics from the miner`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.GetStats(args[0])

		fmt.Println("Statistics information:")
		for _, detail := range result[1:] {
			fmt.Println("  " + strings.ReplaceAll(detail, "=", " = "))
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
