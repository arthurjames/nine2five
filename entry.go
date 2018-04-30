package nine2five

import "time"

// EntryID repesents a timesheet-entry identifier
type EntryID int

// Entry represents an entry for a specific timesheet.
type Entry struct {
	ID          EntryID   `json:"id"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Created     time.Time `json:"created"`
	Deleted     time.Time `json:"deleted"`
	Updated     time.Time `json:"updated"`
}
