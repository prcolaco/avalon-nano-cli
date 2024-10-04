package cmd

import (
	"fmt"
	"strings"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var ledColorCmd = &cobra.Command{
	Use:   "color [flags] miner_host brightness rgb_hex",
	Short: "Sets the color and brightness of the miner led",
	Long:  `Sets the color and brightness of the miner led, brightness in percent, and RGB color in hexadecimal (HTML like)`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		rgbhex := strings.TrimPrefix(args[2], "#")
		result := nano.SetLedColor(args[0], args[1], rgbhex)
		if result {
			fmt.Println("Led brightness set to", args[1], "and color to", args[2])
		}
	},
}

func init() {
	ledCmd.AddCommand(ledColorCmd)
}
