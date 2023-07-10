package osrel

import (
	"bufio"
	"fmt"
	"myexample/basic/methods"
	"os"
	"strconv"
)

func Gorunc123() {
	buf := bufio.NewReader(os.Stdin)

	calc1 := new(methods.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")

	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("input error!")
			os.Exit(0)
		}

		token = token[0 : len(token)-2] // remove \r\n

		switch {
		case token == "q":
			fmt.Println("Caculator stopped.")
			return
		case token >= "0" && token <= "999999":
			i, _ := strconv.Atoi(token)
			calc1.Push(i)

		case token == "+":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p+q)

		case token == "*":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p*q)

		case token == "/":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p/q)

		default:
			fmt.Println("No valid input")
		}
	}
}

/*
在计算机编程中，句柄是一个指向操作系统中某个对象或资源的引用。这个对象或资源可以是任何东西，比如窗口、文件、内存块、设备等。
句柄的主要作用是在程序中标识和访问这些对象或资源。因为操作系统为每个对象或资源分配了一个唯一的标识符，程序可以通过标识符来引用他们。但是这些标识符通常是比较复杂的，而且可能会发生改变。
使用句柄可以使程序更容易的管理和访问这些对象或资源，并且可以避免与标识符相关的问题。
例如，在windows系统中每个窗口都有唯一的句柄。程序可以使用这个句柄来访问窗口，比如来更改窗口大小，位置，标题等。如果没有句柄，程序可能需要使用窗口的标识符来访问它，这可能会很麻烦并且很容易出错。
总之，句柄是一个非常重要的概念，它使程序更容易管理和访问操作系统中的对象或资源。
*/
