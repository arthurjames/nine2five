package nine2five

// Timestamp errors
const (
	ErrTimesheetIDRequired = Error("Timesheet ID required")
	ErrTimesheetNotFound   = Error("Timesheet not found")
)

// Error represents a nine2five error.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}
