package examples

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func Goingf() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
	}
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))

	ffffffffff := fib()
	// Function calls are evaluated left-to-right.
	// println(f(), f(), f(), f(), f())

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	for i := 0; i <= 9; i++ {
		println(i+2, ffffffffff())
	}
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func makeAddSuffix(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}
