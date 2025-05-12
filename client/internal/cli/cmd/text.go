package cmd

import (
	"fmt"

	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/textdto"
	"github.com/spf13/cobra"
)

var (
	textMethod   string
	textID       int
	textData     string
	textMetadata string
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Authenticate with Keeper",
	Run: func(cmd *cobra.Command, args []string) {
		switch textMethod {
		case "post":
			fmt.Println(postText())
		case "put":
			fmt.Println(putText())
		case "get":
			fmt.Println(getText())
		case "delete":
			fmt.Println(deleteText())
		default:
			fmt.Println("Method is not allowed")
			fmt.Println("Use with post, put, get or delete")
		}
	},
}

func init() {
	rootCmd.AddCommand(textCmd)

	textCmd.Flags().StringVarP(&textMethod, "method", "m", "get", "method")
	textCmd.Flags().IntVarP(&textID, "id", "i", 0, "id")
	textCmd.Flags().StringVarP(&textData, "data", "d", "", "data")
	textCmd.Flags().StringVar(&textMetadata, "metadata", "", "metadata")
}

func getText() string {
	accounts, err := textService.Fetch()
	if err != nil {
		return err.Error()
	}
	var ret string
	ret += "Text data:\n"
	for i := range accounts {
		ret += textDataToString(accounts[i])
	}
	return ret
}

func postText() string {
	var ret string
	request := textdto.PostRequest{}
	request.Data = textData
	request.Metadata = textMetadata
	account, err := textService.Create(request)
	if err != nil {
		return err.Error()
	}
	ret += "Create successfull\n"
	ret += textDataToString(*account)
	return ret
}

func putText() string {
	var ret string

	request := textdto.PutRequest{}
	request.ID = textID
	request.Data = textData
	request.Metadata = accountMetadata
	err := textService.Update(request)
	if err != nil {
		return err.Error()
	}
	ret += "Updated successful\n"
	return ret
}

func deleteText() string {
	var ret string
	request := textdto.DeleteRequest{
		ID: textID,
	}
	err := textService.Delete(request)
	if err != nil {
		return err.Error()
	}
	ret += "Deleted successfull"
	return ret
}

func textDataToString(text domain.TextData) string {
	ret := fmt.Sprintf("[%d] data: %s, metadata: %s, created_at: %s, updated_at: %s\n",
		text.ID,
		text.Data,
		text.Metadata,
		text.CreatedAt,
		text.UpdatedAt,
	)
	return ret
}
