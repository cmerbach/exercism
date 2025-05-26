package booking

import "fmt"
import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"

	t, _ := time.Parse(layout,date)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	givenTime, _ := time.Parse(layout, date)
	now := time.Now()

	return givenTime.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	
	appointmentTime, _ := time.Parse(layout, date)
	
	hour := appointmentTime.Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	
	appointmentTime,_ := time.Parse(layout, date)

	formatted := appointmentTime.Format("Monday, January 2, 2006, at 15:04")

	return fmt.Sprintf("You have an appointment on %s.", formatted)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	anniversaryDate := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	
	return anniversaryDate
}
