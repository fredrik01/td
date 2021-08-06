package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var timeLayouts = []string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02 15",
	"2006-01-02",
}

func main() {
	var diffInHours, diffInDays bool
	flag.BoolVar(&diffInHours, "h", false, "diff in hours")
	flag.BoolVar(&diffInDays, "d", false, "diff in days")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Usage: td [flags] [argument]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	timeString := flag.Arg(0)
	timestamp, err := parseTime(timeString, timeLayouts)

	if err != nil {
		log.Fatalf(err.Error())
	}

	now := time.Now()
	diff := now.Sub(timestamp).Round(time.Second)
	var diffString string

	if diffInHours {
		diffString = fmt.Sprintf("%.2f", diff.Hours()) + "h"
	} else if diffInDays {
		diffString = fmt.Sprintf("%.2f", diff.Hours()/24) + "d"
	} else {
		diffString = diff.String()
	}

	fmt.Printf(diffString)
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
