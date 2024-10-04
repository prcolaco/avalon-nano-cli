package cmd

import (
	"fmt"
	"strings"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var wifiCmd = &cobra.Command{
	Use:   "wifi [flags] miner_host",
	Short: "Gets wifi information form the miner",
	Long:  `Gets wifi information from the miner`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.GetWifi(args[0])

		fmt.Println("Wifi information:")
		for _, param := range result {
			fmt.Println("  " + strings.ReplaceAll(param, "=", " = "))
		}
	},
}

func init() {
	rootCmd.AddCommand(wifiCmd)
}
