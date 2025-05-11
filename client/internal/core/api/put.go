package api

import (
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
)

func (c RESTClient) PutAccount(account accountdto.PutRequest) error {
}

func (c *RESTClient) put(path string, body []byte) ([]byte, error) {
	request := c.client.R()
	request.Header.Set("Authorization", "Bearer "+c.jwt)
	request.SetBody(body)
	response, err := request.Put(path)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode() {
	case http.StatusAccepted:
		return response.Body(), nil
	default:
		return nil, err
	}
}
