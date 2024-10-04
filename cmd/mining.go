package cmd

import (
	"fmt"
	"strings"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var miningCmd = &cobra.Command{
	Use:   "mining [flags] miner_host",
	Short: "Gets mining pools information from the miner",
	Long:  `Gets mining pools information from the miner`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.GetMining(args[0])

		fmt.Println("Mining pools information:", result[0])
		for i, pool := range result[1:] {
			fmt.Println("  Pool", i+1)
			for _, param := range strings.Split(pool, ",")[1:] {
				fmt.Println("    " + strings.ReplaceAll(param, "=", " = "))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(miningCmd)
}
