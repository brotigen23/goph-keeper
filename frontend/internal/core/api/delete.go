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

func (c *RESTClient) delete(path string, body []byte) ([]byte, error) {
	request := c.client.R()
	request.Header.Set("Authorization", "Bearer "+c.jwt)
	request.Body = body
	response, err := request.Delete(path)
	if err != nil {
		return nil, err
	}
	switch response.StatusCode() {
	case http.StatusOK:
		return nil, nil
	default:
		return nil, errors.New(string(response.Body()))
	}
}

func (c RESTClient) DeleteAccount(account accountdto.DeleteRequest) (*accountdto.DeleleResponse, error) {
	body, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}
	_, err = c.delete("/user/accounts", body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c RESTClient) DeleteText(text textdto.DeleteRequest) (*textdto.DeleleResponse, error) {
	body, err := json.Marshal(text)
	if err != nil {
		return nil, err
	}
	_, err = c.delete("/user/text", body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c RESTClient) DeleteBinary(binary binarydto.DeleteRequest) (*binarydto.DeleleResponse, error) {
	body, err := json.Marshal(binary)
	if err != nil {
		return nil, err
	}
	_, err = c.delete("/user/binary", body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c RESTClient) DeleteCard(card cardsdto.DeleteRequest) (*cardsdto.DeleleResponse, error) {
	body, err := json.Marshal(card)
	if err != nil {
		return nil, err
	}
	_, err = c.delete("/user/cards", body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
