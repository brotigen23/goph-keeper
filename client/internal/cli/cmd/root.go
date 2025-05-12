package cmd

import (
	"fmt"
	"os"

	"github.com/brotigen23/goph-keeper/client/internal/core/api"
	"github.com/brotigen23/goph-keeper/client/internal/core/service"
	"github.com/spf13/cobra"
)

var accountsService *service.Accounts
var textService *service.Text
var binaryService *service.Binary
var cardsService *service.Cards
var client *api.RESTClient

// Need to init servicies with that client
func Init(c *api.RESTClient) {
	client = c
}

// Init needed servicies
func preRun(cmd *cobra.Command, args []string) error {
	command := os.Args[1]
	if len(os.Args) > 1 {
		switch command {
		case "auth":
			authService = service.NewAuth(client)
		case "accounts":
			accountsService = service.NewAccounts(client)
		case "text":
			textService = service.NewText(client)
		case "binary":
			binaryService = service.NewBinary(client)
		case "cards":
			cardsService = service.NewCards(client)
		}
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:               "keeper <command> [flags]",
	PersistentPreRunE: preRun,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
