package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) > 2 {
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

// Func is exportable because we suppose that somebody will use it in th own project as library.
func FormatTime(timeStr string) (string, error) {
	// This check is necessary because it prevents panic in case of shorter input then expected.
	// If I will get rid of it, then I will have to get rid of lowercase am/pm handling.
	if len(timeStr) != 10 {
		return "", errors.New("wrong time format, expected: 03:04:05PM")
	}

	timeStr = timeStr[:8] + strings.ToUpper(timeStr[8:])

	initTime, err := time.Parse("03:04:05PM", timeStr)

	if err != nil {
		return "", fmt.Errorf("unable to parse time string: %v", err)
	}

	return initTime.Format(time.TimeOnly), nil
}
