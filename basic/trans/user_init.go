package trans

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
}

func Myfunc() {
	fmt.Println("The another golang package.")
}
