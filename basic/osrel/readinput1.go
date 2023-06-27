package osrel

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "69.12 / 5121 / Go"
	format                 = "%f / %d / %s"
)

func Gorun21312() {
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	//从字符串中获取值
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read:", f, i, s)
}
