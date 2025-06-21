package cmd

import (
	"fmt"

	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
	"github.com/spf13/cobra"
)

var (
	method          string
	accountID       int
	accountLogin    string
	accountPassword string
	accountMetadata string
)

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Authenticate with Keeper",
	Run: func(cmd *cobra.Command, args []string) {
		switch method {
		case "post":
			fmt.Println(postAccount())
		case "put":
			fmt.Println(putAccount())
		case "get":
			fmt.Println(getAccounts())
		case "delete":
			fmt.Println(deleteAccount())
		default:
			fmt.Println("Method is not allowed")
			fmt.Println("Use with post, put, get or delete")
		}
	},
}

func init() {
	rootCmd.AddCommand(accountsCmd)

	accountsCmd.Flags().StringVarP(&method, "method", "m", "get", "method")
	accountsCmd.Flags().IntVarP(&accountID, "id", "i", 0, "id")
	accountsCmd.Flags().StringVarP(&accountLogin, "login", "l", "", "login")
	accountsCmd.Flags().StringVarP(&accountPassword, "password", "p", "", "password")
	accountsCmd.Flags().StringVarP(&accountMetadata, "metadata", "d", "", "metadata")
}

func getAccounts() string {
	accounts, err := accountsService.Fetch()
	if err != nil {
		return err.Error()
	}
	var ret string
	for i := range accounts {
		ret += accountToString(accounts[i])
	}
	ret += "Accounts:\n"
	return ret
}

func postAccount() string {
	var ret string
	request := accountdto.PostRequest{}
	request.Login = accountLogin
	request.Password = accountPassword
	request.Metadata = accountMetadata
	account, err := accountsService.Create(request)
	if err != nil {
		return err.Error()
	}
	ret += "Create successfull\n"
	ret += accountToString(*account)
	return ret
}

func putAccount() string {
	var ret string

	request := accountdto.PutRequest{}
	request.ID = accountID
	request.Login = accountLogin
	request.Password = accountPassword
	request.Metadata = accountMetadata
	err := accountsService.Update(request)
	if err != nil {
		return err.Error()
	}
	ret += "Updated successful\n"
	return ret
}

func deleteAccount() string {
	var ret string
	request := accountdto.DeleteRequest{
		ID: accountID,
	}
	err := accountsService.Delete(request)
	if err != nil {
		return err.Error()
	}
	ret += "Deleted successfull"
	return ret
}

func accountToString(account domain.Account) string {
	ret := fmt.Sprintf("[%d] login: %s, password: %s, created_at: %s, updated_at: %s\n",
		account.ID,
		account.Login,
		account.Password,
		account.CreatedAt,
		account.UpdatedAt,
	)
	return ret
}
