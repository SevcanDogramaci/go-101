package main

import "fmt"

// 1) IWD event variables
var eventName string = "IWD Konya 2022"
var eventLocation = "KFAU"
var eventDay string

func main() {
	eventDay = "Saturday"

	// 2) Other variable types
	eventStartHour := 13
	var eventEndHour uint = 16

	fmt.Printf(`
	Event info:
	Name: %s
	Location: %s
	Day: %s
	Hours: %d-%d
	`, eventName, eventLocation, eventDay, eventStartHour, eventEndHour)

	// 3) Get attendee info by function
	fmt.Println(getAttendeeInfo("Sevcan", "Doğramacı", "KFAU", 3, true))
	if attendeeInfo, err := getAttendeeInfoWithError("", "Onur", "NEU", 4, false); err != nil {
		fmt.Println("An error occurred:", err.Error())
	} else {
		fmt.Println(attendeeInfo)
	}

	// 4) Attendee type
	IWDAttendee1 := Attendee{
		name:     "Mehmet",
		surname:  "Söyler",
		school:   "SEU",
		class:    2,
		isFromCS: true,
	}
	fmt.Println(IWDAttendee1.getInfo())

	// 5) Multiple attendees
	IWDAttendees := [2]Attendee{IWDAttendee1, {
		name:     "Fırat",
		surname:  "Onur",
		school:   "KTO Karatay",
		class:    1,
		isFromCS: false,
	}}

	// 6) Print all (standard | range)
	for i, attendee := range IWDAttendees {
		fmt.Println((i + 1), ". attendee:", attendee)
	}

	// 7) Upcoming event algoHack, attendee number can increase
	// 7.1) Only the first attendee of IWD is attending currently
	algoHackAttendees := []Attendee{IWDAttendee1}

	for i, attendee := range algoHackAttendees {
		fmt.Println((i + 1), ". attendee:", attendee)
	}

	// 7.2) New attendee registered
	algoHackAttendees = append(algoHackAttendees, Attendee{
		name:     "Hatice",
		surname:  "Oytun",
		school:   "KFAU",
		class:    4,
		isFromCS: true,
	})

	for i, attendee := range algoHackAttendees {
		fmt.Println((i + 1), ". attendee:", attendee)
	}

	// 8) Keep events in a phone no book like structure
	allAttendees := map[string][]Attendee{
		"IWD Konya 2022": IWDAttendees[:],
		"AlgoHack":       algoHackAttendees,
	}
	for eventName, attendees := range allAttendees {
		fmt.Println(eventName, " attendees:", attendees)
	}
}
