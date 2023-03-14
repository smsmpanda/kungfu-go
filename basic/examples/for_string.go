package examples

import "fmt"

func Going123fssd() {
	s := "\u00ff\u754c"

	for i, c := range s {
		fmt.Printf("%d:%c", i, c)
	}

	var b []byte
	var s1 string = "abc,da,32,df"
	b = append(b, s1...)

	var s2 string = "hello"
	var t = s2[2:4]
	t = "a,3"
	fmt.Println(s2, t)

	s4 := "hello"
	c := []byte(s4)
	c[0] = 'c'
	fmt.Println(c)
}
