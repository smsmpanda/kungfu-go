package examples

import (
	"fmt"
	"time"
)

func main2() {
	fmt.Println("date")
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Day())

	fmt.Println(t.Month())
	fmt.Println(t.Minute())
	fmt.Println(t.Year())

	fmt.Printf("%02d.%02d", 1, 32)
	fmt.Println()
	fmt.Printf("%2d.%02d", 1, 32)
}
