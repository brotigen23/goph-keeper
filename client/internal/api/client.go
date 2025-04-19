package api

import (
	"encoding/json"
	"strings"

	"github.com/brotigen23/goph-keeper/client/internal/dto"
	"github.com/go-resty/resty/v2"
)

type Response struct {
	Body       string
	StatusCode int
	Err        error
}

type Client struct {
	client *resty.Client
	JWT    string
}

func New(baseURL string) *Client {
	client := resty.New().SetBaseURL(baseURL)
	return &Client{
		client: client,
	}
}

func (c Client) Ping() error {
	_, err := c.client.R().Get("/ping")
	return err
}
func (c *Client) Register(login, password string) Response {
	credentials := dto.Login{Login: login, Password: password}
	requestBody, err := json.Marshal(credentials)
	if err != nil {
		return Response{Err: err}
	}
	request := c.client.R()
	request.Body = requestBody
	response, err := request.Post("/register")
	if err != nil {
		return Response{Err: err}
	}
	authHeader := response.Header()
	token := strings.TrimPrefix(authHeader.Get("Authorization"), "Bearer ")
	c.JWT = token

	return Response{Body: string(response.Body()), StatusCode: response.StatusCode(), Err: nil}
}

func (c *Client) Login(login, password string) Response {
	credentials := dto.Login{Login: login, Password: password}
	requestBody, err := json.Marshal(credentials)
	if err != nil {
		return Response{Err: err}
	}
	request := c.client.R()
	request.Body = requestBody
	response, err := request.Post("/login")
	if err != nil {
		return Response{Err: err}
	}
	authHeader := response.Header()
	token := strings.TrimPrefix(authHeader.Get("Authorization"), "Bearer ")
	c.JWT = token

	return Response{Body: string(response.Body()), StatusCode: response.StatusCode(), Err: nil}
}

func (c Client) GetData(path string) Response {
	request := c.client.R()
	request.Header.Set("Authorization", "Bearer "+c.JWT)
	response, err := request.Get(path)
	if err != nil {
		return Response{Err: err}
	}

	return Response{StatusCode: response.StatusCode(), Body: string(response.Body())}
}

func (c Client) GetAccounts() Response {
	request := c.client.R()
	request.Header.Set("Authorization", "Bearer "+c.JWT)
	response, err := request.Get("/user/accounts")
	if err != nil {
		return Response{Err: err}
	}

	return Response{StatusCode: response.StatusCode(), Body: string(response.Body())}
}
