package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/arthurjames/nine2five"
	"github.com/arthurjames/nine2five/storage/internal"
)

func TestMarshalEntry(t *testing.T) {

	v := nine2five.Entry{
		ID:          1,
		Description: "test description",
		Start:       time.Now().UTC(),
		End:         time.Now().UTC(),
		Created:     time.Now().UTC(),
		Deleted:     time.Now().UTC(),
		Updated:     time.Now().UTC(),
	}

	var other nine2five.Entry

	if buf, err := internal.MarshalEntry(&v); err != nil {
		t.Fatal(err)
	} else if err := internal.UnMarshalEntry(buf, &other); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(v, other) {
		t.Fatalf("unexpected copy: %#v", other)
	}
}
