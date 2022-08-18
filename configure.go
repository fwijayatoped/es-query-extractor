package esqueryextractor

import (
	"net/http"

	elasticV6 "gopkg.in/olivere/elastic.v6"
)

type Client struct {
	httpClient      *http.Client
	service         string
	olivereV6Client Olivere6Client
}

type Olivere6Client struct {
	client *elasticV6.Client
}

type ClientOptionFunc func(*Client) error

const DefaultService = "undefined"

func Initialize(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		service:    DefaultService,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return c, nil
}

func SetService(service string) ClientOptionFunc {
	return func(c *Client) error {
		c.service = service
		return nil
	}
}

func SetHttpClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		if httpClient != nil {
			c.httpClient = httpClient
		} else {
			c.httpClient = http.DefaultClient
		}
		return nil
	}
}

func SetOlivereV6Client(elasticV6Client *elasticV6.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.olivereV6Client.client = elasticV6Client
		return nil
	}
}

func (c *Client) GetOlivere6Client() *elasticV6.Client {
	return c.olivereV6Client.client
}

func (c *Client) GetService() string {
	return c.service
}
