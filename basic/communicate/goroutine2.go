package communicate

import (
	"fmt"
)

func GorunGoroutine2() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	//time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
