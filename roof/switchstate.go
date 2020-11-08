package roof

import (
	"fmt"
	"strconv"
)

// Switch structure
type Switch struct {
	value Bits
}

//Bits to store switch flags
type Bits uint16

const (
	powerDown Bits = 1 << iota
	weatherUnsafe
	mountParked
	openSensor
	closeSensor
	openButton
	stopButton
	unused1
	closeButton
	directCommand
	management
	homeRa
	homeDec
	conditions
	unused2
	unused3
)

var switchToString = map[Bits]string{
	powerDown:     "Power is down",
	weatherUnsafe: "Unsafe weather alert",
	mountParked:   "Mount switches at park position",
	openSensor:    "Sensor confirm that roof is open",
	closeSensor:   "Sensor confirm that roof is close",
	openButton:    "Open button",
	stopButton:    "Stop button",
	unused1:       "",
	closeButton:   "Close button",
	directCommand: "Direct command is triggered",
	management:    "Management command triggered",
	homeRa:        "Home RA set",
	homeDec:       "Home DEC set",
	conditions:    "Conditions is disabled",
	unused2:       "",
	unused3:       "",
}

func (b Bits) String() string {
	return switchToString[b]
}

func has(b, flag Bits) bool {
	return b&flag != 0
}

// PrintSwitchesStatus echo the switch values
func (sw Switch) PrintSwitchesStatus() {
	for i, flag := range []Bits{
		powerDown,
		weatherUnsafe,
		mountParked,
		openSensor,
		closeSensor,
		openButton,
		stopButton,
		unused1,
		closeButton,
		directCommand,
		management,
		homeRa,
		homeDec,
		conditions,
		unused2,
		unused3} {
		if has(sw.value, flag) {
			fmt.Println(flag.String())
		}
		i++
		// fmt.Println(i, has(sw.value, flag))
	}
}

// Set Switch value from string value
func (sw Switch) Set(s string) *Switch {
	v, _ := strconv.ParseInt(s, 16, 64)
	sw.value = Bits(v)
	return &sw
}

// IsMountAtPark return true if MAP
func (sw Switch) IsMountAtPark() bool {
	if has(sw.value, mountParked) {
		return true
	}
	return false
}

// IsSystemOnPower return false in case of power failure
func (sw Switch) IsSystemOnPower() bool {
	if has(sw.value, powerDown) {
		return false
	}
	return true
}

// IsWeatherSafe return false if unsafe weather is detected
func (sw Switch) IsWeatherSafe() bool {
	if has(sw.value, weatherUnsafe) {
		return false
	}
	return true
}

// IsRoofOpen return true if sensor confirm that roof is open
func (sw Switch) IsRoofOpen() bool {
	if has(sw.value, openSensor) {
		return true
	}
	return false
}

// IsRoofClose return true if sensor confirm that roof is close
func (sw Switch) IsRoofClose() bool {
	if has(sw.value, closeSensor) {
		return true
	}
	return false
}
