package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/brotigen23/goph-keeper/client/internal/core/dto"
	"github.com/go-resty/resty/v2"
)

const (
	register = iota
	Login
)

const (
	registerPath = "/register"
	loginPath    = "/login"
)

func (c *Client) auth(w int, l, p string) error {
	credentials := dto.Login{Login: l, Password: p}
	requestBody, err := json.Marshal(credentials)
	var response *resty.Response
	if err != nil {
		return err
	}
	request := c.client.R()
	request.Body = requestBody
	switch w {
	case register:
		response, err = request.Post(registerPath)
	case Login:
		response, err = request.Post(loginPath)
	}
	switch response.StatusCode() {
	case http.StatusBadRequest:
		return fmt.Errorf("Bad request%s", response.Body())
	case http.StatusUnauthorized:
		return fmt.Errorf("Unauthorized")
	}
	authHeader := response.Header()
	token := strings.TrimPrefix(authHeader.Get("Authorization"), "Bearer ")
	c.jwt = token
	return nil
}

func (c *Client) Register(login, password string) error {
	return c.auth(register, login, password)
}

func (c *Client) Login(login, password string) error {
	return c.auth(Login, login, password)
}

func (c Client) GetJWT() string {
	return c.jwt
}
func (c *Client) SetJWT(jwt string) {
	c.jwt = jwt
}
