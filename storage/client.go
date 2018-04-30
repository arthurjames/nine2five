package storage

import (
	"time"

	cfg "github.com/arthurjames/kit/config/storage"
	"github.com/arthurjames/kit/storage"
	"github.com/arthurjames/nine2five"
)

// Client represents a client to the underlying database.
type Client struct {

	// Returns the current time.
	Now func() time.Time

	// Services
	timesheetService TimesheetService

	db *storage.Storage
}

func NewClient() *Client {
	c := &Client{Now: time.Now}
	c.timesheetService.client = c
	return c
}

func (c *Client) Open(config *cfg.StorageConfig) error {
	// Open database
	db, err := storage.NewStorage(config)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *Client) Close() error {
	if open, _ := c.db.IsOpen(); open {
		c.db.Close()
	}
	return nil
}

// TimesheetService returns the timesheet service associated with the client.
func (c *Client) TimesheetService() nine2five.TimesheetService {
	return &c.timesheetService
}
