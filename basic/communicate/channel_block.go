package communicate

import (
	"fmt"
	"time"
)

func GorunChannelBlock() {
	// ch1 := make(chan int)
	// go suck(ch1)
	// go pump(ch1)

	// time.Sleep(2 * 1e9)
	// fmt.Println(<-ch1)

	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
