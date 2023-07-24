package communicate

import "fmt"

func Gorun4grr() {
	fmt.Print("hello\n")

	resume = intergers()
	fmt.Println(generateInteger()) //=> 0
	fmt.Println(generateInteger()) //=> 1
	fmt.Println(generateInteger()) //=> 2

}

var resume chan int

func intergers() chan int {
	yield := make(chan int)
	count := 0

	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}
