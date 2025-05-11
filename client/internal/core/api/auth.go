package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/brotigen23/goph-keeper/client/internal/core/dto/auth/logindto"
	"github.com/go-resty/resty/v2"
)

const (
	Register = iota
	Login
)

const (
	registerPath = "/register"
	loginPath    = "/login"
)

func (c *RESTClient) auth(w int, l, p string) error {
	credentials := logindto.PostRequest{Login: l, Password: p}
	requestBody, err := json.Marshal(credentials)
	var response *resty.Response
	if err != nil {
		return err
	}
	request := c.client.R()
	request.Body = requestBody
	switch w {
	case Register:
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

func (c *RESTClient) Register(login, password string) error {
	return c.auth(Register, login, password)
}

func (c *RESTClient) Login(login, password string) error {
	return c.auth(Login, login, password)
}

func (c RESTClient) GetJWT() string {
	return c.jwt
}

func (c *RESTClient) SetJWT(jwt string) {
	c.jwt = jwt
}
