package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
)

func (c *RESTClient) post(path string, body []byte) ([]byte, error) {
	request := c.client.R()
	request.Header.Set("Authorization", "Bearer "+c.jwt)
	request.Body = body
	response, err := request.Post(path)
	if err != nil {
		return nil, err
	}
	switch response.StatusCode() {
	case http.StatusOK:
		return nil, nil
	default:
		return nil, fmt.Errorf(string(response.Body()))
	}
}

func (c RESTClient) PostAccount(account accountdto.PostRequest) (*accountdto.PostResponse, error) {
	body, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}
	response, err := c.post("/user/accounts", body)
	if err != nil {
		return nil, err
	}
	var ret *accountdto.PostResponse
	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
