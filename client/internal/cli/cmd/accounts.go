package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	method string
)

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Authenticate with Keeper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching accounts data...")

		/*
			Check if sqlite db exists

			Check if server has a new data
			if so, fetch data or print current data
			from sqlite and let know how to change
			or create a new data

			All it should work in service
			and command should just take return of this
		*/

		switch method {
		case "get":
			fmt.Println(getAccounts())
		default:
			fmt.Println("Method is not allowed")
		}
	},
}

func init() {
	rootCmd.AddCommand(accountsCmd)

	registerCmd.Flags().StringVarP(&method, "method", "m", "get", "method")
}

func getAccounts() string {
	accounts, err := accountsService.GetAccounts()
	if err != nil {
		return err.Error()
	}
	var ret string
	for i := range accounts {
		ret += fmt.Sprintf(
			"[%d] login: %s, password: %s\n",
			accounts[i].ID,
			accounts[i].Login,
			accounts[i].Password)
	}
	return ret
}
