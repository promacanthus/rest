package example

import (
	"time"

	"github.com/promacanthus/rest/rest"
)

type Interface interface {
	// TODO: customize interface
}

type Client struct {
	restClient rest.Interface
}

// NewClient creates a new lvs client.
func NewClient(c *rest.Config) Interface {
	// Check the configuration for default values.
	if c.UserAgent == "" {
		c.UserAgent = rest.DefaultUserAgent()
	}
	if c.RetryCount == 0 {
		c.RetryCount = 3
	}
	if c.RetryWaitTime == 0 {
		c.RetryWaitTime = 500 * time.Millisecond
	}
	if c.RetryMaxWaitTime == 0 {
		c.RetryMaxWaitTime = 3 * time.Second
	}
	// Create the rest client.
	client := rest.HTTPClientFor(c)
	// Set the base URL.
	client.SetBaseURL(c.BaseURLs["Name"])
	return &Client{restClient: client}
}
