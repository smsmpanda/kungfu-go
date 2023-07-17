package communicate

import (
	"fmt"
	"time"
)

func generate(ch chan int) {
	for i := 2; i < 20; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func Gorun() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	go func() {
		for {
			prime := <-ch
			fmt.Print(prime, " ")
			ch1 := make(chan int)
			go filter(ch, ch1, prime)
			ch = ch1
		}
	}()

	time.Sleep(time.Second * 5)
}
