package nine2five

import "time"

// TimesheetID represents a timesheet identifier.
type TimesheetID int

// CustomerID represents a customer identifier
type CustomerID int

// Timesheet represents a timesheet for a specific customer.
type Timesheet struct {
	ID         TimesheetID `json:"id"`
	CustomerID CustomerID  `json:"customerID"`
	Month      int         `json:"month"`
	Year       int         `json:"year"`
	Locked     bool        `json:"locked"`
	Entries    []EntryID   `json:"entries"`
	Created    time.Time   `json:"created"`
	Updated    time.Time   `json:"updated"`
}

// TimesheetService represents a service for managing timesheets.
type TimesheetService interface {
	Timesheet(id TimesheetID) (*Timesheet, error)
	//	GetTimesheets() ([]*Timesheet, error)
	CreateTimesheet(t *Timesheet) (*Timesheet, error)
	//	DeleteTimesheet(id int) error
	//	SetLocked(id TimesheetID, locked bool) error
	//	AddEntry(id TimesheetID, e *Entry) error
	//	DeleteEntry(id EntryID) error
	//	UpdateEntry(e *Entry) error
}
