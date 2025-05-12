package cmd

import (
	"fmt"

	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
	"github.com/spf13/cobra"
)

var (
	binaryMethod   string
	binaryID       int
	binaryPath     string
	binaryMetadata string

	// Dir path
)

var binaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "Authenticate with Keeper",
	Run: func(cmd *cobra.Command, args []string) {
		switch method {
		case "post":
			fmt.Println(postBinary())
		case "put":
			fmt.Println(putBinary())
		case "get":
			fmt.Println(getBinary())
		case "delete":
			fmt.Println(deleteBinary())
		default:
			fmt.Println("Method is not allowed")
			fmt.Println("Use with post, put, get or delete")
		}
	},
}

func init() {
	rootCmd.AddCommand(binaryCmd)

	binaryCmd.Flags().StringVarP(&binaryMethod, "method", "m", "get", "method")
	binaryCmd.Flags().IntVarP(&binaryID, "id", "i", 0, "id")
	binaryCmd.Flags().StringVarP(&binaryPath, "path", "p", "", "path")
	binaryCmd.Flags().StringVarP(&binaryMetadata, "metadata", "d", "", "metadata")
}

func getBinary() string {
	accounts, err := accountsService.Fetch()
	if err != nil {
		return err.Error()
	}
	var ret string
	for i := range accounts {
		ret += accountToString(accounts[i])
	}
	ret += "Binary data:\n"
	return ret
}

func postBinary() string {
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

func putBinary() string {
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

func deleteBinary() string {
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

func binaryToString(account domain.Account) string {
	ret := fmt.Sprintf("[%d] login: %s, password: %s, created_at: %s, updated_at: %s\n",
		account.ID,
		account.Login,
		account.Password,
		account.CreatedAt,
		account.UpdatedAt,
	)
	return ret
}
