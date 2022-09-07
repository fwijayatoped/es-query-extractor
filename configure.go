package esqueryextractor

import (
	elasticV6 "gopkg.in/olivere/elastic.v6"
)

// Contains all possible client that supported by this package
type Client struct {
	service         Service
	olivereV6Client Olivere6Client
}

type Service string

type Olivere6Client *elasticV6.Client

type ClientOptionFunc func(*Client) error

const DefaultService = "undefined"

func Initialize(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		service: DefaultService,
	}

	for _, option := range options {
		option(c)
	}

	return c, nil
}

func SetService(service Service) ClientOptionFunc {
	return func(c *Client) error {
		c.service = service
		return nil
	}
}

func SetOlivereV6Client(elasticV6Client *elasticV6.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.olivereV6Client = elasticV6Client
		return nil
	}
}

func (c *Client) GetOlivere6Client() *elasticV6.Client {
	return c.olivereV6Client
}

func (c *Client) GetService() Service {
	return c.service
}
