package esqueryextractor

import (
	"sync"
	"time"

	elasticV6 "gopkg.in/olivere/elastic.v6"
)

// Contains all possible client that supported by this package
type Client struct {
	service         Service
	olivereV6Client Olivere6Client
	rateLimiter     *RateLimiter
}

type Service string

type Olivere6Client *elasticV6.Client

type ClientOptionFunc func(*Client) error

type RateLimiter struct {
	// Maximum number of objects that can be created at a given time
	maxSize int
	// Queue holds all the callback functions that are waiting to be executed
	queue []func()
	// Time window for rate limiter
	time time.Duration
	// Timer for each batch of callback functions
	timer *time.Timer
	// Number of callback functions that can be executed per time window
	batchSize int
	// Channel that holds callback functions that are waiting to be executed
	ch chan func()
	// Mutex for protect variable from concurrent access
	mutex sync.Mutex
}

const DefaultService = "undefined"
const DefaultUsecase = "undefined"

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

func SetRateLimiter(maxSize int, sec int32, batchSize int) ClientOptionFunc {
	return func(c *Client) error {
		rateLimiter := &RateLimiter{
			maxSize:   maxSize,
			queue:     make([]func(), 0),
			time:      time.Duration(sec) * time.Second,
			timer:     time.NewTimer(time.Duration(sec) * time.Second),
			ch:        make(chan func(), 1),
			batchSize: batchSize,
			mutex:     sync.Mutex{},
		}
		c.rateLimiter = rateLimiter
		go c.rateLimiter.schedule()
		return nil
	}
}

func (r *RateLimiter) AddFunc(callback func()) {
	r.ch <- callback
}

func (r *RateLimiter) schedule() {
	for {
		select {
		case <-r.timer.C:
			r.flush()
		case data := <-r.ch:
			r.mutex.Lock()
			if len(r.queue) < r.maxSize {
				r.queue = append(r.queue, data)
			}
			r.mutex.Unlock()
		}
	}
}

func (r *RateLimiter) flush() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	n := r.batchSize
	if len(r.queue) < r.batchSize {
		n = len(r.queue)
	}
	sema := make(chan struct{}, n)
	for _, callback := range r.queue {
		sema <- struct{}{}
		go func(callback func()) {
			callback()
			<-sema
		}(callback)
	}
	if r.timer.C != nil && !r.timer.Stop() {
		r.timer.Reset(r.time)
	}
	r.queue = make([]func(), 0)
}

func (c *Client) GetService() Service {
	return c.service
}

func (c *Client) GetOlivere6Client() *elasticV6.Client {
	return c.olivereV6Client
}

func (c *Client) GetRateLimiter() *RateLimiter {
	return c.rateLimiter
}
