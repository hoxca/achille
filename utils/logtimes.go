package utils

import (
	"fmt"
	"time"
)

var (
	// MyLocation define time zone for logtimes package
	MyLocation string
	// Delay is the constant time in seconds before declaring last log line too old to process
	Delay int
)

// IsFreshlog verify that logline is less that AssumeFresh seconds old
func IsFreshlog(logTime string) (bool, int) {

	const (
		logTimeLayout = "2006-01-02 03:04:05 PM"
		timeLayout    = "2006-01-02 15:04:05"
	)
	var fresh = false

	if len(logTime) == 21 {
		t := logTime
		index := 11
		logTime = t[:index] + "0" + t[index:]
	}

	location, err := time.LoadLocation(MyLocation)
	if err != nil {
		fmt.Println(err)
	}

	parsedTime, _ := time.ParseInLocation(logTimeLayout, logTime, location)
	diffTime := time.Since(parsedTime)
	seconds := int(diffTime.Seconds())
	// fmt.Printf("Take off in %f seconds.", seconds)

	if seconds < Delay {
		fresh = true
	}

	return fresh, seconds
}
