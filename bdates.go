package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// isLeapYear checks if a given year is a leap year.
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func main() {
	// Define a flag for the timezone with a default value of "UTC"
	tz := flag.String("tz", "UTC", "Timezone for the date calculations (e.g., 'Europe/Berlin')")

	// Define usage function to display help text
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "%s -tz=\"Timezone\" StartDate EndDate\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Example: %s -tz=\"Europe/Berlin\" 1930-04-03 2017-04-03\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Start flag parsing
	flag.Parse()

	// Check if two arguments were passed after the flags
	if len(flag.Args()) != 2 {
		flag.Usage()
		return
	}

	startDate := flag.Arg(0)
	endDate := flag.Arg(1)

	// Define the layout that matches the input format
	layout := "2006-01-02"

	// Load the timezone from the flag
	loc, err := time.LoadLocation(*tz)
	if err != nil {
		fmt.Printf("Error loading timezone: %v\n", err)
		os.Exit(1)
	}

	// Parse the start and end dates as time.Time
	start, err := time.ParseInLocation(layout, startDate, loc)
	if err != nil {
		fmt.Printf("Error parsing the start date: %v\n", err)
		os.Exit(1)
	}

	end, err := time.ParseInLocation(layout, endDate, loc)
	if err != nil {
		fmt.Printf("Error parsing the end date: %v\n", err)
		os.Exit(1)
	}

	// Get the current time in the specified timezone
	now := time.Now().In(loc)

	// Check if the start date is February 29th
	isStartDateLeapDay := start.Month() == time.February && start.Day() == 29

	// Loop from the start date to the end date
	for d := start; d.Year() <= end.Year(); d = d.AddDate(1, 0, 0) {
		// If the start date is February 29th, only print for leap years
		if isStartDateLeapDay && isLeapYear(d.Year()) {
			verb := "was"
			// Create a new date for the current year to compare with 'now'
			currentYearBirthday := time.Date(d.Year(), start.Month(), start.Day(), 0, 0, 0, 0, loc)
			if currentYearBirthday.After(now) || currentYearBirthday.Equal(now) {
				verb = "is"
			}
			fmt.Printf("The birthday in the year %d %s on a %s\n", d.Year(), verb, d.Weekday())
		} else if !isStartDateLeapDay {
			// If it's not a leap year, print the date for the same day and month as the start date
			birthday := time.Date(d.Year(), start.Month(), start.Day(), 0, 0, 0, 0, loc)
			verb := "was"
			if birthday.After(now) || birthday.Equal(now) {
				verb = "is"
			}
			fmt.Printf("The birthday in the year %d %s on a %s\n", birthday.Year(), verb, birthday.Weekday())
		}
	}
}
