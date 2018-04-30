package storage_test

import (
	"log"
	"reflect"
	"testing"

	cfg "github.com/arthurjames/kit/config/storage"
	"github.com/arthurjames/nine2five"
)

// Ensure timesheet can be created and retrieved.
func TestTimesheetService_CreateTimesheet(t *testing.T) {

	cfg, err := cfg.NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	c := MustOpenClient(cfg)
	defer c.Close()
	s := c.TimesheetService()

	timesheet := &nine2five.Timesheet{
		CustomerID: 1,
		Month:      1,
		Year:       2000,
		Locked:     true,
	}

	// Create new timesheet.
	if timesheet, err = s.CreateTimesheet(timesheet); err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", timesheet)

	// Retrieve timesheet and compare
	other, err := s.Timesheet(timesheet.ID)
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(&timesheet, other) {
		t.Fatalf("Unexpected timesheet: %#v", other)
	}
}
