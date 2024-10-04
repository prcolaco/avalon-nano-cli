package cmd

import (
	"fmt"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var miningPoolCmd = &cobra.Command{
	Use:   "pool [flags] miner_host password pool_index pool_url pool_user pool_pass",
	Short: "Sets the details of one of the miner mining pools",
	Long:  `Sets the details of one of the miner mining pools, url, username, password, etc`,
	Args:  cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.SetMiningPool(args[0], args[1], args[2], args[3], args[4], args[5])
		if result {
			fmt.Println("Mining pool", args[2], "updated")
		}
	},
}

func init() {
	miningCmd.AddCommand(miningPoolCmd)
}
