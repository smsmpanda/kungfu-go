package performence

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func GorunPerfString() {
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)

	fmt.Println(str[3:4])

	str2 := "hello world"
	for i := 0; i < len(str2); i++ {
		fmt.Println(str2[i])
	}
	fmt.Println("--------2-------")
	for _, v := range str2 {
		fmt.Println(v)
	}
	fmt.Println("--------3-------")
	fmt.Println(utf8.RuneCountInString(str2))
	fmt.Println(len([]rune(str2)))
	fmt.Println("--------4-------")
	strs := "shi"
	strl := "li"
	strj := "juan"
	strR := strings.Join([]string{strs, strl, strj}, "2")
	fmt.Println(strR)

	var wt bytes.Buffer
	wt.WriteString("Hellow")
	wt.WriteString("world")
	wt.WriteString("!")
	fmt.Println(wt.String())
}
