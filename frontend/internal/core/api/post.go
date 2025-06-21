package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/binarydto"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/cardsdto"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/textdto"
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
		return response.Body(), nil
	default:
		return nil, errors.New(string(response.Body()))
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

func (c RESTClient) PostText(text textdto.PostRequest) (*textdto.PostResponse, error) {
	body, err := json.Marshal(text)
	if err != nil {
		return nil, err
	}
	response, err := c.post("/user/text", body)
	if err != nil {
		return nil, err
	}
	var ret *textdto.PostResponse
	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c RESTClient) PostBinary(binary binarydto.PostRequest) (*binarydto.PostResponse, error) {
	body, err := json.Marshal(binary)
	if err != nil {
		return nil, err
	}
	response, err := c.post("/user/binary", body)
	if err != nil {
		return nil, err
	}
	var ret *binarydto.PostResponse
	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c RESTClient) PostCard(card cardsdto.PostRequest) (*cardsdto.PostResponse, error) {
	body, err := json.Marshal(card)
	if err != nil {
		return nil, err
	}
	response, err := c.post("/user/binary", body)
	if err != nil {
		return nil, err
	}
	var ret *cardsdto.PostResponse
	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
