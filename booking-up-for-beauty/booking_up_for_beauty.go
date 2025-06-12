package booking

import (
	//"fmt"
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layouts := []string{
		"1/2/2006",
		"1/2/2006 15:04:05",
		"1/02/2006 15:04:05",
		"01/02/2006 15:04:05",
		"January 2, 2006 15:04:05",
		"Monday, January 2, 2006 15:04:05",
	} //patterns of time: Mon Jan 2 15:04:05 -0700 MST 2006
	for _, layout := range layouts {
		t, err := time.Parse(layout, date)
		if err == nil {
			return t
		}
	}
	return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	app := Schedule(date)
	return now.After(app)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	app := Schedule(date)
	if app.Hour() >= 12 && app.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	app := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, at %d:%d.", app.Format("Monday, January 2, 2006"), app.Hour(), app.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.a
// Assuming year is 2067
// 2067-09-15 00:00:00 +0000 UTC
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	date := fmt.Sprintf("09/15/%d", year)
	return Schedule(date)
}
