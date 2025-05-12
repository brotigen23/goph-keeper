package cmd

import (
	"github.com/brotigen23/goph-keeper/client/internal/core/service"
	"github.com/spf13/cobra"
)

var (
	login       string
	password    string
	authService *service.Auth
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Keeper",
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.PersistentFlags().StringVarP(&login, "login", "l", "", "Login")
	authCmd.MarkFlagRequired("login")

	authCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password")
	authCmd.MarkFlagRequired("password")

}
