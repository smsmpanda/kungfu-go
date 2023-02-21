package examples

import (
	"errors"
	"math"
)

func Gofunc12() {

}

// 命名返回值
func MySqrt1(n float64) (ret float64, err error) {
	if n < 0 {
		ret = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		ret = math.Sqrt(n)
	}
	return
}

// 非命名返回值
func MySqrt2(n float64) (float64, error) {
	if n < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	//otherwise use default square root function
	return math.Sqrt(n), nil
}
