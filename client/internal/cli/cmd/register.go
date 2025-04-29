package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Authenticate with Keeper",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			TODO:
			After register execute fetch data
			into sqlite
			and then print information about it
		*/
		login, err := cmd.Flags().GetString("login")
		if err != nil {
			fmt.Println(err)
			return
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(client.Register(login, password))
	},
}

func init() {
	authCmd.AddCommand(registerCmd)
}
