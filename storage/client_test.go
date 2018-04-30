package storage_test

import (
	"time"

	cfg "github.com/arthurjames/kit/config/storage"
	"github.com/arthurjames/nine2five/storage"
)

// Now is the mocked current time for testing.
var Now = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

// Client is a test wrapper for storage.Client.
type Client struct {
	*storage.Client
}

// NewClient returns a new instance of Client.
func NewClient() *Client {
	// Create client wrapper.
	c := &Client{
		Client: storage.NewClient(),
	}
	c.Now = func() time.Time { return Now }
	return c
}

// MustOpenClient returns a new, open instance of Client.
func MustOpenClient(cfg *cfg.StorageConfig) *Client {
	c := NewClient()
	if err := c.Open(cfg); err != nil {
		panic(err)
	}
	return c
}

// Close closes the connection between client and storage.
func (c *Client) Close() error {
	return c.Client.Close()
}
