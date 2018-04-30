package storage

import (
	"fmt"

	"github.com/arthurjames/nine2five"
)

// Ensure TimesheetService implements nine2five.TimesheetService.
var _ nine2five.TimesheetService = &TimesheetService{}

// TimesheetService represents a PostgreSQL implementation of nine2five.TimesheetService.
type TimesheetService struct {
	client *Client
}

// GetTimesheet returns a timesheet for a given id.
func (s *TimesheetService) Timesheet(id nine2five.TimesheetID) (*nine2five.Timesheet, error) {
	// Start read-only transaction.
	tx, err := s.client.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Prepare statement.
	sql := `SELECT id, customer_id, month, year, locked 
                FROM timesheets_vw 
                WHERE id = $1`
	stmt, err := tx.Prepare(sql)

	// Defer close stmt.
	defer stmt.Close()

	// Execute statement.
	var t nine2five.Timesheet
	err = stmt.QueryRow(id).Scan(&t.ID, &t.CustomerID, &t.Month, &t.Year, &t.Locked)

	return &t, nil
}

// CreateTimesheet creates a new timesheet if all necessary fields have proper values.
func (s *TimesheetService) CreateTimesheet(t *nine2five.Timesheet) (*nine2five.Timesheet, error) {
	// Start read-write transaction.
	tx, err := s.client.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Prepare statement.
	sql := `INSERT INTO timesheets (date, customer_id) 
                VALUES (to_date($1, 'YYYY-MM-DD'), $2) 
                RETURNING ID`
	stmt, err := tx.Prepare(sql)

	// Defer close stmt.
	defer stmt.Close()

	// Execute statement.
	err = stmt.QueryRow(
		fmt.Sprintf("%4d-%02d-01", t.Year, t.Month),
		t.CustomerID).Scan(&t.ID)

	return t, nil
}
