package dateutils

import (
	"errors"
	"fmt"
	"time"
)

const (
	// dateParser is a basic date parser for common european date formats
	dateParser = "02/01/2006"
)

// This package provides utilities to calculate how many times a specified day
// occurs on the first of the month, within a defined time interval.
//
// Other helper functions provide a means of calculating this per century (for
// a specified day), or for Sundays in the 20th century.
//
// Further functionality could be added if required.

// countDaysOnFirst counts the number of times (within a provided range)
// the specified day occurs on the first of the month.
func countDaysOnFirst(startDate, endDate time.Time, day time.Weekday) (int, error) {

	// Ensure that startDate is before endDate
	if endDate.Before(startDate) {
		return 0, errors.New("Provided endDate is prior to the startDate")
	}

	// Ensure that the startDate is the 1st of the month
	if startDate.Day() != 1 {
		return 0, fmt.Errorf("Provided startDate is not on the 1st of the month. Got %d", startDate.Day())
	}

	// Keep track of the number
	count := 0
	for {

		// Check that startDate is beforeEndDate
		if startDate.Before(endDate) {

			// Check whether the first of the month is as expected
			if startDate.Weekday() == day {
				count++
			}

			// Increment date by a month
			startDate = startDate.AddDate(0, 1, 0)

		} else {
			break
		}
	}

	return count, nil
}

// CountDaysOnFirstInCentury counts the number of times within the provided
// century, the specified day occurs on the first of the month.
func CountDaysOnFirstInCentury(century int, day time.Weekday) (int, error) {

	// Set initial dates starting at year 1
	startDate, _ := time.Parse(dateParser, "01/01/0001")
	endDate, _ := time.Parse(dateParser, "31/12/0000")

	// Calculate the full year and add add to start/endDate
	startDate = startDate.AddDate((century-1)*100, 0, 0)
	endDate = endDate.AddDate(century*100, 0, 0)

	// Call lower level function
	return countDaysOnFirst(startDate, endDate, day)
}

// CountSundaysOnFirstIn20thCentury counts the number of times within the 20th
// century, Sunday occurs on the first of the month.
func CountSundaysOnFirstIn20thCentury() (int, error) {
	return CountDaysOnFirstInCentury(20, time.Sunday)
}
