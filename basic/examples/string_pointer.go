package main

import "fmt"

func main() {
	s := "good bye"
	var sptr *string = &s
	fmt.Printf("%p address store content is %s", sptr, *sptr)
}
