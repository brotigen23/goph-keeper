package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with Keeper",
	Run: func(cmd *cobra.Command, args []string) {
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
		err = authService.Login(login, password)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("User %s success login\n", login)
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
}
