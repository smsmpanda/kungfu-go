package communicate

import (
	"fmt"
	"time"
)

func generate(ch chan int) {
	for i := 2; i <= 20; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		fmt.Printf("in -> %d\n", i)
		// if i%prime != 0 {
		// 	out <- i // Send 'i' to channel 'out'.
		// }
	}
}

// The prime sieve: Daisy-chain filter processes together.
func Gorunf2() {
	// ch := make(chan int) // Create a new channel.
	// go generate(ch)      // Start generate() as a goroutine.
	// for {
	// 	prime := <-ch
	// 	fmt.Printf("prime -> %d\n", prime)
	// 	ch1 := make(chan int)
	// 	go filter(ch, ch1, prime)
	// 	// ch = ch1
	// }

	// time.Sleep(time.Second * 10000)

	ch := make(chan string)

	go func() {
		ch <- "j"
		time.Sleep(time.Second * 2)
		ch <- "i"
		time.Sleep(time.Second * 1)
		ch <- "a"
		time.Sleep(time.Second * 1)
		ch <- "n"
		time.Sleep(time.Second * 1)
		ch <- "g"
		time.Sleep(time.Second * 1)
		ch <- "r"
		time.Sleep(time.Second * 2)
		ch <- "u"
		time.Sleep(time.Second * 1)
		ch <- "i"
		time.Sleep(time.Second * 1)
		ch <- "h"
		time.Sleep(time.Second * 2)
		ch <- "w"
		time.Sleep(time.Second * 1)
	}()

	go func() {
		for {
			s := <-ch
			fmt.Printf("go1 -> received::->%s\n", s)
			go2(ch, s)
		}
	}()

	for {
		s := <-ch
		fmt.Printf("go3 -> received::->%s\n", s)
	}
}

func go2(c chan string, str string) {
	fmt.Println(str)
	go func() {
		for {
			s := <-c
			fmt.Printf("go2 -> received::->%s\n", s)
		}
	}()
}
