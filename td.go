package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var dateTimeLayouts = []string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02 15",
}

var dateLayouts = []string{
	"2006-01-02",
}

type If bool

func (c If) String(a, b string) string {
	if c {
		return a
	}
	return b
}

func main() {
	var diffInHours, diffInDays, diffInMinutes, diffInSeconds bool
	flag.BoolVar(&diffInSeconds, "s", false, "diff in seconds")
	flag.BoolVar(&diffInMinutes, "m", false, "diff in minutes")
	flag.BoolVar(&diffInHours, "h", false, "diff in hours")
	flag.BoolVar(&diffInDays, "d", false, "diff in days")
	flag.Parse()

	timeString, err := getFirstArgument()

	if err != nil {
		printUsage()
		os.Exit(1)
	}

	hasTime := true
	timestamp, err := parseTime(timeString, dateTimeLayouts)

	if err != nil {
		hasTime = false
		timestamp, err = parseTime(timeString, dateLayouts)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	now := time.Now().Local()
	var output string

	if diffInSeconds {
		output = fmt.Sprintf("%.2f", now.Sub(timestamp).Round(time.Second).Seconds()) + " seconds"
	} else if diffInMinutes {
		output = fmt.Sprintf("%.2f", now.Sub(timestamp).Round(time.Second).Minutes()) + " minutes"
	} else if diffInHours {
		output = fmt.Sprintf("%.2f", now.Sub(timestamp).Round(time.Second).Hours()) + " hours"
	} else if diffInDays {
		output = fmt.Sprintf("%.2f", now.Sub(timestamp).Round(time.Second).Hours()/24) + " days"
	} else {
		year, month, day, hour, min, sec := diff(now, timestamp)
		output = prettyTime(year, month, day, hour, min, sec, hasTime)
	}

	fmt.Printf(output)
}

func getFirstArgument() (string, error) {
	if isInputFromPipe() {
		bytes, err := io.ReadAll(os.Stdin)
		return strings.TrimSpace(string(bytes)), err
	}

	argument := flag.Arg(0)
	if len(argument) == 0 {
		return "", errors.New("Missing argument")
	}

	return argument, nil
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func printUsage() {
	fmt.Println("Usage: td [flags] [argument]")
	flag.PrintDefaults()
}

func parseTime(timeString string, timeLayouts []string) (timestamp time.Time, error error) {
	for _, layout := range timeLayouts {

		timestamp, err := time.ParseInLocation(layout, timeString, time.Local)
		if err == nil {
			return timestamp, nil
		}
	}
	error = errors.New("Could not parse timestamp")
	return
}

func prettyTime(year, month, day, hour, min, sec int, hasTime bool) string {
	var sb strings.Builder

	if year != 0 {
		sb.WriteString(strconv.Itoa(year))
		yearString := If(year == 1).String("year", "years")
		sb.WriteString(fmt.Sprintf(" %s ", yearString))
	}

	if month != 0 {
		sb.WriteString(strconv.Itoa(month))
		monthString := If(month == 1).String("month", "months")
		sb.WriteString(fmt.Sprintf(" %s ", monthString))
	}

	// Catch edge case for when diffing against current day
	emptyOutputAndNoTime := sb.Len() == 0 && !hasTime
	if day != 0 || emptyOutputAndNoTime {
		sb.WriteString(strconv.Itoa(day))
		dayString := If(day == 1).String("day", "days")
		sb.WriteString(fmt.Sprintf(" %s ", dayString))

	}

	if !hasTime {
		return sb.String()
	}

	if hour != 0 {
		sb.WriteString(strconv.Itoa(hour))
		hourString := If(hour == 1).String("hour", "hours")
		sb.WriteString(fmt.Sprintf(" %s ", hourString))
	}

	if min != 0 {
		sb.WriteString(strconv.Itoa(min))
		minString := If(min == 1).String("minute", "minutes")
		sb.WriteString(fmt.Sprintf(" %s ", minString))
	}

	// Catch edge case for when diffing against current date and time
	if sec != 0 || sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(sec))
		secString := If(sec == 1).String("second", "seconds")
		sb.WriteString(fmt.Sprintf(" %s ", secString))
	}

	return sb.String()
}

// https://stackoverflow.com/questions/36530251/time-since-with-months-and-years/36531443#36531443
func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.Local)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
