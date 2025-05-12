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

func (c *RESTClient) put(path string, body []byte) ([]byte, error) {
	request := c.client.R()
	request.Header.Set("Authorization", "Bearer "+c.jwt)
	request.SetBody(body)
	response, err := request.Put(path)
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

func (c *RESTClient) PutAccount(account accountdto.PutRequest) error {
	request, err := json.Marshal(account)
	if err != nil {
		return err
	}
	_, err = c.put("user/accounts", request)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTClient) PutText(text textdto.PutRequest) error {
	request, err := json.Marshal(text)
	if err != nil {
		return err
	}
	_, err = c.put("user/text", request)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTClient) PutBinary(binary binarydto.PutRequest) error {
	request, err := json.Marshal(binary)
	if err != nil {
		return err
	}
	_, err = c.put("user/binary", request)
	if err != nil {
		return err
	}
	return nil
}

func (c *RESTClient) PutCard(card cardsdto.PutRequest) error {
	request, err := json.Marshal(card)
	if err != nil {
		return err
	}
	_, err = c.put("user/cards", request)
	if err != nil {
		return err
	}
	return nil
}
