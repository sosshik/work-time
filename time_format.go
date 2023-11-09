package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	timeStr := os.Args[1]

	milTime, err := FormatTime(timeStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s", milTime)
}

func FormatTime(timeStr string) (string, error) {

	if len(timeStr) != 10 {
		return "", errors.New("wrong time format")
	}

	timeStr = timeStr[:8] + strings.ToUpper(timeStr[8:])

	initTime, err := time.Parse("03:04:05PM", timeStr)

	if err != nil {
		return "", fmt.Errorf("unable to parse time string: %v", err)
	}

	return initTime.Format(time.TimeOnly), nil
}
