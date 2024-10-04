package cmd

import (
	"fmt"

	"github.com/prcolaco/avalon-nano-cli/internal/nano"

	"github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Miner password commands",
	Long:  `Allows to the change password of the miner`,
}

var passwordChangeCmd = &cobra.Command{
	Use:   "change [flags] miner_host current_password new_password",
	Short: "Changes the password of the miner",
	Long:  `Changes the password of the miner, requires the current password to authenticate`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		result := nano.ChangePassword(args[0], args[1], args[2])
		if result {
			fmt.Println("Password changed")
		}
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.AddCommand(passwordChangeCmd)
}
