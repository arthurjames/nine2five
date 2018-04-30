package mock

import (
	"github.com/arthurjames/nine2five"
)

type TimesheetService struct {
	TimesheetFn      func(id nine2five.TimesheetID) (*nine2five.Timesheet, error)
	TimesheetInvoked bool

	//	TimesheetsFn      func() ([]*nine2five.Timesheet, error)
	//	TimesheetsInvoked bool

	CreateTimesheetFn      func(ts *nine2five.Timesheet) error
	CreateTimesheetInvoked bool

	//	DeleteTimesheetFn      func(id nine2five.TimesheetID) error
	//	DeleteTimesheetInvoked bool

	//	SetLockedFn      func(id nine2five.TimesheetID, lock bool) error
	//	SetLockedInvoked bool
}

func (s *TimesheetService) Timesheet(id nine2five.TimesheetID) error {
	s.TimesheetInvoked = true
	return s.TimesheetFn(id)
}

//func (s *TimesheetService) Timesheets() ([]*nine2five.Timesheet, error) {
//	s.TimesheetsInvoked = true
//	return s.TimesheetsFn()
//}
//
func (s *TimesheetService) CreateTimesheet(ts *nine2five.Timesheet) error {
	s.CreateTimesheetInvoked = true
	return s.CreateTimesheetFn(ts)
}

//func (s *TimesheetService) DeleteTimesheet(id nine2five.TimesheetID) error {
//	s.DeleteTimesheetInvoked = true
//	s.DeleteTimesheetFn(id)
//}
//
//func (s *TimesheetService) SetLocked(id nine2five.TimesheetID, lock bool) error {
//	s.SetLockedFn = true
//	return s.SetLockedFn(id, lock)
//}
//
//type CustomerService struct {
//	CustomerFn      func(id nine2five.CustomerID) (*nine2five.Customer, error)
//	CustomerInvoked bool
//
//	CreateCustomerFn      func(cust *Customer) error
//	CreateCustomerInvoked bool
//}
//
//func (s *CustomerService) Customer(id nine2five.CustomerID) (*nine2five.Customer, error) {
//	CustomerInvoked = true
//	return s.CustomerFn(id)
//}
//
//func (s *CustomerService) CreateCustomer(cust nine2five.Customer) error {
//	s.CreateCustomerInvoked = true
//	return s.CreateCustomerFn(cust)
//}
//
