package methods

import (
	"fmt"
	"strconv"
)

type Day int

const (
	Mo Day = iota
	Tu
	We
	Th
	Fr
	Sa
	Su
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

type Celsius float64

func (c Celsius) String() string {
	return "The temperature is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"
}

func Gorungdfg() {
	var c Celsius = 18.36
	fmt.Println(c)

	var th Day = 3
	fmt.Printf("The 3rd day is: %s\n", th)
	// If index > 6: panic: runtime error: index out of range
	// but use the enumerated type to work with valid values:
	var day = Su
	fmt.Println(day) // prints Sunday
	fmt.Println(0, Mo, 1, Tu)

	fmt.Println(EST) // Print* knows about method String() of type TZ
	fmt.Println(0 * HOUR)
	fmt.Println(-6 * HOUR)
}

type TZ int

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

var timeZones = map[TZ]string{
	UTC: "Universal Greenwich time",
	EST: "Eastern Standard time",
	CST: "Central Standard time"}

func (tz TZ) String() string { // Method on TZ (not ptr)
	if zone, ok := timeZones[tz]; ok {
		return zone
	}
	return ""
}
