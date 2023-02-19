package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)

	for _, val := range sl {
		fmt.Printf("%s - \n", val)
	}

	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	fmt.Println(strconv.Atoi("234"))
}
