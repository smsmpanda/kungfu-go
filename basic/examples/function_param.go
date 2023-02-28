package examples

import (
	"fmt"
)

func Goingfunc() {
	callback(10, Add)
	fmt.Println("go-algorithm", len("go-algorithm"))
	fmt.Println("go算法", len("go算法"))

	fc := func(s string) {
		for k, v := range s {
			fmt.Println(k, v)
		}
	}
	fc("abc")
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
