package cmd

import (
	"fmt"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var workCmd = &cobra.Command{
	Use:   "work [flags] miner_host",
	Short: "Gets work level information from the miner",
	Long:  `Gets work level information from the miner, like if the it is working at low, medium or high level`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.GetWorkLevel(args[0])

		fmt.Println("Work level:", result)
	},
}

func init() {
	rootCmd.AddCommand(workCmd)
}
