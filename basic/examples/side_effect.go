package examples

import "fmt"

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func Gorun() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Printf("n's value is %d\n", n)
}
