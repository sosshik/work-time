package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func main() {
	if len(flag.Args()) > 2 {
		panic("panic: Too many args")
	}
	arg := flag.String("time", "", "original time")
	flag.Parse()
	timeStr := *arg

	if timeStr == "" {
		panic("panic: Missing flag -time\n")
	}

	milTime, err := FormatTime(timeStr)
	if err != nil {
		panic(err)
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
