package cmd

import (
	"fmt"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var ledModeCmd = &cobra.Command{
	Use:   "mode [flags] miner_host mode",
	Short: "Sets the mode of the miner led",
	Long:  `Sets the mode of the miner led, possible modes are off, fixed, flash, pulse or loop`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.SetLedMode(args[0], args[1])
		if result {
			fmt.Println("Led mode set to", args[1])
		}
	},
}

func init() {
	ledCmd.AddCommand(ledModeCmd)
}
