package examples

import "fmt"

var a2 = "G"

func main8() {
	n1()
	m1()
	n1()

	fmt.Println(a)
	//结果：G O O
}

func n1() {
	print(a)
}

func m1() {
	a = "O"
	print(a)
}
