package cmd

import (
	"fmt"
	"os"

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

		err = os.WriteFile("./.cred", []byte(client.GetJWT()), 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
}
