package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    return t.Before(time.Now()) 
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    const afternoonStartsHour = 12
    const afternoonEndsHour = 18
    if t.Hour() >= afternoonStartsHour && t.Hour() < afternoonEndsHour {
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    msg := fmt.Sprintf("You have an appointment on %s, %s.",
                       t.Weekday(),
                       t.Format("January 2, 2006, at 15:04"))
    return msg
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(
        time.Now().Year(), 9, 15,
        0, 0, 0, 0,
        time.UTC)
}
