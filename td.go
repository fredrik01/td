package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const layoutDateTime = "2006-01-02 15:04:05"

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
	timestamp, err := time.ParseInLocation(layoutDateTime, timeString, time.Local)

	if err != nil {
		log.Fatalf("Could not parse first timestamp")
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
