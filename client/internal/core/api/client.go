package api

import (
	"github.com/go-resty/resty/v2"
)

type Response struct {
	Body       []byte
	StatusCode int
	Err        error
}

type RESTClient struct {
	client *resty.Client
	jwt    string
}

func New(baseURL string) *RESTClient {
	client := resty.New().
		SetBaseURL(baseURL)

	return &RESTClient{
		client: client,
	}
}

func (c RESTClient) Ping() Response {
	response, err := c.client.R().Get("/ping")
	return Response{
		Body:       response.Body(),
		StatusCode: response.StatusCode(),
		Err:        err,
	}
}
