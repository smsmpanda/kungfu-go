package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"time"
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

// a:=1(称为初始化声明，且只能被用于函数体内)

type MInt int

func main() {
	// a, c := Uint8FromInt(256)
	// fmt.Println(a)
	// fmt.Print(c)

	//e := IntFromFloat64(123.412415125)
	//fmt.Println(e)
	var a int = 3
	var b int16 = 12
	var r int = a + int(b)
	fmt.Print(r)

	var c1 complex64 = 5 + 10i
	var c2 complex64 = 5 + 10i
	fmt.Printf("The value is %v", c1+c2)
	fmt.Println(real(c1), imag(c1))

	fmt.Println(1 & 1)
	fmt.Println(1 & 2)
	fmt.Println(1 & 3)

	type ByteSize float64
	const (
		_           = iota // 通过赋值给空白标识符来忽略值
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}

	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%3.4f / ", 100*rand.Float32())
	}

	var r1, r2 MInt = 3, 4
	r3 := r1 + r2
	fmt.Printf("r3 has the value：%d\n ", r3)

	fmt.Println(HOME)
	fmt.Println(GOROOT)
	fmt.Println(USER)

	var goos string = runtime.GOOS
	fmt.Printf("The opeating system is: %s\n", goos)
	path := os.Getenv("path")
	fmt.Printf("Path is %s\n", path)

	var jr string = "abc"
	fmt.Println("hello, world", jr)

	var q, w, t int
	q, w, t = 5, 5, 1
	fmt.Print(q, w, t, "\n")
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
