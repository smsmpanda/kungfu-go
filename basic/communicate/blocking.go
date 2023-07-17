package communicate

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println("step..3")
	time.Sleep(5 * 1e9)
	fmt.Println(<-in)
}

func Gorun12() {
	// fmt.Println("start...")
	// out := make(chan int)

	// fmt.Println("step..1")
	// go f1(out)

	// fmt.Println("step..2")
	// out <- 5

	// fmt.Println("end")

	//channelBuffer()

	//gosum()
}

func myRead(c <-chan int, done chan<- bool) {
	for num := range c {
		fmt.Printf("received::-> %d", num)
	}
	done <- true
}

func myWrite(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * 10
		fmt.Printf("send::-> %d", c)
	}
	close(c)
}

func sum(x, y int, c chan int) {
	for i := 0; ; i++ {
		c <- x + y
		fmt.Printf("send::-> %d", c)
	}
}

func Gorun123fd() {
	// c := make(chan int)
	// d := make(chan bool)
	// // go sum(12, 13, c)
	// // fmt.Println(<-c) // 25
	// go myWrite(c)
	// go myRead(c, d)

	// <-d
	for {
		fmt.Println(1)
	}
}

func channelBuffer() {
	c := make(chan int, 50)
	go func() {
		time.Sleep(5 * 1e9)
		x := <-c
		fmt.Println(x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
