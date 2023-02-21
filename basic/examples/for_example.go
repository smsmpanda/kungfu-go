package examples

import "fmt"

func main3() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Println("")
	}

	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "我是够浪"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}
