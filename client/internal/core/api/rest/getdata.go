package rest

import (
	"encoding/json"
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/core/domain"
)

func (c *Client) get(path string) ([]byte, error) {
	request := c.client.R()
	request.Header.Set("Authorization", "Bearer "+c.jwt)
	response, err := request.Get(path)
	if err != nil {
		return nil, err
	}
	switch response.StatusCode() {
	case http.StatusOK:
		return response.Body(), nil
	default:
		return nil, err
	}
}

func (c Client) GetAccounts() ([]domain.AccountData, error) {
	ret := []domain.AccountData{}
	resp, err := c.get("/user/accounts")
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
