package cmd

import (
	"fmt"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:   "reboot [flags] miner_host",
	Short: "Reboots the miner",
	Long:  `Reboots the miner, possible level values are low, medium or high`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.Reboot(args[0])
		if result {
			fmt.Println("Miner rebooting, please wait...")
		}
	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)
}
