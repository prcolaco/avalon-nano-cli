package cmd

import (
	"fmt"
	"strings"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var deviceCmd = &cobra.Command{
	Use:   "device [flags] miner_host",
	Short: "Request device information from the miner",
	Long:  `Request device information from the miner and present it in a human readable format`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.GetVersion(args[0])

		fmt.Println("Device information:")
		for _, ver := range result {
			fmt.Println("  " + strings.ReplaceAll(ver, "=", " = "))
		}
	},
}

func init() {
	rootCmd.AddCommand(deviceCmd)
}
