package internal

import (
	"time"

	"github.com/arthurjames/nine2five"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

//go:generate protoc --go_out=plugins=grpc:. internal.proto

// MarshalTimesheet encodes a timesheet to binary format.
func MarshalTimesheet(ts *nine2five.Timesheet) ([]byte, error) {

	var created, updated *types.Timestamp
	var err error

	if created, err = types.TimestampProto(ts.Created); err != nil {
		return []byte{}, err
	}

	if updated, err = types.TimestampProto(ts.Updated); err != nil {
		return []byte{}, err
	}

	var entries []int32
	for _, e := range ts.Entries {
		entries = append(entries, int32(e))
	}

	return proto.Marshal(&Timesheet{
		ID:         int32(ts.ID),
		CustomerID: int32(ts.CustomerID),
		Month:      int32(ts.Month),
		Year:       int32(ts.Year),
		Locked:     ts.Locked,
		Entries:    entries,
		Created:    created,
		Updated:    updated,
	})
}

// UnMarshalTimesheet decodes a timesheet from binary data.
func UnMarshalTimesheet(data []byte, ts *nine2five.Timesheet) error {
	var pb Timesheet
	var err error

	if err = proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	var created, updated time.Time
	if created, err = types.TimestampFromProto(pb.Created); err != nil {
		return err
	}

	if updated, err = types.TimestampFromProto(pb.Updated); err != nil {
		return err
	}

	var entries []nine2five.EntryID
	for _, e := range pb.GetEntries() {
		entries = append(entries, nine2five.EntryID(e))
	}

	ts.ID = nine2five.TimesheetID(pb.GetID())
	ts.CustomerID = nine2five.CustomerID(pb.GetCustomerID())
	ts.Month = int(pb.GetMonth())
	ts.Year = int(pb.GetYear())
	ts.Entries = entries
	ts.Locked = pb.GetLocked()
	ts.Created = created
	ts.Updated = updated

	return nil
}
