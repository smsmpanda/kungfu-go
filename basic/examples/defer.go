package examples

import "fmt"

func Going123() {
	function1()
	deferTest()
	manyDefer()
	b()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Println("Function2: Deferred until the end of the calling function!")
}

func deferTest() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func manyDefer() {
	// for i := 0; i < 5; i++ {
	// 	defer fmt.Printf("%d\n", i)
	// }
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(4)
}

func trace(s string) {
	fmt.Println("entering:", s)
}
func untrace(s string) {
	fmt.Println("leaving:", s)
}

func aaa() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	aaa()
}
