package communicate

import (
	"fmt"
	"time"
)

func GorunSelect() {
	ch1 := make(chan int)
	ch22 := make(chan int)

	go pump1(ch1)
	go pump2(ch22)
	go suck2(ch1, ch22)

	time.Sleep(3 * 1e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck2(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
			break
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
			break
		default:
			fmt.Println("No value received")
			break
		}
	}
}
