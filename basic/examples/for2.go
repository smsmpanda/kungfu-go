package examples

import "fmt"

func main5() {
	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now:%d\n", i)
	}

	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))

	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	str2 := "Chinese: 日本語"
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
