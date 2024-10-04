package cmd

import (
	"fmt"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var workLevelCmd = &cobra.Command{
	Use:   "level [flags] miner_host level",
	Short: "Sets the work level of the miner",
	Long:  `Sets the work level of the miner, possible level values are low, med or high`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.SetWorkLevel(args[0], args[1])
		if result {
			fmt.Println("Work level set to", args[1])
		}
	},
}

func init() {
	workCmd.AddCommand(workLevelCmd)
}
