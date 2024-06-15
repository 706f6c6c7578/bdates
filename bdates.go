package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

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

	// Get the current time
	now := time.Now()

	// Loop from the start date to the end date
	for d := start; d.Before(end) || d.Equal(end); d = d.AddDate(1, 0, 0) {
		// Check if the year is in the future
		verb := "was"
		if d.After(now) {
			verb = "is"
		}

		// Output the weekday for the current date
		fmt.Printf("The birthday in the year %d %s on a %s\n", d.Year(), verb, d.Weekday())
	}
}
