package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func main() {
	arg := flag.String("time", "", "original time")
	flag.Parse()
	timeArgs := flag.Args()
	if len(timeArgs) != 0 {
		panic("Use command like this: go run time_format.go -time=03:04:05am \n")
	}

	timeStr := *arg

	milTime, err := FormatTime(timeStr)
	if err != nil {
		panic(fmt.Errorf("error: %w, use command like this: go run time_format.go -time=03:04:05am", err))
	}

	fmt.Printf("%s\n", milTime)
}

// Func is exportable because we suppose that somebody will use it in their project as library.
func FormatTime(timeStr string) (string, error) {

	timeStr = strings.ToUpper(timeStr)

	initTime, err := time.Parse("03:04:05PM", timeStr)

	if err != nil {
		return "", fmt.Errorf("unable to parse time string: %w", err)
	}

	return initTime.Format(time.TimeOnly), nil
}
