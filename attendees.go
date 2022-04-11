package main

import (
	"errors"
	"fmt"
)

// 3)
func getAttendeeInfo(attendeeName string,
	attendeeSurname string,
	attendeeSchool string,
	attendeeClass uint,
	isFromComputerScience bool) string {
	return fmt.Sprintf(`
	Attendee Info:
	Name: %s
	Surname: %s
	School: %s
	Class: %d
	FromCS: %v
	`, attendeeName,
		attendeeSurname,
		attendeeSchool,
		attendeeClass,
		isFromComputerScience)
}

// 3.1)
func getAttendeeInfoWithError(attendeeName string,
	attendeeSurname string,
	attendeeSchool string,
	attendeeClass uint,
	isFromComputerScience bool) (string, error) {

	if attendeeName == "" || attendeeSurname == "" {
		return "", errors.New("name or surname empty")
	}

	return fmt.Sprintf(`
	Attendee Info:
	Name: %s
	Surname: %s
	School: %s
	Class: %d
	FromCS: %v
	`, attendeeName,
		attendeeSurname,
		attendeeSchool,
		attendeeClass,
		isFromComputerScience), nil
}

// 4
// like class
type Attendee struct {
	name     string
	surname  string
	school   string
	class    uint
	isFromCS bool
}

func (attendee Attendee) getInfo() string { // like class method
	return fmt.Sprintf(`
	Attendee Info:
	Name: %s
	Surname: %s
	School: %s
	Class: %d
	FromCS: %v
	`, attendee.name,
		attendee.surname,
		attendee.school,
		attendee.class,
		attendee.isFromCS)
}
