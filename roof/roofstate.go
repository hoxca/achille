package roof

import (
	"fmt"
	"strconv"
)

// Status structure
type Status struct {
	state int
	ratio float64
}

// State of roof
type state int

// State enum
const (
	Open state = iota
	Close
	Opening
	Closing
	Error
)

var statusToString = map[state]string{
	Open:    "Open",
	Close:   "Close",
	Opening: "Opening",
	Closing: "Closing",
	Error:   "Error",
}

func (stat state) String() string {
	return statusToString[stat]
}

// Set roof structure
func (r Status) Set(logstat string, logratio string) *Status {
	stat, _ := strconv.Atoi(logstat)
	ratio, _ := strconv.ParseFloat(logratio, 32)
	r.state = stat
	r.ratio = ratio
	return &r
}

// PrintStatus print the roof status
func (r Status) PrintStatus() {
	var status = state(r.state)
	fmt.Printf("Roof status is %s @%0.2f%%\n", status.String(), r.ratio)
}

// IsClose roof
func (r Status) IsClose() bool {
	if r.state == 1 {
		return true
	}
	return false
}

// IsOpen roof
func (r Status) IsOpen() bool {
	if r.state == 0 {
		return true
	}
	return false
}

// IsFullyOpen roof
func (r Status) IsFullyOpen() bool {
	if r.state == 0 && r.ratio == 100 {
		return true
	}
	return false
}
