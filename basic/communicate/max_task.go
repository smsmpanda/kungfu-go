package communicate

import (
	"fmt"
	"time"
)

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request1 struct {
	a, b   int
	replyc chan int
}

func process(r *Request1) {
	// do something
}

func handle(r *Request1) {
	sem <- 1 // doesn't matter what we put in it
	process(r)
	<-sem // one empty place in the buffer: the next Request1 can start
}

func server1(service chan *Request1) {
	for {
		Request1 := <-service
		go handle(Request1)
	}
}

func Gorun() {
	// service := make(chan *Request1)
	// go server1(service)

	// select {}

	semqueue := make(chan int)
	go func() {
		for v := range semqueue {
			fmt.Println(v)
		}
		fmt.Print("end...")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			semqueue <- i
			time.Sleep(time.Second * 1)
		}
	}()

	time.Sleep(time.Second * 1000)
}
