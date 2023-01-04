package main

import (
	"fmt"
	"math"
)

func main() {
	// a, c := Uint8FromInt(256)
	// fmt.Println(a)
	// fmt.Print(c)

	e := IntFromFloat64(123.412415125)
	fmt.Println(e)
}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, faction := math.Modf(x)
		if faction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the init32 range", x))
}
