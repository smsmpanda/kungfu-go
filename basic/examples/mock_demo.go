package examples

import (
	"fmt"
)

func Going() {
	a := "helloworld"
	b := make([]byte, 15)
	copy(b, a)

	var c []int = make([]int, 5)
	c = append(c, make([]int, 5)...)

	d := []string{"I", "L", "o", "e", "Y", "o", "u"}
	d = append(d[:3], append([]string{"v"}, d[3:]...)...)

	e, f := split_string("jrslj", 2)
	fmt.Printf("%s , %s\n", e, f)

	fmt.Println(split_string2("jrslj"))

	fmt.Println(string_reverse("google"))

	q29_unique("sweetgirl")
}

func split_string(s string, split int) (s1 string, s2 string) {
	s1 = s[:split]
	s2 = s[split:]
	return
}

func split_string2(str string) string {
	return str[len(str)/2:] + str[:len(str)/2]
}

func string_reverse(str string) string {
	d := []rune(str)
	len := len(d)
	for i := 0; i < len/2; i++ {
		d[i], d[len-1-i] = d[len-1-i], d[i]
	}
	return string(d)
}

func q29_unique(str string) []rune {
	b := make([]rune, 0)
	var c rune
	for _, v := range str {
		if c != v {
			b = append(b, []rune{v}...)
		}
		c = v
	}
	return b
}
