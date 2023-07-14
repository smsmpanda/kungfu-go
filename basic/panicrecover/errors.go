package panicrecover

import (
	"errors"
	"fmt"
	"math"
)

var errNotFound error = errors.New("Not found error")

func Gorunerror() {
	fmt.Printf("error: %v\n", errNotFound)

	fmt.Println(Sqrt(64))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	return math.Sqrt(f), nil
}
