package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const layoutDateTime = "2006-01-02 15:04:05"

func main() {
	var diffInHours bool
	flag.BoolVar(&diffInHours, "h", false, "diff in hours")
	flag.Parse()

	timeString := flag.Arg(0)
	timestamp, err := time.ParseInLocation(layoutDateTime, timeString, time.Local)

	if err != nil {
		log.Fatalf(err.Error())
	}

	now := time.Now()
	diff := now.Sub(timestamp).Round(time.Second)
	var diffString string

	if diffInHours {
		diffString = fmt.Sprintf("%.2f", diff.Hours()) + "h"
	} else {
		diffString = diff.String()
	}

	fmt.Println(diffString)
}
