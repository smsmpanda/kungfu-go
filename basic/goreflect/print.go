package goreflect

import (
	"os"
	"strconv"
)

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + "℃"
}

type Day int

var dayName = []string{"Monday", "TuesDay", "WednesDay", "ThursDay", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func print(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch a := arg.(type) {
		case Stringer:
			os.Stdout.WriteString(a.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
		default:
			os.Stdout.WriteString("???")
		}
	}
}

func Gorunaafd() {
	print(Day(1), "was", Celsius(18.36)) // Tuesday was 18.4 °C
}
