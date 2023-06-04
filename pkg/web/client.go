package web

import (
	"github.com/go-resty/resty/v2"
)

type Client interface {
	Get(url string) (*resty.Response, error)
	GetAndSave(url string, filepath string) (*resty.Response, error)
}

type Response = resty.Response

type client struct {
	*resty.Client
}

func (c *client) Get(url string) (*resty.Response, error) {
	res, err := c.R().
		EnableTrace().
		Get(url)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *client) GetAndSave(url string, filepath string) (*resty.Response, error) {
	res, err := c.R().
		EnableTrace().
		SetOutput(filepath).
		Get(url)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewClient() Client {
	return &client{Client: resty.New()}
}
