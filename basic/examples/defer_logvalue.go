package examples

import (
	"fmt"
	"io"
	"log"
)

func Goingde() {
	func1("Go")
}

func func1(s string) (n int, err error) {
	fmt.Println(1)
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	fmt.Println(2)
	return 7, io.EOF
}
