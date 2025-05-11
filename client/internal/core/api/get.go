package api

import (
	"encoding/json"
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
)

func (c *RESTClient) get(path string) ([]byte, error) {
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

func (c RESTClient) GetAccounts() ([]accountdto.GetResponse, error) {
	c.SetJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc3NTQ0MTMsIklEIjozfQ.YO6YUZb8GzXk54y6pAkn4-rzVbQBRdQCtp-AW37AkV4")
	ret := []accountdto.GetResponse{}
	resp, err := c.get("/user/accounts/fetch")
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
