package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

	timeString, err := getFirstArgument()

	if err != nil {
		printUsage()
		os.Exit(1)
	}

	timestamp, err := parseTime(timeString, timeLayouts)

	if err != nil {
		log.Fatalf(err.Error())
	}

	diff := time.Now().Sub(timestamp).Round(time.Second)
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
