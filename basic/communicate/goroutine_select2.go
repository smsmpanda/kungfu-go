package communicate

import (
	"fmt"
	"os"
	"time"
)

func tel(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	fmt.Println("Is being done!")
	quit <- true
}

func Gorun213() {
	// var ok = true
	// ch := make(chan int)
	// quit := make(chan bool)

	// go tel(ch, quit)
	// for ok {
	// 	select {
	// 	case i := <-ch:
	// 		fmt.Printf("The counter is at %d\n", i)
	// 	case <-quit:
	// 		os.Exit(0)
	// 	}
	// }

	// fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	// run1()

	c := make(chan int)
	// consumer:
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()
	// producer:
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}

func run1() {
	term := 25
	i := 0
	c := make(chan int)
	start := time.Now()

	go fibnterms(term, c)
	for {
		if result, ok := <-c; ok {
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
			i++
		} else {
			end := time.Now()
			delta := end.Sub(start)
			fmt.Printf("longCalculation took this amount of time: %s\n", delta)
			os.Exit(0)
		}
	}
}

func fibnterms(term int, c chan int) {
	for i := 0; i <= term; i++ {
		c <- fibonacci(i)
	}
	close(c)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
