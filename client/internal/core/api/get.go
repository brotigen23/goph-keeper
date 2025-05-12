package api

import (
	"encoding/json"
	"net/http"

	"github.com/brotigen23/goph-keeper/client/internal/core/dto/accountdto"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/binarydto"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/cardsdto"
	"github.com/brotigen23/goph-keeper/client/internal/core/dto/textdto"
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
	ret := []accountdto.GetResponse{}
	resp, err := c.get("/user/accounts/fetch")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c RESTClient) GetText() ([]textdto.GetResponse, error) {
	ret := []textdto.GetResponse{}
	resp, err := c.get("/user/text/fetch")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c RESTClient) GetBinary() ([]binarydto.GetResponse, error) {
	ret := []binarydto.GetResponse{}
	resp, err := c.get("/user/binary/fetch")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c RESTClient) GetCards() ([]cardsdto.GetResponse, error) {
	ret := []cardsdto.GetResponse{}
	resp, err := c.get("/user/binary/fetch")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
