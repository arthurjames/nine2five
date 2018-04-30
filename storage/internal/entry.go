package internal

import (
	"time"

	"github.com/arthurjames/nine2five"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

//go:generate protoc --go_out=plugins=grpc:. internal.proto

// MarshalEntry encodes a timesheet entry to binary format.
func MarshalEntry(e *nine2five.Entry) ([]byte, error) {

	var start, end, created, deleted, updated *types.Timestamp
	var err error

	if start, err = types.TimestampProto(e.Start); err != nil {
		return []byte{}, err
	}
	if end, err = types.TimestampProto(e.End); err != nil {
		return []byte{}, err
	}
	if created, err = types.TimestampProto(e.Created); err != nil {
		return []byte{}, err
	}
	if deleted, err = types.TimestampProto(e.Deleted); err != nil {
		return []byte{}, err
	}
	if updated, err = types.TimestampProto(e.Updated); err != nil {
		return []byte{}, err
	}

	return proto.Marshal(&Entry{
		ID:          int32(e.ID),
		Description: e.Description,
		Start:       start,
		End:         end,
		Created:     created,
		Deleted:     deleted,
		Updated:     updated,
	})
}

// UnMarshalEntry decodes a timesheet entry from binary data.
func UnMarshalEntry(data []byte, e *nine2five.Entry) error {
	var pb Entry
	var err error

	if err = proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	var start, end, created, deleted, updated time.Time
	if start, err = types.TimestampFromProto(pb.Start); err != nil {
		return err
	}
	if end, err = types.TimestampFromProto(pb.End); err != nil {
		return err
	}
	if created, err = types.TimestampFromProto(pb.Created); err != nil {
		return err
	}
	if deleted, err = types.TimestampFromProto(pb.Deleted); err != nil {
		return err
	}
	if updated, err = types.TimestampFromProto(pb.Updated); err != nil {
		return err
	}

	e.ID = nine2five.EntryID(pb.GetID())
	e.Description = pb.GetDescription()
	e.Start = start
	e.End = end
	e.Created = created
	e.Deleted = deleted
	e.Updated = updated

	return nil
}
