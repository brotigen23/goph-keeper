package rest

import (
	"log"
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
)

func (c Client) PutAccount(account domain.AccountData) error {
	request := c.client.R()
	request.Header.Set("Authorization", "Bearer "+c.jwt)
	request.SetBody(account)
	response, err := request.Put("/user/accounts")
	if err != nil {
		return nil
	}
	log.Println(response.StatusCode(), string(response.Body()))
	switch response.StatusCode() {
	case http.StatusAccepted:
		return nil
	default:
		return err
	}
}
