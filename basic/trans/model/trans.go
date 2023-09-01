package model

import "fmt"

type Size byte

const (
	Big int = 12
	Middium
	Small
)

func GetSize() {
	fmt.Printf("Big: %v\n", Big)
}
