package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/arthurjames/nine2five"
	"github.com/arthurjames/nine2five/storage/internal"
)

// Ensure Timesheet can be marshaled and unmarshaled.
func TestMarshalTimesheet(t *testing.T) {

	v := nine2five.Timesheet{
		ID:         1,
		CustomerID: 2,
		Month:      3,
		Year:       2017,
		Locked:     true,
		Entries:    []nine2five.EntryID{1, 2, 3, 4},
		Created:    time.Now().UTC(),
		Updated:    time.Now().UTC(),
	}

	var other nine2five.Timesheet

	if buf, err := internal.MarshalTimesheet(&v); err != nil {
		t.Fatal(err)
	} else if err := internal.UnMarshalTimesheet(buf, &other); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(v, other) {
		t.Fatalf("unexpected copy: %#v - %#v", v, other)
	}
}
