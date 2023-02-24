package examples

import "fmt"

func mytrace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func mya() {
	defer un(mytrace("a"))
	fmt.Println("in a")
}

func myb() {
	defer un(mytrace("b"))
	fmt.Println("in b")
	mya()
}

func Going112121() {
	fmt.Println("defer track go..")
	myb()
}
