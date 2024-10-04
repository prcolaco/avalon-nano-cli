package cmd

import (
	"fmt"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var ledCmd = &cobra.Command{
	Use:   "led [flags] miner_host",
	Short: "Gets led information form the miner",
	Long:  `Gets led information from the miner, like it's mode, brightness or color`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.GetLed(args[0])

		fmt.Println("Led information:")
		fmt.Println("  Mode:", result[0])
		fmt.Println("  Brightness (%):", result[1])
		fmt.Println("  Color (RGB):", "#"+result[2])
	},
}

func init() {
	rootCmd.AddCommand(ledCmd)
}
