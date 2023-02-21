package examples

import (
	"fmt"
)

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func main4() {
	for i := 1; i <= 25; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

	str := "G"
	for i := 0; i < 25; i++ {
		fmt.Println(str)
		str += "G"
	}
	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}

	for i := 1; i <= 100; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Printf(" %d  => 3 & 5\n", i)
		case i%3 == 0:
			fmt.Printf(" %d  => 3\n", i)
		case i%5 == 0:
			fmt.Printf(" %d  => 5\n", i)
		default:
			fmt.Println(" ~ ")
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("*^_^")
		}
		fmt.Println("^-^")
	}
}
